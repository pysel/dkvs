package partition_test

import (
	"context"
	"os"
	"testing"

	"github.com/pysel/dkvs/partition"
	"github.com/pysel/dkvs/prototypes"
	"github.com/pysel/dkvs/testutil"
	"github.com/stretchr/testify/require"
)

func TestGRPCServer(t *testing.T) {
	ctx := context.Background()
	client, closer := testutil.SinglePartitionClient(ctx)
	if closer == nil {
		t.Fatal("Closer should not be nil")
	}

	_, err := client.SetHashrange(ctx, &prototypes.SetHashrangeRequest{
		Min: testutil.DefaultHashrange.Min.Bytes(),
		Max: testutil.DefaultHashrange.Max.Bytes(),
	})

	require.NoError(t, err)

	defer closer()
	defer require.NoError(t, os.RemoveAll(testutil.TestDBPath))
	domainKey := "Partition key" // a hash of this text lays in [from; to]
	nonDomainKey := "Not partition key."

	// Assert that value was stored correctly
	_, err = client.Set(ctx, &prototypes.SetRequest{
		Key:     domainKey,
		Value:   []byte("value"),
		Lamport: 1,
	})
	require.NoError(t, err, "SetMessage should not return error")

	// Assert that value was stored correctly
	getResp, err := client.Get(ctx, &prototypes.GetRequest{Key: domainKey})
	require.NoError(t, err, "GetMessage should not return error")

	expected := partition.ToStoredValue(1, []byte("value"))
	require.Equal(t,
		expected,
		getResp.StoredValue,
		"GetMessage should return correct value",
	)

	// Assert that value was not stored if key is nil
	setResp, err := client.Set(ctx, &prototypes.SetRequest{})
	require.Error(t, err, "SetMessage should return error if key is nil")
	require.Nil(t, setResp, "SetMessage should return nil response if key is nil")

	// Assert that get operation won't succeed if key is nil
	getResp, err = client.Get(ctx, &prototypes.GetRequest{})
	require.Error(t, err, "GetMessage should return error if key is nil")
	require.Nil(t, getResp, "GetMessage should return nil response if key is nil")

	// Assert that value was not stored if key is not in partition's hashrange
	setResp, err = client.Set(ctx, &prototypes.SetRequest{Key: nonDomainKey, Value: []byte("value")})
	require.ErrorContains(t, err, partition.ErrNotThisPartitionKey.Error(), "SetMessage should return error if key is not domain key")
	require.Nil(t, setResp, "SetMessage should return nil response if key is not domain key")

	// Assert that get operation won't succeed if key is not in partition's hashrange
	getResp, err = client.Get(ctx, &prototypes.GetRequest{Key: nonDomainKey})
	require.ErrorContains(t, err, partition.ErrNotThisPartitionKey.Error(), "GetMessage should return error if key is not domain key")
	require.Nil(t, getResp, "GetMessage should return nil response if key is not domain key")

	// Assert that value was deleted correctly
	_, err = client.Delete(ctx, &prototypes.DeleteRequest{Key: domainKey})
	require.NoError(t, err, "DeleteMessage should not return error")

	// Assert that value was not deleted if key is nil
	_, err = client.Delete(ctx, &prototypes.DeleteRequest{})
	require.Error(t, err, "DeleteMessage should return error if key is nil")

	// Assert that value was not deleted if key is not in partition's hashrange
	_, err = client.Delete(ctx, &prototypes.DeleteRequest{Key: nonDomainKey})
	require.ErrorContains(t, err, partition.ErrNotThisPartitionKey.Error(), "DeleteMessage should return error if key is not domain key")

	// Assert that deleted value was removed from partition's state
	getResp, err = client.Get(ctx, &prototypes.GetRequest{Key: domainKey})
	require.NoError(t, err)
	require.Nil(t, getResp.StoredValue)
}
