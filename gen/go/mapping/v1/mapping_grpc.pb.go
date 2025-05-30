// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: mapping/v1/mapping.proto

package mappingv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	MappingService_GetAllMappings_FullMethodName = "/mapping.MappingService/GetAllMappings"
	MappingService_CreateMapping_FullMethodName  = "/mapping.MappingService/CreateMapping"
	MappingService_UpdateMapping_FullMethodName  = "/mapping.MappingService/UpdateMapping"
	MappingService_DeleteMapping_FullMethodName  = "/mapping.MappingService/DeleteMapping"
)

// MappingServiceClient is the client API for MappingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Сервис для управления маппингами станций
type MappingServiceClient interface {
	// Получить все маппинги
	GetAllMappings(ctx context.Context, in *GetAllMappingsRequest, opts ...grpc.CallOption) (*GetAllMappingsResponse, error)
	// Создать новый маппинг
	CreateMapping(ctx context.Context, in *CreateMappingRequest, opts ...grpc.CallOption) (*CreateMappingResponse, error)
	// Обновить существующий маппинг
	UpdateMapping(ctx context.Context, in *UpdateMappingRequest, opts ...grpc.CallOption) (*UpdateMappingResponse, error)
	// Удалить маппинг
	DeleteMapping(ctx context.Context, in *DeleteMappingRequest, opts ...grpc.CallOption) (*DeleteMappingResponse, error)
}

type mappingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMappingServiceClient(cc grpc.ClientConnInterface) MappingServiceClient {
	return &mappingServiceClient{cc}
}

func (c *mappingServiceClient) GetAllMappings(ctx context.Context, in *GetAllMappingsRequest, opts ...grpc.CallOption) (*GetAllMappingsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllMappingsResponse)
	err := c.cc.Invoke(ctx, MappingService_GetAllMappings_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mappingServiceClient) CreateMapping(ctx context.Context, in *CreateMappingRequest, opts ...grpc.CallOption) (*CreateMappingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateMappingResponse)
	err := c.cc.Invoke(ctx, MappingService_CreateMapping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mappingServiceClient) UpdateMapping(ctx context.Context, in *UpdateMappingRequest, opts ...grpc.CallOption) (*UpdateMappingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateMappingResponse)
	err := c.cc.Invoke(ctx, MappingService_UpdateMapping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mappingServiceClient) DeleteMapping(ctx context.Context, in *DeleteMappingRequest, opts ...grpc.CallOption) (*DeleteMappingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteMappingResponse)
	err := c.cc.Invoke(ctx, MappingService_DeleteMapping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MappingServiceServer is the server API for MappingService service.
// All implementations must embed UnimplementedMappingServiceServer
// for forward compatibility.
//
// Сервис для управления маппингами станций
type MappingServiceServer interface {
	// Получить все маппинги
	GetAllMappings(context.Context, *GetAllMappingsRequest) (*GetAllMappingsResponse, error)
	// Создать новый маппинг
	CreateMapping(context.Context, *CreateMappingRequest) (*CreateMappingResponse, error)
	// Обновить существующий маппинг
	UpdateMapping(context.Context, *UpdateMappingRequest) (*UpdateMappingResponse, error)
	// Удалить маппинг
	DeleteMapping(context.Context, *DeleteMappingRequest) (*DeleteMappingResponse, error)
	mustEmbedUnimplementedMappingServiceServer()
}

// UnimplementedMappingServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMappingServiceServer struct{}

func (UnimplementedMappingServiceServer) GetAllMappings(context.Context, *GetAllMappingsRequest) (*GetAllMappingsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllMappings not implemented")
}
func (UnimplementedMappingServiceServer) CreateMapping(context.Context, *CreateMappingRequest) (*CreateMappingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateMapping not implemented")
}
func (UnimplementedMappingServiceServer) UpdateMapping(context.Context, *UpdateMappingRequest) (*UpdateMappingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMapping not implemented")
}
func (UnimplementedMappingServiceServer) DeleteMapping(context.Context, *DeleteMappingRequest) (*DeleteMappingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMapping not implemented")
}
func (UnimplementedMappingServiceServer) mustEmbedUnimplementedMappingServiceServer() {}
func (UnimplementedMappingServiceServer) testEmbeddedByValue()                        {}

// UnsafeMappingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MappingServiceServer will
// result in compilation errors.
type UnsafeMappingServiceServer interface {
	mustEmbedUnimplementedMappingServiceServer()
}

func RegisterMappingServiceServer(s grpc.ServiceRegistrar, srv MappingServiceServer) {
	// If the following call pancis, it indicates UnimplementedMappingServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MappingService_ServiceDesc, srv)
}

func _MappingService_GetAllMappings_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllMappingsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MappingServiceServer).GetAllMappings(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MappingService_GetAllMappings_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MappingServiceServer).GetAllMappings(ctx, req.(*GetAllMappingsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MappingService_CreateMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MappingServiceServer).CreateMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MappingService_CreateMapping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MappingServiceServer).CreateMapping(ctx, req.(*CreateMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MappingService_UpdateMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MappingServiceServer).UpdateMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MappingService_UpdateMapping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MappingServiceServer).UpdateMapping(ctx, req.(*UpdateMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MappingService_DeleteMapping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteMappingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MappingServiceServer).DeleteMapping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MappingService_DeleteMapping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MappingServiceServer).DeleteMapping(ctx, req.(*DeleteMappingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MappingService_ServiceDesc is the grpc.ServiceDesc for MappingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MappingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mapping.MappingService",
	HandlerType: (*MappingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllMappings",
			Handler:    _MappingService_GetAllMappings_Handler,
		},
		{
			MethodName: "CreateMapping",
			Handler:    _MappingService_CreateMapping_Handler,
		},
		{
			MethodName: "UpdateMapping",
			Handler:    _MappingService_UpdateMapping_Handler,
		},
		{
			MethodName: "DeleteMapping",
			Handler:    _MappingService_DeleteMapping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mapping/v1/mapping.proto",
}
