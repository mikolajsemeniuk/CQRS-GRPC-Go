// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: proto/product.proto

package proto

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

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	ListProducts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (ProductService_ListProductsClient, error)
	ReadProduct(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*Product, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) ListProducts(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (ProductService_ListProductsClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProductService_ServiceDesc.Streams[0], "/proto.ProductService/ListProducts", opts...)
	if err != nil {
		return nil, err
	}
	x := &productServiceListProductsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProductService_ListProductsClient interface {
	Recv() (*Product, error)
	grpc.ClientStream
}

type productServiceListProductsClient struct {
	grpc.ClientStream
}

func (x *productServiceListProductsClient) Recv() (*Product, error) {
	m := new(Product)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *productServiceClient) ReadProduct(ctx context.Context, in *ProductId, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/proto.ProductService/ReadProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations should embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	ListProducts(*emptypb.Empty, ProductService_ListProductsServer) error
	ReadProduct(context.Context, *ProductId) (*Product, error)
}

// UnimplementedProductServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) ListProducts(*emptypb.Empty, ProductService_ListProductsServer) error {
	return status.Errorf(codes.Unimplemented, "method ListProducts not implemented")
}
func (UnimplementedProductServiceServer) ReadProduct(context.Context, *ProductId) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadProduct not implemented")
}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_ListProducts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProductServiceServer).ListProducts(m, &productServiceListProductsServer{stream})
}

type ProductService_ListProductsServer interface {
	Send(*Product) error
	grpc.ServerStream
}

type productServiceListProductsServer struct {
	grpc.ServerStream
}

func (x *productServiceListProductsServer) Send(m *Product) error {
	return x.ServerStream.SendMsg(m)
}

func _ProductService_ReadProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).ReadProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProductService/ReadProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).ReadProduct(ctx, req.(*ProductId))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadProduct",
			Handler:    _ProductService_ReadProduct_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListProducts",
			Handler:       _ProductService_ListProducts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/product.proto",
}