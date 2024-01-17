package balancer

import (
	pbbalancer "github.com/pysel/dkvs/prototypes/balancer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type BalancerServer struct {
	pbbalancer.UnimplementedBalancerServiceServer
	*Balancer
}

// RegisterBalancerServer creates a new grpc server and registers the balancer service.
func RegisterBalancerServer(b *Balancer) *grpc.Server {
	s := grpc.NewServer()
	reflection.Register(s)
	pbbalancer.RegisterBalancerServiceServer(s, &BalancerServer{Balancer: b})

	return s
}
