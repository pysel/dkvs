package partition

import (
	"context"
	"encoding/binary"
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
	defer func() { ls.postCRUD(err, req.String()) }()

	// note: if request is not valid, the timestamp will not be incremented
	// TODO: investigate if it is a valid behaviour.
	if err = req.Validate(); err != nil {
		return nil, err
	}

	// process logical timestamp
	fmt.Println("Current timestamp: ", ls.timestamp, "Received timestamp: ", req.Lamport)
	switch err = ls.validateTS(req.Lamport); err.(type) {
	case ErrTimestampIsStale: // stale/already processed request
		return nil, err
	case ErrTimestampNotNext: // replica is not ready to process this request
		ls.backlog.Add(types.BID, req.Lamport, req)
		return nil, err // let balancer know that this replica is not ready for the request
	}

	var value []byte
	value, err = reqToBytes(req)
	if err != nil {
		return nil, err
	}

	shaKey := types.ShaKey(req.Key)
	err = ls.Partition.Set(shaKey[:], value)
	if err != nil {
		return nil, ErrInternal{Reason: err}
	}

	// Log event
	ls.eventHandler.Handle(SetEvent{key: string(req.Key), data: string(req.Value)})

	return &prototypes.SetResponse{}, nil
}

// Get gets a value for a key.
func (ls *ListenServer) Get(ctx context.Context, req *prototypes.GetRequest) (resp *prototypes.GetResponse, err error) {
	defer func() { ls.postCRUD(err, req.String()) }()

	if err = req.Validate(); err != nil {
		return nil, err
	}

	// process logical timestamp
	switch err = ls.validateTS(req.Lamport); err.(type) {
	case ErrTimestampIsStale: // stale/already processed request
		return nil, err
	case ErrTimestampNotNext: // replica is not ready to process this request
		ls.backlog.Add(types.BID, req.Lamport, req)
		return nil, err // let balancer know that this replica is not ready for the request
	}

	shaKey := types.ShaKey(req.Key)

	value, err := ls.Partition.Get(shaKey[:])
	if err != nil {
		return nil, ErrInternal{Reason: err}
	}
	if value == nil {
		return &prototypes.GetResponse{StoredValue: nil}, nil
	}

	var storedValue prototypes.StoredValue
	err = proto.Unmarshal(value, &storedValue)
	if err != nil {
		return nil, err
	}

	// Log event
	ls.eventHandler.Handle(GetEvent{key: string(req.Key), returned: string(storedValue.Value)})

	return &prototypes.GetResponse{StoredValue: &storedValue}, nil
}

// Delete deletes a value for a key.
func (ls *ListenServer) Delete(ctx context.Context, req *prototypes.DeleteRequest) (resp *prototypes.DeleteResponse, err error) {
	defer func() { ls.postCRUD(err, req.String()) }()

	if err = req.Validate(); err != nil {
		return nil, err
	}

	// process logical timestamp
	switch err = ls.validateTS(req.Lamport); err.(type) {
	case ErrTimestampIsStale: // stale/already processed request
		return nil, err
	case ErrTimestampNotNext: // replica is not ready to process this request
		ls.backlog.Add(types.BID, req.Lamport, req)
		return nil, err // let balancer know that this replica is not ready for the request
	}

	shaKey := types.ShaKey(req.Key)

	err = ls.Partition.Delete(shaKey[:])
	if err != nil {
		return nil, ErrInternal{Reason: err}
	}

	// Log event
	ls.eventHandler.Handle(DeleteEvent{key: string(req.Key)})

	return &prototypes.DeleteResponse{}, nil
}

// SetHashrange sets the hashrange for this partition.
func (ls *ListenServer) SetHashrange(ctx context.Context, req *prototypes.SetHashrangeRequest) (res *prototypes.SetHashrangeResponse, err error) {
	defer func() {
		if err != nil {
			ls.eventHandler.Handle(ErrorEvent{err: err})
		} else {
			min := binary.LittleEndian.Uint64(req.Min)
			max := binary.LittleEndian.Uint64(req.Max)
			ls.eventHandler.Handle(SetHashrangeEvent{min: min, max: max})
		}
	}()

	if req == nil {
		err = types.ErrNilRequest
		return
	}

	ls.hashrange = NewRange(req.Min, req.Max)
	return &prototypes.SetHashrangeResponse{}, nil
}

// postCRUD runs functionality that should be run after every CRUD operation.
func (p *Partition) postCRUD(err error, req string) {
	p.ProcessBacklog(err)

	// log error as warning if it is either ErrTimestampIsStale or ErrTimestampNotNext
	// log error as error otherwise
	if err != nil {
		switch typedErr := err.(type) {
		case ErrTimestampIsStale:
			p.eventHandler.Handle(typedErr.ToEvent(req))

		case ErrTimestampNotNext:
			p.eventHandler.Handle(typedErr.ToEvent(req))

		default:
			p.eventHandler.Handle(ErrorEvent{req: req, err: typedErr})
		}
	} else {
		p.IncrTs()
	}
}
