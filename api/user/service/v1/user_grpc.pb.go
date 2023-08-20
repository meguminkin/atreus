// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: user/service/v1/user.proto

package v1

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

const (
	UserService_UserRegister_FullMethodName    = "/user.service.v1.UserService/UserRegister"
	UserService_UserLogin_FullMethodName       = "/user.service.v1.UserService/UserLogin"
	UserService_GetUserInfo_FullMethodName     = "/user.service.v1.UserService/GetUserInfo"
	UserService_GetUserInfos_FullMethodName    = "/user.service.v1.UserService/GetUserInfos"
	UserService_UpdateFollow_FullMethodName    = "/user.service.v1.UserService/UpdateFollow"
	UserService_UpdateFollower_FullMethodName  = "/user.service.v1.UserService/UpdateFollower"
	UserService_UpdateFavorited_FullMethodName = "/user.service.v1.UserService/UpdateFavorited"
	UserService_UpdateWork_FullMethodName      = "/user.service.v1.UserService/UpdateWork"
	UserService_UpdateFavorite_FullMethodName  = "/user.service.v1.UserService/UpdateFavorite"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterReply, error)
	UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginReply, error)
	GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoReply, error)
	GetUserInfos(ctx context.Context, in *UserInfosRequest, opts ...grpc.CallOption) (*UserInfosReply, error)
	UpdateFollow(ctx context.Context, in *UpdateFollowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateFollower(ctx context.Context, in *UpdateFollowerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateFavorited(ctx context.Context, in *UpdateFavoritedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateWork(ctx context.Context, in *UpdateWorkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	UpdateFavorite(ctx context.Context, in *UpdateFavoriteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) UserRegister(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterReply, error) {
	out := new(UserRegisterReply)
	err := c.cc.Invoke(ctx, UserService_UserRegister_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginReply, error) {
	out := new(UserLoginReply)
	err := c.cc.Invoke(ctx, UserService_UserLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoReply, error) {
	out := new(UserInfoReply)
	err := c.cc.Invoke(ctx, UserService_GetUserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserInfos(ctx context.Context, in *UserInfosRequest, opts ...grpc.CallOption) (*UserInfosReply, error) {
	out := new(UserInfosReply)
	err := c.cc.Invoke(ctx, UserService_GetUserInfos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateFollow(ctx context.Context, in *UpdateFollowRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserService_UpdateFollow_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateFollower(ctx context.Context, in *UpdateFollowerRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserService_UpdateFollower_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateFavorited(ctx context.Context, in *UpdateFavoritedRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserService_UpdateFavorited_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateWork(ctx context.Context, in *UpdateWorkRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserService_UpdateWork_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateFavorite(ctx context.Context, in *UpdateFavoriteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserService_UpdateFavorite_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	UserRegister(context.Context, *UserRegisterRequest) (*UserRegisterReply, error)
	UserLogin(context.Context, *UserLoginRequest) (*UserLoginReply, error)
	GetUserInfo(context.Context, *UserInfoRequest) (*UserInfoReply, error)
	GetUserInfos(context.Context, *UserInfosRequest) (*UserInfosReply, error)
	UpdateFollow(context.Context, *UpdateFollowRequest) (*emptypb.Empty, error)
	UpdateFollower(context.Context, *UpdateFollowerRequest) (*emptypb.Empty, error)
	UpdateFavorited(context.Context, *UpdateFavoritedRequest) (*emptypb.Empty, error)
	UpdateWork(context.Context, *UpdateWorkRequest) (*emptypb.Empty, error)
	UpdateFavorite(context.Context, *UpdateFavoriteRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) UserRegister(context.Context, *UserRegisterRequest) (*UserRegisterReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserRegister not implemented")
}
func (UnimplementedUserServiceServer) UserLogin(context.Context, *UserLoginRequest) (*UserLoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserLogin not implemented")
}
func (UnimplementedUserServiceServer) GetUserInfo(context.Context, *UserInfoRequest) (*UserInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserServiceServer) GetUserInfos(context.Context, *UserInfosRequest) (*UserInfosReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfos not implemented")
}
func (UnimplementedUserServiceServer) UpdateFollow(context.Context, *UpdateFollowRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFollow not implemented")
}
func (UnimplementedUserServiceServer) UpdateFollower(context.Context, *UpdateFollowerRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFollower not implemented")
}
func (UnimplementedUserServiceServer) UpdateFavorited(context.Context, *UpdateFavoritedRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFavorited not implemented")
}
func (UnimplementedUserServiceServer) UpdateWork(context.Context, *UpdateWorkRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateWork not implemented")
}
func (UnimplementedUserServiceServer) UpdateFavorite(context.Context, *UpdateFavoriteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFavorite not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_UserRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserRegister_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserRegister(ctx, req.(*UserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UserLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UserLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UserLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UserLogin(ctx, req.(*UserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfo(ctx, req.(*UserInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserInfos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfosRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserInfos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUserInfos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserInfos(ctx, req.(*UserInfosRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateFollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFollowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateFollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateFollow_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateFollow(ctx, req.(*UpdateFollowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateFollower_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFollowerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateFollower(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateFollower_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateFollower(ctx, req.(*UpdateFollowerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateFavorited_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFavoritedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateFavorited(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateFavorited_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateFavorited(ctx, req.(*UpdateFavoritedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateWork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateWorkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateWork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateWork_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateWork(ctx, req.(*UpdateWorkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateFavorite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFavoriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateFavorite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateFavorite_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateFavorite(ctx, req.(*UpdateFavoriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.service.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UserRegister",
			Handler:    _UserService_UserRegister_Handler,
		},
		{
			MethodName: "UserLogin",
			Handler:    _UserService_UserLogin_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _UserService_GetUserInfo_Handler,
		},
		{
			MethodName: "GetUserInfos",
			Handler:    _UserService_GetUserInfos_Handler,
		},
		{
			MethodName: "UpdateFollow",
			Handler:    _UserService_UpdateFollow_Handler,
		},
		{
			MethodName: "UpdateFollower",
			Handler:    _UserService_UpdateFollower_Handler,
		},
		{
			MethodName: "UpdateFavorited",
			Handler:    _UserService_UpdateFavorited_Handler,
		},
		{
			MethodName: "UpdateWork",
			Handler:    _UserService_UpdateWork_Handler,
		},
		{
			MethodName: "UpdateFavorite",
			Handler:    _UserService_UpdateFavorite_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/service/v1/user.proto",
}
