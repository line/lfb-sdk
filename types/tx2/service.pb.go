// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lbm/tx/v1beta1/service.proto

package tx2

import (
	context "context"
	fmt "fmt"
	query "github.com/Finschia/finschia-sdk/types/query"
	tx "github.com/Finschia/finschia-sdk/types/tx"
	_ "github.com/Finschia/ostracon/proto/ostracon/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	types "github.com/tendermint/tendermint/proto/tendermint/types"
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
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// GetBlockWithTxsRequest is the request type for the Service.GetBlockWithTxs
// RPC method.
//
// Since: finschia-sdk 0.47.0
type GetBlockWithTxsRequest struct {
	// height is the height of the block to query.
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	// pagination defines a pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *GetBlockWithTxsRequest) Reset()         { *m = GetBlockWithTxsRequest{} }
func (m *GetBlockWithTxsRequest) String() string { return proto.CompactTextString(m) }
func (*GetBlockWithTxsRequest) ProtoMessage()    {}
func (*GetBlockWithTxsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fc6bd78191bf1b3, []int{0}
}
func (m *GetBlockWithTxsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetBlockWithTxsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetBlockWithTxsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetBlockWithTxsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockWithTxsRequest.Merge(m, src)
}
func (m *GetBlockWithTxsRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetBlockWithTxsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockWithTxsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockWithTxsRequest proto.InternalMessageInfo

func (m *GetBlockWithTxsRequest) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *GetBlockWithTxsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// GetBlockWithTxsResponse is the response type for the Service.GetBlockWithTxs method.
//
// Since: finschia-sdk 0.47.0
type GetBlockWithTxsResponse struct {
	// txs are the transactions in the block.
	Txs     []*tx.Tx       `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
	BlockId *types.BlockID `protobuf:"bytes,2,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	Block   *types.Block   `protobuf:"bytes,3,opt,name=block,proto3" json:"block,omitempty"`
	// pagination defines a pagination for the response.
	Pagination *query.PageResponse `protobuf:"bytes,4,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *GetBlockWithTxsResponse) Reset()         { *m = GetBlockWithTxsResponse{} }
func (m *GetBlockWithTxsResponse) String() string { return proto.CompactTextString(m) }
func (*GetBlockWithTxsResponse) ProtoMessage()    {}
func (*GetBlockWithTxsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6fc6bd78191bf1b3, []int{1}
}
func (m *GetBlockWithTxsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetBlockWithTxsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetBlockWithTxsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetBlockWithTxsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetBlockWithTxsResponse.Merge(m, src)
}
func (m *GetBlockWithTxsResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetBlockWithTxsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetBlockWithTxsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetBlockWithTxsResponse proto.InternalMessageInfo

func (m *GetBlockWithTxsResponse) GetTxs() []*tx.Tx {
	if m != nil {
		return m.Txs
	}
	return nil
}

func (m *GetBlockWithTxsResponse) GetBlockId() *types.BlockID {
	if m != nil {
		return m.BlockId
	}
	return nil
}

func (m *GetBlockWithTxsResponse) GetBlock() *types.Block {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *GetBlockWithTxsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*GetBlockWithTxsRequest)(nil), "lbm.tx.v1beta1.GetBlockWithTxsRequest")
	golang_proto.RegisterType((*GetBlockWithTxsRequest)(nil), "lbm.tx.v1beta1.GetBlockWithTxsRequest")
	proto.RegisterType((*GetBlockWithTxsResponse)(nil), "lbm.tx.v1beta1.GetBlockWithTxsResponse")
	golang_proto.RegisterType((*GetBlockWithTxsResponse)(nil), "lbm.tx.v1beta1.GetBlockWithTxsResponse")
}

func init() { proto.RegisterFile("lbm/tx/v1beta1/service.proto", fileDescriptor_6fc6bd78191bf1b3) }
func init() {
	golang_proto.RegisterFile("lbm/tx/v1beta1/service.proto", fileDescriptor_6fc6bd78191bf1b3)
}

var fileDescriptor_6fc6bd78191bf1b3 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x33, 0x8d, 0xb6, 0x32, 0x05, 0x85, 0x41, 0xdb, 0x18, 0xc2, 0x12, 0x82, 0x34, 0x21,
	0xd0, 0x19, 0x1a, 0xfd, 0x04, 0x45, 0x5a, 0x7b, 0x93, 0xb5, 0x20, 0x78, 0x91, 0xd9, 0xcd, 0xb8,
	0x3b, 0x34, 0x3b, 0xb3, 0xdd, 0x79, 0x29, 0x53, 0xc4, 0x8b, 0x1f, 0x40, 0x04, 0x6f, 0x7e, 0x1a,
	0x8f, 0x1e, 0x0b, 0x5e, 0x3c, 0x4a, 0xe2, 0x07, 0xf0, 0x23, 0x48, 0x66, 0x26, 0x66, 0x5b, 0x5b,
	0x7a, 0x59, 0xde, 0xce, 0xfb, 0xbf, 0xf7, 0xe6, 0x37, 0xef, 0x8f, 0x3b, 0x93, 0xa4, 0x60, 0x60,
	0xd9, 0xd9, 0x5e, 0x22, 0x80, 0xef, 0x31, 0x23, 0xaa, 0x33, 0x99, 0x0a, 0x5a, 0x56, 0x1a, 0x34,
	0xb9, 0x3f, 0x49, 0x0a, 0x0a, 0x96, 0x86, 0x6c, 0xbb, 0x93, 0x69, 0x9d, 0x4d, 0x04, 0xe3, 0xa5,
	0x64, 0x5c, 0x29, 0x0d, 0x1c, 0xa4, 0x56, 0xc6, 0xab, 0xdb, 0x0f, 0x33, 0x9d, 0x69, 0x17, 0xb2,
	0x45, 0x14, 0x4e, 0xdb, 0xa9, 0x36, 0x85, 0x36, 0xf5, 0x21, 0x60, 0x43, 0x6e, 0x18, 0x72, 0x09,
	0x37, 0x82, 0x9d, 0x4e, 0x45, 0x75, 0xfe, 0x4f, 0x53, 0xf2, 0x4c, 0x2a, 0xd7, 0x3e, 0x68, 0x3b,
	0x20, 0xd4, 0x58, 0x54, 0x85, 0x54, 0xc0, 0xe0, 0xbc, 0x14, 0x86, 0x25, 0x13, 0x9d, 0x9e, 0x2c,
	0xa7, 0x68, 0x03, 0x15, 0x4f, 0xb5, 0x0a, 0x39, 0xf7, 0xbd, 0xb1, 0xb2, 0x96, 0xed, 0x59, 0xbc,
	0x75, 0x28, 0x60, 0x7f, 0xd1, 0xeb, 0xb5, 0x84, 0xfc, 0xd8, 0x9a, 0x58, 0x9c, 0x4e, 0x85, 0x01,
	0xb2, 0x85, 0xd7, 0x73, 0x21, 0xb3, 0x1c, 0x5a, 0xa8, 0x8b, 0x06, 0xcd, 0x38, 0xfc, 0x91, 0x03,
	0x8c, 0x57, 0xb7, 0x6b, 0xad, 0x75, 0xd1, 0x60, 0x73, 0xb4, 0x43, 0x3d, 0x0a, 0x5d, 0xa0, 0x50,
	0x87, 0xb2, 0x7c, 0x35, 0xfa, 0x92, 0x67, 0x22, 0xf4, 0x8c, 0x6b, 0x95, 0xbd, 0x3f, 0x08, 0x6f,
	0xff, 0x37, 0xda, 0x94, 0x5a, 0x19, 0x41, 0xfa, 0xb8, 0x09, 0xd6, 0xb4, 0x50, 0xb7, 0x39, 0xd8,
	0x1c, 0x3d, 0x5a, 0x36, 0x5f, 0xad, 0x82, 0x1e, 0xdb, 0x78, 0xa1, 0x20, 0xcf, 0xf0, 0x3d, 0xf7,
	0x0e, 0x6f, 0xe5, 0x38, 0x5c, 0xe5, 0x31, 0x5d, 0xf1, 0x52, 0x4f, 0xea, 0x46, 0x1c, 0x3d, 0x8f,
	0x37, 0x9c, 0xf4, 0x68, 0x4c, 0x76, 0xf1, 0x5d, 0x17, 0xb6, 0x9a, 0xae, 0x64, 0xfb, 0x86, 0x92,
	0xd8, 0xab, 0xc8, 0xe1, 0x25, 0xe2, 0x3b, 0xae, 0xa6, 0x7f, 0x2b, 0xb1, 0x47, 0xa9, 0x23, 0x8f,
	0xbe, 0x22, 0xbc, 0xf1, 0xca, 0x5b, 0x8c, 0x7c, 0x42, 0xf8, 0xc1, 0x15, 0x7c, 0xb2, 0x43, 0x2f,
	0x3b, 0x8e, 0x5e, 0xbf, 0x9a, 0x76, 0xff, 0x56, 0x9d, 0x1f, 0xde, 0x1b, 0x7e, 0xfc, 0xf1, 0xfb,
	0xcb, 0xda, 0x13, 0xd2, 0x63, 0x57, 0x8c, 0x0e, 0x36, 0x18, 0x88, 0xbd, 0xf7, 0x6b, 0xfd, 0xb0,
	0xff, 0xe2, 0xfb, 0x2c, 0x42, 0x17, 0xb3, 0x08, 0xfd, 0x9a, 0x45, 0xe8, 0xf3, 0x3c, 0x6a, 0x7c,
	0x9b, 0x47, 0xe8, 0x62, 0x1e, 0x35, 0x7e, 0xce, 0xa3, 0xc6, 0x9b, 0x61, 0x26, 0x21, 0x9f, 0x26,
	0x34, 0xd5, 0x05, 0x3b, 0x90, 0xca, 0xa4, 0xb9, 0xe4, 0xec, 0x5d, 0x08, 0x76, 0xcd, 0xf8, 0x64,
	0xe9, 0x2d, 0x3b, 0x4a, 0xd6, 0x9d, 0xb5, 0x9e, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x55,
	0xa3, 0x0d, 0x5e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceClient interface {
	// GetBlockWithTxs fetches a block with decoded txs.
	//
	// Since: finschia-sdk 0.47.0
	GetBlockWithTxs(ctx context.Context, in *GetBlockWithTxsRequest, opts ...grpc.CallOption) (*GetBlockWithTxsResponse, error)
}

type serviceClient struct {
	cc grpc1.ClientConn
}

func NewServiceClient(cc grpc1.ClientConn) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) GetBlockWithTxs(ctx context.Context, in *GetBlockWithTxsRequest, opts ...grpc.CallOption) (*GetBlockWithTxsResponse, error) {
	out := new(GetBlockWithTxsResponse)
	err := c.cc.Invoke(ctx, "/lbm.tx.v1beta1.Service/GetBlockWithTxs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
type ServiceServer interface {
	// GetBlockWithTxs fetches a block with decoded txs.
	//
	// Since: finschia-sdk 0.47.0
	GetBlockWithTxs(context.Context, *GetBlockWithTxsRequest) (*GetBlockWithTxsResponse, error)
}

// UnimplementedServiceServer can be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) GetBlockWithTxs(ctx context.Context, req *GetBlockWithTxsRequest) (*GetBlockWithTxsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockWithTxs not implemented")
}

func RegisterServiceServer(s grpc1.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_GetBlockWithTxs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockWithTxsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetBlockWithTxs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lbm.tx.v1beta1.Service/GetBlockWithTxs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetBlockWithTxs(ctx, req.(*GetBlockWithTxsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lbm.tx.v1beta1.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlockWithTxs",
			Handler:    _Service_GetBlockWithTxs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lbm/tx/v1beta1/service.proto",
}

func (m *GetBlockWithTxsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetBlockWithTxsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetBlockWithTxsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Height != 0 {
		i = encodeVarintService(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *GetBlockWithTxsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetBlockWithTxsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetBlockWithTxsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Block != nil {
		{
			size, err := m.Block.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.BlockId != nil {
		{
			size, err := m.BlockId.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Txs) > 0 {
		for iNdEx := len(m.Txs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Txs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintService(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	offset -= sovService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GetBlockWithTxsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovService(uint64(m.Height))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *GetBlockWithTxsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Txs) > 0 {
		for _, e := range m.Txs {
			l = e.Size()
			n += 1 + l + sovService(uint64(l))
		}
	}
	if m.BlockId != nil {
		l = m.BlockId.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.Block != nil {
		l = m.Block.Size()
		n += 1 + l + sovService(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func sovService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GetBlockWithTxsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: GetBlockWithTxsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetBlockWithTxsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func (m *GetBlockWithTxsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
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
			return fmt.Errorf("proto: GetBlockWithTxsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetBlockWithTxsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Txs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Txs = append(m.Txs, &tx.Tx{})
			if err := m.Txs[len(m.Txs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BlockId == nil {
				m.BlockId = &types.BlockID{}
			}
			if err := m.BlockId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Block", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Block == nil {
				m.Block = &types.Block{}
			}
			if err := m.Block.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
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
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
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
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
					return 0, ErrIntOverflowService
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
				return 0, ErrInvalidLengthService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupService = fmt.Errorf("proto: unexpected end of group")
)
