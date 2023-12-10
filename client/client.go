package client

import (
	pbbalancer "github.com/pysel/dkvs/prototypes/balancer"
)

// Client is an object that is responsible for interacting with DKVS.
type Client struct {
	timestamp int

	// a client to balancer server
	balacerClient pbbalancer.BalancerServiceClient
}

func NewClient(balancerAddr string) *Client {
	c := &Client{timestamp: 0}
	c.setupBalancerClient(balancerAddr)
	return c
}
