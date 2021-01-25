// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lbm/authz/v1/authz.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	types1 "github.com/line/lbm-sdk/codec/types"
	github_com_line_lbm_sdk_types "github.com/line/lbm-sdk/types"
	types "github.com/line/lbm-sdk/types"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// SendAuthorization allows the grantee to spend up to spend_limit coins from
// the granter's account.
type SendAuthorization struct {
	SpendLimit github_com_line_lbm_sdk_types.Coins `protobuf:"bytes,1,rep,name=spend_limit,json=spendLimit,proto3,castrepeated=github.com/line/lbm-sdk/types.Coins" json:"spend_limit"`
}

func (m *SendAuthorization) Reset()         { *m = SendAuthorization{} }
func (m *SendAuthorization) String() string { return proto.CompactTextString(m) }
func (*SendAuthorization) ProtoMessage()    {}
func (*SendAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa5041a8d6ee0eda, []int{0}
}
func (m *SendAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendAuthorization.Merge(m, src)
}
func (m *SendAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *SendAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_SendAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_SendAuthorization proto.InternalMessageInfo

func (m *SendAuthorization) GetSpendLimit() github_com_line_lbm_sdk_types.Coins {
	if m != nil {
		return m.SpendLimit
	}
	return nil
}

// GenericAuthorization gives the grantee unrestricted permissions to execute
// the provided method on behalf of the granter's account.
type GenericAuthorization struct {
	// method name to grant unrestricted permissions to execute
	// Note: MethodName() is already a method on `GenericAuthorization` type,
	// we need some custom naming here so using `MessageName`
	MessageName string `protobuf:"bytes,1,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
}

func (m *GenericAuthorization) Reset()         { *m = GenericAuthorization{} }
func (m *GenericAuthorization) String() string { return proto.CompactTextString(m) }
func (*GenericAuthorization) ProtoMessage()    {}
func (*GenericAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa5041a8d6ee0eda, []int{1}
}
func (m *GenericAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenericAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenericAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenericAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenericAuthorization.Merge(m, src)
}
func (m *GenericAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *GenericAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_GenericAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_GenericAuthorization proto.InternalMessageInfo

func (m *GenericAuthorization) GetMessageName() string {
	if m != nil {
		return m.MessageName
	}
	return ""
}

// AuthorizationGrant gives permissions to execute
// the provide method with expiration time.
type AuthorizationGrant struct {
	Authorization *types1.Any `protobuf:"bytes,1,opt,name=authorization,proto3" json:"authorization,omitempty"`
	Expiration    time.Time   `protobuf:"bytes,2,opt,name=expiration,proto3,stdtime" json:"expiration"`
}

func (m *AuthorizationGrant) Reset()         { *m = AuthorizationGrant{} }
func (m *AuthorizationGrant) String() string { return proto.CompactTextString(m) }
func (*AuthorizationGrant) ProtoMessage()    {}
func (*AuthorizationGrant) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa5041a8d6ee0eda, []int{2}
}
func (m *AuthorizationGrant) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthorizationGrant) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthorizationGrant.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthorizationGrant) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationGrant.Merge(m, src)
}
func (m *AuthorizationGrant) XXX_Size() int {
	return m.Size()
}
func (m *AuthorizationGrant) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizationGrant.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizationGrant proto.InternalMessageInfo

func (m *AuthorizationGrant) GetAuthorization() *types1.Any {
	if m != nil {
		return m.Authorization
	}
	return nil
}

func (m *AuthorizationGrant) GetExpiration() time.Time {
	if m != nil {
		return m.Expiration
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*SendAuthorization)(nil), "lbm.authz.v1.SendAuthorization")
	proto.RegisterType((*GenericAuthorization)(nil), "lbm.authz.v1.GenericAuthorization")
	proto.RegisterType((*AuthorizationGrant)(nil), "lbm.authz.v1.AuthorizationGrant")
}

func init() { proto.RegisterFile("lbm/authz/v1/authz.proto", fileDescriptor_aa5041a8d6ee0eda) }

var fileDescriptor_aa5041a8d6ee0eda = []byte{
	// 413 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x41, 0x6b, 0xd4, 0x40,
	0x14, 0xc7, 0x33, 0x0a, 0xa2, 0x13, 0x8b, 0x24, 0x2c, 0xb2, 0xee, 0x21, 0x29, 0x15, 0xa1, 0x20,
	0x9d, 0x71, 0xeb, 0xcd, 0x8b, 0x34, 0x0a, 0xbd, 0x58, 0x0f, 0xd1, 0x93, 0x1e, 0xc2, 0x24, 0x19,
	0x93, 0xc1, 0xcc, 0x4c, 0xc8, 0x4c, 0x96, 0x6e, 0xbf, 0x80, 0xd7, 0x7e, 0x0d, 0x3d, 0xfb, 0x21,
	0x8a, 0xa7, 0xe2, 0xc9, 0x53, 0x2b, 0xd9, 0x2f, 0x22, 0x33, 0xb3, 0x0b, 0xcd, 0xaa, 0xb7, 0xf7,
	0xde, 0xfc, 0xff, 0xbf, 0xf7, 0xf8, 0x27, 0x70, 0xda, 0xe4, 0x1c, 0x93, 0x5e, 0xd7, 0x67, 0x78,
	0x31, 0x77, 0x05, 0x6a, 0x3b, 0xa9, 0x65, 0x78, 0xbf, 0xc9, 0x39, 0x72, 0x83, 0xc5, 0x7c, 0xf6,
	0xd0, 0xe8, 0x72, 0xa2, 0xa8, 0x91, 0x15, 0x92, 0x09, 0xa7, 0x9a, 0x3d, 0x2a, 0xa4, 0xe2, 0x52,
	0x65, 0xb6, 0xc3, 0xae, 0x59, 0x3f, 0xc5, 0x95, 0x94, 0x55, 0x43, 0xb1, 0xed, 0xf2, 0xfe, 0x13,
	0xd6, 0x8c, 0x53, 0xa5, 0x09, 0x6f, 0xd7, 0x82, 0x49, 0x25, 0x2b, 0xe9, 0x8c, 0xa6, 0xda, 0x10,
	0xb7, 0x6d, 0x44, 0x2c, 0xdd, 0xd3, 0xde, 0x17, 0x00, 0x83, 0x77, 0x54, 0x94, 0x47, 0xbd, 0xae,
	0x65, 0xc7, 0xce, 0x88, 0x66, 0x52, 0x84, 0x19, 0xf4, 0x55, 0x4b, 0x45, 0x99, 0x35, 0x8c, 0x33,
	0x3d, 0x05, 0xbb, 0xb7, 0xf7, 0xfd, 0xc3, 0x00, 0x99, 0xf3, 0xcd, 0xc1, 0x68, 0x31, 0x47, 0xaf,
	0x24, 0x13, 0xc9, 0xd3, 0x8b, 0xab, 0xd8, 0xfb, 0x76, 0x1d, 0x3f, 0xae, 0x98, 0xae, 0xfb, 0x1c,
	0x15, 0x92, 0xe3, 0x86, 0x09, 0x8a, 0x9b, 0x9c, 0x1f, 0xa8, 0xf2, 0x33, 0xd6, 0xcb, 0x96, 0x2a,
	0xab, 0x55, 0x29, 0xb4, 0xc8, 0x37, 0x86, 0xf8, 0x22, 0xf8, 0xf9, 0xfd, 0x60, 0x67, 0xb4, 0x73,
	0xef, 0x23, 0x9c, 0x1c, 0x53, 0x41, 0x3b, 0x56, 0x8c, 0x6f, 0x79, 0x06, 0x7d, 0x4e, 0x75, 0x2d,
	0xcb, 0x4c, 0x10, 0x4e, 0xa7, 0x60, 0x17, 0xec, 0xdf, 0x4b, 0x1e, 0x0c, 0x57, 0xb1, 0x7f, 0x42,
	0x95, 0x22, 0x15, 0x7d, 0x4b, 0x38, 0x4d, 0xa1, 0xd3, 0x98, 0xfa, 0x5f, 0xf0, 0xaf, 0x00, 0x86,
	0xa3, 0xc9, 0x71, 0x47, 0x84, 0x0e, 0x4f, 0xe0, 0x0e, 0xb9, 0x39, 0xb5, 0x74, 0xff, 0x70, 0x82,
	0x5c, 0x60, 0x68, 0x13, 0x18, 0x3a, 0x12, 0xcb, 0x24, 0xf8, 0xb1, 0x8d, 0x4d, 0xc7, 0xee, 0xf0,
	0x35, 0x84, 0xf4, 0xb4, 0x65, 0x9d, 0x63, 0xdd, 0xb2, 0xac, 0xd9, 0x5f, 0xac, 0xf7, 0x9b, 0x6f,
	0x96, 0xdc, 0x35, 0xf1, 0x9d, 0x5f, 0xc7, 0x20, 0xbd, 0xe1, 0x4b, 0x5e, 0x5e, 0x0c, 0x11, 0xb8,
	0x1c, 0x22, 0xf0, 0x7b, 0x88, 0xc0, 0xf9, 0x2a, 0xf2, 0x2e, 0x57, 0x91, 0xf7, 0x6b, 0x15, 0x79,
	0x1f, 0x9e, 0xfc, 0x2f, 0xe6, 0xd3, 0xf5, 0xff, 0x66, 0xe3, 0xce, 0xef, 0xd8, 0x55, 0xcf, 0xff,
	0x04, 0x00, 0x00, 0xff, 0xff, 0xde, 0xb0, 0xde, 0x7c, 0x89, 0x02, 0x00, 0x00,
}

func (m *SendAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SendAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SpendLimit) > 0 {
		for iNdEx := len(m.SpendLimit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SpendLimit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAuthz(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GenericAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenericAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenericAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MessageName) > 0 {
		i -= len(m.MessageName)
		copy(dAtA[i:], m.MessageName)
		i = encodeVarintAuthz(dAtA, i, uint64(len(m.MessageName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthorizationGrant) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthorizationGrant) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthorizationGrant) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Expiration, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiration):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintAuthz(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if m.Authorization != nil {
		{
			size, err := m.Authorization.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuthz(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuthz(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuthz(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SendAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SpendLimit) > 0 {
		for _, e := range m.SpendLimit {
			l = e.Size()
			n += 1 + l + sovAuthz(uint64(l))
		}
	}
	return n
}

func (m *GenericAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MessageName)
	if l > 0 {
		n += 1 + l + sovAuthz(uint64(l))
	}
	return n
}

func (m *AuthorizationGrant) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Authorization != nil {
		l = m.Authorization.Size()
		n += 1 + l + sovAuthz(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Expiration)
	n += 1 + l + sovAuthz(uint64(l))
	return n
}

func sovAuthz(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuthz(x uint64) (n int) {
	return sovAuthz(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SendAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: SendAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpendLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpendLimit = append(m.SpendLimit, types.Coin{})
			if err := m.SpendLimit[len(m.SpendLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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
func (m *GenericAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: GenericAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenericAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MessageName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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
func (m *AuthorizationGrant) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuthz
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
			return fmt.Errorf("proto: AuthorizationGrant: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthorizationGrant: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authorization", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Authorization == nil {
				m.Authorization = &types1.Any{}
			}
			if err := m.Authorization.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiration", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuthz
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
				return ErrInvalidLengthAuthz
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuthz
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Expiration, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuthz(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuthz
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
func skipAuthz(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
					return 0, ErrIntOverflowAuthz
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
				return 0, ErrInvalidLengthAuthz
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuthz
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuthz
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuthz        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuthz          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuthz = fmt.Errorf("proto: unexpected end of group")
)
