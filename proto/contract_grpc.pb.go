// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: contract.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (ChatService_ConnectClient, error)
	JoinGroupChat(ctx context.Context, in *JoinGroupChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	LeftGroupChat(ctx context.Context, in *LeftGroupChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateGroupChat(ctx context.Context, in *CreateGroupChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// this method assumes that the username and channel name do not conflict, the
	// username takes precedence.
	SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListChannels(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListChannelsResponse, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Connect(ctx context.Context, in *ConnectRequest, opts ...grpc.CallOption) (ChatService_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], "/chat.ChatService/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_ConnectClient interface {
	Recv() (*ConnectResponse, error)
	grpc.ClientStream
}

type chatServiceConnectClient struct {
	grpc.ClientStream
}

func (x *chatServiceConnectClient) Recv() (*ConnectResponse, error) {
	m := new(ConnectResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) JoinGroupChat(ctx context.Context, in *JoinGroupChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chat.ChatService/JoinGroupChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) LeftGroupChat(ctx context.Context, in *LeftGroupChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chat.ChatService/LeftGroupChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) CreateGroupChat(ctx context.Context, in *CreateGroupChatRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chat.ChatService/CreateGroupChat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) SendMessage(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/chat.ChatService/SendMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) ListChannels(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListChannelsResponse, error) {
	out := new(ListChannelsResponse)
	err := c.cc.Invoke(ctx, "/chat.ChatService/ListChannels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	Connect(*ConnectRequest, ChatService_ConnectServer) error
	JoinGroupChat(context.Context, *JoinGroupChatRequest) (*emptypb.Empty, error)
	LeftGroupChat(context.Context, *LeftGroupChatRequest) (*emptypb.Empty, error)
	CreateGroupChat(context.Context, *CreateGroupChatRequest) (*emptypb.Empty, error)
	// this method assumes that the username and channel name do not conflict, the
	// username takes precedence.
	SendMessage(context.Context, *SendMessageRequest) (*emptypb.Empty, error)
	ListChannels(context.Context, *emptypb.Empty) (*ListChannelsResponse, error)
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) Connect(*ConnectRequest, ChatService_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedChatServiceServer) JoinGroupChat(context.Context, *JoinGroupChatRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinGroupChat not implemented")
}
func (UnimplementedChatServiceServer) LeftGroupChat(context.Context, *LeftGroupChatRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeftGroupChat not implemented")
}
func (UnimplementedChatServiceServer) CreateGroupChat(context.Context, *CreateGroupChatRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroupChat not implemented")
}
func (UnimplementedChatServiceServer) SendMessage(context.Context, *SendMessageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessage not implemented")
}
func (UnimplementedChatServiceServer) ListChannels(context.Context, *emptypb.Empty) (*ListChannelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListChannels not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConnectRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).Connect(m, &chatServiceConnectServer{stream})
}

type ChatService_ConnectServer interface {
	Send(*ConnectResponse) error
	grpc.ServerStream
}

type chatServiceConnectServer struct {
	grpc.ServerStream
}

func (x *chatServiceConnectServer) Send(m *ConnectResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatService_JoinGroupChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinGroupChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).JoinGroupChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/JoinGroupChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).JoinGroupChat(ctx, req.(*JoinGroupChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_LeftGroupChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeftGroupChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).LeftGroupChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/LeftGroupChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).LeftGroupChat(ctx, req.(*LeftGroupChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_CreateGroupChat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).CreateGroupChat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/CreateGroupChat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).CreateGroupChat(ctx, req.(*CreateGroupChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_SendMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).SendMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/SendMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).SendMessage(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_ListChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).ListChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/ListChannels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).ListChannels(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinGroupChat",
			Handler:    _ChatService_JoinGroupChat_Handler,
		},
		{
			MethodName: "LeftGroupChat",
			Handler:    _ChatService_LeftGroupChat_Handler,
		},
		{
			MethodName: "CreateGroupChat",
			Handler:    _ChatService_CreateGroupChat_Handler,
		},
		{
			MethodName: "SendMessage",
			Handler:    _ChatService_SendMessage_Handler,
		},
		{
			MethodName: "ListChannels",
			Handler:    _ChatService_ListChannels_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _ChatService_Connect_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "contract.proto",
}
