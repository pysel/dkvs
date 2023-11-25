// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: dkvs/message/message.proto

package prototypes

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

// StoreMessage is a message that is sent to a node that should store a key
type SetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// a lamport timestamp of this message
	Lamport uint64 `protobuf:"varint,3,opt,name=lamport,proto3" json:"lamport,omitempty"`
}

func (x *SetRequest) Reset() {
	*x = SetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dkvs_message_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRequest) ProtoMessage() {}

func (x *SetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dkvs_message_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRequest.ProtoReflect.Descriptor instead.
func (*SetRequest) Descriptor() ([]byte, []int) {
	return file_dkvs_message_message_proto_rawDescGZIP(), []int{0}
}

func (x *SetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetRequest) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

func (x *SetRequest) GetLamport() uint64 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

type SetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetResponse) Reset() {
	*x = SetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dkvs_message_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetResponse) ProtoMessage() {}

func (x *SetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dkvs_message_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetResponse.ProtoReflect.Descriptor instead.
func (*SetResponse) Descriptor() ([]byte, []int) {
	return file_dkvs_message_message_proto_rawDescGZIP(), []int{1}
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// a lamport timestamp of this message
	Lamport uint64 `protobuf:"varint,2,opt,name=lamport,proto3" json:"lamport,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dkvs_message_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dkvs_message_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_dkvs_message_message_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GetRequest) GetLamport() uint64 {
	if x != nil {
		return x.Lamport
	}
	return 0
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoredValue *StoredValue `protobuf:"bytes,1,opt,name=stored_value,json=storedValue,proto3" json:"stored_value,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dkvs_message_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dkvs_message_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_dkvs_message_message_proto_rawDescGZIP(), []int{3}
}

func (x *GetResponse) GetStoredValue() *StoredValue {
	if x != nil {
		return x.StoredValue
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dkvs_message_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dkvs_message_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_dkvs_message_message_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dkvs_message_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dkvs_message_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_dkvs_message_message_proto_rawDescGZIP(), []int{5}
}

type SetHashrangeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Min []byte `protobuf:"bytes,1,opt,name=min,proto3" json:"min,omitempty"`
	Max []byte `protobuf:"bytes,2,opt,name=max,proto3" json:"max,omitempty"`
}

func (x *SetHashrangeRequest) Reset() {
	*x = SetHashrangeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dkvs_message_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetHashrangeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetHashrangeRequest) ProtoMessage() {}

func (x *SetHashrangeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dkvs_message_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetHashrangeRequest.ProtoReflect.Descriptor instead.
func (*SetHashrangeRequest) Descriptor() ([]byte, []int) {
	return file_dkvs_message_message_proto_rawDescGZIP(), []int{6}
}

func (x *SetHashrangeRequest) GetMin() []byte {
	if x != nil {
		return x.Min
	}
	return nil
}

func (x *SetHashrangeRequest) GetMax() []byte {
	if x != nil {
		return x.Max
	}
	return nil
}

type SetHashrangeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetHashrangeResponse) Reset() {
	*x = SetHashrangeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dkvs_message_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetHashrangeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetHashrangeResponse) ProtoMessage() {}

func (x *SetHashrangeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dkvs_message_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetHashrangeResponse.ProtoReflect.Descriptor instead.
func (*SetHashrangeResponse) Descriptor() ([]byte, []int) {
	return file_dkvs_message_message_proto_rawDescGZIP(), []int{7}
}

var File_dkvs_message_message_proto protoreflect.FileDescriptor

var file_dkvs_message_message_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x64, 0x6b, 0x76, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x64, 0x6b,
	0x76, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x64, 0x6b, 0x76, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x60, 0x0a,
	0x0a, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x7a, 0x02, 0x10, 0x01, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x22,
	0x0d, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x41,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6c, 0x61, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x22, 0x4b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3c, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x6b, 0x76, 0x73, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x2a,
	0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39, 0x0a, 0x13,
	0x53, 0x65, 0x74, 0x48, 0x61, 0x73, 0x68, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x03, 0x6d, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x22, 0x16, 0x0a, 0x14, 0x53, 0x65, 0x74, 0x48, 0x61,
	0x73, 0x68, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x79,
	0x73, 0x65, 0x6c, 0x2f, 0x64, 0x6b, 0x76, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dkvs_message_message_proto_rawDescOnce sync.Once
	file_dkvs_message_message_proto_rawDescData = file_dkvs_message_message_proto_rawDesc
)

func file_dkvs_message_message_proto_rawDescGZIP() []byte {
	file_dkvs_message_message_proto_rawDescOnce.Do(func() {
		file_dkvs_message_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_dkvs_message_message_proto_rawDescData)
	})
	return file_dkvs_message_message_proto_rawDescData
}

var file_dkvs_message_message_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_dkvs_message_message_proto_goTypes = []interface{}{
	(*SetRequest)(nil),           // 0: dkvs.message.SetRequest
	(*SetResponse)(nil),          // 1: dkvs.message.SetResponse
	(*GetRequest)(nil),           // 2: dkvs.message.GetRequest
	(*GetResponse)(nil),          // 3: dkvs.message.GetResponse
	(*DeleteRequest)(nil),        // 4: dkvs.message.DeleteRequest
	(*DeleteResponse)(nil),       // 5: dkvs.message.DeleteResponse
	(*SetHashrangeRequest)(nil),  // 6: dkvs.message.SetHashrangeRequest
	(*SetHashrangeResponse)(nil), // 7: dkvs.message.SetHashrangeResponse
	(*StoredValue)(nil),          // 8: dkvs.message.StoredValue
}
var file_dkvs_message_message_proto_depIdxs = []int32{
	8, // 0: dkvs.message.GetResponse.stored_value:type_name -> dkvs.message.StoredValue
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_dkvs_message_message_proto_init() }
func file_dkvs_message_message_proto_init() {
	if File_dkvs_message_message_proto != nil {
		return
	}
	file_dkvs_message_store_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_dkvs_message_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRequest); i {
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
		file_dkvs_message_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetResponse); i {
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
		file_dkvs_message_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_dkvs_message_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_dkvs_message_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_dkvs_message_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
		file_dkvs_message_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetHashrangeRequest); i {
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
		file_dkvs_message_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetHashrangeResponse); i {
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
			RawDescriptor: file_dkvs_message_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dkvs_message_message_proto_goTypes,
		DependencyIndexes: file_dkvs_message_message_proto_depIdxs,
		MessageInfos:      file_dkvs_message_message_proto_msgTypes,
	}.Build()
	File_dkvs_message_message_proto = out.File
	file_dkvs_message_message_proto_rawDesc = nil
	file_dkvs_message_message_proto_goTypes = nil
	file_dkvs_message_message_proto_depIdxs = nil
}
