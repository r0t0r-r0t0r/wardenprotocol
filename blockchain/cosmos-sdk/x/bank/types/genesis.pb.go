// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/bank/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

// GenesisState defines the bank module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	// balances is an array containing the balances of all the accounts.
	Balances []Balance `protobuf:"bytes,2,rep,name=balances,proto3" json:"balances"`
	// supply represents the total supply. If it is left empty, then supply will be calculated based on the provided
	// balances. Otherwise, it will be used to validate that the sum of the balances equals this amount.
	Supply github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,3,rep,name=supply,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"supply"`
	// denom_metadata defines the metadata of the different coins.
	DenomMetadata []Metadata `protobuf:"bytes,4,rep,name=denom_metadata,json=denomMetadata,proto3" json:"denom_metadata"`
	// send_enabled defines the denoms where send is enabled or disabled.
	//
	// Since: cosmos-sdk 0.47
	SendEnabled []SendEnabled `protobuf:"bytes,5,rep,name=send_enabled,json=sendEnabled,proto3" json:"send_enabled"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f007de11b420c6e, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetBalances() []Balance {
	if m != nil {
		return m.Balances
	}
	return nil
}

func (m *GenesisState) GetSupply() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Supply
	}
	return nil
}

func (m *GenesisState) GetDenomMetadata() []Metadata {
	if m != nil {
		return m.DenomMetadata
	}
	return nil
}

func (m *GenesisState) GetSendEnabled() []SendEnabled {
	if m != nil {
		return m.SendEnabled
	}
	return nil
}

// Balance defines an account address and balance pair used in the bank module's
// genesis state.
type Balance struct {
	// address is the address of the balance holder.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// coins defines the different coins this balance holds.
	Coins github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=coins,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"coins"`
}

func (m *Balance) Reset()         { *m = Balance{} }
func (m *Balance) String() string { return proto.CompactTextString(m) }
func (*Balance) ProtoMessage()    {}
func (*Balance) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f007de11b420c6e, []int{1}
}
func (m *Balance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Balance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Balance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Balance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Balance.Merge(m, src)
}
func (m *Balance) XXX_Size() int {
	return m.Size()
}
func (m *Balance) XXX_DiscardUnknown() {
	xxx_messageInfo_Balance.DiscardUnknown(m)
}

var xxx_messageInfo_Balance proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GenesisState)(nil), "cosmos.bank.v1beta1.GenesisState")
	proto.RegisterType((*Balance)(nil), "cosmos.bank.v1beta1.Balance")
}

func init() { proto.RegisterFile("cosmos/bank/v1beta1/genesis.proto", fileDescriptor_8f007de11b420c6e) }

var fileDescriptor_8f007de11b420c6e = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xbd, 0xae, 0xd3, 0x30,
	0x1c, 0xc5, 0x13, 0xca, 0xed, 0xbd, 0xd7, 0xbd, 0x20, 0x61, 0x3a, 0xa4, 0x05, 0x92, 0xd2, 0xa9,
	0x20, 0x35, 0x51, 0x8b, 0x58, 0x18, 0x90, 0x48, 0x85, 0x98, 0xf8, 0x50, 0xbb, 0xb1, 0x54, 0x4e,
	0x6c, 0xd2, 0xa8, 0x8d, 0x1d, 0xc5, 0x2e, 0xa2, 0x6f, 0xc0, 0xc8, 0x13, 0xa0, 0x8e, 0x88, 0x05,
	0x06, 0x1e, 0xa2, 0x63, 0xc5, 0xc4, 0x04, 0xa8, 0x1d, 0xe0, 0x31, 0x50, 0x6c, 0x37, 0x8d, 0x44,
	0xc4, 0xc4, 0x92, 0x44, 0x3e, 0xe7, 0xfc, 0xfe, 0x27, 0xb6, 0xc1, 0xed, 0x90, 0xf1, 0x84, 0x71,
	0x2f, 0x40, 0x74, 0xee, 0xbd, 0x1e, 0x04, 0x44, 0xa0, 0x81, 0x17, 0x11, 0x4a, 0x78, 0xcc, 0xdd,
	0x34, 0x63, 0x82, 0xc1, 0xeb, 0xca, 0xe2, 0xe6, 0x16, 0x57, 0x5b, 0xda, 0xcd, 0x88, 0x45, 0x4c,
	0xea, 0x5e, 0xfe, 0xa5, 0xac, 0x6d, 0xbb, 0xa0, 0x71, 0x52, 0xd0, 0x42, 0x16, 0xd3, 0xbf, 0xf4,
	0xd2, 0x34, 0xc9, 0x55, 0x7a, 0x4b, 0xe9, 0x53, 0x05, 0xd6, 0x73, 0x95, 0x74, 0x0d, 0x25, 0x31,
	0x65, 0x9e, 0x7c, 0xaa, 0xa5, 0xee, 0xfb, 0x1a, 0xb8, 0x78, 0xa2, 0xaa, 0x4e, 0x04, 0x12, 0x04,
	0x3e, 0x04, 0xf5, 0x14, 0x65, 0x28, 0xe1, 0x96, 0xd9, 0x31, 0x7b, 0x8d, 0xe1, 0x0d, 0xb7, 0xa2,
	0xba, 0xfb, 0x42, 0x5a, 0xfc, 0xf3, 0xcd, 0x77, 0xc7, 0xf8, 0xf0, 0xeb, 0xf3, 0x5d, 0x73, 0xac,
	0x53, 0x70, 0x04, 0xce, 0x02, 0xb4, 0x40, 0x34, 0x24, 0xdc, 0xba, 0xd4, 0xa9, 0xf5, 0x1a, 0xc3,
	0x9b, 0x95, 0x04, 0x5f, 0x99, 0xca, 0x88, 0x22, 0x08, 0x67, 0xa0, 0xce, 0x97, 0x69, 0xba, 0x58,
	0x59, 0x35, 0x89, 0x68, 0x1d, 0x11, 0x9c, 0x14, 0x88, 0x11, 0x8b, 0xa9, 0x7f, 0x3f, 0xcf, 0x7f,
	0xfc, 0xe1, 0xf4, 0xa2, 0x58, 0xcc, 0x96, 0x81, 0x1b, 0xb2, 0x44, 0xff, 0xb4, 0x7e, 0xf5, 0x39,
	0x9e, 0x7b, 0x62, 0x95, 0x12, 0x2e, 0x03, 0x5c, 0xd7, 0x55, 0x7c, 0xf8, 0x1c, 0x5c, 0xc5, 0x84,
	0xb2, 0x64, 0x9a, 0x10, 0x81, 0x30, 0x12, 0xc8, 0xba, 0x2c, 0x27, 0xde, 0xaa, 0x2c, 0xfd, 0x54,
	0x9b, 0xca, 0xad, 0xaf, 0xc8, 0xfc, 0x41, 0x81, 0xcf, 0xc0, 0x05, 0x27, 0x14, 0x4f, 0x09, 0x45,
	0xc1, 0x82, 0x60, 0xeb, 0x44, 0xe2, 0x3a, 0x95, 0xb8, 0x09, 0xa1, 0xf8, 0xb1, 0xf2, 0x95, 0x89,
	0x0d, 0x7e, 0x5c, 0xef, 0x7e, 0x32, 0xc1, 0xa9, 0xde, 0x2b, 0x38, 0x04, 0xa7, 0x08, 0xe3, 0x8c,
	0x70, 0x75, 0x38, 0xe7, 0xbe, 0xf5, 0xf5, 0x4b, 0xbf, 0xa9, 0xc9, 0x8f, 0x94, 0x32, 0x11, 0x59,
	0x4c, 0xa3, 0xf1, 0xc1, 0x08, 0x5f, 0x81, 0x93, 0xfc, 0xf2, 0x1c, 0x0e, 0xe3, 0xff, 0xef, 0xa4,
	0xc2, 0x3f, 0x38, 0x7b, 0xbb, 0x76, 0x8c, 0xdf, 0x6b, 0xc7, 0xf0, 0x47, 0x9b, 0x9d, 0x6d, 0x6e,
	0x77, 0xb6, 0xf9, 0x73, 0x67, 0x9b, 0xef, 0xf6, 0xb6, 0xb1, 0xdd, 0xdb, 0xc6, 0xb7, 0xbd, 0x6d,
	0xbc, 0xbc, 0xf3, 0x4f, 0xf2, 0x1b, 0x75, 0xa5, 0xe5, 0x80, 0xa0, 0x2e, 0xaf, 0xe7, 0xbd, 0x3f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x88, 0xa3, 0x71, 0x5c, 0x03, 0x00, 0x00,
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
	if len(m.SendEnabled) > 0 {
		for iNdEx := len(m.SendEnabled) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SendEnabled[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.DenomMetadata) > 0 {
		for iNdEx := len(m.DenomMetadata) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DenomMetadata[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Supply) > 0 {
		for iNdEx := len(m.Supply) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Supply[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Balances) > 0 {
		for iNdEx := len(m.Balances) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Balances[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Balance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Balance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Balance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Coins) > 0 {
		for iNdEx := len(m.Coins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Coins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Balances) > 0 {
		for _, e := range m.Balances {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Supply) > 0 {
		for _, e := range m.Supply {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DenomMetadata) > 0 {
		for _, e := range m.DenomMetadata {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SendEnabled) > 0 {
		for _, e := range m.SendEnabled {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *Balance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.Coins) > 0 {
		for _, e := range m.Coins {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balances", wireType)
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
			m.Balances = append(m.Balances, Balance{})
			if err := m.Balances[len(m.Balances)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supply", wireType)
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
			m.Supply = append(m.Supply, types.Coin{})
			if err := m.Supply[len(m.Supply)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomMetadata", wireType)
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
			m.DenomMetadata = append(m.DenomMetadata, Metadata{})
			if err := m.DenomMetadata[len(m.DenomMetadata)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SendEnabled", wireType)
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
			m.SendEnabled = append(m.SendEnabled, SendEnabled{})
			if err := m.SendEnabled[len(m.SendEnabled)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *Balance) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Balance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Balance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coins", wireType)
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
			m.Coins = append(m.Coins, types.Coin{})
			if err := m.Coins[len(m.Coins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
