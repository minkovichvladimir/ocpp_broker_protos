// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: central_system/v1/central_system.proto

package central_system_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

// Центральная система
type CentralSystem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	WebsocketUrl  string                 `protobuf:"bytes,3,opt,name=websocket_url,json=websocketUrl,proto3" json:"websocket_url,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CentralSystem) Reset() {
	*x = CentralSystem{}
	mi := &file_central_system_v1_central_system_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CentralSystem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CentralSystem) ProtoMessage() {}

func (x *CentralSystem) ProtoReflect() protoreflect.Message {
	mi := &file_central_system_v1_central_system_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CentralSystem.ProtoReflect.Descriptor instead.
func (*CentralSystem) Descriptor() ([]byte, []int) {
	return file_central_system_v1_central_system_proto_rawDescGZIP(), []int{0}
}

func (x *CentralSystem) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CentralSystem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CentralSystem) GetWebsocketUrl() string {
	if x != nil {
		return x.WebsocketUrl
	}
	return ""
}

func (x *CentralSystem) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CentralSystem) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

// Запрос на создание центральной системы
type CreateCentralSystemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	WebsocketUrl  string                 `protobuf:"bytes,2,opt,name=websocket_url,json=websocketUrl,proto3" json:"websocket_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCentralSystemRequest) Reset() {
	*x = CreateCentralSystemRequest{}
	mi := &file_central_system_v1_central_system_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCentralSystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCentralSystemRequest) ProtoMessage() {}

func (x *CreateCentralSystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_central_system_v1_central_system_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCentralSystemRequest.ProtoReflect.Descriptor instead.
func (*CreateCentralSystemRequest) Descriptor() ([]byte, []int) {
	return file_central_system_v1_central_system_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCentralSystemRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCentralSystemRequest) GetWebsocketUrl() string {
	if x != nil {
		return x.WebsocketUrl
	}
	return ""
}

// Ответ на создание центральной системы
type CreateCentralSystemResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CentralSystem *CentralSystem         `protobuf:"bytes,1,opt,name=central_system,json=centralSystem,proto3" json:"central_system,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateCentralSystemResponse) Reset() {
	*x = CreateCentralSystemResponse{}
	mi := &file_central_system_v1_central_system_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateCentralSystemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCentralSystemResponse) ProtoMessage() {}

func (x *CreateCentralSystemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_central_system_v1_central_system_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCentralSystemResponse.ProtoReflect.Descriptor instead.
func (*CreateCentralSystemResponse) Descriptor() ([]byte, []int) {
	return file_central_system_v1_central_system_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCentralSystemResponse) GetCentralSystem() *CentralSystem {
	if x != nil {
		return x.CentralSystem
	}
	return nil
}

// Запрос на получение центральной системы
type GetCentralSystemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCentralSystemRequest) Reset() {
	*x = GetCentralSystemRequest{}
	mi := &file_central_system_v1_central_system_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCentralSystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCentralSystemRequest) ProtoMessage() {}

func (x *GetCentralSystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_central_system_v1_central_system_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCentralSystemRequest.ProtoReflect.Descriptor instead.
func (*GetCentralSystemRequest) Descriptor() ([]byte, []int) {
	return file_central_system_v1_central_system_proto_rawDescGZIP(), []int{3}
}

func (x *GetCentralSystemRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Ответ на получение центральной системы
type GetCentralSystemResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CentralSystem *CentralSystem         `protobuf:"bytes,1,opt,name=central_system,json=centralSystem,proto3" json:"central_system,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetCentralSystemResponse) Reset() {
	*x = GetCentralSystemResponse{}
	mi := &file_central_system_v1_central_system_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetCentralSystemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCentralSystemResponse) ProtoMessage() {}

func (x *GetCentralSystemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_central_system_v1_central_system_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCentralSystemResponse.ProtoReflect.Descriptor instead.
func (*GetCentralSystemResponse) Descriptor() ([]byte, []int) {
	return file_central_system_v1_central_system_proto_rawDescGZIP(), []int{4}
}

func (x *GetCentralSystemResponse) GetCentralSystem() *CentralSystem {
	if x != nil {
		return x.CentralSystem
	}
	return nil
}

// Запрос на получение списка центральных систем
type ListCentralSystemsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListCentralSystemsRequest) Reset() {
	*x = ListCentralSystemsRequest{}
	mi := &file_central_system_v1_central_system_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCentralSystemsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCentralSystemsRequest) ProtoMessage() {}

func (x *ListCentralSystemsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_central_system_v1_central_system_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCentralSystemsRequest.ProtoReflect.Descriptor instead.
func (*ListCentralSystemsRequest) Descriptor() ([]byte, []int) {
	return file_central_system_v1_central_system_proto_rawDescGZIP(), []int{5}
}

// Ответ на получение списка центральных систем
type ListCentralSystemsResponse struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	CentralSystems []*CentralSystem       `protobuf:"bytes,1,rep,name=central_systems,json=centralSystems,proto3" json:"central_systems,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ListCentralSystemsResponse) Reset() {
	*x = ListCentralSystemsResponse{}
	mi := &file_central_system_v1_central_system_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListCentralSystemsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCentralSystemsResponse) ProtoMessage() {}

func (x *ListCentralSystemsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_central_system_v1_central_system_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCentralSystemsResponse.ProtoReflect.Descriptor instead.
func (*ListCentralSystemsResponse) Descriptor() ([]byte, []int) {
	return file_central_system_v1_central_system_proto_rawDescGZIP(), []int{6}
}

func (x *ListCentralSystemsResponse) GetCentralSystems() []*CentralSystem {
	if x != nil {
		return x.CentralSystems
	}
	return nil
}

// Запрос на обновление центральной системы
type UpdateCentralSystemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	WebsocketUrl  string                 `protobuf:"bytes,3,opt,name=websocket_url,json=websocketUrl,proto3" json:"websocket_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCentralSystemRequest) Reset() {
	*x = UpdateCentralSystemRequest{}
	mi := &file_central_system_v1_central_system_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCentralSystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCentralSystemRequest) ProtoMessage() {}

func (x *UpdateCentralSystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_central_system_v1_central_system_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCentralSystemRequest.ProtoReflect.Descriptor instead.
func (*UpdateCentralSystemRequest) Descriptor() ([]byte, []int) {
	return file_central_system_v1_central_system_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateCentralSystemRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateCentralSystemRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateCentralSystemRequest) GetWebsocketUrl() string {
	if x != nil {
		return x.WebsocketUrl
	}
	return ""
}

// Ответ на обновление центральной системы
type UpdateCentralSystemResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	CentralSystem *CentralSystem         `protobuf:"bytes,1,opt,name=central_system,json=centralSystem,proto3" json:"central_system,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateCentralSystemResponse) Reset() {
	*x = UpdateCentralSystemResponse{}
	mi := &file_central_system_v1_central_system_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateCentralSystemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCentralSystemResponse) ProtoMessage() {}

func (x *UpdateCentralSystemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_central_system_v1_central_system_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCentralSystemResponse.ProtoReflect.Descriptor instead.
func (*UpdateCentralSystemResponse) Descriptor() ([]byte, []int) {
	return file_central_system_v1_central_system_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateCentralSystemResponse) GetCentralSystem() *CentralSystem {
	if x != nil {
		return x.CentralSystem
	}
	return nil
}

// Запрос на удаление центральной системы
type DeleteCentralSystemRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCentralSystemRequest) Reset() {
	*x = DeleteCentralSystemRequest{}
	mi := &file_central_system_v1_central_system_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCentralSystemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCentralSystemRequest) ProtoMessage() {}

func (x *DeleteCentralSystemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_central_system_v1_central_system_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCentralSystemRequest.ProtoReflect.Descriptor instead.
func (*DeleteCentralSystemRequest) Descriptor() ([]byte, []int) {
	return file_central_system_v1_central_system_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteCentralSystemRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Ответ на удаление центральной системы
type DeleteCentralSystemResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteCentralSystemResponse) Reset() {
	*x = DeleteCentralSystemResponse{}
	mi := &file_central_system_v1_central_system_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteCentralSystemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCentralSystemResponse) ProtoMessage() {}

func (x *DeleteCentralSystemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_central_system_v1_central_system_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCentralSystemResponse.ProtoReflect.Descriptor instead.
func (*DeleteCentralSystemResponse) Descriptor() ([]byte, []int) {
	return file_central_system_v1_central_system_proto_rawDescGZIP(), []int{10}
}

var File_central_system_v1_central_system_proto protoreflect.FileDescriptor

var file_central_system_v1_central_system_proto_rawDesc = string([]byte{
	0x0a, 0x26, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x6f, 0x63, 0x70, 0x70, 0x5f, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x22, 0x96, 0x01, 0x0a, 0x0d, 0x43, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a,
	0x0d, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x55,
	0x72, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x55, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61,
	0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f,
	0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x77, 0x65, 0x62, 0x73, 0x6f,
	0x63, 0x6b, 0x65, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x72, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61,
	0x6c, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c,
	0x2e, 0x6f, 0x63, 0x70, 0x70, 0x5f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x63, 0x65, 0x6e,
	0x74, 0x72, 0x61, 0x6c, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x0d, 0x63, 0x65,
	0x6e, 0x74, 0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22, 0x29, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6f, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x43, 0x65, 0x6e,
	0x74, 0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6f, 0x63, 0x70,
	0x70, 0x5f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c,
	0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x6e, 0x74, 0x72,
	0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x0d, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61,
	0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x22, 0x1b, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x73, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x55, 0x0a, 0x0f, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6f, 0x63,
	0x70, 0x70, 0x5f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61,
	0x6c, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x0e, 0x63, 0x65, 0x6e, 0x74, 0x72,
	0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x65, 0x0a, 0x1a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x77,
	0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x77, 0x65, 0x62, 0x73, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x55, 0x72, 0x6c,
	0x22, 0x72, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61,
	0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x53, 0x0a, 0x0e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x6f, 0x63, 0x70, 0x70, 0x5f, 0x62,
	0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x0d, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x22, 0x2c, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x65,
	0x6e, 0x74, 0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x1d, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0xdf, 0x05, 0x0a, 0x14, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x8e, 0x01, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x12, 0x39, 0x2e, 0x6f, 0x63, 0x70, 0x70, 0x5f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e,
	0x6f, 0x63, 0x70, 0x70, 0x5f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x63, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x6c, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x85, 0x01, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x12, 0x36, 0x2e, 0x6f, 0x63, 0x70, 0x70, 0x5f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x63,
	0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x6f, 0x63, 0x70, 0x70, 0x5f,
	0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x8b, 0x01, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x38, 0x2e, 0x6f, 0x63, 0x70,
	0x70, 0x5f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c,
	0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43,
	0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x6f, 0x63, 0x70, 0x70, 0x5f, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x8e, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x39, 0x2e, 0x6f, 0x63, 0x70, 0x70,
	0x5f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x5f,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x6f, 0x63, 0x70, 0x70, 0x5f, 0x62, 0x72, 0x6f, 0x6b,
	0x65, 0x72, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x65, 0x6e, 0x74, 0x72,
	0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x8e, 0x01, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x65, 0x6e,
	0x74, 0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x39, 0x2e, 0x6f, 0x63, 0x70,
	0x70, 0x5f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c,
	0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x6f, 0x63, 0x70, 0x70, 0x5f, 0x62, 0x72, 0x6f,
	0x6b, 0x65, 0x72, 0x2e, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61, 0x6c, 0x5f, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x6c, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x68, 0x5a, 0x66, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6d, 0x69, 0x6e, 0x6b, 0x6f, 0x76, 0x69, 0x63, 0x68, 0x76, 0x6c, 0x61, 0x64, 0x69,
	0x6d, 0x69, 0x72, 0x2f, 0x6f, 0x63, 0x70, 0x70, 0x5f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x6f, 0x63,
	0x70, 0x70, 0x5f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x2f, 0x63, 0x65, 0x6e, 0x74, 0x72, 0x61,
	0x6c, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x65, 0x6e, 0x74,
	0x72, 0x61, 0x6c, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x5f, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_central_system_v1_central_system_proto_rawDescOnce sync.Once
	file_central_system_v1_central_system_proto_rawDescData []byte
)

func file_central_system_v1_central_system_proto_rawDescGZIP() []byte {
	file_central_system_v1_central_system_proto_rawDescOnce.Do(func() {
		file_central_system_v1_central_system_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_central_system_v1_central_system_proto_rawDesc), len(file_central_system_v1_central_system_proto_rawDesc)))
	})
	return file_central_system_v1_central_system_proto_rawDescData
}

var file_central_system_v1_central_system_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_central_system_v1_central_system_proto_goTypes = []any{
	(*CentralSystem)(nil),               // 0: ocpp_broker.central_system.v1.CentralSystem
	(*CreateCentralSystemRequest)(nil),  // 1: ocpp_broker.central_system.v1.CreateCentralSystemRequest
	(*CreateCentralSystemResponse)(nil), // 2: ocpp_broker.central_system.v1.CreateCentralSystemResponse
	(*GetCentralSystemRequest)(nil),     // 3: ocpp_broker.central_system.v1.GetCentralSystemRequest
	(*GetCentralSystemResponse)(nil),    // 4: ocpp_broker.central_system.v1.GetCentralSystemResponse
	(*ListCentralSystemsRequest)(nil),   // 5: ocpp_broker.central_system.v1.ListCentralSystemsRequest
	(*ListCentralSystemsResponse)(nil),  // 6: ocpp_broker.central_system.v1.ListCentralSystemsResponse
	(*UpdateCentralSystemRequest)(nil),  // 7: ocpp_broker.central_system.v1.UpdateCentralSystemRequest
	(*UpdateCentralSystemResponse)(nil), // 8: ocpp_broker.central_system.v1.UpdateCentralSystemResponse
	(*DeleteCentralSystemRequest)(nil),  // 9: ocpp_broker.central_system.v1.DeleteCentralSystemRequest
	(*DeleteCentralSystemResponse)(nil), // 10: ocpp_broker.central_system.v1.DeleteCentralSystemResponse
}
var file_central_system_v1_central_system_proto_depIdxs = []int32{
	0,  // 0: ocpp_broker.central_system.v1.CreateCentralSystemResponse.central_system:type_name -> ocpp_broker.central_system.v1.CentralSystem
	0,  // 1: ocpp_broker.central_system.v1.GetCentralSystemResponse.central_system:type_name -> ocpp_broker.central_system.v1.CentralSystem
	0,  // 2: ocpp_broker.central_system.v1.ListCentralSystemsResponse.central_systems:type_name -> ocpp_broker.central_system.v1.CentralSystem
	0,  // 3: ocpp_broker.central_system.v1.UpdateCentralSystemResponse.central_system:type_name -> ocpp_broker.central_system.v1.CentralSystem
	1,  // 4: ocpp_broker.central_system.v1.CentralSystemService.CreateCentralSystem:input_type -> ocpp_broker.central_system.v1.CreateCentralSystemRequest
	3,  // 5: ocpp_broker.central_system.v1.CentralSystemService.GetCentralSystem:input_type -> ocpp_broker.central_system.v1.GetCentralSystemRequest
	5,  // 6: ocpp_broker.central_system.v1.CentralSystemService.ListCentralSystems:input_type -> ocpp_broker.central_system.v1.ListCentralSystemsRequest
	7,  // 7: ocpp_broker.central_system.v1.CentralSystemService.UpdateCentralSystem:input_type -> ocpp_broker.central_system.v1.UpdateCentralSystemRequest
	9,  // 8: ocpp_broker.central_system.v1.CentralSystemService.DeleteCentralSystem:input_type -> ocpp_broker.central_system.v1.DeleteCentralSystemRequest
	2,  // 9: ocpp_broker.central_system.v1.CentralSystemService.CreateCentralSystem:output_type -> ocpp_broker.central_system.v1.CreateCentralSystemResponse
	4,  // 10: ocpp_broker.central_system.v1.CentralSystemService.GetCentralSystem:output_type -> ocpp_broker.central_system.v1.GetCentralSystemResponse
	6,  // 11: ocpp_broker.central_system.v1.CentralSystemService.ListCentralSystems:output_type -> ocpp_broker.central_system.v1.ListCentralSystemsResponse
	8,  // 12: ocpp_broker.central_system.v1.CentralSystemService.UpdateCentralSystem:output_type -> ocpp_broker.central_system.v1.UpdateCentralSystemResponse
	10, // 13: ocpp_broker.central_system.v1.CentralSystemService.DeleteCentralSystem:output_type -> ocpp_broker.central_system.v1.DeleteCentralSystemResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_central_system_v1_central_system_proto_init() }
func file_central_system_v1_central_system_proto_init() {
	if File_central_system_v1_central_system_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_central_system_v1_central_system_proto_rawDesc), len(file_central_system_v1_central_system_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_central_system_v1_central_system_proto_goTypes,
		DependencyIndexes: file_central_system_v1_central_system_proto_depIdxs,
		MessageInfos:      file_central_system_v1_central_system_proto_msgTypes,
	}.Build()
	File_central_system_v1_central_system_proto = out.File
	file_central_system_v1_central_system_proto_goTypes = nil
	file_central_system_v1_central_system_proto_depIdxs = nil
}
