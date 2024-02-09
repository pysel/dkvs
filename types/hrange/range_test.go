package hrange_test

import (
	"crypto/sha256"
	"math/big"
	"testing"

	"github.com/pysel/dkvs/testutil"
	"github.com/pysel/dkvs/types/hrange"
	"github.com/stretchr/testify/require"
)

func TestNewRange(t *testing.T) {
	tooBig := new(big.Int).Exp(big.NewInt(2), big.NewInt(257), nil)
	tests := []struct {
		min            *big.Int
		max            *big.Int
		expectedRange  *hrange.Range
		expectingPanic bool
	}{
		{
			min:            big.NewInt(-1),
			max:            big.NewInt(0),
			expectedRange:  &hrange.Range{},
			expectingPanic: true,
		},
		{
			min:            big.NewInt(0),
			max:            big.NewInt(0),
			expectedRange:  &hrange.Range{},
			expectingPanic: true,
		},
		{
			min:            big.NewInt(0),
			max:            big.NewInt(1),
			expectedRange:  &hrange.Range{big.NewInt(0), big.NewInt(1)},
			expectingPanic: false,
		},
		{
			min:            big.NewInt(1),
			max:            big.NewInt(0),
			expectedRange:  &hrange.Range{},
			expectingPanic: true,
		},
		{
			min:            big.NewInt(0),
			max:            tooBig,
			expectedRange:  &hrange.Range{},
			expectingPanic: true,
		},
		{
			min:            big.NewInt(0),
			max:            big.NewInt(500),
			expectedRange:  &hrange.Range{big.NewInt(0), big.NewInt(500)},
			expectingPanic: false,
		},
		{
			min:            big.NewInt(500),
			max:            hrange.MaxInt,
			expectedRange:  &hrange.Range{big.NewInt(500), hrange.MaxInt},
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

		got := hrange.NewRange(test.min.Bytes(), test.max.Bytes())
		require.Equal(t, test.expectedRange, got)
	}
}

func TestContains(t *testing.T) {
	nonDomainHash := sha256.Sum256(testutil.NonDomainKey)
	domainHash := sha256.Sum256(testutil.DomainKey)

	tests := []struct {
		name     string
		r        *hrange.Range
		hash     []byte
		expected bool
	}{
		{
			name:     "key is in range",
			r:        testutil.DefaultHashrange,
			hash:     domainHash[:],
			expected: true,
		},
		{
			name:     "key is not in range",
			r:        testutil.DefaultHashrange,
			hash:     nonDomainHash[:],
			expected: false,
		},
		{
			name:     "key is min",
			r:        testutil.DefaultHashrange,
			hash:     testutil.DefaultHashrange.Min.Bytes(),
			expected: true,
		},
		{
			name:     "key is max",
			r:        testutil.DefaultHashrange,
			hash:     testutil.DefaultHashrange.Max.Bytes(),
			expected: true,
		},
		{
			name:     "key is max + 1",
			r:        testutil.DefaultHashrange,
			hash:     new(big.Int).Add(testutil.DefaultHashrange.Max, big.NewInt(1)).Bytes(),
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.r.Contains(test.hash)
			require.Equal(t, test.expected, got)
		})
	}
}

func TestAsString(t *testing.T) {
	tests := []struct {
		name     string
		r        *hrange.Range
		expected hrange.RangeKey
	}{
		{
			name:     "range is default",
			r:        testutil.DefaultHashrange,
			expected: hrange.RangeKey("0" + "; " + testutil.HalfShaDomain.String()),
		},
		{
			name:     "range is full",
			r:        testutil.FullHashrange,
			expected: hrange.RangeKey("0" + "; " + hrange.MaxInt.String()),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.r.AsKey()
			require.Equal(t, test.expected, got)
		})
	}
}

func TestToRange(t *testing.T) {
	tests := []struct {
		name     string
		r        hrange.RangeKey
		expected *hrange.Range
	}{
		{
			name:     "range is default",
			r:        hrange.RangeKey("0" + "; " + testutil.HalfShaDomain.String()),
			expected: testutil.DefaultHashrange,
		},
		{
			name:     "range is full",
			r:        hrange.RangeKey("0" + "; " + hrange.MaxInt.String()),
			expected: testutil.FullHashrange,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := test.r.ToRange()
			require.NoError(t, err)
			require.Equal(t, test.expected, got)
		})
	}
}
