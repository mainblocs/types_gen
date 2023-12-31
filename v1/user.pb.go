// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: user.proto

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"firstName,omitempty"
	FirstName string `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty" dynamodbav:"firstName,omitempty"`
	// @gotags: dynamodbav:"lastName,omitempty"
	LastName string `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty" dynamodbav:"lastName,omitempty"`
	// @gotags: dynamodbav:"userName,omitempty"
	UserName string `protobuf:"bytes,3,opt,name=userName,proto3" json:"userName,omitempty" dynamodbav:"userName,omitempty"`
	// @gotags: dynamodbav:"password,omitempty"
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty" dynamodbav:"password,omitempty"`
	// @gotags: dynamodbav:"confirmPassword,omitempty"
	ConfirmPassword string `protobuf:"bytes,5,opt,name=confirmPassword,proto3" json:"confirmPassword,omitempty" dynamodbav:"confirmPassword,omitempty"`
	// @gotags: dynamodbav:"origin,omitempty"
	Origin string `protobuf:"bytes,6,opt,name=origin,proto3" json:"origin,omitempty" dynamodbav:"origin,omitempty"`
	// @gotags: dynamodbav:"emailVisibility,omitempty"
	EmailVisibility bool `protobuf:"varint,7,opt,name=emailVisibility,proto3" json:"emailVisibility,omitempty" dynamodbav:"emailVisibility,omitempty"`
	// @gotags: dynamodbav:"type,omitempty"
	Type string `protobuf:"bytes,8,opt,name=type,proto3" json:"type,omitempty" dynamodbav:"type,omitempty"`
	// @gotags: dynamodbav:"id,omitempty"
	Id string `protobuf:"bytes,9,opt,name=id,proto3" json:"id,omitempty" dynamodbav:"id,omitempty"`
	// @gotags: dynamodbav:"created,omitempty"
	Created string `protobuf:"bytes,10,opt,name=created,proto3" json:"created,omitempty" dynamodbav:"created,omitempty"`
	// @gotags: dynamodbav:"updated,omitempty"
	Updated string `protobuf:"bytes,11,opt,name=updated,proto3" json:"updated,omitempty" dynamodbav:"updated,omitempty"`
	// @gotags: dynamodbav:"verified,omitempty"
	Verified bool `protobuf:"varint,12,opt,name=verified,proto3" json:"verified,omitempty" dynamodbav:"verified,omitempty"`
	// @gotags: dynamodbav:"avatar,omitempty"
	Avatar string `protobuf:"bytes,13,opt,name=avatar,proto3" json:"avatar,omitempty" dynamodbav:"avatar,omitempty"`
	// @gotags: dynamodbav:"lastResetSentAt,omitempty"
	LastResetSentAt string `protobuf:"bytes,14,opt,name=lastResetSentAt,proto3" json:"lastResetSentAt,omitempty" dynamodbav:"lastResetSentAt,omitempty"`
	// @gotags: dynamodbav:"lastVerificationSentAt,omitempty"
	LastVerificationSentAt string `protobuf:"bytes,15,opt,name=lastVerificationSentAt,proto3" json:"lastVerificationSentAt,omitempty" dynamodbav:"lastVerificationSentAt,omitempty"`
	// @gotags: dynamodbav:"passwordHash,omitempty"
	PasswordHash string `protobuf:"bytes,16,opt,name=passwordHash,proto3" json:"passwordHash,omitempty" dynamodbav:"passwordHash,omitempty"`
	// @gotags: dynamodbav:"tokenKey,omitempty"
	TokenKey string `protobuf:"bytes,17,opt,name=tokenKey,proto3" json:"tokenKey,omitempty" dynamodbav:"tokenKey,omitempty"`
	// @gotags: dynamodbav:"token,omitempty"
	Token string `protobuf:"bytes,18,opt,name=token,proto3" json:"token,omitempty" dynamodbav:"token,omitempty"`
	// @gotags: dynamodbav:"email,omitempty"
	Email string `protobuf:"bytes,19,opt,name=email,proto3" json:"email,omitempty" dynamodbav:"email,omitempty"`
	// @gotags: dynamodbav:"phoneNumber,omitempty"
	PhoneNumber string `protobuf:"bytes,20,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty" dynamodbav:"phoneNumber,omitempty"`
	// @gotags: dynamodbav:"organizationId,omitempty"
	OrganizationId string `protobuf:"bytes,21,opt,name=organizationId,proto3" json:"organizationId,omitempty" dynamodbav:"organizationId,omitempty"`
	// @gotags: dynamodbav:"emailConfirmed,omitempty"
	EmailConfirmed bool `protobuf:"varint,22,opt,name=emailConfirmed,proto3" json:"emailConfirmed,omitempty" dynamodbav:"emailConfirmed,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetConfirmPassword() string {
	if x != nil {
		return x.ConfirmPassword
	}
	return ""
}

func (x *User) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *User) GetEmailVisibility() bool {
	if x != nil {
		return x.EmailVisibility
	}
	return false
}

func (x *User) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetCreated() string {
	if x != nil {
		return x.Created
	}
	return ""
}

func (x *User) GetUpdated() string {
	if x != nil {
		return x.Updated
	}
	return ""
}

func (x *User) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetLastResetSentAt() string {
	if x != nil {
		return x.LastResetSentAt
	}
	return ""
}

func (x *User) GetLastVerificationSentAt() string {
	if x != nil {
		return x.LastVerificationSentAt
	}
	return ""
}

func (x *User) GetPasswordHash() string {
	if x != nil {
		return x.PasswordHash
	}
	return ""
}

func (x *User) GetTokenKey() string {
	if x != nil {
		return x.TokenKey
	}
	return ""
}

func (x *User) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *User) GetOrganizationId() string {
	if x != nil {
		return x.OrganizationId
	}
	return ""
}

func (x *User) GetEmailConfirmed() bool {
	if x != nil {
		return x.EmailConfirmed
	}
	return false
}

type LoginUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginUser) Reset() {
	*x = LoginUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUser) ProtoMessage() {}

func (x *LoginUser) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUser.ProtoReflect.Descriptor instead.
func (*LoginUser) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *LoginUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginUser) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type PocketCollection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *PocketCollection) Reset() {
	*x = PocketCollection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PocketCollection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PocketCollection) ProtoMessage() {}

func (x *PocketCollection) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PocketCollection.ProtoReflect.Descriptor instead.
func (*PocketCollection) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *PocketCollection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PocketCollection) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type FrontEndUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"name,omitempty"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" dynamodbav:"name,omitempty"`
	// @gotags: dynamodbav:"userName,omitempty"
	UserName string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty" dynamodbav:"userName,omitempty"`
	// @gotags: dynamodbav:"origin,omitempty"
	Origin string `protobuf:"bytes,3,opt,name=origin,proto3" json:"origin,omitempty" dynamodbav:"origin,omitempty"`
	// @gotags: dynamodbav:"type,omitempty"
	Type string `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty" dynamodbav:"type,omitempty"`
	// @gotags: dynamodbav:"id,omitempty"
	Id string `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty" dynamodbav:"id,omitempty"`
	// @gotags: dynamodbav:"verified,omitempty"
	Verified bool `protobuf:"varint,6,opt,name=verified,proto3" json:"verified,omitempty" dynamodbav:"verified,omitempty"`
	// @gotags: dynamodbav:"avatar,omitempty"
	Avatar string `protobuf:"bytes,7,opt,name=avatar,proto3" json:"avatar,omitempty" dynamodbav:"avatar,omitempty"`
	// @gotags: dynamodbav:"tokenKey,omitempty"
	TokenKey string `protobuf:"bytes,8,opt,name=tokenKey,proto3" json:"tokenKey,omitempty" dynamodbav:"tokenKey,omitempty"`
	// @gotags: dynamodbav:"token,omitempty"
	Token string `protobuf:"bytes,9,opt,name=token,proto3" json:"token,omitempty" dynamodbav:"token,omitempty"`
	// @gotags: dynamodbav:"email,omitempty"
	Email string `protobuf:"bytes,10,opt,name=email,proto3" json:"email,omitempty" dynamodbav:"email,omitempty"`
}

func (x *FrontEndUser) Reset() {
	*x = FrontEndUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FrontEndUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FrontEndUser) ProtoMessage() {}

func (x *FrontEndUser) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FrontEndUser.ProtoReflect.Descriptor instead.
func (*FrontEndUser) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *FrontEndUser) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FrontEndUser) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *FrontEndUser) GetOrigin() string {
	if x != nil {
		return x.Origin
	}
	return ""
}

func (x *FrontEndUser) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *FrontEndUser) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FrontEndUser) GetVerified() bool {
	if x != nil {
		return x.Verified
	}
	return false
}

func (x *FrontEndUser) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *FrontEndUser) GetTokenKey() string {
	if x != nil {
		return x.TokenKey
	}
	return ""
}

func (x *FrontEndUser) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *FrontEndUser) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_user_proto protoreflect.FileDescriptor

var file_user_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31,
	0x1a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x6c, 0x6f,
	0x63, 0x6b, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x05, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73,
	0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x56, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x28, 0x0a,
	0x0f, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x65, 0x74, 0x53, 0x65, 0x6e, 0x74, 0x41, 0x74,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x53, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x12, 0x36, 0x0a, 0x16, 0x6c, 0x61, 0x73, 0x74, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6e, 0x74, 0x41,
	0x74, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x6c, 0x61, 0x73, 0x74, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x6e, 0x74, 0x41, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x48, 0x61, 0x73, 0x68, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x48,
	0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4b, 0x65, 0x79, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x13,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x26, 0x0a,
	0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x18,
	0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x22, 0x3d, 0x0a,
	0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x3a, 0x0a, 0x10,
	0x70, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xf6, 0x01, 0x0a, 0x0c, 0x66, 0x72, 0x6f,
	0x6e, 0x74, 0x45, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x72, 0x69,
	0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x72, 0x69, 0x67, 0x69,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x4b, 0x65, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x4b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x42, 0x5e, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x6c, 0x6f, 0x63, 0x73, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58,
	0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData = file_user_proto_rawDesc
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_rawDescData)
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_user_proto_goTypes = []interface{}{
	(*User)(nil),             // 0: v1.user
	(*LoginUser)(nil),        // 1: v1.LoginUser
	(*PocketCollection)(nil), // 2: v1.pocketCollection
	(*FrontEndUser)(nil),     // 3: v1.frontEndUser
}
var file_user_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	file_role_proto_init()
	file_password_proto_init()
	file_locked_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginUser); i {
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
		file_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PocketCollection); i {
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
		file_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FrontEndUser); i {
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
			RawDescriptor: file_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_rawDesc = nil
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
