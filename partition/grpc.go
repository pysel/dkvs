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
	pbpartition.UnimplementedPartitionServiceServer
	*Partition
}

func RunPartitionServer(port int64, dbPath string, from *big.Int, to *big.Int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	partition := NewPartition(dbPath, NewRange(from, to))

	grpcServer := grpc.NewServer()
	pbpartition.RegisterPartitionServiceServer(grpcServer, &ListenServer{Partition: partition})
	grpcServer.Serve(lis)
}

func (ls *ListenServer) StoreMessage(ctx context.Context, req *prototypes.StoreMessageRequest) (*prototypes.StoreMessageResponse, error) {
	if req == nil {
		return nil, ErrNilRequest
	}

	if req.Key == nil {
		return nil, ErrNilKey
	}

	shaKey := sha256.Sum256(req.Key)

	err := ls.Set(shaKey[:], req.Value)
	if err != nil {
		return nil, err
	}

	return &prototypes.StoreMessageResponse{}, nil
}

func (ls *ListenServer) GetMessage(ctx context.Context, req *prototypes.GetMessageRequest) (*prototypes.GetMessageResponse, error) {
	if req == nil {
		return nil, ErrNilRequest
	}

	if req.Key == nil {
		return nil, ErrNilKey
	}

	shaKey := sha256.Sum256(req.Key)

	value, err := ls.Get(shaKey[:])
	if err != nil {
		return nil, err
	}

	return &prototypes.GetMessageResponse{Value: value}, nil
}

func (ls *ListenServer) DeleteMessage(ctx context.Context, req *prototypes.DeleteMessageRequest) (*prototypes.DeleteMessageResponse, error) {
	if req == nil {
		return nil, ErrNilRequest
	}

	if req.Key == nil {
		return nil, ErrNilKey
	}

	shaKey := sha256.Sum256(req.Key)

	err := ls.Delete(shaKey[:])
	if err != nil {
		return nil, err
	}

	return &prototypes.DeleteMessageResponse{}, nil
}

func (ls *ListenServer) SetHashrange(ctx context.Context, req *prototypes.SetHashrangeRequest) (*prototypes.SetHashrangeResponse, error) {
	if req == nil {
		return nil, ErrNilRequest
	}

	ls.hashrange = NewRange(new(big.Int).SetBytes(req.Min), new(big.Int).SetBytes(req.Max))
	return &prototypes.SetHashrangeResponse{}, nil
}
