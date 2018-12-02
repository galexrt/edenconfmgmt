// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/galexrt/edenconfmgmt/pkg/apis/beacons/v1alpha/api.proto

package v1alpha

import (
	context "context"
	fmt "fmt"
	v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha"
	v1alpha1 "github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	grpc "google.golang.org/grpc"
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

type Beacon struct {
	// Metadata for Beacon object.
	Metadata *v1alpha.ObjectMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Spec for Beacon.
	Spec                 *BeaconSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Beacon) Reset()         { *m = Beacon{} }
func (m *Beacon) String() string { return proto.CompactTextString(m) }
func (*Beacon) ProtoMessage()    {}
func (*Beacon) Descriptor() ([]byte, []int) {
	return fileDescriptor_d429cd6d4b65179a, []int{0}
}

func (m *Beacon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Beacon.Unmarshal(m, b)
}
func (m *Beacon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Beacon.Marshal(b, m, deterministic)
}
func (m *Beacon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Beacon.Merge(m, src)
}
func (m *Beacon) XXX_Size() int {
	return xxx_messageInfo_Beacon.Size(m)
}
func (m *Beacon) XXX_DiscardUnknown() {
	xxx_messageInfo_Beacon.DiscardUnknown(m)
}

var xxx_messageInfo_Beacon proto.InternalMessageInfo

func (m *Beacon) GetMetadata() *v1alpha.ObjectMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Beacon) GetSpec() *BeaconSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

// TODO
type BeaconSpec struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BeaconSpec) Reset()         { *m = BeaconSpec{} }
func (m *BeaconSpec) String() string { return proto.CompactTextString(m) }
func (*BeaconSpec) ProtoMessage()    {}
func (*BeaconSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_d429cd6d4b65179a, []int{1}
}

func (m *BeaconSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BeaconSpec.Unmarshal(m, b)
}
func (m *BeaconSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BeaconSpec.Marshal(b, m, deterministic)
}
func (m *BeaconSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BeaconSpec.Merge(m, src)
}
func (m *BeaconSpec) XXX_Size() int {
	return xxx_messageInfo_BeaconSpec.Size(m)
}
func (m *BeaconSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_BeaconSpec.DiscardUnknown(m)
}

var xxx_messageInfo_BeaconSpec proto.InternalMessageInfo

// Get request content.
type GetRequest struct {
	// GetOptions options for a GetRequest.
	GetOptions           *v1alpha.GetOptions `protobuf:"bytes,1,opt,name=getOptions,proto3" json:"getOptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d429cd6d4b65179a, []int{2}
}

func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (m *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(m, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

func (m *GetRequest) GetGetOptions() *v1alpha.GetOptions {
	if m != nil {
		return m.GetOptions
	}
	return nil
}

// Get response content.
type GetResponse struct {
	// Beacon object.
	Beacon *Beacon `protobuf:"bytes,1,opt,name=beacon,proto3" json:"beacon,omitempty"`
	// Error object.
	Error                *v1alpha.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetResponse) Reset()         { *m = GetResponse{} }
func (m *GetResponse) String() string { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()    {}
func (*GetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d429cd6d4b65179a, []int{3}
}

func (m *GetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetResponse.Unmarshal(m, b)
}
func (m *GetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetResponse.Marshal(b, m, deterministic)
}
func (m *GetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetResponse.Merge(m, src)
}
func (m *GetResponse) XXX_Size() int {
	return xxx_messageInfo_GetResponse.Size(m)
}
func (m *GetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetResponse proto.InternalMessageInfo

func (m *GetResponse) GetBeacon() *Beacon {
	if m != nil {
		return m.Beacon
	}
	return nil
}

func (m *GetResponse) GetError() *v1alpha.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

// List request content.
type ListRequest struct {
	// ListOptions options for a ListRequest.
	ListOptions          *v1alpha.ListOptions `protobuf:"bytes,1,opt,name=listOptions,proto3" json:"listOptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d429cd6d4b65179a, []int{4}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetListOptions() *v1alpha.ListOptions {
	if m != nil {
		return m.ListOptions
	}
	return nil
}

// List response content.
type ListResponse struct {
	// Beacon list.
	Beacons []*Beacon `protobuf:"bytes,1,rep,name=beacons,proto3" json:"beacons,omitempty"`
	// Error object.
	Error                *v1alpha.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d429cd6d4b65179a, []int{5}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetBeacons() []*Beacon {
	if m != nil {
		return m.Beacons
	}
	return nil
}

func (m *ListResponse) GetError() *v1alpha.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

// Add Beacon request.
type AddRequest struct {
	// Beacon object.
	Beacon               *Beacon  `protobuf:"bytes,1,opt,name=beacon,proto3" json:"beacon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d429cd6d4b65179a, []int{6}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetBeacon() *Beacon {
	if m != nil {
		return m.Beacon
	}
	return nil
}

type AddResponse struct {
	// Beacon object.
	Beacon *Beacon `protobuf:"bytes,1,opt,name=beacon,proto3" json:"beacon,omitempty"`
	// Error object.
	Error                *v1alpha.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d429cd6d4b65179a, []int{7}
}

func (m *AddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResponse.Unmarshal(m, b)
}
func (m *AddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResponse.Marshal(b, m, deterministic)
}
func (m *AddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResponse.Merge(m, src)
}
func (m *AddResponse) XXX_Size() int {
	return xxx_messageInfo_AddResponse.Size(m)
}
func (m *AddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddResponse proto.InternalMessageInfo

func (m *AddResponse) GetBeacon() *Beacon {
	if m != nil {
		return m.Beacon
	}
	return nil
}

func (m *AddResponse) GetError() *v1alpha.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

// Update
type UpdateRequest struct {
	// Beacon object.
	Beacon               *Beacon  `protobuf:"bytes,1,opt,name=beacon,proto3" json:"beacon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d429cd6d4b65179a, []int{8}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetBeacon() *Beacon {
	if m != nil {
		return m.Beacon
	}
	return nil
}

type UpdateResponse struct {
	// Beacon object.
	Beacon *Beacon `protobuf:"bytes,1,opt,name=beacon,proto3" json:"beacon,omitempty"`
	// Error object.
	Error                *v1alpha.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d429cd6d4b65179a, []int{9}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetBeacon() *Beacon {
	if m != nil {
		return m.Beacon
	}
	return nil
}

func (m *UpdateResponse) GetError() *v1alpha.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

// Delete
type DeleteRequest struct {
	// Beacon object.
	Beacon               *Beacon  `protobuf:"bytes,1,opt,name=beacon,proto3" json:"beacon,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d429cd6d4b65179a, []int{10}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetBeacon() *Beacon {
	if m != nil {
		return m.Beacon
	}
	return nil
}

type DeleteResponse struct {
	// Beacon object.
	Beacon *Beacon `protobuf:"bytes,1,opt,name=beacon,proto3" json:"beacon,omitempty"`
	// Error object.
	Error                *v1alpha.Error `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d429cd6d4b65179a, []int{11}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

func (m *DeleteResponse) GetBeacon() *Beacon {
	if m != nil {
		return m.Beacon
	}
	return nil
}

func (m *DeleteResponse) GetError() *v1alpha.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

// Watch
type WatchRequest struct {
	// WatchOptions options for WatchRequest.
	WatchOptions         *v1alpha.WatchOptions `protobuf:"bytes,1,opt,name=watchOptions,proto3" json:"watchOptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *WatchRequest) Reset()         { *m = WatchRequest{} }
func (m *WatchRequest) String() string { return proto.CompactTextString(m) }
func (*WatchRequest) ProtoMessage()    {}
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d429cd6d4b65179a, []int{12}
}

func (m *WatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchRequest.Unmarshal(m, b)
}
func (m *WatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchRequest.Marshal(b, m, deterministic)
}
func (m *WatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchRequest.Merge(m, src)
}
func (m *WatchRequest) XXX_Size() int {
	return xxx_messageInfo_WatchRequest.Size(m)
}
func (m *WatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchRequest proto.InternalMessageInfo

func (m *WatchRequest) GetWatchOptions() *v1alpha.WatchOptions {
	if m != nil {
		return m.WatchOptions
	}
	return nil
}

type WatchResponse struct {
	// Beacon info for watch response.
	Event *v1alpha1.Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	// Beacon for watch response.
	Beacon *Beacon `protobuf:"bytes,2,opt,name=beacon,proto3" json:"beacon,omitempty"`
	// Error object.
	Error                *v1alpha.Error `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *WatchResponse) Reset()         { *m = WatchResponse{} }
func (m *WatchResponse) String() string { return proto.CompactTextString(m) }
func (*WatchResponse) ProtoMessage()    {}
func (*WatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d429cd6d4b65179a, []int{13}
}

func (m *WatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchResponse.Unmarshal(m, b)
}
func (m *WatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchResponse.Marshal(b, m, deterministic)
}
func (m *WatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchResponse.Merge(m, src)
}
func (m *WatchResponse) XXX_Size() int {
	return xxx_messageInfo_WatchResponse.Size(m)
}
func (m *WatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WatchResponse proto.InternalMessageInfo

func (m *WatchResponse) GetEvent() *v1alpha1.Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *WatchResponse) GetBeacon() *Beacon {
	if m != nil {
		return m.Beacon
	}
	return nil
}

func (m *WatchResponse) GetError() *v1alpha.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*Beacon)(nil), "beacons.v1alpha.Beacon")
	proto.RegisterType((*BeaconSpec)(nil), "beacons.v1alpha.BeaconSpec")
	proto.RegisterType((*GetRequest)(nil), "beacons.v1alpha.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "beacons.v1alpha.GetResponse")
	proto.RegisterType((*ListRequest)(nil), "beacons.v1alpha.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "beacons.v1alpha.ListResponse")
	proto.RegisterType((*AddRequest)(nil), "beacons.v1alpha.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "beacons.v1alpha.AddResponse")
	proto.RegisterType((*UpdateRequest)(nil), "beacons.v1alpha.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "beacons.v1alpha.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "beacons.v1alpha.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "beacons.v1alpha.DeleteResponse")
	proto.RegisterType((*WatchRequest)(nil), "beacons.v1alpha.WatchRequest")
	proto.RegisterType((*WatchResponse)(nil), "beacons.v1alpha.WatchResponse")
}

func init() {
	proto.RegisterFile("github.com/galexrt/edenconfmgmt/pkg/apis/beacons/v1alpha/api.proto", fileDescriptor_d429cd6d4b65179a)
}

var fileDescriptor_d429cd6d4b65179a = []byte{
	// 608 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x95, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0xeb, 0xa6, 0x49, 0xd1, 0x24, 0x05, 0xc9, 0x08, 0x91, 0x9a, 0x62, 0x2a, 0x9f, 0x40,
	0x08, 0x9b, 0x96, 0x4b, 0x24, 0x04, 0x24, 0xe1, 0x4f, 0x00, 0x15, 0x2a, 0x19, 0x41, 0x25, 0x6e,
	0x1b, 0x7b, 0xea, 0x18, 0x1c, 0xaf, 0xb1, 0x37, 0x05, 0x6e, 0x3c, 0x02, 0x47, 0x9e, 0x83, 0xa7,
	0xe8, 0x91, 0x23, 0x47, 0x1a, 0x5e, 0x04, 0x79, 0xbd, 0x8e, 0xed, 0xc4, 0x81, 0xa6, 0x48, 0xb9,
	0x79, 0x3d, 0xdf, 0x7e, 0xfb, 0x9b, 0xd5, 0xcc, 0x2c, 0x74, 0x1d, 0x97, 0x0d, 0x46, 0x7d, 0xdd,
	0xa2, 0x43, 0xc3, 0x21, 0x1e, 0x7e, 0x0a, 0x99, 0x81, 0x36, 0xfa, 0x16, 0xf5, 0x0f, 0x87, 0xce,
	0x90, 0x19, 0xc1, 0x7b, 0xc7, 0x20, 0x81, 0x1b, 0x19, 0x7d, 0x24, 0x16, 0xf5, 0x23, 0xe3, 0x68,
	0x87, 0x78, 0xc1, 0x80, 0xc4, 0x3f, 0xf5, 0x20, 0xa4, 0x8c, 0xca, 0x17, 0x44, 0x48, 0x17, 0x21,
	0xe5, 0x56, 0xde, 0x94, 0x3a, 0xd4, 0xe0, 0xba, 0xfe, 0xe8, 0x90, 0xaf, 0xf8, 0x82, 0x7f, 0x25,
	0xfb, 0x95, 0x4d, 0x87, 0x52, 0xc7, 0xc3, 0x4c, 0x45, 0xfc, 0xcf, 0x22, 0xf4, 0xe0, 0xd4, 0x78,
	0x16, 0x0d, 0x71, 0x96, 0x4d, 0xe9, 0x9c, 0xda, 0x00, 0x8f, 0xd0, 0x67, 0x25, 0xe9, 0x69, 0x11,
	0xd4, 0xba, 0x3c, 0x41, 0xb9, 0x05, 0xe7, 0x86, 0xc8, 0x88, 0x4d, 0x18, 0x69, 0x4a, 0xdb, 0xd2,
	0xf5, 0xfa, 0xee, 0x96, 0x1e, 0x9f, 0x9b, 0x26, 0xae, 0xef, 0xf7, 0xdf, 0xa1, 0xc5, 0x5e, 0x08,
	0x8d, 0x39, 0x51, 0xcb, 0x06, 0xac, 0x45, 0x01, 0x5a, 0xcd, 0x55, 0xbe, 0xeb, 0x8a, 0x3e, 0x75,
	0x63, 0x7a, 0x72, 0xc0, 0xab, 0x00, 0x2d, 0x93, 0x0b, 0xb5, 0x06, 0x40, 0xf6, 0x4f, 0x7b, 0x02,
	0xd0, 0x43, 0x66, 0xe2, 0x87, 0x11, 0x46, 0x4c, 0x6e, 0x01, 0x38, 0xc8, 0xf6, 0x03, 0xe6, 0x52,
	0x3f, 0x12, 0x20, 0xcd, 0x22, 0x48, 0x6f, 0x12, 0x37, 0x73, 0x5a, 0xcd, 0x85, 0x3a, 0xf7, 0x89,
	0x02, 0xea, 0x47, 0x28, 0x1b, 0x50, 0x4b, 0x40, 0x84, 0xc9, 0xe5, 0x39, 0x5c, 0xa6, 0x90, 0xc9,
	0x37, 0xa0, 0x8a, 0x61, 0x48, 0x43, 0x91, 0xc7, 0xc5, 0xe2, 0xa1, 0x8f, 0xe3, 0x90, 0x99, 0x28,
	0xb4, 0xe7, 0x50, 0xdf, 0x73, 0xa3, 0x09, 0xf3, 0x5d, 0xa8, 0x7b, 0x6e, 0x34, 0x05, 0xbd, 0x59,
	0xdc, 0xbf, 0x97, 0x09, 0xcc, 0xbc, 0x5a, 0xf3, 0xa0, 0x91, 0x78, 0x09, 0xee, 0x1d, 0x58, 0x17,
	0xa0, 0x4d, 0x69, 0xbb, 0xf2, 0x37, 0xf0, 0x54, 0xb7, 0x08, 0xf9, 0x3d, 0x80, 0x8e, 0x6d, 0xa7,
	0xe0, 0x8b, 0xde, 0x51, 0x7c, 0xc7, 0x7c, 0xfb, 0x12, 0xee, 0xb8, 0x0d, 0x1b, 0xaf, 0x03, 0x9b,
	0x30, 0x3c, 0x33, 0xac, 0x07, 0xe7, 0x53, 0x87, 0xe5, 0xf0, 0x3e, 0x42, 0x0f, 0xff, 0x8f, 0x37,
	0x75, 0x58, 0x02, 0xef, 0x4b, 0x68, 0x1c, 0x10, 0x66, 0x0d, 0x52, 0xdc, 0xfb, 0xd0, 0xf8, 0x18,
	0xaf, 0x8b, 0x55, 0xac, 0x14, 0x1d, 0x0e, 0x72, 0x0a, 0xb3, 0xa0, 0xd7, 0xbe, 0x49, 0xb0, 0x21,
	0x0c, 0x05, 0xfd, 0x4d, 0xa8, 0xf2, 0xb9, 0x23, 0xac, 0x2e, 0xe9, 0xc9, 0x14, 0xca, 0x70, 0xe2,
	0xa5, 0x99, 0x68, 0x72, 0xa9, 0xae, 0x2e, 0x98, 0x6a, 0xe5, 0x5f, 0xa9, 0xee, 0x7e, 0xaf, 0xc0,
	0x7a, 0x57, 0xf4, 0x4a, 0x1b, 0x2a, 0x3d, 0x64, 0xf2, 0xec, 0x94, 0xca, 0x66, 0x90, 0xb2, 0x55,
	0x1e, 0x14, 0x69, 0x3d, 0x84, 0xb5, 0xb8, 0x61, 0xe5, 0x59, 0x55, 0x6e, 0x26, 0x28, 0x57, 0xe7,
	0x44, 0x85, 0x49, 0x1b, 0x2a, 0x1d, 0xdb, 0x2e, 0xc1, 0xc8, 0xba, 0xb3, 0x04, 0x23, 0xdf, 0x7b,
	0xcf, 0xa0, 0x96, 0x54, 0xb7, 0xac, 0xce, 0xe8, 0x0a, 0x8d, 0xa3, 0x5c, 0x9b, 0x1b, 0xcf, 0xac,
	0x92, 0xc2, 0x2b, 0xb1, 0x2a, 0xd4, 0x74, 0x89, 0xd5, 0x54, 0xc5, 0x3e, 0x85, 0x2a, 0x2f, 0x02,
	0x79, 0x36, 0xff, 0x7c, 0xb5, 0x29, 0xea, 0xbc, 0x70, 0xe2, 0x73, 0x5b, 0xea, 0xbe, 0x39, 0x3e,
	0x51, 0xa5, 0x9f, 0x27, 0xea, 0xca, 0x97, 0xb1, 0x2a, 0x1d, 0x8f, 0x55, 0xe9, 0xc7, 0x58, 0x95,
	0x7e, 0x8d, 0x55, 0xe9, 0xeb, 0x6f, 0x75, 0xe5, 0x6d, 0xeb, 0xac, 0x4f, 0x7b, 0xbf, 0xc6, 0x1f,
	0xbe, 0x3b, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xfd, 0x1a, 0x25, 0x1d, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BeaconsClient is the client API for Beacons service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BeaconsClient interface {
	// Get a Beacon.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// List Beacons.
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Add a Beacon.
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	// Update a Beacon.
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete a Beacon.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Watch Beacon.
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (Beacons_WatchClient, error)
}

type beaconsClient struct {
	cc *grpc.ClientConn
}

func NewBeaconsClient(cc *grpc.ClientConn) BeaconsClient {
	return &beaconsClient{cc}
}

func (c *beaconsClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/beacons.v1alpha.Beacons/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconsClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/beacons.v1alpha.Beacons/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconsClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/beacons.v1alpha.Beacons/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconsClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/beacons.v1alpha.Beacons/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconsClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/beacons.v1alpha.Beacons/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *beaconsClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (Beacons_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Beacons_serviceDesc.Streams[0], "/beacons.v1alpha.Beacons/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &beaconsWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Beacons_WatchClient interface {
	Recv() (*WatchResponse, error)
	grpc.ClientStream
}

type beaconsWatchClient struct {
	grpc.ClientStream
}

func (x *beaconsWatchClient) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BeaconsServer is the server API for Beacons service.
type BeaconsServer interface {
	// Get a Beacon.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// List Beacons.
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Add a Beacon.
	Add(context.Context, *AddRequest) (*AddResponse, error)
	// Update a Beacon.
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete a Beacon.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// Watch Beacon.
	Watch(*WatchRequest, Beacons_WatchServer) error
}

func RegisterBeaconsServer(s *grpc.Server, srv BeaconsServer) {
	s.RegisterService(&_Beacons_serviceDesc, srv)
}

func _Beacons_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/beacons.v1alpha.Beacons/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconsServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Beacons_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/beacons.v1alpha.Beacons/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconsServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Beacons_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconsServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/beacons.v1alpha.Beacons/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconsServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Beacons_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/beacons.v1alpha.Beacons/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconsServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Beacons_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BeaconsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/beacons.v1alpha.Beacons/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BeaconsServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Beacons_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BeaconsServer).Watch(m, &beaconsWatchServer{stream})
}

type Beacons_WatchServer interface {
	Send(*WatchResponse) error
	grpc.ServerStream
}

type beaconsWatchServer struct {
	grpc.ServerStream
}

func (x *beaconsWatchServer) Send(m *WatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Beacons_serviceDesc = grpc.ServiceDesc{
	ServiceName: "beacons.v1alpha.Beacons",
	HandlerType: (*BeaconsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Beacons_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Beacons_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Beacons_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Beacons_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Beacons_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Beacons_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/galexrt/edenconfmgmt/pkg/apis/beacons/v1alpha/api.proto",
}
