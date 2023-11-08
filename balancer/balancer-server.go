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

func RunBalancerServer(port int64) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	balancer := NewBalancer()

	s := grpc.NewServer()
	reflection.Register(s)
	pbbalancer.RegisterBalancerServiceServer(s, &BalancerServer{Balancer: balancer})
	go s.Serve(lis)
}