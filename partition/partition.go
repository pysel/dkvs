package partition

import (
	"github.com/pysel/dkvs/db"
)

// Partition is a slave node that stores some range of keys
type Partition struct {
	DB        db.DB
	hashRange *Range
}

// NewPartition creates a new partition instance.
func NewPartition(dbPath string, hashRange *Range) *Partition {
	db, err := db.NewLevelDB(dbPath)
	if err != nil {
		panic(err)
	}

	return &Partition{
		db,
		hashRange,
	}
}

// ---- Database methods ----
// Keys should be sent of 32 length bytes, since SHA-2 produces 256-bit hashes, and be of big endian format.

func (p *Partition) Get(key []byte) ([]byte, error) {
	if err := checkKeyRange(key); err != nil {
		return nil, ErrNotThisPartitionKey
	}

	return p.DB.Get(key)
}
