package balancer

import (
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
)

// PartitionView is an abstraction over a partition server connection.
type PartitionView struct {
	client  *pbpartition.PartitionServiceClient
	lamport uint64
}

func NewPartitionView(client *pbpartition.PartitionServiceClient) *PartitionView {
	return &PartitionView{client: client, lamport: 0}
}
