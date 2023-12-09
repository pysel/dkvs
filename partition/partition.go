package partition

import (
	db "github.com/pysel/dkvs/leveldb"
	"github.com/pysel/dkvs/prototypes"
	"google.golang.org/protobuf/proto"
)

// Partition is a node that is responsible for some range of keys.
type Partition struct {
	hashrange *Range
	db.DB

	// isLocked indicates whether the partition is locked.
	isLocked bool
	// set of messages that could not have been processed yet for some reason.
	backlog *Backlog

	// message that this partition is currently locked in in two-phase commit prepare step.
	lockedMessage proto.Message
}

// NewPartition creates a new partition instance.
func NewPartition(dbPath string) *Partition {
	db, err := db.NewLevelDB(dbPath)
	if err != nil {
		panic(err)
	}

	return &Partition{
		hashrange: nil, // balancer should set this
		DB:        db,
	}
}

// ---- Database methods ----
// Keys should be sent of 32 length bytes, since SHA-2 produces 256-bit hashes, and be of big endian format.

func (p *Partition) Get(key []byte) ([]byte, error) {
	if err := p.checkKeyRange(key); err != nil {
		return nil, ErrNotThisPartitionKey
	}

	return p.DB.Get(key)
}

func (p *Partition) Set(key, value []byte) error {
	if err := p.checkKeyRange(key); err != nil {
		return ErrNotThisPartitionKey
	}

	return p.DB.Set(key, value)
}

func (p *Partition) Delete(key []byte) error {
	if err := p.checkKeyRange(key); err != nil {
		return ErrNotThisPartitionKey
	}

	return p.DB.Delete(key)
}

func (p *Partition) Has(key []byte) (bool, error) {
	if err := p.checkKeyRange(key); err != nil {
		// still error to be able to distinguish between not found and out of range
		return false, ErrNotThisPartitionKey
	}

	return p.DB.Has(key), nil
}

func (p *Partition) Close() error {
	return p.DB.Close()
}

func (p *Partition) SetHashrange(hashrange *Range) {
	p.hashrange = hashrange
}

func (p *Partition) ProcessBacklog() {
	p.isLocked = true
	defer func() {
		p.isLocked = false
	}()

	for {
		message := p.backlog.Pop()
		if message == nil {
			return
		}

		switch m := message.(type) {
		case *prototypes.SetRequest:
			p.Set(m.Key, m.Value)
		case *prototypes.DeleteRequest:
			p.Delete(m.Key)
		}
	}
}
