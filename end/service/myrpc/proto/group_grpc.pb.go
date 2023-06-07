// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: group.proto

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

// GroupClient is the client API for Group service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GroupClient interface {
	// 设备组列表
	GroupGet(ctx context.Context, in *GroupGetReq, opts ...grpc.CallOption) (*GroupGetResp, error)
	// 增加设备组
	GroupPost(ctx context.Context, in *GroupPostReq, opts ...grpc.CallOption) (*GroupPostResp, error)
	// 删除设备组
	GroupDel(ctx context.Context, in *GroupDelReq, opts ...grpc.CallOption) (*GroupDelResp, error)
	// 修改设备组
	GroupPut(ctx context.Context, in *GroupPutReq, opts ...grpc.CallOption) (*GroupPutResp, error)
}

type groupClient struct {
	cc grpc.ClientConnInterface
}

func NewGroupClient(cc grpc.ClientConnInterface) GroupClient {
	return &groupClient{cc}
}

func (c *groupClient) GroupGet(ctx context.Context, in *GroupGetReq, opts ...grpc.CallOption) (*GroupGetResp, error) {
	out := new(GroupGetResp)
	err := c.cc.Invoke(ctx, "/proto.Group/groupGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) GroupPost(ctx context.Context, in *GroupPostReq, opts ...grpc.CallOption) (*GroupPostResp, error) {
	out := new(GroupPostResp)
	err := c.cc.Invoke(ctx, "/proto.Group/groupPost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) GroupDel(ctx context.Context, in *GroupDelReq, opts ...grpc.CallOption) (*GroupDelResp, error) {
	out := new(GroupDelResp)
	err := c.cc.Invoke(ctx, "/proto.Group/groupDel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *groupClient) GroupPut(ctx context.Context, in *GroupPutReq, opts ...grpc.CallOption) (*GroupPutResp, error) {
	out := new(GroupPutResp)
	err := c.cc.Invoke(ctx, "/proto.Group/groupPut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GroupServer is the server API for Group service.
// All implementations must embed UnimplementedGroupServer
// for forward compatibility
type GroupServer interface {
	// 设备组列表
	GroupGet(context.Context, *GroupGetReq) (*GroupGetResp, error)
	// 增加设备组
	GroupPost(context.Context, *GroupPostReq) (*GroupPostResp, error)
	// 删除设备组
	GroupDel(context.Context, *GroupDelReq) (*GroupDelResp, error)
	// 修改设备组
	GroupPut(context.Context, *GroupPutReq) (*GroupPutResp, error)
	mustEmbedUnimplementedGroupServer()
}

// UnimplementedGroupServer must be embedded to have forward compatible implementations.
type UnimplementedGroupServer struct {
}

func (UnimplementedGroupServer) GroupGet(context.Context, *GroupGetReq) (*GroupGetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupGet not implemented")
}
func (UnimplementedGroupServer) GroupPost(context.Context, *GroupPostReq) (*GroupPostResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupPost not implemented")
}
func (UnimplementedGroupServer) GroupDel(context.Context, *GroupDelReq) (*GroupDelResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupDel not implemented")
}
func (UnimplementedGroupServer) GroupPut(context.Context, *GroupPutReq) (*GroupPutResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GroupPut not implemented")
}
func (UnimplementedGroupServer) mustEmbedUnimplementedGroupServer() {}

// UnsafeGroupServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GroupServer will
// result in compilation errors.
type UnsafeGroupServer interface {
	mustEmbedUnimplementedGroupServer()
}

func RegisterGroupServer(s grpc.ServiceRegistrar, srv GroupServer) {
	s.RegisterService(&Group_ServiceDesc, srv)
}

func _Group_GroupGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupGetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).GroupGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Group/groupGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).GroupGet(ctx, req.(*GroupGetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_GroupPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupPostReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).GroupPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Group/groupPost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).GroupPost(ctx, req.(*GroupPostReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_GroupDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupDelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).GroupDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Group/groupDel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).GroupDel(ctx, req.(*GroupDelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Group_GroupPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GroupPutReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GroupServer).GroupPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Group/groupPut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GroupServer).GroupPut(ctx, req.(*GroupPutReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Group_ServiceDesc is the grpc.ServiceDesc for Group service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Group_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Group",
	HandlerType: (*GroupServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "groupGet",
			Handler:    _Group_GroupGet_Handler,
		},
		{
			MethodName: "groupPost",
			Handler:    _Group_GroupPost_Handler,
		},
		{
			MethodName: "groupDel",
			Handler:    _Group_GroupDel_Handler,
		},
		{
			MethodName: "groupPut",
			Handler:    _Group_GroupPut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "group.proto",
}
