package partition

import (
	"errors"
	"strconv"

	"github.com/pysel/dkvs/shared"
)

var (
	ErrNotThisPartitionKey = errors.New("a key provided is not in this partition's range")
	ErrInvalidKeySize      = errors.New("key size should be 32 bytes")
	ErrUnsupported2PCMsg   = errors.New("unsupported 2PC message")
	ErrNoLockedMessage     = errors.New("no locked message")
)

type (
	// ErrTimestampNotNext is returned when a timestamp of a received request is not the current timestamp + 1
	ErrTimestampNotNext struct {
		CurrentTimestamp  uint64
		ReceivedTimestamp uint64
	}

	// ErrInternal is returned when an internal error occurs during attempt to serve a request
	ErrInternal struct {
		Reason error
	}

	ErrTimestampIsStale struct {
		CurrentTimestamp uint64
		StaleTimestamp   uint64
	}
)

func (e ErrTimestampNotNext) Error() string {
	return "timestamp is not the next one, current timestamp: " + strconv.Itoa(int(e.CurrentTimestamp))
}

func (e ErrTimestampNotNext) ToEvent() shared.Event {
	return NotNextRequestEvent{
		currentTimestamp:  e.CurrentTimestamp,
		receivedTimestamp: e.ReceivedTimestamp,
	}
}

func (e ErrInternal) Error() string {
	return "internal error: " + e.Reason.Error()
}

func (e ErrInternal) Unwrap() error {
	return e.Reason
}

func (e ErrTimestampIsStale) Error() string {
	return "timestamp is stale, current timestamp: " + strconv.Itoa(int(e.CurrentTimestamp)) + ", received timestamp: " + strconv.Itoa(int(e.StaleTimestamp))
}

func (e ErrTimestampIsStale) ToEvent() shared.Event {
	return StaleRequestEvent{
		currentTimestamp:  e.CurrentTimestamp,
		receivedTimestamp: e.StaleTimestamp,
	}
}
