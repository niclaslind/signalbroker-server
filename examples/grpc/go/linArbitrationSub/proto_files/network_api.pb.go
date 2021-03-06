// Code generated by protoc-gen-go. DO NOT EDIT.
// source: network_api.proto

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

type SubscriberConfig struct {
	ClientId             *ClientId  `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Signals              *SignalIds `protobuf:"bytes,2,opt,name=signals,proto3" json:"signals,omitempty"`
	OnChange             bool       `protobuf:"varint,3,opt,name=onChange,proto3" json:"onChange,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SubscriberConfig) Reset()         { *m = SubscriberConfig{} }
func (m *SubscriberConfig) String() string { return proto.CompactTextString(m) }
func (*SubscriberConfig) ProtoMessage()    {}
func (*SubscriberConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_22733250f3c9b488, []int{0}
}

func (m *SubscriberConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubscriberConfig.Unmarshal(m, b)
}
func (m *SubscriberConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubscriberConfig.Marshal(b, m, deterministic)
}
func (m *SubscriberConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubscriberConfig.Merge(m, src)
}
func (m *SubscriberConfig) XXX_Size() int {
	return xxx_messageInfo_SubscriberConfig.Size(m)
}
func (m *SubscriberConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SubscriberConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SubscriberConfig proto.InternalMessageInfo

func (m *SubscriberConfig) GetClientId() *ClientId {
	if m != nil {
		return m.ClientId
	}
	return nil
}

func (m *SubscriberConfig) GetSignals() *SignalIds {
	if m != nil {
		return m.Signals
	}
	return nil
}

func (m *SubscriberConfig) GetOnChange() bool {
	if m != nil {
		return m.OnChange
	}
	return false
}

type SignalIds struct {
	SignalId             []*SignalId `protobuf:"bytes,1,rep,name=signalId,proto3" json:"signalId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SignalIds) Reset()         { *m = SignalIds{} }
func (m *SignalIds) String() string { return proto.CompactTextString(m) }
func (*SignalIds) ProtoMessage()    {}
func (*SignalIds) Descriptor() ([]byte, []int) {
	return fileDescriptor_22733250f3c9b488, []int{1}
}

func (m *SignalIds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignalIds.Unmarshal(m, b)
}
func (m *SignalIds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignalIds.Marshal(b, m, deterministic)
}
func (m *SignalIds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignalIds.Merge(m, src)
}
func (m *SignalIds) XXX_Size() int {
	return xxx_messageInfo_SignalIds.Size(m)
}
func (m *SignalIds) XXX_DiscardUnknown() {
	xxx_messageInfo_SignalIds.DiscardUnknown(m)
}

var xxx_messageInfo_SignalIds proto.InternalMessageInfo

func (m *SignalIds) GetSignalId() []*SignalId {
	if m != nil {
		return m.SignalId
	}
	return nil
}

type Signals struct {
	Signal               []*Signal `protobuf:"bytes,1,rep,name=signal,proto3" json:"signal,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Signals) Reset()         { *m = Signals{} }
func (m *Signals) String() string { return proto.CompactTextString(m) }
func (*Signals) ProtoMessage()    {}
func (*Signals) Descriptor() ([]byte, []int) {
	return fileDescriptor_22733250f3c9b488, []int{2}
}

func (m *Signals) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signals.Unmarshal(m, b)
}
func (m *Signals) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signals.Marshal(b, m, deterministic)
}
func (m *Signals) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signals.Merge(m, src)
}
func (m *Signals) XXX_Size() int {
	return xxx_messageInfo_Signals.Size(m)
}
func (m *Signals) XXX_DiscardUnknown() {
	xxx_messageInfo_Signals.DiscardUnknown(m)
}

var xxx_messageInfo_Signals proto.InternalMessageInfo

func (m *Signals) GetSignal() []*Signal {
	if m != nil {
		return m.Signal
	}
	return nil
}

type PublisherConfig struct {
	Signals              *Signals  `protobuf:"bytes,1,opt,name=signals,proto3" json:"signals,omitempty"`
	ClientId             *ClientId `protobuf:"bytes,2,opt,name=clientId,proto3" json:"clientId,omitempty"`
	Frequency            int32     `protobuf:"varint,3,opt,name=frequency,proto3" json:"frequency,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PublisherConfig) Reset()         { *m = PublisherConfig{} }
func (m *PublisherConfig) String() string { return proto.CompactTextString(m) }
func (*PublisherConfig) ProtoMessage()    {}
func (*PublisherConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_22733250f3c9b488, []int{3}
}

func (m *PublisherConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PublisherConfig.Unmarshal(m, b)
}
func (m *PublisherConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PublisherConfig.Marshal(b, m, deterministic)
}
func (m *PublisherConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PublisherConfig.Merge(m, src)
}
func (m *PublisherConfig) XXX_Size() int {
	return xxx_messageInfo_PublisherConfig.Size(m)
}
func (m *PublisherConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_PublisherConfig.DiscardUnknown(m)
}

var xxx_messageInfo_PublisherConfig proto.InternalMessageInfo

func (m *PublisherConfig) GetSignals() *Signals {
	if m != nil {
		return m.Signals
	}
	return nil
}

func (m *PublisherConfig) GetClientId() *ClientId {
	if m != nil {
		return m.ClientId
	}
	return nil
}

func (m *PublisherConfig) GetFrequency() int32 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

type Signal struct {
	Id *SignalId `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*Signal_Integer
	//	*Signal_Double
	//	*Signal_Arbitration
	//	*Signal_Empty
	Payload              isSignal_Payload `protobuf_oneof:"payload"`
	Raw                  []byte           `protobuf:"bytes,5,opt,name=raw,proto3" json:"raw,omitempty"`
	Timestamp            int64            `protobuf:"varint,7,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Signal) Reset()         { *m = Signal{} }
func (m *Signal) String() string { return proto.CompactTextString(m) }
func (*Signal) ProtoMessage()    {}
func (*Signal) Descriptor() ([]byte, []int) {
	return fileDescriptor_22733250f3c9b488, []int{4}
}

func (m *Signal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signal.Unmarshal(m, b)
}
func (m *Signal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signal.Marshal(b, m, deterministic)
}
func (m *Signal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signal.Merge(m, src)
}
func (m *Signal) XXX_Size() int {
	return xxx_messageInfo_Signal.Size(m)
}
func (m *Signal) XXX_DiscardUnknown() {
	xxx_messageInfo_Signal.DiscardUnknown(m)
}

var xxx_messageInfo_Signal proto.InternalMessageInfo

func (m *Signal) GetId() *SignalId {
	if m != nil {
		return m.Id
	}
	return nil
}

type isSignal_Payload interface {
	isSignal_Payload()
}

type Signal_Integer struct {
	Integer int64 `protobuf:"varint,2,opt,name=integer,proto3,oneof"`
}

type Signal_Double struct {
	Double float64 `protobuf:"fixed64,3,opt,name=double,proto3,oneof"`
}

type Signal_Arbitration struct {
	Arbitration bool `protobuf:"varint,4,opt,name=arbitration,proto3,oneof"`
}

type Signal_Empty struct {
	Empty bool `protobuf:"varint,6,opt,name=empty,proto3,oneof"`
}

func (*Signal_Integer) isSignal_Payload() {}

func (*Signal_Double) isSignal_Payload() {}

func (*Signal_Arbitration) isSignal_Payload() {}

func (*Signal_Empty) isSignal_Payload() {}

func (m *Signal) GetPayload() isSignal_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Signal) GetInteger() int64 {
	if x, ok := m.GetPayload().(*Signal_Integer); ok {
		return x.Integer
	}
	return 0
}

func (m *Signal) GetDouble() float64 {
	if x, ok := m.GetPayload().(*Signal_Double); ok {
		return x.Double
	}
	return 0
}

func (m *Signal) GetArbitration() bool {
	if x, ok := m.GetPayload().(*Signal_Arbitration); ok {
		return x.Arbitration
	}
	return false
}

func (m *Signal) GetEmpty() bool {
	if x, ok := m.GetPayload().(*Signal_Empty); ok {
		return x.Empty
	}
	return false
}

func (m *Signal) GetRaw() []byte {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *Signal) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Signal) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Signal_Integer)(nil),
		(*Signal_Double)(nil),
		(*Signal_Arbitration)(nil),
		(*Signal_Empty)(nil),
	}
}

func init() {
	proto.RegisterType((*SubscriberConfig)(nil), "base.SubscriberConfig")
	proto.RegisterType((*SignalIds)(nil), "base.SignalIds")
	proto.RegisterType((*Signals)(nil), "base.Signals")
	proto.RegisterType((*PublisherConfig)(nil), "base.PublisherConfig")
	proto.RegisterType((*Signal)(nil), "base.Signal")
}

func init() { proto.RegisterFile("network_api.proto", fileDescriptor_22733250f3c9b488) }

var fileDescriptor_22733250f3c9b488 = []byte{
	// 440 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x8d, 0xdb, 0x6d, 0xd2, 0x4e, 0x4b, 0x77, 0xb1, 0xc4, 0xca, 0x8a, 0x10, 0xaa, 0x22, 0x24,
	0x0a, 0x87, 0x2e, 0x5a, 0x24, 0x38, 0x22, 0x51, 0x21, 0x75, 0x2f, 0x08, 0xb9, 0xdc, 0x91, 0x93,
	0x78, 0xbb, 0x16, 0x89, 0x1d, 0x6c, 0x97, 0x55, 0xef, 0xf0, 0x53, 0xfc, 0x07, 0xff, 0x83, 0x12,
	0x3b, 0x69, 0x28, 0x87, 0xbd, 0x75, 0xde, 0x9b, 0xe7, 0xe9, 0x7b, 0x33, 0x81, 0xc7, 0x92, 0xdb,
	0x7b, 0xa5, 0xbf, 0x7d, 0x65, 0x95, 0x58, 0x55, 0x5a, 0x59, 0x85, 0xcf, 0x52, 0x66, 0x78, 0x3c,
	0xcb, 0x54, 0x59, 0x2a, 0xe9, 0xb0, 0xe4, 0x17, 0x82, 0x8b, 0xed, 0x3e, 0x35, 0x99, 0x16, 0x29,
	0xd7, 0x6b, 0x25, 0x6f, 0xc5, 0x0e, 0xbf, 0x82, 0x71, 0x56, 0x08, 0x2e, 0xed, 0x4d, 0x4e, 0xd0,
	0x02, 0x2d, 0xa7, 0xd7, 0xf3, 0x55, 0xad, 0x5d, 0xad, 0x3d, 0x4a, 0x3b, 0x1e, 0xbf, 0x84, 0xc8,
	0x88, 0x9d, 0x64, 0x85, 0x21, 0x83, 0xa6, 0xf5, 0xdc, 0xb5, 0x6e, 0x1b, 0xf0, 0x26, 0x37, 0xb4,
	0xe5, 0x71, 0x0c, 0x63, 0x25, 0xd7, 0x77, 0x4c, 0xee, 0x38, 0x19, 0x2e, 0xd0, 0x72, 0x4c, 0xbb,
	0x3a, 0x79, 0x07, 0x93, 0x4e, 0x51, 0xcf, 0x37, 0xbe, 0x20, 0x68, 0x31, 0x3c, 0xce, 0x6f, 0x5b,
	0x68, 0xc7, 0x27, 0x57, 0x10, 0x6d, 0xfd, 0xfb, 0xcf, 0x21, 0x74, 0xb0, 0x17, 0xcd, 0xfa, 0x22,
	0xea, 0xb9, 0xe4, 0x27, 0x82, 0xf3, 0xcf, 0xfb, 0xb4, 0x10, 0xe6, 0xae, 0x33, 0xfc, 0xe2, 0x68,
	0xc2, 0xf9, 0x7d, 0xd4, 0x97, 0xf6, 0x2c, 0xf4, 0x93, 0x19, 0x3c, 0x90, 0xcc, 0x53, 0x98, 0xdc,
	0x6a, 0xfe, 0x7d, 0xcf, 0x65, 0x76, 0x68, 0xfc, 0x8e, 0xe8, 0x11, 0x48, 0xfe, 0x20, 0x08, 0xdd,
	0xf3, 0xf8, 0x19, 0x0c, 0xc4, 0x49, 0xd0, 0x9d, 0xd1, 0x81, 0xc8, 0x71, 0x0c, 0x91, 0x90, 0x96,
	0xef, 0xb8, 0x6e, 0x66, 0x0e, 0x37, 0x01, 0x6d, 0x01, 0x4c, 0x20, 0xcc, 0xd5, 0x3e, 0x2d, 0x5c,
	0xa2, 0x68, 0x13, 0x50, 0x5f, 0xe3, 0x04, 0xa6, 0x4c, 0xa7, 0xc2, 0x6a, 0x66, 0x85, 0x92, 0xe4,
	0xac, 0x0e, 0x7c, 0x13, 0xd0, 0x3e, 0x88, 0x2f, 0x61, 0xc4, 0xcb, 0xca, 0x1e, 0x48, 0xe8, 0x59,
	0x57, 0xe2, 0x0b, 0x18, 0x6a, 0x76, 0x4f, 0x46, 0x0b, 0xb4, 0x9c, 0xd1, 0xfa, 0x67, 0x6d, 0xc6,
	0x8a, 0x92, 0x1b, 0xcb, 0xca, 0x8a, 0x44, 0xf5, 0xbf, 0xa0, 0x47, 0xe0, 0xc3, 0x04, 0xa2, 0x8a,
	0x1d, 0x0a, 0xc5, 0xf2, 0xeb, 0xdf, 0x08, 0xe6, 0x9f, 0xdc, 0xe9, 0x6d, 0xb9, 0xfe, 0x21, 0x32,
	0x8e, 0xdf, 0x03, 0xee, 0x4e, 0xec, 0x8b, 0x6a, 0xb7, 0x75, 0xe9, 0x9d, 0x9e, 0x1c, 0x5f, 0xfc,
	0x6f, 0xf4, 0x49, 0xf0, 0x1a, 0xe1, 0xb7, 0x30, 0xf7, 0x1b, 0x6b, 0xc5, 0x4f, 0x5c, 0xd3, 0xc9,
	0x1e, 0xe3, 0xa9, 0x83, 0x3f, 0xd6, 0x26, 0x92, 0x00, 0x5f, 0xc1, 0x94, 0x72, 0x96, 0xb7, 0xa2,
	0xd3, 0xcb, 0xfc, 0x6f, 0x54, 0x1a, 0x36, 0x1f, 0xc5, 0x9b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x49, 0x8a, 0x0d, 0x19, 0x3d, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetworkServiceClient is the client API for NetworkService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetworkServiceClient interface {
	SubscribeToSignals(ctx context.Context, in *SubscriberConfig, opts ...grpc.CallOption) (NetworkService_SubscribeToSignalsClient, error)
	PublishSignals(ctx context.Context, in *PublisherConfig, opts ...grpc.CallOption) (*Empty, error)
	ReadSignals(ctx context.Context, in *SignalIds, opts ...grpc.CallOption) (*Signals, error)
}

type networkServiceClient struct {
	cc *grpc.ClientConn
}

func NewNetworkServiceClient(cc *grpc.ClientConn) NetworkServiceClient {
	return &networkServiceClient{cc}
}

func (c *networkServiceClient) SubscribeToSignals(ctx context.Context, in *SubscriberConfig, opts ...grpc.CallOption) (NetworkService_SubscribeToSignalsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NetworkService_serviceDesc.Streams[0], "/base.NetworkService/SubscribeToSignals", opts...)
	if err != nil {
		return nil, err
	}
	x := &networkServiceSubscribeToSignalsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NetworkService_SubscribeToSignalsClient interface {
	Recv() (*Signals, error)
	grpc.ClientStream
}

type networkServiceSubscribeToSignalsClient struct {
	grpc.ClientStream
}

func (x *networkServiceSubscribeToSignalsClient) Recv() (*Signals, error) {
	m := new(Signals)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *networkServiceClient) PublishSignals(ctx context.Context, in *PublisherConfig, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/base.NetworkService/PublishSignals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceClient) ReadSignals(ctx context.Context, in *SignalIds, opts ...grpc.CallOption) (*Signals, error) {
	out := new(Signals)
	err := c.cc.Invoke(ctx, "/base.NetworkService/ReadSignals", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NetworkServiceServer is the server API for NetworkService service.
type NetworkServiceServer interface {
	SubscribeToSignals(*SubscriberConfig, NetworkService_SubscribeToSignalsServer) error
	PublishSignals(context.Context, *PublisherConfig) (*Empty, error)
	ReadSignals(context.Context, *SignalIds) (*Signals, error)
}

// UnimplementedNetworkServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNetworkServiceServer struct {
}

func (*UnimplementedNetworkServiceServer) SubscribeToSignals(req *SubscriberConfig, srv NetworkService_SubscribeToSignalsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToSignals not implemented")
}
func (*UnimplementedNetworkServiceServer) PublishSignals(ctx context.Context, req *PublisherConfig) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishSignals not implemented")
}
func (*UnimplementedNetworkServiceServer) ReadSignals(ctx context.Context, req *SignalIds) (*Signals, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadSignals not implemented")
}

func RegisterNetworkServiceServer(s *grpc.Server, srv NetworkServiceServer) {
	s.RegisterService(&_NetworkService_serviceDesc, srv)
}

func _NetworkService_SubscribeToSignals_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscriberConfig)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NetworkServiceServer).SubscribeToSignals(m, &networkServiceSubscribeToSignalsServer{stream})
}

type NetworkService_SubscribeToSignalsServer interface {
	Send(*Signals) error
	grpc.ServerStream
}

type networkServiceSubscribeToSignalsServer struct {
	grpc.ServerStream
}

func (x *networkServiceSubscribeToSignalsServer) Send(m *Signals) error {
	return x.ServerStream.SendMsg(m)
}

func _NetworkService_PublishSignals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublisherConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).PublishSignals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.NetworkService/PublishSignals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).PublishSignals(ctx, req.(*PublisherConfig))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkService_ReadSignals_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignalIds)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceServer).ReadSignals(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/base.NetworkService/ReadSignals",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceServer).ReadSignals(ctx, req.(*SignalIds))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "base.NetworkService",
	HandlerType: (*NetworkServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishSignals",
			Handler:    _NetworkService_PublishSignals_Handler,
		},
		{
			MethodName: "ReadSignals",
			Handler:    _NetworkService_ReadSignals_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToSignals",
			Handler:       _NetworkService_SubscribeToSignals_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "network_api.proto",
}
