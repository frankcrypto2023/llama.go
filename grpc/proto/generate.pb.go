// Code generated by protoc-gen-go. DO NOT EDIT.
// source: generate.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type GenerateRequest struct {
	Prompt               string   `protobuf:"bytes,1,opt,name=prompt,proto3" json:"prompt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateRequest) Reset()         { *m = GenerateRequest{} }
func (m *GenerateRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateRequest) ProtoMessage()    {}
func (*GenerateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_generate_da1ff9b7da0279ba, []int{0}
}
func (m *GenerateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateRequest.Unmarshal(m, b)
}
func (m *GenerateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateRequest.Marshal(b, m, deterministic)
}
func (dst *GenerateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateRequest.Merge(dst, src)
}
func (m *GenerateRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateRequest.Size(m)
}
func (m *GenerateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateRequest proto.InternalMessageInfo

func (m *GenerateRequest) GetPrompt() string {
	if m != nil {
		return m.Prompt
	}
	return ""
}

type GenerateResponse struct {
	Content              string   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateResponse) Reset()         { *m = GenerateResponse{} }
func (m *GenerateResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateResponse) ProtoMessage()    {}
func (*GenerateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_generate_da1ff9b7da0279ba, []int{1}
}
func (m *GenerateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateResponse.Unmarshal(m, b)
}
func (m *GenerateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateResponse.Marshal(b, m, deterministic)
}
func (dst *GenerateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateResponse.Merge(dst, src)
}
func (m *GenerateResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateResponse.Size(m)
}
func (m *GenerateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateResponse proto.InternalMessageInfo

func (m *GenerateResponse) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func init() {
	proto.RegisterType((*GenerateRequest)(nil), "proto.GenerateRequest")
	proto.RegisterType((*GenerateResponse)(nil), "proto.GenerateResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GenerateClient is the client API for Generate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GenerateClient interface {
	Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error)
}

type generateClient struct {
	cc *grpc.ClientConn
}

func NewGenerateClient(cc *grpc.ClientConn) GenerateClient {
	return &generateClient{cc}
}

func (c *generateClient) Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error) {
	out := new(GenerateResponse)
	err := c.cc.Invoke(ctx, "/proto.Generate/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GenerateServer is the server API for Generate service.
type GenerateServer interface {
	Generate(context.Context, *GenerateRequest) (*GenerateResponse, error)
}

func RegisterGenerateServer(s *grpc.Server, srv GenerateServer) {
	s.RegisterService(&_Generate_serviceDesc, srv)
}

func _Generate_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GenerateServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Generate/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GenerateServer).Generate(ctx, req.(*GenerateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Generate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Generate",
	HandlerType: (*GenerateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _Generate_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "generate.proto",
}

func init() { proto.RegisterFile("generate.proto", fileDescriptor_generate_da1ff9b7da0279ba) }

var fileDescriptor_generate_da1ff9b7da0279ba = []byte{
	// 175 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0x4f, 0xcd, 0x4b,
	0x2d, 0x4a, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0x32,
	0xe9, 0xf9, 0xf9, 0xe9, 0x39, 0xa9, 0xfa, 0x89, 0x05, 0x99, 0xfa, 0x89, 0x79, 0x79, 0xf9, 0x25,
	0x89, 0x25, 0x99, 0xf9, 0x79, 0xc5, 0x10, 0x45, 0x4a, 0x9a, 0x5c, 0xfc, 0xee, 0x50, 0x6d, 0x41,
	0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x62, 0x5c, 0x6c, 0x05, 0x45, 0xf9, 0xb9, 0x05, 0x25,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x50, 0x9e, 0x92, 0x0e, 0x97, 0x00, 0x42, 0x69, 0x71,
	0x41, 0x7e, 0x5e, 0x71, 0xaa, 0x90, 0x04, 0x17, 0x7b, 0x72, 0x7e, 0x5e, 0x49, 0x6a, 0x1e, 0x4c,
	0x31, 0x8c, 0x6b, 0x94, 0xc0, 0xc5, 0x01, 0x53, 0x2d, 0x14, 0x82, 0xc4, 0x16, 0x83, 0x58, 0xac,
	0x87, 0x66, 0xab, 0x94, 0x38, 0x86, 0x38, 0xc4, 0x0a, 0x25, 0xf1, 0xa6, 0xcb, 0x4f, 0x26, 0x33,
	0x09, 0x2a, 0xf1, 0xe8, 0x97, 0x19, 0xea, 0xc3, 0xbc, 0x68, 0xc5, 0xa8, 0x95, 0xc4, 0x06, 0xd6,
	0x60, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x2d, 0xec, 0x1a, 0xf8, 0x00, 0x00, 0x00,
}
