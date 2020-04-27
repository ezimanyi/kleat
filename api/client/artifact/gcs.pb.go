// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: artifact/gcs.proto

package artifact

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

// Configuration for the Google Cloud Storage (GCS) artifact provider.
type GcsArtifactProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the GCS artifact provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured GCS accounts.
	Accounts []*GcsArtifactAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *GcsArtifactProvider) Reset() {
	*x = GcsArtifactProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_gcs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcsArtifactProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcsArtifactProvider) ProtoMessage() {}

func (x *GcsArtifactProvider) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_gcs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcsArtifactProvider.ProtoReflect.Descriptor instead.
func (*GcsArtifactProvider) Descriptor() ([]byte, []int) {
	return file_artifact_gcs_proto_rawDescGZIP(), []int{0}
}

func (x *GcsArtifactProvider) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *GcsArtifactProvider) GetAccounts() []*GcsArtifactAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

// Configuration for a GCS artifact account.
type GcsArtifactAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The path to a JSON key for a GCP service account with which to
	// authenticate. The service account must have the `roles/storage.admin`
	// role enabled.
	JsonPath string `protobuf:"bytes,2,opt,name=jsonPath,proto3" json:"jsonPath,omitempty"`
}

func (x *GcsArtifactAccount) Reset() {
	*x = GcsArtifactAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_gcs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GcsArtifactAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GcsArtifactAccount) ProtoMessage() {}

func (x *GcsArtifactAccount) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_gcs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GcsArtifactAccount.ProtoReflect.Descriptor instead.
func (*GcsArtifactAccount) Descriptor() ([]byte, []int) {
	return file_artifact_gcs_proto_rawDescGZIP(), []int{1}
}

func (x *GcsArtifactAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GcsArtifactAccount) GetJsonPath() string {
	if x != nil {
		return x.JsonPath
	}
	return ""
}

var File_artifact_gcs_proto protoreflect.FileDescriptor

var file_artifact_gcs_proto_rawDesc = []byte{
	0x0a, 0x12, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2f, 0x67, 0x63, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x22, 0x6f, 0x0a, 0x13, 0x47, 0x63, 0x73, 0x41, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x3e, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x47, 0x63, 0x73, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x44, 0x0a, 0x12, 0x47, 0x63, 0x73, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x42, 0x30, 0x5a, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61,
	0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_artifact_gcs_proto_rawDescOnce sync.Once
	file_artifact_gcs_proto_rawDescData = file_artifact_gcs_proto_rawDesc
)

func file_artifact_gcs_proto_rawDescGZIP() []byte {
	file_artifact_gcs_proto_rawDescOnce.Do(func() {
		file_artifact_gcs_proto_rawDescData = protoimpl.X.CompressGZIP(file_artifact_gcs_proto_rawDescData)
	})
	return file_artifact_gcs_proto_rawDescData
}

var file_artifact_gcs_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_artifact_gcs_proto_goTypes = []interface{}{
	(*GcsArtifactProvider)(nil), // 0: proto.artifact.GcsArtifactProvider
	(*GcsArtifactAccount)(nil),  // 1: proto.artifact.GcsArtifactAccount
}
var file_artifact_gcs_proto_depIdxs = []int32{
	1, // 0: proto.artifact.GcsArtifactProvider.accounts:type_name -> proto.artifact.GcsArtifactAccount
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_artifact_gcs_proto_init() }
func file_artifact_gcs_proto_init() {
	if File_artifact_gcs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_artifact_gcs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcsArtifactProvider); i {
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
		file_artifact_gcs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GcsArtifactAccount); i {
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
			RawDescriptor: file_artifact_gcs_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_artifact_gcs_proto_goTypes,
		DependencyIndexes: file_artifact_gcs_proto_depIdxs,
		MessageInfos:      file_artifact_gcs_proto_msgTypes,
	}.Build()
	File_artifact_gcs_proto = out.File
	file_artifact_gcs_proto_rawDesc = nil
	file_artifact_gcs_proto_goTypes = nil
	file_artifact_gcs_proto_depIdxs = nil
}
