// Code generated by protoc-gen-go. DO NOT EDIT.
// source: echo.proto

package echopb // import "github.com/vvakame/til/appengine/go111-internal-grpc/echopb"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

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

type SayRequest struct {
	MessageId            string   `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	MessageBody          string   `protobuf:"bytes,2,opt,name=message_body,json=messageBody,proto3" json:"message_body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayRequest) Reset()         { *m = SayRequest{} }
func (m *SayRequest) String() string { return proto.CompactTextString(m) }
func (*SayRequest) ProtoMessage()    {}
func (*SayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_f031054056f06e69, []int{0}
}
func (m *SayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayRequest.Unmarshal(m, b)
}
func (m *SayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayRequest.Marshal(b, m, deterministic)
}
func (dst *SayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayRequest.Merge(dst, src)
}
func (m *SayRequest) XXX_Size() int {
	return xxx_messageInfo_SayRequest.Size(m)
}
func (m *SayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SayRequest proto.InternalMessageInfo

func (m *SayRequest) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *SayRequest) GetMessageBody() string {
	if m != nil {
		return m.MessageBody
	}
	return ""
}

type SayResponse struct {
	MessageId            string               `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	MessageBody          string               `protobuf:"bytes,2,opt,name=message_body,json=messageBody,proto3" json:"message_body,omitempty"`
	Received             *timestamp.Timestamp `protobuf:"bytes,3,opt,name=received,proto3" json:"received,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SayResponse) Reset()         { *m = SayResponse{} }
func (m *SayResponse) String() string { return proto.CompactTextString(m) }
func (*SayResponse) ProtoMessage()    {}
func (*SayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_echo_f031054056f06e69, []int{1}
}
func (m *SayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayResponse.Unmarshal(m, b)
}
func (m *SayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayResponse.Marshal(b, m, deterministic)
}
func (dst *SayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayResponse.Merge(dst, src)
}
func (m *SayResponse) XXX_Size() int {
	return xxx_messageInfo_SayResponse.Size(m)
}
func (m *SayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SayResponse proto.InternalMessageInfo

func (m *SayResponse) GetMessageId() string {
	if m != nil {
		return m.MessageId
	}
	return ""
}

func (m *SayResponse) GetMessageBody() string {
	if m != nil {
		return m.MessageBody
	}
	return ""
}

func (m *SayResponse) GetReceived() *timestamp.Timestamp {
	if m != nil {
		return m.Received
	}
	return nil
}

func init() {
	proto.RegisterType((*SayRequest)(nil), "vvakame.echo.SayRequest")
	proto.RegisterType((*SayResponse)(nil), "vvakame.echo.SayResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EchoClient is the client API for Echo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EchoClient interface {
	Say(ctx context.Context, in *SayRequest, opts ...grpc.CallOption) (*SayResponse, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) Say(ctx context.Context, in *SayRequest, opts ...grpc.CallOption) (*SayResponse, error) {
	out := new(SayResponse)
	err := c.cc.Invoke(ctx, "/vvakame.echo.Echo/Say", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServer is the server API for Echo service.
type EchoServer interface {
	Say(context.Context, *SayRequest) (*SayResponse, error)
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Say(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vvakame.echo.Echo/Say",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Say(ctx, req.(*SayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vvakame.echo.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Say",
			Handler:    _Echo_Say_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echo.proto",
}

func init() { proto.RegisterFile("echo.proto", fileDescriptor_echo_f031054056f06e69) }

var fileDescriptor_echo_f031054056f06e69 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x90, 0x3d, 0x4f, 0xf3, 0x30,
	0x14, 0x85, 0xdf, 0xbc, 0x45, 0x88, 0xde, 0x76, 0xf2, 0x14, 0x22, 0x21, 0x4a, 0xa6, 0x2e, 0xb5,
	0x95, 0x22, 0xb1, 0x00, 0x4b, 0x05, 0x03, 0x0b, 0x43, 0xca, 0xc4, 0x82, 0x9c, 0xe4, 0xe2, 0x58,
	0x24, 0xb9, 0x26, 0x76, 0x22, 0xe5, 0x17, 0xf0, 0xb7, 0x51, 0xf3, 0x01, 0x0c, 0x6c, 0xac, 0xe7,
	0x3c, 0x3a, 0xba, 0xcf, 0x05, 0xc0, 0x34, 0x27, 0x6e, 0x6a, 0x72, 0xc4, 0x96, 0x6d, 0x2b, 0xdf,
	0x64, 0x89, 0xfc, 0x90, 0x05, 0xe7, 0x8a, 0x48, 0x15, 0x28, 0xfa, 0x2e, 0x69, 0x5e, 0x85, 0xd3,
	0x25, 0x5a, 0x27, 0x4b, 0x33, 0xe0, 0xe1, 0x23, 0xc0, 0x5e, 0x76, 0x31, 0xbe, 0x37, 0x68, 0x1d,
	0x3b, 0x03, 0x28, 0xd1, 0x5a, 0xa9, 0xf0, 0x45, 0x67, 0xbe, 0xb7, 0xf2, 0xd6, 0xf3, 0x78, 0x3e,
	0x26, 0x0f, 0x19, 0xbb, 0x80, 0xe5, 0x54, 0x27, 0x94, 0x75, 0xfe, 0xff, 0x1e, 0x58, 0x8c, 0xd9,
	0x8e, 0xb2, 0x2e, 0xfc, 0xf0, 0x60, 0xd1, 0x0f, 0x5a, 0x43, 0x95, 0xc5, 0xbf, 0x2f, 0xb2, 0x2b,
	0x38, 0xa9, 0x31, 0x45, 0xdd, 0x62, 0xe6, 0xcf, 0x56, 0xde, 0x7a, 0xb1, 0x0d, 0xf8, 0x60, 0xc5,
	0x27, 0x2b, 0xfe, 0x34, 0x59, 0xc5, 0x5f, 0xec, 0xf6, 0x0e, 0x8e, 0xee, 0xd3, 0x9c, 0xd8, 0x0d,
	0xcc, 0xf6, 0xb2, 0x63, 0x3e, 0xff, 0xf9, 0x18, 0xfe, 0x2d, 0x1d, 0x9c, 0xfe, 0xd2, 0x0c, 0xd7,
	0x87, 0xff, 0x76, 0xb7, 0xcf, 0xd7, 0x4a, 0xbb, 0xbc, 0x49, 0x78, 0x4a, 0xa5, 0x18, 0x41, 0xe1,
	0x74, 0x21, 0xa4, 0x31, 0x58, 0x29, 0x5d, 0xa1, 0x50, 0x14, 0x45, 0xd1, 0x46, 0x57, 0x0e, 0xeb,
	0x4a, 0x16, 0x1b, 0x55, 0x9b, 0x54, 0x1c, 0xa6, 0x4c, 0x92, 0x1c, 0xf7, 0x27, 0x5e, 0x7e, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x8e, 0x46, 0x7c, 0x5a, 0xa2, 0x01, 0x00, 0x00,
}
