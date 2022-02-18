// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: prod_model.proto

package pbfiles

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ProdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProdId int32 `protobuf:"varint,1,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"` //传入的商品ID
}

func (x *ProdRequest) Reset() {
	*x = ProdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prod_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdRequest) ProtoMessage() {}

func (x *ProdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_prod_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdRequest.ProtoReflect.Descriptor instead.
func (*ProdRequest) Descriptor() ([]byte, []int) {
	return file_prod_model_proto_rawDescGZIP(), []int{0}
}

func (x *ProdRequest) GetProdId() int32 {
	if x != nil {
		return x.ProdId
	}
	return 0
}

//以下是商品相关
type ProdModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`    //商品ID
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` //商品名称
}

func (x *ProdModel) Reset() {
	*x = ProdModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prod_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdModel) ProtoMessage() {}

func (x *ProdModel) ProtoReflect() protoreflect.Message {
	mi := &file_prod_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdModel.ProtoReflect.Descriptor instead.
func (*ProdModel) Descriptor() ([]byte, []int) {
	return file_prod_model_proto_rawDescGZIP(), []int{1}
}

func (x *ProdModel) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProdModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//请求返回的结果
type ProdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *ProdModel `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"` //结果
}

func (x *ProdResponse) Reset() {
	*x = ProdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_prod_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProdResponse) ProtoMessage() {}

func (x *ProdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_prod_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProdResponse.ProtoReflect.Descriptor instead.
func (*ProdResponse) Descriptor() ([]byte, []int) {
	return file_prod_model_proto_rawDescGZIP(), []int{2}
}

func (x *ProdResponse) GetResult() *ProdModel {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_prod_model_proto protoreflect.FileDescriptor

var file_prod_model_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x64, 0x49, 0x64, 0x22, 0x2f, 0x0a, 0x09, 0x50, 0x72,
	0x6f, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x32, 0x0a, 0x0c, 0x50,
	0x72, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42,
	0x17, 0x5a, 0x15, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x70, 0x72, 0x6f, 0x2f, 0x67, 0x73, 0x72, 0x63,
	0x2f, 0x70, 0x62, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_prod_model_proto_rawDescOnce sync.Once
	file_prod_model_proto_rawDescData = file_prod_model_proto_rawDesc
)

func file_prod_model_proto_rawDescGZIP() []byte {
	file_prod_model_proto_rawDescOnce.Do(func() {
		file_prod_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_prod_model_proto_rawDescData)
	})
	return file_prod_model_proto_rawDescData
}

var file_prod_model_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_prod_model_proto_goTypes = []interface{}{
	(*ProdRequest)(nil),  // 0: ProdRequest
	(*ProdModel)(nil),    // 1: ProdModel
	(*ProdResponse)(nil), // 2: ProdResponse
}
var file_prod_model_proto_depIdxs = []int32{
	1, // 0: ProdResponse.result:type_name -> ProdModel
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_prod_model_proto_init() }
func file_prod_model_proto_init() {
	if File_prod_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_prod_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdRequest); i {
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
		file_prod_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdModel); i {
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
		file_prod_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProdResponse); i {
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
			RawDescriptor: file_prod_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_prod_model_proto_goTypes,
		DependencyIndexes: file_prod_model_proto_depIdxs,
		MessageInfos:      file_prod_model_proto_msgTypes,
	}.Build()
	File_prod_model_proto = out.File
	file_prod_model_proto_rawDesc = nil
	file_prod_model_proto_goTypes = nil
	file_prod_model_proto_depIdxs = nil
}
