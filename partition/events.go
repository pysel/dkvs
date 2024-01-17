package partition

import (
	"fmt"
	"math/big"
)

type (
	SetHashrangeEvent struct {
		min *big.Int
		max *big.Int
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
		req string
		err error
	}

	StaleRequestEvent struct {
		req               string
		currentTimestamp  uint64
		receivedTimestamp uint64
	}

	NotNextRequestEvent struct {
		req               string
		currentTimestamp  uint64
		receivedTimestamp uint64
	}

	ServerStartEvent struct {
		port uint64
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
	return fmt.Sprintf("Error for \033[90m{%s}\033[0m request: \033[31m%s\033[0m", e.req, e.err.Error())
}

func (e StaleRequestEvent) Severity() string {
	return "warning"
}

func (e StaleRequestEvent) Message() string {
	return fmt.Sprintf("\033[33mStale Request\033[0m. Request: \033[90m{%s}\033[0m. Current timestamp: \033[32m%d\033[0m, received timestamp: \033[32m%d\033[0m", e.req, e.currentTimestamp, e.receivedTimestamp)
}

func (e NotNextRequestEvent) Severity() string {
	return "warning"
}

func (e NotNextRequestEvent) Message() string {
	return fmt.Sprintf("\033[33mFuture Request\033[0m. Request: \033[90m{%s}\033[0m. Current timestamp: \033[32m%d\033[0m, received timestamp: \033[32m%d\033[0m", e.req, e.currentTimestamp, e.receivedTimestamp)
}

func (e SetHashrangeEvent) Severity() string {
	return "info"
}

func (e SetHashrangeEvent) Message() string {
	return fmt.Sprintf("Set hashrange: \033[32m%s\033[0m -> \033[32m%s\033[0m", e.min.String(), e.max.String())
}

func (e ServerStartEvent) Severity() string {
	return "info"
}

func (e ServerStartEvent) Message() string {
	return fmt.Sprintf("Server started on port: \033[32m%d\033[0m", e.port)
}
