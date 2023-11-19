package balancer

import (
	"math/big"

	"github.com/pysel/dkvs/partition"
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
)

func (b *Balancer) GetTickByValue(value *big.Int) *tick {
	return b.coverage.getTickByValue(value)
}

func (b *Balancer) GetTicksAmount() int {
	return b.coverage.size
}

func (b *Balancer) GetNextPartitionRange() (*partition.Range, error) {
	return b.getNextPartitionRange()
}

func (b *Balancer) SetActivePartitions(amount int) {
	b.activePartitions = amount
}

// NewBalancerTest returns a new balancer instance with an independent coverage every time.
func NewBalancerTest(goalReplicaRanges int) *Balancer {
	b := &Balancer{
		clients:           make(map[*partition.Range][]pbpartition.PartitionServiceClient),
		goalReplicaRanges: goalReplicaRanges,
		activePartitions:  0,
		coverage:          &coverage{},
	}

	b.setupCoverage()

	return b
}
