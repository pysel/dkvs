package balancer_test

import (
	"math/big"
	"net"
	"os"
	"testing"

	"github.com/pysel/dkvs/balancer"
	"github.com/pysel/dkvs/partition"
	"github.com/pysel/dkvs/testutil"
	"github.com/stretchr/testify/require"
)

func SetupPartition(dbPath string) net.Addr {
	lis, s := testutil.PartitionServer(dbPath)
	testutil.RunPartitionServer(lis, s)
	return lis.Addr()
}

func TestRegisterGetPartition(t *testing.T) {
	defer os.RemoveAll(testutil.TestDBPath)

	addr := SetupPartition(testutil.TestDBPath)
	b10 := balancer.NewBalancerTest(10)

	err := b10.RegisterPartition(addr.String(), *testutil.DefaultHashrange)
	require.NoError(t, err)

	domainKey := "Partition key"
	nonDomainKey := "Not partition key."

	keyPartitions := b10.GetPartitions([]byte(domainKey))
	require.Equal(t, 1, len(keyPartitions))

	keyPartitions = b10.GetPartitions([]byte(nonDomainKey))
	require.Equal(t, 0, len(keyPartitions))

	b0 := balancer.NewBalancerTest(0)
	err = b0.RegisterPartition(addr.String(), *testutil.DefaultHashrange)
	require.Error(t, err)
}

func TestBalancerInit(t *testing.T) {
	goalReplicaRanges := 3

	b := balancer.NewBalancerTest(goalReplicaRanges)
	require.Equal(t, b.GetTicksAmount(), goalReplicaRanges+1)

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
	b := balancer.NewBalancerTest(1)
	nextPartitionRange, err := b.GetNextPartitionRange()
	require.NoError(t, err)
	// defaultHashrange is full sha256 domain, in case of 1 node, it's domain should be full
	require.Equal(t, nextPartitionRange, testutil.FullHashrange)

	b2 := balancer.NewBalancerTest(2)
	nextPartitionRange, err = b2.GetNextPartitionRange()
	require.NoError(t, err)
	// defaultHashrange is full sha256 domain, in case of 2 nodes, first node's domain should be half
	require.Equal(t, nextPartitionRange, partition.NewRange(big.NewInt(0), testutil.HalfShaDomain))

	b2.SetActivePartitions(1)
	nextPartitionRange, err = b2.GetNextPartitionRange()
	require.NoError(t, err)
	// defaultHashrange is full sha256 domain, in case of 2 nodes, second node's domain should be the second half
	require.Equal(t, nextPartitionRange, partition.NewRange(testutil.HalfShaDomain, testutil.FullHashrange.Max))

	b2.SetActivePartitions(5)
	nextPartitionRange, err = b2.GetNextPartitionRange()
	require.NoError(t, err)
	// if partitions are already covering the whole domain, next partition should be a replica of the second one
	// (since 5 mod 2 is 1)
	require.Equal(t, nextPartitionRange, partition.NewRange(testutil.HalfShaDomain, testutil.FullHashrange.Max))
}
