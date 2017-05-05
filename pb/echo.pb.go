// Code generated by protoc-gen-go.
// source: pb/echo.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/echo.proto

It has these top-level messages:
	Request
	Response
*/
package pb

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

type Request struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Response struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "pb.Request")
	proto.RegisterType((*Response)(nil), "pb.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Echo service

type EchoClient interface {
	Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) Call(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/pb.Echo/Call", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Echo service

type EchoServer interface {
	Call(context.Context, *Request) (*Response, error)
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Call_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Call(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Echo/Call",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Call(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Call",
			Handler:    _Echo_Call_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/echo.proto",
}

func init() { proto.RegisterFile("pb/echo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 123 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x48, 0xd2, 0x4f,
	0x4d, 0xce, 0xc8, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x92, 0xe7,
	0x62, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29,
	0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x94, 0x14, 0xb8, 0x38, 0x82, 0x52,
	0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0xb1, 0xab, 0x30, 0xd2, 0xe6, 0x62, 0x71, 0x4d, 0xce, 0xc8,
	0x17, 0x52, 0xe6, 0x62, 0x71, 0x4e, 0xcc, 0xc9, 0x11, 0xe2, 0xd6, 0x2b, 0x48, 0xd2, 0x83, 0x1a,
	0x2a, 0xc5, 0x03, 0xe1, 0x40, 0x0c, 0x50, 0x62, 0x48, 0x62, 0x03, 0x5b, 0x6d, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x96, 0x6b, 0x8b, 0x77, 0x8b, 0x00, 0x00, 0x00,
}
