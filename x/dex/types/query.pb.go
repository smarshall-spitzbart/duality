// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dex/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8e98105e6e08a59, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8e98105e6e08a59, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type QueryGetTickMapRequest struct {
	TickIndex string `protobuf:"bytes,1,opt,name=tickIndex,proto3" json:"tickIndex,omitempty"`
}

func (m *QueryGetTickMapRequest) Reset()         { *m = QueryGetTickMapRequest{} }
func (m *QueryGetTickMapRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetTickMapRequest) ProtoMessage()    {}
func (*QueryGetTickMapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8e98105e6e08a59, []int{2}
}
func (m *QueryGetTickMapRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetTickMapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetTickMapRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetTickMapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetTickMapRequest.Merge(m, src)
}
func (m *QueryGetTickMapRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetTickMapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetTickMapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetTickMapRequest proto.InternalMessageInfo

func (m *QueryGetTickMapRequest) GetTickIndex() string {
	if m != nil {
		return m.TickIndex
	}
	return ""
}

type QueryGetTickMapResponse struct {
	TickMap TickMap `protobuf:"bytes,1,opt,name=tickMap,proto3" json:"tickMap"`
}

func (m *QueryGetTickMapResponse) Reset()         { *m = QueryGetTickMapResponse{} }
func (m *QueryGetTickMapResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetTickMapResponse) ProtoMessage()    {}
func (*QueryGetTickMapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8e98105e6e08a59, []int{3}
}
func (m *QueryGetTickMapResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetTickMapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetTickMapResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetTickMapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetTickMapResponse.Merge(m, src)
}
func (m *QueryGetTickMapResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetTickMapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetTickMapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetTickMapResponse proto.InternalMessageInfo

func (m *QueryGetTickMapResponse) GetTickMap() TickMap {
	if m != nil {
		return m.TickMap
	}
	return TickMap{}
}

type QueryAllTickMapRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllTickMapRequest) Reset()         { *m = QueryAllTickMapRequest{} }
func (m *QueryAllTickMapRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllTickMapRequest) ProtoMessage()    {}
func (*QueryAllTickMapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8e98105e6e08a59, []int{4}
}
func (m *QueryAllTickMapRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllTickMapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllTickMapRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllTickMapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllTickMapRequest.Merge(m, src)
}
func (m *QueryAllTickMapRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllTickMapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllTickMapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllTickMapRequest proto.InternalMessageInfo

func (m *QueryAllTickMapRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllTickMapResponse struct {
	TickMap    []TickMap           `protobuf:"bytes,1,rep,name=tickMap,proto3" json:"tickMap"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllTickMapResponse) Reset()         { *m = QueryAllTickMapResponse{} }
func (m *QueryAllTickMapResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllTickMapResponse) ProtoMessage()    {}
func (*QueryAllTickMapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d8e98105e6e08a59, []int{5}
}
func (m *QueryAllTickMapResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllTickMapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllTickMapResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllTickMapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllTickMapResponse.Merge(m, src)
}
func (m *QueryAllTickMapResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllTickMapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllTickMapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllTickMapResponse proto.InternalMessageInfo

func (m *QueryAllTickMapResponse) GetTickMap() []TickMap {
	if m != nil {
		return m.TickMap
	}
	return nil
}

func (m *QueryAllTickMapResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "nicholasdotsol.duality.dex.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "nicholasdotsol.duality.dex.QueryParamsResponse")
	proto.RegisterType((*QueryGetTickMapRequest)(nil), "nicholasdotsol.duality.dex.QueryGetTickMapRequest")
	proto.RegisterType((*QueryGetTickMapResponse)(nil), "nicholasdotsol.duality.dex.QueryGetTickMapResponse")
	proto.RegisterType((*QueryAllTickMapRequest)(nil), "nicholasdotsol.duality.dex.QueryAllTickMapRequest")
	proto.RegisterType((*QueryAllTickMapResponse)(nil), "nicholasdotsol.duality.dex.QueryAllTickMapResponse")
}

func init() { proto.RegisterFile("dex/query.proto", fileDescriptor_d8e98105e6e08a59) }

var fileDescriptor_d8e98105e6e08a59 = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x41, 0x6b, 0x13, 0x41,
	0x14, 0xc7, 0x33, 0xad, 0xa6, 0x74, 0x3c, 0x28, 0x63, 0x51, 0x59, 0xca, 0x2a, 0x63, 0xa9, 0x52,
	0x74, 0xc6, 0xa6, 0x20, 0x1e, 0x6d, 0x15, 0x83, 0x07, 0xa5, 0x46, 0x41, 0xf0, 0xa0, 0x4e, 0x76,
	0x87, 0xed, 0xd2, 0xc9, 0xce, 0x36, 0x33, 0x91, 0x04, 0xf1, 0xe2, 0x27, 0x10, 0xc4, 0x4f, 0xe0,
	0xc1, 0x8b, 0x1f, 0xa4, 0xc7, 0x82, 0x17, 0x4f, 0x22, 0x89, 0x1f, 0xc1, 0x0f, 0x20, 0x3b, 0xf3,
	0x62, 0x52, 0xd7, 0xa6, 0xab, 0xb7, 0xf0, 0xf2, 0xfe, 0xff, 0xf7, 0x7b, 0xf3, 0xfe, 0x09, 0x3e,
	0x1d, 0xcb, 0x3e, 0xdf, 0xeb, 0xc9, 0xee, 0x80, 0xe5, 0x5d, 0x6d, 0x35, 0x09, 0xb2, 0x34, 0xda,
	0xd1, 0x4a, 0x98, 0x58, 0x5b, 0xa3, 0x15, 0x8b, 0x7b, 0x42, 0xa5, 0x76, 0xc0, 0x62, 0xd9, 0x0f,
	0x96, 0x12, 0x9d, 0x68, 0xd7, 0xc6, 0x8b, 0x4f, 0x5e, 0x11, 0x2c, 0x27, 0x5a, 0x27, 0x4a, 0x72,
	0x91, 0xa7, 0x5c, 0x64, 0x99, 0xb6, 0xc2, 0xa6, 0x3a, 0x33, 0xf0, 0xed, 0x5a, 0xa4, 0x4d, 0x47,
	0x1b, 0xde, 0x16, 0x46, 0xfa, 0x41, 0xfc, 0xd5, 0x7a, 0x5b, 0x5a, 0xb1, 0xce, 0x73, 0x91, 0xa4,
	0x99, 0x6b, 0x86, 0xde, 0x33, 0x05, 0x4c, 0x2e, 0xba, 0xa2, 0x33, 0x56, 0x93, 0xa2, 0x62, 0xd3,
	0x68, 0xf7, 0x45, 0x47, 0xe4, 0xbe, 0x46, 0x97, 0x30, 0x79, 0x54, 0xf8, 0x6c, 0xbb, 0xc6, 0x96,
	0xdc, 0xeb, 0x49, 0x63, 0xe9, 0x53, 0x7c, 0xf6, 0x50, 0xd5, 0xe4, 0x3a, 0x33, 0x92, 0xdc, 0xc6,
	0x75, 0x6f, 0x78, 0x01, 0x5d, 0x42, 0x57, 0x4f, 0x35, 0x28, 0x3b, 0x7a, 0x3f, 0xe6, 0xb5, 0x5b,
	0x27, 0xf6, 0xbf, 0x5d, 0xac, 0xb5, 0x40, 0x47, 0x6f, 0xe2, 0x73, 0xce, 0xb8, 0x29, 0xed, 0x93,
	0x34, 0xda, 0x7d, 0x20, 0x72, 0x18, 0x49, 0x96, 0xf1, 0x62, 0x81, 0x76, 0x3f, 0x8b, 0x65, 0xdf,
	0xd9, 0x2f, 0xb6, 0x26, 0x05, 0xfa, 0x1c, 0x9f, 0x2f, 0xe9, 0x00, 0xea, 0x0e, 0x5e, 0xb0, 0xbe,
	0x04, 0x54, 0x97, 0x67, 0x51, 0x81, 0x1a, 0xb0, 0xc6, 0x4a, 0xfa, 0x12, 0xb8, 0x36, 0x95, 0xfa,
	0x83, 0xeb, 0x1e, 0xc6, 0x93, 0xa7, 0x85, 0x09, 0xab, 0xcc, 0xdf, 0x81, 0x15, 0x77, 0x60, 0xfe,
	0xe0, 0x70, 0x07, 0xb6, 0x2d, 0x12, 0x09, 0xda, 0xd6, 0x94, 0x92, 0x7e, 0x42, 0xb0, 0xc2, 0xf4,
	0x88, 0xbf, 0xad, 0x30, 0xff, 0x7f, 0x2b, 0x90, 0xe6, 0x21, 0xd0, 0x39, 0x07, 0x7a, 0xe5, 0x58,
	0x50, 0x4f, 0x30, 0x4d, 0xda, 0xf8, 0x39, 0x8f, 0x4f, 0x3a, 0x52, 0xf2, 0x01, 0xe1, 0xba, 0x3f,
	0x23, 0x61, 0xb3, 0x88, 0xca, 0x09, 0x0a, 0x78, 0xe5, 0x7e, 0x4f, 0x40, 0xd7, 0xde, 0x7e, 0xf9,
	0xf1, 0x7e, 0x6e, 0x85, 0x50, 0xfe, 0x10, 0x84, 0x77, 0xb5, 0x7d, 0xac, 0x15, 0x07, 0x21, 0x9f,
	0xc4, 0x99, 0x7c, 0x46, 0x78, 0x01, 0x5e, 0x81, 0x34, 0x8e, 0x1d, 0x54, 0xca, 0x5a, 0xb0, 0xf1,
	0x4f, 0x1a, 0x00, 0xbc, 0xe5, 0x00, 0x1b, 0xe4, 0xc6, 0x2c, 0xc0, 0xf1, 0xaf, 0x8b, 0xbf, 0xfe,
	0x9d, 0xdd, 0x37, 0xe4, 0x23, 0xc2, 0x18, 0xdc, 0x36, 0x95, 0xaa, 0x40, 0x5c, 0x4a, 0x61, 0x05,
	0xe2, 0x72, 0xac, 0xe8, 0x35, 0x47, 0xbc, 0x4a, 0x56, 0xaa, 0x10, 0x6f, 0x35, 0xf7, 0x87, 0x21,
	0x3a, 0x18, 0x86, 0xe8, 0xfb, 0x30, 0x44, 0xef, 0x46, 0x61, 0xed, 0x60, 0x14, 0xd6, 0xbe, 0x8e,
	0xc2, 0xda, 0xb3, 0xeb, 0x49, 0x6a, 0x77, 0x7a, 0x6d, 0x16, 0xe9, 0xce, 0x51, 0x4e, 0x7d, 0xef,
	0x35, 0xc8, 0xa5, 0x69, 0xd7, 0xdd, 0x3f, 0xcb, 0xc6, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2a,
	0xaf, 0xda, 0x09, 0x0e, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a TickMap by index.
	TickMap(ctx context.Context, in *QueryGetTickMapRequest, opts ...grpc.CallOption) (*QueryGetTickMapResponse, error)
	// Queries a list of TickMap items.
	TickMapAll(ctx context.Context, in *QueryAllTickMapRequest, opts ...grpc.CallOption) (*QueryAllTickMapResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/nicholasdotsol.duality.dex.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TickMap(ctx context.Context, in *QueryGetTickMapRequest, opts ...grpc.CallOption) (*QueryGetTickMapResponse, error) {
	out := new(QueryGetTickMapResponse)
	err := c.cc.Invoke(ctx, "/nicholasdotsol.duality.dex.Query/TickMap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TickMapAll(ctx context.Context, in *QueryAllTickMapRequest, opts ...grpc.CallOption) (*QueryAllTickMapResponse, error) {
	out := new(QueryAllTickMapResponse)
	err := c.cc.Invoke(ctx, "/nicholasdotsol.duality.dex.Query/TickMapAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a TickMap by index.
	TickMap(context.Context, *QueryGetTickMapRequest) (*QueryGetTickMapResponse, error)
	// Queries a list of TickMap items.
	TickMapAll(context.Context, *QueryAllTickMapRequest) (*QueryAllTickMapResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) TickMap(ctx context.Context, req *QueryGetTickMapRequest) (*QueryGetTickMapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TickMap not implemented")
}
func (*UnimplementedQueryServer) TickMapAll(ctx context.Context, req *QueryAllTickMapRequest) (*QueryAllTickMapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TickMapAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nicholasdotsol.duality.dex.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TickMap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetTickMapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TickMap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nicholasdotsol.duality.dex.Query/TickMap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TickMap(ctx, req.(*QueryGetTickMapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TickMapAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllTickMapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TickMapAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/nicholasdotsol.duality.dex.Query/TickMapAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TickMapAll(ctx, req.(*QueryAllTickMapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "nicholasdotsol.duality.dex.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "TickMap",
			Handler:    _Query_TickMap_Handler,
		},
		{
			MethodName: "TickMapAll",
			Handler:    _Query_TickMapAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dex/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryGetTickMapRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetTickMapRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetTickMapRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TickIndex) > 0 {
		i -= len(m.TickIndex)
		copy(dAtA[i:], m.TickIndex)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.TickIndex)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetTickMapResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetTickMapResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetTickMapResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.TickMap.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllTickMapRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllTickMapRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllTickMapRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllTickMapResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllTickMapResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllTickMapResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.TickMap) > 0 {
		for iNdEx := len(m.TickMap) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.TickMap[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryGetTickMapRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TickIndex)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetTickMapResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.TickMap.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllTickMapRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllTickMapResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TickMap) > 0 {
		for _, e := range m.TickMap {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetTickMapRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetTickMapRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetTickMapRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickIndex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TickIndex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetTickMapResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetTickMapResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetTickMapResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickMap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TickMap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllTickMapRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllTickMapRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllTickMapRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllTickMapResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllTickMapResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllTickMapResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TickMap", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TickMap = append(m.TickMap, TickMap{})
			if err := m.TickMap[len(m.TickMap)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
