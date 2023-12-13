package client

import (
	"context"

	"github.com/pysel/dkvs/prototypes"
	pbbalancer "github.com/pysel/dkvs/prototypes/balancer"
	"github.com/pysel/dkvs/types"
)

// Client is an object that is responsible for interacting with DKVS.
type Client struct {
	// go context
	context context.Context

	// logical timestamp
	timestamp uint64

	// a client to balancer server
	balacerClient pbbalancer.BalancerServiceClient

	// a list of messages that might have not yet processed by dkvs.
	nonConfirmedList types.Backlog
}

func NewClient(balancerAddr string) *Client {
	context := context.Background()

	c := &Client{
		timestamp: 0,
		context:   context,
	}

	c.setupBalancerClient(balancerAddr)
	return c
}

// Set sets a value for a key.
func (c *Client) Set(key, value []byte) error {
	req := &prototypes.SetRequest{
		Key:     key,
		Value:   value,
		Lamport: c.timestamp,
	}

	_, err := c.balacerClient.Set(c.context, req)
	if err != nil {
		return err
	}

	c.timestamp++

	return nil
}

// Get gets a value for a key.
func (c *Client) Get(key []byte) ([]byte, error) {
	req := &prototypes.GetRequest{
		Key:     key,
		Lamport: c.timestamp,
	}

	resp, err := c.balacerClient.Get(c.context, req)
	if err != nil {
		return nil, err
	}

	c.timestamp++

	return resp.StoredValue.Value, nil
}

// Delete deletes a value for a key.
func (c *Client) Delete(key []byte) error {
	req := &prototypes.DeleteRequest{
		Key:     key,
		Lamport: c.timestamp,
	}

	_, err := c.balacerClient.Delete(c.context, req)
	if err != nil {
		return err
	}

	c.timestamp++

	return nil
}

// func (c *Client) processGrpcError(err error) {

// }
