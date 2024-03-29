// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.13.0
// source: goods/goods.proto

//proto版本选择

package goods

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

type GoodInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       int64  `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	GoodName string `protobuf:"bytes,20,opt,name=GoodName,proto3" json:"GoodName,omitempty"`
	Amount   string `protobuf:"bytes,30,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Stock    int64  `protobuf:"varint,40,opt,name=Stock,proto3" json:"Stock,omitempty"`
	Image    string `protobuf:"bytes,50,opt,name=Image,proto3" json:"Image,omitempty"`
}

func (x *GoodInfo) Reset() {
	*x = GoodInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GoodInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GoodInfo) ProtoMessage() {}

func (x *GoodInfo) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GoodInfo.ProtoReflect.Descriptor instead.
func (*GoodInfo) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{0}
}

func (x *GoodInfo) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *GoodInfo) GetGoodName() string {
	if x != nil {
		return x.GoodName
	}
	return ""
}

func (x *GoodInfo) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *GoodInfo) GetStock() int64 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *GoodInfo) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

type GetGoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetGoodRequest) Reset() {
	*x = GetGoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodRequest) ProtoMessage() {}

func (x *GetGoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodRequest.ProtoReflect.Descriptor instead.
func (*GetGoodRequest) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{1}
}

func (x *GetGoodRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetGoodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *GoodInfo `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetGoodResponse) Reset() {
	*x = GetGoodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodResponse) ProtoMessage() {}

func (x *GetGoodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodResponse.ProtoReflect.Descriptor instead.
func (*GetGoodResponse) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{2}
}

func (x *GetGoodResponse) GetInfo() *GoodInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type GetGoodsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64 `protobuf:"varint,10,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit  int64 `protobuf:"varint,20,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *GetGoodsRequest) Reset() {
	*x = GetGoodsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsRequest) ProtoMessage() {}

func (x *GetGoodsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsRequest.ProtoReflect.Descriptor instead.
func (*GetGoodsRequest) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{3}
}

func (x *GetGoodsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetGoodsRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetGoodsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Infos []*GoodInfo `protobuf:"bytes,10,rep,name=Infos,proto3" json:"Infos,omitempty"`
	Total int64       `protobuf:"varint,20,opt,name=Total,proto3" json:"Total,omitempty"`
}

func (x *GetGoodsResponse) Reset() {
	*x = GetGoodsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodsResponse) ProtoMessage() {}

func (x *GetGoodsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodsResponse.ProtoReflect.Descriptor instead.
func (*GetGoodsResponse) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{4}
}

func (x *GetGoodsResponse) GetInfos() []*GoodInfo {
	if x != nil {
		return x.Infos
	}
	return nil
}

func (x *GetGoodsResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type CreateGoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *GoodInfo `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateGoodRequest) Reset() {
	*x = CreateGoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGoodRequest) ProtoMessage() {}

func (x *CreateGoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGoodRequest.ProtoReflect.Descriptor instead.
func (*CreateGoodRequest) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{5}
}

func (x *CreateGoodRequest) GetInfo() *GoodInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateGoodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *GoodInfo `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *CreateGoodResponse) Reset() {
	*x = CreateGoodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGoodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGoodResponse) ProtoMessage() {}

func (x *CreateGoodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGoodResponse.ProtoReflect.Descriptor instead.
func (*CreateGoodResponse) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{6}
}

func (x *CreateGoodResponse) GetInfo() *GoodInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateGoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *GoodInfo `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateGoodRequest) Reset() {
	*x = UpdateGoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGoodRequest) ProtoMessage() {}

func (x *UpdateGoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGoodRequest.ProtoReflect.Descriptor instead.
func (*UpdateGoodRequest) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateGoodRequest) GetInfo() *GoodInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateGoodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *GoodInfo `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateGoodResponse) Reset() {
	*x = UpdateGoodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGoodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGoodResponse) ProtoMessage() {}

func (x *UpdateGoodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGoodResponse.ProtoReflect.Descriptor instead.
func (*UpdateGoodResponse) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateGoodResponse) GetInfo() *GoodInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type DeleteGoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *DeleteGoodRequest) Reset() {
	*x = DeleteGoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGoodRequest) ProtoMessage() {}

func (x *DeleteGoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGoodRequest.ProtoReflect.Descriptor instead.
func (*DeleteGoodRequest) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteGoodRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type DeleteGoodResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteGoodResponse) Reset() {
	*x = DeleteGoodResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGoodResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGoodResponse) ProtoMessage() {}

func (x *DeleteGoodResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGoodResponse.ProtoReflect.Descriptor instead.
func (*DeleteGoodResponse) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{10}
}

type GetGoodByGoodnameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Goodname string `protobuf:"bytes,10,opt,name=Goodname,proto3" json:"Goodname,omitempty"`
}

func (x *GetGoodByGoodnameRequest) Reset() {
	*x = GetGoodByGoodnameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodByGoodnameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodByGoodnameRequest) ProtoMessage() {}

func (x *GetGoodByGoodnameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodByGoodnameRequest.ProtoReflect.Descriptor instead.
func (*GetGoodByGoodnameRequest) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{11}
}

func (x *GetGoodByGoodnameRequest) GetGoodname() string {
	if x != nil {
		return x.Goodname
	}
	return ""
}

type GetGoodByGoodnameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *GoodInfo `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *GetGoodByGoodnameResponse) Reset() {
	*x = GetGoodByGoodnameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGoodByGoodnameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGoodByGoodnameResponse) ProtoMessage() {}

func (x *GetGoodByGoodnameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGoodByGoodnameResponse.ProtoReflect.Descriptor instead.
func (*GetGoodByGoodnameResponse) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{12}
}

func (x *GetGoodByGoodnameResponse) GetInfo() *GoodInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type UpdateStockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID  int64 `protobuf:"varint,10,opt,name=ID,proto3" json:"ID,omitempty"`
	Num int64 `protobuf:"varint,20,opt,name=Num,proto3" json:"Num,omitempty"`
}

func (x *UpdateStockRequest) Reset() {
	*x = UpdateStockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStockRequest) ProtoMessage() {}

func (x *UpdateStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStockRequest.ProtoReflect.Descriptor instead.
func (*UpdateStockRequest) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{13}
}

func (x *UpdateStockRequest) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateStockRequest) GetNum() int64 {
	if x != nil {
		return x.Num
	}
	return 0
}

type UpdateStockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Info *GoodInfo `protobuf:"bytes,10,opt,name=Info,proto3" json:"Info,omitempty"`
}

func (x *UpdateStockResponse) Reset() {
	*x = UpdateStockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_goods_goods_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateStockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStockResponse) ProtoMessage() {}

func (x *UpdateStockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goods_goods_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStockResponse.ProtoReflect.Descriptor instead.
func (*UpdateStockResponse) Descriptor() ([]byte, []int) {
	return file_goods_goods_proto_rawDescGZIP(), []int{14}
}

func (x *UpdateStockResponse) GetInfo() *GoodInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

var File_goods_goods_proto protoreflect.FileDescriptor

var file_goods_goods_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x22, 0x7a, 0x0a, 0x08, 0x47, 0x6f,
	0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x1e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74,
	0x6f, 0x63, 0x6b, 0x18, 0x28, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x12, 0x14, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x32, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x36, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47,
	0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x49,
	0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f,
	0x22, 0x3f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x22, 0x4f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f,
	0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x54, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x22, 0x38, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x39, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x38, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x47, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x39, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f,
	0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x23, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49,
	0x44, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x47, 0x6f,
	0x6f, 0x64, 0x42, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x47, 0x6f, 0x6f, 0x64, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x40, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x42, 0x79, 0x47, 0x6f, 0x6f, 0x64,
	0x6e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x04,
	0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x6f, 0x6f,
	0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0x36, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x10, 0x0a, 0x03, 0x4e, 0x75, 0x6d, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x4e, 0x75, 0x6d, 0x22, 0x3a, 0x0a, 0x13, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x23, 0x0a, 0x04, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x6f, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x49, 0x6e, 0x66, 0x6f, 0x32, 0x8c, 0x03, 0x0a, 0x04, 0x47, 0x6f, 0x6f, 0x64, 0x12, 0x38,
	0x0a, 0x07, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x12, 0x15, 0x2e, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x47,
	0x6f, 0x6f, 0x64, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74,
	0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x6f, 0x6f, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47,
	0x6f, 0x6f, 0x64, 0x12, 0x18, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x12, 0x18, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47,
	0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x12, 0x18, 0x2e, 0x67, 0x6f, 0x6f, 0x64,
	0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x47, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44,
	0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x19, 0x2e,
	0x67, 0x6f, 0x6f, 0x64, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x64, 0x73,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x22, 0x5a, 0x20, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x4a, 0x6f, 0x62, 0x4e, 0x69, 0x6e, 0x67, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_goods_goods_proto_rawDescOnce sync.Once
	file_goods_goods_proto_rawDescData = file_goods_goods_proto_rawDesc
)

func file_goods_goods_proto_rawDescGZIP() []byte {
	file_goods_goods_proto_rawDescOnce.Do(func() {
		file_goods_goods_proto_rawDescData = protoimpl.X.CompressGZIP(file_goods_goods_proto_rawDescData)
	})
	return file_goods_goods_proto_rawDescData
}

var file_goods_goods_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_goods_goods_proto_goTypes = []interface{}{
	(*GoodInfo)(nil),                  // 0: goods.GoodInfo
	(*GetGoodRequest)(nil),            // 1: goods.GetGoodRequest
	(*GetGoodResponse)(nil),           // 2: goods.GetGoodResponse
	(*GetGoodsRequest)(nil),           // 3: goods.GetGoodsRequest
	(*GetGoodsResponse)(nil),          // 4: goods.GetGoodsResponse
	(*CreateGoodRequest)(nil),         // 5: goods.CreateGoodRequest
	(*CreateGoodResponse)(nil),        // 6: goods.CreateGoodResponse
	(*UpdateGoodRequest)(nil),         // 7: goods.UpdateGoodRequest
	(*UpdateGoodResponse)(nil),        // 8: goods.UpdateGoodResponse
	(*DeleteGoodRequest)(nil),         // 9: goods.DeleteGoodRequest
	(*DeleteGoodResponse)(nil),        // 10: goods.DeleteGoodResponse
	(*GetGoodByGoodnameRequest)(nil),  // 11: goods.GetGoodByGoodnameRequest
	(*GetGoodByGoodnameResponse)(nil), // 12: goods.GetGoodByGoodnameResponse
	(*UpdateStockRequest)(nil),        // 13: goods.UpdateStockRequest
	(*UpdateStockResponse)(nil),       // 14: goods.UpdateStockResponse
}
var file_goods_goods_proto_depIdxs = []int32{
	0,  // 0: goods.GetGoodResponse.Info:type_name -> goods.GoodInfo
	0,  // 1: goods.GetGoodsResponse.Infos:type_name -> goods.GoodInfo
	0,  // 2: goods.CreateGoodRequest.Info:type_name -> goods.GoodInfo
	0,  // 3: goods.CreateGoodResponse.Info:type_name -> goods.GoodInfo
	0,  // 4: goods.UpdateGoodRequest.Info:type_name -> goods.GoodInfo
	0,  // 5: goods.UpdateGoodResponse.Info:type_name -> goods.GoodInfo
	0,  // 6: goods.GetGoodByGoodnameResponse.Info:type_name -> goods.GoodInfo
	0,  // 7: goods.UpdateStockResponse.Info:type_name -> goods.GoodInfo
	1,  // 8: goods.Good.GetGood:input_type -> goods.GetGoodRequest
	3,  // 9: goods.Good.GetGoods:input_type -> goods.GetGoodsRequest
	5,  // 10: goods.Good.CreateGood:input_type -> goods.CreateGoodRequest
	7,  // 11: goods.Good.UpdateGood:input_type -> goods.UpdateGoodRequest
	9,  // 12: goods.Good.DeleteGood:input_type -> goods.DeleteGoodRequest
	13, // 13: goods.Good.UpdateStock:input_type -> goods.UpdateStockRequest
	2,  // 14: goods.Good.GetGood:output_type -> goods.GetGoodResponse
	4,  // 15: goods.Good.GetGoods:output_type -> goods.GetGoodsResponse
	6,  // 16: goods.Good.CreateGood:output_type -> goods.CreateGoodResponse
	8,  // 17: goods.Good.UpdateGood:output_type -> goods.UpdateGoodResponse
	10, // 18: goods.Good.DeleteGood:output_type -> goods.DeleteGoodResponse
	14, // 19: goods.Good.UpdateStock:output_type -> goods.UpdateStockResponse
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_goods_goods_proto_init() }
func file_goods_goods_proto_init() {
	if File_goods_goods_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_goods_goods_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GoodInfo); i {
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
		file_goods_goods_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodRequest); i {
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
		file_goods_goods_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodResponse); i {
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
		file_goods_goods_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsRequest); i {
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
		file_goods_goods_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodsResponse); i {
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
		file_goods_goods_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGoodRequest); i {
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
		file_goods_goods_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGoodResponse); i {
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
		file_goods_goods_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGoodRequest); i {
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
		file_goods_goods_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGoodResponse); i {
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
		file_goods_goods_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGoodRequest); i {
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
		file_goods_goods_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGoodResponse); i {
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
		file_goods_goods_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodByGoodnameRequest); i {
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
		file_goods_goods_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGoodByGoodnameResponse); i {
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
		file_goods_goods_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStockRequest); i {
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
		file_goods_goods_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateStockResponse); i {
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
			RawDescriptor: file_goods_goods_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goods_goods_proto_goTypes,
		DependencyIndexes: file_goods_goods_proto_depIdxs,
		MessageInfos:      file_goods_goods_proto_msgTypes,
	}.Build()
	File_goods_goods_proto = out.File
	file_goods_goods_proto_rawDesc = nil
	file_goods_goods_proto_goTypes = nil
	file_goods_goods_proto_depIdxs = nil
}
