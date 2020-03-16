// Code generated by protoc-gen-go. DO NOT EDIT.
// source: clouddriver.proto

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

type ClouddriverConfig struct {
	Kubernetes           *KubernetesProvider     `protobuf:"bytes,1,opt,name=kubernetes,proto3" json:"kubernetes,omitempty"`
	Google               *GoogleProvider         `protobuf:"bytes,2,opt,name=google,proto3" json:"google,omitempty"`
	Appengine            *AppengineProvider      `protobuf:"bytes,3,opt,name=appengine,proto3" json:"appengine,omitempty"`
	Aws                  *AwsProvider            `protobuf:"bytes,4,opt,name=aws,proto3" json:"aws,omitempty"`
	Azure                *AzureProvider          `protobuf:"bytes,5,opt,name=azure,proto3" json:"azure,omitempty"`
	Cloudfoundry         *CloudFoundryProvider   `protobuf:"bytes,6,opt,name=cloudfoundry,proto3" json:"cloudfoundry,omitempty"`
	Dcos                 *DcosProvider           `protobuf:"bytes,7,opt,name=dcos,proto3" json:"dcos,omitempty"`
	DockerRegistry       *DockerRegistryProvider `protobuf:"bytes,8,opt,name=dockerRegistry,proto3" json:"dockerRegistry,omitempty"`
	Ecs                  *EcsProvider            `protobuf:"bytes,9,opt,name=ecs,proto3" json:"ecs,omitempty"`
	Huaweicloud          *HuaweiCloudProvider    `protobuf:"bytes,10,opt,name=huaweicloud,proto3" json:"huaweicloud,omitempty"`
	Oracle               *OracleProvider         `protobuf:"bytes,11,opt,name=oracle,proto3" json:"oracle,omitempty"`
	Artifacts            *ArtifactProviders      `protobuf:"bytes,12,opt,name=artifacts,proto3" json:"artifacts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ClouddriverConfig) Reset()         { *m = ClouddriverConfig{} }
func (m *ClouddriverConfig) String() string { return proto.CompactTextString(m) }
func (*ClouddriverConfig) ProtoMessage()    {}
func (*ClouddriverConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_0260082fb56e518f, []int{0}
}

func (m *ClouddriverConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClouddriverConfig.Unmarshal(m, b)
}
func (m *ClouddriverConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClouddriverConfig.Marshal(b, m, deterministic)
}
func (m *ClouddriverConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClouddriverConfig.Merge(m, src)
}
func (m *ClouddriverConfig) XXX_Size() int {
	return xxx_messageInfo_ClouddriverConfig.Size(m)
}
func (m *ClouddriverConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ClouddriverConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ClouddriverConfig proto.InternalMessageInfo

func (m *ClouddriverConfig) GetKubernetes() *KubernetesProvider {
	if m != nil {
		return m.Kubernetes
	}
	return nil
}

func (m *ClouddriverConfig) GetGoogle() *GoogleProvider {
	if m != nil {
		return m.Google
	}
	return nil
}

func (m *ClouddriverConfig) GetAppengine() *AppengineProvider {
	if m != nil {
		return m.Appengine
	}
	return nil
}

func (m *ClouddriverConfig) GetAws() *AwsProvider {
	if m != nil {
		return m.Aws
	}
	return nil
}

func (m *ClouddriverConfig) GetAzure() *AzureProvider {
	if m != nil {
		return m.Azure
	}
	return nil
}

func (m *ClouddriverConfig) GetCloudfoundry() *CloudFoundryProvider {
	if m != nil {
		return m.Cloudfoundry
	}
	return nil
}

func (m *ClouddriverConfig) GetDcos() *DcosProvider {
	if m != nil {
		return m.Dcos
	}
	return nil
}

func (m *ClouddriverConfig) GetDockerRegistry() *DockerRegistryProvider {
	if m != nil {
		return m.DockerRegistry
	}
	return nil
}

func (m *ClouddriverConfig) GetEcs() *EcsProvider {
	if m != nil {
		return m.Ecs
	}
	return nil
}

func (m *ClouddriverConfig) GetHuaweicloud() *HuaweiCloudProvider {
	if m != nil {
		return m.Huaweicloud
	}
	return nil
}

func (m *ClouddriverConfig) GetOracle() *OracleProvider {
	if m != nil {
		return m.Oracle
	}
	return nil
}

func (m *ClouddriverConfig) GetArtifacts() *ArtifactProviders {
	if m != nil {
		return m.Artifacts
	}
	return nil
}

func init() {
	proto.RegisterType((*ClouddriverConfig)(nil), "proto.ClouddriverConfig")
}

func init() { proto.RegisterFile("clouddriver.proto", fileDescriptor_0260082fb56e518f) }

var fileDescriptor_0260082fb56e518f = []byte{
	// 401 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x92, 0x51, 0x4b, 0xf3, 0x30,
	0x14, 0x86, 0xd9, 0xb7, 0x75, 0x9f, 0xcb, 0x86, 0xb2, 0x6c, 0x83, 0xac, 0x43, 0x11, 0x11, 0x14,
	0xc1, 0x5d, 0x28, 0x08, 0x82, 0x20, 0x32, 0xa7, 0x82, 0x17, 0xca, 0xfe, 0xc0, 0xc8, 0xd2, 0xac,
	0x96, 0x8d, 0xa6, 0xa4, 0xed, 0x8a, 0xde, 0xfa, 0xc7, 0xa5, 0x27, 0x69, 0xda, 0x2e, 0x57, 0x83,
	0x73, 0x9e, 0xa7, 0x4b, 0xde, 0x37, 0xa8, 0xcf, 0xb6, 0x22, 0xf5, 0x3c, 0x19, 0xec, 0xb8, 0x9c,
	0x46, 0x52, 0x24, 0x02, 0x3b, 0xf0, 0xe3, 0x8e, 0x37, 0xe9, 0x8a, 0xcb, 0x90, 0x27, 0x3c, 0x5e,
	0x46, 0x52, 0xec, 0x02, 0xaf, 0x20, 0xdc, 0x91, 0x2f, 0x84, 0xbf, 0xe5, 0xfb, 0x63, 0x42, 0xa3,
	0x88, 0x87, 0x7e, 0x10, 0x5a, 0x1b, 0x4c, 0x33, 0xeb, 0x23, 0x43, 0xfa, 0x93, 0x4a, 0x8b, 0x9c,
	0xc0, 0x79, 0xd6, 0x22, 0x0d, 0x3d, 0xf9, 0xbd, 0xbf, 0x1c, 0x78, 0x4c, 0x58, 0xdf, 0x39, 0xf1,
	0x04, 0xdb, 0x70, 0xb9, 0x94, 0xdc, 0x0f, 0xe2, 0xc4, 0x96, 0x30, 0x67, 0x96, 0xe3, 0x7e, 0xa5,
	0x34, 0xe3, 0x01, 0xfc, 0x97, 0x75, 0x39, 0x21, 0x29, 0xb3, 0x2f, 0x77, 0x44, 0x65, 0x12, 0xac,
	0x29, 0x4b, 0x62, 0x35, 0x38, 0xfb, 0x75, 0x50, 0x7f, 0x56, 0x86, 0x37, 0x13, 0xe1, 0x3a, 0xf0,
	0xf1, 0x3d, 0x42, 0x65, 0x6e, 0xa4, 0x71, 0xda, 0xb8, 0xec, 0xde, 0x8c, 0x95, 0x31, 0x7d, 0x37,
	0x8b, 0x4f, 0xfd, 0xed, 0x45, 0x05, 0xc6, 0xd7, 0xa8, 0xad, 0x72, 0x25, 0xff, 0x40, 0x1b, 0x69,
	0xed, 0x15, 0x86, 0x46, 0xd1, 0x10, 0xbe, 0x43, 0x1d, 0x93, 0x37, 0x69, 0x82, 0x41, 0xb4, 0xf1,
	0x54, 0xcc, 0x8d, 0x54, 0xa2, 0xf8, 0x1c, 0x35, 0x69, 0x16, 0x93, 0x16, 0x18, 0xb8, 0x30, 0xb2,
	0xf2, 0x4c, 0xf9, 0x1a, 0x5f, 0x21, 0x07, 0xfa, 0x21, 0x0e, 0x70, 0xc3, 0x82, 0xcb, 0x67, 0x86,
	0x54, 0x08, 0x7e, 0x44, 0xbd, 0x6a, 0x6b, 0xa4, 0x0d, 0xca, 0x44, 0x2b, 0x90, 0xd1, 0x8b, 0x5a,
	0x19, 0xb3, 0x26, 0xe0, 0x0b, 0xd4, 0xca, 0x9b, 0x25, 0xff, 0x41, 0x1c, 0x68, 0xf1, 0x99, 0x89,
	0xf2, 0x50, 0x00, 0xe0, 0x39, 0x3a, 0x54, 0x6d, 0x2f, 0x74, 0xd9, 0xe4, 0x00, 0x94, 0xe3, 0x42,
	0xa9, 0x2d, 0x8d, 0xbc, 0x27, 0xe5, 0x11, 0x70, 0x16, 0x93, 0x4e, 0x2d, 0x82, 0x39, 0xab, 0x44,
	0xc0, 0x59, 0x8c, 0x1f, 0x50, 0xb7, 0xf2, 0x4c, 0x08, 0x02, 0xda, 0xd5, 0xf4, 0x1b, 0x6c, 0xe0,
	0x6e, 0xc6, 0xaa, 0xe2, 0x79, 0x9b, 0xea, 0x21, 0x91, 0x6e, 0xad, 0xcd, 0x0f, 0x18, 0x96, 0x6d,
	0x2a, 0x08, 0xda, 0x2c, 0x1e, 0x18, 0xe9, 0xd5, 0xdb, 0xd4, 0xf3, 0xc2, 0x89, 0x17, 0x25, 0xba,
	0x6a, 0x03, 0x73, 0xfb, 0x17, 0x00, 0x00, 0xff, 0xff, 0x28, 0x25, 0x68, 0x86, 0xc8, 0x03, 0x00,
	0x00,
}
