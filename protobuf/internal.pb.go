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
	return fileDescriptor_internal_3e28187225341aa5, []int{0}
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
	IsOk                 bool     `protobuf:"varint,2,opt,name=is_ok,json=isOk,proto3" json:"is_ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendCheckResultRequest) Reset()         { *m = SendCheckResultRequest{} }
func (m *SendCheckResultRequest) String() string { return proto.CompactTextString(m) }
func (*SendCheckResultRequest) ProtoMessage()    {}
func (*SendCheckResultRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_internal_3e28187225341aa5, []int{1}
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
	return fileDescriptor_internal_3e28187225341aa5, []int{2}
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
	return fileDescriptor_internal_3e28187225341aa5, []int{3}
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
	return fileDescriptor_internal_3e28187225341aa5, []int{4}
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
	return fileDescriptor_internal_3e28187225341aa5, []int{5}
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

func init() {
	proto.RegisterType((*SuccessHashCalcRequest)(nil), "evileye.SuccessHashCalcRequest")
	proto.RegisterType((*SendCheckResultRequest)(nil), "evileye.SendCheckResultRequest")
	proto.RegisterType((*Tarekomis)(nil), "evileye.Tarekomis")
	proto.RegisterType((*SendTxRequest)(nil), "evileye.SendTxRequest")
	proto.RegisterType((*Tx)(nil), "evileye.Tx")
	proto.RegisterType((*Txs)(nil), "evileye.Txs")
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

func init() { proto.RegisterFile("internal.proto", fileDescriptor_internal_3e28187225341aa5) }

var fileDescriptor_internal_3e28187225341aa5 = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x52, 0x4d, 0x0f, 0xd2, 0x40,
	0x10, 0x4d, 0x5b, 0x40, 0x3a, 0x28, 0xca, 0x62, 0x9a, 0xa6, 0xc4, 0x48, 0x1a, 0x0f, 0x9c, 0x4a,
	0xc4, 0x8b, 0x07, 0xf5, 0x42, 0x8c, 0x1f, 0x07, 0x35, 0xa5, 0x9e, 0xc9, 0xd2, 0x8e, 0xb0, 0xe9,
	0xc7, 0xd6, 0xee, 0x16, 0xcb, 0xd5, 0x3f, 0xe0, 0x5f, 0x36, 0xfd, 0xe2, 0xc3, 0xd0, 0xb3, 0xb7,
	0x99, 0xd7, 0x37, 0xef, 0x75, 0x67, 0x1e, 0x8c, 0x59, 0x22, 0x31, 0x4b, 0x68, 0xe4, 0xa4, 0x19,
	0x97, 0x9c, 0x3c, 0xc0, 0x23, 0x8b, 0xf0, 0x84, 0xd6, 0x6c, 0xcf, 0xf9, 0x3e, 0xc2, 0x65, 0x05,
	0xef, 0xf2, 0x1f, 0x4b, 0x8c, 0x53, 0x79, 0xaa, 0x59, 0x96, 0x4e, 0x53, 0x56, 0x97, 0xf6, 0x3b,
	0x30, 0x36, 0xb9, 0xef, 0xa3, 0x10, 0x1f, 0xa9, 0x38, 0xac, 0x69, 0xe4, 0xbb, 0xf8, 0x33, 0x47,
	0x21, 0xc9, 0x18, 0x54, 0x16, 0x98, 0xca, 0x5c, 0x59, 0xe8, 0xae, 0xca, 0x02, 0xf2, 0x14, 0xfa,
	0x09, 0x4f, 0x7c, 0x34, 0xd5, 0x0a, 0xaa, 0x1b, 0xfb, 0x2d, 0x18, 0x1b, 0x4c, 0x82, 0xf5, 0x01,
	0xfd, 0xd0, 0x45, 0x91, 0x47, 0xb2, 0x6b, 0x7e, 0x0a, 0x7d, 0x26, 0xb6, 0x3c, 0xac, 0xe6, 0x87,
	0x6e, 0x8f, 0x89, 0xaf, 0xa1, 0xfd, 0x06, 0x74, 0x8f, 0x66, 0x18, 0xf2, 0x98, 0x09, 0xb2, 0x04,
	0x5d, 0xb6, 0x8d, 0xa9, 0xcc, 0xb5, 0xc5, 0x68, 0x35, 0x71, 0x9a, 0x07, 0x39, 0x2d, 0xcd, 0xbd,
	0x70, 0xec, 0xdf, 0x0a, 0x3c, 0x2a, 0xdd, 0xbd, 0xa2, 0x35, 0x9d, 0x81, 0x9e, 0x0b, 0xcc, 0xb6,
	0x09, 0x8d, 0xb1, 0xf1, 0x1e, 0x96, 0xc0, 0x17, 0x1a, 0x23, 0x79, 0x02, 0x5a, 0x9e, 0x45, 0xcd,
	0xff, 0x97, 0x25, 0x21, 0xd0, 0x0b, 0x50, 0xf8, 0xa6, 0x56, 0x41, 0x55, 0x4d, 0x1c, 0x98, 0xd2,
	0x34, 0xcd, 0xf8, 0x11, 0x83, 0xed, 0x59, 0x4b, 0x98, 0xbd, 0xb9, 0xb6, 0xd0, 0xdd, 0x49, 0xfb,
	0xe9, 0x7b, 0x23, 0x2a, 0xec, 0x5f, 0xa0, 0x7a, 0xc5, 0xff, 0x30, 0x7e, 0x01, 0x9a, 0x57, 0x08,
	0xf2, 0x0c, 0x34, 0x59, 0xb4, 0xfb, 0x1a, 0x5d, 0xf6, 0x55, 0xb8, 0x25, 0xbe, 0xfa, 0xa3, 0xc2,
	0xf0, 0x53, 0x13, 0x12, 0xf2, 0x19, 0x1e, 0xff, 0x73, 0x6d, 0xf2, 0xfc, 0x3c, 0x71, 0x3f, 0x07,
	0x96, 0xe1, 0xd4, 0x51, 0x72, 0xda, 0x28, 0x39, 0xef, 0xcb, 0x28, 0x55, 0x5a, 0xb7, 0x97, 0xbf,
	0xd6, 0xba, 0x9b, 0x89, 0x4e, 0xad, 0x97, 0xa0, 0x7f, 0x40, 0xe9, 0x15, 0xdf, 0x38, 0x8f, 0x48,
	0x07, 0xc9, 0x7a, 0x78, 0xf5, 0x36, 0x41, 0x5e, 0xc3, 0xa0, 0x3e, 0x3d, 0x31, 0x6e, 0x5c, 0xcf,
	0x59, 0xe8, 0x32, 0xdb, 0x0d, 0xaa, 0xfe, 0xd5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x84, 0x9b,
	0x00, 0x4e, 0x3c, 0x03, 0x00, 0x00,
}
