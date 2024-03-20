// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: product.proto

package __

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

type Prodcut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name               string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Seller_ID          uint32 `protobuf:"varint,2,opt,name=Seller_ID,json=SellerID,proto3" json:"Seller_ID,omitempty"`
	Category           uint32 `protobuf:"varint,3,opt,name=Category,proto3" json:"Category,omitempty"`
	Base_Price         uint32 `protobuf:"varint,4,opt,name=Base_Price,json=BasePrice,proto3" json:"Base_Price,omitempty"`
	Description        string `protobuf:"bytes,5,opt,name=Description,proto3" json:"Description,omitempty"`
	Bidding_Type       bool   `protobuf:"varint,6,opt,name=Bidding_Type,json=BiddingType,proto3" json:"Bidding_Type,omitempty"`
	Bidding_Start_Time string `protobuf:"bytes,7,opt,name=Bidding_Start_Time,json=BiddingStartTime,proto3" json:"Bidding_Start_Time,omitempty"`
	Bidiing_End_Time   string `protobuf:"bytes,8,opt,name=Bidiing_End_Time,json=BidiingEndTime,proto3" json:"Bidiing_End_Time,omitempty"`
	Listed_On          string `protobuf:"bytes,9,opt,name=Listed_On,json=ListedOn,proto3" json:"Listed_On,omitempty"`
}

func (x *Prodcut) Reset() {
	*x = Prodcut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prodcut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prodcut) ProtoMessage() {}

func (x *Prodcut) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prodcut.ProtoReflect.Descriptor instead.
func (*Prodcut) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{0}
}

func (x *Prodcut) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Prodcut) GetSeller_ID() uint32 {
	if x != nil {
		return x.Seller_ID
	}
	return 0
}

func (x *Prodcut) GetCategory() uint32 {
	if x != nil {
		return x.Category
	}
	return 0
}

func (x *Prodcut) GetBase_Price() uint32 {
	if x != nil {
		return x.Base_Price
	}
	return 0
}

func (x *Prodcut) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Prodcut) GetBidding_Type() bool {
	if x != nil {
		return x.Bidding_Type
	}
	return false
}

func (x *Prodcut) GetBidding_Start_Time() string {
	if x != nil {
		return x.Bidding_Start_Time
	}
	return ""
}

func (x *Prodcut) GetBidiing_End_Time() string {
	if x != nil {
		return x.Bidiing_End_Time
	}
	return ""
}

func (x *Prodcut) GetListed_On() string {
	if x != nil {
		return x.Listed_On
	}
	return ""
}

type ProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  string `protobuf:"bytes,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Error   string `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *ProductResponse) Reset() {
	*x = ProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductResponse) ProtoMessage() {}

func (x *ProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductResponse.ProtoReflect.Descriptor instead.
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return file_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ProductResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ProductResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_product_proto protoreflect.FileDescriptor

var file_product_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0xaf, 0x02, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x63, 0x75, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x53, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a,
	0x42, 0x61, 0x73, 0x65, 0x5f, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x09, 0x42, 0x61, 0x73, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x44,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a,
	0x0c, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x2c, 0x0a, 0x12, 0x42, 0x69, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x42, 0x69,
	0x64, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28,
	0x0a, 0x10, 0x42, 0x69, 0x64, 0x69, 0x69, 0x6e, 0x67, 0x5f, 0x45, 0x6e, 0x64, 0x5f, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x42, 0x69, 0x64, 0x69, 0x69, 0x6e,
	0x67, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74,
	0x65, 0x64, 0x5f, 0x4f, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4c, 0x69, 0x73,
	0x74, 0x65, 0x64, 0x4f, 0x6e, 0x22, 0x59, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x32, 0x43, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x31, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x63, 0x75, 0x74,
	0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_product_proto_rawDescOnce sync.Once
	file_product_proto_rawDescData = file_product_proto_rawDesc
)

func file_product_proto_rawDescGZIP() []byte {
	file_product_proto_rawDescOnce.Do(func() {
		file_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_product_proto_rawDescData)
	})
	return file_product_proto_rawDescData
}

var file_product_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_product_proto_goTypes = []interface{}{
	(*Prodcut)(nil),         // 0: pb.Prodcut
	(*ProductResponse)(nil), // 1: pb.ProductResponse
}
var file_product_proto_depIdxs = []int32{
	0, // 0: pb.ProductService.CreateProduct:input_type -> pb.Prodcut
	1, // 1: pb.ProductService.CreateProduct:output_type -> pb.ProductResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_product_proto_init() }
func file_product_proto_init() {
	if File_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Prodcut); i {
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
		file_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductResponse); i {
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
			RawDescriptor: file_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_product_proto_goTypes,
		DependencyIndexes: file_product_proto_depIdxs,
		MessageInfos:      file_product_proto_msgTypes,
	}.Build()
	File_product_proto = out.File
	file_product_proto_rawDesc = nil
	file_product_proto_goTypes = nil
	file_product_proto_depIdxs = nil
}