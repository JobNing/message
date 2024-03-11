// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.13.0
// source: goods/goods.proto

//proto版本选择

package goods

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
	Good_GetGood_FullMethodName     = "/goods.Good/GetGood"
	Good_GetGoods_FullMethodName    = "/goods.Good/GetGoods"
	Good_CreateGood_FullMethodName  = "/goods.Good/CreateGood"
	Good_UpdateGood_FullMethodName  = "/goods.Good/UpdateGood"
	Good_DeleteGood_FullMethodName  = "/goods.Good/DeleteGood"
	Good_UpdateStock_FullMethodName = "/goods.Good/UpdateStock"
)

// GoodClient is the client API for Good service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoodClient interface {
	GetGood(ctx context.Context, in *GetGoodRequest, opts ...grpc.CallOption) (*GetGoodResponse, error)
	GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsResponse, error)
	CreateGood(ctx context.Context, in *CreateGoodRequest, opts ...grpc.CallOption) (*CreateGoodResponse, error)
	UpdateGood(ctx context.Context, in *UpdateGoodRequest, opts ...grpc.CallOption) (*UpdateGoodResponse, error)
	DeleteGood(ctx context.Context, in *DeleteGoodRequest, opts ...grpc.CallOption) (*DeleteGoodResponse, error)
	UpdateStock(ctx context.Context, in *UpdateStockRequest, opts ...grpc.CallOption) (*UpdateStockResponse, error)
}

type goodClient struct {
	cc grpc.ClientConnInterface
}

func NewGoodClient(cc grpc.ClientConnInterface) GoodClient {
	return &goodClient{cc}
}

func (c *goodClient) GetGood(ctx context.Context, in *GetGoodRequest, opts ...grpc.CallOption) (*GetGoodResponse, error) {
	out := new(GetGoodResponse)
	err := c.cc.Invoke(ctx, Good_GetGood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodClient) GetGoods(ctx context.Context, in *GetGoodsRequest, opts ...grpc.CallOption) (*GetGoodsResponse, error) {
	out := new(GetGoodsResponse)
	err := c.cc.Invoke(ctx, Good_GetGoods_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodClient) CreateGood(ctx context.Context, in *CreateGoodRequest, opts ...grpc.CallOption) (*CreateGoodResponse, error) {
	out := new(CreateGoodResponse)
	err := c.cc.Invoke(ctx, Good_CreateGood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodClient) UpdateGood(ctx context.Context, in *UpdateGoodRequest, opts ...grpc.CallOption) (*UpdateGoodResponse, error) {
	out := new(UpdateGoodResponse)
	err := c.cc.Invoke(ctx, Good_UpdateGood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodClient) DeleteGood(ctx context.Context, in *DeleteGoodRequest, opts ...grpc.CallOption) (*DeleteGoodResponse, error) {
	out := new(DeleteGoodResponse)
	err := c.cc.Invoke(ctx, Good_DeleteGood_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *goodClient) UpdateStock(ctx context.Context, in *UpdateStockRequest, opts ...grpc.CallOption) (*UpdateStockResponse, error) {
	out := new(UpdateStockResponse)
	err := c.cc.Invoke(ctx, Good_UpdateStock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoodServer is the server API for Good service.
// All implementations must embed UnimplementedGoodServer
// for forward compatibility
type GoodServer interface {
	GetGood(context.Context, *GetGoodRequest) (*GetGoodResponse, error)
	GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsResponse, error)
	CreateGood(context.Context, *CreateGoodRequest) (*CreateGoodResponse, error)
	UpdateGood(context.Context, *UpdateGoodRequest) (*UpdateGoodResponse, error)
	DeleteGood(context.Context, *DeleteGoodRequest) (*DeleteGoodResponse, error)
	UpdateStock(context.Context, *UpdateStockRequest) (*UpdateStockResponse, error)
	mustEmbedUnimplementedGoodServer()
}

// UnimplementedGoodServer must be embedded to have forward compatible implementations.
type UnimplementedGoodServer struct {
}

func (UnimplementedGoodServer) GetGood(context.Context, *GetGoodRequest) (*GetGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGood not implemented")
}
func (UnimplementedGoodServer) GetGoods(context.Context, *GetGoodsRequest) (*GetGoodsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGoods not implemented")
}
func (UnimplementedGoodServer) CreateGood(context.Context, *CreateGoodRequest) (*CreateGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGood not implemented")
}
func (UnimplementedGoodServer) UpdateGood(context.Context, *UpdateGoodRequest) (*UpdateGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGood not implemented")
}
func (UnimplementedGoodServer) DeleteGood(context.Context, *DeleteGoodRequest) (*DeleteGoodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGood not implemented")
}
func (UnimplementedGoodServer) UpdateStock(context.Context, *UpdateStockRequest) (*UpdateStockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStock not implemented")
}
func (UnimplementedGoodServer) mustEmbedUnimplementedGoodServer() {}

// UnsafeGoodServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoodServer will
// result in compilation errors.
type UnsafeGoodServer interface {
	mustEmbedUnimplementedGoodServer()
}

func RegisterGoodServer(s grpc.ServiceRegistrar, srv GoodServer) {
	s.RegisterService(&Good_ServiceDesc, srv)
}

func _Good_GetGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodServer).GetGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Good_GetGood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodServer).GetGood(ctx, req.(*GetGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Good_GetGoods_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGoodsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodServer).GetGoods(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Good_GetGoods_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodServer).GetGoods(ctx, req.(*GetGoodsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Good_CreateGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodServer).CreateGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Good_CreateGood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodServer).CreateGood(ctx, req.(*CreateGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Good_UpdateGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodServer).UpdateGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Good_UpdateGood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodServer).UpdateGood(ctx, req.(*UpdateGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Good_DeleteGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGoodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodServer).DeleteGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Good_DeleteGood_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodServer).DeleteGood(ctx, req.(*DeleteGoodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Good_UpdateStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoodServer).UpdateStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Good_UpdateStock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoodServer).UpdateStock(ctx, req.(*UpdateStockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Good_ServiceDesc is the grpc.ServiceDesc for Good service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Good_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goods.Good",
	HandlerType: (*GoodServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGood",
			Handler:    _Good_GetGood_Handler,
		},
		{
			MethodName: "GetGoods",
			Handler:    _Good_GetGoods_Handler,
		},
		{
			MethodName: "CreateGood",
			Handler:    _Good_CreateGood_Handler,
		},
		{
			MethodName: "UpdateGood",
			Handler:    _Good_UpdateGood_Handler,
		},
		{
			MethodName: "DeleteGood",
			Handler:    _Good_DeleteGood_Handler,
		},
		{
			MethodName: "UpdateStock",
			Handler:    _Good_UpdateStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "goods/goods.proto",
}
