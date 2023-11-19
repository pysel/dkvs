package balancer

import (
	"crypto/sha256"
	"math/big"

	"github.com/pysel/dkvs/partition"
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
)

// Balancer is a node that is responsible for registering partitions and relaying requests to appropriate ones.
type Balancer struct {
	// A mapping from ranges to partitions.
	// Multiple partitions can be mapped to the same range.
	clients map[*partition.Range][]pbpartition.PartitionServiceClient

	// goalReplicaRanges is the number of different sets of replicas that should be created
	goalReplicaRanges int
	// activePartitions is the number of currently registered partitions
	activePartitions int

	// coverage is used for tracking the tracked ranges
	coverage *coverage
}

// NewBalancer returns a new balancer instance.
func NewBalancer(goalReplicaRanges int) *Balancer {
	b := &Balancer{
		clients:           make(map[*partition.Range][]pbpartition.PartitionServiceClient),
		goalReplicaRanges: goalReplicaRanges,
		activePartitions:  0,
		coverage:          GetCoverage(),
	}

	b.setupCoverage()

	return b
}

// AddPartition adds a partition to the balancer.
func (b *Balancer) RegisterPartition(addr string, range_ *partition.Range) error {
	if b.activePartitions == b.goalReplicaRanges {
		return ErrPartitionOverflow
	}

	client := partition.NewPartitionClient(addr)
	b.clients[range_] = append(b.clients[range_], client)
	b.activePartitions++

	return nil
}

// GetPartitions returns a list of partitions that contain the given key.
func (b *Balancer) GetPartitionsByKey(key []byte) []pbpartition.PartitionServiceClient {
	shaKey := sha256.Sum256(key)
	for range_, clients := range b.clients {
		if range_.Contains(shaKey[:]) {
			return clients
		}
	}

	return nil
}

// setupCoverage creates necessary ticks for coverage based on goalReplicaRanges
func (b *Balancer) setupCoverage() {
	if b.goalReplicaRanges == 0 {
		b.coverage.addTick(newTick(big.NewInt(0)), false, false)
		b.coverage.addTick(newTick(partition.MaxInt), false, false)
		return
	}

	// Create a tick for each partition
	for i := 0; i <= b.goalReplicaRanges; i++ {
		numerator := new(big.Int).Mul(big.NewInt(int64(i)), partition.MaxInt)
		value := new(big.Int).Div(numerator, big.NewInt(int64(b.goalReplicaRanges)))
		b.coverage.addTick(newTick(value), false, false)
	}
}

// getNextPartitionRange returns a range to which the next partition should be assigned
func (b *Balancer) getNextPartitionRange() (*partition.Range, error) {
	remainder := b.activePartitions % b.goalReplicaRanges
	tickIndexMin := remainder
	index := 0
	for tick := b.coverage.tick; tick != nil; tick = tick.next() {
		if index == tickIndexMin {
			min := tick.value
			max := tick.next().value
			return partition.NewRange(min, max), nil
		}
		index++
	}

	return nil, ErrCoverageNotProperlySetUp
}

// getRangeFromDigest returns a range to which the given digest belongs
func (b *Balancer) getRangeFromDigest(digest []byte) (*partition.Range, error) {
	for range_ := range b.clients {
		if range_.Contains(digest) {
			return range_, nil
		}
	}

	return nil, ErrDigestNotCovered
}
