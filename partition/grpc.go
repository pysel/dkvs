package partition

import (
	"context"
	"fmt"
	"net"

	"github.com/pysel/dkvs/prototypes"
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
	"github.com/pysel/dkvs/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/proto"
)

type ListenServer struct {
	pbpartition.UnimplementedPartitionServiceServer
	*Partition
}

// RunPartitionServer starts a partition server on the given port.
func RunPartitionServer(port int64, dbPath string) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	partition := NewPartition(dbPath)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pbpartition.RegisterPartitionServiceServer(grpcServer, &ListenServer{Partition: partition})
	fmt.Println("Starting server on port", port)
	return grpcServer.Serve(lis)
}

// Set sets a value for a key.
func (ls *ListenServer) Set(ctx context.Context, req *prototypes.SetRequest) (resp *prototypes.SetResponse, err error) {
	defer ls.postCRUD(err)

	// note: if request is not valid, the timestamp will not be incremented
	// TODO: investigate if it is a valid behaviour.
	if err := req.Validate(); err != nil {
		return nil, err
	}

	// process logical timestamp
	switch ls.validateTS(req.Lamport) {
	case ErrTimestampLessThanCurrent: // wrong: stale request
		return nil, ErrTimestampLessThanCurrent
	case ErrTimestampNotNext{CurrentTimestamp: ls.timestamp}: // replica is not ready to process this request
		ls.backlog.Add(types.BID, req.Lamport, req)
		return nil, ErrTimestampNotNext{CurrentTimestamp: ls.timestamp} // let balancer know that this replica is not ready for the request
	}

	value, err := reqToBytes(req)
	if err != nil {
		return nil, err
	}

	shaKey := types.ShaKey(req.Key)
	err = ls.Partition.Set(shaKey[:], value)
	if err != nil {
		return nil, err
	}

	return &prototypes.SetResponse{}, nil
}

// Get gets a value for a key.
func (ls *ListenServer) Get(ctx context.Context, req *prototypes.GetRequest) (resp *prototypes.GetResponse, err error) {
	defer ls.postCRUD(err)

	if err := req.Validate(); err != nil {
		return nil, err
	}

	shaKey := types.ShaKey(req.Key)

	value, err := ls.Partition.Get(shaKey[:])
	if err != nil {
		return nil, err
	}
	if value == nil {
		return &prototypes.GetResponse{StoredValue: nil}, nil
	}

	var storedValue prototypes.StoredValue
	err = proto.Unmarshal(value, &storedValue)
	if err != nil {
		return nil, err
	}

	return &prototypes.GetResponse{StoredValue: &storedValue}, nil
}

// Delete deletes a value for a key.
func (ls *ListenServer) Delete(ctx context.Context, req *prototypes.DeleteRequest) (resp *prototypes.DeleteResponse, err error) {
	defer ls.postCRUD(err)

	if err := req.Validate(); err != nil {
		ls.IncrTs()
		return nil, err
	}

	// process logical timestamp
	switch ls.validateTS(req.Lamport) {
	case ErrTimestampLessThanCurrent: // stale/already processed request
		return nil, ErrTimestampLessThanCurrent
	case ErrTimestampNotNext{}: // replica is not ready to process this request
		ls.backlog.Add(types.BID, req.Lamport, req)
		return nil, ErrTimestampNotNext{CurrentTimestamp: ls.timestamp} // let balancer know that this replica is not ready for the request
	}

	shaKey := types.ShaKey(req.Key)

	err = ls.Partition.Delete(shaKey[:])
	if err != nil {
		return nil, err
	}

	return &prototypes.DeleteResponse{}, nil
}

// SetHashrange sets the hashrange for this partition.
func (ls *ListenServer) SetHashrange(ctx context.Context, req *prototypes.SetHashrangeRequest) (*prototypes.SetHashrangeResponse, error) {
	if req == nil {
		return nil, types.ErrNilRequest
	}

	ls.hashrange = NewRange(req.Min, req.Max)
	return &prototypes.SetHashrangeResponse{}, nil
}

// postCRUD runs functionality that should be run after every CRUD operation.
func (p *Partition) postCRUD(err error) {
	p.ProcessBacklog(err)
	p.IncrTs()
}
