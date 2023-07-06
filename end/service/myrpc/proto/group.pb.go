// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.6.1
// source: group.proto

package proto

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

type GroupGetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32  `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32  `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Fuzzy  string `protobuf:"bytes,3,opt,name=fuzzy,proto3" json:"fuzzy,omitempty"`
}

func (x *GroupGetReq) Reset() {
	*x = GroupGetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupGetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupGetReq) ProtoMessage() {}

func (x *GroupGetReq) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupGetReq.ProtoReflect.Descriptor instead.
func (*GroupGetReq) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{0}
}

func (x *GroupGetReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GroupGetReq) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GroupGetReq) GetFuzzy() string {
	if x != nil {
		return x.Fuzzy
	}
	return ""
}

type GroupGetResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string       `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Data    []*GroupInfo `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
	Total   int64        `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *GroupGetResp) Reset() {
	*x = GroupGetResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupGetResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupGetResp) ProtoMessage() {}

func (x *GroupGetResp) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupGetResp.ProtoReflect.Descriptor instead.
func (*GroupGetResp) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{1}
}

func (x *GroupGetResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GroupGetResp) GetData() []*GroupInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GroupGetResp) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GroupInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Details    string `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
	Createtime int64  `protobuf:"varint,4,opt,name=createtime,proto3" json:"createtime,omitempty"`
}

func (x *GroupInfo) Reset() {
	*x = GroupInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupInfo) ProtoMessage() {}

func (x *GroupInfo) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupInfo.ProtoReflect.Descriptor instead.
func (*GroupInfo) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{2}
}

func (x *GroupInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GroupInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GroupInfo) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *GroupInfo) GetCreatetime() int64 {
	if x != nil {
		return x.Createtime
	}
	return 0
}

type GroupPostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Details string `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
}

func (x *GroupPostReq) Reset() {
	*x = GroupPostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupPostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupPostReq) ProtoMessage() {}

func (x *GroupPostReq) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupPostReq.ProtoReflect.Descriptor instead.
func (*GroupPostReq) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{3}
}

func (x *GroupPostReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GroupPostReq) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

type GroupPostResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GroupPostResp) Reset() {
	*x = GroupPostResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupPostResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupPostResp) ProtoMessage() {}

func (x *GroupPostResp) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupPostResp.ProtoReflect.Descriptor instead.
func (*GroupPostResp) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{4}
}

func (x *GroupPostResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GroupDelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GroupDelReq) Reset() {
	*x = GroupDelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupDelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupDelReq) ProtoMessage() {}

func (x *GroupDelReq) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupDelReq.ProtoReflect.Descriptor instead.
func (*GroupDelReq) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{5}
}

func (x *GroupDelReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GroupDelResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GroupDelResp) Reset() {
	*x = GroupDelResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupDelResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupDelResp) ProtoMessage() {}

func (x *GroupDelResp) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupDelResp.ProtoReflect.Descriptor instead.
func (*GroupDelResp) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{6}
}

func (x *GroupDelResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GroupPutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Details string `protobuf:"bytes,2,opt,name=details,proto3" json:"details,omitempty"`
	Id      int32  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GroupPutReq) Reset() {
	*x = GroupPutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupPutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupPutReq) ProtoMessage() {}

func (x *GroupPutReq) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupPutReq.ProtoReflect.Descriptor instead.
func (*GroupPutReq) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{7}
}

func (x *GroupPutReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GroupPutReq) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *GroupPutReq) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GroupPutResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GroupPutResp) Reset() {
	*x = GroupPutResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupPutResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupPutResp) ProtoMessage() {}

func (x *GroupPutResp) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupPutResp.ProtoReflect.Descriptor instead.
func (*GroupPutResp) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{8}
}

func (x *GroupPutResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GroupDevicePostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sn      []string `protobuf:"bytes,1,rep,name=sn,proto3" json:"sn,omitempty"`
	GroupID int32    `protobuf:"varint,2,opt,name=groupID,proto3" json:"groupID,omitempty"`
}

func (x *GroupDevicePostReq) Reset() {
	*x = GroupDevicePostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupDevicePostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupDevicePostReq) ProtoMessage() {}

func (x *GroupDevicePostReq) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupDevicePostReq.ProtoReflect.Descriptor instead.
func (*GroupDevicePostReq) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{9}
}

func (x *GroupDevicePostReq) GetSn() []string {
	if x != nil {
		return x.Sn
	}
	return nil
}

func (x *GroupDevicePostReq) GetGroupID() int32 {
	if x != nil {
		return x.GroupID
	}
	return 0
}

type GroupDevicePostResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GroupDevicePostResp) Reset() {
	*x = GroupDevicePostResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupDevicePostResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupDevicePostResp) ProtoMessage() {}

func (x *GroupDevicePostResp) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupDevicePostResp.ProtoReflect.Descriptor instead.
func (*GroupDevicePostResp) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{10}
}

func (x *GroupDevicePostResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GroupDevicePutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sn      []string `protobuf:"bytes,1,rep,name=sn,proto3" json:"sn,omitempty"`
	GroupID int32    `protobuf:"varint,2,opt,name=groupID,proto3" json:"groupID,omitempty"`
}

func (x *GroupDevicePutReq) Reset() {
	*x = GroupDevicePutReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupDevicePutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupDevicePutReq) ProtoMessage() {}

func (x *GroupDevicePutReq) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupDevicePutReq.ProtoReflect.Descriptor instead.
func (*GroupDevicePutReq) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{11}
}

func (x *GroupDevicePutReq) GetSn() []string {
	if x != nil {
		return x.Sn
	}
	return nil
}

func (x *GroupDevicePutReq) GetGroupID() int32 {
	if x != nil {
		return x.GroupID
	}
	return 0
}

type GroupDevicePutResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *GroupDevicePutResp) Reset() {
	*x = GroupDevicePutResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_group_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupDevicePutResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupDevicePutResp) ProtoMessage() {}

func (x *GroupDevicePutResp) ProtoReflect() protoreflect.Message {
	mi := &file_group_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupDevicePutResp.ProtoReflect.Descriptor instead.
func (*GroupDevicePutResp) Descriptor() ([]byte, []int) {
	return file_group_proto_rawDescGZIP(), []int{12}
}

func (x *GroupDevicePutResp) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_group_proto protoreflect.FileDescriptor

var file_group_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x51, 0x0a, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x47, 0x65, 0x74, 0x52, 0x65,
	0x71, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x75, 0x7a, 0x7a, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x66, 0x75, 0x7a, 0x7a, 0x79, 0x22, 0x64, 0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x69, 0x0a, 0x09, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x22, 0x29, 0x0a, 0x0d, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x1d, 0x0a, 0x0b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28,
	0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x4b, 0x0a, 0x0b, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x0c, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x75,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x3e, 0x0a, 0x12, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x02, 0x73, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x22,
	0x2f, 0x0a, 0x13, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x3d, 0x0a, 0x11, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x73, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x02, 0x73, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x22,
	0x2e, 0x0a, 0x12, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x75,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0x82, 0x04, 0x0a, 0x05, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x46, 0x0a, 0x08, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x47, 0x65, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x11,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x4c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x0e, 0x3a, 0x01, 0x2a, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12,
	0x4d, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x6c, 0x12, 0x12, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x1a,
	0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x2a, 0x10, 0x2f, 0x76,
	0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x7b, 0x69, 0x64, 0x3d, 0x2a, 0x7d, 0x12, 0x49,
	0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x75, 0x74, 0x12, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x13,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x50, 0x75, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01, 0x2a, 0x1a, 0x09,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x65, 0x0a, 0x0f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10,
	0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x62, 0x0a, 0x0e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50,
	0x75, 0x74, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x50, 0x75, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x3a,
	0x01, 0x2a, 0x1a, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x42, 0x1a, 0x5a, 0x18, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x6d, 0x79, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_group_proto_rawDescOnce sync.Once
	file_group_proto_rawDescData = file_group_proto_rawDesc
)

func file_group_proto_rawDescGZIP() []byte {
	file_group_proto_rawDescOnce.Do(func() {
		file_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_group_proto_rawDescData)
	})
	return file_group_proto_rawDescData
}

var file_group_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_group_proto_goTypes = []interface{}{
	(*GroupGetReq)(nil),         // 0: proto.groupGetReq
	(*GroupGetResp)(nil),        // 1: proto.groupGetResp
	(*GroupInfo)(nil),           // 2: proto.groupInfo
	(*GroupPostReq)(nil),        // 3: proto.groupPostReq
	(*GroupPostResp)(nil),       // 4: proto.groupPostResp
	(*GroupDelReq)(nil),         // 5: proto.groupDelReq
	(*GroupDelResp)(nil),        // 6: proto.groupDelResp
	(*GroupPutReq)(nil),         // 7: proto.groupPutReq
	(*GroupPutResp)(nil),        // 8: proto.groupPutResp
	(*GroupDevicePostReq)(nil),  // 9: proto.groupDevicePostReq
	(*GroupDevicePostResp)(nil), // 10: proto.groupDevicePostResp
	(*GroupDevicePutReq)(nil),   // 11: proto.groupDevicePutReq
	(*GroupDevicePutResp)(nil),  // 12: proto.groupDevicePutResp
}
var file_group_proto_depIdxs = []int32{
	2,  // 0: proto.groupGetResp.data:type_name -> proto.groupInfo
	0,  // 1: proto.Group.groupGet:input_type -> proto.groupGetReq
	3,  // 2: proto.Group.groupPost:input_type -> proto.groupPostReq
	5,  // 3: proto.Group.groupDel:input_type -> proto.groupDelReq
	7,  // 4: proto.Group.groupPut:input_type -> proto.groupPutReq
	9,  // 5: proto.Group.groupDevicePost:input_type -> proto.groupDevicePostReq
	11, // 6: proto.Group.groupDevicePut:input_type -> proto.groupDevicePutReq
	1,  // 7: proto.Group.groupGet:output_type -> proto.groupGetResp
	4,  // 8: proto.Group.groupPost:output_type -> proto.groupPostResp
	6,  // 9: proto.Group.groupDel:output_type -> proto.groupDelResp
	8,  // 10: proto.Group.groupPut:output_type -> proto.groupPutResp
	10, // 11: proto.Group.groupDevicePost:output_type -> proto.groupDevicePostResp
	12, // 12: proto.Group.groupDevicePut:output_type -> proto.groupDevicePutResp
	7,  // [7:13] is the sub-list for method output_type
	1,  // [1:7] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_group_proto_init() }
func file_group_proto_init() {
	if File_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupGetReq); i {
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
		file_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupGetResp); i {
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
		file_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupInfo); i {
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
		file_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupPostReq); i {
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
		file_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupPostResp); i {
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
		file_group_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupDelReq); i {
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
		file_group_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupDelResp); i {
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
		file_group_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupPutReq); i {
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
		file_group_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupPutResp); i {
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
		file_group_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupDevicePostReq); i {
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
		file_group_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupDevicePostResp); i {
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
		file_group_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupDevicePutReq); i {
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
		file_group_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupDevicePutResp); i {
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
			RawDescriptor: file_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_group_proto_goTypes,
		DependencyIndexes: file_group_proto_depIdxs,
		MessageInfos:      file_group_proto_msgTypes,
	}.Build()
	File_group_proto = out.File
	file_group_proto_rawDesc = nil
	file_group_proto_goTypes = nil
	file_group_proto_depIdxs = nil
}
