package balancer

import "errors"

var (
	ErrPartitionOverflow = errors.New("enough partitions are already registered")
)
