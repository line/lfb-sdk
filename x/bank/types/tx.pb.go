// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lbm/bank/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	github_com_line_lbm_sdk_v2_types "github.com/line/lbm-sdk/v2/types"
	types "github.com/line/lbm-sdk/v2/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgSend represents a message to send coins from one account to another.
type MsgSend struct {
	FromAddress string                                 `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty" yaml:"from_address"`
	ToAddress   string                                 `protobuf:"bytes,2,opt,name=to_address,json=toAddress,proto3" json:"to_address,omitempty" yaml:"to_address"`
	Amount      github_com_line_lbm_sdk_v2_types.Coins `protobuf:"bytes,3,rep,name=amount,proto3,castrepeated=github.com/line/lbm-sdk/v2/types.Coins" json:"amount"`
}

func (m *MsgSend) Reset()         { *m = MsgSend{} }
func (m *MsgSend) String() string { return proto.CompactTextString(m) }
func (*MsgSend) ProtoMessage()    {}
func (*MsgSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_549878a8fb2cb9ef, []int{0}
}
func (m *MsgSend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSend.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSend.Merge(m, src)
}
func (m *MsgSend) XXX_Size() int {
	return m.Size()
}
func (m *MsgSend) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSend.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSend proto.InternalMessageInfo

// MsgSendResponse defines the Msg/Send response type.
type MsgSendResponse struct {
}

func (m *MsgSendResponse) Reset()         { *m = MsgSendResponse{} }
func (m *MsgSendResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSendResponse) ProtoMessage()    {}
func (*MsgSendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_549878a8fb2cb9ef, []int{1}
}
func (m *MsgSendResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendResponse.Merge(m, src)
}
func (m *MsgSendResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendResponse proto.InternalMessageInfo

// MsgMultiSend represents an arbitrary multi-in, multi-out send message.
type MsgMultiSend struct {
	Inputs  []Input  `protobuf:"bytes,1,rep,name=inputs,proto3" json:"inputs"`
	Outputs []Output `protobuf:"bytes,2,rep,name=outputs,proto3" json:"outputs"`
}

func (m *MsgMultiSend) Reset()         { *m = MsgMultiSend{} }
func (m *MsgMultiSend) String() string { return proto.CompactTextString(m) }
func (*MsgMultiSend) ProtoMessage()    {}
func (*MsgMultiSend) Descriptor() ([]byte, []int) {
	return fileDescriptor_549878a8fb2cb9ef, []int{2}
}
func (m *MsgMultiSend) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMultiSend) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMultiSend.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMultiSend) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMultiSend.Merge(m, src)
}
func (m *MsgMultiSend) XXX_Size() int {
	return m.Size()
}
func (m *MsgMultiSend) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMultiSend.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMultiSend proto.InternalMessageInfo

func (m *MsgMultiSend) GetInputs() []Input {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *MsgMultiSend) GetOutputs() []Output {
	if m != nil {
		return m.Outputs
	}
	return nil
}

// MsgMultiSendResponse defines the Msg/MultiSend response type.
type MsgMultiSendResponse struct {
}

func (m *MsgMultiSendResponse) Reset()         { *m = MsgMultiSendResponse{} }
func (m *MsgMultiSendResponse) String() string { return proto.CompactTextString(m) }
func (*MsgMultiSendResponse) ProtoMessage()    {}
func (*MsgMultiSendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_549878a8fb2cb9ef, []int{3}
}
func (m *MsgMultiSendResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgMultiSendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgMultiSendResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgMultiSendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgMultiSendResponse.Merge(m, src)
}
func (m *MsgMultiSendResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgMultiSendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgMultiSendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgMultiSendResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSend)(nil), "lbm.bank.v1beta1.MsgSend")
	proto.RegisterType((*MsgSendResponse)(nil), "lbm.bank.v1beta1.MsgSendResponse")
	proto.RegisterType((*MsgMultiSend)(nil), "lbm.bank.v1beta1.MsgMultiSend")
	proto.RegisterType((*MsgMultiSendResponse)(nil), "lbm.bank.v1beta1.MsgMultiSendResponse")
}

func init() { proto.RegisterFile("lbm/bank/v1beta1/tx.proto", fileDescriptor_549878a8fb2cb9ef) }

var fileDescriptor_549878a8fb2cb9ef = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x8f, 0xd2, 0x40,
	0x00, 0x85, 0x3b, 0x0b, 0x61, 0x65, 0x76, 0x13, 0xdd, 0xba, 0xee, 0xb2, 0x35, 0x69, 0xd7, 0x1e,
	0xd6, 0xbd, 0x38, 0x13, 0x50, 0x13, 0xd3, 0x1b, 0xd5, 0x8b, 0x87, 0xc6, 0xa4, 0xdc, 0x3c, 0x68,
	0x5a, 0x3a, 0xd6, 0x86, 0x76, 0xa6, 0x61, 0xa6, 0x04, 0x7e, 0x80, 0x89, 0x47, 0x7f, 0x02, 0x67,
	0x7f, 0x09, 0x47, 0x8e, 0x9e, 0xd0, 0x40, 0x4c, 0x3c, 0xf3, 0x0b, 0x4c, 0xa7, 0x2d, 0xa0, 0x28,
	0xb7, 0x4e, 0xbe, 0xf7, 0x5e, 0xde, 0xeb, 0x0c, 0xbc, 0x8a, 0xfd, 0x04, 0xfb, 0x1e, 0x1d, 0xe0,
	0x51, 0xdb, 0x27, 0xc2, 0x6b, 0x63, 0x31, 0x46, 0xe9, 0x90, 0x09, 0xa6, 0xde, 0x8b, 0xfd, 0x04,
	0xe5, 0x08, 0x95, 0x48, 0x3b, 0x0f, 0x59, 0xc8, 0x24, 0xc4, 0xf9, 0x57, 0xa1, 0xd3, 0x1e, 0x16,
	0x11, 0x9c, 0x6c, 0x22, 0xfa, 0x2c, 0xa2, 0x7f, 0xc2, 0x9d, 0x7c, 0x99, 0x28, 0xa1, 0xf9, 0x13,
	0xc0, 0x63, 0x87, 0x87, 0x3d, 0x42, 0x03, 0xd5, 0x82, 0xa7, 0x1f, 0x86, 0x2c, 0x79, 0xef, 0x05,
	0xc1, 0x90, 0x70, 0xde, 0x02, 0xd7, 0xe0, 0xb6, 0x69, 0x5f, 0xae, 0x17, 0xc6, 0xfd, 0x89, 0x97,
	0xc4, 0x96, 0xb9, 0x4b, 0x4d, 0xf7, 0x24, 0x3f, 0x76, 0x8b, 0x93, 0xfa, 0x0c, 0x42, 0xc1, 0x36,
	0xce, 0x23, 0xe9, 0x7c, 0xb0, 0x5e, 0x18, 0x67, 0x85, 0x73, 0xcb, 0x4c, 0xb7, 0x29, 0x58, 0xe5,
	0x7a, 0x07, 0x1b, 0x5e, 0xc2, 0x32, 0x2a, 0x5a, 0xb5, 0xeb, 0xda, 0xed, 0x49, 0xe7, 0x02, 0x15,
	0x83, 0x39, 0xa9, 0x06, 0xa3, 0x97, 0x2c, 0xa2, 0x36, 0x9a, 0x2d, 0x0c, 0xe5, 0xeb, 0x77, 0xe3,
	0x26, 0x8c, 0xc4, 0xc7, 0xcc, 0x47, 0x7d, 0x96, 0xe0, 0x38, 0xa2, 0x04, 0xc7, 0x7e, 0xf2, 0x84,
	0x07, 0x03, 0x3c, 0xea, 0x60, 0x31, 0x49, 0x09, 0x97, 0x72, 0xee, 0x96, 0xa9, 0xd6, 0x9d, 0xcf,
	0x53, 0x43, 0xf9, 0x35, 0x35, 0x14, 0xf3, 0x0c, 0xde, 0x2d, 0x67, 0xba, 0x84, 0xa7, 0x8c, 0x72,
	0x62, 0x7e, 0x02, 0xf0, 0xd4, 0xe1, 0xa1, 0x93, 0xc5, 0x22, 0x92, 0xfb, 0x9f, 0xc3, 0x46, 0x44,
	0xd3, 0x4c, 0xe4, 0xcb, 0xf3, 0x36, 0x97, 0xe8, 0xef, 0xdf, 0x8f, 0x5e, 0xe7, 0xdc, 0xae, 0xe7,
	0x75, 0xdc, 0x52, 0xac, 0xbe, 0x80, 0xc7, 0x2c, 0x13, 0xd2, 0x77, 0x24, 0x7d, 0xad, 0x7d, 0xdf,
	0x1b, 0x29, 0x28, 0x8d, 0x95, 0xdc, 0xaa, 0xcb, 0x6a, 0x17, 0xf0, 0x7c, 0xb7, 0x46, 0xd5, 0xaf,
	0x33, 0x05, 0xb0, 0xe6, 0xf0, 0x50, 0x7d, 0x05, 0xeb, 0xb2, 0xde, 0xd5, 0x7e, 0x6c, 0x39, 0x49,
	0x7b, 0xf4, 0x5f, 0x54, 0xa5, 0xa9, 0x3d, 0xd8, 0xdc, 0x2e, 0xd5, 0xff, 0xa9, 0xdf, 0x70, 0xed,
	0xe6, 0x30, 0xaf, 0x42, 0xed, 0xee, 0x6c, 0xa9, 0x83, 0xf9, 0x52, 0x07, 0x3f, 0x96, 0x3a, 0xf8,
	0xb2, 0xd2, 0x95, 0xf9, 0x4a, 0x57, 0xbe, 0xad, 0x74, 0xe5, 0xed, 0xe3, 0x03, 0x37, 0x35, 0x2e,
	0x5e, 0xa3, 0xbc, 0x30, 0xbf, 0x21, 0xdf, 0xe1, 0xd3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xae,
	0xc5, 0xa7, 0xf4, 0x06, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// Send defines a method for sending coins from one account to another account.
	Send(ctx context.Context, in *MsgSend, opts ...grpc.CallOption) (*MsgSendResponse, error)
	// MultiSend defines a method for sending coins from some accounts to other accounts.
	MultiSend(ctx context.Context, in *MsgMultiSend, opts ...grpc.CallOption) (*MsgMultiSendResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Send(ctx context.Context, in *MsgSend, opts ...grpc.CallOption) (*MsgSendResponse, error) {
	out := new(MsgSendResponse)
	err := c.cc.Invoke(ctx, "/lbm.bank.v1beta1.Msg/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) MultiSend(ctx context.Context, in *MsgMultiSend, opts ...grpc.CallOption) (*MsgMultiSendResponse, error) {
	out := new(MsgMultiSendResponse)
	err := c.cc.Invoke(ctx, "/lbm.bank.v1beta1.Msg/MultiSend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// Send defines a method for sending coins from one account to another account.
	Send(context.Context, *MsgSend) (*MsgSendResponse, error)
	// MultiSend defines a method for sending coins from some accounts to other accounts.
	MultiSend(context.Context, *MsgMultiSend) (*MsgMultiSendResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Send(ctx context.Context, req *MsgSend) (*MsgSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (*UnimplementedMsgServer) MultiSend(ctx context.Context, req *MsgMultiSend) (*MsgMultiSendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MultiSend not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSend)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lbm.bank.v1beta1.Msg/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Send(ctx, req.(*MsgSend))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_MultiSend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgMultiSend)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).MultiSend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lbm.bank.v1beta1.Msg/MultiSend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).MultiSend(ctx, req.(*MsgMultiSend))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lbm.bank.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Msg_Send_Handler,
		},
		{
			MethodName: "MultiSend",
			Handler:    _Msg_MultiSend_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lbm/bank/v1beta1/tx.proto",
}

func (m *MsgSend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.ToAddress) > 0 {
		i -= len(m.ToAddress)
		copy(dAtA[i:], m.ToAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ToAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSendResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgMultiSend) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMultiSend) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMultiSend) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Outputs) > 0 {
		for iNdEx := len(m.Outputs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Outputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Inputs) > 0 {
		for iNdEx := len(m.Inputs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Inputs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgMultiSendResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgMultiSendResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgMultiSendResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ToAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgSendResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgMultiSend) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Inputs) > 0 {
		for _, e := range m.Inputs {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.Outputs) > 0 {
		for _, e := range m.Outputs {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgMultiSendResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSend) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSendResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSendResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgMultiSend) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgMultiSend: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMultiSend: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inputs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Inputs = append(m.Inputs, Input{})
			if err := m.Inputs[len(m.Inputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Outputs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Outputs = append(m.Outputs, Output{})
			if err := m.Outputs[len(m.Outputs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgMultiSendResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgMultiSendResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgMultiSendResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
