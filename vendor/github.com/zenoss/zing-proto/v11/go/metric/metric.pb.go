// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zenoss/zing/proto/metric/metric.proto

package metric

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/zenoss/zing-proto/v11/go/cloud/common"
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

// Metric Time-Series struct. Used by OTSDB and RM_5.1 collectors
type MetricTS struct {
	// The metric name
	Metric string `protobuf:"bytes,1,opt,name=metric,proto3" json:"metric,omitempty"`
	// The time at which the value was captured
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The metric value
	Value float64 `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	// Metadata associated with this datapoint.
	Tags map[string]string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The metric tenant id
	Tenant               string   `protobuf:"bytes,5,opt,name=tenant,proto3" json:"tenant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricTS) Reset()         { *m = MetricTS{} }
func (m *MetricTS) String() string { return proto.CompactTextString(m) }
func (*MetricTS) ProtoMessage()    {}
func (*MetricTS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1791175627a454e5, []int{0}
}

func (m *MetricTS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricTS.Unmarshal(m, b)
}
func (m *MetricTS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricTS.Marshal(b, m, deterministic)
}
func (m *MetricTS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricTS.Merge(m, src)
}
func (m *MetricTS) XXX_Size() int {
	return xxx_messageInfo_MetricTS.Size(m)
}
func (m *MetricTS) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricTS.DiscardUnknown(m)
}

var xxx_messageInfo_MetricTS proto.InternalMessageInfo

func (m *MetricTS) GetMetric() string {
	if m != nil {
		return m.Metric
	}
	return ""
}

func (m *MetricTS) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MetricTS) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *MetricTS) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *MetricTS) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

// Compact Metric ZING Time-Series struct. Used for registered metrics
type MetricCTS struct {
	// The time at which the value was captured
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The metric value
	Value float64 `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	// Id of the metric instance that owns this datapoint.
	Id string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// The metric name
	Metric string `protobuf:"bytes,4,opt,name=metric,proto3" json:"metric,omitempty"`
	// The metric tenant id
	Tenant               string   `protobuf:"bytes,5,opt,name=tenant,proto3" json:"tenant,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricCTS) Reset()         { *m = MetricCTS{} }
func (m *MetricCTS) String() string { return proto.CompactTextString(m) }
func (*MetricCTS) ProtoMessage()    {}
func (*MetricCTS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1791175627a454e5, []int{1}
}

func (m *MetricCTS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricCTS.Unmarshal(m, b)
}
func (m *MetricCTS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricCTS.Marshal(b, m, deterministic)
}
func (m *MetricCTS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricCTS.Merge(m, src)
}
func (m *MetricCTS) XXX_Size() int {
	return xxx_messageInfo_MetricCTS.Size(m)
}
func (m *MetricCTS) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricCTS.DiscardUnknown(m)
}

var xxx_messageInfo_MetricCTS proto.InternalMessageInfo

func (m *MetricCTS) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MetricCTS) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *MetricCTS) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MetricCTS) GetMetric() string {
	if m != nil {
		return m.Metric
	}
	return ""
}

func (m *MetricCTS) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

// Metric ZING Time-Series struct. Used for external datastores in ZING
type MetricZTS struct {
	// The metric name
	Metric string `protobuf:"bytes,1,opt,name=metric,proto3" json:"metric,omitempty"`
	// The time at which the value was captured
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The metric value
	Value float64 `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	// Id of the metric instance that owns this datapoint.
	Id                   string   `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricZTS) Reset()         { *m = MetricZTS{} }
func (m *MetricZTS) String() string { return proto.CompactTextString(m) }
func (*MetricZTS) ProtoMessage()    {}
func (*MetricZTS) Descriptor() ([]byte, []int) {
	return fileDescriptor_1791175627a454e5, []int{2}
}

func (m *MetricZTS) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricZTS.Unmarshal(m, b)
}
func (m *MetricZTS) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricZTS.Marshal(b, m, deterministic)
}
func (m *MetricZTS) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricZTS.Merge(m, src)
}
func (m *MetricZTS) XXX_Size() int {
	return xxx_messageInfo_MetricZTS.Size(m)
}
func (m *MetricZTS) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricZTS.DiscardUnknown(m)
}

var xxx_messageInfo_MetricZTS proto.InternalMessageInfo

func (m *MetricZTS) GetMetric() string {
	if m != nil {
		return m.Metric
	}
	return ""
}

func (m *MetricZTS) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *MetricZTS) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *MetricZTS) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// An array of MetricZTS objects
type MetricSeries struct {
	Metrics              []*MetricZTS `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MetricSeries) Reset()         { *m = MetricSeries{} }
func (m *MetricSeries) String() string { return proto.CompactTextString(m) }
func (*MetricSeries) ProtoMessage()    {}
func (*MetricSeries) Descriptor() ([]byte, []int) {
	return fileDescriptor_1791175627a454e5, []int{3}
}

func (m *MetricSeries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricSeries.Unmarshal(m, b)
}
func (m *MetricSeries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricSeries.Marshal(b, m, deterministic)
}
func (m *MetricSeries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricSeries.Merge(m, src)
}
func (m *MetricSeries) XXX_Size() int {
	return xxx_messageInfo_MetricSeries.Size(m)
}
func (m *MetricSeries) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricSeries.DiscardUnknown(m)
}

var xxx_messageInfo_MetricSeries proto.InternalMessageInfo

func (m *MetricSeries) GetMetrics() []*MetricZTS {
	if m != nil {
		return m.Metrics
	}
	return nil
}

// Metric struct used for all metric related processing flows
type Metric struct {
	// The metric name
	Metric string `protobuf:"bytes,1,opt,name=metric,proto3" json:"metric,omitempty"`
	// The time at which the value was captured
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The metric value
	Value float64 `protobuf:"fixed64,3,opt,name=value,proto3" json:"value,omitempty"`
	// Identifies the MetricInstance that owns this datapoint
	Id string `protobuf:"bytes,4,opt,name=id,proto3" json:"id,omitempty"`
	// Dimensions associated with this datapoint.
	Dimensions map[string]string `protobuf:"bytes,5,rep,name=dimensions,proto3" json:"dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The metric tenant id
	Tenant string `protobuf:"bytes,7,opt,name=tenant,proto3" json:"tenant,omitempty"`
	// Metadata associated with this datapoint.
	Metadata             map[string]*common.ScalarArray `protobuf:"bytes,8,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_1791175627a454e5, []int{4}
}

func (m *Metric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metric.Unmarshal(m, b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return xxx_messageInfo_Metric.Size(m)
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

func (m *Metric) GetMetric() string {
	if m != nil {
		return m.Metric
	}
	return ""
}

func (m *Metric) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Metric) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Metric) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Metric) GetDimensions() map[string]string {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *Metric) GetTenant() string {
	if m != nil {
		return m.Tenant
	}
	return ""
}

func (m *Metric) GetMetadata() map[string]*common.ScalarArray {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type MetricWrapper struct {
	// Types that are valid to be assigned to MetricType:
	//	*MetricWrapper_Tagged
	//	*MetricWrapper_Compact
	//	*MetricWrapper_Canonical
	MetricType           isMetricWrapper_MetricType `protobuf_oneof:"metric_type"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *MetricWrapper) Reset()         { *m = MetricWrapper{} }
func (m *MetricWrapper) String() string { return proto.CompactTextString(m) }
func (*MetricWrapper) ProtoMessage()    {}
func (*MetricWrapper) Descriptor() ([]byte, []int) {
	return fileDescriptor_1791175627a454e5, []int{5}
}

func (m *MetricWrapper) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricWrapper.Unmarshal(m, b)
}
func (m *MetricWrapper) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricWrapper.Marshal(b, m, deterministic)
}
func (m *MetricWrapper) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricWrapper.Merge(m, src)
}
func (m *MetricWrapper) XXX_Size() int {
	return xxx_messageInfo_MetricWrapper.Size(m)
}
func (m *MetricWrapper) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricWrapper.DiscardUnknown(m)
}

var xxx_messageInfo_MetricWrapper proto.InternalMessageInfo

type isMetricWrapper_MetricType interface {
	isMetricWrapper_MetricType()
}

type MetricWrapper_Tagged struct {
	Tagged *MetricTS `protobuf:"bytes,1,opt,name=tagged,proto3,oneof"`
}

type MetricWrapper_Compact struct {
	Compact *MetricCTS `protobuf:"bytes,2,opt,name=compact,proto3,oneof"`
}

type MetricWrapper_Canonical struct {
	Canonical *Metric `protobuf:"bytes,3,opt,name=canonical,proto3,oneof"`
}

func (*MetricWrapper_Tagged) isMetricWrapper_MetricType() {}

func (*MetricWrapper_Compact) isMetricWrapper_MetricType() {}

func (*MetricWrapper_Canonical) isMetricWrapper_MetricType() {}

func (m *MetricWrapper) GetMetricType() isMetricWrapper_MetricType {
	if m != nil {
		return m.MetricType
	}
	return nil
}

func (m *MetricWrapper) GetTagged() *MetricTS {
	if x, ok := m.GetMetricType().(*MetricWrapper_Tagged); ok {
		return x.Tagged
	}
	return nil
}

func (m *MetricWrapper) GetCompact() *MetricCTS {
	if x, ok := m.GetMetricType().(*MetricWrapper_Compact); ok {
		return x.Compact
	}
	return nil
}

func (m *MetricWrapper) GetCanonical() *Metric {
	if x, ok := m.GetMetricType().(*MetricWrapper_Canonical); ok {
		return x.Canonical
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*MetricWrapper) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*MetricWrapper_Tagged)(nil),
		(*MetricWrapper_Compact)(nil),
		(*MetricWrapper_Canonical)(nil),
	}
}

func init() {
	proto.RegisterType((*MetricTS)(nil), "metric.MetricTS")
	proto.RegisterMapType((map[string]string)(nil), "metric.MetricTS.TagsEntry")
	proto.RegisterType((*MetricCTS)(nil), "metric.MetricCTS")
	proto.RegisterType((*MetricZTS)(nil), "metric.MetricZTS")
	proto.RegisterType((*MetricSeries)(nil), "metric.MetricSeries")
	proto.RegisterType((*Metric)(nil), "metric.Metric")
	proto.RegisterMapType((map[string]string)(nil), "metric.Metric.DimensionsEntry")
	proto.RegisterMapType((map[string]*common.ScalarArray)(nil), "metric.Metric.MetadataEntry")
	proto.RegisterType((*MetricWrapper)(nil), "metric.MetricWrapper")
}

func init() {
	proto.RegisterFile("zenoss/zing/proto/metric/metric.proto", fileDescriptor_1791175627a454e5)
}

var fileDescriptor_1791175627a454e5 = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xdf, 0x8a, 0xd3, 0x4e,
	0x14, 0xde, 0x49, 0xff, 0xe6, 0xf4, 0xb7, 0xfb, 0xab, 0x83, 0x48, 0x0c, 0x45, 0x4a, 0x50, 0x28,
	0x2b, 0x3b, 0x61, 0xeb, 0x85, 0x8b, 0xa2, 0xe0, 0xae, 0xc2, 0x22, 0x16, 0x24, 0x29, 0x0a, 0xbd,
	0x91, 0xd9, 0x64, 0x88, 0x83, 0xcd, 0x1f, 0x66, 0xa6, 0x0b, 0xdd, 0x1b, 0x1f, 0xc2, 0xa7, 0xf0,
	0x89, 0x7c, 0x1d, 0xe9, 0x4c, 0xd2, 0x26, 0xd5, 0xbd, 0x10, 0xf4, 0xaa, 0x3d, 0x33, 0xdf, 0xf9,
	0xbe, 0xef, 0x9c, 0xcc, 0x39, 0xf0, 0xe8, 0x86, 0x65, 0xb9, 0x94, 0xfe, 0x0d, 0xcf, 0x12, 0xbf,
	0x10, 0xb9, 0xca, 0xfd, 0x94, 0x29, 0xc1, 0xa3, 0xf2, 0x87, 0xe8, 0x33, 0xdc, 0x35, 0x91, 0xfb,
	0xf0, 0x57, 0x78, 0xb4, 0xcc, 0x57, 0xb1, 0x2f, 0x23, 0xba, 0xa4, 0xc2, 0xa0, 0xbd, 0x1f, 0x08,
	0xfa, 0x33, 0x9d, 0x30, 0x0f, 0xf1, 0x3d, 0x28, 0x93, 0x1d, 0x34, 0x46, 0x13, 0x3b, 0x28, 0x23,
	0x3c, 0x02, 0x5b, 0xf1, 0x94, 0x49, 0x45, 0xd3, 0xc2, 0xb1, 0xc6, 0x68, 0xd2, 0x0a, 0x76, 0x07,
	0xf8, 0x2e, 0x74, 0xae, 0xe9, 0x72, 0xc5, 0x9c, 0xd6, 0x18, 0x4d, 0x50, 0x60, 0x02, 0x4c, 0xa0,
	0xad, 0x68, 0x22, 0x9d, 0xf6, 0xb8, 0x35, 0x19, 0x4c, 0x5d, 0x52, 0x7a, 0xac, 0xb4, 0xc8, 0x9c,
	0x26, 0xf2, 0x4d, 0xa6, 0xc4, 0x3a, 0xd0, 0xb8, 0x8d, 0xb6, 0x62, 0x19, 0xcd, 0x94, 0xd3, 0x31,
	0xda, 0x26, 0x72, 0x9f, 0x82, 0xbd, 0x85, 0xe2, 0x21, 0xb4, 0xbe, 0xb0, 0x75, 0xe9, 0x6e, 0xf3,
	0x77, 0x27, 0x6e, 0xe9, 0x33, 0x13, 0x3c, 0xb3, 0xce, 0x90, 0xf7, 0x15, 0x6c, 0x23, 0x76, 0x31,
	0x0f, 0x9b, 0x15, 0xa0, 0x5b, 0x2b, 0xb0, 0xea, 0x15, 0x1c, 0x81, 0xc5, 0x63, 0x5d, 0x94, 0x1d,
	0x58, 0x3c, 0xae, 0x75, 0xa7, 0xdd, 0xe8, 0xce, 0x2d, 0xce, 0xbd, 0xa4, 0x32, 0xb0, 0xf8, 0xcb,
	0xad, 0x35, 0xc6, 0xda, 0x95, 0x31, 0xef, 0x39, 0xfc, 0x67, 0x84, 0x42, 0x26, 0x38, 0x93, 0xf8,
	0x31, 0xf4, 0x0c, 0xbb, 0x74, 0x90, 0xee, 0xfe, 0x9d, 0x66, 0xf7, 0x17, 0xf3, 0x30, 0xa8, 0x10,
	0xde, 0xb7, 0x16, 0x74, 0x67, 0xdb, 0x42, 0xfe, 0x95, 0x47, 0xfc, 0x12, 0x20, 0xe6, 0x29, 0xcb,
	0x24, 0xcf, 0x33, 0xe9, 0x74, 0xb4, 0xad, 0x07, 0x4d, 0x5b, 0xe4, 0xf5, 0x16, 0x60, 0x1e, 0x46,
	0x2d, 0xa3, 0xd6, 0xe4, 0x5e, 0xbd, 0xc9, 0xf8, 0x0c, 0xfa, 0x29, 0x53, 0x34, 0xa6, 0x8a, 0x3a,
	0x7d, 0xcd, 0x3a, 0xda, 0x63, 0x9d, 0x95, 0xd7, 0x86, 0x73, 0x8b, 0x76, 0x5f, 0xc0, 0xff, 0x7b,
	0x82, 0x7f, 0xf2, 0xbc, 0xdc, 0x0f, 0x70, 0xd8, 0x60, 0xfe, 0x4d, 0xb2, 0x5f, 0x4f, 0x1e, 0x4c,
	0xef, 0x13, 0x33, 0x91, 0x44, 0x8f, 0x21, 0x09, 0xf5, 0x18, 0xbe, 0x12, 0x82, 0xae, 0x6b, 0xbc,
	0x6f, 0xdb, 0xfd, 0xee, 0xb0, 0xe7, 0x7d, 0x47, 0x9a, 0x5e, 0xf0, 0xe8, 0xa3, 0xa0, 0x45, 0xc1,
	0x04, 0x3e, 0x86, 0xae, 0xa2, 0x49, 0xc2, 0x62, 0xad, 0x30, 0x98, 0x0e, 0xf7, 0x27, 0xea, 0xf2,
	0x20, 0x28, 0x11, 0xf8, 0x04, 0x7a, 0x51, 0x9e, 0x16, 0x34, 0x52, 0xa5, 0xf4, 0xde, 0x03, 0xb8,
	0xd0, 0xe8, 0x0a, 0x83, 0x09, 0xd8, 0x11, 0xcd, 0xf2, 0x8c, 0x47, 0x74, 0xa9, 0xbf, 0xe2, 0x60,
	0x7a, 0xd4, 0x4c, 0xb8, 0x3c, 0x08, 0x76, 0x90, 0xf3, 0x43, 0x18, 0x98, 0xdb, 0x4f, 0x6a, 0x5d,
	0xb0, 0xf3, 0x77, 0x30, 0xca, 0x45, 0x52, 0x15, 0xb7, 0x59, 0x37, 0x66, 0xb7, 0x94, 0x14, 0xef,
	0xd1, 0xe2, 0x38, 0xe1, 0xea, 0xf3, 0xea, 0x8a, 0x44, 0x79, 0xea, 0xd7, 0xb6, 0xd2, 0x89, 0xd9,
	0x4a, 0xd7, 0xa7, 0xa7, 0x7e, 0x52, 0xed, 0xb2, 0xab, 0xae, 0x3e, 0x7c, 0xf2, 0x33, 0x00, 0x00,
	0xff, 0xff, 0x93, 0x48, 0xe8, 0x54, 0xee, 0x04, 0x00, 0x00,
}
