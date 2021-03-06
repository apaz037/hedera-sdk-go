// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/ContractCallLocal.proto

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

// The log information for an event returned by a smart contract function call. One function call may return several such events.
type ContractLoginfo struct {
	ContractID           *ContractID `protobuf:"bytes,1,opt,name=contractID,proto3" json:"contractID,omitempty"`
	Bloom                []byte      `protobuf:"bytes,2,opt,name=bloom,proto3" json:"bloom,omitempty"`
	Topic                [][]byte    `protobuf:"bytes,3,rep,name=topic,proto3" json:"topic,omitempty"`
	Data                 []byte      `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ContractLoginfo) Reset()         { *m = ContractLoginfo{} }
func (m *ContractLoginfo) String() string { return proto.CompactTextString(m) }
func (*ContractLoginfo) ProtoMessage()    {}
func (*ContractLoginfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b1ca84e1de6fd43, []int{0}
}

func (m *ContractLoginfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractLoginfo.Unmarshal(m, b)
}
func (m *ContractLoginfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractLoginfo.Marshal(b, m, deterministic)
}
func (m *ContractLoginfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractLoginfo.Merge(m, src)
}
func (m *ContractLoginfo) XXX_Size() int {
	return xxx_messageInfo_ContractLoginfo.Size(m)
}
func (m *ContractLoginfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractLoginfo.DiscardUnknown(m)
}

var xxx_messageInfo_ContractLoginfo proto.InternalMessageInfo

func (m *ContractLoginfo) GetContractID() *ContractID {
	if m != nil {
		return m.ContractID
	}
	return nil
}

func (m *ContractLoginfo) GetBloom() []byte {
	if m != nil {
		return m.Bloom
	}
	return nil
}

func (m *ContractLoginfo) GetTopic() [][]byte {
	if m != nil {
		return m.Topic
	}
	return nil
}

func (m *ContractLoginfo) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

// The result returned by a call to a smart contract function. This is part of the response to a ContractCallLocal query, and is in the record for a ContractCall or ContractCreateInstance transaction. The ContractCreateInstance transaction record has the results of the call to the constructor.
type ContractFunctionResult struct {
	ContractID           *ContractID        `protobuf:"bytes,1,opt,name=contractID,proto3" json:"contractID,omitempty"`
	ContractCallResult   []byte             `protobuf:"bytes,2,opt,name=contractCallResult,proto3" json:"contractCallResult,omitempty"`
	ErrorMessage         string             `protobuf:"bytes,3,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	Bloom                []byte             `protobuf:"bytes,4,opt,name=bloom,proto3" json:"bloom,omitempty"`
	GasUsed              uint64             `protobuf:"varint,5,opt,name=gasUsed,proto3" json:"gasUsed,omitempty"`
	LogInfo              []*ContractLoginfo `protobuf:"bytes,6,rep,name=logInfo,proto3" json:"logInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ContractFunctionResult) Reset()         { *m = ContractFunctionResult{} }
func (m *ContractFunctionResult) String() string { return proto.CompactTextString(m) }
func (*ContractFunctionResult) ProtoMessage()    {}
func (*ContractFunctionResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b1ca84e1de6fd43, []int{1}
}

func (m *ContractFunctionResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractFunctionResult.Unmarshal(m, b)
}
func (m *ContractFunctionResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractFunctionResult.Marshal(b, m, deterministic)
}
func (m *ContractFunctionResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractFunctionResult.Merge(m, src)
}
func (m *ContractFunctionResult) XXX_Size() int {
	return xxx_messageInfo_ContractFunctionResult.Size(m)
}
func (m *ContractFunctionResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractFunctionResult.DiscardUnknown(m)
}

var xxx_messageInfo_ContractFunctionResult proto.InternalMessageInfo

func (m *ContractFunctionResult) GetContractID() *ContractID {
	if m != nil {
		return m.ContractID
	}
	return nil
}

func (m *ContractFunctionResult) GetContractCallResult() []byte {
	if m != nil {
		return m.ContractCallResult
	}
	return nil
}

func (m *ContractFunctionResult) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

func (m *ContractFunctionResult) GetBloom() []byte {
	if m != nil {
		return m.Bloom
	}
	return nil
}

func (m *ContractFunctionResult) GetGasUsed() uint64 {
	if m != nil {
		return m.GasUsed
	}
	return 0
}

func (m *ContractFunctionResult) GetLogInfo() []*ContractLoginfo {
	if m != nil {
		return m.LogInfo
	}
	return nil
}

// Call a function of the given smart contract instance, giving it functionParameters as its inputs. It will consume the entire given amount of gas.
//
// This is performed locally on the particular node that the client is communicating with. It cannot change the state of the contract instance (and so, cannot spend anything from the instance's cryptocurrency account). It will not have a consensus timestamp. It cannot generate a record or a receipt. The response will contain the output returned by the function call.  This is useful for calling getter functions, which purely read the state and don't change it. It is faster and cheaper than a normal call, because it is purely local to a single  node.
type ContractCallLocalQuery struct {
	Header               *QueryHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	ContractID           *ContractID  `protobuf:"bytes,2,opt,name=contractID,proto3" json:"contractID,omitempty"`
	Gas                  int64        `protobuf:"varint,3,opt,name=gas,proto3" json:"gas,omitempty"`
	FunctionParameters   []byte       `protobuf:"bytes,4,opt,name=functionParameters,proto3" json:"functionParameters,omitempty"`
	MaxResultSize        int64        `protobuf:"varint,5,opt,name=maxResultSize,proto3" json:"maxResultSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ContractCallLocalQuery) Reset()         { *m = ContractCallLocalQuery{} }
func (m *ContractCallLocalQuery) String() string { return proto.CompactTextString(m) }
func (*ContractCallLocalQuery) ProtoMessage()    {}
func (*ContractCallLocalQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b1ca84e1de6fd43, []int{2}
}

func (m *ContractCallLocalQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractCallLocalQuery.Unmarshal(m, b)
}
func (m *ContractCallLocalQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractCallLocalQuery.Marshal(b, m, deterministic)
}
func (m *ContractCallLocalQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractCallLocalQuery.Merge(m, src)
}
func (m *ContractCallLocalQuery) XXX_Size() int {
	return xxx_messageInfo_ContractCallLocalQuery.Size(m)
}
func (m *ContractCallLocalQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractCallLocalQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ContractCallLocalQuery proto.InternalMessageInfo

func (m *ContractCallLocalQuery) GetHeader() *QueryHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ContractCallLocalQuery) GetContractID() *ContractID {
	if m != nil {
		return m.ContractID
	}
	return nil
}

func (m *ContractCallLocalQuery) GetGas() int64 {
	if m != nil {
		return m.Gas
	}
	return 0
}

func (m *ContractCallLocalQuery) GetFunctionParameters() []byte {
	if m != nil {
		return m.FunctionParameters
	}
	return nil
}

func (m *ContractCallLocalQuery) GetMaxResultSize() int64 {
	if m != nil {
		return m.MaxResultSize
	}
	return 0
}

// Response when the client sends the node ContractCallLocalQuery
type ContractCallLocalResponse struct {
	Header               *ResponseHeader         `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	FunctionResult       *ContractFunctionResult `protobuf:"bytes,2,opt,name=functionResult,proto3" json:"functionResult,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *ContractCallLocalResponse) Reset()         { *m = ContractCallLocalResponse{} }
func (m *ContractCallLocalResponse) String() string { return proto.CompactTextString(m) }
func (*ContractCallLocalResponse) ProtoMessage()    {}
func (*ContractCallLocalResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_3b1ca84e1de6fd43, []int{3}
}

func (m *ContractCallLocalResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContractCallLocalResponse.Unmarshal(m, b)
}
func (m *ContractCallLocalResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContractCallLocalResponse.Marshal(b, m, deterministic)
}
func (m *ContractCallLocalResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractCallLocalResponse.Merge(m, src)
}
func (m *ContractCallLocalResponse) XXX_Size() int {
	return xxx_messageInfo_ContractCallLocalResponse.Size(m)
}
func (m *ContractCallLocalResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractCallLocalResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ContractCallLocalResponse proto.InternalMessageInfo

func (m *ContractCallLocalResponse) GetHeader() *ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ContractCallLocalResponse) GetFunctionResult() *ContractFunctionResult {
	if m != nil {
		return m.FunctionResult
	}
	return nil
}

func init() {
	proto.RegisterType((*ContractLoginfo)(nil), "proto.ContractLoginfo")
	proto.RegisterType((*ContractFunctionResult)(nil), "proto.ContractFunctionResult")
	proto.RegisterType((*ContractCallLocalQuery)(nil), "proto.ContractCallLocalQuery")
	proto.RegisterType((*ContractCallLocalResponse)(nil), "proto.ContractCallLocalResponse")
}

func init() { proto.RegisterFile("proto/ContractCallLocal.proto", fileDescriptor_3b1ca84e1de6fd43) }

var fileDescriptor_3b1ca84e1de6fd43 = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x55, 0x9a, 0xb6, 0x13, 0x77, 0xe3, 0xcb, 0x82, 0x12, 0x2a, 0x4d, 0x8a, 0x22, 0x1e, 0x22,
	0xa4, 0xa6, 0x30, 0xfe, 0xc1, 0x0a, 0x68, 0x93, 0x86, 0x34, 0x0c, 0xbc, 0xf0, 0xe6, 0x3a, 0x6e,
	0x12, 0x91, 0xe4, 0x46, 0xb6, 0x2b, 0x31, 0xde, 0x90, 0xf8, 0x03, 0xfc, 0x42, 0xfe, 0x0a, 0xaa,
	0xed, 0x74, 0x49, 0xe8, 0x03, 0xda, 0x93, 0xed, 0x73, 0x4e, 0x6e, 0xee, 0x39, 0xf7, 0xc2, 0x69,
	0x23, 0x51, 0xe3, 0x72, 0x85, 0xb5, 0x96, 0x8c, 0xeb, 0x15, 0x2b, 0xcb, 0x2b, 0xe4, 0xac, 0x4c,
	0x0c, 0x4e, 0x26, 0xe6, 0x98, 0xcf, 0xac, 0xea, 0x9c, 0xa9, 0x82, 0x7f, 0xbe, 0x69, 0x84, 0xb2,
	0xf4, 0xfc, 0x99, 0xc5, 0x3f, 0x6e, 0x85, 0xbc, 0xb9, 0x10, 0x2c, 0x15, 0xd2, 0x11, 0x73, 0x4b,
	0x50, 0xa1, 0x1a, 0xac, 0x95, 0xe8, 0x72, 0xd1, 0x2f, 0x0f, 0x1e, 0xb6, 0xff, 0xbb, 0xc2, 0xac,
	0xa8, 0x37, 0x48, 0x5e, 0x03, 0x70, 0x07, 0x5d, 0xbe, 0x0d, 0xbc, 0xd0, 0x8b, 0x8f, 0xcf, 0x1e,
	0x5b, 0x7d, 0xb2, 0xda, 0x13, 0xb4, 0x23, 0x22, 0x4f, 0x60, 0xb2, 0x2e, 0x11, 0xab, 0x60, 0x14,
	0x7a, 0xf1, 0x09, 0xb5, 0x8f, 0x1d, 0xaa, 0xb1, 0x29, 0x78, 0xe0, 0x87, 0xfe, 0x0e, 0x35, 0x0f,
	0x42, 0x60, 0x9c, 0x32, 0xcd, 0x82, 0xb1, 0x91, 0x9a, 0x7b, 0xf4, 0x73, 0x04, 0xb3, 0xb6, 0xf4,
	0xfb, 0x6d, 0xcd, 0x75, 0x81, 0x35, 0x15, 0x6a, 0x5b, 0xea, 0xbb, 0x74, 0x93, 0x00, 0xe1, 0x9d,
	0x0c, 0x6d, 0x21, 0xd7, 0xda, 0x01, 0x86, 0x44, 0x70, 0x22, 0xa4, 0x44, 0xf9, 0x41, 0x28, 0xc5,
	0x32, 0x11, 0xf8, 0xa1, 0x17, 0xdf, 0xa3, 0x3d, 0xec, 0xd6, 0xe1, 0xb8, 0xeb, 0x30, 0x80, 0xa3,
	0x8c, 0xa9, 0x2f, 0x4a, 0xa4, 0xc1, 0x24, 0xf4, 0xe2, 0x31, 0x6d, 0x9f, 0xe4, 0x15, 0x1c, 0x95,
	0x98, 0x5d, 0xd6, 0x1b, 0x0c, 0xa6, 0xa1, 0x1f, 0x1f, 0x9f, 0xcd, 0x06, 0x3d, 0xbb, 0xb4, 0x69,
	0x2b, 0x8b, 0xfe, 0x78, 0xb7, 0x19, 0xec, 0x47, 0x6f, 0xa6, 0x49, 0x5e, 0xc2, 0x34, 0x37, 0x53,
	0x73, 0xfe, 0x89, 0xab, 0xd5, 0x99, 0x35, 0x75, 0x8a, 0x41, 0x5e, 0xa3, 0xff, 0xc9, 0xeb, 0x11,
	0xf8, 0x19, 0x53, 0xc6, 0xb6, 0x4f, 0x77, 0xd7, 0x5d, 0x82, 0x1b, 0x37, 0x86, 0x6b, 0x26, 0x59,
	0x25, 0xb4, 0x90, 0xca, 0x59, 0x3f, 0xc0, 0x90, 0x17, 0x70, 0xbf, 0x62, 0xdf, 0x6d, 0x9c, 0x9f,
	0x8a, 0x1f, 0xc2, 0xa4, 0xe1, 0xd3, 0x3e, 0x18, 0xfd, 0xf6, 0xe0, 0xf9, 0x3f, 0x0e, 0xdb, 0xb5,
	0x24, 0x8b, 0x81, 0xc9, 0xa7, 0xae, 0xe9, 0xfe, 0xde, 0xee, 0x7d, 0xbe, 0x83, 0x07, 0x9b, 0xde,
	0xa6, 0x38, 0xaf, 0xa7, 0x03, 0xaf, 0xfd, 0x75, 0xa2, 0x83, 0x8f, 0xce, 0x2f, 0x60, 0xce, 0xb1,
	0x4a, 0x72, 0x91, 0x0a, 0xc9, 0x92, 0x9c, 0xa9, 0x3c, 0x93, 0xac, 0xc9, 0x6d, 0x91, 0x6b, 0xef,
	0x6b, 0x9c, 0x15, 0x3a, 0xdf, 0xae, 0x13, 0x8e, 0xd5, 0x72, 0xcf, 0x2e, 0xad, 0x7c, 0xa1, 0xd2,
	0x6f, 0x8b, 0x0c, 0x97, 0x46, 0xbb, 0x9e, 0x9a, 0xe3, 0xcd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xbe, 0xdd, 0x4a, 0x33, 0xc6, 0x03, 0x00, 0x00,
}
