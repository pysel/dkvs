package balancer

import (
	"bytes"
	"fmt"
	"math/big"

	pbbalancer "github.com/pysel/dkvs/prototypes/balancer"
	"github.com/pysel/dkvs/types/hrange"
)

var CreatedCoverage *coverage

func newTick(value *big.Int, covers int64) *pbbalancer.Tick {
	return &pbbalancer.Tick{
		Covers: covers,
		Value:  value.Bytes(),
	}
}

// coverage is a linked list of initialized ticks.
type coverage struct{ ticks []*pbbalancer.Tick }

func (c *coverage) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Coverage:")
	for _, tick := range c.ticks {
		buffer.WriteString("\n | Tick: ")
		buffer.WriteString("\n | | Value: " + new(big.Int).SetBytes(tick.Value).String())
		buffer.WriteString("\n | | Covers: " + fmt.Sprint(tick.Covers))
		buffer.WriteString("\n |\n |")
	}
	return buffer.String()
}

// GetCoverage returns a coverage.
// Singletone pattern is used here.
func GetCoverage() *coverage {
	if CreatedCoverage == nil {
		CreatedCoverage = &coverage{nil}
	}
	return CreatedCoverage
}

// addTick iterates over the list of ticks until
func (c *coverage) addTick(t *pbbalancer.Tick) {
	if len(c.ticks) == 0 {
		c.ticks = append(c.ticks, t)
		return
	}

	// find the tick that is greater than the new tick
	ind := 0
	for ; ind < len(c.ticks); ind++ {
		if bytes.Compare(c.ticks[ind].Value, t.Value) > 0 {
			break
		}
	}

	// if the tick is not found, append it to the end
	if ind == len(c.ticks) {
		c.ticks = append(c.ticks, t)
	} else {
		// if the tick is found, insert it
		c.ticks = append(append(c.ticks[:ind+1], t), c.ticks[ind+1:]...)
	}
}

// getNextPartitionRange is used when assigning a range to a newly registered partition
func (c *coverage) getNextPartitionRange() (hrange.RangeKey, *pbbalancer.Tick, *pbbalancer.Tick) {
	// initially assume that first interval is minimal
	minCovered := c.ticks[0].Covers
	minLowerTick := c.ticks[0]
	minUpperTick := c.ticks[1]
	minRange := hrange.NewRange(minLowerTick.Value, minUpperTick.Value)
	for ind, tick := range c.ticks[:len(c.ticks)-1] { // no need to cover last
		nextTick := c.ticks[ind+1]
		if tick.Covers < minCovered {
			minRange = hrange.NewRange(tick.Value, nextTick.Value)
			minCovered = tick.Covers
			minLowerTick = tick
			minUpperTick = nextTick
		}
	}

	// minLowerTick and minUpperTick are returned to be increased by 1 if a partition is successfully registered
	return hrange.RangeKey(minRange.AsKey()), minLowerTick, minUpperTick
}

func (c *coverage) bumpTicks(lowerTick *pbbalancer.Tick) {
	lowerTick.Covers++
}

// ToProto converts coverage to protobuf coverage
func (c *coverage) ToProto() *pbbalancer.Coverage {
	return &pbbalancer.Coverage{
		Ticks: c.ticks,
	}
}
