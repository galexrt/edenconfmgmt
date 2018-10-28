// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/galexrt/edenconfmgmt/pkg/apis/tasks/v1alpha/api.proto

package v1alpha

import (
	fmt "fmt"
	v1alpha "github.com/galexrt/edenconfmgmt/pkg/apis/core/v1alpha"
	v1alpha1 "github.com/galexrt/edenconfmgmt/pkg/apis/events/v1alpha"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Task struct {
	// Metadata for Task object.
	Metadata *v1alpha.ObjectMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Spec for Task.
	Spec                 *TaskSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8693f33936b10a7, []int{0}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetMetadata() *v1alpha.ObjectMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Task) GetSpec() *TaskSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type TaskSpec struct {
	// RunOptions.
	RunOptions *RunOptions `protobuf:"bytes,1,opt,name=runOptions,proto3" json:"runOptions,omitempty"`
	// Steps in this Task.
	Steps                []*Step  `protobuf:"bytes,2,rep,name=steps,proto3" json:"steps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskSpec) Reset()         { *m = TaskSpec{} }
func (m *TaskSpec) String() string { return proto.CompactTextString(m) }
func (*TaskSpec) ProtoMessage()    {}
func (*TaskSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8693f33936b10a7, []int{1}
}

func (m *TaskSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskSpec.Unmarshal(m, b)
}
func (m *TaskSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskSpec.Marshal(b, m, deterministic)
}
func (m *TaskSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskSpec.Merge(m, src)
}
func (m *TaskSpec) XXX_Size() int {
	return xxx_messageInfo_TaskSpec.Size(m)
}
func (m *TaskSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskSpec.DiscardUnknown(m)
}

var xxx_messageInfo_TaskSpec proto.InternalMessageInfo

func (m *TaskSpec) GetRunOptions() *RunOptions {
	if m != nil {
		return m.RunOptions
	}
	return nil
}

func (m *TaskSpec) GetSteps() []*Step {
	if m != nil {
		return m.Steps
	}
	return nil
}

type Step struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Spec                 *Action  `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Step) Reset()         { *m = Step{} }
func (m *Step) String() string { return proto.CompactTextString(m) }
func (*Step) ProtoMessage()    {}
func (*Step) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8693f33936b10a7, []int{2}
}

func (m *Step) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Step.Unmarshal(m, b)
}
func (m *Step) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Step.Marshal(b, m, deterministic)
}
func (m *Step) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Step.Merge(m, src)
}
func (m *Step) XXX_Size() int {
	return xxx_messageInfo_Step.Size(m)
}
func (m *Step) XXX_DiscardUnknown() {
	xxx_messageInfo_Step.DiscardUnknown(m)
}

var xxx_messageInfo_Step proto.InternalMessageInfo

func (m *Step) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Step) GetSpec() *Action {
	if m != nil {
		return m.Spec
	}
	return nil
}

type RunOptions struct {
	Sync                 bool       `protobuf:"varint,1,opt,name=sync,proto3" json:"sync,omitempty"`
	Serialize            *Serialize `protobuf:"bytes,2,opt,name=serialize,proto3" json:"serialize,omitempty"`
	Target               *Target    `protobuf:"bytes,3,opt,name=target,proto3" json:"target,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *RunOptions) Reset()         { *m = RunOptions{} }
func (m *RunOptions) String() string { return proto.CompactTextString(m) }
func (*RunOptions) ProtoMessage()    {}
func (*RunOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8693f33936b10a7, []int{3}
}

func (m *RunOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunOptions.Unmarshal(m, b)
}
func (m *RunOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunOptions.Marshal(b, m, deterministic)
}
func (m *RunOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunOptions.Merge(m, src)
}
func (m *RunOptions) XXX_Size() int {
	return xxx_messageInfo_RunOptions.Size(m)
}
func (m *RunOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_RunOptions.DiscardUnknown(m)
}

var xxx_messageInfo_RunOptions proto.InternalMessageInfo

func (m *RunOptions) GetSync() bool {
	if m != nil {
		return m.Sync
	}
	return false
}

func (m *RunOptions) GetSerialize() *Serialize {
	if m != nil {
		return m.Serialize
	}
	return nil
}

func (m *RunOptions) GetTarget() *Target {
	if m != nil {
		return m.Target
	}
	return nil
}

type Serialize struct {
	// Count.
	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// FailAfterAllTargets.
	FailAfterAllTargets  bool     `protobuf:"varint,2,opt,name=failAfterAllTargets,proto3" json:"failAfterAllTargets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Serialize) Reset()         { *m = Serialize{} }
func (m *Serialize) String() string { return proto.CompactTextString(m) }
func (*Serialize) ProtoMessage()    {}
func (*Serialize) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8693f33936b10a7, []int{4}
}

func (m *Serialize) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Serialize.Unmarshal(m, b)
}
func (m *Serialize) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Serialize.Marshal(b, m, deterministic)
}
func (m *Serialize) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Serialize.Merge(m, src)
}
func (m *Serialize) XXX_Size() int {
	return xxx_messageInfo_Serialize.Size(m)
}
func (m *Serialize) XXX_DiscardUnknown() {
	xxx_messageInfo_Serialize.DiscardUnknown(m)
}

var xxx_messageInfo_Serialize proto.InternalMessageInfo

func (m *Serialize) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Serialize) GetFailAfterAllTargets() bool {
	if m != nil {
		return m.FailAfterAllTargets
	}
	return false
}

type Target struct {
	// List of hosts (either name or label selector).
	Hosts []string `protobuf:"bytes,1,rep,name=hosts,proto3" json:"hosts,omitempty"`
	// Limit of how many hosts of the matches should be used.
	Limit                int64    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Target) Reset()         { *m = Target{} }
func (m *Target) String() string { return proto.CompactTextString(m) }
func (*Target) ProtoMessage()    {}
func (*Target) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8693f33936b10a7, []int{5}
}

func (m *Target) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Target.Unmarshal(m, b)
}
func (m *Target) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Target.Marshal(b, m, deterministic)
}
func (m *Target) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Target.Merge(m, src)
}
func (m *Target) XXX_Size() int {
	return xxx_messageInfo_Target.Size(m)
}
func (m *Target) XXX_DiscardUnknown() {
	xxx_messageInfo_Target.DiscardUnknown(m)
}

var xxx_messageInfo_Target proto.InternalMessageInfo

func (m *Target) GetHosts() []string {
	if m != nil {
		return m.Hosts
	}
	return nil
}

func (m *Target) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type Action struct {
	// Action name.
	Action string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	// Conditions.
	Conditions *v1alpha.Conditions `protobuf:"bytes,2,opt,name=conditions,proto3" json:"conditions,omitempty"`
	// RunOptions
	RunOptions *RunOptions `protobuf:"bytes,3,opt,name=runOptions,proto3" json:"runOptions,omitempty"`
	// Map of string interface{} to be given to the action.
	Parameters           []*any.Any `protobuf:"bytes,4,rep,name=parameters,proto3" json:"parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Action) Reset()         { *m = Action{} }
func (m *Action) String() string { return proto.CompactTextString(m) }
func (*Action) ProtoMessage()    {}
func (*Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8693f33936b10a7, []int{6}
}

func (m *Action) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Action.Unmarshal(m, b)
}
func (m *Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Action.Marshal(b, m, deterministic)
}
func (m *Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Action.Merge(m, src)
}
func (m *Action) XXX_Size() int {
	return xxx_messageInfo_Action.Size(m)
}
func (m *Action) XXX_DiscardUnknown() {
	xxx_messageInfo_Action.DiscardUnknown(m)
}

var xxx_messageInfo_Action proto.InternalMessageInfo

func (m *Action) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Action) GetConditions() *v1alpha.Conditions {
	if m != nil {
		return m.Conditions
	}
	return nil
}

func (m *Action) GetRunOptions() *RunOptions {
	if m != nil {
		return m.RunOptions
	}
	return nil
}

func (m *Action) GetParameters() []*any.Any {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type AddTaskResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddTaskResponse) Reset()         { *m = AddTaskResponse{} }
func (m *AddTaskResponse) String() string { return proto.CompactTextString(m) }
func (*AddTaskResponse) ProtoMessage()    {}
func (*AddTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8693f33936b10a7, []int{7}
}

func (m *AddTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddTaskResponse.Unmarshal(m, b)
}
func (m *AddTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddTaskResponse.Marshal(b, m, deterministic)
}
func (m *AddTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddTaskResponse.Merge(m, src)
}
func (m *AddTaskResponse) XXX_Size() int {
	return xxx_messageInfo_AddTaskResponse.Size(m)
}
func (m *AddTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddTaskResponse proto.InternalMessageInfo

type UpdateTaskResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTaskResponse) Reset()         { *m = UpdateTaskResponse{} }
func (m *UpdateTaskResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskResponse) ProtoMessage()    {}
func (*UpdateTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8693f33936b10a7, []int{8}
}

func (m *UpdateTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTaskResponse.Unmarshal(m, b)
}
func (m *UpdateTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTaskResponse.Marshal(b, m, deterministic)
}
func (m *UpdateTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTaskResponse.Merge(m, src)
}
func (m *UpdateTaskResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateTaskResponse.Size(m)
}
func (m *UpdateTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTaskResponse proto.InternalMessageInfo

type DeleteTaskResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteTaskResponse) Reset()         { *m = DeleteTaskResponse{} }
func (m *DeleteTaskResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteTaskResponse) ProtoMessage()    {}
func (*DeleteTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8693f33936b10a7, []int{9}
}

func (m *DeleteTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteTaskResponse.Unmarshal(m, b)
}
func (m *DeleteTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteTaskResponse.Marshal(b, m, deterministic)
}
func (m *DeleteTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteTaskResponse.Merge(m, src)
}
func (m *DeleteTaskResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteTaskResponse.Size(m)
}
func (m *DeleteTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteTaskResponse proto.InternalMessageInfo

type WatchTaskResponse struct {
	// Event info for watch response.
	Event *v1alpha1.Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	// Task for watch response.
	Task                 *Task    `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchTaskResponse) Reset()         { *m = WatchTaskResponse{} }
func (m *WatchTaskResponse) String() string { return proto.CompactTextString(m) }
func (*WatchTaskResponse) ProtoMessage()    {}
func (*WatchTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8693f33936b10a7, []int{10}
}

func (m *WatchTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchTaskResponse.Unmarshal(m, b)
}
func (m *WatchTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchTaskResponse.Marshal(b, m, deterministic)
}
func (m *WatchTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchTaskResponse.Merge(m, src)
}
func (m *WatchTaskResponse) XXX_Size() int {
	return xxx_messageInfo_WatchTaskResponse.Size(m)
}
func (m *WatchTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_WatchTaskResponse proto.InternalMessageInfo

func (m *WatchTaskResponse) GetEvent() *v1alpha1.Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *WatchTaskResponse) GetTask() *Task {
	if m != nil {
		return m.Task
	}
	return nil
}

func init() {
	proto.RegisterType((*Task)(nil), "tasks.v1alpha.Task")
	proto.RegisterType((*TaskSpec)(nil), "tasks.v1alpha.TaskSpec")
	proto.RegisterType((*Step)(nil), "tasks.v1alpha.Step")
	proto.RegisterType((*RunOptions)(nil), "tasks.v1alpha.RunOptions")
	proto.RegisterType((*Serialize)(nil), "tasks.v1alpha.Serialize")
	proto.RegisterType((*Target)(nil), "tasks.v1alpha.Target")
	proto.RegisterType((*Action)(nil), "tasks.v1alpha.Action")
	proto.RegisterType((*AddTaskResponse)(nil), "tasks.v1alpha.AddTaskResponse")
	proto.RegisterType((*UpdateTaskResponse)(nil), "tasks.v1alpha.UpdateTaskResponse")
	proto.RegisterType((*DeleteTaskResponse)(nil), "tasks.v1alpha.DeleteTaskResponse")
	proto.RegisterType((*WatchTaskResponse)(nil), "tasks.v1alpha.WatchTaskResponse")
}

func init() {
	proto.RegisterFile("github.com/galexrt/edenconfmgmt/pkg/apis/tasks/v1alpha/api.proto", fileDescriptor_d8693f33936b10a7)
}

var fileDescriptor_d8693f33936b10a7 = []byte{
	// 707 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xae, 0x9b, 0x34, 0x4d, 0xa6, 0xef, 0xab, 0x57, 0xdd, 0xb6, 0x2f, 0x69, 0x04, 0x56, 0xf1,
	0x85, 0x56, 0x55, 0xed, 0xd2, 0x56, 0x55, 0xb9, 0x00, 0x81, 0x16, 0x4e, 0xa8, 0xd2, 0xa6, 0x80,
	0xc4, 0x6d, 0xe3, 0x4c, 0x1c, 0x53, 0x7f, 0xe1, 0xdd, 0x54, 0x94, 0x13, 0x37, 0xae, 0xfc, 0x01,
	0xfe, 0x4f, 0x25, 0x2e, 0x1c, 0x39, 0xd2, 0xf0, 0x47, 0xd0, 0xae, 0x37, 0x5f, 0x8e, 0x91, 0x0a,
	0xb7, 0x9d, 0x99, 0x67, 0x9f, 0x79, 0x76, 0x3e, 0x6c, 0x78, 0xec, 0xf9, 0xa2, 0xd7, 0x6f, 0xdb,
	0x6e, 0x1c, 0x3a, 0x1e, 0x0b, 0xf0, 0x7d, 0x2a, 0x1c, 0xec, 0x60, 0xe4, 0xc6, 0x51, 0x37, 0xf4,
	0x42, 0xe1, 0x24, 0xe7, 0x9e, 0xc3, 0x12, 0x9f, 0x3b, 0x82, 0xf1, 0x73, 0xee, 0x5c, 0xdc, 0x67,
	0x41, 0xd2, 0x63, 0xd2, 0x65, 0x27, 0x69, 0x2c, 0x62, 0xf2, 0xaf, 0x0a, 0xd8, 0x3a, 0xd0, 0xd8,
	0x99, 0x24, 0x8c, 0xbd, 0xd8, 0x51, 0xa8, 0x76, 0xbf, 0xab, 0x2c, 0x65, 0xa8, 0x53, 0x76, 0xbb,
	0xb1, 0xee, 0xc5, 0xb1, 0x17, 0xe0, 0x18, 0xc5, 0xa2, 0x4b, 0x1d, 0x7a, 0x74, 0x63, 0x69, 0x6e,
	0x9c, 0xe2, 0xac, 0xb2, 0x46, 0xf3, 0xc6, 0x04, 0x78, 0x81, 0x91, 0x28, 0x78, 0x9c, 0x15, 0x42,
	0xf9, 0x8c, 0xf1, 0x73, 0x72, 0x04, 0xd5, 0x10, 0x05, 0xeb, 0x30, 0xc1, 0xea, 0xc6, 0x86, 0xb1,
	0xb9, 0xb4, 0x77, 0xdb, 0x96, 0x59, 0x87, 0xcf, 0xb6, 0x4f, 0xdb, 0x6f, 0xd1, 0x15, 0x2f, 0x34,
	0x86, 0x8e, 0xd0, 0x64, 0x1b, 0xca, 0x3c, 0x41, 0xb7, 0x3e, 0xaf, 0x6e, 0xdd, 0xb2, 0xa7, 0xaa,
	0x65, 0x4b, 0xf2, 0x56, 0x82, 0x2e, 0x55, 0x20, 0x2b, 0x81, 0xea, 0xd0, 0x43, 0x1e, 0x00, 0xa4,
	0xfd, 0xe8, 0x34, 0x11, 0x7e, 0x1c, 0x71, 0x9d, 0x74, 0x3d, 0x77, 0x9d, 0x8e, 0x00, 0x74, 0x02,
	0x4c, 0xb6, 0x60, 0x81, 0x0b, 0x4c, 0x78, 0x7d, 0x7e, 0xa3, 0xb4, 0xb9, 0xb4, 0xb7, 0x92, 0xbb,
	0xd5, 0x12, 0x98, 0xd0, 0x0c, 0x61, 0x9d, 0x40, 0x59, 0x9a, 0x84, 0x40, 0x39, 0x62, 0x21, 0xaa,
	0x3c, 0x35, 0xaa, 0xce, 0x64, 0x6b, 0x4a, 0xfa, 0x5a, 0x8e, 0xa5, 0xe9, 0xca, 0x64, 0x5a, 0xf8,
	0x27, 0x03, 0x60, 0x2c, 0x46, 0xb2, 0xf1, 0xcb, 0xc8, 0x55, 0x6c, 0x55, 0xaa, 0xce, 0xe4, 0x10,
	0x6a, 0x1c, 0x53, 0x9f, 0x05, 0xfe, 0x07, 0xd4, 0x94, 0xf5, 0xbc, 0xb0, 0x61, 0x9c, 0x8e, 0xa1,
	0x64, 0x07, 0x2a, 0x82, 0xa5, 0x1e, 0x8a, 0x7a, 0xa9, 0x50, 0xc7, 0x99, 0x0a, 0x52, 0x0d, 0xb2,
	0x5a, 0x50, 0x1b, 0xd1, 0x90, 0x55, 0x58, 0x70, 0xe3, 0x7e, 0x24, 0x94, 0x90, 0x12, 0xcd, 0x0c,
	0xb2, 0x0b, 0x2b, 0x5d, 0xe6, 0x07, 0xcd, 0xae, 0xc0, 0xb4, 0x19, 0x04, 0x19, 0x01, 0x57, 0x9a,
	0xaa, 0xb4, 0x28, 0x64, 0x1d, 0x40, 0x25, 0x3b, 0x4a, 0xc6, 0x5e, 0xcc, 0x85, 0x6c, 0x48, 0x69,
	0xb3, 0x46, 0x33, 0x43, 0x7a, 0x03, 0x3f, 0xf4, 0x85, 0xe2, 0x28, 0xd1, 0xcc, 0xb0, 0xbe, 0x1a,
	0x50, 0xc9, 0xaa, 0x44, 0xfe, 0x87, 0x0a, 0x53, 0x27, 0x5d, 0x60, 0x6d, 0x91, 0x23, 0x00, 0x37,
	0x8e, 0x3a, 0x7e, 0xd6, 0xe4, 0x61, 0x55, 0xa6, 0x26, 0xeb, 0xe9, 0x28, 0x4e, 0x27, 0xb0, 0xb9,
	0xf1, 0x28, 0xfd, 0xc9, 0x78, 0x1c, 0x00, 0x24, 0x2c, 0x65, 0x21, 0x0a, 0x4c, 0x79, 0xbd, 0xac,
	0x66, 0x64, 0xd5, 0xce, 0x16, 0xd1, 0x1e, 0x2e, 0xa2, 0xdd, 0x8c, 0x2e, 0xe9, 0x04, 0xce, 0x5a,
	0x86, 0xff, 0x9a, 0x9d, 0x8e, 0x1c, 0x4f, 0x8a, 0x3c, 0x89, 0x23, 0x8e, 0xd6, 0x2a, 0x90, 0x97,
	0x49, 0x87, 0x09, 0xcc, 0x7b, 0x8f, 0x31, 0xc0, 0x9c, 0xd7, 0x87, 0xe5, 0xd7, 0x4c, 0xb8, 0xbd,
	0x49, 0x27, 0xd9, 0x86, 0x05, 0xb5, 0x7a, 0x7a, 0xbc, 0xd7, 0xec, 0x6c, 0x11, 0x47, 0x0f, 0x38,
	0x91, 0x26, 0xcd, 0x30, 0xe4, 0x1e, 0x94, 0xe5, 0xf3, 0x74, 0x95, 0x56, 0x0a, 0x36, 0x89, 0x2a,
	0xc0, 0xde, 0x97, 0x12, 0xfc, 0x23, 0x4d, 0xde, 0xc2, 0xf4, 0xc2, 0x77, 0x91, 0x3c, 0x83, 0xc5,
	0x57, 0x98, 0x72, 0x59, 0xf0, 0xdc, 0xda, 0x6a, 0x37, 0xc5, 0x77, 0x7d, 0xe4, 0xa2, 0x71, 0xe7,
	0x37, 0x51, 0x2d, 0x77, 0x1f, 0x16, 0x9f, 0xa3, 0x50, 0x1f, 0x84, 0xa2, 0xf4, 0x8d, 0x22, 0x27,
	0x79, 0x08, 0x8b, 0xba, 0x6e, 0xc5, 0x97, 0xcc, 0xfc, 0x5e, 0x4d, 0x17, 0x99, 0x1c, 0x03, 0x8c,
	0x8b, 0x5c, 0x4c, 0x71, 0x37, 0xe7, 0x9c, 0x6d, 0x8a, 0x64, 0x19, 0x37, 0xe5, 0x66, 0x2c, 0xb3,
	0x4d, 0x24, 0xc7, 0x50, 0x1b, 0x35, 0xb1, 0x98, 0x64, 0x23, 0xe7, 0x9c, 0xe9, 0xf9, 0xae, 0xf1,
	0xe4, 0xec, 0xea, 0xda, 0x34, 0xbe, 0x5f, 0x9b, 0x73, 0x1f, 0x07, 0xa6, 0x71, 0x35, 0x30, 0x8d,
	0x6f, 0x03, 0xd3, 0xf8, 0x31, 0x30, 0x8d, 0xcf, 0x3f, 0xcd, 0xb9, 0x37, 0x87, 0x7f, 0xf7, 0x47,
	0x6a, 0x57, 0xd4, 0xe4, 0xee, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xc2, 0xb4, 0xb2, 0xb1, 0xd2,
	0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TasksServiceClient is the client API for TasksService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TasksServiceClient interface {
	// Version returns the API version.
	Version(ctx context.Context, in *v1alpha.VersionRequest, opts ...grpc.CallOption) (*v1alpha.VersionResponse, error)
	// Get a Task.
	GetTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error)
	// Add a Task.
	AddTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*AddTaskResponse, error)
	// Update a Task.
	UpdateTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*UpdateTaskResponse, error)
	// Delete a Task.
	DeleteTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*DeleteTaskResponse, error)
	// Watch a Task.
	WatchTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (TasksService_WatchTaskClient, error)
}

type tasksServiceClient struct {
	cc *grpc.ClientConn
}

func NewTasksServiceClient(cc *grpc.ClientConn) TasksServiceClient {
	return &tasksServiceClient{cc}
}

func (c *tasksServiceClient) Version(ctx context.Context, in *v1alpha.VersionRequest, opts ...grpc.CallOption) (*v1alpha.VersionResponse, error) {
	out := new(v1alpha.VersionResponse)
	err := c.cc.Invoke(ctx, "/tasks.v1alpha.TasksService/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksServiceClient) GetTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/tasks.v1alpha.TasksService/GetTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksServiceClient) AddTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*AddTaskResponse, error) {
	out := new(AddTaskResponse)
	err := c.cc.Invoke(ctx, "/tasks.v1alpha.TasksService/AddTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksServiceClient) UpdateTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*UpdateTaskResponse, error) {
	out := new(UpdateTaskResponse)
	err := c.cc.Invoke(ctx, "/tasks.v1alpha.TasksService/UpdateTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksServiceClient) DeleteTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (*DeleteTaskResponse, error) {
	out := new(DeleteTaskResponse)
	err := c.cc.Invoke(ctx, "/tasks.v1alpha.TasksService/DeleteTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksServiceClient) WatchTask(ctx context.Context, in *Task, opts ...grpc.CallOption) (TasksService_WatchTaskClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TasksService_serviceDesc.Streams[0], "/tasks.v1alpha.TasksService/WatchTask", opts...)
	if err != nil {
		return nil, err
	}
	x := &tasksServiceWatchTaskClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TasksService_WatchTaskClient interface {
	Recv() (*WatchTaskResponse, error)
	grpc.ClientStream
}

type tasksServiceWatchTaskClient struct {
	grpc.ClientStream
}

func (x *tasksServiceWatchTaskClient) Recv() (*WatchTaskResponse, error) {
	m := new(WatchTaskResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TasksServiceServer is the server API for TasksService service.
type TasksServiceServer interface {
	// Version returns the API version.
	Version(context.Context, *v1alpha.VersionRequest) (*v1alpha.VersionResponse, error)
	// Get a Task.
	GetTask(context.Context, *Task) (*Task, error)
	// Add a Task.
	AddTask(context.Context, *Task) (*AddTaskResponse, error)
	// Update a Task.
	UpdateTask(context.Context, *Task) (*UpdateTaskResponse, error)
	// Delete a Task.
	DeleteTask(context.Context, *Task) (*DeleteTaskResponse, error)
	// Watch a Task.
	WatchTask(*Task, TasksService_WatchTaskServer) error
}

func RegisterTasksServiceServer(s *grpc.Server, srv TasksServiceServer) {
	s.RegisterService(&_TasksService_serviceDesc, srv)
}

func _TasksService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1alpha.VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks.v1alpha.TasksService/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServiceServer).Version(ctx, req.(*v1alpha.VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TasksService_GetTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServiceServer).GetTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks.v1alpha.TasksService/GetTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServiceServer).GetTask(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _TasksService_AddTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServiceServer).AddTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks.v1alpha.TasksService/AddTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServiceServer).AddTask(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _TasksService_UpdateTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServiceServer).UpdateTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks.v1alpha.TasksService/UpdateTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServiceServer).UpdateTask(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _TasksService_DeleteTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Task)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServiceServer).DeleteTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tasks.v1alpha.TasksService/DeleteTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServiceServer).DeleteTask(ctx, req.(*Task))
	}
	return interceptor(ctx, in, info, handler)
}

func _TasksService_WatchTask_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Task)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TasksServiceServer).WatchTask(m, &tasksServiceWatchTaskServer{stream})
}

type TasksService_WatchTaskServer interface {
	Send(*WatchTaskResponse) error
	grpc.ServerStream
}

type tasksServiceWatchTaskServer struct {
	grpc.ServerStream
}

func (x *tasksServiceWatchTaskServer) Send(m *WatchTaskResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _TasksService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tasks.v1alpha.TasksService",
	HandlerType: (*TasksServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _TasksService_Version_Handler,
		},
		{
			MethodName: "GetTask",
			Handler:    _TasksService_GetTask_Handler,
		},
		{
			MethodName: "AddTask",
			Handler:    _TasksService_AddTask_Handler,
		},
		{
			MethodName: "UpdateTask",
			Handler:    _TasksService_UpdateTask_Handler,
		},
		{
			MethodName: "DeleteTask",
			Handler:    _TasksService_DeleteTask_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchTask",
			Handler:       _TasksService_WatchTask_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/galexrt/edenconfmgmt/pkg/apis/tasks/v1alpha/api.proto",
}
