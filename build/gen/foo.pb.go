// Code generated by protoc-gen-go. DO NOT EDIT.
// source: foo.proto

package foo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
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

type IntroRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntroRequest) Reset()         { *m = IntroRequest{} }
func (m *IntroRequest) String() string { return proto.CompactTextString(m) }
func (*IntroRequest) ProtoMessage()    {}
func (*IntroRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce1e2eec643ca48, []int{0}
}

func (m *IntroRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntroRequest.Unmarshal(m, b)
}
func (m *IntroRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntroRequest.Marshal(b, m, deterministic)
}
func (m *IntroRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntroRequest.Merge(m, src)
}
func (m *IntroRequest) XXX_Size() int {
	return xxx_messageInfo_IntroRequest.Size(m)
}
func (m *IntroRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IntroRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IntroRequest proto.InternalMessageInfo

func (m *IntroRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type IntroResponse struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IntroResponse) Reset()         { *m = IntroResponse{} }
func (m *IntroResponse) String() string { return proto.CompactTextString(m) }
func (*IntroResponse) ProtoMessage()    {}
func (*IntroResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7ce1e2eec643ca48, []int{1}
}

func (m *IntroResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntroResponse.Unmarshal(m, b)
}
func (m *IntroResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntroResponse.Marshal(b, m, deterministic)
}
func (m *IntroResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntroResponse.Merge(m, src)
}
func (m *IntroResponse) XXX_Size() int {
	return xxx_messageInfo_IntroResponse.Size(m)
}
func (m *IntroResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IntroResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IntroResponse proto.InternalMessageInfo

func (m *IntroResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*IntroRequest)(nil), "foo.IntroRequest")
	proto.RegisterType((*IntroResponse)(nil), "foo.IntroResponse")
}

func init() { proto.RegisterFile("foo.proto", fileDescriptor_7ce1e2eec643ca48) }

var fileDescriptor_7ce1e2eec643ca48 = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0xcb, 0xcf, 0xd7,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0xcb, 0xcf, 0x97, 0x92, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc,
	0xcf, 0x2b, 0x86, 0x28, 0x51, 0x52, 0xe2, 0xe2, 0xf1, 0xcc, 0x2b, 0x29, 0xca, 0x0f, 0x4a, 0x2d,
	0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x14, 0xb9, 0x78, 0xa1, 0x6a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a,
	0x53, 0x85, 0x04, 0xb8, 0x98, 0x73, 0x8b, 0xd3, 0xa1, 0x6a, 0x40, 0x4c, 0xa3, 0x56, 0x46, 0x2e,
	0x1e, 0x8f, 0xd4, 0x9c, 0x9c, 0xfc, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0x21, 0x3b, 0x2e,
	0x56, 0xb0, 0x1e, 0x21, 0x41, 0x3d, 0x90, 0x7b, 0x90, 0xed, 0x90, 0x12, 0x42, 0x16, 0x82, 0x18,
	0xa9, 0xc4, 0xd7, 0x74, 0xf9, 0xc9, 0x64, 0x26, 0x0e, 0x21, 0x36, 0xfd, 0x0c, 0x90, 0x31, 0x42,
	0xa6, 0x5c, 0x5c, 0x60, 0x05, 0x8e, 0xe9, 0x89, 0x99, 0x79, 0xc4, 0x1a, 0xc2, 0x90, 0xc4, 0x06,
	0xf6, 0x95, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xa8, 0x64, 0xb9, 0x05, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HelloServiceClient interface {
	Intro(ctx context.Context, in *IntroRequest, opts ...grpc.CallOption) (*IntroResponse, error)
	IntroAgain(ctx context.Context, in *IntroRequest, opts ...grpc.CallOption) (*IntroResponse, error)
}

type helloServiceClient struct {
	cc *grpc.ClientConn
}

func NewHelloServiceClient(cc *grpc.ClientConn) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) Intro(ctx context.Context, in *IntroRequest, opts ...grpc.CallOption) (*IntroResponse, error) {
	out := new(IntroResponse)
	err := c.cc.Invoke(ctx, "/foo.HelloService/Intro", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) IntroAgain(ctx context.Context, in *IntroRequest, opts ...grpc.CallOption) (*IntroResponse, error) {
	out := new(IntroResponse)
	err := c.cc.Invoke(ctx, "/foo.HelloService/IntroAgain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HelloServiceServer is the server API for HelloService service.
type HelloServiceServer interface {
	Intro(context.Context, *IntroRequest) (*IntroResponse, error)
	IntroAgain(context.Context, *IntroRequest) (*IntroResponse, error)
}

func RegisterHelloServiceServer(s *grpc.Server, srv HelloServiceServer) {
	s.RegisterService(&_HelloService_serviceDesc, srv)
}

func _HelloService_Intro_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Intro(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foo.HelloService/Intro",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Intro(ctx, req.(*IntroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_IntroAgain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IntroRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).IntroAgain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/foo.HelloService/IntroAgain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).IntroAgain(ctx, req.(*IntroRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "foo.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Intro",
			Handler:    _HelloService_Intro_Handler,
		},
		{
			MethodName: "IntroAgain",
			Handler:    _HelloService_IntroAgain_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "foo.proto",
}
