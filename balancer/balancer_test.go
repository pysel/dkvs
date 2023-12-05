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
	defer os.RemoveAll(testutil.TestDBPath)
	defer os.RemoveAll(TestDBBalancer + t.Name())

	ctx := context.Background()

	addr := testutil.RunPartitionServer(0, testutil.TestDBPath)
	b2 := balancer.NewBalancerTest(t, 2)

	err := b2.RegisterPartition(ctx, addr.String())
	require.NoError(t, err)

	domainKey := "Partition key"
	nonDomainKey := "Not partition key."

	keyPartitions := b2.GetPartitionsByKey([]byte(domainKey))
	require.Equal(t, 1, len(keyPartitions))

	keyPartitions = b2.GetPartitionsByKey([]byte(nonDomainKey))
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
	defer os.RemoveAll(testutil.TestDBPath)
	defer os.RemoveAll(TestDBPath2)
	defer os.RemoveAll(TestDBBalancer + t.Name())

	addr1 := testutil.RunPartitionServer(0, testutil.TestDBPath)
	addr2 := testutil.RunPartitionServer(0, TestDBPath2)

	ctx := context.Background()

	// SUT
	b2 := balancer.NewBalancerTest(t, 2)
	nextPartitionRange, _, _ := b2.GetNextPartitionRange()
	// defaultHashrange is full sha256 domain, in case of 2 nodes, first node's domain should be half
	require.Equal(t, partition.NewRange(big.NewInt(0).Bytes(), testutil.HalfShaDomain.Bytes()), nextPartitionRange)

	// Register first Partition
	b2.RegisterPartition(ctx, addr1.String())

	nextPartitionRange, _, _ = b2.GetNextPartitionRange()
	// defaultHashrange is full sha256 domain, in case of 2 nodes, second node's domain should be the second half
	require.Equal(t, partition.NewRange(testutil.HalfShaDomain.Bytes(), testutil.FullHashrange.Max.Bytes()), nextPartitionRange)

	// Register second Partition
	b2.RegisterPartition(ctx, addr2.String())

	// If all ranges are covered, newer partitions should start coverting the domain from the beginning
	nextPartitionRange, _, _ = b2.GetNextPartitionRange()
	require.Equal(t, nextPartitionRange, partition.NewRange(big.NewInt(0).Bytes(), testutil.HalfShaDomain.Bytes()))

	// Assert that GetNextPartitionRange is non-mutative
	nextPartitionRange, _, _ = b2.GetNextPartitionRange()
	require.Equal(t, nextPartitionRange, partition.NewRange(big.NewInt(0).Bytes(), testutil.HalfShaDomain.Bytes()))
}
