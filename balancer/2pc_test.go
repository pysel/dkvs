package balancer_test

import (
	"context"
	"os"
	"testing"

	"github.com/pysel/dkvs/balancer"
	"github.com/pysel/dkvs/prototypes"
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
	"github.com/pysel/dkvs/testutil"
	"github.com/pysel/dkvs/types"
	"github.com/stretchr/testify/require"
)

var (
	TestDBPath2 = "test2"
)

func TestTwoPhaseCommit(t *testing.T) {
	defer os.RemoveAll(testutil.TestDBPath)
	defer os.RemoveAll(TestDBPath2)

	ctx := context.Background()

	partitionAddr1 := testutil.RunPartitionServer(0, testutil.TestDBPath)
	partitionAddr2 := testutil.RunPartitionServer(0, TestDBPath2)

	b := balancer.NewBalancerTest(2)
	err := b.RegisterPartition(ctx, partitionAddr1.String())
	require.NoError(t, err)

	err = b.RegisterPartition(ctx, partitionAddr2.String())
	require.NoError(t, err)

	domainKey := "Partition key"
	shaKey := types.ShaKey(domainKey)
	range_, err := b.GetRangeFromDigest(shaKey[:])
	require.NoError(t, err)

	msg := &pbpartition.PrepareCommitRequest{
		Message: &pbpartition.PrepareCommitRequest_Set{
			Set: &prototypes.SetRequest{
				Key:     domainKey,
				Value:   []byte("value"),
				Lamport: 0,
			},
		},
	}

	err = b.AtomicMessage(ctx, range_, msg)
	require.NoError(t, err)

	// Assert that value was stored correctly
	getResp, err := b.Get(ctx, domainKey)
	require.NoError(t, err)
	require.Equal(t, []byte("value"), getResp.Value)
}
