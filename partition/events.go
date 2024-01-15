package partition

import "fmt"

type (
	SetEvent struct {
		key  string
		data string
	}

	GetEvent struct {
		key      string
		returned string
	}

	DeleteEvent struct {
		key string
	}

	ErrorEvent struct {
		_type string
		err   error
	}

	StaleRequestEvent struct {
		currentTimestamp  uint64
		receivedTimestamp uint64
	}

	NotNextRequestEvent struct {
		currentTimestamp  uint64
		receivedTimestamp uint64
	}
)

func (e SetEvent) Severity() string {
	return "info"
}

func (e SetEvent) Message() string {
	return fmt.Sprintf("Stored a message: %s -> %s", e.key, e.data)
}

func (e GetEvent) Severity() string {
	return "info"
}

func (e GetEvent) Message() string {
	return fmt.Sprintf("Retrieved a message: %s -> %s", e.key, e.returned)
}

func (e DeleteEvent) Severity() string {
	return "info"
}

func (e DeleteEvent) Message() string {
	return fmt.Sprintf("Deleted a message: %s", e.key)
}

func (e ErrorEvent) Severity() string {
	return "error"
}

func (e ErrorEvent) Message() string {
	return fmt.Sprintf("Error during %s operation: %s", e._type, e.err.Error())
}

func (e StaleRequestEvent) Severity() string {
	return "warning"
}

func (e StaleRequestEvent) Message() string {
	return fmt.Sprintf("Received stale request. Current timestamp: %d, received timestamp: %d", e.currentTimestamp, e.receivedTimestamp)
}

func (e NotNextRequestEvent) Severity() string {
	return "warning"
}

func (e NotNextRequestEvent) Message() string {
	return fmt.Sprintf("Received request with timestamp that is not the next one. Current timestamp: %d, received timestamp: %d", e.currentTimestamp, e.receivedTimestamp)
}
