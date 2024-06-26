// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.4
// source: admin.proto

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

type AdminResponse_Status int32

const (
	AdminResponse_OK    AdminResponse_Status = 0
	AdminResponse_ERROR AdminResponse_Status = 1
)

// Enum value maps for AdminResponse_Status.
var (
	AdminResponse_Status_name = map[int32]string{
		0: "OK",
		1: "ERROR",
	}
	AdminResponse_Status_value = map[string]int32{
		"OK":    0,
		"ERROR": 1,
	}
)

func (x AdminResponse_Status) Enum() *AdminResponse_Status {
	p := new(AdminResponse_Status)
	*p = x
	return p
}

func (x AdminResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdminResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_admin_proto_enumTypes[0].Descriptor()
}

func (AdminResponse_Status) Type() protoreflect.EnumType {
	return &file_admin_proto_enumTypes[0]
}

func (x AdminResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AdminResponse_Status.Descriptor instead.
func (AdminResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{1, 0}
}

type Amount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount float32 `protobuf:"fixed32,1,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *Amount) Reset() {
	*x = Amount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Amount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Amount) ProtoMessage() {}

func (x *Amount) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Amount.ProtoReflect.Descriptor instead.
func (*Amount) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{0}
}

func (x *Amount) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type AdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  AdminResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=pb.AdminResponse_Status" json:"status,omitempty"`
	Message string               `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	// Types that are assignable to Payload:
	//
	//	*AdminResponse_Error
	//	*AdminResponse_Data
	Payload isAdminResponse_Payload `protobuf_oneof:"payload"`
}

func (x *AdminResponse) Reset() {
	*x = AdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdminResponse) ProtoMessage() {}

func (x *AdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdminResponse.ProtoReflect.Descriptor instead.
func (*AdminResponse) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{1}
}

func (x *AdminResponse) GetStatus() AdminResponse_Status {
	if x != nil {
		return x.Status
	}
	return AdminResponse_OK
}

func (x *AdminResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (m *AdminResponse) GetPayload() isAdminResponse_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *AdminResponse) GetError() string {
	if x, ok := x.GetPayload().(*AdminResponse_Error); ok {
		return x.Error
	}
	return ""
}

func (x *AdminResponse) GetData() string {
	if x, ok := x.GetPayload().(*AdminResponse_Data); ok {
		return x.Data
	}
	return ""
}

type isAdminResponse_Payload interface {
	isAdminResponse_Payload()
}

type AdminResponse_Error struct {
	Error string `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

type AdminResponse_Data struct {
	Data string `protobuf:"bytes,4,opt,name=data,proto3,oneof"`
}

func (*AdminResponse_Error) isAdminResponse_Payload() {}

func (*AdminResponse_Data) isAdminResponse_Payload() {}

var File_admin_proto protoreflect.FileDescriptor

var file_admin_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x20, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0xb1, 0x01, 0x0a, 0x0d, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x16, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x1b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x42, 0x09, 0x0a, 0x07,
	0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x32, 0x41, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x54, 0x6f,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x12, 0x0a, 0x2e, 0x70, 0x62,
	0x2e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_proto_rawDescOnce sync.Once
	file_admin_proto_rawDescData = file_admin_proto_rawDesc
)

func file_admin_proto_rawDescGZIP() []byte {
	file_admin_proto_rawDescOnce.Do(func() {
		file_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_proto_rawDescData)
	})
	return file_admin_proto_rawDescData
}

var file_admin_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_admin_proto_goTypes = []interface{}{
	(AdminResponse_Status)(0), // 0: pb.AdminResponse.Status
	(*Amount)(nil),            // 1: pb.Amount
	(*AdminResponse)(nil),     // 2: pb.AdminResponse
}
var file_admin_proto_depIdxs = []int32{
	0, // 0: pb.AdminResponse.status:type_name -> pb.AdminResponse.Status
	1, // 1: pb.AdminService.AddToAdminWallet:input_type -> pb.Amount
	2, // 2: pb.AdminService.AddToAdminWallet:output_type -> pb.AdminResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_admin_proto_init() }
func file_admin_proto_init() {
	if File_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Amount); i {
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
		file_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdminResponse); i {
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
	file_admin_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*AdminResponse_Error)(nil),
		(*AdminResponse_Data)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_admin_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_proto_goTypes,
		DependencyIndexes: file_admin_proto_depIdxs,
		EnumInfos:         file_admin_proto_enumTypes,
		MessageInfos:      file_admin_proto_msgTypes,
	}.Build()
	File_admin_proto = out.File
	file_admin_proto_rawDesc = nil
	file_admin_proto_goTypes = nil
	file_admin_proto_depIdxs = nil
}
