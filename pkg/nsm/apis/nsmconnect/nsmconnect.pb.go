// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nsmconnect.proto

package nsmconnect

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/ligato/networkservicemesh/pkg/nsm/apis/common"
import netmesh "github.com/ligato/networkservicemesh/pkg/nsm/apis/netmesh"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// ConnectionRequest is sent by a NSM client to build a connection.
type ConnectionRequest struct {
	// Since connection request will trigger certain actions
	// executed by NSM for a client to address idempotency, request_id
	// will be tracked.
	RequestId            string              `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Metadata             *common.Metadata    `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	NetworkServiceName   string              `protobuf:"bytes,3,opt,name=network_service_name,json=networkServiceName,proto3" json:"network_service_name,omitempty"`
	LinuxNamespace       string              `protobuf:"bytes,4,opt,name=linux_namespace,json=linuxNamespace,proto3" json:"linux_namespace,omitempty"`
	Interface            []*common.Interface `protobuf:"bytes,5,rep,name=interface,proto3" json:"interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ConnectionRequest) Reset()         { *m = ConnectionRequest{} }
func (m *ConnectionRequest) String() string { return proto.CompactTextString(m) }
func (*ConnectionRequest) ProtoMessage()    {}
func (*ConnectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_nsmconnect_adbb61ef747da1f7, []int{0}
}
func (m *ConnectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionRequest.Unmarshal(m, b)
}
func (m *ConnectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionRequest.Marshal(b, m, deterministic)
}
func (dst *ConnectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionRequest.Merge(dst, src)
}
func (m *ConnectionRequest) XXX_Size() int {
	return xxx_messageInfo_ConnectionRequest.Size(m)
}
func (m *ConnectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionRequest proto.InternalMessageInfo

func (m *ConnectionRequest) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *ConnectionRequest) GetMetadata() *common.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ConnectionRequest) GetNetworkServiceName() string {
	if m != nil {
		return m.NetworkServiceName
	}
	return ""
}

func (m *ConnectionRequest) GetLinuxNamespace() string {
	if m != nil {
		return m.LinuxNamespace
	}
	return ""
}

func (m *ConnectionRequest) GetInterface() []*common.Interface {
	if m != nil {
		return m.Interface
	}
	return nil
}

type ConnectionParameters struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Route                []string `protobuf:"bytes,2,rep,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConnectionParameters) Reset()         { *m = ConnectionParameters{} }
func (m *ConnectionParameters) String() string { return proto.CompactTextString(m) }
func (*ConnectionParameters) ProtoMessage()    {}
func (*ConnectionParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_nsmconnect_adbb61ef747da1f7, []int{1}
}
func (m *ConnectionParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionParameters.Unmarshal(m, b)
}
func (m *ConnectionParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionParameters.Marshal(b, m, deterministic)
}
func (dst *ConnectionParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionParameters.Merge(dst, src)
}
func (m *ConnectionParameters) XXX_Size() int {
	return xxx_messageInfo_ConnectionParameters.Size(m)
}
func (m *ConnectionParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionParameters.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionParameters proto.InternalMessageInfo

func (m *ConnectionParameters) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ConnectionParameters) GetRoute() []string {
	if m != nil {
		return m.Route
	}
	return nil
}

// ConnectionAccept is sent back by NSM as a reply to ConnectionRequest
// accepted true will indicate that the connection is accepted, otherwise false
// indicates that connection was refused and admission_error will provide details
// why connection was refused.
type ConnectionAccept struct {
	Accepted             bool                  `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
	AdmissionError       string                `protobuf:"bytes,2,opt,name=admission_error,json=admissionError,proto3" json:"admission_error,omitempty"`
	ConnectionParameters *ConnectionParameters `protobuf:"bytes,3,opt,name=connection_parameters,json=connectionParameters,proto3" json:"connection_parameters,omitempty"`
	Interface            *common.Interface     `protobuf:"bytes,4,opt,name=interface,proto3" json:"interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ConnectionAccept) Reset()         { *m = ConnectionAccept{} }
func (m *ConnectionAccept) String() string { return proto.CompactTextString(m) }
func (*ConnectionAccept) ProtoMessage()    {}
func (*ConnectionAccept) Descriptor() ([]byte, []int) {
	return fileDescriptor_nsmconnect_adbb61ef747da1f7, []int{2}
}
func (m *ConnectionAccept) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionAccept.Unmarshal(m, b)
}
func (m *ConnectionAccept) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionAccept.Marshal(b, m, deterministic)
}
func (dst *ConnectionAccept) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionAccept.Merge(dst, src)
}
func (m *ConnectionAccept) XXX_Size() int {
	return xxx_messageInfo_ConnectionAccept.Size(m)
}
func (m *ConnectionAccept) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionAccept.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionAccept proto.InternalMessageInfo

func (m *ConnectionAccept) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

func (m *ConnectionAccept) GetAdmissionError() string {
	if m != nil {
		return m.AdmissionError
	}
	return ""
}

func (m *ConnectionAccept) GetConnectionParameters() *ConnectionParameters {
	if m != nil {
		return m.ConnectionParameters
	}
	return nil
}

func (m *ConnectionAccept) GetInterface() *common.Interface {
	if m != nil {
		return m.Interface
	}
	return nil
}

// DiscoveryRequest requests NSM to send back all available/known NetworkServices
type DiscoveryRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscoveryRequest) Reset()         { *m = DiscoveryRequest{} }
func (m *DiscoveryRequest) String() string { return proto.CompactTextString(m) }
func (*DiscoveryRequest) ProtoMessage()    {}
func (*DiscoveryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_nsmconnect_adbb61ef747da1f7, []int{3}
}
func (m *DiscoveryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryRequest.Unmarshal(m, b)
}
func (m *DiscoveryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryRequest.Marshal(b, m, deterministic)
}
func (dst *DiscoveryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryRequest.Merge(dst, src)
}
func (m *DiscoveryRequest) XXX_Size() int {
	return xxx_messageInfo_DiscoveryRequest.Size(m)
}
func (m *DiscoveryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryRequest proto.InternalMessageInfo

// DiscoveryRespons carries a list of all available/known to NSM NetworkServices
type DiscoveryResponse struct {
	NetworkService       []*netmesh.NetworkService `protobuf:"bytes,1,rep,name=network_service,json=networkService,proto3" json:"network_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *DiscoveryResponse) Reset()         { *m = DiscoveryResponse{} }
func (m *DiscoveryResponse) String() string { return proto.CompactTextString(m) }
func (*DiscoveryResponse) ProtoMessage()    {}
func (*DiscoveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_nsmconnect_adbb61ef747da1f7, []int{4}
}
func (m *DiscoveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoveryResponse.Unmarshal(m, b)
}
func (m *DiscoveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoveryResponse.Marshal(b, m, deterministic)
}
func (dst *DiscoveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoveryResponse.Merge(dst, src)
}
func (m *DiscoveryResponse) XXX_Size() int {
	return xxx_messageInfo_DiscoveryResponse.Size(m)
}
func (m *DiscoveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoveryResponse proto.InternalMessageInfo

func (m *DiscoveryResponse) GetNetworkService() []*netmesh.NetworkService {
	if m != nil {
		return m.NetworkService
	}
	return nil
}

// ChannelAdvertiseRequest used by NSE to advertise its available channels
type ChannelAdvertiseRequest struct {
	NetmeshChannel       []*netmesh.NetworkServiceChannel `protobuf:"bytes,1,rep,name=netmesh_channel,json=netmeshChannel,proto3" json:"netmesh_channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *ChannelAdvertiseRequest) Reset()         { *m = ChannelAdvertiseRequest{} }
func (m *ChannelAdvertiseRequest) String() string { return proto.CompactTextString(m) }
func (*ChannelAdvertiseRequest) ProtoMessage()    {}
func (*ChannelAdvertiseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_nsmconnect_adbb61ef747da1f7, []int{5}
}
func (m *ChannelAdvertiseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelAdvertiseRequest.Unmarshal(m, b)
}
func (m *ChannelAdvertiseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelAdvertiseRequest.Marshal(b, m, deterministic)
}
func (dst *ChannelAdvertiseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelAdvertiseRequest.Merge(dst, src)
}
func (m *ChannelAdvertiseRequest) XXX_Size() int {
	return xxx_messageInfo_ChannelAdvertiseRequest.Size(m)
}
func (m *ChannelAdvertiseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelAdvertiseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelAdvertiseRequest proto.InternalMessageInfo

func (m *ChannelAdvertiseRequest) GetNetmeshChannel() []*netmesh.NetworkServiceChannel {
	if m != nil {
		return m.NetmeshChannel
	}
	return nil
}

// ChannelAdvertiseResponse used by NSM to confirm if Channel Object has been created successfully
type ChannelAdvertiseResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChannelAdvertiseResponse) Reset()         { *m = ChannelAdvertiseResponse{} }
func (m *ChannelAdvertiseResponse) String() string { return proto.CompactTextString(m) }
func (*ChannelAdvertiseResponse) ProtoMessage()    {}
func (*ChannelAdvertiseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_nsmconnect_adbb61ef747da1f7, []int{6}
}
func (m *ChannelAdvertiseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChannelAdvertiseResponse.Unmarshal(m, b)
}
func (m *ChannelAdvertiseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChannelAdvertiseResponse.Marshal(b, m, deterministic)
}
func (dst *ChannelAdvertiseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChannelAdvertiseResponse.Merge(dst, src)
}
func (m *ChannelAdvertiseResponse) XXX_Size() int {
	return xxx_messageInfo_ChannelAdvertiseResponse.Size(m)
}
func (m *ChannelAdvertiseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChannelAdvertiseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChannelAdvertiseResponse proto.InternalMessageInfo

func (m *ChannelAdvertiseResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*ConnectionRequest)(nil), "nsmconnect.ConnectionRequest")
	proto.RegisterType((*ConnectionParameters)(nil), "nsmconnect.ConnectionParameters")
	proto.RegisterType((*ConnectionAccept)(nil), "nsmconnect.ConnectionAccept")
	proto.RegisterType((*DiscoveryRequest)(nil), "nsmconnect.DiscoveryRequest")
	proto.RegisterType((*DiscoveryResponse)(nil), "nsmconnect.DiscoveryResponse")
	proto.RegisterType((*ChannelAdvertiseRequest)(nil), "nsmconnect.ChannelAdvertiseRequest")
	proto.RegisterType((*ChannelAdvertiseResponse)(nil), "nsmconnect.ChannelAdvertiseResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClientConnectionClient is the client API for ClientConnection service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientConnectionClient interface {
	RequestConnection(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*ConnectionAccept, error)
	RequestDiscovery(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error)
	RequestAdvertiseChannel(ctx context.Context, in *ChannelAdvertiseRequest, opts ...grpc.CallOption) (*ChannelAdvertiseResponse, error)
}

type clientConnectionClient struct {
	cc *grpc.ClientConn
}

func NewClientConnectionClient(cc *grpc.ClientConn) ClientConnectionClient {
	return &clientConnectionClient{cc}
}

func (c *clientConnectionClient) RequestConnection(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (*ConnectionAccept, error) {
	out := new(ConnectionAccept)
	err := c.cc.Invoke(ctx, "/nsmconnect.ClientConnection/RequestConnection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientConnectionClient) RequestDiscovery(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*DiscoveryResponse, error) {
	out := new(DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/nsmconnect.ClientConnection/RequestDiscovery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientConnectionClient) RequestAdvertiseChannel(ctx context.Context, in *ChannelAdvertiseRequest, opts ...grpc.CallOption) (*ChannelAdvertiseResponse, error) {
	out := new(ChannelAdvertiseResponse)
	err := c.cc.Invoke(ctx, "/nsmconnect.ClientConnection/RequestAdvertiseChannel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientConnectionServer is the server API for ClientConnection service.
type ClientConnectionServer interface {
	RequestConnection(context.Context, *ConnectionRequest) (*ConnectionAccept, error)
	RequestDiscovery(context.Context, *DiscoveryRequest) (*DiscoveryResponse, error)
	RequestAdvertiseChannel(context.Context, *ChannelAdvertiseRequest) (*ChannelAdvertiseResponse, error)
}

func RegisterClientConnectionServer(s *grpc.Server, srv ClientConnectionServer) {
	s.RegisterService(&_ClientConnection_serviceDesc, srv)
}

func _ClientConnection_RequestConnection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientConnectionServer).RequestConnection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nsmconnect.ClientConnection/RequestConnection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientConnectionServer).RequestConnection(ctx, req.(*ConnectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientConnection_RequestDiscovery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientConnectionServer).RequestDiscovery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nsmconnect.ClientConnection/RequestDiscovery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientConnectionServer).RequestDiscovery(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientConnection_RequestAdvertiseChannel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelAdvertiseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientConnectionServer).RequestAdvertiseChannel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nsmconnect.ClientConnection/RequestAdvertiseChannel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientConnectionServer).RequestAdvertiseChannel(ctx, req.(*ChannelAdvertiseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientConnection_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nsmconnect.ClientConnection",
	HandlerType: (*ClientConnectionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestConnection",
			Handler:    _ClientConnection_RequestConnection_Handler,
		},
		{
			MethodName: "RequestDiscovery",
			Handler:    _ClientConnection_RequestDiscovery_Handler,
		},
		{
			MethodName: "RequestAdvertiseChannel",
			Handler:    _ClientConnection_RequestAdvertiseChannel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nsmconnect.proto",
}

func init() { proto.RegisterFile("nsmconnect.proto", fileDescriptor_nsmconnect_adbb61ef747da1f7) }

var fileDescriptor_nsmconnect_adbb61ef747da1f7 = []byte{
	// 550 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0xf6, 0xc3, 0xda, 0x33, 0x69, 0x4b, 0xad, 0xa2, 0x46, 0x11, 0x45, 0x55, 0x40, 0xa2,
	0x17, 0xa8, 0x41, 0x85, 0x07, 0x60, 0x2a, 0x63, 0xda, 0x05, 0x65, 0x0a, 0xda, 0x75, 0xe4, 0x3a,
	0x87, 0xd6, 0x5a, 0x63, 0x07, 0xdb, 0x2d, 0xf0, 0x2c, 0x3c, 0x1c, 0x2f, 0xc0, 0x43, 0xa0, 0x38,
	0x4e, 0xda, 0x8e, 0x6e, 0x12, 0x57, 0xf6, 0xf9, 0xce, 0x97, 0xcf, 0xe7, 0x3b, 0xe7, 0x28, 0xe0,
	0x0b, 0x9d, 0x33, 0x29, 0x04, 0x32, 0x33, 0x2a, 0x94, 0x34, 0x92, 0xc0, 0x06, 0x09, 0xaf, 0xe6,
	0xdc, 0x2c, 0x56, 0xb3, 0x11, 0x93, 0x79, 0xbc, 0xe4, 0x73, 0x6a, 0x64, 0x2c, 0xd0, 0x7c, 0x97,
	0xea, 0x4e, 0xa3, 0x5a, 0x73, 0x86, 0x39, 0xea, 0x45, 0x5c, 0xdc, 0xcd, 0x63, 0xa1, 0xf3, 0x98,
	0x16, 0x5c, 0x97, 0x79, 0x0b, 0xba, 0xb3, 0x12, 0x0d, 0x2f, 0xff, 0x5f, 0x88, 0xc9, 0x3c, 0x97,
	0xc2, 0x1d, 0x95, 0x4c, 0xf4, 0xc7, 0x83, 0xce, 0xa4, 0xaa, 0x8d, 0x4b, 0x91, 0xe0, 0xb7, 0x15,
	0x6a, 0x43, 0xfa, 0x00, 0xaa, 0xba, 0xa6, 0x3c, 0x0b, 0xbc, 0x81, 0x37, 0x6c, 0x27, 0x6d, 0x87,
	0x5c, 0x67, 0xe4, 0x35, 0xb4, 0x72, 0x34, 0x34, 0xa3, 0x86, 0x06, 0x07, 0x03, 0x6f, 0x78, 0x3a,
	0xf6, 0x47, 0x4e, 0xf5, 0x93, 0xc3, 0x93, 0x86, 0x41, 0xde, 0x40, 0xd7, 0x55, 0x96, 0xba, 0xd2,
	0x52, 0x41, 0x73, 0x0c, 0x0e, 0xad, 0x2c, 0x71, 0xb9, 0x2f, 0x55, 0x6a, 0x4a, 0x73, 0x24, 0xaf,
	0xe0, 0x7c, 0xc9, 0xc5, 0xea, 0x87, 0xe5, 0xe9, 0x82, 0x32, 0x0c, 0x8e, 0x2c, 0xf9, 0xcc, 0xc2,
	0xd3, 0x1a, 0x25, 0x31, 0xb4, 0xb9, 0x30, 0xa8, 0xbe, 0x96, 0x94, 0xe3, 0xc1, 0xe1, 0xf0, 0x74,
	0xdc, 0xa9, 0x2b, 0xb9, 0xae, 0x13, 0xc9, 0x86, 0x13, 0x7d, 0x84, 0xee, 0xc6, 0xed, 0x0d, 0x55,
	0x34, 0x47, 0x83, 0x4a, 0x93, 0x00, 0x4e, 0x68, 0x96, 0x29, 0xd4, 0xda, 0xb9, 0xad, 0x43, 0xd2,
	0x85, 0x63, 0x25, 0x57, 0x06, 0x83, 0x83, 0xc1, 0xe1, 0xb0, 0x9d, 0x54, 0x41, 0xf4, 0xdb, 0x03,
	0x7f, 0x23, 0x74, 0xc1, 0x18, 0x16, 0x86, 0x84, 0xd0, 0xa2, 0xf6, 0x86, 0x55, 0xcf, 0x5a, 0x49,
	0x13, 0x97, 0x96, 0x68, 0x96, 0x73, 0xad, 0xb9, 0x14, 0x29, 0x2a, 0x25, 0x95, 0xed, 0x5c, 0x3b,
	0x39, 0x6b, 0xe0, 0xcb, 0x12, 0x25, 0xb7, 0xf0, 0x94, 0x35, 0xc2, 0x69, 0xd1, 0x94, 0x68, 0xdb,
	0x75, 0x3a, 0x1e, 0x8c, 0xb6, 0xd6, 0x6b, 0x9f, 0x95, 0xa4, 0xcb, 0xf6, 0x19, 0xdc, 0xe9, 0xd4,
	0x91, 0x95, 0x7a, 0xbc, 0x53, 0x04, 0xfc, 0x0f, 0x5c, 0x33, 0xb9, 0x46, 0xf5, 0xd3, 0xad, 0x45,
	0x74, 0x0b, 0x9d, 0x2d, 0x4c, 0x17, 0x52, 0x68, 0x24, 0xef, 0xe1, 0xfc, 0xde, 0x78, 0x03, 0xcf,
	0x4e, 0xa2, 0x37, 0xaa, 0x37, 0x76, 0xba, 0x33, 0xe2, 0xe4, 0x6c, 0x77, 0xe4, 0xd1, 0x0c, 0x7a,
	0x93, 0x05, 0x15, 0x02, 0x97, 0x17, 0xd9, 0x1a, 0x95, 0xe1, 0x1a, 0xeb, 0x45, 0xbc, 0xb2, 0xe2,
	0xa5, 0x48, 0xca, 0x2a, 0x8a, 0x13, 0x7f, 0xfe, 0x80, 0xb8, 0x13, 0xb2, 0x6f, 0x94, 0x69, 0x17,
	0x47, 0xef, 0x20, 0xf8, 0xf7, 0x0d, 0xe7, 0x20, 0x80, 0x13, 0xbd, 0x62, 0xac, 0x1e, 0x7e, 0x2b,
	0xa9, 0xc3, 0xf1, 0xaf, 0x03, 0xf0, 0x27, 0x4b, 0x8e, 0xc2, 0x6c, 0x5a, 0x4d, 0x6e, 0xa0, 0xe3,
	0xca, 0xdb, 0x02, 0xfb, 0xfb, 0xe7, 0xe2, 0x88, 0xe1, 0xb3, 0xfd, 0x69, 0xb7, 0x38, 0x9f, 0xc1,
	0x77, 0xc4, 0xa6, 0xbd, 0x64, 0xe7, 0x8b, 0xfb, 0x93, 0x08, 0xfb, 0x0f, 0x64, 0x9d, 0xa3, 0x0c,
	0x7a, 0x8e, 0xd9, 0xb8, 0x75, 0xee, 0xc9, 0x8b, 0x9d, 0x4a, 0xf6, 0xb7, 0x3d, 0x7c, 0xf9, 0x38,
	0xa9, 0x7a, 0x65, 0xf6, 0xc4, 0xfe, 0x42, 0xde, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x2f,
	0xc8, 0x05, 0xf2, 0x04, 0x00, 0x00,
}
