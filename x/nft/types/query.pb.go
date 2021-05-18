// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: shentu/nft/v1alpha1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/irisnet/irismod/modules/nft/types"
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

type QueryAdminRequest struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *QueryAdminRequest) Reset()         { *m = QueryAdminRequest{} }
func (m *QueryAdminRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAdminRequest) ProtoMessage()    {}
func (*QueryAdminRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_07aaa7e667d22cca, []int{0}
}
func (m *QueryAdminRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAdminRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAdminRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAdminRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAdminRequest.Merge(m, src)
}
func (m *QueryAdminRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAdminRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAdminRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAdminRequest proto.InternalMessageInfo

func (m *QueryAdminRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type QueryAdminResponse struct {
	Admin Admin `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin"`
}

func (m *QueryAdminResponse) Reset()         { *m = QueryAdminResponse{} }
func (m *QueryAdminResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAdminResponse) ProtoMessage()    {}
func (*QueryAdminResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_07aaa7e667d22cca, []int{1}
}
func (m *QueryAdminResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAdminResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAdminResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAdminResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAdminResponse.Merge(m, src)
}
func (m *QueryAdminResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAdminResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAdminResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAdminResponse proto.InternalMessageInfo

func (m *QueryAdminResponse) GetAdmin() Admin {
	if m != nil {
		return m.Admin
	}
	return Admin{}
}

type QueryAdminsRequest struct {
}

func (m *QueryAdminsRequest) Reset()         { *m = QueryAdminsRequest{} }
func (m *QueryAdminsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAdminsRequest) ProtoMessage()    {}
func (*QueryAdminsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_07aaa7e667d22cca, []int{2}
}
func (m *QueryAdminsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAdminsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAdminsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAdminsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAdminsRequest.Merge(m, src)
}
func (m *QueryAdminsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAdminsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAdminsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAdminsRequest proto.InternalMessageInfo

type QueryAdminsResponse struct {
	Admins []Admin `protobuf:"bytes,1,rep,name=admins,proto3" json:"admins"`
}

func (m *QueryAdminsResponse) Reset()         { *m = QueryAdminsResponse{} }
func (m *QueryAdminsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAdminsResponse) ProtoMessage()    {}
func (*QueryAdminsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_07aaa7e667d22cca, []int{3}
}
func (m *QueryAdminsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAdminsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAdminsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAdminsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAdminsResponse.Merge(m, src)
}
func (m *QueryAdminsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAdminsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAdminsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAdminsResponse proto.InternalMessageInfo

func (m *QueryAdminsResponse) GetAdmins() []Admin {
	if m != nil {
		return m.Admins
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryAdminRequest)(nil), "shentu.nft.v1alpha1.QueryAdminRequest")
	proto.RegisterType((*QueryAdminResponse)(nil), "shentu.nft.v1alpha1.QueryAdminResponse")
	proto.RegisterType((*QueryAdminsRequest)(nil), "shentu.nft.v1alpha1.QueryAdminsRequest")
	proto.RegisterType((*QueryAdminsResponse)(nil), "shentu.nft.v1alpha1.QueryAdminsResponse")
}

func init() { proto.RegisterFile("shentu/nft/v1alpha1/query.proto", fileDescriptor_07aaa7e667d22cca) }

var fileDescriptor_07aaa7e667d22cca = []byte{
	// 568 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x86, 0x63, 0x4a, 0x52, 0x31, 0xac, 0x7a, 0x52, 0x44, 0x31, 0xad, 0x13, 0xa5, 0xf4, 0xc2,
	0xcd, 0xa3, 0x06, 0x09, 0xb1, 0xa5, 0xa0, 0x8a, 0x05, 0x6a, 0x45, 0xe8, 0x8a, 0x0d, 0x72, 0xe3,
	0x49, 0x62, 0x35, 0x99, 0x71, 0x3d, 0x63, 0x68, 0x14, 0x55, 0x95, 0x10, 0x0f, 0x80, 0xc4, 0x0b,
	0xb1, 0xec, 0xb2, 0x12, 0x1b, 0x56, 0x08, 0x25, 0x3c, 0x08, 0xf2, 0xf1, 0xb8, 0xd8, 0xb2, 0x93,
	0x74, 0x67, 0xcf, 0xff, 0xcf, 0xf9, 0xfe, 0x33, 0x37, 0x52, 0x93, 0x3d, 0xc6, 0x55, 0x48, 0x79,
	0x47, 0xd1, 0x4f, 0x3b, 0x4e, 0xdf, 0xef, 0x39, 0x3b, 0xf4, 0x24, 0x64, 0xc1, 0xd0, 0xf6, 0x03,
	0xa1, 0x04, 0x54, 0x63, 0x83, 0xcd, 0x3b, 0xca, 0x4e, 0x0c, 0xe6, 0x72, 0x57, 0x74, 0x05, 0xea,
	0x34, 0xfa, 0x8a, 0xad, 0xe6, 0x6a, 0x57, 0x88, 0x6e, 0x9f, 0x51, 0xc7, 0xf7, 0xa8, 0xc3, 0xb9,
	0x50, 0x8e, 0xf2, 0x04, 0x97, 0x5a, 0xbd, 0xeb, 0x05, 0x9e, 0x1c, 0x08, 0x17, 0x51, 0x29, 0x82,
	0xb9, 0x56, 0x14, 0x21, 0xc2, 0xa1, 0xdc, 0x78, 0x4a, 0x96, 0xde, 0x45, 0xee, 0x97, 0xee, 0xc0,
	0xe3, 0x2d, 0x76, 0x12, 0x32, 0xa9, 0x60, 0x85, 0x2c, 0x3a, 0xae, 0x1b, 0x30, 0x29, 0x57, 0x8c,
	0xba, 0xb1, 0x7d, 0xab, 0x95, 0xfc, 0x36, 0xde, 0x12, 0x48, 0xdb, 0xa5, 0x2f, 0xb8, 0x64, 0xf0,
	0x9c, 0x94, 0x9d, 0x68, 0x00, 0xdd, 0xb7, 0x9b, 0xa6, 0x5d, 0xd0, 0x95, 0x8d, 0x53, 0x76, 0x6f,
	0x5e, 0xfc, 0xae, 0x95, 0x5a, 0xb1, 0xbd, 0xb1, 0x9c, 0xae, 0x26, 0x35, 0xbd, 0x71, 0x40, 0xaa,
	0x99, 0x51, 0x0d, 0x79, 0x41, 0x2a, 0x38, 0x2b, 0xca, 0xb4, 0x70, 0x2d, 0x8a, 0xf6, 0x37, 0x7f,
	0x2c, 0x92, 0x32, 0x56, 0x84, 0x73, 0x52, 0x79, 0x1f, 0xfa, 0x7e, 0x7f, 0x08, 0x35, 0x5b, 0x2f,
	0x18, 0x4e, 0x47, 0x35, 0x56, 0x74, 0x0a, 0xb3, 0x3e, 0xdd, 0x10, 0x07, 0x6a, 0x34, 0xbf, 0xfc,
	0xfc, 0xfb, 0xfd, 0xc6, 0x13, 0x78, 0x44, 0xd3, 0x6b, 0xdf, 0x16, 0xfd, 0x3e, 0x6b, 0xe3, 0xd6,
	0xd0, 0x91, 0xcb, 0xb8, 0x18, 0x7c, 0xf4, 0xdc, 0x33, 0x2a, 0x63, 0x6c, 0x9b, 0x94, 0x0f, 0x3e,
	0x73, 0x16, 0x80, 0x95, 0x2f, 0x8f, 0x42, 0x82, 0xaf, 0x4d, 0xd5, 0x35, 0xfd, 0x1e, 0xd2, 0xab,
	0xb0, 0x94, 0xa1, 0xf3, 0x8e, 0x92, 0xf0, 0xd5, 0x20, 0xe4, 0xd5, 0x55, 0x0c, 0x58, 0xcf, 0x97,
	0xfa, 0xaf, 0x26, 0xbc, 0x07, 0xb3, 0x4d, 0x1a, 0xfa, 0x18, 0xa1, 0x1b, 0xb0, 0x7e, 0x8d, 0x96,
	0xc1, 0x27, 0xe5, 0xd7, 0xd1, 0x4f, 0x51, 0xaf, 0x28, 0xcc, 0xe8, 0x55, 0xeb, 0x1a, 0xbb, 0x89,
	0xd8, 0x3a, 0x58, 0x19, 0x2c, 0x92, 0x32, 0xc4, 0x1e, 0xa9, 0xe0, 0x44, 0x09, 0xd3, 0x4a, 0xca,
	0x19, 0xdb, 0x9b, 0x18, 0x34, 0xf4, 0x3e, 0x42, 0xef, 0x40, 0xb5, 0x00, 0x0a, 0x92, 0x2c, 0xec,
	0xef, 0x1d, 0xc2, 0x6a, 0xbe, 0xca, 0xfe, 0xde, 0x61, 0xc2, 0x58, 0x9b, 0xa2, 0x6a, 0x00, 0x45,
	0xc0, 0x43, 0xd8, 0xca, 0xed, 0x60, 0xfa, 0xe0, 0x8c, 0x94, 0x38, 0x66, 0x1c, 0xdb, 0x3b, 0x27,
	0x65, 0x3c, 0xde, 0xb0, 0x59, 0x78, 0xf4, 0x73, 0xf7, 0xd8, 0xdc, 0x9a, 0xeb, 0xd3, 0x51, 0x36,
	0x30, 0x4a, 0x0d, 0xd6, 0x32, 0x51, 0xe2, 0xeb, 0x43, 0x47, 0xfa, 0xf2, 0x9f, 0xc1, 0x29, 0xa9,
	0xc4, 0x97, 0x12, 0xe6, 0x55, 0xbe, 0x5a, 0xe7, 0xed, 0xf9, 0xc6, 0x99, 0xeb, 0x1d, 0x67, 0xd8,
	0x7d, 0x73, 0x31, 0xb6, 0x8c, 0xcb, 0xb1, 0x65, 0xfc, 0x19, 0x5b, 0xc6, 0xb7, 0x89, 0x55, 0xba,
	0x9c, 0x58, 0xa5, 0x5f, 0x13, 0xab, 0xf4, 0xc1, 0xee, 0x7a, 0xaa, 0x17, 0x1e, 0xd9, 0x6d, 0x31,
	0xa0, 0x6d, 0x16, 0x28, 0xef, 0xb8, 0x23, 0x42, 0xee, 0xe2, 0xe3, 0x48, 0xf5, 0xdb, 0x77, 0x8a,
	0xa5, 0xd4, 0xd0, 0x67, 0xf2, 0xa8, 0x82, 0xef, 0xde, 0xb3, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x11, 0xab, 0x3b, 0xc8, 0x9b, 0x05, 0x00, 0x00,
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
	// Supply queries the total supply of a given denom or owner
	Supply(ctx context.Context, in *types.QuerySupplyRequest, opts ...grpc.CallOption) (*types.QuerySupplyResponse, error)
	// Owner queries the NFTs of the specified owner
	Owner(ctx context.Context, in *types.QueryOwnerRequest, opts ...grpc.CallOption) (*types.QueryOwnerResponse, error)
	// Collection queries the NFTs of the specified denom
	Collection(ctx context.Context, in *types.QueryCollectionRequest, opts ...grpc.CallOption) (*types.QueryCollectionResponse, error)
	// Denom queries the definition of a given denom
	Denom(ctx context.Context, in *types.QueryDenomRequest, opts ...grpc.CallOption) (*types.QueryDenomResponse, error)
	// Denoms queries all the denoms
	Denoms(ctx context.Context, in *types.QueryDenomsRequest, opts ...grpc.CallOption) (*types.QueryDenomsResponse, error)
	// NFT queries the NFT for the given denom and token ID
	NFT(ctx context.Context, in *types.QueryNFTRequest, opts ...grpc.CallOption) (*types.QueryNFTResponse, error)
	Admin(ctx context.Context, in *QueryAdminRequest, opts ...grpc.CallOption) (*QueryAdminResponse, error)
	Admins(ctx context.Context, in *QueryAdminsRequest, opts ...grpc.CallOption) (*QueryAdminsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Supply(ctx context.Context, in *types.QuerySupplyRequest, opts ...grpc.CallOption) (*types.QuerySupplyResponse, error) {
	out := new(types.QuerySupplyResponse)
	err := c.cc.Invoke(ctx, "/shentu.nft.v1alpha1.Query/Supply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Owner(ctx context.Context, in *types.QueryOwnerRequest, opts ...grpc.CallOption) (*types.QueryOwnerResponse, error) {
	out := new(types.QueryOwnerResponse)
	err := c.cc.Invoke(ctx, "/shentu.nft.v1alpha1.Query/Owner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Collection(ctx context.Context, in *types.QueryCollectionRequest, opts ...grpc.CallOption) (*types.QueryCollectionResponse, error) {
	out := new(types.QueryCollectionResponse)
	err := c.cc.Invoke(ctx, "/shentu.nft.v1alpha1.Query/Collection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Denom(ctx context.Context, in *types.QueryDenomRequest, opts ...grpc.CallOption) (*types.QueryDenomResponse, error) {
	out := new(types.QueryDenomResponse)
	err := c.cc.Invoke(ctx, "/shentu.nft.v1alpha1.Query/Denom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Denoms(ctx context.Context, in *types.QueryDenomsRequest, opts ...grpc.CallOption) (*types.QueryDenomsResponse, error) {
	out := new(types.QueryDenomsResponse)
	err := c.cc.Invoke(ctx, "/shentu.nft.v1alpha1.Query/Denoms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) NFT(ctx context.Context, in *types.QueryNFTRequest, opts ...grpc.CallOption) (*types.QueryNFTResponse, error) {
	out := new(types.QueryNFTResponse)
	err := c.cc.Invoke(ctx, "/shentu.nft.v1alpha1.Query/NFT", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Admin(ctx context.Context, in *QueryAdminRequest, opts ...grpc.CallOption) (*QueryAdminResponse, error) {
	out := new(QueryAdminResponse)
	err := c.cc.Invoke(ctx, "/shentu.nft.v1alpha1.Query/Admin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) Admins(ctx context.Context, in *QueryAdminsRequest, opts ...grpc.CallOption) (*QueryAdminsResponse, error) {
	out := new(QueryAdminsResponse)
	err := c.cc.Invoke(ctx, "/shentu.nft.v1alpha1.Query/Admins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Supply queries the total supply of a given denom or owner
	Supply(context.Context, *types.QuerySupplyRequest) (*types.QuerySupplyResponse, error)
	// Owner queries the NFTs of the specified owner
	Owner(context.Context, *types.QueryOwnerRequest) (*types.QueryOwnerResponse, error)
	// Collection queries the NFTs of the specified denom
	Collection(context.Context, *types.QueryCollectionRequest) (*types.QueryCollectionResponse, error)
	// Denom queries the definition of a given denom
	Denom(context.Context, *types.QueryDenomRequest) (*types.QueryDenomResponse, error)
	// Denoms queries all the denoms
	Denoms(context.Context, *types.QueryDenomsRequest) (*types.QueryDenomsResponse, error)
	// NFT queries the NFT for the given denom and token ID
	NFT(context.Context, *types.QueryNFTRequest) (*types.QueryNFTResponse, error)
	Admin(context.Context, *QueryAdminRequest) (*QueryAdminResponse, error)
	Admins(context.Context, *QueryAdminsRequest) (*QueryAdminsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Supply(ctx context.Context, req *types.QuerySupplyRequest) (*types.QuerySupplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Supply not implemented")
}
func (*UnimplementedQueryServer) Owner(ctx context.Context, req *types.QueryOwnerRequest) (*types.QueryOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Owner not implemented")
}
func (*UnimplementedQueryServer) Collection(ctx context.Context, req *types.QueryCollectionRequest) (*types.QueryCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Collection not implemented")
}
func (*UnimplementedQueryServer) Denom(ctx context.Context, req *types.QueryDenomRequest) (*types.QueryDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Denom not implemented")
}
func (*UnimplementedQueryServer) Denoms(ctx context.Context, req *types.QueryDenomsRequest) (*types.QueryDenomsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Denoms not implemented")
}
func (*UnimplementedQueryServer) NFT(ctx context.Context, req *types.QueryNFTRequest) (*types.QueryNFTResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NFT not implemented")
}
func (*UnimplementedQueryServer) Admin(ctx context.Context, req *QueryAdminRequest) (*QueryAdminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Admin not implemented")
}
func (*UnimplementedQueryServer) Admins(ctx context.Context, req *QueryAdminsRequest) (*QueryAdminsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Admins not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Supply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.QuerySupplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Supply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.nft.v1alpha1.Query/Supply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Supply(ctx, req.(*types.QuerySupplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Owner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.QueryOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Owner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.nft.v1alpha1.Query/Owner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Owner(ctx, req.(*types.QueryOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Collection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.QueryCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Collection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.nft.v1alpha1.Query/Collection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Collection(ctx, req.(*types.QueryCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Denom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.QueryDenomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Denom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.nft.v1alpha1.Query/Denom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Denom(ctx, req.(*types.QueryDenomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Denoms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.QueryDenomsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Denoms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.nft.v1alpha1.Query/Denoms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Denoms(ctx, req.(*types.QueryDenomsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_NFT_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.QueryNFTRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).NFT(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.nft.v1alpha1.Query/NFT",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).NFT(ctx, req.(*types.QueryNFTRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Admin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Admin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.nft.v1alpha1.Query/Admin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Admin(ctx, req.(*QueryAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_Admins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAdminsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Admins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/shentu.nft.v1alpha1.Query/Admins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Admins(ctx, req.(*QueryAdminsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "shentu.nft.v1alpha1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Supply",
			Handler:    _Query_Supply_Handler,
		},
		{
			MethodName: "Owner",
			Handler:    _Query_Owner_Handler,
		},
		{
			MethodName: "Collection",
			Handler:    _Query_Collection_Handler,
		},
		{
			MethodName: "Denom",
			Handler:    _Query_Denom_Handler,
		},
		{
			MethodName: "Denoms",
			Handler:    _Query_Denoms_Handler,
		},
		{
			MethodName: "NFT",
			Handler:    _Query_NFT_Handler,
		},
		{
			MethodName: "Admin",
			Handler:    _Query_Admin_Handler,
		},
		{
			MethodName: "Admins",
			Handler:    _Query_Admins_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shentu/nft/v1alpha1/query.proto",
}

func (m *QueryAdminRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAdminRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAdminRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAdminResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAdminResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAdminResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Admin.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryAdminsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAdminsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAdminsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryAdminsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAdminsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAdminsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Admins) > 0 {
		for iNdEx := len(m.Admins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Admins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryAdminRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAdminResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Admin.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAdminsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryAdminsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Admins) > 0 {
		for _, e := range m.Admins {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryAdminRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAdminRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAdminRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryAdminResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAdminResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAdminResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
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
			if err := m.Admin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryAdminsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAdminsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAdminsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
func (m *QueryAdminsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryAdminsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAdminsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admins", wireType)
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
			m.Admins = append(m.Admins, Admin{})
			if err := m.Admins[len(m.Admins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
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
