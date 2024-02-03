package types

import (
	"google.golang.org/protobuf/proto"
)

// Backlog is a list of messages and corresponding timestamps.
type Backlog ([]struct {
	timestamp uint64
	msg       proto.Message
})

func NewBacklog() *Backlog {
	return new(Backlog)
}

// Add appends a message to messages and preserves the ascending order of timestamps
func (b *Backlog) Add(ts uint64, message proto.Message) {
	element := struct {
		timestamp uint64
		msg       proto.Message
	}{ts, message}
	for i, msg := range *b {
		if msg.timestamp > ts {
			*b = append(
				append((*b)[:i], element), // append element to the slice of a messages with lower timestamps
				(*b)[i:]...,               // append the rest of the messages
			)
			return
		}
	}

	// if the message has the highest timestamp, append it to the end
	*b = append(*b, element)
}

// Pop returns the first message (the one with the smallest timestamp) and removes it from the backlog.
func (b *Backlog) Pop() proto.Message {
	if len(*b) == 0 {
		return nil
	}

	msg := (*b)[0].msg
	*b = (*b)[1:]
	return msg
}
