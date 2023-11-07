package balancer

import (
	"context"

	pbbalancer "github.com/pysel/dkvs/prototypes/balancer"
)

func (b *BalancerServer) RegisterPartition(ctx context.Context, req *pbbalancer.RegisterPartitionRequest) (*pbbalancer.RegisterPartitionResponse, error) {
	return nil, nil
}
