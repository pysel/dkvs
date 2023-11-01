package e2e

import (
	"context"
	"crypto/sha256"
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
	from *big.Int
	to   *big.Int
)

func init() {
	from = new(big.Int).SetInt64(0)
	to_bz := make([]byte, 32)
	for i := 0; i < 32; i++ {
		to_bz[i] = 0xFF
	}

	to = new(big.Int).SetBytes(to_bz)
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
	fmt.Println(sha256.Sum256(domainKey))

	_, err := client.StoreMessage(ctx, &prototypes.StoreMessageRequest{
		Key:   domainKey,
		Value: []byte("value"),
	})
	require.NoError(t, err, "StoreMessage should not return error")
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
