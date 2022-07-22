// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: zenoss/zing/proto/cloud/yamr.proto

package yamr

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

// YamrServiceClient is the client API for YamrService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YamrServiceClient interface {
	// Put ingests a single yamr item
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	// PutBulk ingests multiple yamr items
	PutBulk(ctx context.Context, in *PutBulkRequest, opts ...grpc.CallOption) (*PutBulkResponse, error)
	// Search queries for yamr items
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	// StreamingSearch performs a query for yamr items and return a stream of results
	StreamingSearch(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (YamrService_StreamingSearchClient, error)
	// Get returns a single item by its id
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// GetBulk returns multiple items by their ids
	GetBulk(ctx context.Context, in *GetBulkRequest, opts ...grpc.CallOption) (*GetBulkResponse, error)
	// Count returns the number of ids that applies to each value for a time
	Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error)
}

type yamrServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewYamrServiceClient(cc grpc.ClientConnInterface) YamrServiceClient {
	return &yamrServiceClient{cc}
}

func (c *yamrServiceClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/zenoss.cloud.YamrService/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yamrServiceClient) PutBulk(ctx context.Context, in *PutBulkRequest, opts ...grpc.CallOption) (*PutBulkResponse, error) {
	out := new(PutBulkResponse)
	err := c.cc.Invoke(ctx, "/zenoss.cloud.YamrService/PutBulk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yamrServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/zenoss.cloud.YamrService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yamrServiceClient) StreamingSearch(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (YamrService_StreamingSearchClient, error) {
	stream, err := c.cc.NewStream(ctx, &YamrService_ServiceDesc.Streams[0], "/zenoss.cloud.YamrService/StreamingSearch", opts...)
	if err != nil {
		return nil, err
	}
	x := &yamrServiceStreamingSearchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type YamrService_StreamingSearchClient interface {
	Recv() (*SearchResponse, error)
	grpc.ClientStream
}

type yamrServiceStreamingSearchClient struct {
	grpc.ClientStream
}

func (x *yamrServiceStreamingSearchClient) Recv() (*SearchResponse, error) {
	m := new(SearchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *yamrServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/zenoss.cloud.YamrService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yamrServiceClient) GetBulk(ctx context.Context, in *GetBulkRequest, opts ...grpc.CallOption) (*GetBulkResponse, error) {
	out := new(GetBulkResponse)
	err := c.cc.Invoke(ctx, "/zenoss.cloud.YamrService/GetBulk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yamrServiceClient) Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error) {
	out := new(CountResponse)
	err := c.cc.Invoke(ctx, "/zenoss.cloud.YamrService/Count", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YamrServiceServer is the server API for YamrService service.
// All implementations must embed UnimplementedYamrServiceServer
// for forward compatibility
type YamrServiceServer interface {
	// Put ingests a single yamr item
	Put(context.Context, *PutRequest) (*PutResponse, error)
	// PutBulk ingests multiple yamr items
	PutBulk(context.Context, *PutBulkRequest) (*PutBulkResponse, error)
	// Search queries for yamr items
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	// StreamingSearch performs a query for yamr items and return a stream of results
	StreamingSearch(*SearchRequest, YamrService_StreamingSearchServer) error
	// Get returns a single item by its id
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// GetBulk returns multiple items by their ids
	GetBulk(context.Context, *GetBulkRequest) (*GetBulkResponse, error)
	// Count returns the number of ids that applies to each value for a time
	Count(context.Context, *CountRequest) (*CountResponse, error)
	mustEmbedUnimplementedYamrServiceServer()
}

// UnimplementedYamrServiceServer must be embedded to have forward compatible implementations.
type UnimplementedYamrServiceServer struct {
}

func (UnimplementedYamrServiceServer) Put(context.Context, *PutRequest) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedYamrServiceServer) PutBulk(context.Context, *PutBulkRequest) (*PutBulkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutBulk not implemented")
}
func (UnimplementedYamrServiceServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedYamrServiceServer) StreamingSearch(*SearchRequest, YamrService_StreamingSearchServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamingSearch not implemented")
}
func (UnimplementedYamrServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedYamrServiceServer) GetBulk(context.Context, *GetBulkRequest) (*GetBulkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBulk not implemented")
}
func (UnimplementedYamrServiceServer) Count(context.Context, *CountRequest) (*CountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Count not implemented")
}
func (UnimplementedYamrServiceServer) mustEmbedUnimplementedYamrServiceServer() {}

// UnsafeYamrServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YamrServiceServer will
// result in compilation errors.
type UnsafeYamrServiceServer interface {
	mustEmbedUnimplementedYamrServiceServer()
}

func RegisterYamrServiceServer(s grpc.ServiceRegistrar, srv YamrServiceServer) {
	s.RegisterService(&YamrService_ServiceDesc, srv)
}

func _YamrService_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YamrServiceServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenoss.cloud.YamrService/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YamrServiceServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YamrService_PutBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutBulkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YamrServiceServer).PutBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenoss.cloud.YamrService/PutBulk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YamrServiceServer).PutBulk(ctx, req.(*PutBulkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YamrService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YamrServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenoss.cloud.YamrService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YamrServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YamrService_StreamingSearch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(YamrServiceServer).StreamingSearch(m, &yamrServiceStreamingSearchServer{stream})
}

type YamrService_StreamingSearchServer interface {
	Send(*SearchResponse) error
	grpc.ServerStream
}

type yamrServiceStreamingSearchServer struct {
	grpc.ServerStream
}

func (x *yamrServiceStreamingSearchServer) Send(m *SearchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _YamrService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YamrServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenoss.cloud.YamrService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YamrServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YamrService_GetBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBulkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YamrServiceServer).GetBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenoss.cloud.YamrService/GetBulk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YamrServiceServer).GetBulk(ctx, req.(*GetBulkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YamrService_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YamrServiceServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenoss.cloud.YamrService/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YamrServiceServer).Count(ctx, req.(*CountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// YamrService_ServiceDesc is the grpc.ServiceDesc for YamrService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YamrService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zenoss.cloud.YamrService",
	HandlerType: (*YamrServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _YamrService_Put_Handler,
		},
		{
			MethodName: "PutBulk",
			Handler:    _YamrService_PutBulk_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _YamrService_Search_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _YamrService_Get_Handler,
		},
		{
			MethodName: "GetBulk",
			Handler:    _YamrService_GetBulk_Handler,
		},
		{
			MethodName: "Count",
			Handler:    _YamrService_Count_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamingSearch",
			Handler:       _YamrService_StreamingSearch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "zenoss/zing/proto/cloud/yamr.proto",
}

// YamrIngestClient is the client API for YamrIngest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YamrIngestClient interface {
	// Put ingests a single yamr item
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	// PutBulk ingests multiple yamr items
	PutBulk(ctx context.Context, in *PutBulkRequest, opts ...grpc.CallOption) (*PutBulkResponse, error)
}

type yamrIngestClient struct {
	cc grpc.ClientConnInterface
}

func NewYamrIngestClient(cc grpc.ClientConnInterface) YamrIngestClient {
	return &yamrIngestClient{cc}
}

func (c *yamrIngestClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/zenoss.cloud.YamrIngest/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yamrIngestClient) PutBulk(ctx context.Context, in *PutBulkRequest, opts ...grpc.CallOption) (*PutBulkResponse, error) {
	out := new(PutBulkResponse)
	err := c.cc.Invoke(ctx, "/zenoss.cloud.YamrIngest/PutBulk", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YamrIngestServer is the server API for YamrIngest service.
// All implementations must embed UnimplementedYamrIngestServer
// for forward compatibility
type YamrIngestServer interface {
	// Put ingests a single yamr item
	Put(context.Context, *PutRequest) (*PutResponse, error)
	// PutBulk ingests multiple yamr items
	PutBulk(context.Context, *PutBulkRequest) (*PutBulkResponse, error)
	mustEmbedUnimplementedYamrIngestServer()
}

// UnimplementedYamrIngestServer must be embedded to have forward compatible implementations.
type UnimplementedYamrIngestServer struct {
}

func (UnimplementedYamrIngestServer) Put(context.Context, *PutRequest) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedYamrIngestServer) PutBulk(context.Context, *PutBulkRequest) (*PutBulkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutBulk not implemented")
}
func (UnimplementedYamrIngestServer) mustEmbedUnimplementedYamrIngestServer() {}

// UnsafeYamrIngestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YamrIngestServer will
// result in compilation errors.
type UnsafeYamrIngestServer interface {
	mustEmbedUnimplementedYamrIngestServer()
}

func RegisterYamrIngestServer(s grpc.ServiceRegistrar, srv YamrIngestServer) {
	s.RegisterService(&YamrIngest_ServiceDesc, srv)
}

func _YamrIngest_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YamrIngestServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenoss.cloud.YamrIngest/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YamrIngestServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YamrIngest_PutBulk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutBulkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YamrIngestServer).PutBulk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenoss.cloud.YamrIngest/PutBulk",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YamrIngestServer).PutBulk(ctx, req.(*PutBulkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// YamrIngest_ServiceDesc is the grpc.ServiceDesc for YamrIngest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YamrIngest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zenoss.cloud.YamrIngest",
	HandlerType: (*YamrIngestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Put",
			Handler:    _YamrIngest_Put_Handler,
		},
		{
			MethodName: "PutBulk",
			Handler:    _YamrIngest_PutBulk_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zenoss/zing/proto/cloud/yamr.proto",
}
