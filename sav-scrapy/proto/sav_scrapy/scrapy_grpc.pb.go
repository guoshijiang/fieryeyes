// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: scrapy.proto

package sav_scrapy

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

// GiantWhaleServiceClient is the client API for GiantWhaleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GiantWhaleServiceClient interface {
	GetSupportChain(ctx context.Context, in *SupportChainReq, opts ...grpc.CallOption) (*SupportChainRep, error)
	SetGiantWhale(ctx context.Context, in *SetGiantWhaleReq, opts ...grpc.CallOption) (*SetGiantWhaleRep, error)
	GetGiantWhale(ctx context.Context, in *GetGiantWhaleReq, opts ...grpc.CallOption) (*GetGiantWhaleRep, error)
	RemoveGiantWhale(ctx context.Context, in *RemoveGiantWhaleReq, opts ...grpc.CallOption) (*RemoveGiantWhaleRep, error)
}

type giantWhaleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGiantWhaleServiceClient(cc grpc.ClientConnInterface) GiantWhaleServiceClient {
	return &giantWhaleServiceClient{cc}
}

func (c *giantWhaleServiceClient) GetSupportChain(ctx context.Context, in *SupportChainReq, opts ...grpc.CallOption) (*SupportChainRep, error) {
	out := new(SupportChainRep)
	err := c.cc.Invoke(ctx, "/savourrpc.keylocker.GiantWhaleService/getSupportChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giantWhaleServiceClient) SetGiantWhale(ctx context.Context, in *SetGiantWhaleReq, opts ...grpc.CallOption) (*SetGiantWhaleRep, error) {
	out := new(SetGiantWhaleRep)
	err := c.cc.Invoke(ctx, "/savourrpc.keylocker.GiantWhaleService/setGiantWhale", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giantWhaleServiceClient) GetGiantWhale(ctx context.Context, in *GetGiantWhaleReq, opts ...grpc.CallOption) (*GetGiantWhaleRep, error) {
	out := new(GetGiantWhaleRep)
	err := c.cc.Invoke(ctx, "/savourrpc.keylocker.GiantWhaleService/getGiantWhale", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *giantWhaleServiceClient) RemoveGiantWhale(ctx context.Context, in *RemoveGiantWhaleReq, opts ...grpc.CallOption) (*RemoveGiantWhaleRep, error) {
	out := new(RemoveGiantWhaleRep)
	err := c.cc.Invoke(ctx, "/savourrpc.keylocker.GiantWhaleService/removeGiantWhale", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GiantWhaleServiceServer is the server API for GiantWhaleService service.
// All implementations must embed UnimplementedGiantWhaleServiceServer
// for forward compatibility
type GiantWhaleServiceServer interface {
	GetSupportChain(context.Context, *SupportChainReq) (*SupportChainRep, error)
	SetGiantWhale(context.Context, *SetGiantWhaleReq) (*SetGiantWhaleRep, error)
	GetGiantWhale(context.Context, *GetGiantWhaleReq) (*GetGiantWhaleRep, error)
	RemoveGiantWhale(context.Context, *RemoveGiantWhaleReq) (*RemoveGiantWhaleRep, error)
}

// UnimplementedGiantWhaleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGiantWhaleServiceServer struct {
}

func (UnimplementedGiantWhaleServiceServer) GetSupportChain(context.Context, *SupportChainReq) (*SupportChainRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSupportChain not implemented")
}
func (UnimplementedGiantWhaleServiceServer) SetGiantWhale(context.Context, *SetGiantWhaleReq) (*SetGiantWhaleRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGiantWhale not implemented")
}
func (UnimplementedGiantWhaleServiceServer) GetGiantWhale(context.Context, *GetGiantWhaleReq) (*GetGiantWhaleRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGiantWhale not implemented")
}
func (UnimplementedGiantWhaleServiceServer) RemoveGiantWhale(context.Context, *RemoveGiantWhaleReq) (*RemoveGiantWhaleRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveGiantWhale not implemented")
}
func (UnimplementedGiantWhaleServiceServer) mustEmbedUnimplementedGiantWhaleServiceServer() {}

// UnsafeGiantWhaleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GiantWhaleServiceServer will
// result in compilation errors.
type UnsafeGiantWhaleServiceServer interface {
	mustEmbedUnimplementedGiantWhaleServiceServer()
}

func RegisterGiantWhaleServiceServer(s grpc.ServiceRegistrar, srv GiantWhaleServiceServer) {
	s.RegisterService(&GiantWhaleService_ServiceDesc, srv)
}

func _GiantWhaleService_GetSupportChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupportChainReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GiantWhaleServiceServer).GetSupportChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/savourrpc.keylocker.GiantWhaleService/getSupportChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GiantWhaleServiceServer).GetSupportChain(ctx, req.(*SupportChainReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GiantWhaleService_SetGiantWhale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGiantWhaleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GiantWhaleServiceServer).SetGiantWhale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/savourrpc.keylocker.GiantWhaleService/setGiantWhale",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GiantWhaleServiceServer).SetGiantWhale(ctx, req.(*SetGiantWhaleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GiantWhaleService_GetGiantWhale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGiantWhaleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GiantWhaleServiceServer).GetGiantWhale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/savourrpc.keylocker.GiantWhaleService/getGiantWhale",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GiantWhaleServiceServer).GetGiantWhale(ctx, req.(*GetGiantWhaleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _GiantWhaleService_RemoveGiantWhale_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveGiantWhaleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GiantWhaleServiceServer).RemoveGiantWhale(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/savourrpc.keylocker.GiantWhaleService/removeGiantWhale",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GiantWhaleServiceServer).RemoveGiantWhale(ctx, req.(*RemoveGiantWhaleReq))
	}
	return interceptor(ctx, in, info, handler)
}

// GiantWhaleService_ServiceDesc is the grpc.ServiceDesc for GiantWhaleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GiantWhaleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "savourrpc.keylocker.GiantWhaleService",
	HandlerType: (*GiantWhaleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getSupportChain",
			Handler:    _GiantWhaleService_GetSupportChain_Handler,
		},
		{
			MethodName: "setGiantWhale",
			Handler:    _GiantWhaleService_SetGiantWhale_Handler,
		},
		{
			MethodName: "getGiantWhale",
			Handler:    _GiantWhaleService_GetGiantWhale_Handler,
		},
		{
			MethodName: "removeGiantWhale",
			Handler:    _GiantWhaleService_RemoveGiantWhale_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scrapy.proto",
}