package balancer

import (
	"context"
	"crypto/sha256"
	"math/big"

	"github.com/pysel/dkvs/partition"
	"github.com/pysel/dkvs/prototypes"
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
	"github.com/pysel/dkvs/types"
)

// Balancer is a node that is responsible for registering partitions and relaying requests to appropriate ones.
type Balancer struct {
	// A mapping from ranges to partitions.
	// Multiple partitions can be mapped to the same range.
	clients map[*partition.Range][]pbpartition.PartitionServiceClient

	// goalReplicaRanges is the number of different sets of replicas that should be created
	goalReplicaRanges int
	// activePartitions is the number of currently registered partitions
	activePartitions int

	// coverage is used for tracking the tracked ranges
	coverage *coverage
}

// NewBalancer returns a new balancer instance.
func NewBalancer(goalReplicaRanges int) *Balancer {
	b := &Balancer{
		clients:           make(map[*partition.Range][]pbpartition.PartitionServiceClient),
		goalReplicaRanges: goalReplicaRanges,
		activePartitions:  0,
		coverage:          GetCoverage(),
	}

	b.setupCoverage()

	return b
}

// AddPartition adds a partition to the balancer.
func (b *Balancer) RegisterPartition(ctx context.Context, addr string) error {
	client := partition.NewPartitionClient(addr)

	partitionRange, lowerTick, upperTick := b.coverage.getNextPartitionRange()
	_, err := client.SetHashrange(ctx, &prototypes.SetHashrangeRequest{Min: partitionRange.Min.Bytes(), Max: partitionRange.Max.Bytes()})
	if err != nil {
		return err
	}

	b.clients[partitionRange] = append(b.clients[partitionRange], client)
	b.activePartitions++

	// on sucess, inrease min and max values of ticks
	b.coverage.bumpTicks(lowerTick, upperTick)

	return nil
}

// GetPartitions returns a list of partitions that contain the given key.
func (b *Balancer) GetPartitionsByKey(key []byte) []pbpartition.PartitionServiceClient {
	shaKey := sha256.Sum256(key)
	for range_, clients := range b.clients {
		if range_.Contains(shaKey[:]) {
			return clients
		}
	}

	return nil
}

func (b *Balancer) Get(ctx context.Context, key string) (*prototypes.GetResponse, error) {
	shaKey := types.ShaKey(key)
	range_, err := b.getRangeFromDigest(shaKey[:])
	if err != nil {
		return nil, err
	}

	responsibleClients := b.clients[range_]
	if len(responsibleClients) == 0 {
		return nil, ErrRangeNotYetCovered
	}

	var response *prototypes.GetResponse
	maxLamport := uint64(0)
	for _, client := range responsibleClients {
		resp, err := client.Get(ctx, &prototypes.GetRequest{Key: key})
		if err != nil {
			continue
		}

		// since returned value will be a tuple of lamport timestamp and value, check which returned value
		// has the highest lamport timestamp

		if resp.StoredValue.Lamport >= maxLamport {
			maxLamport = resp.StoredValue.Lamport
			response = &prototypes.GetResponse{StoredValue: resp.StoredValue}
		}
	}

	if response == nil {
		return nil, ErrAllReplicasFailed
	}

	return response, nil
}

// setupCoverage creates necessary ticks for coverage based on goalReplicaRanges
func (b *Balancer) setupCoverage() {
	if b.goalReplicaRanges == 0 {
		b.coverage.addTick(newTick(big.NewInt(0)), false, false)
		b.coverage.addTick(newTick(partition.MaxInt), false, false)
		return
	}

	// Create a tick for each partition
	for i := 0; i <= b.goalReplicaRanges; i++ {
		numerator := new(big.Int).Mul(big.NewInt(int64(i)), partition.MaxInt)
		value := new(big.Int).Div(numerator, big.NewInt(int64(b.goalReplicaRanges)))
		b.coverage.addTick(newTick(value), false, false)
	}
}

// getRangeFromDigest returns a range to which the given digest belongs
func (b *Balancer) getRangeFromDigest(digest []byte) (*partition.Range, error) {
	for range_ := range b.clients {
		if range_.Contains(digest) {
			return range_, nil
		}
	}

	return nil, ErrDigestNotCovered
}
