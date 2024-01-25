package balancer

import (
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
)

// RangeView is an abstraction over clients to partitions responsible for the same range
type RangeView struct {
	clients   []*pbpartition.PartitionServiceClient
	addresses []string
	lamport   uint64
}

func NewRangeView(clients []*pbpartition.PartitionServiceClient, addresses []string) *RangeView {
	return &RangeView{clients: clients, addresses: addresses, lamport: 0}
}

// AddPartitionClient adds a client to the set of clients in a range view
func (rv *RangeView) AddPartitionData(client *pbpartition.PartitionServiceClient, address string) {
	rv.clients = append(rv.clients, client)
	rv.addresses = append(rv.addresses, address)
}

func (rv *RangeView) GetResponsibleClients() []*pbpartition.PartitionServiceClient {
	return rv.clients
}

// RemovePartition removes a partition from the balancer's registry.
func (rv *RangeView) removePartition(addr string) error {
	for i, address := range rv.addresses {
		if address == addr {
			rv.clients = append(rv.clients[:i], rv.clients[i+1:]...)
			rv.addresses = append(rv.addresses[:i], rv.addresses[i+1:]...)
			return nil
		}
	}
	return ErrPartitionAtAddressNotExist
}
