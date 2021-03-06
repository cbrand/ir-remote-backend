// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protocol

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// RemoteServiceClient is the client API for RemoteService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemoteServiceClient interface {
	GetRemote(ctx context.Context, in *RemoteQueryParams, opts ...grpc.CallOption) (*Remote, error)
	GetTheatersOf(ctx context.Context, in *RemoteQueryParams, opts ...grpc.CallOption) (*TheatersResponse, error)
	GetIscpStatusOf(ctx context.Context, in *RemoteQueryParams, opts ...grpc.CallOption) (*IscpStatusResponse, error)
	SetTheater(ctx context.Context, in *SetTheaterQueryParams, opts ...grpc.CallOption) (*TheatersResponse, error)
	RemoveTheater(ctx context.Context, in *RemoveTheaterQueryParams, opts ...grpc.CallOption) (*TheatersResponse, error)
	GetStatus(ctx context.Context, in *RemoteQueryParams, opts ...grpc.CallOption) (*RemoteStatus, error)
	PlayScene(ctx context.Context, in *PlaySceneParams, opts ...grpc.CallOption) (*PlayResponse, error)
	PlayCommand(ctx context.Context, in *PlayCommandParams, opts ...grpc.CallOption) (*PlayResponse, error)
}

type remoteServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRemoteServiceClient(cc grpc.ClientConnInterface) RemoteServiceClient {
	return &remoteServiceClient{cc}
}

func (c *remoteServiceClient) GetRemote(ctx context.Context, in *RemoteQueryParams, opts ...grpc.CallOption) (*Remote, error) {
	out := new(Remote)
	err := c.cc.Invoke(ctx, "/remote.remoteService/getRemote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteServiceClient) GetTheatersOf(ctx context.Context, in *RemoteQueryParams, opts ...grpc.CallOption) (*TheatersResponse, error) {
	out := new(TheatersResponse)
	err := c.cc.Invoke(ctx, "/remote.remoteService/getTheatersOf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteServiceClient) GetIscpStatusOf(ctx context.Context, in *RemoteQueryParams, opts ...grpc.CallOption) (*IscpStatusResponse, error) {
	out := new(IscpStatusResponse)
	err := c.cc.Invoke(ctx, "/remote.remoteService/getIscpStatusOf", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteServiceClient) SetTheater(ctx context.Context, in *SetTheaterQueryParams, opts ...grpc.CallOption) (*TheatersResponse, error) {
	out := new(TheatersResponse)
	err := c.cc.Invoke(ctx, "/remote.remoteService/setTheater", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteServiceClient) RemoveTheater(ctx context.Context, in *RemoveTheaterQueryParams, opts ...grpc.CallOption) (*TheatersResponse, error) {
	out := new(TheatersResponse)
	err := c.cc.Invoke(ctx, "/remote.remoteService/removeTheater", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteServiceClient) GetStatus(ctx context.Context, in *RemoteQueryParams, opts ...grpc.CallOption) (*RemoteStatus, error) {
	out := new(RemoteStatus)
	err := c.cc.Invoke(ctx, "/remote.remoteService/getStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteServiceClient) PlayScene(ctx context.Context, in *PlaySceneParams, opts ...grpc.CallOption) (*PlayResponse, error) {
	out := new(PlayResponse)
	err := c.cc.Invoke(ctx, "/remote.remoteService/playScene", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteServiceClient) PlayCommand(ctx context.Context, in *PlayCommandParams, opts ...grpc.CallOption) (*PlayResponse, error) {
	out := new(PlayResponse)
	err := c.cc.Invoke(ctx, "/remote.remoteService/playCommand", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteServiceServer is the server API for RemoteService service.
// All implementations must embed UnimplementedRemoteServiceServer
// for forward compatibility
type RemoteServiceServer interface {
	GetRemote(context.Context, *RemoteQueryParams) (*Remote, error)
	GetTheatersOf(context.Context, *RemoteQueryParams) (*TheatersResponse, error)
	GetIscpStatusOf(context.Context, *RemoteQueryParams) (*IscpStatusResponse, error)
	SetTheater(context.Context, *SetTheaterQueryParams) (*TheatersResponse, error)
	RemoveTheater(context.Context, *RemoveTheaterQueryParams) (*TheatersResponse, error)
	GetStatus(context.Context, *RemoteQueryParams) (*RemoteStatus, error)
	PlayScene(context.Context, *PlaySceneParams) (*PlayResponse, error)
	PlayCommand(context.Context, *PlayCommandParams) (*PlayResponse, error)
	mustEmbedUnimplementedRemoteServiceServer()
}

// UnimplementedRemoteServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRemoteServiceServer struct {
}

func (UnimplementedRemoteServiceServer) GetRemote(context.Context, *RemoteQueryParams) (*Remote, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRemote not implemented")
}
func (UnimplementedRemoteServiceServer) GetTheatersOf(context.Context, *RemoteQueryParams) (*TheatersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTheatersOf not implemented")
}
func (UnimplementedRemoteServiceServer) GetIscpStatusOf(context.Context, *RemoteQueryParams) (*IscpStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIscpStatusOf not implemented")
}
func (UnimplementedRemoteServiceServer) SetTheater(context.Context, *SetTheaterQueryParams) (*TheatersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTheater not implemented")
}
func (UnimplementedRemoteServiceServer) RemoveTheater(context.Context, *RemoveTheaterQueryParams) (*TheatersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveTheater not implemented")
}
func (UnimplementedRemoteServiceServer) GetStatus(context.Context, *RemoteQueryParams) (*RemoteStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedRemoteServiceServer) PlayScene(context.Context, *PlaySceneParams) (*PlayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayScene not implemented")
}
func (UnimplementedRemoteServiceServer) PlayCommand(context.Context, *PlayCommandParams) (*PlayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlayCommand not implemented")
}
func (UnimplementedRemoteServiceServer) mustEmbedUnimplementedRemoteServiceServer() {}

// UnsafeRemoteServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemoteServiceServer will
// result in compilation errors.
type UnsafeRemoteServiceServer interface {
	mustEmbedUnimplementedRemoteServiceServer()
}

func RegisterRemoteServiceServer(s grpc.ServiceRegistrar, srv RemoteServiceServer) {
	s.RegisterService(&_RemoteService_serviceDesc, srv)
}

func _RemoteService_GetRemote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteQueryParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteServiceServer).GetRemote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.remoteService/getRemote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteServiceServer).GetRemote(ctx, req.(*RemoteQueryParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteService_GetTheatersOf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteQueryParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteServiceServer).GetTheatersOf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.remoteService/getTheatersOf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteServiceServer).GetTheatersOf(ctx, req.(*RemoteQueryParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteService_GetIscpStatusOf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteQueryParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteServiceServer).GetIscpStatusOf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.remoteService/getIscpStatusOf",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteServiceServer).GetIscpStatusOf(ctx, req.(*RemoteQueryParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteService_SetTheater_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTheaterQueryParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteServiceServer).SetTheater(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.remoteService/setTheater",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteServiceServer).SetTheater(ctx, req.(*SetTheaterQueryParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteService_RemoveTheater_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveTheaterQueryParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteServiceServer).RemoveTheater(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.remoteService/removeTheater",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteServiceServer).RemoveTheater(ctx, req.(*RemoveTheaterQueryParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteQueryParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.remoteService/getStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteServiceServer).GetStatus(ctx, req.(*RemoteQueryParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteService_PlayScene_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaySceneParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteServiceServer).PlayScene(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.remoteService/playScene",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteServiceServer).PlayScene(ctx, req.(*PlaySceneParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteService_PlayCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlayCommandParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteServiceServer).PlayCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/remote.remoteService/playCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteServiceServer).PlayCommand(ctx, req.(*PlayCommandParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _RemoteService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "remote.remoteService",
	HandlerType: (*RemoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getRemote",
			Handler:    _RemoteService_GetRemote_Handler,
		},
		{
			MethodName: "getTheatersOf",
			Handler:    _RemoteService_GetTheatersOf_Handler,
		},
		{
			MethodName: "getIscpStatusOf",
			Handler:    _RemoteService_GetIscpStatusOf_Handler,
		},
		{
			MethodName: "setTheater",
			Handler:    _RemoteService_SetTheater_Handler,
		},
		{
			MethodName: "removeTheater",
			Handler:    _RemoteService_RemoveTheater_Handler,
		},
		{
			MethodName: "getStatus",
			Handler:    _RemoteService_GetStatus_Handler,
		},
		{
			MethodName: "playScene",
			Handler:    _RemoteService_PlayScene_Handler,
		},
		{
			MethodName: "playCommand",
			Handler:    _RemoteService_PlayCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "definition.proto",
}
