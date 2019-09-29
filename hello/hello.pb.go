// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hello/hello.proto

/*
Package hello is a generated protocol buffer package.

It is generated from these files:
	hello/hello.proto

It has these top-level messages:
	HelloRequest
	HelloReply
*/
package hello

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// The request message containing the user's name.
type HelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloRequest) Reset()                    { *m = HelloRequest{} }
func (m *HelloRequest) String() string            { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()               {}
func (*HelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *HelloReply) Reset()                    { *m = HelloReply{} }
func (m *HelloReply) String() string            { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()               {}
func (*HelloReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "hello.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "hello.HelloReply")
}

func init() { proto.RegisterFile("hello/hello.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 138 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcc, 0x48, 0xcd, 0xc9,
	0xc9, 0xd7, 0x07, 0x93, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xac, 0x60, 0x8e, 0x92, 0x12,
	0x17, 0x8f, 0x07, 0x88, 0x11, 0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc4, 0xc5, 0x92,
	0x97, 0x98, 0x9b, 0x2a, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x2b, 0xa9, 0x71, 0x71,
	0x41, 0xd5, 0x14, 0xe4, 0x54, 0x0a, 0x49, 0x70, 0xb1, 0xe7, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xc3,
	0x14, 0xc1, 0xb8, 0x46, 0xf6, 0x5c, 0xec, 0xee, 0x45, 0xa9, 0xa9, 0x25, 0xa9, 0x45, 0x42, 0x26,
	0x5c, 0x1c, 0xc1, 0x89, 0x95, 0x60, 0x5d, 0x42, 0xc2, 0x7a, 0x10, 0x7b, 0x91, 0xed, 0x91, 0x12,
	0x44, 0x15, 0x2c, 0xc8, 0xa9, 0x54, 0x62, 0x48, 0x62, 0x03, 0x3b, 0xcd, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x00, 0xcc, 0xb4, 0x38, 0xaf, 0x00, 0x00, 0x00,
}
