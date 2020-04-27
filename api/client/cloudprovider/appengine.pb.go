// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: cloudprovider/appengine.proto

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

// Configuration for the Google App Engine (GAE) provider.
type AppengineProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured accounts.
	Accounts []*AppengineAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	// The name of the primary account.
	PrimaryAccount string `protobuf:"bytes,3,opt,name=primaryAccount,proto3" json:"primaryAccount,omitempty"`
}

func (x *AppengineProvider) Reset() {
	*x = AppengineProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_appengine_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppengineProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppengineProvider) ProtoMessage() {}

func (x *AppengineProvider) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_appengine_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppengineProvider.ProtoReflect.Descriptor instead.
func (*AppengineProvider) Descriptor() ([]byte, []int) {
	return file_cloudprovider_appengine_proto_rawDescGZIP(), []int{0}
}

func (x *AppengineProvider) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *AppengineProvider) GetAccounts() []*AppengineAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

func (x *AppengineProvider) GetPrimaryAccount() string {
	if x != nil {
		return x.PrimaryAccount
	}
	return ""
}

// Configuration for an App Engine account.
type AppengineAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The interval in seconds at which Spinnaker will poll for updates in
	// your App Engine clusters.
	CachingIntervalSeconds int32 `protobuf:"varint,1,opt,name=cachingIntervalSeconds,proto3" json:"cachingIntervalSeconds,omitempty"`
	// The environment name for the account. Many accounts can share the
	// same environment (e.g., dev, test, prod).
	Environment string `protobuf:"bytes,2,opt,name=environment,proto3" json:"environment,omitempty"`
	// The gcloud release track (`ALPHA`, `BETA`, or `STABLE`) that Spinnaker
	// will use when deploying to App Engine.
	GcloudReleaseTrack string `protobuf:"bytes,3,opt,name=gcloudReleaseTrack,proto3" json:"gcloudReleaseTrack,omitempty"`
	// A username to be used when connecting to a remote git repository
	// server over HTTPS. If set, `gitHttpsPassword` must also be set.
	GitHttpsUsername string `protobuf:"bytes,4,opt,name=gitHttpsUsername,proto3" json:"gitHttpsUsername,omitempty"`
	//  A password to be used when connecting to a remote git repository
	//  server over HTTPS. If set, `gitHttpsUsername` must also be set.
	GitHttpsPassword string `protobuf:"bytes,5,opt,name=gitHttpsPassword,proto3" json:"gitHttpsPassword,omitempty"`
	// An OAuth token provided by Github for connecting to a git repository
	// over HTTPS. See
	// https://help.github.com/articles/creating-an-access-token-for-command-line-use
	// for more information.
	GithubOAuthAccessToken string `protobuf:"bytes,6,opt,name=githubOAuthAccessToken,proto3" json:"githubOAuthAccessToken,omitempty"`
	// The path to a JSON service account that Spinnaker will use as
	// credentials. This is only needed if Spinnaker is not deployed on a
	// Google Compute Engine VM, or needs permissions not afforded to the VM
	// it is running on. See
	// https://cloud.google.com/compute/docs/access/service-accounts for
	// more information.
	JsonPath string `protobuf:"bytes,7,opt,name=jsonPath,proto3" json:"jsonPath,omitempty"`
	// A local directory to be used to stage source files for App Engine
	// deployments within Clouddriver.
	// Defaults to `/var/tmp/clouddriver`.
	LocalRepositoryDirectory string `protobuf:"bytes,8,opt,name=localRepositoryDirectory,proto3" json:"localRepositoryDirectory,omitempty"`
	// A list of regular expressions. Any service matching one of these
	// regexes will be ignored by Spinnaker.
	OmitServices []string `protobuf:"bytes,9,rep,name=omitServices,proto3" json:"omitServices,omitempty"`
	//  A list of regular expressions. Any version matching one of these
	//  regexes will be ignored by Spinnaker.
	OmitVersions []string `protobuf:"bytes,10,rep,name=omitVersions,proto3" json:"omitVersions,omitempty"`
	// The Google Cloud Platform project this Spinnaker account will manage.
	Project string `protobuf:"bytes,11,opt,name=project,proto3" json:"project,omitempty"`
	// Fiat permissions configuration.
	Permissions *client.Permissions `protobuf:"bytes,12,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// (Deprecated): List of required Fiat permission groups. Configure
	// `permissions` instead.
	RequiredGroupMemberships []string `protobuf:"bytes,13,rep,name=requiredGroupMemberships,proto3" json:"requiredGroupMemberships,omitempty"`
	// A list of regular expressions. Any service matching one of these
	// regexes will be indexed by Spinnaker (unless the service also
	// matches a regex in `omitServices`).
	Services []string `protobuf:"bytes,14,rep,name=services,proto3" json:"services,omitempty"`
	// The path to a `known_hosts` file to be used when connecting with a
	// remote git repository over SSH.
	SshKnownHostsFilePath string `protobuf:"bytes,15,opt,name=sshKnownHostsFilePath,proto3" json:"sshKnownHostsFilePath,omitempty"`
	// The path to an SSH private key to be used when connecting with a
	// remote git repository over SSH. If set, `sshPrivateKeyPassphrase` must
	// also be set.
	SshPrivateKeyFilePath string `protobuf:"bytes,16,opt,name=sshPrivateKeyFilePath,proto3" json:"sshPrivateKeyFilePath,omitempty"`
	// The passphrase to an SSH private key to be used when connecting with
	// a remote git repository over SSH. If set, `sshPrivateKeyFilePath` must
	// also be set.
	SshPrivateKeyPassphrase string `protobuf:"bytes,17,opt,name=sshPrivateKeyPassphrase,proto3" json:"sshPrivateKeyPassphrase,omitempty"`
	// Enabling this flag will allow Spinnaker to connect with a remote git
	// repository over SSH without verifying the server's IP address against
	// a `known_hosts` file. Defaults to false.
	SshTrustUnknownHosts bool `protobuf:"varint,18,opt,name=sshTrustUnknownHosts,proto3" json:"sshTrustUnknownHosts,omitempty"`
	// A list of regular expressions. Any version matching one of these
	// regexes will be indexed by Spinnaker (unless the version also matches
	// a regex in `omitVersions`).
	Versions []string `protobuf:"bytes,19,rep,name=versions,proto3" json:"versions,omitempty"`
	// The name of the account.
	Name string `protobuf:"bytes,20,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AppengineAccount) Reset() {
	*x = AppengineAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_appengine_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppengineAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppengineAccount) ProtoMessage() {}

func (x *AppengineAccount) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_appengine_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppengineAccount.ProtoReflect.Descriptor instead.
func (*AppengineAccount) Descriptor() ([]byte, []int) {
	return file_cloudprovider_appengine_proto_rawDescGZIP(), []int{1}
}

func (x *AppengineAccount) GetCachingIntervalSeconds() int32 {
	if x != nil {
		return x.CachingIntervalSeconds
	}
	return 0
}

func (x *AppengineAccount) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *AppengineAccount) GetGcloudReleaseTrack() string {
	if x != nil {
		return x.GcloudReleaseTrack
	}
	return ""
}

func (x *AppengineAccount) GetGitHttpsUsername() string {
	if x != nil {
		return x.GitHttpsUsername
	}
	return ""
}

func (x *AppengineAccount) GetGitHttpsPassword() string {
	if x != nil {
		return x.GitHttpsPassword
	}
	return ""
}

func (x *AppengineAccount) GetGithubOAuthAccessToken() string {
	if x != nil {
		return x.GithubOAuthAccessToken
	}
	return ""
}

func (x *AppengineAccount) GetJsonPath() string {
	if x != nil {
		return x.JsonPath
	}
	return ""
}

func (x *AppengineAccount) GetLocalRepositoryDirectory() string {
	if x != nil {
		return x.LocalRepositoryDirectory
	}
	return ""
}

func (x *AppengineAccount) GetOmitServices() []string {
	if x != nil {
		return x.OmitServices
	}
	return nil
}

func (x *AppengineAccount) GetOmitVersions() []string {
	if x != nil {
		return x.OmitVersions
	}
	return nil
}

func (x *AppengineAccount) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *AppengineAccount) GetPermissions() *client.Permissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *AppengineAccount) GetRequiredGroupMemberships() []string {
	if x != nil {
		return x.RequiredGroupMemberships
	}
	return nil
}

func (x *AppengineAccount) GetServices() []string {
	if x != nil {
		return x.Services
	}
	return nil
}

func (x *AppengineAccount) GetSshKnownHostsFilePath() string {
	if x != nil {
		return x.SshKnownHostsFilePath
	}
	return ""
}

func (x *AppengineAccount) GetSshPrivateKeyFilePath() string {
	if x != nil {
		return x.SshPrivateKeyFilePath
	}
	return ""
}

func (x *AppengineAccount) GetSshPrivateKeyPassphrase() string {
	if x != nil {
		return x.SshPrivateKeyPassphrase
	}
	return ""
}

func (x *AppengineAccount) GetSshTrustUnknownHosts() bool {
	if x != nil {
		return x.SshTrustUnknownHosts
	}
	return false
}

func (x *AppengineAccount) GetVersions() []string {
	if x != nil {
		return x.Versions
	}
	return nil
}

func (x *AppengineAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_cloudprovider_appengine_proto protoreflect.FileDescriptor

var file_cloudprovider_appengine_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x1a, 0x11, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x11, 0x41, 0x70, 0x70, 0x65,
	0x6e, 0x67, 0x69, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x41, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x41, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0xfe, 0x06, 0x0a, 0x10, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x36, 0x0a, 0x16, 0x63, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x67, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x16, 0x63, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x67,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x2e, 0x0a, 0x12, 0x67, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x52, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x67,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x54, 0x72, 0x61, 0x63,
	0x6b, 0x12, 0x2a, 0x0a, 0x10, 0x67, 0x69, 0x74, 0x48, 0x74, 0x74, 0x70, 0x73, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x67, 0x69, 0x74,
	0x48, 0x74, 0x74, 0x70, 0x73, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a,
	0x10, 0x67, 0x69, 0x74, 0x48, 0x74, 0x74, 0x70, 0x73, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x67, 0x69, 0x74, 0x48, 0x74, 0x74, 0x70,
	0x73, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x36, 0x0a, 0x16, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x12, 0x3a, 0x0a,
	0x18, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x18, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x6d, 0x69,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0c, 0x6f, 0x6d, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x22, 0x0a,
	0x0c, 0x6f, 0x6d, 0x69, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x6d, 0x69, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x34, 0x0a, 0x0b, 0x70,
	0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x3a, 0x0a, 0x18, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18, 0x0d, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x18, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x1a, 0x0a,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x34, 0x0a, 0x15, 0x73, 0x73, 0x68,
	0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x73, 0x73, 0x68, 0x4b, 0x6e, 0x6f,
	0x77, 0x6e, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x12,
	0x34, 0x0a, 0x15, 0x73, 0x73, 0x68, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15,
	0x73, 0x73, 0x68, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x46, 0x69, 0x6c,
	0x65, 0x50, 0x61, 0x74, 0x68, 0x12, 0x38, 0x0a, 0x17, 0x73, 0x73, 0x68, 0x50, 0x72, 0x69, 0x76,
	0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x73, 0x73, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65,
	0x18, 0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x17, 0x73, 0x73, 0x68, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x4b, 0x65, 0x79, 0x50, 0x61, 0x73, 0x73, 0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x14, 0x73, 0x73, 0x68, 0x54, 0x72, 0x75, 0x73, 0x74, 0x55, 0x6e, 0x6b, 0x6e, 0x6f,
	0x77, 0x6e, 0x48, 0x6f, 0x73, 0x74, 0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14, 0x73,
	0x73, 0x68, 0x54, 0x72, 0x75, 0x73, 0x74, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x48, 0x6f,
	0x73, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x13, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61,
	0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cloudprovider_appengine_proto_rawDescOnce sync.Once
	file_cloudprovider_appengine_proto_rawDescData = file_cloudprovider_appengine_proto_rawDesc
)

func file_cloudprovider_appengine_proto_rawDescGZIP() []byte {
	file_cloudprovider_appengine_proto_rawDescOnce.Do(func() {
		file_cloudprovider_appengine_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudprovider_appengine_proto_rawDescData)
	})
	return file_cloudprovider_appengine_proto_rawDescData
}

var file_cloudprovider_appengine_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cloudprovider_appengine_proto_goTypes = []interface{}{
	(*AppengineProvider)(nil),  // 0: proto.cloudprovider.AppengineProvider
	(*AppengineAccount)(nil),   // 1: proto.cloudprovider.AppengineAccount
	(*client.Permissions)(nil), // 2: proto.Permissions
}
var file_cloudprovider_appengine_proto_depIdxs = []int32{
	1, // 0: proto.cloudprovider.AppengineProvider.accounts:type_name -> proto.cloudprovider.AppengineAccount
	2, // 1: proto.cloudprovider.AppengineAccount.permissions:type_name -> proto.Permissions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cloudprovider_appengine_proto_init() }
func file_cloudprovider_appengine_proto_init() {
	if File_cloudprovider_appengine_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloudprovider_appengine_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppengineProvider); i {
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
		file_cloudprovider_appengine_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppengineAccount); i {
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
			RawDescriptor: file_cloudprovider_appengine_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloudprovider_appengine_proto_goTypes,
		DependencyIndexes: file_cloudprovider_appengine_proto_depIdxs,
		MessageInfos:      file_cloudprovider_appengine_proto_msgTypes,
	}.Build()
	File_cloudprovider_appengine_proto = out.File
	file_cloudprovider_appengine_proto_rawDesc = nil
	file_cloudprovider_appengine_proto_goTypes = nil
	file_cloudprovider_appengine_proto_depIdxs = nil
}
