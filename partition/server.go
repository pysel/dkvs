package partition

import (
	pbpartition "github.com/pysel/dkvs/prototypes/partition"
	"github.com/pysel/dkvs/shared"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type ListenServer struct {
	pbpartition.UnimplementedPartitionServiceServer

	// event handler
	eventHandler *shared.EventHandler

	*Partition
}

func RegisterPartitionServer(partition *Partition, eventHandler *shared.EventHandler) *grpc.Server {
	s := grpc.NewServer()
	reflection.Register(s)
	pbpartition.RegisterPartitionServiceServer(s, &ListenServer{Partition: partition, eventHandler: eventHandler})

	return s
}
