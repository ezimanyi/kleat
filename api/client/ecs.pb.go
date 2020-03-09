// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ecs.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Configuration for the ECS provider.
type Ecs struct {
	// Whether the provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured accounts.
	Accounts []*EcsAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	// The name of the primary account.
	PrimaryAccount       string   `protobuf:"bytes,3,opt,name=primaryAccount,proto3" json:"primaryAccount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ecs) Reset()         { *m = Ecs{} }
func (m *Ecs) String() string { return proto.CompactTextString(m) }
func (*Ecs) ProtoMessage()    {}
func (*Ecs) Descriptor() ([]byte, []int) {
	return fileDescriptor_290e148f4db3a904, []int{0}
}

func (m *Ecs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ecs.Unmarshal(m, b)
}
func (m *Ecs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ecs.Marshal(b, m, deterministic)
}
func (m *Ecs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ecs.Merge(m, src)
}
func (m *Ecs) XXX_Size() int {
	return xxx_messageInfo_Ecs.Size(m)
}
func (m *Ecs) XXX_DiscardUnknown() {
	xxx_messageInfo_Ecs.DiscardUnknown(m)
}

var xxx_messageInfo_Ecs proto.InternalMessageInfo

func (m *Ecs) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Ecs) GetAccounts() []*EcsAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *Ecs) GetPrimaryAccount() string {
	if m != nil {
		return m.PrimaryAccount
	}
	return ""
}

// Configuration for an ECS account.
type EcsAccount struct {
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
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *EcsAccount) Reset()         { *m = EcsAccount{} }
func (m *EcsAccount) String() string { return proto.CompactTextString(m) }
func (*EcsAccount) ProtoMessage()    {}
func (*EcsAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_290e148f4db3a904, []int{1}
}

func (m *EcsAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EcsAccount.Unmarshal(m, b)
}
func (m *EcsAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EcsAccount.Marshal(b, m, deterministic)
}
func (m *EcsAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EcsAccount.Merge(m, src)
}
func (m *EcsAccount) XXX_Size() int {
	return xxx_messageInfo_EcsAccount.Size(m)
}
func (m *EcsAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_EcsAccount.DiscardUnknown(m)
}

var xxx_messageInfo_EcsAccount proto.InternalMessageInfo

func (m *EcsAccount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EcsAccount) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *EcsAccount) GetAwsAccount() string {
	if m != nil {
		return m.AwsAccount
	}
	return ""
}

func (m *EcsAccount) GetPermissions() *Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *EcsAccount) GetRequiredGroupMemberships() []string {
	if m != nil {
		return m.RequiredGroupMemberships
	}
	return nil
}

func init() {
	proto.RegisterType((*Ecs)(nil), "proto.Ecs")
	proto.RegisterType((*EcsAccount)(nil), "proto.EcsAccount")
}

func init() { proto.RegisterFile("ecs.proto", fileDescriptor_290e148f4db3a904) }

var fileDescriptor_290e148f4db3a904 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0xd9, 0x6e, 0xab, 0xdd, 0x59, 0x10, 0x3a, 0xa7, 0xe0, 0x41, 0x42, 0x0f, 0x92, 0x8b,
	0x3d, 0x54, 0x4f, 0xde, 0x3c, 0x14, 0x4f, 0x82, 0xe4, 0x1b, 0x64, 0xd3, 0x01, 0x03, 0xe6, 0x8f,
	0x99, 0xdd, 0x8a, 0xdf, 0xd3, 0x0f, 0x24, 0xa4, 0xab, 0x5d, 0x05, 0x4f, 0x49, 0xde, 0xef, 0xe5,
	0xbd, 0x19, 0x68, 0xc8, 0xf2, 0x26, 0xe5, 0xd8, 0x47, 0x5c, 0x94, 0xe3, 0x72, 0x95, 0x28, 0x7b,
	0xc7, 0xec, 0x62, 0x18, 0xc9, 0xfa, 0x00, 0xf5, 0xce, 0x32, 0x0a, 0x38, 0xa7, 0x60, 0xba, 0x57,
	0xda, 0x8b, 0x4a, 0x56, 0x6a, 0xa9, 0xbf, 0x9f, 0x78, 0x03, 0x4b, 0x63, 0x6d, 0x1c, 0x42, 0xcf,
	0x62, 0x26, 0x6b, 0xd5, 0x6e, 0x57, 0xc7, 0xaf, 0x9b, 0x9d, 0xe5, 0x87, 0x23, 0xd1, 0x3f, 0x16,
	0xbc, 0x86, 0x8b, 0x94, 0x9d, 0x37, 0xf9, 0x63, 0x64, 0xa2, 0x96, 0x95, 0x6a, 0xf4, 0x1f, 0x75,
	0xfd, 0x59, 0x01, 0x9c, 0x02, 0x10, 0x61, 0x1e, 0x8c, 0xa7, 0x52, 0xde, 0xe8, 0x72, 0x47, 0x09,
	0x2d, 0x85, 0x83, 0xcb, 0x31, 0x78, 0x0a, 0xbd, 0x98, 0x15, 0x34, 0x95, 0xf0, 0x0a, 0xc0, 0xbc,
	0xf3, 0xef, 0xa2, 0x89, 0x82, 0x77, 0xd0, 0x4e, 0x36, 0x16, 0x73, 0x59, 0xa9, 0x76, 0x8b, 0xe3,
	0xf8, 0xcf, 0x27, 0xa2, 0xa7, 0x36, 0xbc, 0x07, 0x91, 0xe9, 0x6d, 0x70, 0x99, 0xf6, 0x8f, 0x39,
	0x0e, 0xe9, 0x89, 0x7c, 0x47, 0x99, 0x5f, 0x5c, 0x62, 0xb1, 0x90, 0xb5, 0x6a, 0xf4, 0xbf, 0xbc,
	0x3b, 0x2b, 0xd9, 0xb7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xe4, 0xc5, 0x7f, 0x7c, 0x01,
	0x00, 0x00,
}
