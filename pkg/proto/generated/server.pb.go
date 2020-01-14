// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server.proto

package k3sd

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type K3SServer struct {
	NetworkDevice        string   `protobuf:"bytes,1,opt,name=NetworkDevice,proto3" json:"NetworkDevice,omitempty"`
	TLSSan               string   `protobuf:"bytes,2,opt,name=TLSSan,proto3" json:"TLSSan,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *K3SServer) Reset()         { *m = K3SServer{} }
func (m *K3SServer) String() string { return proto.CompactTextString(m) }
func (*K3SServer) ProtoMessage()    {}
func (*K3SServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{0}
}

func (m *K3SServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_K3SServer.Unmarshal(m, b)
}
func (m *K3SServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_K3SServer.Marshal(b, m, deterministic)
}
func (m *K3SServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_K3SServer.Merge(m, src)
}
func (m *K3SServer) XXX_Size() int {
	return xxx_messageInfo_K3SServer.Size(m)
}
func (m *K3SServer) XXX_DiscardUnknown() {
	xxx_messageInfo_K3SServer.DiscardUnknown(m)
}

var xxx_messageInfo_K3SServer proto.InternalMessageInfo

func (m *K3SServer) GetNetworkDevice() string {
	if m != nil {
		return m.NetworkDevice
	}
	return ""
}

func (m *K3SServer) GetTLSSan() string {
	if m != nil {
		return m.TLSSan
	}
	return ""
}

type K3SServerManagedId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *K3SServerManagedId) Reset()         { *m = K3SServerManagedId{} }
func (m *K3SServerManagedId) String() string { return proto.CompactTextString(m) }
func (*K3SServerManagedId) ProtoMessage()    {}
func (*K3SServerManagedId) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{1}
}

func (m *K3SServerManagedId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_K3SServerManagedId.Unmarshal(m, b)
}
func (m *K3SServerManagedId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_K3SServerManagedId.Marshal(b, m, deterministic)
}
func (m *K3SServerManagedId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_K3SServerManagedId.Merge(m, src)
}
func (m *K3SServerManagedId) XXX_Size() int {
	return xxx_messageInfo_K3SServerManagedId.Size(m)
}
func (m *K3SServerManagedId) XXX_DiscardUnknown() {
	xxx_messageInfo_K3SServerManagedId.DiscardUnknown(m)
}

var xxx_messageInfo_K3SServerManagedId proto.InternalMessageInfo

func (m *K3SServerManagedId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*K3SServer)(nil), "k3sd.K3SServer")
	proto.RegisterType((*K3SServerManagedId)(nil), "k3sd.K3SServerManagedId")
}

func init() { proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7) }

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xc9, 0x36, 0x2e, 0x4e, 0x51, 0xf2,
	0xe4, 0xe2, 0xf4, 0x36, 0x0e, 0x0e, 0x06, 0x4b, 0x08, 0xa9, 0x70, 0xf1, 0xfa, 0xa5, 0x96, 0x94,
	0xe7, 0x17, 0x65, 0xbb, 0xa4, 0x96, 0x65, 0x26, 0xa7, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0xa1, 0x0a, 0x0a, 0x89, 0x71, 0xb1, 0x85, 0xf8, 0x04, 0x07, 0x27, 0xe6, 0x49, 0x30, 0x81, 0xa5,
	0xa1, 0x3c, 0x25, 0x15, 0x2e, 0x21, 0xb8, 0x51, 0xbe, 0x89, 0x79, 0x89, 0xe9, 0xa9, 0x29, 0x9e,
	0x29, 0x42, 0x7c, 0x5c, 0x4c, 0x9e, 0x29, 0x50, 0x83, 0x98, 0x3c, 0x53, 0x8c, 0x3c, 0xb8, 0x04,
	0xd0, 0x54, 0x15, 0x09, 0x99, 0x70, 0xb1, 0x06, 0x97, 0x24, 0x16, 0x95, 0x08, 0xf1, 0xeb, 0x81,
	0x1c, 0xa5, 0x07, 0x57, 0x20, 0x25, 0x81, 0x26, 0x00, 0x37, 0x57, 0x89, 0x21, 0x89, 0x0d, 0xec,
	0x0f, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x85, 0xc5, 0x9c, 0xba, 0xd7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// K3SServerManagerClient is the client API for K3SServerManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type K3SServerManagerClient interface {
	Start(ctx context.Context, in *K3SServer, opts ...grpc.CallOption) (*K3SServerManagedId, error)
}

type k3SServerManagerClient struct {
	cc *grpc.ClientConn
}

func NewK3SServerManagerClient(cc *grpc.ClientConn) K3SServerManagerClient {
	return &k3SServerManagerClient{cc}
}

func (c *k3SServerManagerClient) Start(ctx context.Context, in *K3SServer, opts ...grpc.CallOption) (*K3SServerManagedId, error) {
	out := new(K3SServerManagedId)
	err := c.cc.Invoke(ctx, "/k3sd.K3SServerManager/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// K3SServerManagerServer is the server API for K3SServerManager service.
type K3SServerManagerServer interface {
	Start(context.Context, *K3SServer) (*K3SServerManagedId, error)
}

// UnimplementedK3SServerManagerServer can be embedded to have forward compatible implementations.
type UnimplementedK3SServerManagerServer struct {
}

func (*UnimplementedK3SServerManagerServer) Start(ctx context.Context, req *K3SServer) (*K3SServerManagedId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}

func RegisterK3SServerManagerServer(s *grpc.Server, srv K3SServerManagerServer) {
	s.RegisterService(&_K3SServerManager_serviceDesc, srv)
}

func _K3SServerManager_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(K3SServer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K3SServerManagerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/k3sd.K3SServerManager/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K3SServerManagerServer).Start(ctx, req.(*K3SServer))
	}
	return interceptor(ctx, in, info, handler)
}

var _K3SServerManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "k3sd.K3SServerManager",
	HandlerType: (*K3SServerManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _K3SServerManager_Start_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
