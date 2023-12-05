// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: confirm_email.proto

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

type ConfirmEmail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"email,omitempty"
	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty" dynamodbav:"email,omitempty"`
	// @gotags: dynamodbav:"userName,omitempty"
	UserName string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty" dynamodbav:"userName,omitempty"`
	// @gotags: dynamodbav:"confirmationCode,omitempty"
	ConfirmationCode string `protobuf:"bytes,3,opt,name=confirmationCode,proto3" json:"confirmationCode,omitempty" dynamodbav:"confirmationCode,omitempty"`
	// @gotags: dynamodbav:"token,omitempty"
	Token string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty" dynamodbav:"token,omitempty"`
	// @gotags: dynamodbav:"businessId,omitempty"
	BusinessId string `protobuf:"bytes,5,opt,name=businessId,proto3" json:"businessId,omitempty" dynamodbav:"businessId,omitempty"`
}

func (x *ConfirmEmail) Reset() {
	*x = ConfirmEmail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_confirm_email_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfirmEmail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfirmEmail) ProtoMessage() {}

func (x *ConfirmEmail) ProtoReflect() protoreflect.Message {
	mi := &file_confirm_email_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfirmEmail.ProtoReflect.Descriptor instead.
func (*ConfirmEmail) Descriptor() ([]byte, []int) {
	return file_confirm_email_proto_rawDescGZIP(), []int{0}
}

func (x *ConfirmEmail) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *ConfirmEmail) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *ConfirmEmail) GetConfirmationCode() string {
	if x != nil {
		return x.ConfirmationCode
	}
	return ""
}

func (x *ConfirmEmail) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ConfirmEmail) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

var File_confirm_email_proto protoreflect.FileDescriptor

var file_confirm_email_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x22, 0xa2, 0x01, 0x0a, 0x0c, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x49, 0x64, 0x42, 0x66,
	0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72,
	0x6d, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6d, 0x73, 0x66, 0x6f,
	0x72, 0x64, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31,
	0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_confirm_email_proto_rawDescOnce sync.Once
	file_confirm_email_proto_rawDescData = file_confirm_email_proto_rawDesc
)

func file_confirm_email_proto_rawDescGZIP() []byte {
	file_confirm_email_proto_rawDescOnce.Do(func() {
		file_confirm_email_proto_rawDescData = protoimpl.X.CompressGZIP(file_confirm_email_proto_rawDescData)
	})
	return file_confirm_email_proto_rawDescData
}

var file_confirm_email_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_confirm_email_proto_goTypes = []interface{}{
	(*ConfirmEmail)(nil), // 0: v1.confirmEmail
}
var file_confirm_email_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_confirm_email_proto_init() }
func file_confirm_email_proto_init() {
	if File_confirm_email_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_confirm_email_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfirmEmail); i {
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
			RawDescriptor: file_confirm_email_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_confirm_email_proto_goTypes,
		DependencyIndexes: file_confirm_email_proto_depIdxs,
		MessageInfos:      file_confirm_email_proto_msgTypes,
	}.Build()
	File_confirm_email_proto = out.File
	file_confirm_email_proto_rawDesc = nil
	file_confirm_email_proto_goTypes = nil
	file_confirm_email_proto_depIdxs = nil
}
