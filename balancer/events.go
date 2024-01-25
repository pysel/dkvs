package balancer

import (
	"fmt"

	"github.com/pysel/dkvs/shared"
)

type (
	// RegisterPartition registers a partition in the balancer.
	RegisterPartitionEvent struct {
		Address string
	}

	GetEvent struct {
		msg string
	}

	SetEvent struct {
		msg string
	}

	DeleteEvent struct {
		msg string
	}

	PartitionOfflineEvent struct { // TODO: use it
		Address string
		Err     error
	}

	TwoPhaseCommitFailedEvent struct {
		Reason string
	}
)

func (e *RegisterPartitionEvent) Severity() string {
	return "info"
}

func (e *RegisterPartitionEvent) Message() string {
	return fmt.Sprintf("Registered partition on address %s", shared.GreenWrap(e.Address))
}

func (e *PartitionOfflineEvent) Severity() string {
	return "warning"
}

func (e *PartitionOfflineEvent) Message() string {
	return fmt.Sprintf("Lost connection to partition %s", shared.YellowWrap(e.Address))
}

func (e *GetEvent) Severity() string {
	return "info"
}

func (e *GetEvent) Message() string {
	return fmt.Sprintf("Relayed GET request: %s", shared.GreyWrap(e.msg))
}

func (e *SetEvent) Severity() string {
	return "info"
}

func (e *SetEvent) Message() string {
	return fmt.Sprintf("Relayed SET request: %s", shared.GreyWrap(e.msg))
}

func (e *DeleteEvent) Severity() string {
	return "info"
}

func (e *DeleteEvent) Message() string {
	return fmt.Sprintf("Relayed DELETE request: %s", shared.GreyWrap(e.msg))
}

func (e *TwoPhaseCommitFailedEvent) Severity() string {
	return "error"
}

func (e *TwoPhaseCommitFailedEvent) Message() string {
	return fmt.Sprintf("Two phase commit failed: %s", shared.RedWrap(e.Reason))
}
