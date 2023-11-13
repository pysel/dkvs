package partition

import (
	"encoding/binary"

	pbpartition "github.com/pysel/dkvs/prototypes/partition"
)

// toStoredValue converts a value with lamport timestamp to a stored value.
func toStoredValue(lamport uint64, value []byte) *pbpartition.StoredValue {
	lamportBz := make([]byte, 8)
	binary.BigEndian.PutUint64(lamportBz, lamport)
	return &pbpartition.StoredValue{
		Lamport: lamportBz,
		Value:   value,
	}
}
