// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: admin.proto

package __

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

const (
	AdminService_AddToAdminWallet_FullMethodName = "/pb.AdminService/AddToAdminWallet"
)

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	AddToAdminWallet(ctx context.Context, in *Amount, opts ...grpc.CallOption) (*AdminResponse, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) AddToAdminWallet(ctx context.Context, in *Amount, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, AdminService_AddToAdminWallet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	AddToAdminWallet(context.Context, *Amount) (*AdminResponse, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) AddToAdminWallet(context.Context, *Amount) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToAdminWallet not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_AddToAdminWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Amount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AddToAdminWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_AddToAdminWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AddToAdminWallet(ctx, req.(*Amount))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddToAdminWallet",
			Handler:    _AdminService_AddToAdminWallet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}
