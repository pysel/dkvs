package balancer

import (
	"fmt"
	"net"

	pbbalancer "github.com/pysel/dkvs/prototypes/balancer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type BalancerServer struct {
	pbbalancer.UnimplementedBalancerServiceServer
	*Balancer
}

// RunBalancerServer creates a new grpc server and registers the balancer service.
func RegisterBalancerServer(b *Balancer) *grpc.Server {
	s := grpc.NewServer()
	reflection.Register(s)
	pbbalancer.RegisterBalancerServiceServer(s, &BalancerServer{Balancer: b})

	return s
}

// listenOnPort starts a grpc server listening on the given port.
func listenOnPort(s *grpc.Server, port int64) net.Addr {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	go s.Serve(lis)
	return lis.Addr()
}
