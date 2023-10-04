// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bech32ibc/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	_ "google.golang.org/protobuf/types/known/durationpb"
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

type QueryHrpIbcRecordsRequest struct {
}

func (m *QueryHrpIbcRecordsRequest) Reset()         { *m = QueryHrpIbcRecordsRequest{} }
func (m *QueryHrpIbcRecordsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryHrpIbcRecordsRequest) ProtoMessage()    {}
func (*QueryHrpIbcRecordsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_004dff02d981ce19, []int{0}
}
func (m *QueryHrpIbcRecordsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryHrpIbcRecordsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryHrpIbcRecordsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHrpIbcRecordsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHrpIbcRecordsRequest.Merge(m, src)
}
func (m *QueryHrpIbcRecordsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryHrpIbcRecordsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHrpIbcRecordsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHrpIbcRecordsRequest proto.InternalMessageInfo

type QueryHrpIbcRecordsResponse struct {
	HrpIbcRecords []HrpIbcRecord `protobuf:"bytes,1,rep,name=hrp_ibc_records,json=hrpIbcRecords,proto3" json:"hrp_ibc_records" yaml:"hrp_ibc_records"`
}

func (m *QueryHrpIbcRecordsResponse) Reset()         { *m = QueryHrpIbcRecordsResponse{} }
func (m *QueryHrpIbcRecordsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryHrpIbcRecordsResponse) ProtoMessage()    {}
func (*QueryHrpIbcRecordsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_004dff02d981ce19, []int{1}
}
func (m *QueryHrpIbcRecordsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryHrpIbcRecordsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryHrpIbcRecordsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHrpIbcRecordsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHrpIbcRecordsResponse.Merge(m, src)
}
func (m *QueryHrpIbcRecordsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryHrpIbcRecordsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHrpIbcRecordsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHrpIbcRecordsResponse proto.InternalMessageInfo

func (m *QueryHrpIbcRecordsResponse) GetHrpIbcRecords() []HrpIbcRecord {
	if m != nil {
		return m.HrpIbcRecords
	}
	return nil
}

type QueryHrpIbcRecordRequest struct {
	Hrp string `protobuf:"bytes,1,opt,name=hrp,proto3" json:"hrp,omitempty" yaml:"hrp"`
}

func (m *QueryHrpIbcRecordRequest) Reset()         { *m = QueryHrpIbcRecordRequest{} }
func (m *QueryHrpIbcRecordRequest) String() string { return proto.CompactTextString(m) }
func (*QueryHrpIbcRecordRequest) ProtoMessage()    {}
func (*QueryHrpIbcRecordRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_004dff02d981ce19, []int{2}
}
func (m *QueryHrpIbcRecordRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryHrpIbcRecordRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryHrpIbcRecordRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHrpIbcRecordRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHrpIbcRecordRequest.Merge(m, src)
}
func (m *QueryHrpIbcRecordRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryHrpIbcRecordRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHrpIbcRecordRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHrpIbcRecordRequest proto.InternalMessageInfo

func (m *QueryHrpIbcRecordRequest) GetHrp() string {
	if m != nil {
		return m.Hrp
	}
	return ""
}

type QueryHrpIbcRecordResponse struct {
	HrpIbcRecord HrpIbcRecord `protobuf:"bytes,1,opt,name=hrp_ibc_record,json=hrpIbcRecord,proto3" json:"hrp_ibc_record" yaml:"hrp_ibc_record"`
}

func (m *QueryHrpIbcRecordResponse) Reset()         { *m = QueryHrpIbcRecordResponse{} }
func (m *QueryHrpIbcRecordResponse) String() string { return proto.CompactTextString(m) }
func (*QueryHrpIbcRecordResponse) ProtoMessage()    {}
func (*QueryHrpIbcRecordResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_004dff02d981ce19, []int{3}
}
func (m *QueryHrpIbcRecordResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryHrpIbcRecordResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryHrpIbcRecordResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryHrpIbcRecordResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryHrpIbcRecordResponse.Merge(m, src)
}
func (m *QueryHrpIbcRecordResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryHrpIbcRecordResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryHrpIbcRecordResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryHrpIbcRecordResponse proto.InternalMessageInfo

func (m *QueryHrpIbcRecordResponse) GetHrpIbcRecord() HrpIbcRecord {
	if m != nil {
		return m.HrpIbcRecord
	}
	return HrpIbcRecord{}
}

type QueryNativeHrpRequest struct {
}

func (m *QueryNativeHrpRequest) Reset()         { *m = QueryNativeHrpRequest{} }
func (m *QueryNativeHrpRequest) String() string { return proto.CompactTextString(m) }
func (*QueryNativeHrpRequest) ProtoMessage()    {}
func (*QueryNativeHrpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_004dff02d981ce19, []int{4}
}
func (m *QueryNativeHrpRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryNativeHrpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryNativeHrpRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryNativeHrpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryNativeHrpRequest.Merge(m, src)
}
func (m *QueryNativeHrpRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryNativeHrpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryNativeHrpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryNativeHrpRequest proto.InternalMessageInfo

type QueryNativeHrpResponse struct {
	NativeHrp string `protobuf:"bytes,1,opt,name=native_hrp,json=nativeHrp,proto3" json:"native_hrp,omitempty" yaml:"native_hrp"`
}

func (m *QueryNativeHrpResponse) Reset()         { *m = QueryNativeHrpResponse{} }
func (m *QueryNativeHrpResponse) String() string { return proto.CompactTextString(m) }
func (*QueryNativeHrpResponse) ProtoMessage()    {}
func (*QueryNativeHrpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_004dff02d981ce19, []int{5}
}
func (m *QueryNativeHrpResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryNativeHrpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryNativeHrpResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryNativeHrpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryNativeHrpResponse.Merge(m, src)
}
func (m *QueryNativeHrpResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryNativeHrpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryNativeHrpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryNativeHrpResponse proto.InternalMessageInfo

func (m *QueryNativeHrpResponse) GetNativeHrp() string {
	if m != nil {
		return m.NativeHrp
	}
	return ""
}

func init() {
	proto.RegisterType((*QueryHrpIbcRecordsRequest)(nil), "bech32ibc.bech32ibc.v1beta1.QueryHrpIbcRecordsRequest")
	proto.RegisterType((*QueryHrpIbcRecordsResponse)(nil), "bech32ibc.bech32ibc.v1beta1.QueryHrpIbcRecordsResponse")
	proto.RegisterType((*QueryHrpIbcRecordRequest)(nil), "bech32ibc.bech32ibc.v1beta1.QueryHrpIbcRecordRequest")
	proto.RegisterType((*QueryHrpIbcRecordResponse)(nil), "bech32ibc.bech32ibc.v1beta1.QueryHrpIbcRecordResponse")
	proto.RegisterType((*QueryNativeHrpRequest)(nil), "bech32ibc.bech32ibc.v1beta1.QueryNativeHrpRequest")
	proto.RegisterType((*QueryNativeHrpResponse)(nil), "bech32ibc.bech32ibc.v1beta1.QueryNativeHrpResponse")
}

func init() { proto.RegisterFile("bech32ibc/v1beta1/query.proto", fileDescriptor_004dff02d981ce19) }

var fileDescriptor_004dff02d981ce19 = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xc1, 0x6a, 0x13, 0x41,
	0x18, 0xc7, 0x33, 0x56, 0x85, 0x8c, 0x6d, 0xc5, 0xc1, 0xd4, 0x98, 0xda, 0x4d, 0x18, 0x04, 0x5b,
	0xa1, 0x19, 0xbb, 0xd1, 0x2a, 0xe2, 0x29, 0xa7, 0x0a, 0x52, 0x70, 0x8f, 0x5e, 0xc2, 0xee, 0x76,
	0x9a, 0x5d, 0x48, 0x66, 0x36, 0x33, 0xb3, 0xa5, 0x41, 0xbc, 0x78, 0xf5, 0xa0, 0xe0, 0xc5, 0xa3,
	0xaf, 0xe0, 0xc9, 0x57, 0xe8, 0xb1, 0xe0, 0xc5, 0x53, 0x90, 0xc4, 0x27, 0xc8, 0x13, 0x48, 0x66,
	0x27, 0xbb, 0x69, 0x4c, 0x53, 0xf6, 0x36, 0xc9, 0x7f, 0xbe, 0xff, 0xf7, 0xfb, 0xef, 0x7c, 0x7c,
	0x70, 0xcb, 0xa3, 0x7e, 0xd0, 0xb0, 0x43, 0xcf, 0x27, 0x27, 0x7b, 0x1e, 0x55, 0xee, 0x1e, 0xe9,
	0xc5, 0x54, 0xf4, 0xeb, 0x91, 0xe0, 0x8a, 0xa3, 0xcd, 0x54, 0xae, 0x67, 0x27, 0x73, 0xb1, 0x72,
	0xb7, 0xcd, 0xdb, 0x5c, 0xdf, 0x23, 0x93, 0x53, 0x52, 0x52, 0x79, 0xd0, 0xe6, 0xbc, 0xdd, 0xa1,
	0xc4, 0x8d, 0x42, 0xe2, 0x32, 0xc6, 0x95, 0xab, 0x42, 0xce, 0xa4, 0x51, 0x2d, 0xa3, 0xea, 0x5f,
	0x5e, 0x7c, 0x4c, 0x8e, 0x62, 0xa1, 0x2f, 0x18, 0x7d, 0x01, 0x8f, 0xea, 0x47, 0xd4, 0x94, 0xe3,
	0x4d, 0x78, 0xff, 0xed, 0x04, 0xef, 0x40, 0x44, 0xaf, 0x3d, 0xdf, 0xa1, 0x3e, 0x17, 0x47, 0xd2,
	0xa1, 0xbd, 0x98, 0x4a, 0x85, 0x3f, 0x03, 0x58, 0x59, 0xa4, 0xca, 0x88, 0x33, 0x49, 0x51, 0x0f,
	0xde, 0x0e, 0x44, 0xd4, 0x0a, 0x3d, 0xbf, 0x25, 0x12, 0xa9, 0x0c, 0x6a, 0x2b, 0xdb, 0xb7, 0xec,
	0x9d, 0xfa, 0x92, 0x94, 0xf5, 0x59, 0xb3, 0xa6, 0x75, 0x36, 0xa8, 0x16, 0xc6, 0x83, 0xea, 0x46,
	0xdf, 0xed, 0x76, 0x5e, 0xe2, 0x39, 0x3f, 0xec, 0xac, 0x05, 0xb3, 0xad, 0xf1, 0x2b, 0x58, 0xfe,
	0x0f, 0xc8, 0xd0, 0xa2, 0x1a, 0x5c, 0x09, 0x44, 0x54, 0x06, 0x35, 0xb0, 0x5d, 0x6c, 0xae, 0x8f,
	0x07, 0x55, 0x98, 0x7a, 0x62, 0x67, 0x22, 0xe1, 0x4f, 0x60, 0x41, 0xda, 0x34, 0x0e, 0x83, 0xeb,
	0x17, 0xdb, 0x6b, 0xab, 0x5c, 0x69, 0xb6, 0x4c, 0x9a, 0xd2, 0xa2, 0x34, 0xd8, 0x59, 0x9d, 0x0d,
	0x83, 0xef, 0xc1, 0x92, 0x86, 0x39, 0x74, 0x55, 0x78, 0x42, 0x0f, 0x44, 0x34, 0xfd, 0xec, 0x87,
	0x70, 0x63, 0x5e, 0x30, 0x88, 0x4f, 0x21, 0x64, 0xfa, 0xcf, 0x56, 0x96, 0xb4, 0x34, 0x1e, 0x54,
	0xef, 0x24, 0xfd, 0x32, 0x0d, 0x3b, 0x45, 0x36, 0xad, 0xb6, 0xbf, 0x5d, 0x87, 0x37, 0xb4, 0x21,
	0xfa, 0x01, 0xe0, 0xda, 0x85, 0xb7, 0x44, 0xfb, 0x4b, 0xc3, 0x5d, 0x3a, 0x1a, 0x95, 0xe7, 0xb9,
	0xeb, 0x92, 0x08, 0x98, 0x7c, 0xfc, 0xf5, 0xf7, 0xeb, 0xb5, 0x1d, 0xf4, 0x88, 0x70, 0xd9, 0xe5,
	0x32, 0x94, 0x44, 0x9d, 0x1e, 0x53, 0x2a, 0xd3, 0xe9, 0x9c, 0x1b, 0x01, 0xf4, 0x13, 0xc0, 0xd5,
	0x59, 0x2b, 0xf4, 0x2c, 0x5f, 0xeb, 0x29, 0xf1, 0x7e, 0xde, 0x32, 0x03, 0xfc, 0x42, 0x03, 0xdb,
	0xe8, 0xc9, 0x32, 0x60, 0xc9, 0x63, 0xe1, 0xd3, 0x96, 0x1f, 0xb8, 0x8c, 0xd1, 0x0e, 0x79, 0x1f,
	0x88, 0xe8, 0x03, 0xfa, 0x0e, 0x60, 0x31, 0x7d, 0x43, 0x64, 0x5f, 0xdd, 0x7f, 0x7e, 0x12, 0x2a,
	0x8d, 0x5c, 0x35, 0x06, 0xf8, 0xb1, 0x06, 0x7e, 0x88, 0xf0, 0x65, 0xc0, 0xd9, 0x98, 0x34, 0xdf,
	0x9c, 0x0d, 0x2d, 0x70, 0x3e, 0xb4, 0xc0, 0x9f, 0xa1, 0x05, 0xbe, 0x8c, 0xac, 0xc2, 0xf9, 0xc8,
	0x2a, 0xfc, 0x1e, 0x59, 0x85, 0x77, 0x76, 0x3b, 0x54, 0x41, 0xec, 0xd5, 0x7d, 0xde, 0x25, 0x6e,
	0x47, 0x05, 0xd4, 0xdd, 0x65, 0x54, 0x91, 0x84, 0x62, 0x77, 0xb2, 0x4e, 0x4e, 0x49, 0xb6, 0x5a,
	0xf4, 0x4a, 0xf1, 0x6e, 0xea, 0x9d, 0xd2, 0xf8, 0x17, 0x00, 0x00, 0xff, 0xff, 0xe1, 0xa9, 0x1c,
	0x36, 0x04, 0x05, 0x00, 0x00,
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
	// HrpIbcRecords returns to full list of records
	HrpIbcRecords(ctx context.Context, in *QueryHrpIbcRecordsRequest, opts ...grpc.CallOption) (*QueryHrpIbcRecordsResponse, error)
	// HrpIbcRecord returns the record for a requested HRP
	HrpIbcRecord(ctx context.Context, in *QueryHrpIbcRecordRequest, opts ...grpc.CallOption) (*QueryHrpIbcRecordResponse, error)
	// NativeHrp returns the chain's native HRP
	NativeHrp(ctx context.Context, in *QueryNativeHrpRequest, opts ...grpc.CallOption) (*QueryNativeHrpResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) HrpIbcRecords(ctx context.Context, in *QueryHrpIbcRecordsRequest, opts ...grpc.CallOption) (*QueryHrpIbcRecordsResponse, error) {
	out := new(QueryHrpIbcRecordsResponse)
	err := c.cc.Invoke(ctx, "/bech32ibc.bech32ibc.v1beta1.Query/HrpIbcRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) HrpIbcRecord(ctx context.Context, in *QueryHrpIbcRecordRequest, opts ...grpc.CallOption) (*QueryHrpIbcRecordResponse, error) {
	out := new(QueryHrpIbcRecordResponse)
	err := c.cc.Invoke(ctx, "/bech32ibc.bech32ibc.v1beta1.Query/HrpIbcRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) NativeHrp(ctx context.Context, in *QueryNativeHrpRequest, opts ...grpc.CallOption) (*QueryNativeHrpResponse, error) {
	out := new(QueryNativeHrpResponse)
	err := c.cc.Invoke(ctx, "/bech32ibc.bech32ibc.v1beta1.Query/NativeHrp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// HrpIbcRecords returns to full list of records
	HrpIbcRecords(context.Context, *QueryHrpIbcRecordsRequest) (*QueryHrpIbcRecordsResponse, error)
	// HrpIbcRecord returns the record for a requested HRP
	HrpIbcRecord(context.Context, *QueryHrpIbcRecordRequest) (*QueryHrpIbcRecordResponse, error)
	// NativeHrp returns the chain's native HRP
	NativeHrp(context.Context, *QueryNativeHrpRequest) (*QueryNativeHrpResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) HrpIbcRecords(ctx context.Context, req *QueryHrpIbcRecordsRequest) (*QueryHrpIbcRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HrpIbcRecords not implemented")
}
func (*UnimplementedQueryServer) HrpIbcRecord(ctx context.Context, req *QueryHrpIbcRecordRequest) (*QueryHrpIbcRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HrpIbcRecord not implemented")
}
func (*UnimplementedQueryServer) NativeHrp(ctx context.Context, req *QueryNativeHrpRequest) (*QueryNativeHrpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NativeHrp not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_HrpIbcRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHrpIbcRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HrpIbcRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bech32ibc.bech32ibc.v1beta1.Query/HrpIbcRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HrpIbcRecords(ctx, req.(*QueryHrpIbcRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_HrpIbcRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryHrpIbcRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).HrpIbcRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bech32ibc.bech32ibc.v1beta1.Query/HrpIbcRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).HrpIbcRecord(ctx, req.(*QueryHrpIbcRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_NativeHrp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryNativeHrpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).NativeHrp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bech32ibc.bech32ibc.v1beta1.Query/NativeHrp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).NativeHrp(ctx, req.(*QueryNativeHrpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bech32ibc.bech32ibc.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HrpIbcRecords",
			Handler:    _Query_HrpIbcRecords_Handler,
		},
		{
			MethodName: "HrpIbcRecord",
			Handler:    _Query_HrpIbcRecord_Handler,
		},
		{
			MethodName: "NativeHrp",
			Handler:    _Query_NativeHrp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bech32ibc/v1beta1/query.proto",
}

func (m *QueryHrpIbcRecordsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHrpIbcRecordsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryHrpIbcRecordsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryHrpIbcRecordsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHrpIbcRecordsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryHrpIbcRecordsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.HrpIbcRecords) > 0 {
		for iNdEx := len(m.HrpIbcRecords) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.HrpIbcRecords[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryHrpIbcRecordRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHrpIbcRecordRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryHrpIbcRecordRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hrp) > 0 {
		i -= len(m.Hrp)
		copy(dAtA[i:], m.Hrp)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Hrp)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryHrpIbcRecordResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryHrpIbcRecordResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryHrpIbcRecordResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.HrpIbcRecord.MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryNativeHrpRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryNativeHrpRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryNativeHrpRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryNativeHrpResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryNativeHrpResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryNativeHrpResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NativeHrp) > 0 {
		i -= len(m.NativeHrp)
		copy(dAtA[i:], m.NativeHrp)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.NativeHrp)))
		i--
		dAtA[i] = 0xa
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
func (m *QueryHrpIbcRecordsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryHrpIbcRecordsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.HrpIbcRecords) > 0 {
		for _, e := range m.HrpIbcRecords {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryHrpIbcRecordRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hrp)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryHrpIbcRecordResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.HrpIbcRecord.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryNativeHrpRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryNativeHrpResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.NativeHrp)
	if l > 0 {
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
func (m *QueryHrpIbcRecordsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryHrpIbcRecordsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryHrpIbcRecordsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryHrpIbcRecordsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryHrpIbcRecordsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryHrpIbcRecordsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HrpIbcRecords", wireType)
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
			m.HrpIbcRecords = append(m.HrpIbcRecords, HrpIbcRecord{})
			if err := m.HrpIbcRecords[len(m.HrpIbcRecords)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryHrpIbcRecordRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryHrpIbcRecordRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryHrpIbcRecordRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hrp", wireType)
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
			m.Hrp = string(dAtA[iNdEx:postIndex])
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
func (m *QueryHrpIbcRecordResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryHrpIbcRecordResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryHrpIbcRecordResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HrpIbcRecord", wireType)
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
			if err := m.HrpIbcRecord.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryNativeHrpRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryNativeHrpRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryNativeHrpRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryNativeHrpResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryNativeHrpResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryNativeHrpResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NativeHrp", wireType)
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
			m.NativeHrp = string(dAtA[iNdEx:postIndex])
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
