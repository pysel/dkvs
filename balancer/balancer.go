package balancer

import (
	"crypto/sha256"
	"math/big"

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

	// coverage is used for tracking the tracked ranges
	coverage *coverage
}

// NewBalancer returns a new balancer instance.
func NewBalancer(goalPartitions int) *Balancer {
	b := &Balancer{
		clients:          make(map[partition.Range][]pbpartition.PartitionServiceClient),
		goalPartitions:   goalPartitions,
		activePartitions: 0,
		coverage:         GetCoverage(),
	}

	b.setupCoverage()

	return b
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

// setupCoverage creates necessary ticks for coverage based on goalPartitions
func (b *Balancer) setupCoverage() {
	if b.goalPartitions == 0 {
		b.coverage.addTick(newTick(big.NewInt(0)), false, false)
		b.coverage.addTick(newTick(partition.MaxInt), false, false)
		return
	}
	// Create a tick for each partition
	for i := 0; i <= b.goalPartitions; i++ {
		numerator := new(big.Int).Mul(big.NewInt(int64(i)), partition.MaxInt)
		value := new(big.Int).Div(numerator, big.NewInt(int64(b.goalPartitions)))
		b.coverage.addTick(newTick(value), false, false)
	}
}
