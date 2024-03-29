package testutil

import (
	"context"
	"log"
	"math/big"
	"net"
	"os"
	"strconv"
	"testing"

	"github.com/pysel/dkvs/balancer"

	"github.com/pysel/dkvs/partition"
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
	"github.com/pysel/dkvs/shared"
	hashrange "github.com/pysel/dkvs/types/hashrange"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

const (
	bufSize    = 1024 * 1024
	TestDBPath = "test"
)

var (
	Min              *big.Int // zero
	HalfShaDomain    *big.Int // half of sha-2 domain
	DefaultHashrange *hashrange.Range
	FullHashrange    *hashrange.Range
)

func init() {
	Min = new(big.Int).SetInt64(0)
	to_bz := make([]byte, 32)
	for i := 0; i < 32; i++ {
		to_bz[i] = 0xFF
	}
	full_range := new(big.Int).SetBytes(to_bz) // 2^256 - 1

	HalfShaDomain = new(big.Int).Div(full_range, big.NewInt(2)) // half of 2^256 - 1
	DefaultHashrange = &hashrange.Range{
		Min: Min,
		Max: HalfShaDomain,
	}

	FullHashrange = &hashrange.Range{
		Min: big.NewInt(0),
		Max: full_range,
	}
}

// BufferedPartitionServer creates a listener and a server for the partition service.
func BufferedPartitionServer(dbPath string) (*bufconn.Listener, *grpc.Server) {
	lis := bufconn.Listen(bufSize)
	s := grpc.NewServer()
	p := partition.NewPartition(dbPath)

	pbpartition.RegisterPartitionServiceServer(s, &partition.ListenServer{Partition: p})

	return lis, s
}

// RunBufferedPartitionServer runs a goroutine that constantly serves requests on the given listener.
func RunBufferedPartitionServer(lis *bufconn.Listener, s *grpc.Server) {
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Server exited with error: %v", err)
		}
	}()
}

// StartPartitionClientToBufferedServer sets up a partition server and partition client. Returns client and closer function.
// Client is used to test the rpc calls.
func StartPartitionClientToBufferedServer(ctx context.Context) (net.Addr, pbpartition.PartitionServiceClient, func()) {
	lis, s := BufferedPartitionServer(TestDBPath)
	RunBufferedPartitionServer(lis, s)

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
	return lis.Addr(), pbpartition.NewPartitionServiceClient(conn), closer
}

func StartXPartitionServers(t *testing.T, x int) ([]net.Addr, []string) {
	addrs := make([]net.Addr, x)
	dbPaths := make([]string, x)
	for i := 0; i < x; i++ {
		path := TestDBPath + strconv.Itoa(i) + "test" + t.Name()
		p := partition.NewPartition(path)
		s := partition.RegisterPartitionServer(p)
		_, addr := shared.StartListeningOnPort(s, 0)
		addrs[i] = addr
		dbPaths[i] = path
	}

	return addrs, dbPaths
}

func BalancerClientWith2Partitions(t *testing.T) (net.Addr, func()) {
	ctx := context.Background()
	addrs, dbPaths := StartXPartitionServers(t, 2)

	// register partitions
	b := balancer.NewBalancer(balancer.BalancerDBPath+t.Name(), 2)
	_ = b.RegisterPartition(ctx, addrs[0].String())
	_ = b.RegisterPartition(ctx, addrs[1].String())

	server := balancer.RegisterBalancerServer(b)
	_, addr := shared.StartListeningOnPort(server, 0)

	return addr, func() {
		// remove all databases - one for balancer and one for each partitin
		os.RemoveAll(balancer.BalancerDBPath + t.Name())
		for _, path := range dbPaths {
			err := os.RemoveAll(path)
			if err != nil {
				log.Printf("error removing path: %v", err)
			}
		}
	}
}
