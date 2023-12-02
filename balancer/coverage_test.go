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
			expected: defaultCoverage_.Tick,
		},
		"Get tick at the end": {
			value:    fullInt,
			coverage: defaultCoverage_,
			expected: defaultCoverage_.Tick.NextInitialized.NextInitialized.NextInitialized.NextInitialized,
		},
		"Get tick in the middle": {
			value:    quarterInt,
			coverage: defaultCoverage_,
			expected: defaultCoverage_.Tick.NextInitialized,
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
		isMin             bool
		isMax             bool
		expectedTick      *pbbalancer.Tick
		expectedTickValue *big.Int
	}{
		"Add tick at the beginning": {
			toAdd:    newTick(new(big.Int).SetInt64(1)),
			coverage: &coverage{nil, 0},
			isMin:    true,
			isMax:    false,
			expectedTick: &pbbalancer.Tick{
				MinOf: 1,
				MaxOf: 0,
				Value: new(big.Int).SetInt64(1).Bytes(),
			},
			expectedTickValue: new(big.Int).SetInt64(1),
		},
		"Add tick that already exists": {
			toAdd:    newTick(zeroInt),
			coverage: defaultCoverage_,
			isMin:    true,
			isMax:    false,
			expectedTick: &pbbalancer.Tick{
				MinOf: 2,
				MaxOf: 0,
				Value: zeroInt.Bytes(),
			},
			expectedTickValue: zeroInt,
		},
		"Add tick at the end and already exists": {
			toAdd:    newTick(fullInt),
			coverage: defaultCoverage_,
			isMin:    false,
			isMax:    true,
			expectedTick: &pbbalancer.Tick{
				MinOf: 0,
				MaxOf: 2,
				Value: fullInt.Bytes(),
			},
			expectedTickValue: fullInt,
		},
		"Add tick at the end": {
			toAdd:    newTick(new(big.Int).Mul(fullInt, big.NewInt(2))),
			coverage: defaultCoverage_,
			isMin:    false,
			isMax:    true,
			expectedTick: &pbbalancer.Tick{
				MinOf: 0,
				MaxOf: 1,
				Value: new(big.Int).Mul(fullInt, big.NewInt(2)).Bytes(),
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
	one64 := int64(1)
	zero64 := int64(0)

	firstTick := c.Tick
	require.Nil(t, firstTick.PreviousInitialized)
	require.Equal(t, zeroInt.Bytes(), firstTick.Value)
	require.Equal(t, one64, firstTick.MinOf)
	require.Equal(t, zero64, firstTick.MaxOf)

	secondTick := firstTick.NextInitialized
	require.Equal(t, firstTick, secondTick.PreviousInitialized)
	require.Equal(t, firstTick.NextInitialized, secondTick)
	require.Equal(t, quarterInt.Bytes(), secondTick.Value)
	require.Equal(t, one64, secondTick.MinOf)
	require.Equal(t, one64, secondTick.MaxOf)

	thirdTick := secondTick.NextInitialized
	require.Equal(t, secondTick, thirdTick.PreviousInitialized)
	require.Equal(t, secondTick.NextInitialized, thirdTick)
	require.Equal(t, halfInt.Bytes(), thirdTick.Value)
	require.Equal(t, zero64, thirdTick.MinOf)
	require.Equal(t, one64, thirdTick.MaxOf)

	fourthTick := thirdTick.NextInitialized
	require.Equal(t, thirdTick, fourthTick.PreviousInitialized)
	require.Equal(t, thirdTick.NextInitialized, fourthTick)
	require.Equal(t, threeQuartersInt.Bytes(), fourthTick.Value)
	require.Equal(t, one64, fourthTick.MinOf)
	require.Equal(t, zero64, fourthTick.MaxOf)

	fifthTick := fourthTick.NextInitialized
	require.Equal(t, fourthTick, fifthTick.PreviousInitialized)
	require.Nil(t, fifthTick.NextInitialized)
	require.Equal(t, fullInt.Bytes(), fifthTick.Value)
	require.Equal(t, zero64, fifthTick.MinOf)
	require.Equal(t, one64, fifthTick.MaxOf)
}

func tickDeepEqual(t *testing.T, expected, actual *pbbalancer.Tick) {
	require.Equal(t, bytes.Compare(expected.Value, actual.Value), 0)
	require.Equal(t, expected.MinOf, actual.MinOf)
	require.Equal(t, expected.MaxOf, actual.MaxOf)
}

func (c *coverage) getTickByValue(value *big.Int) *pbbalancer.Tick {
	curTick := c.Tick
	vbytes := value.Bytes()
	if bytes.Equal(curTick.Value, vbytes) {
		return curTick
	}

	for curTick != nil && bytes.Compare(curTick.Value, vbytes) != 1 {
		if bytes.Equal(curTick.Value, vbytes) {
			return curTick
		}
		curTick = curTick.NextInitialized
	}

	return nil
}
