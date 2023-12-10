package client

import "github.com/pysel/dkvs/balancer"

func (c *Client) setupBalancerClient(addr string) {
	c.balacerClient = balancer.NewBalancerClient(addr)
}
