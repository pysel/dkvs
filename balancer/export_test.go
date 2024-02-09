package balancer

import (
	"math/big"
	"testing"

	coverage "github.com/pysel/dkvs/balancer/coverage"
	"github.com/pysel/dkvs/balancer/rangeview"
	leveldb "github.com/pysel/dkvs/db/leveldb"
	pbbalancer "github.com/pysel/dkvs/prototypes/balancer"
	"github.com/pysel/dkvs/types/hashrange"
	"github.com/stretchr/testify/require"
)

var balancerName = "balancer"

type (
	ClientIdToLamport clientIdToLamport
)

func (b *Balancer) GetTickByValue(value *big.Int) *pbbalancer.Tick {
	return b.coverage.GetTickByValue(value)
}

func (b *Balancer) GetCoverageSize() int {
	return len(b.coverage.Ticks)
}

func (b *Balancer) GetNextPartitionRange() (hashrange.RangeKey, *pbbalancer.Tick, *pbbalancer.Tick) {
	return b.coverage.GetNextPartitionRange()
}

func (b *Balancer) GetRangeFromDigest(digest []byte) (*hashrange.Range, error) {
	return b.getRangeFromDigest(digest)
}

func (b *Balancer) GetRangeToViews() map[hashrange.RangeKey]*rangeview.RangeView {
	return b.rangeToViews
}

// NewBalancerTest returns a new balancer instance with an independent Coverage every time.
func NewBalancerTest(t *testing.T, goalReplicaRanges int) *Balancer {
	balancerName = "balancer" + t.Name()
	db, err := leveldb.NewLevelDB(balancerName)

	require.NoError(t, err)

	b := &Balancer{
		DB:                db,
		rangeToViews:      make(map[hashrange.RangeKey]*rangeview.RangeView),
		coverage:          &coverage.Coverage{},
		clientIdToLamport: NewClientIdToLamport(),
	}

	require.NoError(t, b.setupCoverage(goalReplicaRanges))
	return b
}
