package partition

import (
	"fmt"
	"sync"

	db "github.com/pysel/dkvs/leveldb"
	"github.com/pysel/dkvs/prototypes"
	"github.com/pysel/dkvs/shared"
	"github.com/pysel/dkvs/types"
	"google.golang.org/protobuf/proto"
)

// Partition is a node that is responsible for some range of keys.
type Partition struct {
	// hashrange is a range of keys that this partition is responsible for.
	hashrange *Range

	// Database instance
	db.DB

	// read-write mutex
	rwmutex sync.RWMutex

	// set of messages that could not have been processed yet for some reason.
	backlog *types.Backlog

	// timestamp of the last message that was processed.
	timestamp uint64

	// message that this partition is currently locked in two-phase commit prepare step.
	lockedMessage proto.Message

	// event handler
	eventHandler *shared.EventHandler
}

// NewPartition creates a new partition instance.
func NewPartition(dbPath string) *Partition {
	db, err := db.NewLevelDB(dbPath)
	if err != nil {
		panic(err)
	}

	eventHandler := shared.NewEventHandler()

	return &Partition{
		eventHandler: eventHandler,
		hashrange:    nil, // balancer should set this
		DB:           db,
		timestamp:    0,
		backlog:      types.NewBacklog(),
	}
}

// ---- Database methods ----
// Keys should be sent of 32 length bytes, since SHA-2 produces 256-bit hashes, and be of big endian format.

func (p *Partition) Get(key []byte) ([]byte, error) {
	defer p.rwmutex.RUnlock()
	p.rwmutex.RLock()

	if err := p.checkKeyRange(key); err != nil {
		return nil, ErrNotThisPartitionKey
	}

	return p.DB.Get(key)
}

func (p *Partition) Set(key, value []byte) error {
	defer p.rwmutex.Unlock()
	p.rwmutex.Lock()

	if err := p.checkKeyRange(key); err != nil {
		return ErrNotThisPartitionKey
	}

	return p.DB.Set(key, value)
}

func (p *Partition) Delete(key []byte) error {
	defer p.rwmutex.Unlock()
	p.rwmutex.Lock()

	if err := p.checkKeyRange(key); err != nil {
		return ErrNotThisPartitionKey
	}

	return p.DB.Delete(key)
}

func (p *Partition) Close() error {
	return p.DB.Close()
}

func (p *Partition) SetHashrange(hashrange *Range) {
	p.hashrange = hashrange
}

// validate TS checks the timestamp of received message against local timestamp
func (p *Partition) validateTS(ts uint64) error {
	if ts <= p.timestamp {
		return ErrTimestampIsStale{CurrentTimestamp: p.timestamp, StaleTimestamp: ts}
	} else if ts > p.timestamp+1 { // timestamp is not the next one
		return ErrTimestampNotNext{CurrentTimestamp: p.timestamp, ReceivedTimestamp: ts}
	}

	return nil
}

func (p *Partition) IncrTs() {
	p.timestamp++
}

// ProcessBacklog processes messages in backlog.
func (p *Partition) ProcessBacklog(err error) error {
	if err != nil {
		fmt.Println(err)
		return nil // TODO: should return error?
	}

	var latestTimestamp uint64
	for {
		message := p.backlog.Pop(types.BID, p.timestamp)
		if message == nil {
			break
		}

		var err error
		switch m := message.(type) {
		case *prototypes.SetRequest:
			latestTimestamp = m.Lamport
			err = p.Set(m.Key, m.Value)
		case *prototypes.DeleteRequest:
			latestTimestamp = m.Lamport
			err = p.Delete(m.Key)
		default:
			fmt.Println("Unknown message type") // TODO: think of something better here.
		}

		if err != nil {
			return err
		}
	}

	if latestTimestamp != 0 { // aka: if some message was processed
		p.timestamp = latestTimestamp
	}

	return nil
}
