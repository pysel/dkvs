// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.9
// source: dkvs/partition/rpc.proto

package partition

import (
	prototypes "github.com/pysel/dkvs/prototypes"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_dkvs_partition_rpc_proto protoreflect.FileDescriptor

var file_dkvs_partition_rpc_proto_rawDesc = []byte{
	0x0a, 0x18, 0x64, 0x6b, 0x76, 0x73, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x64, 0x6b, 0x76, 0x73,
	0x2e, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x1a, 0x1a, 0x64, 0x6b, 0x76, 0x73, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb9, 0x01, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0c, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x2e, 0x64, 0x6b, 0x76, 0x73,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x64,
	0x6b, 0x76, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4f, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f,
	0x2e, 0x64, 0x6b, 0x76, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x64, 0x6b, 0x76, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x70, 0x79, 0x73, 0x65, 0x6c, 0x2f, 0x64, 0x6b, 0x76, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_dkvs_partition_rpc_proto_goTypes = []interface{}{
	(*prototypes.StoreMessageRequest)(nil),  // 0: dkvs.message.StoreMessageRequest
	(*prototypes.GetMessageRequest)(nil),    // 1: dkvs.message.GetMessageRequest
	(*prototypes.StoreMessageResponse)(nil), // 2: dkvs.message.StoreMessageResponse
	(*prototypes.GetMessageResponse)(nil),   // 3: dkvs.message.GetMessageResponse
}
var file_dkvs_partition_rpc_proto_depIdxs = []int32{
	0, // 0: dkvs.balancer.CommandsService.StoreMessage:input_type -> dkvs.message.StoreMessageRequest
	1, // 1: dkvs.balancer.CommandsService.GetMessage:input_type -> dkvs.message.GetMessageRequest
	2, // 2: dkvs.balancer.CommandsService.StoreMessage:output_type -> dkvs.message.StoreMessageResponse
	3, // 3: dkvs.balancer.CommandsService.GetMessage:output_type -> dkvs.message.GetMessageResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dkvs_partition_rpc_proto_init() }
func file_dkvs_partition_rpc_proto_init() {
	if File_dkvs_partition_rpc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_dkvs_partition_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dkvs_partition_rpc_proto_goTypes,
		DependencyIndexes: file_dkvs_partition_rpc_proto_depIdxs,
	}.Build()
	File_dkvs_partition_rpc_proto = out.File
	file_dkvs_partition_rpc_proto_rawDesc = nil
	file_dkvs_partition_rpc_proto_goTypes = nil
	file_dkvs_partition_rpc_proto_depIdxs = nil
}
