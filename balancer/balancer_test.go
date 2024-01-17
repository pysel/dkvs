package balancer_test

import (
	"context"
	"math/big"
	"os"
	"testing"

	"github.com/pysel/dkvs/balancer"
	"github.com/pysel/dkvs/partition"
	"github.com/pysel/dkvs/testutil"
	"github.com/stretchr/testify/require"
)

func TestRegisterGetPartition(t *testing.T) {
	defer os.RemoveAll(TestDBBalancer + t.Name())
	addrs, paths := testutil.StartXPartitionServers(1)
	defer testutil.RemovePaths(paths)

	ctx := context.Background()

	addr := addrs[0]
	b2 := balancer.NewBalancerTest(t, 2)

	err := b2.RegisterPartition(ctx, addr.String())
	require.NoError(t, err)

	keyPartitions := b2.GetPartitionsByKey(testutil.DomainKey)
	require.Equal(t, 1, len(keyPartitions))

	keyPartitions = b2.GetPartitionsByKey(testutil.NonDomainKey)
	require.Equal(t, 0, len(keyPartitions))
}

func TestBalancerInit(t *testing.T) {
	defer os.RemoveAll(TestDBBalancer + t.Name())

	goalReplicaRanges := 3

	b := balancer.NewBalancerTest(t, goalReplicaRanges)
	require.Equal(t, b.GetCoverageSize(), goalReplicaRanges+1)

	expectedFirstTickValue := big.NewInt(0)
	require.NotNil(t, b.GetTickByValue(expectedFirstTickValue))

	expectedSecondTickValue := new(big.Int).Div(partition.MaxInt, big.NewInt(3))
	require.NotNil(t, b.GetTickByValue(expectedSecondTickValue))

	expectedThirdTickNumerator := new(big.Int).Mul(partition.MaxInt, big.NewInt(2))
	expectedThirdTickValue := new(big.Int).Div(expectedThirdTickNumerator, big.NewInt(3))
	require.NotNil(t, b.GetTickByValue(expectedThirdTickValue))

	expectedFourthTickValue := partition.MaxInt
	require.NotNil(t, b.GetTickByValue(expectedFourthTickValue))
}

func TestGetNextPartitionRange(t *testing.T) {
	defer os.RemoveAll(TestDBBalancer + t.Name())
	addrs, paths := testutil.StartXPartitionServers(2)
	defer testutil.RemovePaths(paths)

	addr1, addr2 := addrs[0], addrs[1]

	ctx := context.Background()

	// SUT
	b2 := balancer.NewBalancerTest(t, 2)
	nextPartitionRange, _, _ := b2.GetNextPartitionRange()
	// defaultHashrange is full sha256 domain, in case of 2 nodes, first node's domain should be half
	require.Equal(t, partition.NewRange(big.NewInt(0).Bytes(), testutil.HalfShaDomain.Bytes()), nextPartitionRange)

	// Register first Partition
	require.NoError(t, b2.RegisterPartition(ctx, addr1.String()))

	nextPartitionRange, _, _ = b2.GetNextPartitionRange()
	// defaultHashrange is full sha256 domain, in case of 2 nodes, second node's domain should be the second half
	require.Equal(t, partition.NewRange(testutil.HalfShaDomain.Bytes(), testutil.FullHashrange.Max.Bytes()), nextPartitionRange)

	// Register second Partition
	require.NoError(t, b2.RegisterPartition(ctx, addr2.String()))

	// If all ranges are covered, newer partitions should start coverting the domain from the beginning
	nextPartitionRange, _, _ = b2.GetNextPartitionRange()
	require.Equal(t, nextPartitionRange, partition.NewRange(big.NewInt(0).Bytes(), testutil.HalfShaDomain.Bytes()))

	// Assert that GetNextPartitionRange is non-mutative
	nextPartitionRange, _, _ = b2.GetNextPartitionRange()
	require.Equal(t, nextPartitionRange, partition.NewRange(big.NewInt(0).Bytes(), testutil.HalfShaDomain.Bytes()))
}
