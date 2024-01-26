package balancer

import (
	"context"
	"crypto/sha256"
	"math/big"

	db "github.com/pysel/dkvs/leveldb"
	"github.com/pysel/dkvs/partition"
	"github.com/pysel/dkvs/prototypes"
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
	"github.com/pysel/dkvs/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Balancer is a node that is responsible for registering partitions and relaying requests to appropriate ones.
type Balancer struct {
	// Database instance
	db.DB

	// A registry, which is a mapping from ranges to partitions.
	// Multiple partitions can be mapped to the same range.
	rangeToViews map[partition.RangeKey]*RangeView

	// coverage is used for tracking the tracked ranges
	coverage *coverage
}

// NewBalancer returns a new balancer instance.
func NewBalancer(goalReplicaRanges int) *Balancer {
	db, err := db.NewLevelDB("balancer-db")
	if err != nil {
		panic(err)
	}

	b := &Balancer{
		DB:           db,
		rangeToViews: make(map[partition.RangeKey]*RangeView),
		coverage:     GetCoverage(),
	}

	err = b.setupCoverage(goalReplicaRanges)
	if err != nil {
		panic(err)
	}

	return b
}

// RegisterPartition adds a partition to the balancer.
func (b *Balancer) RegisterPartition(ctx context.Context, addr string) error {
	client := partition.NewPartitionClient(addr)

	nextPartitionRangeKey, lowerTick, _ := b.coverage.getNextPartitionRange()
	partitionRange, _ := nextPartitionRangeKey.ToRange() // TODO: err

	_, err := client.SetHashrange(ctx, &prototypes.SetHashrangeRequest{Min: partitionRange.Min.Bytes(), Max: partitionRange.Max.Bytes()})
	if err != nil {
		return err
	}

	rangeView := b.rangeToViews[nextPartitionRangeKey]
	if rangeView == nil { // means that the range is not yet covered, initialize a new range view
		rangeView = NewRangeView([]*pbpartition.PartitionServiceClient{}, []string{})
		b.rangeToViews[nextPartitionRangeKey] = rangeView
	}

	rangeView.AddPartitionData(&client, addr)

	// on sucess, inrease min and max values of ticks
	b.coverage.bumpTicks(lowerTick)

	return b.saveCoverage()
}

// GetPartitionsByKey returns a range view of partitions that contain the given key.
func (b *Balancer) GetPartitionsByKey(key []byte) *RangeView {
	shaKey := sha256.Sum256(key)
	for rangeKey, rangeView := range b.rangeToViews {
		range_, _ := rangeKey.ToRange() // Todo: err
		if range_.Contains(shaKey[:]) {
			return rangeView
		}
	}

	return nil
}

// Get returns the most up to date value between responsible replicas for a given key.
func (b *Balancer) Get(ctx context.Context, key []byte) (*prototypes.GetResponse, error) {
	shaKey := types.ShaKey(key)
	range_, err := b.getRangeFromDigest(shaKey[:])
	if err != nil {
		return nil, err
	}

	rangeView := b.rangeToViews[range_.AsKey()]
	if len(rangeView.clients) == 0 {
		return nil, ErrRangeNotYetCovered
	}

	var response *prototypes.GetResponse
	maxLamport := uint64(0)
	offlineAddressesErr := ErrPartitionsOffline{Addresses: []string{}, Errors: []error{}}

	rangeView.lamport++ // increase lamport timestamp so that we account for get request we are sending here
	requestLamport := rangeView.lamport
	for i, client := range rangeView.clients {
		resp, err := (*client).Get(ctx, &prototypes.GetRequest{Key: key, Lamport: requestLamport})
		if err != nil {
			// remove the partition if it is offline
			if s, ok := status.FromError(err); ok {
				if s.Code() == codes.Unavailable {
					offlineAddressesErr.Addresses = append(offlineAddressesErr.Addresses, rangeView.addresses[i])
					offlineAddressesErr.Errors = append(offlineAddressesErr.Errors, err)

					// TODO: consider tombstoning before removing
					rangeView.removePartition(rangeView.addresses[i])
				}
			}
			continue
		} else if resp.StoredValue == nil {
			response = resp
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

	return response, offlineAddressesErr.ErrOrNil()
}

// setupCoverage creates necessary ticks for coverage based on goalReplicaRanges
func (b *Balancer) setupCoverage(goalReplicaRanges int) error {
	if goalReplicaRanges == 0 {
		b.coverage.addTick(newTick(big.NewInt(0), 0))
		b.coverage.addTick(newTick(partition.MaxInt, 0))
		return nil
	}

	// Create a tick for each partition
	for i := 0; i <= goalReplicaRanges; i++ {
		numerator := new(big.Int).Mul(big.NewInt(int64(i)), partition.MaxInt)
		value := new(big.Int).Div(numerator, big.NewInt(int64(goalReplicaRanges)))
		b.coverage.addTick(newTick(value, 0))
	}

	return b.saveCoverage()
}

// getRangeFromDigest returns a range to which the given digest belongs
func (b *Balancer) getRangeFromDigest(digest []byte) (*partition.Range, error) {
	for rangeKey := range b.rangeToViews {
		range_, _ := rangeKey.ToRange() // TODO: err
		if range_.Contains(digest) {
			return range_, nil
		}
	}

	return nil, ErrDigestNotCovered
}

// saveCoverage saves the current coverage to disk
func (b *Balancer) saveCoverage() error {
	coverageBz, err := proto.Marshal(b.coverage.ToProto())
	if err != nil {
		return err
	}

	return b.DB.Set(CoverageKey, coverageBz)
}

// GetNextLamportForKey returns the next lamport timestamp for a given key based on the digest of the key.
func (b *Balancer) GetNextLamportForKey(key []byte) uint64 {
	shaKey := types.ShaKey(key)
	range_, err := b.getRangeFromDigest(shaKey[:])
	if err != nil {
		return 0
	}

	return b.rangeToViews[range_.AsKey()].lamport + 1
}
