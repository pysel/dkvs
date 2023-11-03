package balancer

import (
	pclient "github.com/pysel/dkvs/balancer/partition-client"
	"github.com/pysel/dkvs/partition"
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
)

type Balancer struct {
	// A mapping from ranges to partitions.
	// Multiple partitions can be mapped to the same range.
	clients map[partition.Range][]pbpartition.PartitionServiceClient
}

func NewBalancer() *Balancer {
	return &Balancer{
		clients: make(map[partition.Range][]pbpartition.PartitionServiceClient),
	}
}

// AddPartition adds a partition to the balancer.
func (b *Balancer) AddPartition(addr string, range_ partition.Range) {
	client := pclient.NewPartitionClient(addr)
	b.clients[range_] = append(b.clients[range_], client)
}

// GetPartitions returns a list of partitions that contain the given key.
func (b *Balancer) GetPartitions(key []byte) []pbpartition.PartitionServiceClient {
	for range_, clients := range b.clients {
		if range_.Contains(key) {
			return clients
		}
	}

	return nil
}
