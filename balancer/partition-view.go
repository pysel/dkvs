package balancer

import (
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
)

// RangeView is an abstraction over clients to partitions responsible for the same range
type RangeView struct {
	clients []*pbpartition.PartitionServiceClient
	lamport uint64
}

func NewRangeView(clients []*pbpartition.PartitionServiceClient) *RangeView {
	return &RangeView{clients: clients, lamport: 0}
}

// AddPartitionClient adds a client to the set of clients in a range view
func (rv *RangeView) AddPartitionClient(client *pbpartition.PartitionServiceClient) {
	rv.clients = append(rv.clients, client)
}

func (rv *RangeView) GetResponsibleClients() []*pbpartition.PartitionServiceClient {
	return rv.clients
}
