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
	return fileDescriptor_disgover_d7fe0bb1883d4e13, []int{0}
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
	return fileDescriptor_disgover_d7fe0bb1883d4e13, []int{1}
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
	Hash                 string    `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Address              string    `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
	Signature            string    `protobuf:"bytes,3,opt,name=Signature,proto3" json:"Signature,omitempty"`
	Endpoint             *Endpoint `protobuf:"bytes,4,opt,name=Endpoint,proto3" json:"Endpoint,omitempty"`
	Type                 string    `protobuf:"bytes,5,opt,name=Type,proto3" json:"Type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_disgover_d7fe0bb1883d4e13, []int{2}
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

func (m *Node) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Node) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Node) GetSignature() string {
	if m != nil {
		return m.Signature
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
	return fileDescriptor_disgover_d7fe0bb1883d4e13, []int{3}
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

func init() { proto.RegisterFile("disgover.proto", fileDescriptor_disgover_d7fe0bb1883d4e13) }

var fileDescriptor_disgover_d7fe0bb1883d4e13 = []byte{
	// 273 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0xc1, 0x4e, 0xeb, 0x30,
	0x10, 0xac, 0x5f, 0xd2, 0xd7, 0x64, 0x5b, 0x15, 0x69, 0x4f, 0x56, 0xc5, 0x21, 0xca, 0x29, 0x07,
	0x94, 0x43, 0xca, 0x81, 0x2b, 0x52, 0x2b, 0x38, 0x20, 0x54, 0xa5, 0xf0, 0x01, 0x01, 0xaf, 0x52,
	0x4b, 0x10, 0x5b, 0xb6, 0x41, 0xca, 0x8f, 0xf0, 0xbd, 0xc8, 0x86, 0x34, 0x14, 0x71, 0x9b, 0xd9,
	0x99, 0xf1, 0x8e, 0x6d, 0x58, 0x0a, 0x69, 0x5b, 0xf5, 0x4e, 0xa6, 0xd4, 0x46, 0x39, 0x85, 0xc9,
	0xc0, 0xf3, 0x19, 0x4c, 0xb7, 0xaf, 0xda, 0xf5, 0x79, 0x05, 0xc9, 0xb6, 0x13, 0x5a, 0xc9, 0xce,
	0x21, 0x42, 0x7c, 0xab, 0xac, 0xe3, 0x2c, 0x63, 0x45, 0x5a, 0x07, 0xec, 0x67, 0x3b, 0x65, 0x1c,
	0xff, 0x97, 0xb1, 0x22, 0xaa, 0x03, 0xce, 0x3f, 0x18, 0xc4, 0xf7, 0x4a, 0x50, 0x08, 0x34, 0xf6,
	0x70, 0x0c, 0x34, 0xf6, 0x80, 0x1c, 0x66, 0xd7, 0x42, 0x18, 0xb2, 0x36, 0x64, 0xd2, 0x7a, 0xa0,
	0x78, 0x0e, 0xe9, 0x5e, 0xb6, 0x5d, 0xe3, 0xde, 0x0c, 0xf1, 0x28, 0x68, 0xe3, 0x00, 0xcb, 0xb1,
	0x08, 0x8f, 0x33, 0x56, 0xcc, 0x2b, 0x2c, 0x8f, 0xf5, 0x07, 0xa5, 0x3e, 0x29, 0xfb, 0xd0, 0x6b,
	0xe2, 0xd3, 0xaf, 0xdd, 0x1e, 0xe7, 0x57, 0x90, 0xf8, 0x5e, 0x77, 0xd2, 0x3a, 0xbc, 0x80, 0x74,
	0x43, 0x2f, 0xd4, 0x36, 0x8e, 0x2c, 0x67, 0x59, 0x54, 0xcc, 0xab, 0xe5, 0x78, 0xa0, 0xb7, 0xd5,
	0xa3, 0xa1, 0xea, 0x61, 0xb1, 0xf9, 0xd6, 0x6e, 0x8c, 0x7e, 0xc6, 0x4b, 0x58, 0xec, 0x64, 0xd7,
	0xee, 0x89, 0x44, 0xe0, 0xbf, 0xa2, 0x2b, 0x3c, 0xe5, 0x7e, 0x63, 0x3e, 0xc1, 0x35, 0xc0, 0xa3,
	0x16, 0x8d, 0xa3, 0x90, 0xf9, 0xc3, 0xb3, 0x3a, 0xfb, 0x71, 0xa7, 0xf0, 0xfe, 0x93, 0xa7, 0xff,
	0xe1, 0x6f, 0xd6, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xec, 0x20, 0xdd, 0x07, 0xad, 0x01, 0x00,
	0x00,
}
