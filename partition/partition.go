package partition

import "github.com/pysel/dkvs/db"

type Partition struct {
	DB        db.DB
	hashRange *Range
}

func NewPartition(db db.DB, hashRange *Range) *Partition {
	if db == nil {
		// db should always be set on creation.
		panic("db is nil")
	}

	return &Partition{
		db,
		hashRange,
	}
}
