// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.2
// source: artifact/artifacts.proto

package artifact

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

type ArtifactProviders struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bitbucket *BitbucketArtifactProvider `protobuf:"bytes,1,opt,name=bitbucket,proto3" json:"bitbucket,omitempty"`
	Gcs       *GcsArtifactProvider       `protobuf:"bytes,2,opt,name=gcs,proto3" json:"gcs,omitempty"`
	Github    *GitHubArtifactProvider    `protobuf:"bytes,3,opt,name=github,proto3" json:"github,omitempty"`
	Gitlab    *GitLabArtifactProvider    `protobuf:"bytes,4,opt,name=gitlab,proto3" json:"gitlab,omitempty"`
	Gitrepo   *GitRepoArtifactProvider   `protobuf:"bytes,5,opt,name=gitrepo,proto3" json:"gitrepo,omitempty"`
	Helm      *HelmArtifactProvider      `protobuf:"bytes,6,opt,name=helm,proto3" json:"helm,omitempty"`
	Http      *HttpArtifactProvider      `protobuf:"bytes,7,opt,name=http,proto3" json:"http,omitempty"`
	Maven     *MavenArtifactProvider     `protobuf:"bytes,8,opt,name=maven,proto3" json:"maven,omitempty"`
	Oracle    *OracleArtifactProvider    `protobuf:"bytes,9,opt,name=oracle,proto3" json:"oracle,omitempty"`
	S3        *S3ArtifactProvider        `protobuf:"bytes,10,opt,name=s3,proto3" json:"s3,omitempty"`
	Templates []*ArtifactTemplate        `protobuf:"bytes,11,rep,name=templates,proto3" json:"templates,omitempty"`
}

func (x *ArtifactProviders) Reset() {
	*x = ArtifactProviders{}
	if protoimpl.UnsafeEnabled {
		mi := &file_artifact_artifacts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArtifactProviders) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArtifactProviders) ProtoMessage() {}

func (x *ArtifactProviders) ProtoReflect() protoreflect.Message {
	mi := &file_artifact_artifacts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArtifactProviders.ProtoReflect.Descriptor instead.
func (*ArtifactProviders) Descriptor() ([]byte, []int) {
	return file_artifact_artifacts_proto_rawDescGZIP(), []int{0}
}

func (x *ArtifactProviders) GetBitbucket() *BitbucketArtifactProvider {
	if x != nil {
		return x.Bitbucket
	}
	return nil
}

func (x *ArtifactProviders) GetGcs() *GcsArtifactProvider {
	if x != nil {
		return x.Gcs
	}
	return nil
}

func (x *ArtifactProviders) GetGithub() *GitHubArtifactProvider {
	if x != nil {
		return x.Github
	}
	return nil
}

func (x *ArtifactProviders) GetGitlab() *GitLabArtifactProvider {
	if x != nil {
		return x.Gitlab
	}
	return nil
}

func (x *ArtifactProviders) GetGitrepo() *GitRepoArtifactProvider {
	if x != nil {
		return x.Gitrepo
	}
	return nil
}

func (x *ArtifactProviders) GetHelm() *HelmArtifactProvider {
	if x != nil {
		return x.Helm
	}
	return nil
}

func (x *ArtifactProviders) GetHttp() *HttpArtifactProvider {
	if x != nil {
		return x.Http
	}
	return nil
}

func (x *ArtifactProviders) GetMaven() *MavenArtifactProvider {
	if x != nil {
		return x.Maven
	}
	return nil
}

func (x *ArtifactProviders) GetOracle() *OracleArtifactProvider {
	if x != nil {
		return x.Oracle
	}
	return nil
}

func (x *ArtifactProviders) GetS3() *S3ArtifactProvider {
	if x != nil {
		return x.S3
	}
	return nil
}

func (x *ArtifactProviders) GetTemplates() []*ArtifactTemplate {
	if x != nil {
		return x.Templates
	}
	return nil
}

var File_artifact_artifacts_proto protoreflect.FileDescriptor

var file_artifact_artifacts_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x1a, 0x18, 0x61, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x2f, 0x62, 0x69, 0x74, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2f, 0x67,
	0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2f, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x2f, 0x67, 0x69, 0x74, 0x72, 0x65, 0x70, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2f, 0x68, 0x65, 0x6c, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2f, 0x68, 0x74,
	0x74, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x2f, 0x6d, 0x61, 0x76, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15,
	0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2f, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2f,
	0x73, 0x33, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x5f, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x05, 0x0a, 0x11, 0x41,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73,
	0x12, 0x47, 0x0a, 0x09, 0x62, 0x69, 0x74, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x2e, 0x42, 0x69, 0x74, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x41, 0x72,
	0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x09,
	0x62, 0x69, 0x74, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x12, 0x35, 0x0a, 0x03, 0x67, 0x63, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x47, 0x63, 0x73, 0x41, 0x72, 0x74, 0x69, 0x66,
	0x61, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x03, 0x67, 0x63, 0x73,
	0x12, 0x3e, 0x0a, 0x06, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x2e, 0x47, 0x69, 0x74, 0x48, 0x75, 0x62, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x12, 0x3e, 0x0a, 0x06, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x2e, 0x47, 0x69, 0x74, 0x4c, 0x61, 0x62, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x12, 0x41, 0x0a, 0x07, 0x67, 0x69, 0x74, 0x72, 0x65, 0x70, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x2e, 0x47, 0x69, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x07, 0x67, 0x69, 0x74, 0x72,
	0x65, 0x70, 0x6f, 0x12, 0x38, 0x0a, 0x04, 0x68, 0x65, 0x6c, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61,
	0x63, 0x74, 0x2e, 0x48, 0x65, 0x6c, 0x6d, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x04, 0x68, 0x65, 0x6c, 0x6d, 0x12, 0x38, 0x0a,
	0x04, 0x68, 0x74, 0x74, 0x70, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x48, 0x74, 0x74,
	0x70, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x52, 0x04, 0x68, 0x74, 0x74, 0x70, 0x12, 0x3b, 0x0a, 0x05, 0x6d, 0x61, 0x76, 0x65, 0x6e,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61,
	0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x4d, 0x61, 0x76, 0x65, 0x6e, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6d,
	0x61, 0x76, 0x65, 0x6e, 0x12, 0x3e, 0x0a, 0x06, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x41, 0x72, 0x74, 0x69,
	0x66, 0x61, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x06, 0x6f, 0x72,
	0x61, 0x63, 0x6c, 0x65, 0x12, 0x32, 0x0a, 0x02, 0x73, 0x33, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63,
	0x74, 0x2e, 0x53, 0x33, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x52, 0x02, 0x73, 0x33, 0x12, 0x3e, 0x0a, 0x09, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x66, 0x61, 0x63, 0x74, 0x54, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x52, 0x09, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x70, 0x69, 0x6e, 0x6e, 0x61, 0x6b, 0x65, 0x72,
	0x2f, 0x6b, 0x6c, 0x65, 0x61, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_artifact_artifacts_proto_rawDescOnce sync.Once
	file_artifact_artifacts_proto_rawDescData = file_artifact_artifacts_proto_rawDesc
)

func file_artifact_artifacts_proto_rawDescGZIP() []byte {
	file_artifact_artifacts_proto_rawDescOnce.Do(func() {
		file_artifact_artifacts_proto_rawDescData = protoimpl.X.CompressGZIP(file_artifact_artifacts_proto_rawDescData)
	})
	return file_artifact_artifacts_proto_rawDescData
}

var file_artifact_artifacts_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_artifact_artifacts_proto_goTypes = []interface{}{
	(*ArtifactProviders)(nil),         // 0: proto.artifact.ArtifactProviders
	(*BitbucketArtifactProvider)(nil), // 1: proto.artifact.BitbucketArtifactProvider
	(*GcsArtifactProvider)(nil),       // 2: proto.artifact.GcsArtifactProvider
	(*GitHubArtifactProvider)(nil),    // 3: proto.artifact.GitHubArtifactProvider
	(*GitLabArtifactProvider)(nil),    // 4: proto.artifact.GitLabArtifactProvider
	(*GitRepoArtifactProvider)(nil),   // 5: proto.artifact.GitRepoArtifactProvider
	(*HelmArtifactProvider)(nil),      // 6: proto.artifact.HelmArtifactProvider
	(*HttpArtifactProvider)(nil),      // 7: proto.artifact.HttpArtifactProvider
	(*MavenArtifactProvider)(nil),     // 8: proto.artifact.MavenArtifactProvider
	(*OracleArtifactProvider)(nil),    // 9: proto.artifact.OracleArtifactProvider
	(*S3ArtifactProvider)(nil),        // 10: proto.artifact.S3ArtifactProvider
	(*ArtifactTemplate)(nil),          // 11: proto.artifact.ArtifactTemplate
}
var file_artifact_artifacts_proto_depIdxs = []int32{
	1,  // 0: proto.artifact.ArtifactProviders.bitbucket:type_name -> proto.artifact.BitbucketArtifactProvider
	2,  // 1: proto.artifact.ArtifactProviders.gcs:type_name -> proto.artifact.GcsArtifactProvider
	3,  // 2: proto.artifact.ArtifactProviders.github:type_name -> proto.artifact.GitHubArtifactProvider
	4,  // 3: proto.artifact.ArtifactProviders.gitlab:type_name -> proto.artifact.GitLabArtifactProvider
	5,  // 4: proto.artifact.ArtifactProviders.gitrepo:type_name -> proto.artifact.GitRepoArtifactProvider
	6,  // 5: proto.artifact.ArtifactProviders.helm:type_name -> proto.artifact.HelmArtifactProvider
	7,  // 6: proto.artifact.ArtifactProviders.http:type_name -> proto.artifact.HttpArtifactProvider
	8,  // 7: proto.artifact.ArtifactProviders.maven:type_name -> proto.artifact.MavenArtifactProvider
	9,  // 8: proto.artifact.ArtifactProviders.oracle:type_name -> proto.artifact.OracleArtifactProvider
	10, // 9: proto.artifact.ArtifactProviders.s3:type_name -> proto.artifact.S3ArtifactProvider
	11, // 10: proto.artifact.ArtifactProviders.templates:type_name -> proto.artifact.ArtifactTemplate
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_artifact_artifacts_proto_init() }
func file_artifact_artifacts_proto_init() {
	if File_artifact_artifacts_proto != nil {
		return
	}
	file_artifact_bitbucket_proto_init()
	file_artifact_gcs_proto_init()
	file_artifact_github_proto_init()
	file_artifact_gitlab_proto_init()
	file_artifact_gitrepo_proto_init()
	file_artifact_helm_proto_init()
	file_artifact_http_proto_init()
	file_artifact_maven_proto_init()
	file_artifact_oracle_proto_init()
	file_artifact_s3_proto_init()
	file_artifact_artifact_template_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_artifact_artifacts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArtifactProviders); i {
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
			RawDescriptor: file_artifact_artifacts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_artifact_artifacts_proto_goTypes,
		DependencyIndexes: file_artifact_artifacts_proto_depIdxs,
		MessageInfos:      file_artifact_artifacts_proto_msgTypes,
	}.Build()
	File_artifact_artifacts_proto = out.File
	file_artifact_artifacts_proto_rawDesc = nil
	file_artifact_artifacts_proto_goTypes = nil
	file_artifact_artifacts_proto_depIdxs = nil
}
