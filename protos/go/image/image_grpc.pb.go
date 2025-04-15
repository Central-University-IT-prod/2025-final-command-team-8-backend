// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: image.proto

package __

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
	Image_UploadImage_FullMethodName = "/image.Image/UploadImage"
	Image_GetImage_FullMethodName    = "/image.Image/GetImage"
	Image_DeleteImage_FullMethodName = "/image.Image/DeleteImage"
)

// ImageClient is the client API for Image service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageClient interface {
	UploadImage(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadImageRequest, UploadImageResponse], error)
	GetImage(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetImageResponse], error)
	DeleteImage(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type imageClient struct {
	cc grpc.ClientConnInterface
}

func NewImageClient(cc grpc.ClientConnInterface) ImageClient {
	return &imageClient{cc}
}

func (c *imageClient) UploadImage(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[UploadImageRequest, UploadImageResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Image_ServiceDesc.Streams[0], Image_UploadImage_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[UploadImageRequest, UploadImageResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Image_UploadImageClient = grpc.ClientStreamingClient[UploadImageRequest, UploadImageResponse]

func (c *imageClient) GetImage(ctx context.Context, in *GetImageRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[GetImageResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &Image_ServiceDesc.Streams[1], Image_GetImage_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[GetImageRequest, GetImageResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Image_GetImageClient = grpc.ServerStreamingClient[GetImageResponse]

func (c *imageClient) DeleteImage(ctx context.Context, in *DeleteImageRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Image_DeleteImage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageServer is the server API for Image service.
// All implementations must embed UnimplementedImageServer
// for forward compatibility.
type ImageServer interface {
	UploadImage(grpc.ClientStreamingServer[UploadImageRequest, UploadImageResponse]) error
	GetImage(*GetImageRequest, grpc.ServerStreamingServer[GetImageResponse]) error
	DeleteImage(context.Context, *DeleteImageRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedImageServer()
}

// UnimplementedImageServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedImageServer struct{}

func (UnimplementedImageServer) UploadImage(grpc.ClientStreamingServer[UploadImageRequest, UploadImageResponse]) error {
	return status.Errorf(codes.Unimplemented, "method UploadImage not implemented")
}
func (UnimplementedImageServer) GetImage(*GetImageRequest, grpc.ServerStreamingServer[GetImageResponse]) error {
	return status.Errorf(codes.Unimplemented, "method GetImage not implemented")
}
func (UnimplementedImageServer) DeleteImage(context.Context, *DeleteImageRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteImage not implemented")
}
func (UnimplementedImageServer) mustEmbedUnimplementedImageServer() {}
func (UnimplementedImageServer) testEmbeddedByValue()               {}

// UnsafeImageServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageServer will
// result in compilation errors.
type UnsafeImageServer interface {
	mustEmbedUnimplementedImageServer()
}

func RegisterImageServer(s grpc.ServiceRegistrar, srv ImageServer) {
	// If the following call pancis, it indicates UnimplementedImageServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Image_ServiceDesc, srv)
}

func _Image_UploadImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ImageServer).UploadImage(&grpc.GenericServerStream[UploadImageRequest, UploadImageResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Image_UploadImageServer = grpc.ClientStreamingServer[UploadImageRequest, UploadImageResponse]

func _Image_GetImage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetImageRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ImageServer).GetImage(m, &grpc.GenericServerStream[GetImageRequest, GetImageResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type Image_GetImageServer = grpc.ServerStreamingServer[GetImageResponse]

func _Image_DeleteImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServer).DeleteImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Image_DeleteImage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServer).DeleteImage(ctx, req.(*DeleteImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Image_ServiceDesc is the grpc.ServiceDesc for Image service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Image_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "image.Image",
	HandlerType: (*ImageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteImage",
			Handler:    _Image_DeleteImage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadImage",
			Handler:       _Image_UploadImage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetImage",
			Handler:       _Image_GetImage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "image.proto",
}
