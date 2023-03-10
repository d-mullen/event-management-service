// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: zenoss/zing/proto/cloud/event_query_v2_grpc.proto

package eventquery

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

// EventQueryServiceClient is the client API for EventQueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventQueryServiceClient interface {
	GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*GetEventResponse, error)
	GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error)
	Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error)
	Frequency(ctx context.Context, in *FrequencyRequest, opts ...grpc.CallOption) (*FrequencyResponse, error)
	SearchStream(ctx context.Context, in *SearchStreamRequest, opts ...grpc.CallOption) (EventQueryService_SearchStreamClient, error)
	EventsWithCountsStream(ctx context.Context, in *EventsWithCountsStreamRequest, opts ...grpc.CallOption) (EventQueryService_EventsWithCountsStreamClient, error)
}

type eventQueryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventQueryServiceClient(cc grpc.ClientConnInterface) EventQueryServiceClient {
	return &eventQueryServiceClient{cc}
}

func (c *eventQueryServiceClient) GetEvent(ctx context.Context, in *GetEventRequest, opts ...grpc.CallOption) (*GetEventResponse, error) {
	out := new(GetEventResponse)
	err := c.cc.Invoke(ctx, "/zenoss.cloud.event_query.v2.EventQueryService/GetEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventQueryServiceClient) GetEvents(ctx context.Context, in *GetEventsRequest, opts ...grpc.CallOption) (*GetEventsResponse, error) {
	out := new(GetEventsResponse)
	err := c.cc.Invoke(ctx, "/zenoss.cloud.event_query.v2.EventQueryService/GetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventQueryServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResponse, error) {
	out := new(SearchResponse)
	err := c.cc.Invoke(ctx, "/zenoss.cloud.event_query.v2.EventQueryService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventQueryServiceClient) Count(ctx context.Context, in *CountRequest, opts ...grpc.CallOption) (*CountResponse, error) {
	out := new(CountResponse)
	err := c.cc.Invoke(ctx, "/zenoss.cloud.event_query.v2.EventQueryService/Count", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventQueryServiceClient) Frequency(ctx context.Context, in *FrequencyRequest, opts ...grpc.CallOption) (*FrequencyResponse, error) {
	out := new(FrequencyResponse)
	err := c.cc.Invoke(ctx, "/zenoss.cloud.event_query.v2.EventQueryService/Frequency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventQueryServiceClient) SearchStream(ctx context.Context, in *SearchStreamRequest, opts ...grpc.CallOption) (EventQueryService_SearchStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventQueryService_ServiceDesc.Streams[0], "/zenoss.cloud.event_query.v2.EventQueryService/SearchStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventQueryServiceSearchStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventQueryService_SearchStreamClient interface {
	Recv() (*SearchStreamResponse, error)
	grpc.ClientStream
}

type eventQueryServiceSearchStreamClient struct {
	grpc.ClientStream
}

func (x *eventQueryServiceSearchStreamClient) Recv() (*SearchStreamResponse, error) {
	m := new(SearchStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventQueryServiceClient) EventsWithCountsStream(ctx context.Context, in *EventsWithCountsStreamRequest, opts ...grpc.CallOption) (EventQueryService_EventsWithCountsStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventQueryService_ServiceDesc.Streams[1], "/zenoss.cloud.event_query.v2.EventQueryService/EventsWithCountsStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventQueryServiceEventsWithCountsStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventQueryService_EventsWithCountsStreamClient interface {
	Recv() (*EventsWithCountsStreamResponse, error)
	grpc.ClientStream
}

type eventQueryServiceEventsWithCountsStreamClient struct {
	grpc.ClientStream
}

func (x *eventQueryServiceEventsWithCountsStreamClient) Recv() (*EventsWithCountsStreamResponse, error) {
	m := new(EventsWithCountsStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventQueryServiceServer is the server API for EventQueryService service.
// All implementations must embed UnimplementedEventQueryServiceServer
// for forward compatibility
type EventQueryServiceServer interface {
	GetEvent(context.Context, *GetEventRequest) (*GetEventResponse, error)
	GetEvents(context.Context, *GetEventsRequest) (*GetEventsResponse, error)
	Search(context.Context, *SearchRequest) (*SearchResponse, error)
	Count(context.Context, *CountRequest) (*CountResponse, error)
	Frequency(context.Context, *FrequencyRequest) (*FrequencyResponse, error)
	SearchStream(*SearchStreamRequest, EventQueryService_SearchStreamServer) error
	EventsWithCountsStream(*EventsWithCountsStreamRequest, EventQueryService_EventsWithCountsStreamServer) error
	mustEmbedUnimplementedEventQueryServiceServer()
}

// UnimplementedEventQueryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEventQueryServiceServer struct {
}

func (UnimplementedEventQueryServiceServer) GetEvent(context.Context, *GetEventRequest) (*GetEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvent not implemented")
}
func (UnimplementedEventQueryServiceServer) GetEvents(context.Context, *GetEventsRequest) (*GetEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (UnimplementedEventQueryServiceServer) Search(context.Context, *SearchRequest) (*SearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedEventQueryServiceServer) Count(context.Context, *CountRequest) (*CountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Count not implemented")
}
func (UnimplementedEventQueryServiceServer) Frequency(context.Context, *FrequencyRequest) (*FrequencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Frequency not implemented")
}
func (UnimplementedEventQueryServiceServer) SearchStream(*SearchStreamRequest, EventQueryService_SearchStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SearchStream not implemented")
}
func (UnimplementedEventQueryServiceServer) EventsWithCountsStream(*EventsWithCountsStreamRequest, EventQueryService_EventsWithCountsStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method EventsWithCountsStream not implemented")
}
func (UnimplementedEventQueryServiceServer) mustEmbedUnimplementedEventQueryServiceServer() {}

// UnsafeEventQueryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventQueryServiceServer will
// result in compilation errors.
type UnsafeEventQueryServiceServer interface {
	mustEmbedUnimplementedEventQueryServiceServer()
}

func RegisterEventQueryServiceServer(s grpc.ServiceRegistrar, srv EventQueryServiceServer) {
	s.RegisterService(&EventQueryService_ServiceDesc, srv)
}

func _EventQueryService_GetEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventQueryServiceServer).GetEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenoss.cloud.event_query.v2.EventQueryService/GetEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventQueryServiceServer).GetEvent(ctx, req.(*GetEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventQueryService_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventQueryServiceServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenoss.cloud.event_query.v2.EventQueryService/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventQueryServiceServer).GetEvents(ctx, req.(*GetEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventQueryService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventQueryServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenoss.cloud.event_query.v2.EventQueryService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventQueryServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventQueryService_Count_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventQueryServiceServer).Count(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenoss.cloud.event_query.v2.EventQueryService/Count",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventQueryServiceServer).Count(ctx, req.(*CountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventQueryService_Frequency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FrequencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventQueryServiceServer).Frequency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenoss.cloud.event_query.v2.EventQueryService/Frequency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventQueryServiceServer).Frequency(ctx, req.(*FrequencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventQueryService_SearchStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SearchStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventQueryServiceServer).SearchStream(m, &eventQueryServiceSearchStreamServer{stream})
}

type EventQueryService_SearchStreamServer interface {
	Send(*SearchStreamResponse) error
	grpc.ServerStream
}

type eventQueryServiceSearchStreamServer struct {
	grpc.ServerStream
}

func (x *eventQueryServiceSearchStreamServer) Send(m *SearchStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _EventQueryService_EventsWithCountsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventsWithCountsStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventQueryServiceServer).EventsWithCountsStream(m, &eventQueryServiceEventsWithCountsStreamServer{stream})
}

type EventQueryService_EventsWithCountsStreamServer interface {
	Send(*EventsWithCountsStreamResponse) error
	grpc.ServerStream
}

type eventQueryServiceEventsWithCountsStreamServer struct {
	grpc.ServerStream
}

func (x *eventQueryServiceEventsWithCountsStreamServer) Send(m *EventsWithCountsStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

// EventQueryService_ServiceDesc is the grpc.ServiceDesc for EventQueryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventQueryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zenoss.cloud.event_query.v2.EventQueryService",
	HandlerType: (*EventQueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEvent",
			Handler:    _EventQueryService_GetEvent_Handler,
		},
		{
			MethodName: "GetEvents",
			Handler:    _EventQueryService_GetEvents_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _EventQueryService_Search_Handler,
		},
		{
			MethodName: "Count",
			Handler:    _EventQueryService_Count_Handler,
		},
		{
			MethodName: "Frequency",
			Handler:    _EventQueryService_Frequency_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SearchStream",
			Handler:       _EventQueryService_SearchStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "EventsWithCountsStream",
			Handler:       _EventQueryService_EventsWithCountsStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "zenoss/zing/proto/cloud/event_query_v2_grpc.proto",
}
