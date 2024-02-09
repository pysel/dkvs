package client_test

import (
	"testing"

	"github.com/pysel/dkvs/client"
	"github.com/pysel/dkvs/testutil"
	"github.com/stretchr/testify/require"
)

var (
	// zero = uint64(0)
	one = uint64(1)
)

func TestClient(t *testing.T) {
	// setup balancer server to which the client will be connected
	balancerAddress, closer := testutil.BalancerClientWith2Partitions()

	defer closer()

	// setup client
	c := client.NewClient(balancerAddress.String())

	// check fields that should be non-zero
	require.Zero(t, c.GetTimestamp())        // timestamp is initially set to 0 to indicate that no messages were yet processed
	require.Equal(t, one, c.GetId())         // id is never 0
	require.NotNil(t, c.GetBalancerClient()) // address is never nil
	require.NotNil(t, c.GetContext())        // context is never nil

	// valid client requests
	err := c.Set([]byte("key"), []byte("value"))
	require.NoError(t, err)

	err = c.Delete([]byte("key"))
	require.NoError(t, err)

	value, err := c.Get([]byte("key"))
	require.NoError(t, err)
	require.Nil(t, value)
}
