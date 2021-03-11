// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zenoss/zing/proto/cloud/event_management_service.proto

package event_management

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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
	StatusList           map[string]*EMEventStatus `protobuf:"bytes,1,rep,name=status_list,json=statusList,proto3" json:"status_list,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
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

func (m *EventStatusRequest) GetStatusList() map[string]*EMEventStatus {
	if m != nil {
		return m.StatusList
	}
	return nil
}

type EMEventStatus struct {
	EventId              string                 `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	Acknowledge          *wrappers.BoolValue    `protobuf:"bytes,2,opt,name=acknowledge,proto3" json:"acknowledge,omitempty"`
	StatusWrapper        *EMEventStatus_Wrapper `protobuf:"bytes,3,opt,name=status_wrapper,json=statusWrapper,proto3" json:"status_wrapper,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
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

func (m *EMEventStatus) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *EMEventStatus) GetAcknowledge() *wrappers.BoolValue {
	if m != nil {
		return m.Acknowledge
	}
	return nil
}

func (m *EMEventStatus) GetStatusWrapper() *EMEventStatus_Wrapper {
	if m != nil {
		return m.StatusWrapper
	}
	return nil
}

type EMEventStatus_Wrapper struct {
	Status               EMStatus `protobuf:"varint,1,opt,name=status,proto3,enum=zenoss.cloud.EMStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EMEventStatus_Wrapper) Reset()         { *m = EMEventStatus_Wrapper{} }
func (m *EMEventStatus_Wrapper) String() string { return proto.CompactTextString(m) }
func (*EMEventStatus_Wrapper) ProtoMessage()    {}
func (*EMEventStatus_Wrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_e190cb14ee571f53, []int{1, 0}
}

func (m *EMEventStatus_Wrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EMEventStatus_Wrapper.Unmarshal(m, b)
}
func (m *EMEventStatus_Wrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EMEventStatus_Wrapper.Marshal(b, m, deterministic)
}
func (m *EMEventStatus_Wrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EMEventStatus_Wrapper.Merge(m, src)
}
func (m *EMEventStatus_Wrapper) XXX_Size() int {
	return xxx_messageInfo_EMEventStatus_Wrapper.Size(m)
}
func (m *EMEventStatus_Wrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_EMEventStatus_Wrapper.DiscardUnknown(m)
}

var xxx_messageInfo_EMEventStatus_Wrapper proto.InternalMessageInfo

func (m *EMEventStatus_Wrapper) GetStatus() EMStatus {
	if m != nil {
		return m.Status
	}
	return EMStatus_EM_STATUS_DEFAULT
}

type EventStatusResponse struct {
	SuccessList          map[string]bool `protobuf:"bytes,1,rep,name=success_list,json=successList,proto3" json:"success_list,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
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
	AnnotationList       map[string]*Annotation `protobuf:"bytes,1,rep,name=annotation_list,json=annotationList,proto3" json:"annotation_list,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
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

func (m *EventAnnotationRequest) GetAnnotationList() map[string]*Annotation {
	if m != nil {
		return m.AnnotationList
	}
	return nil
}

type Annotation struct {
	EventId              string   `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	AnnotationId         string   `protobuf:"bytes,2,opt,name=annotation_id,json=annotationId,proto3" json:"annotation_id,omitempty"`
	Annotation           string   `protobuf:"bytes,3,opt,name=annotation,proto3" json:"annotation,omitempty"`
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

func (m *Annotation) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *Annotation) GetAnnotationId() string {
	if m != nil {
		return m.AnnotationId
	}
	return ""
}

func (m *Annotation) GetAnnotation() string {
	if m != nil {
		return m.Annotation
	}
	return ""
}

type AnnotationResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	AnnotationId         string   `protobuf:"bytes,2,opt,name=annotation_id,json=annotationId,proto3" json:"annotation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnnotationResponse) Reset()         { *m = AnnotationResponse{} }
func (m *AnnotationResponse) String() string { return proto.CompactTextString(m) }
func (*AnnotationResponse) ProtoMessage()    {}
func (*AnnotationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e190cb14ee571f53, []int{5}
}

func (m *AnnotationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnnotationResponse.Unmarshal(m, b)
}
func (m *AnnotationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnnotationResponse.Marshal(b, m, deterministic)
}
func (m *AnnotationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnnotationResponse.Merge(m, src)
}
func (m *AnnotationResponse) XXX_Size() int {
	return xxx_messageInfo_AnnotationResponse.Size(m)
}
func (m *AnnotationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AnnotationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AnnotationResponse proto.InternalMessageInfo

func (m *AnnotationResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *AnnotationResponse) GetAnnotationId() string {
	if m != nil {
		return m.AnnotationId
	}
	return ""
}

type EventAnnotationResponse struct {
	Success                bool                           `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	AnnotationResponseList map[string]*AnnotationResponse `protobuf:"bytes,2,rep,name=annotation_response_list,json=annotationResponseList,proto3" json:"annotation_response_list,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral   struct{}                       `json:"-"`
	XXX_unrecognized       []byte                         `json:"-"`
	XXX_sizecache          int32                          `json:"-"`
}

func (m *EventAnnotationResponse) Reset()         { *m = EventAnnotationResponse{} }
func (m *EventAnnotationResponse) String() string { return proto.CompactTextString(m) }
func (*EventAnnotationResponse) ProtoMessage()    {}
func (*EventAnnotationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e190cb14ee571f53, []int{6}
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

func (m *EventAnnotationResponse) GetAnnotationResponseList() map[string]*AnnotationResponse {
	if m != nil {
		return m.AnnotationResponseList
	}
	return nil
}

func init() {
	proto.RegisterEnum("zenoss.cloud.EMStatus", EMStatus_name, EMStatus_value)
	proto.RegisterType((*EventStatusRequest)(nil), "zenoss.cloud.EventStatusRequest")
	proto.RegisterMapType((map[string]*EMEventStatus)(nil), "zenoss.cloud.EventStatusRequest.StatusListEntry")
	proto.RegisterType((*EMEventStatus)(nil), "zenoss.cloud.EMEventStatus")
	proto.RegisterType((*EMEventStatus_Wrapper)(nil), "zenoss.cloud.EMEventStatus.Wrapper")
	proto.RegisterType((*EventStatusResponse)(nil), "zenoss.cloud.EventStatusResponse")
	proto.RegisterMapType((map[string]bool)(nil), "zenoss.cloud.EventStatusResponse.SuccessListEntry")
	proto.RegisterType((*EventAnnotationRequest)(nil), "zenoss.cloud.EventAnnotationRequest")
	proto.RegisterMapType((map[string]*Annotation)(nil), "zenoss.cloud.EventAnnotationRequest.AnnotationListEntry")
	proto.RegisterType((*Annotation)(nil), "zenoss.cloud.Annotation")
	proto.RegisterType((*AnnotationResponse)(nil), "zenoss.cloud.AnnotationResponse")
	proto.RegisterType((*EventAnnotationResponse)(nil), "zenoss.cloud.EventAnnotationResponse")
	proto.RegisterMapType((map[string]*AnnotationResponse)(nil), "zenoss.cloud.EventAnnotationResponse.AnnotationResponseListEntry")
}

func init() {
	proto.RegisterFile("zenoss/zing/proto/cloud/event_management_service.proto", fileDescriptor_e190cb14ee571f53)
}

var fileDescriptor_e190cb14ee571f53 = []byte{
	// 691 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x55, 0xef, 0x4e, 0xd3, 0x5e,
	0x18, 0xfe, 0x75, 0xcb, 0x0f, 0xb6, 0x77, 0x30, 0xe6, 0x01, 0xb1, 0x96, 0x84, 0xcc, 0xa1, 0x09,
	0xd1, 0x78, 0x2a, 0x33, 0x21, 0x48, 0x8c, 0xc9, 0x90, 0x9a, 0x60, 0x40, 0x66, 0xcb, 0x24, 0xc1,
	0x0f, 0x4b, 0xd9, 0x8e, 0xb5, 0x59, 0xe9, 0x99, 0x3d, 0xed, 0x08, 0x78, 0x2d, 0xde, 0x81, 0x97,
	0xe0, 0x37, 0x6f, 0xc2, 0xbb, 0xf0, 0x16, 0xcc, 0x7a, 0x4e, 0xe9, 0xe9, 0x36, 0x36, 0xbe, 0x9d,
	0xf3, 0xbc, 0xcf, 0xfb, 0xf7, 0x39, 0x7f, 0x60, 0xfb, 0x9a, 0xf8, 0x94, 0x31, 0xfd, 0xda, 0xf5,
	0x1d, 0xbd, 0x1f, 0xd0, 0x90, 0xea, 0x1d, 0x8f, 0x46, 0x5d, 0x9d, 0x0c, 0x88, 0x1f, 0xb6, 0x2f,
	0x6c, 0xdf, 0x76, 0xc8, 0xc5, 0x70, 0xc9, 0x48, 0x30, 0x70, 0x3b, 0x04, 0xc7, 0x24, 0xb4, 0xc0,
	0xfd, 0x70, 0x4c, 0xd6, 0xd6, 0x1d, 0x4a, 0x1d, 0x8f, 0xf0, 0x00, 0xe7, 0xd1, 0x17, 0xfd, 0x32,
	0xb0, 0xfb, 0x7d, 0x12, 0x30, 0xce, 0xae, 0xfd, 0x56, 0x00, 0x19, 0xc3, 0x80, 0x56, 0x68, 0x87,
	0x11, 0x33, 0xc9, 0xb7, 0x88, 0xb0, 0x10, 0x7d, 0x84, 0x12, 0x8b, 0x81, 0xb6, 0xe7, 0xb2, 0x50,
	0x55, 0xaa, 0xf9, 0xcd, 0x52, 0xfd, 0x05, 0x96, 0x43, 0xe3, 0x71, 0x37, 0xcc, 0x77, 0x87, 0x2e,
	0x0b, 0x0d, 0x3f, 0x0c, 0xae, 0x4c, 0x60, 0x37, 0x80, 0x76, 0x06, 0x4b, 0x23, 0x66, 0x54, 0x81,
	0x7c, 0x8f, 0x5c, 0xa9, 0x4a, 0x55, 0xd9, 0x2c, 0x9a, 0xc3, 0x25, 0xda, 0x82, 0xff, 0x07, 0xb6,
	0x17, 0x11, 0x35, 0x57, 0x55, 0x36, 0x4b, 0xf5, 0xb5, 0x91, 0x8c, 0x47, 0x72, 0x4e, 0xce, 0xdc,
	0xcd, 0xed, 0x28, 0xb5, 0xbf, 0x0a, 0x2c, 0x66, 0x8c, 0xe8, 0x21, 0x14, 0xf8, 0x9c, 0xdc, 0xae,
	0x88, 0x3f, 0x1f, 0xef, 0x0f, 0xba, 0xe8, 0x35, 0x94, 0xec, 0x4e, 0xcf, 0xa7, 0x97, 0x1e, 0xe9,
	0x3a, 0x49, 0x26, 0x0d, 0xf3, 0x41, 0xe1, 0x64, 0x50, 0x78, 0x8f, 0x52, 0xef, 0xd3, 0x30, 0x83,
	0x29, 0xd3, 0xd1, 0x7b, 0x28, 0x8b, 0xc9, 0x88, 0x49, 0xaa, 0xf9, 0x38, 0xc0, 0xc6, 0x94, 0x52,
	0xf1, 0x29, 0xa7, 0x9a, 0x8b, 0xdc, 0x55, 0x6c, 0xb5, 0x57, 0x30, 0x2f, 0x96, 0x08, 0xc3, 0x1c,
	0xb7, 0xc5, 0xd5, 0x96, 0xeb, 0xab, 0xa3, 0xe1, 0x44, 0xd3, 0x82, 0x55, 0xfb, 0xa9, 0xc0, 0x72,
	0x46, 0x00, 0xd6, 0xa7, 0x3e, 0x23, 0xa8, 0x05, 0x0b, 0x2c, 0xea, 0x74, 0x08, 0xcb, 0x28, 0x57,
	0x9f, 0xa2, 0x1c, 0x77, 0xc4, 0x16, 0xf7, 0x4a, 0xb5, 0x2b, 0xb1, 0x14, 0xd1, 0xde, 0x40, 0x65,
	0x94, 0x30, 0x41, 0xbd, 0x15, 0x59, 0xbd, 0x82, 0x2c, 0xd0, 0x1f, 0x05, 0x56, 0xe3, 0xac, 0x0d,
	0xdf, 0xa7, 0xa1, 0x1d, 0xba, 0xd4, 0x4f, 0x8e, 0x9a, 0x0d, 0x4b, 0xf6, 0x0d, 0x28, 0x17, 0xbd,
	0x33, 0xa1, 0xe8, 0x31, 0x77, 0x9c, 0x22, 0x69, 0xe9, 0x65, 0x3b, 0x03, 0x6a, 0x9f, 0x61, 0x79,
	0x02, 0x6d, 0x42, 0x03, 0x38, 0x7b, 0xfc, 0xd4, 0x6c, 0x05, 0x52, 0x72, 0xa9, 0x35, 0x0f, 0x20,
	0x35, 0x4c, 0x3b, 0x77, 0x1b, 0xb0, 0x28, 0x35, 0xea, 0x76, 0xe3, 0x24, 0x45, 0x73, 0x21, 0x05,
	0x0f, 0xba, 0x68, 0x1d, 0x20, 0xdd, 0xc7, 0x47, 0xab, 0x68, 0x4a, 0x48, 0xcd, 0x02, 0x24, 0xcf,
	0x40, 0xa8, 0xae, 0xc2, 0xbc, 0x50, 0x2b, 0x4e, 0x5a, 0x30, 0x93, 0xed, 0x9d, 0x92, 0xd6, 0x7e,
	0xe4, 0xe0, 0xc1, 0xd8, 0x78, 0x67, 0x86, 0xfe, 0x0e, 0xaa, 0x14, 0x3a, 0x10, 0x0e, 0x5c, 0xc1,
	0x5c, 0xac, 0x60, 0x63, 0x86, 0x82, 0xe2, 0xe8, 0x8d, 0x43, 0xa9, 0x94, 0xab, 0xf6, 0x44, 0xa3,
	0xd6, 0x83, 0xb5, 0x29, 0x6e, 0x13, 0xa4, 0xdd, 0xce, 0x4a, 0x5b, 0xbd, 0x55, 0x5a, 0x11, 0x4b,
	0x92, 0xf8, 0x29, 0x81, 0x42, 0x72, 0x01, 0xd1, 0x7d, 0xb8, 0x67, 0x1c, 0xb5, 0xad, 0x93, 0xc6,
	0x49, 0xcb, 0x6a, 0xef, 0x1b, 0xef, 0x1a, 0xad, 0xc3, 0x93, 0xca, 0x7f, 0x08, 0x41, 0x39, 0x85,
	0x8f, 0x9b, 0xc6, 0x87, 0x8a, 0x82, 0x54, 0x58, 0x49, 0x31, 0xab, 0xd5, 0x6c, 0x9a, 0x86, 0x65,
	0x19, 0xfb, 0x95, 0x1c, 0x5a, 0x81, 0x4a, 0x6a, 0x79, 0x7b, 0x78, 0x3c, 0x44, 0xf3, 0xf5, 0x5f,
	0x0a, 0x2c, 0xc5, 0x33, 0x3a, 0xba, 0x79, 0xdb, 0x51, 0x13, 0x8a, 0x16, 0x49, 0x1e, 0xb5, 0xea,
	0xac, 0x07, 0x58, 0x7b, 0x34, 0xf3, 0xa2, 0xa3, 0x53, 0x28, 0x88, 0x6e, 0x09, 0x7a, 0x7c, 0x97,
	0x2b, 0xa6, 0x3d, 0xb9, 0x93, 0x8c, 0x7b, 0x3d, 0x78, 0x46, 0x03, 0x27, 0xe1, 0x0e, 0xbf, 0x2d,
	0xfe, 0x98, 0x0a, 0x37, 0x92, 0xed, 0xac, 0xa9, 0x9c, 0xed, 0x3a, 0x6e, 0xf8, 0x35, 0x3a, 0xc7,
	0x1d, 0x7a, 0xa1, 0x4b, 0x9f, 0xdd, 0x73, 0xfe, 0xd9, 0x0d, 0xb6, 0xb6, 0x74, 0xe7, 0xb6, 0x3f,
	0xef, 0x7c, 0x2e, 0x26, 0xbd, 0xfc, 0x17, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x57, 0x91, 0x22, 0x26,
	0x07, 0x00, 0x00,
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
