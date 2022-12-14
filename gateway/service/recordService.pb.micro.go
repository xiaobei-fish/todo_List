// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: recordService.proto

package service

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for RecordService service

func NewRecordServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RecordService service

type RecordService interface {
	FormRecord(ctx context.Context, in *RecordRequest, opts ...client.CallOption) (*RecordInfoResponse, error)
	GetRecordsList(ctx context.Context, in *RecordRequest, opts ...client.CallOption) (*RecordListResponse, error)
	GetRecord(ctx context.Context, in *RecordRequest, opts ...client.CallOption) (*RecordInfoResponse, error)
	UpdateRecord(ctx context.Context, in *RecordRequest, opts ...client.CallOption) (*RecordInfoResponse, error)
	DeleteRecord(ctx context.Context, in *RecordRequest, opts ...client.CallOption) (*RecordInfoResponse, error)
	OpHistory(ctx context.Context, in *HistoryRequest, opts ...client.CallOption) (*HistoryInfo, error)
}

type recordService struct {
	c    client.Client
	name string
}

func NewRecordService(name string, c client.Client) RecordService {
	return &recordService{
		c:    c,
		name: name,
	}
}

func (c *recordService) FormRecord(ctx context.Context, in *RecordRequest, opts ...client.CallOption) (*RecordInfoResponse, error) {
	req := c.c.NewRequest(c.name, "RecordService.FormRecord", in)
	out := new(RecordInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordService) GetRecordsList(ctx context.Context, in *RecordRequest, opts ...client.CallOption) (*RecordListResponse, error) {
	req := c.c.NewRequest(c.name, "RecordService.GetRecordsList", in)
	out := new(RecordListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordService) GetRecord(ctx context.Context, in *RecordRequest, opts ...client.CallOption) (*RecordInfoResponse, error) {
	req := c.c.NewRequest(c.name, "RecordService.GetRecord", in)
	out := new(RecordInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordService) UpdateRecord(ctx context.Context, in *RecordRequest, opts ...client.CallOption) (*RecordInfoResponse, error) {
	req := c.c.NewRequest(c.name, "RecordService.UpdateRecord", in)
	out := new(RecordInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordService) DeleteRecord(ctx context.Context, in *RecordRequest, opts ...client.CallOption) (*RecordInfoResponse, error) {
	req := c.c.NewRequest(c.name, "RecordService.DeleteRecord", in)
	out := new(RecordInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *recordService) OpHistory(ctx context.Context, in *HistoryRequest, opts ...client.CallOption) (*HistoryInfo, error) {
	req := c.c.NewRequest(c.name, "RecordService.OpHistory", in)
	out := new(HistoryInfo)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RecordService service

type RecordServiceHandler interface {
	FormRecord(context.Context, *RecordRequest, *RecordInfoResponse) error
	GetRecordsList(context.Context, *RecordRequest, *RecordListResponse) error
	GetRecord(context.Context, *RecordRequest, *RecordInfoResponse) error
	UpdateRecord(context.Context, *RecordRequest, *RecordInfoResponse) error
	DeleteRecord(context.Context, *RecordRequest, *RecordInfoResponse) error
	OpHistory(context.Context, *HistoryRequest, *HistoryInfo) error
}

func RegisterRecordServiceHandler(s server.Server, hdlr RecordServiceHandler, opts ...server.HandlerOption) error {
	type recordService interface {
		FormRecord(ctx context.Context, in *RecordRequest, out *RecordInfoResponse) error
		GetRecordsList(ctx context.Context, in *RecordRequest, out *RecordListResponse) error
		GetRecord(ctx context.Context, in *RecordRequest, out *RecordInfoResponse) error
		UpdateRecord(ctx context.Context, in *RecordRequest, out *RecordInfoResponse) error
		DeleteRecord(ctx context.Context, in *RecordRequest, out *RecordInfoResponse) error
		OpHistory(ctx context.Context, in *HistoryRequest, out *HistoryInfo) error
	}
	type RecordService struct {
		recordService
	}
	h := &recordServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RecordService{h}, opts...))
}

type recordServiceHandler struct {
	RecordServiceHandler
}

func (h *recordServiceHandler) FormRecord(ctx context.Context, in *RecordRequest, out *RecordInfoResponse) error {
	return h.RecordServiceHandler.FormRecord(ctx, in, out)
}

func (h *recordServiceHandler) GetRecordsList(ctx context.Context, in *RecordRequest, out *RecordListResponse) error {
	return h.RecordServiceHandler.GetRecordsList(ctx, in, out)
}

func (h *recordServiceHandler) GetRecord(ctx context.Context, in *RecordRequest, out *RecordInfoResponse) error {
	return h.RecordServiceHandler.GetRecord(ctx, in, out)
}

func (h *recordServiceHandler) UpdateRecord(ctx context.Context, in *RecordRequest, out *RecordInfoResponse) error {
	return h.RecordServiceHandler.UpdateRecord(ctx, in, out)
}

func (h *recordServiceHandler) DeleteRecord(ctx context.Context, in *RecordRequest, out *RecordInfoResponse) error {
	return h.RecordServiceHandler.DeleteRecord(ctx, in, out)
}

func (h *recordServiceHandler) OpHistory(ctx context.Context, in *HistoryRequest, out *HistoryInfo) error {
	return h.RecordServiceHandler.OpHistory(ctx, in, out)
}
