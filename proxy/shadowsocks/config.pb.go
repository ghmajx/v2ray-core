// Code generated by protoc-gen-go.
// source: v2ray.com/core/proxy/shadowsocks/config.proto
// DO NOT EDIT!

/*
Package shadowsocks is a generated protocol buffer package.

It is generated from these files:
	v2ray.com/core/proxy/shadowsocks/config.proto

It has these top-level messages:
	Account
	ServerConfig
	ClientConfig
*/
package shadowsocks

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_common_protocol "v2ray.com/core/common/protocol"
import v2ray_core_common_protocol1 "v2ray.com/core/common/protocol"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CipherType int32

const (
	CipherType_UNKNOWN       CipherType = 0
	CipherType_AES_128_CFB   CipherType = 1
	CipherType_AES_256_CFB   CipherType = 2
	CipherType_CHACHA20      CipherType = 3
	CipherType_CHACHA20_IEFT CipherType = 4
)

var CipherType_name = map[int32]string{
	0: "UNKNOWN",
	1: "AES_128_CFB",
	2: "AES_256_CFB",
	3: "CHACHA20",
	4: "CHACHA20_IEFT",
}
var CipherType_value = map[string]int32{
	"UNKNOWN":       0,
	"AES_128_CFB":   1,
	"AES_256_CFB":   2,
	"CHACHA20":      3,
	"CHACHA20_IEFT": 4,
}

func (x CipherType) String() string {
	return proto.EnumName(CipherType_name, int32(x))
}
func (CipherType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Account_OneTimeAuth int32

const (
	Account_Auto     Account_OneTimeAuth = 0
	Account_Disabled Account_OneTimeAuth = 1
	Account_Enabled  Account_OneTimeAuth = 2
)

var Account_OneTimeAuth_name = map[int32]string{
	0: "Auto",
	1: "Disabled",
	2: "Enabled",
}
var Account_OneTimeAuth_value = map[string]int32{
	"Auto":     0,
	"Disabled": 1,
	"Enabled":  2,
}

func (x Account_OneTimeAuth) String() string {
	return proto.EnumName(Account_OneTimeAuth_name, int32(x))
}
func (Account_OneTimeAuth) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type Account struct {
	Password   string              `protobuf:"bytes,1,opt,name=password" json:"password,omitempty"`
	CipherType CipherType          `protobuf:"varint,2,opt,name=cipher_type,json=cipherType,enum=v2ray.core.proxy.shadowsocks.CipherType" json:"cipher_type,omitempty"`
	Ota        Account_OneTimeAuth `protobuf:"varint,3,opt,name=ota,enum=v2ray.core.proxy.shadowsocks.Account_OneTimeAuth" json:"ota,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ServerConfig struct {
	UdpEnabled bool                             `protobuf:"varint,1,opt,name=udp_enabled,json=udpEnabled" json:"udp_enabled,omitempty"`
	User       *v2ray_core_common_protocol.User `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
}

func (m *ServerConfig) Reset()                    { *m = ServerConfig{} }
func (m *ServerConfig) String() string            { return proto.CompactTextString(m) }
func (*ServerConfig) ProtoMessage()               {}
func (*ServerConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ServerConfig) GetUser() *v2ray_core_common_protocol.User {
	if m != nil {
		return m.User
	}
	return nil
}

type ClientConfig struct {
	Server []*v2ray_core_common_protocol1.ServerEndpoint `protobuf:"bytes,1,rep,name=server" json:"server,omitempty"`
}

func (m *ClientConfig) Reset()                    { *m = ClientConfig{} }
func (m *ClientConfig) String() string            { return proto.CompactTextString(m) }
func (*ClientConfig) ProtoMessage()               {}
func (*ClientConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ClientConfig) GetServer() []*v2ray_core_common_protocol1.ServerEndpoint {
	if m != nil {
		return m.Server
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "v2ray.core.proxy.shadowsocks.Account")
	proto.RegisterType((*ServerConfig)(nil), "v2ray.core.proxy.shadowsocks.ServerConfig")
	proto.RegisterType((*ClientConfig)(nil), "v2ray.core.proxy.shadowsocks.ClientConfig")
	proto.RegisterEnum("v2ray.core.proxy.shadowsocks.CipherType", CipherType_name, CipherType_value)
	proto.RegisterEnum("v2ray.core.proxy.shadowsocks.Account_OneTimeAuth", Account_OneTimeAuth_name, Account_OneTimeAuth_value)
}

func init() { proto.RegisterFile("v2ray.com/core/proxy/shadowsocks/config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 431 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0xeb, 0x24, 0x6a, 0xc3, 0x6c, 0x00, 0xb3, 0xa7, 0x28, 0x42, 0xc2, 0xca, 0x29, 0x54,
	0x62, 0xdd, 0x9a, 0x3f, 0xe2, 0xc0, 0xc5, 0x31, 0xa9, 0x5a, 0x21, 0xa5, 0xc8, 0x4d, 0x85, 0x84,
	0x90, 0x2c, 0x77, 0x3d, 0x10, 0x8b, 0xd8, 0xb3, 0xda, 0xb5, 0x5b, 0xf2, 0x91, 0xf9, 0x16, 0xc8,
	0xeb, 0x24, 0x44, 0x1c, 0xc2, 0xcd, 0x33, 0x7e, 0xef, 0xf9, 0xcd, 0xcf, 0xf0, 0xea, 0x3e, 0xd0,
	0xe9, 0x5a, 0x48, 0x2a, 0x7c, 0x49, 0x1a, 0x7d, 0xa5, 0xe9, 0xd7, 0xda, 0x37, 0xcb, 0x34, 0xa3,
	0x07, 0x43, 0xf2, 0xa7, 0xf1, 0x25, 0x95, 0xdf, 0xf3, 0x1f, 0x42, 0x69, 0xaa, 0x88, 0x3f, 0xdf,
	0xca, 0x35, 0x0a, 0x2b, 0x15, 0x7b, 0xd2, 0xd1, 0xcb, 0x7f, 0xc2, 0x24, 0x15, 0x05, 0x95, 0xbe,
	0xb5, 0x4a, 0x5a, 0xf9, 0xb5, 0x41, 0xdd, 0x06, 0x8d, 0xce, 0xfe, 0x23, 0x35, 0xa8, 0xef, 0x51,
	0x27, 0x46, 0xa1, 0x6c, 0x1d, 0xe3, 0xdf, 0x0e, 0x9c, 0x84, 0x52, 0x52, 0x5d, 0x56, 0x7c, 0x04,
	0x7d, 0x95, 0x1a, 0xf3, 0x40, 0x3a, 0x1b, 0x3a, 0x9e, 0x33, 0x79, 0x14, 0xef, 0x66, 0x7e, 0x05,
	0x4c, 0xe6, 0x6a, 0x89, 0x3a, 0xa9, 0xd6, 0x0a, 0x87, 0x1d, 0xcf, 0x99, 0x3c, 0x09, 0x26, 0xe2,
	0x50, 0x71, 0x11, 0x59, 0xc3, 0x62, 0xad, 0x30, 0x06, 0xb9, 0x7b, 0xe6, 0x11, 0x74, 0xa9, 0x4a,
	0x87, 0x5d, 0x1b, 0x71, 0x7e, 0x38, 0x62, 0x53, 0x4d, 0x5c, 0x97, 0xb8, 0xc8, 0x0b, 0x0c, 0xeb,
	0x6a, 0x19, 0x37, 0xee, 0x71, 0x00, 0x6c, 0x6f, 0xc7, 0xfb, 0xd0, 0x0b, 0xeb, 0x8a, 0xdc, 0x23,
	0x3e, 0x80, 0xfe, 0xc7, 0xdc, 0xa4, 0x77, 0x2b, 0xcc, 0x5c, 0x87, 0x33, 0x38, 0x99, 0x95, 0xed,
	0xd0, 0x19, 0x23, 0x0c, 0x6e, 0x2c, 0x80, 0xc8, 0xc2, 0xe7, 0x2f, 0x80, 0xd5, 0x99, 0x4a, 0xb0,
	0x15, 0xd8, 0x93, 0xfb, 0x31, 0xd4, 0x99, 0xda, 0x58, 0xf8, 0x1b, 0xe8, 0x35, 0x70, 0xed, 0xb5,
	0x2c, 0xf0, 0xf6, 0xab, 0xb6, 0x64, 0xc5, 0x96, 0xac, 0xb8, 0x35, 0xa8, 0x63, 0xab, 0x1e, 0xc7,
	0x30, 0x88, 0x56, 0x39, 0x96, 0xd5, 0xe6, 0x33, 0x53, 0x38, 0x6e, 0xb9, 0x0f, 0x1d, 0xaf, 0x3b,
	0x61, 0xc1, 0xe9, 0xa1, 0x9c, 0xb6, 0xe0, 0xac, 0xcc, 0x14, 0xe5, 0x65, 0x15, 0x6f, 0x9c, 0xa7,
	0xdf, 0x00, 0xfe, 0xd2, 0x6c, 0xae, 0xba, 0x9d, 0x7f, 0x9a, 0x5f, 0x7f, 0x99, 0xbb, 0x47, 0xfc,
	0x29, 0xb0, 0x70, 0x76, 0x93, 0x9c, 0x07, 0xef, 0x93, 0xe8, 0x62, 0xea, 0x3a, 0xdb, 0x45, 0xf0,
	0xf6, 0x9d, 0x5d, 0x74, 0x1a, 0x24, 0xd1, 0x65, 0x18, 0x5d, 0x86, 0xc1, 0x99, 0xdb, 0xe5, 0xcf,
	0xe0, 0xf1, 0x76, 0x4a, 0xae, 0x66, 0x17, 0x0b, 0xb7, 0x37, 0xfd, 0x00, 0x9e, 0xa4, 0xe2, 0xe0,
	0x9f, 0x98, 0xb2, 0xf6, 0x9a, 0xcf, 0x4d, 0xd1, 0xaf, 0x6c, 0xef, 0xcd, 0xdd, 0xb1, 0x2d, 0xff,
	0xfa, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x59, 0xf6, 0xcd, 0xed, 0xf5, 0x02, 0x00, 0x00,
}