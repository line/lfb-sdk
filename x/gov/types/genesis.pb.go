// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lbm/gov/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the gov module's genesis state.
type GenesisState struct {
	// starting_proposal_id is the ID of the starting proposal.
	StartingProposalId uint64 `protobuf:"varint,1,opt,name=starting_proposal_id,json=startingProposalId,proto3" json:"starting_proposal_id,omitempty" yaml:"starting_proposal_id"`
	// deposits defines all the deposits present at genesis.
	Deposits Deposits `protobuf:"bytes,2,rep,name=deposits,proto3,castrepeated=Deposits" json:"deposits"`
	// votes defines all the votes present at genesis.
	Votes Votes `protobuf:"bytes,3,rep,name=votes,proto3,castrepeated=Votes" json:"votes"`
	// proposals defines all the proposals present at genesis.
	Proposals Proposals `protobuf:"bytes,4,rep,name=proposals,proto3,castrepeated=Proposals" json:"proposals"`
	// params defines all the paramaters of related to deposit.
	DepositParams DepositParams `protobuf:"bytes,5,opt,name=deposit_params,json=depositParams,proto3" json:"deposit_params" yaml:"deposit_params"`
	// params defines all the paramaters of related to voting.
	VotingParams VotingParams `protobuf:"bytes,6,opt,name=voting_params,json=votingParams,proto3" json:"voting_params" yaml:"voting_params"`
	// params defines all the paramaters of related to tally.
	TallyParams TallyParams `protobuf:"bytes,7,opt,name=tally_params,json=tallyParams,proto3" json:"tally_params" yaml:"tally_params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_ccdc73bc8806a7f1, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetStartingProposalId() uint64 {
	if m != nil {
		return m.StartingProposalId
	}
	return 0
}

func (m *GenesisState) GetDeposits() Deposits {
	if m != nil {
		return m.Deposits
	}
	return nil
}

func (m *GenesisState) GetVotes() Votes {
	if m != nil {
		return m.Votes
	}
	return nil
}

func (m *GenesisState) GetProposals() Proposals {
	if m != nil {
		return m.Proposals
	}
	return nil
}

func (m *GenesisState) GetDepositParams() DepositParams {
	if m != nil {
		return m.DepositParams
	}
	return DepositParams{}
}

func (m *GenesisState) GetVotingParams() VotingParams {
	if m != nil {
		return m.VotingParams
	}
	return VotingParams{}
}

func (m *GenesisState) GetTallyParams() TallyParams {
	if m != nil {
		return m.TallyParams
	}
	return TallyParams{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "lbm.gov.v1beta1.GenesisState")
}

func init() { proto.RegisterFile("lbm/gov/v1beta1/genesis.proto", fileDescriptor_ccdc73bc8806a7f1) }

var fileDescriptor_ccdc73bc8806a7f1 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcd, 0x8e, 0xd3, 0x30,
	0x14, 0x85, 0x13, 0xa6, 0x1d, 0x66, 0xdc, 0x96, 0x1f, 0xd3, 0x91, 0x32, 0x3f, 0x4d, 0xaa, 0x2c,
	0x50, 0x37, 0xc4, 0x9a, 0xb2, 0x9b, 0x15, 0x8a, 0x10, 0x08, 0x56, 0x43, 0x40, 0x2c, 0x10, 0x52,
	0x71, 0x88, 0x15, 0x22, 0x9c, 0x3a, 0x8a, 0x8d, 0x45, 0xdf, 0x82, 0xe7, 0xe0, 0x49, 0x66, 0x39,
	0x3b, 0x58, 0x15, 0xd4, 0xbe, 0x41, 0x9f, 0x00, 0xc5, 0x76, 0x20, 0x25, 0x9d, 0x5d, 0x72, 0xef,
	0x39, 0xdf, 0xb9, 0xbe, 0xba, 0x60, 0x44, 0xe3, 0x1c, 0xa5, 0x4c, 0x22, 0x79, 0x1e, 0x13, 0x81,
	0xcf, 0x51, 0x4a, 0xe6, 0x84, 0x67, 0x3c, 0x28, 0x4a, 0x26, 0x18, 0xbc, 0x4b, 0xe3, 0x3c, 0x48,
	0x99, 0x0c, 0x4c, 0xfb, 0x64, 0x98, 0xb2, 0x94, 0xa9, 0x1e, 0xaa, 0xbe, 0xb4, 0xec, 0xe4, 0xb8,
	0x45, 0x61, 0x52, 0xb7, 0xfc, 0x1f, 0x1d, 0xd0, 0x7f, 0xae, 0x99, 0xaf, 0x05, 0x16, 0x04, 0xbe,
	0x02, 0x43, 0x2e, 0x70, 0x29, 0xb2, 0x79, 0x3a, 0x2b, 0x4a, 0x56, 0x30, 0x8e, 0xe9, 0x2c, 0x4b,
	0x1c, 0x7b, 0x6c, 0x4f, 0x3a, 0xa1, 0xb7, 0x59, 0x7a, 0xa7, 0x0b, 0x9c, 0xd3, 0x0b, 0x7f, 0x97,
	0xca, 0x8f, 0x60, 0x5d, 0xbe, 0x34, 0xd5, 0x17, 0x09, 0x7c, 0x06, 0x0e, 0x12, 0x52, 0x30, 0x9e,
	0x09, 0xee, 0xdc, 0x1a, 0xef, 0x4d, 0x7a, 0x53, 0x27, 0xf8, 0x6f, 0xf0, 0xe0, 0xa9, 0x16, 0x84,
	0xf7, 0xae, 0x96, 0x9e, 0xf5, 0xfd, 0x97, 0x77, 0x60, 0x0a, 0x3c, 0xfa, 0xeb, 0x85, 0x17, 0xa0,
	0x2b, 0x99, 0x20, 0xdc, 0xd9, 0x53, 0x90, 0xa3, 0x16, 0xe4, 0x2d, 0x13, 0x24, 0x1c, 0x18, 0x42,
	0xb7, 0xfa, 0xe3, 0x91, 0xb6, 0xc0, 0x97, 0xe0, 0xb0, 0x9e, 0x93, 0x3b, 0x1d, 0xe5, 0x3f, 0x6e,
	0xf9, 0xeb, 0x99, 0xc3, 0xfb, 0x86, 0x71, 0x58, 0x57, 0x78, 0xf4, 0xcf, 0x0e, 0x13, 0x70, 0xc7,
	0xcc, 0x34, 0x2b, 0x70, 0x89, 0x73, 0xee, 0x74, 0xc7, 0xf6, 0xa4, 0x37, 0x75, 0x6f, 0x7a, 0xd5,
	0xa5, 0x52, 0x85, 0xa3, 0x8a, 0xba, 0x59, 0x7a, 0x47, 0x7a, 0x81, 0xdb, 0x0c, 0x3f, 0x1a, 0x24,
	0x4d, 0x35, 0xfc, 0x00, 0x06, 0x92, 0xe9, 0x05, 0xeb, 0x90, 0x7d, 0x15, 0x32, 0xda, 0xf5, 0xea,
	0x6a, 0xdf, 0x3a, 0xe3, 0xcc, 0x64, 0x0c, 0x75, 0xc6, 0x16, 0xc1, 0x8f, 0xfa, 0xb2, 0xa1, 0x85,
	0xef, 0x41, 0x5f, 0x60, 0x4a, 0x17, 0x75, 0xc0, 0x6d, 0x15, 0x70, 0xd6, 0x0a, 0x78, 0x53, 0x89,
	0x0c, 0xff, 0xd4, 0xf0, 0x1f, 0x68, 0x7e, 0xd3, 0xef, 0x47, 0x3d, 0xd1, 0x50, 0x3e, 0xb9, 0x5a,
	0xb9, 0xf6, 0xf5, 0xca, 0xb5, 0x7f, 0xaf, 0x5c, 0xfb, 0xdb, 0xda, 0xb5, 0xae, 0xd7, 0xae, 0xf5,
	0x73, 0xed, 0x5a, 0xef, 0x1e, 0xa6, 0x99, 0xf8, 0xf4, 0x25, 0x0e, 0x3e, 0xb2, 0x1c, 0xd1, 0x6c,
	0x4e, 0x10, 0x8d, 0xf3, 0x47, 0x3c, 0xf9, 0x8c, 0xe4, 0x14, 0x7d, 0x55, 0x77, 0x2a, 0x16, 0x05,
	0xe1, 0xf1, 0xbe, 0x3a, 0xd1, 0xc7, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x64, 0x6e, 0xb6, 0xcc,
	0x05, 0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.TallyParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size, err := m.VotingParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size, err := m.DepositParams.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Proposals) > 0 {
		for iNdEx := len(m.Proposals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Proposals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Votes) > 0 {
		for iNdEx := len(m.Votes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Votes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Deposits) > 0 {
		for iNdEx := len(m.Deposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.StartingProposalId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.StartingProposalId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.StartingProposalId != 0 {
		n += 1 + sovGenesis(uint64(m.StartingProposalId))
	}
	if len(m.Deposits) > 0 {
		for _, e := range m.Deposits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Votes) > 0 {
		for _, e := range m.Votes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Proposals) > 0 {
		for _, e := range m.Proposals {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	l = m.DepositParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.VotingParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
	l = m.TallyParams.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartingProposalId", wireType)
			}
			m.StartingProposalId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartingProposalId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deposits = append(m.Deposits, Deposit{})
			if err := m.Deposits[len(m.Deposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Votes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Votes = append(m.Votes, Vote{})
			if err := m.Votes[len(m.Votes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proposals", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proposals = append(m.Proposals, Proposal{})
			if err := m.Proposals[len(m.Proposals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DepositParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DepositParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VotingParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.VotingParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TallyParams", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TallyParams.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
