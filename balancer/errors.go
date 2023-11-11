package balancer

import "errors"

var (
	ErrPartitionOverflow        = errors.New("enough partitions are already registered")
	ErrCoverageNotProperlySetUp = errors.New("coverage is not properly set up")
)
