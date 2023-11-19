package balancer_test

import (
	"context"
	"os"
	"testing"

	"github.com/pysel/dkvs/balancer"
	"github.com/pysel/dkvs/testutil"
	"github.com/stretchr/testify/require"
)

var TestDBPath2 = "test2"

func TestTwoPhaseCommit(t *testing.T) {
	defer os.RemoveAll(testutil.TestDBPath)
	defer os.RemoveAll(TestDBPath2)

	ctx := context.Background()

	partitionAddr1 := SetupPartition(testutil.TestDBPath)
	partitionAddr2 := SetupPartition(TestDBPath2)

	balancer := balancer.NewBalancerTest(2)
	err := balancer.RegisterPartition(partitionAddr1.String(), testutil.DefaultHashrange)
	require.NoError(t, err)

	err = balancer.RegisterPartition(partitionAddr2.String(), testutil.DefaultHashrange)
	require.NoError(t, err)

	domainKey := "Partition key"

	err = balancer.SetAtomic(ctx, domainKey, []byte("value"))
	require.NoError(t, err)
}
