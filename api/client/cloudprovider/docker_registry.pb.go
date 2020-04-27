// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: cloudprovider/docker_registry.proto

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

// Configuration for the Docker Registry provider.
type DockerRegistryProvider struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether the provider is enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The list of configured accounts.
	Accounts []*DockerRegistryAccount `protobuf:"bytes,2,rep,name=accounts,proto3" json:"accounts,omitempty"`
	// The name of the primary account.
	PrimaryAccount string `protobuf:"bytes,3,opt,name=primaryAccount,proto3" json:"primaryAccount,omitempty"`
}

func (x *DockerRegistryProvider) Reset() {
	*x = DockerRegistryProvider{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_docker_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DockerRegistryProvider) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DockerRegistryProvider) ProtoMessage() {}

func (x *DockerRegistryProvider) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_docker_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DockerRegistryProvider.ProtoReflect.Descriptor instead.
func (*DockerRegistryProvider) Descriptor() ([]byte, []int) {
	return file_cloudprovider_docker_registry_proto_rawDescGZIP(), []int{0}
}

func (x *DockerRegistryProvider) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *DockerRegistryProvider) GetAccounts() []*DockerRegistryAccount {
	if x != nil {
		return x.Accounts
	}
	return nil
}

func (x *DockerRegistryProvider) GetPrimaryAccount() string {
	if x != nil {
		return x.PrimaryAccount
	}
	return ""
}

// A credential able to authenticate against a set of Docker repositories.
type DockerRegistryAccount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The name of the account.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// (Required) The registry address from which to pull and deploy images
	// (e.g., `https://index.docker.io`).
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// The number of seconds between polling the Docker registry. Certain
	// registries are sensitive to over-polling, and larger intervals
	// (e.g., 10 minutes = 600 seconds) are desirable if you experience rate
	// limiting. Defaults to `30`.
	CacheIntervalSeconds int32 `protobuf:"varint,3,opt,name=cacheIntervalSeconds,proto3" json:"cacheIntervalSeconds,omitempty"`
	// The number of threads on which to cache all provided repositories.
	// Really only useful if you have a ton of repos. Defaults to 1.
	CacheThreads int32 `protobuf:"varint,4,opt,name=cacheThreads,proto3" json:"cacheThreads,omitempty"`
	// Timeout in milliseconds for provided repositories. Defaults to
	// `60,000`.
	ClientTimeoutMillis int32 `protobuf:"varint,5,opt,name=clientTimeoutMillis,proto3" json:"clientTimeoutMillis,omitempty"`
	// The email associated with your Docker registry. Often this only needs to
	// be well-formed, rather than be a real address.
	Email string `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	// The environment name for the account. Many accounts can share the
	// same environment (e.g., dev, test, prod).
	Environment string `protobuf:"bytes,7,opt,name=environment,proto3" json:"environment,omitempty"`
	// If `true`, Spinnaker will treat the Docker registry as insecure and not
	// validate the SSL certificate. Defaults to `false`.
	InsecureRegistry bool `protobuf:"varint,8,opt,name=insecureRegistry,proto3" json:"insecureRegistry,omitempty"`
	// Pagination size for the Docker `repository _catalog` endpoint. Defaults
	// to `100`.
	PaginateSize int32 `protobuf:"varint,9,opt,name=paginateSize,proto3" json:"paginateSize,omitempty"`
	// The Docker registry password. Only one of `password`, `passwordCommand`,
	// and `passwordFile` should be specified.
	Password string `protobuf:"bytes,10,opt,name=password,proto3" json:"password,omitempty"`
	// Command to retrieve Docker token/password. The command must be available
	// in the environment. Only one of `password`, `passwordCommand`, and
	// `passwordFile` should be specified.
	PasswordCommand string `protobuf:"bytes,11,opt,name=passwordCommand,proto3" json:"passwordCommand,omitempty"`
	// The path to a file containing your Docker password in plaintext (not a
	// Docker `config.json` file). Only one of `password`, `passwordCommand`,
	// and `passwordFile` should be specified.
	PasswordFile string `protobuf:"bytes,12,opt,name=passwordFile,proto3" json:"passwordFile,omitempty"`
	// Fiat permissions configuration.
	Permissions *client.Permissions `protobuf:"bytes,13,opt,name=permissions,proto3" json:"permissions,omitempty"`
	// (Deprecated) List of required Fiat permission groups. Configure
	// `permissions` instead.
	RequiredGroupMemberships []string `protobuf:"bytes,14,rep,name=requiredGroupMemberships,proto3" json:"requiredGroupMemberships,omitempty"`
	// An optional list of repositories from which to cache images. If not
	// provided, Spinnaker will attempt to read accessible repositories from the
	// `registries _catalog` endpoint.
	Repositories []string `protobuf:"bytes,15,rep,name=repositories,proto3" json:"repositories,omitempty"`
	// If `true`, Spinnaker will sort tags by creation date. Defaults to
	// `false`. Not recommended for use with large registries; sorting
	// performance scales poorly due to limitations of the Docker V2 API.
	SortTagsByDate bool `protobuf:"varint,16,opt,name=sortTagsByDate,proto3" json:"sortTagsByDate,omitempty"`
	// If `true`, Spinnaker will track digest changes. This is not recommended
	// because it greatly increases queries to the registry, and most
	// registries are flaky. Defaults to `false`.
	TrackDigests bool `protobuf:"varint,17,opt,name=trackDigests,proto3" json:"trackDigests,omitempty"`
	// The username associated with this Docker registry.
	Username string `protobuf:"bytes,18,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *DockerRegistryAccount) Reset() {
	*x = DockerRegistryAccount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cloudprovider_docker_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DockerRegistryAccount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DockerRegistryAccount) ProtoMessage() {}

func (x *DockerRegistryAccount) ProtoReflect() protoreflect.Message {
	mi := &file_cloudprovider_docker_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DockerRegistryAccount.ProtoReflect.Descriptor instead.
func (*DockerRegistryAccount) Descriptor() ([]byte, []int) {
	return file_cloudprovider_docker_registry_proto_rawDescGZIP(), []int{1}
}

func (x *DockerRegistryAccount) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DockerRegistryAccount) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *DockerRegistryAccount) GetCacheIntervalSeconds() int32 {
	if x != nil {
		return x.CacheIntervalSeconds
	}
	return 0
}

func (x *DockerRegistryAccount) GetCacheThreads() int32 {
	if x != nil {
		return x.CacheThreads
	}
	return 0
}

func (x *DockerRegistryAccount) GetClientTimeoutMillis() int32 {
	if x != nil {
		return x.ClientTimeoutMillis
	}
	return 0
}

func (x *DockerRegistryAccount) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *DockerRegistryAccount) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *DockerRegistryAccount) GetInsecureRegistry() bool {
	if x != nil {
		return x.InsecureRegistry
	}
	return false
}

func (x *DockerRegistryAccount) GetPaginateSize() int32 {
	if x != nil {
		return x.PaginateSize
	}
	return 0
}

func (x *DockerRegistryAccount) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *DockerRegistryAccount) GetPasswordCommand() string {
	if x != nil {
		return x.PasswordCommand
	}
	return ""
}

func (x *DockerRegistryAccount) GetPasswordFile() string {
	if x != nil {
		return x.PasswordFile
	}
	return ""
}

func (x *DockerRegistryAccount) GetPermissions() *client.Permissions {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *DockerRegistryAccount) GetRequiredGroupMemberships() []string {
	if x != nil {
		return x.RequiredGroupMemberships
	}
	return nil
}

func (x *DockerRegistryAccount) GetRepositories() []string {
	if x != nil {
		return x.Repositories
	}
	return nil
}

func (x *DockerRegistryAccount) GetSortTagsByDate() bool {
	if x != nil {
		return x.SortTagsByDate
	}
	return false
}

func (x *DockerRegistryAccount) GetTrackDigests() bool {
	if x != nil {
		return x.TrackDigests
	}
	return false
}

func (x *DockerRegistryAccount) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

var File_cloudprovider_docker_registry_proto protoreflect.FileDescriptor

var file_cloudprovider_docker_registry_proto_rawDesc = []byte{
	0x0a, 0x23, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f,
	0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x1a, 0x11, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01,
	0x0a, 0x16, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x12, 0x46, 0x0a, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x6f, 0x63, 0x6b, 0x65,
	0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72,
	0x69, 0x6d, 0x61, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x22, 0xbf, 0x05, 0x0a, 0x15, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x79, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x32, 0x0a, 0x14, 0x63, 0x61,
	0x63, 0x68, 0x65, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e,
	0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x14, 0x63, 0x61, 0x63, 0x68, 0x65, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x22,
	0x0a, 0x0c, 0x63, 0x61, 0x63, 0x68, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x63, 0x61, 0x63, 0x68, 0x65, 0x54, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x73, 0x12, 0x30, 0x0a, 0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x4d, 0x69, 0x6c, 0x6c, 0x69, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4d, 0x69,
	0x6c, 0x6c, 0x69, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e,
	0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x10,
	0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x65,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x34, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52,
	0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3a, 0x0a, 0x18,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x09, 0x52, 0x18,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x0e,
	0x73, 0x6f, 0x72, 0x74, 0x54, 0x61, 0x67, 0x73, 0x42, 0x79, 0x44, 0x61, 0x74, 0x65, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x73, 0x6f, 0x72, 0x74, 0x54, 0x61, 0x67, 0x73, 0x42, 0x79,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x44, 0x69, 0x67,
	0x65, 0x73, 0x74, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65,
	0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cloudprovider_docker_registry_proto_rawDescOnce sync.Once
	file_cloudprovider_docker_registry_proto_rawDescData = file_cloudprovider_docker_registry_proto_rawDesc
)

func file_cloudprovider_docker_registry_proto_rawDescGZIP() []byte {
	file_cloudprovider_docker_registry_proto_rawDescOnce.Do(func() {
		file_cloudprovider_docker_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_cloudprovider_docker_registry_proto_rawDescData)
	})
	return file_cloudprovider_docker_registry_proto_rawDescData
}

var file_cloudprovider_docker_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cloudprovider_docker_registry_proto_goTypes = []interface{}{
	(*DockerRegistryProvider)(nil), // 0: proto.cloudprovider.DockerRegistryProvider
	(*DockerRegistryAccount)(nil),  // 1: proto.cloudprovider.DockerRegistryAccount
	(*client.Permissions)(nil),     // 2: proto.Permissions
}
var file_cloudprovider_docker_registry_proto_depIdxs = []int32{
	1, // 0: proto.cloudprovider.DockerRegistryProvider.accounts:type_name -> proto.cloudprovider.DockerRegistryAccount
	2, // 1: proto.cloudprovider.DockerRegistryAccount.permissions:type_name -> proto.Permissions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cloudprovider_docker_registry_proto_init() }
func file_cloudprovider_docker_registry_proto_init() {
	if File_cloudprovider_docker_registry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cloudprovider_docker_registry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DockerRegistryProvider); i {
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
		file_cloudprovider_docker_registry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DockerRegistryAccount); i {
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
			RawDescriptor: file_cloudprovider_docker_registry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cloudprovider_docker_registry_proto_goTypes,
		DependencyIndexes: file_cloudprovider_docker_registry_proto_depIdxs,
		MessageInfos:      file_cloudprovider_docker_registry_proto_msgTypes,
	}.Build()
	File_cloudprovider_docker_registry_proto = out.File
	file_cloudprovider_docker_registry_proto_rawDesc = nil
	file_cloudprovider_docker_registry_proto_goTypes = nil
	file_cloudprovider_docker_registry_proto_depIdxs = nil
}
