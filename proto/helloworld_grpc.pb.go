// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: helloworld.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GreeterClient interface {
	// Unary RPC.
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	Plus(ctx context.Context, in *PlusRequest, opts ...grpc.CallOption) (*PlusReply, error)
	SayHelloAfterDelay(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
	// Server side streaming.
	SayHelloServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_SayHelloServerStreamClient, error)
	// Client side streaming.
	SayHelloClientStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloClientStreamClient, error)
	// Bidirectional streaming.
	SayHelloBidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloBidirectionalStreamClient, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) Plus(ctx context.Context, in *PlusRequest, opts ...grpc.CallOption) (*PlusReply, error) {
	out := new(PlusReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/Plus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayHelloAfterDelay(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHelloAfterDelay", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterClient) SayHelloServerStream(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (Greeter_SayHelloServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[0], "/helloworld.Greeter/SayHelloServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHelloServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Greeter_SayHelloServerStreamClient interface {
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterSayHelloServerStreamClient struct {
	grpc.ClientStream
}

func (x *greeterSayHelloServerStreamClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) SayHelloClientStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[1], "/helloworld.Greeter/SayHelloClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHelloClientStreamClient{stream}
	return x, nil
}

type Greeter_SayHelloClientStreamClient interface {
	Send(*HelloRequest) error
	CloseAndRecv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterSayHelloClientStreamClient struct {
	grpc.ClientStream
}

func (x *greeterSayHelloClientStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterSayHelloClientStreamClient) CloseAndRecv() (*HelloReply, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *greeterClient) SayHelloBidirectionalStream(ctx context.Context, opts ...grpc.CallOption) (Greeter_SayHelloBidirectionalStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Greeter_ServiceDesc.Streams[2], "/helloworld.Greeter/SayHelloBidirectionalStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &greeterSayHelloBidirectionalStreamClient{stream}
	return x, nil
}

type Greeter_SayHelloBidirectionalStreamClient interface {
	Send(*HelloRequest) error
	Recv() (*HelloReply, error)
	grpc.ClientStream
}

type greeterSayHelloBidirectionalStreamClient struct {
	grpc.ClientStream
}

func (x *greeterSayHelloBidirectionalStreamClient) Send(m *HelloRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *greeterSayHelloBidirectionalStreamClient) Recv() (*HelloReply, error) {
	m := new(HelloReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GreeterServer is the server API for Greeter service.
// All implementations must embed UnimplementedGreeterServer
// for forward compatibility
type GreeterServer interface {
	// Unary RPC.
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	Plus(context.Context, *PlusRequest) (*PlusReply, error)
	SayHelloAfterDelay(context.Context, *HelloRequest) (*HelloReply, error)
	// Server side streaming.
	SayHelloServerStream(*HelloRequest, Greeter_SayHelloServerStreamServer) error
	// Client side streaming.
	SayHelloClientStream(Greeter_SayHelloClientStreamServer) error
	// Bidirectional streaming.
	SayHelloBidirectionalStream(Greeter_SayHelloBidirectionalStreamServer) error
	mustEmbedUnimplementedGreeterServer()
}

// UnimplementedGreeterServer must be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (UnimplementedGreeterServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedGreeterServer) Plus(context.Context, *PlusRequest) (*PlusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Plus not implemented")
}
func (UnimplementedGreeterServer) SayHelloAfterDelay(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHelloAfterDelay not implemented")
}
func (UnimplementedGreeterServer) SayHelloServerStream(*HelloRequest, Greeter_SayHelloServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloServerStream not implemented")
}
func (UnimplementedGreeterServer) SayHelloClientStream(Greeter_SayHelloClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloClientStream not implemented")
}
func (UnimplementedGreeterServer) SayHelloBidirectionalStream(Greeter_SayHelloBidirectionalStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloBidirectionalStream not implemented")
}
func (UnimplementedGreeterServer) mustEmbedUnimplementedGreeterServer() {}

// UnsafeGreeterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GreeterServer will
// result in compilation errors.
type UnsafeGreeterServer interface {
	mustEmbedUnimplementedGreeterServer()
}

func RegisterGreeterServer(s grpc.ServiceRegistrar, srv GreeterServer) {
	s.RegisterService(&Greeter_ServiceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_Plus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Plus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/Plus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Plus(ctx, req.(*PlusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHelloAfterDelay_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHelloAfterDelay(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHelloAfterDelay",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHelloAfterDelay(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Greeter_SayHelloServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GreeterServer).SayHelloServerStream(m, &greeterSayHelloServerStreamServer{stream})
}

type Greeter_SayHelloServerStreamServer interface {
	Send(*HelloReply) error
	grpc.ServerStream
}

type greeterSayHelloServerStreamServer struct {
	grpc.ServerStream
}

func (x *greeterSayHelloServerStreamServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func _Greeter_SayHelloClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).SayHelloClientStream(&greeterSayHelloClientStreamServer{stream})
}

type Greeter_SayHelloClientStreamServer interface {
	SendAndClose(*HelloReply) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterSayHelloClientStreamServer struct {
	grpc.ServerStream
}

func (x *greeterSayHelloClientStreamServer) SendAndClose(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterSayHelloClientStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Greeter_SayHelloBidirectionalStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GreeterServer).SayHelloBidirectionalStream(&greeterSayHelloBidirectionalStreamServer{stream})
}

type Greeter_SayHelloBidirectionalStreamServer interface {
	Send(*HelloReply) error
	Recv() (*HelloRequest, error)
	grpc.ServerStream
}

type greeterSayHelloBidirectionalStreamServer struct {
	grpc.ServerStream
}

func (x *greeterSayHelloBidirectionalStreamServer) Send(m *HelloReply) error {
	return x.ServerStream.SendMsg(m)
}

func (x *greeterSayHelloBidirectionalStreamServer) Recv() (*HelloRequest, error) {
	m := new(HelloRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Greeter_ServiceDesc is the grpc.ServiceDesc for Greeter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Greeter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
		{
			MethodName: "Plus",
			Handler:    _Greeter_Plus_Handler,
		},
		{
			MethodName: "SayHelloAfterDelay",
			Handler:    _Greeter_SayHelloAfterDelay_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SayHelloServerStream",
			Handler:       _Greeter_SayHelloServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayHelloClientStream",
			Handler:       _Greeter_SayHelloClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SayHelloBidirectionalStream",
			Handler:       _Greeter_SayHelloBidirectionalStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "helloworld.proto",
}
