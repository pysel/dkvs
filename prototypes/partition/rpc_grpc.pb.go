// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
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

// PartitionServiceClient is the client API for PartitionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PartitionServiceClient interface {
	Set(ctx context.Context, in *prototypes.SetRequest, opts ...grpc.CallOption) (*prototypes.SetResponse, error)
	Get(ctx context.Context, in *prototypes.GetRequest, opts ...grpc.CallOption) (*prototypes.GetResponse, error)
	Delete(ctx context.Context, in *prototypes.DeleteRequest, opts ...grpc.CallOption) (*prototypes.DeleteResponse, error)
	// Two-phase commit
	PrepareCommit(ctx context.Context, in *PrepareCommitRequest, opts ...grpc.CallOption) (*PrepareCommitResponse, error)
	AbortCommit(ctx context.Context, in *AbortCommitRequest, opts ...grpc.CallOption) (*AbortCommitResponse, error)
	Commit(ctx context.Context, in *CommitRequest, opts ...grpc.CallOption) (*CommitResponse, error)
	// SetHashrange sets this node's hashrange to the given range.
	SetHashrange(ctx context.Context, in *prototypes.SetHashrangeRequest, opts ...grpc.CallOption) (*prototypes.SetHashrangeResponse, error)
}

type partitionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPartitionServiceClient(cc grpc.ClientConnInterface) PartitionServiceClient {
	return &partitionServiceClient{cc}
}

func (c *partitionServiceClient) Set(ctx context.Context, in *prototypes.SetRequest, opts ...grpc.CallOption) (*prototypes.SetResponse, error) {
	out := new(prototypes.SetResponse)
	err := c.cc.Invoke(ctx, "/dkvs.partition.PartitionService/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partitionServiceClient) Get(ctx context.Context, in *prototypes.GetRequest, opts ...grpc.CallOption) (*prototypes.GetResponse, error) {
	out := new(prototypes.GetResponse)
	err := c.cc.Invoke(ctx, "/dkvs.partition.PartitionService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partitionServiceClient) Delete(ctx context.Context, in *prototypes.DeleteRequest, opts ...grpc.CallOption) (*prototypes.DeleteResponse, error) {
	out := new(prototypes.DeleteResponse)
	err := c.cc.Invoke(ctx, "/dkvs.partition.PartitionService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partitionServiceClient) PrepareCommit(ctx context.Context, in *PrepareCommitRequest, opts ...grpc.CallOption) (*PrepareCommitResponse, error) {
	out := new(PrepareCommitResponse)
	err := c.cc.Invoke(ctx, "/dkvs.partition.PartitionService/PrepareCommit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partitionServiceClient) AbortCommit(ctx context.Context, in *AbortCommitRequest, opts ...grpc.CallOption) (*AbortCommitResponse, error) {
	out := new(AbortCommitResponse)
	err := c.cc.Invoke(ctx, "/dkvs.partition.PartitionService/AbortCommit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partitionServiceClient) Commit(ctx context.Context, in *CommitRequest, opts ...grpc.CallOption) (*CommitResponse, error) {
	out := new(CommitResponse)
	err := c.cc.Invoke(ctx, "/dkvs.partition.PartitionService/Commit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partitionServiceClient) SetHashrange(ctx context.Context, in *prototypes.SetHashrangeRequest, opts ...grpc.CallOption) (*prototypes.SetHashrangeResponse, error) {
	out := new(prototypes.SetHashrangeResponse)
	err := c.cc.Invoke(ctx, "/dkvs.partition.PartitionService/SetHashrange", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PartitionServiceServer is the server API for PartitionService service.
// All implementations should embed UnimplementedPartitionServiceServer
// for forward compatibility
type PartitionServiceServer interface {
	Set(context.Context, *prototypes.SetRequest) (*prototypes.SetResponse, error)
	Get(context.Context, *prototypes.GetRequest) (*prototypes.GetResponse, error)
	Delete(context.Context, *prototypes.DeleteRequest) (*prototypes.DeleteResponse, error)
	// Two-phase commit
	PrepareCommit(context.Context, *PrepareCommitRequest) (*PrepareCommitResponse, error)
	AbortCommit(context.Context, *AbortCommitRequest) (*AbortCommitResponse, error)
	Commit(context.Context, *CommitRequest) (*CommitResponse, error)
	// SetHashrange sets this node's hashrange to the given range.
	SetHashrange(context.Context, *prototypes.SetHashrangeRequest) (*prototypes.SetHashrangeResponse, error)
}

// UnimplementedPartitionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPartitionServiceServer struct {
}

func (UnimplementedPartitionServiceServer) Set(context.Context, *prototypes.SetRequest) (*prototypes.SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedPartitionServiceServer) Get(context.Context, *prototypes.GetRequest) (*prototypes.GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedPartitionServiceServer) Delete(context.Context, *prototypes.DeleteRequest) (*prototypes.DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedPartitionServiceServer) PrepareCommit(context.Context, *PrepareCommitRequest) (*PrepareCommitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareCommit not implemented")
}
func (UnimplementedPartitionServiceServer) AbortCommit(context.Context, *AbortCommitRequest) (*AbortCommitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AbortCommit not implemented")
}
func (UnimplementedPartitionServiceServer) Commit(context.Context, *CommitRequest) (*CommitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commit not implemented")
}
func (UnimplementedPartitionServiceServer) SetHashrange(context.Context, *prototypes.SetHashrangeRequest) (*prototypes.SetHashrangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetHashrange not implemented")
}

// UnsafePartitionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PartitionServiceServer will
// result in compilation errors.
type UnsafePartitionServiceServer interface {
	mustEmbedUnimplementedPartitionServiceServer()
}

func RegisterPartitionServiceServer(s grpc.ServiceRegistrar, srv PartitionServiceServer) {
	s.RegisterService(&PartitionService_ServiceDesc, srv)
}

func _PartitionService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(prototypes.SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartitionServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkvs.partition.PartitionService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartitionServiceServer).Set(ctx, req.(*prototypes.SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartitionService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(prototypes.GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartitionServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkvs.partition.PartitionService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartitionServiceServer).Get(ctx, req.(*prototypes.GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartitionService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(prototypes.DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartitionServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkvs.partition.PartitionService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartitionServiceServer).Delete(ctx, req.(*prototypes.DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartitionService_PrepareCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareCommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartitionServiceServer).PrepareCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkvs.partition.PartitionService/PrepareCommit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartitionServiceServer).PrepareCommit(ctx, req.(*PrepareCommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartitionService_AbortCommit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AbortCommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartitionServiceServer).AbortCommit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkvs.partition.PartitionService/AbortCommit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartitionServiceServer).AbortCommit(ctx, req.(*AbortCommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartitionService_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartitionServiceServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkvs.partition.PartitionService/Commit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartitionServiceServer).Commit(ctx, req.(*CommitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartitionService_SetHashrange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(prototypes.SetHashrangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartitionServiceServer).SetHashrange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dkvs.partition.PartitionService/SetHashrange",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartitionServiceServer).SetHashrange(ctx, req.(*prototypes.SetHashrangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PartitionService_ServiceDesc is the grpc.ServiceDesc for PartitionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PartitionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dkvs.partition.PartitionService",
	HandlerType: (*PartitionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _PartitionService_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _PartitionService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PartitionService_Delete_Handler,
		},
		{
			MethodName: "PrepareCommit",
			Handler:    _PartitionService_PrepareCommit_Handler,
		},
		{
			MethodName: "AbortCommit",
			Handler:    _PartitionService_AbortCommit_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _PartitionService_Commit_Handler,
		},
		{
			MethodName: "SetHashrange",
			Handler:    _PartitionService_SetHashrange_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dkvs/partition/rpc.proto",
}
