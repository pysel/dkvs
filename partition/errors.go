package partition

import (
	"errors"
	"strconv"
)

var (
	ErrNotThisPartitionKey      = errors.New("a key provided is not in this partition's range")
	ErrInvalidKeySize           = errors.New("key size should be 32 bytes")
	ErrUnsupported2PCMsg        = errors.New("unsupported 2PC message")
	ErrNoLockedMessage          = errors.New("no locked message")
	ErrTimestampLessThanCurrent = errors.New("timestamp is less than current timestamp")
)

type (
	ErrTimestampNotNext struct {
		CurrentTimestamp uint64
	}
)

func (e ErrTimestampNotNext) Error() string {
	return "timestamp is not the next one, current timestamp: " + strconv.Itoa(int(e.CurrentTimestamp))
}
