// Code generated by protoc-gen-go. DO NOT EDIT.
// source: disgover.proto

package disgover

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_disgover_a6811bf8d9f5ed5a, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Endpoint struct {
	Host                 string   `protobuf:"bytes,1,opt,name=Host,proto3" json:"Host,omitempty"`
	Port                 int64    `protobuf:"varint,2,opt,name=Port,proto3" json:"Port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_disgover_a6811bf8d9f5ed5a, []int{1}
}
func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (dst *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(dst, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *Endpoint) GetPort() int64 {
	if m != nil {
		return m.Port
	}
	return 0
}

type Node struct {
	Address              string    `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	Endpoint             *Endpoint `protobuf:"bytes,2,opt,name=Endpoint,proto3" json:"Endpoint,omitempty"`
	Type                 string    `protobuf:"bytes,3,opt,name=Type,proto3" json:"Type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_disgover_a6811bf8d9f5ed5a, []int{2}
}
func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (dst *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(dst, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Node) GetEndpoint() *Endpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *Node) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type NodeList struct {
	Delegates            []*Node  `protobuf:"bytes,1,rep,name=Delegates,proto3" json:"Delegates,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeList) Reset()         { *m = NodeList{} }
func (m *NodeList) String() string { return proto.CompactTextString(m) }
func (*NodeList) ProtoMessage()    {}
func (*NodeList) Descriptor() ([]byte, []int) {
	return fileDescriptor_disgover_a6811bf8d9f5ed5a, []int{3}
}
func (m *NodeList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeList.Unmarshal(m, b)
}
func (m *NodeList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeList.Marshal(b, m, deterministic)
}
func (dst *NodeList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeList.Merge(dst, src)
}
func (m *NodeList) XXX_Size() int {
	return xxx_messageInfo_NodeList.Size(m)
}
func (m *NodeList) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeList.DiscardUnknown(m)
}

var xxx_messageInfo_NodeList proto.InternalMessageInfo

func (m *NodeList) GetDelegates() []*Node {
	if m != nil {
		return m.Delegates
	}
	return nil
}

func init() {
	proto.RegisterType((*Empty)(nil), "disgover.Empty")
	proto.RegisterType((*Endpoint)(nil), "disgover.Endpoint")
	proto.RegisterType((*Node)(nil), "disgover.Node")
	proto.RegisterType((*NodeList)(nil), "disgover.NodeList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DisgoverGrpcClient is the client API for DisgoverGrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DisgoverGrpcClient interface {
	PingSeedGrpc(ctx context.Context, in *Node, opts ...grpc.CallOption) (*NodeList, error)
	UpdateGrpc(ctx context.Context, in *NodeList, opts ...grpc.CallOption) (*Empty, error)
}

type disgoverGrpcClient struct {
	cc *grpc.ClientConn
}

func NewDisgoverGrpcClient(cc *grpc.ClientConn) DisgoverGrpcClient {
	return &disgoverGrpcClient{cc}
}

func (c *disgoverGrpcClient) PingSeedGrpc(ctx context.Context, in *Node, opts ...grpc.CallOption) (*NodeList, error) {
	out := new(NodeList)
	err := c.cc.Invoke(ctx, "/disgover.DisgoverGrpc/PingSeedGrpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *disgoverGrpcClient) UpdateGrpc(ctx context.Context, in *NodeList, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/disgover.DisgoverGrpc/UpdateGrpc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DisgoverGrpcServer is the server API for DisgoverGrpc service.
type DisgoverGrpcServer interface {
	PingSeedGrpc(context.Context, *Node) (*NodeList, error)
	UpdateGrpc(context.Context, *NodeList) (*Empty, error)
}

func RegisterDisgoverGrpcServer(s *grpc.Server, srv DisgoverGrpcServer) {
	s.RegisterService(&_DisgoverGrpc_serviceDesc, srv)
}

func _DisgoverGrpc_PingSeedGrpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisgoverGrpcServer).PingSeedGrpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/disgover.DisgoverGrpc/PingSeedGrpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisgoverGrpcServer).PingSeedGrpc(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _DisgoverGrpc_UpdateGrpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NodeList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisgoverGrpcServer).UpdateGrpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/disgover.DisgoverGrpc/UpdateGrpc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisgoverGrpcServer).UpdateGrpc(ctx, req.(*NodeList))
	}
	return interceptor(ctx, in, info, handler)
}

var _DisgoverGrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "disgover.DisgoverGrpc",
	HandlerType: (*DisgoverGrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PingSeedGrpc",
			Handler:    _DisgoverGrpc_PingSeedGrpc_Handler,
		},
		{
			MethodName: "UpdateGrpc",
			Handler:    _DisgoverGrpc_UpdateGrpc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "disgover.proto",
}

func init() { proto.RegisterFile("disgover.proto", fileDescriptor_disgover_a6811bf8d9f5ed5a) }

var fileDescriptor_disgover_a6811bf8d9f5ed5a = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x1b, 0x53, 0x6d, 0x3a, 0x2d, 0x15, 0xe6, 0x14, 0x7a, 0x2a, 0x7b, 0xca, 0x41, 0x72,
	0x48, 0x3d, 0x78, 0x15, 0x5a, 0xf4, 0x20, 0x52, 0xa2, 0xfe, 0x80, 0xea, 0x0c, 0x61, 0x41, 0xb3,
	0xcb, 0xee, 0x20, 0xe4, 0xdf, 0xcb, 0xae, 0x8d, 0x31, 0xe2, 0xed, 0xbd, 0xcc, 0x7b, 0xf3, 0x65,
	0x07, 0x56, 0xa4, 0x7d, 0x63, 0x3e, 0xd9, 0x95, 0xd6, 0x19, 0x31, 0x98, 0xf5, 0x5e, 0xcd, 0xe0,
	0x7c, 0xff, 0x61, 0xa5, 0x53, 0x15, 0x64, 0xfb, 0x96, 0xac, 0xd1, 0xad, 0x20, 0xc2, 0xf4, 0xde,
	0x78, 0xc9, 0x93, 0x4d, 0x52, 0xcc, 0xeb, 0xa8, 0xc3, 0xb7, 0x83, 0x71, 0x92, 0x9f, 0x6d, 0x92,
	0x22, 0xad, 0xa3, 0x56, 0x04, 0xd3, 0x47, 0x43, 0x8c, 0x39, 0xcc, 0x6e, 0x89, 0x1c, 0x7b, 0x7f,
	0xaa, 0xf4, 0x16, 0xcb, 0x61, 0x6b, 0x6c, 0x2e, 0x2a, 0x2c, 0x7f, 0xfe, 0xa5, 0x9f, 0xd4, 0x23,
	0xf2, 0x73, 0x67, 0x39, 0x4f, 0xbf, 0xc9, 0x41, 0xab, 0x1b, 0xc8, 0x02, 0xe5, 0x41, 0x7b, 0xc1,
	0x2b, 0x98, 0xef, 0xf8, 0x9d, 0x9b, 0xa3, 0x70, 0x60, 0xa5, 0xc5, 0xa2, 0x5a, 0x0d, 0x0b, 0x43,
	0xac, 0x1e, 0x02, 0x55, 0x07, 0xcb, 0xdd, 0x69, 0x76, 0xe7, 0xec, 0x1b, 0x5e, 0xc3, 0xf2, 0xa0,
	0xdb, 0xe6, 0x89, 0x99, 0xa2, 0xff, 0x53, 0x5d, 0xe3, 0xd8, 0x07, 0xa2, 0x9a, 0xe0, 0x16, 0xe0,
	0xc5, 0xd2, 0x51, 0x38, 0x76, 0xfe, 0xc9, 0xac, 0x2f, 0x7f, 0xbd, 0x29, 0x1e, 0x73, 0xf2, 0x7a,
	0x11, 0x0f, 0xbd, 0xfd, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x63, 0x14, 0x81, 0x5a, 0x7a, 0x01, 0x00,
	0x00,
}
