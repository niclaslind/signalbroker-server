// Code generated by protoc-gen-go. DO NOT EDIT.
// source: system_api.proto

package base

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

func init() { proto.RegisterFile("system_api.proto", fileDescriptor_a502bc89b858c2cc) }

var fileDescriptor_a502bc89b858c2cc = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0xae, 0x2c, 0x2e,
	0x49, 0xcd, 0x8d, 0x4f, 0x2c, 0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0x4a,
	0x2c, 0x4e, 0x95, 0xe2, 0x49, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x83, 0x88, 0x19, 0x95, 0x73, 0xf1,
	0x06, 0x83, 0xd5, 0x05, 0xa7, 0x16, 0x95, 0x65, 0x26, 0xa7, 0x0a, 0x99, 0x71, 0x09, 0xb8, 0xa7,
	0x96, 0x38, 0xe7, 0xe7, 0xa5, 0x65, 0xa6, 0x97, 0x16, 0x25, 0x96, 0x64, 0xe6, 0xe7, 0x09, 0x71,
	0xeb, 0x81, 0x74, 0xea, 0xb9, 0xe6, 0x16, 0x94, 0x54, 0x4a, 0x09, 0x43, 0x38, 0x28, 0x2a, 0x94,
	0x18, 0x84, 0xf4, 0xb8, 0xb8, 0x7d, 0x32, 0x8b, 0x4b, 0x82, 0x33, 0xd3, 0xf3, 0x12, 0x73, 0x8a,
	0x85, 0xf8, 0x21, 0xaa, 0xfc, 0x12, 0x73, 0x53, 0x83, 0x0b, 0x12, 0x93, 0x53, 0xa5, 0x78, 0x20,
	0x02, 0x6e, 0x45, 0x89, 0xb9, 0xa9, 0xc5, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0xfb, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x65, 0xf6, 0xe7, 0x00, 0xa7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SystemServiceClient is the client API for SystemService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SystemServiceClient interface {
	GetConfiguration(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Configuration, error)
	ListSignals(ctx context.Context, in *NameSpace, opts ...grpc.CallOption) (*Frames, error)
}

type systemServiceClient struct {
	cc *grpc.ClientConn
}

func NewSystemServiceClient(cc *grpc.ClientConn) SystemServiceClient {
	return &systemServiceClient{cc}
}

func (c *systemServiceClient) GetConfiguration(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Configuration, error) {
	out := new(Configuration)
	err := c.cc.Invoke(ctx, "/base.SystemService/GetConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *systemServiceClient) ListSignals(ctx context.Context, in *NameSpace, opts ...grpc.CallOption) (*Frames, error) {
	out := new(Frames)
	err := c.cc.Invoke(ctx, "/base.SystemService/ListSignals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SystemServiceServer is the server API for SystemService service.
type SystemServiceServer interface {
	GetConfiguration(context.Context, *Empty) (*Configuration, error)
	ListSignals(context.Context, *NameSpace) (*Frames, error)
}

// UnimplementedSystemServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSystemServiceServer struct {
}

func (*UnimplementedSystemServiceServer) GetConfiguration(ctx context.Context, req *Empty) (*Configuration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfiguration not implemented")
}
func (*UnimplementedSystemServiceServer) ListSignals(ctx context.Context, req *NameSpace) (*Frames, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSignals not implemented")
}

func RegisterSystemServiceServer(s *grpc.Server, srv SystemServiceServer) {
	s.RegisterService(&_SystemService_serviceDesc, srv)
}

func _SystemService_GetConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).GetConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.SystemService/GetConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).GetConfiguration(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SystemService_ListSignals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NameSpace)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SystemServiceServer).ListSignals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.SystemService/ListSignals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SystemServiceServer).ListSignals(ctx, req.(*NameSpace))
	}
	return interceptor(ctx, in, info, handler)
}

var _SystemService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "base.SystemService",
	HandlerType: (*SystemServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfiguration",
			Handler:    _SystemService_GetConfiguration_Handler,
		},
		{
			MethodName: "ListSignals",
			Handler:    _SystemService_ListSignals_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system_api.proto",
}
