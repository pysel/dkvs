package partition

import (
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
)

// toStoredValue converts a value with lamport timestamp to a stored value.
func toStoredValue(lamport uint64, value []byte) *pbpartition.StoredValue {
	return &pbpartition.StoredValue{
		Lamport: lamport,
		Value:   value,
	}
}
