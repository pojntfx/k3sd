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

type K3SServerEmptyArgs struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *K3SServerEmptyArgs) Reset()         { *m = K3SServerEmptyArgs{} }
func (m *K3SServerEmptyArgs) String() string { return proto.CompactTextString(m) }
func (*K3SServerEmptyArgs) ProtoMessage()    {}
func (*K3SServerEmptyArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{1}
}

func (m *K3SServerEmptyArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_K3SServerEmptyArgs.Unmarshal(m, b)
}
func (m *K3SServerEmptyArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_K3SServerEmptyArgs.Marshal(b, m, deterministic)
}
func (m *K3SServerEmptyArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_K3SServerEmptyArgs.Merge(m, src)
}
func (m *K3SServerEmptyArgs) XXX_Size() int {
	return xxx_messageInfo_K3SServerEmptyArgs.Size(m)
}
func (m *K3SServerEmptyArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_K3SServerEmptyArgs.DiscardUnknown(m)
}

var xxx_messageInfo_K3SServerEmptyArgs proto.InternalMessageInfo

type K3SServerState struct {
	Running              bool     `protobuf:"varint,1,opt,name=Running,proto3" json:"Running,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *K3SServerState) Reset()         { *m = K3SServerState{} }
func (m *K3SServerState) String() string { return proto.CompactTextString(m) }
func (*K3SServerState) ProtoMessage()    {}
func (*K3SServerState) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{2}
}

func (m *K3SServerState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_K3SServerState.Unmarshal(m, b)
}
func (m *K3SServerState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_K3SServerState.Marshal(b, m, deterministic)
}
func (m *K3SServerState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_K3SServerState.Merge(m, src)
}
func (m *K3SServerState) XXX_Size() int {
	return xxx_messageInfo_K3SServerState.Size(m)
}
func (m *K3SServerState) XXX_DiscardUnknown() {
	xxx_messageInfo_K3SServerState.DiscardUnknown(m)
}

var xxx_messageInfo_K3SServerState proto.InternalMessageInfo

func (m *K3SServerState) GetRunning() bool {
	if m != nil {
		return m.Running
	}
	return false
}

type K3SServerDeletionState struct {
	Deleted              bool     `protobuf:"varint,1,opt,name=Deleted,proto3" json:"Deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *K3SServerDeletionState) Reset()         { *m = K3SServerDeletionState{} }
func (m *K3SServerDeletionState) String() string { return proto.CompactTextString(m) }
func (*K3SServerDeletionState) ProtoMessage()    {}
func (*K3SServerDeletionState) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{3}
}

func (m *K3SServerDeletionState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_K3SServerDeletionState.Unmarshal(m, b)
}
func (m *K3SServerDeletionState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_K3SServerDeletionState.Marshal(b, m, deterministic)
}
func (m *K3SServerDeletionState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_K3SServerDeletionState.Merge(m, src)
}
func (m *K3SServerDeletionState) XXX_Size() int {
	return xxx_messageInfo_K3SServerDeletionState.Size(m)
}
func (m *K3SServerDeletionState) XXX_DiscardUnknown() {
	xxx_messageInfo_K3SServerDeletionState.DiscardUnknown(m)
}

var xxx_messageInfo_K3SServerDeletionState proto.InternalMessageInfo

func (m *K3SServerDeletionState) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

type K3SServerKubeconfig struct {
	Kubeconfig           string   `protobuf:"bytes,1,opt,name=Kubeconfig,proto3" json:"Kubeconfig,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *K3SServerKubeconfig) Reset()         { *m = K3SServerKubeconfig{} }
func (m *K3SServerKubeconfig) String() string { return proto.CompactTextString(m) }
func (*K3SServerKubeconfig) ProtoMessage()    {}
func (*K3SServerKubeconfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{4}
}

func (m *K3SServerKubeconfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_K3SServerKubeconfig.Unmarshal(m, b)
}
func (m *K3SServerKubeconfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_K3SServerKubeconfig.Marshal(b, m, deterministic)
}
func (m *K3SServerKubeconfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_K3SServerKubeconfig.Merge(m, src)
}
func (m *K3SServerKubeconfig) XXX_Size() int {
	return xxx_messageInfo_K3SServerKubeconfig.Size(m)
}
func (m *K3SServerKubeconfig) XXX_DiscardUnknown() {
	xxx_messageInfo_K3SServerKubeconfig.DiscardUnknown(m)
}

var xxx_messageInfo_K3SServerKubeconfig proto.InternalMessageInfo

func (m *K3SServerKubeconfig) GetKubeconfig() string {
	if m != nil {
		return m.Kubeconfig
	}
	return ""
}

type K3SServerToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *K3SServerToken) Reset()         { *m = K3SServerToken{} }
func (m *K3SServerToken) String() string { return proto.CompactTextString(m) }
func (*K3SServerToken) ProtoMessage()    {}
func (*K3SServerToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad098daeda4239f7, []int{5}
}

func (m *K3SServerToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_K3SServerToken.Unmarshal(m, b)
}
func (m *K3SServerToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_K3SServerToken.Marshal(b, m, deterministic)
}
func (m *K3SServerToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_K3SServerToken.Merge(m, src)
}
func (m *K3SServerToken) XXX_Size() int {
	return xxx_messageInfo_K3SServerToken.Size(m)
}
func (m *K3SServerToken) XXX_DiscardUnknown() {
	xxx_messageInfo_K3SServerToken.DiscardUnknown(m)
}

var xxx_messageInfo_K3SServerToken proto.InternalMessageInfo

func (m *K3SServerToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*K3SServer)(nil), "k3sd.K3SServer")
	proto.RegisterType((*K3SServerEmptyArgs)(nil), "k3sd.K3SServerEmptyArgs")
	proto.RegisterType((*K3SServerState)(nil), "k3sd.K3SServerState")
	proto.RegisterType((*K3SServerDeletionState)(nil), "k3sd.K3SServerDeletionState")
	proto.RegisterType((*K3SServerKubeconfig)(nil), "k3sd.K3SServerKubeconfig")
	proto.RegisterType((*K3SServerToken)(nil), "k3sd.K3SServerToken")
}

func init() {
	proto.RegisterFile("server.proto", fileDescriptor_ad098daeda4239f7)
}

var fileDescriptor_ad098daeda4239f7 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x5d, 0x4b, 0x32, 0x41,
	0x14, 0x5e, 0xc5, 0xcf, 0xc3, 0xeb, 0x5b, 0x9c, 0x44, 0xb6, 0x88, 0x88, 0x21, 0x22, 0xba, 0x90,
	0x50, 0x82, 0x2e, 0xba, 0x09, 0x2d, 0x09, 0xab, 0x0b, 0xc7, 0x3f, 0x30, 0xea, 0x69, 0x59, 0xd6,
	0x66, 0x96, 0xd9, 0x59, 0xa3, 0x9f, 0xd8, 0xbf, 0x8a, 0xc6, 0x75, 0x5a, 0x17, 0x8c, 0xee, 0xe6,
	0xf9, 0x3a, 0x9c, 0xf3, 0x30, 0xf0, 0x2f, 0x21, 0xbd, 0x22, 0xdd, 0x8d, 0xb5, 0x32, 0x0a, 0x2b,
	0x51, 0x3f, 0x59, 0xb0, 0x47, 0x68, 0x8e, 0xfb, 0x9c, 0x5b, 0x01, 0xcf, 0xa0, 0xf5, 0x42, 0xe6,
	0x5d, 0xe9, 0x68, 0x48, 0xab, 0x70, 0x4e, 0x7e, 0xe9, 0xb4, 0x74, 0xd1, 0x9c, 0x6c, 0x93, 0xd8,
	0x81, 0xda, 0xf4, 0x89, 0x73, 0x21, 0xfd, 0xb2, 0x95, 0x33, 0xc4, 0xda, 0x80, 0x6e, 0xd4, 0xfd,
	0x5b, 0x6c, 0x3e, 0xee, 0x74, 0x90, 0xb0, 0x4b, 0xf8, 0xef, 0x58, 0x6e, 0x84, 0x21, 0xf4, 0xa1,
	0x3e, 0x49, 0xa5, 0x0c, 0x65, 0x60, 0xe7, 0x37, 0x26, 0x1b, 0xc8, 0x7a, 0xd0, 0x71, 0xde, 0x21,
	0x2d, 0xc9, 0x84, 0x4a, 0xba, 0x8c, 0x25, 0x68, 0xb1, 0xc9, 0x64, 0x90, 0x5d, 0xc3, 0x81, 0xcb,
	0x8c, 0xd3, 0x19, 0xcd, 0x95, 0x7c, 0x0d, 0x03, 0x3c, 0x01, 0xf8, 0x41, 0xd9, 0x1d, 0x39, 0x86,
	0x9d, 0xe7, 0xd6, 0x9a, 0xaa, 0x88, 0x24, 0xb6, 0xa1, 0x6a, 0x1f, 0x99, 0x79, 0x0d, 0x7a, 0x9f,
	0x65, 0xd8, 0x77, 0xc6, 0x67, 0x21, 0x45, 0x40, 0x1a, 0xaf, 0xa0, 0xca, 0x8d, 0xd0, 0x06, 0xf7,
	0xba, 0xdf, 0x25, 0x76, 0x9d, 0xe1, 0xa8, 0x5d, 0x20, 0xec, 0xf6, 0xcc, 0xc3, 0x1b, 0xa8, 0x70,
	0xa3, 0x62, 0xf4, 0x0b, 0xba, 0xeb, 0x69, 0x67, 0x72, 0x00, 0xf5, 0xc1, 0x92, 0x84, 0x4c, 0x7f,
	0x0b, 0x1f, 0x17, 0x94, 0xad, 0xf2, 0x98, 0x87, 0x0f, 0xd0, 0x1a, 0x91, 0xc9, 0xd5, 0xb3, 0x7b,
	0xd4, 0x61, 0x41, 0xc9, 0x75, 0xe6, 0xe1, 0x2d, 0x34, 0x46, 0x64, 0xd6, 0x7d, 0xfd, 0xfd, 0x14,
	0xeb, 0x67, 0xde, 0xac, 0x66, 0x3f, 0x5e, 0xff, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xee, 0xac, 0x2b,
	0x9a, 0x88, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// K3SServerManagerClient is the client API for K3SServerManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type K3SServerManagerClient interface {
	Start(ctx context.Context, in *K3SServer, opts ...grpc.CallOption) (*K3SServerState, error)
	Stop(ctx context.Context, in *K3SServerEmptyArgs, opts ...grpc.CallOption) (*K3SServerState, error)
	Cleanup(ctx context.Context, in *K3SServerEmptyArgs, opts ...grpc.CallOption) (*K3SServerDeletionState, error)
	GetKubeconfig(ctx context.Context, in *K3SServerEmptyArgs, opts ...grpc.CallOption) (*K3SServerKubeconfig, error)
	GetToken(ctx context.Context, in *K3SServerEmptyArgs, opts ...grpc.CallOption) (*K3SServerToken, error)
}

type k3SServerManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewK3SServerManagerClient(cc grpc.ClientConnInterface) K3SServerManagerClient {
	return &k3SServerManagerClient{cc}
}

func (c *k3SServerManagerClient) Start(ctx context.Context, in *K3SServer, opts ...grpc.CallOption) (*K3SServerState, error) {
	out := new(K3SServerState)
	err := c.cc.Invoke(ctx, "/k3sd.K3SServerManager/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k3SServerManagerClient) Stop(ctx context.Context, in *K3SServerEmptyArgs, opts ...grpc.CallOption) (*K3SServerState, error) {
	out := new(K3SServerState)
	err := c.cc.Invoke(ctx, "/k3sd.K3SServerManager/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k3SServerManagerClient) Cleanup(ctx context.Context, in *K3SServerEmptyArgs, opts ...grpc.CallOption) (*K3SServerDeletionState, error) {
	out := new(K3SServerDeletionState)
	err := c.cc.Invoke(ctx, "/k3sd.K3SServerManager/Cleanup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k3SServerManagerClient) GetKubeconfig(ctx context.Context, in *K3SServerEmptyArgs, opts ...grpc.CallOption) (*K3SServerKubeconfig, error) {
	out := new(K3SServerKubeconfig)
	err := c.cc.Invoke(ctx, "/k3sd.K3SServerManager/GetKubeconfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k3SServerManagerClient) GetToken(ctx context.Context, in *K3SServerEmptyArgs, opts ...grpc.CallOption) (*K3SServerToken, error) {
	out := new(K3SServerToken)
	err := c.cc.Invoke(ctx, "/k3sd.K3SServerManager/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// K3SServerManagerServer is the server API for K3SServerManager service.
type K3SServerManagerServer interface {
	Start(context.Context, *K3SServer) (*K3SServerState, error)
	Stop(context.Context, *K3SServerEmptyArgs) (*K3SServerState, error)
	Cleanup(context.Context, *K3SServerEmptyArgs) (*K3SServerDeletionState, error)
	GetKubeconfig(context.Context, *K3SServerEmptyArgs) (*K3SServerKubeconfig, error)
	GetToken(context.Context, *K3SServerEmptyArgs) (*K3SServerToken, error)
}

// UnimplementedK3SServerManagerServer can be embedded to have forward compatible implementations.
type UnimplementedK3SServerManagerServer struct {
}

func (*UnimplementedK3SServerManagerServer) Start(ctx context.Context, req *K3SServer) (*K3SServerState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (*UnimplementedK3SServerManagerServer) Stop(ctx context.Context, req *K3SServerEmptyArgs) (*K3SServerState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (*UnimplementedK3SServerManagerServer) Cleanup(ctx context.Context, req *K3SServerEmptyArgs) (*K3SServerDeletionState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cleanup not implemented")
}
func (*UnimplementedK3SServerManagerServer) GetKubeconfig(ctx context.Context, req *K3SServerEmptyArgs) (*K3SServerKubeconfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKubeconfig not implemented")
}
func (*UnimplementedK3SServerManagerServer) GetToken(ctx context.Context, req *K3SServerEmptyArgs) (*K3SServerToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
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

func _K3SServerManager_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(K3SServerEmptyArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K3SServerManagerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/k3sd.K3SServerManager/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K3SServerManagerServer).Stop(ctx, req.(*K3SServerEmptyArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _K3SServerManager_Cleanup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(K3SServerEmptyArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K3SServerManagerServer).Cleanup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/k3sd.K3SServerManager/Cleanup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K3SServerManagerServer).Cleanup(ctx, req.(*K3SServerEmptyArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _K3SServerManager_GetKubeconfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(K3SServerEmptyArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K3SServerManagerServer).GetKubeconfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/k3sd.K3SServerManager/GetKubeconfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K3SServerManagerServer).GetKubeconfig(ctx, req.(*K3SServerEmptyArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _K3SServerManager_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(K3SServerEmptyArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K3SServerManagerServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/k3sd.K3SServerManager/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K3SServerManagerServer).GetToken(ctx, req.(*K3SServerEmptyArgs))
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
		{
			MethodName: "Stop",
			Handler:    _K3SServerManager_Stop_Handler,
		},
		{
			MethodName: "Cleanup",
			Handler:    _K3SServerManager_Cleanup_Handler,
		},
		{
			MethodName: "GetKubeconfig",
			Handler:    _K3SServerManager_GetKubeconfig_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _K3SServerManager_GetToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server.proto",
}
