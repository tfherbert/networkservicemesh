// Code generated by protoc-gen-go. DO NOT EDIT.
// source: netmesh.proto

package netmesh

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/ligato/networkservicemesh/pkg/nsm/apis/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type NetworkServiceChannel struct {
	Metadata             *common.Metadata    `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	NetworkServiceName   string              `protobuf:"bytes,2,opt,name=network_service_name,json=networkServiceName,proto3" json:"network_service_name,omitempty"`
	NseProviderName      string              `protobuf:"bytes,3,opt,name=nse_provider_name,json=nseProviderName,proto3" json:"nse_provider_name,omitempty"`
	Payload              string              `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	SocketLocation       string              `protobuf:"bytes,5,opt,name=socket_location,json=socketLocation,proto3" json:"socket_location,omitempty"`
	Interface            []*common.Interface `protobuf:"bytes,6,rep,name=interface,proto3" json:"interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *NetworkServiceChannel) Reset()         { *m = NetworkServiceChannel{} }
func (m *NetworkServiceChannel) String() string { return proto.CompactTextString(m) }
func (*NetworkServiceChannel) ProtoMessage()    {}
func (*NetworkServiceChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_netmesh_22b23a6c883e53c9, []int{0}
}
func (m *NetworkServiceChannel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkServiceChannel.Unmarshal(m, b)
}
func (m *NetworkServiceChannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkServiceChannel.Marshal(b, m, deterministic)
}
func (dst *NetworkServiceChannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkServiceChannel.Merge(dst, src)
}
func (m *NetworkServiceChannel) XXX_Size() int {
	return xxx_messageInfo_NetworkServiceChannel.Size(m)
}
func (m *NetworkServiceChannel) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkServiceChannel.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkServiceChannel proto.InternalMessageInfo

func (m *NetworkServiceChannel) GetMetadata() *common.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *NetworkServiceChannel) GetNetworkServiceName() string {
	if m != nil {
		return m.NetworkServiceName
	}
	return ""
}

func (m *NetworkServiceChannel) GetNseProviderName() string {
	if m != nil {
		return m.NseProviderName
	}
	return ""
}

func (m *NetworkServiceChannel) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *NetworkServiceChannel) GetSocketLocation() string {
	if m != nil {
		return m.SocketLocation
	}
	return ""
}

func (m *NetworkServiceChannel) GetInterface() []*common.Interface {
	if m != nil {
		return m.Interface
	}
	return nil
}

type NetworkServiceEndpoint struct {
	Metadata             *common.Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *NetworkServiceEndpoint) Reset()         { *m = NetworkServiceEndpoint{} }
func (m *NetworkServiceEndpoint) String() string { return proto.CompactTextString(m) }
func (*NetworkServiceEndpoint) ProtoMessage()    {}
func (*NetworkServiceEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_netmesh_22b23a6c883e53c9, []int{1}
}
func (m *NetworkServiceEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkServiceEndpoint.Unmarshal(m, b)
}
func (m *NetworkServiceEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkServiceEndpoint.Marshal(b, m, deterministic)
}
func (dst *NetworkServiceEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkServiceEndpoint.Merge(dst, src)
}
func (m *NetworkServiceEndpoint) XXX_Size() int {
	return xxx_messageInfo_NetworkServiceEndpoint.Size(m)
}
func (m *NetworkServiceEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkServiceEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkServiceEndpoint proto.InternalMessageInfo

func (m *NetworkServiceEndpoint) GetMetadata() *common.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type NetworkService struct {
	Metadata             *common.Metadata         `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Channel              []*NetworkServiceChannel `protobuf:"bytes,2,rep,name=channel,proto3" json:"channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *NetworkService) Reset()         { *m = NetworkService{} }
func (m *NetworkService) String() string { return proto.CompactTextString(m) }
func (*NetworkService) ProtoMessage()    {}
func (*NetworkService) Descriptor() ([]byte, []int) {
	return fileDescriptor_netmesh_22b23a6c883e53c9, []int{2}
}
func (m *NetworkService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkService.Unmarshal(m, b)
}
func (m *NetworkService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkService.Marshal(b, m, deterministic)
}
func (dst *NetworkService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkService.Merge(dst, src)
}
func (m *NetworkService) XXX_Size() int {
	return xxx_messageInfo_NetworkService.Size(m)
}
func (m *NetworkService) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkService.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkService proto.InternalMessageInfo

func (m *NetworkService) GetMetadata() *common.Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *NetworkService) GetChannel() []*NetworkServiceChannel {
	if m != nil {
		return m.Channel
	}
	return nil
}

func init() {
	proto.RegisterType((*NetworkServiceChannel)(nil), "netmesh.NetworkServiceChannel")
	proto.RegisterType((*NetworkServiceEndpoint)(nil), "netmesh.NetworkServiceEndpoint")
	proto.RegisterType((*NetworkService)(nil), "netmesh.NetworkService")
}

func init() { proto.RegisterFile("netmesh.proto", fileDescriptor_netmesh_22b23a6c883e53c9) }

var fileDescriptor_netmesh_22b23a6c883e53c9 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x69, 0xab, 0xad, 0x4d, 0xb1, 0xb5, 0x41, 0x25, 0x78, 0x90, 0xd2, 0x8b, 0x45, 0x64,
	0x23, 0xf5, 0xe2, 0x5d, 0x2a, 0x08, 0x5a, 0x64, 0x7d, 0x80, 0x25, 0xcd, 0x8e, 0x6d, 0xe8, 0x66,
	0x66, 0xd9, 0xc4, 0xaa, 0x4f, 0xe2, 0xeb, 0x8a, 0xbb, 0x59, 0xa5, 0xe0, 0xa5, 0xa7, 0x65, 0xfe,
	0xfd, 0x66, 0x26, 0xf9, 0xff, 0xb0, 0x43, 0x04, 0x6f, 0xc1, 0xad, 0xa2, 0xbc, 0x20, 0x4f, 0xbc,
	0x13, 0xca, 0xb3, 0xd9, 0xd2, 0xf8, 0xd5, 0xdb, 0x22, 0xd2, 0x64, 0x65, 0x66, 0x96, 0xca, 0x93,
	0x44, 0xf0, 0xef, 0x54, 0xac, 0x1d, 0x14, 0x1b, 0xa3, 0xe1, 0x87, 0x92, 0xf9, 0x7a, 0x29, 0xd1,
	0x59, 0xa9, 0x72, 0xe3, 0xa4, 0x26, 0x6b, 0x09, 0xc3, 0xa7, 0x9a, 0x37, 0xfe, 0x6a, 0xb2, 0x93,
	0x79, 0xd5, 0xf7, 0x52, 0xf5, 0xdd, 0xad, 0x14, 0x22, 0x64, 0xfc, 0x8a, 0x1d, 0x58, 0xf0, 0x2a,
	0x55, 0x5e, 0x89, 0xc6, 0xa8, 0x31, 0xe9, 0x4d, 0x8f, 0xa2, 0xd0, 0xfa, 0x14, 0xf4, 0xf8, 0x97,
	0xe0, 0xd7, 0xec, 0x38, 0xac, 0x4f, 0xc2, 0xfe, 0x04, 0x95, 0x05, 0xd1, 0x1c, 0x35, 0x26, 0xdd,
	0x98, 0xe3, 0xd6, 0x8a, 0xb9, 0xb2, 0xc0, 0x2f, 0xd9, 0x10, 0x1d, 0x24, 0x79, 0x41, 0x1b, 0x93,
	0x42, 0x51, 0xe1, 0xad, 0x12, 0x1f, 0xa0, 0x83, 0xe7, 0xa0, 0x97, 0xac, 0x60, 0x9d, 0x5c, 0x7d,
	0x66, 0xa4, 0x52, 0xb1, 0x57, 0x12, 0x75, 0xc9, 0x2f, 0xd8, 0xc0, 0x91, 0x5e, 0x83, 0x4f, 0x32,
	0xd2, 0xca, 0x1b, 0x42, 0xb1, 0x5f, 0x12, 0xfd, 0x4a, 0x7e, 0x0c, 0x2a, 0x97, 0xac, 0x6b, 0xd0,
	0x43, 0xf1, 0xaa, 0x34, 0x88, 0xf6, 0xa8, 0x35, 0xe9, 0x4d, 0x87, 0xf5, 0x7d, 0x1e, 0xea, 0x1f,
	0xf1, 0x1f, 0x33, 0xbe, 0x67, 0xa7, 0xdb, 0xc6, 0xcc, 0x30, 0xcd, 0xc9, 0xa0, 0xdf, 0xcd, 0x99,
	0xf1, 0x07, 0xeb, 0x6f, 0xcf, 0xd9, 0xd1, 0xd9, 0x5b, 0xd6, 0xd1, 0x55, 0x24, 0xa2, 0x59, 0x1e,
	0xfb, 0x3c, 0xaa, 0x9f, 0xc4, 0xbf, 0xc1, 0xc5, 0x35, 0xbe, 0x68, 0x97, 0x11, 0xdf, 0x7c, 0x07,
	0x00, 0x00, 0xff, 0xff, 0xd8, 0xed, 0x17, 0xb9, 0x43, 0x02, 0x00, 0x00,
}
