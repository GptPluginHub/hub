// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: hub.io/api/plugin/v1alpha1/rpc.proto

package v1alpha1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PluginServiceClient is the client API for PluginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PluginServiceClient interface {
	ListPlugins(ctx context.Context, in *ListPluginRequest, opts ...grpc.CallOption) (*ListPluginResponse, error)
	CreatePlugin(ctx context.Context, in *CreatePluginRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreatePluginScore(ctx context.Context, in *CreatePluginScoreRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type pluginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginServiceClient(cc grpc.ClientConnInterface) PluginServiceClient {
	return &pluginServiceClient{cc}
}

func (c *pluginServiceClient) ListPlugins(ctx context.Context, in *ListPluginRequest, opts ...grpc.CallOption) (*ListPluginResponse, error) {
	out := new(ListPluginResponse)
	err := c.cc.Invoke(ctx, "/hub.io.api.plugin.v1alpha1.PluginService/ListPlugins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginServiceClient) CreatePlugin(ctx context.Context, in *CreatePluginRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/hub.io.api.plugin.v1alpha1.PluginService/CreatePlugin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pluginServiceClient) CreatePluginScore(ctx context.Context, in *CreatePluginScoreRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/hub.io.api.plugin.v1alpha1.PluginService/CreatePluginScore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginServiceServer is the server API for PluginService service.
// All implementations must embed UnimplementedPluginServiceServer
// for forward compatibility
type PluginServiceServer interface {
	ListPlugins(context.Context, *ListPluginRequest) (*ListPluginResponse, error)
	CreatePlugin(context.Context, *CreatePluginRequest) (*emptypb.Empty, error)
	CreatePluginScore(context.Context, *CreatePluginScoreRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedPluginServiceServer()
}

// UnimplementedPluginServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPluginServiceServer struct {
}

func (UnimplementedPluginServiceServer) ListPlugins(context.Context, *ListPluginRequest) (*ListPluginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPlugins not implemented")
}
func (UnimplementedPluginServiceServer) CreatePlugin(context.Context, *CreatePluginRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlugin not implemented")
}
func (UnimplementedPluginServiceServer) CreatePluginScore(context.Context, *CreatePluginScoreRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePluginScore not implemented")
}
func (UnimplementedPluginServiceServer) mustEmbedUnimplementedPluginServiceServer() {}

// UnsafePluginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PluginServiceServer will
// result in compilation errors.
type UnsafePluginServiceServer interface {
	mustEmbedUnimplementedPluginServiceServer()
}

func RegisterPluginServiceServer(s grpc.ServiceRegistrar, srv PluginServiceServer) {
	s.RegisterService(&PluginService_ServiceDesc, srv)
}

func _PluginService_ListPlugins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServiceServer).ListPlugins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hub.io.api.plugin.v1alpha1.PluginService/ListPlugins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServiceServer).ListPlugins(ctx, req.(*ListPluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginService_CreatePlugin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePluginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServiceServer).CreatePlugin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hub.io.api.plugin.v1alpha1.PluginService/CreatePlugin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServiceServer).CreatePlugin(ctx, req.(*CreatePluginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PluginService_CreatePluginScore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePluginScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginServiceServer).CreatePluginScore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hub.io.api.plugin.v1alpha1.PluginService/CreatePluginScore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginServiceServer).CreatePluginScore(ctx, req.(*CreatePluginScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PluginService_ServiceDesc is the grpc.ServiceDesc for PluginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PluginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hub.io.api.plugin.v1alpha1.PluginService",
	HandlerType: (*PluginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPlugins",
			Handler:    _PluginService_ListPlugins_Handler,
		},
		{
			MethodName: "CreatePlugin",
			Handler:    _PluginService_CreatePlugin_Handler,
		},
		{
			MethodName: "CreatePluginScore",
			Handler:    _PluginService_CreatePluginScore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hub.io/api/plugin/v1alpha1/rpc.proto",
}
