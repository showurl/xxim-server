// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: appmgmt.proto

package pb

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

// AppMgmtServiceClient is the client API for AppMgmtService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppMgmtServiceClient interface {
	GetAllAppMgmtConfig(ctx context.Context, in *GetAllAppMgmtConfigReq, opts ...grpc.CallOption) (*GetAllAppMgmtConfigResp, error)
	UpdateAppMgmtConfig(ctx context.Context, in *UpdateAppMgmtConfigReq, opts ...grpc.CallOption) (*UpdateAppMgmtConfigResp, error)
}

type appMgmtServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppMgmtServiceClient(cc grpc.ClientConnInterface) AppMgmtServiceClient {
	return &appMgmtServiceClient{cc}
}

func (c *appMgmtServiceClient) GetAllAppMgmtConfig(ctx context.Context, in *GetAllAppMgmtConfigReq, opts ...grpc.CallOption) (*GetAllAppMgmtConfigResp, error) {
	out := new(GetAllAppMgmtConfigResp)
	err := c.cc.Invoke(ctx, "/pb.appMgmtService/GetAllAppMgmtConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appMgmtServiceClient) UpdateAppMgmtConfig(ctx context.Context, in *UpdateAppMgmtConfigReq, opts ...grpc.CallOption) (*UpdateAppMgmtConfigResp, error) {
	out := new(UpdateAppMgmtConfigResp)
	err := c.cc.Invoke(ctx, "/pb.appMgmtService/UpdateAppMgmtConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppMgmtServiceServer is the server API for AppMgmtService service.
// All implementations must embed UnimplementedAppMgmtServiceServer
// for forward compatibility
type AppMgmtServiceServer interface {
	GetAllAppMgmtConfig(context.Context, *GetAllAppMgmtConfigReq) (*GetAllAppMgmtConfigResp, error)
	UpdateAppMgmtConfig(context.Context, *UpdateAppMgmtConfigReq) (*UpdateAppMgmtConfigResp, error)
	mustEmbedUnimplementedAppMgmtServiceServer()
}

// UnimplementedAppMgmtServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAppMgmtServiceServer struct {
}

func (UnimplementedAppMgmtServiceServer) GetAllAppMgmtConfig(context.Context, *GetAllAppMgmtConfigReq) (*GetAllAppMgmtConfigResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAppMgmtConfig not implemented")
}
func (UnimplementedAppMgmtServiceServer) UpdateAppMgmtConfig(context.Context, *UpdateAppMgmtConfigReq) (*UpdateAppMgmtConfigResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAppMgmtConfig not implemented")
}
func (UnimplementedAppMgmtServiceServer) mustEmbedUnimplementedAppMgmtServiceServer() {}

// UnsafeAppMgmtServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppMgmtServiceServer will
// result in compilation errors.
type UnsafeAppMgmtServiceServer interface {
	mustEmbedUnimplementedAppMgmtServiceServer()
}

func RegisterAppMgmtServiceServer(s grpc.ServiceRegistrar, srv AppMgmtServiceServer) {
	s.RegisterService(&AppMgmtService_ServiceDesc, srv)
}

func _AppMgmtService_GetAllAppMgmtConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllAppMgmtConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppMgmtServiceServer).GetAllAppMgmtConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.appMgmtService/GetAllAppMgmtConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppMgmtServiceServer).GetAllAppMgmtConfig(ctx, req.(*GetAllAppMgmtConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppMgmtService_UpdateAppMgmtConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAppMgmtConfigReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppMgmtServiceServer).UpdateAppMgmtConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.appMgmtService/UpdateAppMgmtConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppMgmtServiceServer).UpdateAppMgmtConfig(ctx, req.(*UpdateAppMgmtConfigReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AppMgmtService_ServiceDesc is the grpc.ServiceDesc for AppMgmtService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppMgmtService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.appMgmtService",
	HandlerType: (*AppMgmtServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllAppMgmtConfig",
			Handler:    _AppMgmtService_GetAllAppMgmtConfig_Handler,
		},
		{
			MethodName: "UpdateAppMgmtConfig",
			Handler:    _AppMgmtService_UpdateAppMgmtConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "appmgmt.proto",
}