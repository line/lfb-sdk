// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lbm/fswap/v1/event.proto

package types

import (
	fmt "fmt"
	types "github.com/Finschia/finschia-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type EventSwapCoins struct {
	// holder's address
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// amount of the old currency
	OldCoinAmount types.Coin `protobuf:"bytes,2,opt,name=old_coin_amount,json=oldCoinAmount,proto3,castrepeated=github.com/Finschia/finschia-sdk/types.Coin" json:"old_coin_amount"`
	// amount of the new currency
	NewCoinAmount types.Coin `protobuf:"bytes,3,opt,name=new_coin_amount,json=newCoinAmount,proto3,castrepeated=github.com/Finschia/finschia-sdk/types.Coin" json:"new_coin_amount"`
}

func (m *EventSwapCoins) Reset()         { *m = EventSwapCoins{} }
func (m *EventSwapCoins) String() string { return proto.CompactTextString(m) }
func (*EventSwapCoins) ProtoMessage()    {}
func (*EventSwapCoins) Descriptor() ([]byte, []int) {
	return fileDescriptor_92d5edbd64a725af, []int{0}
}
func (m *EventSwapCoins) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventSwapCoins) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventSwapCoins.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventSwapCoins) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSwapCoins.Merge(m, src)
}
func (m *EventSwapCoins) XXX_Size() int {
	return m.Size()
}
func (m *EventSwapCoins) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSwapCoins.DiscardUnknown(m)
}

var xxx_messageInfo_EventSwapCoins proto.InternalMessageInfo

func (m *EventSwapCoins) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *EventSwapCoins) GetOldCoinAmount() types.Coin {
	if m != nil {
		return m.OldCoinAmount
	}
	return types.Coin{}
}

func (m *EventSwapCoins) GetNewCoinAmount() types.Coin {
	if m != nil {
		return m.NewCoinAmount
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*EventSwapCoins)(nil), "lbm.fswap.v1.EventSwapCoins")
}

func init() { proto.RegisterFile("lbm/fswap/v1/event.proto", fileDescriptor_92d5edbd64a725af) }

var fileDescriptor_92d5edbd64a725af = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x91, 0x3f, 0x4e, 0xc3, 0x30,
	0x1c, 0x85, 0xe3, 0x22, 0x81, 0x08, 0xff, 0xa4, 0x88, 0x21, 0x74, 0x70, 0x2b, 0xa6, 0x4a, 0x08,
	0x5b, 0xa1, 0x27, 0xa0, 0x08, 0xc4, 0x5c, 0x36, 0x96, 0xca, 0x4e, 0xdc, 0x34, 0x22, 0xf1, 0x2f,
	0xaa, 0x5d, 0x07, 0x6e, 0xc1, 0xcc, 0x11, 0x38, 0x49, 0xc7, 0x8e, 0x4c, 0x80, 0x92, 0x8b, 0x20,
	0x3b, 0x41, 0x88, 0x89, 0x89, 0xed, 0x59, 0xcf, 0xef, 0xfb, 0x24, 0xdb, 0x0f, 0x73, 0x5e, 0xd0,
	0xb9, 0xaa, 0x58, 0x49, 0x4d, 0x44, 0x85, 0x11, 0x52, 0x93, 0x72, 0x09, 0x1a, 0x82, 0xfd, 0x9c,
	0x17, 0xc4, 0x35, 0xc4, 0x44, 0xfd, 0xe3, 0x14, 0x52, 0x70, 0x05, 0xb5, 0xa9, 0xbd, 0xd3, 0xc7,
	0x31, 0xa8, 0x02, 0x14, 0xe5, 0x4c, 0x09, 0x6a, 0x22, 0x2e, 0x34, 0x8b, 0x68, 0x0c, 0x99, 0x6c,
	0xfb, 0xd3, 0x97, 0x9e, 0x7f, 0x78, 0x6d, 0x99, 0x77, 0x15, 0x2b, 0xaf, 0x20, 0x93, 0x2a, 0x08,
	0xfd, 0x1d, 0x96, 0x24, 0x4b, 0xa1, 0x54, 0x88, 0x86, 0x68, 0xb4, 0x3b, 0xfd, 0x3e, 0x06, 0xc6,
	0x3f, 0x82, 0x3c, 0x99, 0xd9, 0xf9, 0x8c, 0x15, 0xb0, 0x92, 0x3a, 0xec, 0x0d, 0xd1, 0x68, 0xef,
	0xe2, 0x84, 0xb4, 0x1a, 0x62, 0x35, 0xa4, 0xd3, 0x10, 0x8b, 0x9b, 0x8c, 0xd7, 0xef, 0x03, 0xef,
	0xf5, 0x63, 0x70, 0x96, 0x66, 0x7a, 0xb1, 0xe2, 0x24, 0x86, 0x82, 0xde, 0x64, 0x52, 0xc5, 0x8b,
	0x8c, 0xd1, 0x79, 0x17, 0xce, 0x55, 0xf2, 0x40, 0xf5, 0x53, 0x29, 0x94, 0x1b, 0x4d, 0x0f, 0x20,
	0x4f, 0x6c, 0xb8, 0x74, 0x12, 0xeb, 0x95, 0xa2, 0xfa, 0xe5, 0xdd, 0xfa, 0x1f, 0xaf, 0x14, 0xd5,
	0x8f, 0x77, 0x72, 0xbb, 0xae, 0x31, 0xda, 0xd4, 0x18, 0x7d, 0xd6, 0x18, 0x3d, 0x37, 0xd8, 0xdb,
	0x34, 0xd8, 0x7b, 0x6b, 0xb0, 0x77, 0x4f, 0xfe, 0xa4, 0x3e, 0x76, 0x7f, 0xe6, 0xe8, 0x7c, 0xdb,
	0xbd, 0xf6, 0xf8, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x08, 0x6b, 0x7c, 0xcd, 0x01, 0x00, 0x00,
}

func (m *EventSwapCoins) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSwapCoins) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventSwapCoins) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.OldCoinAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintEvent(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventSwapCoins) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = m.OldCoinAmount.Size()
	n += 1 + l + sovEvent(uint64(l))
	l = m.NewCoinAmount.Size()
	n += 1 + l + sovEvent(uint64(l))
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventSwapCoins) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventSwapCoins: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSwapCoins: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OldCoinAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OldCoinAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewCoinAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
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
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
