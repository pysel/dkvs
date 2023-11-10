package balancer_test

import (
	"net"
	"os"
	"testing"

	"github.com/pysel/dkvs/balancer"
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
}
