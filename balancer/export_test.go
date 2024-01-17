package balancer

import (
	"math/big"
	"testing"

	db "github.com/pysel/dkvs/leveldb"
	"github.com/pysel/dkvs/partition"
	pbbalancer "github.com/pysel/dkvs/prototypes/balancer"
	"github.com/stretchr/testify/require"
)

var balancerName = "balancer"

func init() {

}

func (b *Balancer) GetTickByValue(value *big.Int) *pbbalancer.Tick {
	return b.coverage.getTickByValue(value)
}

func (b *Balancer) GetCoverageSize() int {
	return len(b.coverage.ticks)
}

func (b *Balancer) GetNextPartitionRange() (partition.RangeKey, *pbbalancer.Tick, *pbbalancer.Tick) {
	return b.coverage.getNextPartitionRange()
}

func (b *Balancer) GetRangeFromDigest(digest []byte) (*partition.Range, error) {
	return b.getRangeFromDigest(digest)
}

func (b *Balancer) GetRangeToPartitions() map[partition.RangeKey][]*PartitionView {
	return b.rangeToPartitions
}

// NewBalancerTest returns a new balancer instance with an independent coverage every time.
func NewBalancerTest(t *testing.T, goalReplicaRanges int) *Balancer {
	balancerName = "balancer" + t.Name()
	db, err := db.NewLevelDB(balancerName)
	require.NoError(t, err)

	b := &Balancer{
		DB:                db,
		rangeToPartitions: make(map[partition.RangeKey][]*PartitionView),
		coverage:          &coverage{},
	}

	require.NoError(t, b.setupCoverage(goalReplicaRanges))

	return b
}
