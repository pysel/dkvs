package partition

import (
	"context"
	"fmt"
	"net"

	"github.com/pysel/dkvs/prototypes"
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
	"google.golang.org/grpc"
)

var ListeningPort int64

type ListenServer struct {
	pbpartition.UnimplementedCommandsServiceServer
}

func init() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", ListeningPort))
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pbpartition.RegisterCommandsServiceServer(grpcServer, &ListenServer{})
	grpcServer.Serve(lis)
}

func (ls *ListenServer) StoreMessage(ctx context.Context, req *prototypes.StoreMessageRequest) (*prototypes.StoreMessageResponse, error) {
	return nil, nil
}
