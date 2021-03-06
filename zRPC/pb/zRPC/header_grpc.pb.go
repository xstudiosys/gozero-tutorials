// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: header.proto

package zRPC

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

// HeaderServiceClient is the client API for HeaderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HeaderServiceClient interface {
	Add(ctx context.Context, in *Header, opts ...grpc.CallOption) (*AddHeaderResponse, error)
	Modify(ctx context.Context, in *Header, opts ...grpc.CallOption) (*ModifyHeaderResponse, error)
	Retrieve(ctx context.Context, in *RetrieveHeaderRequest, opts ...grpc.CallOption) (*Header, error)
}

type headerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHeaderServiceClient(cc grpc.ClientConnInterface) HeaderServiceClient {
	return &headerServiceClient{cc}
}

func (c *headerServiceClient) Add(ctx context.Context, in *Header, opts ...grpc.CallOption) (*AddHeaderResponse, error) {
	out := new(AddHeaderResponse)
	err := c.cc.Invoke(ctx, "/header.HeaderService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headerServiceClient) Modify(ctx context.Context, in *Header, opts ...grpc.CallOption) (*ModifyHeaderResponse, error) {
	out := new(ModifyHeaderResponse)
	err := c.cc.Invoke(ctx, "/header.HeaderService/Modify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *headerServiceClient) Retrieve(ctx context.Context, in *RetrieveHeaderRequest, opts ...grpc.CallOption) (*Header, error) {
	out := new(Header)
	err := c.cc.Invoke(ctx, "/header.HeaderService/Retrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeaderServiceServer is the server API for HeaderService service.
// All implementations must embed UnimplementedHeaderServiceServer
// for forward compatibility
type HeaderServiceServer interface {
	Add(context.Context, *Header) (*AddHeaderResponse, error)
	Modify(context.Context, *Header) (*ModifyHeaderResponse, error)
	Retrieve(context.Context, *RetrieveHeaderRequest) (*Header, error)
	mustEmbedUnimplementedHeaderServiceServer()
}

// UnimplementedHeaderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHeaderServiceServer struct {
}

func (UnimplementedHeaderServiceServer) Add(context.Context, *Header) (*AddHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedHeaderServiceServer) Modify(context.Context, *Header) (*ModifyHeaderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Modify not implemented")
}
func (UnimplementedHeaderServiceServer) Retrieve(context.Context, *RetrieveHeaderRequest) (*Header, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Retrieve not implemented")
}
func (UnimplementedHeaderServiceServer) mustEmbedUnimplementedHeaderServiceServer() {}

// UnsafeHeaderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HeaderServiceServer will
// result in compilation errors.
type UnsafeHeaderServiceServer interface {
	mustEmbedUnimplementedHeaderServiceServer()
}

func RegisterHeaderServiceServer(s grpc.ServiceRegistrar, srv HeaderServiceServer) {
	s.RegisterService(&HeaderService_ServiceDesc, srv)
}

func _HeaderService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Header)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeaderServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/header.HeaderService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeaderServiceServer).Add(ctx, req.(*Header))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeaderService_Modify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Header)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeaderServiceServer).Modify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/header.HeaderService/Modify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeaderServiceServer).Modify(ctx, req.(*Header))
	}
	return interceptor(ctx, in, info, handler)
}

func _HeaderService_Retrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveHeaderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeaderServiceServer).Retrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/header.HeaderService/Retrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeaderServiceServer).Retrieve(ctx, req.(*RetrieveHeaderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HeaderService_ServiceDesc is the grpc.ServiceDesc for HeaderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HeaderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "header.HeaderService",
	HandlerType: (*HeaderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _HeaderService_Add_Handler,
		},
		{
			MethodName: "Modify",
			Handler:    _HeaderService_Modify_Handler,
		},
		{
			MethodName: "Retrieve",
			Handler:    _HeaderService_Retrieve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "header.proto",
}
