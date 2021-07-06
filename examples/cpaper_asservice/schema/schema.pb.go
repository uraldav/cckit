// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: cpaper_asservice/schema/schema.proto

package schema

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/mwitkow/go-proto-validators"
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

type CommercialPaper_State int32

const (
	CommercialPaper_ISSUED   CommercialPaper_State = 0
	CommercialPaper_TRADING  CommercialPaper_State = 1
	CommercialPaper_REDEEMED CommercialPaper_State = 2
)

// Enum value maps for CommercialPaper_State.
var (
	CommercialPaper_State_name = map[int32]string{
		0: "ISSUED",
		1: "TRADING",
		2: "REDEEMED",
	}
	CommercialPaper_State_value = map[string]int32{
		"ISSUED":   0,
		"TRADING":  1,
		"REDEEMED": 2,
	}
)

func (x CommercialPaper_State) Enum() *CommercialPaper_State {
	p := new(CommercialPaper_State)
	*p = x
	return p
}

func (x CommercialPaper_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommercialPaper_State) Descriptor() protoreflect.EnumDescriptor {
	return file_cpaper_asservice_schema_schema_proto_enumTypes[0].Descriptor()
}

func (CommercialPaper_State) Type() protoreflect.EnumType {
	return &file_cpaper_asservice_schema_schema_proto_enumTypes[0]
}

func (x CommercialPaper_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommercialPaper_State.Descriptor instead.
func (CommercialPaper_State) EnumDescriptor() ([]byte, []int) {
	return file_cpaper_asservice_schema_schema_proto_rawDescGZIP(), []int{0, 0}
}

// Commercial Paper state entry
type CommercialPaper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Issuer and Paper number comprises composite primary key of Commercial paper entry
	Issuer       string                `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	PaperNumber  string                `protobuf:"bytes,2,opt,name=paper_number,json=paperNumber,proto3" json:"paper_number,omitempty"`
	Owner        string                `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	IssueDate    *timestamp.Timestamp  `protobuf:"bytes,4,opt,name=issue_date,json=issueDate,proto3" json:"issue_date,omitempty"`
	MaturityDate *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=maturity_date,json=maturityDate,proto3" json:"maturity_date,omitempty"`
	FaceValue    int32                 `protobuf:"varint,6,opt,name=face_value,json=faceValue,proto3" json:"face_value,omitempty"`
	State        CommercialPaper_State `protobuf:"varint,7,opt,name=state,proto3,enum=cpaper_asservice.schema.CommercialPaper_State" json:"state,omitempty"`
	// Additional unique field for entry
	ExternalId string `protobuf:"bytes,8,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
}

func (x *CommercialPaper) Reset() {
	*x = CommercialPaper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cpaper_asservice_schema_schema_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommercialPaper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommercialPaper) ProtoMessage() {}

func (x *CommercialPaper) ProtoReflect() protoreflect.Message {
	mi := &file_cpaper_asservice_schema_schema_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommercialPaper.ProtoReflect.Descriptor instead.
func (*CommercialPaper) Descriptor() ([]byte, []int) {
	return file_cpaper_asservice_schema_schema_proto_rawDescGZIP(), []int{0}
}

func (x *CommercialPaper) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *CommercialPaper) GetPaperNumber() string {
	if x != nil {
		return x.PaperNumber
	}
	return ""
}

func (x *CommercialPaper) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *CommercialPaper) GetIssueDate() *timestamp.Timestamp {
	if x != nil {
		return x.IssueDate
	}
	return nil
}

func (x *CommercialPaper) GetMaturityDate() *timestamp.Timestamp {
	if x != nil {
		return x.MaturityDate
	}
	return nil
}

func (x *CommercialPaper) GetFaceValue() int32 {
	if x != nil {
		return x.FaceValue
	}
	return 0
}

func (x *CommercialPaper) GetState() CommercialPaper_State {
	if x != nil {
		return x.State
	}
	return CommercialPaper_ISSUED
}

func (x *CommercialPaper) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

// CommercialPaperId identifier part
type CommercialPaperId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Issuer      string `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	PaperNumber string `protobuf:"bytes,2,opt,name=paper_number,json=paperNumber,proto3" json:"paper_number,omitempty"`
}

func (x *CommercialPaperId) Reset() {
	*x = CommercialPaperId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cpaper_asservice_schema_schema_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommercialPaperId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommercialPaperId) ProtoMessage() {}

func (x *CommercialPaperId) ProtoReflect() protoreflect.Message {
	mi := &file_cpaper_asservice_schema_schema_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommercialPaperId.ProtoReflect.Descriptor instead.
func (*CommercialPaperId) Descriptor() ([]byte, []int) {
	return file_cpaper_asservice_schema_schema_proto_rawDescGZIP(), []int{1}
}

func (x *CommercialPaperId) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *CommercialPaperId) GetPaperNumber() string {
	if x != nil {
		return x.PaperNumber
	}
	return ""
}

// ExternalId
type ExternalId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ExternalId) Reset() {
	*x = ExternalId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cpaper_asservice_schema_schema_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalId) ProtoMessage() {}

func (x *ExternalId) ProtoReflect() protoreflect.Message {
	mi := &file_cpaper_asservice_schema_schema_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalId.ProtoReflect.Descriptor instead.
func (*ExternalId) Descriptor() ([]byte, []int) {
	return file_cpaper_asservice_schema_schema_proto_rawDescGZIP(), []int{2}
}

func (x *ExternalId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Container for returning multiple entities
type CommercialPaperList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*CommercialPaper `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *CommercialPaperList) Reset() {
	*x = CommercialPaperList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cpaper_asservice_schema_schema_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommercialPaperList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommercialPaperList) ProtoMessage() {}

func (x *CommercialPaperList) ProtoReflect() protoreflect.Message {
	mi := &file_cpaper_asservice_schema_schema_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommercialPaperList.ProtoReflect.Descriptor instead.
func (*CommercialPaperList) Descriptor() ([]byte, []int) {
	return file_cpaper_asservice_schema_schema_proto_rawDescGZIP(), []int{3}
}

func (x *CommercialPaperList) GetItems() []*CommercialPaper {
	if x != nil {
		return x.Items
	}
	return nil
}

// IssueCommercialPaper event
type IssueCommercialPaper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Issuer       string               `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	PaperNumber  string               `protobuf:"bytes,2,opt,name=paper_number,json=paperNumber,proto3" json:"paper_number,omitempty"`
	IssueDate    *timestamp.Timestamp `protobuf:"bytes,3,opt,name=issue_date,json=issueDate,proto3" json:"issue_date,omitempty"`
	MaturityDate *timestamp.Timestamp `protobuf:"bytes,4,opt,name=maturity_date,json=maturityDate,proto3" json:"maturity_date,omitempty"`
	FaceValue    int32                `protobuf:"varint,5,opt,name=face_value,json=faceValue,proto3" json:"face_value,omitempty"`
	// external_id  - once more uniq id of state entry
	ExternalId string `protobuf:"bytes,6,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
}

func (x *IssueCommercialPaper) Reset() {
	*x = IssueCommercialPaper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cpaper_asservice_schema_schema_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueCommercialPaper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueCommercialPaper) ProtoMessage() {}

func (x *IssueCommercialPaper) ProtoReflect() protoreflect.Message {
	mi := &file_cpaper_asservice_schema_schema_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueCommercialPaper.ProtoReflect.Descriptor instead.
func (*IssueCommercialPaper) Descriptor() ([]byte, []int) {
	return file_cpaper_asservice_schema_schema_proto_rawDescGZIP(), []int{4}
}

func (x *IssueCommercialPaper) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *IssueCommercialPaper) GetPaperNumber() string {
	if x != nil {
		return x.PaperNumber
	}
	return ""
}

func (x *IssueCommercialPaper) GetIssueDate() *timestamp.Timestamp {
	if x != nil {
		return x.IssueDate
	}
	return nil
}

func (x *IssueCommercialPaper) GetMaturityDate() *timestamp.Timestamp {
	if x != nil {
		return x.MaturityDate
	}
	return nil
}

func (x *IssueCommercialPaper) GetFaceValue() int32 {
	if x != nil {
		return x.FaceValue
	}
	return 0
}

func (x *IssueCommercialPaper) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

// BuyCommercialPaper event
type BuyCommercialPaper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Issuer       string               `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	PaperNumber  string               `protobuf:"bytes,2,opt,name=paper_number,json=paperNumber,proto3" json:"paper_number,omitempty"`
	CurrentOwner string               `protobuf:"bytes,3,opt,name=current_owner,json=currentOwner,proto3" json:"current_owner,omitempty"`
	NewOwner     string               `protobuf:"bytes,4,opt,name=new_owner,json=newOwner,proto3" json:"new_owner,omitempty"`
	Price        int32                `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	PurchaseDate *timestamp.Timestamp `protobuf:"bytes,6,opt,name=purchase_date,json=purchaseDate,proto3" json:"purchase_date,omitempty"`
}

func (x *BuyCommercialPaper) Reset() {
	*x = BuyCommercialPaper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cpaper_asservice_schema_schema_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyCommercialPaper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyCommercialPaper) ProtoMessage() {}

func (x *BuyCommercialPaper) ProtoReflect() protoreflect.Message {
	mi := &file_cpaper_asservice_schema_schema_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyCommercialPaper.ProtoReflect.Descriptor instead.
func (*BuyCommercialPaper) Descriptor() ([]byte, []int) {
	return file_cpaper_asservice_schema_schema_proto_rawDescGZIP(), []int{5}
}

func (x *BuyCommercialPaper) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *BuyCommercialPaper) GetPaperNumber() string {
	if x != nil {
		return x.PaperNumber
	}
	return ""
}

func (x *BuyCommercialPaper) GetCurrentOwner() string {
	if x != nil {
		return x.CurrentOwner
	}
	return ""
}

func (x *BuyCommercialPaper) GetNewOwner() string {
	if x != nil {
		return x.NewOwner
	}
	return ""
}

func (x *BuyCommercialPaper) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *BuyCommercialPaper) GetPurchaseDate() *timestamp.Timestamp {
	if x != nil {
		return x.PurchaseDate
	}
	return nil
}

// RedeemCommercialPaper event
type RedeemCommercialPaper struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Issuer         string               `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	PaperNumber    string               `protobuf:"bytes,2,opt,name=paper_number,json=paperNumber,proto3" json:"paper_number,omitempty"`
	RedeemingOwner string               `protobuf:"bytes,3,opt,name=redeeming_owner,json=redeemingOwner,proto3" json:"redeeming_owner,omitempty"`
	RedeemDate     *timestamp.Timestamp `protobuf:"bytes,4,opt,name=redeem_date,json=redeemDate,proto3" json:"redeem_date,omitempty"`
}

func (x *RedeemCommercialPaper) Reset() {
	*x = RedeemCommercialPaper{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cpaper_asservice_schema_schema_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedeemCommercialPaper) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedeemCommercialPaper) ProtoMessage() {}

func (x *RedeemCommercialPaper) ProtoReflect() protoreflect.Message {
	mi := &file_cpaper_asservice_schema_schema_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedeemCommercialPaper.ProtoReflect.Descriptor instead.
func (*RedeemCommercialPaper) Descriptor() ([]byte, []int) {
	return file_cpaper_asservice_schema_schema_proto_rawDescGZIP(), []int{6}
}

func (x *RedeemCommercialPaper) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *RedeemCommercialPaper) GetPaperNumber() string {
	if x != nil {
		return x.PaperNumber
	}
	return ""
}

func (x *RedeemCommercialPaper) GetRedeemingOwner() string {
	if x != nil {
		return x.RedeemingOwner
	}
	return ""
}

func (x *RedeemCommercialPaper) GetRedeemDate() *timestamp.Timestamp {
	if x != nil {
		return x.RedeemDate
	}
	return nil
}

var File_cpaper_asservice_schema_schema_proto protoreflect.FileDescriptor

var file_cpaper_asservice_schema_schema_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x63, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x61,
	0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x1a,
	0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69,
	0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x94, 0x03, 0x0a, 0x0f, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x72, 0x63, 0x69, 0x61, 0x6c, 0x50, 0x61, 0x70, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x70, 0x65,
	0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x39, 0x0a,
	0x0a, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3f, 0x0a, 0x0d, 0x6d, 0x61, 0x74, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6d, 0x61, 0x74,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x61, 0x63,
	0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x66,
	0x61, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x44, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x63, 0x70, 0x61, 0x70, 0x65, 0x72,
	0x5f, 0x61, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x69, 0x61, 0x6c, 0x50, 0x61, 0x70, 0x65,
	0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x22,
	0x2e, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x49, 0x53, 0x53, 0x55,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x54, 0x52, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x44, 0x45, 0x45, 0x4d, 0x45, 0x44, 0x10, 0x02, 0x22,
	0x4e, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x69, 0x61, 0x6c, 0x50, 0x61, 0x70,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x70, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22,
	0x1c, 0x0a, 0x0a, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x49, 0x64, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x55, 0x0a,
	0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x72, 0x63, 0x69, 0x61, 0x6c, 0x50, 0x61, 0x70, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x61, 0x73, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x72, 0x63, 0x69, 0x61, 0x6c, 0x50, 0x61, 0x70, 0x65, 0x72, 0x52, 0x05, 0x69,
	0x74, 0x65, 0x6d, 0x73, 0x22, 0xbd, 0x02, 0x0a, 0x14, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x72, 0x63, 0x69, 0x61, 0x6c, 0x50, 0x61, 0x70, 0x65, 0x72, 0x12, 0x1e, 0x0a,
	0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2,
	0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x29, 0x0a,
	0x0c, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x0b, 0x70, 0x61, 0x70,
	0x65, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x0a, 0x69, 0x73, 0x73, 0x75,
	0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01,
	0x52, 0x09, 0x69, 0x73, 0x73, 0x75, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x47, 0x0a, 0x0d, 0x6d,
	0x61, 0x74, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06,
	0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x0c, 0x6d, 0x61, 0x74, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x44, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x0a, 0x66, 0x61, 0x63, 0x65, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x10, 0x00,
	0x52, 0x09, 0x66, 0x61, 0x63, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x27, 0x0a, 0x0b, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x0a, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x49, 0x64, 0x22, 0x98, 0x02, 0x0a, 0x12, 0x42, 0x75, 0x79, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x72, 0x63, 0x69, 0x61, 0x6c, 0x50, 0x61, 0x70, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x06, 0x69,
	0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f,
	0x02, 0x58, 0x01, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x0c, 0x70,
	0x61, 0x70, 0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x0b, 0x70, 0x61, 0x70, 0x65, 0x72,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2b, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2,
	0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x77,
	0x6e, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x09, 0x6e, 0x65, 0x77, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x08,
	0x6e, 0x65, 0x77, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x10, 0x00, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0d, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20,
	0x01, 0x52, 0x0c, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22,
	0xd8, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x72,
	0x63, 0x69, 0x61, 0x6c, 0x50, 0x61, 0x70, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x06, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58,
	0x01, 0x52, 0x06, 0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x0c, 0x70, 0x61, 0x70,
	0x65, 0x72, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x0b, 0x70, 0x61, 0x70, 0x65, 0x72, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x0f, 0x72, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x69, 0x6e,
	0x67, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2,
	0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x0e, 0x72, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x69, 0x6e, 0x67,
	0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x43, 0x0a, 0x0b, 0x72, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x20, 0x01, 0x52, 0x0a,
	0x72, 0x65, 0x64, 0x65, 0x65, 0x6d, 0x44, 0x61, 0x74, 0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x37, 0x74, 0x65, 0x63, 0x68, 0x6c,
	0x61, 0x62, 0x2f, 0x63, 0x63, 0x6b, 0x69, 0x74, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x73, 0x2f, 0x63, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_cpaper_asservice_schema_schema_proto_rawDescOnce sync.Once
	file_cpaper_asservice_schema_schema_proto_rawDescData = file_cpaper_asservice_schema_schema_proto_rawDesc
)

func file_cpaper_asservice_schema_schema_proto_rawDescGZIP() []byte {
	file_cpaper_asservice_schema_schema_proto_rawDescOnce.Do(func() {
		file_cpaper_asservice_schema_schema_proto_rawDescData = protoimpl.X.CompressGZIP(file_cpaper_asservice_schema_schema_proto_rawDescData)
	})
	return file_cpaper_asservice_schema_schema_proto_rawDescData
}

var file_cpaper_asservice_schema_schema_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_cpaper_asservice_schema_schema_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cpaper_asservice_schema_schema_proto_goTypes = []interface{}{
	(CommercialPaper_State)(0),    // 0: cpaper_asservice.schema.CommercialPaper.State
	(*CommercialPaper)(nil),       // 1: cpaper_asservice.schema.CommercialPaper
	(*CommercialPaperId)(nil),     // 2: cpaper_asservice.schema.CommercialPaperId
	(*ExternalId)(nil),            // 3: cpaper_asservice.schema.ExternalId
	(*CommercialPaperList)(nil),   // 4: cpaper_asservice.schema.CommercialPaperList
	(*IssueCommercialPaper)(nil),  // 5: cpaper_asservice.schema.IssueCommercialPaper
	(*BuyCommercialPaper)(nil),    // 6: cpaper_asservice.schema.BuyCommercialPaper
	(*RedeemCommercialPaper)(nil), // 7: cpaper_asservice.schema.RedeemCommercialPaper
	(*timestamp.Timestamp)(nil),   // 8: google.protobuf.Timestamp
}
var file_cpaper_asservice_schema_schema_proto_depIdxs = []int32{
	8, // 0: cpaper_asservice.schema.CommercialPaper.issue_date:type_name -> google.protobuf.Timestamp
	8, // 1: cpaper_asservice.schema.CommercialPaper.maturity_date:type_name -> google.protobuf.Timestamp
	0, // 2: cpaper_asservice.schema.CommercialPaper.state:type_name -> cpaper_asservice.schema.CommercialPaper.State
	1, // 3: cpaper_asservice.schema.CommercialPaperList.items:type_name -> cpaper_asservice.schema.CommercialPaper
	8, // 4: cpaper_asservice.schema.IssueCommercialPaper.issue_date:type_name -> google.protobuf.Timestamp
	8, // 5: cpaper_asservice.schema.IssueCommercialPaper.maturity_date:type_name -> google.protobuf.Timestamp
	8, // 6: cpaper_asservice.schema.BuyCommercialPaper.purchase_date:type_name -> google.protobuf.Timestamp
	8, // 7: cpaper_asservice.schema.RedeemCommercialPaper.redeem_date:type_name -> google.protobuf.Timestamp
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_cpaper_asservice_schema_schema_proto_init() }
func file_cpaper_asservice_schema_schema_proto_init() {
	if File_cpaper_asservice_schema_schema_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cpaper_asservice_schema_schema_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommercialPaper); i {
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
		file_cpaper_asservice_schema_schema_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommercialPaperId); i {
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
		file_cpaper_asservice_schema_schema_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalId); i {
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
		file_cpaper_asservice_schema_schema_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommercialPaperList); i {
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
		file_cpaper_asservice_schema_schema_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueCommercialPaper); i {
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
		file_cpaper_asservice_schema_schema_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyCommercialPaper); i {
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
		file_cpaper_asservice_schema_schema_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedeemCommercialPaper); i {
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
			RawDescriptor: file_cpaper_asservice_schema_schema_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cpaper_asservice_schema_schema_proto_goTypes,
		DependencyIndexes: file_cpaper_asservice_schema_schema_proto_depIdxs,
		EnumInfos:         file_cpaper_asservice_schema_schema_proto_enumTypes,
		MessageInfos:      file_cpaper_asservice_schema_schema_proto_msgTypes,
	}.Build()
	File_cpaper_asservice_schema_schema_proto = out.File
	file_cpaper_asservice_schema_schema_proto_rawDesc = nil
	file_cpaper_asservice_schema_schema_proto_goTypes = nil
	file_cpaper_asservice_schema_schema_proto_depIdxs = nil
}
