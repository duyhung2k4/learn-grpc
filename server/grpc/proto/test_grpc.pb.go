// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/test.proto

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

// DataServiceClient is the client API for DataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataServiceClient interface {
	SendTextOneToOne(ctx context.Context, in *DataReq, opts ...grpc.CallOption) (*DataRes, error)
	SendTextOneToMany(ctx context.Context, in *DataReq, opts ...grpc.CallOption) (DataService_SendTextOneToManyClient, error)
	SendTextManyToOne(ctx context.Context, opts ...grpc.CallOption) (DataService_SendTextManyToOneClient, error)
	SendTextManyToMany(ctx context.Context, opts ...grpc.CallOption) (DataService_SendTextManyToManyClient, error)
}

type dataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDataServiceClient(cc grpc.ClientConnInterface) DataServiceClient {
	return &dataServiceClient{cc}
}

func (c *dataServiceClient) SendTextOneToOne(ctx context.Context, in *DataReq, opts ...grpc.CallOption) (*DataRes, error) {
	out := new(DataRes)
	err := c.cc.Invoke(ctx, "/proto.DataService/SendTextOneToOne", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dataServiceClient) SendTextOneToMany(ctx context.Context, in *DataReq, opts ...grpc.CallOption) (DataService_SendTextOneToManyClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataService_ServiceDesc.Streams[0], "/proto.DataService/SendTextOneToMany", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataServiceSendTextOneToManyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DataService_SendTextOneToManyClient interface {
	Recv() (*DataRes, error)
	grpc.ClientStream
}

type dataServiceSendTextOneToManyClient struct {
	grpc.ClientStream
}

func (x *dataServiceSendTextOneToManyClient) Recv() (*DataRes, error) {
	m := new(DataRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataServiceClient) SendTextManyToOne(ctx context.Context, opts ...grpc.CallOption) (DataService_SendTextManyToOneClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataService_ServiceDesc.Streams[1], "/proto.DataService/SendTextManyToOne", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataServiceSendTextManyToOneClient{stream}
	return x, nil
}

type DataService_SendTextManyToOneClient interface {
	Send(*DataReq) error
	CloseAndRecv() (*DataRes, error)
	grpc.ClientStream
}

type dataServiceSendTextManyToOneClient struct {
	grpc.ClientStream
}

func (x *dataServiceSendTextManyToOneClient) Send(m *DataReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataServiceSendTextManyToOneClient) CloseAndRecv() (*DataRes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(DataRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dataServiceClient) SendTextManyToMany(ctx context.Context, opts ...grpc.CallOption) (DataService_SendTextManyToManyClient, error) {
	stream, err := c.cc.NewStream(ctx, &DataService_ServiceDesc.Streams[2], "/proto.DataService/SendTextManyToMany", opts...)
	if err != nil {
		return nil, err
	}
	x := &dataServiceSendTextManyToManyClient{stream}
	return x, nil
}

type DataService_SendTextManyToManyClient interface {
	Send(*DataReq) error
	Recv() (*DataRes, error)
	grpc.ClientStream
}

type dataServiceSendTextManyToManyClient struct {
	grpc.ClientStream
}

func (x *dataServiceSendTextManyToManyClient) Send(m *DataReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dataServiceSendTextManyToManyClient) Recv() (*DataRes, error) {
	m := new(DataRes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataServiceServer is the server API for DataService service.
// All implementations must embed UnimplementedDataServiceServer
// for forward compatibility
type DataServiceServer interface {
	SendTextOneToOne(context.Context, *DataReq) (*DataRes, error)
	SendTextOneToMany(*DataReq, DataService_SendTextOneToManyServer) error
	SendTextManyToOne(DataService_SendTextManyToOneServer) error
	SendTextManyToMany(DataService_SendTextManyToManyServer) error
	mustEmbedUnimplementedDataServiceServer()
}

// UnimplementedDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDataServiceServer struct {
}

func (UnimplementedDataServiceServer) SendTextOneToOne(context.Context, *DataReq) (*DataRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTextOneToOne not implemented")
}
func (UnimplementedDataServiceServer) SendTextOneToMany(*DataReq, DataService_SendTextOneToManyServer) error {
	return status.Errorf(codes.Unimplemented, "method SendTextOneToMany not implemented")
}
func (UnimplementedDataServiceServer) SendTextManyToOne(DataService_SendTextManyToOneServer) error {
	return status.Errorf(codes.Unimplemented, "method SendTextManyToOne not implemented")
}
func (UnimplementedDataServiceServer) SendTextManyToMany(DataService_SendTextManyToManyServer) error {
	return status.Errorf(codes.Unimplemented, "method SendTextManyToMany not implemented")
}
func (UnimplementedDataServiceServer) mustEmbedUnimplementedDataServiceServer() {}

// UnsafeDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataServiceServer will
// result in compilation errors.
type UnsafeDataServiceServer interface {
	mustEmbedUnimplementedDataServiceServer()
}

func RegisterDataServiceServer(s grpc.ServiceRegistrar, srv DataServiceServer) {
	s.RegisterService(&DataService_ServiceDesc, srv)
}

func _DataService_SendTextOneToOne_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServiceServer).SendTextOneToOne(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.DataService/SendTextOneToOne",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServiceServer).SendTextOneToOne(ctx, req.(*DataReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _DataService_SendTextOneToMany_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DataReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DataServiceServer).SendTextOneToMany(m, &dataServiceSendTextOneToManyServer{stream})
}

type DataService_SendTextOneToManyServer interface {
	Send(*DataRes) error
	grpc.ServerStream
}

type dataServiceSendTextOneToManyServer struct {
	grpc.ServerStream
}

func (x *dataServiceSendTextOneToManyServer) Send(m *DataRes) error {
	return x.ServerStream.SendMsg(m)
}

func _DataService_SendTextManyToOne_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataServiceServer).SendTextManyToOne(&dataServiceSendTextManyToOneServer{stream})
}

type DataService_SendTextManyToOneServer interface {
	SendAndClose(*DataRes) error
	Recv() (*DataReq, error)
	grpc.ServerStream
}

type dataServiceSendTextManyToOneServer struct {
	grpc.ServerStream
}

func (x *dataServiceSendTextManyToOneServer) SendAndClose(m *DataRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataServiceSendTextManyToOneServer) Recv() (*DataReq, error) {
	m := new(DataReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DataService_SendTextManyToMany_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DataServiceServer).SendTextManyToMany(&dataServiceSendTextManyToManyServer{stream})
}

type DataService_SendTextManyToManyServer interface {
	Send(*DataRes) error
	Recv() (*DataReq, error)
	grpc.ServerStream
}

type dataServiceSendTextManyToManyServer struct {
	grpc.ServerStream
}

func (x *dataServiceSendTextManyToManyServer) Send(m *DataRes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dataServiceSendTextManyToManyServer) Recv() (*DataReq, error) {
	m := new(DataReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DataService_ServiceDesc is the grpc.ServiceDesc for DataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DataService",
	HandlerType: (*DataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendTextOneToOne",
			Handler:    _DataService_SendTextOneToOne_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendTextOneToMany",
			Handler:       _DataService_SendTextOneToMany_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SendTextManyToOne",
			Handler:       _DataService_SendTextManyToOne_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "SendTextManyToMany",
			Handler:       _DataService_SendTextManyToMany_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/test.proto",
}
