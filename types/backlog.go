package types

import (
	"google.golang.org/protobuf/proto"
)

type messages ([]struct {
	timestamp uint64
	msg       proto.Message
})

// Add appends a message to messages and preserves the ascending order of timestamps
func (m *messages) Add(ts uint64, message proto.Message) {
	for i, msg := range *m {
		if msg.timestamp > ts {
			*m = append((*m)[:i], append([]struct {
				timestamp uint64
				msg       proto.Message
			}{{ts, message}}, (*m)[i:]...)...)
			return
		}
	}
}

// Backlog is a mapping of client id to messages.
// Messages are stored in ascending order of timestamps.
type Backlog (map[string]messages)

func NewBacklog() *Backlog {
	return new(Backlog)
}

// Add appends a message to ID's backlog and preserves the ascending order of timestamps.
func (b *Backlog) Add(id string, ts uint64, message proto.Message) {
	messages := (*b)[id]
	messages.Add(ts, message)
}

// Pop removes and returns the first message with timestamp less than or equal to ts.
func (b *Backlog) Pop(id string, ts uint64) proto.Message {
	messages := (*b)[id]
	for i, msg := range messages {
		if msg.timestamp <= ts {
			(*b)[id] = append(messages[:i], messages[i+1:]...)
			return msg.msg
		}
	}

	return nil
}
