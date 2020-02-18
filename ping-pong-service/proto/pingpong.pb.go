// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pingpong.proto

/*
Package pingpong is a generated protocol buffer package.

It is generated from these files:
	pingpong.proto

It has these top-level messages:
	Ping
	Pong
*/
package pingpong

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

// The ping request containing a message.
type Ping struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Ping) Reset()                    { *m = Ping{} }
func (m *Ping) String() string            { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()               {}
func (*Ping) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Ping) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

// The pong response containing a message.
type Pong struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Pong) Reset()                    { *m = Pong{} }
func (m *Pong) String() string            { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()               {}
func (*Pong) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Pong) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Ping)(nil), "pingpong.Ping")
	proto.RegisterType((*Pong)(nil), "pingpong.Pong")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PingPong service

type PingPongClient interface {
	// Sends a ping
	SendPing(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error)
}

type pingPongClient struct {
	cc *grpc.ClientConn
}

func NewPingPongClient(cc *grpc.ClientConn) PingPongClient {
	return &pingPongClient{cc}
}

func (c *pingPongClient) SendPing(ctx context.Context, in *Ping, opts ...grpc.CallOption) (*Pong, error) {
	out := new(Pong)
	err := grpc.Invoke(ctx, "/pingpong.PingPong/SendPing", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PingPong service

type PingPongServer interface {
	// Sends a ping
	SendPing(context.Context, *Ping) (*Pong, error)
}

func RegisterPingPongServer(s *grpc.Server, srv PingPongServer) {
	s.RegisterService(&_PingPong_serviceDesc, srv)
}

func _PingPong_SendPing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Ping)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingPongServer).SendPing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pingpong.PingPong/SendPing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingPongServer).SendPing(ctx, req.(*Ping))
	}
	return interceptor(ctx, in, info, handler)
}

var _PingPong_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pingpong.PingPong",
	HandlerType: (*PingPongServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendPing",
			Handler:    _PingPong_SendPing_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pingpong.proto",
}

func init() { proto.RegisterFile("pingpong.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 111 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0xc8, 0xcc, 0x4b,
	0x2f, 0xc8, 0xcf, 0x4b, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x80, 0xf1, 0x95, 0x14,
	0xb8, 0x58, 0x02, 0x32, 0xf3, 0xd2, 0x85, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3,
	0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0xb0, 0x8a, 0x7c, 0x7c, 0x2a, 0x8c,
	0x2c, 0xb8, 0x38, 0x40, 0x66, 0x80, 0x55, 0xe9, 0x70, 0x71, 0x04, 0xa7, 0xe6, 0xa5, 0x80, 0xcd,
	0xe4, 0xd3, 0x83, 0x5b, 0x0b, 0xe2, 0x4b, 0x21, 0xf3, 0x41, 0x76, 0x33, 0x24, 0xb1, 0x81, 0x9d,
	0x63, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xff, 0x45, 0x8e, 0x97, 0xa0, 0x00, 0x00, 0x00,
}