package partition

import (
	"github.com/pysel/dkvs/prototypes"
)

// toStoredValue converts a value with lamport timestamp to a stored value.
func ToStoredValue(lamport uint64, value []byte) *prototypes.StoredValue {
	return &prototypes.StoredValue{
		Lamport: lamport,
		Value:   value,
	}
}
