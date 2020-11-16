// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: config/rosco.proto

package config

import (
	proto "github.com/golang/protobuf/proto"
	cloudprovider "github.com/spinnaker/kleat/api/client/cloudprovider"
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

// Configuration for the rosco microservice.
type Rosco struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Google      *cloudprovider.GoogleComputeEngine `protobuf:"bytes,1,opt,name=google,proto3" json:"google,omitempty"`
	Aws         *cloudprovider.Aws                 `protobuf:"bytes,2,opt,name=aws,proto3" json:"aws,omitempty"`
	Azure       *cloudprovider.Azure               `protobuf:"bytes,3,opt,name=azure,proto3" json:"azure,omitempty"`
	Huaweicloud *cloudprovider.HuaweiCloud         `protobuf:"bytes,4,opt,name=huaweicloud,proto3" json:"huaweicloud,omitempty"`
	Oracle      *cloudprovider.Oracle              `protobuf:"bytes,5,opt,name=oracle,proto3" json:"oracle,omitempty"`
}

func (x *Rosco) Reset() {
	*x = Rosco{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_rosco_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rosco) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rosco) ProtoMessage() {}

func (x *Rosco) ProtoReflect() protoreflect.Message {
	mi := &file_config_rosco_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rosco.ProtoReflect.Descriptor instead.
func (*Rosco) Descriptor() ([]byte, []int) {
	return file_config_rosco_proto_rawDescGZIP(), []int{0}
}

func (x *Rosco) GetGoogle() *cloudprovider.GoogleComputeEngine {
	if x != nil {
		return x.Google
	}
	return nil
}

func (x *Rosco) GetAws() *cloudprovider.Aws {
	if x != nil {
		return x.Aws
	}
	return nil
}

func (x *Rosco) GetAzure() *cloudprovider.Azure {
	if x != nil {
		return x.Azure
	}
	return nil
}

func (x *Rosco) GetHuaweicloud() *cloudprovider.HuaweiCloud {
	if x != nil {
		return x.Huaweicloud
	}
	return nil
}

func (x *Rosco) GetOracle() *cloudprovider.Oracle {
	if x != nil {
		return x.Oracle
	}
	return nil
}

var File_config_rosco_proto protoreflect.FileDescriptor

var file_config_rosco_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x72, 0x6f, 0x73, 0x63, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x1a, 0x17, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2f, 0x61, 0x77, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x61, 0x7a, 0x75, 0x72, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2f, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xa0, 0x02, 0x0a, 0x05, 0x52, 0x6f, 0x73, 0x63, 0x6f, 0x12, 0x40, 0x0a, 0x06, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x75, 0x74, 0x65, 0x45, 0x6e, 0x67,
	0x69, 0x6e, 0x65, 0x52, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x03, 0x61,
	0x77, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x41,
	0x77, 0x73, 0x52, 0x03, 0x61, 0x77, 0x73, 0x12, 0x30, 0x0a, 0x05, 0x61, 0x7a, 0x75, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x41, 0x7a, 0x75,
	0x72, 0x65, 0x52, 0x05, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x68, 0x75, 0x61,
	0x77, 0x65, 0x69, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x48, 0x75, 0x61, 0x77, 0x65, 0x69, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x52, 0x0b, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x12, 0x33, 0x0a,
	0x06, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x52, 0x06, 0x6f, 0x72, 0x61, 0x63,
	0x6c, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_rosco_proto_rawDescOnce sync.Once
	file_config_rosco_proto_rawDescData = file_config_rosco_proto_rawDesc
)

func file_config_rosco_proto_rawDescGZIP() []byte {
	file_config_rosco_proto_rawDescOnce.Do(func() {
		file_config_rosco_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_rosco_proto_rawDescData)
	})
	return file_config_rosco_proto_rawDescData
}

var file_config_rosco_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_config_rosco_proto_goTypes = []interface{}{
	(*Rosco)(nil), // 0: proto.config.Rosco
	(*cloudprovider.GoogleComputeEngine)(nil), // 1: proto.cloudprovider.GoogleComputeEngine
	(*cloudprovider.Aws)(nil),                 // 2: proto.cloudprovider.Aws
	(*cloudprovider.Azure)(nil),               // 3: proto.cloudprovider.Azure
	(*cloudprovider.HuaweiCloud)(nil),         // 4: proto.cloudprovider.HuaweiCloud
	(*cloudprovider.Oracle)(nil),              // 5: proto.cloudprovider.Oracle
}
var file_config_rosco_proto_depIdxs = []int32{
	1, // 0: proto.config.Rosco.google:type_name -> proto.cloudprovider.GoogleComputeEngine
	2, // 1: proto.config.Rosco.aws:type_name -> proto.cloudprovider.Aws
	3, // 2: proto.config.Rosco.azure:type_name -> proto.cloudprovider.Azure
	4, // 3: proto.config.Rosco.huaweicloud:type_name -> proto.cloudprovider.HuaweiCloud
	5, // 4: proto.config.Rosco.oracle:type_name -> proto.cloudprovider.Oracle
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_config_rosco_proto_init() }
func file_config_rosco_proto_init() {
	if File_config_rosco_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_rosco_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rosco); i {
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
			RawDescriptor: file_config_rosco_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_rosco_proto_goTypes,
		DependencyIndexes: file_config_rosco_proto_depIdxs,
		MessageInfos:      file_config_rosco_proto_msgTypes,
	}.Build()
	File_config_rosco_proto = out.File
	file_config_rosco_proto_rawDesc = nil
	file_config_rosco_proto_goTypes = nil
	file_config_rosco_proto_depIdxs = nil
}
