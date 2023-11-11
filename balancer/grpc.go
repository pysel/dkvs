package balancer

import (
	"context"

	pbbalancer "github.com/pysel/dkvs/prototypes/balancer"
)

func (b *BalancerServer) RegisterPartition(ctx context.Context, req *pbbalancer.RegisterPartitionRequest) (*pbbalancer.RegisterPartitionResponse, error) {
	range_, err := b.getNextPartitionRange()
	if err != nil {
		return nil, err
	}

	err = b.Balancer.RegisterPartition(req.Address, *range_)
	if err != nil {
		return nil, err
	}

	// partition successfully registered
	return &pbbalancer.RegisterPartitionResponse{}, nil
}
