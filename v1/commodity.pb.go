// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: commodity.proto

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

type CommodityService int32

const (
	CommodityService_unknownCommodityService CommodityService = 0
	CommodityService_protectFromFreeze       CommodityService = 1
	CommodityService_sortAndSegregate        CommodityService = 2
	CommodityService_guaranteed              CommodityService = 3
	CommodityService_hazardous               CommodityService = 4
	CommodityService_stackable               CommodityService = 5
)

// Enum value maps for CommodityService.
var (
	CommodityService_name = map[int32]string{
		0: "unknownCommodityService",
		1: "protectFromFreeze",
		2: "sortAndSegregate",
		3: "guaranteed",
		4: "hazardous",
		5: "stackable",
	}
	CommodityService_value = map[string]int32{
		"unknownCommodityService": 0,
		"protectFromFreeze":       1,
		"sortAndSegregate":        2,
		"guaranteed":              3,
		"hazardous":               4,
		"stackable":               5,
	}
)

func (x CommodityService) Enum() *CommodityService {
	p := new(CommodityService)
	*p = x
	return p
}

func (x CommodityService) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommodityService) Descriptor() protoreflect.EnumDescriptor {
	return file_commodity_proto_enumTypes[0].Descriptor()
}

func (CommodityService) Type() protoreflect.EnumType {
	return &file_commodity_proto_enumTypes[0]
}

func (x CommodityService) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommodityService.Descriptor instead.
func (CommodityService) EnumDescriptor() ([]byte, []int) {
	return file_commodity_proto_rawDescGZIP(), []int{0}
}

type PackageType int32

const (
	PackageType_unknownPackageType PackageType = 0
	PackageType_pallets48x40       PackageType = 1
	PackageType_pallets48x48       PackageType = 2
	PackageType_bags               PackageType = 3
	PackageType_bales              PackageType = 4
	PackageType_boxes              PackageType = 5
	PackageType_buckets            PackageType = 6
	PackageType_bundles            PackageType = 7
	PackageType_cans               PackageType = 8
	PackageType_cartons            PackageType = 9
	PackageType_cases              PackageType = 10
	PackageType_coils              PackageType = 11
	PackageType_crates             PackageType = 12
	PackageType_cylinders          PackageType = 13
	PackageType_drums              PackageType = 14
	PackageType_pails              PackageType = 15
	PackageType_reels              PackageType = 16
	PackageType_rolls              PackageType = 17
	PackageType_totes              PackageType = 18
	PackageType_tubes              PackageType = 19
	PackageType_units              PackageType = 20
)

// Enum value maps for PackageType.
var (
	PackageType_name = map[int32]string{
		0:  "unknownPackageType",
		1:  "pallets48x40",
		2:  "pallets48x48",
		3:  "bags",
		4:  "bales",
		5:  "boxes",
		6:  "buckets",
		7:  "bundles",
		8:  "cans",
		9:  "cartons",
		10: "cases",
		11: "coils",
		12: "crates",
		13: "cylinders",
		14: "drums",
		15: "pails",
		16: "reels",
		17: "rolls",
		18: "totes",
		19: "tubes",
		20: "units",
	}
	PackageType_value = map[string]int32{
		"unknownPackageType": 0,
		"pallets48x40":       1,
		"pallets48x48":       2,
		"bags":               3,
		"bales":              4,
		"boxes":              5,
		"buckets":            6,
		"bundles":            7,
		"cans":               8,
		"cartons":            9,
		"cases":              10,
		"coils":              11,
		"crates":             12,
		"cylinders":          13,
		"drums":              14,
		"pails":              15,
		"reels":              16,
		"rolls":              17,
		"totes":              18,
		"tubes":              19,
		"units":              20,
	}
)

func (x PackageType) Enum() *PackageType {
	p := new(PackageType)
	*p = x
	return p
}

func (x PackageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PackageType) Descriptor() protoreflect.EnumDescriptor {
	return file_commodity_proto_enumTypes[1].Descriptor()
}

func (PackageType) Type() protoreflect.EnumType {
	return &file_commodity_proto_enumTypes[1]
}

func (x PackageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PackageType.Descriptor instead.
func (PackageType) EnumDescriptor() ([]byte, []int) {
	return file_commodity_proto_rawDescGZIP(), []int{1}
}

type FreightClass int32

const (
	FreightClass_unknownFreightClass FreightClass = 0
	FreightClass_class50             FreightClass = 1
	FreightClass_class55             FreightClass = 2
	FreightClass_class60             FreightClass = 3
	FreightClass_class65             FreightClass = 4
	FreightClass_class70             FreightClass = 5
	FreightClass_class775            FreightClass = 6
	FreightClass_class85             FreightClass = 7
	FreightClass_class925            FreightClass = 8
	FreightClass_class100            FreightClass = 9
	FreightClass_class110            FreightClass = 10
	FreightClass_class125            FreightClass = 11
	FreightClass_class150            FreightClass = 12
	FreightClass_class175            FreightClass = 13
	FreightClass_class200            FreightClass = 14
	FreightClass_class250            FreightClass = 15
	FreightClass_class300            FreightClass = 16
	FreightClass_class400            FreightClass = 17
	FreightClass_class500            FreightClass = 18
)

// Enum value maps for FreightClass.
var (
	FreightClass_name = map[int32]string{
		0:  "unknownFreightClass",
		1:  "class50",
		2:  "class55",
		3:  "class60",
		4:  "class65",
		5:  "class70",
		6:  "class775",
		7:  "class85",
		8:  "class925",
		9:  "class100",
		10: "class110",
		11: "class125",
		12: "class150",
		13: "class175",
		14: "class200",
		15: "class250",
		16: "class300",
		17: "class400",
		18: "class500",
	}
	FreightClass_value = map[string]int32{
		"unknownFreightClass": 0,
		"class50":             1,
		"class55":             2,
		"class60":             3,
		"class65":             4,
		"class70":             5,
		"class775":            6,
		"class85":             7,
		"class925":            8,
		"class100":            9,
		"class110":            10,
		"class125":            11,
		"class150":            12,
		"class175":            13,
		"class200":            14,
		"class250":            15,
		"class300":            16,
		"class400":            17,
		"class500":            18,
	}
)

func (x FreightClass) Enum() *FreightClass {
	p := new(FreightClass)
	*p = x
	return p
}

func (x FreightClass) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FreightClass) Descriptor() protoreflect.EnumDescriptor {
	return file_commodity_proto_enumTypes[2].Descriptor()
}

func (FreightClass) Type() protoreflect.EnumType {
	return &file_commodity_proto_enumTypes[2]
}

func (x FreightClass) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FreightClass.Descriptor instead.
func (FreightClass) EnumDescriptor() ([]byte, []int) {
	return file_commodity_proto_rawDescGZIP(), []int{2}
}

type DimensionUOM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"INCH,omitempty"
	INCH bool `protobuf:"varint,1,opt,name=INCH,proto3" json:"INCH,omitempty" dynamodbav:"INCH,omitempty"`
	// @gotags: dynamodbav:"CM,omitempty"
	CM bool `protobuf:"varint,2,opt,name=CM,proto3" json:"CM,omitempty" dynamodbav:"CM,omitempty"`
}

func (x *DimensionUOM) Reset() {
	*x = DimensionUOM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commodity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DimensionUOM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DimensionUOM) ProtoMessage() {}

func (x *DimensionUOM) ProtoReflect() protoreflect.Message {
	mi := &file_commodity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DimensionUOM.ProtoReflect.Descriptor instead.
func (*DimensionUOM) Descriptor() ([]byte, []int) {
	return file_commodity_proto_rawDescGZIP(), []int{0}
}

func (x *DimensionUOM) GetINCH() bool {
	if x != nil {
		return x.INCH
	}
	return false
}

func (x *DimensionUOM) GetCM() bool {
	if x != nil {
		return x.CM
	}
	return false
}

type WeightUOM struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"LB,omitempty"
	LB bool `protobuf:"varint,1,opt,name=LB,proto3" json:"LB,omitempty" dynamodbav:"LB,omitempty"`
	// @gotags: dynamodbav:"KG,omitempty"
	KG bool `protobuf:"varint,2,opt,name=KG,proto3" json:"KG,omitempty" dynamodbav:"KG,omitempty"`
}

func (x *WeightUOM) Reset() {
	*x = WeightUOM{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commodity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeightUOM) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeightUOM) ProtoMessage() {}

func (x *WeightUOM) ProtoReflect() protoreflect.Message {
	mi := &file_commodity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeightUOM.ProtoReflect.Descriptor instead.
func (*WeightUOM) Descriptor() ([]byte, []int) {
	return file_commodity_proto_rawDescGZIP(), []int{1}
}

func (x *WeightUOM) GetLB() bool {
	if x != nil {
		return x.LB
	}
	return false
}

func (x *WeightUOM) GetKG() bool {
	if x != nil {
		return x.KG
	}
	return false
}

type Commodity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: dynamodbav:"density,omitempty"
	Density float32 `protobuf:"fixed32,1,opt,name=density,proto3" json:"density,omitempty" dynamodbav:"density,omitempty"`
	// @gotags: dynamodbav:"length,omitempty"
	Length float32 `protobuf:"fixed32,2,opt,name=length,proto3" json:"length,omitempty" dynamodbav:"length,omitempty"`
	// @gotags: dynamodbav:"width,omitempty"
	Width float32 `protobuf:"fixed32,3,opt,name=width,proto3" json:"width,omitempty" dynamodbav:"width,omitempty"`
	// @gotags: dynamodbav:"height,omitempty"
	Height float32 `protobuf:"fixed32,4,opt,name=height,proto3" json:"height,omitempty" dynamodbav:"height,omitempty"`
	// @gotags: dynamodbav:"weight,omitempty"
	Weight float32 `protobuf:"fixed32,5,opt,name=weight,proto3" json:"weight,omitempty" dynamodbav:"weight,omitempty"`
	// @gotags: dynamodbav:"dimensionUOM,omitempty"
	DimensionUOM *DimensionUOM `protobuf:"bytes,6,opt,name=dimensionUOM,proto3" json:"dimensionUOM,omitempty" dynamodbav:"dimensionUOM,omitempty"`
	// @gotags: dynamodbav:"weightUOM,omitempty"
	WeightUOM *WeightUOM `protobuf:"bytes,7,opt,name=weightUOM,proto3" json:"weightUOM,omitempty" dynamodbav:"weightUOM,omitempty"`
	// @gotags: dynamodbav:"dimensionDisplay,omitempty"
	DimensionDisplay string `protobuf:"bytes,8,opt,name=dimensionDisplay,proto3" json:"dimensionDisplay,omitempty" dynamodbav:"dimensionDisplay,omitempty"`
	// @gotags: dynamodbav:"packageType,omitempty"
	PackageType PackageType `protobuf:"varint,9,opt,name=packageType,proto3,enum=v1.PackageType" json:"packageType,omitempty" dynamodbav:"packageType,omitempty"`
	// @gotags: dynamodbav:"quantity,omitempty"
	Quantity int32 `protobuf:"varint,10,opt,name=quantity,proto3" json:"quantity,omitempty" dynamodbav:"quantity,omitempty"`
	// @gotags: dynamodbav:"freightClass,omitempty"
	FreightClass FreightClass `protobuf:"varint,11,opt,name=freightClass,proto3,enum=v1.FreightClass" json:"freightClass,omitempty" dynamodbav:"freightClass,omitempty"`
	// @gotags: dynamodbav:"instructions,omitempty"
	Instructions string `protobuf:"bytes,12,opt,name=instructions,proto3" json:"instructions,omitempty" dynamodbav:"instructions,omitempty"`
	// @gotags: dynamodbav:"index,omitempty"
	Index int32 `protobuf:"varint,14,opt,name=index,proto3" json:"index,omitempty" dynamodbav:"index,omitempty"`
	// @gotags: dynamodbav:"description,omitempty"
	Description string `protobuf:"bytes,15,opt,name=description,proto3" json:"description,omitempty" dynamodbav:"description,omitempty"`
	// @gotags: dynamodbav:"commodityServices,omitempty"
	CommodityServices []CommodityService `protobuf:"varint,16,rep,packed,name=commodityServices,proto3,enum=v1.CommodityService" json:"commodityServices,omitempty" dynamodbav:"commodityServices,omitempty"`
}

func (x *Commodity) Reset() {
	*x = Commodity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commodity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Commodity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Commodity) ProtoMessage() {}

func (x *Commodity) ProtoReflect() protoreflect.Message {
	mi := &file_commodity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Commodity.ProtoReflect.Descriptor instead.
func (*Commodity) Descriptor() ([]byte, []int) {
	return file_commodity_proto_rawDescGZIP(), []int{2}
}

func (x *Commodity) GetDensity() float32 {
	if x != nil {
		return x.Density
	}
	return 0
}

func (x *Commodity) GetLength() float32 {
	if x != nil {
		return x.Length
	}
	return 0
}

func (x *Commodity) GetWidth() float32 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *Commodity) GetHeight() float32 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *Commodity) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *Commodity) GetDimensionUOM() *DimensionUOM {
	if x != nil {
		return x.DimensionUOM
	}
	return nil
}

func (x *Commodity) GetWeightUOM() *WeightUOM {
	if x != nil {
		return x.WeightUOM
	}
	return nil
}

func (x *Commodity) GetDimensionDisplay() string {
	if x != nil {
		return x.DimensionDisplay
	}
	return ""
}

func (x *Commodity) GetPackageType() PackageType {
	if x != nil {
		return x.PackageType
	}
	return PackageType_unknownPackageType
}

func (x *Commodity) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Commodity) GetFreightClass() FreightClass {
	if x != nil {
		return x.FreightClass
	}
	return FreightClass_unknownFreightClass
}

func (x *Commodity) GetInstructions() string {
	if x != nil {
		return x.Instructions
	}
	return ""
}

func (x *Commodity) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *Commodity) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Commodity) GetCommodityServices() []CommodityService {
	if x != nil {
		return x.CommodityServices
	}
	return nil
}

var File_commodity_proto protoreflect.FileDescriptor

var file_commodity_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x32,
	0x0a, 0x0c, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x4f, 0x4d, 0x12, 0x12,
	0x0a, 0x04, 0x49, 0x4e, 0x43, 0x48, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x49, 0x4e,
	0x43, 0x48, 0x12, 0x0e, 0x0a, 0x02, 0x43, 0x4d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02,
	0x43, 0x4d, 0x22, 0x2b, 0x0a, 0x09, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x55, 0x4f, 0x4d, 0x12,
	0x0e, 0x0a, 0x02, 0x4c, 0x42, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x4c, 0x42, 0x12,
	0x0e, 0x0a, 0x02, 0x4b, 0x47, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x4b, 0x47, 0x22,
	0x97, 0x05, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a,
	0x07, 0x64, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07,
	0x64, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x42, 0x0a, 0xba, 0xe9, 0xc0, 0x03, 0x05, 0xa2, 0x01,
	0x02, 0x08, 0x01, 0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x12, 0x20, 0x0a, 0x05, 0x77,
	0x69, 0x64, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x42, 0x0a, 0xba, 0xe9, 0xc0, 0x03,
	0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x05, 0x77, 0x69, 0x64, 0x74, 0x68, 0x12, 0x22, 0x0a,
	0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x42, 0x0a, 0xba,
	0xe9, 0xc0, 0x03, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x22, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x42, 0x0a, 0xba, 0xe9, 0xc0, 0x03, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x06, 0x77,
	0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x40, 0x0a, 0x0c, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x55, 0x4f, 0x4d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x76, 0x31,
	0x2e, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x55, 0x4f, 0x4d, 0x42, 0x0a, 0xba,
	0xe9, 0xc0, 0x03, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0c, 0x64, 0x69, 0x6d, 0x65, 0x6e,
	0x73, 0x69, 0x6f, 0x6e, 0x55, 0x4f, 0x4d, 0x12, 0x37, 0x0a, 0x09, 0x77, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x55, 0x4f, 0x4d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e,
	0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x55, 0x4f, 0x4d, 0x42, 0x0a, 0xba, 0xe9, 0xc0, 0x03, 0x05,
	0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x09, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x55, 0x4f, 0x4d,
	0x12, 0x2a, 0x0a, 0x10, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x69, 0x6d, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x44, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x12, 0x31, 0x0a, 0x0b,
	0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x26, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x0a, 0xba, 0xe9, 0xc0, 0x03, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x40, 0x0a, 0x0c, 0x66, 0x72, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x76, 0x31, 0x2e, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x42,
	0x0a, 0xba, 0xe9, 0xc0, 0x03, 0x05, 0xa2, 0x01, 0x02, 0x08, 0x01, 0x52, 0x0c, 0x66, 0x72, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x69, 0x6e, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x42, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69,
	0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x10, 0x20, 0x03, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x76, 0x31, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x11, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2a, 0x8a, 0x01, 0x0a, 0x10, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1b,
	0x0a, 0x17, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x69,
	0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x70,
	0x72, 0x6f, 0x74, 0x65, 0x63, 0x74, 0x46, 0x72, 0x6f, 0x6d, 0x46, 0x72, 0x65, 0x65, 0x7a, 0x65,
	0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x73, 0x6f, 0x72, 0x74, 0x41, 0x6e, 0x64, 0x53, 0x65, 0x67,
	0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x67, 0x75, 0x61, 0x72,
	0x61, 0x6e, 0x74, 0x65, 0x65, 0x64, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x68, 0x61, 0x7a, 0x61,
	0x72, 0x64, 0x6f, 0x75, 0x73, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x63, 0x6b,
	0x61, 0x62, 0x6c, 0x65, 0x10, 0x05, 0x2a, 0x98, 0x02, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x12, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77,
	0x6e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x10, 0x00, 0x12, 0x10,
	0x0a, 0x0c, 0x70, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x34, 0x38, 0x78, 0x34, 0x30, 0x10, 0x01,
	0x12, 0x10, 0x0a, 0x0c, 0x70, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x73, 0x34, 0x38, 0x78, 0x34, 0x38,
	0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x62, 0x61, 0x67, 0x73, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05,
	0x62, 0x61, 0x6c, 0x65, 0x73, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x62, 0x6f, 0x78, 0x65, 0x73,
	0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x10, 0x06, 0x12,
	0x0b, 0x0a, 0x07, 0x62, 0x75, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x10, 0x07, 0x12, 0x08, 0x0a, 0x04,
	0x63, 0x61, 0x6e, 0x73, 0x10, 0x08, 0x12, 0x0b, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x74, 0x6f, 0x6e,
	0x73, 0x10, 0x09, 0x12, 0x09, 0x0a, 0x05, 0x63, 0x61, 0x73, 0x65, 0x73, 0x10, 0x0a, 0x12, 0x09,
	0x0a, 0x05, 0x63, 0x6f, 0x69, 0x6c, 0x73, 0x10, 0x0b, 0x12, 0x0a, 0x0a, 0x06, 0x63, 0x72, 0x61,
	0x74, 0x65, 0x73, 0x10, 0x0c, 0x12, 0x0d, 0x0a, 0x09, 0x63, 0x79, 0x6c, 0x69, 0x6e, 0x64, 0x65,
	0x72, 0x73, 0x10, 0x0d, 0x12, 0x09, 0x0a, 0x05, 0x64, 0x72, 0x75, 0x6d, 0x73, 0x10, 0x0e, 0x12,
	0x09, 0x0a, 0x05, 0x70, 0x61, 0x69, 0x6c, 0x73, 0x10, 0x0f, 0x12, 0x09, 0x0a, 0x05, 0x72, 0x65,
	0x65, 0x6c, 0x73, 0x10, 0x10, 0x12, 0x09, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x6c, 0x73, 0x10, 0x11,
	0x12, 0x09, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x65, 0x73, 0x10, 0x12, 0x12, 0x09, 0x0a, 0x05, 0x74,
	0x75, 0x62, 0x65, 0x73, 0x10, 0x13, 0x12, 0x09, 0x0a, 0x05, 0x75, 0x6e, 0x69, 0x74, 0x73, 0x10,
	0x14, 0x2a, 0x9d, 0x02, 0x0a, 0x0c, 0x66, 0x72, 0x65, 0x69, 0x67, 0x68, 0x74, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x12, 0x17, 0x0a, 0x13, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x46, 0x72, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x35, 0x30, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x35, 0x35, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x36, 0x30,
	0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x36, 0x35, 0x10, 0x04, 0x12,
	0x0b, 0x0a, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x37, 0x30, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x37, 0x37, 0x35, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x63, 0x6c,
	0x61, 0x73, 0x73, 0x38, 0x35, 0x10, 0x07, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x39, 0x32, 0x35, 0x10, 0x08, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x31, 0x30,
	0x30, 0x10, 0x09, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x31, 0x31, 0x30, 0x10,
	0x0a, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x31, 0x32, 0x35, 0x10, 0x0b, 0x12,
	0x0c, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x31, 0x35, 0x30, 0x10, 0x0c, 0x12, 0x0c, 0x0a,
	0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x31, 0x37, 0x35, 0x10, 0x0d, 0x12, 0x0c, 0x0a, 0x08, 0x63,
	0x6c, 0x61, 0x73, 0x73, 0x32, 0x30, 0x30, 0x10, 0x0e, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x32, 0x35, 0x30, 0x10, 0x0f, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x33, 0x30, 0x30, 0x10, 0x10, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x34, 0x30,
	0x30, 0x10, 0x11, 0x12, 0x0c, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x35, 0x30, 0x30, 0x10,
	0x12, 0x42, 0x63, 0x0a, 0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x62, 0x6c,
	0x6f, 0x63, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31,
	0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02, 0x56, 0x31,
	0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_commodity_proto_rawDescOnce sync.Once
	file_commodity_proto_rawDescData = file_commodity_proto_rawDesc
)

func file_commodity_proto_rawDescGZIP() []byte {
	file_commodity_proto_rawDescOnce.Do(func() {
		file_commodity_proto_rawDescData = protoimpl.X.CompressGZIP(file_commodity_proto_rawDescData)
	})
	return file_commodity_proto_rawDescData
}

var file_commodity_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_commodity_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_commodity_proto_goTypes = []interface{}{
	(CommodityService)(0), // 0: v1.commodityService
	(PackageType)(0),      // 1: v1.packageType
	(FreightClass)(0),     // 2: v1.freightClass
	(*DimensionUOM)(nil),  // 3: v1.dimensionUOM
	(*WeightUOM)(nil),     // 4: v1.weightUOM
	(*Commodity)(nil),     // 5: v1.commodity
}
var file_commodity_proto_depIdxs = []int32{
	3, // 0: v1.commodity.dimensionUOM:type_name -> v1.dimensionUOM
	4, // 1: v1.commodity.weightUOM:type_name -> v1.weightUOM
	1, // 2: v1.commodity.packageType:type_name -> v1.packageType
	2, // 3: v1.commodity.freightClass:type_name -> v1.freightClass
	0, // 4: v1.commodity.commodityServices:type_name -> v1.commodityService
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_commodity_proto_init() }
func file_commodity_proto_init() {
	if File_commodity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commodity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DimensionUOM); i {
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
		file_commodity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeightUOM); i {
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
		file_commodity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Commodity); i {
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
			RawDescriptor: file_commodity_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_commodity_proto_goTypes,
		DependencyIndexes: file_commodity_proto_depIdxs,
		EnumInfos:         file_commodity_proto_enumTypes,
		MessageInfos:      file_commodity_proto_msgTypes,
	}.Build()
	File_commodity_proto = out.File
	file_commodity_proto_rawDesc = nil
	file_commodity_proto_goTypes = nil
	file_commodity_proto_depIdxs = nil
}
