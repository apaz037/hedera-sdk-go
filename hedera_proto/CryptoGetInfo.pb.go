// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CryptoGetInfo.proto

package hedera_proto

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

// Get all the information about an account, including the balance. This does not get the list of account records.
type CryptoGetInfoQuery struct {
	Header               *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	AccountID            *AccountID   `protobuf:"bytes,2,opt,name=accountID,proto3" json:"accountID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CryptoGetInfoQuery) Reset()         { *m = CryptoGetInfoQuery{} }
func (m *CryptoGetInfoQuery) String() string { return proto.CompactTextString(m) }
func (*CryptoGetInfoQuery) ProtoMessage()    {}
func (*CryptoGetInfoQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb659bcbcf0fa194, []int{0}
}

func (m *CryptoGetInfoQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoGetInfoQuery.Unmarshal(m, b)
}
func (m *CryptoGetInfoQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoGetInfoQuery.Marshal(b, m, deterministic)
}
func (m *CryptoGetInfoQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoGetInfoQuery.Merge(m, src)
}
func (m *CryptoGetInfoQuery) XXX_Size() int {
	return xxx_messageInfo_CryptoGetInfoQuery.Size(m)
}
func (m *CryptoGetInfoQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoGetInfoQuery.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoGetInfoQuery proto.InternalMessageInfo

func (m *CryptoGetInfoQuery) GetHeader() *QueryHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CryptoGetInfoQuery) GetAccountID() *AccountID {
	if m != nil {
		return m.AccountID
	}
	return nil
}

// Response when the client sends the node CryptoGetInfoQuery
type CryptoGetInfoResponse struct {
	Header               *ResponseHeader                    `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	AccountInfo          *CryptoGetInfoResponse_AccountInfo `protobuf:"bytes,2,opt,name=accountInfo,proto3" json:"accountInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *CryptoGetInfoResponse) Reset()         { *m = CryptoGetInfoResponse{} }
func (m *CryptoGetInfoResponse) String() string { return proto.CompactTextString(m) }
func (*CryptoGetInfoResponse) ProtoMessage()    {}
func (*CryptoGetInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb659bcbcf0fa194, []int{1}
}

func (m *CryptoGetInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoGetInfoResponse.Unmarshal(m, b)
}
func (m *CryptoGetInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoGetInfoResponse.Marshal(b, m, deterministic)
}
func (m *CryptoGetInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoGetInfoResponse.Merge(m, src)
}
func (m *CryptoGetInfoResponse) XXX_Size() int {
	return xxx_messageInfo_CryptoGetInfoResponse.Size(m)
}
func (m *CryptoGetInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoGetInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoGetInfoResponse proto.InternalMessageInfo

func (m *CryptoGetInfoResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CryptoGetInfoResponse) GetAccountInfo() *CryptoGetInfoResponse_AccountInfo {
	if m != nil {
		return m.AccountInfo
	}
	return nil
}

type CryptoGetInfoResponse_AccountInfo struct {
	AccountID                      *AccountID `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	ContractAccountID              string     `protobuf:"bytes,2,opt,name=contractAccountID,proto3" json:"contractAccountID,omitempty"`
	Deleted                        bool       `protobuf:"varint,3,opt,name=deleted,proto3" json:"deleted,omitempty"`
	ProxyAccountID                 *AccountID `protobuf:"bytes,4,opt,name=proxyAccountID,proto3" json:"proxyAccountID,omitempty"`
	ProxyReceived                  int64      `protobuf:"varint,6,opt,name=proxyReceived,proto3" json:"proxyReceived,omitempty"`
	Key                            *Key       `protobuf:"bytes,7,opt,name=key,proto3" json:"key,omitempty"`
	Balance                        uint64     `protobuf:"varint,8,opt,name=balance,proto3" json:"balance,omitempty"`
	GenerateSendRecordThreshold    uint64     `protobuf:"varint,9,opt,name=generateSendRecordThreshold,proto3" json:"generateSendRecordThreshold,omitempty"`
	GenerateReceiveRecordThreshold uint64     `protobuf:"varint,10,opt,name=generateReceiveRecordThreshold,proto3" json:"generateReceiveRecordThreshold,omitempty"`
	ReceiverSigRequired            bool       `protobuf:"varint,11,opt,name=receiverSigRequired,proto3" json:"receiverSigRequired,omitempty"`
	ExpirationTime                 *Timestamp `protobuf:"bytes,12,opt,name=expirationTime,proto3" json:"expirationTime,omitempty"`
	AutoRenewPeriod                *Duration  `protobuf:"bytes,13,opt,name=autoRenewPeriod,proto3" json:"autoRenewPeriod,omitempty"`
	Claims                         []*Claim   `protobuf:"bytes,14,rep,name=claims,proto3" json:"claims,omitempty"`
	XXX_NoUnkeyedLiteral           struct{}   `json:"-"`
	XXX_unrecognized               []byte     `json:"-"`
	XXX_sizecache                  int32      `json:"-"`
}

func (m *CryptoGetInfoResponse_AccountInfo) Reset()         { *m = CryptoGetInfoResponse_AccountInfo{} }
func (m *CryptoGetInfoResponse_AccountInfo) String() string { return proto.CompactTextString(m) }
func (*CryptoGetInfoResponse_AccountInfo) ProtoMessage()    {}
func (*CryptoGetInfoResponse_AccountInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb659bcbcf0fa194, []int{1, 0}
}

func (m *CryptoGetInfoResponse_AccountInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoGetInfoResponse_AccountInfo.Unmarshal(m, b)
}
func (m *CryptoGetInfoResponse_AccountInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoGetInfoResponse_AccountInfo.Marshal(b, m, deterministic)
}
func (m *CryptoGetInfoResponse_AccountInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoGetInfoResponse_AccountInfo.Merge(m, src)
}
func (m *CryptoGetInfoResponse_AccountInfo) XXX_Size() int {
	return xxx_messageInfo_CryptoGetInfoResponse_AccountInfo.Size(m)
}
func (m *CryptoGetInfoResponse_AccountInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoGetInfoResponse_AccountInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoGetInfoResponse_AccountInfo proto.InternalMessageInfo

func (m *CryptoGetInfoResponse_AccountInfo) GetAccountID() *AccountID {
	if m != nil {
		return m.AccountID
	}
	return nil
}

func (m *CryptoGetInfoResponse_AccountInfo) GetContractAccountID() string {
	if m != nil {
		return m.ContractAccountID
	}
	return ""
}

func (m *CryptoGetInfoResponse_AccountInfo) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

func (m *CryptoGetInfoResponse_AccountInfo) GetProxyAccountID() *AccountID {
	if m != nil {
		return m.ProxyAccountID
	}
	return nil
}

func (m *CryptoGetInfoResponse_AccountInfo) GetProxyReceived() int64 {
	if m != nil {
		return m.ProxyReceived
	}
	return 0
}

func (m *CryptoGetInfoResponse_AccountInfo) GetKey() *Key {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *CryptoGetInfoResponse_AccountInfo) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *CryptoGetInfoResponse_AccountInfo) GetGenerateSendRecordThreshold() uint64 {
	if m != nil {
		return m.GenerateSendRecordThreshold
	}
	return 0
}

func (m *CryptoGetInfoResponse_AccountInfo) GetGenerateReceiveRecordThreshold() uint64 {
	if m != nil {
		return m.GenerateReceiveRecordThreshold
	}
	return 0
}

func (m *CryptoGetInfoResponse_AccountInfo) GetReceiverSigRequired() bool {
	if m != nil {
		return m.ReceiverSigRequired
	}
	return false
}

func (m *CryptoGetInfoResponse_AccountInfo) GetExpirationTime() *Timestamp {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

func (m *CryptoGetInfoResponse_AccountInfo) GetAutoRenewPeriod() *Duration {
	if m != nil {
		return m.AutoRenewPeriod
	}
	return nil
}

func (m *CryptoGetInfoResponse_AccountInfo) GetClaims() []*Claim {
	if m != nil {
		return m.Claims
	}
	return nil
}

func init() {
	proto.RegisterType((*CryptoGetInfoQuery)(nil), "proto.CryptoGetInfoQuery")
	proto.RegisterType((*CryptoGetInfoResponse)(nil), "proto.CryptoGetInfoResponse")
	proto.RegisterType((*CryptoGetInfoResponse_AccountInfo)(nil), "proto.CryptoGetInfoResponse.AccountInfo")
}

func init() { proto.RegisterFile("CryptoGetInfo.proto", fileDescriptor_bb659bcbcf0fa194) }

var fileDescriptor_bb659bcbcf0fa194 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0xc1, 0x6e, 0x13, 0x3f,
	0x10, 0xc6, 0xb5, 0xff, 0xe4, 0x9f, 0x36, 0x93, 0x34, 0x69, 0x5d, 0x2a, 0x59, 0x01, 0xa1, 0x28,
	0xea, 0x61, 0x85, 0x20, 0x42, 0x85, 0x03, 0xdc, 0x48, 0x5b, 0x01, 0x85, 0x4b, 0x71, 0x73, 0xe2,
	0x82, 0xdc, 0xf5, 0x34, 0x6b, 0x48, 0x6c, 0xe3, 0x75, 0x4a, 0xf7, 0x11, 0x78, 0x5f, 0x1e, 0x00,
	0xad, 0xd7, 0x9b, 0x66, 0x43, 0xa1, 0xa7, 0x95, 0xbf, 0xf9, 0xcd, 0xf7, 0x8d, 0xad, 0x1d, 0xd8,
	0x3f, 0xb1, 0xb9, 0x71, 0xfa, 0x1d, 0xba, 0x33, 0x75, 0xa5, 0xc7, 0xc6, 0x6a, 0xa7, 0xc9, 0xff,
	0xfe, 0x33, 0xe8, 0x4f, 0xe5, 0x02, 0x33, 0xc7, 0x17, 0xa6, 0xd4, 0x07, 0xbd, 0xd3, 0xa5, 0xe5,
	0x4e, 0x6a, 0x15, 0xce, 0xbb, 0xc7, 0x3c, 0x93, 0xc9, 0x34, 0x37, 0x98, 0x05, 0x65, 0xef, 0xd3,
	0x12, 0x6d, 0xfe, 0x1e, 0xb9, 0x40, 0x1b, 0xa4, 0x07, 0x0c, 0x33, 0xa3, 0x55, 0x86, 0x75, 0xb5,
	0xcc, 0x9d, 0x08, 0x71, 0x32, 0xe7, 0x72, 0x51, 0xaa, 0x23, 0x03, 0xa4, 0x36, 0x8f, 0x77, 0x23,
	0x4f, 0xa0, 0x95, 0xfa, 0x5e, 0x1a, 0x0d, 0xa3, 0xb8, 0x73, 0x44, 0x4a, 0x7a, 0xbc, 0x96, 0xc5,
	0x02, 0x41, 0xc6, 0xd0, 0xe6, 0x49, 0xa2, 0x97, 0xca, 0x9d, 0x9d, 0xd2, 0xff, 0x3c, 0xbe, 0x1b,
	0xf0, 0x49, 0xa5, 0xb3, 0x5b, 0x64, 0xf4, 0xb3, 0x05, 0x07, 0xb5, 0xc8, 0x6a, 0x5a, 0xf2, 0x6c,
	0x23, 0xf5, 0x20, 0xd8, 0xd4, 0xaf, 0xb3, 0x0a, 0xfe, 0x00, 0x9d, 0xca, 0x55, 0x5d, 0xe9, 0x10,
	0x1d, 0x87, 0x9e, 0x3b, 0x13, 0x56, 0x03, 0x15, 0xda, 0x7a, 0xf3, 0xe0, 0x57, 0x13, 0x3a, 0x6b,
	0xc5, 0xfa, 0xa5, 0xa2, 0x7b, 0x2f, 0x45, 0x9e, 0xc2, 0x5e, 0xa2, 0x95, 0xb3, 0x3c, 0x71, 0x93,
	0xda, 0x63, 0xb4, 0xd9, 0x9f, 0x05, 0x42, 0x61, 0x4b, 0xe0, 0x1c, 0x1d, 0x0a, 0xda, 0x18, 0x46,
	0xf1, 0x36, 0xab, 0x8e, 0xe4, 0x15, 0xf4, 0x8c, 0xd5, 0x37, 0xf9, 0xad, 0x49, 0xf3, 0x2f, 0xe1,
	0x1b, 0x1c, 0x39, 0x84, 0x1d, 0xaf, 0x30, 0x4c, 0x50, 0x5e, 0xa3, 0xa0, 0xad, 0x61, 0x14, 0x37,
	0x58, 0x5d, 0x24, 0x8f, 0xa0, 0xf1, 0x0d, 0x73, 0xba, 0xe5, 0x4d, 0x21, 0x98, 0x7e, 0xc4, 0x9c,
	0x15, 0x72, 0x31, 0xd7, 0x25, 0x9f, 0x73, 0x95, 0x20, 0xdd, 0x1e, 0x46, 0x71, 0x93, 0x55, 0x47,
	0xf2, 0x06, 0x1e, 0xce, 0x50, 0xa1, 0xe5, 0x0e, 0x2f, 0x50, 0x09, 0x86, 0x89, 0xb6, 0x62, 0x9a,
	0x5a, 0xcc, 0x52, 0x3d, 0x17, 0xb4, 0xed, 0xe9, 0x7f, 0x21, 0xe4, 0x2d, 0x3c, 0xae, 0xca, 0x61,
	0x9a, 0x4d, 0x13, 0xf0, 0x26, 0xf7, 0x50, 0xe4, 0x39, 0xec, 0xdb, 0xb2, 0x62, 0x2f, 0xe4, 0x8c,
	0xe1, 0xf7, 0xa5, 0xb4, 0x28, 0x68, 0xc7, 0xbf, 0xe3, 0x5d, 0xa5, 0xe2, 0x4d, 0xf1, 0xc6, 0xc8,
	0x72, 0x8f, 0x8a, 0x05, 0xa3, 0xdd, 0xda, 0x9b, 0xae, 0x76, 0x8e, 0x6d, 0x70, 0xe4, 0x35, 0xf4,
	0xf9, 0xd2, 0x69, 0x86, 0x0a, 0x7f, 0x9c, 0xa3, 0x95, 0x5a, 0xd0, 0x1d, 0xdf, 0xda, 0x0f, 0xad,
	0xd5, 0x76, 0xb2, 0x4d, 0x8e, 0x1c, 0x42, 0x2b, 0x29, 0xd6, 0x2c, 0xa3, 0xbd, 0x61, 0x23, 0xee,
	0x1c, 0x75, 0xab, 0xff, 0xb2, 0x10, 0x59, 0xa8, 0x1d, 0xbf, 0x84, 0x51, 0xa2, 0x17, 0xe3, 0x14,
	0x05, 0x5a, 0x9e, 0xf2, 0x2c, 0x9d, 0x59, 0x6e, 0xd2, 0x31, 0x37, 0x32, 0xe0, 0x5f, 0xf9, 0x35,
	0x3f, 0x8f, 0x3e, 0x77, 0x4b, 0xe2, 0x8b, 0x17, 0x2f, 0x5b, 0xfe, 0xf3, 0xe2, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x20, 0x9b, 0x1f, 0xb3, 0x4a, 0x04, 0x00, 0x00,
}