// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        (unknown)
// source: types/v1/payment_method.proto

package typesv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PaymentMethod_Card_Brand int32

const (
	PaymentMethod_Card_unknown    PaymentMethod_Card_Brand = 0
	PaymentMethod_Card_amex       PaymentMethod_Card_Brand = 1
	PaymentMethod_Card_diners     PaymentMethod_Card_Brand = 2
	PaymentMethod_Card_discover   PaymentMethod_Card_Brand = 3
	PaymentMethod_Card_jcb        PaymentMethod_Card_Brand = 4
	PaymentMethod_Card_mastercard PaymentMethod_Card_Brand = 5
	PaymentMethod_Card_unionpay   PaymentMethod_Card_Brand = 6
	PaymentMethod_Card_visa       PaymentMethod_Card_Brand = 7
)

// Enum value maps for PaymentMethod_Card_Brand.
var (
	PaymentMethod_Card_Brand_name = map[int32]string{
		0: "unknown",
		1: "amex",
		2: "diners",
		3: "discover",
		4: "jcb",
		5: "mastercard",
		6: "unionpay",
		7: "visa",
	}
	PaymentMethod_Card_Brand_value = map[string]int32{
		"unknown":    0,
		"amex":       1,
		"diners":     2,
		"discover":   3,
		"jcb":        4,
		"mastercard": 5,
		"unionpay":   6,
		"visa":       7,
	}
)

func (x PaymentMethod_Card_Brand) Enum() *PaymentMethod_Card_Brand {
	p := new(PaymentMethod_Card_Brand)
	*p = x
	return p
}

func (x PaymentMethod_Card_Brand) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PaymentMethod_Card_Brand) Descriptor() protoreflect.EnumDescriptor {
	return file_types_v1_payment_method_proto_enumTypes[0].Descriptor()
}

func (PaymentMethod_Card_Brand) Type() protoreflect.EnumType {
	return &file_types_v1_payment_method_proto_enumTypes[0]
}

func (x PaymentMethod_Card_Brand) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PaymentMethod_Card_Brand.Descriptor instead.
func (PaymentMethod_Card_Brand) EnumDescriptor() ([]byte, []int) {
	return file_types_v1_payment_method_proto_rawDescGZIP(), []int{1, 0, 0}
}

type Address struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	City          string                 `protobuf:"bytes,1,opt,name=city,proto3" json:"city,omitempty"`
	Country       string                 `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	Line1         string                 `protobuf:"bytes,3,opt,name=line1,proto3" json:"line1,omitempty"`
	Line2         string                 `protobuf:"bytes,4,opt,name=line2,proto3" json:"line2,omitempty"`
	PostalCode    string                 `protobuf:"bytes,5,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	State         string                 `protobuf:"bytes,6,opt,name=state,proto3" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Address) Reset() {
	*x = Address{}
	mi := &file_types_v1_payment_method_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_payment_method_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_types_v1_payment_method_proto_rawDescGZIP(), []int{0}
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetLine1() string {
	if x != nil {
		return x.Line1
	}
	return ""
}

func (x *Address) GetLine2() string {
	if x != nil {
		return x.Line2
	}
	return ""
}

func (x *Address) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type PaymentMethod struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,100,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	StripeId  string                 `protobuf:"bytes,1,opt,name=stripe_id,json=stripeId,proto3" json:"stripe_id,omitempty"`
	// Types that are valid to be assigned to Type:
	//
	//	*PaymentMethod_Card_
	Type           isPaymentMethod_Type          `protobuf_oneof:"type"`
	BillingDetails *PaymentMethod_BillingDetails `protobuf:"bytes,3,opt,name=billing_details,json=billingDetails,proto3" json:"billing_details,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *PaymentMethod) Reset() {
	*x = PaymentMethod{}
	mi := &file_types_v1_payment_method_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentMethod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentMethod) ProtoMessage() {}

func (x *PaymentMethod) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_payment_method_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentMethod.ProtoReflect.Descriptor instead.
func (*PaymentMethod) Descriptor() ([]byte, []int) {
	return file_types_v1_payment_method_proto_rawDescGZIP(), []int{1}
}

func (x *PaymentMethod) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *PaymentMethod) GetStripeId() string {
	if x != nil {
		return x.StripeId
	}
	return ""
}

func (x *PaymentMethod) GetType() isPaymentMethod_Type {
	if x != nil {
		return x.Type
	}
	return nil
}

func (x *PaymentMethod) GetCard() *PaymentMethod_Card {
	if x != nil {
		if x, ok := x.Type.(*PaymentMethod_Card_); ok {
			return x.Card
		}
	}
	return nil
}

func (x *PaymentMethod) GetBillingDetails() *PaymentMethod_BillingDetails {
	if x != nil {
		return x.BillingDetails
	}
	return nil
}

type isPaymentMethod_Type interface {
	isPaymentMethod_Type()
}

type PaymentMethod_Card_ struct {
	Card *PaymentMethod_Card `protobuf:"bytes,200,opt,name=card,proto3,oneof"`
}

func (*PaymentMethod_Card_) isPaymentMethod_Type() {}

type PaymentMethod_Card struct {
	state         protoimpl.MessageState   `protogen:"open.v1"`
	Brand         PaymentMethod_Card_Brand `protobuf:"varint,1,opt,name=brand,proto3,enum=types.v1.PaymentMethod_Card_Brand" json:"brand,omitempty"`
	Country       string                   `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	DisplayBrandh string                   `protobuf:"bytes,3,opt,name=display_brandh,json=displayBrandh,proto3" json:"display_brandh,omitempty"`
	ExpMonth      int64                    `protobuf:"varint,4,opt,name=exp_month,json=expMonth,proto3" json:"exp_month,omitempty"`
	ExpYear       int64                    `protobuf:"varint,5,opt,name=exp_year,json=expYear,proto3" json:"exp_year,omitempty"`
	Last4         string                   `protobuf:"bytes,6,opt,name=last4,proto3" json:"last4,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaymentMethod_Card) Reset() {
	*x = PaymentMethod_Card{}
	mi := &file_types_v1_payment_method_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentMethod_Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentMethod_Card) ProtoMessage() {}

func (x *PaymentMethod_Card) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_payment_method_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentMethod_Card.ProtoReflect.Descriptor instead.
func (*PaymentMethod_Card) Descriptor() ([]byte, []int) {
	return file_types_v1_payment_method_proto_rawDescGZIP(), []int{1, 0}
}

func (x *PaymentMethod_Card) GetBrand() PaymentMethod_Card_Brand {
	if x != nil {
		return x.Brand
	}
	return PaymentMethod_Card_unknown
}

func (x *PaymentMethod_Card) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *PaymentMethod_Card) GetDisplayBrandh() string {
	if x != nil {
		return x.DisplayBrandh
	}
	return ""
}

func (x *PaymentMethod_Card) GetExpMonth() int64 {
	if x != nil {
		return x.ExpMonth
	}
	return 0
}

func (x *PaymentMethod_Card) GetExpYear() int64 {
	if x != nil {
		return x.ExpYear
	}
	return 0
}

func (x *PaymentMethod_Card) GetLast4() string {
	if x != nil {
		return x.Last4
	}
	return ""
}

type PaymentMethod_BillingDetails struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Address       *Address               `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name          string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Phone         string                 `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaymentMethod_BillingDetails) Reset() {
	*x = PaymentMethod_BillingDetails{}
	mi := &file_types_v1_payment_method_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaymentMethod_BillingDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentMethod_BillingDetails) ProtoMessage() {}

func (x *PaymentMethod_BillingDetails) ProtoReflect() protoreflect.Message {
	mi := &file_types_v1_payment_method_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentMethod_BillingDetails.ProtoReflect.Descriptor instead.
func (*PaymentMethod_BillingDetails) Descriptor() ([]byte, []int) {
	return file_types_v1_payment_method_proto_rawDescGZIP(), []int{1, 1}
}

func (x *PaymentMethod_BillingDetails) GetAddress() *Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *PaymentMethod_BillingDetails) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *PaymentMethod_BillingDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PaymentMethod_BillingDetails) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

var File_types_v1_payment_method_proto protoreflect.FileDescriptor

var file_types_v1_payment_method_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9a, 0x01, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x31, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x31, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6e, 0x65, 0x32, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x32,
	0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0xb1, 0x05, 0x0a, 0x0d, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x49,
	0x64, 0x12, 0x33, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x64, 0x18, 0xc8, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x48, 0x00,
	0x52, 0x04, 0x63, 0x61, 0x72, 0x64, 0x12, 0x4f, 0x0a, 0x0f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e,
	0x67, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x0e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0xba, 0x02, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64,
	0x12, 0x38, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x22, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x2e, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f,
	0x62, 0x72, 0x61, 0x6e, 0x64, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x65,
	0x78, 0x70, 0x5f, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x65, 0x78, 0x70, 0x4d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x5f,
	0x79, 0x65, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x78, 0x70, 0x59,
	0x65, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x73, 0x74, 0x34, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x73, 0x74, 0x34, 0x22, 0x69, 0x0a, 0x05, 0x42, 0x72, 0x61,
	0x6e, 0x64, 0x12, 0x0b, 0x0a, 0x07, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x12,
	0x08, 0x0a, 0x04, 0x61, 0x6d, 0x65, 0x78, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x64, 0x69, 0x6e,
	0x65, 0x72, 0x73, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03, 0x6a, 0x63, 0x62, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x63, 0x61, 0x72, 0x64, 0x10, 0x05, 0x12, 0x0c, 0x0a, 0x08,
	0x75, 0x6e, 0x69, 0x6f, 0x6e, 0x70, 0x61, 0x79, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x76, 0x69,
	0x73, 0x61, 0x10, 0x07, 0x1a, 0x7d, 0x0a, 0x0e, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x2b, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x92, 0x01, 0x0a, 0x0c,
	0x63, 0x6f, 0x6d, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68,
	0x75, 0x6d, 0x61, 0x6e, 0x6c, 0x6f, 0x67, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x76,
	0x31, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x08, 0x54, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x56, 0x31, 0xca, 0x02, 0x08, 0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x54, 0x79, 0x70, 0x65, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_types_v1_payment_method_proto_rawDescOnce sync.Once
	file_types_v1_payment_method_proto_rawDescData []byte
)

func file_types_v1_payment_method_proto_rawDescGZIP() []byte {
	file_types_v1_payment_method_proto_rawDescOnce.Do(func() {
		file_types_v1_payment_method_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_types_v1_payment_method_proto_rawDesc), len(file_types_v1_payment_method_proto_rawDesc)))
	})
	return file_types_v1_payment_method_proto_rawDescData
}

var file_types_v1_payment_method_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_types_v1_payment_method_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_types_v1_payment_method_proto_goTypes = []any{
	(PaymentMethod_Card_Brand)(0),        // 0: types.v1.PaymentMethod.Card.Brand
	(*Address)(nil),                      // 1: types.v1.Address
	(*PaymentMethod)(nil),                // 2: types.v1.PaymentMethod
	(*PaymentMethod_Card)(nil),           // 3: types.v1.PaymentMethod.Card
	(*PaymentMethod_BillingDetails)(nil), // 4: types.v1.PaymentMethod.BillingDetails
	(*timestamppb.Timestamp)(nil),        // 5: google.protobuf.Timestamp
}
var file_types_v1_payment_method_proto_depIdxs = []int32{
	5, // 0: types.v1.PaymentMethod.created_at:type_name -> google.protobuf.Timestamp
	3, // 1: types.v1.PaymentMethod.card:type_name -> types.v1.PaymentMethod.Card
	4, // 2: types.v1.PaymentMethod.billing_details:type_name -> types.v1.PaymentMethod.BillingDetails
	0, // 3: types.v1.PaymentMethod.Card.brand:type_name -> types.v1.PaymentMethod.Card.Brand
	1, // 4: types.v1.PaymentMethod.BillingDetails.address:type_name -> types.v1.Address
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_types_v1_payment_method_proto_init() }
func file_types_v1_payment_method_proto_init() {
	if File_types_v1_payment_method_proto != nil {
		return
	}
	file_types_v1_payment_method_proto_msgTypes[1].OneofWrappers = []any{
		(*PaymentMethod_Card_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_types_v1_payment_method_proto_rawDesc), len(file_types_v1_payment_method_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_types_v1_payment_method_proto_goTypes,
		DependencyIndexes: file_types_v1_payment_method_proto_depIdxs,
		EnumInfos:         file_types_v1_payment_method_proto_enumTypes,
		MessageInfos:      file_types_v1_payment_method_proto_msgTypes,
	}.Build()
	File_types_v1_payment_method_proto = out.File
	file_types_v1_payment_method_proto_goTypes = nil
	file_types_v1_payment_method_proto_depIdxs = nil
}
