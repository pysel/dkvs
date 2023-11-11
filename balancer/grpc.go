package balancer

import (
	"context"
	"crypto/sha256"

	"github.com/pysel/dkvs/prototypes"
	pbbalancer "github.com/pysel/dkvs/prototypes/balancer"

	"github.com/pysel/dkvs/types"
)

func (bs *BalancerServer) RegisterPartition(ctx context.Context, req *pbbalancer.RegisterPartitionRequest) (*pbbalancer.RegisterPartitionResponse, error) {
	range_, err := bs.getNextPartitionRange()
	if err != nil {
		return nil, err
	}

	err = bs.Balancer.RegisterPartition(req.Address, *range_)
	if err != nil {
		return nil, err
	}

	// partition successfully registered
	return &pbbalancer.RegisterPartitionResponse{}, nil
}

// ----- To be relayed requests -----

func (bs *BalancerServer) Get(ctx context.Context, req *prototypes.GetRequest) (*prototypes.GetResponse, error) {
	if req == nil {
		return nil, types.ErrNilRequest
	}

	key := req.Key

	if key == "" {
		return nil, types.ErrNilKey
	}

	shaKey := sha256.Sum256([]byte(key))
	range_, err := bs.getRangeFromDigest(shaKey[:])
	if err != nil {
		return nil, err
	}

	responsibleClients := bs.clients[*range_]
	if len(responsibleClients) == 0 {
		return nil, ErrRangeNotYetCovered
	}

	var response *prototypes.GetResponse
	errCounter := 0
	for _, client := range responsibleClients {
		resp, err := client.Get(ctx, &prototypes.GetRequest{Key: key})
		if err != nil {
			errCounter++
			continue
		}

		response = resp
		break // break here since other replicas would return the same value
	}

	if errCounter == len(responsibleClients) {
		return nil, ErrAllReplicasFailed
	}

	return response, nil
}
