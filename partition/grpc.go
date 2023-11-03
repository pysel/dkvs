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
	"google.golang.org/grpc/reflection"
)

type ListenServer struct {
	pbpartition.UnimplementedPartitionServiceServer
	*Partition
}

// RunPartitionServer starts a partition server on the given port.
func RunPartitionServer(port int64, dbPath string) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	partition := NewPartition(dbPath)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pbpartition.RegisterPartitionServiceServer(grpcServer, &ListenServer{Partition: partition})
	fmt.Println("Starting server on port", port)
	grpcServer.Serve(lis)
}

// SetMessage sets a value for a key.
func (ls *ListenServer) SetMessage(ctx context.Context, req *prototypes.SetMessageRequest) (*prototypes.SetMessageResponse, error) {
	if req == nil {
		return nil, ErrNilRequest
	}

	if req.Key == "" {
		return nil, ErrNilKey
	}

	if req.Value == nil {
		return nil, ErrNilValue
	}

	shaKey := shaKey(req.Key)
	valueBz := []byte(req.Value)

	err := ls.Set(shaKey[:], valueBz)
	if err != nil {
		return nil, err
	}

	return &prototypes.SetMessageResponse{}, nil
}

// GetMessage gets a value for a key.
func (ls *ListenServer) GetMessage(ctx context.Context, req *prototypes.GetMessageRequest) (*prototypes.GetMessageResponse, error) {
	if req == nil {
		return nil, ErrNilRequest
	}

	if req.Key == "" {
		return nil, ErrNilKey
	}

	shaKey := shaKey(req.Key)

	value, err := ls.Get(shaKey[:])
	if err != nil {
		return nil, err
	}

	return &prototypes.GetMessageResponse{Value: value}, nil
}

// DeleteMessage deletes a value for a key.
func (ls *ListenServer) DeleteMessage(ctx context.Context, req *prototypes.DeleteMessageRequest) (*prototypes.DeleteMessageResponse, error) {
	if req == nil {
		return nil, ErrNilRequest
	}

	if req.Key == "" {
		return nil, ErrNilKey
	}
	shaKey := shaKey(req.Key)

	err := ls.Delete(shaKey[:])
	if err != nil {
		return nil, err
	}

	return &prototypes.DeleteMessageResponse{}, nil
}

// SetHashrange sets the hashrange for this partition.
func (ls *ListenServer) SetHashrange(ctx context.Context, req *prototypes.SetHashrangeRequest) (*prototypes.SetHashrangeResponse, error) {
	if req == nil {
		return nil, ErrNilRequest
	}

	ls.hashrange = NewRange(new(big.Int).SetBytes(req.Min), new(big.Int).SetBytes(req.Max))
	return &prototypes.SetHashrangeResponse{}, nil
}

func shaKey(key string) []byte {
	checksum := sha256.Sum256([]byte(key))
	return checksum[:]
}
