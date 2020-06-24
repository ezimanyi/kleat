// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: config/deck_env.proto

package config

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

// Environment variables to be set when running deck.
type DeckEnv struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Path to the .crt file containing deck's SSL certificate.
	DeckCert string `protobuf:"bytes,1,opt,name=deckCert,json=DECK_CERT,proto3" json:"deckCert,omitempty"`
	// Path to the .key file for deck's SSL certificate.
	DeckKey string `protobuf:"bytes,2,opt,name=deckKey,json=DECK_KEY,proto3" json:"deckKey,omitempty"`
	// The passphrase for deck's SSL certificate.
	Passphrase string `protobuf:"bytes,3,opt,name=passphrase,json=PASSPHRASE,proto3" json:"passphrase,omitempty"`
}

func (x *DeckEnv) Reset() {
	*x = DeckEnv{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_deck_env_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeckEnv) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeckEnv) ProtoMessage() {}

func (x *DeckEnv) ProtoReflect() protoreflect.Message {
	mi := &file_config_deck_env_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeckEnv.ProtoReflect.Descriptor instead.
func (*DeckEnv) Descriptor() ([]byte, []int) {
	return file_config_deck_env_proto_rawDescGZIP(), []int{0}
}

func (x *DeckEnv) GetDeckCert() string {
	if x != nil {
		return x.DeckCert
	}
	return ""
}

func (x *DeckEnv) GetDeckKey() string {
	if x != nil {
		return x.DeckKey
	}
	return ""
}

func (x *DeckEnv) GetPassphrase() string {
	if x != nil {
		return x.Passphrase
	}
	return ""
}

var File_config_deck_env_proto protoreflect.FileDescriptor

var file_config_deck_env_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x64, 0x65, 0x63, 0x6b, 0x5f, 0x65, 0x6e,
	0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x61, 0x0a, 0x07, 0x44, 0x65, 0x63, 0x6b, 0x45, 0x6e, 0x76,
	0x12, 0x1b, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x6b, 0x43, 0x65, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x44, 0x45, 0x43, 0x4b, 0x5f, 0x43, 0x45, 0x52, 0x54, 0x12, 0x19, 0x0a,
	0x07, 0x64, 0x65, 0x63, 0x6b, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x44, 0x45, 0x43, 0x4b, 0x5f, 0x4b, 0x45, 0x59, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x73, 0x73,
	0x70, 0x68, 0x72, 0x61, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x41,
	0x53, 0x53, 0x50, 0x48, 0x52, 0x41, 0x53, 0x45, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72,
	0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_deck_env_proto_rawDescOnce sync.Once
	file_config_deck_env_proto_rawDescData = file_config_deck_env_proto_rawDesc
)

func file_config_deck_env_proto_rawDescGZIP() []byte {
	file_config_deck_env_proto_rawDescOnce.Do(func() {
		file_config_deck_env_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_deck_env_proto_rawDescData)
	})
	return file_config_deck_env_proto_rawDescData
}

var file_config_deck_env_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_config_deck_env_proto_goTypes = []interface{}{
	(*DeckEnv)(nil), // 0: proto.config.DeckEnv
}
var file_config_deck_env_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_config_deck_env_proto_init() }
func file_config_deck_env_proto_init() {
	if File_config_deck_env_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_deck_env_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeckEnv); i {
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
			RawDescriptor: file_config_deck_env_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_config_deck_env_proto_goTypes,
		DependencyIndexes: file_config_deck_env_proto_depIdxs,
		MessageInfos:      file_config_deck_env_proto_msgTypes,
	}.Build()
	File_config_deck_env_proto = out.File
	file_config_deck_env_proto_rawDesc = nil
	file_config_deck_env_proto_goTypes = nil
	file_config_deck_env_proto_depIdxs = nil
}
