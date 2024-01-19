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

func (bs *BalancerServer) Get(ctx context.Context, req *prototypes.GetRequest) (*prototypes.GetResponse, error) {
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

func (bs *BalancerServer) Set(ctx context.Context, req *prototypes.SetRequest) (*prototypes.SetResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	shaKey := types.ShaKey(req.Key)
	range_, err := bs.getRangeFromDigest(shaKey[:])
	if err != nil {
		return nil, err
	}

	msg := shared.NewPrepareCommitMessage_Set(req.Key, req.Value)

	err = bs.AtomicMessage(ctx, range_, msg)
	if err != nil {
		return nil, err
	}

	bs.eventHandler.Emit(&SetEvent{msg: req.String()})

	return &prototypes.SetResponse{}, nil
}

func (bs *BalancerServer) Delete(ctx context.Context, req *prototypes.DeleteRequest) (*prototypes.DeleteResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	shaKey := types.ShaKey(req.Key)
	range_, err := bs.getRangeFromDigest(shaKey[:])
	if err != nil {
		return nil, err
	}

	msg := shared.NewPrepareCommitMessage_Delete(req.Key)

	err = bs.AtomicMessage(ctx, range_, msg)
	if err != nil {
		return nil, err
	}

	bs.eventHandler.Emit(&DeleteEvent{msg: req.String()})

	return &prototypes.DeleteResponse{}, nil
}
