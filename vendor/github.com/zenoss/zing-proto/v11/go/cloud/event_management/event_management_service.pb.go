// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zenoss/zing/proto/cloud/event_management_service.proto

package event_management

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type EMStatus int32

const (
	EMStatus_EM_STATUS_DEFAULT    EMStatus = 0
	EMStatus_EM_STATUS_OPEN       EMStatus = 1
	EMStatus_EM_STATUS_SUPPRESSED EMStatus = 2
	EMStatus_EM_STATUS_CLOSED     EMStatus = 3
)

var EMStatus_name = map[int32]string{
	0: "EM_STATUS_DEFAULT",
	1: "EM_STATUS_OPEN",
	2: "EM_STATUS_SUPPRESSED",
	3: "EM_STATUS_CLOSED",
}

var EMStatus_value = map[string]int32{
	"EM_STATUS_DEFAULT":    0,
	"EM_STATUS_OPEN":       1,
	"EM_STATUS_SUPPRESSED": 2,
	"EM_STATUS_CLOSED":     3,
}

func (x EMStatus) String() string {
	return proto.EnumName(EMStatus_name, int32(x))
}

func (EMStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e190cb14ee571f53, []int{0}
}

type EventStatusRequest struct {
	Tenant               string                    `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	StatusList           map[string]*EMEventStatus `protobuf:"bytes,2,rep,name=statusList,proto3" json:"statusList,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *EventStatusRequest) Reset()         { *m = EventStatusRequest{} }
func (m *EventStatusRequest) String() string { return proto.CompactTextString(m) }
func (*EventStatusRequest) ProtoMessage()    {}
func (*EventStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e190cb14ee571f53, []int{0}
}

func (m *EventStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventStatusRequest.Unmarshal(m, b)
}
func (m *EventStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventStatusRequest.Marshal(b, m, deterministic)
}
func (m *EventStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventStatusRequest.Merge(m, src)
}
func (m *EventStatusRequest) XXX_Size() int {
	return xxx_messageInfo_EventStatusRequest.Size(m)
}
func (m *EventStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventStatusRequest proto.InternalMessageInfo

func (m *EventStatusRequest) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *EventStatusRequest) GetStatusList() map[string]*EMEventStatus {
	if m != nil {
		return m.StatusList
	}
	return nil
}

type EMEventStatus struct {
	Acknowledge          bool     `protobuf:"varint,1,opt,name=acknowledge,proto3" json:"acknowledge,omitempty"`
	Status               EMStatus `protobuf:"varint,2,opt,name=status,proto3,enum=zenoss.cloud.EMStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EMEventStatus) Reset()         { *m = EMEventStatus{} }
func (m *EMEventStatus) String() string { return proto.CompactTextString(m) }
func (*EMEventStatus) ProtoMessage()    {}
func (*EMEventStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_e190cb14ee571f53, []int{1}
}

func (m *EMEventStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EMEventStatus.Unmarshal(m, b)
}
func (m *EMEventStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EMEventStatus.Marshal(b, m, deterministic)
}
func (m *EMEventStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EMEventStatus.Merge(m, src)
}
func (m *EMEventStatus) XXX_Size() int {
	return xxx_messageInfo_EMEventStatus.Size(m)
}
func (m *EMEventStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_EMEventStatus.DiscardUnknown(m)
}

var xxx_messageInfo_EMEventStatus proto.InternalMessageInfo

func (m *EMEventStatus) GetAcknowledge() bool {
	if m != nil {
		return m.Acknowledge
	}
	return false
}

func (m *EMEventStatus) GetStatus() EMStatus {
	if m != nil {
		return m.Status
	}
	return EMStatus_EM_STATUS_DEFAULT
}

type EventStatusResponse struct {
	SuccessList          map[string]bool `protobuf:"bytes,2,rep,name=successList,proto3" json:"successList,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *EventStatusResponse) Reset()         { *m = EventStatusResponse{} }
func (m *EventStatusResponse) String() string { return proto.CompactTextString(m) }
func (*EventStatusResponse) ProtoMessage()    {}
func (*EventStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e190cb14ee571f53, []int{2}
}

func (m *EventStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventStatusResponse.Unmarshal(m, b)
}
func (m *EventStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventStatusResponse.Marshal(b, m, deterministic)
}
func (m *EventStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventStatusResponse.Merge(m, src)
}
func (m *EventStatusResponse) XXX_Size() int {
	return xxx_messageInfo_EventStatusResponse.Size(m)
}
func (m *EventStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventStatusResponse proto.InternalMessageInfo

func (m *EventStatusResponse) GetSuccessList() map[string]bool {
	if m != nil {
		return m.SuccessList
	}
	return nil
}

type EventAnnotationRequest struct {
	Tenant               string                 `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	EventId              string                 `protobuf:"bytes,2,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	AnnotationList       map[string]*Annotation `protobuf:"bytes,3,rep,name=annotationList,proto3" json:"annotationList,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *EventAnnotationRequest) Reset()         { *m = EventAnnotationRequest{} }
func (m *EventAnnotationRequest) String() string { return proto.CompactTextString(m) }
func (*EventAnnotationRequest) ProtoMessage()    {}
func (*EventAnnotationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e190cb14ee571f53, []int{3}
}

func (m *EventAnnotationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventAnnotationRequest.Unmarshal(m, b)
}
func (m *EventAnnotationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventAnnotationRequest.Marshal(b, m, deterministic)
}
func (m *EventAnnotationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAnnotationRequest.Merge(m, src)
}
func (m *EventAnnotationRequest) XXX_Size() int {
	return xxx_messageInfo_EventAnnotationRequest.Size(m)
}
func (m *EventAnnotationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAnnotationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EventAnnotationRequest proto.InternalMessageInfo

func (m *EventAnnotationRequest) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *EventAnnotationRequest) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *EventAnnotationRequest) GetAnnotationList() map[string]*Annotation {
	if m != nil {
		return m.AnnotationList
	}
	return nil
}

type Annotation struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Annotation           string   `protobuf:"bytes,2,opt,name=annotation,proto3" json:"annotation,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Annotation) Reset()         { *m = Annotation{} }
func (m *Annotation) String() string { return proto.CompactTextString(m) }
func (*Annotation) ProtoMessage()    {}
func (*Annotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_e190cb14ee571f53, []int{4}
}

func (m *Annotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Annotation.Unmarshal(m, b)
}
func (m *Annotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Annotation.Marshal(b, m, deterministic)
}
func (m *Annotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Annotation.Merge(m, src)
}
func (m *Annotation) XXX_Size() int {
	return xxx_messageInfo_Annotation.Size(m)
}
func (m *Annotation) XXX_DiscardUnknown() {
	xxx_messageInfo_Annotation.DiscardUnknown(m)
}

var xxx_messageInfo_Annotation proto.InternalMessageInfo

func (m *Annotation) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Annotation) GetAnnotation() string {
	if m != nil {
		return m.Annotation
	}
	return ""
}

type EventAnnotationResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	NoteId               string   `protobuf:"bytes,2,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventAnnotationResponse) Reset()         { *m = EventAnnotationResponse{} }
func (m *EventAnnotationResponse) String() string { return proto.CompactTextString(m) }
func (*EventAnnotationResponse) ProtoMessage()    {}
func (*EventAnnotationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e190cb14ee571f53, []int{5}
}

func (m *EventAnnotationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventAnnotationResponse.Unmarshal(m, b)
}
func (m *EventAnnotationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventAnnotationResponse.Marshal(b, m, deterministic)
}
func (m *EventAnnotationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAnnotationResponse.Merge(m, src)
}
func (m *EventAnnotationResponse) XXX_Size() int {
	return xxx_messageInfo_EventAnnotationResponse.Size(m)
}
func (m *EventAnnotationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAnnotationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventAnnotationResponse proto.InternalMessageInfo

func (m *EventAnnotationResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *EventAnnotationResponse) GetNoteId() string {
	if m != nil {
		return m.NoteId
	}
	return ""
}

func init() {
	proto.RegisterEnum("zenoss.cloud.EMStatus", EMStatus_name, EMStatus_value)
	proto.RegisterType((*EventStatusRequest)(nil), "zenoss.cloud.EventStatusRequest")
	proto.RegisterMapType((map[string]*EMEventStatus)(nil), "zenoss.cloud.EventStatusRequest.StatusListEntry")
	proto.RegisterType((*EMEventStatus)(nil), "zenoss.cloud.EMEventStatus")
	proto.RegisterType((*EventStatusResponse)(nil), "zenoss.cloud.EventStatusResponse")
	proto.RegisterMapType((map[string]bool)(nil), "zenoss.cloud.EventStatusResponse.SuccessListEntry")
	proto.RegisterType((*EventAnnotationRequest)(nil), "zenoss.cloud.EventAnnotationRequest")
	proto.RegisterMapType((map[string]*Annotation)(nil), "zenoss.cloud.EventAnnotationRequest.AnnotationListEntry")
	proto.RegisterType((*Annotation)(nil), "zenoss.cloud.Annotation")
	proto.RegisterType((*EventAnnotationResponse)(nil), "zenoss.cloud.EventAnnotationResponse")
}

func init() {
	proto.RegisterFile("zenoss/zing/proto/cloud/event_management_service.proto", fileDescriptor_e190cb14ee571f53)
}

var fileDescriptor_e190cb14ee571f53 = []byte{
	// 582 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xed, 0x8e, 0xd2, 0x40,
	0x14, 0xb5, 0x25, 0xf2, 0x71, 0x51, 0xb6, 0xce, 0x22, 0x5b, 0x31, 0x31, 0x48, 0x34, 0x21, 0x1a,
	0x5b, 0xc1, 0xc4, 0x6c, 0x36, 0xc6, 0x04, 0xdd, 0x9a, 0x6c, 0x02, 0x6e, 0xd3, 0x42, 0x4c, 0xd6,
	0x1f, 0xd8, 0x6d, 0x27, 0xb5, 0x01, 0xa6, 0x2b, 0x9d, 0x62, 0x76, 0x1f, 0xc1, 0x57, 0xf1, 0x15,
	0x7c, 0x12, 0x9f, 0xc6, 0x74, 0x3a, 0xd0, 0x29, 0xa0, 0xf8, 0xaf, 0x33, 0x73, 0xee, 0x39, 0x73,
	0xce, 0xdc, 0x5e, 0x78, 0x7d, 0x83, 0x49, 0x18, 0x45, 0xfa, 0x4d, 0x40, 0x7c, 0xfd, 0x6a, 0x11,
	0xd2, 0x50, 0x77, 0x67, 0x61, 0xec, 0xe9, 0x78, 0x89, 0x09, 0x9d, 0xcc, 0x1d, 0xe2, 0xf8, 0x78,
	0x9e, 0x7c, 0x46, 0x78, 0xb1, 0x0c, 0x5c, 0xac, 0x31, 0x10, 0xba, 0x93, 0xd6, 0x69, 0x0c, 0xdc,
	0xfe, 0x2d, 0x01, 0x32, 0x92, 0x02, 0x9b, 0x3a, 0x34, 0x8e, 0x2c, 0xfc, 0x2d, 0xc6, 0x11, 0x45,
	0x0d, 0x28, 0x52, 0x4c, 0x1c, 0x42, 0x55, 0xa9, 0x25, 0x75, 0x2a, 0x16, 0x5f, 0x21, 0x13, 0x20,
	0x62, 0xc0, 0x41, 0x10, 0x51, 0x55, 0x6e, 0x15, 0x3a, 0xd5, 0xde, 0x4b, 0x4d, 0x64, 0xd4, 0xb6,
	0xd9, 0x34, 0x7b, 0x5d, 0x62, 0x10, 0xba, 0xb8, 0xb6, 0x04, 0x8e, 0xe6, 0x05, 0x1c, 0x6c, 0x1c,
	0x23, 0x05, 0x0a, 0x53, 0x7c, 0xcd, 0x95, 0x93, 0x4f, 0xd4, 0x85, 0xdb, 0x4b, 0x67, 0x16, 0x63,
	0x55, 0x6e, 0x49, 0x9d, 0x6a, 0xef, 0xe1, 0x86, 0xe2, 0x50, 0xd4, 0x4c, 0x91, 0x27, 0xf2, 0xb1,
	0xd4, 0x76, 0xe0, 0x6e, 0xee, 0x0c, 0xb5, 0xa0, 0xea, 0xb8, 0x53, 0x12, 0x7e, 0x9f, 0x61, 0xcf,
	0xc7, 0x4c, 0xa1, 0x6c, 0x89, 0x5b, 0x48, 0x83, 0x62, 0x7a, 0x39, 0x26, 0x55, 0xeb, 0x35, 0x36,
	0xa5, 0xb8, 0x0a, 0x47, 0xb5, 0x7f, 0x4a, 0x70, 0x98, 0x73, 0x1c, 0x5d, 0x85, 0x24, 0xc2, 0x68,
	0x04, 0xd5, 0x28, 0x76, 0x5d, 0x1c, 0x89, 0x49, 0xf5, 0xfe, 0x91, 0x54, 0x5a, 0xa7, 0xd9, 0x59,
	0x51, 0x9a, 0x95, 0x48, 0xd3, 0x7c, 0x0b, 0xca, 0x26, 0x60, 0x47, 0x5a, 0x75, 0x31, 0xad, 0xb2,
	0x18, 0xc8, 0x0f, 0x19, 0x1a, 0x4c, 0xb5, 0x4f, 0x48, 0x48, 0x1d, 0x1a, 0x84, 0x64, 0xdf, 0x8b,
	0x3f, 0x80, 0x72, 0xda, 0x50, 0x81, 0xc7, 0xf8, 0x2a, 0x56, 0x89, 0xad, 0xcf, 0x3c, 0xf4, 0x05,
	0x6a, 0xce, 0x9a, 0x87, 0xd9, 0x2c, 0x30, 0x9b, 0xc7, 0x3b, 0x6c, 0x6e, 0x09, 0x6a, 0xfd, 0x5c,
	0x69, 0x6a, 0x76, 0x83, 0xaf, 0xf9, 0x19, 0x0e, 0x77, 0xc0, 0x76, 0x58, 0xd6, 0xf2, 0x0d, 0xa2,
	0xe6, 0x6f, 0x20, 0x88, 0x0b, 0x61, 0xbc, 0x01, 0xc8, 0x0e, 0x50, 0x0d, 0xe4, 0xc0, 0xe3, 0x94,
	0x72, 0xe0, 0xa1, 0x47, 0x00, 0xd9, 0x65, 0xb8, 0x73, 0x61, 0xa7, 0x3d, 0x80, 0xa3, 0x2d, 0x63,
	0xfc, 0xed, 0x55, 0x28, 0xf1, 0x47, 0xe3, 0x1d, 0xb6, 0x5a, 0xa2, 0x23, 0x28, 0x91, 0x90, 0xe2,
	0x2c, 0xcb, 0x62, 0xb2, 0x3c, 0xf3, 0x9e, 0x61, 0x28, 0xaf, 0x5a, 0x0b, 0xdd, 0x87, 0x7b, 0xc6,
	0x70, 0x62, 0x8f, 0xfa, 0xa3, 0xb1, 0x3d, 0x39, 0x35, 0x3e, 0xf4, 0xc7, 0x83, 0x91, 0x72, 0x0b,
	0x21, 0xa8, 0x65, 0xdb, 0xe7, 0xa6, 0xf1, 0x51, 0x91, 0x90, 0x0a, 0xf5, 0x6c, 0xcf, 0x1e, 0x9b,
	0xa6, 0x65, 0xd8, 0xb6, 0x71, 0xaa, 0xc8, 0xa8, 0x0e, 0x4a, 0x76, 0xf2, 0x7e, 0x70, 0x9e, 0xec,
	0x16, 0x7a, 0xbf, 0x24, 0x38, 0x60, 0xb7, 0x1e, 0xae, 0xa7, 0x03, 0x32, 0xa1, 0x62, 0xe3, 0xf5,
	0x0f, 0xb2, 0xef, 0x5f, 0x6e, 0x3e, 0xde, 0xdb, 0xc3, 0xe8, 0x13, 0x94, 0x79, 0x2a, 0x18, 0x3d,
	0xf9, 0x9f, 0x5e, 0x68, 0x3e, 0xdd, 0x83, 0x4a, 0x89, 0xdf, 0x4d, 0xe1, 0x79, 0xb8, 0xf0, 0x57,
	0xd8, 0x64, 0xf0, 0xa5, 0x33, 0x8d, 0x97, 0xe1, 0xbc, 0x33, 0x53, 0xba, 0x38, 0xf1, 0x03, 0xfa,
	0x35, 0xbe, 0xd4, 0xdc, 0x70, 0xae, 0x0b, 0xe3, 0xf2, 0x45, 0x3a, 0x2e, 0x97, 0xdd, 0xae, 0xee,
	0xff, 0x6d, 0x6a, 0x5e, 0x16, 0x19, 0xe8, 0xd5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x78, 0xa4,
	0xc7, 0xbc, 0x68, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventManagementClient is the client API for EventManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventManagementClient interface {
	SetStatus(ctx context.Context, in *EventStatusRequest, opts ...grpc.CallOption) (*EventStatusResponse, error)
	Annotate(ctx context.Context, in *EventAnnotationRequest, opts ...grpc.CallOption) (*EventAnnotationResponse, error)
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

// EventManagementServer is the server API for EventManagement service.
type EventManagementServer interface {
	SetStatus(context.Context, *EventStatusRequest) (*EventStatusResponse, error)
	Annotate(context.Context, *EventAnnotationRequest) (*EventAnnotationResponse, error)
}

// UnimplementedEventManagementServer can be embedded to have forward compatible implementations.
type UnimplementedEventManagementServer struct {
}

func (*UnimplementedEventManagementServer) SetStatus(ctx context.Context, req *EventStatusRequest) (*EventStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetStatus not implemented")
}
func (*UnimplementedEventManagementServer) Annotate(ctx context.Context, req *EventAnnotationRequest) (*EventAnnotationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Annotate not implemented")
}

func RegisterEventManagementServer(s *grpc.Server, srv EventManagementServer) {
	s.RegisterService(&_EventManagement_serviceDesc, srv)
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

var _EventManagement_serviceDesc = grpc.ServiceDesc{
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zenoss/zing/proto/cloud/event_management_service.proto",
}
