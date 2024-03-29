// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: dkvs/balancer/rpc.proto

package balancer

import (
	prototypes "github.com/pysel/dkvs/prototypes"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegisterPartitionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *RegisterPartitionRequest) Reset() {
	*x = RegisterPartitionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dkvs_balancer_rpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPartitionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPartitionRequest) ProtoMessage() {}

func (x *RegisterPartitionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dkvs_balancer_rpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPartitionRequest.ProtoReflect.Descriptor instead.
func (*RegisterPartitionRequest) Descriptor() ([]byte, []int) {
	return file_dkvs_balancer_rpc_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterPartitionRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type RegisterPartitionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterPartitionResponse) Reset() {
	*x = RegisterPartitionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dkvs_balancer_rpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterPartitionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterPartitionResponse) ProtoMessage() {}

func (x *RegisterPartitionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dkvs_balancer_rpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterPartitionResponse.ProtoReflect.Descriptor instead.
func (*RegisterPartitionResponse) Descriptor() ([]byte, []int) {
	return file_dkvs_balancer_rpc_proto_rawDescGZIP(), []int{1}
}

type GetIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetIdRequest) Reset() {
	*x = GetIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dkvs_balancer_rpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdRequest) ProtoMessage() {}

func (x *GetIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dkvs_balancer_rpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdRequest.ProtoReflect.Descriptor instead.
func (*GetIdRequest) Descriptor() ([]byte, []int) {
	return file_dkvs_balancer_rpc_proto_rawDescGZIP(), []int{2}
}

type GetIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetIdResponse) Reset() {
	*x = GetIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dkvs_balancer_rpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetIdResponse) ProtoMessage() {}

func (x *GetIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dkvs_balancer_rpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetIdResponse.ProtoReflect.Descriptor instead.
func (*GetIdResponse) Descriptor() ([]byte, []int) {
	return file_dkvs_balancer_rpc_proto_rawDescGZIP(), []int{3}
}

func (x *GetIdResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_dkvs_balancer_rpc_proto protoreflect.FileDescriptor

var file_dkvs_balancer_rpc_proto_rawDesc = []byte{
	0x0a, 0x17, 0x64, 0x6b, 0x76, 0x73, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2f,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x64, 0x6b, 0x76, 0x73, 0x2e,
	0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x1a, 0x1a, 0x64, 0x6b, 0x76, 0x73, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a, 0x18, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x1b, 0x0a, 0x19, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x32, 0x84, 0x03, 0x0a, 0x0f, 0x42, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x64, 0x6b, 0x76, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x64, 0x6b, 0x76, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x03, 0x53, 0x65,
	0x74, 0x12, 0x18, 0x2e, 0x64, 0x6b, 0x76, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x6b,
	0x76, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x1b, 0x2e, 0x64, 0x6b, 0x76, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x64, 0x6b, 0x76, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x68, 0x0a, 0x11, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x74, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x27, 0x2e, 0x64, 0x6b, 0x76, 0x73, 0x2e, 0x62, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e,
	0x64, 0x6b, 0x76, 0x73, 0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x50, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x05, 0x47, 0x65, 0x74,
	0x49, 0x64, 0x12, 0x1b, 0x2e, 0x64, 0x6b, 0x76, 0x73, 0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1c, 0x2e, 0x64, 0x6b, 0x76, 0x73, 0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x79,
	0x73, 0x65, 0x6c, 0x2f, 0x64, 0x6b, 0x76, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dkvs_balancer_rpc_proto_rawDescOnce sync.Once
	file_dkvs_balancer_rpc_proto_rawDescData = file_dkvs_balancer_rpc_proto_rawDesc
)

func file_dkvs_balancer_rpc_proto_rawDescGZIP() []byte {
	file_dkvs_balancer_rpc_proto_rawDescOnce.Do(func() {
		file_dkvs_balancer_rpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_dkvs_balancer_rpc_proto_rawDescData)
	})
	return file_dkvs_balancer_rpc_proto_rawDescData
}

var file_dkvs_balancer_rpc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_dkvs_balancer_rpc_proto_goTypes = []interface{}{
	(*RegisterPartitionRequest)(nil),  // 0: dkvs.balancer.RegisterPartitionRequest
	(*RegisterPartitionResponse)(nil), // 1: dkvs.balancer.RegisterPartitionResponse
	(*GetIdRequest)(nil),              // 2: dkvs.balancer.GetIdRequest
	(*GetIdResponse)(nil),             // 3: dkvs.balancer.GetIdResponse
	(*prototypes.GetRequest)(nil),     // 4: dkvs.message.GetRequest
	(*prototypes.SetRequest)(nil),     // 5: dkvs.message.SetRequest
	(*prototypes.DeleteRequest)(nil),  // 6: dkvs.message.DeleteRequest
	(*prototypes.GetResponse)(nil),    // 7: dkvs.message.GetResponse
	(*prototypes.SetResponse)(nil),    // 8: dkvs.message.SetResponse
	(*prototypes.DeleteResponse)(nil), // 9: dkvs.message.DeleteResponse
}
var file_dkvs_balancer_rpc_proto_depIdxs = []int32{
	4, // 0: dkvs.balancer.BalancerService.Get:input_type -> dkvs.message.GetRequest
	5, // 1: dkvs.balancer.BalancerService.Set:input_type -> dkvs.message.SetRequest
	6, // 2: dkvs.balancer.BalancerService.Delete:input_type -> dkvs.message.DeleteRequest
	0, // 3: dkvs.balancer.BalancerService.RegisterPartition:input_type -> dkvs.balancer.RegisterPartitionRequest
	2, // 4: dkvs.balancer.BalancerService.GetId:input_type -> dkvs.balancer.GetIdRequest
	7, // 5: dkvs.balancer.BalancerService.Get:output_type -> dkvs.message.GetResponse
	8, // 6: dkvs.balancer.BalancerService.Set:output_type -> dkvs.message.SetResponse
	9, // 7: dkvs.balancer.BalancerService.Delete:output_type -> dkvs.message.DeleteResponse
	1, // 8: dkvs.balancer.BalancerService.RegisterPartition:output_type -> dkvs.balancer.RegisterPartitionResponse
	3, // 9: dkvs.balancer.BalancerService.GetId:output_type -> dkvs.balancer.GetIdResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dkvs_balancer_rpc_proto_init() }
func file_dkvs_balancer_rpc_proto_init() {
	if File_dkvs_balancer_rpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dkvs_balancer_rpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPartitionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dkvs_balancer_rpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterPartitionResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dkvs_balancer_rpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_dkvs_balancer_rpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetIdResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dkvs_balancer_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dkvs_balancer_rpc_proto_goTypes,
		DependencyIndexes: file_dkvs_balancer_rpc_proto_depIdxs,
		MessageInfos:      file_dkvs_balancer_rpc_proto_msgTypes,
	}.Build()
	File_dkvs_balancer_rpc_proto = out.File
	file_dkvs_balancer_rpc_proto_rawDesc = nil
	file_dkvs_balancer_rpc_proto_goTypes = nil
	file_dkvs_balancer_rpc_proto_depIdxs = nil
}
