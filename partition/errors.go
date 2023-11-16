package partition

import "errors"

var (
	ErrNotThisPartitionKey = errors.New("a key provided is not in this partition's range")
	ErrInvalidKeySize      = errors.New("key size should be 32 bytes")
	ErrUnsupported2PCMsg   = errors.New("unsupported 2PC message")
	ErrNoLockedMessage     = errors.New("no locked message")
)
