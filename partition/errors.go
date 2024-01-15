package partition

import (
	"errors"
	"strconv"
)

var (
	ErrNotThisPartitionKey = errors.New("a key provided is not in this partition's range")
	ErrInvalidKeySize      = errors.New("key size should be 32 bytes")
	ErrUnsupported2PCMsg   = errors.New("unsupported 2PC message")
	ErrNoLockedMessage     = errors.New("no locked message")
	ErrTimestampIsStale    = errors.New("timestamp is less than current timestamp")
)

type (
	// ErrTimestampNotNext is returned when a timestamp of a received request is not the current timestamp + 1
	ErrTimestampNotNext struct {
		CurrentTimestamp uint64
	}

	// ErrInternal is returned when an internal error occurs during attempt to serve a request
	ErrInternal struct {
		Reason error
	}
)

func (e ErrTimestampNotNext) Error() string {
	return "timestamp is not the next one, current timestamp: " + strconv.Itoa(int(e.CurrentTimestamp))
}

func (e ErrInternal) Error() string {
	return "internal error: " + e.Reason.Error()
}

func (e ErrInternal) Unwrap() error {
	return e.Reason
}
