// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: cloudprovider/oracle.proto

package cloudprovider

import (
	proto "github.com/golang/protobuf/proto"
	client "github.com/spinnaker/kleat/api/client"
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

// Configuration for the Oracle provider.
type Oracle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured accounts.
	Accounts []*OracleAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	// The name of the primary account.
	PrimaryAccount string `protobuf:"bytes,3,opt,name=primaryAccount,proto3" json:"primaryAccount,omitempty"`
	// Configuration for Spinnaker's image bakery.
	BakeryDefaults *OracleBakeryDefaults `protobuf:"bytes,4,opt,name=bakeryDefaults,proto3" json:"bakeryDefaults,omitempty"`
}

func (x *Oracle) Reset() {
	*x = Oracle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_oracle_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Oracle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Oracle) ProtoMessage() {}

func (x *Oracle) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_oracle_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Oracle.ProtoReflect.Descriptor instead.
func (*Oracle) Descriptor() ([]byte, []int) {
	return file_cloudprovider_oracle_proto_rawDescGZIP(), []int{0}
}

func (x *Oracle) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *Oracle) GetAccounts() []*OracleAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

func (x *Oracle) GetPrimaryAccount() string {
	if x != nil {
		return x.PrimaryAccount
	}
	return ""
}

func (x *Oracle) GetBakeryDefaults() *OracleBakeryDefaults {
	if x != nil {
		return x.BakeryDefaults
	}
	return nil
}

// Configuration for an Oracle account. An account maps to an Oracle Cloud
// Infrastructure (OCI) user.
type OracleAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// (Deprecated) List of required Fiat permission groups. Configure
	// `permissions` instead.
	RequiredGroupMemberships []string `protobuf:"bytes,2,rep,name=requiredGroupMemberships,proto3" json:"requiredGroupMemberships,omitempty"`
	// Fiat permissions configuration.
	Permissions *client.Permissions `protobuf:"bytes,3,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// (Required) The OCID of the Oracle Compartment to use.
	CompartmentId string `protobuf:"bytes,4,opt,name=compartmentId,proto3" json:"compartmentId,omitempty"`
	// The environment name for the account. Many accounts can share the
	// same environment (e.g., dev, test, prod).
	Environment string `protobuf:"bytes,5,opt,name=environment,proto3" json:"environment,omitempty"`
	// (Required) Fingerprint of the public key.
	Fingerprint string `protobuf:"bytes,6,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
	// Passphrase used for the private key, if it is encrypted.
	PrivateKeyPassphrase string `protobuf:"bytes,7,opt,name=privateKeyPassphrase,proto3" json:"privateKeyPassphrase,omitempty"`
	// (Required) An Oracle region (e.g., `us-phoenix-1`).
	Region string `protobuf:"bytes,8,opt,name=region,proto3" json:"region,omitempty"`
	// (Required) Path to the private key in PEM format.
	SshPrivateKeyFilePath string `protobuf:"bytes,9,opt,name=sshPrivateKeyFilePath,proto3" json:"sshPrivateKeyFilePath,omitempty"`
	// (Required) The OCID of the Oracle Tenancy to use.
	TenancyId string `protobuf:"bytes,10,opt,name=tenancyId,proto3" json:"tenancyId,omitempty"`
	// (Required) The OCID of the Oracle User with which to authenticate.
	UserId string `protobuf:"bytes,11,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *OracleAccount) Reset() {
	*x = OracleAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_oracle_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OracleAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OracleAccount) ProtoMessage() {}

func (x *OracleAccount) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_oracle_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OracleAccount.ProtoReflect.Descriptor instead.
func (*OracleAccount) Descriptor() ([]byte, []int) {
	return file_cloudprovider_oracle_proto_rawDescGZIP(), []int{1}
}

func (x *OracleAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OracleAccount) GetRequiredGroupMemberships() []string {
	if x != nil {
		return x.RequiredGroupMemberships
	}
	return nil
}

func (x *OracleAccount) GetPermissions() *client.Permissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *OracleAccount) GetCompartmentId() string {
	if x != nil {
		return x.CompartmentId
	}
	return ""
}

func (x *OracleAccount) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *OracleAccount) GetFingerprint() string {
	if x != nil {
		return x.Fingerprint
	}
	return ""
}

func (x *OracleAccount) GetPrivateKeyPassphrase() string {
	if x != nil {
		return x.PrivateKeyPassphrase
	}
	return ""
}

func (x *OracleAccount) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *OracleAccount) GetSshPrivateKeyFilePath() string {
	if x != nil {
		return x.SshPrivateKeyFilePath
	}
	return ""
}

func (x *OracleAccount) GetTenancyId() string {
	if x != nil {
		return x.TenancyId
	}
	return ""
}

func (x *OracleAccount) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// Configuration for Spinnaker's image bakery.
type OracleBakeryDefaults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the Packer template that will be used to bake images from
	// this base image. The template file must be found in this list:
	// https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer,
	// or supplied as described here: https://spinnaker.io/setup/bakery/.
	TemplateFile string `protobuf:"bytes,1,opt,name=templateFile,proto3" json:"templateFile,omitempty"`
	// List of configured base images.
	BaseImages []*OracleBaseImageSettings `protobuf:"bytes,2,rep,name=baseImages,proto3" json:"baseImages,omitempty"`
	// (Required) The name of the Availability Domain within which a new instance
	// is launched and provisioned.
	AvailabilityDomain string `protobuf:"bytes,3,opt,name=availabilityDomain,proto3" json:"availabilityDomain,omitempty"`
	// (Required) The shape for a newly created instance.
	InstanceShape string `protobuf:"bytes,4,opt,name=instanceShape,proto3" json:"instanceShape,omitempty"`
	// (Required) The name of the subnet within which a new instance is launched
	// and provisioned.
	SubnetId string `protobuf:"bytes,5,opt,name=subnetId,proto3" json:"subnetId,omitempty"`
}

func (x *OracleBakeryDefaults) Reset() {
	*x = OracleBakeryDefaults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_oracle_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OracleBakeryDefaults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OracleBakeryDefaults) ProtoMessage() {}

func (x *OracleBakeryDefaults) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_oracle_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OracleBakeryDefaults.ProtoReflect.Descriptor instead.
func (*OracleBakeryDefaults) Descriptor() ([]byte, []int) {
	return file_cloudprovider_oracle_proto_rawDescGZIP(), []int{2}
}

func (x *OracleBakeryDefaults) GetTemplateFile() string {
	if x != nil {
		return x.TemplateFile
	}
	return ""
}

func (x *OracleBakeryDefaults) GetBaseImages() []*OracleBaseImageSettings {
	if x != nil {
		return x.BaseImages
	}
	return nil
}

func (x *OracleBakeryDefaults) GetAvailabilityDomain() string {
	if x != nil {
		return x.AvailabilityDomain
	}
	return ""
}

func (x *OracleBakeryDefaults) GetInstanceShape() string {
	if x != nil {
		return x.InstanceShape
	}
	return ""
}

func (x *OracleBakeryDefaults) GetSubnetId() string {
	if x != nil {
		return x.SubnetId
	}
	return ""
}

// Configuration for a base image for the Oracle provider's bakery.
type OracleBaseImageSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Oracle base image configuration.
	BaseImage *OracleBaseImage `protobuf:"bytes,1,opt,name=baseImage,proto3" json:"baseImage,omitempty"`
	// Oracle virtualization settings.
	VirtualizationSettings *OracleVirtualizationSettings `protobuf:"bytes,2,opt,name=virtualizationSettings,proto3" json:"virtualizationSettings,omitempty"`
}

func (x *OracleBaseImageSettings) Reset() {
	*x = OracleBaseImageSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_oracle_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OracleBaseImageSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OracleBaseImageSettings) ProtoMessage() {}

func (x *OracleBaseImageSettings) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_oracle_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OracleBaseImageSettings.ProtoReflect.Descriptor instead.
func (*OracleBaseImageSettings) Descriptor() ([]byte, []int) {
	return file_cloudprovider_oracle_proto_rawDescGZIP(), []int{3}
}

func (x *OracleBaseImageSettings) GetBaseImage() *OracleBaseImage {
	if x != nil {
		return x.BaseImage
	}
	return nil
}

func (x *OracleBaseImageSettings) GetVirtualizationSettings() *OracleVirtualizationSettings {
	if x != nil {
		return x.VirtualizationSettings
	}
	return nil
}

// Oracle base image configuration.
type OracleBaseImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the base image.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// A short description to help human operators identify the
	// image.
	ShortDescription string `protobuf:"bytes,2,opt,name=shortDescription,proto3" json:"shortDescription,omitempty"`
	// A long description to help human operators identify the
	// image.
	DetailedDescription string `protobuf:"bytes,3,opt,name=detailedDescription,proto3" json:"detailedDescription,omitempty"`
	//  This is used to help Spinnaker's bakery download the build
	//  artifacts you supply it with. For example, specifying deb
	//  indicates that your artifacts will need to be fetched from a
	//  debian repository.
	PackageType string `protobuf:"bytes,4,opt,name=packageType,proto3" json:"packageType,omitempty"`
	// The name of the Packer template that will be used to bake images from
	// this base image. The template file must be found in this list:
	// https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer,
	// or supplied as described here: https://spinnaker.io/setup/bakery/.
	TemplateFile string `protobuf:"bytes,5,opt,name=templateFile,proto3" json:"templateFile,omitempty"`
}

func (x *OracleBaseImage) Reset() {
	*x = OracleBaseImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_oracle_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OracleBaseImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OracleBaseImage) ProtoMessage() {}

func (x *OracleBaseImage) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_oracle_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OracleBaseImage.ProtoReflect.Descriptor instead.
func (*OracleBaseImage) Descriptor() ([]byte, []int) {
	return file_cloudprovider_oracle_proto_rawDescGZIP(), []int{4}
}

func (x *OracleBaseImage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OracleBaseImage) GetShortDescription() string {
	if x != nil {
		return x.ShortDescription
	}
	return ""
}

func (x *OracleBaseImage) GetDetailedDescription() string {
	if x != nil {
		return x.DetailedDescription
	}
	return ""
}

func (x *OracleBaseImage) GetPackageType() string {
	if x != nil {
		return x.PackageType
	}
	return ""
}

func (x *OracleBaseImage) GetTemplateFile() string {
	if x != nil {
		return x.TemplateFile
	}
	return ""
}

// Oracle virtualization settings.
type OracleVirtualizationSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The OCID of the base image ID for the baking configuration.
	BaseImageId string `protobuf:"bytes,1,opt,name=baseImageId,proto3" json:"baseImageId,omitempty"`
	// (Required) The ssh username for the baking configuration.
	SshUserName string `protobuf:"bytes,2,opt,name=sshUserName,proto3" json:"sshUserName,omitempty"`
}

func (x *OracleVirtualizationSettings) Reset() {
	*x = OracleVirtualizationSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_oracle_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OracleVirtualizationSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OracleVirtualizationSettings) ProtoMessage() {}

func (x *OracleVirtualizationSettings) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_oracle_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OracleVirtualizationSettings.ProtoReflect.Descriptor instead.
func (*OracleVirtualizationSettings) Descriptor() ([]byte, []int) {
	return file_cloudprovider_oracle_proto_rawDescGZIP(), []int{5}
}

func (x *OracleVirtualizationSettings) GetBaseImageId() string {
	if x != nil {
		return x.BaseImageId
	}
	return ""
}

func (x *OracleVirtualizationSettings) GetSshUserName() string {
	if x != nil {
		return x.SshUserName
	}
	return ""
}

var File_cloudprovider_oracle_proto protoreflect.FileDescriptor

var file_cloudprovider_oracle_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f,
	0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x1a, 0x11, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd, 0x01, 0x0a, 0x06, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x3e, 0x0a, 0x08, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x69,
	0x6d, 0x61, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x51, 0x0a, 0x0e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x79, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x42, 0x61, 0x6b, 0x65, 0x72, 0x79, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x73, 0x52, 0x0e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x79, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x73, 0x22, 0xb7, 0x03, 0x0a, 0x0d, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x18, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x18, 0x72, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x34, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x24, 0x0a, 0x0d,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x66, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x70, 0x72,
	0x69, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x66, 0x69, 0x6e, 0x67, 0x65,
	0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x14, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x73, 0x73, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x50, 0x61, 0x73, 0x73, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x12, 0x34, 0x0a, 0x15, 0x73, 0x73, 0x68, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65,
	0x4b, 0x65, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x15, 0x73, 0x73, 0x68, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x61,
	0x6e, 0x63, 0x79, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x6e,
	0x61, 0x6e, 0x63, 0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0xfa,
	0x01, 0x0a, 0x14, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x42, 0x61, 0x6b, 0x65, 0x72, 0x79, 0x44,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x4c, 0x0a, 0x0a, 0x62,
	0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x42, 0x61, 0x73, 0x65,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0a, 0x62,
	0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x2e, 0x0a, 0x12, 0x61, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x68, 0x61, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x53, 0x68, 0x61, 0x70, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x22, 0xc8, 0x01, 0x0a, 0x17,
	0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x42, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2e, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x69, 0x0a, 0x16, 0x76,
	0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x2e, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x16,
	0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0xc5, 0x01, 0x0a, 0x0f, 0x4f, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x68,
	0x6f, 0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x13, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x13, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x65,
	0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x62,
	0x0a, 0x1c, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x20,
	0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0b, 0x73, 0x73, 0x68, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x73, 0x68, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61,
	0x6d, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_cloudprovider_oracle_proto_rawDescOnce sync.Once
	file_cloudprovider_oracle_proto_rawDescData = file_cloudprovider_oracle_proto_rawDesc
)

func file_cloudprovider_oracle_proto_rawDescGZIP() []byte {
	file_cloudprovider_oracle_proto_rawDescOnce.Do(func() {
		file_cloudprovider_oracle_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudprovider_oracle_proto_rawDescData)
	})
	return file_cloudprovider_oracle_proto_rawDescData
}

var file_cloudprovider_oracle_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cloudprovider_oracle_proto_goTypes = []interface{}{
	(*Oracle)(nil),                       // 0: proto.cloudprovider.Oracle
	(*OracleAccount)(nil),                // 1: proto.cloudprovider.OracleAccount
	(*OracleBakeryDefaults)(nil),         // 2: proto.cloudprovider.OracleBakeryDefaults
	(*OracleBaseImageSettings)(nil),      // 3: proto.cloudprovider.OracleBaseImageSettings
	(*OracleBaseImage)(nil),              // 4: proto.cloudprovider.OracleBaseImage
	(*OracleVirtualizationSettings)(nil), // 5: proto.cloudprovider.OracleVirtualizationSettings
	(*client.Permissions)(nil),           // 6: proto.Permissions
}
var file_cloudprovider_oracle_proto_depIdxs = []int32{
	1, // 0: proto.cloudprovider.Oracle.accounts:type_name -> proto.cloudprovider.OracleAccount
	2, // 1: proto.cloudprovider.Oracle.bakeryDefaults:type_name -> proto.cloudprovider.OracleBakeryDefaults
	6, // 2: proto.cloudprovider.OracleAccount.permissions:type_name -> proto.Permissions
	3, // 3: proto.cloudprovider.OracleBakeryDefaults.baseImages:type_name -> proto.cloudprovider.OracleBaseImageSettings
	4, // 4: proto.cloudprovider.OracleBaseImageSettings.baseImage:type_name -> proto.cloudprovider.OracleBaseImage
	5, // 5: proto.cloudprovider.OracleBaseImageSettings.virtualizationSettings:type_name -> proto.cloudprovider.OracleVirtualizationSettings
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cloudprovider_oracle_proto_init() }
func file_cloudprovider_oracle_proto_init() {
	if File_cloudprovider_oracle_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloudprovider_oracle_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Oracle); i {
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
		file_cloudprovider_oracle_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OracleAccount); i {
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
		file_cloudprovider_oracle_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OracleBakeryDefaults); i {
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
		file_cloudprovider_oracle_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OracleBaseImageSettings); i {
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
		file_cloudprovider_oracle_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OracleBaseImage); i {
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
		file_cloudprovider_oracle_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OracleVirtualizationSettings); i {
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
			RawDescriptor: file_cloudprovider_oracle_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloudprovider_oracle_proto_goTypes,
		DependencyIndexes: file_cloudprovider_oracle_proto_depIdxs,
		MessageInfos:      file_cloudprovider_oracle_proto_msgTypes,
	}.Build()
	File_cloudprovider_oracle_proto = out.File
	file_cloudprovider_oracle_proto_rawDesc = nil
	file_cloudprovider_oracle_proto_goTypes = nil
	file_cloudprovider_oracle_proto_depIdxs = nil
}
