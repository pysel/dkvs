package balancer

import "errors"

var (
	ErrPartitionOverflow        = errors.New("enough partitions are already registered")
	ErrCoverageNotProperlySetUp = errors.New("coverage is not properly set up")
	ErrDigestNotCovered         = errors.New("digest is not covered by any range")
	ErrRangeNotYetCovered       = errors.New("range is not yet covered by any partition")
	ErrAllReplicasFailed        = errors.New("all replicas failed to process request")
)

// GRPC errors
var (
	ErrNilRequest = errors.New("request is nil")
)
