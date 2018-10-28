// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/galexrt/edenconfmgmt/pkg/apis/eventreactors/v1alpha/api.proto

package v1alpha

import (
	fmt "fmt"
	v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha"
	v1alpha1 "github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	context "golang.org/x/net/context"
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

// EventReactor Reaction to certain events.
type EventReactor struct {
	// Metadata for EventReactor object.
	Metadata *v1alpha.ObjectMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Spec for EventReactor.
	Spec                 *EventReactorSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *EventReactor) Reset()         { *m = EventReactor{} }
func (m *EventReactor) String() string { return proto.CompactTextString(m) }
func (*EventReactor) ProtoMessage()    {}
func (*EventReactor) Descriptor() ([]byte, []int) {
	return fileDescriptor_92ae376702498914, []int{0}
}

func (m *EventReactor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventReactor.Unmarshal(m, b)
}
func (m *EventReactor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventReactor.Marshal(b, m, deterministic)
}
func (m *EventReactor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventReactor.Merge(m, src)
}
func (m *EventReactor) XXX_Size() int {
	return xxx_messageInfo_EventReactor.Size(m)
}
func (m *EventReactor) XXX_DiscardUnknown() {
	xxx_messageInfo_EventReactor.DiscardUnknown(m)
}

var xxx_messageInfo_EventReactor proto.InternalMessageInfo

func (m *EventReactor) GetMetadata() *v1alpha.ObjectMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *EventReactor) GetSpec() *EventReactorSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type EventReactorSpec struct {
	// Reaction.
	Reaction *Reaction `protobuf:"bytes,1,opt,name=reaction,proto3" json:"reaction,omitempty"`
	// Conditions.
	Task *v1alpha.Conditions `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	// TriggerLimiter.
	TriggerLimiter       *TriggerLimiter `protobuf:"bytes,3,opt,name=triggerLimiter,proto3" json:"triggerLimiter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *EventReactorSpec) Reset()         { *m = EventReactorSpec{} }
func (m *EventReactorSpec) String() string { return proto.CompactTextString(m) }
func (*EventReactorSpec) ProtoMessage()    {}
func (*EventReactorSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_92ae376702498914, []int{1}
}

func (m *EventReactorSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventReactorSpec.Unmarshal(m, b)
}
func (m *EventReactorSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventReactorSpec.Marshal(b, m, deterministic)
}
func (m *EventReactorSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventReactorSpec.Merge(m, src)
}
func (m *EventReactorSpec) XXX_Size() int {
	return xxx_messageInfo_EventReactorSpec.Size(m)
}
func (m *EventReactorSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_EventReactorSpec.DiscardUnknown(m)
}

var xxx_messageInfo_EventReactorSpec proto.InternalMessageInfo

func (m *EventReactorSpec) GetReaction() *Reaction {
	if m != nil {
		return m.Reaction
	}
	return nil
}

func (m *EventReactorSpec) GetTask() *v1alpha.Conditions {
	if m != nil {
		return m.Task
	}
	return nil
}

func (m *EventReactorSpec) GetTriggerLimiter() *TriggerLimiter {
	if m != nil {
		return m.TriggerLimiter
	}
	return nil
}

// Reaction currently only jobs can be triggered when an event occurs.
type Reaction struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Reaction) Reset()         { *m = Reaction{} }
func (m *Reaction) String() string { return proto.CompactTextString(m) }
func (*Reaction) ProtoMessage()    {}
func (*Reaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_92ae376702498914, []int{2}
}

func (m *Reaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Reaction.Unmarshal(m, b)
}
func (m *Reaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Reaction.Marshal(b, m, deterministic)
}
func (m *Reaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Reaction.Merge(m, src)
}
func (m *Reaction) XXX_Size() int {
	return xxx_messageInfo_Reaction.Size(m)
}
func (m *Reaction) XXX_DiscardUnknown() {
	xxx_messageInfo_Reaction.DiscardUnknown(m)
}

var xxx_messageInfo_Reaction proto.InternalMessageInfo

// TriggerLimiter trigger limiter options.
type TriggerLimiter struct {
	// MinOccurences
	MinOccurences int64 `protobuf:"varint,1,opt,name=MinOccurences,proto3" json:"MinOccurences,omitempty"`
	// ReReactDelay.
	ReReactDelay *duration.Duration `protobuf:"bytes,2,opt,name=reReactDelay,proto3" json:"reReactDelay,omitempty"`
	// WaitForMore.
	WaitForMore          *WaitForMore `protobuf:"bytes,3,opt,name=waitForMore,proto3" json:"waitForMore,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *TriggerLimiter) Reset()         { *m = TriggerLimiter{} }
func (m *TriggerLimiter) String() string { return proto.CompactTextString(m) }
func (*TriggerLimiter) ProtoMessage()    {}
func (*TriggerLimiter) Descriptor() ([]byte, []int) {
	return fileDescriptor_92ae376702498914, []int{3}
}

func (m *TriggerLimiter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TriggerLimiter.Unmarshal(m, b)
}
func (m *TriggerLimiter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TriggerLimiter.Marshal(b, m, deterministic)
}
func (m *TriggerLimiter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TriggerLimiter.Merge(m, src)
}
func (m *TriggerLimiter) XXX_Size() int {
	return xxx_messageInfo_TriggerLimiter.Size(m)
}
func (m *TriggerLimiter) XXX_DiscardUnknown() {
	xxx_messageInfo_TriggerLimiter.DiscardUnknown(m)
}

var xxx_messageInfo_TriggerLimiter proto.InternalMessageInfo

func (m *TriggerLimiter) GetMinOccurences() int64 {
	if m != nil {
		return m.MinOccurences
	}
	return 0
}

func (m *TriggerLimiter) GetReReactDelay() *duration.Duration {
	if m != nil {
		return m.ReReactDelay
	}
	return nil
}

func (m *TriggerLimiter) GetWaitForMore() *WaitForMore {
	if m != nil {
		return m.WaitForMore
	}
	return nil
}

// WaitForMore trigger limiter WaitForMore options.
type WaitForMore struct {
	// Timeout.
	Timeout *duration.Duration `protobuf:"bytes,1,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// ResetTimeoutOnMore
	ResetTimeoutOnMore   bool     `protobuf:"varint,2,opt,name=resetTimeoutOnMore,proto3" json:"resetTimeoutOnMore,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WaitForMore) Reset()         { *m = WaitForMore{} }
func (m *WaitForMore) String() string { return proto.CompactTextString(m) }
func (*WaitForMore) ProtoMessage()    {}
func (*WaitForMore) Descriptor() ([]byte, []int) {
	return fileDescriptor_92ae376702498914, []int{4}
}

func (m *WaitForMore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WaitForMore.Unmarshal(m, b)
}
func (m *WaitForMore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WaitForMore.Marshal(b, m, deterministic)
}
func (m *WaitForMore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WaitForMore.Merge(m, src)
}
func (m *WaitForMore) XXX_Size() int {
	return xxx_messageInfo_WaitForMore.Size(m)
}
func (m *WaitForMore) XXX_DiscardUnknown() {
	xxx_messageInfo_WaitForMore.DiscardUnknown(m)
}

var xxx_messageInfo_WaitForMore proto.InternalMessageInfo

func (m *WaitForMore) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *WaitForMore) GetResetTimeoutOnMore() bool {
	if m != nil {
		return m.ResetTimeoutOnMore
	}
	return false
}

type AddEventReactorResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddEventReactorResponse) Reset()         { *m = AddEventReactorResponse{} }
func (m *AddEventReactorResponse) String() string { return proto.CompactTextString(m) }
func (*AddEventReactorResponse) ProtoMessage()    {}
func (*AddEventReactorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92ae376702498914, []int{5}
}

func (m *AddEventReactorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddEventReactorResponse.Unmarshal(m, b)
}
func (m *AddEventReactorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddEventReactorResponse.Marshal(b, m, deterministic)
}
func (m *AddEventReactorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddEventReactorResponse.Merge(m, src)
}
func (m *AddEventReactorResponse) XXX_Size() int {
	return xxx_messageInfo_AddEventReactorResponse.Size(m)
}
func (m *AddEventReactorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddEventReactorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddEventReactorResponse proto.InternalMessageInfo

type UpdateEventReactorResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateEventReactorResponse) Reset()         { *m = UpdateEventReactorResponse{} }
func (m *UpdateEventReactorResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateEventReactorResponse) ProtoMessage()    {}
func (*UpdateEventReactorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92ae376702498914, []int{6}
}

func (m *UpdateEventReactorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEventReactorResponse.Unmarshal(m, b)
}
func (m *UpdateEventReactorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEventReactorResponse.Marshal(b, m, deterministic)
}
func (m *UpdateEventReactorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEventReactorResponse.Merge(m, src)
}
func (m *UpdateEventReactorResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateEventReactorResponse.Size(m)
}
func (m *UpdateEventReactorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEventReactorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEventReactorResponse proto.InternalMessageInfo

type DeleteEventReactorResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteEventReactorResponse) Reset()         { *m = DeleteEventReactorResponse{} }
func (m *DeleteEventReactorResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteEventReactorResponse) ProtoMessage()    {}
func (*DeleteEventReactorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92ae376702498914, []int{7}
}

func (m *DeleteEventReactorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEventReactorResponse.Unmarshal(m, b)
}
func (m *DeleteEventReactorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEventReactorResponse.Marshal(b, m, deterministic)
}
func (m *DeleteEventReactorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEventReactorResponse.Merge(m, src)
}
func (m *DeleteEventReactorResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteEventReactorResponse.Size(m)
}
func (m *DeleteEventReactorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEventReactorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEventReactorResponse proto.InternalMessageInfo

type WatchEventReactorResponse struct {
	// Event info for watch response.
	Event *v1alpha1.Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	// EventReactor for watch response.
	EventReactor         *EventReactor `protobuf:"bytes,2,opt,name=EventReactor,proto3" json:"EventReactor,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *WatchEventReactorResponse) Reset()         { *m = WatchEventReactorResponse{} }
func (m *WatchEventReactorResponse) String() string { return proto.CompactTextString(m) }
func (*WatchEventReactorResponse) ProtoMessage()    {}
func (*WatchEventReactorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_92ae376702498914, []int{8}
}

func (m *WatchEventReactorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchEventReactorResponse.Unmarshal(m, b)
}
func (m *WatchEventReactorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchEventReactorResponse.Marshal(b, m, deterministic)
}
func (m *WatchEventReactorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchEventReactorResponse.Merge(m, src)
}
func (m *WatchEventReactorResponse) XXX_Size() int {
	return xxx_messageInfo_WatchEventReactorResponse.Size(m)
}
func (m *WatchEventReactorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchEventReactorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WatchEventReactorResponse proto.InternalMessageInfo

func (m *WatchEventReactorResponse) GetEvent() *v1alpha1.Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *WatchEventReactorResponse) GetEventReactor() *EventReactor {
	if m != nil {
		return m.EventReactor
	}
	return nil
}

func init() {
	proto.RegisterType((*EventReactor)(nil), "eventreactors.v1alpha.EventReactor")
	proto.RegisterType((*EventReactorSpec)(nil), "eventreactors.v1alpha.EventReactorSpec")
	proto.RegisterType((*Reaction)(nil), "eventreactors.v1alpha.Reaction")
	proto.RegisterType((*TriggerLimiter)(nil), "eventreactors.v1alpha.TriggerLimiter")
	proto.RegisterType((*WaitForMore)(nil), "eventreactors.v1alpha.WaitForMore")
	proto.RegisterType((*AddEventReactorResponse)(nil), "eventreactors.v1alpha.AddEventReactorResponse")
	proto.RegisterType((*UpdateEventReactorResponse)(nil), "eventreactors.v1alpha.UpdateEventReactorResponse")
	proto.RegisterType((*DeleteEventReactorResponse)(nil), "eventreactors.v1alpha.DeleteEventReactorResponse")
	proto.RegisterType((*WatchEventReactorResponse)(nil), "eventreactors.v1alpha.WatchEventReactorResponse")
}

func init() {
	proto.RegisterFile("github.com/galexrt/edenconfmgmt/pkg/apis/eventreactors/v1alpha/api.proto", fileDescriptor_92ae376702498914)
}

var fileDescriptor_92ae376702498914 = []byte{
	// 650 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x6b, 0x5a, 0x68, 0x34, 0x2d, 0x2d, 0xac, 0xa8, 0x48, 0xa3, 0x62, 0x90, 0x01, 0x81,
	0x04, 0xd8, 0xfd, 0x73, 0x41, 0xaa, 0x00, 0x15, 0x42, 0xcb, 0x81, 0xa8, 0xd2, 0xb6, 0x50, 0x09,
	0xb8, 0x6c, 0xec, 0xa9, 0x6b, 0xea, 0x78, 0xcd, 0x7a, 0x53, 0xe0, 0xc6, 0x85, 0x3b, 0x3c, 0x0f,
	0x2f, 0xd0, 0x03, 0x07, 0x8e, 0x1c, 0x69, 0x38, 0xf3, 0x0e, 0xc8, 0x9b, 0x4d, 0x1b, 0x27, 0x59,
	0x64, 0xd4, 0x5b, 0x76, 0xe7, 0x37, 0x3b, 0xf3, 0xed, 0x7c, 0x1b, 0xc3, 0xb3, 0x30, 0x92, 0x7b,
	0xed, 0xa6, 0xeb, 0xf3, 0x96, 0x17, 0xb2, 0x18, 0x3f, 0x08, 0xe9, 0x61, 0x80, 0x89, 0xcf, 0x93,
	0xdd, 0x56, 0xd8, 0x92, 0x5e, 0xba, 0x1f, 0x7a, 0x2c, 0x8d, 0x32, 0x0f, 0x0f, 0x30, 0x91, 0x02,
	0x99, 0x2f, 0xb9, 0xc8, 0xbc, 0x83, 0x25, 0x16, 0xa7, 0x7b, 0x2c, 0x0f, 0xb9, 0xa9, 0xe0, 0x92,
	0x93, 0xb9, 0x02, 0xe0, 0x6a, 0xa0, 0x76, 0xaf, 0xbf, 0x00, 0x0f, 0xb9, 0xa7, 0xe8, 0x66, 0x7b,
	0x57, 0xad, 0xd4, 0x42, 0xfd, 0xea, 0x9e, 0x52, 0xb3, 0x43, 0xce, 0xc3, 0x18, 0x4f, 0xa8, 0xa0,
	0x2d, 0x98, 0x8c, 0x78, 0xa2, 0xe3, 0x8f, 0x4a, 0xf7, 0xeb, 0x73, 0x81, 0xc3, 0x6d, 0xd6, 0xd6,
	0xfe, 0x4f, 0xf0, 0x08, 0xa5, 0xce, 0x67, 0x0b, 0xa6, 0x9f, 0xe6, 0x41, 0xda, 0x15, 0x4b, 0xee,
	0x43, 0xa5, 0x85, 0x92, 0x05, 0x4c, 0xb2, 0xaa, 0x75, 0xcd, 0xba, 0x3d, 0xb5, 0xbc, 0xe0, 0xe6,
	0xe5, 0x7b, 0x97, 0xe0, 0x6e, 0x36, 0xdf, 0xa2, 0x2f, 0x1b, 0x9a, 0xa1, 0xc7, 0x34, 0x59, 0x85,
	0x89, 0x2c, 0x45, 0xbf, 0x7a, 0x46, 0x65, 0xdd, 0x72, 0x47, 0xde, 0xa1, 0xdb, 0x5f, 0x6c, 0x2b,
	0x45, 0x9f, 0xaa, 0x24, 0xe7, 0xbb, 0x05, 0x17, 0x06, 0x43, 0x64, 0x15, 0x2a, 0x2a, 0x3f, 0xe2,
	0x89, 0xee, 0xe5, 0xaa, 0xe1, 0x54, 0xaa, 0x31, 0x7a, 0x9c, 0x40, 0xee, 0xc2, 0x84, 0x64, 0xd9,
	0xbe, 0x6e, 0xa7, 0x5a, 0x14, 0xf1, 0x84, 0x27, 0x41, 0x94, 0x63, 0x19, 0x55, 0x14, 0x69, 0xc0,
	0x8c, 0x14, 0x51, 0x18, 0xa2, 0x78, 0x1e, 0xb5, 0x22, 0x89, 0xa2, 0x3a, 0xae, 0xf2, 0x6e, 0x1a,
	0x0a, 0x6e, 0x17, 0x60, 0x3a, 0x90, 0xec, 0x00, 0x54, 0x7a, 0x2d, 0x39, 0xdf, 0x2c, 0x98, 0x29,
	0xe2, 0xe4, 0x06, 0x9c, 0x6f, 0x44, 0xc9, 0xa6, 0xef, 0xb7, 0x05, 0x26, 0x3e, 0x66, 0x4a, 0xdd,
	0x38, 0x2d, 0x6e, 0x92, 0x07, 0x30, 0x2d, 0x50, 0x1d, 0x53, 0xc7, 0x98, 0x7d, 0xd4, 0x4a, 0xe6,
	0xdd, 0xae, 0xad, 0xdc, 0x9e, 0xad, 0xdc, 0xba, 0xb6, 0x15, 0x2d, 0xe0, 0xa4, 0x0e, 0x53, 0xef,
	0x59, 0x24, 0xd7, 0xb9, 0x68, 0x70, 0x81, 0x5a, 0x8f, 0x63, 0xd0, 0xb3, 0x73, 0x42, 0xd2, 0xfe,
	0x34, 0x47, 0xc0, 0x54, 0x5f, 0x8c, 0xac, 0xc0, 0xa4, 0x8c, 0x5a, 0xc8, 0xdb, 0x52, 0x4f, 0xe4,
	0x1f, 0xed, 0xf4, 0x48, 0xe2, 0x02, 0x11, 0x98, 0xa1, 0xdc, 0xee, 0xae, 0x37, 0x13, 0xd5, 0x50,
	0x2e, 0xa7, 0x42, 0x47, 0x44, 0x9c, 0x79, 0xb8, 0xbc, 0x16, 0x04, 0xfd, 0x76, 0xa0, 0x98, 0xa5,
	0x3c, 0xc9, 0xd0, 0x59, 0x80, 0xda, 0x8b, 0x34, 0x60, 0x12, 0x4d, 0xd1, 0x3a, 0xc6, 0x68, 0x88,
	0x7e, 0xb5, 0x60, 0x7e, 0x87, 0x49, 0x7f, 0x6f, 0x54, 0x94, 0xdc, 0x81, 0xb3, 0xea, 0x6a, 0xb4,
	0xae, 0xb9, 0xee, 0x45, 0x0d, 0x1a, 0xb7, 0xcb, 0x90, 0x8d, 0xe2, 0xab, 0xd1, 0xa3, 0xb9, 0x5e,
	0xc2, 0xf3, 0xb4, 0x90, 0xb8, 0xfc, 0x67, 0x02, 0x2e, 0xf5, 0x6f, 0x64, 0x5b, 0x28, 0x0e, 0x22,
	0x1f, 0xc9, 0x3a, 0x4c, 0xbe, 0x44, 0x91, 0xe5, 0x4e, 0x1e, 0x78, 0x80, 0x7a, 0x9b, 0xe2, 0xbb,
	0x36, 0x66, 0xb2, 0x76, 0xc5, 0x10, 0xd5, 0xb2, 0x5e, 0xc3, 0xec, 0x06, 0xca, 0xc2, 0x13, 0x2f,
	0xd3, 0x66, 0xad, 0x0c, 0x44, 0x76, 0x61, 0x76, 0x60, 0x50, 0xe5, 0x0e, 0x77, 0x0d, 0x90, 0x61,
	0xea, 0x24, 0x06, 0x32, 0x3c, 0xf5, 0x72, 0xa5, 0x96, 0x0c, 0x90, 0xd9, 0x45, 0x79, 0xb5, 0x61,
	0x17, 0x9d, 0xae, 0x9a, 0xd9, 0x95, 0x24, 0x86, 0x8b, 0x43, 0xa6, 0x2c, 0x57, 0x6c, 0xd1, 0xf8,
	0x96, 0x0d, 0x1e, 0x5f, 0xb4, 0x1e, 0xbf, 0x39, 0x3c, 0xb2, 0xad, 0x9f, 0x47, 0xf6, 0xd8, 0xa7,
	0x8e, 0x6d, 0x1d, 0x76, 0x6c, 0xeb, 0x47, 0xc7, 0xb6, 0x7e, 0x75, 0x6c, 0xeb, 0xcb, 0x6f, 0x7b,
	0xec, 0xd5, 0xc3, 0xd3, 0x7d, 0x41, 0x9b, 0xe7, 0xd4, 0x9f, 0xc0, 0xca, 0xdf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xf1, 0x32, 0xef, 0x0d, 0x8a, 0x07, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventReactorsServiceClient is the client API for EventReactorsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventReactorsServiceClient interface {
	// Version returns the API version.
	Version(ctx context.Context, in *v1alpha.VersionRequest, opts ...grpc.CallOption) (*v1alpha.VersionResponse, error)
	// Get a EventReactor.
	GetEventReactor(ctx context.Context, in *EventReactor, opts ...grpc.CallOption) (*EventReactor, error)
	// Add a EventReactor.
	AddEventReactor(ctx context.Context, in *EventReactor, opts ...grpc.CallOption) (*AddEventReactorResponse, error)
	// Update a EventReactor.
	UpdateEventReactor(ctx context.Context, in *EventReactor, opts ...grpc.CallOption) (*UpdateEventReactorResponse, error)
	// Delete a EventReactor.
	DeleteEventReactor(ctx context.Context, in *EventReactor, opts ...grpc.CallOption) (*DeleteEventReactorResponse, error)
	// Watch a EventReactor.
	WatchEventReactor(ctx context.Context, in *EventReactor, opts ...grpc.CallOption) (EventReactorsService_WatchEventReactorClient, error)
}

type eventReactorsServiceClient struct {
	cc *grpc.ClientConn
}

func NewEventReactorsServiceClient(cc *grpc.ClientConn) EventReactorsServiceClient {
	return &eventReactorsServiceClient{cc}
}

func (c *eventReactorsServiceClient) Version(ctx context.Context, in *v1alpha.VersionRequest, opts ...grpc.CallOption) (*v1alpha.VersionResponse, error) {
	out := new(v1alpha.VersionResponse)
	err := c.cc.Invoke(ctx, "/eventreactors.v1alpha.EventReactorsService/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventReactorsServiceClient) GetEventReactor(ctx context.Context, in *EventReactor, opts ...grpc.CallOption) (*EventReactor, error) {
	out := new(EventReactor)
	err := c.cc.Invoke(ctx, "/eventreactors.v1alpha.EventReactorsService/GetEventReactor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventReactorsServiceClient) AddEventReactor(ctx context.Context, in *EventReactor, opts ...grpc.CallOption) (*AddEventReactorResponse, error) {
	out := new(AddEventReactorResponse)
	err := c.cc.Invoke(ctx, "/eventreactors.v1alpha.EventReactorsService/AddEventReactor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventReactorsServiceClient) UpdateEventReactor(ctx context.Context, in *EventReactor, opts ...grpc.CallOption) (*UpdateEventReactorResponse, error) {
	out := new(UpdateEventReactorResponse)
	err := c.cc.Invoke(ctx, "/eventreactors.v1alpha.EventReactorsService/UpdateEventReactor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventReactorsServiceClient) DeleteEventReactor(ctx context.Context, in *EventReactor, opts ...grpc.CallOption) (*DeleteEventReactorResponse, error) {
	out := new(DeleteEventReactorResponse)
	err := c.cc.Invoke(ctx, "/eventreactors.v1alpha.EventReactorsService/DeleteEventReactor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventReactorsServiceClient) WatchEventReactor(ctx context.Context, in *EventReactor, opts ...grpc.CallOption) (EventReactorsService_WatchEventReactorClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EventReactorsService_serviceDesc.Streams[0], "/eventreactors.v1alpha.EventReactorsService/WatchEventReactor", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventReactorsServiceWatchEventReactorClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventReactorsService_WatchEventReactorClient interface {
	Recv() (*WatchEventReactorResponse, error)
	grpc.ClientStream
}

type eventReactorsServiceWatchEventReactorClient struct {
	grpc.ClientStream
}

func (x *eventReactorsServiceWatchEventReactorClient) Recv() (*WatchEventReactorResponse, error) {
	m := new(WatchEventReactorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EventReactorsServiceServer is the server API for EventReactorsService service.
type EventReactorsServiceServer interface {
	// Version returns the API version.
	Version(context.Context, *v1alpha.VersionRequest) (*v1alpha.VersionResponse, error)
	// Get a EventReactor.
	GetEventReactor(context.Context, *EventReactor) (*EventReactor, error)
	// Add a EventReactor.
	AddEventReactor(context.Context, *EventReactor) (*AddEventReactorResponse, error)
	// Update a EventReactor.
	UpdateEventReactor(context.Context, *EventReactor) (*UpdateEventReactorResponse, error)
	// Delete a EventReactor.
	DeleteEventReactor(context.Context, *EventReactor) (*DeleteEventReactorResponse, error)
	// Watch a EventReactor.
	WatchEventReactor(*EventReactor, EventReactorsService_WatchEventReactorServer) error
}

func RegisterEventReactorsServiceServer(s *grpc.Server, srv EventReactorsServiceServer) {
	s.RegisterService(&_EventReactorsService_serviceDesc, srv)
}

func _EventReactorsService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha.VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventReactorsServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventreactors.v1alpha.EventReactorsService/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventReactorsServiceServer).Version(ctx, req.(*v1alpha.VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventReactorsService_GetEventReactor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventReactor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventReactorsServiceServer).GetEventReactor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventreactors.v1alpha.EventReactorsService/GetEventReactor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventReactorsServiceServer).GetEventReactor(ctx, req.(*EventReactor))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventReactorsService_AddEventReactor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventReactor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventReactorsServiceServer).AddEventReactor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventreactors.v1alpha.EventReactorsService/AddEventReactor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventReactorsServiceServer).AddEventReactor(ctx, req.(*EventReactor))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventReactorsService_UpdateEventReactor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventReactor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventReactorsServiceServer).UpdateEventReactor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventreactors.v1alpha.EventReactorsService/UpdateEventReactor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventReactorsServiceServer).UpdateEventReactor(ctx, req.(*EventReactor))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventReactorsService_DeleteEventReactor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EventReactor)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventReactorsServiceServer).DeleteEventReactor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventreactors.v1alpha.EventReactorsService/DeleteEventReactor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventReactorsServiceServer).DeleteEventReactor(ctx, req.(*EventReactor))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventReactorsService_WatchEventReactor_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EventReactor)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventReactorsServiceServer).WatchEventReactor(m, &eventReactorsServiceWatchEventReactorServer{stream})
}

type EventReactorsService_WatchEventReactorServer interface {
	Send(*WatchEventReactorResponse) error
	grpc.ServerStream
}

type eventReactorsServiceWatchEventReactorServer struct {
	grpc.ServerStream
}

func (x *eventReactorsServiceWatchEventReactorServer) Send(m *WatchEventReactorResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _EventReactorsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eventreactors.v1alpha.EventReactorsService",
	HandlerType: (*EventReactorsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _EventReactorsService_Version_Handler,
		},
		{
			MethodName: "GetEventReactor",
			Handler:    _EventReactorsService_GetEventReactor_Handler,
		},
		{
			MethodName: "AddEventReactor",
			Handler:    _EventReactorsService_AddEventReactor_Handler,
		},
		{
			MethodName: "UpdateEventReactor",
			Handler:    _EventReactorsService_UpdateEventReactor_Handler,
		},
		{
			MethodName: "DeleteEventReactor",
			Handler:    _EventReactorsService_DeleteEventReactor_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchEventReactor",
			Handler:       _EventReactorsService_WatchEventReactor_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/galexrt/edenconfmgmt/pkg/apis/eventreactors/v1alpha/api.proto",
}
