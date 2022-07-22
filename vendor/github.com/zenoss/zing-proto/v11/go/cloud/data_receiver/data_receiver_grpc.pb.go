// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: zenoss/cloud/data_receiver.proto

package data_receiver

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

// DataReceiverServiceClient is the client API for DataReceiverService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataReceiverServiceClient interface {
	// Send Events
	PutEvents(ctx context.Context, in *Events, opts ...grpc.CallOption) (*EventStatusResult, error)
	// Stream Events of any type.
	PutEvent(ctx context.Context, opts ...grpc.CallOption) (DataReceiverService_PutEventClient, error)
	// Send Metrics
	PutMetrics(ctx context.Context, in *Metrics, opts ...grpc.CallOption) (*StatusResult, error)
	// Stream Metric of any type
	PutMetric(ctx context.Context, opts ...grpc.CallOption) (DataReceiverService_PutMetricClient, error)
	// Send batch of models
	PutModels(ctx context.Context, in *Models, opts ...grpc.CallOption) (*ModelStatusResult, error)
}

type dataReceiverServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataReceiverServiceClient(cc grpc.ClientConnInterface) DataReceiverServiceClient {
	return &dataReceiverServiceClient{cc}
}

func (c *dataReceiverServiceClient) PutEvents(ctx context.Context, in *Events, opts ...grpc.CallOption) (*EventStatusResult, error) {
	out := new(EventStatusResult)
	err := c.cc.Invoke(ctx, "/zenoss.cloud.DataReceiverService/PutEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataReceiverServiceClient) PutEvent(ctx context.Context, opts ...grpc.CallOption) (DataReceiverService_PutEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataReceiverService_ServiceDesc.Streams[0], "/zenoss.cloud.DataReceiverService/PutEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataReceiverServicePutEventClient{stream}
	return x, nil
}

type DataReceiverService_PutEventClient interface {
	Send(*EventWrapper) error
	CloseAndRecv() (*Void, error)
	grpc.ClientStream
}

type dataReceiverServicePutEventClient struct {
	grpc.ClientStream
}

func (x *dataReceiverServicePutEventClient) Send(m *EventWrapper) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataReceiverServicePutEventClient) CloseAndRecv() (*Void, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Void)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataReceiverServiceClient) PutMetrics(ctx context.Context, in *Metrics, opts ...grpc.CallOption) (*StatusResult, error) {
	out := new(StatusResult)
	err := c.cc.Invoke(ctx, "/zenoss.cloud.DataReceiverService/PutMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataReceiverServiceClient) PutMetric(ctx context.Context, opts ...grpc.CallOption) (DataReceiverService_PutMetricClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataReceiverService_ServiceDesc.Streams[1], "/zenoss.cloud.DataReceiverService/PutMetric", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataReceiverServicePutMetricClient{stream}
	return x, nil
}

type DataReceiverService_PutMetricClient interface {
	Send(*MetricWrapper) error
	CloseAndRecv() (*Void, error)
	grpc.ClientStream
}

type dataReceiverServicePutMetricClient struct {
	grpc.ClientStream
}

func (x *dataReceiverServicePutMetricClient) Send(m *MetricWrapper) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataReceiverServicePutMetricClient) CloseAndRecv() (*Void, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Void)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataReceiverServiceClient) PutModels(ctx context.Context, in *Models, opts ...grpc.CallOption) (*ModelStatusResult, error) {
	out := new(ModelStatusResult)
	err := c.cc.Invoke(ctx, "/zenoss.cloud.DataReceiverService/PutModels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataReceiverServiceServer is the server API for DataReceiverService service.
// All implementations must embed UnimplementedDataReceiverServiceServer
// for forward compatibility
type DataReceiverServiceServer interface {
	// Send Events
	PutEvents(context.Context, *Events) (*EventStatusResult, error)
	// Stream Events of any type.
	PutEvent(DataReceiverService_PutEventServer) error
	// Send Metrics
	PutMetrics(context.Context, *Metrics) (*StatusResult, error)
	// Stream Metric of any type
	PutMetric(DataReceiverService_PutMetricServer) error
	// Send batch of models
	PutModels(context.Context, *Models) (*ModelStatusResult, error)
	mustEmbedUnimplementedDataReceiverServiceServer()
}

// UnimplementedDataReceiverServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataReceiverServiceServer struct {
}

func (UnimplementedDataReceiverServiceServer) PutEvents(context.Context, *Events) (*EventStatusResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutEvents not implemented")
}
func (UnimplementedDataReceiverServiceServer) PutEvent(DataReceiverService_PutEventServer) error {
	return status.Errorf(codes.Unimplemented, "method PutEvent not implemented")
}
func (UnimplementedDataReceiverServiceServer) PutMetrics(context.Context, *Metrics) (*StatusResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutMetrics not implemented")
}
func (UnimplementedDataReceiverServiceServer) PutMetric(DataReceiverService_PutMetricServer) error {
	return status.Errorf(codes.Unimplemented, "method PutMetric not implemented")
}
func (UnimplementedDataReceiverServiceServer) PutModels(context.Context, *Models) (*ModelStatusResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutModels not implemented")
}
func (UnimplementedDataReceiverServiceServer) mustEmbedUnimplementedDataReceiverServiceServer() {}

// UnsafeDataReceiverServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataReceiverServiceServer will
// result in compilation errors.
type UnsafeDataReceiverServiceServer interface {
	mustEmbedUnimplementedDataReceiverServiceServer()
}

func RegisterDataReceiverServiceServer(s grpc.ServiceRegistrar, srv DataReceiverServiceServer) {
	s.RegisterService(&DataReceiverService_ServiceDesc, srv)
}

func _DataReceiverService_PutEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Events)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataReceiverServiceServer).PutEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenoss.cloud.DataReceiverService/PutEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataReceiverServiceServer).PutEvents(ctx, req.(*Events))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataReceiverService_PutEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataReceiverServiceServer).PutEvent(&dataReceiverServicePutEventServer{stream})
}

type DataReceiverService_PutEventServer interface {
	SendAndClose(*Void) error
	Recv() (*EventWrapper, error)
	grpc.ServerStream
}

type dataReceiverServicePutEventServer struct {
	grpc.ServerStream
}

func (x *dataReceiverServicePutEventServer) SendAndClose(m *Void) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataReceiverServicePutEventServer) Recv() (*EventWrapper, error) {
	m := new(EventWrapper)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DataReceiverService_PutMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Metrics)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataReceiverServiceServer).PutMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenoss.cloud.DataReceiverService/PutMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataReceiverServiceServer).PutMetrics(ctx, req.(*Metrics))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataReceiverService_PutMetric_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataReceiverServiceServer).PutMetric(&dataReceiverServicePutMetricServer{stream})
}

type DataReceiverService_PutMetricServer interface {
	SendAndClose(*Void) error
	Recv() (*MetricWrapper, error)
	grpc.ServerStream
}

type dataReceiverServicePutMetricServer struct {
	grpc.ServerStream
}

func (x *dataReceiverServicePutMetricServer) SendAndClose(m *Void) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataReceiverServicePutMetricServer) Recv() (*MetricWrapper, error) {
	m := new(MetricWrapper)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DataReceiverService_PutModels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Models)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataReceiverServiceServer).PutModels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenoss.cloud.DataReceiverService/PutModels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataReceiverServiceServer).PutModels(ctx, req.(*Models))
	}
	return interceptor(ctx, in, info, handler)
}

// DataReceiverService_ServiceDesc is the grpc.ServiceDesc for DataReceiverService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataReceiverService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zenoss.cloud.DataReceiverService",
	HandlerType: (*DataReceiverServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutEvents",
			Handler:    _DataReceiverService_PutEvents_Handler,
		},
		{
			MethodName: "PutMetrics",
			Handler:    _DataReceiverService_PutMetrics_Handler,
		},
		{
			MethodName: "PutModels",
			Handler:    _DataReceiverService_PutModels_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PutEvent",
			Handler:       _DataReceiverService_PutEvent_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "PutMetric",
			Handler:       _DataReceiverService_PutMetric_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "zenoss/cloud/data_receiver.proto",
}
