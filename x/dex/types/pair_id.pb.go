// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: duality/dex/pair_id.proto

package types

import (
	fmt "fmt"
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

type PairId struct {
	Token0 string `protobuf:"bytes,1,opt,name=token0,proto3" json:"token0,omitempty"`
	Token1 string `protobuf:"bytes,2,opt,name=token1,proto3" json:"token1,omitempty"`
}

func (m *PairId) Reset()         { *m = PairId{} }
func (m *PairId) String() string { return proto.CompactTextString(m) }
func (*PairId) ProtoMessage()    {}
func (*PairId) Descriptor() ([]byte, []int) {
	return fileDescriptor_1919813da3dc14c8, []int{0}
}
func (m *PairId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PairId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PairId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PairId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PairId.Merge(m, src)
}
func (m *PairId) XXX_Size() int {
	return m.Size()
}
func (m *PairId) XXX_DiscardUnknown() {
	xxx_messageInfo_PairId.DiscardUnknown(m)
}

var xxx_messageInfo_PairId proto.InternalMessageInfo

func (m *PairId) GetToken0() string {
	if m != nil {
		return m.Token0
	}
	return ""
}

func (m *PairId) GetToken1() string {
	if m != nil {
		return m.Token1
	}
	return ""
}

func init() {
	proto.RegisterType((*PairId)(nil), "nicholasdotsol.duality.dex.PairId")
}

func init() { proto.RegisterFile("duality/dex/pair_id.proto", fileDescriptor_1919813da3dc14c8) }

var fileDescriptor_1919813da3dc14c8 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4c, 0x29, 0x4d, 0xcc,
	0xc9, 0x2c, 0xa9, 0xd4, 0x4f, 0x49, 0xad, 0xd0, 0x2f, 0x48, 0xcc, 0x2c, 0x8a, 0xcf, 0x4c, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xca, 0xcb, 0x4c, 0xce, 0xc8, 0xcf, 0x49, 0x2c, 0x4e,
	0xc9, 0x2f, 0x29, 0xce, 0xcf, 0xd1, 0x83, 0xaa, 0xd4, 0x4b, 0x49, 0xad, 0x50, 0xb2, 0xe0, 0x62,
	0x0b, 0x48, 0xcc, 0x2c, 0xf2, 0x4c, 0x11, 0x12, 0xe3, 0x62, 0x2b, 0xc9, 0xcf, 0x4e, 0xcd, 0x33,
	0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0xe0, 0xe2, 0x86, 0x12, 0x4c, 0x48, 0xe2,
	0x86, 0x4e, 0xee, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3,
	0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x9b, 0x9e,
	0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0xef, 0x07, 0xb5, 0xda, 0x25, 0xbf, 0x24,
	0x38, 0x3f, 0x47, 0x1f, 0xe6, 0xc8, 0x0a, 0xb0, 0x33, 0x4b, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8,
	0xc0, 0xae, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x51, 0xb2, 0x9a, 0x72, 0xc2, 0x00, 0x00,
	0x00,
}

func (m *PairId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PairId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PairId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Token1) > 0 {
		i -= len(m.Token1)
		copy(dAtA[i:], m.Token1)
		i = encodeVarintPairId(dAtA, i, uint64(len(m.Token1)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Token0) > 0 {
		i -= len(m.Token0)
		copy(dAtA[i:], m.Token0)
		i = encodeVarintPairId(dAtA, i, uint64(len(m.Token0)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPairId(dAtA []byte, offset int, v uint64) int {
	offset -= sovPairId(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PairId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Token0)
	if l > 0 {
		n += 1 + l + sovPairId(uint64(l))
	}
	l = len(m.Token1)
	if l > 0 {
		n += 1 + l + sovPairId(uint64(l))
	}
	return n
}

func sovPairId(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPairId(x uint64) (n int) {
	return sovPairId(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PairId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPairId
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
			return fmt.Errorf("proto: PairId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PairId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token0", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPairId
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
				return ErrInvalidLengthPairId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPairId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token0 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPairId
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
				return ErrInvalidLengthPairId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPairId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPairId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPairId
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
func skipPairId(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPairId
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
					return 0, ErrIntOverflowPairId
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
					return 0, ErrIntOverflowPairId
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
				return 0, ErrInvalidLengthPairId
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPairId
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPairId
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPairId        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPairId          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPairId = fmt.Errorf("proto: unexpected end of group")
)