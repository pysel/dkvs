package client_test

import (
	"sync"
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
	balancerAddress, closer := testutil.BalancerClientWith2Partitions(t)

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

func TestClientParallel(t *testing.T) {
	// setup balancer server to which the client will be connected
	balancerAddress, closer := testutil.BalancerClientWith2Partitions(t)

	defer closer()

	// generate load
	var wg sync.WaitGroup
	load := generateLoad(2)
	channel := make(chan grpcError, len(load))

	wg.Add(len(load))
	for _, f := range load {
		c := client.NewClient(balancerAddress.String())
		go f(c, channel, &wg)
	}

	require.Zero(t, len(channel))
}

type grpcError struct {
	errSet    error
	errDelete error
	errGet    error
}

func (ge *grpcError) ok() bool {
	return ge.errSet == nil && ge.errDelete == nil && ge.errGet == nil
}

func generateLoad(goroutineNumber int) []func(*client.Client, chan grpcError, *sync.WaitGroup) {
	var load []func(*client.Client, chan grpcError, *sync.WaitGroup)
	for i := 0; i < goroutineNumber; i++ {
		load = append(load, func(c *client.Client, channel chan grpcError, wg *sync.WaitGroup) {
			defer wg.Done()

			errSet := c.Set([]byte("key"+string(rune(i))), []byte("value"+string(rune(i))))
			errDelete := c.Delete([]byte("key" + string(rune(i))))
			_, errGet := c.Get([]byte("key" + string(rune(i))))

			grpcError := grpcError{
				errSet:    errSet,
				errDelete: errDelete,
				errGet:    errGet,
			}

			if grpcError.ok() {
				return
			}

			channel <- grpcError
		})
	}

	return load
}
