// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lbm/fswap/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/Finschia/finschia-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type QuerySwappedRequest struct {
}

func (m *QuerySwappedRequest) Reset()         { *m = QuerySwappedRequest{} }
func (m *QuerySwappedRequest) String() string { return proto.CompactTextString(m) }
func (*QuerySwappedRequest) ProtoMessage()    {}
func (*QuerySwappedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01deae9da7816d6a, []int{0}
}
func (m *QuerySwappedRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySwappedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySwappedRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySwappedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySwappedRequest.Merge(m, src)
}
func (m *QuerySwappedRequest) XXX_Size() int {
	return m.Size()
}
func (m *QuerySwappedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySwappedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySwappedRequest proto.InternalMessageInfo

type QuerySwappedResponse struct {
	OldCoinAmount types.Coin `protobuf:"bytes,1,opt,name=old_coin_amount,json=oldCoinAmount,proto3,castrepeated=github.com/Finschia/finschia-sdk/types.Coin" json:"old_coin_amount"`
	NewCoinAmount types.Coin `protobuf:"bytes,2,opt,name=new_coin_amount,json=newCoinAmount,proto3,castrepeated=github.com/Finschia/finschia-sdk/types.Coin" json:"new_coin_amount"`
}

func (m *QuerySwappedResponse) Reset()         { *m = QuerySwappedResponse{} }
func (m *QuerySwappedResponse) String() string { return proto.CompactTextString(m) }
func (*QuerySwappedResponse) ProtoMessage()    {}
func (*QuerySwappedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01deae9da7816d6a, []int{1}
}
func (m *QuerySwappedResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QuerySwappedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QuerySwappedResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QuerySwappedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuerySwappedResponse.Merge(m, src)
}
func (m *QuerySwappedResponse) XXX_Size() int {
	return m.Size()
}
func (m *QuerySwappedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuerySwappedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuerySwappedResponse proto.InternalMessageInfo

func (m *QuerySwappedResponse) GetOldCoinAmount() types.Coin {
	if m != nil {
		return m.OldCoinAmount
	}
	return types.Coin{}
}

func (m *QuerySwappedResponse) GetNewCoinAmount() types.Coin {
	if m != nil {
		return m.NewCoinAmount
	}
	return types.Coin{}
}

type QueryTotalSwappableAmountRequest struct {
}

func (m *QueryTotalSwappableAmountRequest) Reset()         { *m = QueryTotalSwappableAmountRequest{} }
func (m *QueryTotalSwappableAmountRequest) String() string { return proto.CompactTextString(m) }
func (*QueryTotalSwappableAmountRequest) ProtoMessage()    {}
func (*QueryTotalSwappableAmountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_01deae9da7816d6a, []int{2}
}
func (m *QueryTotalSwappableAmountRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTotalSwappableAmountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTotalSwappableAmountRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTotalSwappableAmountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTotalSwappableAmountRequest.Merge(m, src)
}
func (m *QueryTotalSwappableAmountRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryTotalSwappableAmountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTotalSwappableAmountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTotalSwappableAmountRequest proto.InternalMessageInfo

type QueryTotalSwappableAmountResponse struct {
	SwappableNewCoin types.Coin `protobuf:"bytes,1,opt,name=swappable_new_coin,json=swappableNewCoin,proto3,castrepeated=github.com/Finschia/finschia-sdk/types.Coin" json:"swappable_new_coin"`
}

func (m *QueryTotalSwappableAmountResponse) Reset()         { *m = QueryTotalSwappableAmountResponse{} }
func (m *QueryTotalSwappableAmountResponse) String() string { return proto.CompactTextString(m) }
func (*QueryTotalSwappableAmountResponse) ProtoMessage()    {}
func (*QueryTotalSwappableAmountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_01deae9da7816d6a, []int{3}
}
func (m *QueryTotalSwappableAmountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryTotalSwappableAmountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryTotalSwappableAmountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryTotalSwappableAmountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryTotalSwappableAmountResponse.Merge(m, src)
}
func (m *QueryTotalSwappableAmountResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryTotalSwappableAmountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryTotalSwappableAmountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryTotalSwappableAmountResponse proto.InternalMessageInfo

func (m *QueryTotalSwappableAmountResponse) GetSwappableNewCoin() types.Coin {
	if m != nil {
		return m.SwappableNewCoin
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*QuerySwappedRequest)(nil), "lbm.fswap.v1.QuerySwappedRequest")
	proto.RegisterType((*QuerySwappedResponse)(nil), "lbm.fswap.v1.QuerySwappedResponse")
	proto.RegisterType((*QueryTotalSwappableAmountRequest)(nil), "lbm.fswap.v1.QueryTotalSwappableAmountRequest")
	proto.RegisterType((*QueryTotalSwappableAmountResponse)(nil), "lbm.fswap.v1.QueryTotalSwappableAmountResponse")
}

func init() { proto.RegisterFile("lbm/fswap/v1/query.proto", fileDescriptor_01deae9da7816d6a) }

var fileDescriptor_01deae9da7816d6a = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x93, 0xc1, 0xaa, 0xd3, 0x40,
	0x14, 0x86, 0x93, 0x80, 0x0a, 0xa3, 0xa2, 0x8c, 0xbd, 0x78, 0x0d, 0x9a, 0x7b, 0x6f, 0x36, 0x2a,
	0xe2, 0x0c, 0xb9, 0xf7, 0x09, 0xac, 0x20, 0xae, 0x04, 0xab, 0x2b, 0x37, 0x65, 0x92, 0x4c, 0xd3,
	0x60, 0x32, 0x27, 0xed, 0x4c, 0x12, 0x0b, 0xae, 0x5c, 0xb8, 0x16, 0x7c, 0x02, 0x71, 0x23, 0x3e,
	0x49, 0x97, 0x05, 0x37, 0xae, 0x54, 0x5a, 0x1f, 0x44, 0x32, 0x99, 0xaa, 0x81, 0x50, 0xdd, 0x74,
	0x77, 0x98, 0xff, 0x64, 0xbe, 0x7f, 0xfe, 0x73, 0x82, 0x0e, 0xb3, 0x30, 0xa7, 0x13, 0x59, 0xb3,
	0x82, 0x56, 0x01, 0x9d, 0x95, 0x7c, 0xbe, 0x20, 0xc5, 0x1c, 0x14, 0xe0, 0x4b, 0x59, 0x98, 0x13,
	0xad, 0x90, 0x2a, 0x70, 0x6f, 0x26, 0x00, 0x49, 0xc6, 0x29, 0x2b, 0x52, 0xca, 0x84, 0x00, 0xc5,
	0x54, 0x0a, 0x42, 0xb6, 0xbd, 0xee, 0x20, 0x81, 0x04, 0x74, 0x49, 0x9b, 0xca, 0x9c, 0x7a, 0x11,
	0xc8, 0x1c, 0x24, 0x0d, 0x99, 0xe4, 0xb4, 0x0a, 0x42, 0xae, 0x58, 0x40, 0x23, 0x48, 0x85, 0xd1,
	0xbb, 0xec, 0x16, 0xa5, 0x15, 0xff, 0x00, 0x5d, 0x7b, 0xda, 0x58, 0x79, 0x56, 0xb3, 0xa2, 0xe0,
	0xf1, 0x88, 0xcf, 0x4a, 0x2e, 0x95, 0xff, 0xd6, 0x41, 0x83, 0xee, 0xb9, 0x2c, 0x40, 0x48, 0x8e,
	0x2b, 0x74, 0x05, 0xb2, 0x78, 0xdc, 0xdc, 0x3d, 0x66, 0x39, 0x94, 0x42, 0x1d, 0xda, 0xc7, 0xf6,
	0x9d, 0x8b, 0xa7, 0x37, 0x48, 0xeb, 0x81, 0x34, 0x1e, 0x88, 0xf1, 0x40, 0x1e, 0x42, 0x2a, 0x86,
	0x67, 0xcb, 0x6f, 0x47, 0xd6, 0xe7, 0xef, 0x47, 0xf7, 0x92, 0x54, 0x4d, 0xcb, 0x90, 0x44, 0x90,
	0xd3, 0x47, 0xa9, 0x90, 0xd1, 0x34, 0x65, 0x74, 0x62, 0x8a, 0xfb, 0x32, 0x7e, 0x49, 0xd5, 0xa2,
	0xe0, 0x52, 0x7f, 0x34, 0xba, 0x0c, 0x59, 0xdc, 0x14, 0x0f, 0x34, 0xa4, 0xe1, 0x0a, 0x5e, 0x77,
	0xb8, 0xce, 0x7e, 0xb8, 0x82, 0xd7, 0x7f, 0xb8, 0xbe, 0x8f, 0x8e, 0x75, 0x0e, 0xcf, 0x41, 0xb1,
	0x4c, 0x87, 0xc1, 0xc2, 0x8c, 0xb7, 0xe2, 0x36, 0xac, 0x0f, 0x36, 0x3a, 0xd9, 0xd1, 0x64, 0x92,
	0x7b, 0x8d, 0xb0, 0xdc, 0x4a, 0xe3, 0xed, 0x5b, 0xf6, 0x14, 0xde, 0xd5, 0xdf, 0xa4, 0x27, 0xed,
	0x6b, 0x4e, 0x3f, 0x3a, 0xe8, 0x9c, 0xf6, 0x88, 0x01, 0x5d, 0x30, 0x43, 0xc5, 0x27, 0xe4, 0xef,
	0xcd, 0x23, 0x3d, 0x8b, 0xe0, 0xfa, 0xbb, 0x5a, 0xda, 0x97, 0xf9, 0xb7, 0xde, 0x7c, 0xf9, 0xf9,
	0xde, 0xb9, 0x8e, 0x0f, 0x68, 0x67, 0xcd, 0xa4, 0xa1, 0x7c, 0xb2, 0xd1, 0xa0, 0x2f, 0x19, 0x4c,
	0x7a, 0xee, 0xde, 0x91, 0xb3, 0x4b, 0xff, 0xbb, 0xdf, 0x18, 0xa3, 0xda, 0xd8, 0x5d, 0x7c, 0xbb,
	0xc7, 0x58, 0x67, 0x0c, 0x66, 0xa5, 0x86, 0x8f, 0x97, 0x6b, 0xcf, 0x5e, 0xad, 0x3d, 0xfb, 0xc7,
	0xda, 0xb3, 0xdf, 0x6d, 0x3c, 0x6b, 0xb5, 0xf1, 0xac, 0xaf, 0x1b, 0xcf, 0x7a, 0x41, 0xfe, 0x19,
	0xff, 0x2b, 0x03, 0xd0, 0x63, 0x08, 0xcf, 0xeb, 0xdf, 0xeb, 0xec, 0x57, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x63, 0x79, 0x56, 0x3c, 0xf6, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Swapped queries the current swapped status that includes a burnt amount of old coin and a minted amount of new
	// coin.
	Swapped(ctx context.Context, in *QuerySwappedRequest, opts ...grpc.CallOption) (*QuerySwappedResponse, error)
	// TotalSwappableAmount queries the current swappable amount for new coin.
	TotalSwappableAmount(ctx context.Context, in *QueryTotalSwappableAmountRequest, opts ...grpc.CallOption) (*QueryTotalSwappableAmountResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Swapped(ctx context.Context, in *QuerySwappedRequest, opts ...grpc.CallOption) (*QuerySwappedResponse, error) {
	out := new(QuerySwappedResponse)
	err := c.cc.Invoke(ctx, "/lbm.fswap.v1.Query/Swapped", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TotalSwappableAmount(ctx context.Context, in *QueryTotalSwappableAmountRequest, opts ...grpc.CallOption) (*QueryTotalSwappableAmountResponse, error) {
	out := new(QueryTotalSwappableAmountResponse)
	err := c.cc.Invoke(ctx, "/lbm.fswap.v1.Query/TotalSwappableAmount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Swapped queries the current swapped status that includes a burnt amount of old coin and a minted amount of new
	// coin.
	Swapped(context.Context, *QuerySwappedRequest) (*QuerySwappedResponse, error)
	// TotalSwappableAmount queries the current swappable amount for new coin.
	TotalSwappableAmount(context.Context, *QueryTotalSwappableAmountRequest) (*QueryTotalSwappableAmountResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Swapped(ctx context.Context, req *QuerySwappedRequest) (*QuerySwappedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Swapped not implemented")
}
func (*UnimplementedQueryServer) TotalSwappableAmount(ctx context.Context, req *QueryTotalSwappableAmountRequest) (*QueryTotalSwappableAmountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TotalSwappableAmount not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Swapped_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuerySwappedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Swapped(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lbm.fswap.v1.Query/Swapped",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Swapped(ctx, req.(*QuerySwappedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TotalSwappableAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTotalSwappableAmountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TotalSwappableAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lbm.fswap.v1.Query/TotalSwappableAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TotalSwappableAmount(ctx, req.(*QueryTotalSwappableAmountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lbm.fswap.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Swapped",
			Handler:    _Query_Swapped_Handler,
		},
		{
			MethodName: "TotalSwappableAmount",
			Handler:    _Query_TotalSwappableAmount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lbm/fswap/v1/query.proto",
}

func (m *QuerySwappedRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySwappedRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySwappedRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QuerySwappedResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QuerySwappedResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QuerySwappedResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.NewCoinAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.OldCoinAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryTotalSwappableAmountRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTotalSwappableAmountRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTotalSwappableAmountRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryTotalSwappableAmountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryTotalSwappableAmountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryTotalSwappableAmountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.SwappableNewCoin.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QuerySwappedRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QuerySwappedResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.OldCoinAmount.Size()
	n += 1 + l + sovQuery(uint64(l))
	l = m.NewCoinAmount.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryTotalSwappableAmountRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryTotalSwappableAmountResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.SwappableNewCoin.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QuerySwappedRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QuerySwappedRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySwappedRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QuerySwappedResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QuerySwappedResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QuerySwappedResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OldCoinAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OldCoinAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewCoinAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NewCoinAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryTotalSwappableAmountRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryTotalSwappableAmountRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTotalSwappableAmountRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryTotalSwappableAmountResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryTotalSwappableAmountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryTotalSwappableAmountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SwappableNewCoin", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SwappableNewCoin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
