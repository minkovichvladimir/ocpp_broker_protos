// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/ocpp.proto

package proto

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

type OCPPMessageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StationId     string                 `protobuf:"bytes,1,opt,name=station_id,json=stationId,proto3" json:"station_id,omitempty"`
	Message       []byte                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	MessageType   string                 `protobuf:"bytes,3,opt,name=message_type,json=messageType,proto3" json:"message_type,omitempty"`
	RequestId     string                 `protobuf:"bytes,4,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OCPPMessageRequest) Reset() {
	*x = OCPPMessageRequest{}
	mi := &file_proto_ocpp_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OCPPMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OCPPMessageRequest) ProtoMessage() {}

func (x *OCPPMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ocpp_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OCPPMessageRequest.ProtoReflect.Descriptor instead.
func (*OCPPMessageRequest) Descriptor() ([]byte, []int) {
	return file_proto_ocpp_proto_rawDescGZIP(), []int{0}
}

func (x *OCPPMessageRequest) GetStationId() string {
	if x != nil {
		return x.StationId
	}
	return ""
}

func (x *OCPPMessageRequest) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *OCPPMessageRequest) GetMessageType() string {
	if x != nil {
		return x.MessageType
	}
	return ""
}

func (x *OCPPMessageRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

type OCPPMessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       []byte                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	RequestId     string                 `protobuf:"bytes,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *OCPPMessageResponse) Reset() {
	*x = OCPPMessageResponse{}
	mi := &file_proto_ocpp_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *OCPPMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OCPPMessageResponse) ProtoMessage() {}

func (x *OCPPMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_ocpp_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OCPPMessageResponse.ProtoReflect.Descriptor instead.
func (*OCPPMessageResponse) Descriptor() ([]byte, []int) {
	return file_proto_ocpp_proto_rawDescGZIP(), []int{1}
}

func (x *OCPPMessageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *OCPPMessageResponse) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *OCPPMessageResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

var File_proto_ocpp_proto protoreflect.FileDescriptor

var file_proto_ocpp_proto_rawDesc = string([]byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x63, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x6f, 0x63, 0x70, 0x70, 0x22, 0x8f, 0x01, 0x0a, 0x12, 0x4f, 0x43, 0x50,
	0x50, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x22, 0x68, 0x0a, 0x13, 0x4f, 0x43,
	0x50, 0x50, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x32, 0x57, 0x0a, 0x0b, 0x4f, 0x43, 0x50, 0x50, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x43, 0x50, 0x50, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x2e, 0x6f, 0x63, 0x70, 0x70, 0x2e, 0x4f, 0x43,
	0x50, 0x50, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x6f, 0x63, 0x70, 0x70, 0x2e, 0x4f, 0x43, 0x50, 0x50, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x37, 0x5a,
	0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x69, 0x6e, 0x6b,
	0x6f, 0x76, 0x69, 0x63, 0x68, 0x76, 0x6c, 0x61, 0x64, 0x69, 0x6d, 0x69, 0x72, 0x2f, 0x6f, 0x63,
	0x70, 0x70, 0x5f, 0x62, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_proto_ocpp_proto_rawDescOnce sync.Once
	file_proto_ocpp_proto_rawDescData []byte
)

func file_proto_ocpp_proto_rawDescGZIP() []byte {
	file_proto_ocpp_proto_rawDescOnce.Do(func() {
		file_proto_ocpp_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_ocpp_proto_rawDesc), len(file_proto_ocpp_proto_rawDesc)))
	})
	return file_proto_ocpp_proto_rawDescData
}

var file_proto_ocpp_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_ocpp_proto_goTypes = []any{
	(*OCPPMessageRequest)(nil),  // 0: ocpp.OCPPMessageRequest
	(*OCPPMessageResponse)(nil), // 1: ocpp.OCPPMessageResponse
}
var file_proto_ocpp_proto_depIdxs = []int32{
	0, // 0: ocpp.OCPPService.SendOCPPMessage:input_type -> ocpp.OCPPMessageRequest
	1, // 1: ocpp.OCPPService.SendOCPPMessage:output_type -> ocpp.OCPPMessageResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_ocpp_proto_init() }
func file_proto_ocpp_proto_init() {
	if File_proto_ocpp_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_ocpp_proto_rawDesc), len(file_proto_ocpp_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_ocpp_proto_goTypes,
		DependencyIndexes: file_proto_ocpp_proto_depIdxs,
		MessageInfos:      file_proto_ocpp_proto_msgTypes,
	}.Build()
	File_proto_ocpp_proto = out.File
	file_proto_ocpp_proto_goTypes = nil
	file_proto_ocpp_proto_depIdxs = nil
}
