// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.1
// source: getmodel/getmodel.proto

package grpc

import (
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

// The request message containing the user's name.
type HashRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *HashRequest) Reset() {
	*x = HashRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_getmodel_getmodel_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashRequest) ProtoMessage() {}

func (x *HashRequest) ProtoReflect() protoreflect.Message {
	mi := &file_getmodel_getmodel_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashRequest.ProtoReflect.Descriptor instead.
func (*HashRequest) Descriptor() ([]byte, []int) {
	return file_getmodel_getmodel_proto_rawDescGZIP(), []int{0}
}

// The response message containing the greetings
type HashReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash string `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (x *HashReply) Reset() {
	*x = HashReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_getmodel_getmodel_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HashReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HashReply) ProtoMessage() {}

func (x *HashReply) ProtoReflect() protoreflect.Message {
	mi := &file_getmodel_getmodel_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HashReply.ProtoReflect.Descriptor instead.
func (*HashReply) Descriptor() ([]byte, []int) {
	return file_getmodel_getmodel_proto_rawDescGZIP(), []int{1}
}

func (x *HashReply) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

var File_getmodel_getmodel_proto protoreflect.FileDescriptor

var file_getmodel_getmodel_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x65, 0x74, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x67, 0x65, 0x74, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x66, 0x65, 0x74, 0x63, 0x68,
	0x48, 0x61, 0x73, 0x68, 0x22, 0x0d, 0x0a, 0x0b, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x1f, 0x0a, 0x09, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x32, 0x44, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12,
	0x39, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x2e, 0x66, 0x65, 0x74,
	0x63, 0x68, 0x48, 0x61, 0x73, 0x68, 0x2e, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x66, 0x65, 0x74, 0x63, 0x68, 0x48, 0x61, 0x73, 0x68, 0x2e, 0x48,
	0x61, 0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x30, 0x0a, 0x07, 0x69, 0x6f,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x42, 0x0d, 0x47, 0x65, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x14, 0x69, 0x70, 0x66, 0x73, 0x2d, 0x66, 0x65, 0x74,
	0x63, 0x68, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_getmodel_getmodel_proto_rawDescOnce sync.Once
	file_getmodel_getmodel_proto_rawDescData = file_getmodel_getmodel_proto_rawDesc
)

func file_getmodel_getmodel_proto_rawDescGZIP() []byte {
	file_getmodel_getmodel_proto_rawDescOnce.Do(func() {
		file_getmodel_getmodel_proto_rawDescData = protoimpl.X.CompressGZIP(file_getmodel_getmodel_proto_rawDescData)
	})
	return file_getmodel_getmodel_proto_rawDescData
}

var file_getmodel_getmodel_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_getmodel_getmodel_proto_goTypes = []interface{}{
	(*HashRequest)(nil), // 0: fetchHash.HashRequest
	(*HashReply)(nil),   // 1: fetchHash.HashReply
}
var file_getmodel_getmodel_proto_depIdxs = []int32{
	0, // 0: fetchHash.Greeter.GetHash:input_type -> fetchHash.HashRequest
	1, // 1: fetchHash.Greeter.GetHash:output_type -> fetchHash.HashReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_getmodel_getmodel_proto_init() }
func file_getmodel_getmodel_proto_init() {
	if File_getmodel_getmodel_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_getmodel_getmodel_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashRequest); i {
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
		file_getmodel_getmodel_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HashReply); i {
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
			RawDescriptor: file_getmodel_getmodel_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_getmodel_getmodel_proto_goTypes,
		DependencyIndexes: file_getmodel_getmodel_proto_depIdxs,
		MessageInfos:      file_getmodel_getmodel_proto_msgTypes,
	}.Build()
	File_getmodel_getmodel_proto = out.File
	file_getmodel_getmodel_proto_rawDesc = nil
	file_getmodel_getmodel_proto_goTypes = nil
	file_getmodel_getmodel_proto_depIdxs = nil
}
