// Code generated by protoc-gen-go. DO NOT EDIT.
// source: openpitrix/types/confd.proto

package pbtypes // import "openpitrix.io/openpitrix/pkg/pb/types"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ConfdConfig struct {
	ProcessorConfig      *ConfdProcessorConfig `protobuf:"bytes,1,opt,name=processor_config,json=processorConfig,proto3" json:"processor_config"`
	BackendConfig        *ConfdBackendConfig   `protobuf:"bytes,2,opt,name=backend_config,json=backendConfig,proto3" json:"backend_config"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ConfdConfig) Reset()         { *m = ConfdConfig{} }
func (m *ConfdConfig) String() string { return proto.CompactTextString(m) }
func (*ConfdConfig) ProtoMessage()    {}
func (*ConfdConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_confd_35e27e7922cf6243, []int{0}
}
func (m *ConfdConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfdConfig.Unmarshal(m, b)
}
func (m *ConfdConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfdConfig.Marshal(b, m, deterministic)
}
func (dst *ConfdConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfdConfig.Merge(dst, src)
}
func (m *ConfdConfig) XXX_Size() int {
	return xxx_messageInfo_ConfdConfig.Size(m)
}
func (m *ConfdConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfdConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ConfdConfig proto.InternalMessageInfo

func (m *ConfdConfig) GetProcessorConfig() *ConfdProcessorConfig {
	if m != nil {
		return m.ProcessorConfig
	}
	return nil
}

func (m *ConfdConfig) GetBackendConfig() *ConfdBackendConfig {
	if m != nil {
		return m.BackendConfig
	}
	return nil
}

type ConfdEndpoint struct {
	FrontgateId          string   `protobuf:"bytes,1,opt,name=frontgate_id,json=frontgateId,proto3" json:"frontgate_id"`
	DroneIp              string   `protobuf:"bytes,2,opt,name=drone_ip,json=droneIp,proto3" json:"drone_ip"`
	DronePort            int32    `protobuf:"varint,3,opt,name=drone_port,json=dronePort,proto3" json:"drone_port"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfdEndpoint) Reset()         { *m = ConfdEndpoint{} }
func (m *ConfdEndpoint) String() string { return proto.CompactTextString(m) }
func (*ConfdEndpoint) ProtoMessage()    {}
func (*ConfdEndpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_confd_35e27e7922cf6243, []int{1}
}
func (m *ConfdEndpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfdEndpoint.Unmarshal(m, b)
}
func (m *ConfdEndpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfdEndpoint.Marshal(b, m, deterministic)
}
func (dst *ConfdEndpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfdEndpoint.Merge(dst, src)
}
func (m *ConfdEndpoint) XXX_Size() int {
	return xxx_messageInfo_ConfdEndpoint.Size(m)
}
func (m *ConfdEndpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfdEndpoint.DiscardUnknown(m)
}

var xxx_messageInfo_ConfdEndpoint proto.InternalMessageInfo

func (m *ConfdEndpoint) GetFrontgateId() string {
	if m != nil {
		return m.FrontgateId
	}
	return ""
}

func (m *ConfdEndpoint) GetDroneIp() string {
	if m != nil {
		return m.DroneIp
	}
	return ""
}

func (m *ConfdEndpoint) GetDronePort() int32 {
	if m != nil {
		return m.DronePort
	}
	return 0
}

// Keep same as libconfd.Config
// See https://godoc.org/openpitrix.io/libconfd#Config
type ConfdProcessorConfig struct {
	Confdir              string   `protobuf:"bytes,1,opt,name=confdir,proto3" json:"confdir"`
	Interval             int32    `protobuf:"varint,2,opt,name=interval,proto3" json:"interval"`
	Noop                 bool     `protobuf:"varint,3,opt,name=noop,proto3" json:"noop"`
	Prefix               string   `protobuf:"bytes,4,opt,name=prefix,proto3" json:"prefix"`
	SyncOnly             bool     `protobuf:"varint,5,opt,name=sync_only,json=syncOnly,proto3" json:"sync_only"`
	LogLevel             string   `protobuf:"bytes,6,opt,name=log_level,json=logLevel,proto3" json:"log_level"`
	Onetime              bool     `protobuf:"varint,7,opt,name=onetime,proto3" json:"onetime"`
	Watch                bool     `protobuf:"varint,8,opt,name=watch,proto3" json:"watch"`
	KeepStageFile        bool     `protobuf:"varint,9,opt,name=keep_stage_file,json=keepStageFile,proto3" json:"keep_stage_file"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfdProcessorConfig) Reset()         { *m = ConfdProcessorConfig{} }
func (m *ConfdProcessorConfig) String() string { return proto.CompactTextString(m) }
func (*ConfdProcessorConfig) ProtoMessage()    {}
func (*ConfdProcessorConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_confd_35e27e7922cf6243, []int{2}
}
func (m *ConfdProcessorConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfdProcessorConfig.Unmarshal(m, b)
}
func (m *ConfdProcessorConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfdProcessorConfig.Marshal(b, m, deterministic)
}
func (dst *ConfdProcessorConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfdProcessorConfig.Merge(dst, src)
}
func (m *ConfdProcessorConfig) XXX_Size() int {
	return xxx_messageInfo_ConfdProcessorConfig.Size(m)
}
func (m *ConfdProcessorConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfdProcessorConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ConfdProcessorConfig proto.InternalMessageInfo

func (m *ConfdProcessorConfig) GetConfdir() string {
	if m != nil {
		return m.Confdir
	}
	return ""
}

func (m *ConfdProcessorConfig) GetInterval() int32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

func (m *ConfdProcessorConfig) GetNoop() bool {
	if m != nil {
		return m.Noop
	}
	return false
}

func (m *ConfdProcessorConfig) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *ConfdProcessorConfig) GetSyncOnly() bool {
	if m != nil {
		return m.SyncOnly
	}
	return false
}

func (m *ConfdProcessorConfig) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

func (m *ConfdProcessorConfig) GetOnetime() bool {
	if m != nil {
		return m.Onetime
	}
	return false
}

func (m *ConfdProcessorConfig) GetWatch() bool {
	if m != nil {
		return m.Watch
	}
	return false
}

func (m *ConfdProcessorConfig) GetKeepStageFile() bool {
	if m != nil {
		return m.KeepStageFile
	}
	return false
}

// Keep same as libconfd.BackendConfig
// See https://godoc.org/openpitrix.io/libconfd#BackendConfig
type ConfdBackendConfig struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type"`
	Host                 []string `protobuf:"bytes,2,rep,name=host,proto3" json:"host"`
	User                 string   `protobuf:"bytes,3,opt,name=user,proto3" json:"user"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password"`
	ClientCaKeys         string   `protobuf:"bytes,5,opt,name=client_ca_keys,json=clientCaKeys,proto3" json:"client_ca_keys"`
	ClientCert           string   `protobuf:"bytes,6,opt,name=client_cert,json=clientCert,proto3" json:"client_cert"`
	ClientKey            string   `protobuf:"bytes,7,opt,name=client_key,json=clientKey,proto3" json:"client_key"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfdBackendConfig) Reset()         { *m = ConfdBackendConfig{} }
func (m *ConfdBackendConfig) String() string { return proto.CompactTextString(m) }
func (*ConfdBackendConfig) ProtoMessage()    {}
func (*ConfdBackendConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_confd_35e27e7922cf6243, []int{3}
}
func (m *ConfdBackendConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfdBackendConfig.Unmarshal(m, b)
}
func (m *ConfdBackendConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfdBackendConfig.Marshal(b, m, deterministic)
}
func (dst *ConfdBackendConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfdBackendConfig.Merge(dst, src)
}
func (m *ConfdBackendConfig) XXX_Size() int {
	return xxx_messageInfo_ConfdBackendConfig.Size(m)
}
func (m *ConfdBackendConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfdBackendConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ConfdBackendConfig proto.InternalMessageInfo

func (m *ConfdBackendConfig) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ConfdBackendConfig) GetHost() []string {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *ConfdBackendConfig) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ConfdBackendConfig) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *ConfdBackendConfig) GetClientCaKeys() string {
	if m != nil {
		return m.ClientCaKeys
	}
	return ""
}

func (m *ConfdBackendConfig) GetClientCert() string {
	if m != nil {
		return m.ClientCert
	}
	return ""
}

func (m *ConfdBackendConfig) GetClientKey() string {
	if m != nil {
		return m.ClientKey
	}
	return ""
}

type ConfdStatus struct {
	ProcessId            int32                `protobuf:"varint,1,opt,name=process_id,json=processId,proto3" json:"process_id"`
	UpTime               *timestamp.Timestamp `protobuf:"bytes,2,opt,name=up_time,json=upTime,proto3" json:"up_time"`
	Status               string               `protobuf:"bytes,3,opt,name=status,proto3" json:"status"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ConfdStatus) Reset()         { *m = ConfdStatus{} }
func (m *ConfdStatus) String() string { return proto.CompactTextString(m) }
func (*ConfdStatus) ProtoMessage()    {}
func (*ConfdStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_confd_35e27e7922cf6243, []int{4}
}
func (m *ConfdStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfdStatus.Unmarshal(m, b)
}
func (m *ConfdStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfdStatus.Marshal(b, m, deterministic)
}
func (dst *ConfdStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfdStatus.Merge(dst, src)
}
func (m *ConfdStatus) XXX_Size() int {
	return xxx_messageInfo_ConfdStatus.Size(m)
}
func (m *ConfdStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfdStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ConfdStatus proto.InternalMessageInfo

func (m *ConfdStatus) GetProcessId() int32 {
	if m != nil {
		return m.ProcessId
	}
	return 0
}

func (m *ConfdStatus) GetUpTime() *timestamp.Timestamp {
	if m != nil {
		return m.UpTime
	}
	return nil
}

func (m *ConfdStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type SetConfdConfigRequest struct {
	Endpoint             *ConfdEndpoint `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint"`
	Config               *ConfdConfig   `protobuf:"bytes,2,opt,name=config,proto3" json:"config"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *SetConfdConfigRequest) Reset()         { *m = SetConfdConfigRequest{} }
func (m *SetConfdConfigRequest) String() string { return proto.CompactTextString(m) }
func (*SetConfdConfigRequest) ProtoMessage()    {}
func (*SetConfdConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_confd_35e27e7922cf6243, []int{5}
}
func (m *SetConfdConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetConfdConfigRequest.Unmarshal(m, b)
}
func (m *SetConfdConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetConfdConfigRequest.Marshal(b, m, deterministic)
}
func (dst *SetConfdConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetConfdConfigRequest.Merge(dst, src)
}
func (m *SetConfdConfigRequest) XXX_Size() int {
	return xxx_messageInfo_SetConfdConfigRequest.Size(m)
}
func (m *SetConfdConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetConfdConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetConfdConfigRequest proto.InternalMessageInfo

func (m *SetConfdConfigRequest) GetEndpoint() *ConfdEndpoint {
	if m != nil {
		return m.Endpoint
	}
	return nil
}

func (m *SetConfdConfigRequest) GetConfig() *ConfdConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func init() {
	proto.RegisterType((*ConfdConfig)(nil), "openpitrix.types.ConfdConfig")
	proto.RegisterType((*ConfdEndpoint)(nil), "openpitrix.types.ConfdEndpoint")
	proto.RegisterType((*ConfdProcessorConfig)(nil), "openpitrix.types.ConfdProcessorConfig")
	proto.RegisterType((*ConfdBackendConfig)(nil), "openpitrix.types.ConfdBackendConfig")
	proto.RegisterType((*ConfdStatus)(nil), "openpitrix.types.ConfdStatus")
	proto.RegisterType((*SetConfdConfigRequest)(nil), "openpitrix.types.SetConfdConfigRequest")
}

func init() { proto.RegisterFile("openpitrix/types/confd.proto", fileDescriptor_confd_35e27e7922cf6243) }

var fileDescriptor_confd_35e27e7922cf6243 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x54, 0xdd, 0x4e, 0xdb, 0x30,
	0x14, 0x56, 0x81, 0xb6, 0xc9, 0x29, 0x7f, 0xb2, 0xd8, 0x94, 0xb1, 0x21, 0x58, 0x85, 0x10, 0x37,
	0x4b, 0xa4, 0xa1, 0x5d, 0x71, 0x07, 0xda, 0x24, 0xc4, 0xa4, 0xb1, 0xb0, 0xab, 0xdd, 0x44, 0x69,
	0x72, 0x1a, 0xac, 0x06, 0xdb, 0xb3, 0x5d, 0x20, 0xaf, 0xb0, 0x3d, 0xcb, 0x5e, 0x66, 0x4f, 0x34,
	0xf9, 0xc4, 0xe9, 0xca, 0x36, 0x6e, 0x5a, 0x7f, 0xdf, 0x39, 0xe7, 0x73, 0xce, 0x9f, 0xe1, 0x95,
	0x54, 0x28, 0x14, 0xb7, 0x9a, 0x3f, 0x24, 0xb6, 0x51, 0x68, 0x92, 0x42, 0x8a, 0x69, 0x19, 0x2b,
	0x2d, 0xad, 0x64, 0xdb, 0x7f, 0xac, 0x31, 0x59, 0x77, 0xf7, 0x2b, 0x29, 0xab, 0x1a, 0x13, 0xb2,
	0x4f, 0xe6, 0xd3, 0xc4, 0xf2, 0x5b, 0x34, 0x36, 0xbf, 0x55, 0x6d, 0xc8, 0xf8, 0x67, 0x0f, 0x46,
	0xe7, 0x4e, 0xc2, 0xfd, 0xf0, 0x8a, 0x7d, 0x86, 0x6d, 0xa5, 0x65, 0x81, 0xc6, 0x48, 0x9d, 0x15,
	0xc4, 0x45, 0xbd, 0x83, 0xde, 0xf1, 0xe8, 0xed, 0x51, 0xfc, 0xb7, 0x7a, 0x4c, 0x81, 0x57, 0x9d,
	0x7b, 0xab, 0x90, 0x6e, 0xa9, 0xc7, 0x04, 0xbb, 0x84, 0xcd, 0x49, 0x5e, 0xcc, 0x50, 0x94, 0x9d,
	0xe0, 0x0a, 0x09, 0x1e, 0x3e, 0x21, 0x78, 0xd6, 0x3a, 0x7b, 0xb9, 0x8d, 0xc9, 0x32, 0x1c, 0xd7,
	0xb0, 0x41, 0x4e, 0xef, 0x45, 0xa9, 0x24, 0x17, 0x96, 0xbd, 0x86, 0xf5, 0xa9, 0x96, 0xc2, 0x56,
	0xb9, 0xc5, 0x8c, 0x97, 0xf4, 0xb1, 0x61, 0x3a, 0x5a, 0x70, 0x17, 0x25, 0x7b, 0x01, 0x41, 0xa9,
	0xa5, 0xc0, 0x8c, 0x2b, 0xba, 0x3a, 0x4c, 0x87, 0x84, 0x2f, 0x14, 0xdb, 0x03, 0x68, 0x4d, 0x4a,
	0x6a, 0x1b, 0xad, 0x1e, 0xf4, 0x8e, 0xfb, 0x69, 0x48, 0xcc, 0x95, 0xd4, 0x76, 0xfc, 0x7d, 0x05,
	0x76, 0xfe, 0x97, 0x24, 0x8b, 0x60, 0x48, 0x85, 0xe7, 0xda, 0x5f, 0xd8, 0x41, 0xb6, 0x0b, 0x01,
	0x17, 0x16, 0xf5, 0x5d, 0x5e, 0xd3, 0x65, 0xfd, 0x74, 0x81, 0x19, 0x83, 0x35, 0x21, 0xa5, 0xa2,
	0x7b, 0x82, 0x94, 0xce, 0xec, 0x39, 0x0c, 0x94, 0xc6, 0x29, 0x7f, 0x88, 0xd6, 0x48, 0xc8, 0x23,
	0xf6, 0x12, 0x42, 0xd3, 0x88, 0x22, 0x93, 0xa2, 0x6e, 0xa2, 0x3e, 0x05, 0x04, 0x8e, 0xf8, 0x24,
	0xea, 0xc6, 0x19, 0x6b, 0x59, 0x65, 0x35, 0xde, 0x61, 0x1d, 0x0d, 0x28, 0x2e, 0xa8, 0x65, 0xf5,
	0xd1, 0x61, 0xf7, 0x6d, 0x52, 0xa0, 0x6b, 0x74, 0x34, 0xa4, 0xb8, 0x0e, 0xb2, 0x1d, 0xe8, 0xdf,
	0xe7, 0xb6, 0xb8, 0x89, 0x02, 0xe2, 0x5b, 0xc0, 0x8e, 0x60, 0x6b, 0x86, 0xa8, 0x32, 0x63, 0xf3,
	0x0a, 0xb3, 0x29, 0xaf, 0x31, 0x0a, 0xc9, 0xbe, 0xe1, 0xe8, 0x6b, 0xc7, 0x7e, 0xe0, 0x35, 0x8e,
	0x7f, 0xf5, 0x80, 0xfd, 0xdb, 0x20, 0x97, 0x94, 0x6b, 0x9e, 0xaf, 0x03, 0x9d, 0x1d, 0x77, 0x23,
	0x8d, 0x8d, 0x56, 0x0e, 0x56, 0x1d, 0xe7, 0xce, 0x8e, 0x9b, 0x1b, 0xd4, 0x94, 0x7c, 0x98, 0xd2,
	0xd9, 0x15, 0x4b, 0xe5, 0xc6, 0xdc, 0x4b, 0x5d, 0xfa, 0xf4, 0x17, 0x98, 0x1d, 0xc2, 0x66, 0x51,
	0x73, 0x14, 0x36, 0x2b, 0xf2, 0x6c, 0x86, 0x8d, 0xa1, 0x2a, 0x84, 0xe9, 0x7a, 0xcb, 0x9e, 0xe7,
	0x97, 0xd8, 0x18, 0xb6, 0x0f, 0xa3, 0xce, 0x0b, 0xb5, 0xf5, 0xb5, 0x00, 0xef, 0x82, 0xda, 0xba,
	0x0e, 0x7b, 0x87, 0x19, 0x36, 0x54, 0x90, 0x30, 0x0d, 0x5b, 0xe6, 0x12, 0x9b, 0x71, 0xe3, 0xc7,
	0xff, 0xda, 0xe6, 0x76, 0x6e, 0x9c, 0xb7, 0x1f, 0xdf, 0x6e, 0x96, 0xfa, 0x69, 0xe8, 0x99, 0x8b,
	0x92, 0x9d, 0xc0, 0x70, 0xae, 0x32, 0x2a, 0x6d, 0x3b, 0xc3, 0xbb, 0x71, 0xbb, 0x60, 0x71, 0xb7,
	0x60, 0xf1, 0x97, 0x6e, 0xc1, 0xd2, 0xc1, 0x5c, 0x39, 0xe0, 0x3a, 0x6c, 0x48, 0xdd, 0xa7, 0xee,
	0xd1, 0xf8, 0x47, 0x0f, 0x9e, 0x5d, 0xa3, 0x5d, 0xda, 0xbe, 0x14, 0xbf, 0xcd, 0xd1, 0x58, 0x76,
	0x0a, 0x01, 0xfa, 0xf9, 0xf6, 0xcb, 0xb7, 0xff, 0xc4, 0xae, 0x74, 0x6b, 0x90, 0x2e, 0x02, 0xd8,
	0x3b, 0x18, 0x3c, 0x5a, 0xb3, 0xbd, 0x27, 0x42, 0xfd, 0x95, 0xde, 0xf9, 0x2c, 0xf9, 0xfa, 0x66,
	0xc9, 0x8f, 0xcb, 0x64, 0xe9, 0xa5, 0x51, 0xb3, 0x2a, 0x51, 0x93, 0xf6, 0xc1, 0x39, 0x55, 0x13,
	0xfa, 0x9f, 0x0c, 0x28, 0xe5, 0x93, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x71, 0x3c, 0xe8, 0xa7,
	0x93, 0x04, 0x00, 0x00,
}
