package balancer

import (
	"fmt"
	"net"

	"log"

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

// startListeningOnPort starts a grpc server listening on the given port.
func startListeningOnPort(s *grpc.Server, port int64) net.Addr {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	go func() {
		err := s.Serve(lis)
		if err != nil {
			log.Fatalf("Balancer server exited with error: %v", err)
		}
	}()

	return lis.Addr()
}
