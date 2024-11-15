// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: types/v1/version_artifact.proto

package typesv1

import (
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

type VersionArtifact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url             string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Sha256          string `protobuf:"bytes,2,opt,name=sha256,proto3" json:"sha256,omitempty"`
	Signature       string `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	Architecture    string `protobuf:"bytes,4,opt,name=architecture,proto3" json:"architecture,omitempty"`
	OperatingSystem string `protobuf:"bytes,5,opt,name=operating_system,json=operatingSystem,proto3" json:"operating_system,omitempty"`
}

func (x *VersionArtifact) Reset() {
	*x = VersionArtifact{}
	mi := &file_types_v1_version_artifact_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VersionArtifact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VersionArtifact) ProtoMessage() {}

func (x *VersionArtifact) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_version_artifact_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VersionArtifact.ProtoReflect.Descriptor instead.
func (*VersionArtifact) Descriptor() ([]byte, []int) {
	return file_types_v1_version_artifact_proto_rawDescGZIP(), []int{0}
}

func (x *VersionArtifact) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *VersionArtifact) GetSha256() string {
	if x != nil {
		return x.Sha256
	}
	return ""
}

func (x *VersionArtifact) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *VersionArtifact) GetArchitecture() string {
	if x != nil {
		return x.Architecture
	}
	return ""
}

func (x *VersionArtifact) GetOperatingSystem() string {
	if x != nil {
		return x.OperatingSystem
	}
	return ""
}

var File_types_v1_version_artifact_proto protoreflect.FileDescriptor

var file_types_v1_version_artifact_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x61, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x22, 0xa8, 0x01, 0x0a, 0x0f,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69,
	0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x74, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x6f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x42, 0x94, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x41, 0x72, 0x74, 0x69, 0x66, 0x61, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x6d, 0x61,
	0x6e, 0x6c, 0x6f, 0x67, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x08, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x09, 0x54, 0x79, 0x70, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_v1_version_artifact_proto_rawDescOnce sync.Once
	file_types_v1_version_artifact_proto_rawDescData = file_types_v1_version_artifact_proto_rawDesc
)

func file_types_v1_version_artifact_proto_rawDescGZIP() []byte {
	file_types_v1_version_artifact_proto_rawDescOnce.Do(func() {
		file_types_v1_version_artifact_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_v1_version_artifact_proto_rawDescData)
	})
	return file_types_v1_version_artifact_proto_rawDescData
}

var file_types_v1_version_artifact_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_types_v1_version_artifact_proto_goTypes = []any{
	(*VersionArtifact)(nil), // 0: types.v1.VersionArtifact
}
var file_types_v1_version_artifact_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_types_v1_version_artifact_proto_init() }
func file_types_v1_version_artifact_proto_init() {
	if File_types_v1_version_artifact_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_types_v1_version_artifact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_v1_version_artifact_proto_goTypes,
		DependencyIndexes: file_types_v1_version_artifact_proto_depIdxs,
		MessageInfos:      file_types_v1_version_artifact_proto_msgTypes,
	}.Build()
	File_types_v1_version_artifact_proto = out.File
	file_types_v1_version_artifact_proto_rawDesc = nil
	file_types_v1_version_artifact_proto_goTypes = nil
	file_types_v1_version_artifact_proto_depIdxs = nil
}
