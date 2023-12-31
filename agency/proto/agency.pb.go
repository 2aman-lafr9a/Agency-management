// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/agency.proto

package proto

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

type AgencyItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AgencyItem) Reset() {
	*x = AgencyItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_agency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgencyItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgencyItem) ProtoMessage() {}

func (x *AgencyItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_agency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgencyItem.ProtoReflect.Descriptor instead.
func (*AgencyItem) Descriptor() ([]byte, []int) {
	return file_proto_agency_proto_rawDescGZIP(), []int{0}
}

func (x *AgencyItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AgencyItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetAgenciesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAgenciesRequest) Reset() {
	*x = GetAgenciesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_agency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAgenciesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAgenciesRequest) ProtoMessage() {}

func (x *GetAgenciesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_agency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAgenciesRequest.ProtoReflect.Descriptor instead.
func (*GetAgenciesRequest) Descriptor() ([]byte, []int) {
	return file_proto_agency_proto_rawDescGZIP(), []int{1}
}

type GetAgenciesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Agencies []*AgencyItem `protobuf:"bytes,1,rep,name=agencies,proto3" json:"agencies,omitempty"`
}

func (x *GetAgenciesResponse) Reset() {
	*x = GetAgenciesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_agency_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAgenciesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAgenciesResponse) ProtoMessage() {}

func (x *GetAgenciesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_agency_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAgenciesResponse.ProtoReflect.Descriptor instead.
func (*GetAgenciesResponse) Descriptor() ([]byte, []int) {
	return file_proto_agency_proto_rawDescGZIP(), []int{2}
}

func (x *GetAgenciesResponse) GetAgencies() []*AgencyItem {
	if x != nil {
		return x.Agencies
	}
	return nil
}

type GetAgencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetAgencyRequest) Reset() {
	*x = GetAgencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_agency_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAgencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAgencyRequest) ProtoMessage() {}

func (x *GetAgencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_agency_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAgencyRequest.ProtoReflect.Descriptor instead.
func (*GetAgencyRequest) Descriptor() ([]byte, []int) {
	return file_proto_agency_proto_rawDescGZIP(), []int{3}
}

func (x *GetAgencyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetAgencyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetAgencyResponse) Reset() {
	*x = GetAgencyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_agency_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAgencyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAgencyResponse) ProtoMessage() {}

func (x *GetAgencyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_agency_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAgencyResponse.ProtoReflect.Descriptor instead.
func (*GetAgencyResponse) Descriptor() ([]byte, []int) {
	return file_proto_agency_proto_rawDescGZIP(), []int{4}
}

func (x *GetAgencyResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetAgencyResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateAgencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateAgencyRequest) Reset() {
	*x = CreateAgencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_agency_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAgencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAgencyRequest) ProtoMessage() {}

func (x *CreateAgencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_agency_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAgencyRequest.ProtoReflect.Descriptor instead.
func (*CreateAgencyRequest) Descriptor() ([]byte, []int) {
	return file_proto_agency_proto_rawDescGZIP(), []int{5}
}

func (x *CreateAgencyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateAgencyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateAgencyResponse) Reset() {
	*x = CreateAgencyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_agency_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAgencyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAgencyResponse) ProtoMessage() {}

func (x *CreateAgencyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_agency_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAgencyResponse.ProtoReflect.Descriptor instead.
func (*CreateAgencyResponse) Descriptor() ([]byte, []int) {
	return file_proto_agency_proto_rawDescGZIP(), []int{6}
}

func (x *CreateAgencyResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateAgencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateAgencyRequest) Reset() {
	*x = UpdateAgencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_agency_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAgencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAgencyRequest) ProtoMessage() {}

func (x *UpdateAgencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_agency_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAgencyRequest.ProtoReflect.Descriptor instead.
func (*UpdateAgencyRequest) Descriptor() ([]byte, []int) {
	return file_proto_agency_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateAgencyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateAgencyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UpdateAgencyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UpdateAgencyResponse) Reset() {
	*x = UpdateAgencyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_agency_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAgencyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAgencyResponse) ProtoMessage() {}

func (x *UpdateAgencyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_agency_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAgencyResponse.ProtoReflect.Descriptor instead.
func (*UpdateAgencyResponse) Descriptor() ([]byte, []int) {
	return file_proto_agency_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateAgencyResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateAgencyResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteAgencyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteAgencyRequest) Reset() {
	*x = DeleteAgencyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_agency_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAgencyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAgencyRequest) ProtoMessage() {}

func (x *DeleteAgencyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_agency_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAgencyRequest.ProtoReflect.Descriptor instead.
func (*DeleteAgencyRequest) Descriptor() ([]byte, []int) {
	return file_proto_agency_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteAgencyRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteAgencyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteAgencyResponse) Reset() {
	*x = DeleteAgencyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_agency_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAgencyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAgencyResponse) ProtoMessage() {}

func (x *DeleteAgencyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_agency_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAgencyResponse.ProtoReflect.Descriptor instead.
func (*DeleteAgencyResponse) Descriptor() ([]byte, []int) {
	return file_proto_agency_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteAgencyResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DeleteAgencyResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proto_agency_proto protoreflect.FileDescriptor

var file_proto_agency_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x0a, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65,
	0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3e, 0x0a, 0x13,
	0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x74,
	0x65, 0x6d, 0x52, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x22, 0x26, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x37, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x63,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x29, 0x0a,
	0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2a, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x67,
	0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x3a, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x25, 0x0a, 0x13, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x3a, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e,
	0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0xad,
	0x02, 0x0a, 0x06, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x38, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x41, 0x67, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x12, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x67,
	0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79,
	0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x41, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x67,
	0x65, 0x6e, 0x63, 0x79, 0x12, 0x14, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x67, 0x65,
	0x6e, 0x63, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x63,
	0x79, 0x12, 0x14, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x67, 0x65, 0x6e, 0x63, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_agency_proto_rawDescOnce sync.Once
	file_proto_agency_proto_rawDescData = file_proto_agency_proto_rawDesc
)

func file_proto_agency_proto_rawDescGZIP() []byte {
	file_proto_agency_proto_rawDescOnce.Do(func() {
		file_proto_agency_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_agency_proto_rawDescData)
	})
	return file_proto_agency_proto_rawDescData
}

var file_proto_agency_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_agency_proto_goTypes = []interface{}{
	(*AgencyItem)(nil),           // 0: AgencyItem
	(*GetAgenciesRequest)(nil),   // 1: GetAgenciesRequest
	(*GetAgenciesResponse)(nil),  // 2: GetAgenciesResponse
	(*GetAgencyRequest)(nil),     // 3: GetAgencyRequest
	(*GetAgencyResponse)(nil),    // 4: GetAgencyResponse
	(*CreateAgencyRequest)(nil),  // 5: CreateAgencyRequest
	(*CreateAgencyResponse)(nil), // 6: CreateAgencyResponse
	(*UpdateAgencyRequest)(nil),  // 7: UpdateAgencyRequest
	(*UpdateAgencyResponse)(nil), // 8: UpdateAgencyResponse
	(*DeleteAgencyRequest)(nil),  // 9: DeleteAgencyRequest
	(*DeleteAgencyResponse)(nil), // 10: DeleteAgencyResponse
}
var file_proto_agency_proto_depIdxs = []int32{
	0,  // 0: GetAgenciesResponse.agencies:type_name -> AgencyItem
	1,  // 1: Agency.GetAgencies:input_type -> GetAgenciesRequest
	3,  // 2: Agency.GetAgency:input_type -> GetAgencyRequest
	5,  // 3: Agency.CreateAgency:input_type -> CreateAgencyRequest
	7,  // 4: Agency.UpdateAgency:input_type -> UpdateAgencyRequest
	9,  // 5: Agency.DeleteAgency:input_type -> DeleteAgencyRequest
	2,  // 6: Agency.GetAgencies:output_type -> GetAgenciesResponse
	4,  // 7: Agency.GetAgency:output_type -> GetAgencyResponse
	6,  // 8: Agency.CreateAgency:output_type -> CreateAgencyResponse
	8,  // 9: Agency.UpdateAgency:output_type -> UpdateAgencyResponse
	10, // 10: Agency.DeleteAgency:output_type -> DeleteAgencyResponse
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_proto_agency_proto_init() }
func file_proto_agency_proto_init() {
	if File_proto_agency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_agency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgencyItem); i {
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
		file_proto_agency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAgenciesRequest); i {
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
		file_proto_agency_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAgenciesResponse); i {
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
		file_proto_agency_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAgencyRequest); i {
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
		file_proto_agency_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAgencyResponse); i {
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
		file_proto_agency_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAgencyRequest); i {
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
		file_proto_agency_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAgencyResponse); i {
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
		file_proto_agency_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAgencyRequest); i {
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
		file_proto_agency_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAgencyResponse); i {
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
		file_proto_agency_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAgencyRequest); i {
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
		file_proto_agency_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAgencyResponse); i {
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
			RawDescriptor: file_proto_agency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_agency_proto_goTypes,
		DependencyIndexes: file_proto_agency_proto_depIdxs,
		MessageInfos:      file_proto_agency_proto_msgTypes,
	}.Build()
	File_proto_agency_proto = out.File
	file_proto_agency_proto_rawDesc = nil
	file_proto_agency_proto_goTypes = nil
	file_proto_agency_proto_depIdxs = nil
}
