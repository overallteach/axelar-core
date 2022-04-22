// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: axelar/tss/v1beta1/service.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/axelarnetwork/axelar-core/x/snapshot/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("axelar/tss/v1beta1/service.proto", fileDescriptor_03d0b5f2955aa837) }
func init() {
	golang_proto.RegisterFile("axelar/tss/v1beta1/service.proto", fileDescriptor_03d0b5f2955aa837)
}

var fileDescriptor_03d0b5f2955aa837 = []byte{
	// 711 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x95, 0xcf, 0x4f, 0xd4, 0x4e,
	0x14, 0xc0, 0xb7, 0xdf, 0x6f, 0xc2, 0x37, 0xcc, 0x57, 0x2f, 0x13, 0x88, 0x64, 0xc1, 0x8a, 0x45,
	0xf0, 0x47, 0xdc, 0x96, 0x1f, 0x1a, 0x13, 0x13, 0x0f, 0x12, 0x4d, 0x34, 0x04, 0x82, 0xac, 0xe1,
	0xc0, 0x85, 0x4c, 0xbb, 0x8f, 0xee, 0x84, 0xa5, 0xb3, 0xcc, 0x4c, 0x71, 0x0b, 0x21, 0x46, 0xe2,
	0xc5, 0x83, 0x89, 0xc6, 0x3f, 0xc2, 0x3f, 0xc2, 0x8b, 0x47, 0x8f, 0x24, 0x7a, 0xf0, 0x68, 0x58,
	0xff, 0x10, 0x33, 0xdd, 0x69, 0xe9, 0xae, 0xc3, 0x2e, 0xdc, 0x60, 0xdf, 0xe7, 0xbd, 0xf7, 0xe9,
	0x9b, 0xd7, 0x29, 0x9a, 0x24, 0x2d, 0x68, 0x10, 0xee, 0x49, 0x21, 0xbc, 0xbd, 0x39, 0x1f, 0x24,
	0x99, 0xf3, 0x04, 0xf0, 0x3d, 0x1a, 0x80, 0xdb, 0xe4, 0x4c, 0x32, 0x8c, 0x3b, 0x84, 0x2b, 0x85,
	0x70, 0x35, 0x51, 0x1e, 0x09, 0x59, 0xc8, 0xd2, 0xb0, 0xa7, 0xfe, 0xea, 0x90, 0xe5, 0x89, 0x90,
	0xb1, 0xb0, 0x01, 0x1e, 0x69, 0x52, 0x8f, 0x44, 0x11, 0x93, 0x44, 0x52, 0x16, 0x09, 0x1d, 0xcd,
	0x3a, 0x89, 0x88, 0x34, 0x45, 0x9d, 0xc9, 0xbc, 0x9d, 0x6c, 0x69, 0x62, 0xdc, 0xe0, 0x92, 0x07,
	0x6d, 0x43, 0x70, 0x37, 0x06, 0x9e, 0x74, 0xe2, 0xf3, 0x1f, 0x11, 0x42, 0xcb, 0x22, 0xac, 0x76,
	0xdc, 0xf1, 0x67, 0x0b, 0x8d, 0xac, 0x41, 0x48, 0x85, 0x04, 0xfe, 0xb4, 0x25, 0x81, 0x47, 0xa4,
	0xb1, 0x04, 0x89, 0xc0, 0x9e, 0xfb, 0xf7, 0xf3, 0xb8, 0x26, 0x72, 0x0d, 0x76, 0x63, 0x10, 0xb2,
	0x3c, 0x7b, 0xfe, 0x04, 0xd1, 0x64, 0x91, 0x00, 0xe7, 0xee, 0xd1, 0xf7, 0xdf, 0x9f, 0xfe, 0x99,
	0x71, 0xae, 0x7b, 0x05, 0x67, 0xae, 0x33, 0x2a, 0xa0, 0x53, 0x2a, 0xdb, 0x90, 0x3c, 0xb4, 0xee,
	0xe0, 0x7d, 0x34, 0xfc, 0x0c, 0x08, 0x97, 0x8b, 0x40, 0x24, 0xbe, 0x61, 0x6a, 0x96, 0x87, 0x33,
	0xa5, 0xe9, 0x01, 0x94, 0xf6, 0x98, 0x4c, 0x3d, 0xca, 0xce, 0x68, 0xd1, 0xa3, 0xae, 0x30, 0x1f,
	0x88, 0x54, 0xbd, 0x8f, 0x2c, 0xf4, 0x7f, 0x55, 0x12, 0x2e, 0x97, 0x20, 0x09, 0x21, 0xc2, 0x33,
	0xa6, 0xc2, 0x05, 0x20, 0x13, 0xb8, 0x39, 0x90, 0xd3, 0x0a, 0x4e, 0xaa, 0x30, 0xe1, 0x5c, 0x29,
	0x2a, 0x88, 0x53, 0x50, 0x49, 0xbc, 0xb1, 0xd0, 0xc8, 0x2a, 0x67, 0x01, 0x08, 0xd1, 0xf9, 0xf1,
	0x25, 0x27, 0x5b, 0x5b, 0x34, 0x30, 0x1f, 0x95, 0x89, 0xec, 0x7b, 0x54, 0xe6, 0x04, 0xed, 0x37,
	0x94, 0xfa, 0x95, 0xf0, 0x6b, 0x34, 0xbc, 0xa6, 0xf6, 0x15, 0x96, 0x20, 0x31, 0x1f, 0x42, 0x1e,
	0xee, 0x7b, 0x08, 0x05, 0x4a, 0x77, 0x98, 0x4e, 0x3b, 0x5c, 0x73, 0xca, 0xc5, 0x09, 0x10, 0x21,
	0x68, 0x18, 0x79, 0x07, 0x41, 0x9d, 0xd0, 0xe8, 0x50, 0x0d, 0x21, 0x40, 0x68, 0x9d, 0x49, 0x58,
	0x8d, 0x7d, 0x65, 0x60, 0xac, 0x7d, 0x1a, 0xcf, 0x14, 0x66, 0x06, 0x61, 0x3d, 0x4f, 0x79, 0x80,
	0xb0, 0x9e, 0x46, 0x95, 0x86, 0xf9, 0x98, 0x2b, 0x7d, 0xa6, 0x56, 0xe0, 0xb2, 0xa6, 0xee, 0x79,
	0xf1, 0x9e, 0xe6, 0x1b, 0xe8, 0x3f, 0xa5, 0x56, 0xa5, 0x21, 0x76, 0xce, 0xf2, 0xae, 0xd2, 0x30,
	0x6b, 0x33, 0xd5, 0x97, 0xe9, 0xa9, 0xfd, 0xd6, 0x42, 0xa3, 0xd5, 0xd8, 0xdf, 0xa1, 0x72, 0x39,
	0x6e, 0x48, 0x2a, 0x68, 0xd8, 0x99, 0x80, 0xc0, 0xc6, 0x95, 0x30, 0xa2, 0x59, 0xe3, 0xb9, 0x0b,
	0x64, 0xf4, 0x68, 0xbc, 0xb7, 0xd0, 0x58, 0x37, 0xa9, 0x06, 0x42, 0x64, 0xcc, 0x41, 0xe0, 0x85,
	0xc1, 0x75, 0x4f, 0xe9, 0x4c, 0xe6, 0xde, 0xc5, 0x92, 0xba, 0x7d, 0xe6, 0x7f, 0xfc, 0x8b, 0x2e,
	0xbd, 0x50, 0x77, 0x64, 0x76, 0x2b, 0xee, 0xa3, 0xe1, 0x15, 0x68, 0xa9, 0x77, 0xef, 0xf9, 0x13,
	0xf3, 0x9a, 0xe7, 0xe1, 0xbe, 0x6b, 0x5e, 0xa0, 0xba, 0xef, 0x1a, 0x3c, 0xd6, 0x75, 0x41, 0x47,
	0xd0, 0x92, 0x9b, 0xdb, 0x90, 0x6c, 0xd2, 0x1a, 0x7e, 0x67, 0xa1, 0xcb, 0x8f, 0xd3, 0xbd, 0x27,
	0x7e, 0x23, 0x7d, 0xcf, 0x6e, 0x99, 0x4a, 0x77, 0x21, 0x99, 0xc4, 0xed, 0x73, 0x90, 0x5a, 0x64,
	0x2a, 0x15, 0xb9, 0x8a, 0xc7, 0xbb, 0x44, 0x48, 0xce, 0x2a, 0x1d, 0xfc, 0xc5, 0x42, 0xa3, 0xeb,
	0xa4, 0x41, 0x6b, 0x44, 0x32, 0x9e, 0x0d, 0xf2, 0xec, 0x7d, 0x31, 0xa2, 0x7d, 0xf7, 0xe5, 0x8c,
	0x0c, 0xed, 0xf8, 0x28, 0x75, 0x7c, 0x80, 0xef, 0x7b, 0x86, 0x8f, 0xda, 0x5e, 0x96, 0xba, 0xb9,
	0xa3, 0x73, 0x95, 0xb2, 0xf0, 0x0e, 0x48, 0xad, 0xc6, 0x41, 0x88, 0xc3, 0xc5, 0x95, 0x6f, 0x27,
	0xb6, 0x75, 0x7c, 0x62, 0x5b, 0xbf, 0x4e, 0x6c, 0xeb, 0x43, 0xdb, 0x2e, 0x7d, 0x6d, 0xdb, 0xd6,
	0x71, 0xdb, 0x2e, 0xfd, 0x6c, 0xdb, 0xa5, 0x8d, 0xd9, 0x90, 0xca, 0x7a, 0xec, 0xbb, 0x01, 0xdb,
	0xd1, 0xe5, 0x23, 0x90, 0xaf, 0x18, 0xdf, 0xd6, 0xff, 0x55, 0x02, 0xc6, 0xc1, 0x6b, 0xa5, 0x3d,
	0x65, 0xd2, 0x04, 0xe1, 0x0f, 0xa5, 0x5f, 0xd0, 0x85, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd5,
	0x2d, 0xcd, 0x84, 0x0c, 0x08, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgServiceClient is the client API for MsgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgServiceClient interface {
	RegisterExternalKeys(ctx context.Context, in *RegisterExternalKeysRequest, opts ...grpc.CallOption) (*RegisterExternalKeysResponse, error)
	HeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*HeartBeatResponse, error)
	StartKeygen(ctx context.Context, in *StartKeygenRequest, opts ...grpc.CallOption) (*StartKeygenResponse, error)
	ProcessKeygenTraffic(ctx context.Context, in *ProcessKeygenTrafficRequest, opts ...grpc.CallOption) (*ProcessKeygenTrafficResponse, error)
	RotateKey(ctx context.Context, in *RotateKeyRequest, opts ...grpc.CallOption) (*RotateKeyResponse, error)
	VotePubKey(ctx context.Context, in *VotePubKeyRequest, opts ...grpc.CallOption) (*VotePubKeyResponse, error)
	ProcessSignTraffic(ctx context.Context, in *ProcessSignTrafficRequest, opts ...grpc.CallOption) (*ProcessSignTrafficResponse, error)
	VoteSig(ctx context.Context, in *VoteSigRequest, opts ...grpc.CallOption) (*VoteSigResponse, error)
	SubmitMultisigPubKeys(ctx context.Context, in *SubmitMultisigPubKeysRequest, opts ...grpc.CallOption) (*SubmitMultisigPubKeysResponse, error)
	SubmitMultisigSignatures(ctx context.Context, in *SubmitMultisigSignaturesRequest, opts ...grpc.CallOption) (*SubmitMultisigSignaturesResponse, error)
}

type msgServiceClient struct {
	cc grpc1.ClientConn
}

func NewMsgServiceClient(cc grpc1.ClientConn) MsgServiceClient {
	return &msgServiceClient{cc}
}

func (c *msgServiceClient) RegisterExternalKeys(ctx context.Context, in *RegisterExternalKeysRequest, opts ...grpc.CallOption) (*RegisterExternalKeysResponse, error) {
	out := new(RegisterExternalKeysResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.MsgService/RegisterExternalKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) HeartBeat(ctx context.Context, in *HeartBeatRequest, opts ...grpc.CallOption) (*HeartBeatResponse, error) {
	out := new(HeartBeatResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.MsgService/HeartBeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) StartKeygen(ctx context.Context, in *StartKeygenRequest, opts ...grpc.CallOption) (*StartKeygenResponse, error) {
	out := new(StartKeygenResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.MsgService/StartKeygen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) ProcessKeygenTraffic(ctx context.Context, in *ProcessKeygenTrafficRequest, opts ...grpc.CallOption) (*ProcessKeygenTrafficResponse, error) {
	out := new(ProcessKeygenTrafficResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.MsgService/ProcessKeygenTraffic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) RotateKey(ctx context.Context, in *RotateKeyRequest, opts ...grpc.CallOption) (*RotateKeyResponse, error) {
	out := new(RotateKeyResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.MsgService/RotateKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) VotePubKey(ctx context.Context, in *VotePubKeyRequest, opts ...grpc.CallOption) (*VotePubKeyResponse, error) {
	out := new(VotePubKeyResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.MsgService/VotePubKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) ProcessSignTraffic(ctx context.Context, in *ProcessSignTrafficRequest, opts ...grpc.CallOption) (*ProcessSignTrafficResponse, error) {
	out := new(ProcessSignTrafficResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.MsgService/ProcessSignTraffic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) VoteSig(ctx context.Context, in *VoteSigRequest, opts ...grpc.CallOption) (*VoteSigResponse, error) {
	out := new(VoteSigResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.MsgService/VoteSig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) SubmitMultisigPubKeys(ctx context.Context, in *SubmitMultisigPubKeysRequest, opts ...grpc.CallOption) (*SubmitMultisigPubKeysResponse, error) {
	out := new(SubmitMultisigPubKeysResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.MsgService/SubmitMultisigPubKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgServiceClient) SubmitMultisigSignatures(ctx context.Context, in *SubmitMultisigSignaturesRequest, opts ...grpc.CallOption) (*SubmitMultisigSignaturesResponse, error) {
	out := new(SubmitMultisigSignaturesResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.MsgService/SubmitMultisigSignatures", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServiceServer is the server API for MsgService service.
type MsgServiceServer interface {
	RegisterExternalKeys(context.Context, *RegisterExternalKeysRequest) (*RegisterExternalKeysResponse, error)
	HeartBeat(context.Context, *HeartBeatRequest) (*HeartBeatResponse, error)
	StartKeygen(context.Context, *StartKeygenRequest) (*StartKeygenResponse, error)
	ProcessKeygenTraffic(context.Context, *ProcessKeygenTrafficRequest) (*ProcessKeygenTrafficResponse, error)
	RotateKey(context.Context, *RotateKeyRequest) (*RotateKeyResponse, error)
	VotePubKey(context.Context, *VotePubKeyRequest) (*VotePubKeyResponse, error)
	ProcessSignTraffic(context.Context, *ProcessSignTrafficRequest) (*ProcessSignTrafficResponse, error)
	VoteSig(context.Context, *VoteSigRequest) (*VoteSigResponse, error)
	SubmitMultisigPubKeys(context.Context, *SubmitMultisigPubKeysRequest) (*SubmitMultisigPubKeysResponse, error)
	SubmitMultisigSignatures(context.Context, *SubmitMultisigSignaturesRequest) (*SubmitMultisigSignaturesResponse, error)
}

// UnimplementedMsgServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServiceServer struct {
}

func (*UnimplementedMsgServiceServer) RegisterExternalKeys(ctx context.Context, req *RegisterExternalKeysRequest) (*RegisterExternalKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterExternalKeys not implemented")
}
func (*UnimplementedMsgServiceServer) HeartBeat(ctx context.Context, req *HeartBeatRequest) (*HeartBeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HeartBeat not implemented")
}
func (*UnimplementedMsgServiceServer) StartKeygen(ctx context.Context, req *StartKeygenRequest) (*StartKeygenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartKeygen not implemented")
}
func (*UnimplementedMsgServiceServer) ProcessKeygenTraffic(ctx context.Context, req *ProcessKeygenTrafficRequest) (*ProcessKeygenTrafficResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessKeygenTraffic not implemented")
}
func (*UnimplementedMsgServiceServer) RotateKey(ctx context.Context, req *RotateKeyRequest) (*RotateKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RotateKey not implemented")
}
func (*UnimplementedMsgServiceServer) VotePubKey(ctx context.Context, req *VotePubKeyRequest) (*VotePubKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VotePubKey not implemented")
}
func (*UnimplementedMsgServiceServer) ProcessSignTraffic(ctx context.Context, req *ProcessSignTrafficRequest) (*ProcessSignTrafficResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessSignTraffic not implemented")
}
func (*UnimplementedMsgServiceServer) VoteSig(ctx context.Context, req *VoteSigRequest) (*VoteSigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteSig not implemented")
}
func (*UnimplementedMsgServiceServer) SubmitMultisigPubKeys(ctx context.Context, req *SubmitMultisigPubKeysRequest) (*SubmitMultisigPubKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitMultisigPubKeys not implemented")
}
func (*UnimplementedMsgServiceServer) SubmitMultisigSignatures(ctx context.Context, req *SubmitMultisigSignaturesRequest) (*SubmitMultisigSignaturesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitMultisigSignatures not implemented")
}

func RegisterMsgServiceServer(s grpc1.Server, srv MsgServiceServer) {
	s.RegisterService(&_MsgService_serviceDesc, srv)
}

func _MsgService_RegisterExternalKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterExternalKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).RegisterExternalKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.MsgService/RegisterExternalKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).RegisterExternalKeys(ctx, req.(*RegisterExternalKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_HeartBeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartBeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).HeartBeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.MsgService/HeartBeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).HeartBeat(ctx, req.(*HeartBeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_StartKeygen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartKeygenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).StartKeygen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.MsgService/StartKeygen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).StartKeygen(ctx, req.(*StartKeygenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_ProcessKeygenTraffic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessKeygenTrafficRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).ProcessKeygenTraffic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.MsgService/ProcessKeygenTraffic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).ProcessKeygenTraffic(ctx, req.(*ProcessKeygenTrafficRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_RotateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RotateKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).RotateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.MsgService/RotateKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).RotateKey(ctx, req.(*RotateKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_VotePubKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VotePubKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).VotePubKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.MsgService/VotePubKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).VotePubKey(ctx, req.(*VotePubKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_ProcessSignTraffic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessSignTrafficRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).ProcessSignTraffic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.MsgService/ProcessSignTraffic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).ProcessSignTraffic(ctx, req.(*ProcessSignTrafficRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_VoteSig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteSigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).VoteSig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.MsgService/VoteSig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).VoteSig(ctx, req.(*VoteSigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_SubmitMultisigPubKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitMultisigPubKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).SubmitMultisigPubKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.MsgService/SubmitMultisigPubKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).SubmitMultisigPubKeys(ctx, req.(*SubmitMultisigPubKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MsgService_SubmitMultisigSignatures_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitMultisigSignaturesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServiceServer).SubmitMultisigSignatures(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.MsgService/SubmitMultisigSignatures",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServiceServer).SubmitMultisigSignatures(ctx, req.(*SubmitMultisigSignaturesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MsgService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "axelar.tss.v1beta1.MsgService",
	HandlerType: (*MsgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterExternalKeys",
			Handler:    _MsgService_RegisterExternalKeys_Handler,
		},
		{
			MethodName: "HeartBeat",
			Handler:    _MsgService_HeartBeat_Handler,
		},
		{
			MethodName: "StartKeygen",
			Handler:    _MsgService_StartKeygen_Handler,
		},
		{
			MethodName: "ProcessKeygenTraffic",
			Handler:    _MsgService_ProcessKeygenTraffic_Handler,
		},
		{
			MethodName: "RotateKey",
			Handler:    _MsgService_RotateKey_Handler,
		},
		{
			MethodName: "VotePubKey",
			Handler:    _MsgService_VotePubKey_Handler,
		},
		{
			MethodName: "ProcessSignTraffic",
			Handler:    _MsgService_ProcessSignTraffic_Handler,
		},
		{
			MethodName: "VoteSig",
			Handler:    _MsgService_VoteSig_Handler,
		},
		{
			MethodName: "SubmitMultisigPubKeys",
			Handler:    _MsgService_SubmitMultisigPubKeys_Handler,
		},
		{
			MethodName: "SubmitMultisigSignatures",
			Handler:    _MsgService_SubmitMultisigSignatures_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "axelar/tss/v1beta1/service.proto",
}

// QueryServiceClient is the client API for QueryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryServiceClient interface {
	// NextKeyID returns the key ID assigned for the next rotation on a given
	// chain and for the given key role
	NextKeyID(ctx context.Context, in *NextKeyIDRequest, opts ...grpc.CallOption) (*NextKeyIDResponse, error)
	// AssignableKey returns true if there is no assigned key for the
	// next rotation on a given chain, and false otherwise
	AssignableKey(ctx context.Context, in *AssignableKeyRequest, opts ...grpc.CallOption) (*AssignableKeyResponse, error)
	// ValidatorMultisigKeys returns the validator's multisig pubkeys
	// corresponding to each active key ID
	ValidatorMultisigKeys(ctx context.Context, in *ValidatorMultisigKeysRequest, opts ...grpc.CallOption) (*ValidatorMultisigKeysResponse, error)
}

type queryServiceClient struct {
	cc grpc1.ClientConn
}

func NewQueryServiceClient(cc grpc1.ClientConn) QueryServiceClient {
	return &queryServiceClient{cc}
}

func (c *queryServiceClient) NextKeyID(ctx context.Context, in *NextKeyIDRequest, opts ...grpc.CallOption) (*NextKeyIDResponse, error) {
	out := new(NextKeyIDResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.QueryService/NextKeyID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) AssignableKey(ctx context.Context, in *AssignableKeyRequest, opts ...grpc.CallOption) (*AssignableKeyResponse, error) {
	out := new(AssignableKeyResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.QueryService/AssignableKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceClient) ValidatorMultisigKeys(ctx context.Context, in *ValidatorMultisigKeysRequest, opts ...grpc.CallOption) (*ValidatorMultisigKeysResponse, error) {
	out := new(ValidatorMultisigKeysResponse)
	err := c.cc.Invoke(ctx, "/axelar.tss.v1beta1.QueryService/ValidatorMultisigKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServiceServer is the server API for QueryService service.
type QueryServiceServer interface {
	// NextKeyID returns the key ID assigned for the next rotation on a given
	// chain and for the given key role
	NextKeyID(context.Context, *NextKeyIDRequest) (*NextKeyIDResponse, error)
	// AssignableKey returns true if there is no assigned key for the
	// next rotation on a given chain, and false otherwise
	AssignableKey(context.Context, *AssignableKeyRequest) (*AssignableKeyResponse, error)
	// ValidatorMultisigKeys returns the validator's multisig pubkeys
	// corresponding to each active key ID
	ValidatorMultisigKeys(context.Context, *ValidatorMultisigKeysRequest) (*ValidatorMultisigKeysResponse, error)
}

// UnimplementedQueryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServiceServer struct {
}

func (*UnimplementedQueryServiceServer) NextKeyID(ctx context.Context, req *NextKeyIDRequest) (*NextKeyIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NextKeyID not implemented")
}
func (*UnimplementedQueryServiceServer) AssignableKey(ctx context.Context, req *AssignableKeyRequest) (*AssignableKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignableKey not implemented")
}
func (*UnimplementedQueryServiceServer) ValidatorMultisigKeys(ctx context.Context, req *ValidatorMultisigKeysRequest) (*ValidatorMultisigKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatorMultisigKeys not implemented")
}

func RegisterQueryServiceServer(s grpc1.Server, srv QueryServiceServer) {
	s.RegisterService(&_QueryService_serviceDesc, srv)
}

func _QueryService_NextKeyID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NextKeyIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).NextKeyID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.QueryService/NextKeyID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).NextKeyID(ctx, req.(*NextKeyIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_AssignableKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignableKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).AssignableKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.QueryService/AssignableKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).AssignableKey(ctx, req.(*AssignableKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryService_ValidatorMultisigKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidatorMultisigKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceServer).ValidatorMultisigKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/axelar.tss.v1beta1.QueryService/ValidatorMultisigKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceServer).ValidatorMultisigKeys(ctx, req.(*ValidatorMultisigKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _QueryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "axelar.tss.v1beta1.QueryService",
	HandlerType: (*QueryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NextKeyID",
			Handler:    _QueryService_NextKeyID_Handler,
		},
		{
			MethodName: "AssignableKey",
			Handler:    _QueryService_AssignableKey_Handler,
		},
		{
			MethodName: "ValidatorMultisigKeys",
			Handler:    _QueryService_ValidatorMultisigKeys_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "axelar/tss/v1beta1/service.proto",
}
