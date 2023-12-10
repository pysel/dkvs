package balancer

import (
	"log"

	pbbalancer "github.com/pysel/dkvs/prototypes/balancer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewBalancerClient creates a new client to a balancer.
func NewBalancerClient(addr string) pbbalancer.BalancerServiceClient {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	client := pbbalancer.NewBalancerServiceClient(conn)
	return client
}
