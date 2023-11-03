package partition

import (
	"math/big"
)

// A range of keys this partition is responsible for. Total range is [0; 2^256].
type Range struct {
	Min *big.Int
	Max *big.Int
}

var (
	MinInt *big.Int
	MaxInt *big.Int
)

func init() {
	MinInt = new(big.Int).SetInt64(0)
	MaxInt_bz := make([]byte, 32)
	for i := 0; i < 32; i++ {
		MaxInt_bz[i] = 0xFF // a byte with all bits set to 1
	}

	MaxInt = new(big.Int).SetBytes(MaxInt_bz)
}

// NewRange is a constructor for Range.
func NewRange(min, max *big.Int) *Range {
	if min.Cmp(MinInt) == -1 {
		// min should be >= 0, since SHA-2 only produces positive hashes.
		panic("min is negative")
	}

	if max.Cmp(MaxInt) == 1 {
		// max should be lower than maximum possible hash.
		panic("max is greater than 2^256")
	}

	if max.Cmp(min) == 0 {
		// min and max should be different.
		panic("min and max are equal")
	}

	if max.Cmp(min) == -1 {
		// max should be greater than min.
		panic("max is less than min")
	}

	return &Range{min, max}
}

func (r *Range) Contains(key []byte) bool {
	keyInt := new(big.Int).SetBytes(key)

	return r.Min.Cmp(keyInt) <= 0 && r.Max.Cmp(keyInt) >= 0
}
