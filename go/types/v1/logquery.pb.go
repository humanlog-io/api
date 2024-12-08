// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        (unknown)
// source: types/v1/logquery.proto

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

type UnaryOp_Operator int32

const (
	UnaryOp_INVALID UnaryOp_Operator = 0
	UnaryOp_NOT     UnaryOp_Operator = 1
)

// Enum value maps for UnaryOp_Operator.
var (
	UnaryOp_Operator_name = map[int32]string{
		0: "INVALID",
		1: "NOT",
	}
	UnaryOp_Operator_value = map[string]int32{
		"INVALID": 0,
		"NOT":     1,
	}
)

func (x UnaryOp_Operator) Enum() *UnaryOp_Operator {
	p := new(UnaryOp_Operator)
	*p = x
	return p
}

func (x UnaryOp_Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UnaryOp_Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_types_v1_logquery_proto_enumTypes[0].Descriptor()
}

func (UnaryOp_Operator) Type() protoreflect.EnumType {
	return &file_types_v1_logquery_proto_enumTypes[0]
}

func (x UnaryOp_Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UnaryOp_Operator.Descriptor instead.
func (UnaryOp_Operator) EnumDescriptor() ([]byte, []int) {
	return file_types_v1_logquery_proto_rawDescGZIP(), []int{2, 0}
}

type BinaryOp_Operator int32

const (
	BinaryOp_INVALID   BinaryOp_Operator = 0
	BinaryOp_LOG_AND   BinaryOp_Operator = 101
	BinaryOp_LOG_OR    BinaryOp_Operator = 102
	BinaryOp_LOG_XOR   BinaryOp_Operator = 201
	BinaryOp_NUM_ADD   BinaryOp_Operator = 202
	BinaryOp_NUM_SUB   BinaryOp_Operator = 203
	BinaryOp_NUM_DIV   BinaryOp_Operator = 204
	BinaryOp_NUM_MUL   BinaryOp_Operator = 205
	BinaryOp_CMP_EQ    BinaryOp_Operator = 301
	BinaryOp_CMP_NOTEQ BinaryOp_Operator = 302
	BinaryOp_CMP_GT    BinaryOp_Operator = 303
	BinaryOp_CMP_GTE   BinaryOp_Operator = 304
	BinaryOp_CMP_LT    BinaryOp_Operator = 305
	BinaryOp_CMP_LTE   BinaryOp_Operator = 306
)

// Enum value maps for BinaryOp_Operator.
var (
	BinaryOp_Operator_name = map[int32]string{
		0:   "INVALID",
		101: "LOG_AND",
		102: "LOG_OR",
		201: "LOG_XOR",
		202: "NUM_ADD",
		203: "NUM_SUB",
		204: "NUM_DIV",
		205: "NUM_MUL",
		301: "CMP_EQ",
		302: "CMP_NOTEQ",
		303: "CMP_GT",
		304: "CMP_GTE",
		305: "CMP_LT",
		306: "CMP_LTE",
	}
	BinaryOp_Operator_value = map[string]int32{
		"INVALID":   0,
		"LOG_AND":   101,
		"LOG_OR":    102,
		"LOG_XOR":   201,
		"NUM_ADD":   202,
		"NUM_SUB":   203,
		"NUM_DIV":   204,
		"NUM_MUL":   205,
		"CMP_EQ":    301,
		"CMP_NOTEQ": 302,
		"CMP_GT":    303,
		"CMP_GTE":   304,
		"CMP_LT":    305,
		"CMP_LTE":   306,
	}
)

func (x BinaryOp_Operator) Enum() *BinaryOp_Operator {
	p := new(BinaryOp_Operator)
	*p = x
	return p
}

func (x BinaryOp_Operator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BinaryOp_Operator) Descriptor() protoreflect.EnumDescriptor {
	return file_types_v1_logquery_proto_enumTypes[1].Descriptor()
}

func (BinaryOp_Operator) Type() protoreflect.EnumType {
	return &file_types_v1_logquery_proto_enumTypes[1]
}

func (x BinaryOp_Operator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BinaryOp_Operator.Descriptor instead.
func (BinaryOp_Operator) EnumDescriptor() ([]byte, []int) {
	return file_types_v1_logquery_proto_rawDescGZIP(), []int{3, 0}
}

type LogQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From  *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To    *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Query *Expr                  `protobuf:"bytes,100,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *LogQuery) Reset() {
	*x = LogQuery{}
	mi := &file_types_v1_logquery_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogQuery) ProtoMessage() {}

func (x *LogQuery) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_logquery_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogQuery.ProtoReflect.Descriptor instead.
func (*LogQuery) Descriptor() ([]byte, []int) {
	return file_types_v1_logquery_proto_rawDescGZIP(), []int{0}
}

func (x *LogQuery) GetFrom() *timestamppb.Timestamp {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *LogQuery) GetTo() *timestamppb.Timestamp {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *LogQuery) GetQuery() *Expr {
	if x != nil {
		return x.Query
	}
	return nil
}

type Expr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Expr:
	//
	//	*Expr_Literal
	//	*Expr_Unary
	//	*Expr_Binary
	//	*Expr_FuncCall
	Expr isExpr_Expr `protobuf_oneof:"expr"`
	Next *Expr       `protobuf:"bytes,5,opt,name=next,proto3,oneof" json:"next,omitempty"`
}

func (x *Expr) Reset() {
	*x = Expr{}
	mi := &file_types_v1_logquery_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Expr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Expr) ProtoMessage() {}

func (x *Expr) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_logquery_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Expr.ProtoReflect.Descriptor instead.
func (*Expr) Descriptor() ([]byte, []int) {
	return file_types_v1_logquery_proto_rawDescGZIP(), []int{1}
}

func (m *Expr) GetExpr() isExpr_Expr {
	if m != nil {
		return m.Expr
	}
	return nil
}

func (x *Expr) GetLiteral() *Val {
	if x, ok := x.GetExpr().(*Expr_Literal); ok {
		return x.Literal
	}
	return nil
}

func (x *Expr) GetUnary() *UnaryOp {
	if x, ok := x.GetExpr().(*Expr_Unary); ok {
		return x.Unary
	}
	return nil
}

func (x *Expr) GetBinary() *BinaryOp {
	if x, ok := x.GetExpr().(*Expr_Binary); ok {
		return x.Binary
	}
	return nil
}

func (x *Expr) GetFuncCall() *FuncCall {
	if x, ok := x.GetExpr().(*Expr_FuncCall); ok {
		return x.FuncCall
	}
	return nil
}

func (x *Expr) GetNext() *Expr {
	if x != nil {
		return x.Next
	}
	return nil
}

type isExpr_Expr interface {
	isExpr_Expr()
}

type Expr_Literal struct {
	Literal *Val `protobuf:"bytes,1,opt,name=literal,proto3,oneof"`
}

type Expr_Unary struct {
	Unary *UnaryOp `protobuf:"bytes,2,opt,name=unary,proto3,oneof"`
}

type Expr_Binary struct {
	Binary *BinaryOp `protobuf:"bytes,3,opt,name=binary,proto3,oneof"`
}

type Expr_FuncCall struct {
	FuncCall *FuncCall `protobuf:"bytes,4,opt,name=func_call,json=funcCall,proto3,oneof"`
}

func (*Expr_Literal) isExpr_Expr() {}

func (*Expr_Unary) isExpr_Expr() {}

func (*Expr_Binary) isExpr_Expr() {}

func (*Expr_FuncCall) isExpr_Expr() {}

type UnaryOp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op  UnaryOp_Operator `protobuf:"varint,1,opt,name=op,proto3,enum=types.v1.UnaryOp_Operator" json:"op,omitempty"`
	Arg *Expr            `protobuf:"bytes,2,opt,name=arg,proto3" json:"arg,omitempty"`
}

func (x *UnaryOp) Reset() {
	*x = UnaryOp{}
	mi := &file_types_v1_logquery_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UnaryOp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnaryOp) ProtoMessage() {}

func (x *UnaryOp) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_logquery_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnaryOp.ProtoReflect.Descriptor instead.
func (*UnaryOp) Descriptor() ([]byte, []int) {
	return file_types_v1_logquery_proto_rawDescGZIP(), []int{2}
}

func (x *UnaryOp) GetOp() UnaryOp_Operator {
	if x != nil {
		return x.Op
	}
	return UnaryOp_INVALID
}

func (x *UnaryOp) GetArg() *Expr {
	if x != nil {
		return x.Arg
	}
	return nil
}

type BinaryOp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lhs *Expr             `protobuf:"bytes,1,opt,name=lhs,proto3" json:"lhs,omitempty"`
	Op  BinaryOp_Operator `protobuf:"varint,2,opt,name=op,proto3,enum=types.v1.BinaryOp_Operator" json:"op,omitempty"`
	Rhs *Expr             `protobuf:"bytes,3,opt,name=rhs,proto3" json:"rhs,omitempty"`
}

func (x *BinaryOp) Reset() {
	*x = BinaryOp{}
	mi := &file_types_v1_logquery_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BinaryOp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryOp) ProtoMessage() {}

func (x *BinaryOp) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_logquery_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryOp.ProtoReflect.Descriptor instead.
func (*BinaryOp) Descriptor() ([]byte, []int) {
	return file_types_v1_logquery_proto_rawDescGZIP(), []int{3}
}

func (x *BinaryOp) GetLhs() *Expr {
	if x != nil {
		return x.Lhs
	}
	return nil
}

func (x *BinaryOp) GetOp() BinaryOp_Operator {
	if x != nil {
		return x.Op
	}
	return BinaryOp_INVALID
}

func (x *BinaryOp) GetRhs() *Expr {
	if x != nil {
		return x.Rhs
	}
	return nil
}

type FuncCall struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Args []*Expr `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
}

func (x *FuncCall) Reset() {
	*x = FuncCall{}
	mi := &file_types_v1_logquery_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FuncCall) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FuncCall) ProtoMessage() {}

func (x *FuncCall) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_logquery_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FuncCall.ProtoReflect.Descriptor instead.
func (*FuncCall) Descriptor() ([]byte, []int) {
	return file_types_v1_logquery_proto_rawDescGZIP(), []int{4}
}

func (x *FuncCall) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FuncCall) GetArgs() []*Expr {
	if x != nil {
		return x.Args
	}
	return nil
}

var File_types_v1_logquery_proto protoreflect.FileDescriptor

var file_types_v1_logquery_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x08, 0x4c,
	0x6f, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x2e, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x02, 0x74, 0x6f, 0x12, 0x24, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78,
	0x70, 0x72, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0xf7, 0x01, 0x0a, 0x04, 0x45, 0x78,
	0x70, 0x72, 0x12, 0x29, 0x0a, 0x07, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56,
	0x61, 0x6c, 0x48, 0x00, 0x52, 0x07, 0x6c, 0x69, 0x74, 0x65, 0x72, 0x61, 0x6c, 0x12, 0x29, 0x0a,
	0x05, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x4f, 0x70, 0x48,
	0x00, 0x52, 0x05, 0x75, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x2c, 0x0a, 0x06, 0x62, 0x69, 0x6e, 0x61,
	0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x4f, 0x70, 0x48, 0x00, 0x52, 0x06,
	0x62, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x12, 0x31, 0x0a, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x5f, 0x63,
	0x61, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x43, 0x61, 0x6c, 0x6c, 0x48, 0x00, 0x52,
	0x08, 0x66, 0x75, 0x6e, 0x63, 0x43, 0x61, 0x6c, 0x6c, 0x12, 0x27, 0x0a, 0x04, 0x6e, 0x65, 0x78,
	0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x65, 0x78, 0x74, 0x88,
	0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x65, 0x78, 0x70, 0x72, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e,
	0x65, 0x78, 0x74, 0x22, 0x79, 0x0a, 0x07, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x4f, 0x70, 0x12, 0x2a,
	0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x4f, 0x70, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x20, 0x0a, 0x03, 0x61, 0x72,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x52, 0x03, 0x61, 0x72, 0x67, 0x22, 0x20, 0x0a, 0x08,
	0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x4f, 0x54, 0x10, 0x01, 0x22, 0xc7,
	0x02, 0x0a, 0x08, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x4f, 0x70, 0x12, 0x20, 0x0a, 0x03, 0x6c,
	0x68, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x52, 0x03, 0x6c, 0x68, 0x73, 0x12, 0x2b, 0x0a,
	0x02, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x4f, 0x70, 0x2e, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x02, 0x6f, 0x70, 0x12, 0x20, 0x0a, 0x03, 0x72, 0x68,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x52, 0x03, 0x72, 0x68, 0x73, 0x22, 0xc9, 0x01, 0x0a,
	0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x4f, 0x47, 0x5f, 0x41, 0x4e,
	0x44, 0x10, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x4f, 0x47, 0x5f, 0x4f, 0x52, 0x10, 0x66, 0x12,
	0x0c, 0x0a, 0x07, 0x4c, 0x4f, 0x47, 0x5f, 0x58, 0x4f, 0x52, 0x10, 0xc9, 0x01, 0x12, 0x0c, 0x0a,
	0x07, 0x4e, 0x55, 0x4d, 0x5f, 0x41, 0x44, 0x44, 0x10, 0xca, 0x01, 0x12, 0x0c, 0x0a, 0x07, 0x4e,
	0x55, 0x4d, 0x5f, 0x53, 0x55, 0x42, 0x10, 0xcb, 0x01, 0x12, 0x0c, 0x0a, 0x07, 0x4e, 0x55, 0x4d,
	0x5f, 0x44, 0x49, 0x56, 0x10, 0xcc, 0x01, 0x12, 0x0c, 0x0a, 0x07, 0x4e, 0x55, 0x4d, 0x5f, 0x4d,
	0x55, 0x4c, 0x10, 0xcd, 0x01, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x50, 0x5f, 0x45, 0x51, 0x10,
	0xad, 0x02, 0x12, 0x0e, 0x0a, 0x09, 0x43, 0x4d, 0x50, 0x5f, 0x4e, 0x4f, 0x54, 0x45, 0x51, 0x10,
	0xae, 0x02, 0x12, 0x0b, 0x0a, 0x06, 0x43, 0x4d, 0x50, 0x5f, 0x47, 0x54, 0x10, 0xaf, 0x02, 0x12,
	0x0c, 0x0a, 0x07, 0x43, 0x4d, 0x50, 0x5f, 0x47, 0x54, 0x45, 0x10, 0xb0, 0x02, 0x12, 0x0b, 0x0a,
	0x06, 0x43, 0x4d, 0x50, 0x5f, 0x4c, 0x54, 0x10, 0xb1, 0x02, 0x12, 0x0c, 0x0a, 0x07, 0x43, 0x4d,
	0x50, 0x5f, 0x4c, 0x54, 0x45, 0x10, 0xb2, 0x02, 0x22, 0x42, 0x0a, 0x08, 0x46, 0x75, 0x6e, 0x63,
	0x43, 0x61, 0x6c, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x78, 0x70, 0x72, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x42, 0x8d, 0x01, 0x0a,
	0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x4c,
	0x6f, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2d,
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
	file_types_v1_logquery_proto_rawDescOnce sync.Once
	file_types_v1_logquery_proto_rawDescData = file_types_v1_logquery_proto_rawDesc
)

func file_types_v1_logquery_proto_rawDescGZIP() []byte {
	file_types_v1_logquery_proto_rawDescOnce.Do(func() {
		file_types_v1_logquery_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_v1_logquery_proto_rawDescData)
	})
	return file_types_v1_logquery_proto_rawDescData
}

var file_types_v1_logquery_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_types_v1_logquery_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_types_v1_logquery_proto_goTypes = []any{
	(UnaryOp_Operator)(0),         // 0: types.v1.UnaryOp.Operator
	(BinaryOp_Operator)(0),        // 1: types.v1.BinaryOp.Operator
	(*LogQuery)(nil),              // 2: types.v1.LogQuery
	(*Expr)(nil),                  // 3: types.v1.Expr
	(*UnaryOp)(nil),               // 4: types.v1.UnaryOp
	(*BinaryOp)(nil),              // 5: types.v1.BinaryOp
	(*FuncCall)(nil),              // 6: types.v1.FuncCall
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
	(*Val)(nil),                   // 8: types.v1.Val
}
var file_types_v1_logquery_proto_depIdxs = []int32{
	7,  // 0: types.v1.LogQuery.from:type_name -> google.protobuf.Timestamp
	7,  // 1: types.v1.LogQuery.to:type_name -> google.protobuf.Timestamp
	3,  // 2: types.v1.LogQuery.query:type_name -> types.v1.Expr
	8,  // 3: types.v1.Expr.literal:type_name -> types.v1.Val
	4,  // 4: types.v1.Expr.unary:type_name -> types.v1.UnaryOp
	5,  // 5: types.v1.Expr.binary:type_name -> types.v1.BinaryOp
	6,  // 6: types.v1.Expr.func_call:type_name -> types.v1.FuncCall
	3,  // 7: types.v1.Expr.next:type_name -> types.v1.Expr
	0,  // 8: types.v1.UnaryOp.op:type_name -> types.v1.UnaryOp.Operator
	3,  // 9: types.v1.UnaryOp.arg:type_name -> types.v1.Expr
	3,  // 10: types.v1.BinaryOp.lhs:type_name -> types.v1.Expr
	1,  // 11: types.v1.BinaryOp.op:type_name -> types.v1.BinaryOp.Operator
	3,  // 12: types.v1.BinaryOp.rhs:type_name -> types.v1.Expr
	3,  // 13: types.v1.FuncCall.args:type_name -> types.v1.Expr
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_types_v1_logquery_proto_init() }
func file_types_v1_logquery_proto_init() {
	if File_types_v1_logquery_proto != nil {
		return
	}
	file_types_v1_types_proto_init()
	file_types_v1_logquery_proto_msgTypes[1].OneofWrappers = []any{
		(*Expr_Literal)(nil),
		(*Expr_Unary)(nil),
		(*Expr_Binary)(nil),
		(*Expr_FuncCall)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_types_v1_logquery_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_v1_logquery_proto_goTypes,
		DependencyIndexes: file_types_v1_logquery_proto_depIdxs,
		EnumInfos:         file_types_v1_logquery_proto_enumTypes,
		MessageInfos:      file_types_v1_logquery_proto_msgTypes,
	}.Build()
	File_types_v1_logquery_proto = out.File
	file_types_v1_logquery_proto_rawDesc = nil
	file_types_v1_logquery_proto_goTypes = nil
	file_types_v1_logquery_proto_depIdxs = nil
}
