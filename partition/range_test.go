package partition_test

import (
	"crypto/sha256"
	"math/big"
	"testing"

	"github.com/pysel/dkvs/partition"
	"github.com/pysel/dkvs/testutil"
	"github.com/stretchr/testify/require"
)

func TestNewRange(t *testing.T) {
	tooBig := new(big.Int).Exp(big.NewInt(2), big.NewInt(257), nil)
	tests := []struct {
		min            *big.Int
		max            *big.Int
		expectedRange  *partition.Range
		expectingPanic bool
	}{
		{
			min:            big.NewInt(-1),
			max:            big.NewInt(0),
			expectedRange:  &partition.Range{},
			expectingPanic: true,
		},
		{
			min:            big.NewInt(0),
			max:            big.NewInt(0),
			expectedRange:  &partition.Range{},
			expectingPanic: true,
		},
		{
			min:            big.NewInt(0),
			max:            big.NewInt(1),
			expectedRange:  &partition.Range{big.NewInt(0), big.NewInt(1)},
			expectingPanic: false,
		},
		{
			min:            big.NewInt(1),
			max:            big.NewInt(0),
			expectedRange:  &partition.Range{},
			expectingPanic: true,
		},
		{
			min:            big.NewInt(0),
			max:            tooBig,
			expectedRange:  &partition.Range{},
			expectingPanic: true,
		},
		{
			min:            big.NewInt(0),
			max:            big.NewInt(500),
			expectedRange:  &partition.Range{big.NewInt(0), big.NewInt(500)},
			expectingPanic: false,
		},
		{
			min:            big.NewInt(500),
			max:            partition.MaxInt,
			expectedRange:  &partition.Range{big.NewInt(500), partition.MaxInt},
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

		got := partition.NewRange(test.min.Bytes(), test.max.Bytes())
		require.Equal(t, test.expectedRange, got)
	}
}

func TestContains(t *testing.T) {
	nonDomainHash := sha256.Sum256(testutil.NonDomainKey)
	domainHash := sha256.Sum256(testutil.DomainKey)

	tests := []struct {
		name     string
		r        *partition.Range
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
		r        *partition.Range
		expected partition.RangeKey
	}{
		{
			name:     "range is default",
			r:        testutil.DefaultHashrange,
			expected: partition.RangeKey("0" + "; " + testutil.HalfShaDomain.String()),
		},
		{
			name:     "range is full",
			r:        testutil.FullHashrange,
			expected: partition.RangeKey("0" + "; " + partition.MaxInt.String()),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.r.AsString()
			require.Equal(t, test.expected, got)
		})
	}
}

func TestToRange(t *testing.T) {
	tests := []struct {
		name     string
		r        partition.RangeKey
		expected *partition.Range
	}{
		{
			name:     "range is default",
			r:        partition.RangeKey("0" + "; " + testutil.HalfShaDomain.String()),
			expected: testutil.DefaultHashrange,
		},
		{
			name:     "range is full",
			r:        partition.RangeKey("0" + "; " + partition.MaxInt.String()),
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
