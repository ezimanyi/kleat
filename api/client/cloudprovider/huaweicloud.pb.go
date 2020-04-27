// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: cloudprovider/huaweicloud.proto

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

// Configuration for the Huawei Cloud provider.
type HuaweiCloud struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured accounts.
	Accounts []*HuaweiCloudAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	// The name of the primary account.
	PrimaryAccount string `protobuf:"bytes,3,opt,name=primaryAccount,proto3" json:"primaryAccount,omitempty"`
	// Configuration for Spinnaker's image bakery.
	BakeryDefaults *HuaweiCloudBakeryDefaults `protobuf:"bytes,4,opt,name=bakeryDefaults,proto3" json:"bakeryDefaults,omitempty"`
}

func (x *HuaweiCloud) Reset() {
	*x = HuaweiCloud{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_huaweicloud_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HuaweiCloud) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HuaweiCloud) ProtoMessage() {}

func (x *HuaweiCloud) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_huaweicloud_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HuaweiCloud.ProtoReflect.Descriptor instead.
func (*HuaweiCloud) Descriptor() ([]byte, []int) {
	return file_cloudprovider_huaweicloud_proto_rawDescGZIP(), []int{0}
}

func (x *HuaweiCloud) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *HuaweiCloud) GetAccounts() []*HuaweiCloudAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

func (x *HuaweiCloud) GetPrimaryAccount() string {
	if x != nil {
		return x.PrimaryAccount
	}
	return ""
}

func (x *HuaweiCloud) GetBakeryDefaults() *HuaweiCloudBakeryDefaults {
	if x != nil {
		return x.BakeryDefaults
	}
	return nil
}

// Configuration for a Huawei Cloud account.
type HuaweiCloudAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The type of account.
	AccountType string `protobuf:"bytes,2,opt,name=accountType,proto3" json:"accountType,omitempty"`
	// (Deprecated) List of required Fiat permission groups. Configure
	// `permissions` instead.
	RequiredGroupMemberships []string `protobuf:"bytes,3,rep,name=requiredGroupMemberships,proto3" json:"requiredGroupMemberships,omitempty"`
	// Fiat permissions configuration.
	Permissions *client.Permissions `protobuf:"bytes,4,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// (Required) The auth URL of the cloud.
	AuthUrl string `protobuf:"bytes,5,opt,name=authUrl,proto3" json:"authUrl,omitempty"`
	// (Required) The domain name of the cloud.
	DomainName string `protobuf:"bytes,6,opt,name=domainName,proto3" json:"domainName,omitempty"`
	// The environment name for the account. Many accounts can share the
	// same environment (e.g., dev, test, prod).
	Environment string `protobuf:"bytes,7,opt,name=environment,proto3" json:"environment,omitempty"`
	// If `true`, disables certificate validation on SSL connections. Needed if
	// certificates are self-signed. Defaults to `false`.
	Insecure bool `protobuf:"varint,8,opt,name=insecure,proto3" json:"insecure,omitempty"`
	// (Required) The password used to access the cloud.
	Password string `protobuf:"bytes,9,opt,name=password,proto3" json:"password,omitempty"`
	// (Required) The name of the project within the cloud.
	ProjectName string `protobuf:"bytes,10,opt,name=projectName,proto3" json:"projectName,omitempty"`
	// (Required) The region(s) of the cloud.
	Regions []string `protobuf:"bytes,11,rep,name=regions,proto3" json:"regions,omitempty"`
	// (Required) The username used to access the cloud.
	Username string `protobuf:"bytes,12,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *HuaweiCloudAccount) Reset() {
	*x = HuaweiCloudAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_huaweicloud_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HuaweiCloudAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HuaweiCloudAccount) ProtoMessage() {}

func (x *HuaweiCloudAccount) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_huaweicloud_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HuaweiCloudAccount.ProtoReflect.Descriptor instead.
func (*HuaweiCloudAccount) Descriptor() ([]byte, []int) {
	return file_cloudprovider_huaweicloud_proto_rawDescGZIP(), []int{1}
}

func (x *HuaweiCloudAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HuaweiCloudAccount) GetAccountType() string {
	if x != nil {
		return x.AccountType
	}
	return ""
}

func (x *HuaweiCloudAccount) GetRequiredGroupMemberships() []string {
	if x != nil {
		return x.RequiredGroupMemberships
	}
	return nil
}

func (x *HuaweiCloudAccount) GetPermissions() *client.Permissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *HuaweiCloudAccount) GetAuthUrl() string {
	if x != nil {
		return x.AuthUrl
	}
	return ""
}

func (x *HuaweiCloudAccount) GetDomainName() string {
	if x != nil {
		return x.DomainName
	}
	return ""
}

func (x *HuaweiCloudAccount) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *HuaweiCloudAccount) GetInsecure() bool {
	if x != nil {
		return x.Insecure
	}
	return false
}

func (x *HuaweiCloudAccount) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *HuaweiCloudAccount) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *HuaweiCloudAccount) GetRegions() []string {
	if x != nil {
		return x.Regions
	}
	return nil
}

func (x *HuaweiCloudAccount) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

// Configuration for Spinnaker's image bakery.
type HuaweiCloudBakeryDefaults struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of configured base images.
	BaseImages []*HuaweiCloudBaseImageSettings `protobuf:"bytes,1,rep,name=baseImages,proto3" json:"baseImages,omitempty"`
	// This is the name of the packer template that will be used to bake images
	// from this base image. The template file must be found in this list:
	// https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer, or
	// supplied as described here: https://spinnaker.io/setup/bakery/.
	TemplateFile string `protobuf:"bytes,2,opt,name=templateFile,proto3" json:"templateFile,omitempty"`
	// (Required) The default auth URL in which images will be baked.
	AuthUrl string `protobuf:"bytes,3,opt,name=authUrl,proto3" json:"authUrl,omitempty"`
	// (Required) The default username with which images will be baked.
	Username string `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	// (Required) The default password with which images will be baked.
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	// The name of the default project in which images will be baked.
	ProjectName string `protobuf:"bytes,6,opt,name=projectName,proto3" json:"projectName,omitempty"`
	// (Required) The default domain name in which images will be baked.
	DomainName string `protobuf:"bytes,7,opt,name=domainName,proto3" json:"domainName,omitempty"`
	// The security setting for connecting to the Huawei Cloud account. Defaults
	// to `false`.
	Insecure bool `protobuf:"varint,8,opt,name=insecure,proto3" json:"insecure,omitempty"`
	// (Required) The VPC in which images will be baked.
	VpcId string `protobuf:"bytes,9,opt,name=vpcId,proto3" json:"vpcId,omitempty"`
	//  (Required) The subnet in which images will be baked.
	SubnetId string `protobuf:"bytes,10,opt,name=subnetId,proto3" json:"subnetId,omitempty"`
	// (Required) The default security group in which images will be baked.
	SecurityGroup string `protobuf:"bytes,11,opt,name=securityGroup,proto3" json:"securityGroup,omitempty"`
	// (Required) The bandwidth size of EIP in which images will be baked.
	EipBandwidthSize int32 `protobuf:"varint,12,opt,name=eipBandwidthSize,proto3" json:"eipBandwidthSize,omitempty"`
}

func (x *HuaweiCloudBakeryDefaults) Reset() {
	*x = HuaweiCloudBakeryDefaults{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_huaweicloud_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HuaweiCloudBakeryDefaults) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HuaweiCloudBakeryDefaults) ProtoMessage() {}

func (x *HuaweiCloudBakeryDefaults) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_huaweicloud_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HuaweiCloudBakeryDefaults.ProtoReflect.Descriptor instead.
func (*HuaweiCloudBakeryDefaults) Descriptor() ([]byte, []int) {
	return file_cloudprovider_huaweicloud_proto_rawDescGZIP(), []int{2}
}

func (x *HuaweiCloudBakeryDefaults) GetBaseImages() []*HuaweiCloudBaseImageSettings {
	if x != nil {
		return x.BaseImages
	}
	return nil
}

func (x *HuaweiCloudBakeryDefaults) GetTemplateFile() string {
	if x != nil {
		return x.TemplateFile
	}
	return ""
}

func (x *HuaweiCloudBakeryDefaults) GetAuthUrl() string {
	if x != nil {
		return x.AuthUrl
	}
	return ""
}

func (x *HuaweiCloudBakeryDefaults) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *HuaweiCloudBakeryDefaults) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *HuaweiCloudBakeryDefaults) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *HuaweiCloudBakeryDefaults) GetDomainName() string {
	if x != nil {
		return x.DomainName
	}
	return ""
}

func (x *HuaweiCloudBakeryDefaults) GetInsecure() bool {
	if x != nil {
		return x.Insecure
	}
	return false
}

func (x *HuaweiCloudBakeryDefaults) GetVpcId() string {
	if x != nil {
		return x.VpcId
	}
	return ""
}

func (x *HuaweiCloudBakeryDefaults) GetSubnetId() string {
	if x != nil {
		return x.SubnetId
	}
	return ""
}

func (x *HuaweiCloudBakeryDefaults) GetSecurityGroup() string {
	if x != nil {
		return x.SecurityGroup
	}
	return ""
}

func (x *HuaweiCloudBakeryDefaults) GetEipBandwidthSize() int32 {
	if x != nil {
		return x.EipBandwidthSize
	}
	return 0
}

// Configuration for a base image for the Huawei Cloud provider's bakery.
type HuaweiCloudBaseImageSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Base image configuration.
	BaseImage *HuaweiCloudBaseImage `protobuf:"bytes,1,opt,name=baseImage,proto3" json:"baseImage,omitempty"`
	// Image source configuration.
	VirtualizationSettings []*HuaweiCloudVirtualizationSettings `protobuf:"bytes,2,rep,name=virtualizationSettings,proto3" json:"virtualizationSettings,omitempty"`
}

func (x *HuaweiCloudBaseImageSettings) Reset() {
	*x = HuaweiCloudBaseImageSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_huaweicloud_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HuaweiCloudBaseImageSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HuaweiCloudBaseImageSettings) ProtoMessage() {}

func (x *HuaweiCloudBaseImageSettings) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_huaweicloud_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HuaweiCloudBaseImageSettings.ProtoReflect.Descriptor instead.
func (*HuaweiCloudBaseImageSettings) Descriptor() ([]byte, []int) {
	return file_cloudprovider_huaweicloud_proto_rawDescGZIP(), []int{3}
}

func (x *HuaweiCloudBaseImageSettings) GetBaseImage() *HuaweiCloudBaseImage {
	if x != nil {
		return x.BaseImage
	}
	return nil
}

func (x *HuaweiCloudBaseImageSettings) GetVirtualizationSettings() []*HuaweiCloudVirtualizationSettings {
	if x != nil {
		return x.VirtualizationSettings
	}
	return nil
}

// Huawei Cloud base image settings.
type HuaweiCloudBaseImage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the base image.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// This is used to help Spinnaker's bakery download the build artifacts you
	// supply it with. For example, specifying `deb` indicates that your
	// artifacts will need to be fetched from a debian repository.
	PackageType string `protobuf:"bytes,2,opt,name=packageType,proto3" json:"packageType,omitempty"`
	// This is the name of the packer template that will be used to bake images
	// from this base image. The template file must be found in this list:
	// https://github.com/spinnaker/rosco/tree/master/rosco-web/config/packer, or
	// supplied as described here: https://spinnaker.io/setup/bakery/.
	TemplateFile string `protobuf:"bytes,3,opt,name=templateFile,proto3" json:"templateFile,omitempty"`
	// A short description to help human operators identify the
	// image.
	ShortDescription string `protobuf:"bytes,4,opt,name=shortDescription,proto3" json:"shortDescription,omitempty"`
	// A long description to help human operators identify the
	// image.
	DetailedDescription string `protobuf:"bytes,5,opt,name=detailedDescription,proto3" json:"detailedDescription,omitempty"`
}

func (x *HuaweiCloudBaseImage) Reset() {
	*x = HuaweiCloudBaseImage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_huaweicloud_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HuaweiCloudBaseImage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HuaweiCloudBaseImage) ProtoMessage() {}

func (x *HuaweiCloudBaseImage) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_huaweicloud_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HuaweiCloudBaseImage.ProtoReflect.Descriptor instead.
func (*HuaweiCloudBaseImage) Descriptor() ([]byte, []int) {
	return file_cloudprovider_huaweicloud_proto_rawDescGZIP(), []int{4}
}

func (x *HuaweiCloudBaseImage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *HuaweiCloudBaseImage) GetPackageType() string {
	if x != nil {
		return x.PackageType
	}
	return ""
}

func (x *HuaweiCloudBaseImage) GetTemplateFile() string {
	if x != nil {
		return x.TemplateFile
	}
	return ""
}

func (x *HuaweiCloudBaseImage) GetShortDescription() string {
	if x != nil {
		return x.ShortDescription
	}
	return ""
}

func (x *HuaweiCloudBaseImage) GetDetailedDescription() string {
	if x != nil {
		return x.DetailedDescription
	}
	return ""
}

// Huawei Cloud virtualization settings.
type HuaweiCloudVirtualizationSettings struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// (Required) The region for the baking configuration.
	Region string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	// (Required) The instance type for the baking configuration.
	InstanceType string `protobuf:"bytes,2,opt,name=instanceType,proto3" json:"instanceType,omitempty"`
	// (Required) The source image ID for the baking configuration.
	SourceImageId string `protobuf:"bytes,3,opt,name=sourceImageId,proto3" json:"sourceImageId,omitempty"`
	// (Required) The SSH username for the baking configuration.
	SshUserName string `protobuf:"bytes,4,opt,name=sshUserName,proto3" json:"sshUserName,omitempty"`
	// (Required) The EIP type for the baking configuration. See the API doc to
	// get its value.
	EipType string `protobuf:"bytes,5,opt,name=eipType,proto3" json:"eipType,omitempty"`
}

func (x *HuaweiCloudVirtualizationSettings) Reset() {
	*x = HuaweiCloudVirtualizationSettings{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_huaweicloud_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HuaweiCloudVirtualizationSettings) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HuaweiCloudVirtualizationSettings) ProtoMessage() {}

func (x *HuaweiCloudVirtualizationSettings) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_huaweicloud_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HuaweiCloudVirtualizationSettings.ProtoReflect.Descriptor instead.
func (*HuaweiCloudVirtualizationSettings) Descriptor() ([]byte, []int) {
	return file_cloudprovider_huaweicloud_proto_rawDescGZIP(), []int{5}
}

func (x *HuaweiCloudVirtualizationSettings) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *HuaweiCloudVirtualizationSettings) GetInstanceType() string {
	if x != nil {
		return x.InstanceType
	}
	return ""
}

func (x *HuaweiCloudVirtualizationSettings) GetSourceImageId() string {
	if x != nil {
		return x.SourceImageId
	}
	return ""
}

func (x *HuaweiCloudVirtualizationSettings) GetSshUserName() string {
	if x != nil {
		return x.SshUserName
	}
	return ""
}

func (x *HuaweiCloudVirtualizationSettings) GetEipType() string {
	if x != nil {
		return x.EipType
	}
	return ""
}

var File_cloudprovider_huaweicloud_proto protoreflect.FileDescriptor

var file_cloudprovider_huaweicloud_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f,
	0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x1a, 0x11, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x01, 0x0a, 0x0b, 0x48, 0x75,
	0x61, 0x77, 0x65, 0x69, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x12, 0x43, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x48, 0x75, 0x61, 0x77,
	0x65, 0x69, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x6d,
	0x61, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x56, 0x0a, 0x0e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x79, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c,
	0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x48,
	0x75, 0x61, 0x77, 0x65, 0x69, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x61, 0x6b, 0x65, 0x72, 0x79,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x0e, 0x62, 0x61, 0x6b, 0x65, 0x72, 0x79,
	0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x22, 0xa8, 0x03, 0x0a, 0x12, 0x48, 0x75, 0x61,
	0x77, 0x65, 0x69, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3a, 0x0a, 0x18, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x18, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70,
	0x73, 0x12, 0x34, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x55,
	0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x55, 0x72,
	0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0xc6, 0x03, 0x0a, 0x19, 0x48, 0x75, 0x61, 0x77, 0x65, 0x69, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x42, 0x61, 0x6b, 0x65, 0x72, 0x79, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x73, 0x12, 0x51, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x48, 0x75, 0x61, 0x77,
	0x65, 0x69, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6d,
	0x61, 0x67, 0x65, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x75, 0x74, 0x68,
	0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x55,
	0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x70, 0x63, 0x49,
	0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x70, 0x63, 0x49, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x12, 0x2a, 0x0a, 0x10, 0x65, 0x69, 0x70, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x65, 0x69, 0x70, 0x42,
	0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xd7, 0x01, 0x0a,
	0x1c, 0x48, 0x75, 0x61, 0x77, 0x65, 0x69, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x61, 0x73, 0x65,
	0x49, 0x6d, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x47, 0x0a,
	0x09, 0x62, 0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x48, 0x75, 0x61, 0x77, 0x65, 0x69, 0x43, 0x6c, 0x6f,
	0x75, 0x64, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52, 0x09, 0x62, 0x61, 0x73,
	0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x6e, 0x0a, 0x16, 0x76, 0x69, 0x72, 0x74, 0x75, 0x61,
	0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x48, 0x75, 0x61,
	0x77, 0x65, 0x69, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x69,
	0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x16,
	0x76, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x22, 0xca, 0x01, 0x0a, 0x14, 0x48, 0x75, 0x61, 0x77, 0x65,
	0x69, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x42, 0x61, 0x73, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x20, 0x0a, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x10, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x30, 0x0a, 0x13, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xc1, 0x01, 0x0a, 0x21, 0x48, 0x75, 0x61, 0x77, 0x65, 0x69, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x56, 0x69, 0x72, 0x74, 0x75, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x73,
	0x73, 0x68, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x73, 0x68, 0x55, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x65, 0x69, 0x70, 0x54, 0x79, 0x70, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f,
	0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cloudprovider_huaweicloud_proto_rawDescOnce sync.Once
	file_cloudprovider_huaweicloud_proto_rawDescData = file_cloudprovider_huaweicloud_proto_rawDesc
)

func file_cloudprovider_huaweicloud_proto_rawDescGZIP() []byte {
	file_cloudprovider_huaweicloud_proto_rawDescOnce.Do(func() {
		file_cloudprovider_huaweicloud_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudprovider_huaweicloud_proto_rawDescData)
	})
	return file_cloudprovider_huaweicloud_proto_rawDescData
}

var file_cloudprovider_huaweicloud_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cloudprovider_huaweicloud_proto_goTypes = []interface{}{
	(*HuaweiCloud)(nil),                       // 0: proto.cloudprovider.HuaweiCloud
	(*HuaweiCloudAccount)(nil),                // 1: proto.cloudprovider.HuaweiCloudAccount
	(*HuaweiCloudBakeryDefaults)(nil),         // 2: proto.cloudprovider.HuaweiCloudBakeryDefaults
	(*HuaweiCloudBaseImageSettings)(nil),      // 3: proto.cloudprovider.HuaweiCloudBaseImageSettings
	(*HuaweiCloudBaseImage)(nil),              // 4: proto.cloudprovider.HuaweiCloudBaseImage
	(*HuaweiCloudVirtualizationSettings)(nil), // 5: proto.cloudprovider.HuaweiCloudVirtualizationSettings
	(*client.Permissions)(nil),                // 6: proto.Permissions
}
var file_cloudprovider_huaweicloud_proto_depIdxs = []int32{
	1, // 0: proto.cloudprovider.HuaweiCloud.accounts:type_name -> proto.cloudprovider.HuaweiCloudAccount
	2, // 1: proto.cloudprovider.HuaweiCloud.bakeryDefaults:type_name -> proto.cloudprovider.HuaweiCloudBakeryDefaults
	6, // 2: proto.cloudprovider.HuaweiCloudAccount.permissions:type_name -> proto.Permissions
	3, // 3: proto.cloudprovider.HuaweiCloudBakeryDefaults.baseImages:type_name -> proto.cloudprovider.HuaweiCloudBaseImageSettings
	4, // 4: proto.cloudprovider.HuaweiCloudBaseImageSettings.baseImage:type_name -> proto.cloudprovider.HuaweiCloudBaseImage
	5, // 5: proto.cloudprovider.HuaweiCloudBaseImageSettings.virtualizationSettings:type_name -> proto.cloudprovider.HuaweiCloudVirtualizationSettings
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cloudprovider_huaweicloud_proto_init() }
func file_cloudprovider_huaweicloud_proto_init() {
	if File_cloudprovider_huaweicloud_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloudprovider_huaweicloud_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HuaweiCloud); i {
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
		file_cloudprovider_huaweicloud_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HuaweiCloudAccount); i {
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
		file_cloudprovider_huaweicloud_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HuaweiCloudBakeryDefaults); i {
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
		file_cloudprovider_huaweicloud_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HuaweiCloudBaseImageSettings); i {
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
		file_cloudprovider_huaweicloud_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HuaweiCloudBaseImage); i {
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
		file_cloudprovider_huaweicloud_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HuaweiCloudVirtualizationSettings); i {
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
			RawDescriptor: file_cloudprovider_huaweicloud_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloudprovider_huaweicloud_proto_goTypes,
		DependencyIndexes: file_cloudprovider_huaweicloud_proto_depIdxs,
		MessageInfos:      file_cloudprovider_huaweicloud_proto_msgTypes,
	}.Build()
	File_cloudprovider_huaweicloud_proto = out.File
	file_cloudprovider_huaweicloud_proto_rawDesc = nil
	file_cloudprovider_huaweicloud_proto_goTypes = nil
	file_cloudprovider_huaweicloud_proto_depIdxs = nil
}
