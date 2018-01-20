// Code generated by protoc-gen-go. DO NOT EDIT.
// source: twitterarchive.proto

/*
Package twitterarchive is a generated protocol buffer package.

It is generated from these files:
	twitterarchive.proto

It has these top-level messages:
	CreateRequest
	CreateResponse
*/
package twitterarchive

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

type CreateRequest struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CreateResponse struct {
	Ok bool `protobuf:"varint,1,opt,name=ok" json:"ok,omitempty"`
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "twitterarchive.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "twitterarchive.CreateResponse")
}

func init() { proto.RegisterFile("twitterarchive.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 163 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x29, 0xcf, 0x2c,
	0x29, 0x49, 0x2d, 0x4a, 0x2c, 0x4a, 0xce, 0xc8, 0x2c, 0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0xe2, 0x43, 0x15, 0x55, 0xb2, 0xe5, 0xe2, 0x75, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x0d, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x24, 0xb8, 0xd8, 0x73, 0x53, 0x8b, 0x8b, 0x13, 0xd3,
	0x53, 0x25, 0x98, 0xc0, 0xc2, 0x30, 0xae, 0x92, 0x02, 0x17, 0x1f, 0x4c, 0x7b, 0x71, 0x41, 0x7e,
	0x5e, 0x71, 0xaa, 0x10, 0x1f, 0x17, 0x53, 0x7e, 0x36, 0x58, 0x37, 0x47, 0x10, 0x53, 0x7e, 0xb6,
	0x51, 0x34, 0x17, 0x5f, 0x08, 0xc4, 0x4a, 0x47, 0x88, 0x95, 0x42, 0x9e, 0x5c, 0x6c, 0x10, 0x3d,
	0x42, 0xb2, 0x7a, 0x68, 0x6e, 0x44, 0x71, 0x8a, 0x94, 0x1c, 0x2e, 0x69, 0x88, 0x55, 0x4a, 0x0c,
	0x49, 0x6c, 0x60, 0x4f, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xef, 0x8e, 0x7f, 0x55, 0xec,
	0x00, 0x00, 0x00,
}
