package balancer

import (
	"math/big"
	"testing"

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
		expected *tick
	}{
		"Get tick at the beginning": {
			value:    zeroInt,
			coverage: defaultCoverage_,
			expected: defaultCoverage_.tick,
		},
		"Get tick at the end": {
			value:    fullInt,
			coverage: defaultCoverage_,
			expected: defaultCoverage_.tick.next().next().next().next(),
		},
		"Get tick in the middle": {
			value:    quarterInt,
			coverage: defaultCoverage_,
			expected: defaultCoverage_.tick.next(),
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
		toAdd             *tick
		coverage          *coverage
		isMin             bool
		isMax             bool
		expectedTick      *tick
		expectedTickValue *big.Int
	}{
		"Add tick at the beginning": {
			toAdd:    newTick(new(big.Int).SetInt64(1)),
			coverage: &coverage{nil, 0},
			isMin:    true,
			isMax:    false,
			expectedTick: &tick{
				minOf: 1,
				maxOf: 0,
				value: new(big.Int).SetInt64(1),
			},
			expectedTickValue: new(big.Int).SetInt64(1),
		},
		"Add tick that already exists": {
			toAdd:    newTick(zeroInt),
			coverage: defaultCoverage_,
			isMin:    true,
			isMax:    false,
			expectedTick: &tick{
				minOf: 2,
				maxOf: 0,
				value: zeroInt,
			},
			expectedTickValue: zeroInt,
		},
		"Add tick at the end and already exists": {
			toAdd:    newTick(fullInt),
			coverage: defaultCoverage_,
			isMin:    false,
			isMax:    true,
			expectedTick: &tick{
				minOf: 0,
				maxOf: 2,
				value: fullInt,
			},
			expectedTickValue: fullInt,
		},
		"Add tick at the end": {
			toAdd:    newTick(new(big.Int).Mul(fullInt, big.NewInt(2))),
			coverage: defaultCoverage_,
			isMin:    false,
			isMax:    true,
			expectedTick: &tick{
				minOf: 0,
				maxOf: 1,
				value: new(big.Int).Mul(fullInt, big.NewInt(2)),
			},
			expectedTickValue: new(big.Int).Mul(fullInt, big.NewInt(2)),
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			test.coverage.addTick(test.toAdd, test.isMin, test.isMax)
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
	coverage := &coverage{nil, 0}

	// zeroNull corresponds to tick at 0
	zeroNull := newTick(zeroInt)

	// tickQuarter corresponds to tick at 1/4 of the domain
	tickQuarter := newTick(quarterInt)

	// tickHalf corresponds to tick at 1/2 of the domain
	tickHalf := newTick(halfInt)

	// tickThreeQuarters corresponds to tick at 3/4 of the domain
	tickThreeQuarters := newTick(threeQuartersInt)

	// tickFull corresponds to tick at the end of the domain
	tickFull := newTick(fullInt)

	coverage.addTick(zeroNull, true, false)
	coverage.addTick(tickQuarter, true, true)
	coverage.addTick(tickHalf, false, true)
	coverage.addTick(tickThreeQuarters, true, false)
	coverage.addTick(tickFull, false, true)

	assertDefaultcoverage(t, coverage)

	return coverage
}

func assertDefaultcoverage(t *testing.T, c *coverage) {
	require.Equal(t, 5, c.size)

	firstTick := c.tick
	require.Nil(t, firstTick.previousInitialized)
	require.Equal(t, zeroInt, firstTick.value)
	require.Equal(t, 1, firstTick.minOf)
	require.Equal(t, 0, firstTick.maxOf)

	secondTick := firstTick.next()
	require.Equal(t, firstTick, secondTick.previousInitialized)
	require.Equal(t, firstTick.nextInitialized, secondTick)
	require.Equal(t, quarterInt, secondTick.value)
	require.Equal(t, 1, secondTick.minOf)
	require.Equal(t, 1, secondTick.maxOf)

	thirdTick := secondTick.next()
	require.Equal(t, secondTick, thirdTick.previousInitialized)
	require.Equal(t, secondTick.nextInitialized, thirdTick)
	require.Equal(t, halfInt, thirdTick.value)
	require.Equal(t, 0, thirdTick.minOf)
	require.Equal(t, 1, thirdTick.maxOf)

	fourthTick := thirdTick.next()
	require.Equal(t, thirdTick, fourthTick.previousInitialized)
	require.Equal(t, thirdTick.nextInitialized, fourthTick)
	require.Equal(t, threeQuartersInt, fourthTick.value)
	require.Equal(t, 1, fourthTick.minOf)
	require.Equal(t, 0, fourthTick.maxOf)

	fifthTick := fourthTick.next()
	require.Equal(t, fourthTick, fifthTick.previousInitialized)
	require.Nil(t, fifthTick.nextInitialized)
	require.Equal(t, fullInt, fifthTick.value)
	require.Equal(t, 0, fifthTick.minOf)
	require.Equal(t, 1, fifthTick.maxOf)
}

func tickDeepEqual(t *testing.T, expected, actual *tick) {
	require.Equal(t, expected.value.Cmp(actual.value), 0)
	require.Equal(t, expected.minOf, actual.minOf)
	require.Equal(t, expected.maxOf, actual.maxOf)
}
