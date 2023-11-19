package testutil

import (
	"context"
	"log"
	"math/big"
	"net"

	"github.com/pysel/dkvs/partition"
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const (
	bufSize    = 1024 * 1024
	TestDBPath = "test"
)

var (
	Min10            *big.Int // zero
	HalfShaDomain    *big.Int // half of sha-2 domain
	DefaultHashrange *partition.Range
	FullHashrange    *partition.Range
)

func init() {
	Min10 = new(big.Int).SetInt64(10)
	to_bz := make([]byte, 32)
	for i := 0; i < 32; i++ {
		to_bz[i] = 0xFF
	}
	full_range := new(big.Int).SetBytes(to_bz) // 2^256 - 1

	HalfShaDomain = new(big.Int).Div(full_range, big.NewInt(2)) // half of 2^256 - 1
	DefaultHashrange = &partition.Range{
		Min: Min10,
		Max: HalfShaDomain,
	}

	FullHashrange = &partition.Range{
		Min: big.NewInt(0),
		Max: full_range,
	}
}

// PartitionServer creates a listener and a server for the partition service.
func PartitionServer(dbPath string) (*bufconn.Listener, *grpc.Server) {
	lis := bufconn.Listen(bufSize)
	s := grpc.NewServer()
	p := partition.NewPartition(dbPath)

	pbpartition.RegisterPartitionServiceServer(s, &partition.ListenServer{Partition: p})

	return lis, s
}

// RunPartitionServer runs a goroutine that constantly serves requests on the given listener.
func RunPartitionServer(lis *bufconn.Listener, s *grpc.Server) {
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

// partitionClient set ups a partition server and partition client. Returns client and closer function.
// Client is used to test the rpc calls.
func SinglePartitionClient(ctx context.Context) (pbpartition.PartitionServiceClient, func()) {
	lis, s := PartitionServer(TestDBPath)
	RunPartitionServer(lis, s)

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
	return pbpartition.NewPartitionServiceClient(conn), closer
}
