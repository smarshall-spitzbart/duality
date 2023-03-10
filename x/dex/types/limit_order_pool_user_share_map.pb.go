// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/limit_order_pool_user_share_map.proto

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

type LimitOrderTrancheUserShareMap struct {
	PairId      string                                 `protobuf:"bytes,1,opt,name=pairId,proto3" json:"pairId,omitempty"`
	Token       string                                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	TickIndex   int64                                  `protobuf:"varint,3,opt,name=tickIndex,proto3" json:"tickIndex,omitempty"`
	Count       uint64                                 `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	Address     string                                 `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	SharesOwned github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=sharesOwned,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"sharesOwned" yaml:"sharesOwned"`
}

func (m *LimitOrderTrancheUserShareMap) Reset()         { *m = LimitOrderTrancheUserShareMap{} }
func (m *LimitOrderTrancheUserShareMap) String() string { return proto.CompactTextString(m) }
func (*LimitOrderTrancheUserShareMap) ProtoMessage()    {}
func (*LimitOrderTrancheUserShareMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_948654eb44426d87, []int{0}
}
func (m *LimitOrderTrancheUserShareMap) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LimitOrderTrancheUserShareMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LimitOrderTrancheUserShareMap.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LimitOrderTrancheUserShareMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimitOrderTrancheUserShareMap.Merge(m, src)
}
func (m *LimitOrderTrancheUserShareMap) XXX_Size() int {
	return m.Size()
}
func (m *LimitOrderTrancheUserShareMap) XXX_DiscardUnknown() {
	xxx_messageInfo_LimitOrderTrancheUserShareMap.DiscardUnknown(m)
}

var xxx_messageInfo_LimitOrderTrancheUserShareMap proto.InternalMessageInfo

func (m *LimitOrderTrancheUserShareMap) GetPairId() string {
	if m != nil {
		return m.PairId
	}
	return ""
}

func (m *LimitOrderTrancheUserShareMap) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LimitOrderTrancheUserShareMap) GetTickIndex() int64 {
	if m != nil {
		return m.TickIndex
	}
	return 0
}

func (m *LimitOrderTrancheUserShareMap) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *LimitOrderTrancheUserShareMap) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*LimitOrderTrancheUserShareMap)(nil), "nicholasdotsol.duality.dex.LimitOrderTrancheUserShareMap")
}

func init() {
	proto.RegisterFile("dex/limit_order_pool_user_share_map.proto", fileDescriptor_948654eb44426d87)
}

var fileDescriptor_948654eb44426d87 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x91, 0xbf, 0x6a, 0xc3, 0x30,
	0x10, 0xc6, 0xed, 0xfc, 0x2b, 0x71, 0x37, 0x13, 0x8a, 0x09, 0xc5, 0x0e, 0x19, 0x4a, 0x3a, 0xc4,
	0x1a, 0xba, 0x75, 0x0c, 0x81, 0x12, 0x68, 0x9b, 0xe2, 0xd0, 0xa5, 0x8b, 0x51, 0x2c, 0x91, 0x88,
	0xc8, 0x3e, 0x23, 0xc9, 0xd4, 0x79, 0x8a, 0xf6, 0xb1, 0x32, 0x66, 0x2c, 0x1d, 0x4c, 0x49, 0xb6,
	0x8e, 0x7d, 0x82, 0x22, 0xc7, 0xa5, 0x9e, 0xa4, 0xef, 0xd3, 0xdd, 0xef, 0xd0, 0x77, 0xd6, 0x35,
	0xa1, 0x39, 0xe2, 0x2c, 0x66, 0x2a, 0x04, 0x41, 0xa8, 0x08, 0x53, 0x00, 0x1e, 0x66, 0x92, 0x8a,
	0x50, 0xae, 0xb1, 0xa0, 0x61, 0x8c, 0x53, 0x3f, 0x15, 0xa0, 0xc0, 0xee, 0x27, 0x2c, 0x5a, 0x03,
	0xc7, 0x92, 0x80, 0x92, 0xc0, 0x7d, 0x92, 0x61, 0xce, 0xd4, 0xd6, 0x27, 0x34, 0xef, 0xf7, 0x56,
	0xb0, 0x82, 0xb2, 0x0c, 0xe9, 0xdb, 0xa9, 0x63, 0xf8, 0xd6, 0xb0, 0xfa, 0xf7, 0x9a, 0x3d, 0xd7,
	0xe8, 0x27, 0x00, 0xfe, 0x2c, 0xa9, 0x58, 0x68, 0xee, 0x03, 0x4e, 0xed, 0x0b, 0xab, 0x93, 0x62,
	0x26, 0x66, 0xc4, 0x31, 0x07, 0xe6, 0xa8, 0x1b, 0x54, 0xca, 0xee, 0x59, 0x6d, 0x05, 0x1b, 0x9a,
	0x38, 0x8d, 0xd2, 0x3e, 0x09, 0xfb, 0xd2, 0xea, 0x2a, 0x16, 0x6d, 0x66, 0x09, 0xa1, 0xb9, 0xd3,
	0x1c, 0x98, 0xa3, 0x66, 0xf0, 0x6f, 0xe8, 0x9e, 0x08, 0xb2, 0x44, 0x39, 0xad, 0x81, 0x39, 0x6a,
	0x05, 0x27, 0x61, 0x3b, 0xd6, 0x19, 0x26, 0x44, 0x50, 0x29, 0x9d, 0x76, 0xc9, 0xfa, 0x93, 0x76,
	0x66, 0x9d, 0x97, 0xff, 0x93, 0xf3, 0xd7, 0x84, 0x12, 0xa7, 0xa3, 0x5f, 0x27, 0x8b, 0x5d, 0xe1,
	0x19, 0x9f, 0x85, 0x77, 0xb5, 0x62, 0x6a, 0x9d, 0x2d, 0xfd, 0x08, 0x62, 0x14, 0x81, 0x8c, 0x41,
	0x56, 0xc7, 0x58, 0x92, 0x0d, 0x52, 0xdb, 0x94, 0x4a, 0x7f, 0x4a, 0xa3, 0xef, 0xc2, 0xab, 0x43,
	0x7e, 0x0a, 0xcf, 0xde, 0xe2, 0x98, 0xdf, 0x0e, 0x6b, 0xe6, 0x30, 0xa8, 0x97, 0x4c, 0xee, 0x76,
	0x07, 0xd7, 0xdc, 0x1f, 0x5c, 0xf3, 0xeb, 0xe0, 0x9a, 0xef, 0x47, 0xd7, 0xd8, 0x1f, 0x5d, 0xe3,
	0xe3, 0xe8, 0x1a, 0x2f, 0xe3, 0xda, 0xcc, 0xc7, 0x2a, 0xe8, 0x29, 0xa8, 0x05, 0x70, 0x54, 0x05,
	0x8d, 0x72, 0xa4, 0x97, 0x55, 0x8e, 0x5f, 0x76, 0xca, 0x84, 0x6f, 0x7e, 0x03, 0x00, 0x00, 0xff,
	0xff, 0x7b, 0x1c, 0x0e, 0x51, 0xc0, 0x01, 0x00, 0x00,
}

func (m *LimitOrderTrancheUserShareMap) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LimitOrderTrancheUserShareMap) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LimitOrderTrancheUserShareMap) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SharesOwned.Size()
		i -= size
		if _, err := m.SharesOwned.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTrancheUserShareMap(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintLimitOrderTrancheUserShareMap(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Count != 0 {
		i = encodeVarintLimitOrderTrancheUserShareMap(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x20
	}
	if m.TickIndex != 0 {
		i = encodeVarintLimitOrderTrancheUserShareMap(dAtA, i, uint64(m.TickIndex))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintLimitOrderTrancheUserShareMap(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PairId) > 0 {
		i -= len(m.PairId)
		copy(dAtA[i:], m.PairId)
		i = encodeVarintLimitOrderTrancheUserShareMap(dAtA, i, uint64(len(m.PairId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLimitOrderTrancheUserShareMap(dAtA []byte, offset int, v uint64) int {
	offset -= sovLimitOrderTrancheUserShareMap(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LimitOrderTrancheUserShareMap) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PairId)
	if l > 0 {
		n += 1 + l + sovLimitOrderTrancheUserShareMap(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovLimitOrderTrancheUserShareMap(uint64(l))
	}
	if m.TickIndex != 0 {
		n += 1 + sovLimitOrderTrancheUserShareMap(uint64(m.TickIndex))
	}
	if m.Count != 0 {
		n += 1 + sovLimitOrderTrancheUserShareMap(uint64(m.Count))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovLimitOrderTrancheUserShareMap(uint64(l))
	}
	l = m.SharesOwned.Size()
	n += 1 + l + sovLimitOrderTrancheUserShareMap(uint64(l))
	return n
}

func sovLimitOrderTrancheUserShareMap(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLimitOrderTrancheUserShareMap(x uint64) (n int) {
	return sovLimitOrderTrancheUserShareMap(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LimitOrderTrancheUserShareMap) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLimitOrderTrancheUserShareMap
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
			return fmt.Errorf("proto: LimitOrderTrancheUserShareMap: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LimitOrderTrancheUserShareMap: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheUserShareMap
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
				return ErrInvalidLengthLimitOrderTrancheUserShareMap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTrancheUserShareMap
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
					return ErrIntOverflowLimitOrderTrancheUserShareMap
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
				return ErrInvalidLengthLimitOrderTrancheUserShareMap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTrancheUserShareMap
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
					return ErrIntOverflowLimitOrderTrancheUserShareMap
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
					return ErrIntOverflowLimitOrderTrancheUserShareMap
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
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheUserShareMap
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
				return ErrInvalidLengthLimitOrderTrancheUserShareMap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTrancheUserShareMap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SharesOwned", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTrancheUserShareMap
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
				return ErrInvalidLengthLimitOrderTrancheUserShareMap
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTrancheUserShareMap
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SharesOwned.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLimitOrderTrancheUserShareMap(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLimitOrderTrancheUserShareMap
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
func skipLimitOrderTrancheUserShareMap(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLimitOrderTrancheUserShareMap
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
					return 0, ErrIntOverflowLimitOrderTrancheUserShareMap
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
					return 0, ErrIntOverflowLimitOrderTrancheUserShareMap
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
				return 0, ErrInvalidLengthLimitOrderTrancheUserShareMap
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLimitOrderTrancheUserShareMap
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLimitOrderTrancheUserShareMap
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLimitOrderTrancheUserShareMap        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLimitOrderTrancheUserShareMap          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLimitOrderTrancheUserShareMap = fmt.Errorf("proto: unexpected end of group")
)
