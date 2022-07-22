// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: zenoss/zing/proto/cloud/event_management_service.proto

package event_management

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

// EventManagementClient is the client API for EventManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventManagementClient interface {
	//
	// sets status for event(s). Acknowledge and/or status can be updated.
	SetStatus(ctx context.Context, in *EventStatusRequest, opts ...grpc.CallOption) (*EventStatusResponse, error)
	//
	// add or edit annotations for event(s).
	Annotate(ctx context.Context, in *EventAnnotationRequest, opts ...grpc.CallOption) (*EventAnnotationResponse, error)
	//
	// delete annotations for event(s).
	DeleteAnnotations(ctx context.Context, in *EventAnnotationRequest, opts ...grpc.CallOption) (*EventAnnotationResponse, error)
}

type eventManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewEventManagementClient(cc grpc.ClientConnInterface) EventManagementClient {
	return &eventManagementClient{cc}
}

func (c *eventManagementClient) SetStatus(ctx context.Context, in *EventStatusRequest, opts ...grpc.CallOption) (*EventStatusResponse, error) {
	out := new(EventStatusResponse)
	err := c.cc.Invoke(ctx, "/zenoss.cloud.EventManagement/SetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventManagementClient) Annotate(ctx context.Context, in *EventAnnotationRequest, opts ...grpc.CallOption) (*EventAnnotationResponse, error) {
	out := new(EventAnnotationResponse)
	err := c.cc.Invoke(ctx, "/zenoss.cloud.EventManagement/Annotate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventManagementClient) DeleteAnnotations(ctx context.Context, in *EventAnnotationRequest, opts ...grpc.CallOption) (*EventAnnotationResponse, error) {
	out := new(EventAnnotationResponse)
	err := c.cc.Invoke(ctx, "/zenoss.cloud.EventManagement/DeleteAnnotations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventManagementServer is the server API for EventManagement service.
// All implementations must embed UnimplementedEventManagementServer
// for forward compatibility
type EventManagementServer interface {
	//
	// sets status for event(s). Acknowledge and/or status can be updated.
	SetStatus(context.Context, *EventStatusRequest) (*EventStatusResponse, error)
	//
	// add or edit annotations for event(s).
	Annotate(context.Context, *EventAnnotationRequest) (*EventAnnotationResponse, error)
	//
	// delete annotations for event(s).
	DeleteAnnotations(context.Context, *EventAnnotationRequest) (*EventAnnotationResponse, error)
	mustEmbedUnimplementedEventManagementServer()
}

// UnimplementedEventManagementServer must be embedded to have forward compatible implementations.
type UnimplementedEventManagementServer struct {
}

func (UnimplementedEventManagementServer) SetStatus(context.Context, *EventStatusRequest) (*EventStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetStatus not implemented")
}
func (UnimplementedEventManagementServer) Annotate(context.Context, *EventAnnotationRequest) (*EventAnnotationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Annotate not implemented")
}
func (UnimplementedEventManagementServer) DeleteAnnotations(context.Context, *EventAnnotationRequest) (*EventAnnotationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAnnotations not implemented")
}
func (UnimplementedEventManagementServer) mustEmbedUnimplementedEventManagementServer() {}

// UnsafeEventManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventManagementServer will
// result in compilation errors.
type UnsafeEventManagementServer interface {
	mustEmbedUnimplementedEventManagementServer()
}

func RegisterEventManagementServer(s grpc.ServiceRegistrar, srv EventManagementServer) {
	s.RegisterService(&EventManagement_ServiceDesc, srv)
}

func _EventManagement_SetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventManagementServer).SetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenoss.cloud.EventManagement/SetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventManagementServer).SetStatus(ctx, req.(*EventStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventManagement_Annotate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventAnnotationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventManagementServer).Annotate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenoss.cloud.EventManagement/Annotate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventManagementServer).Annotate(ctx, req.(*EventAnnotationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventManagement_DeleteAnnotations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventAnnotationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventManagementServer).DeleteAnnotations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/zenoss.cloud.EventManagement/DeleteAnnotations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventManagementServer).DeleteAnnotations(ctx, req.(*EventAnnotationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventManagement_ServiceDesc is the grpc.ServiceDesc for EventManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "zenoss.cloud.EventManagement",
	HandlerType: (*EventManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetStatus",
			Handler:    _EventManagement_SetStatus_Handler,
		},
		{
			MethodName: "Annotate",
			Handler:    _EventManagement_Annotate_Handler,
		},
		{
			MethodName: "DeleteAnnotations",
			Handler:    _EventManagement_DeleteAnnotations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zenoss/zing/proto/cloud/event_management_service.proto",
}
