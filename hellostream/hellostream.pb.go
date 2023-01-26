// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: hellostream/hellostream.proto

package hellostream

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

// 请求消息
type HelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *HelloRequest) Reset() {
	*x = HelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hellostream_hellostream_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRequest) ProtoMessage() {}

func (x *HelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hellostream_hellostream_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRequest.ProtoReflect.Descriptor instead.
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return file_hellostream_hellostream_proto_rawDescGZIP(), []int{0}
}

func (x *HelloRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

// 响应消息
type HelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *HelloResponse) Reset() {
	*x = HelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hellostream_hellostream_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloResponse) ProtoMessage() {}

func (x *HelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_hellostream_hellostream_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloResponse.ProtoReflect.Descriptor instead.
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return file_hellostream_hellostream_proto_rawDescGZIP(), []int{1}
}

func (x *HelloResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

var File_hellostream_hellostream_proto protoreflect.FileDescriptor

var file_hellostream_hellostream_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2f, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x22, 0x22, 0x0a, 0x0c,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x23, 0x0a, 0x0d, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xa8, 0x02, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65,
	0x72, 0x12, 0x43, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x19, 0x2e,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x48, 0x65, 0x6c, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x19, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x46,
	0x0a, 0x09, 0x53, 0x65, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x19, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x48, 0x0a, 0x09, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x12, 0x19, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01,
	0x42, 0x1a, 0x5a, 0x18, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hellostream_hellostream_proto_rawDescOnce sync.Once
	file_hellostream_hellostream_proto_rawDescData = file_hellostream_hellostream_proto_rawDesc
)

func file_hellostream_hellostream_proto_rawDescGZIP() []byte {
	file_hellostream_hellostream_proto_rawDescOnce.Do(func() {
		file_hellostream_hellostream_proto_rawDescData = protoimpl.X.CompressGZIP(file_hellostream_hellostream_proto_rawDescData)
	})
	return file_hellostream_hellostream_proto_rawDescData
}

var file_hellostream_hellostream_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_hellostream_hellostream_proto_goTypes = []interface{}{
	(*HelloRequest)(nil),  // 0: hellostream.HelloRequest
	(*HelloResponse)(nil), // 1: hellostream.HelloResponse
}
var file_hellostream_hellostream_proto_depIdxs = []int32{
	0, // 0: hellostream.Greeter.SayHello:input_type -> hellostream.HelloRequest
	0, // 1: hellostream.Greeter.GetStream:input_type -> hellostream.HelloRequest
	0, // 2: hellostream.Greeter.SetStream:input_type -> hellostream.HelloRequest
	0, // 3: hellostream.Greeter.AllStream:input_type -> hellostream.HelloRequest
	1, // 4: hellostream.Greeter.SayHello:output_type -> hellostream.HelloResponse
	1, // 5: hellostream.Greeter.GetStream:output_type -> hellostream.HelloResponse
	1, // 6: hellostream.Greeter.SetStream:output_type -> hellostream.HelloResponse
	1, // 7: hellostream.Greeter.AllStream:output_type -> hellostream.HelloResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hellostream_hellostream_proto_init() }
func file_hellostream_hellostream_proto_init() {
	if File_hellostream_hellostream_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hellostream_hellostream_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRequest); i {
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
		file_hellostream_hellostream_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloResponse); i {
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
			RawDescriptor: file_hellostream_hellostream_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hellostream_hellostream_proto_goTypes,
		DependencyIndexes: file_hellostream_hellostream_proto_depIdxs,
		MessageInfos:      file_hellostream_hellostream_proto_msgTypes,
	}.Build()
	File_hellostream_hellostream_proto = out.File
	file_hellostream_hellostream_proto_rawDesc = nil
	file_hellostream_hellostream_proto_goTypes = nil
	file_hellostream_hellostream_proto_depIdxs = nil
}
