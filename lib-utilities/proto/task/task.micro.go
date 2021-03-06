// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: task.proto

package task

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for GetTaskService service

type GetTaskService interface {
	DeleteTask(ctx context.Context, in *GetTaskRequest, opts ...client.CallOption) (*TaskResponse, error)
	GetTasks(ctx context.Context, in *GetTaskRequest, opts ...client.CallOption) (*TaskResponse, error)
	GetSubTasks(ctx context.Context, in *GetTaskRequest, opts ...client.CallOption) (*TaskResponse, error)
	GetSubTask(ctx context.Context, in *GetTaskRequest, opts ...client.CallOption) (*TaskResponse, error)
	TaskCollection(ctx context.Context, in *GetTaskRequest, opts ...client.CallOption) (*TaskResponse, error)
	GetTaskService(ctx context.Context, in *GetTaskRequest, opts ...client.CallOption) (*TaskResponse, error)
	GetTaskMonitor(ctx context.Context, in *GetTaskRequest, opts ...client.CallOption) (*TaskResponse, error)
	CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...client.CallOption) (*CreateTaskResponse, error)
	CreateChildTask(ctx context.Context, in *CreateTaskRequest, opts ...client.CallOption) (*CreateTaskResponse, error)
	UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...client.CallOption) (*UpdateTaskResponse, error)
}

type getTaskService struct {
	c    client.Client
	name string
}

func NewGetTaskService(name string, c client.Client) GetTaskService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "gettaskservice"
	}
	return &getTaskService{
		c:    c,
		name: name,
	}
}

func (c *getTaskService) DeleteTask(ctx context.Context, in *GetTaskRequest, opts ...client.CallOption) (*TaskResponse, error) {
	req := c.c.NewRequest(c.name, "GetTaskService.DeleteTask", in)
	out := new(TaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getTaskService) GetTasks(ctx context.Context, in *GetTaskRequest, opts ...client.CallOption) (*TaskResponse, error) {
	req := c.c.NewRequest(c.name, "GetTaskService.GetTasks", in)
	out := new(TaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getTaskService) GetSubTasks(ctx context.Context, in *GetTaskRequest, opts ...client.CallOption) (*TaskResponse, error) {
	req := c.c.NewRequest(c.name, "GetTaskService.GetSubTasks", in)
	out := new(TaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getTaskService) GetSubTask(ctx context.Context, in *GetTaskRequest, opts ...client.CallOption) (*TaskResponse, error) {
	req := c.c.NewRequest(c.name, "GetTaskService.GetSubTask", in)
	out := new(TaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getTaskService) TaskCollection(ctx context.Context, in *GetTaskRequest, opts ...client.CallOption) (*TaskResponse, error) {
	req := c.c.NewRequest(c.name, "GetTaskService.TaskCollection", in)
	out := new(TaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getTaskService) GetTaskService(ctx context.Context, in *GetTaskRequest, opts ...client.CallOption) (*TaskResponse, error) {
	req := c.c.NewRequest(c.name, "GetTaskService.GetTaskService", in)
	out := new(TaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getTaskService) GetTaskMonitor(ctx context.Context, in *GetTaskRequest, opts ...client.CallOption) (*TaskResponse, error) {
	req := c.c.NewRequest(c.name, "GetTaskService.GetTaskMonitor", in)
	out := new(TaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getTaskService) CreateTask(ctx context.Context, in *CreateTaskRequest, opts ...client.CallOption) (*CreateTaskResponse, error) {
	req := c.c.NewRequest(c.name, "GetTaskService.CreateTask", in)
	out := new(CreateTaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getTaskService) CreateChildTask(ctx context.Context, in *CreateTaskRequest, opts ...client.CallOption) (*CreateTaskResponse, error) {
	req := c.c.NewRequest(c.name, "GetTaskService.CreateChildTask", in)
	out := new(CreateTaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *getTaskService) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...client.CallOption) (*UpdateTaskResponse, error) {
	req := c.c.NewRequest(c.name, "GetTaskService.UpdateTask", in)
	out := new(UpdateTaskResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for GetTaskService service

type GetTaskServiceHandler interface {
	DeleteTask(context.Context, *GetTaskRequest, *TaskResponse) error
	GetTasks(context.Context, *GetTaskRequest, *TaskResponse) error
	GetSubTasks(context.Context, *GetTaskRequest, *TaskResponse) error
	GetSubTask(context.Context, *GetTaskRequest, *TaskResponse) error
	TaskCollection(context.Context, *GetTaskRequest, *TaskResponse) error
	GetTaskService(context.Context, *GetTaskRequest, *TaskResponse) error
	GetTaskMonitor(context.Context, *GetTaskRequest, *TaskResponse) error
	CreateTask(context.Context, *CreateTaskRequest, *CreateTaskResponse) error
	CreateChildTask(context.Context, *CreateTaskRequest, *CreateTaskResponse) error
	UpdateTask(context.Context, *UpdateTaskRequest, *UpdateTaskResponse) error
}

func RegisterGetTaskServiceHandler(s server.Server, hdlr GetTaskServiceHandler, opts ...server.HandlerOption) error {
	type getTaskService interface {
		DeleteTask(ctx context.Context, in *GetTaskRequest, out *TaskResponse) error
		GetTasks(ctx context.Context, in *GetTaskRequest, out *TaskResponse) error
		GetSubTasks(ctx context.Context, in *GetTaskRequest, out *TaskResponse) error
		GetSubTask(ctx context.Context, in *GetTaskRequest, out *TaskResponse) error
		TaskCollection(ctx context.Context, in *GetTaskRequest, out *TaskResponse) error
		GetTaskService(ctx context.Context, in *GetTaskRequest, out *TaskResponse) error
		GetTaskMonitor(ctx context.Context, in *GetTaskRequest, out *TaskResponse) error
		CreateTask(ctx context.Context, in *CreateTaskRequest, out *CreateTaskResponse) error
		CreateChildTask(ctx context.Context, in *CreateTaskRequest, out *CreateTaskResponse) error
		UpdateTask(ctx context.Context, in *UpdateTaskRequest, out *UpdateTaskResponse) error
	}
	type GetTaskService struct {
		getTaskService
	}
	h := &getTaskServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&GetTaskService{h}, opts...))
}

type getTaskServiceHandler struct {
	GetTaskServiceHandler
}

func (h *getTaskServiceHandler) DeleteTask(ctx context.Context, in *GetTaskRequest, out *TaskResponse) error {
	return h.GetTaskServiceHandler.DeleteTask(ctx, in, out)
}

func (h *getTaskServiceHandler) GetTasks(ctx context.Context, in *GetTaskRequest, out *TaskResponse) error {
	return h.GetTaskServiceHandler.GetTasks(ctx, in, out)
}

func (h *getTaskServiceHandler) GetSubTasks(ctx context.Context, in *GetTaskRequest, out *TaskResponse) error {
	return h.GetTaskServiceHandler.GetSubTasks(ctx, in, out)
}

func (h *getTaskServiceHandler) GetSubTask(ctx context.Context, in *GetTaskRequest, out *TaskResponse) error {
	return h.GetTaskServiceHandler.GetSubTask(ctx, in, out)
}

func (h *getTaskServiceHandler) TaskCollection(ctx context.Context, in *GetTaskRequest, out *TaskResponse) error {
	return h.GetTaskServiceHandler.TaskCollection(ctx, in, out)
}

func (h *getTaskServiceHandler) GetTaskService(ctx context.Context, in *GetTaskRequest, out *TaskResponse) error {
	return h.GetTaskServiceHandler.GetTaskService(ctx, in, out)
}

func (h *getTaskServiceHandler) GetTaskMonitor(ctx context.Context, in *GetTaskRequest, out *TaskResponse) error {
	return h.GetTaskServiceHandler.GetTaskMonitor(ctx, in, out)
}

func (h *getTaskServiceHandler) CreateTask(ctx context.Context, in *CreateTaskRequest, out *CreateTaskResponse) error {
	return h.GetTaskServiceHandler.CreateTask(ctx, in, out)
}

func (h *getTaskServiceHandler) CreateChildTask(ctx context.Context, in *CreateTaskRequest, out *CreateTaskResponse) error {
	return h.GetTaskServiceHandler.CreateChildTask(ctx, in, out)
}

func (h *getTaskServiceHandler) UpdateTask(ctx context.Context, in *UpdateTaskRequest, out *UpdateTaskResponse) error {
	return h.GetTaskServiceHandler.UpdateTask(ctx, in, out)
}
