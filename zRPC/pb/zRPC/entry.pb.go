// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: entry.proto

package zRPC

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

type Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Time     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Prefix   string                 `protobuf:"bytes,3,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Suffix   string                 `protobuf:"bytes,4,opt,name=suffix,proto3" json:"suffix,omitempty"`
	Title    string                 `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Url      string                 `protobuf:"bytes,6,opt,name=url,proto3" json:"url,omitempty"`
	Config   string                 `protobuf:"bytes,7,opt,name=config,proto3" json:"config,omitempty"`
	Operator string                 `protobuf:"bytes,8,opt,name=operator,proto3" json:"operator,omitempty"`
}

func (x *Entry) Reset() {
	*x = Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entry) ProtoMessage() {}

func (x *Entry) ProtoReflect() protoreflect.Message {
	mi := &file_entry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entry.ProtoReflect.Descriptor instead.
func (*Entry) Descriptor() ([]byte, []int) {
	return file_entry_proto_rawDescGZIP(), []int{0}
}

func (x *Entry) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Entry) GetTime() *timestamppb.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

func (x *Entry) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

func (x *Entry) GetSuffix() string {
	if x != nil {
		return x.Suffix
	}
	return ""
}

func (x *Entry) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Entry) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Entry) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

func (x *Entry) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

type ModifyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ModifyResponse) Reset() {
	*x = ModifyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyResponse) ProtoMessage() {}

func (x *ModifyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_entry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyResponse.ProtoReflect.Descriptor instead.
func (*ModifyResponse) Descriptor() ([]byte, []int) {
	return file_entry_proto_rawDescGZIP(), []int{1}
}

func (x *ModifyResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type RetrieveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *RetrieveRequest) Reset() {
	*x = RetrieveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrieveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrieveRequest) ProtoMessage() {}

func (x *RetrieveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrieveRequest.ProtoReflect.Descriptor instead.
func (*RetrieveRequest) Descriptor() ([]byte, []int) {
	return file_entry_proto_rawDescGZIP(), []int{2}
}

func (x *RetrieveRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type InsertResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *InsertResponse) Reset() {
	*x = InsertResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertResponse) ProtoMessage() {}

func (x *InsertResponse) ProtoReflect() protoreflect.Message {
	mi := &file_entry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertResponse.ProtoReflect.Descriptor instead.
func (*InsertResponse) Descriptor() ([]byte, []int) {
	return file_entry_proto_rawDescGZIP(), []int{3}
}

func (x *InsertResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entry_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_entry_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_entry_proto_rawDescGZIP(), []int{4}
}

func (x *ListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type Entries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *Entries) Reset() {
	*x = Entries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entry_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Entries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Entries) ProtoMessage() {}

func (x *Entries) ProtoReflect() protoreflect.Message {
	mi := &file_entry_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Entries.ProtoReflect.Descriptor instead.
func (*Entries) Descriptor() ([]byte, []int) {
	return file_entry_proto_rawDescGZIP(), []int{5}
}

func (x *Entries) GetEntries() []*Entry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type EntryQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *EntryQuery) Reset() {
	*x = EntryQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entry_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntryQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntryQuery) ProtoMessage() {}

func (x *EntryQuery) ProtoReflect() protoreflect.Message {
	mi := &file_entry_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntryQuery.ProtoReflect.Descriptor instead.
func (*EntryQuery) Descriptor() ([]byte, []int) {
	return file_entry_proto_rawDescGZIP(), []int{6}
}

func (x *EntryQuery) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

var File_entry_proto protoreflect.FileDescriptor

var file_entry_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd3, 0x01, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69,
	0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x66, 0x66, 0x69, 0x78, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x20, 0x0a, 0x0e, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x21, 0x0a,
	0x0f, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x20, 0x0a, 0x0e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x25, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x31, 0x0a, 0x07, 0x45, 0x6e, 0x74,
	0x72, 0x69, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x07, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x24, 0x0a, 0x0a,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66,
	0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73,
	0x65, 0x74, 0x32, 0xd3, 0x01, 0x0a, 0x0d, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x06, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x0c,
	0x2e, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x15, 0x2e, 0x65,
	0x6e, 0x74, 0x72, 0x79, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x12,
	0x0c, 0x2e, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x15, 0x2e,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x08, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65,
	0x76, 0x65, 0x12, 0x16, 0x2e, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x12, 0x2e, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x45,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x7a, 0x52,
	0x50, 0x43, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_entry_proto_rawDescOnce sync.Once
	file_entry_proto_rawDescData = file_entry_proto_rawDesc
)

func file_entry_proto_rawDescGZIP() []byte {
	file_entry_proto_rawDescOnce.Do(func() {
		file_entry_proto_rawDescData = protoimpl.X.CompressGZIP(file_entry_proto_rawDescData)
	})
	return file_entry_proto_rawDescData
}

var file_entry_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_entry_proto_goTypes = []interface{}{
	(*Entry)(nil),                 // 0: entry.Entry
	(*ModifyResponse)(nil),        // 1: entry.ModifyResponse
	(*RetrieveRequest)(nil),       // 2: entry.RetrieveRequest
	(*InsertResponse)(nil),        // 3: entry.InsertResponse
	(*ListRequest)(nil),           // 4: entry.ListRequest
	(*Entries)(nil),               // 5: entry.Entries
	(*EntryQuery)(nil),            // 6: entry.EntryQuery
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_entry_proto_depIdxs = []int32{
	7, // 0: entry.Entry.time:type_name -> google.protobuf.Timestamp
	0, // 1: entry.Entries.entries:type_name -> entry.Entry
	0, // 2: entry.EntryServices.Insert:input_type -> entry.Entry
	0, // 3: entry.EntryServices.Modify:input_type -> entry.Entry
	2, // 4: entry.EntryServices.Retrieve:input_type -> entry.RetrieveRequest
	4, // 5: entry.EntryServices.List:input_type -> entry.ListRequest
	3, // 6: entry.EntryServices.Insert:output_type -> entry.InsertResponse
	1, // 7: entry.EntryServices.Modify:output_type -> entry.ModifyResponse
	0, // 8: entry.EntryServices.Retrieve:output_type -> entry.Entry
	5, // 9: entry.EntryServices.List:output_type -> entry.Entries
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_entry_proto_init() }
func file_entry_proto_init() {
	if File_entry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_entry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entry); i {
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
		file_entry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyResponse); i {
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
		file_entry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetrieveRequest); i {
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
		file_entry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertResponse); i {
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
		file_entry_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_entry_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Entries); i {
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
		file_entry_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntryQuery); i {
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
			RawDescriptor: file_entry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_entry_proto_goTypes,
		DependencyIndexes: file_entry_proto_depIdxs,
		MessageInfos:      file_entry_proto_msgTypes,
	}.Build()
	File_entry_proto = out.File
	file_entry_proto_rawDesc = nil
	file_entry_proto_goTypes = nil
	file_entry_proto_depIdxs = nil
}
