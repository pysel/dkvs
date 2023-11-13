package balancer

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/pysel/dkvs/prototypes"
	pbbalancer "github.com/pysel/dkvs/prototypes/balancer"
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
	"google.golang.org/protobuf/proto"

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
	maxLamport := uint64(0)
	for _, client := range responsibleClients {
		resp, err := client.Get(ctx, &prototypes.GetRequest{Key: key})
		if err != nil {
			continue
		}

		// since returned value will be a tuple of lamport timestamp and value, check which returned value
		// has the highest lamport timestamp
		var storedValue pbpartition.StoredValue
		err = proto.Unmarshal(resp.Value, &storedValue)
		if err != nil {
			// TODO: partition is in incorrect state, should remove it from active set
			fmt.Println("Error unmarshalling value from partition", err)
			continue
		}

		if storedValue.Lamport > maxLamport {
			maxLamport = storedValue.Lamport
			response = resp
		}
	}

	if response == nil {
		return nil, ErrAllReplicasFailed
	}

	return response, nil
}
