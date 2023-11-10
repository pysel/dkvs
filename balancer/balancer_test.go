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

func SetupPartition() net.Addr {
	lis, s := testutil.PartitionServer()
	testutil.RunPartitionServer(lis, s)
	return lis.Addr()
}

func TestRegisterGetPartition(t *testing.T) {
	defer os.RemoveAll(testutil.TestDBPath)

	addr := SetupPartition()
	b10 := balancer.NewBalancer(10)

	err := b10.RegisterPartition(addr.String(), *testutil.DefaultHashrange)
	require.NoError(t, err)

	domainKey := "Partition key"
	nonDomainKey := "Not partition key."

	keyPartitions := b10.GetPartitions([]byte(domainKey))
	require.Equal(t, 1, len(keyPartitions))

	keyPartitions = b10.GetPartitions([]byte(nonDomainKey))
	require.Equal(t, 0, len(keyPartitions))

	b0 := balancer.NewBalancer(0)
	err = b0.RegisterPartition(addr.String(), *testutil.DefaultHashrange)
	require.Error(t, err)
}

func TestBalancerInit(t *testing.T) {
	goalPartitions := 3

	b := balancer.NewBalancer(goalPartitions)
	require.Equal(t, b.GetTicksAmount(), goalPartitions+1)

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
