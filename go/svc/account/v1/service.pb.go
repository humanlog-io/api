// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: svc/account/v1/service.proto

package accountv1

import (
	v1 "github.com/humanlogio/api/go/types/v1"
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

type ListMachineRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cursor    *v1.Cursor `protobuf:"bytes,1,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Limit     int32      `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	AccountId int64      `protobuf:"varint,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *ListMachineRequest) Reset() {
	*x = ListMachineRequest{}
	mi := &file_svc_account_v1_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListMachineRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMachineRequest) ProtoMessage() {}

func (x *ListMachineRequest) ProtoReflect() protoreflect.Message {
	mi := &file_svc_account_v1_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMachineRequest.ProtoReflect.Descriptor instead.
func (*ListMachineRequest) Descriptor() ([]byte, []int) {
	return file_svc_account_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListMachineRequest) GetCursor() *v1.Cursor {
	if x != nil {
		return x.Cursor
	}
	return nil
}

func (x *ListMachineRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *ListMachineRequest) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

type ListMachineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Next  *v1.Cursor                      `protobuf:"bytes,1,opt,name=next,proto3" json:"next,omitempty"`
	Items []*ListMachineResponse_ListItem `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListMachineResponse) Reset() {
	*x = ListMachineResponse{}
	mi := &file_svc_account_v1_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListMachineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMachineResponse) ProtoMessage() {}

func (x *ListMachineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_svc_account_v1_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMachineResponse.ProtoReflect.Descriptor instead.
func (*ListMachineResponse) Descriptor() ([]byte, []int) {
	return file_svc_account_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListMachineResponse) GetNext() *v1.Cursor {
	if x != nil {
		return x.Next
	}
	return nil
}

func (x *ListMachineResponse) GetItems() []*ListMachineResponse_ListItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type ListMachineResponse_ListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Machine *v1.Machine `protobuf:"bytes,1,opt,name=machine,proto3" json:"machine,omitempty"`
}

func (x *ListMachineResponse_ListItem) Reset() {
	*x = ListMachineResponse_ListItem{}
	mi := &file_svc_account_v1_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListMachineResponse_ListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMachineResponse_ListItem) ProtoMessage() {}

func (x *ListMachineResponse_ListItem) ProtoReflect() protoreflect.Message {
	mi := &file_svc_account_v1_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMachineResponse_ListItem.ProtoReflect.Descriptor instead.
func (*ListMachineResponse_ListItem) Descriptor() ([]byte, []int) {
	return file_svc_account_v1_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ListMachineResponse_ListItem) GetMachine() *v1.Machine {
	if x != nil {
		return x.Machine
	}
	return nil
}

var File_svc_account_v1_service_proto protoreflect.FileDescriptor

var file_svc_account_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x76, 0x63, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x73, 0x76, 0x63, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x16,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x63,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x52, 0x06, 0x63,
	0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xb8, 0x01, 0x0a, 0x13, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x75, 0x72, 0x73,
	0x6f, 0x72, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x12, 0x42, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x63,
	0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x1a, 0x37, 0x0a, 0x08,
	0x4c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x2b, 0x0a, 0x07, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x07, 0x6d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x32, 0x68, 0x0a, 0x0e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x56, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x22, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x73, 0x76, 0x63,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0xb3, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x76, 0x63, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x6c, 0x6f, 0x67, 0x69, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x76, 0x63, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x3b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x53, 0x41, 0x58, 0xaa, 0x02, 0x0e, 0x53, 0x76, 0x63, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0e, 0x53, 0x76, 0x63, 0x5c, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1a, 0x53, 0x76, 0x63, 0x5c, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x10, 0x53, 0x76, 0x63, 0x3a, 0x3a, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_svc_account_v1_service_proto_rawDescOnce sync.Once
	file_svc_account_v1_service_proto_rawDescData = file_svc_account_v1_service_proto_rawDesc
)

func file_svc_account_v1_service_proto_rawDescGZIP() []byte {
	file_svc_account_v1_service_proto_rawDescOnce.Do(func() {
		file_svc_account_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_svc_account_v1_service_proto_rawDescData)
	})
	return file_svc_account_v1_service_proto_rawDescData
}

var file_svc_account_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_svc_account_v1_service_proto_goTypes = []any{
	(*ListMachineRequest)(nil),           // 0: svc.account.v1.ListMachineRequest
	(*ListMachineResponse)(nil),          // 1: svc.account.v1.ListMachineResponse
	(*ListMachineResponse_ListItem)(nil), // 2: svc.account.v1.ListMachineResponse.ListItem
	(*v1.Cursor)(nil),                    // 3: types.v1.Cursor
	(*v1.Machine)(nil),                   // 4: types.v1.Machine
}
var file_svc_account_v1_service_proto_depIdxs = []int32{
	3, // 0: svc.account.v1.ListMachineRequest.cursor:type_name -> types.v1.Cursor
	3, // 1: svc.account.v1.ListMachineResponse.next:type_name -> types.v1.Cursor
	2, // 2: svc.account.v1.ListMachineResponse.items:type_name -> svc.account.v1.ListMachineResponse.ListItem
	4, // 3: svc.account.v1.ListMachineResponse.ListItem.machine:type_name -> types.v1.Machine
	0, // 4: svc.account.v1.AccountService.ListMachine:input_type -> svc.account.v1.ListMachineRequest
	1, // 5: svc.account.v1.AccountService.ListMachine:output_type -> svc.account.v1.ListMachineResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_svc_account_v1_service_proto_init() }
func file_svc_account_v1_service_proto_init() {
	if File_svc_account_v1_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_svc_account_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_svc_account_v1_service_proto_goTypes,
		DependencyIndexes: file_svc_account_v1_service_proto_depIdxs,
		MessageInfos:      file_svc_account_v1_service_proto_msgTypes,
	}.Build()
	File_svc_account_v1_service_proto = out.File
	file_svc_account_v1_service_proto_rawDesc = nil
	file_svc_account_v1_service_proto_goTypes = nil
	file_svc_account_v1_service_proto_depIdxs = nil
}
