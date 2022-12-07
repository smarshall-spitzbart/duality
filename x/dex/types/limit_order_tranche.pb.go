// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/limit_order_tranche.proto

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

type LimitOrderTranche struct {
	ReservesTokenIn  github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=reservesTokenIn,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reservesTokenIn" yaml:"reservesTokenIn"`
	ReservesTokenOut github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=reservesTokenOut,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reservesTokenOut" yaml:"reservesTokenOut"`
	TotalTokenIn     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=totalTokenIn,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"totalTokenIn" yaml:"totalTokenIn"`
	TotalTokenOut    github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=totalTokenOut,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"totalTokenOut" yaml:"totalTokenOut"`
}

func (m *LimitOrderTranche) Reset()         { *m = LimitOrderTranche{} }
func (m *LimitOrderTranche) String() string { return proto.CompactTextString(m) }
func (*LimitOrderTranche) ProtoMessage()    {}
func (*LimitOrderTranche) Descriptor() ([]byte, []int) {
	return fileDescriptor_a83217ee8b948359, []int{0}
}
func (m *LimitOrderTranche) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LimitOrderTranche) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LimitOrderTranche.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LimitOrderTranche) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimitOrderTranche.Merge(m, src)
}
func (m *LimitOrderTranche) XXX_Size() int {
	return m.Size()
}
func (m *LimitOrderTranche) XXX_DiscardUnknown() {
	xxx_messageInfo_LimitOrderTranche.DiscardUnknown(m)
}

var xxx_messageInfo_LimitOrderTranche proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LimitOrderTranche)(nil), "nicholasdotsol.duality.dex.LimitOrderTranche")
}

func init() { proto.RegisterFile("dex/limit_order_tranche.proto", fileDescriptor_a83217ee8b948359) }

var fileDescriptor_a83217ee8b948359 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x49, 0xad, 0xd0,
	0xcf, 0xc9, 0xcc, 0xcd, 0x2c, 0x89, 0xcf, 0x2f, 0x4a, 0x49, 0x2d, 0x8a, 0x2f, 0x29, 0x4a, 0xcc,
	0x4b, 0xce, 0x48, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0xca, 0xcb, 0x4c, 0xce, 0xc8,
	0xcf, 0x49, 0x2c, 0x4e, 0xc9, 0x2f, 0x29, 0xce, 0xcf, 0xd1, 0x4b, 0x29, 0x4d, 0xcc, 0xc9, 0x2c,
	0xa9, 0xd4, 0x4b, 0x49, 0xad, 0x90, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x2b, 0xd3, 0x07, 0xb1,
	0x20, 0x3a, 0x94, 0x16, 0xb2, 0x70, 0x09, 0xfa, 0x80, 0xcc, 0xf3, 0x07, 0x19, 0x17, 0x02, 0x31,
	0x4d, 0xa8, 0x95, 0x91, 0x8b, 0xbf, 0x28, 0xb5, 0x38, 0xb5, 0xa8, 0x2c, 0xb5, 0x38, 0x24, 0x3f,
	0x3b, 0x35, 0xcf, 0x33, 0x4f, 0x82, 0x55, 0x81, 0x51, 0x83, 0xd3, 0x29, 0xfa, 0xc4, 0x3d, 0x79,
	0x86, 0x5b, 0xf7, 0xe4, 0xd5, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5,
	0x93, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0xa1, 0x94, 0x6e, 0x71, 0x4a, 0xb6, 0x7e, 0x49, 0x65, 0x41,
	0x6a, 0xb1, 0x9e, 0x4b, 0x6a, 0xf2, 0xab, 0x7b, 0xf2, 0xe8, 0x06, 0x7d, 0xba, 0x27, 0x2f, 0x56,
	0x99, 0x98, 0x9b, 0x63, 0xa5, 0x84, 0x26, 0xa1, 0x14, 0x84, 0xae, 0x54, 0xa8, 0x93, 0x91, 0x4b,
	0x00, 0x45, 0xcc, 0xbf, 0xb4, 0x44, 0x82, 0x0d, 0xec, 0x90, 0x58, 0x92, 0x1d, 0x82, 0x61, 0xd2,
	0xa7, 0x7b, 0xf2, 0xe2, 0x58, 0x5c, 0xe2, 0x5f, 0x5a, 0xa2, 0x14, 0x84, 0xa1, 0x58, 0xa8, 0x92,
	0x8b, 0xa7, 0x24, 0xbf, 0x24, 0x31, 0x07, 0x16, 0x1e, 0xec, 0x60, 0x67, 0x84, 0x92, 0xec, 0x0c,
	0x14, 0x53, 0x3e, 0xdd, 0x93, 0x17, 0x86, 0x38, 0x01, 0x59, 0x54, 0x29, 0x08, 0x45, 0x91, 0x50,
	0x2d, 0x17, 0x2f, 0x82, 0x0f, 0x0a, 0x02, 0x0e, 0xb0, 0xdd, 0xe1, 0x24, 0xdb, 0x8d, 0x6a, 0xcc,
	0xa7, 0x7b, 0xf2, 0x22, 0xe8, 0x96, 0x83, 0x3d, 0x8f, 0xaa, 0xcc, 0xc9, 0xfd, 0xc4, 0x23, 0x39,
	0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63,
	0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x74, 0x91, 0x6c, 0xf6, 0x83, 0x26, 0x3d, 0x97, 0xfc,
	0x92, 0xe0, 0xfc, 0x1c, 0x7d, 0x68, 0xd2, 0xd3, 0xaf, 0xd0, 0x07, 0x25, 0x59, 0xb0, 0x23, 0x92,
	0xd8, 0xc0, 0x69, 0xce, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x63, 0xdd, 0x78, 0xa7, 0xc6, 0x02,
	0x00, 0x00,
}

func (m *LimitOrderTranche) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LimitOrderTranche) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LimitOrderTranche) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.TotalTokenOut.Size()
		i -= size
		if _, err := m.TotalTokenOut.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTranche(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.TotalTokenIn.Size()
		i -= size
		if _, err := m.TotalTokenIn.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTranche(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.ReservesTokenOut.Size()
		i -= size
		if _, err := m.ReservesTokenOut.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTranche(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.ReservesTokenIn.Size()
		i -= size
		if _, err := m.ReservesTokenIn.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderTranche(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	return len(dAtA) - i, nil
}

func encodeVarintLimitOrderTranche(dAtA []byte, offset int, v uint64) int {
	offset -= sovLimitOrderTranche(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LimitOrderTranche) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.ReservesTokenIn.Size()
	n += 1 + l + sovLimitOrderTranche(uint64(l))
	l = m.ReservesTokenOut.Size()
	n += 1 + l + sovLimitOrderTranche(uint64(l))
	l = m.TotalTokenIn.Size()
	n += 1 + l + sovLimitOrderTranche(uint64(l))
	l = m.TotalTokenOut.Size()
	n += 1 + l + sovLimitOrderTranche(uint64(l))
	return n
}

func sovLimitOrderTranche(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLimitOrderTranche(x uint64) (n int) {
	return sovLimitOrderTranche(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LimitOrderTranche) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLimitOrderTranche
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
			return fmt.Errorf("proto: LimitOrderTranche: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LimitOrderTranche: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReservesTokenIn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
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
				return ErrInvalidLengthLimitOrderTranche
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTranche
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReservesTokenIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReservesTokenOut", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
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
				return ErrInvalidLengthLimitOrderTranche
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTranche
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReservesTokenOut.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalTokenIn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
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
				return ErrInvalidLengthLimitOrderTranche
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTranche
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalTokenIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalTokenOut", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderTranche
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
				return ErrInvalidLengthLimitOrderTranche
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderTranche
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TotalTokenOut.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLimitOrderTranche(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLimitOrderTranche
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
func skipLimitOrderTranche(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLimitOrderTranche
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
					return 0, ErrIntOverflowLimitOrderTranche
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
					return 0, ErrIntOverflowLimitOrderTranche
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
				return 0, ErrInvalidLengthLimitOrderTranche
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLimitOrderTranche
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLimitOrderTranche
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLimitOrderTranche        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLimitOrderTranche          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLimitOrderTranche = fmt.Errorf("proto: unexpected end of group")
)
