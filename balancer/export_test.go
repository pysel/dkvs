package balancer

import (
	"math/big"
	"testing"

	db "github.com/pysel/dkvs/leveldb"
	"github.com/pysel/dkvs/partition"
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
	"github.com/stretchr/testify/require"
)

var balancerName = "balancer"

func init() {

}

func (b *Balancer) GetTickByValue(value *big.Int) *tick {
	return b.coverage.getTickByValue(value)
}

func (b *Balancer) GetTicksAmount() int {
	return b.coverage.size
}

func (b *Balancer) GetNextPartitionRange() (*partition.Range, *tick, *tick) {
	return b.coverage.getNextPartitionRange()
}

func (b *Balancer) GetRangeFromDigest(digest []byte) (*partition.Range, error) {
	return b.getRangeFromDigest(digest)
}

// NewBalancerTest returns a new balancer instance with an independent coverage every time.
func NewBalancerTest(t *testing.T, goalReplicaRanges int) *Balancer {
	balancerName = "balancer" + t.Name()
	db, err := db.NewLevelDB(balancerName)
	require.NoError(t, err)

	b := &Balancer{
		DB:                db,
		clients:           make(map[*partition.Range][]pbpartition.PartitionServiceClient),
		goalReplicaRanges: goalReplicaRanges,
		activePartitions:  0,
		coverage:          &coverage{},
	}

	b.setupCoverage()

	return b
}
