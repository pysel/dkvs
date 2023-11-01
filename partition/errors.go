package partition

import "errors"

var (
	ErrNotThisPartitionKey = errors.New("a key provided is not in this partition's range")
	ErrInvalidKeySize      = errors.New("key size should be 32 bytes")
	ErrNilKey              = errors.New("key is nil")
	ErrNilRequest          = errors.New("request is nil")
)
