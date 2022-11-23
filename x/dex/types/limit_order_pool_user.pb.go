// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/limit_order_pool_user.proto

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

type LimitOrderPoolUser struct {
	PairId          string                                 `protobuf:"bytes,1,opt,name=pairId,proto3" json:"pairId,omitempty"`
	Token           string                                 `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	TickIndex       int64                                  `protobuf:"varint,3,opt,name=tickIndex,proto3" json:"tickIndex,omitempty"`
	Count           uint64                                 `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	Address         string                                 `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	SharesOwned     github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=sharesOwned,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"sharesOwned" yaml:"sharesOwned"`
	SharesWithdrawn github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,7,opt,name=sharesWithdrawn,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"sharesWithdrawn" yaml:"sharesWithdrawn"`
	SharesCancelled github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,8,opt,name=sharesCancelled,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"sharesCancelled" yaml:"sharesCancelled"`
}

func (m *LimitOrderPoolUser) Reset()         { *m = LimitOrderPoolUser{} }
func (m *LimitOrderPoolUser) String() string { return proto.CompactTextString(m) }
func (*LimitOrderPoolUser) ProtoMessage()    {}
func (*LimitOrderPoolUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_27a55c4773c00f10, []int{0}
}
func (m *LimitOrderPoolUser) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LimitOrderPoolUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LimitOrderPoolUser.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LimitOrderPoolUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LimitOrderPoolUser.Merge(m, src)
}
func (m *LimitOrderPoolUser) XXX_Size() int {
	return m.Size()
}
func (m *LimitOrderPoolUser) XXX_DiscardUnknown() {
	xxx_messageInfo_LimitOrderPoolUser.DiscardUnknown(m)
}

var xxx_messageInfo_LimitOrderPoolUser proto.InternalMessageInfo

func (m *LimitOrderPoolUser) GetPairId() string {
	if m != nil {
		return m.PairId
	}
	return ""
}

func (m *LimitOrderPoolUser) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *LimitOrderPoolUser) GetTickIndex() int64 {
	if m != nil {
		return m.TickIndex
	}
	return 0
}

func (m *LimitOrderPoolUser) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *LimitOrderPoolUser) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*LimitOrderPoolUser)(nil), "nicholasdotsol.duality.dex.LimitOrderPoolUser")
}

func init() { proto.RegisterFile("dex/limit_order_pool_user.proto", fileDescriptor_27a55c4773c00f10) }

var fileDescriptor_27a55c4773c00f10 = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xbf, 0x6b, 0xdb, 0x40,
	0x18, 0xd5, 0x35, 0x8e, 0xd3, 0x5c, 0x87, 0xc2, 0x11, 0xc2, 0x11, 0x8a, 0x64, 0x34, 0x14, 0x2f,
	0x91, 0x86, 0x6e, 0x1d, 0xd3, 0x40, 0x09, 0x94, 0xa6, 0x28, 0x94, 0x42, 0x3b, 0x98, 0xb3, 0xee,
	0xb0, 0x0e, 0x9f, 0xf4, 0x89, 0xbb, 0x13, 0x96, 0xff, 0x80, 0xee, 0xdd, 0xfa, 0x2f, 0x79, 0xf4,
	0x58, 0x3a, 0x88, 0x62, 0x6f, 0x1d, 0xfd, 0x17, 0x14, 0xfd, 0xa8, 0x2a, 0xbc, 0x95, 0x4c, 0x77,
	0xef, 0x7d, 0xdf, 0xf7, 0xde, 0x1b, 0x1e, 0xf6, 0xb8, 0x28, 0x43, 0x25, 0x53, 0x69, 0x67, 0xa0,
	0xb9, 0xd0, 0xb3, 0x1c, 0x40, 0xcd, 0x0a, 0x23, 0x74, 0x90, 0x6b, 0xb0, 0x40, 0xae, 0x32, 0x19,
	0x27, 0xa0, 0x98, 0xe1, 0x60, 0x0d, 0xa8, 0x80, 0x17, 0x4c, 0x49, 0xbb, 0x0e, 0xb8, 0x28, 0xaf,
	0x2e, 0x16, 0xb0, 0x80, 0x66, 0x2d, 0xac, 0x7f, 0xed, 0x85, 0xff, 0x7d, 0x84, 0xc9, 0xbb, 0x5a,
	0xf1, 0xbe, 0x16, 0xfc, 0x00, 0xa0, 0x3e, 0x1a, 0xa1, 0xc9, 0x25, 0x1e, 0xe7, 0x4c, 0xea, 0x3b,
	0x4e, 0xd1, 0x04, 0x4d, 0xcf, 0xa3, 0x0e, 0x91, 0x0b, 0x7c, 0x6a, 0x61, 0x29, 0x32, 0xfa, 0xa4,
	0xa1, 0x5b, 0x40, 0x5e, 0xe0, 0x73, 0x2b, 0xe3, 0xe5, 0x5d, 0xc6, 0x45, 0x49, 0x4f, 0x26, 0x68,
	0x7a, 0x12, 0xfd, 0x23, 0xea, 0x9b, 0x18, 0x8a, 0xcc, 0xd2, 0xd1, 0x04, 0x4d, 0x47, 0x51, 0x0b,
	0x08, 0xc5, 0x67, 0x8c, 0x73, 0x2d, 0x8c, 0xa1, 0xa7, 0x8d, 0xd6, 0x5f, 0x48, 0x0a, 0xfc, 0xcc,
	0x24, 0x4c, 0x0b, 0x73, 0xbf, 0xca, 0x04, 0xa7, 0xe3, 0x7a, 0x7a, 0xf3, 0xb0, 0xa9, 0x3c, 0xe7,
	0x67, 0xe5, 0xbd, 0x5c, 0x48, 0x9b, 0x14, 0xf3, 0x20, 0x86, 0x34, 0x8c, 0xc1, 0xa4, 0x60, 0xba,
	0xe7, 0xda, 0xf0, 0x65, 0x68, 0xd7, 0xb9, 0x30, 0xc1, 0xad, 0x88, 0x7f, 0x57, 0xde, 0x50, 0xe4,
	0x50, 0x79, 0x64, 0xcd, 0x52, 0xf5, 0xda, 0x1f, 0x90, 0x7e, 0x34, 0x5c, 0x21, 0x5f, 0x11, 0x7e,
	0xde, 0xe2, 0x4f, 0xd2, 0x26, 0x5c, 0xb3, 0x55, 0x46, 0xcf, 0x1a, 0xef, 0x2f, 0xff, 0xed, 0x7d,
	0x2c, 0x74, 0xa8, 0xbc, 0xcb, 0xa1, 0x7f, 0x3f, 0xf0, 0xa3, 0xe3, 0xd5, 0x41, 0x8e, 0x37, 0x2c,
	0x8b, 0x85, 0x52, 0x82, 0xd3, 0xa7, 0x8f, 0xcb, 0xd1, 0x0b, 0x1d, 0xe7, 0xe8, 0x07, 0x7d, 0x8e,
	0x9e, 0xb9, 0x79, 0xbb, 0xd9, 0xb9, 0x68, 0xbb, 0x73, 0xd1, 0xaf, 0x9d, 0x8b, 0xbe, 0xed, 0x5d,
	0x67, 0xbb, 0x77, 0x9d, 0x1f, 0x7b, 0xd7, 0xf9, 0x7c, 0x3d, 0xf0, 0x7f, 0xdf, 0x15, 0xee, 0x16,
	0xec, 0x03, 0xa8, 0xb0, 0x2b, 0x5c, 0x58, 0x86, 0x75, 0x55, 0x9b, 0x28, 0xf3, 0x71, 0xd3, 0xb4,
	0x57, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x68, 0x62, 0x01, 0xf9, 0xbe, 0x02, 0x00, 0x00,
}

func (m *LimitOrderPoolUser) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LimitOrderPoolUser) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LimitOrderPoolUser) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.SharesCancelled.Size()
		i -= size
		if _, err := m.SharesCancelled.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderPoolUser(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.SharesWithdrawn.Size()
		i -= size
		if _, err := m.SharesWithdrawn.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderPoolUser(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.SharesOwned.Size()
		i -= size
		if _, err := m.SharesOwned.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLimitOrderPoolUser(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintLimitOrderPoolUser(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Count != 0 {
		i = encodeVarintLimitOrderPoolUser(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x20
	}
	if m.TickIndex != 0 {
		i = encodeVarintLimitOrderPoolUser(dAtA, i, uint64(m.TickIndex))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintLimitOrderPoolUser(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PairId) > 0 {
		i -= len(m.PairId)
		copy(dAtA[i:], m.PairId)
		i = encodeVarintLimitOrderPoolUser(dAtA, i, uint64(len(m.PairId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintLimitOrderPoolUser(dAtA []byte, offset int, v uint64) int {
	offset -= sovLimitOrderPoolUser(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LimitOrderPoolUser) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PairId)
	if l > 0 {
		n += 1 + l + sovLimitOrderPoolUser(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovLimitOrderPoolUser(uint64(l))
	}
	if m.TickIndex != 0 {
		n += 1 + sovLimitOrderPoolUser(uint64(m.TickIndex))
	}
	if m.Count != 0 {
		n += 1 + sovLimitOrderPoolUser(uint64(m.Count))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovLimitOrderPoolUser(uint64(l))
	}
	l = m.SharesOwned.Size()
	n += 1 + l + sovLimitOrderPoolUser(uint64(l))
	l = m.SharesWithdrawn.Size()
	n += 1 + l + sovLimitOrderPoolUser(uint64(l))
	l = m.SharesCancelled.Size()
	n += 1 + l + sovLimitOrderPoolUser(uint64(l))
	return n
}

func sovLimitOrderPoolUser(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLimitOrderPoolUser(x uint64) (n int) {
	return sovLimitOrderPoolUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LimitOrderPoolUser) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLimitOrderPoolUser
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
			return fmt.Errorf("proto: LimitOrderPoolUser: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LimitOrderPoolUser: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderPoolUser
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
				return ErrInvalidLengthLimitOrderPoolUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderPoolUser
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
					return ErrIntOverflowLimitOrderPoolUser
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
				return ErrInvalidLengthLimitOrderPoolUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderPoolUser
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
					return ErrIntOverflowLimitOrderPoolUser
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
					return ErrIntOverflowLimitOrderPoolUser
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
					return ErrIntOverflowLimitOrderPoolUser
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
				return ErrInvalidLengthLimitOrderPoolUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderPoolUser
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
					return ErrIntOverflowLimitOrderPoolUser
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
				return ErrInvalidLengthLimitOrderPoolUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderPoolUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SharesOwned.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SharesWithdrawn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderPoolUser
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
				return ErrInvalidLengthLimitOrderPoolUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderPoolUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SharesWithdrawn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SharesCancelled", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLimitOrderPoolUser
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
				return ErrInvalidLengthLimitOrderPoolUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLimitOrderPoolUser
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SharesCancelled.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLimitOrderPoolUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLimitOrderPoolUser
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
func skipLimitOrderPoolUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLimitOrderPoolUser
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
					return 0, ErrIntOverflowLimitOrderPoolUser
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
					return 0, ErrIntOverflowLimitOrderPoolUser
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
				return 0, ErrInvalidLengthLimitOrderPoolUser
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLimitOrderPoolUser
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLimitOrderPoolUser
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLimitOrderPoolUser        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLimitOrderPoolUser          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLimitOrderPoolUser = fmt.Errorf("proto: unexpected end of group")
)