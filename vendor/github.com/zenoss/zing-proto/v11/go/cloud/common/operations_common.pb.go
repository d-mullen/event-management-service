// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: zenoss/zing/proto/cloud/operations_common.proto

package common

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Specify comparing operations names
type Operation int32

const (
	Operation_OP_EXISTS        Operation = 0
	Operation_OP_EQUALS        Operation = 1
	Operation_OP_NOT_EQUALS    Operation = 2
	Operation_OP_CONTAINS      Operation = 3
	Operation_OP_LESS          Operation = 4
	Operation_OP_GREATER       Operation = 5
	Operation_OP_LESS_OR_EQ    Operation = 6
	Operation_OP_GREATER_OR_EQ Operation = 7
	Operation_OP_STARTS_WITH   Operation = 8
	Operation_OP_IN            Operation = 9
	Operation_OP_NOT_IN        Operation = 10
	Operation_OP_END_WITH      Operation = 11
	Operation_OP_NOT_CONTAINS  Operation = 12
)

// Enum value maps for Operation.
var (
	Operation_name = map[int32]string{
		0:  "OP_EXISTS",
		1:  "OP_EQUALS",
		2:  "OP_NOT_EQUALS",
		3:  "OP_CONTAINS",
		4:  "OP_LESS",
		5:  "OP_GREATER",
		6:  "OP_LESS_OR_EQ",
		7:  "OP_GREATER_OR_EQ",
		8:  "OP_STARTS_WITH",
		9:  "OP_IN",
		10: "OP_NOT_IN",
		11: "OP_END_WITH",
		12: "OP_NOT_CONTAINS",
	}
	Operation_value = map[string]int32{
		"OP_EXISTS":        0,
		"OP_EQUALS":        1,
		"OP_NOT_EQUALS":    2,
		"OP_CONTAINS":      3,
		"OP_LESS":          4,
		"OP_GREATER":       5,
		"OP_LESS_OR_EQ":    6,
		"OP_GREATER_OR_EQ": 7,
		"OP_STARTS_WITH":   8,
		"OP_IN":            9,
		"OP_NOT_IN":        10,
		"OP_END_WITH":      11,
		"OP_NOT_CONTAINS":  12,
	}
)

func (x Operation) Enum() *Operation {
	p := new(Operation)
	*p = x
	return p
}

func (x Operation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Operation) Descriptor() protoreflect.EnumDescriptor {
	return file_zenoss_zing_proto_cloud_operations_common_proto_enumTypes[0].Descriptor()
}

func (Operation) Type() protoreflect.EnumType {
	return &file_zenoss_zing_proto_cloud_operations_common_proto_enumTypes[0]
}

func (x Operation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Operation.Descriptor instead.
func (Operation) EnumDescriptor() ([]byte, []int) {
	return file_zenoss_zing_proto_cloud_operations_common_proto_rawDescGZIP(), []int{0}
}

var File_zenoss_zing_proto_cloud_operations_common_proto protoreflect.FileDescriptor

var file_zenoss_zing_proto_cloud_operations_common_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x7a, 0x65, 0x6e, 0x6f, 0x73, 0x73, 0x2f, 0x7a, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x7a, 0x65, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2a,
	0xe7, 0x01, 0x0a, 0x09, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0d, 0x0a,
	0x09, 0x4f, 0x50, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x4f, 0x50, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x53, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x4f,
	0x50, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x45, 0x51, 0x55, 0x41, 0x4c, 0x53, 0x10, 0x02, 0x12, 0x0f,
	0x0a, 0x0b, 0x4f, 0x50, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x53, 0x10, 0x03, 0x12,
	0x0b, 0x0a, 0x07, 0x4f, 0x50, 0x5f, 0x4c, 0x45, 0x53, 0x53, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a,
	0x4f, 0x50, 0x5f, 0x47, 0x52, 0x45, 0x41, 0x54, 0x45, 0x52, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d,
	0x4f, 0x50, 0x5f, 0x4c, 0x45, 0x53, 0x53, 0x5f, 0x4f, 0x52, 0x5f, 0x45, 0x51, 0x10, 0x06, 0x12,
	0x14, 0x0a, 0x10, 0x4f, 0x50, 0x5f, 0x47, 0x52, 0x45, 0x41, 0x54, 0x45, 0x52, 0x5f, 0x4f, 0x52,
	0x5f, 0x45, 0x51, 0x10, 0x07, 0x12, 0x12, 0x0a, 0x0e, 0x4f, 0x50, 0x5f, 0x53, 0x54, 0x41, 0x52,
	0x54, 0x53, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x10, 0x08, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x50, 0x5f,
	0x49, 0x4e, 0x10, 0x09, 0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x50, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x49,
	0x4e, 0x10, 0x0a, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x50, 0x5f, 0x45, 0x4e, 0x44, 0x5f, 0x57, 0x49,
	0x54, 0x48, 0x10, 0x0b, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x50, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x43,
	0x4f, 0x4e, 0x54, 0x41, 0x49, 0x4e, 0x53, 0x10, 0x0c, 0x42, 0x58, 0x0a, 0x22, 0x6f, 0x72, 0x67,
	0x2e, 0x7a, 0x65, 0x6e, 0x6f, 0x73, 0x73, 0x2e, 0x7a, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50,
	0x01, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x65,
	0x6e, 0x6f, 0x73, 0x73, 0x2f, 0x7a, 0x69, 0x6e, 0x67, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x76, 0x31, 0x31, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zenoss_zing_proto_cloud_operations_common_proto_rawDescOnce sync.Once
	file_zenoss_zing_proto_cloud_operations_common_proto_rawDescData = file_zenoss_zing_proto_cloud_operations_common_proto_rawDesc
)

func file_zenoss_zing_proto_cloud_operations_common_proto_rawDescGZIP() []byte {
	file_zenoss_zing_proto_cloud_operations_common_proto_rawDescOnce.Do(func() {
		file_zenoss_zing_proto_cloud_operations_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_zenoss_zing_proto_cloud_operations_common_proto_rawDescData)
	})
	return file_zenoss_zing_proto_cloud_operations_common_proto_rawDescData
}

var file_zenoss_zing_proto_cloud_operations_common_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_zenoss_zing_proto_cloud_operations_common_proto_goTypes = []interface{}{
	(Operation)(0), // 0: zenoss.cloud.Operation
}
var file_zenoss_zing_proto_cloud_operations_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_zenoss_zing_proto_cloud_operations_common_proto_init() }
func file_zenoss_zing_proto_cloud_operations_common_proto_init() {
	if File_zenoss_zing_proto_cloud_operations_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_zenoss_zing_proto_cloud_operations_common_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_zenoss_zing_proto_cloud_operations_common_proto_goTypes,
		DependencyIndexes: file_zenoss_zing_proto_cloud_operations_common_proto_depIdxs,
		EnumInfos:         file_zenoss_zing_proto_cloud_operations_common_proto_enumTypes,
	}.Build()
	File_zenoss_zing_proto_cloud_operations_common_proto = out.File
	file_zenoss_zing_proto_cloud_operations_common_proto_rawDesc = nil
	file_zenoss_zing_proto_cloud_operations_common_proto_goTypes = nil
	file_zenoss_zing_proto_cloud_operations_common_proto_depIdxs = nil
}
