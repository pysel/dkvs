package balancer

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/pysel/dkvs/partition"
	pbbalancer "github.com/pysel/dkvs/prototypes/balancer"
)

var CreatedCoverage *coverage

func newTick(value *big.Int) *pbbalancer.Tick {
	return &pbbalancer.Tick{
		MinOf:               0,
		MaxOf:               0,
		Value:               value.Bytes(),
		NextInitialized:     nil,
		PreviousInitialized: nil,
	}
}

// coverage is a linked list of initialized ticks.
type coverage struct {
	*pbbalancer.Tick
	size int // TODO: consider removing this
}

func (c *coverage) String() string {
	curTick := c.Tick
	str := ""
	for curTick != nil {
		bigInt := new(big.Int).SetBytes(curTick.Value)
		str += fmt.Sprintf("%v", bigInt)
		curTick = curTick.NextInitialized
		if curTick != nil {
			str += " -> "
		}
	}
	return str
}

// GetCoverage returns a coverage.
// Singletone pattern is used here.
func GetCoverage() *coverage {
	if CreatedCoverage == nil {
		CreatedCoverage = &coverage{nil, 0}
	}
	return CreatedCoverage
}

// addTick iterates over the list of ticks until
func (c *coverage) addTick(t *pbbalancer.Tick, isMin, isMax bool) {
	if isMin {
		t.MinOf++
	}

	if isMax {
		t.MaxOf++
	}

	// Cover case when to-add tick is the first one
	if c.Tick == nil {
		c.size++
		c.Tick = t
		return
	}

	curTick := c.Tick
	if bytes.Equal(curTick.Value, t.Value) {
		curTick.MinOf += t.MinOf
		curTick.MaxOf += t.MaxOf
		return
	} else if bytes.Compare(curTick.Value, t.Value) == 1 { // means tick is lower than every value in list => add to the beginning
		curTick.PreviousInitialized = t
		t.NextInitialized = curTick
		c.size++
		c.Tick = t
		return
	}
	nextTick := curTick.NextInitialized

	for nextTick != nil {
		if bytes.Compare(nextTick.Value, t.Value) == 1 {
			curTick.NextInitialized = t
			nextTick.PreviousInitialized = t

			t.NextInitialized = nextTick
			t.PreviousInitialized = curTick

			c.size++
			return
		} else if bytes.Equal(nextTick.Value, t.Value) { // means tick is already covered
			nextTick.MinOf += t.MinOf
			nextTick.MaxOf += t.MaxOf
			return
		}
		curTick = curTick.NextInitialized
		nextTick = curTick.NextInitialized
	}

	// if not returned at this point, means the value of a t is higher than every value in list => add to the end
	curTick.NextInitialized = t
	t.PreviousInitialized = curTick
	c.size++
}

// getNextPartitionRange is used when assigning a range to a newly registered partition
func (c *coverage) getNextPartitionRange() (*partition.Range, *pbbalancer.Tick, *pbbalancer.Tick) {
	// initially assume that first interval is minimal
	minCovered := c.Tick.MinOf + c.Tick.NextInitialized.MaxOf
	minLowerTick := c.Tick
	minUpperTick := c.Tick.NextInitialized
	minRange := partition.NewRange(minLowerTick.Value, minUpperTick.Value)
	for tick := c.Tick; tick.NextInitialized != nil; tick = tick.NextInitialized {
		coveredBy := tick.MinOf + tick.NextInitialized.MaxOf
		if coveredBy < minCovered {
			minRange = partition.NewRange(tick.Value, tick.NextInitialized.Value)
			minCovered = coveredBy
			minLowerTick = tick
			minUpperTick = tick.NextInitialized
		}
	}

	// minLowerTick and minUpperTick are returned to be increased by 1 if a partition is successfully registered
	return minRange, minLowerTick, minUpperTick
}

func (c *coverage) bumpTicks(lowerTick, upperTick *pbbalancer.Tick) {
	lowerTick.MinOf++
	upperTick.MaxOf++
}

// ToProto converts coverage to protobuf coverage
func (c *coverage) ToProto() *pbbalancer.Coverage {
	return &pbbalancer.Coverage{
		Tick: c.Tick,
		Size: int64(c.size),
	}
}
