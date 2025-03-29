// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: LiiklusService.proto

package liiklus

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LiiklusService_Publish_FullMethodName       = "/com.github.bsideup.liiklus.LiiklusService/Publish"
	LiiklusService_Subscribe_FullMethodName     = "/com.github.bsideup.liiklus.LiiklusService/Subscribe"
	LiiklusService_Receive_FullMethodName       = "/com.github.bsideup.liiklus.LiiklusService/Receive"
	LiiklusService_Ack_FullMethodName           = "/com.github.bsideup.liiklus.LiiklusService/Ack"
	LiiklusService_GetOffsets_FullMethodName    = "/com.github.bsideup.liiklus.LiiklusService/GetOffsets"
	LiiklusService_GetEndOffsets_FullMethodName = "/com.github.bsideup.liiklus.LiiklusService/GetEndOffsets"
)

// LiiklusServiceClient is the client API for LiiklusService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LiiklusServiceClient interface {
	Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishReply, error)
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[SubscribeReply], error)
	Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ReceiveReply], error)
	Ack(ctx context.Context, in *AckRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetOffsets(ctx context.Context, in *GetOffsetsRequest, opts ...grpc.CallOption) (*GetOffsetsReply, error)
	GetEndOffsets(ctx context.Context, in *GetEndOffsetsRequest, opts ...grpc.CallOption) (*GetEndOffsetsReply, error)
}

type liiklusServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLiiklusServiceClient(cc grpc.ClientConnInterface) LiiklusServiceClient {
	return &liiklusServiceClient{cc}
}

func (c *liiklusServiceClient) Publish(ctx context.Context, in *PublishRequest, opts ...grpc.CallOption) (*PublishReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PublishReply)
	err := c.cc.Invoke(ctx, LiiklusService_Publish_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liiklusServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[SubscribeReply], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &LiiklusService_ServiceDesc.Streams[0], LiiklusService_Subscribe_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SubscribeRequest, SubscribeReply]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LiiklusService_SubscribeClient = grpc.ServerStreamingClient[SubscribeReply]

func (c *liiklusServiceClient) Receive(ctx context.Context, in *ReceiveRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ReceiveReply], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &LiiklusService_ServiceDesc.Streams[1], LiiklusService_Receive_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ReceiveRequest, ReceiveReply]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LiiklusService_ReceiveClient = grpc.ServerStreamingClient[ReceiveReply]

func (c *liiklusServiceClient) Ack(ctx context.Context, in *AckRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, LiiklusService_Ack_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liiklusServiceClient) GetOffsets(ctx context.Context, in *GetOffsetsRequest, opts ...grpc.CallOption) (*GetOffsetsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOffsetsReply)
	err := c.cc.Invoke(ctx, LiiklusService_GetOffsets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *liiklusServiceClient) GetEndOffsets(ctx context.Context, in *GetEndOffsetsRequest, opts ...grpc.CallOption) (*GetEndOffsetsReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetEndOffsetsReply)
	err := c.cc.Invoke(ctx, LiiklusService_GetEndOffsets_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LiiklusServiceServer is the server API for LiiklusService service.
// All implementations must embed UnimplementedLiiklusServiceServer
// for forward compatibility.
type LiiklusServiceServer interface {
	Publish(context.Context, *PublishRequest) (*PublishReply, error)
	Subscribe(*SubscribeRequest, grpc.ServerStreamingServer[SubscribeReply]) error
	Receive(*ReceiveRequest, grpc.ServerStreamingServer[ReceiveReply]) error
	Ack(context.Context, *AckRequest) (*emptypb.Empty, error)
	GetOffsets(context.Context, *GetOffsetsRequest) (*GetOffsetsReply, error)
	GetEndOffsets(context.Context, *GetEndOffsetsRequest) (*GetEndOffsetsReply, error)
	mustEmbedUnimplementedLiiklusServiceServer()
}

// UnimplementedLiiklusServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLiiklusServiceServer struct{}

func (UnimplementedLiiklusServiceServer) Publish(context.Context, *PublishRequest) (*PublishReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Publish not implemented")
}
func (UnimplementedLiiklusServiceServer) Subscribe(*SubscribeRequest, grpc.ServerStreamingServer[SubscribeReply]) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedLiiklusServiceServer) Receive(*ReceiveRequest, grpc.ServerStreamingServer[ReceiveReply]) error {
	return status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (UnimplementedLiiklusServiceServer) Ack(context.Context, *AckRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ack not implemented")
}
func (UnimplementedLiiklusServiceServer) GetOffsets(context.Context, *GetOffsetsRequest) (*GetOffsetsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOffsets not implemented")
}
func (UnimplementedLiiklusServiceServer) GetEndOffsets(context.Context, *GetEndOffsetsRequest) (*GetEndOffsetsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEndOffsets not implemented")
}
func (UnimplementedLiiklusServiceServer) mustEmbedUnimplementedLiiklusServiceServer() {}
func (UnimplementedLiiklusServiceServer) testEmbeddedByValue()                        {}

// UnsafeLiiklusServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LiiklusServiceServer will
// result in compilation errors.
type UnsafeLiiklusServiceServer interface {
	mustEmbedUnimplementedLiiklusServiceServer()
}

func RegisterLiiklusServiceServer(s grpc.ServiceRegistrar, srv LiiklusServiceServer) {
	// If the following call pancis, it indicates UnimplementedLiiklusServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LiiklusService_ServiceDesc, srv)
}

func _LiiklusService_Publish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiiklusServiceServer).Publish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiiklusService_Publish_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiiklusServiceServer).Publish(ctx, req.(*PublishRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiiklusService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LiiklusServiceServer).Subscribe(m, &grpc.GenericServerStream[SubscribeRequest, SubscribeReply]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LiiklusService_SubscribeServer = grpc.ServerStreamingServer[SubscribeReply]

func _LiiklusService_Receive_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReceiveRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LiiklusServiceServer).Receive(m, &grpc.GenericServerStream[ReceiveRequest, ReceiveReply]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LiiklusService_ReceiveServer = grpc.ServerStreamingServer[ReceiveReply]

func _LiiklusService_Ack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiiklusServiceServer).Ack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiiklusService_Ack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiiklusServiceServer).Ack(ctx, req.(*AckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiiklusService_GetOffsets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOffsetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiiklusServiceServer).GetOffsets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiiklusService_GetOffsets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiiklusServiceServer).GetOffsets(ctx, req.(*GetOffsetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LiiklusService_GetEndOffsets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEndOffsetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LiiklusServiceServer).GetEndOffsets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LiiklusService_GetEndOffsets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LiiklusServiceServer).GetEndOffsets(ctx, req.(*GetEndOffsetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LiiklusService_ServiceDesc is the grpc.ServiceDesc for LiiklusService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LiiklusService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.github.bsideup.liiklus.LiiklusService",
	HandlerType: (*LiiklusServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Publish",
			Handler:    _LiiklusService_Publish_Handler,
		},
		{
			MethodName: "Ack",
			Handler:    _LiiklusService_Ack_Handler,
		},
		{
			MethodName: "GetOffsets",
			Handler:    _LiiklusService_GetOffsets_Handler,
		},
		{
			MethodName: "GetEndOffsets",
			Handler:    _LiiklusService_GetEndOffsets_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _LiiklusService_Subscribe_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Receive",
			Handler:       _LiiklusService_Receive_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "LiiklusService.proto",
}
