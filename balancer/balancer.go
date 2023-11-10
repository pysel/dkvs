package balancer

import (
	"crypto/sha256"

	"github.com/pysel/dkvs/partition"
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
)

type Balancer struct {
	// A mapping from ranges to partitions.
	// Multiple partitions can be mapped to the same range.
	clients map[partition.Range][]pbpartition.PartitionServiceClient

	// goalPartitions is the number of partitions we are expecting
	goalPartitions int
	// activePartitions is the number of currently registered partitions
	activePartitions int
}

func NewBalancer(goalPartitions int) *Balancer {
	return &Balancer{
		clients:          make(map[partition.Range][]pbpartition.PartitionServiceClient),
		goalPartitions:   goalPartitions,
		activePartitions: 0,
	}
}

// AddPartition adds a partition to the balancer.
func (b *Balancer) RegisterPartition(addr string, range_ partition.Range) error {
	if b.activePartitions == b.goalPartitions {
		return ErrPartitionOverflow
	}

	client := NewPartitionClient(addr)
	b.clients[range_] = append(b.clients[range_], client)
	return nil
}

// GetPartitions returns a list of partitions that contain the given key.
func (b *Balancer) GetPartitions(key []byte) []pbpartition.PartitionServiceClient {
	shaKey := sha256.Sum256(key)
	for range_, clients := range b.clients {
		if range_.Contains(shaKey[:]) {
			return clients
		}
	}

	return nil
}
