package balancer

import (
	"fmt"
	"math/big"
)

var CreatedCoverage *coverage

type tick struct {
	value               *big.Int
	minOf               int
	maxOf               int
	previousInitialized *tick
	nextInitialized     *tick
}

func newTick(value *big.Int) *tick {
	return &tick{
		minOf:               0,
		maxOf:               0,
		value:               value,
		nextInitialized:     nil,
		previousInitialized: nil,
	}
}

func (t *tick) next() *tick {
	return t.nextInitialized
}

// coverage is a linked list of initialized ticks.
type coverage struct {
	*tick
	size int
}

func (c *coverage) String() string {
	curTick := c.tick
	str := ""
	for curTick != nil {
		str += fmt.Sprintf("%v", curTick.value)
		curTick = curTick.next()
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
func (c *coverage) addTick(t *tick, isMin, isMax bool) {
	if isMin {
		t.minOf++
	}

	if isMax {
		t.maxOf++
	}

	// Cover case when to-add tick is the first one
	if c.tick == nil {
		c.size++
		c.tick = t
		return
	}

	curTick := c.tick
	if curTick.value.Cmp(t.value) == 0 {
		curTick.minOf += t.minOf
		curTick.maxOf += t.maxOf
		return
	} else if curTick.value.Cmp(t.value) == 1 { // means tick is lower than every value in list => add to the beginning
		curTick.previousInitialized = t
		t.nextInitialized = curTick
		c.size++
		c.tick = t
		return
	}
	nextTick := curTick.next()

	for nextTick != nil {
		if nextTick.value.Cmp(t.value) == 1 {
			curTick.nextInitialized = t
			nextTick.previousInitialized = t

			t.nextInitialized = nextTick
			t.previousInitialized = curTick

			c.size++
			return
		} else if nextTick.value.Cmp(t.value) == 0 { // means tick is already covered
			nextTick.minOf += t.minOf
			nextTick.maxOf += t.maxOf
			return
		}
		curTick = curTick.next()
		nextTick = curTick.next()
	}

	// if not returned at this point, means the value of a t is higher than every value in list => add to the end
	curTick.nextInitialized = t
	t.previousInitialized = curTick
	c.size++
}

func (c *coverage) getTickByValue(value *big.Int) *tick {
	curTick := c.tick
	if curTick.value.Cmp(value) == 0 {
		return curTick
	}

	for curTick != nil && curTick.value.Cmp(value) != 1 {
		if curTick.value.Cmp(value) == 0 {
			return curTick
		}
		curTick = curTick.next()
	}

	return nil
}
