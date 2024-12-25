// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        (unknown)
// source: types/v1/query.proto

package typesv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Data struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Schema schema = 1;
	//
	// Types that are valid to be assigned to Shape:
	//
	//	*Data_Tabular
	//	*Data_ScalarTimeseries
	//	*Data_VectorTimeseries
	Shape         isData_Shape `protobuf_oneof:"shape"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Data) Reset() {
	*x = Data{}
	mi := &file_types_v1_query_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Data) ProtoMessage() {}

func (x *Data) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_query_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Data.ProtoReflect.Descriptor instead.
func (*Data) Descriptor() ([]byte, []int) {
	return file_types_v1_query_proto_rawDescGZIP(), []int{0}
}

func (x *Data) GetShape() isData_Shape {
	if x != nil {
		return x.Shape
	}
	return nil
}

func (x *Data) GetTabular() *Tabular {
	if x != nil {
		if x, ok := x.Shape.(*Data_Tabular); ok {
			return x.Tabular
		}
	}
	return nil
}

func (x *Data) GetScalarTimeseries() *ScalarTimeseries {
	if x != nil {
		if x, ok := x.Shape.(*Data_ScalarTimeseries); ok {
			return x.ScalarTimeseries
		}
	}
	return nil
}

func (x *Data) GetVectorTimeseries() *VectorTimeseries {
	if x != nil {
		if x, ok := x.Shape.(*Data_VectorTimeseries); ok {
			return x.VectorTimeseries
		}
	}
	return nil
}

type isData_Shape interface {
	isData_Shape()
}

type Data_Tabular struct {
	// tabular is data that has a table header, followed by rows that match the
	// header
	Tabular *Tabular `protobuf:"bytes,301,opt,name=tabular,proto3,oneof"`
}

type Data_ScalarTimeseries struct {
	// scalar timeseries data is a list of [ts, scalar] where scalar is usually
	// a numeric value
	ScalarTimeseries *ScalarTimeseries `protobuf:"bytes,401,opt,name=scalar_timeseries,json=scalarTimeseries,proto3,oneof"`
}

type Data_VectorTimeseries struct {
	// vector timeseries data is a list of [ts, [scalar0, ..., scalarN]] where
	// scalar is usually a numeric value
	VectorTimeseries *VectorTimeseries `protobuf:"bytes,402,opt,name=vector_timeseries,json=vectorTimeseries,proto3,oneof"`
}

func (*Data_Tabular) isData_Shape() {}

func (*Data_ScalarTimeseries) isData_Shape() {}

func (*Data_VectorTimeseries) isData_Shape() {}

type Tabular struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Shape:
	//
	//	*Tabular_LogEvents
	//	*Tabular_FreeForm
	Shape         isTabular_Shape `protobuf_oneof:"shape"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Tabular) Reset() {
	*x = Tabular{}
	mi := &file_types_v1_query_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Tabular) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tabular) ProtoMessage() {}

func (x *Tabular) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_query_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tabular.ProtoReflect.Descriptor instead.
func (*Tabular) Descriptor() ([]byte, []int) {
	return file_types_v1_query_proto_rawDescGZIP(), []int{1}
}

func (x *Tabular) GetShape() isTabular_Shape {
	if x != nil {
		return x.Shape
	}
	return nil
}

func (x *Tabular) GetLogEvents() *LogEventGroup {
	if x != nil {
		if x, ok := x.Shape.(*Tabular_LogEvents); ok {
			return x.LogEvents
		}
	}
	return nil
}

func (x *Tabular) GetFreeForm() *Table {
	if x != nil {
		if x, ok := x.Shape.(*Tabular_FreeForm); ok {
			return x.FreeForm
		}
	}
	return nil
}

type isTabular_Shape interface {
	isTabular_Shape()
}

type Tabular_LogEvents struct {
	LogEvents *LogEventGroup `protobuf:"bytes,201,opt,name=log_events,json=logEvents,proto3,oneof"`
}

type Tabular_FreeForm struct {
	FreeForm *Table `protobuf:"bytes,202,opt,name=free_form,json=freeForm,proto3,oneof"`
}

func (*Tabular_LogEvents) isTabular_Shape() {}

func (*Tabular_FreeForm) isTabular_Shape() {}

type ScalarTimeseries struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// scalars will contain only this type of value
	Type          *VarType           `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Scalars       []*ScalarTimestamp `protobuf:"bytes,3,rep,name=scalars,proto3" json:"scalars,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScalarTimeseries) Reset() {
	*x = ScalarTimeseries{}
	mi := &file_types_v1_query_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScalarTimeseries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScalarTimeseries) ProtoMessage() {}

func (x *ScalarTimeseries) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_query_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScalarTimeseries.ProtoReflect.Descriptor instead.
func (*ScalarTimeseries) Descriptor() ([]byte, []int) {
	return file_types_v1_query_proto_rawDescGZIP(), []int{2}
}

func (x *ScalarTimeseries) GetType() *VarType {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *ScalarTimeseries) GetScalars() []*ScalarTimestamp {
	if x != nil {
		return x.Scalars
	}
	return nil
}

type ScalarTimestamp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ts            *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Scalar        *Scalar                `protobuf:"bytes,2,opt,name=scalar,proto3" json:"scalar,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ScalarTimestamp) Reset() {
	*x = ScalarTimestamp{}
	mi := &file_types_v1_query_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScalarTimestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScalarTimestamp) ProtoMessage() {}

func (x *ScalarTimestamp) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_query_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScalarTimestamp.ProtoReflect.Descriptor instead.
func (*ScalarTimestamp) Descriptor() ([]byte, []int) {
	return file_types_v1_query_proto_rawDescGZIP(), []int{3}
}

func (x *ScalarTimestamp) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *ScalarTimestamp) GetScalar() *Scalar {
	if x != nil {
		return x.Scalar
	}
	return nil
}

type VectorTimeseries struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Context *Obj                   `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	// vectors will contain only this type of value
	Type          *VarType           `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Vectors       []*VectorTimestamp `protobuf:"bytes,3,rep,name=vectors,proto3" json:"vectors,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VectorTimeseries) Reset() {
	*x = VectorTimeseries{}
	mi := &file_types_v1_query_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VectorTimeseries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VectorTimeseries) ProtoMessage() {}

func (x *VectorTimeseries) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_query_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VectorTimeseries.ProtoReflect.Descriptor instead.
func (*VectorTimeseries) Descriptor() ([]byte, []int) {
	return file_types_v1_query_proto_rawDescGZIP(), []int{4}
}

func (x *VectorTimeseries) GetContext() *Obj {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *VectorTimeseries) GetType() *VarType {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *VectorTimeseries) GetVectors() []*VectorTimestamp {
	if x != nil {
		return x.Vectors
	}
	return nil
}

type VectorTimestamp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Ts            *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	Vector        []*Scalar              `protobuf:"bytes,2,rep,name=vector,proto3" json:"vector,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VectorTimestamp) Reset() {
	*x = VectorTimestamp{}
	mi := &file_types_v1_query_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VectorTimestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VectorTimestamp) ProtoMessage() {}

func (x *VectorTimestamp) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_query_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VectorTimestamp.ProtoReflect.Descriptor instead.
func (*VectorTimestamp) Descriptor() ([]byte, []int) {
	return file_types_v1_query_proto_rawDescGZIP(), []int{5}
}

func (x *VectorTimestamp) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *VectorTimestamp) GetVector() []*Scalar {
	if x != nil {
		return x.Vector
	}
	return nil
}

var File_types_v1_query_proto protoreflect.FileDescriptor

var file_types_v1_query_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd7, 0x01, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x2e, 0x0a, 0x07, 0x74, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x18, 0xad, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x54, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x48, 0x00, 0x52, 0x07, 0x74, 0x61, 0x62, 0x75, 0x6c,
	0x61, 0x72, 0x12, 0x4a, 0x0a, 0x11, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x18, 0x91, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x48, 0x00, 0x52, 0x10, 0x73, 0x63,
	0x61, 0x6c, 0x61, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x4a,
	0x0a, 0x11, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x92, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x48, 0x00, 0x52, 0x10, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x42, 0x07, 0x0a, 0x05, 0x73, 0x68,
	0x61, 0x70, 0x65, 0x22, 0x7e, 0x0a, 0x07, 0x54, 0x61, 0x62, 0x75, 0x6c, 0x61, 0x72, 0x12, 0x39,
	0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0xc9, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x48, 0x00, 0x52, 0x09,
	0x6c, 0x6f, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2f, 0x0a, 0x09, 0x66, 0x72, 0x65,
	0x65, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0xca, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x00,
	0x52, 0x08, 0x66, 0x72, 0x65, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x42, 0x07, 0x0a, 0x05, 0x73, 0x68,
	0x61, 0x70, 0x65, 0x22, 0x6e, 0x0a, 0x10, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x56, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x33,
	0x0a, 0x07, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x61,
	0x72, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x73, 0x63, 0x61, 0x6c,
	0x61, 0x72, 0x73, 0x22, 0x67, 0x0a, 0x0f, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02,
	0x74, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x63,
	0x61, 0x6c, 0x61, 0x72, 0x52, 0x06, 0x73, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x22, 0x97, 0x01, 0x0a,
	0x10, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x65, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x27, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x62,
	0x6a, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x33, 0x0a, 0x07, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x65,
	0x63, 0x74, 0x6f, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x07, 0x76,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x22, 0x67, 0x0a, 0x0f, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x28, 0x0a, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x53, 0x63, 0x61, 0x6c, 0x61, 0x72, 0x52, 0x06, 0x76, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x42,
	0x8a, 0x01, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x42, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x75, 0x6d, 0x61, 0x6e,
	0x6c, 0x6f, 0x67, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x54, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x08, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x09, 0x54, 0x79, 0x70, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_types_v1_query_proto_rawDescOnce sync.Once
	file_types_v1_query_proto_rawDescData = file_types_v1_query_proto_rawDesc
)

func file_types_v1_query_proto_rawDescGZIP() []byte {
	file_types_v1_query_proto_rawDescOnce.Do(func() {
		file_types_v1_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_v1_query_proto_rawDescData)
	})
	return file_types_v1_query_proto_rawDescData
}

var file_types_v1_query_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_types_v1_query_proto_goTypes = []any{
	(*Data)(nil),                  // 0: types.v1.Data
	(*Tabular)(nil),               // 1: types.v1.Tabular
	(*ScalarTimeseries)(nil),      // 2: types.v1.ScalarTimeseries
	(*ScalarTimestamp)(nil),       // 3: types.v1.ScalarTimestamp
	(*VectorTimeseries)(nil),      // 4: types.v1.VectorTimeseries
	(*VectorTimestamp)(nil),       // 5: types.v1.VectorTimestamp
	(*LogEventGroup)(nil),         // 6: types.v1.LogEventGroup
	(*Table)(nil),                 // 7: types.v1.Table
	(*VarType)(nil),               // 8: types.v1.VarType
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*Scalar)(nil),                // 10: types.v1.Scalar
	(*Obj)(nil),                   // 11: types.v1.Obj
}
var file_types_v1_query_proto_depIdxs = []int32{
	1,  // 0: types.v1.Data.tabular:type_name -> types.v1.Tabular
	2,  // 1: types.v1.Data.scalar_timeseries:type_name -> types.v1.ScalarTimeseries
	4,  // 2: types.v1.Data.vector_timeseries:type_name -> types.v1.VectorTimeseries
	6,  // 3: types.v1.Tabular.log_events:type_name -> types.v1.LogEventGroup
	7,  // 4: types.v1.Tabular.free_form:type_name -> types.v1.Table
	8,  // 5: types.v1.ScalarTimeseries.type:type_name -> types.v1.VarType
	3,  // 6: types.v1.ScalarTimeseries.scalars:type_name -> types.v1.ScalarTimestamp
	9,  // 7: types.v1.ScalarTimestamp.ts:type_name -> google.protobuf.Timestamp
	10, // 8: types.v1.ScalarTimestamp.scalar:type_name -> types.v1.Scalar
	11, // 9: types.v1.VectorTimeseries.context:type_name -> types.v1.Obj
	8,  // 10: types.v1.VectorTimeseries.type:type_name -> types.v1.VarType
	5,  // 11: types.v1.VectorTimeseries.vectors:type_name -> types.v1.VectorTimestamp
	9,  // 12: types.v1.VectorTimestamp.ts:type_name -> google.protobuf.Timestamp
	10, // 13: types.v1.VectorTimestamp.vector:type_name -> types.v1.Scalar
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_types_v1_query_proto_init() }
func file_types_v1_query_proto_init() {
	if File_types_v1_query_proto != nil {
		return
	}
	file_types_v1_logevent_proto_init()
	file_types_v1_logquery_proto_init()
	file_types_v1_types_proto_init()
	file_types_v1_query_proto_msgTypes[0].OneofWrappers = []any{
		(*Data_Tabular)(nil),
		(*Data_ScalarTimeseries)(nil),
		(*Data_VectorTimeseries)(nil),
	}
	file_types_v1_query_proto_msgTypes[1].OneofWrappers = []any{
		(*Tabular_LogEvents)(nil),
		(*Tabular_FreeForm)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_types_v1_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_v1_query_proto_goTypes,
		DependencyIndexes: file_types_v1_query_proto_depIdxs,
		MessageInfos:      file_types_v1_query_proto_msgTypes,
	}.Build()
	File_types_v1_query_proto = out.File
	file_types_v1_query_proto_rawDesc = nil
	file_types_v1_query_proto_goTypes = nil
	file_types_v1_query_proto_depIdxs = nil
}
