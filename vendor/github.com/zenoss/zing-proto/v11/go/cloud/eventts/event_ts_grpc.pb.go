// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: zenoss/zing/proto/cloud/event_ts.proto

package eventts

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

// EventTSServiceClient is the client API for EventTSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventTSServiceClient interface {
	GetEvents(ctx context.Context, in *EventTSRequest, opts ...grpc.CallOption) (*EventTSResponse, error)
	GetEventsStream(ctx context.Context, in *EventTSRequest, opts ...grpc.CallOption) (EventTSService_GetEventsStreamClient, error)
	GetEventCounts(ctx context.Context, in *EventTSCountsRequest, opts ...grpc.CallOption) (*EventTSCountsResponse, error)
	GetEventCountsStream(ctx context.Context, in *EventTSCountsRequest, opts ...grpc.CallOption) (EventTSService_GetEventCountsStreamClient, error)
	GetEventFrequency(ctx context.Context, in *EventTSFrequencyRequest, opts ...grpc.CallOption) (*EventTSFrequencyResponse, error)
	EventsWithCountsStream(ctx context.Context, in *EventsWithCountsRequest, opts ...grpc.CallOption) (EventTSService_EventsWithCountsStreamClient, error)
	GetRawEvents(ctx context.Context, in *EventTSRequest, opts ...grpc.CallOption) (*RawEventsResponse, error)
}

type eventTSServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventTSServiceClient(cc grpc.ClientConnInterface) EventTSServiceClient {
	return &eventTSServiceClient{cc}
}

func (c *eventTSServiceClient) GetEvents(ctx context.Context, in *EventTSRequest, opts ...grpc.CallOption) (*EventTSResponse, error) {
	out := new(EventTSResponse)
	err := c.cc.Invoke(ctx, "/zenoss.cloud.EventTSService/GetEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventTSServiceClient) GetEventsStream(ctx context.Context, in *EventTSRequest, opts ...grpc.CallOption) (EventTSService_GetEventsStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventTSService_ServiceDesc.Streams[0], "/zenoss.cloud.EventTSService/GetEventsStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventTSServiceGetEventsStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventTSService_GetEventsStreamClient interface {
	Recv() (*EventTSResponse, error)
	grpc.ClientStream
}

type eventTSServiceGetEventsStreamClient struct {
	grpc.ClientStream
}

func (x *eventTSServiceGetEventsStreamClient) Recv() (*EventTSResponse, error) {
	m := new(EventTSResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventTSServiceClient) GetEventCounts(ctx context.Context, in *EventTSCountsRequest, opts ...grpc.CallOption) (*EventTSCountsResponse, error) {
	out := new(EventTSCountsResponse)
	err := c.cc.Invoke(ctx, "/zenoss.cloud.EventTSService/GetEventCounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventTSServiceClient) GetEventCountsStream(ctx context.Context, in *EventTSCountsRequest, opts ...grpc.CallOption) (EventTSService_GetEventCountsStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventTSService_ServiceDesc.Streams[1], "/zenoss.cloud.EventTSService/GetEventCountsStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventTSServiceGetEventCountsStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventTSService_GetEventCountsStreamClient interface {
	Recv() (*EventTSCountsResponse, error)
	grpc.ClientStream
}

type eventTSServiceGetEventCountsStreamClient struct {
	grpc.ClientStream
}

func (x *eventTSServiceGetEventCountsStreamClient) Recv() (*EventTSCountsResponse, error) {
	m := new(EventTSCountsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventTSServiceClient) GetEventFrequency(ctx context.Context, in *EventTSFrequencyRequest, opts ...grpc.CallOption) (*EventTSFrequencyResponse, error) {
	out := new(EventTSFrequencyResponse)
	err := c.cc.Invoke(ctx, "/zenoss.cloud.EventTSService/GetEventFrequency", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventTSServiceClient) EventsWithCountsStream(ctx context.Context, in *EventsWithCountsRequest, opts ...grpc.CallOption) (EventTSService_EventsWithCountsStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventTSService_ServiceDesc.Streams[2], "/zenoss.cloud.EventTSService/EventsWithCountsStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventTSServiceEventsWithCountsStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventTSService_EventsWithCountsStreamClient interface {
	Recv() (*EventsWithCountsResponse, error)
	grpc.ClientStream
}

type eventTSServiceEventsWithCountsStreamClient struct {
	grpc.ClientStream
}

func (x *eventTSServiceEventsWithCountsStreamClient) Recv() (*EventsWithCountsResponse, error) {
	m := new(EventsWithCountsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventTSServiceClient) GetRawEvents(ctx context.Context, in *EventTSRequest, opts ...grpc.CallOption) (*RawEventsResponse, error) {
	out := new(RawEventsResponse)
	err := c.cc.Invoke(ctx, "/zenoss.cloud.EventTSService/GetRawEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventTSServiceServer is the server API for EventTSService service.
// All implementations must embed UnimplementedEventTSServiceServer
// for forward compatibility
type EventTSServiceServer interface {
	GetEvents(context.Context, *EventTSRequest) (*EventTSResponse, error)
	GetEventsStream(*EventTSRequest, EventTSService_GetEventsStreamServer) error
	GetEventCounts(context.Context, *EventTSCountsRequest) (*EventTSCountsResponse, error)
	GetEventCountsStream(*EventTSCountsRequest, EventTSService_GetEventCountsStreamServer) error
	GetEventFrequency(context.Context, *EventTSFrequencyRequest) (*EventTSFrequencyResponse, error)
	EventsWithCountsStream(*EventsWithCountsRequest, EventTSService_EventsWithCountsStreamServer) error
	GetRawEvents(context.Context, *EventTSRequest) (*RawEventsResponse, error)
	mustEmbedUnimplementedEventTSServiceServer()
}

// UnimplementedEventTSServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEventTSServiceServer struct {
}

func (UnimplementedEventTSServiceServer) GetEvents(context.Context, *EventTSRequest) (*EventTSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEvents not implemented")
}
func (UnimplementedEventTSServiceServer) GetEventsStream(*EventTSRequest, EventTSService_GetEventsStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEventsStream not implemented")
}
func (UnimplementedEventTSServiceServer) GetEventCounts(context.Context, *EventTSCountsRequest) (*EventTSCountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventCounts not implemented")
}
func (UnimplementedEventTSServiceServer) GetEventCountsStream(*EventTSCountsRequest, EventTSService_GetEventCountsStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetEventCountsStream not implemented")
}
func (UnimplementedEventTSServiceServer) GetEventFrequency(context.Context, *EventTSFrequencyRequest) (*EventTSFrequencyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEventFrequency not implemented")
}
func (UnimplementedEventTSServiceServer) EventsWithCountsStream(*EventsWithCountsRequest, EventTSService_EventsWithCountsStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method EventsWithCountsStream not implemented")
}
func (UnimplementedEventTSServiceServer) GetRawEvents(context.Context, *EventTSRequest) (*RawEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRawEvents not implemented")
}
func (UnimplementedEventTSServiceServer) mustEmbedUnimplementedEventTSServiceServer() {}

// UnsafeEventTSServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventTSServiceServer will
// result in compilation errors.
type UnsafeEventTSServiceServer interface {
	mustEmbedUnimplementedEventTSServiceServer()
}

func RegisterEventTSServiceServer(s grpc.ServiceRegistrar, srv EventTSServiceServer) {
	s.RegisterService(&EventTSService_ServiceDesc, srv)
}

func _EventTSService_GetEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventTSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventTSServiceServer).GetEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenoss.cloud.EventTSService/GetEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventTSServiceServer).GetEvents(ctx, req.(*EventTSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventTSService_GetEventsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventTSRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventTSServiceServer).GetEventsStream(m, &eventTSServiceGetEventsStreamServer{stream})
}

type EventTSService_GetEventsStreamServer interface {
	Send(*EventTSResponse) error
	grpc.ServerStream
}

type eventTSServiceGetEventsStreamServer struct {
	grpc.ServerStream
}

func (x *eventTSServiceGetEventsStreamServer) Send(m *EventTSResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _EventTSService_GetEventCounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventTSCountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventTSServiceServer).GetEventCounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenoss.cloud.EventTSService/GetEventCounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventTSServiceServer).GetEventCounts(ctx, req.(*EventTSCountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventTSService_GetEventCountsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventTSCountsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventTSServiceServer).GetEventCountsStream(m, &eventTSServiceGetEventCountsStreamServer{stream})
}

type EventTSService_GetEventCountsStreamServer interface {
	Send(*EventTSCountsResponse) error
	grpc.ServerStream
}

type eventTSServiceGetEventCountsStreamServer struct {
	grpc.ServerStream
}

func (x *eventTSServiceGetEventCountsStreamServer) Send(m *EventTSCountsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _EventTSService_GetEventFrequency_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventTSFrequencyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventTSServiceServer).GetEventFrequency(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenoss.cloud.EventTSService/GetEventFrequency",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventTSServiceServer).GetEventFrequency(ctx, req.(*EventTSFrequencyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventTSService_EventsWithCountsStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventsWithCountsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventTSServiceServer).EventsWithCountsStream(m, &eventTSServiceEventsWithCountsStreamServer{stream})
}

type EventTSService_EventsWithCountsStreamServer interface {
	Send(*EventsWithCountsResponse) error
	grpc.ServerStream
}

type eventTSServiceEventsWithCountsStreamServer struct {
	grpc.ServerStream
}

func (x *eventTSServiceEventsWithCountsStreamServer) Send(m *EventsWithCountsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _EventTSService_GetRawEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventTSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventTSServiceServer).GetRawEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenoss.cloud.EventTSService/GetRawEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventTSServiceServer).GetRawEvents(ctx, req.(*EventTSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventTSService_ServiceDesc is the grpc.ServiceDesc for EventTSService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventTSService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zenoss.cloud.EventTSService",
	HandlerType: (*EventTSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEvents",
			Handler:    _EventTSService_GetEvents_Handler,
		},
		{
			MethodName: "GetEventCounts",
			Handler:    _EventTSService_GetEventCounts_Handler,
		},
		{
			MethodName: "GetEventFrequency",
			Handler:    _EventTSService_GetEventFrequency_Handler,
		},
		{
			MethodName: "GetRawEvents",
			Handler:    _EventTSService_GetRawEvents_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetEventsStream",
			Handler:       _EventTSService_GetEventsStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetEventCountsStream",
			Handler:       _EventTSService_GetEventCountsStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "EventsWithCountsStream",
			Handler:       _EventTSService_EventsWithCountsStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "zenoss/zing/proto/cloud/event_ts.proto",
}
