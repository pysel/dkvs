package balancer

import (
	"math/big"

	"github.com/pysel/dkvs/partition"
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
