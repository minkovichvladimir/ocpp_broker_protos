// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: central_system/v1/central_system.proto

package cs_v1

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
	CentralSystemService_CreateCentralSystem_FullMethodName = "/cs.CentralSystemService/CreateCentralSystem"
	CentralSystemService_GetCentralSystem_FullMethodName    = "/cs.CentralSystemService/GetCentralSystem"
	CentralSystemService_ListCentralSystems_FullMethodName  = "/cs.CentralSystemService/ListCentralSystems"
	CentralSystemService_UpdateCentralSystem_FullMethodName = "/cs.CentralSystemService/UpdateCentralSystem"
	CentralSystemService_DeleteCentralSystem_FullMethodName = "/cs.CentralSystemService/DeleteCentralSystem"
)

// CentralSystemServiceClient is the client API for CentralSystemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Сервис для управления центральными системами
type CentralSystemServiceClient interface {
	// Создает новую центральную систему
	CreateCentralSystem(ctx context.Context, in *CreateCentralSystemRequest, opts ...grpc.CallOption) (*CreateCentralSystemResponse, error)
	// Получает центральную систему по ID
	GetCentralSystem(ctx context.Context, in *GetCentralSystemRequest, opts ...grpc.CallOption) (*GetCentralSystemResponse, error)
	// Получает список всех центральных систем
	ListCentralSystems(ctx context.Context, in *ListCentralSystemsRequest, opts ...grpc.CallOption) (*ListCentralSystemsResponse, error)
	// Обновляет центральную систему
	UpdateCentralSystem(ctx context.Context, in *UpdateCentralSystemRequest, opts ...grpc.CallOption) (*UpdateCentralSystemResponse, error)
	// Удаляет центральную систему
	DeleteCentralSystem(ctx context.Context, in *DeleteCentralSystemRequest, opts ...grpc.CallOption) (*DeleteCentralSystemResponse, error)
}

type centralSystemServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCentralSystemServiceClient(cc grpc.ClientConnInterface) CentralSystemServiceClient {
	return &centralSystemServiceClient{cc}
}

func (c *centralSystemServiceClient) CreateCentralSystem(ctx context.Context, in *CreateCentralSystemRequest, opts ...grpc.CallOption) (*CreateCentralSystemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateCentralSystemResponse)
	err := c.cc.Invoke(ctx, CentralSystemService_CreateCentralSystem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centralSystemServiceClient) GetCentralSystem(ctx context.Context, in *GetCentralSystemRequest, opts ...grpc.CallOption) (*GetCentralSystemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCentralSystemResponse)
	err := c.cc.Invoke(ctx, CentralSystemService_GetCentralSystem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centralSystemServiceClient) ListCentralSystems(ctx context.Context, in *ListCentralSystemsRequest, opts ...grpc.CallOption) (*ListCentralSystemsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListCentralSystemsResponse)
	err := c.cc.Invoke(ctx, CentralSystemService_ListCentralSystems_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centralSystemServiceClient) UpdateCentralSystem(ctx context.Context, in *UpdateCentralSystemRequest, opts ...grpc.CallOption) (*UpdateCentralSystemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateCentralSystemResponse)
	err := c.cc.Invoke(ctx, CentralSystemService_UpdateCentralSystem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *centralSystemServiceClient) DeleteCentralSystem(ctx context.Context, in *DeleteCentralSystemRequest, opts ...grpc.CallOption) (*DeleteCentralSystemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteCentralSystemResponse)
	err := c.cc.Invoke(ctx, CentralSystemService_DeleteCentralSystem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CentralSystemServiceServer is the server API for CentralSystemService service.
// All implementations must embed UnimplementedCentralSystemServiceServer
// for forward compatibility.
//
// Сервис для управления центральными системами
type CentralSystemServiceServer interface {
	// Создает новую центральную систему
	CreateCentralSystem(context.Context, *CreateCentralSystemRequest) (*CreateCentralSystemResponse, error)
	// Получает центральную систему по ID
	GetCentralSystem(context.Context, *GetCentralSystemRequest) (*GetCentralSystemResponse, error)
	// Получает список всех центральных систем
	ListCentralSystems(context.Context, *ListCentralSystemsRequest) (*ListCentralSystemsResponse, error)
	// Обновляет центральную систему
	UpdateCentralSystem(context.Context, *UpdateCentralSystemRequest) (*UpdateCentralSystemResponse, error)
	// Удаляет центральную систему
	DeleteCentralSystem(context.Context, *DeleteCentralSystemRequest) (*DeleteCentralSystemResponse, error)
	mustEmbedUnimplementedCentralSystemServiceServer()
}

// UnimplementedCentralSystemServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCentralSystemServiceServer struct{}

func (UnimplementedCentralSystemServiceServer) CreateCentralSystem(context.Context, *CreateCentralSystemRequest) (*CreateCentralSystemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCentralSystem not implemented")
}
func (UnimplementedCentralSystemServiceServer) GetCentralSystem(context.Context, *GetCentralSystemRequest) (*GetCentralSystemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCentralSystem not implemented")
}
func (UnimplementedCentralSystemServiceServer) ListCentralSystems(context.Context, *ListCentralSystemsRequest) (*ListCentralSystemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCentralSystems not implemented")
}
func (UnimplementedCentralSystemServiceServer) UpdateCentralSystem(context.Context, *UpdateCentralSystemRequest) (*UpdateCentralSystemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCentralSystem not implemented")
}
func (UnimplementedCentralSystemServiceServer) DeleteCentralSystem(context.Context, *DeleteCentralSystemRequest) (*DeleteCentralSystemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCentralSystem not implemented")
}
func (UnimplementedCentralSystemServiceServer) mustEmbedUnimplementedCentralSystemServiceServer() {}
func (UnimplementedCentralSystemServiceServer) testEmbeddedByValue()                              {}

// UnsafeCentralSystemServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CentralSystemServiceServer will
// result in compilation errors.
type UnsafeCentralSystemServiceServer interface {
	mustEmbedUnimplementedCentralSystemServiceServer()
}

func RegisterCentralSystemServiceServer(s grpc.ServiceRegistrar, srv CentralSystemServiceServer) {
	// If the following call pancis, it indicates UnimplementedCentralSystemServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CentralSystemService_ServiceDesc, srv)
}

func _CentralSystemService_CreateCentralSystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCentralSystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentralSystemServiceServer).CreateCentralSystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CentralSystemService_CreateCentralSystem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentralSystemServiceServer).CreateCentralSystem(ctx, req.(*CreateCentralSystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentralSystemService_GetCentralSystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCentralSystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentralSystemServiceServer).GetCentralSystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CentralSystemService_GetCentralSystem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentralSystemServiceServer).GetCentralSystem(ctx, req.(*GetCentralSystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentralSystemService_ListCentralSystems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCentralSystemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentralSystemServiceServer).ListCentralSystems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CentralSystemService_ListCentralSystems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentralSystemServiceServer).ListCentralSystems(ctx, req.(*ListCentralSystemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentralSystemService_UpdateCentralSystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCentralSystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentralSystemServiceServer).UpdateCentralSystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CentralSystemService_UpdateCentralSystem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentralSystemServiceServer).UpdateCentralSystem(ctx, req.(*UpdateCentralSystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CentralSystemService_DeleteCentralSystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCentralSystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CentralSystemServiceServer).DeleteCentralSystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CentralSystemService_DeleteCentralSystem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CentralSystemServiceServer).DeleteCentralSystem(ctx, req.(*DeleteCentralSystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CentralSystemService_ServiceDesc is the grpc.ServiceDesc for CentralSystemService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CentralSystemService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cs.CentralSystemService",
	HandlerType: (*CentralSystemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCentralSystem",
			Handler:    _CentralSystemService_CreateCentralSystem_Handler,
		},
		{
			MethodName: "GetCentralSystem",
			Handler:    _CentralSystemService_GetCentralSystem_Handler,
		},
		{
			MethodName: "ListCentralSystems",
			Handler:    _CentralSystemService_ListCentralSystems_Handler,
		},
		{
			MethodName: "UpdateCentralSystem",
			Handler:    _CentralSystemService_UpdateCentralSystem_Handler,
		},
		{
			MethodName: "DeleteCentralSystem",
			Handler:    _CentralSystemService_DeleteCentralSystem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "central_system/v1/central_system.proto",
}
