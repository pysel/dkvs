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

type (
	ClientIdToLamport clientIdToLamport
)

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

func (b *Balancer) GetrangeToViews() map[partition.RangeKey]*RangeView {
	return b.rangeToViews
}

func (rv *RangeView) RemovePartition(addr string) error {
	return rv.removePartition(addr)
}

func (rv *RangeView) GetAddresses() []string {
	return rv.addresses
}

// NewBalancerTest returns a new balancer instance with an independent coverage every time.
func NewBalancerTest(t *testing.T, goalReplicaRanges int) *Balancer {
	balancerName = "balancer" + t.Name()
	db, err := db.NewLevelDB(balancerName)
	require.NoError(t, err)

	b := &Balancer{
		DB:                db,
		rangeToViews:      make(map[partition.RangeKey]*RangeView),
		coverage:          &coverage{},
		clientIdToLamport: make(clientIdToLamport),
	}

	require.NoError(t, b.setupCoverage(goalReplicaRanges))
	return b
}
