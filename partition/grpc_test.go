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

	client.SetHashrange(ctx, &prototypes.SetHashrangeRequest{
		Min: testutil.DefaultHashrange.Min.Bytes(),
		Max: testutil.DefaultHashrange.Max.Bytes(),
	})

	defer closer()
	defer require.NoError(t, os.RemoveAll(testutil.TestDBPath))
	domainKey := "Partition key" // a hash of this text lays in [from; to]
	nonDomainKey := "Not partition key."

	// Assert that value was stored correctly
	_, err := client.SetMessage(ctx, &prototypes.SetMessageRequest{
		Key:   domainKey,
		Value: []byte("value"),
	})
	require.NoError(t, err, "SetMessage should not return error")

	// Assert that value was stored correctly
	getResp, err := client.GetMessage(ctx, &prototypes.GetMessageRequest{Key: domainKey})
	require.NoError(t, err, "GetMessage should not return error")
	require.Equal(t, []byte("value"), getResp.Value, "GetMessage should return correct value")

	// Assert that value was not stored if key is nil
	setResp, err := client.SetMessage(ctx, &prototypes.SetMessageRequest{})
	require.ErrorContains(t, err, partition.ErrNilKey.Error(), "SetMessage should return error if key is nil")
	require.Nil(t, setResp, "SetMessage should return nil response if key is nil")

	// Assert that get operation won't succeed if key is nil
	getResp, err = client.GetMessage(ctx, &prototypes.GetMessageRequest{})
	require.ErrorContains(t, err, partition.ErrNilKey.Error(), "GetMessage should return error if key is nil")
	require.Nil(t, getResp, "GetMessage should return nil response if key is nil")

	// Assert that value was not stored if key is not in partition's hashrange
	setResp, err = client.SetMessage(ctx, &prototypes.SetMessageRequest{Key: nonDomainKey, Value: []byte("value")})
	require.ErrorContains(t, err, partition.ErrNotThisPartitionKey.Error(), "SetMessage should return error if key is not domain key")
	require.Nil(t, setResp, "SetMessage should return nil response if key is not domain key")

	// Assert that get operation won't succeed if key is not in partition's hashrange
	getResp, err = client.GetMessage(ctx, &prototypes.GetMessageRequest{Key: nonDomainKey})
	require.ErrorContains(t, err, partition.ErrNotThisPartitionKey.Error(), "GetMessage should return error if key is not domain key")
	require.Nil(t, getResp, "GetMessage should return nil response if key is not domain key")
}
