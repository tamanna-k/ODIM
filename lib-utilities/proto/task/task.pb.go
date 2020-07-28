// Code generated by protoc-gen-go. DO NOT EDIT.
// source: lib-utilities/proto/task/task.proto

package task

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Payload struct {
	HTTPHeaders          map[string]string `protobuf:"bytes,1,rep,name=HTTPHeaders,proto3" json:"HTTPHeaders,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	HTTPOperation        string            `protobuf:"bytes,2,opt,name=HTTPOperation,proto3" json:"HTTPOperation,omitempty"`
	JSONBody             []byte            `protobuf:"bytes,3,opt,name=JSONBody,proto3" json:"JSONBody,omitempty"`
	StatusCode           int32             `protobuf:"varint,4,opt,name=StatusCode,proto3" json:"StatusCode,omitempty"`
	TargetURI            string            `protobuf:"bytes,5,opt,name=TargetURI,proto3" json:"TargetURI,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Payload) Reset()         { *m = Payload{} }
func (m *Payload) String() string { return proto.CompactTextString(m) }
func (*Payload) ProtoMessage()    {}
func (*Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_9778bdeffbdbac56, []int{0}
}

func (m *Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Payload.Unmarshal(m, b)
}
func (m *Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Payload.Marshal(b, m, deterministic)
}
func (m *Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payload.Merge(m, src)
}
func (m *Payload) XXX_Size() int {
	return xxx_messageInfo_Payload.Size(m)
}
func (m *Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Payload proto.InternalMessageInfo

func (m *Payload) GetHTTPHeaders() map[string]string {
	if m != nil {
		return m.HTTPHeaders
	}
	return nil
}

func (m *Payload) GetHTTPOperation() string {
	if m != nil {
		return m.HTTPOperation
	}
	return ""
}

func (m *Payload) GetJSONBody() []byte {
	if m != nil {
		return m.JSONBody
	}
	return nil
}

func (m *Payload) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *Payload) GetTargetURI() string {
	if m != nil {
		return m.TargetURI
	}
	return ""
}

type GetTaskRequest struct {
	TaskID               string   `protobuf:"bytes,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
	SubTaskID            string   `protobuf:"bytes,2,opt,name=subTaskID,proto3" json:"subTaskID,omitempty"`
	SessionToken         string   `protobuf:"bytes,3,opt,name=sessionToken,proto3" json:"sessionToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTaskRequest) Reset()         { *m = GetTaskRequest{} }
func (m *GetTaskRequest) String() string { return proto.CompactTextString(m) }
func (*GetTaskRequest) ProtoMessage()    {}
func (*GetTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9778bdeffbdbac56, []int{1}
}

func (m *GetTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTaskRequest.Unmarshal(m, b)
}
func (m *GetTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTaskRequest.Marshal(b, m, deterministic)
}
func (m *GetTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTaskRequest.Merge(m, src)
}
func (m *GetTaskRequest) XXX_Size() int {
	return xxx_messageInfo_GetTaskRequest.Size(m)
}
func (m *GetTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTaskRequest proto.InternalMessageInfo

func (m *GetTaskRequest) GetTaskID() string {
	if m != nil {
		return m.TaskID
	}
	return ""
}

func (m *GetTaskRequest) GetSubTaskID() string {
	if m != nil {
		return m.SubTaskID
	}
	return ""
}

func (m *GetTaskRequest) GetSessionToken() string {
	if m != nil {
		return m.SessionToken
	}
	return ""
}

type TaskResponse struct {
	StatusCode           int32             `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	StatusMessage        string            `protobuf:"bytes,2,opt,name=statusMessage,proto3" json:"statusMessage,omitempty"`
	Header               map[string]string `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body                 []byte            `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TaskResponse) Reset()         { *m = TaskResponse{} }
func (m *TaskResponse) String() string { return proto.CompactTextString(m) }
func (*TaskResponse) ProtoMessage()    {}
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9778bdeffbdbac56, []int{2}
}

func (m *TaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskResponse.Unmarshal(m, b)
}
func (m *TaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskResponse.Marshal(b, m, deterministic)
}
func (m *TaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskResponse.Merge(m, src)
}
func (m *TaskResponse) XXX_Size() int {
	return xxx_messageInfo_TaskResponse.Size(m)
}
func (m *TaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TaskResponse proto.InternalMessageInfo

func (m *TaskResponse) GetStatusCode() int32 {
	if m != nil {
		return m.StatusCode
	}
	return 0
}

func (m *TaskResponse) GetStatusMessage() string {
	if m != nil {
		return m.StatusMessage
	}
	return ""
}

func (m *TaskResponse) GetHeader() map[string]string {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *TaskResponse) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type CreateTaskRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	ParentTaskID         string   `protobuf:"bytes,2,opt,name=parentTaskID,proto3" json:"parentTaskID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTaskRequest) Reset()         { *m = CreateTaskRequest{} }
func (m *CreateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTaskRequest) ProtoMessage()    {}
func (*CreateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9778bdeffbdbac56, []int{3}
}

func (m *CreateTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskRequest.Unmarshal(m, b)
}
func (m *CreateTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskRequest.Marshal(b, m, deterministic)
}
func (m *CreateTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskRequest.Merge(m, src)
}
func (m *CreateTaskRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTaskRequest.Size(m)
}
func (m *CreateTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskRequest proto.InternalMessageInfo

func (m *CreateTaskRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *CreateTaskRequest) GetParentTaskID() string {
	if m != nil {
		return m.ParentTaskID
	}
	return ""
}

type CreateTaskResponse struct {
	TaskURI              string   `protobuf:"bytes,1,opt,name=taskURI,proto3" json:"taskURI,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTaskResponse) Reset()         { *m = CreateTaskResponse{} }
func (m *CreateTaskResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTaskResponse) ProtoMessage()    {}
func (*CreateTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9778bdeffbdbac56, []int{4}
}

func (m *CreateTaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskResponse.Unmarshal(m, b)
}
func (m *CreateTaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskResponse.Marshal(b, m, deterministic)
}
func (m *CreateTaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskResponse.Merge(m, src)
}
func (m *CreateTaskResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTaskResponse.Size(m)
}
func (m *CreateTaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskResponse proto.InternalMessageInfo

func (m *CreateTaskResponse) GetTaskURI() string {
	if m != nil {
		return m.TaskURI
	}
	return ""
}

type UpdateTaskRequest struct {
	TaskID               string               `protobuf:"bytes,1,opt,name=taskID,proto3" json:"taskID,omitempty"`
	TaskState            string               `protobuf:"bytes,2,opt,name=taskState,proto3" json:"taskState,omitempty"`
	TaskStatus           string               `protobuf:"bytes,3,opt,name=taskStatus,proto3" json:"taskStatus,omitempty"`
	PercentComplete      int32                `protobuf:"varint,4,opt,name=percentComplete,proto3" json:"percentComplete,omitempty"`
	PayLoad              *Payload             `protobuf:"bytes,5,opt,name=payLoad,proto3" json:"payLoad,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,6,opt,name=endTime,proto3" json:"endTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UpdateTaskRequest) Reset()         { *m = UpdateTaskRequest{} }
func (m *UpdateTaskRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskRequest) ProtoMessage()    {}
func (*UpdateTaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9778bdeffbdbac56, []int{5}
}

func (m *UpdateTaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTaskRequest.Unmarshal(m, b)
}
func (m *UpdateTaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTaskRequest.Marshal(b, m, deterministic)
}
func (m *UpdateTaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTaskRequest.Merge(m, src)
}
func (m *UpdateTaskRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateTaskRequest.Size(m)
}
func (m *UpdateTaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTaskRequest proto.InternalMessageInfo

func (m *UpdateTaskRequest) GetTaskID() string {
	if m != nil {
		return m.TaskID
	}
	return ""
}

func (m *UpdateTaskRequest) GetTaskState() string {
	if m != nil {
		return m.TaskState
	}
	return ""
}

func (m *UpdateTaskRequest) GetTaskStatus() string {
	if m != nil {
		return m.TaskStatus
	}
	return ""
}

func (m *UpdateTaskRequest) GetPercentComplete() int32 {
	if m != nil {
		return m.PercentComplete
	}
	return 0
}

func (m *UpdateTaskRequest) GetPayLoad() *Payload {
	if m != nil {
		return m.PayLoad
	}
	return nil
}

func (m *UpdateTaskRequest) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

type UpdateTaskResponse struct {
	StatusMessage        string   `protobuf:"bytes,1,opt,name=statusMessage,proto3" json:"statusMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateTaskResponse) Reset()         { *m = UpdateTaskResponse{} }
func (m *UpdateTaskResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskResponse) ProtoMessage()    {}
func (*UpdateTaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9778bdeffbdbac56, []int{6}
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

func (m *UpdateTaskResponse) GetStatusMessage() string {
	if m != nil {
		return m.StatusMessage
	}
	return ""
}

func init() {
	proto.RegisterType((*Payload)(nil), "Payload")
	proto.RegisterMapType((map[string]string)(nil), "Payload.HTTPHeadersEntry")
	proto.RegisterType((*GetTaskRequest)(nil), "GetTaskRequest")
	proto.RegisterType((*TaskResponse)(nil), "TaskResponse")
	proto.RegisterMapType((map[string]string)(nil), "TaskResponse.HeaderEntry")
	proto.RegisterType((*CreateTaskRequest)(nil), "CreateTaskRequest")
	proto.RegisterType((*CreateTaskResponse)(nil), "CreateTaskResponse")
	proto.RegisterType((*UpdateTaskRequest)(nil), "UpdateTaskRequest")
	proto.RegisterType((*UpdateTaskResponse)(nil), "UpdateTaskResponse")
}

func init() {
	proto.RegisterFile("lib-utilities/proto/task/task.proto", fileDescriptor_9778bdeffbdbac56)
}

var fileDescriptor_9778bdeffbdbac56 = []byte{
	// 656 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x5d, 0x6e, 0xd3, 0x40,
	0x10, 0xae, 0x9b, 0x34, 0x49, 0x27, 0xfd, 0xdd, 0x22, 0x64, 0x2c, 0x04, 0x91, 0xe1, 0x21, 0x0f,
	0xb0, 0x11, 0x01, 0xa9, 0x50, 0x10, 0x0f, 0xa4, 0xa8, 0x2d, 0xa2, 0x3f, 0x72, 0xdc, 0x03, 0x6c,
	0xea, 0x21, 0x35, 0x71, 0xbc, 0xc6, 0xbb, 0xae, 0x94, 0x0b, 0x70, 0x09, 0xee, 0xc5, 0x2d, 0xb8,
	0x03, 0x5a, 0xaf, 0x5d, 0xdb, 0x09, 0x42, 0xc9, 0x4b, 0xb4, 0xf3, 0xcd, 0x8f, 0x67, 0xbe, 0xf9,
	0x26, 0xf0, 0x2c, 0xf0, 0x47, 0x2f, 0x13, 0xe9, 0x07, 0xbe, 0xf4, 0x51, 0xf4, 0xa2, 0x98, 0x4b,
	0xde, 0x93, 0x4c, 0x4c, 0xd2, 0x1f, 0x9a, 0xda, 0xd6, 0xd3, 0x31, 0xe7, 0xe3, 0x00, 0xb5, 0x77,
	0x94, 0x7c, 0xeb, 0x49, 0x7f, 0x8a, 0x42, 0xb2, 0x69, 0xa4, 0x03, 0xec, 0x9f, 0xeb, 0xd0, 0xbc,
	0x62, 0xb3, 0x80, 0x33, 0x8f, 0xbc, 0x87, 0xf6, 0xa9, 0xeb, 0x5e, 0x9d, 0x22, 0xf3, 0x30, 0x16,
	0xa6, 0xd1, 0xa9, 0x75, 0xdb, 0xfd, 0x47, 0x34, 0x73, 0xd3, 0x92, 0xef, 0x73, 0x28, 0xe3, 0x99,
	0x53, 0x8e, 0x26, 0xcf, 0x61, 0x5b, 0x99, 0x97, 0x11, 0xc6, 0x4c, 0xfa, 0x3c, 0x34, 0xd7, 0x3b,
	0x46, 0x77, 0xd3, 0xa9, 0x82, 0xc4, 0x82, 0xd6, 0x97, 0xe1, 0xe5, 0xc5, 0x27, 0xee, 0xcd, 0xcc,
	0x5a, 0xc7, 0xe8, 0x6e, 0x39, 0xf7, 0x36, 0x79, 0x02, 0x30, 0x94, 0x4c, 0x26, 0x62, 0xc0, 0x3d,
	0x34, 0xeb, 0x1d, 0xa3, 0xbb, 0xe1, 0x94, 0x10, 0xf2, 0x18, 0x36, 0x5d, 0x16, 0x8f, 0x51, 0x5e,
	0x3b, 0x67, 0xe6, 0x46, 0x5a, 0xbd, 0x00, 0xac, 0x8f, 0xb0, 0x37, 0xdf, 0x20, 0xd9, 0x83, 0xda,
	0x04, 0x67, 0xa6, 0x91, 0xc6, 0xaa, 0x27, 0x79, 0x00, 0x1b, 0x77, 0x2c, 0x48, 0x30, 0xeb, 0x4e,
	0x1b, 0x47, 0xeb, 0x6f, 0x0d, 0xfb, 0x3b, 0xec, 0x9c, 0xa0, 0x74, 0x99, 0x98, 0x38, 0xf8, 0x23,
	0x41, 0x21, 0xc9, 0x43, 0x68, 0x28, 0x26, 0xcf, 0x8e, 0xb3, 0x02, 0x99, 0xa5, 0xfa, 0x10, 0xc9,
	0xc8, 0xd5, 0x2e, 0x5d, 0xa7, 0x00, 0x88, 0x0d, 0x5b, 0x02, 0x85, 0xf0, 0x79, 0xe8, 0xf2, 0x09,
	0x86, 0xe9, 0x94, 0x9b, 0x4e, 0x05, 0xb3, 0x7f, 0x1b, 0xb0, 0xa5, 0xbf, 0x24, 0x22, 0x1e, 0x0a,
	0x54, 0xa3, 0x8b, 0x62, 0x74, 0x43, 0x8f, 0x5e, 0x20, 0x8a, 0x5c, 0x6d, 0x9d, 0xa3, 0x10, 0x6c,
	0x9c, 0xb7, 0x5f, 0x05, 0xc9, 0x2b, 0x68, 0xdc, 0xa6, 0xe3, 0x9b, 0xb5, 0x6c, 0x75, 0xe5, 0x8f,
	0x50, 0x4d, 0x8d, 0x5e, 0x5d, 0x16, 0x48, 0x08, 0xd4, 0x47, 0x6a, 0x17, 0xf5, 0x74, 0x17, 0xe9,
	0xdb, 0x7a, 0x07, 0xed, 0x52, 0xe8, 0x4a, 0x24, 0x0e, 0x61, 0x7f, 0x10, 0x23, 0x93, 0x58, 0xe6,
	0xd1, 0x82, 0x56, 0x22, 0x30, 0xbe, 0x60, 0x53, 0xcc, 0xaa, 0xdc, 0xdb, 0x8a, 0xad, 0x88, 0xc5,
	0x18, 0xca, 0x0a, 0x9d, 0x15, 0xcc, 0xa6, 0x40, 0xca, 0x45, 0x33, 0xca, 0x4c, 0x68, 0xaa, 0x7d,
	0x28, 0x2d, 0xe8, 0xa2, 0xb9, 0x69, 0xff, 0x31, 0x60, 0xff, 0x3a, 0xf2, 0xe6, 0xba, 0xf8, 0xcf,
	0x36, 0xd5, 0x4b, 0xe9, 0x2c, 0x1f, 0xa8, 0x00, 0xd4, 0x62, 0x72, 0x23, 0x11, 0xd9, 0x2e, 0x4b,
	0x08, 0xe9, 0xc2, 0x6e, 0x84, 0xf1, 0x0d, 0x86, 0x72, 0xc0, 0xa7, 0x51, 0x80, 0x32, 0x17, 0xee,
	0x3c, 0x4c, 0x6c, 0x68, 0x46, 0x6c, 0xf6, 0x95, 0x33, 0x2f, 0xd5, 0x6e, 0xbb, 0xdf, 0xca, 0x0f,
	0xcb, 0xc9, 0x1d, 0xe4, 0x0d, 0x34, 0x31, 0xf4, 0x5c, 0x7f, 0x8a, 0x66, 0x23, 0x8d, 0xb1, 0xa8,
	0xbe, 0x5f, 0x9a, 0xdf, 0x2f, 0x75, 0xf3, 0xfb, 0x75, 0xf2, 0x50, 0xfb, 0x08, 0x48, 0x79, 0xdc,
	0x8c, 0x9f, 0x05, 0xc9, 0x18, 0xff, 0x90, 0x4c, 0xff, 0x57, 0xfd, 0x5e, 0xf6, 0x43, 0x8c, 0xef,
	0xfc, 0x1b, 0x24, 0x14, 0xe0, 0x18, 0x55, 0xcb, 0x0a, 0x24, 0xbb, 0xb4, 0x7a, 0x15, 0xd6, 0x76,
	0x45, 0x54, 0xf6, 0x1a, 0x79, 0x01, 0xad, 0x2c, 0x44, 0x2c, 0x11, 0xdd, 0x83, 0xf6, 0x09, 0xca,
	0xa1, 0x3e, 0x97, 0x65, 0x12, 0x28, 0x40, 0x91, 0xb0, 0x44, 0x7c, 0x1f, 0x76, 0x14, 0x32, 0xe0,
	0x41, 0x80, 0x37, 0xe9, 0x7f, 0xce, 0x52, 0x39, 0x73, 0x24, 0xac, 0x92, 0x73, 0xce, 0x43, 0x5f,
	0xf2, 0x78, 0x89, 0x9c, 0x43, 0x80, 0x42, 0xc9, 0x84, 0xd0, 0x85, 0x5b, 0xb1, 0x0e, 0xe8, 0xa2,
	0xd4, 0xed, 0x35, 0xf2, 0x01, 0x76, 0x35, 0x3e, 0xb8, 0xf5, 0x03, 0x6f, 0xd5, 0xec, 0x43, 0x80,
	0x42, 0x20, 0x84, 0xd0, 0x85, 0xe3, 0xb0, 0x0e, 0xe8, 0xa2, 0x82, 0xec, 0xb5, 0x51, 0x23, 0x95,
	0xdd, 0xeb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x79, 0x8c, 0x30, 0x2a, 0x6b, 0x06, 0x00, 0x00,
}