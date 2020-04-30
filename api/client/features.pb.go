// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: features.proto

package client

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

// Feature flags
type Features struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Enable pipeline template support. Read more at https://github.com/spinnaker/dcd-spec.
	PipelineTemplates bool `protobuf:"varint,1,opt,name=pipelineTemplates,proto3" json:"pipelineTemplates,omitempty"`
	Auth              bool `protobuf:"varint,2,opt,name=auth,proto3" json:"auth,omitempty"`
	Fiat              bool `protobuf:"varint,3,opt,name=fiat,proto3" json:"fiat,omitempty"`
	Chaos             bool `protobuf:"varint,4,opt,name=chaos,proto3" json:"chaos,omitempty"`
	EntityTags        bool `protobuf:"varint,5,opt,name=entityTags,proto3" json:"entityTags,omitempty"`
	// Enable canary support. For this to work, you'll need a canary judge configured.
	MineCanary                   bool `protobuf:"varint,6,opt,name=mineCanary,proto3" json:"mineCanary,omitempty"`
	ManagedPipelineTemplatesV2UI bool `protobuf:"varint,7,opt,name=managedPipelineTemplatesV2UI,proto3" json:"managedPipelineTemplatesV2UI,omitempty"`
}

func (x *Features) Reset() {
	*x = Features{}
	if protoimpl.UnsafeEnabled {
		mi := &file_features_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Features) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Features) ProtoMessage() {}

func (x *Features) ProtoReflect() protoreflect.Message {
	mi := &file_features_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Features.ProtoReflect.Descriptor instead.
func (*Features) Descriptor() ([]byte, []int) {
	return file_features_proto_rawDescGZIP(), []int{0}
}

func (x *Features) GetPipelineTemplates() bool {
	if x != nil {
		return x.PipelineTemplates
	}
	return false
}

func (x *Features) GetAuth() bool {
	if x != nil {
		return x.Auth
	}
	return false
}

func (x *Features) GetFiat() bool {
	if x != nil {
		return x.Fiat
	}
	return false
}

func (x *Features) GetChaos() bool {
	if x != nil {
		return x.Chaos
	}
	return false
}

func (x *Features) GetEntityTags() bool {
	if x != nil {
		return x.EntityTags
	}
	return false
}

func (x *Features) GetMineCanary() bool {
	if x != nil {
		return x.MineCanary
	}
	return false
}

func (x *Features) GetManagedPipelineTemplatesV2UI() bool {
	if x != nil {
		return x.ManagedPipelineTemplatesV2UI
	}
	return false
}

var File_features_proto protoreflect.FileDescriptor

var file_features_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x01, 0x0a, 0x08, 0x46, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x12, 0x2c, 0x0a, 0x11, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x11, 0x70, 0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x04, 0x61, 0x75, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x61, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x66, 0x69, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68,
	0x61, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x63, 0x68, 0x61, 0x6f, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x61, 0x67, 0x73, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x54, 0x61, 0x67, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x69, 0x6e, 0x65, 0x43, 0x61, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x6d, 0x69, 0x6e, 0x65, 0x43, 0x61, 0x6e, 0x61, 0x72, 0x79,
	0x12, 0x42, 0x0a, 0x1c, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x50, 0x69, 0x70, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x56, 0x32, 0x55, 0x49,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1c, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x64, 0x50,
	0x69, 0x70, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x56, 0x32, 0x55, 0x49, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65,
	0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_features_proto_rawDescOnce sync.Once
	file_features_proto_rawDescData = file_features_proto_rawDesc
)

func file_features_proto_rawDescGZIP() []byte {
	file_features_proto_rawDescOnce.Do(func() {
		file_features_proto_rawDescData = protoimpl.X.CompressGZIP(file_features_proto_rawDescData)
	})
	return file_features_proto_rawDescData
}

var file_features_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_features_proto_goTypes = []interface{}{
	(*Features)(nil), // 0: proto.Features
}
var file_features_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_features_proto_init() }
func file_features_proto_init() {
	if File_features_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_features_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Features); i {
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
			RawDescriptor: file_features_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_features_proto_goTypes,
		DependencyIndexes: file_features_proto_depIdxs,
		MessageInfos:      file_features_proto_msgTypes,
	}.Build()
	File_features_proto = out.File
	file_features_proto_rawDesc = nil
	file_features_proto_goTypes = nil
	file_features_proto_depIdxs = nil
}
