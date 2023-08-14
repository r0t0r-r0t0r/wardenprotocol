// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ethermint/types/v1/dynamic_fee.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// ExtensionOptionDynamicFeeTx is an extension option that specifies the
// maxPrioPrice for cosmos tx
type ExtensionOptionDynamicFeeTx struct {
	// max_priority_price is the same as `max_priority_fee_per_gas` in eip-1559
	// spec
	MaxPriorityPrice github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=max_priority_price,json=maxPriorityPrice,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"max_priority_price"`
}

func (m *ExtensionOptionDynamicFeeTx) Reset()         { *m = ExtensionOptionDynamicFeeTx{} }
func (m *ExtensionOptionDynamicFeeTx) String() string { return proto.CompactTextString(m) }
func (*ExtensionOptionDynamicFeeTx) ProtoMessage()    {}
func (*ExtensionOptionDynamicFeeTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_9d7cf05c9992c480, []int{0}
}
func (m *ExtensionOptionDynamicFeeTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExtensionOptionDynamicFeeTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExtensionOptionDynamicFeeTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExtensionOptionDynamicFeeTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtensionOptionDynamicFeeTx.Merge(m, src)
}
func (m *ExtensionOptionDynamicFeeTx) XXX_Size() int {
	return m.Size()
}
func (m *ExtensionOptionDynamicFeeTx) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtensionOptionDynamicFeeTx.DiscardUnknown(m)
}

var xxx_messageInfo_ExtensionOptionDynamicFeeTx proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExtensionOptionDynamicFeeTx)(nil), "ethermint.types.v1.ExtensionOptionDynamicFeeTx")
}

func init() {
	proto.RegisterFile("ethermint/types/v1/dynamic_fee.proto", fileDescriptor_9d7cf05c9992c480)
}

var fileDescriptor_9d7cf05c9992c480 = []byte{
	// 253 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0xdb, 0x8b, 0x60, 0x4f, 0x52, 0x3c, 0x88, 0x42, 0x26, 0x22, 0xe2, 0x41, 0x13, 0x86,
	0x6f, 0x30, 0x74, 0xe0, 0xc9, 0x21, 0x9e, 0x44, 0x18, 0x69, 0xfa, 0xad, 0x0d, 0x9a, 0x7c, 0x5d,
	0xf2, 0x6d, 0xb4, 0xf8, 0x12, 0x3e, 0xd6, 0x8e, 0x3b, 0x8a, 0x87, 0x21, 0xed, 0x8b, 0x48, 0xd3,
	0xe1, 0x29, 0x3f, 0xc8, 0x8f, 0x1f, 0xff, 0x2f, 0xb9, 0x04, 0x2a, 0xc1, 0x19, 0x6d, 0x49, 0x50,
	0x53, 0x81, 0x17, 0xeb, 0xb1, 0xc8, 0x1b, 0x2b, 0x8d, 0x56, 0xf3, 0x05, 0x00, 0xaf, 0x1c, 0x12,
	0xa6, 0xe9, 0xbf, 0xc5, 0x83, 0xc5, 0xd7, 0xe3, 0xd3, 0xe3, 0x02, 0x0b, 0x0c, 0xdf, 0xa2, 0xa7,
	0xc1, 0xbc, 0xf8, 0x4c, 0xce, 0x1e, 0x6a, 0x02, 0xeb, 0x35, 0xda, 0xa7, 0x8a, 0x34, 0xda, 0xfb,
	0xa1, 0x36, 0x05, 0x78, 0xa9, 0xd3, 0xb7, 0x24, 0x35, 0xb2, 0x9e, 0x57, 0x4e, 0xa3, 0xd3, 0xd4,
	0xf4, 0xa0, 0xe0, 0x24, 0x3e, 0x8f, 0xaf, 0x0f, 0x27, 0x7c, 0xb3, 0x1b, 0x45, 0x3f, 0xbb, 0xd1,
	0x55, 0xa1, 0xa9, 0x5c, 0x65, 0x5c, 0xa1, 0x11, 0x0a, 0xbd, 0x41, 0xbf, 0x7f, 0x6e, 0x7d, 0xfe,
	0x3e, 0xac, 0xe4, 0x8f, 0x96, 0x9e, 0x8f, 0x8c, 0xac, 0x67, 0xfb, 0xd0, 0xac, 0xef, 0x4c, 0xa6,
	0x9b, 0x96, 0xc5, 0xdb, 0x96, 0xc5, 0xbf, 0x2d, 0x8b, 0xbf, 0x3a, 0x16, 0x6d, 0x3b, 0x16, 0x7d,
	0x77, 0x2c, 0x7a, 0xbd, 0x29, 0x34, 0x7d, 0xc8, 0x8c, 0x2f, 0x1d, 0xe4, 0x18, 0xca, 0x4b, 0x97,
	0xa3, 0x2a, 0xa5, 0xb6, 0x62, 0xb1, 0xea, 0xc7, 0x0e, 0x1c, 0xea, 0xd9, 0x41, 0xb8, 0xe5, 0xee,
	0x2f, 0x00, 0x00, 0xff, 0xff, 0x0c, 0xf0, 0xd3, 0x48, 0x1d, 0x01, 0x00, 0x00,
}

func (m *ExtensionOptionDynamicFeeTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExtensionOptionDynamicFeeTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExtensionOptionDynamicFeeTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MaxPriorityPrice.Size()
		i -= size
		if _, err := m.MaxPriorityPrice.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintDynamicFee(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintDynamicFee(dAtA []byte, offset int, v uint64) int {
	offset -= sovDynamicFee(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExtensionOptionDynamicFeeTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.MaxPriorityPrice.Size()
	n += 1 + l + sovDynamicFee(uint64(l))
	return n
}

func sovDynamicFee(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDynamicFee(x uint64) (n int) {
	return sovDynamicFee(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExtensionOptionDynamicFeeTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDynamicFee
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
			return fmt.Errorf("proto: ExtensionOptionDynamicFeeTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExtensionOptionDynamicFeeTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPriorityPrice", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDynamicFee
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
				return ErrInvalidLengthDynamicFee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDynamicFee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.MaxPriorityPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDynamicFee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDynamicFee
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
func skipDynamicFee(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDynamicFee
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
					return 0, ErrIntOverflowDynamicFee
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
					return 0, ErrIntOverflowDynamicFee
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
				return 0, ErrInvalidLengthDynamicFee
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDynamicFee
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDynamicFee
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDynamicFee        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDynamicFee          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDynamicFee = fmt.Errorf("proto: unexpected end of group")
)
