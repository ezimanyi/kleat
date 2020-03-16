// Code generated by protoc-gen-go. DO NOT EDIT.
// source: s3_artifact_provider.proto

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

// Configuration for the S3 artifact provider.
type S3ArtifactProvider struct {
	// Whether the S3 artifact provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured S3 artifact accounts.
	Accounts             []*S3ArtifactAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *S3ArtifactProvider) Reset()         { *m = S3ArtifactProvider{} }
func (m *S3ArtifactProvider) String() string { return proto.CompactTextString(m) }
func (*S3ArtifactProvider) ProtoMessage()    {}
func (*S3ArtifactProvider) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d1b7390c451dbc2, []int{0}
}

func (m *S3ArtifactProvider) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S3ArtifactProvider.Unmarshal(m, b)
}
func (m *S3ArtifactProvider) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S3ArtifactProvider.Marshal(b, m, deterministic)
}
func (m *S3ArtifactProvider) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S3ArtifactProvider.Merge(m, src)
}
func (m *S3ArtifactProvider) XXX_Size() int {
	return xxx_messageInfo_S3ArtifactProvider.Size(m)
}
func (m *S3ArtifactProvider) XXX_DiscardUnknown() {
	xxx_messageInfo_S3ArtifactProvider.DiscardUnknown(m)
}

var xxx_messageInfo_S3ArtifactProvider proto.InternalMessageInfo

func (m *S3ArtifactProvider) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *S3ArtifactProvider) GetAccounts() []*S3ArtifactAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
}

// Configuration for an S3 artifact account.
type S3ArtifactAccount struct {
	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The S3 API endpoint. Only required when using an S3 clone such as Minio.
	ApiEndpoint string `protobuf:"bytes,2,opt,name=apiEndpoint,proto3" json:"apiEndpoint,omitempty"`
	// The S3 API region. Only required when using an S3 clone such as Minio.
	ApiRegion string `protobuf:"bytes,3,opt,name=apiRegion,proto3" json:"apiRegion,omitempty"`
	// The AWS Access Key ID. If not provided, Spinnaker will try to find AWS
	// credentials as described at
	// http://docs.aws.amazon.com/sdk-for-java/v1/developer-guide/credentials.html#credentials-default.
	AwsAccessKeyId string `protobuf:"bytes,4,opt,name=awsAccessKeyId,proto3" json:"awsAccessKeyId,omitempty"`
	// The AWS Secret Key.
	AwsSecretAccessKey string `protobuf:"bytes,5,opt,name=awsSecretAccessKey,proto3" json:"awsSecretAccessKey,omitempty"`
	// The S3 region.
	Region               string   `protobuf:"bytes,6,opt,name=region,proto3" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *S3ArtifactAccount) Reset()         { *m = S3ArtifactAccount{} }
func (m *S3ArtifactAccount) String() string { return proto.CompactTextString(m) }
func (*S3ArtifactAccount) ProtoMessage()    {}
func (*S3ArtifactAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d1b7390c451dbc2, []int{1}
}

func (m *S3ArtifactAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S3ArtifactAccount.Unmarshal(m, b)
}
func (m *S3ArtifactAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S3ArtifactAccount.Marshal(b, m, deterministic)
}
func (m *S3ArtifactAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S3ArtifactAccount.Merge(m, src)
}
func (m *S3ArtifactAccount) XXX_Size() int {
	return xxx_messageInfo_S3ArtifactAccount.Size(m)
}
func (m *S3ArtifactAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_S3ArtifactAccount.DiscardUnknown(m)
}

var xxx_messageInfo_S3ArtifactAccount proto.InternalMessageInfo

func (m *S3ArtifactAccount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *S3ArtifactAccount) GetApiEndpoint() string {
	if m != nil {
		return m.ApiEndpoint
	}
	return ""
}

func (m *S3ArtifactAccount) GetApiRegion() string {
	if m != nil {
		return m.ApiRegion
	}
	return ""
}

func (m *S3ArtifactAccount) GetAwsAccessKeyId() string {
	if m != nil {
		return m.AwsAccessKeyId
	}
	return ""
}

func (m *S3ArtifactAccount) GetAwsSecretAccessKey() string {
	if m != nil {
		return m.AwsSecretAccessKey
	}
	return ""
}

func (m *S3ArtifactAccount) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func init() {
	proto.RegisterType((*S3ArtifactProvider)(nil), "proto.S3ArtifactProvider")
	proto.RegisterType((*S3ArtifactAccount)(nil), "proto.S3ArtifactAccount")
}

func init() { proto.RegisterFile("s3_artifact_provider.proto", fileDescriptor_2d1b7390c451dbc2) }

var fileDescriptor_2d1b7390c451dbc2 = []byte{
	// 236 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0xd9, 0xfe, 0x59, 0xdb, 0x29, 0x08, 0xce, 0x41, 0x82, 0x78, 0x58, 0x7a, 0x90, 0x9e,
	0xf6, 0xe0, 0xfa, 0x02, 0x7b, 0xf0, 0x20, 0x5e, 0x64, 0xfb, 0x00, 0x65, 0x9a, 0x8c, 0x12, 0xd0,
	0x24, 0x24, 0xd1, 0xe2, 0x93, 0xfa, 0x3a, 0xc2, 0xec, 0xda, 0x15, 0xed, 0x29, 0x99, 0xef, 0xf7,
	0x63, 0x3e, 0x06, 0xae, 0x52, 0xb3, 0xa3, 0x98, 0xed, 0x33, 0xe9, 0xbc, 0x0b, 0xd1, 0x7f, 0x58,
	0xc3, 0xb1, 0x0e, 0xd1, 0x67, 0x8f, 0x73, 0x79, 0xd6, 0x06, 0x70, 0xdb, 0xb4, 0x83, 0xf3, 0x34,
	0x28, 0xa8, 0xe0, 0x8c, 0x1d, 0xed, 0x5f, 0xd9, 0xa8, 0xa2, 0x2a, 0x36, 0x8b, 0xee, 0x67, 0xc4,
	0x3b, 0x58, 0x90, 0xd6, 0xfe, 0xdd, 0xe5, 0xa4, 0x26, 0xd5, 0x74, 0xb3, 0xba, 0x55, 0xfd, 0xc2,
	0x7a, 0x5c, 0xd3, 0xf6, 0x42, 0x77, 0x34, 0xd7, 0x5f, 0x05, 0x5c, 0xfc, 0xe3, 0x88, 0x30, 0x73,
	0xf4, 0xc6, 0x52, 0xb1, 0xec, 0xe4, 0x8f, 0x15, 0xac, 0x28, 0xd8, 0x7b, 0x67, 0x82, 0xb7, 0x2e,
	0xab, 0x89, 0xa0, 0xdf, 0x11, 0x5e, 0xc3, 0x92, 0x82, 0xed, 0xf8, 0xc5, 0x7a, 0xa7, 0xa6, 0xc2,
	0xc7, 0x00, 0x6f, 0xe0, 0x9c, 0x0e, 0xa9, 0xd5, 0x9a, 0x53, 0x7a, 0xe4, 0xcf, 0x07, 0xa3, 0x66,
	0xa2, 0xfc, 0x49, 0xb1, 0x06, 0xa4, 0x43, 0xda, 0xb2, 0x8e, 0x9c, 0x8f, 0xb9, 0x9a, 0x8b, 0x7b,
	0x82, 0xe0, 0x25, 0x94, 0xb1, 0xaf, 0x2c, 0xc5, 0x19, 0xa6, 0x7d, 0x29, 0xc7, 0x37, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x53, 0x78, 0x55, 0x07, 0x6b, 0x01, 0x00, 0x00,
}
