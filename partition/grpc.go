package partition

import (
	"context"
	"crypto/sha256"
	"fmt"
	"math/big"
	"net"

	"github.com/pysel/dkvs/prototypes"
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
	"google.golang.org/grpc"
)

type ListenServer struct {
	pbpartition.UnimplementedCommandsServiceServer
	*Partition
}

func RunPartitionServer(port int64, dbPath string, from *big.Int, to *big.Int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	partition := NewPartition(dbPath, NewRange(from, to))

	grpcServer := grpc.NewServer()
	pbpartition.RegisterCommandsServiceServer(grpcServer, &ListenServer{Partition: partition})
	grpcServer.Serve(lis)
}

func (ls *ListenServer) StoreMessage(ctx context.Context, req *prototypes.StoreMessageRequest) (*prototypes.StoreMessageResponse, error) {
	shaKey := sha256.Sum256(req.Key)

	err := ls.Set(shaKey[:], req.Value)
	if err != nil {
		return nil, err
	}

	return &prototypes.StoreMessageResponse{}, nil
}
