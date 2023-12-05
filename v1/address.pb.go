// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: address.proto

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

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"addressLine1,omitempty"
	AddressLine1 string `protobuf:"bytes,1,opt,name=addressLine1,proto3" json:"addressLine1,omitempty" dynamodbav:"addressLine1,omitempty"`
	// @gotags: dynamodbav:"addressLine2,omitempty"
	AddressLine2 string `protobuf:"bytes,2,opt,name=addressLine2,proto3" json:"addressLine2,omitempty" dynamodbav:"addressLine2,omitempty"`
	// @gotags: dynamodbav:"streetName,omitempty"
	StreetName string `protobuf:"bytes,3,opt,name=streetName,proto3" json:"streetName,omitempty" dynamodbav:"streetName,omitempty"`
	// @gotags: dynamodbav:"city,omitempty"
	City string `protobuf:"bytes,4,opt,name=city,proto3" json:"city,omitempty" dynamodbav:"city,omitempty"`
	// @gotags: dynamodbav:"county,omitempty"
	County string `protobuf:"bytes,5,opt,name=county,proto3" json:"county,omitempty" dynamodbav:"county,omitempty"`
	// @gotags: dynamodbav:"zipCode,omitempty"
	ZipCode string `protobuf:"bytes,6,opt,name=zipCode,proto3" json:"zipCode,omitempty" dynamodbav:"zipCode,omitempty"`
	// @gotags: dynamodbav:"state,omitempty"
	State string `protobuf:"bytes,7,opt,name=state,proto3" json:"state,omitempty" dynamodbav:"state,omitempty"`
	// @gotags: dynamodbav:"stateCode,omitempty"
	StateCode string `protobuf:"bytes,8,opt,name=stateCode,proto3" json:"stateCode,omitempty" dynamodbav:"stateCode,omitempty"`
	// @gotags: dynamodbav:"extendedZipCode,omitempty"
	ExtendedZipCode string `protobuf:"bytes,9,opt,name=extendedZipCode,proto3" json:"extendedZipCode,omitempty" dynamodbav:"extendedZipCode,omitempty"`
	// @gotags: dynamodbav:"country,omitempty"
	Country string `protobuf:"bytes,10,opt,name=country,proto3" json:"country,omitempty" dynamodbav:"country,omitempty"`
	// @gotags: dynamodbav:"countryCode,omitempty"
	CountryCode string `protobuf:"bytes,11,opt,name=countryCode,proto3" json:"countryCode,omitempty" dynamodbav:"countryCode,omitempty"`
	// @gotags: dynamodbav:"freeFormAddress,omitempty"
	FreeFormAddress string `protobuf:"bytes,12,opt,name=freeFormAddress,proto3" json:"freeFormAddress,omitempty" dynamodbav:"freeFormAddress,omitempty"`
	// @gotags: dynamodbav:"localName,omitempty"
	LocalName string `protobuf:"bytes,13,opt,name=localName,proto3" json:"localName,omitempty" dynamodbav:"localName,omitempty"`
	// @gotags: dynamodbav:"lat,omitempty"
	Lat float32 `protobuf:"fixed32,14,opt,name=lat,proto3" json:"lat,omitempty" dynamodbav:"lat,omitempty"`
	// @gotags: dynamodbav:"long,omitempty"
	Long float32 `protobuf:"fixed32,15,opt,name=long,proto3" json:"long,omitempty" dynamodbav:"long,omitempty"`
	// @gotags: dynamodbav:"businessId,omitempty"
	BusinessId string `protobuf:"bytes,16,opt,name=businessId,proto3" json:"businessId,omitempty" dynamodbav:"businessId,omitempty"`
	// @gotags: dynamodbav:"type,omitempty"
	Type string `protobuf:"bytes,17,opt,name=type,proto3" json:"type,omitempty" dynamodbav:"type,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_address_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_address_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_address_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetAddressLine1() string {
	if x != nil {
		return x.AddressLine1
	}
	return ""
}

func (x *Address) GetAddressLine2() string {
	if x != nil {
		return x.AddressLine2
	}
	return ""
}

func (x *Address) GetStreetName() string {
	if x != nil {
		return x.StreetName
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetCounty() string {
	if x != nil {
		return x.County
	}
	return ""
}

func (x *Address) GetZipCode() string {
	if x != nil {
		return x.ZipCode
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetStateCode() string {
	if x != nil {
		return x.StateCode
	}
	return ""
}

func (x *Address) GetExtendedZipCode() string {
	if x != nil {
		return x.ExtendedZipCode
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetCountryCode() string {
	if x != nil {
		return x.CountryCode
	}
	return ""
}

func (x *Address) GetFreeFormAddress() string {
	if x != nil {
		return x.FreeFormAddress
	}
	return ""
}

func (x *Address) GetLocalName() string {
	if x != nil {
		return x.LocalName
	}
	return ""
}

func (x *Address) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Address) GetLong() float32 {
	if x != nil {
		return x.Long
	}
	return 0
}

func (x *Address) GetBusinessId() string {
	if x != nil {
		return x.BusinessId
	}
	return ""
}

func (x *Address) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

var File_address_proto protoreflect.FileDescriptor

var file_address_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x76, 0x31, 0x22, 0xf3, 0x03, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x22, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x31, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69,
	0x6e, 0x65, 0x31, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x4c, 0x69,
	0x6e, 0x65, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x4c, 0x69, 0x6e, 0x65, 0x32, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x65, 0x65,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x72,
	0x65, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x7a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x28, 0x0a, 0x0f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5a, 0x69, 0x70,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x64, 0x5a, 0x69, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x66, 0x72, 0x65, 0x65, 0x46,
	0x6f, 0x72, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x66, 0x72, 0x65, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73,
	0x73, 0x49, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x11, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x61, 0x0a, 0x06, 0x63, 0x6f, 0x6d,
	0x2e, 0x76, 0x31, 0x42, 0x0c, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x72, 0x61, 0x6d, 0x73, 0x66, 0x6f, 0x72, 0x64, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x5f,
	0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56,
	0x31, 0xca, 0x02, 0x02, 0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_address_proto_rawDescOnce sync.Once
	file_address_proto_rawDescData = file_address_proto_rawDesc
)

func file_address_proto_rawDescGZIP() []byte {
	file_address_proto_rawDescOnce.Do(func() {
		file_address_proto_rawDescData = protoimpl.X.CompressGZIP(file_address_proto_rawDescData)
	})
	return file_address_proto_rawDescData
}

var file_address_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_address_proto_goTypes = []interface{}{
	(*Address)(nil), // 0: v1.address
}
var file_address_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_address_proto_init() }
func file_address_proto_init() {
	if File_address_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_address_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
			RawDescriptor: file_address_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_address_proto_goTypes,
		DependencyIndexes: file_address_proto_depIdxs,
		MessageInfos:      file_address_proto_msgTypes,
	}.Build()
	File_address_proto = out.File
	file_address_proto_rawDesc = nil
	file_address_proto_goTypes = nil
	file_address_proto_depIdxs = nil
}
