// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/tick.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Tick struct {
	PairId                string                                 `protobuf:"bytes,1,opt,name=pairId,proto3" json:"pairId,omitempty"`
	TickIndex             int64                                  `protobuf:"varint,2,opt,name=tickIndex,proto3" json:"tickIndex,omitempty"`
	TickData              *TickDataType                          `protobuf:"bytes,3,opt,name=tickData,proto3" json:"tickData,omitempty"`
	LimitOrderTranche0To1 *LimitTrancheIndexes                   `protobuf:"bytes,4,opt,name=LimitOrderTranche0to1,proto3" json:"LimitOrderTranche0to1,omitempty"`
	LimitOrderTranche1To0 *LimitTrancheIndexes                   `protobuf:"bytes,5,opt,name=LimitOrderTranche1to0,proto3" json:"LimitOrderTranche1to0,omitempty"`
	Price0To1             github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=price0To1,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"price0To1" yaml:"price0To1"`
}

func (m *Tick) Reset()         { *m = Tick{} }
func (m *Tick) String() string { return proto.CompactTextString(m) }
func (*Tick) ProtoMessage()    {}
func (*Tick) Descriptor() ([]byte, []int) {
	return fileDescriptor_120cef08c2828891, []int{0}
}
func (m *Tick) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Tick) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Tick.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Tick) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tick.Merge(m, src)
}
func (m *Tick) XXX_Size() int {
	return m.Size()
}
func (m *Tick) XXX_DiscardUnknown() {
	xxx_messageInfo_Tick.DiscardUnknown(m)
}

var xxx_messageInfo_Tick proto.InternalMessageInfo

func (m *Tick) GetPairId() string {
	if m != nil {
		return m.PairId
	}
	return ""
}

func (m *Tick) GetTickIndex() int64 {
	if m != nil {
		return m.TickIndex
	}
	return 0
}

func (m *Tick) GetTickData() *TickDataType {
	if m != nil {
		return m.TickData
	}
	return nil
}

func (m *Tick) GetLimitOrderTranche0To1() *LimitTrancheIndexes {
	if m != nil {
		return m.LimitOrderTranche0To1
	}
	return nil
}

func (m *Tick) GetLimitOrderTranche1To0() *LimitTrancheIndexes {
	if m != nil {
		return m.LimitOrderTranche1To0
	}
	return nil
}

func init() {
	proto.RegisterType((*Tick)(nil), "nicholasdotsol.duality.dex.Tick")
}

func init() { proto.RegisterFile("dex/tick.proto", fileDescriptor_120cef08c2828891) }

var fileDescriptor_120cef08c2828891 = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x3d, 0x6b, 0xe3, 0x40,
	0x10, 0x95, 0xce, 0x3e, 0x73, 0xd6, 0xc1, 0x71, 0x88, 0xbb, 0x20, 0x4c, 0x90, 0x8c, 0x8b, 0xa0,
	0xc6, 0x5a, 0x3b, 0xe9, 0x52, 0x1a, 0x41, 0x30, 0x84, 0x18, 0x14, 0x55, 0x69, 0xc4, 0x7a, 0x77,
	0xb1, 0x17, 0x4b, 0x1e, 0x21, 0xad, 0x41, 0xfa, 0x17, 0xf9, 0x47, 0x69, 0x5d, 0xba, 0x0c, 0x29,
	0x44, 0xb0, 0xbb, 0x94, 0xf9, 0x05, 0x61, 0x25, 0xc5, 0x4e, 0xe1, 0xa4, 0x4a, 0xa5, 0xf9, 0x7a,
	0xef, 0xcd, 0xd3, 0x8e, 0xf6, 0x87, 0xb2, 0x0c, 0x09, 0x4e, 0x16, 0x4e, 0x9c, 0x80, 0x00, 0xbd,
	0xb3, 0xe4, 0x64, 0x0e, 0x21, 0x4e, 0x29, 0x88, 0x14, 0x42, 0x87, 0xae, 0x70, 0xc8, 0x45, 0xee,
	0x50, 0x96, 0x75, 0x8c, 0xf7, 0xd9, 0x80, 0x62, 0x81, 0x03, 0x91, 0xc7, 0xac, 0x42, 0x75, 0x2c,
	0xd9, 0x09, 0x79, 0xc4, 0x45, 0x20, 0x12, 0xbc, 0x24, 0x73, 0x16, 0xf0, 0x25, 0x65, 0x19, 0x4b,
	0xeb, 0x81, 0x7f, 0x33, 0x98, 0x41, 0x19, 0x22, 0x19, 0x55, 0xd5, 0xde, 0x43, 0x43, 0x6b, 0xfa,
	0x9c, 0x2c, 0xf4, 0x13, 0xad, 0x15, 0x63, 0x9e, 0x8c, 0xa9, 0xa1, 0x76, 0x55, 0xbb, 0xed, 0xd5,
	0x99, 0x7e, 0xaa, 0xb5, 0xa5, 0xde, 0x58, 0x72, 0x19, 0x3f, 0xba, 0xaa, 0xdd, 0xf0, 0x0e, 0x05,
	0xdd, 0xd5, 0x7e, 0xc9, 0xc4, 0xc5, 0x02, 0x1b, 0x8d, 0xae, 0x6a, 0xff, 0x3e, 0xb7, 0x9d, 0xcf,
	0xd7, 0x77, 0xfc, 0x7a, 0xd6, 0xcf, 0x63, 0xe6, 0xed, 0x91, 0x3a, 0xd3, 0xfe, 0x5f, 0xcb, 0xcd,
	0x27, 0x09, 0x65, 0x89, 0x5f, 0x6d, 0x3f, 0x10, 0x30, 0x34, 0x9a, 0x25, 0x25, 0xfa, 0x8a, 0xb2,
	0x04, 0xd6, 0x98, 0x71, 0x65, 0xd8, 0x3b, 0xce, 0x76, 0x54, 0x66, 0x28, 0x60, 0x60, 0xfc, 0xfc,
	0x2e, 0x19, 0xc9, 0xa6, 0x47, 0x5a, 0x3b, 0x4e, 0x38, 0x61, 0x03, 0x1f, 0x86, 0x46, 0x4b, 0xfe,
	0xcc, 0xd1, 0x64, 0x5d, 0x58, 0xca, 0x53, 0x61, 0x9d, 0xcd, 0xb8, 0x98, 0xaf, 0xa6, 0x0e, 0x81,
	0x08, 0x11, 0x48, 0x23, 0x48, 0xeb, 0x4f, 0x3f, 0xa5, 0x0b, 0x24, 0x5f, 0x33, 0x75, 0x5c, 0x46,
	0x5e, 0x0a, 0xeb, 0x40, 0xf1, 0x5a, 0x58, 0x7f, 0x73, 0x1c, 0x85, 0x97, 0xbd, 0x7d, 0xa9, 0xe7,
	0x1d, 0xda, 0xa3, 0xab, 0xf5, 0xd6, 0x54, 0x37, 0x5b, 0x53, 0x7d, 0xde, 0x9a, 0xea, 0xfd, 0xce,
	0x54, 0x36, 0x3b, 0x53, 0x79, 0xdc, 0x99, 0xca, 0x5d, 0xff, 0x83, 0xda, 0x4d, 0x6d, 0xcd, 0x05,
	0x71, 0x0b, 0x21, 0xaa, 0xad, 0xa1, 0x0c, 0x95, 0x07, 0x25, 0x85, 0xa7, 0xad, 0xf2, 0x22, 0x2e,
	0xde, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xb3, 0x5a, 0xd5, 0x90, 0x02, 0x00, 0x00,
}

func (m *Tick) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Tick) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Tick) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Price0To1.Size()
		i -= size
		if _, err := m.Price0To1.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintTick(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if m.LimitOrderTranche1To0 != nil {
		{
			size, err := m.LimitOrderTranche1To0.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTick(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.LimitOrderTranche0To1 != nil {
		{
			size, err := m.LimitOrderTranche0To1.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTick(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.TickData != nil {
		{
			size, err := m.TickData.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTick(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.TickIndex != 0 {
		i = encodeVarintTick(dAtA, i, uint64(m.TickIndex))
		i--
		dAtA[i] = 0x10
	}
	if len(m.PairId) > 0 {
		i -= len(m.PairId)
		copy(dAtA[i:], m.PairId)
		i = encodeVarintTick(dAtA, i, uint64(len(m.PairId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTick(dAtA []byte, offset int, v uint64) int {
	offset -= sovTick(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Tick) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PairId)
	if l > 0 {
		n += 1 + l + sovTick(uint64(l))
	}
	if m.TickIndex != 0 {
		n += 1 + sovTick(uint64(m.TickIndex))
	}
	if m.TickData != nil {
		l = m.TickData.Size()
		n += 1 + l + sovTick(uint64(l))
	}
	if m.LimitOrderTranche0To1 != nil {
		l = m.LimitOrderTranche0To1.Size()
		n += 1 + l + sovTick(uint64(l))
	}
	if m.LimitOrderTranche1To0 != nil {
		l = m.LimitOrderTranche1To0.Size()
		n += 1 + l + sovTick(uint64(l))
	}
	l = m.Price0To1.Size()
	n += 1 + l + sovTick(uint64(l))
	return n
}

func sovTick(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTick(x uint64) (n int) {
	return sovTick(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Tick) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTick
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
			return fmt.Errorf("proto: Tick: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Tick: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTick
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
				return ErrInvalidLengthTick
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTick
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PairId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickIndex", wireType)
			}
			m.TickIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTick
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TickIndex |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTick
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
				return ErrInvalidLengthTick
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTick
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TickData == nil {
				m.TickData = &TickDataType{}
			}
			if err := m.TickData.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitOrderTranche0To1", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTick
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
				return ErrInvalidLengthTick
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTick
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LimitOrderTranche0To1 == nil {
				m.LimitOrderTranche0To1 = &LimitTrancheIndexes{}
			}
			if err := m.LimitOrderTranche0To1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitOrderTranche1To0", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTick
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
				return ErrInvalidLengthTick
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTick
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LimitOrderTranche1To0 == nil {
				m.LimitOrderTranche1To0 = &LimitTrancheIndexes{}
			}
			if err := m.LimitOrderTranche1To0.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price0To1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTick
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
				return ErrInvalidLengthTick
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTick
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Price0To1.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTick(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTick
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
func skipTick(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTick
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
					return 0, ErrIntOverflowTick
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
					return 0, ErrIntOverflowTick
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
				return 0, ErrInvalidLengthTick
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTick
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTick
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTick        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTick          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTick = fmt.Errorf("proto: unexpected end of group")
)
