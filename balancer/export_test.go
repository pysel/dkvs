package balancer

import "math/big"

func (b *Balancer) GetTickByValue(value *big.Int) *tick {
	return b.coverage.getTickByValue(value)
}

func (b *Balancer) GetTicksAmount() int {
	return b.coverage.size
}
