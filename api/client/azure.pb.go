// Code generated by protoc-gen-go. DO NOT EDIT.
// source: azure.proto

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

// Configuration for the Azure provider.
type Azure struct {
	// Whether the provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured accounts.
	Accounts []*AzureAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	// The name of the primary account.
	PrimaryAccount string `protobuf:"bytes,3,opt,name=primaryAccount,proto3" json:"primaryAccount,omitempty"`
	// Configuration for Spinnaker's image bakery.
	BakeryDefaults       *AzureBakeryDefaults `protobuf:"bytes,4,opt,name=bakeryDefaults,proto3" json:"bakeryDefaults,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Azure) Reset()         { *m = Azure{} }
func (m *Azure) String() string { return proto.CompactTextString(m) }
func (*Azure) ProtoMessage()    {}
func (*Azure) Descriptor() ([]byte, []int) {
	return fileDescriptor_a258c4f0d0c4b1e7, []int{0}
}

func (m *Azure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Azure.Unmarshal(m, b)
}
func (m *Azure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Azure.Marshal(b, m, deterministic)
}
func (m *Azure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Azure.Merge(m, src)
}
func (m *Azure) XXX_Size() int {
	return xxx_messageInfo_Azure.Size(m)
}
func (m *Azure) XXX_DiscardUnknown() {
	xxx_messageInfo_Azure.DiscardUnknown(m)
}

var xxx_messageInfo_Azure proto.InternalMessageInfo

func (m *Azure) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Azure) GetAccounts() []*AzureAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
}

func (m *Azure) GetPrimaryAccount() string {
	if m != nil {
		return m.PrimaryAccount
	}
	return ""
}

func (m *Azure) GetBakeryDefaults() *AzureBakeryDefaults {
	if m != nil {
		return m.BakeryDefaults
	}
	return nil
}

type AzureAccount struct {
	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// (Required) The `appKey` (password) of your service principal.
	AppKey string `protobuf:"bytes,2,opt,name=appKey,proto3" json:"appKey,omitempty"`
	// (Required) The `clientId` (also called `appId`) of your service principal.
	ClientId string `protobuf:"bytes,3,opt,name=clientId,proto3" json:"clientId,omitempty"`
	// (Required) The name of a KeyVault that contains the user name, password,
	// and ssh public key used to create VMs
	DefaultKeyVault string `protobuf:"bytes,4,opt,name=defaultKeyVault,proto3" json:"defaultKeyVault,omitempty"`
	// (Required) The default resource group to contain any non-application
	// specific resources.
	DefaultResourceGroup string `protobuf:"bytes,5,opt,name=defaultResourceGroup,proto3" json:"defaultResourceGroup,omitempty"`
	// The environment name for the account. Many accounts can share the
	// same environment (e.g., dev, test, prod).
	Environment string `protobuf:"bytes,6,opt,name=environment,proto3" json:"environment,omitempty"`
	// The `objectId` of your service principal. This is only required if using
	// Packer to bake Windows images.
	ObjectId string `protobuf:"bytes,7,opt,name=objectId,proto3" json:"objectId,omitempty"`
	// The resource group to use if baking images with Packer.
	PackerResourceGroup string `protobuf:"bytes,8,opt,name=packerResourceGroup,proto3" json:"packerResourceGroup,omitempty"`
	// The storage account to use if baking images with Packer.
	PackerStorageAccount string `protobuf:"bytes,9,opt,name=packerStorageAccount,proto3" json:"packerStorageAccount,omitempty"`
	// Fiat permissions configuration.
	Permissions *Permissions `protobuf:"bytes,10,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// (Deprecated): List of required Fiat permission groups. Configure
	// `permissions` instead.
	RequiredGroupMemberships []string `protobuf:"bytes,11,rep,name=requiredGroupMemberships,proto3" json:"requiredGroupMemberships,omitempty"`
	// The Azure regions this Spinnaker account will manage.
	// TODO(mneterval): Halyard defaults to `[westus, eastus]`. Move to
	// Clouddriver.
	Regions []string `protobuf:"bytes,12,rep,name=regions,proto3" json:"regions,omitempty"`
	// (Required) The `subscriptionId` to which your service principal is
	// assigned.
	SubscriptionId string `protobuf:"bytes,13,opt,name=subscriptionId,proto3" json:"subscriptionId,omitempty"`
	// (Required) The `tenantId` to which your service principal is assigned.
	TenantId string `protobuf:"bytes,14,opt,name=tenantId,proto3" json:"tenantId,omitempty"`
	// If true, the  SSH public key is used to provision the linux VM.
	// If false, the password is used instead.
	// TODO(mneterval): Halyard defaults to `true`. Move to Clouddriver.
	UseSshPublicKey      bool     `protobuf:"varint,15,opt,name=useSshPublicKey,proto3" json:"useSshPublicKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AzureAccount) Reset()         { *m = AzureAccount{} }
func (m *AzureAccount) String() string { return proto.CompactTextString(m) }
func (*AzureAccount) ProtoMessage()    {}
func (*AzureAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_a258c4f0d0c4b1e7, []int{1}
}

func (m *AzureAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureAccount.Unmarshal(m, b)
}
func (m *AzureAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureAccount.Marshal(b, m, deterministic)
}
func (m *AzureAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureAccount.Merge(m, src)
}
func (m *AzureAccount) XXX_Size() int {
	return xxx_messageInfo_AzureAccount.Size(m)
}
func (m *AzureAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureAccount.DiscardUnknown(m)
}

var xxx_messageInfo_AzureAccount proto.InternalMessageInfo

func (m *AzureAccount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AzureAccount) GetAppKey() string {
	if m != nil {
		return m.AppKey
	}
	return ""
}

func (m *AzureAccount) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *AzureAccount) GetDefaultKeyVault() string {
	if m != nil {
		return m.DefaultKeyVault
	}
	return ""
}

func (m *AzureAccount) GetDefaultResourceGroup() string {
	if m != nil {
		return m.DefaultResourceGroup
	}
	return ""
}

func (m *AzureAccount) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *AzureAccount) GetObjectId() string {
	if m != nil {
		return m.ObjectId
	}
	return ""
}

func (m *AzureAccount) GetPackerResourceGroup() string {
	if m != nil {
		return m.PackerResourceGroup
	}
	return ""
}

func (m *AzureAccount) GetPackerStorageAccount() string {
	if m != nil {
		return m.PackerStorageAccount
	}
	return ""
}

func (m *AzureAccount) GetPermissions() *Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *AzureAccount) GetRequiredGroupMemberships() []string {
	if m != nil {
		return m.RequiredGroupMemberships
	}
	return nil
}

func (m *AzureAccount) GetRegions() []string {
	if m != nil {
		return m.Regions
	}
	return nil
}

func (m *AzureAccount) GetSubscriptionId() string {
	if m != nil {
		return m.SubscriptionId
	}
	return ""
}

func (m *AzureAccount) GetTenantId() string {
	if m != nil {
		return m.TenantId
	}
	return ""
}

func (m *AzureAccount) GetUseSshPublicKey() bool {
	if m != nil {
		return m.UseSshPublicKey
	}
	return false
}

// Configuration for Spinnaker's image bakery.
type AzureBakeryDefaults struct {
	// List of configured base images.
	BaseImages           []*AzureBaseImageSettings `protobuf:"bytes,1,rep,name=baseImages,proto3" json:"baseImages,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *AzureBakeryDefaults) Reset()         { *m = AzureBakeryDefaults{} }
func (m *AzureBakeryDefaults) String() string { return proto.CompactTextString(m) }
func (*AzureBakeryDefaults) ProtoMessage()    {}
func (*AzureBakeryDefaults) Descriptor() ([]byte, []int) {
	return fileDescriptor_a258c4f0d0c4b1e7, []int{2}
}

func (m *AzureBakeryDefaults) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureBakeryDefaults.Unmarshal(m, b)
}
func (m *AzureBakeryDefaults) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureBakeryDefaults.Marshal(b, m, deterministic)
}
func (m *AzureBakeryDefaults) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureBakeryDefaults.Merge(m, src)
}
func (m *AzureBakeryDefaults) XXX_Size() int {
	return xxx_messageInfo_AzureBakeryDefaults.Size(m)
}
func (m *AzureBakeryDefaults) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureBakeryDefaults.DiscardUnknown(m)
}

var xxx_messageInfo_AzureBakeryDefaults proto.InternalMessageInfo

func (m *AzureBakeryDefaults) GetBaseImages() []*AzureBaseImageSettings {
	if m != nil {
		return m.BaseImages
	}
	return nil
}

// Configuration for a base image for the Azure provider's bakery.
type AzureBaseImageSettings struct {
	// Base image configuration.
	BaseImage            *AzureBaseImage `protobuf:"bytes,1,opt,name=baseImage,proto3" json:"baseImage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AzureBaseImageSettings) Reset()         { *m = AzureBaseImageSettings{} }
func (m *AzureBaseImageSettings) String() string { return proto.CompactTextString(m) }
func (*AzureBaseImageSettings) ProtoMessage()    {}
func (*AzureBaseImageSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_a258c4f0d0c4b1e7, []int{3}
}

func (m *AzureBaseImageSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureBaseImageSettings.Unmarshal(m, b)
}
func (m *AzureBaseImageSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureBaseImageSettings.Marshal(b, m, deterministic)
}
func (m *AzureBaseImageSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureBaseImageSettings.Merge(m, src)
}
func (m *AzureBaseImageSettings) XXX_Size() int {
	return xxx_messageInfo_AzureBaseImageSettings.Size(m)
}
func (m *AzureBaseImageSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureBaseImageSettings.DiscardUnknown(m)
}

var xxx_messageInfo_AzureBaseImageSettings proto.InternalMessageInfo

func (m *AzureBaseImageSettings) GetBaseImage() *AzureBaseImage {
	if m != nil {
		return m.BaseImage
	}
	return nil
}

// Base image configuration.
type AzureBaseImage struct {
	// A short description to help human operators identify the
	// image.
	ShortDescription string `protobuf:"bytes,1,opt,name=shortDescription,proto3" json:"shortDescription,omitempty"`
	// A long description to help human operators identify the
	// image.
	DetailedDescription string `protobuf:"bytes,2,opt,name=detailedDescription,proto3" json:"detailedDescription,omitempty"`
	// (Required) The Publisher name for your base image. See
	// https://aka.ms/azspinimage to get a list of images.
	Publisher string `protobuf:"bytes,3,opt,name=publisher,proto3" json:"publisher,omitempty"`
	// (Required) The offer for your base image. See https://aka.ms/azspinimage
	// to get a list of images.
	Offer string `protobuf:"bytes,4,opt,name=offer,proto3" json:"offer,omitempty"`
	// (Required) The SKU for your base image. See https://aka.ms/azspinimage to
	// get a list of images.
	Sku string `protobuf:"bytes,5,opt,name=sku,proto3" json:"sku,omitempty"`
	// The version of your base image. This defaults to `latest` if not
	// specified.
	Version string `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
	// This is used to help Spinnaker's bakery download the build artifacts you
	// supply it with. For example, specifying `deb` indicates that your
	// artifacts will need to be fetched from a debian repository.
	PackageType string `protobuf:"bytes,7,opt,name=packageType,proto3" json:"packageType,omitempty"`
	// This is the name of the packer template that will be used to bake images
	// from this base image. The template file must be found in this list:
	// https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer,
	// or supplied as described here: https://spinnaker.io/setup/bakery/.
	TemplateFile         string   `protobuf:"bytes,8,opt,name=templateFile,proto3" json:"templateFile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AzureBaseImage) Reset()         { *m = AzureBaseImage{} }
func (m *AzureBaseImage) String() string { return proto.CompactTextString(m) }
func (*AzureBaseImage) ProtoMessage()    {}
func (*AzureBaseImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_a258c4f0d0c4b1e7, []int{4}
}

func (m *AzureBaseImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureBaseImage.Unmarshal(m, b)
}
func (m *AzureBaseImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureBaseImage.Marshal(b, m, deterministic)
}
func (m *AzureBaseImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureBaseImage.Merge(m, src)
}
func (m *AzureBaseImage) XXX_Size() int {
	return xxx_messageInfo_AzureBaseImage.Size(m)
}
func (m *AzureBaseImage) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureBaseImage.DiscardUnknown(m)
}

var xxx_messageInfo_AzureBaseImage proto.InternalMessageInfo

func (m *AzureBaseImage) GetShortDescription() string {
	if m != nil {
		return m.ShortDescription
	}
	return ""
}

func (m *AzureBaseImage) GetDetailedDescription() string {
	if m != nil {
		return m.DetailedDescription
	}
	return ""
}

func (m *AzureBaseImage) GetPublisher() string {
	if m != nil {
		return m.Publisher
	}
	return ""
}

func (m *AzureBaseImage) GetOffer() string {
	if m != nil {
		return m.Offer
	}
	return ""
}

func (m *AzureBaseImage) GetSku() string {
	if m != nil {
		return m.Sku
	}
	return ""
}

func (m *AzureBaseImage) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *AzureBaseImage) GetPackageType() string {
	if m != nil {
		return m.PackageType
	}
	return ""
}

func (m *AzureBaseImage) GetTemplateFile() string {
	if m != nil {
		return m.TemplateFile
	}
	return ""
}

func init() {
	proto.RegisterType((*Azure)(nil), "proto.Azure")
	proto.RegisterType((*AzureAccount)(nil), "proto.AzureAccount")
	proto.RegisterType((*AzureBakeryDefaults)(nil), "proto.AzureBakeryDefaults")
	proto.RegisterType((*AzureBaseImageSettings)(nil), "proto.AzureBaseImageSettings")
	proto.RegisterType((*AzureBaseImage)(nil), "proto.AzureBaseImage")
}

func init() { proto.RegisterFile("azure.proto", fileDescriptor_a258c4f0d0c4b1e7) }

var fileDescriptor_a258c4f0d0c4b1e7 = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0xd1, 0x4e, 0xdb, 0x30,
	0x14, 0x55, 0x29, 0x85, 0xe6, 0x86, 0x15, 0x66, 0x18, 0xb2, 0xd0, 0x26, 0x55, 0x7d, 0x98, 0xa2,
	0x3d, 0xb0, 0xa9, 0xec, 0x69, 0xd2, 0x1e, 0x40, 0x68, 0x13, 0x42, 0x48, 0x28, 0xa0, 0xbd, 0x3b,
	0xc9, 0xa5, 0xf5, 0x48, 0x62, 0xcf, 0x76, 0x90, 0xba, 0xef, 0xd8, 0xf7, 0xec, 0x23, 0xf6, 0x45,
	0x93, 0x1d, 0xa7, 0x24, 0x5d, 0xf6, 0x94, 0xdc, 0x73, 0x8e, 0xe3, 0x93, 0x7b, 0xcf, 0x85, 0x90,
	0xfd, 0xac, 0x14, 0x9e, 0x4a, 0x25, 0x8c, 0x20, 0x23, 0xf7, 0x38, 0x79, 0x29, 0x51, 0x15, 0x5c,
	0x6b, 0x2e, 0x4a, 0x5d, 0x33, 0xb3, 0xdf, 0x03, 0x18, 0x9d, 0x5b, 0x25, 0xa1, 0xb0, 0x8b, 0x25,
	0x4b, 0x72, 0xcc, 0xe8, 0x60, 0x3a, 0x88, 0xc6, 0x71, 0x53, 0x92, 0xf7, 0x30, 0x66, 0x69, 0x2a,
	0xaa, 0xd2, 0x68, 0xba, 0x35, 0x1d, 0x46, 0xe1, 0xfc, 0xb0, 0x3e, 0x7d, 0xea, 0x4e, 0x9e, 0xd7,
	0x5c, 0xbc, 0x16, 0x91, 0xb7, 0x30, 0x91, 0x8a, 0x17, 0x4c, 0xad, 0x3c, 0x47, 0x87, 0xd3, 0x41,
	0x14, 0xc4, 0x1b, 0x28, 0xb9, 0x80, 0x49, 0xc2, 0x1e, 0x51, 0xad, 0x2e, 0xf1, 0x81, 0x55, 0xb9,
	0xd1, 0x74, 0x7b, 0x3a, 0x88, 0xc2, 0xf9, 0x49, 0xfb, 0xf3, 0x17, 0x1d, 0x45, 0xbc, 0x71, 0x62,
	0xf6, 0x67, 0x1b, 0xf6, 0xda, 0x36, 0x08, 0x81, 0xed, 0x92, 0x15, 0xe8, 0x7e, 0x22, 0x88, 0xdd,
	0x3b, 0x39, 0x86, 0x1d, 0x26, 0xe5, 0x35, 0xae, 0xe8, 0x96, 0x43, 0x7d, 0x45, 0x4e, 0x60, 0x9c,
	0xe6, 0x1c, 0x4b, 0x73, 0x95, 0x79, 0x8b, 0xeb, 0x9a, 0x44, 0xb0, 0x9f, 0xd5, 0x97, 0x5c, 0xe3,
	0xea, 0x9b, 0x7d, 0x3a, 0x77, 0x41, 0xbc, 0x09, 0x93, 0x39, 0x1c, 0x79, 0x28, 0x46, 0x2d, 0x2a,
	0x95, 0xe2, 0x57, 0x25, 0x2a, 0x49, 0x47, 0x4e, 0xde, 0xcb, 0x91, 0x29, 0x84, 0x58, 0x3e, 0x71,
	0x25, 0xca, 0x02, 0x4b, 0x43, 0x77, 0x9c, 0xb4, 0x0d, 0x59, 0x6f, 0x22, 0xf9, 0x8e, 0xa9, 0xf5,
	0xb6, 0x5b, 0x7b, 0x6b, 0x6a, 0xf2, 0x01, 0x0e, 0x25, 0x4b, 0x1f, 0x51, 0x75, 0x2f, 0x1c, 0x3b,
	0x59, 0x1f, 0x65, 0x3d, 0xd6, 0xf0, 0x9d, 0x11, 0x8a, 0x2d, 0x9a, 0x6e, 0xd1, 0xa0, 0xf6, 0xd8,
	0xc7, 0x91, 0x8f, 0x10, 0xb6, 0x02, 0x43, 0xc1, 0xcd, 0x86, 0xf8, 0xd9, 0xdc, 0x3e, 0x33, 0x71,
	0x5b, 0x46, 0x3e, 0x01, 0x55, 0xf8, 0xa3, 0xe2, 0x0a, 0x33, 0x77, 0xf5, 0x0d, 0x16, 0x09, 0x2a,
	0xbd, 0xe4, 0x52, 0xd3, 0x70, 0x3a, 0x8c, 0x82, 0xf8, 0xbf, 0xbc, 0xcd, 0xa0, 0xc2, 0x85, 0xbb,
	0x6d, 0xcf, 0x49, 0x9b, 0xd2, 0x46, 0x4a, 0x57, 0x89, 0x4e, 0x15, 0x97, 0x86, 0x8b, 0xf2, 0x2a,
	0xa3, 0x2f, 0xea, 0x48, 0x75, 0x51, 0xdb, 0x35, 0x83, 0x25, 0x73, 0x13, 0x9d, 0xd4, 0x5d, 0x6b,
	0x6a, 0x3b, 0xd1, 0x4a, 0xe3, 0x9d, 0x5e, 0xde, 0x56, 0x49, 0xce, 0x53, 0x1b, 0x87, 0x7d, 0x97,
	0xf4, 0x4d, 0x78, 0x76, 0x0f, 0x87, 0x3d, 0xd9, 0x23, 0x9f, 0x01, 0x12, 0xa6, 0xf1, 0xaa, 0x60,
	0x0b, 0xd4, 0x74, 0xe0, 0x56, 0xe1, 0x4d, 0x37, 0xab, 0x9e, 0xbd, 0x43, 0x63, 0x78, 0xb9, 0xd0,
	0x71, 0xeb, 0xc0, 0xec, 0x06, 0x8e, 0xfb, 0x55, 0xe4, 0x0c, 0x82, 0xb5, 0xce, 0x05, 0x37, 0x9c,
	0xbf, 0xea, 0xfd, 0x6e, 0xfc, 0xac, 0x9b, 0xfd, 0xda, 0x82, 0x49, 0x97, 0x25, 0xef, 0xe0, 0x40,
	0x2f, 0x85, 0x32, 0x97, 0xb8, 0xee, 0x89, 0xdf, 0x83, 0x7f, 0x70, 0x9b, 0xa1, 0x0c, 0x0d, 0xe3,
	0x39, 0x66, 0x6d, 0x79, 0xbd, 0x20, 0x7d, 0x14, 0x79, 0x0d, 0x81, 0xb4, 0x2d, 0xd2, 0x4b, 0x54,
	0x7e, 0x5d, 0x9e, 0x01, 0x72, 0x04, 0x23, 0xf1, 0xf0, 0x80, 0xca, 0x6f, 0x49, 0x5d, 0x90, 0x03,
	0x18, 0xea, 0xc7, 0xca, 0xaf, 0x82, 0x7d, 0xb5, 0x33, 0x7e, 0x42, 0x65, 0xb3, 0xe2, 0x53, 0xdf,
	0x94, 0x76, 0x27, 0x6c, 0x0e, 0xd9, 0x02, 0xef, 0x57, 0x12, 0x7d, 0xe8, 0xdb, 0x10, 0x99, 0xc1,
	0x9e, 0xc1, 0x42, 0xe6, 0xcc, 0xe0, 0x17, 0x9e, 0xa3, 0x0f, 0x7c, 0x07, 0x4b, 0x76, 0x5c, 0xdf,
	0xce, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x29, 0x1c, 0xe3, 0x83, 0x01, 0x05, 0x00, 0x00,
}
