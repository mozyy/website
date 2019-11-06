// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datamanage/database.proto

package database

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	message "go.yyue.dev/common/message"
	lianjia "go.yyue.dev/crawler/parser/lianjia"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type ConnectRequest struct {
	Database             string   `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectRequest) Reset()         { *m = ConnectRequest{} }
func (m *ConnectRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectRequest) ProtoMessage()    {}
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4ad0b9d61c6647e, []int{0}
}

func (m *ConnectRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectRequest.Unmarshal(m, b)
}
func (m *ConnectRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectRequest.Marshal(b, m, deterministic)
}
func (m *ConnectRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectRequest.Merge(m, src)
}
func (m *ConnectRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectRequest.Size(m)
}
func (m *ConnectRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectRequest proto.InternalMessageInfo

func (m *ConnectRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

type InsertRequest struct {
	HouseInfo            *lianjia.HouseInfo `protobuf:"bytes,1,opt,name=house_info,json=houseInfo,proto3" json:"house_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *InsertRequest) Reset()         { *m = InsertRequest{} }
func (m *InsertRequest) String() string { return proto.CompactTextString(m) }
func (*InsertRequest) ProtoMessage()    {}
func (*InsertRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d4ad0b9d61c6647e, []int{1}
}

func (m *InsertRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertRequest.Unmarshal(m, b)
}
func (m *InsertRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertRequest.Marshal(b, m, deterministic)
}
func (m *InsertRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertRequest.Merge(m, src)
}
func (m *InsertRequest) XXX_Size() int {
	return xxx_messageInfo_InsertRequest.Size(m)
}
func (m *InsertRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InsertRequest proto.InternalMessageInfo

func (m *InsertRequest) GetHouseInfo() *lianjia.HouseInfo {
	if m != nil {
		return m.HouseInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*ConnectRequest)(nil), "database.ConnectRequest")
	proto.RegisterType((*InsertRequest)(nil), "database.InsertRequest")
}

func init() { proto.RegisterFile("datamanage/database.proto", fileDescriptor_d4ad0b9d61c6647e) }

var fileDescriptor_d4ad0b9d61c6647e = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x49, 0x2c, 0x49,
	0xcc, 0x4d, 0xcc, 0x4b, 0x4c, 0x4f, 0xd5, 0x07, 0x31, 0x93, 0x12, 0x8b, 0x53, 0xf5, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0x38, 0x60, 0x7c, 0x29, 0x91, 0xe4, 0xfc, 0xdc, 0xdc, 0xfc, 0x3c, 0xfd,
	0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0x74, 0xa8, 0xbc, 0x94, 0x68, 0x72, 0x51, 0x62, 0x79, 0x4e, 0x6a,
	0x91, 0x7e, 0x4e, 0x66, 0x62, 0x5e, 0x56, 0x66, 0x22, 0x44, 0x58, 0x49, 0x87, 0x8b, 0xcf, 0x39,
	0x3f, 0x2f, 0x2f, 0x35, 0xb9, 0x24, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x8a, 0x0b,
	0x6e, 0x94, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x9c, 0xaf, 0xe4, 0xc4, 0xc5, 0xeb, 0x99,
	0x57, 0x9c, 0x5a, 0x04, 0x57, 0x6c, 0xc8, 0xc5, 0x95, 0x91, 0x5f, 0x5a, 0x9c, 0x1a, 0x9f, 0x99,
	0x97, 0x96, 0x0f, 0x56, 0xce, 0x6d, 0x24, 0xa4, 0x07, 0xb3, 0xc2, 0x03, 0x24, 0xe5, 0x99, 0x97,
	0x96, 0x1f, 0xc4, 0x99, 0x01, 0x63, 0x1a, 0x35, 0x32, 0x72, 0xf1, 0xbb, 0x40, 0x0d, 0x0c, 0x4e,
	0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x32, 0xe7, 0x62, 0x87, 0xba, 0x42, 0x48, 0x42, 0x0f, 0xee,
	0x31, 0x54, 0x87, 0x49, 0x09, 0xe8, 0xc1, 0x7c, 0xe4, 0x0b, 0xa1, 0x95, 0x18, 0x84, 0x4c, 0xb9,
	0xd8, 0x20, 0x0e, 0x12, 0x12, 0x47, 0xe8, 0x43, 0x71, 0x22, 0x36, 0x6d, 0x4e, 0x8a, 0x51, 0xf2,
	0xe9, 0xf9, 0x7a, 0x95, 0x95, 0xa5, 0xa9, 0x7a, 0x29, 0xa9, 0x65, 0xfa, 0x58, 0x42, 0x35, 0x89,
	0x0d, 0x1c, 0x3e, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xc0, 0xa1, 0x93, 0x73, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for DatabaseService service

type DatabaseServiceClient interface {
	Connect(ctx context.Context, in *ConnectRequest, opts ...client.CallOption) (*message.Message, error)
	Insert(ctx context.Context, in *InsertRequest, opts ...client.CallOption) (*message.Message, error)
}

type databaseServiceClient struct {
	c           client.Client
	serviceName string
}

func NewDatabaseServiceClient(serviceName string, c client.Client) DatabaseServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "database"
	}
	return &databaseServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *databaseServiceClient) Connect(ctx context.Context, in *ConnectRequest, opts ...client.CallOption) (*message.Message, error) {
	req := c.c.NewRequest(c.serviceName, "DatabaseService.Connect", in)
	out := new(message.Message)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) Insert(ctx context.Context, in *InsertRequest, opts ...client.CallOption) (*message.Message, error) {
	req := c.c.NewRequest(c.serviceName, "DatabaseService.Insert", in)
	out := new(message.Message)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DatabaseService service

type DatabaseServiceHandler interface {
	Connect(context.Context, *ConnectRequest, *message.Message) error
	Insert(context.Context, *InsertRequest, *message.Message) error
}

func RegisterDatabaseServiceHandler(s server.Server, hdlr DatabaseServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&DatabaseService{hdlr}, opts...))
}

type DatabaseService struct {
	DatabaseServiceHandler
}

func (h *DatabaseService) Connect(ctx context.Context, in *ConnectRequest, out *message.Message) error {
	return h.DatabaseServiceHandler.Connect(ctx, in, out)
}

func (h *DatabaseService) Insert(ctx context.Context, in *InsertRequest, out *message.Message) error {
	return h.DatabaseServiceHandler.Insert(ctx, in, out)
}
