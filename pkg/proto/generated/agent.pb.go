// Code generated by protoc-gen-go. DO NOT EDIT.
// source: agent.proto

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

type K3SAgent struct {
	NetworkDevice        string   `protobuf:"bytes,1,opt,name=NetworkDevice,proto3" json:"NetworkDevice,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	ServerURL            string   `protobuf:"bytes,3,opt,name=ServerURL,proto3" json:"ServerURL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *K3SAgent) Reset()         { *m = K3SAgent{} }
func (m *K3SAgent) String() string { return proto.CompactTextString(m) }
func (*K3SAgent) ProtoMessage()    {}
func (*K3SAgent) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{0}
}

func (m *K3SAgent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_K3SAgent.Unmarshal(m, b)
}
func (m *K3SAgent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_K3SAgent.Marshal(b, m, deterministic)
}
func (m *K3SAgent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_K3SAgent.Merge(m, src)
}
func (m *K3SAgent) XXX_Size() int {
	return xxx_messageInfo_K3SAgent.Size(m)
}
func (m *K3SAgent) XXX_DiscardUnknown() {
	xxx_messageInfo_K3SAgent.DiscardUnknown(m)
}

var xxx_messageInfo_K3SAgent proto.InternalMessageInfo

func (m *K3SAgent) GetNetworkDevice() string {
	if m != nil {
		return m.NetworkDevice
	}
	return ""
}

func (m *K3SAgent) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *K3SAgent) GetServerURL() string {
	if m != nil {
		return m.ServerURL
	}
	return ""
}

type K3SAgentEmptyArgs struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *K3SAgentEmptyArgs) Reset()         { *m = K3SAgentEmptyArgs{} }
func (m *K3SAgentEmptyArgs) String() string { return proto.CompactTextString(m) }
func (*K3SAgentEmptyArgs) ProtoMessage()    {}
func (*K3SAgentEmptyArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{1}
}

func (m *K3SAgentEmptyArgs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_K3SAgentEmptyArgs.Unmarshal(m, b)
}
func (m *K3SAgentEmptyArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_K3SAgentEmptyArgs.Marshal(b, m, deterministic)
}
func (m *K3SAgentEmptyArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_K3SAgentEmptyArgs.Merge(m, src)
}
func (m *K3SAgentEmptyArgs) XXX_Size() int {
	return xxx_messageInfo_K3SAgentEmptyArgs.Size(m)
}
func (m *K3SAgentEmptyArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_K3SAgentEmptyArgs.DiscardUnknown(m)
}

var xxx_messageInfo_K3SAgentEmptyArgs proto.InternalMessageInfo

type K3SAgentState struct {
	Running              bool     `protobuf:"varint,1,opt,name=Running,proto3" json:"Running,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *K3SAgentState) Reset()         { *m = K3SAgentState{} }
func (m *K3SAgentState) String() string { return proto.CompactTextString(m) }
func (*K3SAgentState) ProtoMessage()    {}
func (*K3SAgentState) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{2}
}

func (m *K3SAgentState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_K3SAgentState.Unmarshal(m, b)
}
func (m *K3SAgentState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_K3SAgentState.Marshal(b, m, deterministic)
}
func (m *K3SAgentState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_K3SAgentState.Merge(m, src)
}
func (m *K3SAgentState) XXX_Size() int {
	return xxx_messageInfo_K3SAgentState.Size(m)
}
func (m *K3SAgentState) XXX_DiscardUnknown() {
	xxx_messageInfo_K3SAgentState.DiscardUnknown(m)
}

var xxx_messageInfo_K3SAgentState proto.InternalMessageInfo

func (m *K3SAgentState) GetRunning() bool {
	if m != nil {
		return m.Running
	}
	return false
}

type K3SAgentDeletionState struct {
	Deleted              bool     `protobuf:"varint,1,opt,name=Deleted,proto3" json:"Deleted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *K3SAgentDeletionState) Reset()         { *m = K3SAgentDeletionState{} }
func (m *K3SAgentDeletionState) String() string { return proto.CompactTextString(m) }
func (*K3SAgentDeletionState) ProtoMessage()    {}
func (*K3SAgentDeletionState) Descriptor() ([]byte, []int) {
	return fileDescriptor_56ede974c0020f77, []int{3}
}

func (m *K3SAgentDeletionState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_K3SAgentDeletionState.Unmarshal(m, b)
}
func (m *K3SAgentDeletionState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_K3SAgentDeletionState.Marshal(b, m, deterministic)
}
func (m *K3SAgentDeletionState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_K3SAgentDeletionState.Merge(m, src)
}
func (m *K3SAgentDeletionState) XXX_Size() int {
	return xxx_messageInfo_K3SAgentDeletionState.Size(m)
}
func (m *K3SAgentDeletionState) XXX_DiscardUnknown() {
	xxx_messageInfo_K3SAgentDeletionState.DiscardUnknown(m)
}

var xxx_messageInfo_K3SAgentDeletionState proto.InternalMessageInfo

func (m *K3SAgentDeletionState) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func init() {
	proto.RegisterType((*K3SAgent)(nil), "k3sd.K3SAgent")
	proto.RegisterType((*K3SAgentEmptyArgs)(nil), "k3sd.K3SAgentEmptyArgs")
	proto.RegisterType((*K3SAgentState)(nil), "k3sd.K3SAgentState")
	proto.RegisterType((*K3SAgentDeletionState)(nil), "k3sd.K3SAgentDeletionState")
}

func init() {
	proto.RegisterFile("agent.proto", fileDescriptor_56ede974c0020f77)
}

var fileDescriptor_56ede974c0020f77 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x1b, 0x6d, 0x6d, 0x7b, 0xa5, 0x8a, 0xb7, 0x8a, 0x41, 0x5d, 0xc8, 0xe0, 0x42, 0x37,
	0x01, 0x0d, 0xb8, 0x0f, 0xd6, 0x95, 0x3f, 0x8b, 0x8c, 0x3e, 0xc0, 0x68, 0x2e, 0x43, 0x48, 0x9d,
	0x09, 0xd3, 0xdb, 0x8a, 0xef, 0xe6, 0xc3, 0x49, 0x47, 0x27, 0x32, 0x0b, 0x5d, 0x9e, 0xf3, 0xe5,
	0xcb, 0xe5, 0x30, 0xb0, 0xad, 0x34, 0x19, 0xce, 0x5a, 0x67, 0xd9, 0x62, 0xbf, 0xc9, 0x17, 0x95,
	0xa8, 0x60, 0x74, 0x97, 0xcb, 0x62, 0xdd, 0xe3, 0x19, 0x4c, 0x1e, 0x89, 0xdf, 0xad, 0x6b, 0x66,
	0xb4, 0xaa, 0x5f, 0x29, 0x4d, 0x4e, 0x93, 0xf3, 0x71, 0x19, 0x97, 0xb8, 0x0f, 0x83, 0x27, 0xdb,
	0x90, 0x49, 0x37, 0x3c, 0xfd, 0x0e, 0x78, 0x02, 0x63, 0x49, 0x6e, 0x45, 0xee, 0xb9, 0xbc, 0x4f,
	0x37, 0x3d, 0xf9, 0x2d, 0xc4, 0x14, 0xf6, 0xc2, 0x95, 0xdb, 0xb7, 0x96, 0x3f, 0x0a, 0xa7, 0x17,
	0xe2, 0x02, 0x26, 0xa1, 0x94, 0xac, 0x98, 0x30, 0x85, 0x61, 0xb9, 0x34, 0xa6, 0x36, 0xda, 0x5f,
	0x1e, 0x95, 0x21, 0x8a, 0x4b, 0x38, 0x08, 0x9f, 0xce, 0x68, 0x4e, 0x5c, 0x5b, 0xd3, 0x29, 0xbe,
	0xa0, 0x2a, 0x28, 0x3f, 0xf1, 0xea, 0x33, 0x81, 0xdd, 0xe0, 0x3c, 0x28, 0xa3, 0x34, 0x39, 0xcc,
	0x60, 0x20, 0x59, 0x39, 0xc6, 0x9d, 0x6c, 0x3d, 0x3e, 0x0b, 0xfc, 0x68, 0x1a, 0x67, 0xff, 0x6f,
	0xd1, 0xc3, 0x6b, 0xe8, 0x4b, 0xb6, 0x2d, 0x1e, 0xc6, 0xb8, 0x9b, 0xf0, 0x97, 0x57, 0xc0, 0xf0,
	0x66, 0x4e, 0xca, 0x2c, 0xff, 0x51, 0x8f, 0x63, 0x10, 0xcd, 0x12, 0xbd, 0x97, 0x2d, 0xff, 0x48,
	0xf9, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xf6, 0x1c, 0x6c, 0xb3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// K3SAgentManagerClient is the client API for K3SAgentManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type K3SAgentManagerClient interface {
	Start(ctx context.Context, in *K3SAgent, opts ...grpc.CallOption) (*K3SAgentState, error)
	Stop(ctx context.Context, in *K3SAgentEmptyArgs, opts ...grpc.CallOption) (*K3SAgentState, error)
	Cleanup(ctx context.Context, in *K3SAgentEmptyArgs, opts ...grpc.CallOption) (*K3SAgentDeletionState, error)
}

type k3SAgentManagerClient struct {
	cc grpc.ClientConnInterface
}

func NewK3SAgentManagerClient(cc grpc.ClientConnInterface) K3SAgentManagerClient {
	return &k3SAgentManagerClient{cc}
}

func (c *k3SAgentManagerClient) Start(ctx context.Context, in *K3SAgent, opts ...grpc.CallOption) (*K3SAgentState, error) {
	out := new(K3SAgentState)
	err := c.cc.Invoke(ctx, "/k3sd.K3SAgentManager/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k3SAgentManagerClient) Stop(ctx context.Context, in *K3SAgentEmptyArgs, opts ...grpc.CallOption) (*K3SAgentState, error) {
	out := new(K3SAgentState)
	err := c.cc.Invoke(ctx, "/k3sd.K3SAgentManager/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *k3SAgentManagerClient) Cleanup(ctx context.Context, in *K3SAgentEmptyArgs, opts ...grpc.CallOption) (*K3SAgentDeletionState, error) {
	out := new(K3SAgentDeletionState)
	err := c.cc.Invoke(ctx, "/k3sd.K3SAgentManager/Cleanup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// K3SAgentManagerServer is the server API for K3SAgentManager service.
type K3SAgentManagerServer interface {
	Start(context.Context, *K3SAgent) (*K3SAgentState, error)
	Stop(context.Context, *K3SAgentEmptyArgs) (*K3SAgentState, error)
	Cleanup(context.Context, *K3SAgentEmptyArgs) (*K3SAgentDeletionState, error)
}

// UnimplementedK3SAgentManagerServer can be embedded to have forward compatible implementations.
type UnimplementedK3SAgentManagerServer struct {
}

func (*UnimplementedK3SAgentManagerServer) Start(ctx context.Context, req *K3SAgent) (*K3SAgentState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (*UnimplementedK3SAgentManagerServer) Stop(ctx context.Context, req *K3SAgentEmptyArgs) (*K3SAgentState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (*UnimplementedK3SAgentManagerServer) Cleanup(ctx context.Context, req *K3SAgentEmptyArgs) (*K3SAgentDeletionState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cleanup not implemented")
}

func RegisterK3SAgentManagerServer(s *grpc.Server, srv K3SAgentManagerServer) {
	s.RegisterService(&_K3SAgentManager_serviceDesc, srv)
}

func _K3SAgentManager_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(K3SAgent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K3SAgentManagerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/k3sd.K3SAgentManager/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K3SAgentManagerServer).Start(ctx, req.(*K3SAgent))
	}
	return interceptor(ctx, in, info, handler)
}

func _K3SAgentManager_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(K3SAgentEmptyArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K3SAgentManagerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/k3sd.K3SAgentManager/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K3SAgentManagerServer).Stop(ctx, req.(*K3SAgentEmptyArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _K3SAgentManager_Cleanup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(K3SAgentEmptyArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(K3SAgentManagerServer).Cleanup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/k3sd.K3SAgentManager/Cleanup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(K3SAgentManagerServer).Cleanup(ctx, req.(*K3SAgentEmptyArgs))
	}
	return interceptor(ctx, in, info, handler)
}

var _K3SAgentManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "k3sd.K3SAgentManager",
	HandlerType: (*K3SAgentManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _K3SAgentManager_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _K3SAgentManager_Stop_Handler,
		},
		{
			MethodName: "Cleanup",
			Handler:    _K3SAgentManager_Cleanup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "agent.proto",
}
