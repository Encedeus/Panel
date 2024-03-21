// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.4
// source: servers_api.proto

package protoapi

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

type ServersCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Crater        string `protobuf:"bytes,2,opt,name=crater,proto3" json:"crater,omitempty"`
	CraterVariant string `protobuf:"bytes,3,opt,name=crater_variant,json=craterVariant,proto3" json:"crater_variant,omitempty"`
	Owner         *UUID  `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty"`
	Node          *UUID  `protobuf:"bytes,5,opt,name=node,proto3" json:"node,omitempty"`
	Ram           uint64 `protobuf:"varint,6,opt,name=ram,proto3" json:"ram,omitempty"`
	Storage       uint64 `protobuf:"varint,7,opt,name=storage,proto3" json:"storage,omitempty"`
	LogicalCores  uint32 `protobuf:"varint,8,opt,name=logical_cores,json=logicalCores,proto3" json:"logical_cores,omitempty"`
}

func (x *ServersCreateRequest) Reset() {
	*x = ServersCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servers_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServersCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServersCreateRequest) ProtoMessage() {}

func (x *ServersCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_servers_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServersCreateRequest.ProtoReflect.Descriptor instead.
func (*ServersCreateRequest) Descriptor() ([]byte, []int) {
	return file_servers_api_proto_rawDescGZIP(), []int{0}
}

func (x *ServersCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ServersCreateRequest) GetCrater() string {
	if x != nil {
		return x.Crater
	}
	return ""
}

func (x *ServersCreateRequest) GetCraterVariant() string {
	if x != nil {
		return x.CraterVariant
	}
	return ""
}

func (x *ServersCreateRequest) GetOwner() *UUID {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *ServersCreateRequest) GetNode() *UUID {
	if x != nil {
		return x.Node
	}
	return nil
}

func (x *ServersCreateRequest) GetRam() uint64 {
	if x != nil {
		return x.Ram
	}
	return 0
}

func (x *ServersCreateRequest) GetStorage() uint64 {
	if x != nil {
		return x.Storage
	}
	return 0
}

func (x *ServersCreateRequest) GetLogicalCores() uint32 {
	if x != nil {
		return x.LogicalCores
	}
	return 0
}

type ServersCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Servers []*Server `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
}

func (x *ServersCreateResponse) Reset() {
	*x = ServersCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servers_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServersCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServersCreateResponse) ProtoMessage() {}

func (x *ServersCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_servers_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServersCreateResponse.ProtoReflect.Descriptor instead.
func (*ServersCreateResponse) Descriptor() ([]byte, []int) {
	return file_servers_api_proto_rawDescGZIP(), []int{1}
}

func (x *ServersCreateResponse) GetServers() []*Server {
	if x != nil {
		return x.Servers
	}
	return nil
}

type ServersFindAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BackendOnly  bool `protobuf:"varint,1,opt,name=backendOnly,proto3" json:"backendOnly,omitempty"`
	FrontendOnly bool `protobuf:"varint,2,opt,name=frontendOnly,proto3" json:"frontendOnly,omitempty"`
}

func (x *ServersFindAllRequest) Reset() {
	*x = ServersFindAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servers_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServersFindAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServersFindAllRequest) ProtoMessage() {}

func (x *ServersFindAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_servers_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServersFindAllRequest.ProtoReflect.Descriptor instead.
func (*ServersFindAllRequest) Descriptor() ([]byte, []int) {
	return file_servers_api_proto_rawDescGZIP(), []int{2}
}

func (x *ServersFindAllRequest) GetBackendOnly() bool {
	if x != nil {
		return x.BackendOnly
	}
	return false
}

func (x *ServersFindAllRequest) GetFrontendOnly() bool {
	if x != nil {
		return x.FrontendOnly
	}
	return false
}

type ServersFindAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Servers []*Server `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
}

func (x *ServersFindAllResponse) Reset() {
	*x = ServersFindAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servers_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServersFindAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServersFindAllResponse) ProtoMessage() {}

func (x *ServersFindAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_servers_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServersFindAllResponse.ProtoReflect.Descriptor instead.
func (*ServersFindAllResponse) Descriptor() ([]byte, []int) {
	return file_servers_api_proto_rawDescGZIP(), []int{3}
}

func (x *ServersFindAllResponse) GetServers() []*Server {
	if x != nil {
		return x.Servers
	}
	return nil
}

type ServersFindOneRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *UUID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ServersFindOneRequest) Reset() {
	*x = ServersFindOneRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servers_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServersFindOneRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServersFindOneRequest) ProtoMessage() {}

func (x *ServersFindOneRequest) ProtoReflect() protoreflect.Message {
	mi := &file_servers_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServersFindOneRequest.ProtoReflect.Descriptor instead.
func (*ServersFindOneRequest) Descriptor() ([]byte, []int) {
	return file_servers_api_proto_rawDescGZIP(), []int{4}
}

func (x *ServersFindOneRequest) GetId() *UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

type ServersFindOneResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Servers []*Server `protobuf:"bytes,1,rep,name=servers,proto3" json:"servers,omitempty"`
}

func (x *ServersFindOneResponse) Reset() {
	*x = ServersFindOneResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servers_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServersFindOneResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServersFindOneResponse) ProtoMessage() {}

func (x *ServersFindOneResponse) ProtoReflect() protoreflect.Message {
	mi := &file_servers_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServersFindOneResponse.ProtoReflect.Descriptor instead.
func (*ServersFindOneResponse) Descriptor() ([]byte, []int) {
	return file_servers_api_proto_rawDescGZIP(), []int{5}
}

func (x *ServersFindOneResponse) GetServers() []*Server {
	if x != nil {
		return x.Servers
	}
	return nil
}

type ServersDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *UUID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ServersDeleteRequest) Reset() {
	*x = ServersDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servers_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServersDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServersDeleteRequest) ProtoMessage() {}

func (x *ServersDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_servers_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServersDeleteRequest.ProtoReflect.Descriptor instead.
func (*ServersDeleteRequest) Descriptor() ([]byte, []int) {
	return file_servers_api_proto_rawDescGZIP(), []int{6}
}

func (x *ServersDeleteRequest) GetId() *UUID {
	if x != nil {
		return x.Id
	}
	return nil
}

type ServersDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServersDeleteResponse) Reset() {
	*x = ServersDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servers_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServersDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServersDeleteResponse) ProtoMessage() {}

func (x *ServersDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_servers_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServersDeleteResponse.ProtoReflect.Descriptor instead.
func (*ServersDeleteResponse) Descriptor() ([]byte, []int) {
	return file_servers_api_proto_rawDescGZIP(), []int{7}
}

type ServersGetStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ServersGetStatusResponse) Reset() {
	*x = ServersGetStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_servers_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServersGetStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServersGetStatusResponse) ProtoMessage() {}

func (x *ServersGetStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_servers_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServersGetStatusResponse.ProtoReflect.Descriptor instead.
func (*ServersGetStatusResponse) Descriptor() ([]byte, []int) {
	return file_servers_api_proto_rawDescGZIP(), []int{8}
}

func (x *ServersGetStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_servers_api_proto protoreflect.FileDescriptor

var file_servers_api_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xf2, 0x01, 0x0a, 0x14, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x63, 0x72, 0x61, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63,
	0x72, 0x61, 0x74, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x72, 0x61, 0x74, 0x65, 0x72, 0x5f,
	0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x72, 0x61, 0x74, 0x65, 0x72, 0x56, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x55,
	0x49, 0x44, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x04, 0x6e, 0x6f, 0x64,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x04,
	0x6e, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x61, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x72, 0x61, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x72, 0x65,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c,
	0x43, 0x6f, 0x72, 0x65, 0x73, 0x22, 0x3a, 0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21,
	0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x73, 0x22, 0x5d, 0x0a, 0x15, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x46, 0x69, 0x6e, 0x64,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x61,
	0x63, 0x6b, 0x65, 0x6e, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0b, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x12, 0x22, 0x0a, 0x0c,
	0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x4f, 0x6e, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0c, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x64, 0x4f, 0x6e, 0x6c, 0x79,
	0x22, 0x3b, 0x0a, 0x16, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x46, 0x69, 0x6e, 0x64, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x22, 0x2e, 0x0a,
	0x15, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x05, 0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3b, 0x0a,
	0x16, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x22, 0x2d, 0x0a, 0x14, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x15, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05,
	0x2e, 0x55, 0x55, 0x49, 0x44, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x32, 0x0a, 0x18, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x67, 0x6f, 0x3b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_servers_api_proto_rawDescOnce sync.Once
	file_servers_api_proto_rawDescData = file_servers_api_proto_rawDesc
)

func file_servers_api_proto_rawDescGZIP() []byte {
	file_servers_api_proto_rawDescOnce.Do(func() {
		file_servers_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_servers_api_proto_rawDescData)
	})
	return file_servers_api_proto_rawDescData
}

var file_servers_api_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_servers_api_proto_goTypes = []interface{}{
	(*ServersCreateRequest)(nil),     // 0: ServersCreateRequest
	(*ServersCreateResponse)(nil),    // 1: ServersCreateResponse
	(*ServersFindAllRequest)(nil),    // 2: ServersFindAllRequest
	(*ServersFindAllResponse)(nil),   // 3: ServersFindAllResponse
	(*ServersFindOneRequest)(nil),    // 4: ServersFindOneRequest
	(*ServersFindOneResponse)(nil),   // 5: ServersFindOneResponse
	(*ServersDeleteRequest)(nil),     // 6: ServersDeleteRequest
	(*ServersDeleteResponse)(nil),    // 7: ServersDeleteResponse
	(*ServersGetStatusResponse)(nil), // 8: ServersGetStatusResponse
	(*UUID)(nil),                     // 9: UUID
	(*Server)(nil),                   // 10: Server
}
var file_servers_api_proto_depIdxs = []int32{
	9,  // 0: ServersCreateRequest.owner:type_name -> UUID
	9,  // 1: ServersCreateRequest.node:type_name -> UUID
	10, // 2: ServersCreateResponse.servers:type_name -> Server
	10, // 3: ServersFindAllResponse.servers:type_name -> Server
	9,  // 4: ServersFindOneRequest.id:type_name -> UUID
	10, // 5: ServersFindOneResponse.servers:type_name -> Server
	9,  // 6: ServersDeleteRequest.id:type_name -> UUID
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_servers_api_proto_init() }
func file_servers_api_proto_init() {
	if File_servers_api_proto != nil {
		return
	}
	file_generic_proto_init()
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_servers_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServersCreateRequest); i {
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
		file_servers_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServersCreateResponse); i {
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
		file_servers_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServersFindAllRequest); i {
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
		file_servers_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServersFindAllResponse); i {
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
		file_servers_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServersFindOneRequest); i {
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
		file_servers_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServersFindOneResponse); i {
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
		file_servers_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServersDeleteRequest); i {
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
		file_servers_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServersDeleteResponse); i {
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
		file_servers_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServersGetStatusResponse); i {
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
			RawDescriptor: file_servers_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_servers_api_proto_goTypes,
		DependencyIndexes: file_servers_api_proto_depIdxs,
		MessageInfos:      file_servers_api_proto_msgTypes,
	}.Build()
	File_servers_api_proto = out.File
	file_servers_api_proto_rawDesc = nil
	file_servers_api_proto_goTypes = nil
	file_servers_api_proto_depIdxs = nil
}
