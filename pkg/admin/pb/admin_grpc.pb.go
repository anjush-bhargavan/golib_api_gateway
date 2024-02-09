// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: admin.proto

package __

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

const (
	AdminService_AdminLogin_FullMethodName          = "/pbA.AdminService/AdminLogin"
	AdminService_FindUserByEmailFn_FullMethodName   = "/pbA.AdminService/FindUserByEmailFn"
	AdminService_FindUserByIDFn_FullMethodName      = "/pbA.AdminService/FindUserByIDFn"
	AdminService_FindAllUsersFn_FullMethodName      = "/pbA.AdminService/FindAllUsersFn"
	AdminService_EditUserFn_FullMethodName          = "/pbA.AdminService/EditUserFn"
	AdminService_DeleteUserFn_FullMethodName        = "/pbA.AdminService/DeleteUserFn"
	AdminService_UserCreateBook_FullMethodName      = "/pbA.AdminService/UserCreateBook"
	AdminService_UserFetchBookByID_FullMethodName   = "/pbA.AdminService/UserFetchBookByID"
	AdminService_UserFetchBookByName_FullMethodName = "/pbA.AdminService/UserFetchBookByName"
	AdminService_USerFetchAllBooks_FullMethodName   = "/pbA.AdminService/USerFetchAllBooks"
)

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AdminServiceClient interface {
	AdminLogin(ctx context.Context, in *AdminRequest, opts ...grpc.CallOption) (*AdminResponse, error)
	FindUserByEmailFn(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*User, error)
	FindUserByIDFn(ctx context.Context, in *AUserID, opts ...grpc.CallOption) (*User, error)
	FindAllUsersFn(ctx context.Context, in *AdminNoParam, opts ...grpc.CallOption) (*UserList, error)
	EditUserFn(ctx context.Context, in *UserModel, opts ...grpc.CallOption) (*AdminResponse, error)
	DeleteUserFn(ctx context.Context, in *AUserID, opts ...grpc.CallOption) (*AdminResponse, error)
	UserCreateBook(ctx context.Context, in *ABookModel, opts ...grpc.CallOption) (*AdminResponse, error)
	UserFetchBookByID(ctx context.Context, in *ABookID, opts ...grpc.CallOption) (*ABookModel, error)
	UserFetchBookByName(ctx context.Context, in *ABookName, opts ...grpc.CallOption) (*ABookModel, error)
	USerFetchAllBooks(ctx context.Context, in *AdminNoParam, opts ...grpc.CallOption) (*ABookList, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) AdminLogin(ctx context.Context, in *AdminRequest, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, AdminService_AdminLogin_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) FindUserByEmailFn(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, AdminService_FindUserByEmailFn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) FindUserByIDFn(ctx context.Context, in *AUserID, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, AdminService_FindUserByIDFn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) FindAllUsersFn(ctx context.Context, in *AdminNoParam, opts ...grpc.CallOption) (*UserList, error) {
	out := new(UserList)
	err := c.cc.Invoke(ctx, AdminService_FindAllUsersFn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) EditUserFn(ctx context.Context, in *UserModel, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, AdminService_EditUserFn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) DeleteUserFn(ctx context.Context, in *AUserID, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, AdminService_DeleteUserFn_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) UserCreateBook(ctx context.Context, in *ABookModel, opts ...grpc.CallOption) (*AdminResponse, error) {
	out := new(AdminResponse)
	err := c.cc.Invoke(ctx, AdminService_UserCreateBook_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) UserFetchBookByID(ctx context.Context, in *ABookID, opts ...grpc.CallOption) (*ABookModel, error) {
	out := new(ABookModel)
	err := c.cc.Invoke(ctx, AdminService_UserFetchBookByID_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) UserFetchBookByName(ctx context.Context, in *ABookName, opts ...grpc.CallOption) (*ABookModel, error) {
	out := new(ABookModel)
	err := c.cc.Invoke(ctx, AdminService_UserFetchBookByName_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) USerFetchAllBooks(ctx context.Context, in *AdminNoParam, opts ...grpc.CallOption) (*ABookList, error) {
	out := new(ABookList)
	err := c.cc.Invoke(ctx, AdminService_USerFetchAllBooks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServiceServer is the server API for AdminService service.
// All implementations must embed UnimplementedAdminServiceServer
// for forward compatibility
type AdminServiceServer interface {
	AdminLogin(context.Context, *AdminRequest) (*AdminResponse, error)
	FindUserByEmailFn(context.Context, *UserRequest) (*User, error)
	FindUserByIDFn(context.Context, *AUserID) (*User, error)
	FindAllUsersFn(context.Context, *AdminNoParam) (*UserList, error)
	EditUserFn(context.Context, *UserModel) (*AdminResponse, error)
	DeleteUserFn(context.Context, *AUserID) (*AdminResponse, error)
	UserCreateBook(context.Context, *ABookModel) (*AdminResponse, error)
	UserFetchBookByID(context.Context, *ABookID) (*ABookModel, error)
	UserFetchBookByName(context.Context, *ABookName) (*ABookModel, error)
	USerFetchAllBooks(context.Context, *AdminNoParam) (*ABookList, error)
	mustEmbedUnimplementedAdminServiceServer()
}

// UnimplementedAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (UnimplementedAdminServiceServer) AdminLogin(context.Context, *AdminRequest) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminLogin not implemented")
}
func (UnimplementedAdminServiceServer) FindUserByEmailFn(context.Context, *UserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserByEmailFn not implemented")
}
func (UnimplementedAdminServiceServer) FindUserByIDFn(context.Context, *AUserID) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindUserByIDFn not implemented")
}
func (UnimplementedAdminServiceServer) FindAllUsersFn(context.Context, *AdminNoParam) (*UserList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAllUsersFn not implemented")
}
func (UnimplementedAdminServiceServer) EditUserFn(context.Context, *UserModel) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditUserFn not implemented")
}
func (UnimplementedAdminServiceServer) DeleteUserFn(context.Context, *AUserID) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserFn not implemented")
}
func (UnimplementedAdminServiceServer) UserCreateBook(context.Context, *ABookModel) (*AdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserCreateBook not implemented")
}
func (UnimplementedAdminServiceServer) UserFetchBookByID(context.Context, *ABookID) (*ABookModel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserFetchBookByID not implemented")
}
func (UnimplementedAdminServiceServer) UserFetchBookByName(context.Context, *ABookName) (*ABookModel, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserFetchBookByName not implemented")
}
func (UnimplementedAdminServiceServer) USerFetchAllBooks(context.Context, *AdminNoParam) (*ABookList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method USerFetchAllBooks not implemented")
}
func (UnimplementedAdminServiceServer) mustEmbedUnimplementedAdminServiceServer() {}

// UnsafeAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AdminServiceServer will
// result in compilation errors.
type UnsafeAdminServiceServer interface {
	mustEmbedUnimplementedAdminServiceServer()
}

func RegisterAdminServiceServer(s grpc.ServiceRegistrar, srv AdminServiceServer) {
	s.RegisterService(&AdminService_ServiceDesc, srv)
}

func _AdminService_AdminLogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminLogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_AdminLogin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminLogin(ctx, req.(*AdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_FindUserByEmailFn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).FindUserByEmailFn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_FindUserByEmailFn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).FindUserByEmailFn(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_FindUserByIDFn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AUserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).FindUserByIDFn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_FindUserByIDFn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).FindUserByIDFn(ctx, req.(*AUserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_FindAllUsersFn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminNoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).FindAllUsersFn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_FindAllUsersFn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).FindAllUsersFn(ctx, req.(*AdminNoParam))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_EditUserFn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).EditUserFn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_EditUserFn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).EditUserFn(ctx, req.(*UserModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_DeleteUserFn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AUserID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).DeleteUserFn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_DeleteUserFn_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).DeleteUserFn(ctx, req.(*AUserID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_UserCreateBook_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ABookModel)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).UserCreateBook(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_UserCreateBook_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).UserCreateBook(ctx, req.(*ABookModel))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_UserFetchBookByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ABookID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).UserFetchBookByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_UserFetchBookByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).UserFetchBookByID(ctx, req.(*ABookID))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_UserFetchBookByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ABookName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).UserFetchBookByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_UserFetchBookByName_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).UserFetchBookByName(ctx, req.(*ABookName))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_USerFetchAllBooks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminNoParam)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).USerFetchAllBooks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AdminService_USerFetchAllBooks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).USerFetchAllBooks(ctx, req.(*AdminNoParam))
	}
	return interceptor(ctx, in, info, handler)
}

// AdminService_ServiceDesc is the grpc.ServiceDesc for AdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pbA.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AdminLogin",
			Handler:    _AdminService_AdminLogin_Handler,
		},
		{
			MethodName: "FindUserByEmailFn",
			Handler:    _AdminService_FindUserByEmailFn_Handler,
		},
		{
			MethodName: "FindUserByIDFn",
			Handler:    _AdminService_FindUserByIDFn_Handler,
		},
		{
			MethodName: "FindAllUsersFn",
			Handler:    _AdminService_FindAllUsersFn_Handler,
		},
		{
			MethodName: "EditUserFn",
			Handler:    _AdminService_EditUserFn_Handler,
		},
		{
			MethodName: "DeleteUserFn",
			Handler:    _AdminService_DeleteUserFn_Handler,
		},
		{
			MethodName: "UserCreateBook",
			Handler:    _AdminService_UserCreateBook_Handler,
		},
		{
			MethodName: "UserFetchBookByID",
			Handler:    _AdminService_UserFetchBookByID_Handler,
		},
		{
			MethodName: "UserFetchBookByName",
			Handler:    _AdminService_UserFetchBookByName_Handler,
		},
		{
			MethodName: "USerFetchAllBooks",
			Handler:    _AdminService_USerFetchAllBooks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}