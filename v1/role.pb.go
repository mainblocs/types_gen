// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: role.proto

package v1

import (
	_ "github.com/mainblocs/types_gen/v1/validate"
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

type Role int32

const (
	Role_unknownRole       Role = 0
	Role_adminRole         Role = 1
	Role_userRole          Role = 2
	Role_staffRole         Role = 3
	Role_cashierRole       Role = 4
	Role_managerRole       Role = 5
	Role_systemAdminRole   Role = 6
	Role_systemStaffRole   Role = 7
	Role_systemManagerRole Role = 8
)

// Enum value maps for Role.
var (
	Role_name = map[int32]string{
		0: "unknownRole",
		1: "adminRole",
		2: "userRole",
		3: "staffRole",
		4: "cashierRole",
		5: "managerRole",
		6: "systemAdminRole",
		7: "systemStaffRole",
		8: "systemManagerRole",
	}
	Role_value = map[string]int32{
		"unknownRole":       0,
		"adminRole":         1,
		"userRole":          2,
		"staffRole":         3,
		"cashierRole":       4,
		"managerRole":       5,
		"systemAdminRole":   6,
		"systemStaffRole":   7,
		"systemManagerRole": 8,
	}
)

func (x Role) Enum() *Role {
	p := new(Role)
	*p = x
	return p
}

func (x Role) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Role) Descriptor() protoreflect.EnumDescriptor {
	return file_role_proto_enumTypes[0].Descriptor()
}

func (Role) Type() protoreflect.EnumType {
	return &file_role_proto_enumTypes[0]
}

func (x Role) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Role.Descriptor instead.
func (Role) EnumDescriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{0}
}

type UpdateUserRole struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"token,omitempty"
	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty" dynamodbav:"token,omitempty"`
	// @gotags: dynamodbav:"newRole,omitempty"
	NewRole []Role `protobuf:"varint,2,rep,packed,name=newRole,proto3,enum=v1.Role" json:"newRole,omitempty" dynamodbav:"newRole,omitempty"`
	// @gotags: dynamodbav:"staffEmail,omitempty"
	StaffEmail string `protobuf:"bytes,3,opt,name=staffEmail,proto3" json:"staffEmail,omitempty" dynamodbav:"staffEmail,omitempty"`
}

func (x *UpdateUserRole) Reset() {
	*x = UpdateUserRole{}
	if protoimpl.UnsafeEnabled {
		mi := &file_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserRole) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRole) ProtoMessage() {}

func (x *UpdateUserRole) ProtoReflect() protoreflect.Message {
	mi := &file_role_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRole.ProtoReflect.Descriptor instead.
func (*UpdateUserRole) Descriptor() ([]byte, []int) {
	return file_role_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateUserRole) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *UpdateUserRole) GetNewRole() []Role {
	if x != nil {
		return x.NewRole
	}
	return nil
}

func (x *UpdateUserRole) GetStaffEmail() string {
	if x != nil {
		return x.StaffEmail
	}
	return ""
}

var File_role_proto protoreflect.FileDescriptor

var file_role_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31,
	0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x01, 0x0a, 0x0e, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0xe9, 0xc0,
	0x03, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2e,
	0x0a, 0x07, 0x6e, 0x65, 0x77, 0x52, 0x6f, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x08, 0x2e, 0x76, 0x31, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x42, 0x0a, 0xba, 0xe9, 0xc0, 0x03, 0x05,
	0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x29,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x66, 0x66, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x09, 0xba, 0xe9, 0xc0, 0x03, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x66, 0x66, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x2a, 0xa6, 0x01, 0x0a, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x52, 0x6f, 0x6c,
	0x65, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x6f, 0x6c, 0x65,
	0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x10, 0x02,
	0x12, 0x0d, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x66, 0x66, 0x52, 0x6f, 0x6c, 0x65, 0x10, 0x03, 0x12,
	0x0f, 0x0a, 0x0b, 0x63, 0x61, 0x73, 0x68, 0x69, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x10, 0x04,
	0x12, 0x0f, 0x0a, 0x0b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x10,
	0x05, 0x12, 0x13, 0x0a, 0x0f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x52, 0x6f, 0x6c, 0x65, 0x10, 0x06, 0x12, 0x13, 0x0a, 0x0f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x53, 0x74, 0x61, 0x66, 0x66, 0x52, 0x6f, 0x6c, 0x65, 0x10, 0x07, 0x12, 0x15, 0x0a, 0x11, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65,
	0x10, 0x08, 0x42, 0x5e, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x52, 0x6f,
	0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x6d, 0x73, 0x66, 0x6f, 0x72, 0x64, 0x73, 0x2f,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56,
	0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02,
	0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_role_proto_rawDescOnce sync.Once
	file_role_proto_rawDescData = file_role_proto_rawDesc
)

func file_role_proto_rawDescGZIP() []byte {
	file_role_proto_rawDescOnce.Do(func() {
		file_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_role_proto_rawDescData)
	})
	return file_role_proto_rawDescData
}

var file_role_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_role_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_role_proto_goTypes = []interface{}{
	(Role)(0),              // 0: v1.role
	(*UpdateUserRole)(nil), // 1: v1.updateUserRole
}
var file_role_proto_depIdxs = []int32{
	0, // 0: v1.updateUserRole.newRole:type_name -> v1.role
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_role_proto_init() }
func file_role_proto_init() {
	if File_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserRole); i {
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
			RawDescriptor: file_role_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_role_proto_goTypes,
		DependencyIndexes: file_role_proto_depIdxs,
		EnumInfos:         file_role_proto_enumTypes,
		MessageInfos:      file_role_proto_msgTypes,
	}.Build()
	File_role_proto = out.File
	file_role_proto_rawDesc = nil
	file_role_proto_goTypes = nil
	file_role_proto_depIdxs = nil
}
