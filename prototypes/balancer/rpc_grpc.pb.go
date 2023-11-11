// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: dkvs/balancer/rpc.proto

package balancer

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

// BalancerServiceClient is the client API for BalancerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BalancerServiceClient interface {
	// ----- To be relayed requests -----
	Get(ctx context.Context, in *prototypes.GetRequest, opts ...grpc.CallOption) (*prototypes.GetResponse, error)
	Set(ctx context.Context, in *prototypes.SetRequest, opts ...grpc.CallOption) (*prototypes.SetResponse, error)
	Delete(ctx context.Context, in *prototypes.DeleteRequest, opts ...grpc.CallOption) (*prototypes.DeleteResponse, error)
	// RegisterPartition is called by a partition to register itself with the balancer
	// The balancer will set partition's range and run a new client of this partition's server
	RegisterPartition(ctx context.Context, in *RegisterPartitionRequest, opts ...grpc.CallOption) (*RegisterPartitionResponse, error)
}

type balancerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBalancerServiceClient(cc grpc.ClientConnInterface) BalancerServiceClient {
	return &balancerServiceClient{cc}
}

func (c *balancerServiceClient) Get(ctx context.Context, in *prototypes.GetRequest, opts ...grpc.CallOption) (*prototypes.GetResponse, error) {
	out := new(prototypes.GetResponse)
	err := c.cc.Invoke(ctx, "/dkvs.balancer.BalancerService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balancerServiceClient) Set(ctx context.Context, in *prototypes.SetRequest, opts ...grpc.CallOption) (*prototypes.SetResponse, error) {
	out := new(prototypes.SetResponse)
	err := c.cc.Invoke(ctx, "/dkvs.balancer.BalancerService/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balancerServiceClient) Delete(ctx context.Context, in *prototypes.DeleteRequest, opts ...grpc.CallOption) (*prototypes.DeleteResponse, error) {
	out := new(prototypes.DeleteResponse)
	err := c.cc.Invoke(ctx, "/dkvs.balancer.BalancerService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balancerServiceClient) RegisterPartition(ctx context.Context, in *RegisterPartitionRequest, opts ...grpc.CallOption) (*RegisterPartitionResponse, error) {
	out := new(RegisterPartitionResponse)
	err := c.cc.Invoke(ctx, "/dkvs.balancer.BalancerService/RegisterPartition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BalancerServiceServer is the server API for BalancerService service.
// All implementations must embed UnimplementedBalancerServiceServer
// for forward compatibility
type BalancerServiceServer interface {
	// ----- To be relayed requests -----
	Get(context.Context, *prototypes.GetRequest) (*prototypes.GetResponse, error)
	Set(context.Context, *prototypes.SetRequest) (*prototypes.SetResponse, error)
	Delete(context.Context, *prototypes.DeleteRequest) (*prototypes.DeleteResponse, error)
	// RegisterPartition is called by a partition to register itself with the balancer
	// The balancer will set partition's range and run a new client of this partition's server
	RegisterPartition(context.Context, *RegisterPartitionRequest) (*RegisterPartitionResponse, error)
	mustEmbedUnimplementedBalancerServiceServer()
}

// UnimplementedBalancerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBalancerServiceServer struct {
}

func (UnimplementedBalancerServiceServer) Get(context.Context, *prototypes.GetRequest) (*prototypes.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedBalancerServiceServer) Set(context.Context, *prototypes.SetRequest) (*prototypes.SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedBalancerServiceServer) Delete(context.Context, *prototypes.DeleteRequest) (*prototypes.DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBalancerServiceServer) RegisterPartition(context.Context, *RegisterPartitionRequest) (*RegisterPartitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterPartition not implemented")
}
func (UnimplementedBalancerServiceServer) mustEmbedUnimplementedBalancerServiceServer() {}

// UnsafeBalancerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BalancerServiceServer will
// result in compilation errors.
type UnsafeBalancerServiceServer interface {
	mustEmbedUnimplementedBalancerServiceServer()
}

func RegisterBalancerServiceServer(s grpc.ServiceRegistrar, srv BalancerServiceServer) {
	s.RegisterService(&BalancerService_ServiceDesc, srv)
}

func _BalancerService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(prototypes.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancerServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkvs.balancer.BalancerService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancerServiceServer).Get(ctx, req.(*prototypes.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BalancerService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(prototypes.SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancerServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkvs.balancer.BalancerService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancerServiceServer).Set(ctx, req.(*prototypes.SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BalancerService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(prototypes.DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancerServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkvs.balancer.BalancerService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancerServiceServer).Delete(ctx, req.(*prototypes.DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BalancerService_RegisterPartition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterPartitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancerServiceServer).RegisterPartition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkvs.balancer.BalancerService/RegisterPartition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancerServiceServer).RegisterPartition(ctx, req.(*RegisterPartitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BalancerService_ServiceDesc is the grpc.ServiceDesc for BalancerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BalancerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dkvs.balancer.BalancerService",
	HandlerType: (*BalancerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _BalancerService_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _BalancerService_Set_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BalancerService_Delete_Handler,
		},
		{
			MethodName: "RegisterPartition",
			Handler:    _BalancerService_RegisterPartition_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dkvs/balancer/rpc.proto",
}
