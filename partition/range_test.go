package partition

import (
	"math/big"
	"testing"
)

func TestNewRange(t *testing.T) {
	tooBig := new(big.Int).Exp(big.NewInt(2), big.NewInt(257), nil)
	tests := []struct {
		min            *big.Int
		max            *big.Int
		expectedRange  *Range
		expectingPanic bool
	}{
		{
			min:            big.NewInt(-1),
			max:            big.NewInt(0),
			expectedRange:  &Range{},
			expectingPanic: true,
		},
		{
			min:            big.NewInt(0),
			max:            big.NewInt(0),
			expectedRange:  &Range{},
			expectingPanic: true,
		},
		{
			min:            big.NewInt(0),
			max:            big.NewInt(1),
			expectedRange:  &Range{big.NewInt(0), big.NewInt(1)},
			expectingPanic: false,
		},
		{
			min:            big.NewInt(1),
			max:            big.NewInt(0),
			expectedRange:  &Range{},
			expectingPanic: true,
		},
		{
			min:            big.NewInt(0),
			max:            tooBig,
			expectedRange:  &Range{},
			expectingPanic: true,
		},
		{
			min:            big.NewInt(0),
			max:            big.NewInt(500),
			expectedRange:  &Range{big.NewInt(0), big.NewInt(500)},
			expectingPanic: false,
		},
		{
			min:            big.NewInt(500),
			max:            maxInt,
			expectedRange:  &Range{big.NewInt(500), maxInt},
			expectingPanic: false,
		},
	}

	for _, test := range tests {
		if test.expectingPanic {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("NewRange(%v, %v) should panic", test.min, test.max)
				}
			}()
		}

		NewRange(test.min, test.max)
	}
}
