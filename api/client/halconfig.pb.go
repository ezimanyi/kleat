// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: halconfig.proto

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

type HalConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersistentStorage *HalConfig_PersistentStorage `protobuf:"bytes,1,opt,name=persistentStorage,proto3" json:"persistentStorage,omitempty"`
	Providers         *HalConfig_Providers         `protobuf:"bytes,2,opt,name=providers,proto3" json:"providers,omitempty"`
	Artifacts         *ArtifactProviders           `protobuf:"bytes,3,opt,name=artifacts,proto3" json:"artifacts,omitempty"`
	Notifications     *HalConfig_Notifications     `protobuf:"bytes,4,opt,name=notifications,proto3" json:"notifications,omitempty"`
}

func (x *HalConfig) Reset() {
	*x = HalConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_halconfig_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HalConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HalConfig) ProtoMessage() {}

func (x *HalConfig) ProtoReflect() protoreflect.Message {
	mi := &file_halconfig_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HalConfig.ProtoReflect.Descriptor instead.
func (*HalConfig) Descriptor() ([]byte, []int) {
	return file_halconfig_proto_rawDescGZIP(), []int{0}
}

func (x *HalConfig) GetPersistentStorage() *HalConfig_PersistentStorage {
	if x != nil {
		return x.PersistentStorage
	}
	return nil
}

func (x *HalConfig) GetProviders() *HalConfig_Providers {
	if x != nil {
		return x.Providers
	}
	return nil
}

func (x *HalConfig) GetArtifacts() *ArtifactProviders {
	if x != nil {
		return x.Artifacts
	}
	return nil
}

func (x *HalConfig) GetNotifications() *HalConfig_Notifications {
	if x != nil {
		return x.Notifications
	}
	return nil
}

type HalConfig_PersistentStorage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersistentStoreType string `protobuf:"bytes,1,opt,name=persistentStoreType,proto3" json:"persistentStoreType,omitempty"`
	Gcs                 *GCS   `protobuf:"bytes,2,opt,name=gcs,proto3" json:"gcs,omitempty"`
}

func (x *HalConfig_PersistentStorage) Reset() {
	*x = HalConfig_PersistentStorage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_halconfig_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HalConfig_PersistentStorage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HalConfig_PersistentStorage) ProtoMessage() {}

func (x *HalConfig_PersistentStorage) ProtoReflect() protoreflect.Message {
	mi := &file_halconfig_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HalConfig_PersistentStorage.ProtoReflect.Descriptor instead.
func (*HalConfig_PersistentStorage) Descriptor() ([]byte, []int) {
	return file_halconfig_proto_rawDescGZIP(), []int{0, 0}
}

func (x *HalConfig_PersistentStorage) GetPersistentStoreType() string {
	if x != nil {
		return x.PersistentStoreType
	}
	return ""
}

func (x *HalConfig_PersistentStorage) GetGcs() *GCS {
	if x != nil {
		return x.Gcs
	}
	return nil
}

type HalConfig_Providers struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Kubernetes     *KubernetesProvider     `protobuf:"bytes,1,opt,name=kubernetes,proto3" json:"kubernetes,omitempty"`
	Google         *GoogleProvider         `protobuf:"bytes,2,opt,name=google,proto3" json:"google,omitempty"`
	Appengine      *AppengineProvider      `protobuf:"bytes,3,opt,name=appengine,proto3" json:"appengine,omitempty"`
	Aws            *AwsProvider            `protobuf:"bytes,4,opt,name=aws,proto3" json:"aws,omitempty"`
	Azure          *AzureProvider          `protobuf:"bytes,5,opt,name=azure,proto3" json:"azure,omitempty"`
	Cloudfoundry   *CloudFoundryProvider   `protobuf:"bytes,6,opt,name=cloudfoundry,proto3" json:"cloudfoundry,omitempty"`
	Dcos           *DcosProvider           `protobuf:"bytes,7,opt,name=dcos,proto3" json:"dcos,omitempty"`
	DockerRegistry *DockerRegistryProvider `protobuf:"bytes,8,opt,name=dockerRegistry,proto3" json:"dockerRegistry,omitempty"`
	Ecs            *EcsProvider            `protobuf:"bytes,9,opt,name=ecs,proto3" json:"ecs,omitempty"`
	Huaweicloud    *HuaweiCloudProvider    `protobuf:"bytes,10,opt,name=huaweicloud,proto3" json:"huaweicloud,omitempty"`
	Oracle         *OracleProvider         `protobuf:"bytes,11,opt,name=oracle,proto3" json:"oracle,omitempty"`
}

func (x *HalConfig_Providers) Reset() {
	*x = HalConfig_Providers{}
	if protoimpl.UnsafeEnabled {
		mi := &file_halconfig_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HalConfig_Providers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HalConfig_Providers) ProtoMessage() {}

func (x *HalConfig_Providers) ProtoReflect() protoreflect.Message {
	mi := &file_halconfig_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HalConfig_Providers.ProtoReflect.Descriptor instead.
func (*HalConfig_Providers) Descriptor() ([]byte, []int) {
	return file_halconfig_proto_rawDescGZIP(), []int{0, 1}
}

func (x *HalConfig_Providers) GetKubernetes() *KubernetesProvider {
	if x != nil {
		return x.Kubernetes
	}
	return nil
}

func (x *HalConfig_Providers) GetGoogle() *GoogleProvider {
	if x != nil {
		return x.Google
	}
	return nil
}

func (x *HalConfig_Providers) GetAppengine() *AppengineProvider {
	if x != nil {
		return x.Appengine
	}
	return nil
}

func (x *HalConfig_Providers) GetAws() *AwsProvider {
	if x != nil {
		return x.Aws
	}
	return nil
}

func (x *HalConfig_Providers) GetAzure() *AzureProvider {
	if x != nil {
		return x.Azure
	}
	return nil
}

func (x *HalConfig_Providers) GetCloudfoundry() *CloudFoundryProvider {
	if x != nil {
		return x.Cloudfoundry
	}
	return nil
}

func (x *HalConfig_Providers) GetDcos() *DcosProvider {
	if x != nil {
		return x.Dcos
	}
	return nil
}

func (x *HalConfig_Providers) GetDockerRegistry() *DockerRegistryProvider {
	if x != nil {
		return x.DockerRegistry
	}
	return nil
}

func (x *HalConfig_Providers) GetEcs() *EcsProvider {
	if x != nil {
		return x.Ecs
	}
	return nil
}

func (x *HalConfig_Providers) GetHuaweicloud() *HuaweiCloudProvider {
	if x != nil {
		return x.Huaweicloud
	}
	return nil
}

func (x *HalConfig_Providers) GetOracle() *OracleProvider {
	if x != nil {
		return x.Oracle
	}
	return nil
}

type HalConfig_Notifications struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Slack        *SlackNotification        `protobuf:"bytes,1,opt,name=slack,proto3" json:"slack,omitempty"`
	Twilio       *TwilioNotification       `protobuf:"bytes,2,opt,name=twilio,proto3" json:"twilio,omitempty"`
	GithubStatus *GithubStatusNotification `protobuf:"bytes,3,opt,name=githubStatus,proto3" json:"githubStatus,omitempty"`
}

func (x *HalConfig_Notifications) Reset() {
	*x = HalConfig_Notifications{}
	if protoimpl.UnsafeEnabled {
		mi := &file_halconfig_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HalConfig_Notifications) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HalConfig_Notifications) ProtoMessage() {}

func (x *HalConfig_Notifications) ProtoReflect() protoreflect.Message {
	mi := &file_halconfig_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HalConfig_Notifications.ProtoReflect.Descriptor instead.
func (*HalConfig_Notifications) Descriptor() ([]byte, []int) {
	return file_halconfig_proto_rawDescGZIP(), []int{0, 2}
}

func (x *HalConfig_Notifications) GetSlack() *SlackNotification {
	if x != nil {
		return x.Slack
	}
	return nil
}

func (x *HalConfig_Notifications) GetTwilio() *TwilioNotification {
	if x != nil {
		return x.Twilio
	}
	return nil
}

func (x *HalConfig_Notifications) GetGithubStatus() *GithubStatusNotification {
	if x != nil {
		return x.GithubStatus
	}
	return nil
}

var File_halconfig_proto protoreflect.FileDescriptor

var file_halconfig_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x68, 0x61, 0x6c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x5f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12,
	0x61, 0x77, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x14, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x64, 0x63, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x64, 0x6f, 0x63, 0x6b,
	0x65, 0x72, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x65, 0x63, 0x73, 0x5f,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a,
	0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x6f, 0x72, 0x61, 0x63,
	0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x13, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x08, 0x0a, 0x09, 0x48, 0x61, 0x6c, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x50, 0x0a, 0x11, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x52, 0x11, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x48, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x36, 0x0a, 0x09, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x52, 0x09,
	0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x73, 0x12, 0x44, 0x0a, 0x0d, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x61, 0x6c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x0d, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a,
	0x63, 0x0a, 0x11, 0x50, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x12, 0x30, 0x0a, 0x13, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x13, 0x70, 0x65, 0x72, 0x73, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x03, 0x67, 0x63, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x43, 0x53, 0x52,
	0x03, 0x67, 0x63, 0x73, 0x1a, 0xc3, 0x04, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x12, 0x2d, 0x0a,
	0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x12, 0x36, 0x0a, 0x09,
	0x61, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x70, 0x70, 0x65, 0x6e, 0x67, 0x69, 0x6e,
	0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x09, 0x61, 0x70, 0x70, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x12, 0x24, 0x0a, 0x03, 0x61, 0x77, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x41, 0x77, 0x73, 0x50, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x03, 0x61, 0x77, 0x73, 0x12, 0x2a, 0x0a, 0x05, 0x61, 0x7a,
	0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x7a, 0x75, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52,
	0x05, 0x61, 0x7a, 0x75, 0x72, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x72,
	0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x0c, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x72, 0x79, 0x12, 0x27, 0x0a, 0x04, 0x64, 0x63, 0x6f, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x63,
	0x6f, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x04, 0x64, 0x63, 0x6f, 0x73,
	0x12, 0x45, 0x0a, 0x0e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x44, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x0e, 0x64, 0x6f, 0x63, 0x6b, 0x65, 0x72, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x03, 0x65, 0x63, 0x73, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x63, 0x73,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x03, 0x65, 0x63, 0x73, 0x12, 0x3c, 0x0a,
	0x0b, 0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x75, 0x61, 0x77, 0x65,
	0x69, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x0b,
	0x68, 0x75, 0x61, 0x77, 0x65, 0x69, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x12, 0x2d, 0x0a, 0x06, 0x6f,
	0x72, 0x61, 0x63, 0x6c, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x52, 0x06, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x1a, 0xb7, 0x01, 0x0a, 0x0d, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2e, 0x0a, 0x05,
	0x73, 0x6c, 0x61, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x12, 0x31, 0x0a, 0x06,
	0x74, 0x77, 0x69, 0x6c, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x77, 0x69, 0x6c, 0x69, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x74, 0x77, 0x69, 0x6c, 0x69, 0x6f, 0x12,
	0x43, 0x0a, 0x0c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65,
	0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_halconfig_proto_rawDescOnce sync.Once
	file_halconfig_proto_rawDescData = file_halconfig_proto_rawDesc
)

func file_halconfig_proto_rawDescGZIP() []byte {
	file_halconfig_proto_rawDescOnce.Do(func() {
		file_halconfig_proto_rawDescData = protoimpl.X.CompressGZIP(file_halconfig_proto_rawDescData)
	})
	return file_halconfig_proto_rawDescData
}

var file_halconfig_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_halconfig_proto_goTypes = []interface{}{
	(*HalConfig)(nil),                   // 0: proto.HalConfig
	(*HalConfig_PersistentStorage)(nil), // 1: proto.HalConfig.PersistentStorage
	(*HalConfig_Providers)(nil),         // 2: proto.HalConfig.Providers
	(*HalConfig_Notifications)(nil),     // 3: proto.HalConfig.Notifications
	(*ArtifactProviders)(nil),           // 4: proto.ArtifactProviders
	(*GCS)(nil),                         // 5: proto.GCS
	(*KubernetesProvider)(nil),          // 6: proto.KubernetesProvider
	(*GoogleProvider)(nil),              // 7: proto.GoogleProvider
	(*AppengineProvider)(nil),           // 8: proto.AppengineProvider
	(*AwsProvider)(nil),                 // 9: proto.AwsProvider
	(*AzureProvider)(nil),               // 10: proto.AzureProvider
	(*CloudFoundryProvider)(nil),        // 11: proto.CloudFoundryProvider
	(*DcosProvider)(nil),                // 12: proto.DcosProvider
	(*DockerRegistryProvider)(nil),      // 13: proto.DockerRegistryProvider
	(*EcsProvider)(nil),                 // 14: proto.EcsProvider
	(*HuaweiCloudProvider)(nil),         // 15: proto.HuaweiCloudProvider
	(*OracleProvider)(nil),              // 16: proto.OracleProvider
	(*SlackNotification)(nil),           // 17: proto.SlackNotification
	(*TwilioNotification)(nil),          // 18: proto.TwilioNotification
	(*GithubStatusNotification)(nil),    // 19: proto.GithubStatusNotification
}
var file_halconfig_proto_depIdxs = []int32{
	1,  // 0: proto.HalConfig.persistentStorage:type_name -> proto.HalConfig.PersistentStorage
	2,  // 1: proto.HalConfig.providers:type_name -> proto.HalConfig.Providers
	4,  // 2: proto.HalConfig.artifacts:type_name -> proto.ArtifactProviders
	3,  // 3: proto.HalConfig.notifications:type_name -> proto.HalConfig.Notifications
	5,  // 4: proto.HalConfig.PersistentStorage.gcs:type_name -> proto.GCS
	6,  // 5: proto.HalConfig.Providers.kubernetes:type_name -> proto.KubernetesProvider
	7,  // 6: proto.HalConfig.Providers.google:type_name -> proto.GoogleProvider
	8,  // 7: proto.HalConfig.Providers.appengine:type_name -> proto.AppengineProvider
	9,  // 8: proto.HalConfig.Providers.aws:type_name -> proto.AwsProvider
	10, // 9: proto.HalConfig.Providers.azure:type_name -> proto.AzureProvider
	11, // 10: proto.HalConfig.Providers.cloudfoundry:type_name -> proto.CloudFoundryProvider
	12, // 11: proto.HalConfig.Providers.dcos:type_name -> proto.DcosProvider
	13, // 12: proto.HalConfig.Providers.dockerRegistry:type_name -> proto.DockerRegistryProvider
	14, // 13: proto.HalConfig.Providers.ecs:type_name -> proto.EcsProvider
	15, // 14: proto.HalConfig.Providers.huaweicloud:type_name -> proto.HuaweiCloudProvider
	16, // 15: proto.HalConfig.Providers.oracle:type_name -> proto.OracleProvider
	17, // 16: proto.HalConfig.Notifications.slack:type_name -> proto.SlackNotification
	18, // 17: proto.HalConfig.Notifications.twilio:type_name -> proto.TwilioNotification
	19, // 18: proto.HalConfig.Notifications.githubStatus:type_name -> proto.GithubStatusNotification
	19, // [19:19] is the sub-list for method output_type
	19, // [19:19] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_halconfig_proto_init() }
func file_halconfig_proto_init() {
	if File_halconfig_proto != nil {
		return
	}
	file_persistent_storage_proto_init()
	file_kubernetes_provider_proto_init()
	file_google_provider_proto_init()
	file_appengine_provider_proto_init()
	file_aws_provider_proto_init()
	file_azure_provider_proto_init()
	file_cloudfoundry_provider_proto_init()
	file_dcos_provider_proto_init()
	file_docker_registry_provider_proto_init()
	file_ecs_provider_proto_init()
	file_huaweicloud_provider_proto_init()
	file_oracle_provider_proto_init()
	file_artifacts_proto_init()
	file_notifications_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_halconfig_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HalConfig); i {
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
		file_halconfig_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HalConfig_PersistentStorage); i {
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
		file_halconfig_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HalConfig_Providers); i {
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
		file_halconfig_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HalConfig_Notifications); i {
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
			RawDescriptor: file_halconfig_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_halconfig_proto_goTypes,
		DependencyIndexes: file_halconfig_proto_depIdxs,
		MessageInfos:      file_halconfig_proto_msgTypes,
	}.Build()
	File_halconfig_proto = out.File
	file_halconfig_proto_rawDesc = nil
	file_halconfig_proto_goTypes = nil
	file_halconfig_proto_depIdxs = nil
}
