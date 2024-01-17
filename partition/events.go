package partition

import (
	"fmt"
	"math/big"

	"github.com/pysel/dkvs/shared"
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

	TwoPCPrepareCommitEvent struct {
		msg string
	}

	TwoPCAbortEvent struct{}
)

func (e SetEvent) Severity() string {
	return "info"
}

func (e SetEvent) Message() string {
	return fmt.Sprintf("Stored a message: %s -> %s", shared.GreenWrap(e.key), shared.GreenWrap(e.data))
}

func (e GetEvent) Severity() string {
	return "info"
}

func (e GetEvent) Message() string {
	return fmt.Sprintf("Retrieved a message: %s -> %s", shared.GreenWrap(e.key), shared.GreenWrap(e.returned))
}

func (e DeleteEvent) Severity() string {
	return "info"
}

func (e DeleteEvent) Message() string {
	return fmt.Sprintf("Deleted a message: %s", shared.GreenWrap(e.key))
}

func (e ErrorEvent) Severity() string {
	return "error"
}

func (e ErrorEvent) Message() string {
	return fmt.Sprintf("Error for {%s} request: %s", shared.GreyWrap(e.req), shared.RedWrap(e.err.Error()))
}

func (e StaleRequestEvent) Severity() string {
	return "warning"
}

func (e StaleRequestEvent) Message() string {
	return fmt.Sprintf("\033[33mStale Request\033[0m. Request: {%s}. Current timestamp: \033[32m%d\033[0m, received timestamp: \033[32m%d\033[0m", shared.GreyWrap(e.req), e.currentTimestamp, e.receivedTimestamp)
}

func (e NotNextRequestEvent) Severity() string {
	return "warning"
}

func (e NotNextRequestEvent) Message() string {
	return fmt.Sprintf("\033[33mFuture Request\033[0m. Request: {%s}. Current timestamp: \033[32m%d\033[0m, received timestamp: \033[32m%d\033[0m", shared.GreyWrap(e.req), e.currentTimestamp, e.receivedTimestamp)
}

func (e SetHashrangeEvent) Severity() string {
	return "info"
}

func (e SetHashrangeEvent) Message() string {
	return fmt.Sprintf("Set hashrange: %s -> %s", shared.GreenWrap(e.min.String()), shared.GreenWrap(e.max.String()))
}

func (e ServerStartEvent) Severity() string {
	return "info"
}

func (e ServerStartEvent) Message() string {
	return fmt.Sprintf("Server started on port: \033[32m%d\033[0m", e.port)
}

func (e TwoPCPrepareCommitEvent) Severity() string {
	return "info"
}

func (e TwoPCPrepareCommitEvent) Message() string {
	return fmt.Sprintf("2PC prepare commit locked message: %s", shared.GreenWrap(e.msg))
}

func (e TwoPCAbortEvent) Severity() string {
	return "info"
}

func (e TwoPCAbortEvent) Message() string {
	return shared.RedWrap("2PC was aborted")
}
