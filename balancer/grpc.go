package balancer

import (
	"context"

	"github.com/pysel/dkvs/prototypes"
	pbbalancer "github.com/pysel/dkvs/prototypes/balancer"
	"github.com/pysel/dkvs/shared"
	"github.com/pysel/dkvs/types"
)

// RegisterPartition registers a partition in the balancer.
func (bs *BalancerServer) RegisterPartition(ctx context.Context, req *pbbalancer.RegisterPartitionRequest) (*pbbalancer.RegisterPartitionResponse, error) {
	err := bs.Balancer.RegisterPartition(ctx, req.Address)
	if err != nil {
		return nil, err
	}

	bs.eventHandler.Emit(&RegisterPartitionEvent{Address: req.Address})

	// partition successfully registered
	return &pbbalancer.RegisterPartitionResponse{}, nil
}

// ----- To be relayed requests -----

func (bs *BalancerServer) Get(ctx context.Context, req *prototypes.GetRequest) (res *prototypes.GetResponse, err error) {
	defer func() { bs.postCRUD(err, req.String()) }()

	if err := req.Validate(); err != nil {
		return nil, err
	}

	response, err := bs.Balancer.Get(ctx, req.Key)
	if err != nil {
		return nil, err
	}

	bs.eventHandler.Emit(&GetEvent{msg: req.String()})

	return response, nil
}

func (bs *BalancerServer) Set(ctx context.Context, req *prototypes.SetRequest) (res *prototypes.SetResponse, err error) {
	defer func() { bs.postCRUD(err, req.String()) }()

	if err := req.Validate(); err != nil {
		return nil, err
	}

	shaKey := types.ShaKey(req.Key)
	range_, err := bs.getRangeFromDigest(shaKey[:])
	if err != nil {
		return nil, err
	}

	lamport := bs.Balancer.GetNextLamportForKey(req.Key)
	msg := shared.NewPrepareCommitMessage_Set(req.Key, req.Value, lamport)

	err = bs.AtomicMessage(ctx, range_, msg)
	if err != nil {
		return nil, err
	}

	bs.eventHandler.Emit(&SetEvent{msg: req.String()})

	return &prototypes.SetResponse{}, nil
}

func (bs *BalancerServer) Delete(ctx context.Context, req *prototypes.DeleteRequest) (res *prototypes.DeleteResponse, err error) {
	defer func() { bs.postCRUD(err, req.String()) }()

	if err := req.Validate(); err != nil {
		return nil, err
	}

	shaKey := types.ShaKey(req.Key)
	range_, err := bs.getRangeFromDigest(shaKey[:])
	if err != nil {
		return nil, err
	}

	lamport := bs.Balancer.GetNextLamportForKey(req.Key)
	msg := shared.NewPrepareCommitMessage_Delete(req.Key, lamport)

	err = bs.AtomicMessage(ctx, range_, msg)
	if err != nil {
		return nil, err
	}

	bs.eventHandler.Emit(&DeleteEvent{msg: req.String()})

	return &prototypes.DeleteResponse{}, nil
}

func (ls *BalancerServer) postCRUD(err error, req string) {
	if err != nil {
		if eventError, ok := err.(shared.IsWarningEventError); ok {
			ls.eventHandler.Emit(eventError.WarningErrorToEvent(req))
			return
		}

		ls.eventHandler.Emit(&shared.ErrorEvent{Req: req, Err: err})
	}
}
