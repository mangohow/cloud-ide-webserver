// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: pb/proto/service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ResourceLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cpu     string `protobuf:"bytes,1,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory  string `protobuf:"bytes,2,opt,name=Memory,proto3" json:"Memory,omitempty"`
	Storage string `protobuf:"bytes,3,opt,name=Storage,proto3" json:"Storage,omitempty"`
}

func (x *ResourceLimit) Reset() {
	*x = ResourceLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceLimit) ProtoMessage() {}

func (x *ResourceLimit) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceLimit.ProtoReflect.Descriptor instead.
func (*ResourceLimit) Descriptor() ([]byte, []int) {
	return file_pb_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceLimit) GetCpu() string {
	if x != nil {
		return x.Cpu
	}
	return ""
}

func (x *ResourceLimit) GetMemory() string {
	if x != nil {
		return x.Memory
	}
	return ""
}

func (x *ResourceLimit) GetStorage() string {
	if x != nil {
		return x.Storage
	}
	return ""
}

type PodInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name          string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace     string         `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Image         string         `protobuf:"bytes,3,opt,name=image,proto3" json:"image,omitempty"`
	Port          uint32         `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	ResourceLimit *ResourceLimit `protobuf:"bytes,5,opt,name=resourceLimit,proto3" json:"resourceLimit,omitempty"`
}

func (x *PodInfo) Reset() {
	*x = PodInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodInfo) ProtoMessage() {}

func (x *PodInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodInfo.ProtoReflect.Descriptor instead.
func (*PodInfo) Descriptor() ([]byte, []int) {
	return file_pb_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *PodInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PodInfo) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *PodInfo) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *PodInfo) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *PodInfo) GetResourceLimit() *ResourceLimit {
	if x != nil {
		return x.ResourceLimit
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_pb_proto_service_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type QueryOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *QueryOption) Reset() {
	*x = QueryOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryOption) ProtoMessage() {}

func (x *QueryOption) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryOption.ProtoReflect.Descriptor instead.
func (*QueryOption) Descriptor() ([]byte, []int) {
	return file_pb_proto_service_proto_rawDescGZIP(), []int{3}
}

func (x *QueryOption) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *QueryOption) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type PodStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PodStatus) Reset() {
	*x = PodStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodStatus) ProtoMessage() {}

func (x *PodStatus) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodStatus.ProtoReflect.Descriptor instead.
func (*PodStatus) Descriptor() ([]byte, []int) {
	return file_pb_proto_service_proto_rawDescGZIP(), []int{4}
}

func (x *PodStatus) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *PodStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type PodSpaceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName string `protobuf:"bytes,1,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	Ip       string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Port     int32  `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *PodSpaceInfo) Reset() {
	*x = PodSpaceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_proto_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PodSpaceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodSpaceInfo) ProtoMessage() {}

func (x *PodSpaceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pb_proto_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodSpaceInfo.ProtoReflect.Descriptor instead.
func (*PodSpaceInfo) Descriptor() ([]byte, []int) {
	return file_pb_proto_service_proto_rawDescGZIP(), []int{5}
}

func (x *PodSpaceInfo) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *PodSpaceInfo) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *PodSpaceInfo) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_pb_proto_service_proto protoreflect.FileDescriptor

var file_pb_proto_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x53, 0x0a, 0x0d,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x63, 0x70, 0x75, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12,
	0x16, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x22, 0x9e, 0x01, 0x0a, 0x07, 0x50, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x37, 0x0a, 0x0d, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69,
	0x6d, 0x69, 0x74, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x22, 0x3c, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x3f, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x22, 0x3d, 0x0a, 0x09, 0x50, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x4e, 0x0a, 0x0c, 0x50, 0x6f, 0x64, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74,
	0x32, 0xb1, 0x02, 0x0a, 0x0f, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x49, 0x64, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x70,
	0x61, 0x63, 0x65, 0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f,
	0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x64, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x2b, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x12, 0x0b, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x10, 0x2e,
	0x70, 0x62, 0x2e, 0x50, 0x6f, 0x64, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x2c, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x0f,
	0x2e, 0x70, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a,
	0x09, 0x73, 0x74, 0x6f, 0x70, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0c, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x11, 0x67, 0x65, 0x74,
	0x50, 0x6f, 0x64, 0x53, 0x70, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f,
	0x2e, 0x70, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x1a,
	0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x34,
	0x0a, 0x0f, 0x67, 0x65, 0x74, 0x50, 0x6f, 0x64, 0x53, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x64, 0x53, 0x70, 0x61, 0x63, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_proto_service_proto_rawDescOnce sync.Once
	file_pb_proto_service_proto_rawDescData = file_pb_proto_service_proto_rawDesc
)

func file_pb_proto_service_proto_rawDescGZIP() []byte {
	file_pb_proto_service_proto_rawDescOnce.Do(func() {
		file_pb_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_proto_service_proto_rawDescData)
	})
	return file_pb_proto_service_proto_rawDescData
}

var file_pb_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pb_proto_service_proto_goTypes = []interface{}{
	(*ResourceLimit)(nil), // 0: pb.ResourceLimit
	(*PodInfo)(nil),       // 1: pb.PodInfo
	(*Response)(nil),      // 2: pb.Response
	(*QueryOption)(nil),   // 3: pb.QueryOption
	(*PodStatus)(nil),     // 4: pb.PodStatus
	(*PodSpaceInfo)(nil),  // 5: pb.PodSpaceInfo
}
var file_pb_proto_service_proto_depIdxs = []int32{
	0, // 0: pb.PodInfo.resourceLimit:type_name -> pb.ResourceLimit
	1, // 1: pb.CloudIdeService.createSpace:input_type -> pb.PodInfo
	1, // 2: pb.CloudIdeService.startSpace:input_type -> pb.PodInfo
	3, // 3: pb.CloudIdeService.deleteSpace:input_type -> pb.QueryOption
	3, // 4: pb.CloudIdeService.stopSpace:input_type -> pb.QueryOption
	3, // 5: pb.CloudIdeService.getPodSpaceStatus:input_type -> pb.QueryOption
	3, // 6: pb.CloudIdeService.getPodSpaceInfo:input_type -> pb.QueryOption
	5, // 7: pb.CloudIdeService.createSpace:output_type -> pb.PodSpaceInfo
	5, // 8: pb.CloudIdeService.startSpace:output_type -> pb.PodSpaceInfo
	2, // 9: pb.CloudIdeService.deleteSpace:output_type -> pb.Response
	2, // 10: pb.CloudIdeService.stopSpace:output_type -> pb.Response
	4, // 11: pb.CloudIdeService.getPodSpaceStatus:output_type -> pb.PodStatus
	5, // 12: pb.CloudIdeService.getPodSpaceInfo:output_type -> pb.PodSpaceInfo
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_proto_service_proto_init() }
func file_pb_proto_service_proto_init() {
	if File_pb_proto_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_proto_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceLimit); i {
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
		file_pb_proto_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodInfo); i {
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
		file_pb_proto_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_pb_proto_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryOption); i {
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
		file_pb_proto_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodStatus); i {
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
		file_pb_proto_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PodSpaceInfo); i {
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
			RawDescriptor: file_pb_proto_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_proto_service_proto_goTypes,
		DependencyIndexes: file_pb_proto_service_proto_depIdxs,
		MessageInfos:      file_pb_proto_service_proto_msgTypes,
	}.Build()
	File_pb_proto_service_proto = out.File
	file_pb_proto_service_proto_rawDesc = nil
	file_pb_proto_service_proto_goTypes = nil
	file_pb_proto_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CloudIdeServiceClient is the client API for CloudIdeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CloudIdeServiceClient interface {
	// 创建云IDE空间并等待Pod状态变为Running,第一次创建,需要挂载存储卷
	CreateSpace(ctx context.Context, in *PodInfo, opts ...grpc.CallOption) (*PodSpaceInfo, error)
	// 启动(创建)云IDE空间,非第一次创建,无需挂载存储卷,使用之前的存储卷
	StartSpace(ctx context.Context, in *PodInfo, opts ...grpc.CallOption) (*PodSpaceInfo, error)
	// 删除云IDE空间,需要删除存储卷
	DeleteSpace(ctx context.Context, in *QueryOption, opts ...grpc.CallOption) (*Response, error)
	// 停止(删除)云工作空间,无需删除存储卷
	StopSpace(ctx context.Context, in *QueryOption, opts ...grpc.CallOption) (*Response, error)
	// 获取Pod运行状态
	GetPodSpaceStatus(ctx context.Context, in *QueryOption, opts ...grpc.CallOption) (*PodStatus, error)
	// 获取云IDE空间Pod的信息
	GetPodSpaceInfo(ctx context.Context, in *QueryOption, opts ...grpc.CallOption) (*PodSpaceInfo, error)
}

type cloudIdeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudIdeServiceClient(cc grpc.ClientConnInterface) CloudIdeServiceClient {
	return &cloudIdeServiceClient{cc}
}

func (c *cloudIdeServiceClient) CreateSpace(ctx context.Context, in *PodInfo, opts ...grpc.CallOption) (*PodSpaceInfo, error) {
	out := new(PodSpaceInfo)
	err := c.cc.Invoke(ctx, "/pb.CloudIdeService/createSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudIdeServiceClient) StartSpace(ctx context.Context, in *PodInfo, opts ...grpc.CallOption) (*PodSpaceInfo, error) {
	out := new(PodSpaceInfo)
	err := c.cc.Invoke(ctx, "/pb.CloudIdeService/startSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudIdeServiceClient) DeleteSpace(ctx context.Context, in *QueryOption, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.CloudIdeService/deleteSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudIdeServiceClient) StopSpace(ctx context.Context, in *QueryOption, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/pb.CloudIdeService/stopSpace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudIdeServiceClient) GetPodSpaceStatus(ctx context.Context, in *QueryOption, opts ...grpc.CallOption) (*PodStatus, error) {
	out := new(PodStatus)
	err := c.cc.Invoke(ctx, "/pb.CloudIdeService/getPodSpaceStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudIdeServiceClient) GetPodSpaceInfo(ctx context.Context, in *QueryOption, opts ...grpc.CallOption) (*PodSpaceInfo, error) {
	out := new(PodSpaceInfo)
	err := c.cc.Invoke(ctx, "/pb.CloudIdeService/getPodSpaceInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudIdeServiceServer is the server API for CloudIdeService service.
type CloudIdeServiceServer interface {
	// 创建云IDE空间并等待Pod状态变为Running,第一次创建,需要挂载存储卷
	CreateSpace(context.Context, *PodInfo) (*PodSpaceInfo, error)
	// 启动(创建)云IDE空间,非第一次创建,无需挂载存储卷,使用之前的存储卷
	StartSpace(context.Context, *PodInfo) (*PodSpaceInfo, error)
	// 删除云IDE空间,需要删除存储卷
	DeleteSpace(context.Context, *QueryOption) (*Response, error)
	// 停止(删除)云工作空间,无需删除存储卷
	StopSpace(context.Context, *QueryOption) (*Response, error)
	// 获取Pod运行状态
	GetPodSpaceStatus(context.Context, *QueryOption) (*PodStatus, error)
	// 获取云IDE空间Pod的信息
	GetPodSpaceInfo(context.Context, *QueryOption) (*PodSpaceInfo, error)
}

// UnimplementedCloudIdeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCloudIdeServiceServer struct {
}

func (*UnimplementedCloudIdeServiceServer) CreateSpace(context.Context, *PodInfo) (*PodSpaceInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSpace not implemented")
}
func (*UnimplementedCloudIdeServiceServer) StartSpace(context.Context, *PodInfo) (*PodSpaceInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartSpace not implemented")
}
func (*UnimplementedCloudIdeServiceServer) DeleteSpace(context.Context, *QueryOption) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSpace not implemented")
}
func (*UnimplementedCloudIdeServiceServer) StopSpace(context.Context, *QueryOption) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopSpace not implemented")
}
func (*UnimplementedCloudIdeServiceServer) GetPodSpaceStatus(context.Context, *QueryOption) (*PodStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPodSpaceStatus not implemented")
}
func (*UnimplementedCloudIdeServiceServer) GetPodSpaceInfo(context.Context, *QueryOption) (*PodSpaceInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPodSpaceInfo not implemented")
}

func RegisterCloudIdeServiceServer(s *grpc.Server, srv CloudIdeServiceServer) {
	s.RegisterService(&_CloudIdeService_serviceDesc, srv)
}

func _CloudIdeService_CreateSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudIdeServiceServer).CreateSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CloudIdeService/CreateSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudIdeServiceServer).CreateSpace(ctx, req.(*PodInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudIdeService_StartSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PodInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudIdeServiceServer).StartSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CloudIdeService/StartSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudIdeServiceServer).StartSpace(ctx, req.(*PodInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudIdeService_DeleteSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOption)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudIdeServiceServer).DeleteSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CloudIdeService/DeleteSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudIdeServiceServer).DeleteSpace(ctx, req.(*QueryOption))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudIdeService_StopSpace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOption)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudIdeServiceServer).StopSpace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CloudIdeService/StopSpace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudIdeServiceServer).StopSpace(ctx, req.(*QueryOption))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudIdeService_GetPodSpaceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOption)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudIdeServiceServer).GetPodSpaceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CloudIdeService/GetPodSpaceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudIdeServiceServer).GetPodSpaceStatus(ctx, req.(*QueryOption))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudIdeService_GetPodSpaceInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryOption)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudIdeServiceServer).GetPodSpaceInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CloudIdeService/GetPodSpaceInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudIdeServiceServer).GetPodSpaceInfo(ctx, req.(*QueryOption))
	}
	return interceptor(ctx, in, info, handler)
}

var _CloudIdeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CloudIdeService",
	HandlerType: (*CloudIdeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createSpace",
			Handler:    _CloudIdeService_CreateSpace_Handler,
		},
		{
			MethodName: "startSpace",
			Handler:    _CloudIdeService_StartSpace_Handler,
		},
		{
			MethodName: "deleteSpace",
			Handler:    _CloudIdeService_DeleteSpace_Handler,
		},
		{
			MethodName: "stopSpace",
			Handler:    _CloudIdeService_StopSpace_Handler,
		},
		{
			MethodName: "getPodSpaceStatus",
			Handler:    _CloudIdeService_GetPodSpaceStatus_Handler,
		},
		{
			MethodName: "getPodSpaceInfo",
			Handler:    _CloudIdeService_GetPodSpaceInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/proto/service.proto",
}
