package balancer_test

import (
	"context"
	"os"
	"testing"

	"github.com/pysel/dkvs/balancer"
	"github.com/pysel/dkvs/partition"
	"github.com/pysel/dkvs/prototypes"
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
	"github.com/pysel/dkvs/testutil"
	"github.com/pysel/dkvs/types"
	"github.com/stretchr/testify/require"
)

var (
	TestDBPath2    = "test2"
	TestDBBalancer = "balancer"
)

func TestTwoPhaseCommit(t *testing.T) {
	defer os.RemoveAll(testutil.TestDBPath)
	defer os.RemoveAll(TestDBPath2)
	defer os.RemoveAll(TestDBBalancer + t.Name())

	ctx := context.Background()

	partitionAddr1 := testutil.RunPartitionServer(0, testutil.TestDBPath)
	partitionAddr2 := testutil.RunPartitionServer(0, TestDBPath2)

	b := balancer.NewBalancerTest(t, 2)
	err := b.RegisterPartition(ctx, partitionAddr1.String())
	require.NoError(t, err)

	err = b.RegisterPartition(ctx, partitionAddr2.String())
	require.NoError(t, err)

	domainKey := []byte("Partition key")
	shaKey := types.ShaKey(domainKey)
	range_, err := b.GetRangeFromDigest(shaKey[:])
	require.NoError(t, err)

	msgSet := &pbpartition.PrepareCommitRequest{
		Message: &pbpartition.PrepareCommitRequest_Set{
			Set: &prototypes.SetRequest{
				Key:     domainKey,
				Value:   []byte("value"),
				Lamport: 0,
			},
		},
	}

	err = b.AtomicMessage(ctx, range_, msgSet)
	require.NoError(t, err)

	// Assert that value was stored correctly
	getResp, err := b.Get(ctx, domainKey)
	require.NoError(t, err)

	expected := partition.ToStoredValue(0, []byte("value"))
	require.Equal(t, expected, getResp.StoredValue)

	msgDelete := &pbpartition.PrepareCommitRequest{
		Message: &pbpartition.PrepareCommitRequest_Delete{
			Delete: &prototypes.DeleteRequest{
				Key: domainKey,
			},
		},
	}

	err = b.AtomicMessage(ctx, range_, msgDelete)
	require.NoError(t, err)

	// Assert that value was deleted correctly
	getResp, err = b.Get(ctx, domainKey)
	require.NoError(t, err)

	require.Nil(t, getResp.StoredValue)
}
