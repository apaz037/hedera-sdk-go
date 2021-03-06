// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/CryptoTransfer.proto

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

// An account, and the amount that it sends or receives during a cryptocurrency transfer.
type AccountAmount struct {
	AccountID            *AccountID `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"`
	Amount               int64      `protobuf:"zigzag64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *AccountAmount) Reset()         { *m = AccountAmount{} }
func (m *AccountAmount) String() string { return proto.CompactTextString(m) }
func (*AccountAmount) ProtoMessage()    {}
func (*AccountAmount) Descriptor() ([]byte, []int) {
	return fileDescriptor_645768f75ce75fc6, []int{0}
}

func (m *AccountAmount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountAmount.Unmarshal(m, b)
}
func (m *AccountAmount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountAmount.Marshal(b, m, deterministic)
}
func (m *AccountAmount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountAmount.Merge(m, src)
}
func (m *AccountAmount) XXX_Size() int {
	return xxx_messageInfo_AccountAmount.Size(m)
}
func (m *AccountAmount) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountAmount.DiscardUnknown(m)
}

var xxx_messageInfo_AccountAmount proto.InternalMessageInfo

func (m *AccountAmount) GetAccountID() *AccountID {
	if m != nil {
		return m.AccountID
	}
	return nil
}

func (m *AccountAmount) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

// A list of accounts and amounts to transfer out of each account (negative) or into it (positive).
type TransferList struct {
	AccountAmounts       []*AccountAmount `protobuf:"bytes,1,rep,name=accountAmounts,proto3" json:"accountAmounts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TransferList) Reset()         { *m = TransferList{} }
func (m *TransferList) String() string { return proto.CompactTextString(m) }
func (*TransferList) ProtoMessage()    {}
func (*TransferList) Descriptor() ([]byte, []int) {
	return fileDescriptor_645768f75ce75fc6, []int{1}
}

func (m *TransferList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferList.Unmarshal(m, b)
}
func (m *TransferList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferList.Marshal(b, m, deterministic)
}
func (m *TransferList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferList.Merge(m, src)
}
func (m *TransferList) XXX_Size() int {
	return xxx_messageInfo_TransferList.Size(m)
}
func (m *TransferList) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferList.DiscardUnknown(m)
}

var xxx_messageInfo_TransferList proto.InternalMessageInfo

func (m *TransferList) GetAccountAmounts() []*AccountAmount {
	if m != nil {
		return m.AccountAmounts
	}
	return nil
}

// Transfer cryptocurrency from some accounts to other accounts. The accounts list can contain up to 10 accounts. The amounts list must be the same length as the accounts list. Each negative amount is withdrawn from the corresponding account (a sender), and each positive one is added to the corresponding account (a receiver). The amounts list must sum to zero. Each amount is a number of tinyBars (there are 100,000,000 tinyBars in one Hbar). If any sender account fails to have sufficient hbars to do the withdrawal, then the entire transaction fails, and none of those transfers occur, though the transaction fee is still charged. This transaction must be signed by the keys for all the sending accounts, and for any receiving accounts that have receiverSigRequired == true. The signatures are in the same order as the accounts, skipping those accounts that don't need a signature.
type CryptoTransferTransactionBody struct {
	Transfers            *TransferList `protobuf:"bytes,1,opt,name=transfers,proto3" json:"transfers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *CryptoTransferTransactionBody) Reset()         { *m = CryptoTransferTransactionBody{} }
func (m *CryptoTransferTransactionBody) String() string { return proto.CompactTextString(m) }
func (*CryptoTransferTransactionBody) ProtoMessage()    {}
func (*CryptoTransferTransactionBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_645768f75ce75fc6, []int{2}
}

func (m *CryptoTransferTransactionBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptoTransferTransactionBody.Unmarshal(m, b)
}
func (m *CryptoTransferTransactionBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptoTransferTransactionBody.Marshal(b, m, deterministic)
}
func (m *CryptoTransferTransactionBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptoTransferTransactionBody.Merge(m, src)
}
func (m *CryptoTransferTransactionBody) XXX_Size() int {
	return xxx_messageInfo_CryptoTransferTransactionBody.Size(m)
}
func (m *CryptoTransferTransactionBody) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptoTransferTransactionBody.DiscardUnknown(m)
}

var xxx_messageInfo_CryptoTransferTransactionBody proto.InternalMessageInfo

func (m *CryptoTransferTransactionBody) GetTransfers() *TransferList {
	if m != nil {
		return m.Transfers
	}
	return nil
}

func init() {
	proto.RegisterType((*AccountAmount)(nil), "proto.AccountAmount")
	proto.RegisterType((*TransferList)(nil), "proto.TransferList")
	proto.RegisterType((*CryptoTransferTransactionBody)(nil), "proto.CryptoTransferTransactionBody")
}

func init() { proto.RegisterFile("proto/CryptoTransfer.proto", fileDescriptor_645768f75ce75fc6) }

var fileDescriptor_645768f75ce75fc6 = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0x89, 0xe2, 0x60, 0x6f, 0x2a, 0x12, 0x65, 0x8c, 0x82, 0x50, 0x7a, 0xea, 0x65, 0x29,
	0xce, 0xab, 0x97, 0x56, 0x0f, 0x0a, 0x3b, 0x48, 0x18, 0x08, 0xde, 0xde, 0xd2, 0xda, 0x14, 0x69,
	0x5f, 0x49, 0xb2, 0x43, 0xff, 0x7b, 0x31, 0xe9, 0xdc, 0xe6, 0x25, 0x1f, 0xf9, 0xbe, 0xdf, 0x7b,
	0xf9, 0x02, 0x51, 0x6f, 0xc8, 0x51, 0xf6, 0x6c, 0x86, 0xde, 0xd1, 0xc6, 0x60, 0x67, 0xbf, 0x2a,
	0x23, 0xbc, 0xc9, 0x2f, 0xbc, 0x44, 0xf3, 0x80, 0x14, 0x68, 0x1b, 0xb5, 0x19, 0xfa, 0xca, 0x86,
	0x38, 0xf9, 0x80, 0xab, 0x5c, 0x29, 0xda, 0x75, 0x2e, 0x6f, 0x7f, 0x4f, 0x2e, 0x60, 0x8a, 0xc1,
	0x78, 0x7b, 0x59, 0xb0, 0x98, 0xa5, 0xb3, 0xd5, 0x4d, 0x60, 0x45, 0xbe, 0xf7, 0xe5, 0x01, 0xe1,
	0x73, 0x98, 0xa0, 0x9f, 0x5c, 0x9c, 0xc5, 0x2c, 0xe5, 0x72, 0xbc, 0x25, 0x6b, 0xb8, 0xdc, 0x37,
	0x59, 0x37, 0xd6, 0xf1, 0x27, 0xb8, 0xc6, 0xe3, 0x87, 0xec, 0x82, 0xc5, 0xe7, 0xe9, 0x6c, 0x75,
	0x77, 0xba, 0x3c, 0x84, 0xf2, 0x1f, 0x9b, 0x48, 0xb8, 0x3f, 0xfd, 0x9d, 0x57, 0x54, 0xae, 0xa1,
	0xae, 0xa0, 0x72, 0xe0, 0x0f, 0x30, 0x75, 0x63, 0x64, 0xc7, 0xda, 0xb7, 0xe3, 0xe6, 0xe3, 0x1a,
	0xf2, 0x40, 0x15, 0xaf, 0x10, 0x29, 0x6a, 0x85, 0xae, 0xca, 0xca, 0xa0, 0xd0, 0x68, 0x75, 0x6d,
	0xb0, 0xd7, 0x61, 0xea, 0x9d, 0x7d, 0xa6, 0x75, 0xe3, 0xf4, 0x6e, 0x2b, 0x14, 0xb5, 0xd9, 0x5f,
	0x9a, 0x05, 0x7c, 0x69, 0xcb, 0xef, 0x65, 0x4d, 0x99, 0x67, 0xb7, 0x13, 0x2f, 0x8f, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x8e, 0xc8, 0x73, 0x23, 0x88, 0x01, 0x00, 0x00,
}
