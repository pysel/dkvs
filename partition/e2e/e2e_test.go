package e2e

import (
	"context"
	"log"
	"math/big"
	"net"
	"os"
	"testing"

	"github.com/pysel/dkvs/partition"
	"github.com/pysel/dkvs/prototypes"
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const (
	bufSize    = 1024 * 1024
	testDBPath = "test"
)

var (
	from *big.Int
	to   *big.Int
)

func init() {
	from = new(big.Int).SetInt64(0)
	to_bz := make([]byte, 32)
	for i := 0; i < 32; i++ {
		to_bz[i] = 0xFF
	}
	full_range := new(big.Int).SetBytes(to_bz) // 2^256 - 1

	to = new(big.Int).Div(full_range, big.NewInt(2)) // half of 2^256 - 1
}

func TestGRPCServer(t *testing.T) {
	ctx := context.Background()

	client, closer := server(ctx)
	if closer == nil {
		t.Fatal("Closer should not be nil")
	}

	defer closer()
	defer os.RemoveAll(testDBPath)
	domainKey := []byte("Partition key")
	nonDomainKey := []byte("Not partition key.")

	// Assert that value was stored correctly
	_, err := client.StoreMessage(ctx, &prototypes.StoreMessageRequest{
		Key:   domainKey,
		Value: []byte("value"),
	})
	require.NoError(t, err, "StoreMessage should not return error")

	// Assert that value was stored correctly
	getResp, err := client.GetMessage(ctx, &prototypes.GetMessageRequest{Key: domainKey})
	require.NoError(t, err, "GetMessage should not return error")
	require.Equal(t, []byte("value"), getResp.Value, "GetMessage should return correct value")

	// Assert that value was not stored if key is nil
	setResp, err := client.StoreMessage(ctx, &prototypes.StoreMessageRequest{})
	require.ErrorContains(t, err, partition.ErrNilKey.Error(), "StoreMessage should return error if key is nil")
	require.Nil(t, setResp, "StoreMessage should return nil response if key is nil")

	// Assert that get operation won't succeed if key is nil
	getResp, err = client.GetMessage(ctx, &prototypes.GetMessageRequest{})
	require.ErrorContains(t, err, partition.ErrNilKey.Error(), "GetMessage should return error if key is nil")
	require.Nil(t, getResp, "GetMessage should return nil response if key is nil")

	// Assert that value was not stored if key is not in partition's hashrange
	setResp, err = client.StoreMessage(ctx, &prototypes.StoreMessageRequest{Key: nonDomainKey})
	require.ErrorContains(t, err, partition.ErrNotThisPartitionKey.Error(), "StoreMessage should return error if key is not domain key")
	require.Nil(t, setResp, "StoreMessage should return nil response if key is not domain key")

	// Assert that get operation won't succeed if key is not in partition's hashrange
	getResp, err = client.GetMessage(ctx, &prototypes.GetMessageRequest{Key: nonDomainKey})
	require.ErrorContains(t, err, partition.ErrNotThisPartitionKey.Error(), "GetMessage should return error if key is not domain key")
	require.Nil(t, getResp, "GetMessage should return nil response if key is not domain key")
}

func server(ctx context.Context) (pbpartition.CommandsServiceClient, func()) {
	lis := bufconn.Listen(bufSize)
	s := grpc.NewServer()
	p := partition.NewPartition(testDBPath, partition.NewRange(from, to))

	pbpartition.RegisterCommandsServiceServer(s, &partition.ListenServer{Partition: p})
	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		s.Stop()
	}
	return pbpartition.NewCommandsServiceClient(conn), closer
}
