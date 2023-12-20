// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/agency.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// AgencyServiceClient is the client API for AgencyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AgencyServiceClient interface {
	GetAgencies(ctx context.Context, in *GetAgenciesRequest, opts ...grpc.CallOption) (*GetAgenciesResponse, error)
	GetAgency(ctx context.Context, in *GetAgencyRequest, opts ...grpc.CallOption) (*GetAgencyResponse, error)
	CreateAgency(ctx context.Context, in *CreateAgencyRequest, opts ...grpc.CallOption) (*CreateAgencyResponse, error)
	UpdateAgency(ctx context.Context, in *UpdateAgencyRequest, opts ...grpc.CallOption) (*UpdateAgencyResponse, error)
	DeleteAgency(ctx context.Context, in *DeleteAgencyRequest, opts ...grpc.CallOption) (*DeleteAgencyResponse, error)
}

type agencyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgencyServiceClient(cc grpc.ClientConnInterface) AgencyServiceClient {
	return &agencyServiceClient{cc}
}

func (c *agencyServiceClient) GetAgencies(ctx context.Context, in *GetAgenciesRequest, opts ...grpc.CallOption) (*GetAgenciesResponse, error) {
	out := new(GetAgenciesResponse)
	err := c.cc.Invoke(ctx, "/AgencyService/GetAgencies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agencyServiceClient) GetAgency(ctx context.Context, in *GetAgencyRequest, opts ...grpc.CallOption) (*GetAgencyResponse, error) {
	out := new(GetAgencyResponse)
	err := c.cc.Invoke(ctx, "/AgencyService/GetAgency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agencyServiceClient) CreateAgency(ctx context.Context, in *CreateAgencyRequest, opts ...grpc.CallOption) (*CreateAgencyResponse, error) {
	out := new(CreateAgencyResponse)
	err := c.cc.Invoke(ctx, "/AgencyService/CreateAgency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agencyServiceClient) UpdateAgency(ctx context.Context, in *UpdateAgencyRequest, opts ...grpc.CallOption) (*UpdateAgencyResponse, error) {
	out := new(UpdateAgencyResponse)
	err := c.cc.Invoke(ctx, "/AgencyService/UpdateAgency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agencyServiceClient) DeleteAgency(ctx context.Context, in *DeleteAgencyRequest, opts ...grpc.CallOption) (*DeleteAgencyResponse, error) {
	out := new(DeleteAgencyResponse)
	err := c.cc.Invoke(ctx, "/AgencyService/DeleteAgency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgencyServiceServer is the server API for AgencyService service.
// All implementations must embed UnimplementedAgencyServiceServer
// for forward compatibility
type AgencyServiceServer interface {
	GetAgencies(context.Context, *GetAgenciesRequest) (*GetAgenciesResponse, error)
	GetAgency(context.Context, *GetAgencyRequest) (*GetAgencyResponse, error)
	CreateAgency(context.Context, *CreateAgencyRequest) (*CreateAgencyResponse, error)
	UpdateAgency(context.Context, *UpdateAgencyRequest) (*UpdateAgencyResponse, error)
	DeleteAgency(context.Context, *DeleteAgencyRequest) (*DeleteAgencyResponse, error)
	mustEmbedUnimplementedAgencyServiceServer()
}

// UnimplementedAgencyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAgencyServiceServer struct {
}

func (UnimplementedAgencyServiceServer) GetAgencies(context.Context, *GetAgenciesRequest) (*GetAgenciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgencies not implemented")
}
func (UnimplementedAgencyServiceServer) GetAgency(context.Context, *GetAgencyRequest) (*GetAgencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAgency not implemented")
}
func (UnimplementedAgencyServiceServer) CreateAgency(context.Context, *CreateAgencyRequest) (*CreateAgencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAgency not implemented")
}
func (UnimplementedAgencyServiceServer) UpdateAgency(context.Context, *UpdateAgencyRequest) (*UpdateAgencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAgency not implemented")
}
func (UnimplementedAgencyServiceServer) DeleteAgency(context.Context, *DeleteAgencyRequest) (*DeleteAgencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAgency not implemented")
}
func (UnimplementedAgencyServiceServer) mustEmbedUnimplementedAgencyServiceServer() {}

// UnsafeAgencyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AgencyServiceServer will
// result in compilation errors.
type UnsafeAgencyServiceServer interface {
	mustEmbedUnimplementedAgencyServiceServer()
}

func RegisterAgencyServiceServer(s grpc.ServiceRegistrar, srv AgencyServiceServer) {
	s.RegisterService(&AgencyService_ServiceDesc, srv)
}

func _AgencyService_GetAgencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgenciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgencyServiceServer).GetAgencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AgencyService/GetAgencies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgencyServiceServer).GetAgencies(ctx, req.(*GetAgenciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgencyService_GetAgency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAgencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgencyServiceServer).GetAgency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AgencyService/GetAgency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgencyServiceServer).GetAgency(ctx, req.(*GetAgencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgencyService_CreateAgency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAgencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgencyServiceServer).CreateAgency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AgencyService/CreateAgency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgencyServiceServer).CreateAgency(ctx, req.(*CreateAgencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgencyService_UpdateAgency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAgencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgencyServiceServer).UpdateAgency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AgencyService/UpdateAgency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgencyServiceServer).UpdateAgency(ctx, req.(*UpdateAgencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AgencyService_DeleteAgency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAgencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgencyServiceServer).DeleteAgency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/AgencyService/DeleteAgency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgencyServiceServer).DeleteAgency(ctx, req.(*DeleteAgencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AgencyService_ServiceDesc is the grpc.ServiceDesc for AgencyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AgencyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AgencyService",
	HandlerType: (*AgencyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAgencies",
			Handler:    _AgencyService_GetAgencies_Handler,
		},
		{
			MethodName: "GetAgency",
			Handler:    _AgencyService_GetAgency_Handler,
		},
		{
			MethodName: "CreateAgency",
			Handler:    _AgencyService_CreateAgency_Handler,
		},
		{
			MethodName: "UpdateAgency",
			Handler:    _AgencyService_UpdateAgency_Handler,
		},
		{
			MethodName: "DeleteAgency",
			Handler:    _AgencyService_DeleteAgency_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/agency.proto",
}