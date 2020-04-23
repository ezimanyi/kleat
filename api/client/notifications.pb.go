// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: notifications.proto

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

// Configuration for Slack notifications.
type SlackNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether Slack notifications are enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// The name of your Slack bot.
	BotName string `protobuf:"bytes,2,opt,name=botName,proto3" json:"botName,omitempty"`
	// Your Slack bot token.
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	// Slack endpoint. Optional, can only be set if using a compatible API.
	BaseUrl string `protobuf:"bytes,4,opt,name=baseUrl,proto3" json:"baseUrl,omitempty"`
	// Force usage of incoming webhooks endpoint for slack. Optional, only set if
	// using a compatible API.
	ForceUseIncomingWebhook bool `protobuf:"varint,5,opt,name=forceUseIncomingWebhook,proto3" json:"forceUseIncomingWebhook,omitempty"`
}

func (x *SlackNotification) Reset() {
	*x = SlackNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SlackNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SlackNotification) ProtoMessage() {}

func (x *SlackNotification) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SlackNotification.ProtoReflect.Descriptor instead.
func (*SlackNotification) Descriptor() ([]byte, []int) {
	return file_notifications_proto_rawDescGZIP(), []int{0}
}

func (x *SlackNotification) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *SlackNotification) GetBotName() string {
	if x != nil {
		return x.BotName
	}
	return ""
}

func (x *SlackNotification) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SlackNotification) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

func (x *SlackNotification) GetForceUseIncomingWebhook() bool {
	if x != nil {
		return x.ForceUseIncomingWebhook
	}
	return false
}

// Configuration for Twilio notifications.
type TwilioNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether Twilio notifications are enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Your Twilio account SID.
	Account string `protobuf:"bytes,2,opt,name=account,proto3" json:"account,omitempty"`
	// Your Twilio auth token.
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	// The endpoint of the Twilio API. Optional, only set if overriding the default.
	BaseUrl string `protobuf:"bytes,4,opt,name=baseUrl,proto3" json:"baseUrl,omitempty"`
	// The phone number from which the SMS will be sent (i.e. +1234-567-8910).
	From string `protobuf:"bytes,5,opt,name=from,proto3" json:"from,omitempty"`
}

func (x *TwilioNotification) Reset() {
	*x = TwilioNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TwilioNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TwilioNotification) ProtoMessage() {}

func (x *TwilioNotification) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TwilioNotification.ProtoReflect.Descriptor instead.
func (*TwilioNotification) Descriptor() ([]byte, []int) {
	return file_notifications_proto_rawDescGZIP(), []int{1}
}

func (x *TwilioNotification) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *TwilioNotification) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *TwilioNotification) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *TwilioNotification) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

func (x *TwilioNotification) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

// Configuration for Github status notifications.
type GithubStatusNotification struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Whether Github status notifications are enabled.
	Enabled bool `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	// Your github account token.
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GithubStatusNotification) Reset() {
	*x = GithubStatusNotification{}
	if protoimpl.UnsafeEnabled {
		mi := &file_notifications_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GithubStatusNotification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GithubStatusNotification) ProtoMessage() {}

func (x *GithubStatusNotification) ProtoReflect() protoreflect.Message {
	mi := &file_notifications_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GithubStatusNotification.ProtoReflect.Descriptor instead.
func (*GithubStatusNotification) Descriptor() ([]byte, []int) {
	return file_notifications_proto_rawDescGZIP(), []int{2}
}

func (x *GithubStatusNotification) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *GithubStatusNotification) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

var File_notifications_proto protoreflect.FileDescriptor

var file_notifications_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a,
	0x11, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x62, 0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62,
	0x6f, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x62, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62,
	0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x38, 0x0a, 0x17, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x55,
	0x73, 0x65, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f,
	0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x17, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x55, 0x73,
	0x65, 0x49, 0x6e, 0x63, 0x6f, 0x6d, 0x69, 0x6e, 0x67, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x22, 0x8c, 0x01, 0x0a, 0x12, 0x54, 0x77, 0x69, 0x6c, 0x69, 0x6f, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x22,
	0x4a, 0x0a, 0x18, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x27, 0x5a, 0x25, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61,
	0x6b, 0x65, 0x72, 0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_notifications_proto_rawDescOnce sync.Once
	file_notifications_proto_rawDescData = file_notifications_proto_rawDesc
)

func file_notifications_proto_rawDescGZIP() []byte {
	file_notifications_proto_rawDescOnce.Do(func() {
		file_notifications_proto_rawDescData = protoimpl.X.CompressGZIP(file_notifications_proto_rawDescData)
	})
	return file_notifications_proto_rawDescData
}

var file_notifications_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_notifications_proto_goTypes = []interface{}{
	(*SlackNotification)(nil),        // 0: proto.SlackNotification
	(*TwilioNotification)(nil),       // 1: proto.TwilioNotification
	(*GithubStatusNotification)(nil), // 2: proto.GithubStatusNotification
}
var file_notifications_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_notifications_proto_init() }
func file_notifications_proto_init() {
	if File_notifications_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_notifications_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SlackNotification); i {
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
		file_notifications_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TwilioNotification); i {
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
		file_notifications_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GithubStatusNotification); i {
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
			RawDescriptor: file_notifications_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_notifications_proto_goTypes,
		DependencyIndexes: file_notifications_proto_depIdxs,
		MessageInfos:      file_notifications_proto_msgTypes,
	}.Build()
	File_notifications_proto = out.File
	file_notifications_proto_rawDesc = nil
	file_notifications_proto_goTypes = nil
	file_notifications_proto_depIdxs = nil
}
