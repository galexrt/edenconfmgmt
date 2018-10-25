// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha/api.proto

package v1alpha

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ObjectMetadata struct {
	// ID of object.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of object.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Labels of object.
	Labels map[string]string `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Annotations of object.
	Annotations          map[string]string `protobuf:"bytes,4,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ObjectMetadata) Reset()         { *m = ObjectMetadata{} }
func (m *ObjectMetadata) String() string { return proto.CompactTextString(m) }
func (*ObjectMetadata) ProtoMessage()    {}
func (*ObjectMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_30dbcc7c2d394087, []int{0}
}

func (m *ObjectMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectMetadata.Unmarshal(m, b)
}
func (m *ObjectMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectMetadata.Marshal(b, m, deterministic)
}
func (m *ObjectMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectMetadata.Merge(m, src)
}
func (m *ObjectMetadata) XXX_Size() int {
	return xxx_messageInfo_ObjectMetadata.Size(m)
}
func (m *ObjectMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectMetadata proto.InternalMessageInfo

func (m *ObjectMetadata) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ObjectMetadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ObjectMetadata) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ObjectMetadata) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

type VersionRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionRequest) Reset()         { *m = VersionRequest{} }
func (m *VersionRequest) String() string { return proto.CompactTextString(m) }
func (*VersionRequest) ProtoMessage()    {}
func (*VersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_30dbcc7c2d394087, []int{1}
}

func (m *VersionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionRequest.Unmarshal(m, b)
}
func (m *VersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionRequest.Marshal(b, m, deterministic)
}
func (m *VersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionRequest.Merge(m, src)
}
func (m *VersionRequest) XXX_Size() int {
	return xxx_messageInfo_VersionRequest.Size(m)
}
func (m *VersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VersionRequest proto.InternalMessageInfo

type VersionResponse struct {
	// Semversion compatible version number.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Semversion major number.
	Major int64 `protobuf:"varint,2,opt,name=major,proto3" json:"major,omitempty"`
	// Semversion minor number.
	Minor int64 `protobuf:"varint,3,opt,name=minor,proto3" json:"minor,omitempty"`
	// Semversion patch number.
	Patch int64 `protobuf:"varint,4,opt,name=patch,proto3" json:"patch,omitempty"`
	// API state (e.g., alpha, beta, stable).
	State                string   `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VersionResponse) Reset()         { *m = VersionResponse{} }
func (m *VersionResponse) String() string { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()    {}
func (*VersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_30dbcc7c2d394087, []int{2}
}

func (m *VersionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VersionResponse.Unmarshal(m, b)
}
func (m *VersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VersionResponse.Marshal(b, m, deterministic)
}
func (m *VersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionResponse.Merge(m, src)
}
func (m *VersionResponse) XXX_Size() int {
	return xxx_messageInfo_VersionResponse.Size(m)
}
func (m *VersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VersionResponse proto.InternalMessageInfo

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *VersionResponse) GetMajor() int64 {
	if m != nil {
		return m.Major
	}
	return 0
}

func (m *VersionResponse) GetMinor() int64 {
	if m != nil {
		return m.Minor
	}
	return 0
}

func (m *VersionResponse) GetPatch() int64 {
	if m != nil {
		return m.Patch
	}
	return 0
}

func (m *VersionResponse) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

type Include struct {
	// Name of include (if no paths are given, will be used as file name/path).
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// List of paths.
	Paths                []string `protobuf:"bytes,2,rep,name=paths,proto3" json:"paths,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Include) Reset()         { *m = Include{} }
func (m *Include) String() string { return proto.CompactTextString(m) }
func (*Include) ProtoMessage()    {}
func (*Include) Descriptor() ([]byte, []int) {
	return fileDescriptor_30dbcc7c2d394087, []int{3}
}

func (m *Include) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Include.Unmarshal(m, b)
}
func (m *Include) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Include.Marshal(b, m, deterministic)
}
func (m *Include) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Include.Merge(m, src)
}
func (m *Include) XXX_Size() int {
	return xxx_messageInfo_Include.Size(m)
}
func (m *Include) XXX_DiscardUnknown() {
	xxx_messageInfo_Include.DiscardUnknown(m)
}

var xxx_messageInfo_Include proto.InternalMessageInfo

func (m *Include) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Include) GetPaths() []string {
	if m != nil {
		return m.Paths
	}
	return nil
}

type Conditions struct {
	// `when` condition.
	When *Condition `protobuf:"bytes,1,opt,name=when,proto3" json:"when,omitempty"`
	// `success` condition.
	Success              *Condition `protobuf:"bytes,2,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Conditions) Reset()         { *m = Conditions{} }
func (m *Conditions) String() string { return proto.CompactTextString(m) }
func (*Conditions) ProtoMessage()    {}
func (*Conditions) Descriptor() ([]byte, []int) {
	return fileDescriptor_30dbcc7c2d394087, []int{4}
}

func (m *Conditions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Conditions.Unmarshal(m, b)
}
func (m *Conditions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Conditions.Marshal(b, m, deterministic)
}
func (m *Conditions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Conditions.Merge(m, src)
}
func (m *Conditions) XXX_Size() int {
	return xxx_messageInfo_Conditions.Size(m)
}
func (m *Conditions) XXX_DiscardUnknown() {
	xxx_messageInfo_Conditions.DiscardUnknown(m)
}

var xxx_messageInfo_Conditions proto.InternalMessageInfo

func (m *Conditions) GetWhen() *Condition {
	if m != nil {
		return m.When
	}
	return nil
}

func (m *Conditions) GetSuccess() *Condition {
	if m != nil {
		return m.Success
	}
	return nil
}

type Condition struct {
	Condition            string   `protobuf:"bytes,1,opt,name=condition,proto3" json:"condition,omitempty"`
	Retry                *Retry   `protobuf:"bytes,2,opt,name=retry,proto3" json:"retry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Condition) Reset()         { *m = Condition{} }
func (m *Condition) String() string { return proto.CompactTextString(m) }
func (*Condition) ProtoMessage()    {}
func (*Condition) Descriptor() ([]byte, []int) {
	return fileDescriptor_30dbcc7c2d394087, []int{5}
}

func (m *Condition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Condition.Unmarshal(m, b)
}
func (m *Condition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Condition.Marshal(b, m, deterministic)
}
func (m *Condition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Condition.Merge(m, src)
}
func (m *Condition) XXX_Size() int {
	return xxx_messageInfo_Condition.Size(m)
}
func (m *Condition) XXX_DiscardUnknown() {
	xxx_messageInfo_Condition.DiscardUnknown(m)
}

var xxx_messageInfo_Condition proto.InternalMessageInfo

func (m *Condition) GetCondition() string {
	if m != nil {
		return m.Condition
	}
	return ""
}

func (m *Condition) GetRetry() *Retry {
	if m != nil {
		return m.Retry
	}
	return nil
}

type Retry struct {
	Interval             *duration.Duration `protobuf:"bytes,1,opt,name=interval,proto3" json:"interval,omitempty"`
	Limit                int64              `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Retry) Reset()         { *m = Retry{} }
func (m *Retry) String() string { return proto.CompactTextString(m) }
func (*Retry) ProtoMessage()    {}
func (*Retry) Descriptor() ([]byte, []int) {
	return fileDescriptor_30dbcc7c2d394087, []int{6}
}

func (m *Retry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Retry.Unmarshal(m, b)
}
func (m *Retry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Retry.Marshal(b, m, deterministic)
}
func (m *Retry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Retry.Merge(m, src)
}
func (m *Retry) XXX_Size() int {
	return xxx_messageInfo_Retry.Size(m)
}
func (m *Retry) XXX_DiscardUnknown() {
	xxx_messageInfo_Retry.DiscardUnknown(m)
}

var xxx_messageInfo_Retry proto.InternalMessageInfo

func (m *Retry) GetInterval() *duration.Duration {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *Retry) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func init() {
	proto.RegisterType((*ObjectMetadata)(nil), "core.v1alpha.ObjectMetadata")
	proto.RegisterMapType((map[string]string)(nil), "core.v1alpha.ObjectMetadata.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "core.v1alpha.ObjectMetadata.LabelsEntry")
	proto.RegisterType((*VersionRequest)(nil), "core.v1alpha.VersionRequest")
	proto.RegisterType((*VersionResponse)(nil), "core.v1alpha.VersionResponse")
	proto.RegisterType((*Include)(nil), "core.v1alpha.Include")
	proto.RegisterType((*Conditions)(nil), "core.v1alpha.Conditions")
	proto.RegisterType((*Condition)(nil), "core.v1alpha.Condition")
	proto.RegisterType((*Retry)(nil), "core.v1alpha.Retry")
}

func init() {
	proto.RegisterFile("github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha/api.proto", fileDescriptor_30dbcc7c2d394087)
}

var fileDescriptor_30dbcc7c2d394087 = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x97, 0xa4, 0x5d, 0xe9, 0x5b, 0x54, 0x2a, 0x83, 0x44, 0xa8, 0x50, 0x54, 0xe5, 0x54,
	0x84, 0x96, 0x68, 0x9b, 0x26, 0x01, 0x07, 0xfe, 0x73, 0x40, 0x02, 0x4d, 0x0a, 0x13, 0x07, 0x6e,
	0x6e, 0xe2, 0xa5, 0xde, 0x12, 0x3b, 0xd8, 0x4e, 0xa1, 0x37, 0x24, 0xbe, 0x00, 0x9f, 0x89, 0xd3,
	0x8e, 0x1c, 0x39, 0xb2, 0xf2, 0x45, 0x50, 0xec, 0xa4, 0x4d, 0x39, 0x4c, 0x82, 0x9b, 0x9f, 0x27,
	0x8f, 0x7f, 0x79, 0x5f, 0xbf, 0x7a, 0xe1, 0x49, 0x4a, 0xd5, 0xbc, 0x9c, 0x05, 0x31, 0xcf, 0xc3,
	0x14, 0x67, 0xe4, 0xb3, 0x50, 0x21, 0x49, 0x08, 0x8b, 0x39, 0x3b, 0xcd, 0xd3, 0x5c, 0x85, 0xc5,
	0x79, 0x1a, 0xe2, 0x82, 0xca, 0x30, 0xe6, 0x82, 0x84, 0x8b, 0x7d, 0x9c, 0x15, 0x73, 0x5c, 0x39,
	0x41, 0x21, 0xb8, 0xe2, 0xe8, 0x7a, 0xe5, 0x07, 0xb5, 0x3f, 0xde, 0x6b, 0xe3, 0x78, 0xca, 0x43,
	0x1d, 0x9a, 0x95, 0xa7, 0x5a, 0x69, 0xa1, 0x4f, 0xe6, 0xf2, 0xd8, 0x4b, 0x39, 0x4f, 0x33, 0xb2,
	0x49, 0x25, 0xa5, 0xc0, 0x8a, 0x72, 0x66, 0xbe, 0xfb, 0xdf, 0x6d, 0x18, 0x1e, 0xcf, 0xce, 0x48,
	0xac, 0xde, 0x12, 0x85, 0x13, 0xac, 0x30, 0x1a, 0x82, 0x4d, 0x13, 0xd7, 0x9a, 0x58, 0xd3, 0x7e,
	0x64, 0xd3, 0x04, 0x21, 0xe8, 0x30, 0x9c, 0x13, 0xd7, 0xd6, 0x8e, 0x3e, 0xa3, 0xa7, 0xb0, 0x9b,
	0xe1, 0x19, 0xc9, 0xa4, 0xeb, 0x4c, 0x9c, 0xe9, 0xe0, 0x60, 0x1a, 0xb4, 0x8b, 0x0c, 0xb6, 0x89,
	0xc1, 0x1b, 0x1d, 0x7d, 0xc5, 0x94, 0x58, 0x46, 0xf5, 0x3d, 0x74, 0x0c, 0x03, 0xcc, 0x18, 0x57,
	0xba, 0x18, 0xe9, 0x76, 0x34, 0x66, 0xef, 0x4a, 0xcc, 0xb3, 0x4d, 0xde, 0xb0, 0xda, 0x84, 0xf1,
	0x43, 0x18, 0xb4, 0xfe, 0x83, 0x46, 0xe0, 0x9c, 0x93, 0x65, 0xdd, 0x46, 0x75, 0x44, 0xb7, 0xa0,
	0xbb, 0xc0, 0x59, 0xd9, 0x34, 0x62, 0xc4, 0x23, 0xfb, 0x81, 0x35, 0x7e, 0x0c, 0xa3, 0xbf, 0xd9,
	0xff, 0x72, 0xdf, 0x1f, 0xc1, 0xf0, 0x3d, 0x11, 0x92, 0x72, 0x16, 0x91, 0x8f, 0x25, 0x91, 0xca,
	0xff, 0x6a, 0xc1, 0x8d, 0xb5, 0x25, 0x0b, 0xce, 0x24, 0x41, 0x2e, 0xf4, 0x16, 0xc6, 0xaa, 0xa9,
	0x8d, 0xac, 0xc8, 0x39, 0x3e, 0xe3, 0x42, 0x93, 0x9d, 0xc8, 0x08, 0xed, 0x52, 0xc6, 0x85, 0xeb,
	0xd4, 0x6e, 0x25, 0x2a, 0xb7, 0xc0, 0x2a, 0x9e, 0xbb, 0x1d, 0xe3, 0x6a, 0x51, 0xb9, 0x52, 0x61,
	0x45, 0xdc, 0xae, 0xa9, 0x4d, 0x0b, 0xff, 0x10, 0x7a, 0xaf, 0x59, 0x9c, 0x95, 0x09, 0x59, 0x0f,
	0xd1, 0x6a, 0x0d, 0xd1, 0xa0, 0xe6, 0xd2, 0xb5, 0x27, 0x4e, 0x75, 0x49, 0x0b, 0x3f, 0x03, 0x78,
	0xc1, 0x59, 0x42, 0xf5, 0x5b, 0xa0, 0xfb, 0xd0, 0xf9, 0x34, 0x27, 0xa6, 0xe2, 0xc1, 0xc1, 0xed,
	0xed, 0xf9, 0xac, 0x73, 0x91, 0x0e, 0xa1, 0x7d, 0xe8, 0xc9, 0x32, 0x8e, 0x89, 0x94, 0xba, 0x93,
	0x2b, 0xf2, 0x4d, 0xce, 0x3f, 0x81, 0xfe, 0xda, 0x45, 0x77, 0xa1, 0x1f, 0x37, 0xa2, 0xae, 0x74,
	0x63, 0xa0, 0x7b, 0xd0, 0x15, 0x44, 0x89, 0x65, 0xcd, 0xbe, 0xb9, 0xcd, 0x8e, 0xaa, 0x4f, 0x91,
	0x49, 0xf8, 0x27, 0xd0, 0xd5, 0x1a, 0x1d, 0xc1, 0x35, 0xca, 0x14, 0x11, 0x0b, 0x9c, 0xd5, 0x2d,
	0xdc, 0x09, 0xcc, 0x46, 0x04, 0xcd, 0x46, 0x04, 0x2f, 0xeb, 0x8d, 0x88, 0xd6, 0xd1, 0xea, 0x65,
	0x32, 0x9a, 0x53, 0xd5, 0x0c, 0x44, 0x8b, 0xe7, 0xef, 0x2e, 0x2e, 0x3d, 0xeb, 0xe7, 0xa5, 0xb7,
	0xf3, 0x65, 0xe5, 0x59, 0x17, 0x2b, 0xcf, 0xfa, 0xb1, 0xf2, 0xac, 0x5f, 0x2b, 0xcf, 0xfa, 0xf6,
	0xdb, 0xdb, 0xf9, 0x70, 0xf4, 0x5f, 0x7b, 0x3e, 0xdb, 0xd5, 0x75, 0x1c, 0xfe, 0x09, 0x00, 0x00,
	0xff, 0xff, 0xa2, 0xc9, 0xdc, 0x53, 0x27, 0x04, 0x00, 0x00,
}
