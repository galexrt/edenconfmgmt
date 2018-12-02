// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/galexrt/edenconfmgmt/pkg/apis/configs/v1alpha/api.proto

package v1alpha

import (
	context "context"
	fmt "fmt"
	v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha"
	v1alpha1 "github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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

// Config object base.
type Config struct {
	// Metadata for Config object.
	Metadata *v1alpha.ObjectMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Spec for Config.
	Spec                 *ConfigSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fa3e7f761d5b74b, []int{0}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetMetadata() *v1alpha.ObjectMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Config) GetSpec() *ConfigSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type ConfigSpec struct {
	// DataStore.
	DataStore            *DataStore `protobuf:"bytes,1,opt,name=dataStore,proto3" json:"dataStore,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ConfigSpec) Reset()         { *m = ConfigSpec{} }
func (m *ConfigSpec) String() string { return proto.CompactTextString(m) }
func (*ConfigSpec) ProtoMessage()    {}
func (*ConfigSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fa3e7f761d5b74b, []int{1}
}

func (m *ConfigSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigSpec.Unmarshal(m, b)
}
func (m *ConfigSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigSpec.Marshal(b, m, deterministic)
}
func (m *ConfigSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigSpec.Merge(m, src)
}
func (m *ConfigSpec) XXX_Size() int {
	return xxx_messageInfo_ConfigSpec.Size(m)
}
func (m *ConfigSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigSpec proto.InternalMessageInfo

func (m *ConfigSpec) GetDataStore() *DataStore {
	if m != nil {
		return m.DataStore
	}
	return nil
}

// DataStore config for the data store
type DataStore struct {
	// Type.
	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Count.
	Count int64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	// ETCD config options.
	Etcd                 *ETCD    `protobuf:"bytes,3,opt,name=etcd,proto3" json:"etcd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataStore) Reset()         { *m = DataStore{} }
func (m *DataStore) String() string { return proto.CompactTextString(m) }
func (*DataStore) ProtoMessage()    {}
func (*DataStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fa3e7f761d5b74b, []int{2}
}

func (m *DataStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataStore.Unmarshal(m, b)
}
func (m *DataStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataStore.Marshal(b, m, deterministic)
}
func (m *DataStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataStore.Merge(m, src)
}
func (m *DataStore) XXX_Size() int {
	return xxx_messageInfo_DataStore.Size(m)
}
func (m *DataStore) XXX_DiscardUnknown() {
	xxx_messageInfo_DataStore.DiscardUnknown(m)
}

var xxx_messageInfo_DataStore proto.InternalMessageInfo

func (m *DataStore) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DataStore) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *DataStore) GetEtcd() *ETCD {
	if m != nil {
		return m.Etcd
	}
	return nil
}

type ETCD struct {
	// List of endpoints if running your own ETCD server (should only be used for testing).
	Endpoints            string   `protobuf:"bytes,1,opt,name=endpoints,proto3" json:"endpoints,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ETCD) Reset()         { *m = ETCD{} }
func (m *ETCD) String() string { return proto.CompactTextString(m) }
func (*ETCD) ProtoMessage()    {}
func (*ETCD) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fa3e7f761d5b74b, []int{3}
}

func (m *ETCD) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ETCD.Unmarshal(m, b)
}
func (m *ETCD) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ETCD.Marshal(b, m, deterministic)
}
func (m *ETCD) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ETCD.Merge(m, src)
}
func (m *ETCD) XXX_Size() int {
	return xxx_messageInfo_ETCD.Size(m)
}
func (m *ETCD) XXX_DiscardUnknown() {
	xxx_messageInfo_ETCD.DiscardUnknown(m)
}

var xxx_messageInfo_ETCD proto.InternalMessageInfo

func (m *ETCD) GetEndpoints() string {
	if m != nil {
		return m.Endpoints
	}
	return ""
}

// Get
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
	return fileDescriptor_8fa3e7f761d5b74b, []int{4}
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

type GetResponse struct {
	// Config object.
	Config *Config `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
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
	return fileDescriptor_8fa3e7f761d5b74b, []int{5}
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

func (m *GetResponse) GetConfig() *Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *GetResponse) GetError() *v1alpha.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

// List
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
	return fileDescriptor_8fa3e7f761d5b74b, []int{6}
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

type ListResponse struct {
	// Config list.
	Configs []*Config `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty"`
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
	return fileDescriptor_8fa3e7f761d5b74b, []int{7}
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

func (m *ListResponse) GetConfigs() []*Config {
	if m != nil {
		return m.Configs
	}
	return nil
}

func (m *ListResponse) GetError() *v1alpha.Error {
	if m != nil {
		return m.Error
	}
	return nil
}

// Add
type AddRequest struct {
	// Config object.
	Config               *Config  `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fa3e7f761d5b74b, []int{8}
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

func (m *AddRequest) GetConfig() *Config {
	if m != nil {
		return m.Config
	}
	return nil
}

type AddResponse struct {
	// Config object.
	Config *Config `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
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
	return fileDescriptor_8fa3e7f761d5b74b, []int{9}
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

func (m *AddResponse) GetConfig() *Config {
	if m != nil {
		return m.Config
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
	// Config object.
	Config               *Config  `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fa3e7f761d5b74b, []int{10}
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

func (m *UpdateRequest) GetConfig() *Config {
	if m != nil {
		return m.Config
	}
	return nil
}

type UpdateResponse struct {
	// Config object.
	Config *Config `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
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
	return fileDescriptor_8fa3e7f761d5b74b, []int{11}
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

func (m *UpdateResponse) GetConfig() *Config {
	if m != nil {
		return m.Config
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
	// Config object.
	Config               *Config  `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8fa3e7f761d5b74b, []int{12}
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

func (m *DeleteRequest) GetConfig() *Config {
	if m != nil {
		return m.Config
	}
	return nil
}

type DeleteResponse struct {
	// Config object.
	Config *Config `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
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
	return fileDescriptor_8fa3e7f761d5b74b, []int{13}
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

func (m *DeleteResponse) GetConfig() *Config {
	if m != nil {
		return m.Config
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
	return fileDescriptor_8fa3e7f761d5b74b, []int{14}
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
	// Event info for watch response.
	Event *v1alpha1.Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	// Config object.
	Config *Config `protobuf:"bytes,2,opt,name=config,proto3" json:"config,omitempty"`
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
	return fileDescriptor_8fa3e7f761d5b74b, []int{15}
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

func (m *WatchResponse) GetConfig() *Config {
	if m != nil {
		return m.Config
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
	proto.RegisterType((*Config)(nil), "configs.v1alpha.Config")
	proto.RegisterType((*ConfigSpec)(nil), "configs.v1alpha.ConfigSpec")
	proto.RegisterType((*DataStore)(nil), "configs.v1alpha.DataStore")
	proto.RegisterType((*ETCD)(nil), "configs.v1alpha.ETCD")
	proto.RegisterType((*GetRequest)(nil), "configs.v1alpha.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "configs.v1alpha.GetResponse")
	proto.RegisterType((*ListRequest)(nil), "configs.v1alpha.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "configs.v1alpha.ListResponse")
	proto.RegisterType((*AddRequest)(nil), "configs.v1alpha.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "configs.v1alpha.AddResponse")
	proto.RegisterType((*UpdateRequest)(nil), "configs.v1alpha.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "configs.v1alpha.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "configs.v1alpha.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "configs.v1alpha.DeleteResponse")
	proto.RegisterType((*WatchRequest)(nil), "configs.v1alpha.WatchRequest")
	proto.RegisterType((*WatchResponse)(nil), "configs.v1alpha.WatchResponse")
}

func init() {
	proto.RegisterFile("github.com/galexrt/edenconfmgmt/pkg/apis/configs/v1alpha/api.proto", fileDescriptor_8fa3e7f761d5b74b)
}

var fileDescriptor_8fa3e7f761d5b74b = []byte{
	// 679 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0xeb, 0x24, 0x25, 0x93, 0x16, 0xa4, 0x85, 0x8a, 0x60, 0x8a, 0xa9, 0x2c, 0x0e, 0x54,
	0x08, 0x9b, 0x96, 0x4b, 0x24, 0x04, 0xb4, 0x4d, 0x4b, 0x01, 0x15, 0x2a, 0x6d, 0x81, 0x4a, 0x9c,
	0x70, 0xec, 0xad, 0x6b, 0x48, 0xbc, 0xc6, 0xde, 0x14, 0xb8, 0xf1, 0x09, 0x1c, 0xf9, 0x0e, 0xbe,
	0xa2, 0x47, 0x8e, 0x1c, 0x69, 0xf8, 0x11, 0xe4, 0xdd, 0x75, 0x6c, 0xd7, 0x09, 0x90, 0x22, 0xf5,
	0xe6, 0xd9, 0x79, 0xfb, 0xe6, 0xbd, 0xe9, 0xec, 0x34, 0xb0, 0xee, 0xf9, 0xec, 0xa0, 0xdf, 0x31,
	0x1d, 0xda, 0xb3, 0x3c, 0xbb, 0x4b, 0x3e, 0x46, 0xcc, 0x22, 0x2e, 0x09, 0x1c, 0x1a, 0xec, 0xf7,
	0xbc, 0x1e, 0xb3, 0xc2, 0x77, 0x9e, 0x65, 0x87, 0x7e, 0x6c, 0x25, 0x27, 0xbe, 0x17, 0x5b, 0x87,
	0xcb, 0x76, 0x37, 0x3c, 0xb0, 0x93, 0x43, 0x33, 0x8c, 0x28, 0xa3, 0xe8, 0x82, 0x4c, 0x99, 0x32,
	0xa5, 0xdd, 0xce, 0x93, 0x52, 0x8f, 0x5a, 0x1c, 0xd7, 0xe9, 0xef, 0xf3, 0x88, 0x07, 0xfc, 0x4b,
	0xdc, 0xd7, 0x1e, 0x4e, 0xa0, 0x21, 0x22, 0x65, 0x01, 0xda, 0xda, 0x3f, 0x13, 0x90, 0x43, 0x12,
	0xb0, 0x11, 0x1e, 0x8c, 0x18, 0x6a, 0x6d, 0xee, 0x02, 0xb5, 0xe0, 0x5c, 0x8f, 0x30, 0xdb, 0xb5,
	0x99, 0xdd, 0x54, 0x16, 0x95, 0x9b, 0x8d, 0x95, 0x05, 0x33, 0xa9, 0x9b, 0xba, 0x33, 0x77, 0x3a,
	0x6f, 0x89, 0xc3, 0x9e, 0x49, 0x0c, 0x1e, 0xa2, 0x91, 0x05, 0x95, 0x38, 0x24, 0x4e, 0x73, 0x9a,
	0xdf, 0xba, 0x6a, 0x9e, 0x68, 0x8b, 0x29, 0x0a, 0xec, 0x86, 0xc4, 0xc1, 0x1c, 0x68, 0x3c, 0x02,
	0xc8, 0xce, 0x50, 0x0b, 0xea, 0x09, 0xcd, 0x2e, 0xa3, 0x11, 0x91, 0x95, 0xb5, 0x12, 0xc7, 0x46,
	0x8a, 0xc0, 0x19, 0xd8, 0x78, 0x03, 0xf5, 0xe1, 0x39, 0x42, 0x50, 0x61, 0x9f, 0x42, 0xc1, 0x50,
	0xc7, 0xfc, 0x1b, 0x5d, 0x82, 0xaa, 0x43, 0xfb, 0x01, 0xe3, 0xd2, 0x54, 0x2c, 0x02, 0xb4, 0x04,
	0x15, 0xc2, 0x1c, 0xb7, 0xa9, 0xf2, 0x5a, 0xf3, 0xa5, 0x5a, 0x9b, 0x2f, 0xda, 0x1b, 0x98, 0x43,
	0x8c, 0x1b, 0x50, 0x49, 0x22, 0xb4, 0x00, 0x75, 0x12, 0xb8, 0x21, 0xf5, 0x03, 0x16, 0xcb, 0x0a,
	0xd9, 0x41, 0xe2, 0x67, 0x8b, 0x30, 0x4c, 0xde, 0xf7, 0x49, 0xcc, 0x50, 0x0b, 0xc0, 0x23, 0x6c,
	0x27, 0x64, 0x3e, 0x0d, 0x62, 0x69, 0xa8, 0x59, 0x6c, 0xe5, 0xd6, 0x30, 0x8f, 0x73, 0x58, 0xc3,
	0x87, 0x06, 0xe7, 0x89, 0x43, 0x1a, 0xc4, 0x04, 0x59, 0x50, 0x13, 0xd2, 0x24, 0xc9, 0xe5, 0x31,
	0x9d, 0xc5, 0x12, 0x86, 0x96, 0xa0, 0x4a, 0xa2, 0x88, 0x46, 0xf2, 0x2f, 0x71, 0xb1, 0x58, 0x74,
	0x33, 0x49, 0x61, 0x81, 0x30, 0x9e, 0x42, 0x63, 0xdb, 0x8f, 0x87, 0x9a, 0xef, 0x41, 0xa3, 0xeb,
	0xc7, 0x27, 0x44, 0x5f, 0x29, 0xde, 0xdf, 0xce, 0x00, 0x38, 0x8f, 0x36, 0xba, 0x30, 0x2b, 0xb8,
	0xa4, 0xee, 0x65, 0x98, 0x91, 0x42, 0x9b, 0xca, 0xa2, 0xfa, 0x27, 0xe1, 0x29, 0x6e, 0x12, 0xe5,
	0xf7, 0x01, 0xd6, 0x5c, 0x37, 0x15, 0x3e, 0x69, 0x8f, 0x92, 0x1e, 0xf3, 0xeb, 0x67, 0xd0, 0xe3,
	0x55, 0x98, 0x7b, 0x19, 0xba, 0x36, 0x23, 0xa7, 0x16, 0xdb, 0x85, 0xf3, 0x29, 0xc3, 0xd9, 0xe8,
	0xdd, 0x20, 0x5d, 0xf2, 0x7f, 0x7a, 0x53, 0x86, 0x33, 0xd0, 0xfb, 0x1c, 0x66, 0xf7, 0x6c, 0xe6,
	0x1c, 0xa4, 0x72, 0x1f, 0xc0, 0xec, 0x87, 0x24, 0x2e, 0x4e, 0xb1, 0x56, 0x64, 0xd8, 0xcb, 0x21,
	0x70, 0x01, 0x6f, 0x7c, 0x55, 0x60, 0x4e, 0x12, 0x4a, 0xf5, 0xb7, 0xa0, 0xca, 0x37, 0xa7, 0xa4,
	0x9a, 0x37, 0xc5, 0x1e, 0xcd, 0xe4, 0x24, 0x21, 0x16, 0x98, 0x9c, 0xd5, 0xe9, 0x09, 0xad, 0xaa,
	0x7f, 0xb3, 0xba, 0xf2, 0x4d, 0x85, 0x99, 0xb6, 0x7c, 0x2b, 0xab, 0xa0, 0x6e, 0x11, 0x86, 0xca,
	0x7b, 0x36, 0xdb, 0x41, 0xda, 0xc2, 0xe8, 0xa4, 0xb4, 0xd5, 0x86, 0x4a, 0xf2, 0x60, 0x51, 0x19,
	0x95, 0xdb, 0x09, 0xda, 0xb5, 0x31, 0x59, 0x49, 0xb2, 0x0a, 0xea, 0x9a, 0xeb, 0x8e, 0x90, 0x91,
	0xbd, 0xce, 0x11, 0x32, 0xf2, 0x6f, 0xef, 0x09, 0xd4, 0xc4, 0x74, 0x23, 0xbd, 0x84, 0x2b, 0x3c,
	0x1c, 0xed, 0xfa, 0xd8, 0x7c, 0x46, 0x25, 0x06, 0x6f, 0x04, 0x55, 0x61, 0xa6, 0x47, 0x50, 0x9d,
	0x98, 0xd8, 0xc7, 0x50, 0xe5, 0x43, 0x80, 0xca, 0xfe, 0xf3, 0xd3, 0xa6, 0xe9, 0xe3, 0xd2, 0x82,
	0xe7, 0x8e, 0xb2, 0xfe, 0xea, 0xe8, 0x58, 0x57, 0x7e, 0x1c, 0xeb, 0x53, 0x9f, 0x07, 0xba, 0x72,
	0x34, 0xd0, 0x95, 0xef, 0x03, 0x5d, 0xf9, 0x39, 0xd0, 0x95, 0x2f, 0xbf, 0xf4, 0xa9, 0xd7, 0xad,
	0xd3, 0xfe, 0x02, 0xe9, 0xd4, 0xf8, 0xbf, 0xee, 0xbb, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x62,
	0x0b, 0x3c, 0x04, 0xc4, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfigsClient is the client API for Configs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigsClient interface {
	// Get a Config.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// List Configs.
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// Add a Config.
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error)
	// Update a Config.
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Delete a Config.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Watch Config.
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (Configs_WatchClient, error)
}

type configsClient struct {
	cc *grpc.ClientConn
}

func NewConfigsClient(cc *grpc.ClientConn) ConfigsClient {
	return &configsClient{cc}
}

func (c *configsClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/configs.v1alpha.Configs/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configsClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/configs.v1alpha.Configs/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configsClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddResponse, error) {
	out := new(AddResponse)
	err := c.cc.Invoke(ctx, "/configs.v1alpha.Configs/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configsClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/configs.v1alpha.Configs/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configsClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/configs.v1alpha.Configs/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configsClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (Configs_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Configs_serviceDesc.Streams[0], "/configs.v1alpha.Configs/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &configsWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Configs_WatchClient interface {
	Recv() (*WatchResponse, error)
	grpc.ClientStream
}

type configsWatchClient struct {
	grpc.ClientStream
}

func (x *configsWatchClient) Recv() (*WatchResponse, error) {
	m := new(WatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ConfigsServer is the server API for Configs service.
type ConfigsServer interface {
	// Get a Config.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// List Configs.
	List(context.Context, *ListRequest) (*ListResponse, error)
	// Add a Config.
	Add(context.Context, *AddRequest) (*AddResponse, error)
	// Update a Config.
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Delete a Config.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// Watch Config.
	Watch(*WatchRequest, Configs_WatchServer) error
}

func RegisterConfigsServer(s *grpc.Server, srv ConfigsServer) {
	s.RegisterService(&_Configs_serviceDesc, srv)
}

func _Configs_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigsServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configs.v1alpha.Configs/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigsServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configs_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configs.v1alpha.Configs/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigsServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configs_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigsServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configs.v1alpha.Configs/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigsServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configs_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigsServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configs.v1alpha.Configs/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigsServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configs_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigsServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/configs.v1alpha.Configs/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigsServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Configs_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ConfigsServer).Watch(m, &configsWatchServer{stream})
}

type Configs_WatchServer interface {
	Send(*WatchResponse) error
	grpc.ServerStream
}

type configsWatchServer struct {
	grpc.ServerStream
}

func (x *configsWatchServer) Send(m *WatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Configs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "configs.v1alpha.Configs",
	HandlerType: (*ConfigsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Configs_Get_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Configs_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Configs_Add_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Configs_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Configs_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _Configs_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/galexrt/edenconfmgmt/pkg/apis/configs/v1alpha/api.proto",
}
