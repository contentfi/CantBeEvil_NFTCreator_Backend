// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: api.proto

package api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type GetCollectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCollectionRequest) Reset() {
	*x = GetCollectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCollectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCollectionRequest) ProtoMessage() {}

func (x *GetCollectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCollectionRequest.ProtoReflect.Descriptor instead.
func (*GetCollectionRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *GetCollectionRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListCollectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AnchorId int64 `protobuf:"varint,1,opt,name=anchor_id,json=anchorId,proto3" json:"anchor_id,omitempty"`
	Size     int32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Reverse  bool  `protobuf:"varint,3,opt,name=reverse,proto3" json:"reverse,omitempty"`
}

func (x *ListCollectionRequest) Reset() {
	*x = ListCollectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCollectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCollectionRequest) ProtoMessage() {}

func (x *ListCollectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCollectionRequest.ProtoReflect.Descriptor instead.
func (*ListCollectionRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *ListCollectionRequest) GetAnchorId() int64 {
	if x != nil {
		return x.AnchorId
	}
	return 0
}

func (x *ListCollectionRequest) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *ListCollectionRequest) GetReverse() bool {
	if x != nil {
		return x.Reverse
	}
	return false
}

type ListCollectionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collections []*Collection `protobuf:"bytes,1,rep,name=collections,proto3" json:"collections,omitempty"`
	AnchorId    int64         `protobuf:"varint,2,opt,name=anchor_id,json=anchorId,proto3" json:"anchor_id,omitempty"`
	IsEnd       bool          `protobuf:"varint,3,opt,name=is_end,json=isEnd,proto3" json:"is_end,omitempty"`
}

func (x *ListCollectionReply) Reset() {
	*x = ListCollectionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCollectionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCollectionReply) ProtoMessage() {}

func (x *ListCollectionReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCollectionReply.ProtoReflect.Descriptor instead.
func (*ListCollectionReply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *ListCollectionReply) GetCollections() []*Collection {
	if x != nil {
		return x.Collections
	}
	return nil
}

func (x *ListCollectionReply) GetAnchorId() int64 {
	if x != nil {
		return x.AnchorId
	}
	return 0
}

func (x *ListCollectionReply) GetIsEnd() bool {
	if x != nil {
		return x.IsEnd
	}
	return false
}

// The request message containing the user's name.
type Collection struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Logo    string `protobuf:"bytes,3,opt,name=logo,proto3" json:"logo,omitempty"`
	Desc    string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	License string `protobuf:"bytes,5,opt,name=license,proto3" json:"license,omitempty"`
	Address string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	ChainId int64  `protobuf:"varint,7,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Symbol  string `protobuf:"bytes,8,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Mtime   int64  `protobuf:"varint,9,opt,name=Mtime,proto3" json:"Mtime,omitempty"`
}

func (x *Collection) Reset() {
	*x = Collection{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Collection) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Collection) ProtoMessage() {}

func (x *Collection) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Collection.ProtoReflect.Descriptor instead.
func (*Collection) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *Collection) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Collection) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Collection) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *Collection) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Collection) GetLicense() string {
	if x != nil {
		return x.License
	}
	return ""
}

func (x *Collection) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Collection) GetChainId() int64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *Collection) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Collection) GetMtime() int64 {
	if x != nil {
		return x.Mtime
	}
	return 0
}

// The response message containing the greetings
type CreateCollectionReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Id      int64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateCollectionReply) Reset() {
	*x = CreateCollectionReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCollectionReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCollectionReply) ProtoMessage() {}

func (x *CreateCollectionReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCollectionReply.ProtoReflect.Descriptor instead.
func (*CreateCollectionReply) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *CreateCollectionReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateCollectionReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x77, 0x69, 0x72, 0x65, 0x6d, 0x61, 0x6e, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x26, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x62, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61,
	0x6e, 0x63, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72,
	0x65, 0x76, 0x65, 0x72, 0x73, 0x65, 0x22, 0x88, 0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3d,
	0x0a, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x77, 0x69,
	0x72, 0x65, 0x6d, 0x61, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1b, 0x0a,
	0x09, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73,
	0x5f, 0x65, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x45, 0x6e,
	0x64, 0x22, 0xd5, 0x01, 0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x18, 0x0a, 0x07,
	0x6c, 0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c,
	0x69, 0x63, 0x65, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x4d, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x41, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32, 0xde, 0x02, 0x0a,
	0x07, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x12, 0x6f, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x2e, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x77, 0x69, 0x72, 0x65, 0x6d, 0x61, 0x6e, 0x2e, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x26, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x77, 0x69, 0x72, 0x65, 0x6d, 0x61, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x22, 0x0b, 0x2f, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x01, 0x2a, 0x12, 0x73, 0x0a, 0x0e, 0x4c, 0x69, 0x73,
	0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x77, 0x69, 0x72, 0x65, 0x6d, 0x61, 0x6e, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x77, 0x69,
	0x72, 0x65, 0x6d, 0x61, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0d, 0x12, 0x0b, 0x2f, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6d,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x25, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x77, 0x69, 0x72, 0x65, 0x6d, 0x61,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x77, 0x69, 0x72, 0x65, 0x6d, 0x61, 0x6e, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x3f, 0x5a,
	0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x66, 0x69, 0x2f, 0x43, 0x61, 0x6e, 0x74, 0x42, 0x65, 0x45, 0x76, 0x69, 0x6c,
	0x5f, 0x4e, 0x46, 0x54, 0x43, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x42, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_api_proto_goTypes = []interface{}{
	(*GetCollectionRequest)(nil),  // 0: service.wireman.GetCollectionRequest
	(*ListCollectionRequest)(nil), // 1: service.wireman.ListCollectionRequest
	(*ListCollectionReply)(nil),   // 2: service.wireman.ListCollectionReply
	(*Collection)(nil),            // 3: service.wireman.Collection
	(*CreateCollectionReply)(nil), // 4: service.wireman.CreateCollectionReply
}
var file_api_proto_depIdxs = []int32{
	3, // 0: service.wireman.ListCollectionReply.collections:type_name -> service.wireman.Collection
	3, // 1: service.wireman.Greeter.CreateCollection:input_type -> service.wireman.Collection
	1, // 2: service.wireman.Greeter.ListCollection:input_type -> service.wireman.ListCollectionRequest
	0, // 3: service.wireman.Greeter.GetCollection:input_type -> service.wireman.GetCollectionRequest
	4, // 4: service.wireman.Greeter.CreateCollection:output_type -> service.wireman.CreateCollectionReply
	2, // 5: service.wireman.Greeter.ListCollection:output_type -> service.wireman.ListCollectionReply
	3, // 6: service.wireman.Greeter.GetCollection:output_type -> service.wireman.Collection
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCollectionRequest); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCollectionRequest); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCollectionReply); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Collection); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCollectionReply); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}