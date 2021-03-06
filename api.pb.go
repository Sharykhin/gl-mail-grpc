// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	Empty
	CountResponse
	FailMailRequest
	FailMailResponse
	FailMailFilter
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

// Empty request
type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CountResponse struct {
	Total int64 `protobuf:"varint,1,opt,name=total" json:"total,omitempty"`
}

func (m *CountResponse) Reset()                    { *m = CountResponse{} }
func (m *CountResponse) String() string            { return proto.CompactTextString(m) }
func (*CountResponse) ProtoMessage()               {}
func (*CountResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CountResponse) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

// Request message for creating a new fail mail
type FailMailRequest struct {
	Action  string `protobuf:"bytes,1,opt,name=action" json:"action,omitempty"`
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	Reason  string `protobuf:"bytes,3,opt,name=reason" json:"reason,omitempty"`
}

func (m *FailMailRequest) Reset()                    { *m = FailMailRequest{} }
func (m *FailMailRequest) String() string            { return proto.CompactTextString(m) }
func (*FailMailRequest) ProtoMessage()               {}
func (*FailMailRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FailMailRequest) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *FailMailRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *FailMailRequest) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

// Response of fail mail
type FailMailResponse struct {
	ID        int64  `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	Action    string `protobuf:"bytes,2,opt,name=action" json:"action,omitempty"`
	Payload   []byte `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Reason    string `protobuf:"bytes,4,opt,name=reason" json:"reason,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt" json:"created_at,omitempty"`
	DeletedAt string `protobuf:"bytes,6,opt,name=deleted_at,json=deletedAt" json:"deleted_at,omitempty"`
}

func (m *FailMailResponse) Reset()                    { *m = FailMailResponse{} }
func (m *FailMailResponse) String() string            { return proto.CompactTextString(m) }
func (*FailMailResponse) ProtoMessage()               {}
func (*FailMailResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *FailMailResponse) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *FailMailResponse) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *FailMailResponse) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *FailMailResponse) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *FailMailResponse) GetCreatedAt() string {
	if m != nil {
		return m.CreatedAt
	}
	return ""
}

func (m *FailMailResponse) GetDeletedAt() string {
	if m != nil {
		return m.DeletedAt
	}
	return ""
}

// Filter for streaming fail mails
type FailMailFilter struct {
	Limit  int64 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	Offset int64 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
}

func (m *FailMailFilter) Reset()                    { *m = FailMailFilter{} }
func (m *FailMailFilter) String() string            { return proto.CompactTextString(m) }
func (*FailMailFilter) ProtoMessage()               {}
func (*FailMailFilter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *FailMailFilter) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *FailMailFilter) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "api.Empty")
	proto.RegisterType((*CountResponse)(nil), "api.CountResponse")
	proto.RegisterType((*FailMailRequest)(nil), "api.FailMailRequest")
	proto.RegisterType((*FailMailResponse)(nil), "api.FailMailResponse")
	proto.RegisterType((*FailMailFilter)(nil), "api.FailMailFilter")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for FailMail service

type FailMailClient interface {
	CreateFailMail(ctx context.Context, in *FailMailRequest, opts ...grpc.CallOption) (*FailMailResponse, error)
	GetFailMails(ctx context.Context, in *FailMailFilter, opts ...grpc.CallOption) (FailMail_GetFailMailsClient, error)
	CountFailMails(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CountResponse, error)
}

type failMailClient struct {
	cc *grpc.ClientConn
}

func NewFailMailClient(cc *grpc.ClientConn) FailMailClient {
	return &failMailClient{cc}
}

func (c *failMailClient) CreateFailMail(ctx context.Context, in *FailMailRequest, opts ...grpc.CallOption) (*FailMailResponse, error) {
	out := new(FailMailResponse)
	err := grpc.Invoke(ctx, "/api.FailMail/CreateFailMail", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *failMailClient) GetFailMails(ctx context.Context, in *FailMailFilter, opts ...grpc.CallOption) (FailMail_GetFailMailsClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FailMail_serviceDesc.Streams[0], c.cc, "/api.FailMail/GetFailMails", opts...)
	if err != nil {
		return nil, err
	}
	x := &failMailGetFailMailsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FailMail_GetFailMailsClient interface {
	Recv() (*FailMailResponse, error)
	grpc.ClientStream
}

type failMailGetFailMailsClient struct {
	grpc.ClientStream
}

func (x *failMailGetFailMailsClient) Recv() (*FailMailResponse, error) {
	m := new(FailMailResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *failMailClient) CountFailMails(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*CountResponse, error) {
	out := new(CountResponse)
	err := grpc.Invoke(ctx, "/api.FailMail/CountFailMails", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FailMail service

type FailMailServer interface {
	CreateFailMail(context.Context, *FailMailRequest) (*FailMailResponse, error)
	GetFailMails(*FailMailFilter, FailMail_GetFailMailsServer) error
	CountFailMails(context.Context, *Empty) (*CountResponse, error)
}

func RegisterFailMailServer(s *grpc.Server, srv FailMailServer) {
	s.RegisterService(&_FailMail_serviceDesc, srv)
}

func _FailMail_CreateFailMail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FailMailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FailMailServer).CreateFailMail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.FailMail/CreateFailMail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FailMailServer).CreateFailMail(ctx, req.(*FailMailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FailMail_GetFailMails_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FailMailFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FailMailServer).GetFailMails(m, &failMailGetFailMailsServer{stream})
}

type FailMail_GetFailMailsServer interface {
	Send(*FailMailResponse) error
	grpc.ServerStream
}

type failMailGetFailMailsServer struct {
	grpc.ServerStream
}

func (x *failMailGetFailMailsServer) Send(m *FailMailResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _FailMail_CountFailMails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FailMailServer).CountFailMails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.FailMail/CountFailMails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FailMailServer).CountFailMails(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _FailMail_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.FailMail",
	HandlerType: (*FailMailServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFailMail",
			Handler:    _FailMail_CreateFailMail_Handler,
		},
		{
			MethodName: "CountFailMails",
			Handler:    _FailMail_CountFailMails_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetFailMails",
			Handler:       _FailMail_GetFailMails_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x6b, 0xc2, 0x40,
	0x10, 0x35, 0x49, 0xd5, 0x3a, 0xd8, 0xb4, 0x6c, 0x6d, 0x09, 0x42, 0x41, 0x16, 0x0a, 0x9e, 0xa4,
	0xd8, 0xbb, 0x45, 0xb4, 0x16, 0x0f, 0xbd, 0xe4, 0xda, 0x43, 0xd9, 0xea, 0x08, 0x0b, 0x6b, 0x36,
	0xcd, 0x8e, 0x07, 0xff, 0x52, 0xff, 0x44, 0xff, 0x5a, 0xc9, 0xee, 0xa6, 0x7e, 0x40, 0x6e, 0xfb,
	0xe6, 0xcd, 0xcc, 0x7b, 0xbc, 0x1d, 0xe8, 0x88, 0x5c, 0x8e, 0xf2, 0x42, 0x93, 0x66, 0x91, 0xc8,
	0x25, 0x6f, 0x43, 0xf3, 0x75, 0x9b, 0xd3, 0x9e, 0x3f, 0xc2, 0xd5, 0x4c, 0xef, 0x32, 0x4a, 0xd1,
	0xe4, 0x3a, 0x33, 0xc8, 0x7a, 0xd0, 0x24, 0x4d, 0x42, 0x25, 0xc1, 0x20, 0x18, 0x46, 0xa9, 0x03,
	0xfc, 0x03, 0xae, 0x17, 0x42, 0xaa, 0x77, 0x21, 0x55, 0x8a, 0xdf, 0x3b, 0x34, 0xc4, 0xee, 0xa1,
	0x25, 0x56, 0x24, 0x75, 0x66, 0x3b, 0x3b, 0xa9, 0x47, 0x2c, 0x81, 0x76, 0x2e, 0xf6, 0x4a, 0x8b,
	0x75, 0x12, 0x0e, 0x82, 0x61, 0x37, 0xad, 0x60, 0x39, 0x51, 0xa0, 0x30, 0x3a, 0x4b, 0x22, 0x37,
	0xe1, 0x10, 0xff, 0x09, 0xe0, 0xe6, 0xb0, 0xdd, 0xfb, 0x88, 0x21, 0x5c, 0xce, 0xbd, 0x89, 0x70,
	0x39, 0x3f, 0x92, 0x0b, 0xeb, 0xe4, 0xa2, 0x3a, 0xb9, 0x8b, 0x63, 0x39, 0xf6, 0x00, 0xb0, 0x2a,
	0x50, 0x10, 0xae, 0x3f, 0x05, 0x25, 0x4d, 0xcb, 0x75, 0x7c, 0x65, 0x4a, 0x25, 0xbd, 0x46, 0x85,
	0x9e, 0x6e, 0x39, 0xda, 0x57, 0xa6, 0xc4, 0x27, 0x10, 0x57, 0x5e, 0x17, 0x52, 0x11, 0x16, 0x65,
	0x62, 0x4a, 0x6e, 0x25, 0x55, 0x89, 0x59, 0x50, 0xaa, 0xeb, 0xcd, 0xc6, 0x20, 0x59, 0xbf, 0x51,
	0xea, 0xd1, 0xf8, 0x37, 0x80, 0xcb, 0x6a, 0x01, 0x7b, 0x81, 0x78, 0x66, 0x85, 0xff, 0x2b, 0xbd,
	0x51, 0xf9, 0x53, 0x67, 0x59, 0xf7, 0xef, 0xce, 0xaa, 0x2e, 0x23, 0xde, 0x60, 0x13, 0xe8, 0xbe,
	0x21, 0x55, 0x84, 0x61, 0xb7, 0x27, 0x8d, 0xce, 0x60, 0xed, 0xf4, 0x53, 0xc0, 0xc6, 0x10, 0xdb,
	0xef, 0x3f, 0x6c, 0x00, 0xdb, 0x6c, 0x8f, 0xa3, 0xcf, 0xec, 0xfb, 0xe4, 0x3e, 0x78, 0xe3, 0xab,
	0x65, 0xef, 0xe8, 0xf9, 0x2f, 0x00, 0x00, 0xff, 0xff, 0xe0, 0x80, 0xf7, 0xb7, 0x54, 0x02, 0x00,
	0x00,
}
