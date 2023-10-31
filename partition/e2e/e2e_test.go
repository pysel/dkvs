package e2e

import (
	"context"
	"fmt"
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
	from = new(big.Int).SetInt64(0)
	to   = new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil) // half of SHA-2 domain
)

func server(ctx context.Context) (pbpartition.CommandsServiceClient, func()) {
	lis := bufconn.Listen(bufSize)
	s := grpc.NewServer()
	fmt.Println("TO: ", to)
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

func TestGRPCServer(t *testing.T) {
	ctx := context.Background()

	client, closer := server(ctx)
	if closer == nil {
		t.Fatal("Closer should not be nil")
	}

	defer closer()
	defer os.RemoveAll(testDBPath)
	domainKey := []byte("Partition key")

	res, err := client.StoreMessage(ctx, &prototypes.StoreMessageRequest{
		Key:   domainKey,
		Value: []byte("value"),
	})
	require.NoError(t, err, "StoreMessage should not return error")
	require.Equal(t, res, &prototypes.StoreMessageResponse{}, "StoreMessage should return empty response")

}
