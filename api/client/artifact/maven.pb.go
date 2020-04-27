// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: artifact/maven.proto

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

// Configuration for the Maven artifact provider.
type MavenArtifactProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the Maven artifact provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured Maven accounts.
	Accounts []*MavenArtifactAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *MavenArtifactProvider) Reset() {
	*x = MavenArtifactProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_maven_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MavenArtifactProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MavenArtifactProvider) ProtoMessage() {}

func (x *MavenArtifactProvider) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_maven_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MavenArtifactProvider.ProtoReflect.Descriptor instead.
func (*MavenArtifactProvider) Descriptor() ([]byte, []int) {
	return file_artifact_maven_proto_rawDescGZIP(), []int{0}
}

func (x *MavenArtifactProvider) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *MavenArtifactProvider) GetAccounts() []*MavenArtifactAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

// Configuration for a Maven artifact account.
type MavenArtifactAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// (Required) Full URI for the Maven repository (e.g.,
	// `http://some.host.com/repository/path`).
	RepositoryUrl string `protobuf:"bytes,2,opt,name=repositoryUrl,proto3" json:"repositoryUrl,omitempty"`
}

func (x *MavenArtifactAccount) Reset() {
	*x = MavenArtifactAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_maven_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MavenArtifactAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MavenArtifactAccount) ProtoMessage() {}

func (x *MavenArtifactAccount) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_maven_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MavenArtifactAccount.ProtoReflect.Descriptor instead.
func (*MavenArtifactAccount) Descriptor() ([]byte, []int) {
	return file_artifact_maven_proto_rawDescGZIP(), []int{1}
}

func (x *MavenArtifactAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MavenArtifactAccount) GetRepositoryUrl() string {
	if x != nil {
		return x.RepositoryUrl
	}
	return ""
}

var File_artifact_maven_proto protoreflect.FileDescriptor

var file_artifact_maven_proto_rawDesc = []byte{
	0x0a, 0x14, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2f, 0x6d, 0x61, 0x76, 0x65, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x22, 0x73, 0x0a, 0x15, 0x4d, 0x61, 0x76, 0x65, 0x6e, 0x41,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x40, 0x0a, 0x08, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x4d, 0x61, 0x76,
	0x65, 0x6e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22, 0x50, 0x0a, 0x14, 0x4d,
	0x61, 0x76, 0x65, 0x6e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x55, 0x72, 0x6c, 0x42, 0x30, 0x5a,
	0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e,
	0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_artifact_maven_proto_rawDescOnce sync.Once
	file_artifact_maven_proto_rawDescData = file_artifact_maven_proto_rawDesc
)

func file_artifact_maven_proto_rawDescGZIP() []byte {
	file_artifact_maven_proto_rawDescOnce.Do(func() {
		file_artifact_maven_proto_rawDescData = protoimpl.X.CompressGZIP(file_artifact_maven_proto_rawDescData)
	})
	return file_artifact_maven_proto_rawDescData
}

var file_artifact_maven_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_artifact_maven_proto_goTypes = []interface{}{
	(*MavenArtifactProvider)(nil), // 0: proto.artifact.MavenArtifactProvider
	(*MavenArtifactAccount)(nil),  // 1: proto.artifact.MavenArtifactAccount
}
var file_artifact_maven_proto_depIdxs = []int32{
	1, // 0: proto.artifact.MavenArtifactProvider.accounts:type_name -> proto.artifact.MavenArtifactAccount
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_artifact_maven_proto_init() }
func file_artifact_maven_proto_init() {
	if File_artifact_maven_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_artifact_maven_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MavenArtifactProvider); i {
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
		file_artifact_maven_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MavenArtifactAccount); i {
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
			RawDescriptor: file_artifact_maven_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_artifact_maven_proto_goTypes,
		DependencyIndexes: file_artifact_maven_proto_depIdxs,
		MessageInfos:      file_artifact_maven_proto_msgTypes,
	}.Build()
	File_artifact_maven_proto = out.File
	file_artifact_maven_proto_rawDesc = nil
	file_artifact_maven_proto_goTypes = nil
	file_artifact_maven_proto_depIdxs = nil
}
