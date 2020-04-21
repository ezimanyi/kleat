// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: ecs_provider.proto

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

// Configuration for the ECS provider.
type EcsProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured accounts.
	Accounts []*EcsAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	// The name of the primary account.
	PrimaryAccount string `protobuf:"bytes,3,opt,name=primaryAccount,proto3" json:"primaryAccount,omitempty"`
}

func (x *EcsProvider) Reset() {
	*x = EcsProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ecs_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EcsProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EcsProvider) ProtoMessage() {}

func (x *EcsProvider) ProtoReflect() protoreflect.Message {
	mi := &file_ecs_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EcsProvider.ProtoReflect.Descriptor instead.
func (*EcsProvider) Descriptor() ([]byte, []int) {
	return file_ecs_provider_proto_rawDescGZIP(), []int{0}
}

func (x *EcsProvider) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *EcsProvider) GetAccounts() []*EcsAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

func (x *EcsProvider) GetPrimaryAccount() string {
	if x != nil {
		return x.PrimaryAccount
	}
	return ""
}

// Configuration for an ECS account.
type EcsAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The environment name for the account. Many accounts can share the
	// same environment (e.g., dev, test, prod).
	Environment string `protobuf:"bytes,2,opt,name=environment,proto3" json:"environment,omitempty"`
	// (Required) Provide the name of the AWS account associated with this ECS
	// account. See
	// https://github.com/spinnaker/clouddriver/blob/master/clouddriver-ecs/README.md
	// for more information.
	AwsAccount string `protobuf:"bytes,3,opt,name=awsAccount,proto3" json:"awsAccount,omitempty"`
	// Fiat permissions configuration.
	Permissions *Permissions `protobuf:"bytes,4,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// (Deprecated) List of required Fiat permission groups. Configure
	// `permissions` instead.
	RequiredGroupMemberships []string `protobuf:"bytes,5,rep,name=requiredGroupMemberships,proto3" json:"requiredGroupMemberships,omitempty"`
}

func (x *EcsAccount) Reset() {
	*x = EcsAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ecs_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EcsAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EcsAccount) ProtoMessage() {}

func (x *EcsAccount) ProtoReflect() protoreflect.Message {
	mi := &file_ecs_provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EcsAccount.ProtoReflect.Descriptor instead.
func (*EcsAccount) Descriptor() ([]byte, []int) {
	return file_ecs_provider_proto_rawDescGZIP(), []int{1}
}

func (x *EcsAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EcsAccount) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *EcsAccount) GetAwsAccount() string {
	if x != nil {
		return x.AwsAccount
	}
	return ""
}

func (x *EcsAccount) GetPermissions() *Permissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *EcsAccount) GetRequiredGroupMemberships() []string {
	if x != nil {
		return x.RequiredGroupMemberships
	}
	return nil
}

var File_ecs_provider_proto protoreflect.FileDescriptor

var file_ecs_provider_proto_rawDesc = []byte{
	0x0a, 0x12, 0x65, 0x63, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e,
	0x0a, 0x0b, 0x45, 0x63, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x2d, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x63, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72,
	0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xd4,
	0x01, 0x0a, 0x0a, 0x45, 0x63, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x77, 0x73, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x77, 0x73, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3a, 0x0a, 0x18, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x18, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x73, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c,
	0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ecs_provider_proto_rawDescOnce sync.Once
	file_ecs_provider_proto_rawDescData = file_ecs_provider_proto_rawDesc
)

func file_ecs_provider_proto_rawDescGZIP() []byte {
	file_ecs_provider_proto_rawDescOnce.Do(func() {
		file_ecs_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_ecs_provider_proto_rawDescData)
	})
	return file_ecs_provider_proto_rawDescData
}

var file_ecs_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ecs_provider_proto_goTypes = []interface{}{
	(*EcsProvider)(nil), // 0: proto.EcsProvider
	(*EcsAccount)(nil),  // 1: proto.EcsAccount
	(*Permissions)(nil), // 2: proto.Permissions
}
var file_ecs_provider_proto_depIdxs = []int32{
	1, // 0: proto.EcsProvider.accounts:type_name -> proto.EcsAccount
	2, // 1: proto.EcsAccount.permissions:type_name -> proto.Permissions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ecs_provider_proto_init() }
func file_ecs_provider_proto_init() {
	if File_ecs_provider_proto != nil {
		return
	}
	file_permissions_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ecs_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EcsProvider); i {
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
		file_ecs_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EcsAccount); i {
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
			RawDescriptor: file_ecs_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ecs_provider_proto_goTypes,
		DependencyIndexes: file_ecs_provider_proto_depIdxs,
		MessageInfos:      file_ecs_provider_proto_msgTypes,
	}.Build()
	File_ecs_provider_proto = out.File
	file_ecs_provider_proto_rawDesc = nil
	file_ecs_provider_proto_goTypes = nil
	file_ecs_provider_proto_depIdxs = nil
}
