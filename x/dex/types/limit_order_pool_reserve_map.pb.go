// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/limit_order_pool_reserve_map.proto

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

type LimitOrderTrancheReserveMap struct {
	PairId    string                                 `protobuf:"bytes,1,opt,name=pairId,proto3" json:"pairId,omitempty"`
	Token     string                                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	TickIndex int64                                  `protobuf:"varint,3,opt,name=tickIndex,proto3" json:"tickIndex,omitempty"`
	Count     uint64                                 `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	Reserves  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=reserves,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reserves" yaml:"reserves"`
}

func (m *LimitOrderTrancheReserveMap) Reset()         { *m = LimitOrderTrancheReserveMap{} }
func (m *LimitOrderTrancheReserveMap) String() string { return proto.CompactTextString(m) }
func (*LimitOrderTrancheReserveMap) ProtoMessage()    {}
func (*LimitOrderTrancheReserveMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c0f8247f33ec8ae, []int{0}
}
func (m *LimitOrderTrancheReserveMap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LimitOrderTrancheReserveMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LimitOrderTrancheReserveMap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LimitOrderTrancheReserveMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimitOrderTrancheReserveMap.Merge(m, src)
}
func (m *LimitOrderTrancheReserveMap) XXX_Size() int {
	return m.Size()
}
func (m *LimitOrderTrancheReserveMap) XXX_DiscardUnknown() {
	xxx_messageInfo_LimitOrderTrancheReserveMap.DiscardUnknown(m)
}

var xxx_messageInfo_LimitOrderTrancheReserveMap proto.InternalMessageInfo

func (m *LimitOrderTrancheReserveMap) GetPairId() string {
	if m != nil {
		return m.PairId
	}
	return ""
}

func (m *LimitOrderTrancheReserveMap) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LimitOrderTrancheReserveMap) GetTickIndex() int64 {
	if m != nil {
		return m.TickIndex
	}
	return 0
}

func (m *LimitOrderTrancheReserveMap) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*LimitOrderTrancheReserveMap)(nil), "nicholasdotsol.duality.dex.LimitOrderTrancheReserveMap")
}

func init() {
	proto.RegisterFile("dex/limit_order_pool_reserve_map.proto", fileDescriptor_9c0f8247f33ec8ae)
}

var fileDescriptor_9c0f8247f33ec8ae = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x1b, 0xf7, 0x07, 0xd7, 0x8b, 0x50, 0x86, 0x94, 0x21, 0xed, 0xd8, 0x61, 0xec, 0xb2,
	0xe6, 0xe0, 0xcd, 0xe3, 0x18, 0xc8, 0xc0, 0xa9, 0xd4, 0x9b, 0x97, 0xd2, 0x35, 0x61, 0x0b, 0x4b,
	0xfb, 0x86, 0x24, 0x93, 0xee, 0x5b, 0xf8, 0xb1, 0x76, 0xdc, 0x51, 0x3c, 0x14, 0xd9, 0x0e, 0x82,
	0x47, 0x3f, 0x81, 0xa4, 0xad, 0xd3, 0x53, 0xf2, 0x3c, 0xc9, 0xfb, 0x7b, 0xc8, 0x13, 0x7b, 0x48,
	0x68, 0x8e, 0x39, 0x4b, 0x99, 0x8e, 0x40, 0x12, 0x2a, 0x23, 0x01, 0xc0, 0x23, 0x49, 0x15, 0x95,
	0x2f, 0x34, 0x4a, 0x63, 0x11, 0x08, 0x09, 0x1a, 0x9c, 0x5e, 0xc6, 0x92, 0x15, 0xf0, 0x58, 0x11,
	0xd0, 0x0a, 0x78, 0x40, 0x36, 0x31, 0x67, 0x7a, 0x1b, 0x10, 0x9a, 0xf7, 0xba, 0x4b, 0x58, 0x42,
	0x79, 0x0d, 0x9b, 0x5d, 0x35, 0x31, 0xf8, 0x44, 0xb6, 0x7b, 0x67, 0xc0, 0x0f, 0x86, 0xfb, 0x08,
	0xc0, 0xc3, 0x8a, 0x3a, 0x8f, 0x85, 0x73, 0x69, 0xb7, 0x45, 0xcc, 0xe4, 0x8c, 0xb8, 0xa8, 0x8f,
	0x46, 0x9d, 0xb0, 0x56, 0x4e, 0xd7, 0x6e, 0x69, 0x58, 0xd3, 0xcc, 0x3d, 0x2b, 0xed, 0x4a, 0x38,
	0x57, 0x76, 0x47, 0xb3, 0x64, 0x3d, 0xcb, 0x08, 0xcd, 0xdd, 0x46, 0x1f, 0x8d, 0x1a, 0xe1, 0x9f,
	0x61, 0x66, 0x12, 0xd8, 0x64, 0xda, 0x6d, 0xf6, 0xd1, 0xa8, 0x19, 0x56, 0xc2, 0x61, 0xf6, 0x79,
	0xfd, 0x0a, 0xe5, 0xb6, 0x0c, 0x6c, 0x32, 0xdf, 0x15, 0xbe, 0xf5, 0x5e, 0xf8, 0xc3, 0x25, 0xd3,
	0xab, 0xcd, 0x22, 0x48, 0x20, 0xc5, 0x09, 0xa8, 0x14, 0x54, 0xbd, 0x8c, 0x15, 0x59, 0x63, 0xbd,
	0x15, 0x54, 0x05, 0x53, 0x9a, 0x7c, 0x15, 0xfe, 0x89, 0xf0, 0x5d, 0xf8, 0x17, 0xdb, 0x38, 0xe5,
	0x37, 0x83, 0x5f, 0x67, 0x10, 0x9e, 0x0e, 0x27, 0xb7, 0xbb, 0x83, 0x87, 0xf6, 0x07, 0x0f, 0x7d,
	0x1c, 0x3c, 0xf4, 0x7a, 0xf4, 0xac, 0xfd, 0xd1, 0xb3, 0xde, 0x8e, 0x9e, 0xf5, 0x3c, 0xfe, 0x17,
	0x75, 0x5f, 0x17, 0x38, 0x05, 0xfd, 0x04, 0x1c, 0xd7, 0x05, 0xe2, 0x1c, 0x9b, 0x1f, 0x28, 0x53,
	0x17, 0xed, 0xb2, 0xb9, 0xeb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x97, 0xad, 0x46, 0x4b, 0x95,
	0x01, 0x00, 0x00,
}

func (m *LimitOrderTrancheReserveMap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LimitOrderTrancheReserveMap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LimitOrderTrancheReserveMap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Reserves.Size()
		i -= size
		if _, err := m.Reserves.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTrancheReserveMap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.Count != 0 {
		i = encodeVarintLimitOrderTrancheReserveMap(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x20
	}
	if m.TickIndex != 0 {
		i = encodeVarintLimitOrderTrancheReserveMap(dAtA, i, uint64(m.TickIndex))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintLimitOrderTrancheReserveMap(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PairId) > 0 {
		i -= len(m.PairId)
		copy(dAtA[i:], m.PairId)
		i = encodeVarintLimitOrderTrancheReserveMap(dAtA, i, uint64(len(m.PairId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLimitOrderTrancheReserveMap(dAtA []byte, offset int, v uint64) int {
	offset -= sovLimitOrderTrancheReserveMap(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LimitOrderTrancheReserveMap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PairId)
	if l > 0 {
		n += 1 + l + sovLimitOrderTrancheReserveMap(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovLimitOrderTrancheReserveMap(uint64(l))
	}
	if m.TickIndex != 0 {
		n += 1 + sovLimitOrderTrancheReserveMap(uint64(m.TickIndex))
	}
	if m.Count != 0 {
		n += 1 + sovLimitOrderTrancheReserveMap(uint64(m.Count))
	}
	l = m.Reserves.Size()
	n += 1 + l + sovLimitOrderTrancheReserveMap(uint64(l))
	return n
}

func sovLimitOrderTrancheReserveMap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLimitOrderTrancheReserveMap(x uint64) (n int) {
	return sovLimitOrderTrancheReserveMap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LimitOrderTrancheReserveMap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLimitOrderTrancheReserveMap
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
			return fmt.Errorf("proto: LimitOrderTrancheReserveMap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LimitOrderTrancheReserveMap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheReserveMap
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
				return ErrInvalidLengthLimitOrderTrancheReserveMap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTrancheReserveMap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PairId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheReserveMap
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
				return ErrInvalidLengthLimitOrderTrancheReserveMap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTrancheReserveMap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickIndex", wireType)
			}
			m.TickIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheReserveMap
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
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheReserveMap
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reserves", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheReserveMap
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
				return ErrInvalidLengthLimitOrderTrancheReserveMap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTrancheReserveMap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Reserves.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLimitOrderTrancheReserveMap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLimitOrderTrancheReserveMap
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
func skipLimitOrderTrancheReserveMap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLimitOrderTrancheReserveMap
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
					return 0, ErrIntOverflowLimitOrderTrancheReserveMap
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
					return 0, ErrIntOverflowLimitOrderTrancheReserveMap
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
				return 0, ErrInvalidLengthLimitOrderTrancheReserveMap
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLimitOrderTrancheReserveMap
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLimitOrderTrancheReserveMap
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLimitOrderTrancheReserveMap        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLimitOrderTrancheReserveMap          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLimitOrderTrancheReserveMap = fmt.Errorf("proto: unexpected end of group")
)
