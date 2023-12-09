package partition

import "google.golang.org/protobuf/proto"

type Backlog ([]proto.Message)

func (b *Backlog) Add(message proto.Message) {
	*b = append(*b, message)
}

func (b *Backlog) Pop() proto.Message {
	if len(*b) == 0 {
		return nil
	}

	message := (*b)[0]
	*b = (*b)[1:]
	return message
}
