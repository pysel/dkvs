package partition

import (
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
)

func ToStoredValue(lamport uint64, value []byte) *pbpartition.StoredValue {
	return toStoredValue(lamport, value)
}
