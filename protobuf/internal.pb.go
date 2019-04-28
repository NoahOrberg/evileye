// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal.proto

package evileye

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SuccessHashCalcRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nonce                string   `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SuccessHashCalcRequest) Reset()         { *m = SuccessHashCalcRequest{} }
func (m *SuccessHashCalcRequest) String() string { return proto.CompactTextString(m) }
func (*SuccessHashCalcRequest) ProtoMessage()    {}
func (*SuccessHashCalcRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_internal_2988fb781b00cf03, []int{0}
}
func (m *SuccessHashCalcRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SuccessHashCalcRequest.Unmarshal(m, b)
}
func (m *SuccessHashCalcRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SuccessHashCalcRequest.Marshal(b, m, deterministic)
}
func (dst *SuccessHashCalcRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SuccessHashCalcRequest.Merge(dst, src)
}
func (m *SuccessHashCalcRequest) XXX_Size() int {
	return xxx_messageInfo_SuccessHashCalcRequest.Size(m)
}
func (m *SuccessHashCalcRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SuccessHashCalcRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SuccessHashCalcRequest proto.InternalMessageInfo

func (m *SuccessHashCalcRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SuccessHashCalcRequest) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

type SendCheckResultRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Nonce                string   `protobuf:"bytes,2,opt,name=nonce,proto3" json:"nonce,omitempty"`
	IsOk                 bool     `protobuf:"varint,3,opt,name=is_ok,json=isOk,proto3" json:"is_ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendCheckResultRequest) Reset()         { *m = SendCheckResultRequest{} }
func (m *SendCheckResultRequest) String() string { return proto.CompactTextString(m) }
func (*SendCheckResultRequest) ProtoMessage()    {}
func (*SendCheckResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_internal_2988fb781b00cf03, []int{1}
}
func (m *SendCheckResultRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendCheckResultRequest.Unmarshal(m, b)
}
func (m *SendCheckResultRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendCheckResultRequest.Marshal(b, m, deterministic)
}
func (dst *SendCheckResultRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendCheckResultRequest.Merge(dst, src)
}
func (m *SendCheckResultRequest) XXX_Size() int {
	return xxx_messageInfo_SendCheckResultRequest.Size(m)
}
func (m *SendCheckResultRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendCheckResultRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendCheckResultRequest proto.InternalMessageInfo

func (m *SendCheckResultRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SendCheckResultRequest) GetNonce() string {
	if m != nil {
		return m.Nonce
	}
	return ""
}

func (m *SendCheckResultRequest) GetIsOk() bool {
	if m != nil {
		return m.IsOk
	}
	return false
}

type Tarekomis struct {
	Tarekomis            []*Tarekomi `protobuf:"bytes,1,rep,name=tarekomis,proto3" json:"tarekomis,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Tarekomis) Reset()         { *m = Tarekomis{} }
func (m *Tarekomis) String() string { return proto.CompactTextString(m) }
func (*Tarekomis) ProtoMessage()    {}
func (*Tarekomis) Descriptor() ([]byte, []int) {
	return fileDescriptor_internal_2988fb781b00cf03, []int{2}
}
func (m *Tarekomis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tarekomis.Unmarshal(m, b)
}
func (m *Tarekomis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tarekomis.Marshal(b, m, deterministic)
}
func (dst *Tarekomis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tarekomis.Merge(dst, src)
}
func (m *Tarekomis) XXX_Size() int {
	return xxx_messageInfo_Tarekomis.Size(m)
}
func (m *Tarekomis) XXX_DiscardUnknown() {
	xxx_messageInfo_Tarekomis.DiscardUnknown(m)
}

var xxx_messageInfo_Tarekomis proto.InternalMessageInfo

func (m *Tarekomis) GetTarekomis() []*Tarekomi {
	if m != nil {
		return m.Tarekomis
	}
	return nil
}

type SendTxRequest struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Desc                 string   `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	ApprovedUserNames    []string `protobuf:"bytes,4,rep,name=approved_user_names,json=approvedUserNames,proto3" json:"approved_user_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendTxRequest) Reset()         { *m = SendTxRequest{} }
func (m *SendTxRequest) String() string { return proto.CompactTextString(m) }
func (*SendTxRequest) ProtoMessage()    {}
func (*SendTxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_internal_2988fb781b00cf03, []int{3}
}
func (m *SendTxRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendTxRequest.Unmarshal(m, b)
}
func (m *SendTxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendTxRequest.Marshal(b, m, deterministic)
}
func (dst *SendTxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendTxRequest.Merge(dst, src)
}
func (m *SendTxRequest) XXX_Size() int {
	return xxx_messageInfo_SendTxRequest.Size(m)
}
func (m *SendTxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendTxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendTxRequest proto.InternalMessageInfo

func (m *SendTxRequest) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *SendTxRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *SendTxRequest) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *SendTxRequest) GetApprovedUserNames() []string {
	if m != nil {
		return m.ApprovedUserNames
	}
	return nil
}

type Tx struct {
	UserName             string   `protobuf:"bytes,1,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Desc                 string   `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	ApprovedUserNames    []string `protobuf:"bytes,4,rep,name=approved_user_names,json=approvedUserNames,proto3" json:"approved_user_names,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tx) Reset()         { *m = Tx{} }
func (m *Tx) String() string { return proto.CompactTextString(m) }
func (*Tx) ProtoMessage()    {}
func (*Tx) Descriptor() ([]byte, []int) {
	return fileDescriptor_internal_2988fb781b00cf03, []int{4}
}
func (m *Tx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tx.Unmarshal(m, b)
}
func (m *Tx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tx.Marshal(b, m, deterministic)
}
func (dst *Tx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tx.Merge(dst, src)
}
func (m *Tx) XXX_Size() int {
	return xxx_messageInfo_Tx.Size(m)
}
func (m *Tx) XXX_DiscardUnknown() {
	xxx_messageInfo_Tx.DiscardUnknown(m)
}

var xxx_messageInfo_Tx proto.InternalMessageInfo

func (m *Tx) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Tx) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Tx) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func (m *Tx) GetApprovedUserNames() []string {
	if m != nil {
		return m.ApprovedUserNames
	}
	return nil
}

type Txs struct {
	Txs                  []*Tx    `protobuf:"bytes,1,rep,name=txs,proto3" json:"txs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Txs) Reset()         { *m = Txs{} }
func (m *Txs) String() string { return proto.CompactTextString(m) }
func (*Txs) ProtoMessage()    {}
func (*Txs) Descriptor() ([]byte, []int) {
	return fileDescriptor_internal_2988fb781b00cf03, []int{5}
}
func (m *Txs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Txs.Unmarshal(m, b)
}
func (m *Txs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Txs.Marshal(b, m, deterministic)
}
func (dst *Txs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Txs.Merge(dst, src)
}
func (m *Txs) XXX_Size() int {
	return xxx_messageInfo_Txs.Size(m)
}
func (m *Txs) XXX_DiscardUnknown() {
	xxx_messageInfo_Txs.DiscardUnknown(m)
}

var xxx_messageInfo_Txs proto.InternalMessageInfo

func (m *Txs) GetTxs() []*Tx {
	if m != nil {
		return m.Txs
	}
	return nil
}

type HealthCheckResponse struct {
	Hash                 string   `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckResponse) Reset()         { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()    {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_internal_2988fb781b00cf03, []int{6}
}
func (m *HealthCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckResponse.Unmarshal(m, b)
}
func (m *HealthCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckResponse.Marshal(b, m, deterministic)
}
func (dst *HealthCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckResponse.Merge(dst, src)
}
func (m *HealthCheckResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckResponse.Size(m)
}
func (m *HealthCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckResponse proto.InternalMessageInfo

func (m *HealthCheckResponse) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*SuccessHashCalcRequest)(nil), "evileye.SuccessHashCalcRequest")
	proto.RegisterType((*SendCheckResultRequest)(nil), "evileye.SendCheckResultRequest")
	proto.RegisterType((*Tarekomis)(nil), "evileye.Tarekomis")
	proto.RegisterType((*SendTxRequest)(nil), "evileye.SendTxRequest")
	proto.RegisterType((*Tx)(nil), "evileye.Tx")
	proto.RegisterType((*Txs)(nil), "evileye.Txs")
	proto.RegisterType((*HealthCheckResponse)(nil), "evileye.HealthCheckResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// InternalClient is the client API for Internal service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type InternalClient interface {
	HealthCheck(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	//
	// SuccessHashCalc
	//
	// ハッシュ計算が成功した際に他ノードに通知する
	// 成功者がid(UUIDv4)を振り、nonceとともに他ノードにBroadCastする
	SuccessHashCalc(ctx context.Context, in *SuccessHashCalcRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	//
	// SendCheckResult
	//
	// N-1のノードが計算をした結果をそれぞれ自分以外にBroadCastする
	// この際、idを付与することでそのidのOK, NotOKの数をそれぞれがカウントできるようにする
	// それぞれがOKのカウントがしきい値を超えたら、リーダーノードからTxPoolを取得して勝手にブロックを生成する
	SendCheckResult(ctx context.Context, in *SendCheckResultRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	//
	// GetTxPool
	//
	// ノードがブロック生成時にリーダーノード宛にリクエストを行う
	// ここには、Approvedになった黒歴史がリストで詰まっている
	//
	// リーダーノードはTxPoolの内容をすべて返す
	GetTxPool(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Txs, error)
	//
	// SendTx
	//
	// ApproveになったTarekomiを投げる. (リーダーノードのみ!)
	SendTx(ctx context.Context, in *SendTxRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type internalClient struct {
	cc *grpc.ClientConn
}

func NewInternalClient(cc *grpc.ClientConn) InternalClient {
	return &internalClient{cc}
}

func (c *internalClient) HealthCheck(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/evileye.Internal/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalClient) SuccessHashCalc(ctx context.Context, in *SuccessHashCalcRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/evileye.Internal/SuccessHashCalc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalClient) SendCheckResult(ctx context.Context, in *SendCheckResultRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/evileye.Internal/SendCheckResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalClient) GetTxPool(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*Txs, error) {
	out := new(Txs)
	err := c.cc.Invoke(ctx, "/evileye.Internal/GetTxPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalClient) SendTx(ctx context.Context, in *SendTxRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/evileye.Internal/SendTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// InternalServer is the server API for Internal service.
type InternalServer interface {
	HealthCheck(context.Context, *empty.Empty) (*HealthCheckResponse, error)
	//
	// SuccessHashCalc
	//
	// ハッシュ計算が成功した際に他ノードに通知する
	// 成功者がid(UUIDv4)を振り、nonceとともに他ノードにBroadCastする
	SuccessHashCalc(context.Context, *SuccessHashCalcRequest) (*empty.Empty, error)
	//
	// SendCheckResult
	//
	// N-1のノードが計算をした結果をそれぞれ自分以外にBroadCastする
	// この際、idを付与することでそのidのOK, NotOKの数をそれぞれがカウントできるようにする
	// それぞれがOKのカウントがしきい値を超えたら、リーダーノードからTxPoolを取得して勝手にブロックを生成する
	SendCheckResult(context.Context, *SendCheckResultRequest) (*empty.Empty, error)
	//
	// GetTxPool
	//
	// ノードがブロック生成時にリーダーノード宛にリクエストを行う
	// ここには、Approvedになった黒歴史がリストで詰まっている
	//
	// リーダーノードはTxPoolの内容をすべて返す
	GetTxPool(context.Context, *empty.Empty) (*Txs, error)
	//
	// SendTx
	//
	// ApproveになったTarekomiを投げる. (リーダーノードのみ!)
	SendTx(context.Context, *SendTxRequest) (*empty.Empty, error)
}

func RegisterInternalServer(s *grpc.Server, srv InternalServer) {
	s.RegisterService(&_Internal_serviceDesc, srv)
}

func _Internal_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evileye.Internal/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServer).HealthCheck(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Internal_SuccessHashCalc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuccessHashCalcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServer).SuccessHashCalc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evileye.Internal/SuccessHashCalc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServer).SuccessHashCalc(ctx, req.(*SuccessHashCalcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Internal_SendCheckResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendCheckResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServer).SendCheckResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evileye.Internal/SendCheckResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServer).SendCheckResult(ctx, req.(*SendCheckResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Internal_GetTxPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServer).GetTxPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evileye.Internal/GetTxPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServer).GetTxPool(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Internal_SendTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalServer).SendTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/evileye.Internal/SendTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalServer).SendTx(ctx, req.(*SendTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Internal_serviceDesc = grpc.ServiceDesc{
	ServiceName: "evileye.Internal",
	HandlerType: (*InternalServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _Internal_HealthCheck_Handler,
		},
		{
			MethodName: "SuccessHashCalc",
			Handler:    _Internal_SuccessHashCalc_Handler,
		},
		{
			MethodName: "SendCheckResult",
			Handler:    _Internal_SendCheckResult_Handler,
		},
		{
			MethodName: "GetTxPool",
			Handler:    _Internal_GetTxPool_Handler,
		},
		{
			MethodName: "SendTx",
			Handler:    _Internal_SendTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal.proto",
}

func init() { proto.RegisterFile("internal.proto", fileDescriptor_internal_2988fb781b00cf03) }

var fileDescriptor_internal_2988fb781b00cf03 = []byte{
	// 433 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x92, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x65, 0x3b, 0x2d, 0xf1, 0x04, 0x0a, 0xdd, 0xa0, 0xc8, 0x4a, 0x41, 0x44, 0x16, 0x87,
	0x70, 0x71, 0x44, 0xb9, 0x70, 0x40, 0x5c, 0x22, 0x44, 0xe1, 0x00, 0x68, 0x63, 0xce, 0xd1, 0xd6,
	0x1e, 0xea, 0x55, 0x36, 0x5e, 0xe3, 0x59, 0x17, 0xf7, 0xca, 0x33, 0xf2, 0x40, 0xc8, 0xff, 0xd2,
	0x14, 0xc5, 0x07, 0x4e, 0xbd, 0xcd, 0x8c, 0x67, 0xe6, 0xfb, 0x76, 0xfc, 0x83, 0x13, 0x99, 0x1a,
	0xcc, 0x53, 0xa1, 0x82, 0x2c, 0xd7, 0x46, 0xb3, 0x07, 0x78, 0x2d, 0x15, 0xde, 0xe0, 0xf4, 0xec,
	0x4a, 0xeb, 0x2b, 0x85, 0x8b, 0xba, 0x7c, 0x59, 0xfc, 0x58, 0xe0, 0x36, 0x33, 0x37, 0x4d, 0xd7,
	0xd4, 0x15, 0x99, 0x6c, 0x42, 0xff, 0x3d, 0x4c, 0x56, 0x45, 0x14, 0x21, 0xd1, 0x85, 0xa0, 0x64,
	0x29, 0x54, 0xc4, 0xf1, 0x67, 0x81, 0x64, 0xd8, 0x09, 0xd8, 0x32, 0xf6, 0xac, 0x99, 0x35, 0x77,
	0xb9, 0x2d, 0x63, 0xf6, 0x14, 0x8e, 0x52, 0x9d, 0x46, 0xe8, 0xd9, 0x75, 0xa9, 0x49, 0xfc, 0x15,
	0x4c, 0x56, 0x98, 0xc6, 0xcb, 0x04, 0xa3, 0x0d, 0x47, 0x2a, 0x94, 0xf9, 0xaf, 0x79, 0x36, 0x86,
	0x23, 0x49, 0x6b, 0xbd, 0xf1, 0x9c, 0x99, 0x35, 0x1f, 0xf2, 0x81, 0xa4, 0xaf, 0x1b, 0xff, 0x1d,
	0xb8, 0xa1, 0xc8, 0x71, 0xa3, 0xb7, 0x92, 0xd8, 0x02, 0x5c, 0xd3, 0x25, 0x9e, 0x35, 0x73, 0xe6,
	0xa3, 0xf3, 0xd3, 0xa0, 0x7d, 0x66, 0xd0, 0xb5, 0xf1, 0xdb, 0x1e, 0xff, 0xb7, 0x05, 0x8f, 0x2a,
	0x4f, 0x61, 0xd9, 0x59, 0x39, 0x03, 0xb7, 0x20, 0xcc, 0xd7, 0xa9, 0xd8, 0x62, 0xeb, 0x68, 0x58,
	0x15, 0xbe, 0x88, 0x2d, 0xb2, 0x27, 0xe0, 0x14, 0xb9, 0x6a, 0x5d, 0x55, 0x21, 0x63, 0x30, 0x88,
	0x91, 0xa2, 0xda, 0x92, 0xcb, 0xeb, 0x98, 0x05, 0x30, 0x16, 0x59, 0x96, 0xeb, 0x6b, 0x8c, 0xd7,
	0xbb, 0x5d, 0xe4, 0x0d, 0x66, 0xce, 0xdc, 0xe5, 0xa7, 0xdd, 0xa7, 0xef, 0xed, 0x52, 0xf2, 0x7f,
	0x81, 0x1d, 0x96, 0xf7, 0x21, 0xfc, 0x12, 0x9c, 0xb0, 0x24, 0xf6, 0x1c, 0x1c, 0x53, 0x76, 0xf7,
	0x1a, 0xdd, 0xde, 0xab, 0xe4, 0x55, 0xdd, 0x7f, 0x05, 0xe3, 0x0b, 0x14, 0xca, 0x24, 0xdd, 0x8f,
	0xcb, 0x74, 0x4a, 0x58, 0x19, 0x48, 0x04, 0x25, 0xad, 0xd5, 0x3a, 0x3e, 0xff, 0x63, 0xc3, 0xf0,
	0x53, 0x4b, 0x19, 0x5b, 0xc2, 0x68, 0x6f, 0x8e, 0x4d, 0x82, 0x06, 0xb3, 0xa0, 0xc3, 0x2c, 0xf8,
	0x50, 0x61, 0x36, 0x7d, 0xb6, 0x13, 0x3c, 0xa4, 0xf2, 0x19, 0x1e, 0xff, 0xc3, 0x1c, 0x7b, 0xb1,
	0x1b, 0x38, 0x4c, 0xe3, 0xb4, 0x47, 0xa9, 0xde, 0x75, 0x97, 0xbf, 0xfd, 0x5d, 0x07, 0xc9, 0xec,
	0xdd, 0xf5, 0x1a, 0xdc, 0x8f, 0x68, 0xc2, 0xf2, 0x9b, 0xd6, 0xaa, 0xf7, 0x69, 0x0f, 0xf7, 0x6e,
	0x49, 0xec, 0x2d, 0x1c, 0x37, 0xa8, 0xb1, 0xc9, 0x1d, 0xd5, 0x1d, 0x7b, 0x7d, 0x62, 0x97, 0xc7,
	0x75, 0xfe, 0xe6, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x6d, 0xdc, 0xd6, 0xc2, 0x03, 0x00,
	0x00,
}
