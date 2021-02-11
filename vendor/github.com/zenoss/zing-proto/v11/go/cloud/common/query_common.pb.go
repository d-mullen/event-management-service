// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zenoss/zing/proto/cloud/query_common.proto

package common

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type SortOrder int32

const (
	SortOrder_ASC  SortOrder = 0
	SortOrder_DESC SortOrder = 1
)

var SortOrder_name = map[int32]string{
	0: "ASC",
	1: "DESC",
}

var SortOrder_value = map[string]int32{
	"ASC":  0,
	"DESC": 1,
}

func (x SortOrder) String() string {
	return proto.EnumName(SortOrder_name, int32(x))
}

func (SortOrder) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e39dd0e52a9d3f00, []int{0}
}

type SortBySeries_AggregateOp int32

const (
	SortBySeries_LAST SortBySeries_AggregateOp = 0
	SortBySeries_AVG  SortBySeries_AggregateOp = 1
)

var SortBySeries_AggregateOp_name = map[int32]string{
	0: "LAST",
	1: "AVG",
}

var SortBySeries_AggregateOp_value = map[string]int32{
	"LAST": 0,
	"AVG":  1,
}

func (x SortBySeries_AggregateOp) String() string {
	return proto.EnumName(SortBySeries_AggregateOp_name, int32(x))
}

func (SortBySeries_AggregateOp) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e39dd0e52a9d3f00, []int{5, 0}
}

// Specify pagination for results using cursors.
// Logically, given a result-set for query_id, specifies the open interval
// of results specified by the cursors after and before, sliced by first and
// last.  I.e., query_id:(after..before)[-last:first]
// - At least one of after and before must be specified.
// - If after or before are not specified, the beginning or end (respecively)
//   of the data set are used in their place.
type PagingInput struct {
	// Return results after this curser
	After string `protobuf:"bytes,1,opt,name=after,proto3" json:"after,omitempty"`
	// Return results before this cursor
	Before string `protobuf:"bytes,2,opt,name=before,proto3" json:"before,omitempty"`
	// Return the first N results; only legal in concert with after
	First uint64 `protobuf:"varint,3,opt,name=first,proto3" json:"first,omitempty"`
	// Return the lastt N results; only legal in concert with before
	Last                 uint64   `protobuf:"varint,4,opt,name=last,proto3" json:"last,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PagingInput) Reset()         { *m = PagingInput{} }
func (m *PagingInput) String() string { return proto.CompactTextString(m) }
func (*PagingInput) ProtoMessage()    {}
func (*PagingInput) Descriptor() ([]byte, []int) {
	return fileDescriptor_e39dd0e52a9d3f00, []int{0}
}

func (m *PagingInput) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PagingInput.Unmarshal(m, b)
}
func (m *PagingInput) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PagingInput.Marshal(b, m, deterministic)
}
func (m *PagingInput) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PagingInput.Merge(m, src)
}
func (m *PagingInput) XXX_Size() int {
	return xxx_messageInfo_PagingInput.Size(m)
}
func (m *PagingInput) XXX_DiscardUnknown() {
	xxx_messageInfo_PagingInput.DiscardUnknown(m)
}

var xxx_messageInfo_PagingInput proto.InternalMessageInfo

func (m *PagingInput) GetAfter() string {
	if m != nil {
		return m.After
	}
	return ""
}

func (m *PagingInput) GetBefore() string {
	if m != nil {
		return m.Before
	}
	return ""
}

func (m *PagingInput) GetFirst() uint64 {
	if m != nil {
		return m.First
	}
	return 0
}

func (m *PagingInput) GetLast() uint64 {
	if m != nil {
		return m.Last
	}
	return 0
}

type PageInfo struct {
	// Cursor for the first item in the page; for subsequent use as PagingInput.after
	StartCursor string `protobuf:"bytes,1,opt,name=start_cursor,json=startCursor,proto3" json:"start_cursor,omitempty"`
	// Cursor for the last item in the page; for subsequently use as PagingInput.after
	EndCursor string `protobuf:"bytes,2,opt,name=end_cursor,json=endCursor,proto3" json:"end_cursor,omitempty"`
	// Total number of results returned
	TotalCount uint64 `protobuf:"varint,3,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	// Whether there are results subsequent to this page
	HasNext bool `protobuf:"varint,4,opt,name=has_next,json=hasNext,proto3" json:"has_next,omitempty"`
	// Whether there are results preceding this page
	HasPrev              bool     `protobuf:"varint,5,opt,name=has_prev,json=hasPrev,proto3" json:"has_prev,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageInfo) Reset()         { *m = PageInfo{} }
func (m *PageInfo) String() string { return proto.CompactTextString(m) }
func (*PageInfo) ProtoMessage()    {}
func (*PageInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e39dd0e52a9d3f00, []int{1}
}

func (m *PageInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageInfo.Unmarshal(m, b)
}
func (m *PageInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageInfo.Marshal(b, m, deterministic)
}
func (m *PageInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageInfo.Merge(m, src)
}
func (m *PageInfo) XXX_Size() int {
	return xxx_messageInfo_PageInfo.Size(m)
}
func (m *PageInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PageInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PageInfo proto.InternalMessageInfo

func (m *PageInfo) GetStartCursor() string {
	if m != nil {
		return m.StartCursor
	}
	return ""
}

func (m *PageInfo) GetEndCursor() string {
	if m != nil {
		return m.EndCursor
	}
	return ""
}

func (m *PageInfo) GetTotalCount() uint64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *PageInfo) GetHasNext() bool {
	if m != nil {
		return m.HasNext
	}
	return false
}

func (m *PageInfo) GetHasPrev() bool {
	if m != nil {
		return m.HasPrev
	}
	return false
}

// Fields needed to search on metadata
type MetadataQuery struct {
	// An id to fetch
	Id string `protobuf:"bytes,7,opt,name=id,proto3" json:"id,omitempty"`
	// A string to filter by name
	NameFilter string `protobuf:"bytes,4,opt,name=name_filter,json=nameFilter,proto3" json:"name_filter,omitempty"`
	// Fields we want in our result.  If omitted, results will only include minimal information
	ResultFields []string `protobuf:"bytes,5,rep,name=result_fields,json=resultFields,proto3" json:"result_fields,omitempty"`
	// The time range for this query
	TimeRange *TimeRange `protobuf:"bytes,6,opt,name=time_range,json=timeRange,proto3" json:"time_range,omitempty"`
	// The sort criteria for results
	Sort []*SortBy `protobuf:"bytes,8,rep,name=sort,proto3" json:"sort,omitempty"`
	// If true, will return all fields for each result, regardless of result_fields
	ReturnAllFields bool `protobuf:"varint,9,opt,name=return_all_fields,json=returnAllFields,proto3" json:"return_all_fields,omitempty"`
	// A map of field filters
	Filters map[string]*Scalar `protobuf:"bytes,10,rep,name=filters,proto3" json:"filters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// A map of dimension filters
	DimensionFilters     map[string]*Scalar `protobuf:"bytes,11,rep,name=dimension_filters,json=dimensionFilters,proto3" json:"dimension_filters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MetadataQuery) Reset()         { *m = MetadataQuery{} }
func (m *MetadataQuery) String() string { return proto.CompactTextString(m) }
func (*MetadataQuery) ProtoMessage()    {}
func (*MetadataQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_e39dd0e52a9d3f00, []int{2}
}

func (m *MetadataQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataQuery.Unmarshal(m, b)
}
func (m *MetadataQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataQuery.Marshal(b, m, deterministic)
}
func (m *MetadataQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataQuery.Merge(m, src)
}
func (m *MetadataQuery) XXX_Size() int {
	return xxx_messageInfo_MetadataQuery.Size(m)
}
func (m *MetadataQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataQuery.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataQuery proto.InternalMessageInfo

func (m *MetadataQuery) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MetadataQuery) GetNameFilter() string {
	if m != nil {
		return m.NameFilter
	}
	return ""
}

func (m *MetadataQuery) GetResultFields() []string {
	if m != nil {
		return m.ResultFields
	}
	return nil
}

func (m *MetadataQuery) GetTimeRange() *TimeRange {
	if m != nil {
		return m.TimeRange
	}
	return nil
}

func (m *MetadataQuery) GetSort() []*SortBy {
	if m != nil {
		return m.Sort
	}
	return nil
}

func (m *MetadataQuery) GetReturnAllFields() bool {
	if m != nil {
		return m.ReturnAllFields
	}
	return false
}

func (m *MetadataQuery) GetFilters() map[string]*Scalar {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *MetadataQuery) GetDimensionFilters() map[string]*Scalar {
	if m != nil {
		return m.DimensionFilters
	}
	return nil
}

// The start and end timestamp for a query
type TimeRange struct {
	Start                int64    `protobuf:"varint,1,opt,name=start,proto3" json:"start,omitempty"`
	End                  int64    `protobuf:"varint,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeRange) Reset()         { *m = TimeRange{} }
func (m *TimeRange) String() string { return proto.CompactTextString(m) }
func (*TimeRange) ProtoMessage()    {}
func (*TimeRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_e39dd0e52a9d3f00, []int{3}
}

func (m *TimeRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeRange.Unmarshal(m, b)
}
func (m *TimeRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeRange.Marshal(b, m, deterministic)
}
func (m *TimeRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeRange.Merge(m, src)
}
func (m *TimeRange) XXX_Size() int {
	return xxx_messageInfo_TimeRange.Size(m)
}
func (m *TimeRange) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeRange.DiscardUnknown(m)
}

var xxx_messageInfo_TimeRange proto.InternalMessageInfo

func (m *TimeRange) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *TimeRange) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

// The top-level, dimension or metadata property to sort by
type SortByField struct {
	// Types that are valid to be assigned to SortField:
	//	*SortByField_Property
	//	*SortByField_Dimension
	//	*SortByField_Metadata
	SortField            isSortByField_SortField `protobuf_oneof:"sort_field"`
	Order                SortOrder               `protobuf:"varint,4,opt,name=order,proto3,enum=zenoss.cloud.SortOrder" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SortByField) Reset()         { *m = SortByField{} }
func (m *SortByField) String() string { return proto.CompactTextString(m) }
func (*SortByField) ProtoMessage()    {}
func (*SortByField) Descriptor() ([]byte, []int) {
	return fileDescriptor_e39dd0e52a9d3f00, []int{4}
}

func (m *SortByField) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SortByField.Unmarshal(m, b)
}
func (m *SortByField) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SortByField.Marshal(b, m, deterministic)
}
func (m *SortByField) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SortByField.Merge(m, src)
}
func (m *SortByField) XXX_Size() int {
	return xxx_messageInfo_SortByField.Size(m)
}
func (m *SortByField) XXX_DiscardUnknown() {
	xxx_messageInfo_SortByField.DiscardUnknown(m)
}

var xxx_messageInfo_SortByField proto.InternalMessageInfo

type isSortByField_SortField interface {
	isSortByField_SortField()
}

type SortByField_Property struct {
	Property string `protobuf:"bytes,1,opt,name=property,proto3,oneof"`
}

type SortByField_Dimension struct {
	Dimension string `protobuf:"bytes,2,opt,name=dimension,proto3,oneof"`
}

type SortByField_Metadata struct {
	Metadata string `protobuf:"bytes,3,opt,name=metadata,proto3,oneof"`
}

func (*SortByField_Property) isSortByField_SortField() {}

func (*SortByField_Dimension) isSortByField_SortField() {}

func (*SortByField_Metadata) isSortByField_SortField() {}

func (m *SortByField) GetSortField() isSortByField_SortField {
	if m != nil {
		return m.SortField
	}
	return nil
}

func (m *SortByField) GetProperty() string {
	if x, ok := m.GetSortField().(*SortByField_Property); ok {
		return x.Property
	}
	return ""
}

func (m *SortByField) GetDimension() string {
	if x, ok := m.GetSortField().(*SortByField_Dimension); ok {
		return x.Dimension
	}
	return ""
}

func (m *SortByField) GetMetadata() string {
	if x, ok := m.GetSortField().(*SortByField_Metadata); ok {
		return x.Metadata
	}
	return ""
}

func (m *SortByField) GetOrder() SortOrder {
	if m != nil {
		return m.Order
	}
	return SortOrder_ASC
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SortByField) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SortByField_Property)(nil),
		(*SortByField_Dimension)(nil),
		(*SortByField_Metadata)(nil),
	}
}

// The series name and aggregate operation to sort by
type SortBySeries struct {
	SeriesName           string                   `protobuf:"bytes,1,opt,name=series_name,json=seriesName,proto3" json:"series_name,omitempty"`
	Operation            SortBySeries_AggregateOp `protobuf:"varint,2,opt,name=operation,proto3,enum=zenoss.cloud.SortBySeries_AggregateOp" json:"operation,omitempty"`
	Order                SortOrder                `protobuf:"varint,3,opt,name=order,proto3,enum=zenoss.cloud.SortOrder" json:"order,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *SortBySeries) Reset()         { *m = SortBySeries{} }
func (m *SortBySeries) String() string { return proto.CompactTextString(m) }
func (*SortBySeries) ProtoMessage()    {}
func (*SortBySeries) Descriptor() ([]byte, []int) {
	return fileDescriptor_e39dd0e52a9d3f00, []int{5}
}

func (m *SortBySeries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SortBySeries.Unmarshal(m, b)
}
func (m *SortBySeries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SortBySeries.Marshal(b, m, deterministic)
}
func (m *SortBySeries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SortBySeries.Merge(m, src)
}
func (m *SortBySeries) XXX_Size() int {
	return xxx_messageInfo_SortBySeries.Size(m)
}
func (m *SortBySeries) XXX_DiscardUnknown() {
	xxx_messageInfo_SortBySeries.DiscardUnknown(m)
}

var xxx_messageInfo_SortBySeries proto.InternalMessageInfo

func (m *SortBySeries) GetSeriesName() string {
	if m != nil {
		return m.SeriesName
	}
	return ""
}

func (m *SortBySeries) GetOperation() SortBySeries_AggregateOp {
	if m != nil {
		return m.Operation
	}
	return SortBySeries_LAST
}

func (m *SortBySeries) GetOrder() SortOrder {
	if m != nil {
		return m.Order
	}
	return SortOrder_ASC
}

// Query results can be sorted by a field or metric series
type SortBy struct {
	// Types that are valid to be assigned to SortType:
	//	*SortBy_ByField
	//	*SortBy_BySeries
	SortType             isSortBy_SortType `protobuf_oneof:"sort_type"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SortBy) Reset()         { *m = SortBy{} }
func (m *SortBy) String() string { return proto.CompactTextString(m) }
func (*SortBy) ProtoMessage()    {}
func (*SortBy) Descriptor() ([]byte, []int) {
	return fileDescriptor_e39dd0e52a9d3f00, []int{6}
}

func (m *SortBy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SortBy.Unmarshal(m, b)
}
func (m *SortBy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SortBy.Marshal(b, m, deterministic)
}
func (m *SortBy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SortBy.Merge(m, src)
}
func (m *SortBy) XXX_Size() int {
	return xxx_messageInfo_SortBy.Size(m)
}
func (m *SortBy) XXX_DiscardUnknown() {
	xxx_messageInfo_SortBy.DiscardUnknown(m)
}

var xxx_messageInfo_SortBy proto.InternalMessageInfo

type isSortBy_SortType interface {
	isSortBy_SortType()
}

type SortBy_ByField struct {
	ByField *SortByField `protobuf:"bytes,1,opt,name=by_field,json=byField,proto3,oneof"`
}

type SortBy_BySeries struct {
	BySeries *SortBySeries `protobuf:"bytes,2,opt,name=by_series,json=bySeries,proto3,oneof"`
}

func (*SortBy_ByField) isSortBy_SortType() {}

func (*SortBy_BySeries) isSortBy_SortType() {}

func (m *SortBy) GetSortType() isSortBy_SortType {
	if m != nil {
		return m.SortType
	}
	return nil
}

func (m *SortBy) GetByField() *SortByField {
	if x, ok := m.GetSortType().(*SortBy_ByField); ok {
		return x.ByField
	}
	return nil
}

func (m *SortBy) GetBySeries() *SortBySeries {
	if x, ok := m.GetSortType().(*SortBy_BySeries); ok {
		return x.BySeries
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SortBy) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SortBy_ByField)(nil),
		(*SortBy_BySeries)(nil),
	}
}

type EntityResult struct {
	// The entity id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The entity name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Facts associated with this entity.
	FactIds []string `protobuf:"bytes,3,rep,name=fact_ids,json=factIds,proto3" json:"fact_ids,omitempty"`
	// Schemas associated with this entity.
	SchemaIds []string `protobuf:"bytes,4,rep,name=schema_ids,json=schemaIds,proto3" json:"schema_ids,omitempty"`
	// Metadata associated with this entity.
	Metadata map[string]*ScalarArray `protobuf:"bytes,7,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Dimensions associated with this entity.
	Dimensions           map[string]*Scalar `protobuf:"bytes,8,rep,name=dimensions,proto3" json:"dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *EntityResult) Reset()         { *m = EntityResult{} }
func (m *EntityResult) String() string { return proto.CompactTextString(m) }
func (*EntityResult) ProtoMessage()    {}
func (*EntityResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_e39dd0e52a9d3f00, []int{7}
}

func (m *EntityResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntityResult.Unmarshal(m, b)
}
func (m *EntityResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntityResult.Marshal(b, m, deterministic)
}
func (m *EntityResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityResult.Merge(m, src)
}
func (m *EntityResult) XXX_Size() int {
	return xxx_messageInfo_EntityResult.Size(m)
}
func (m *EntityResult) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityResult.DiscardUnknown(m)
}

var xxx_messageInfo_EntityResult proto.InternalMessageInfo

func (m *EntityResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EntityResult) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EntityResult) GetFactIds() []string {
	if m != nil {
		return m.FactIds
	}
	return nil
}

func (m *EntityResult) GetSchemaIds() []string {
	if m != nil {
		return m.SchemaIds
	}
	return nil
}

func (m *EntityResult) GetMetadata() map[string]*ScalarArray {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *EntityResult) GetDimensions() map[string]*Scalar {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

type EventResult struct {
	// The event id
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The event name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Starttime.
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Metadata associated with this event.
	Metadata map[string]*ScalarArray `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Dimensions associated with this event.
	Dimensions           map[string]*Scalar `protobuf:"bytes,7,rep,name=dimensions,proto3" json:"dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *EventResult) Reset()         { *m = EventResult{} }
func (m *EventResult) String() string { return proto.CompactTextString(m) }
func (*EventResult) ProtoMessage()    {}
func (*EventResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_e39dd0e52a9d3f00, []int{8}
}

func (m *EventResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventResult.Unmarshal(m, b)
}
func (m *EventResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventResult.Marshal(b, m, deterministic)
}
func (m *EventResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventResult.Merge(m, src)
}
func (m *EventResult) XXX_Size() int {
	return xxx_messageInfo_EventResult.Size(m)
}
func (m *EventResult) XXX_DiscardUnknown() {
	xxx_messageInfo_EventResult.DiscardUnknown(m)
}

var xxx_messageInfo_EventResult proto.InternalMessageInfo

func (m *EventResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EventResult) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EventResult) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *EventResult) GetMetadata() map[string]*ScalarArray {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *EventResult) GetDimensions() map[string]*Scalar {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

type MetricResult struct {
	// Identifies the Instance that owns this result
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Dimensions associated with this result.
	Dimensions map[string]*Scalar `protobuf:"bytes,5,rep,name=dimensions,proto3" json:"dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Metadata associated with this result.
	Metadata             map[string]*ScalarArray `protobuf:"bytes,6,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *MetricResult) Reset()         { *m = MetricResult{} }
func (m *MetricResult) String() string { return proto.CompactTextString(m) }
func (*MetricResult) ProtoMessage()    {}
func (*MetricResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_e39dd0e52a9d3f00, []int{9}
}

func (m *MetricResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricResult.Unmarshal(m, b)
}
func (m *MetricResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricResult.Marshal(b, m, deterministic)
}
func (m *MetricResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricResult.Merge(m, src)
}
func (m *MetricResult) XXX_Size() int {
	return xxx_messageInfo_MetricResult.Size(m)
}
func (m *MetricResult) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricResult.DiscardUnknown(m)
}

var xxx_messageInfo_MetricResult proto.InternalMessageInfo

func (m *MetricResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MetricResult) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetricResult) GetDimensions() map[string]*Scalar {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *MetricResult) GetMetadata() map[string]*ScalarArray {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type AnomalyResult struct {
	// Identifies the instance that owns this result
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// The metric id associated with this result
	MetricId string `protobuf:"bytes,3,opt,name=metric_id,json=metricId,proto3" json:"metric_id,omitempty"`
	// Dimensions associated with this result
	Dimensions map[string]*Scalar `protobuf:"bytes,6,rep,name=dimensions,proto3" json:"dimensions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Metadata associated with this result.
	Metadata             map[string]*ScalarArray `protobuf:"bytes,7,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *AnomalyResult) Reset()         { *m = AnomalyResult{} }
func (m *AnomalyResult) String() string { return proto.CompactTextString(m) }
func (*AnomalyResult) ProtoMessage()    {}
func (*AnomalyResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_e39dd0e52a9d3f00, []int{10}
}

func (m *AnomalyResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnomalyResult.Unmarshal(m, b)
}
func (m *AnomalyResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnomalyResult.Marshal(b, m, deterministic)
}
func (m *AnomalyResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnomalyResult.Merge(m, src)
}
func (m *AnomalyResult) XXX_Size() int {
	return xxx_messageInfo_AnomalyResult.Size(m)
}
func (m *AnomalyResult) XXX_DiscardUnknown() {
	xxx_messageInfo_AnomalyResult.DiscardUnknown(m)
}

var xxx_messageInfo_AnomalyResult proto.InternalMessageInfo

func (m *AnomalyResult) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AnomalyResult) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AnomalyResult) GetMetricId() string {
	if m != nil {
		return m.MetricId
	}
	return ""
}

func (m *AnomalyResult) GetDimensions() map[string]*Scalar {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *AnomalyResult) GetMetadata() map[string]*ScalarArray {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type AggregationResult struct {
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Value                *Scalar  `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AggregationResult) Reset()         { *m = AggregationResult{} }
func (m *AggregationResult) String() string { return proto.CompactTextString(m) }
func (*AggregationResult) ProtoMessage()    {}
func (*AggregationResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_e39dd0e52a9d3f00, []int{11}
}

func (m *AggregationResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AggregationResult.Unmarshal(m, b)
}
func (m *AggregationResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AggregationResult.Marshal(b, m, deterministic)
}
func (m *AggregationResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AggregationResult.Merge(m, src)
}
func (m *AggregationResult) XXX_Size() int {
	return xxx_messageInfo_AggregationResult.Size(m)
}
func (m *AggregationResult) XXX_DiscardUnknown() {
	xxx_messageInfo_AggregationResult.DiscardUnknown(m)
}

var xxx_messageInfo_AggregationResult proto.InternalMessageInfo

func (m *AggregationResult) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *AggregationResult) GetValue() *Scalar {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterEnum("zenoss.cloud.SortOrder", SortOrder_name, SortOrder_value)
	proto.RegisterEnum("zenoss.cloud.SortBySeries_AggregateOp", SortBySeries_AggregateOp_name, SortBySeries_AggregateOp_value)
	proto.RegisterType((*PagingInput)(nil), "zenoss.cloud.PagingInput")
	proto.RegisterType((*PageInfo)(nil), "zenoss.cloud.PageInfo")
	proto.RegisterType((*MetadataQuery)(nil), "zenoss.cloud.MetadataQuery")
	proto.RegisterMapType((map[string]*Scalar)(nil), "zenoss.cloud.MetadataQuery.DimensionFiltersEntry")
	proto.RegisterMapType((map[string]*Scalar)(nil), "zenoss.cloud.MetadataQuery.FiltersEntry")
	proto.RegisterType((*TimeRange)(nil), "zenoss.cloud.TimeRange")
	proto.RegisterType((*SortByField)(nil), "zenoss.cloud.SortByField")
	proto.RegisterType((*SortBySeries)(nil), "zenoss.cloud.SortBySeries")
	proto.RegisterType((*SortBy)(nil), "zenoss.cloud.SortBy")
	proto.RegisterType((*EntityResult)(nil), "zenoss.cloud.EntityResult")
	proto.RegisterMapType((map[string]*Scalar)(nil), "zenoss.cloud.EntityResult.DimensionsEntry")
	proto.RegisterMapType((map[string]*ScalarArray)(nil), "zenoss.cloud.EntityResult.MetadataEntry")
	proto.RegisterType((*EventResult)(nil), "zenoss.cloud.EventResult")
	proto.RegisterMapType((map[string]*Scalar)(nil), "zenoss.cloud.EventResult.DimensionsEntry")
	proto.RegisterMapType((map[string]*ScalarArray)(nil), "zenoss.cloud.EventResult.MetadataEntry")
	proto.RegisterType((*MetricResult)(nil), "zenoss.cloud.MetricResult")
	proto.RegisterMapType((map[string]*Scalar)(nil), "zenoss.cloud.MetricResult.DimensionsEntry")
	proto.RegisterMapType((map[string]*ScalarArray)(nil), "zenoss.cloud.MetricResult.MetadataEntry")
	proto.RegisterType((*AnomalyResult)(nil), "zenoss.cloud.AnomalyResult")
	proto.RegisterMapType((map[string]*Scalar)(nil), "zenoss.cloud.AnomalyResult.DimensionsEntry")
	proto.RegisterMapType((map[string]*ScalarArray)(nil), "zenoss.cloud.AnomalyResult.MetadataEntry")
	proto.RegisterType((*AggregationResult)(nil), "zenoss.cloud.AggregationResult")
}

func init() {
	proto.RegisterFile("zenoss/zing/proto/cloud/query_common.proto", fileDescriptor_e39dd0e52a9d3f00)
}

var fileDescriptor_e39dd0e52a9d3f00 = []byte{
	// 1140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x57, 0xdd, 0x6e, 0xe3, 0xc4,
	0x17, 0xaf, 0x63, 0xe7, 0xc3, 0xc7, 0x69, 0xd7, 0x1d, 0xf5, 0xff, 0xc7, 0x2d, 0xcb, 0x12, 0x0c,
	0x82, 0x6c, 0xd1, 0x26, 0xb4, 0x2b, 0xad, 0x80, 0xbb, 0xf4, 0x63, 0x69, 0x0b, 0xbb, 0x5b, 0x9c,
	0xd5, 0x8a, 0x0f, 0x09, 0x6b, 0x12, 0x4f, 0x5c, 0x0b, 0x7f, 0x84, 0x99, 0x49, 0xb5, 0xd9, 0x27,
	0xe0, 0x11, 0xb8, 0xe1, 0x0a, 0x89, 0x57, 0xe0, 0x0d, 0xb8, 0xe6, 0x85, 0x90, 0xd0, 0xcc, 0xd8,
	0x49, 0xdc, 0x4d, 0x0a, 0x15, 0xec, 0xcd, 0xde, 0x8d, 0xcf, 0x99, 0xf3, 0x3b, 0x67, 0xce, 0xef,
	0x9c, 0x33, 0x1e, 0xd8, 0x7d, 0x41, 0xd2, 0x8c, 0xb1, 0xee, 0x8b, 0x28, 0x0d, 0xbb, 0x63, 0x9a,
	0xf1, 0xac, 0x3b, 0x8c, 0xb3, 0x49, 0xd0, 0xfd, 0x61, 0x42, 0xe8, 0xd4, 0x1f, 0x66, 0x49, 0x92,
	0xa5, 0x1d, 0xa9, 0x40, 0x4d, 0xb5, 0xb7, 0x23, 0x37, 0xec, 0xbc, 0xb7, 0xca, 0x92, 0x0d, 0x71,
	0x8c, 0xa9, 0xb2, 0x71, 0x09, 0x58, 0xe7, 0x38, 0x8c, 0xd2, 0xf0, 0x34, 0x1d, 0x4f, 0x38, 0xda,
	0x82, 0x2a, 0x1e, 0x71, 0x42, 0x1d, 0xad, 0xa5, 0xb5, 0x4d, 0x4f, 0x7d, 0xa0, 0xff, 0x43, 0x6d,
	0x40, 0x46, 0x19, 0x25, 0x4e, 0x45, 0x8a, 0xf3, 0x2f, 0xb1, 0x7b, 0x14, 0x51, 0xc6, 0x1d, 0xbd,
	0xa5, 0xb5, 0x0d, 0x4f, 0x7d, 0x20, 0x04, 0x46, 0x8c, 0x19, 0x77, 0x0c, 0x29, 0x94, 0x6b, 0xf7,
	0x17, 0x0d, 0x1a, 0xe7, 0x38, 0x24, 0xa7, 0xe9, 0x28, 0x43, 0xef, 0x40, 0x93, 0x71, 0x4c, 0xb9,
	0x3f, 0x9c, 0x50, 0x96, 0x15, 0xbe, 0x2c, 0x29, 0x3b, 0x94, 0x22, 0xf4, 0x16, 0x00, 0x49, 0x83,
	0x62, 0x83, 0xf2, 0x6a, 0x92, 0x34, 0xc8, 0xd5, 0x6f, 0x83, 0xc5, 0x33, 0x8e, 0x63, 0x7f, 0x98,
	0x4d, 0xd2, 0xc2, 0x3d, 0x48, 0xd1, 0xa1, 0x90, 0xa0, 0x6d, 0x68, 0x5c, 0x60, 0xe6, 0xa7, 0xe4,
	0xb9, 0x8a, 0xa3, 0xe1, 0xd5, 0x2f, 0x30, 0x7b, 0x4c, 0x9e, 0xcf, 0x54, 0x63, 0x4a, 0x2e, 0x9d,
	0xea, 0x4c, 0x75, 0x4e, 0xc9, 0xa5, 0xfb, 0x87, 0x01, 0xeb, 0x8f, 0x08, 0xc7, 0x01, 0xe6, 0xf8,
	0x4b, 0x91, 0x5f, 0xb4, 0x01, 0x95, 0x28, 0x70, 0xea, 0xd2, 0x7f, 0x25, 0x0a, 0x84, 0xe3, 0x14,
	0x27, 0xc4, 0x1f, 0x45, 0xb1, 0xc8, 0x92, 0x21, 0x15, 0x20, 0x44, 0x0f, 0xa5, 0x04, 0xbd, 0x0b,
	0xeb, 0x94, 0xb0, 0x49, 0xcc, 0xfd, 0x51, 0x44, 0xe2, 0x80, 0x39, 0xd5, 0x96, 0xde, 0x36, 0xbd,
	0xa6, 0x12, 0x3e, 0x94, 0x32, 0xf4, 0x00, 0x80, 0x47, 0x09, 0xf1, 0x29, 0x4e, 0x43, 0xe2, 0xd4,
	0x5a, 0x5a, 0xdb, 0xda, 0x7f, 0xa3, 0xb3, 0xc8, 0x5e, 0xe7, 0x69, 0x94, 0x10, 0x4f, 0xa8, 0x3d,
	0x93, 0x17, 0x4b, 0xd4, 0x06, 0x83, 0x65, 0x94, 0x3b, 0x8d, 0x96, 0xde, 0xb6, 0xf6, 0xb7, 0xca,
	0x16, 0xfd, 0x8c, 0xf2, 0x83, 0xa9, 0x27, 0x77, 0xa0, 0x5d, 0xd8, 0xa4, 0x84, 0x4f, 0x68, 0xea,
	0xe3, 0x38, 0x2e, 0x42, 0x31, 0xe5, 0x69, 0x6f, 0x29, 0x45, 0x2f, 0x8e, 0xf3, 0x68, 0x0e, 0xa0,
	0xae, 0x8e, 0xc3, 0x1c, 0x90, 0xc0, 0xed, 0x32, 0x70, 0x29, 0x23, 0x1d, 0x75, 0x4e, 0x76, 0x9c,
	0x72, 0x3a, 0xf5, 0x0a, 0x43, 0xf4, 0x1d, 0x6c, 0x06, 0x51, 0x42, 0x52, 0x16, 0x65, 0xa9, 0x5f,
	0xa0, 0x59, 0x12, 0x6d, 0xef, 0x3a, 0xb4, 0xa3, 0xc2, 0xa8, 0x04, 0x6b, 0x07, 0x57, 0xc4, 0x3b,
	0xe7, 0xd0, 0x5c, 0xdc, 0x81, 0x6c, 0xd0, 0xbf, 0x27, 0xd3, 0xbc, 0x72, 0xc4, 0x12, 0xed, 0x42,
	0xf5, 0x12, 0xc7, 0x13, 0x55, 0xa2, 0x2f, 0x27, 0x47, 0xd6, 0xbc, 0xa7, 0xb6, 0x7c, 0x5a, 0xf9,
	0x58, 0xdb, 0xf9, 0x1a, 0xfe, 0xb7, 0xd4, 0xf9, 0xbf, 0x87, 0x3e, 0x33, 0x1a, 0x9a, 0x5d, 0x39,
	0x33, 0x1a, 0xba, 0x6d, 0xb8, 0xf7, 0xc1, 0x9c, 0x51, 0x29, 0xfa, 0x45, 0x16, 0xb9, 0x04, 0xd7,
	0x3d, 0xf5, 0x21, 0x1c, 0x92, 0x34, 0x90, 0xe0, 0xba, 0x27, 0x96, 0xee, 0xaf, 0x1a, 0x58, 0x8a,
	0x4e, 0x49, 0x11, 0xba, 0x0d, 0x8d, 0x31, 0xcd, 0xc6, 0x84, 0xf2, 0x3c, 0xae, 0x93, 0x35, 0x6f,
	0x26, 0x41, 0x77, 0xc0, 0x9c, 0xe5, 0x4b, 0xb5, 0xca, 0xc9, 0x9a, 0x37, 0x17, 0x09, 0xeb, 0x24,
	0x4f, 0xba, 0xec, 0x14, 0x69, 0x5d, 0x48, 0xd0, 0x3d, 0xa8, 0x66, 0x34, 0xc8, 0x6b, 0x79, 0xe3,
	0x6a, 0x19, 0x8a, 0x28, 0x9e, 0x08, 0xb5, 0xa7, 0x76, 0x1d, 0x34, 0x01, 0x44, 0x81, 0xa9, 0x92,
	0x72, 0x7f, 0xd7, 0xa0, 0xa9, 0x02, 0xed, 0x13, 0x1a, 0x11, 0x26, 0xfa, 0x83, 0xc9, 0x95, 0x2f,
	0x7a, 0x22, 0x4f, 0x22, 0x28, 0xd1, 0x63, 0x9c, 0x10, 0x74, 0x04, 0xa6, 0x08, 0x1b, 0xf3, 0x22,
	0xd8, 0x8d, 0xfd, 0xf7, 0x97, 0xd5, 0xb1, 0xc2, 0xeb, 0xf4, 0xc2, 0x90, 0x92, 0x10, 0x73, 0xf2,
	0x64, 0xec, 0xcd, 0x0d, 0xe7, 0x41, 0xeb, 0xff, 0x24, 0x68, 0xb7, 0x05, 0xd6, 0x02, 0x10, 0x6a,
	0x80, 0xf1, 0x45, 0xaf, 0xff, 0xd4, 0x5e, 0x43, 0x75, 0xd0, 0x7b, 0xcf, 0x3e, 0xb3, 0x35, 0xf7,
	0x47, 0x0d, 0x6a, 0xca, 0x31, 0x7a, 0x00, 0x8d, 0xc1, 0x54, 0x9d, 0x4f, 0xc6, 0x6f, 0xed, 0x6f,
	0x2f, 0x0b, 0x50, 0x32, 0x73, 0xb2, 0xe6, 0xd5, 0x07, 0x39, 0x49, 0x9f, 0x80, 0x39, 0x98, 0xfa,
	0xea, 0xa8, 0x79, 0xa5, 0xec, 0xac, 0x3e, 0x99, 0xe0, 0x60, 0x90, 0xaf, 0x0f, 0x2c, 0x30, 0x65,
	0x52, 0xf9, 0x74, 0x4c, 0xdc, 0xdf, 0x74, 0x68, 0x1e, 0xa7, 0x3c, 0xe2, 0x53, 0x4f, 0xce, 0x8c,
	0x7c, 0x06, 0x69, 0xb3, 0x19, 0x84, 0xc0, 0x90, 0xc9, 0x55, 0x53, 0x51, 0xae, 0xc5, 0x50, 0x1b,
	0xe1, 0x21, 0xf7, 0xa3, 0x80, 0x39, 0xba, 0x9c, 0x38, 0x75, 0xf1, 0x7d, 0x1a, 0x30, 0x31, 0x4a,
	0xd9, 0xf0, 0x82, 0x24, 0x58, 0x2a, 0x0d, 0xa9, 0x34, 0x95, 0x44, 0xa8, 0x8f, 0x16, 0xaa, 0xa3,
	0xbe, 0xac, 0xfd, 0x17, 0x63, 0x99, 0x75, 0xaf, 0xea, 0xd3, 0x79, 0x15, 0x9d, 0x01, 0xcc, 0x0a,
	0x8e, 0xe5, 0xf3, 0x69, 0xf7, 0x1a, 0x9c, 0x59, 0xeb, 0xe5, 0x1d, 0xbf, 0x60, 0xbd, 0xf3, 0x6c,
	0x3e, 0x84, 0x57, 0x75, 0x64, 0xb7, 0xdc, 0x91, 0xdb, 0xcb, 0x3a, 0xb2, 0x47, 0x29, 0x9e, 0x2e,
	0x76, 0x7c, 0x1f, 0x6e, 0x5d, 0x71, 0xfb, 0x9f, 0xf4, 0x7a, 0xd5, 0xae, 0x9d, 0x19, 0x8d, 0x9a,
	0x5d, 0x77, 0x7f, 0xd2, 0xc1, 0x3a, 0xbe, 0x24, 0x29, 0xbf, 0x01, 0x71, 0xb7, 0x41, 0xce, 0x77,
	0xc6, 0x71, 0x32, 0x96, 0xd5, 0xac, 0x7b, 0x73, 0x01, 0x3a, 0x5c, 0x20, 0xa7, 0x26, 0x93, 0xfa,
	0xc1, 0x95, 0xa4, 0xce, 0xdd, 0xad, 0xe4, 0xe6, 0xb4, 0xc4, 0x8d, 0xe2, 0xf8, 0xee, 0x6a, 0x98,
	0xd7, 0x88, 0x1a, 0xc3, 0xae, 0x2a, 0x82, 0xdc, 0x3f, 0x2b, 0xd0, 0x7c, 0x44, 0x38, 0x8d, 0x86,
	0x37, 0xe0, 0xa6, 0x5c, 0xd4, 0xd5, 0x65, 0x45, 0xbd, 0x88, 0x79, 0x5d, 0xe6, 0x4a, 0x6d, 0x56,
	0x5b, 0x71, 0xcb, 0xce, 0x91, 0x56, 0x50, 0xf9, 0x4a, 0xf2, 0xf4, 0xaa, 0x48, 0x55, 0x17, 0xa0,
	0x62, 0xc1, 0xfd, 0x59, 0x87, 0xf5, 0x5e, 0x9a, 0x25, 0x38, 0xbe, 0xc9, 0x54, 0x7b, 0x13, 0xcc,
	0x44, 0xa6, 0xc5, 0x8f, 0x02, 0x75, 0x75, 0xc9, 0x5c, 0xd0, 0x68, 0x78, 0x1a, 0xa0, 0xcf, 0x4b,
	0xec, 0xa8, 0x9c, 0x7e, 0x58, 0x0e, 0xac, 0xe4, 0xf1, 0x5a, 0x7a, 0x8e, 0x5f, 0x9a, 0x82, 0x77,
	0xaf, 0x83, 0x7a, 0x4d, 0xf8, 0x99, 0xf7, 0xc7, 0xb7, 0xb0, 0x59, 0xdc, 0x90, 0x51, 0x96, 0xe6,
	0x14, 0x6d, 0x41, 0x55, 0xfd, 0x5f, 0xab, 0x5f, 0x13, 0xf5, 0x31, 0x0f, 0x5f, 0xff, 0xdb, 0xf0,
	0xd5, 0x9f, 0xd0, 0xee, 0x1d, 0x30, 0x67, 0x57, 0xb2, 0xbc, 0x72, 0xfb, 0x87, 0xf6, 0x9a, 0xb8,
	0x85, 0x8f, 0x8e, 0xfb, 0x87, 0xb6, 0x76, 0xf0, 0x15, 0xb8, 0x19, 0x0d, 0x0b, 0x1c, 0xf1, 0x5e,
	0x51, 0x8f, 0x93, 0x1c, 0x52, 0xbd, 0x71, 0xce, 0xb5, 0x6f, 0x3e, 0x0a, 0x23, 0x7e, 0x31, 0x19,
	0x08, 0x41, 0x77, 0xe1, 0x71, 0x73, 0x4f, 0x3d, 0x6e, 0x2e, 0xf7, 0xf6, 0xba, 0x61, 0xf1, 0xc6,
	0x51, 0x36, 0x83, 0x9a, 0x54, 0xdd, 0xff, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x28, 0xc2, 0x61,
	0x46, 0x0d, 0x00, 0x00,
}
