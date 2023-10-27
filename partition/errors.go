package partition

import "errors"

var (
	ErrNotThisPartitionKey = errors.New("a key provided is not in this partition's range")
)
