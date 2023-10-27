package balancer

import "net"

// A Partition is a slave node that stores some range of keys.
// Stores only the necessary information for Balancer.
type BPartition struct {
	ip   net.IP
	port int64
}

func NewBPartition(ip net.IP, port int64) *BPartition {
	return &BPartition{
		ip,
		port,
	}
}
