// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/ContractGetRecords.proto

package proto

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

// Get all the records for a smart contract instance, for any function call (or the constructor call) during the last 25 hours, for which a Record was requested.
type ContractGetRecordsQuery struct {
	Header               *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	ContractID           *ContractID  `protobuf:"bytes,2,opt,name=contractID,proto3" json:"contractID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ContractGetRecordsQuery) Reset()         { *m = ContractGetRecordsQuery{} }
func (m *ContractGetRecordsQuery) String() string { return proto.CompactTextString(m) }
func (*ContractGetRecordsQuery) ProtoMessage()    {}
func (*ContractGetRecordsQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7807842e4904992, []int{0}
}

func (m *ContractGetRecordsQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractGetRecordsQuery.Unmarshal(m, b)
}
func (m *ContractGetRecordsQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractGetRecordsQuery.Marshal(b, m, deterministic)
}
func (m *ContractGetRecordsQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractGetRecordsQuery.Merge(m, src)
}
func (m *ContractGetRecordsQuery) XXX_Size() int {
	return xxx_messageInfo_ContractGetRecordsQuery.Size(m)
}
func (m *ContractGetRecordsQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractGetRecordsQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ContractGetRecordsQuery proto.InternalMessageInfo

func (m *ContractGetRecordsQuery) GetHeader() *QueryHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ContractGetRecordsQuery) GetContractID() *ContractID {
	if m != nil {
		return m.ContractID
	}
	return nil
}

// Response when the client sends the node ContractGetRecordsQuery
type ContractGetRecordsResponse struct {
	Header               *ResponseHeader      `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	ContractID           *ContractID          `protobuf:"bytes,2,opt,name=contractID,proto3" json:"contractID,omitempty"`
	Records              []*TransactionRecord `protobuf:"bytes,3,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ContractGetRecordsResponse) Reset()         { *m = ContractGetRecordsResponse{} }
func (m *ContractGetRecordsResponse) String() string { return proto.CompactTextString(m) }
func (*ContractGetRecordsResponse) ProtoMessage()    {}
func (*ContractGetRecordsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7807842e4904992, []int{1}
}

func (m *ContractGetRecordsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractGetRecordsResponse.Unmarshal(m, b)
}
func (m *ContractGetRecordsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractGetRecordsResponse.Marshal(b, m, deterministic)
}
func (m *ContractGetRecordsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractGetRecordsResponse.Merge(m, src)
}
func (m *ContractGetRecordsResponse) XXX_Size() int {
	return xxx_messageInfo_ContractGetRecordsResponse.Size(m)
}
func (m *ContractGetRecordsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractGetRecordsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContractGetRecordsResponse proto.InternalMessageInfo

func (m *ContractGetRecordsResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ContractGetRecordsResponse) GetContractID() *ContractID {
	if m != nil {
		return m.ContractID
	}
	return nil
}

func (m *ContractGetRecordsResponse) GetRecords() []*TransactionRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

func init() {
	proto.RegisterType((*ContractGetRecordsQuery)(nil), "proto.ContractGetRecordsQuery")
	proto.RegisterType((*ContractGetRecordsResponse)(nil), "proto.ContractGetRecordsResponse")
}

func init() { proto.RegisterFile("proto/ContractGetRecords.proto", fileDescriptor_c7807842e4904992) }

var fileDescriptor_c7807842e4904992 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xa9, 0xc3, 0x09, 0xd9, 0xc9, 0x80, 0xae, 0x14, 0x94, 0xb1, 0x53, 0x11, 0x9a, 0x62,
	0xfd, 0x06, 0x9b, 0xe0, 0xbc, 0x69, 0xd8, 0xc9, 0x5b, 0x96, 0x3e, 0x9a, 0x22, 0x6b, 0xca, 0x4b,
	0x06, 0xee, 0x6b, 0xf9, 0x09, 0x85, 0xbc, 0x6c, 0x54, 0x7b, 0xf4, 0xf4, 0xa0, 0xbf, 0xdf, 0x7b,
	0xff, 0x7f, 0xc3, 0xee, 0x7b, 0xb4, 0xde, 0x96, 0x6b, 0xdb, 0x79, 0x54, 0xda, 0xbf, 0x80, 0x97,
	0xa0, 0x2d, 0xd6, 0x4e, 0x04, 0xc0, 0x2f, 0xc3, 0xc8, 0x6e, 0x49, 0x5b, 0x29, 0xd7, 0xea, 0xed,
	0xb1, 0x87, 0x88, 0xb3, 0x3b, 0xfa, 0xbe, 0x45, 0xd5, 0x39, 0xa5, 0x7d, 0x6b, 0x3b, 0x5a, 0x8f,
	0x78, 0x4e, 0xf8, 0xfd, 0x00, 0x78, 0xdc, 0x80, 0xaa, 0x01, 0x23, 0xc8, 0x08, 0x48, 0x70, 0xbd,
	0xed, 0x1c, 0x0c, 0xd9, 0xf2, 0x8b, 0xcd, 0xc7, 0x75, 0xc2, 0x09, 0xfe, 0xc0, 0xa6, 0x26, 0xa8,
	0x69, 0xb2, 0x48, 0xf2, 0x59, 0xc5, 0x69, 0x45, 0x0c, 0x02, 0x64, 0x34, 0xf8, 0x23, 0x63, 0x3a,
	0x9e, 0x79, 0x7d, 0x4e, 0x2f, 0x82, 0x7f, 0x1d, 0xfd, 0xf5, 0x19, 0xc8, 0x81, 0xb4, 0xfc, 0x4e,
	0x58, 0x36, 0x8e, 0x3e, 0x95, 0xe4, 0xc5, 0x9f, 0xf4, 0x9b, 0x78, 0xed, 0xf7, 0x5f, 0xfc, 0xa3,
	0x00, 0xaf, 0xd8, 0x15, 0x52, 0x68, 0x3a, 0x59, 0x4c, 0xf2, 0x59, 0x95, 0x46, 0x7f, 0xf4, 0xc0,
	0xf2, 0x24, 0xae, 0x36, 0x2c, 0xd3, 0x76, 0x2f, 0x0c, 0xd4, 0x80, 0x4a, 0x18, 0xe5, 0x4c, 0x83,
	0xaa, 0x37, 0xb4, 0xf8, 0x96, 0x7c, 0xe4, 0x4d, 0xeb, 0xcd, 0x61, 0x27, 0xb4, 0xdd, 0x97, 0x67,
	0x5a, 0x92, 0x5e, 0xb8, 0xfa, 0xb3, 0x68, 0x6c, 0x19, 0xdc, 0xdd, 0x34, 0x8c, 0xa7, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xf2, 0x49, 0x5f, 0x0d, 0x14, 0x02, 0x00, 0x00,
}
