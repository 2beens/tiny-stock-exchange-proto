// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: tiny_stock_exchange.proto

package tiny_stock_exchange_proto

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

// TinyStockExchangeClient is the client API for TinyStockExchange service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TinyStockExchangeClient interface {
	NewStock(ctx context.Context, in *Stock, opts ...grpc.CallOption) (*Result, error)
	RemoveStock(ctx context.Context, in *Stock, opts ...grpc.CallOption) (*Result, error)
	UpdateStock(ctx context.Context, in *Stock, opts ...grpc.CallOption) (*Result, error)
	ListStocks(ctx context.Context, in *ListStocksRequest, opts ...grpc.CallOption) (TinyStockExchange_ListStocksClient, error)
	NewValueDelta(ctx context.Context, in *StockValueDelta, opts ...grpc.CallOption) (*Result, error)
	ListStockValueDeltas(ctx context.Context, in *ListStockValueDeltasRequest, opts ...grpc.CallOption) (TinyStockExchange_ListStockValueDeltasClient, error)
}

type tinyStockExchangeClient struct {
	cc grpc.ClientConnInterface
}

func NewTinyStockExchangeClient(cc grpc.ClientConnInterface) TinyStockExchangeClient {
	return &tinyStockExchangeClient{cc}
}

func (c *tinyStockExchangeClient) NewStock(ctx context.Context, in *Stock, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/tiny_stock_exchange.TinyStockExchange/NewStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyStockExchangeClient) RemoveStock(ctx context.Context, in *Stock, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/tiny_stock_exchange.TinyStockExchange/RemoveStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyStockExchangeClient) UpdateStock(ctx context.Context, in *Stock, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/tiny_stock_exchange.TinyStockExchange/UpdateStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyStockExchangeClient) ListStocks(ctx context.Context, in *ListStocksRequest, opts ...grpc.CallOption) (TinyStockExchange_ListStocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &TinyStockExchange_ServiceDesc.Streams[0], "/tiny_stock_exchange.TinyStockExchange/ListStocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &tinyStockExchangeListStocksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TinyStockExchange_ListStocksClient interface {
	Recv() (*Stock, error)
	grpc.ClientStream
}

type tinyStockExchangeListStocksClient struct {
	grpc.ClientStream
}

func (x *tinyStockExchangeListStocksClient) Recv() (*Stock, error) {
	m := new(Stock)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tinyStockExchangeClient) NewValueDelta(ctx context.Context, in *StockValueDelta, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/tiny_stock_exchange.TinyStockExchange/NewValueDelta", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tinyStockExchangeClient) ListStockValueDeltas(ctx context.Context, in *ListStockValueDeltasRequest, opts ...grpc.CallOption) (TinyStockExchange_ListStockValueDeltasClient, error) {
	stream, err := c.cc.NewStream(ctx, &TinyStockExchange_ServiceDesc.Streams[1], "/tiny_stock_exchange.TinyStockExchange/ListStockValueDeltas", opts...)
	if err != nil {
		return nil, err
	}
	x := &tinyStockExchangeListStockValueDeltasClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TinyStockExchange_ListStockValueDeltasClient interface {
	Recv() (*StockValueDelta, error)
	grpc.ClientStream
}

type tinyStockExchangeListStockValueDeltasClient struct {
	grpc.ClientStream
}

func (x *tinyStockExchangeListStockValueDeltasClient) Recv() (*StockValueDelta, error) {
	m := new(StockValueDelta)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TinyStockExchangeServer is the server API for TinyStockExchange service.
// All implementations must embed UnimplementedTinyStockExchangeServer
// for forward compatibility
type TinyStockExchangeServer interface {
	NewStock(context.Context, *Stock) (*Result, error)
	RemoveStock(context.Context, *Stock) (*Result, error)
	UpdateStock(context.Context, *Stock) (*Result, error)
	ListStocks(*ListStocksRequest, TinyStockExchange_ListStocksServer) error
	NewValueDelta(context.Context, *StockValueDelta) (*Result, error)
	ListStockValueDeltas(*ListStockValueDeltasRequest, TinyStockExchange_ListStockValueDeltasServer) error
	mustEmbedUnimplementedTinyStockExchangeServer()
}

// UnimplementedTinyStockExchangeServer must be embedded to have forward compatible implementations.
type UnimplementedTinyStockExchangeServer struct {
}

func (UnimplementedTinyStockExchangeServer) NewStock(context.Context, *Stock) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewStock not implemented")
}
func (UnimplementedTinyStockExchangeServer) RemoveStock(context.Context, *Stock) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveStock not implemented")
}
func (UnimplementedTinyStockExchangeServer) UpdateStock(context.Context, *Stock) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStock not implemented")
}
func (UnimplementedTinyStockExchangeServer) ListStocks(*ListStocksRequest, TinyStockExchange_ListStocksServer) error {
	return status.Errorf(codes.Unimplemented, "method ListStocks not implemented")
}
func (UnimplementedTinyStockExchangeServer) NewValueDelta(context.Context, *StockValueDelta) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewValueDelta not implemented")
}
func (UnimplementedTinyStockExchangeServer) ListStockValueDeltas(*ListStockValueDeltasRequest, TinyStockExchange_ListStockValueDeltasServer) error {
	return status.Errorf(codes.Unimplemented, "method ListStockValueDeltas not implemented")
}
func (UnimplementedTinyStockExchangeServer) mustEmbedUnimplementedTinyStockExchangeServer() {}

// UnsafeTinyStockExchangeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TinyStockExchangeServer will
// result in compilation errors.
type UnsafeTinyStockExchangeServer interface {
	mustEmbedUnimplementedTinyStockExchangeServer()
}

func RegisterTinyStockExchangeServer(s grpc.ServiceRegistrar, srv TinyStockExchangeServer) {
	s.RegisterService(&TinyStockExchange_ServiceDesc, srv)
}

func _TinyStockExchange_NewStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Stock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyStockExchangeServer).NewStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tiny_stock_exchange.TinyStockExchange/NewStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyStockExchangeServer).NewStock(ctx, req.(*Stock))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyStockExchange_RemoveStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Stock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyStockExchangeServer).RemoveStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tiny_stock_exchange.TinyStockExchange/RemoveStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyStockExchangeServer).RemoveStock(ctx, req.(*Stock))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyStockExchange_UpdateStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Stock)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyStockExchangeServer).UpdateStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tiny_stock_exchange.TinyStockExchange/UpdateStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyStockExchangeServer).UpdateStock(ctx, req.(*Stock))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyStockExchange_ListStocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListStocksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TinyStockExchangeServer).ListStocks(m, &tinyStockExchangeListStocksServer{stream})
}

type TinyStockExchange_ListStocksServer interface {
	Send(*Stock) error
	grpc.ServerStream
}

type tinyStockExchangeListStocksServer struct {
	grpc.ServerStream
}

func (x *tinyStockExchangeListStocksServer) Send(m *Stock) error {
	return x.ServerStream.SendMsg(m)
}

func _TinyStockExchange_NewValueDelta_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StockValueDelta)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TinyStockExchangeServer).NewValueDelta(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tiny_stock_exchange.TinyStockExchange/NewValueDelta",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TinyStockExchangeServer).NewValueDelta(ctx, req.(*StockValueDelta))
	}
	return interceptor(ctx, in, info, handler)
}

func _TinyStockExchange_ListStockValueDeltas_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListStockValueDeltasRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TinyStockExchangeServer).ListStockValueDeltas(m, &tinyStockExchangeListStockValueDeltasServer{stream})
}

type TinyStockExchange_ListStockValueDeltasServer interface {
	Send(*StockValueDelta) error
	grpc.ServerStream
}

type tinyStockExchangeListStockValueDeltasServer struct {
	grpc.ServerStream
}

func (x *tinyStockExchangeListStockValueDeltasServer) Send(m *StockValueDelta) error {
	return x.ServerStream.SendMsg(m)
}

// TinyStockExchange_ServiceDesc is the grpc.ServiceDesc for TinyStockExchange service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TinyStockExchange_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tiny_stock_exchange.TinyStockExchange",
	HandlerType: (*TinyStockExchangeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewStock",
			Handler:    _TinyStockExchange_NewStock_Handler,
		},
		{
			MethodName: "RemoveStock",
			Handler:    _TinyStockExchange_RemoveStock_Handler,
		},
		{
			MethodName: "UpdateStock",
			Handler:    _TinyStockExchange_UpdateStock_Handler,
		},
		{
			MethodName: "NewValueDelta",
			Handler:    _TinyStockExchange_NewValueDelta_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListStocks",
			Handler:       _TinyStockExchange_ListStocks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ListStockValueDeltas",
			Handler:       _TinyStockExchange_ListStockValueDeltas_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "tiny_stock_exchange.proto",
}
