// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: log_out.proto

package v1

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

type Logout struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"userId,omitempty"
	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty" dynamodbav:"userId,omitempty"`
}

func (x *Logout) Reset() {
	*x = Logout{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_out_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Logout) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Logout) ProtoMessage() {}

func (x *Logout) ProtoReflect() protoreflect.Message {
	mi := &file_log_out_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Logout.ProtoReflect.Descriptor instead.
func (*Logout) Descriptor() ([]byte, []int) {
	return file_log_out_proto_rawDescGZIP(), []int{0}
}

func (x *Logout) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_log_out_proto protoreflect.FileDescriptor

var file_log_out_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6c, 0x6f, 0x67, 0x5f, 0x6f, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x76, 0x31, 0x22, 0x20, 0x0a, 0x06, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x42, 0x60, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42,
	0x0b, 0x4c, 0x6f, 0x67, 0x4f, 0x75, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x62,
	0x6c, 0x6f, 0x63, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56,
	0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_log_out_proto_rawDescOnce sync.Once
	file_log_out_proto_rawDescData = file_log_out_proto_rawDesc
)

func file_log_out_proto_rawDescGZIP() []byte {
	file_log_out_proto_rawDescOnce.Do(func() {
		file_log_out_proto_rawDescData = protoimpl.X.CompressGZIP(file_log_out_proto_rawDescData)
	})
	return file_log_out_proto_rawDescData
}

var file_log_out_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_log_out_proto_goTypes = []interface{}{
	(*Logout)(nil), // 0: v1.logout
}
var file_log_out_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_log_out_proto_init() }
func file_log_out_proto_init() {
	if File_log_out_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_log_out_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Logout); i {
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
			RawDescriptor: file_log_out_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_log_out_proto_goTypes,
		DependencyIndexes: file_log_out_proto_depIdxs,
		MessageInfos:      file_log_out_proto_msgTypes,
	}.Build()
	File_log_out_proto = out.File
	file_log_out_proto_rawDesc = nil
	file_log_out_proto_goTypes = nil
	file_log_out_proto_depIdxs = nil
}
