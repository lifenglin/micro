// Code generated by protoc-gen-go. DO NOT EDIT.
// source: micro/micro/debug/log/proto/log.proto

package go_micro_debug_log

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Record struct {
	// timestamp of log record
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// record metadata
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// record value
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_23adf446d3f28816, []int{0}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Record) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Record) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReadRequest struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_23adf446d3f28816, []int{1}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *ReadRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type ReadResponse struct {
	Records              []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_23adf446d3f28816, []int{2}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

func init() {
	proto.RegisterType((*Record)(nil), "go.micro.debug.log.Record")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.debug.log.Record.MetadataEntry")
	proto.RegisterType((*ReadRequest)(nil), "go.micro.debug.log.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "go.micro.debug.log.ReadResponse")
}

func init() {
	proto.RegisterFile("micro/micro/debug/log/proto/log.proto", fileDescriptor_23adf446d3f28816)
}

var fileDescriptor_23adf446d3f28816 = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0x41, 0x4b, 0xf4, 0x30,
	0x14, 0xfc, 0xb2, 0xfd, 0xdc, 0xb5, 0x6f, 0x15, 0x24, 0x78, 0x08, 0x45, 0xb0, 0x14, 0x84, 0x9e,
	0x52, 0x58, 0x3d, 0x88, 0x9e, 0x84, 0xf5, 0xa4, 0x5e, 0xf2, 0x0f, 0xb2, 0xed, 0x23, 0x14, 0xdb,
	0xa6, 0x26, 0x69, 0x61, 0x7f, 0x9f, 0x7f, 0x4c, 0x9a, 0xb6, 0x2b, 0xa2, 0x7b, 0x09, 0x6f, 0x5e,
	0x66, 0x26, 0x33, 0x04, 0x6e, 0xea, 0x32, 0x37, 0x3a, 0x1b, 0xcf, 0x02, 0x77, 0x9d, 0xca, 0x2a,
	0xad, 0xb2, 0xd6, 0x68, 0xa7, 0x87, 0x89, 0xfb, 0x89, 0x52, 0xa5, 0xb9, 0xe7, 0x70, 0xcf, 0xe1,
	0x95, 0x56, 0xc9, 0x27, 0x81, 0xa5, 0xc0, 0x5c, 0x9b, 0x82, 0x5e, 0x41, 0xe8, 0xca, 0x1a, 0xad,
	0x93, 0x75, 0xcb, 0x48, 0x4c, 0xd2, 0x40, 0x7c, 0x2f, 0xe8, 0x16, 0x4e, 0x6b, 0x74, 0xb2, 0x90,
	0x4e, 0xb2, 0x45, 0x1c, 0xa4, 0xeb, 0x4d, 0xca, 0x7f, 0xfb, 0xf1, 0xd1, 0x8b, 0xbf, 0x4d, 0xd4,
	0xe7, 0xc6, 0x99, 0xbd, 0x38, 0x28, 0x29, 0x83, 0x55, 0x8d, 0xd6, 0x4a, 0x85, 0x2c, 0x88, 0x49,
	0x1a, 0x8a, 0x19, 0x46, 0x8f, 0x70, 0xfe, 0x43, 0x44, 0x2f, 0x20, 0x78, 0xc7, 0xbd, 0x0f, 0x12,
	0x8a, 0x61, 0xa4, 0x97, 0x70, 0xd2, 0xcb, 0xaa, 0x43, 0xb6, 0xf0, 0xbb, 0x11, 0x3c, 0x2c, 0xee,
	0x49, 0xf2, 0x04, 0x6b, 0x81, 0xb2, 0x10, 0xf8, 0xd1, 0xa1, 0x75, 0xc3, 0x2b, 0x16, 0x4d, 0x5f,
	0xe6, 0x38, 0xc9, 0x67, 0x38, 0xdc, 0xf4, 0x68, 0x6c, 0xa9, 0x9b, 0xc9, 0x64, 0x86, 0xc9, 0x16,
	0xce, 0x46, 0x0b, 0xdb, 0xea, 0xc6, 0x22, 0xbd, 0x83, 0x95, 0xf1, 0x5d, 0x2c, 0x23, 0xbe, 0x6e,
	0x74, 0xbc, 0xae, 0x98, 0xa9, 0x1b, 0x01, 0xc1, 0xab, 0x56, 0xf4, 0x05, 0xfe, 0x0f, 0x66, 0xf4,
	0xfa, 0x6f, 0xcd, 0x21, 0x69, 0x14, 0x1f, 0x27, 0x8c, 0x39, 0x92, 0x7f, 0xbb, 0xa5, 0xff, 0xbd,
	0xdb, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x86, 0xd3, 0x6b, 0x8e, 0xe6, 0x01, 0x00, 0x00,
}
