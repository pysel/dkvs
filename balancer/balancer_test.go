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

func TestAddGetPartition(t *testing.T) {
	defer os.RemoveAll(testutil.TestDBPath)

	addr := SetupPartition()
	b := balancer.NewBalancer()

	b.AddPartition(addr.String(), *testutil.DefaultHashrange)

	domainKey := "Partition key"
	nonDomainKey := "Not partition key."

	keyPartitions := b.GetPartitions([]byte(domainKey))
	require.Equal(t, 1, len(keyPartitions))

	keyPartitions = b.GetPartitions([]byte(nonDomainKey))
	require.Equal(t, 0, len(keyPartitions))
}
