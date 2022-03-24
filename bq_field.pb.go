// Copyright 2014 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.6
// source: bq_field.proto

package robot_data

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Message containing options related to BigQuery schema generation
// and management via Protobuf.
type BigQueryFieldOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Flag to specify that a field should be marked as 'REQUIRED' when
	// used to generate schema for BigQuery.
	Require bool `protobuf:"varint,1,opt,name=require,proto3" json:"require,omitempty"`
	// Optionally override whatever type is resolved by the schema
	// generator. This is useful, for example, to store epoch timestamps
	// with the underlying 'TIMESTAMP' type, when normally, they would
	// be structured as 'INTEGER' fields.
	TypeOverride string `protobuf:"bytes,2,opt,name=type_override,json=typeOverride,proto3" json:"type_override,omitempty"`
	// Optionally omit a field from BigQuery schema.
	Ignore bool `protobuf:"varint,3,opt,name=ignore,proto3" json:"ignore,omitempty"`
	// Set the description for a field in BigQuery schema.
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	// Customize the name of the field in the BigQuery schema.
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *BigQueryFieldOptions) Reset() {
	*x = BigQueryFieldOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bq_field_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigQueryFieldOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigQueryFieldOptions) ProtoMessage() {}

func (x *BigQueryFieldOptions) ProtoReflect() protoreflect.Message {
	mi := &file_bq_field_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigQueryFieldOptions.ProtoReflect.Descriptor instead.
func (*BigQueryFieldOptions) Descriptor() ([]byte, []int) {
	return file_bq_field_proto_rawDescGZIP(), []int{0}
}

func (x *BigQueryFieldOptions) GetRequire() bool {
	if x != nil {
		return x.Require
	}
	return false
}

func (x *BigQueryFieldOptions) GetTypeOverride() string {
	if x != nil {
		return x.TypeOverride
	}
	return ""
}

func (x *BigQueryFieldOptions) GetIgnore() bool {
	if x != nil {
		return x.Ignore
	}
	return false
}

func (x *BigQueryFieldOptions) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *BigQueryFieldOptions) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var file_bq_field_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*BigQueryFieldOptions)(nil),
		Field:         1021,
		Name:          "robot_data.bigquery",
		Tag:           "bytes,1021,opt,name=bigquery",
		Filename:      "bq_field.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// BigQuery field schema generation options.
	//
	// optional robot_data.BigQueryFieldOptions bigquery = 1021;
	E_Bigquery = &file_bq_field_proto_extTypes[0]
)

var File_bq_field_proto protoreflect.FileDescriptor

var file_bq_field_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x62, 0x71, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0a, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3,
	0x01, 0x0a, 0x14, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x69,
	0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x4f, 0x76,
	0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x5c, 0x0a, 0x08, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xfd, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x5f, 0x64,
	0x61, 0x74, 0x61, 0x2e, 0x42, 0x69, 0x67, 0x51, 0x75, 0x65, 0x72, 0x79, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x08, 0x62, 0x69, 0x67, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x53, 0x50, 0x6f, 0x6f, 0x4e, 0x71, 0x69, 0x72, 0x2f, 0x72, 0x6f, 0x62, 0x6f, 0x74, 0x2d,
	0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bq_field_proto_rawDescOnce sync.Once
	file_bq_field_proto_rawDescData = file_bq_field_proto_rawDesc
)

func file_bq_field_proto_rawDescGZIP() []byte {
	file_bq_field_proto_rawDescOnce.Do(func() {
		file_bq_field_proto_rawDescData = protoimpl.X.CompressGZIP(file_bq_field_proto_rawDescData)
	})
	return file_bq_field_proto_rawDescData
}

var file_bq_field_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_bq_field_proto_goTypes = []interface{}{
	(*BigQueryFieldOptions)(nil),      // 0: robot_data.BigQueryFieldOptions
	(*descriptorpb.FieldOptions)(nil), // 1: google.protobuf.FieldOptions
}
var file_bq_field_proto_depIdxs = []int32{
	1, // 0: robot_data.bigquery:extendee -> google.protobuf.FieldOptions
	0, // 1: robot_data.bigquery:type_name -> robot_data.BigQueryFieldOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_bq_field_proto_init() }
func file_bq_field_proto_init() {
	if File_bq_field_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bq_field_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigQueryFieldOptions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bq_field_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_bq_field_proto_goTypes,
		DependencyIndexes: file_bq_field_proto_depIdxs,
		MessageInfos:      file_bq_field_proto_msgTypes,
		ExtensionInfos:    file_bq_field_proto_extTypes,
	}.Build()
	File_bq_field_proto = out.File
	file_bq_field_proto_rawDesc = nil
	file_bq_field_proto_goTypes = nil
	file_bq_field_proto_depIdxs = nil
}
