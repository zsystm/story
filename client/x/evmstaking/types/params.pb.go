// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: client/x/evmstaking/types/params.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params defines the parameters for the module.
type Params struct {
	MaxWithdrawalPerBlock      uint32 `protobuf:"varint,1,opt,name=max_withdrawal_per_block,json=maxWithdrawalPerBlock,proto3" json:"max_withdrawal_per_block,omitempty" yaml:"max_withdrawal_per_block"`
	MaxSweepPerBlock           uint32 `protobuf:"varint,2,opt,name=max_sweep_per_block,json=maxSweepPerBlock,proto3" json:"max_sweep_per_block,omitempty" yaml:"max_sweep_per_block"`
	MinPartialWithdrawalAmount uint64 `protobuf:"varint,3,opt,name=min_partial_withdrawal_amount,json=minPartialWithdrawalAmount,proto3" json:"min_partial_withdrawal_amount,omitempty" yaml:"min_partial_withdrawal_amount"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_dddf03d6f1b350f8, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMaxWithdrawalPerBlock() uint32 {
	if m != nil {
		return m.MaxWithdrawalPerBlock
	}
	return 0
}

func (m *Params) GetMaxSweepPerBlock() uint32 {
	if m != nil {
		return m.MaxSweepPerBlock
	}
	return 0
}

func (m *Params) GetMinPartialWithdrawalAmount() uint64 {
	if m != nil {
		return m.MinPartialWithdrawalAmount
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "client.x.evmstaking.types.Params")
}

func init() {
	proto.RegisterFile("client/x/evmstaking/types/params.proto", fileDescriptor_dddf03d6f1b350f8)
}

var fileDescriptor_dddf03d6f1b350f8 = []byte{
	// 288 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4b, 0xce, 0xc9, 0x4c,
	0xcd, 0x2b, 0xd1, 0xaf, 0xd0, 0x4f, 0x2d, 0xcb, 0x2d, 0x2e, 0x49, 0xcc, 0xce, 0xcc, 0x4b, 0xd7,
	0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0xd6, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x92, 0x84, 0xa8, 0xd3, 0xab, 0xd0, 0x43, 0xa8, 0xd3, 0x03, 0xab, 0x93, 0x12,
	0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xab, 0xd2, 0x07, 0xb1, 0x20, 0x1a, 0x94, 0x16, 0x31, 0x71, 0xb1,
	0x05, 0x80, 0x4d, 0x10, 0x8a, 0xe1, 0x92, 0xc8, 0x4d, 0xac, 0x88, 0x2f, 0xcf, 0x2c, 0xc9, 0x48,
	0x29, 0x4a, 0x2c, 0x4f, 0xcc, 0x89, 0x2f, 0x48, 0x2d, 0x8a, 0x4f, 0xca, 0xc9, 0x4f, 0xce, 0x96,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x75, 0x52, 0xfe, 0x74, 0x4f, 0x5e, 0xbe, 0x32, 0x31, 0x37, 0xc7,
	0x4a, 0x09, 0x97, 0x4a, 0xa5, 0x20, 0xd1, 0xdc, 0xc4, 0x8a, 0x70, 0xb8, 0x4c, 0x40, 0x6a, 0x91,
	0x13, 0x48, 0x5c, 0xc8, 0x97, 0x4b, 0x18, 0xa4, 0xa7, 0xb8, 0x3c, 0x35, 0xb5, 0x00, 0xc9, 0x60,
	0x26, 0xb0, 0xc1, 0x72, 0x9f, 0xee, 0xc9, 0x4b, 0x21, 0x0c, 0x46, 0x53, 0xa4, 0x14, 0x24, 0x90,
	0x9b, 0x58, 0x11, 0x0c, 0x12, 0x84, 0x1b, 0x97, 0xcd, 0x25, 0x9b, 0x9b, 0x99, 0x17, 0x5f, 0x90,
	0x58, 0x54, 0x92, 0x99, 0x98, 0x83, 0xec, 0x94, 0xc4, 0xdc, 0xfc, 0xd2, 0xbc, 0x12, 0x09, 0x66,
	0x05, 0x46, 0x0d, 0x16, 0x27, 0x8d, 0x4f, 0xf7, 0xe4, 0x55, 0xa0, 0x06, 0xe3, 0x53, 0xae, 0x14,
	0x24, 0x95, 0x9b, 0x99, 0x17, 0x00, 0x91, 0x46, 0xb8, 0xde, 0x11, 0x2c, 0xe9, 0x64, 0x7c, 0xe2,
	0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70,
	0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x92, 0x38, 0xe3, 0x25, 0x89, 0x0d, 0x1c,
	0xc0, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x45, 0xda, 0xba, 0xbb, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MinPartialWithdrawalAmount != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinPartialWithdrawalAmount))
		i--
		dAtA[i] = 0x18
	}
	if m.MaxSweepPerBlock != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxSweepPerBlock))
		i--
		dAtA[i] = 0x10
	}
	if m.MaxWithdrawalPerBlock != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxWithdrawalPerBlock))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MaxWithdrawalPerBlock != 0 {
		n += 1 + sovParams(uint64(m.MaxWithdrawalPerBlock))
	}
	if m.MaxSweepPerBlock != 0 {
		n += 1 + sovParams(uint64(m.MaxSweepPerBlock))
	}
	if m.MinPartialWithdrawalAmount != 0 {
		n += 1 + sovParams(uint64(m.MinPartialWithdrawalAmount))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxWithdrawalPerBlock", wireType)
			}
			m.MaxWithdrawalPerBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxWithdrawalPerBlock |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxSweepPerBlock", wireType)
			}
			m.MaxSweepPerBlock = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxSweepPerBlock |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinPartialWithdrawalAmount", wireType)
			}
			m.MinPartialWithdrawalAmount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinPartialWithdrawalAmount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
