package partition

import "fmt"

type (
	SetHashrangeEvent struct {
		min uint64
		max uint64
	}

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
	return fmt.Sprintf("Stored a message: \033[32m%s\033[0m -> \033[32m%s\033[0m", e.key, e.data)
}

func (e GetEvent) Severity() string {
	return "info"
}

func (e GetEvent) Message() string {
	return fmt.Sprintf("Retrieved a message: \033[32m%s\033[0m -> \033[32m%s\033[0m", e.key, e.returned)
}

func (e DeleteEvent) Severity() string {
	return "info"
}

func (e DeleteEvent) Message() string {
	return fmt.Sprintf("Deleted a message: \033[32m%s\033[0m", e.key)
}

func (e ErrorEvent) Severity() string {
	return "error"
}

func (e ErrorEvent) Message() string {
	return fmt.Sprintf("Error during %s operation: \033[31m%s\033[0m", e._type, e.err.Error())
}

func (e StaleRequestEvent) Severity() string {
	return "warning"
}

func (e StaleRequestEvent) Message() string {
	return fmt.Sprintf("\033[33mStale Request\033[0m. Current timestamp: \033[32m%d\033[0m, received timestamp: \033[32m%d\033[0m", e.currentTimestamp, e.receivedTimestamp)
}

func (e NotNextRequestEvent) Severity() string {
	return "warning"
}

func (e NotNextRequestEvent) Message() string {
	return fmt.Sprintf("\033[33mFuture Request\033[0m. Current timestamp: \033[32m%d\033[0m, received timestamp: \033[32m%d\033[0m", e.currentTimestamp, e.receivedTimestamp)
}

func (e SetHashrangeEvent) Severity() string {
	return "info"
}

func (e SetHashrangeEvent) Message() string {
	return fmt.Sprintf("Set hashrange: \033[32m%d\033[0m -> \033[32m%d\033[0m", e.min, e.max)
}
