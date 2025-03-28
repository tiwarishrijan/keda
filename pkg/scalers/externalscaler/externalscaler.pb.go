// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.3
// source: externalscaler.proto

package externalscaler

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

type ScaledObjectRef struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Name           string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace      string                 `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	ScalerMetadata map[string]string      `protobuf:"bytes,3,rep,name=scalerMetadata,proto3" json:"scalerMetadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *ScaledObjectRef) Reset() {
	*x = ScaledObjectRef{}
	mi := &file_externalscaler_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ScaledObjectRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScaledObjectRef) ProtoMessage() {}

func (x *ScaledObjectRef) ProtoReflect() protoreflect.Message {
	mi := &file_externalscaler_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScaledObjectRef.ProtoReflect.Descriptor instead.
func (*ScaledObjectRef) Descriptor() ([]byte, []int) {
	return file_externalscaler_proto_rawDescGZIP(), []int{0}
}

func (x *ScaledObjectRef) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ScaledObjectRef) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ScaledObjectRef) GetScalerMetadata() map[string]string {
	if x != nil {
		return x.ScalerMetadata
	}
	return nil
}

type IsActiveResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Result        bool                   `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *IsActiveResponse) Reset() {
	*x = IsActiveResponse{}
	mi := &file_externalscaler_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *IsActiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsActiveResponse) ProtoMessage() {}

func (x *IsActiveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_externalscaler_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsActiveResponse.ProtoReflect.Descriptor instead.
func (*IsActiveResponse) Descriptor() ([]byte, []int) {
	return file_externalscaler_proto_rawDescGZIP(), []int{1}
}

func (x *IsActiveResponse) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type GetMetricSpecResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MetricSpecs   []*MetricSpec          `protobuf:"bytes,1,rep,name=metricSpecs,proto3" json:"metricSpecs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMetricSpecResponse) Reset() {
	*x = GetMetricSpecResponse{}
	mi := &file_externalscaler_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMetricSpecResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetricSpecResponse) ProtoMessage() {}

func (x *GetMetricSpecResponse) ProtoReflect() protoreflect.Message {
	mi := &file_externalscaler_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetricSpecResponse.ProtoReflect.Descriptor instead.
func (*GetMetricSpecResponse) Descriptor() ([]byte, []int) {
	return file_externalscaler_proto_rawDescGZIP(), []int{2}
}

func (x *GetMetricSpecResponse) GetMetricSpecs() []*MetricSpec {
	if x != nil {
		return x.MetricSpecs
	}
	return nil
}

type MetricSpec struct {
	state      protoimpl.MessageState `protogen:"open.v1"`
	MetricName string                 `protobuf:"bytes,1,opt,name=metricName,proto3" json:"metricName,omitempty"`
	// deprecated, use targetSizeFloat instead
	TargetSize      int64   `protobuf:"varint,2,opt,name=targetSize,proto3" json:"targetSize,omitempty"`
	TargetSizeFloat float64 `protobuf:"fixed64,3,opt,name=targetSizeFloat,proto3" json:"targetSizeFloat,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *MetricSpec) Reset() {
	*x = MetricSpec{}
	mi := &file_externalscaler_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricSpec) ProtoMessage() {}

func (x *MetricSpec) ProtoReflect() protoreflect.Message {
	mi := &file_externalscaler_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricSpec.ProtoReflect.Descriptor instead.
func (*MetricSpec) Descriptor() ([]byte, []int) {
	return file_externalscaler_proto_rawDescGZIP(), []int{3}
}

func (x *MetricSpec) GetMetricName() string {
	if x != nil {
		return x.MetricName
	}
	return ""
}

func (x *MetricSpec) GetTargetSize() int64 {
	if x != nil {
		return x.TargetSize
	}
	return 0
}

func (x *MetricSpec) GetTargetSizeFloat() float64 {
	if x != nil {
		return x.TargetSizeFloat
	}
	return 0
}

type GetMetricsRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	ScaledObjectRef *ScaledObjectRef       `protobuf:"bytes,1,opt,name=scaledObjectRef,proto3" json:"scaledObjectRef,omitempty"`
	MetricName      string                 `protobuf:"bytes,2,opt,name=metricName,proto3" json:"metricName,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *GetMetricsRequest) Reset() {
	*x = GetMetricsRequest{}
	mi := &file_externalscaler_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetricsRequest) ProtoMessage() {}

func (x *GetMetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_externalscaler_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetricsRequest.ProtoReflect.Descriptor instead.
func (*GetMetricsRequest) Descriptor() ([]byte, []int) {
	return file_externalscaler_proto_rawDescGZIP(), []int{4}
}

func (x *GetMetricsRequest) GetScaledObjectRef() *ScaledObjectRef {
	if x != nil {
		return x.ScaledObjectRef
	}
	return nil
}

func (x *GetMetricsRequest) GetMetricName() string {
	if x != nil {
		return x.MetricName
	}
	return ""
}

type GetMetricsResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MetricValues  []*MetricValue         `protobuf:"bytes,1,rep,name=metricValues,proto3" json:"metricValues,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetMetricsResponse) Reset() {
	*x = GetMetricsResponse{}
	mi := &file_externalscaler_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetMetricsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMetricsResponse) ProtoMessage() {}

func (x *GetMetricsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_externalscaler_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMetricsResponse.ProtoReflect.Descriptor instead.
func (*GetMetricsResponse) Descriptor() ([]byte, []int) {
	return file_externalscaler_proto_rawDescGZIP(), []int{5}
}

func (x *GetMetricsResponse) GetMetricValues() []*MetricValue {
	if x != nil {
		return x.MetricValues
	}
	return nil
}

type MetricValue struct {
	state      protoimpl.MessageState `protogen:"open.v1"`
	MetricName string                 `protobuf:"bytes,1,opt,name=metricName,proto3" json:"metricName,omitempty"`
	// deprecated, use metricValueFloat instead
	MetricValue      int64   `protobuf:"varint,2,opt,name=metricValue,proto3" json:"metricValue,omitempty"`
	MetricValueFloat float64 `protobuf:"fixed64,3,opt,name=metricValueFloat,proto3" json:"metricValueFloat,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *MetricValue) Reset() {
	*x = MetricValue{}
	mi := &file_externalscaler_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricValue) ProtoMessage() {}

func (x *MetricValue) ProtoReflect() protoreflect.Message {
	mi := &file_externalscaler_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricValue.ProtoReflect.Descriptor instead.
func (*MetricValue) Descriptor() ([]byte, []int) {
	return file_externalscaler_proto_rawDescGZIP(), []int{6}
}

func (x *MetricValue) GetMetricName() string {
	if x != nil {
		return x.MetricName
	}
	return ""
}

func (x *MetricValue) GetMetricValue() int64 {
	if x != nil {
		return x.MetricValue
	}
	return 0
}

func (x *MetricValue) GetMetricValueFloat() float64 {
	if x != nil {
		return x.MetricValueFloat
	}
	return 0
}

var File_externalscaler_proto protoreflect.FileDescriptor

var file_externalscaler_proto_rawDesc = []byte{
	0x0a, 0x14, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x22, 0xe3, 0x01, 0x0a, 0x0f, 0x53, 0x63, 0x61, 0x6c, 0x65,
	0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x5b, 0x0a, 0x0e,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x66, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0e, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x41, 0x0a, 0x13, 0x53, 0x63, 0x61,
	0x6c, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2a, 0x0a, 0x10,
	0x49, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x55, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x70, 0x65, 0x63, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x70,
	0x65, 0x63, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x70, 0x65, 0x63, 0x73, 0x22,
	0x76, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1e, 0x0a,
	0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x28, 0x0a,
	0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x69, 0x7a, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x53, 0x69,
	0x7a, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x22, 0x7e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x49, 0x0a, 0x0f,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x64, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x52, 0x0f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x64, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x12, 0x1e, 0x0a, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x55, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a,
	0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x72, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x7b,
	0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x2a, 0x0a, 0x10, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x6c,
	0x6f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x32, 0xec, 0x02, 0x0a, 0x0e,
	0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x12, 0x4f,
	0x0a, 0x08, 0x49, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1f, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x61, 0x6c,
	0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x1a, 0x20, 0x2e, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x49, 0x73, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x57, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x49, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x12, 0x1f, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x72, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x66, 0x1a, 0x20, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x72, 0x2e, 0x49, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x59, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x70, 0x65, 0x63, 0x12, 0x1f, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x53, 0x63, 0x61, 0x6c, 0x65,
	0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x66, 0x1a, 0x25, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x70, 0x65, 0x63, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x12, 0x21, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x63, 0x61, 0x6c,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x3b,
	0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_externalscaler_proto_rawDescOnce sync.Once
	file_externalscaler_proto_rawDescData = file_externalscaler_proto_rawDesc
)

func file_externalscaler_proto_rawDescGZIP() []byte {
	file_externalscaler_proto_rawDescOnce.Do(func() {
		file_externalscaler_proto_rawDescData = protoimpl.X.CompressGZIP(file_externalscaler_proto_rawDescData)
	})
	return file_externalscaler_proto_rawDescData
}

var file_externalscaler_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_externalscaler_proto_goTypes = []any{
	(*ScaledObjectRef)(nil),       // 0: externalscaler.ScaledObjectRef
	(*IsActiveResponse)(nil),      // 1: externalscaler.IsActiveResponse
	(*GetMetricSpecResponse)(nil), // 2: externalscaler.GetMetricSpecResponse
	(*MetricSpec)(nil),            // 3: externalscaler.MetricSpec
	(*GetMetricsRequest)(nil),     // 4: externalscaler.GetMetricsRequest
	(*GetMetricsResponse)(nil),    // 5: externalscaler.GetMetricsResponse
	(*MetricValue)(nil),           // 6: externalscaler.MetricValue
	nil,                           // 7: externalscaler.ScaledObjectRef.ScalerMetadataEntry
}
var file_externalscaler_proto_depIdxs = []int32{
	7, // 0: externalscaler.ScaledObjectRef.scalerMetadata:type_name -> externalscaler.ScaledObjectRef.ScalerMetadataEntry
	3, // 1: externalscaler.GetMetricSpecResponse.metricSpecs:type_name -> externalscaler.MetricSpec
	0, // 2: externalscaler.GetMetricsRequest.scaledObjectRef:type_name -> externalscaler.ScaledObjectRef
	6, // 3: externalscaler.GetMetricsResponse.metricValues:type_name -> externalscaler.MetricValue
	0, // 4: externalscaler.ExternalScaler.IsActive:input_type -> externalscaler.ScaledObjectRef
	0, // 5: externalscaler.ExternalScaler.StreamIsActive:input_type -> externalscaler.ScaledObjectRef
	0, // 6: externalscaler.ExternalScaler.GetMetricSpec:input_type -> externalscaler.ScaledObjectRef
	4, // 7: externalscaler.ExternalScaler.GetMetrics:input_type -> externalscaler.GetMetricsRequest
	1, // 8: externalscaler.ExternalScaler.IsActive:output_type -> externalscaler.IsActiveResponse
	1, // 9: externalscaler.ExternalScaler.StreamIsActive:output_type -> externalscaler.IsActiveResponse
	2, // 10: externalscaler.ExternalScaler.GetMetricSpec:output_type -> externalscaler.GetMetricSpecResponse
	5, // 11: externalscaler.ExternalScaler.GetMetrics:output_type -> externalscaler.GetMetricsResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_externalscaler_proto_init() }
func file_externalscaler_proto_init() {
	if File_externalscaler_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_externalscaler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_externalscaler_proto_goTypes,
		DependencyIndexes: file_externalscaler_proto_depIdxs,
		MessageInfos:      file_externalscaler_proto_msgTypes,
	}.Build()
	File_externalscaler_proto = out.File
	file_externalscaler_proto_rawDesc = nil
	file_externalscaler_proto_goTypes = nil
	file_externalscaler_proto_depIdxs = nil
}
