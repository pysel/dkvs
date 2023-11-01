// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: dkvs/partition/rpc.proto

package partition

import (
	context "context"
	prototypes "github.com/pysel/dkvs/prototypes"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CommandsServiceClient is the client API for CommandsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommandsServiceClient interface {
	StoreMessage(ctx context.Context, in *prototypes.StoreMessageRequest, opts ...grpc.CallOption) (*prototypes.StoreMessageResponse, error)
	GetMessage(ctx context.Context, in *prototypes.GetMessageRequest, opts ...grpc.CallOption) (*prototypes.GetMessageResponse, error)
}

type commandsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommandsServiceClient(cc grpc.ClientConnInterface) CommandsServiceClient {
	return &commandsServiceClient{cc}
}

func (c *commandsServiceClient) StoreMessage(ctx context.Context, in *prototypes.StoreMessageRequest, opts ...grpc.CallOption) (*prototypes.StoreMessageResponse, error) {
	out := new(prototypes.StoreMessageResponse)
	err := c.cc.Invoke(ctx, "/dkvs.partition.CommandsService/StoreMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandsServiceClient) GetMessage(ctx context.Context, in *prototypes.GetMessageRequest, opts ...grpc.CallOption) (*prototypes.GetMessageResponse, error) {
	out := new(prototypes.GetMessageResponse)
	err := c.cc.Invoke(ctx, "/dkvs.partition.CommandsService/GetMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommandsServiceServer is the server API for CommandsService service.
// All implementations must embed UnimplementedCommandsServiceServer
// for forward compatibility
type CommandsServiceServer interface {
	StoreMessage(context.Context, *prototypes.StoreMessageRequest) (*prototypes.StoreMessageResponse, error)
	GetMessage(context.Context, *prototypes.GetMessageRequest) (*prototypes.GetMessageResponse, error)
	mustEmbedUnimplementedCommandsServiceServer()
}

// UnimplementedCommandsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommandsServiceServer struct {
}

func (UnimplementedCommandsServiceServer) StoreMessage(context.Context, *prototypes.StoreMessageRequest) (*prototypes.StoreMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreMessage not implemented")
}
func (UnimplementedCommandsServiceServer) GetMessage(context.Context, *prototypes.GetMessageRequest) (*prototypes.GetMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedCommandsServiceServer) mustEmbedUnimplementedCommandsServiceServer() {}

// UnsafeCommandsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommandsServiceServer will
// result in compilation errors.
type UnsafeCommandsServiceServer interface {
	mustEmbedUnimplementedCommandsServiceServer()
}

func RegisterCommandsServiceServer(s grpc.ServiceRegistrar, srv CommandsServiceServer) {
	s.RegisterService(&CommandsService_ServiceDesc, srv)
}

func _CommandsService_StoreMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(prototypes.StoreMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandsServiceServer).StoreMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkvs.partition.CommandsService/StoreMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandsServiceServer).StoreMessage(ctx, req.(*prototypes.StoreMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandsService_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(prototypes.GetMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandsServiceServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkvs.partition.CommandsService/GetMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandsServiceServer).GetMessage(ctx, req.(*prototypes.GetMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommandsService_ServiceDesc is the grpc.ServiceDesc for CommandsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommandsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dkvs.partition.CommandsService",
	HandlerType: (*CommandsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreMessage",
			Handler:    _CommandsService_StoreMessage_Handler,
		},
		{
			MethodName: "GetMessage",
			Handler:    _CommandsService_GetMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dkvs/partition/rpc.proto",
}
