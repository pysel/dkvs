package balancer

import (
	"bytes"
	"math/big"
	"testing"

	pbbalancer "github.com/pysel/dkvs/prototypes/balancer"
	"github.com/pysel/dkvs/testutil"
	"github.com/stretchr/testify/require"
)

var (
	zeroInt          = new(big.Int).SetInt64(0)
	quarterInt       = new(big.Int).Div(testutil.HalfShaDomain, big.NewInt(2))
	halfInt          = testutil.HalfShaDomain
	threeQuartersInt = new(big.Int).Mul(quarterInt, big.NewInt(3))
	fullInt          = new(big.Int).Mul(testutil.HalfShaDomain, big.NewInt(2))
)

func TestGetTickByValue(t *testing.T) {
	defaultCoverage_ := defaulCoverage(t)

	tests := map[string]struct {
		value    *big.Int
		coverage *coverage
		expected *pbbalancer.Tick
	}{
		"Get tick at the beginning": {
			value:    zeroInt,
			coverage: defaultCoverage_,
			expected: defaultCoverage_.ticks[0],
		},
		"Get tick at the end": {
			value:    fullInt,
			coverage: defaultCoverage_,
			expected: defaultCoverage_.ticks[len(defaultCoverage_.ticks)-1],
		},
		"Get second tick": {
			value:    quarterInt,
			coverage: defaultCoverage_,
			expected: defaultCoverage_.ticks[1],
		},
		"Get tick that doesn't exist": {
			value:    new(big.Int).SetInt64(-1),
			coverage: defaultCoverage_,
			expected: nil,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			tick := test.coverage.getTickByValue(test.value)
			if test.expected != nil {
				tickDeepEqual(t, tick, test.expected)
			} else {
				require.Nil(t, test.expected)
			}
		})
	}
}

func TestAddTick(t *testing.T) {
	defaultCoverage_ := defaulCoverage(t)

	tests := map[string]struct {
		toAdd             *pbbalancer.Tick
		coverage          *coverage
		expectedTick      *pbbalancer.Tick
		expectedTickValue *big.Int
	}{
		"Add tick at the beginning": {
			toAdd:    newTick(new(big.Int).SetInt64(1), 1),
			coverage: &coverage{nil},
			expectedTick: &pbbalancer.Tick{
				Covers: 1,
				Value:  new(big.Int).SetInt64(1).Bytes(),
			},
			expectedTickValue: new(big.Int).SetInt64(1),
		},
		"Add tick at the end": {
			toAdd:    newTick(new(big.Int).Mul(fullInt, big.NewInt(2)), 0),
			coverage: defaultCoverage_,
			expectedTick: &pbbalancer.Tick{
				Value: new(big.Int).Mul(fullInt, big.NewInt(2)).Bytes(),
			},
			expectedTickValue: new(big.Int).Mul(fullInt, big.NewInt(2)),
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			test.coverage.addTick(test.toAdd)
			tickDeepEqual(t, test.coverage.getTickByValue(test.expectedTickValue), test.expectedTick)
		})
	}
}

// defaulCoverage creates a coverage with the following ticks:
// - tick at 0 [min]
// - tick at 1/4 of the domain [min and max]
// - tick at 1/2 of the domain [max]
// - tick at 3/4 of the domain [min and max]
// - tick at the end of the domain [max]
//
// Visually ("-" denotes areas that are covered):
//
// 0-----1/4-----1/2     3/4-----1
func defaulCoverage(t *testing.T) *coverage {
	coverage := &coverage{nil}

	// zeroNull corresponds to tick at 0
	zeroNull := newTick(zeroInt, 1)

	// tickQuarter corresponds to tick at 1/4 of the domain
	tickQuarter := newTick(quarterInt, 1)

	// tickHalf corresponds to tick at 1/2 of the domain
	tickHalf := newTick(halfInt, 0)

	// tickThreeQuarters corresponds to tick at 3/4 of the domain
	tickThreeQuarters := newTick(threeQuartersInt, 1)

	// tickFull corresponds to tick at the end of the domain
	tickFull := newTick(fullInt, 0)

	coverage.addTick(zeroNull)
	coverage.addTick(tickQuarter)
	coverage.addTick(tickHalf)
	coverage.addTick(tickThreeQuarters)
	coverage.addTick(tickFull)

	assertDefaultcoverage(t, coverage)

	return coverage
}

func assertDefaultcoverage(t *testing.T, c *coverage) {
	require.Equal(t, 5, len(c.ticks))
	one64 := int64(1)
	zero64 := int64(0)

	firstTick := c.ticks[0]
	require.Equal(t, zeroInt.Bytes(), firstTick.Value)
	require.Equal(t, one64, firstTick.Covers)

	secondTick := c.ticks[1]
	require.Equal(t, quarterInt.Bytes(), secondTick.Value)
	require.Equal(t, one64, secondTick.Covers)

	thirdTick := c.ticks[2]
	require.Equal(t, halfInt.Bytes(), thirdTick.Value)
	require.Equal(t, zero64, thirdTick.Covers)

	fourthTick := c.ticks[3]
	require.Equal(t, threeQuartersInt.Bytes(), fourthTick.Value)
	require.Equal(t, one64, fourthTick.Covers)

	fifthTick := c.ticks[4]
	require.Equal(t, fullInt.Bytes(), fifthTick.Value)
	require.Equal(t, zero64, fifthTick.Covers)
}

func tickDeepEqual(t *testing.T, expected, actual *pbbalancer.Tick) {
	require.Equal(t, bytes.Compare(expected.Value, actual.Value), 0)
	require.Equal(t, expected.Covers, actual.Covers)
}

func (c *coverage) getTickByValue(value *big.Int) *pbbalancer.Tick {
	for _, tick := range c.ticks {
		if bytes.Equal(tick.Value, value.Bytes()) {
			return tick
		}
	}

	return nil
}
