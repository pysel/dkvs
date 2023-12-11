package partition

import "errors"

var (
	ErrNotThisPartitionKey      = errors.New("a key provided is not in this partition's range")
	ErrInvalidKeySize           = errors.New("key size should be 32 bytes")
	ErrUnsupported2PCMsg        = errors.New("unsupported 2PC message")
	ErrNoLockedMessage          = errors.New("no locked message")
	ErrTimestampLessThanCurrent = errors.New("timestamp is less than current timestamp")
	ErrTimestampNotNext         = errors.New("timestamp is not the next one")
	ErrNotBalancerID            = errors.New("not a balancer id")
)
