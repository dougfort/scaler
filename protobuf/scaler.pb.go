// Code generated by protoc-gen-go. DO NOT EDIT.
// source: scaler.proto

package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// DeploymentsResponse contains all known deploymwnt IDs
type DeploymentsResponse struct {
	DeploymentId         []string `protobuf:"bytes,1,rep,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeploymentsResponse) Reset()         { *m = DeploymentsResponse{} }
func (m *DeploymentsResponse) String() string { return proto.CompactTextString(m) }
func (*DeploymentsResponse) ProtoMessage()    {}
func (*DeploymentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_scaler_f957fac56d0ce2e1, []int{0}
}
func (m *DeploymentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeploymentsResponse.Unmarshal(m, b)
}
func (m *DeploymentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeploymentsResponse.Marshal(b, m, deterministic)
}
func (dst *DeploymentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeploymentsResponse.Merge(dst, src)
}
func (m *DeploymentsResponse) XXX_Size() int {
	return xxx_messageInfo_DeploymentsResponse.Size(m)
}
func (m *DeploymentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeploymentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeploymentsResponse proto.InternalMessageInfo

func (m *DeploymentsResponse) GetDeploymentId() []string {
	if m != nil {
		return m.DeploymentId
	}
	return nil
}

// PlatformsResponse contains all known platform IDs
type PlatformsResponse struct {
	PlatformId           []string `protobuf:"bytes,1,rep,name=platform_id,json=platformId,proto3" json:"platform_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlatformsResponse) Reset()         { *m = PlatformsResponse{} }
func (m *PlatformsResponse) String() string { return proto.CompactTextString(m) }
func (*PlatformsResponse) ProtoMessage()    {}
func (*PlatformsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_scaler_f957fac56d0ce2e1, []int{1}
}
func (m *PlatformsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlatformsResponse.Unmarshal(m, b)
}
func (m *PlatformsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlatformsResponse.Marshal(b, m, deterministic)
}
func (dst *PlatformsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlatformsResponse.Merge(dst, src)
}
func (m *PlatformsResponse) XXX_Size() int {
	return xxx_messageInfo_PlatformsResponse.Size(m)
}
func (m *PlatformsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PlatformsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PlatformsResponse proto.InternalMessageInfo

func (m *PlatformsResponse) GetPlatformId() []string {
	if m != nil {
		return m.PlatformId
	}
	return nil
}

// GetDeploymentStatusRequest requests the status of a deployment
type GetDeploymentStatusRequest struct {
	DeploymentId         string   `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDeploymentStatusRequest) Reset()         { *m = GetDeploymentStatusRequest{} }
func (m *GetDeploymentStatusRequest) String() string { return proto.CompactTextString(m) }
func (*GetDeploymentStatusRequest) ProtoMessage()    {}
func (*GetDeploymentStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_scaler_f957fac56d0ce2e1, []int{2}
}
func (m *GetDeploymentStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDeploymentStatusRequest.Unmarshal(m, b)
}
func (m *GetDeploymentStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDeploymentStatusRequest.Marshal(b, m, deterministic)
}
func (dst *GetDeploymentStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDeploymentStatusRequest.Merge(dst, src)
}
func (m *GetDeploymentStatusRequest) XXX_Size() int {
	return xxx_messageInfo_GetDeploymentStatusRequest.Size(m)
}
func (m *GetDeploymentStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDeploymentStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDeploymentStatusRequest proto.InternalMessageInfo

func (m *GetDeploymentStatusRequest) GetDeploymentId() string {
	if m != nil {
		return m.DeploymentId
	}
	return ""
}

// DeploymentStatusResponse lists deployment status on all known platforms
type DeploymentStatusResponse struct {
	DeploymentStatus     []*DeploymentStatus `protobuf:"bytes,1,rep,name=deployment_status,json=deploymentStatus,proto3" json:"deployment_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *DeploymentStatusResponse) Reset()         { *m = DeploymentStatusResponse{} }
func (m *DeploymentStatusResponse) String() string { return proto.CompactTextString(m) }
func (*DeploymentStatusResponse) ProtoMessage()    {}
func (*DeploymentStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_scaler_f957fac56d0ce2e1, []int{3}
}
func (m *DeploymentStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeploymentStatusResponse.Unmarshal(m, b)
}
func (m *DeploymentStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeploymentStatusResponse.Marshal(b, m, deterministic)
}
func (dst *DeploymentStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeploymentStatusResponse.Merge(dst, src)
}
func (m *DeploymentStatusResponse) XXX_Size() int {
	return xxx_messageInfo_DeploymentStatusResponse.Size(m)
}
func (m *DeploymentStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeploymentStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeploymentStatusResponse proto.InternalMessageInfo

func (m *DeploymentStatusResponse) GetDeploymentStatus() []*DeploymentStatus {
	if m != nil {
		return m.DeploymentStatus
	}
	return nil
}

type SetDeploymentStatusRequest struct {
	DeploymentId         string              `protobuf:"bytes,1,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	DeploymentStatus     []*DeploymentStatus `protobuf:"bytes,2,rep,name=deployment_status,json=deploymentStatus,proto3" json:"deployment_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SetDeploymentStatusRequest) Reset()         { *m = SetDeploymentStatusRequest{} }
func (m *SetDeploymentStatusRequest) String() string { return proto.CompactTextString(m) }
func (*SetDeploymentStatusRequest) ProtoMessage()    {}
func (*SetDeploymentStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_scaler_f957fac56d0ce2e1, []int{4}
}
func (m *SetDeploymentStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDeploymentStatusRequest.Unmarshal(m, b)
}
func (m *SetDeploymentStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDeploymentStatusRequest.Marshal(b, m, deterministic)
}
func (dst *SetDeploymentStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDeploymentStatusRequest.Merge(dst, src)
}
func (m *SetDeploymentStatusRequest) XXX_Size() int {
	return xxx_messageInfo_SetDeploymentStatusRequest.Size(m)
}
func (m *SetDeploymentStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDeploymentStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetDeploymentStatusRequest proto.InternalMessageInfo

func (m *SetDeploymentStatusRequest) GetDeploymentId() string {
	if m != nil {
		return m.DeploymentId
	}
	return ""
}

func (m *SetDeploymentStatusRequest) GetDeploymentStatus() []*DeploymentStatus {
	if m != nil {
		return m.DeploymentStatus
	}
	return nil
}

// DeploymentStatus is the information about a deployment on an individual platform
type DeploymentStatus struct {
	PlatformId           string   `protobuf:"bytes,1,opt,name=platform_id,json=platformId,proto3" json:"platform_id,omitempty"`
	InstanceCount        int32    `protobuf:"varint,2,opt,name=instance_count,json=instanceCount,proto3" json:"instance_count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeploymentStatus) Reset()         { *m = DeploymentStatus{} }
func (m *DeploymentStatus) String() string { return proto.CompactTextString(m) }
func (*DeploymentStatus) ProtoMessage()    {}
func (*DeploymentStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_scaler_f957fac56d0ce2e1, []int{5}
}
func (m *DeploymentStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeploymentStatus.Unmarshal(m, b)
}
func (m *DeploymentStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeploymentStatus.Marshal(b, m, deterministic)
}
func (dst *DeploymentStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeploymentStatus.Merge(dst, src)
}
func (m *DeploymentStatus) XXX_Size() int {
	return xxx_messageInfo_DeploymentStatus.Size(m)
}
func (m *DeploymentStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_DeploymentStatus.DiscardUnknown(m)
}

var xxx_messageInfo_DeploymentStatus proto.InternalMessageInfo

func (m *DeploymentStatus) GetPlatformId() string {
	if m != nil {
		return m.PlatformId
	}
	return ""
}

func (m *DeploymentStatus) GetInstanceCount() int32 {
	if m != nil {
		return m.InstanceCount
	}
	return 0
}

func init() {
	proto.RegisterType((*DeploymentsResponse)(nil), "protobuf.DeploymentsResponse")
	proto.RegisterType((*PlatformsResponse)(nil), "protobuf.PlatformsResponse")
	proto.RegisterType((*GetDeploymentStatusRequest)(nil), "protobuf.GetDeploymentStatusRequest")
	proto.RegisterType((*DeploymentStatusResponse)(nil), "protobuf.DeploymentStatusResponse")
	proto.RegisterType((*SetDeploymentStatusRequest)(nil), "protobuf.SetDeploymentStatusRequest")
	proto.RegisterType((*DeploymentStatus)(nil), "protobuf.DeploymentStatus")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ScalerClient is the client API for Scaler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ScalerClient interface {
	// EnumerateDeployments lists all known deployment IDs
	EnumerateDeployments(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*DeploymentsResponse, error)
	// EnumeratePlatforms lists all known platform IDs
	EnumeratePlatforms(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PlatformsResponse, error)
	// GetDeploymentStatus lists the status of the deployment on all platforms
	GetDeploymentStatus(ctx context.Context, in *GetDeploymentStatusRequest, opts ...grpc.CallOption) (*DeploymentStatusResponse, error)
	// SetDeploymentStatus sets the desired deployment status
	// Returns the known deployment status at time of execution
	SetDeploymentStatus(ctx context.Context, in *SetDeploymentStatusRequest, opts ...grpc.CallOption) (*DeploymentStatusResponse, error)
}

type scalerClient struct {
	cc *grpc.ClientConn
}

func NewScalerClient(cc *grpc.ClientConn) ScalerClient {
	return &scalerClient{cc}
}

func (c *scalerClient) EnumerateDeployments(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*DeploymentsResponse, error) {
	out := new(DeploymentsResponse)
	err := c.cc.Invoke(ctx, "/protobuf.Scaler/EnumerateDeployments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scalerClient) EnumeratePlatforms(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*PlatformsResponse, error) {
	out := new(PlatformsResponse)
	err := c.cc.Invoke(ctx, "/protobuf.Scaler/EnumeratePlatforms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scalerClient) GetDeploymentStatus(ctx context.Context, in *GetDeploymentStatusRequest, opts ...grpc.CallOption) (*DeploymentStatusResponse, error) {
	out := new(DeploymentStatusResponse)
	err := c.cc.Invoke(ctx, "/protobuf.Scaler/GetDeploymentStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scalerClient) SetDeploymentStatus(ctx context.Context, in *SetDeploymentStatusRequest, opts ...grpc.CallOption) (*DeploymentStatusResponse, error) {
	out := new(DeploymentStatusResponse)
	err := c.cc.Invoke(ctx, "/protobuf.Scaler/SetDeploymentStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ScalerServer is the server API for Scaler service.
type ScalerServer interface {
	// EnumerateDeployments lists all known deployment IDs
	EnumerateDeployments(context.Context, *empty.Empty) (*DeploymentsResponse, error)
	// EnumeratePlatforms lists all known platform IDs
	EnumeratePlatforms(context.Context, *empty.Empty) (*PlatformsResponse, error)
	// GetDeploymentStatus lists the status of the deployment on all platforms
	GetDeploymentStatus(context.Context, *GetDeploymentStatusRequest) (*DeploymentStatusResponse, error)
	// SetDeploymentStatus sets the desired deployment status
	// Returns the known deployment status at time of execution
	SetDeploymentStatus(context.Context, *SetDeploymentStatusRequest) (*DeploymentStatusResponse, error)
}

func RegisterScalerServer(s *grpc.Server, srv ScalerServer) {
	s.RegisterService(&_Scaler_serviceDesc, srv)
}

func _Scaler_EnumerateDeployments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScalerServer).EnumerateDeployments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Scaler/EnumerateDeployments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScalerServer).EnumerateDeployments(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scaler_EnumeratePlatforms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScalerServer).EnumeratePlatforms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Scaler/EnumeratePlatforms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScalerServer).EnumeratePlatforms(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scaler_GetDeploymentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeploymentStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScalerServer).GetDeploymentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Scaler/GetDeploymentStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScalerServer).GetDeploymentStatus(ctx, req.(*GetDeploymentStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Scaler_SetDeploymentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDeploymentStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ScalerServer).SetDeploymentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Scaler/SetDeploymentStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ScalerServer).SetDeploymentStatus(ctx, req.(*SetDeploymentStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Scaler_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Scaler",
	HandlerType: (*ScalerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EnumerateDeployments",
			Handler:    _Scaler_EnumerateDeployments_Handler,
		},
		{
			MethodName: "EnumeratePlatforms",
			Handler:    _Scaler_EnumeratePlatforms_Handler,
		},
		{
			MethodName: "GetDeploymentStatus",
			Handler:    _Scaler_GetDeploymentStatus_Handler,
		},
		{
			MethodName: "SetDeploymentStatus",
			Handler:    _Scaler_SetDeploymentStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "scaler.proto",
}

func init() { proto.RegisterFile("scaler.proto", fileDescriptor_scaler_f957fac56d0ce2e1) }

var fileDescriptor_scaler_f957fac56d0ce2e1 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0xaa, 0xd3, 0x40,
	0x14, 0xc6, 0x99, 0xaa, 0x17, 0x3d, 0xf7, 0x5e, 0x69, 0xa7, 0x22, 0x65, 0x5a, 0xb5, 0xc4, 0x2a,
	0x45, 0x24, 0xc1, 0x5a, 0x10, 0xba, 0x13, 0x2d, 0xa5, 0x3b, 0x49, 0x76, 0x6e, 0xca, 0x34, 0x99,
	0x96, 0x40, 0x32, 0x13, 0x33, 0x13, 0xa1, 0x88, 0x1b, 0x97, 0x6e, 0x45, 0xf0, 0x1d, 0x7c, 0x1c,
	0x5f, 0xc1, 0x07, 0x91, 0x4e, 0x9a, 0xa4, 0x6d, 0xd2, 0x80, 0xb7, 0xab, 0xc0, 0x77, 0xfe, 0x7c,
	0xbf, 0x9c, 0xf9, 0xe0, 0x4a, 0xba, 0x34, 0x60, 0xb1, 0x19, 0xc5, 0x42, 0x09, 0x7c, 0x57, 0x7f,
	0x96, 0xc9, 0x8a, 0xf4, 0xd6, 0x42, 0xac, 0x03, 0x66, 0xd1, 0xc8, 0xb7, 0x28, 0xe7, 0x42, 0x51,
	0xe5, 0x0b, 0x2e, 0xd3, 0x3e, 0xd2, 0xdd, 0x55, 0xb3, 0x76, 0x8b, 0x85, 0x91, 0xda, 0xa4, 0x45,
	0x63, 0x02, 0xed, 0xf7, 0x2c, 0x0a, 0xc4, 0x26, 0x64, 0x5c, 0x49, 0x9b, 0xc9, 0x48, 0x70, 0xc9,
	0xf0, 0x53, 0xb8, 0xf6, 0x72, 0x79, 0xe1, 0x7b, 0x1d, 0xd4, 0xbf, 0x35, 0xbc, 0x67, 0x5f, 0x15,
	0xe2, 0xdc, 0x33, 0xc6, 0xd0, 0xfa, 0x10, 0x50, 0xb5, 0x12, 0x71, 0x58, 0x4c, 0x3e, 0x81, 0xcb,
	0x68, 0x27, 0x16, 0x73, 0x90, 0x49, 0x73, 0xcf, 0x78, 0x0b, 0x64, 0xc6, 0x54, 0x61, 0xea, 0x28,
	0xaa, 0x12, 0x69, 0xb3, 0x4f, 0x09, 0x93, 0xaa, 0xca, 0x18, 0x95, 0x8c, 0x5d, 0xe8, 0x94, 0xe7,
	0x77, 0xfe, 0x33, 0x68, 0xed, 0x2d, 0x90, 0xba, 0xa8, 0x29, 0x2e, 0x47, 0xc4, 0xcc, 0x4e, 0x60,
	0x96, 0xc6, 0x9b, 0xde, 0x91, 0x62, 0x7c, 0x47, 0x40, 0x9c, 0xf3, 0x40, 0xab, 0x61, 0x1a, 0x37,
	0x80, 0xf9, 0x08, 0xcd, 0xe3, 0xae, 0xf2, 0xa5, 0xd1, 0xe1, 0xa5, 0xf1, 0x33, 0xb8, 0xef, 0x73,
	0xa9, 0x28, 0x77, 0xd9, 0xc2, 0x15, 0x09, 0x57, 0x9d, 0x46, 0x1f, 0x0d, 0xef, 0xd8, 0xd7, 0x99,
	0xfa, 0x6e, 0x2b, 0x8e, 0x7e, 0xdf, 0x86, 0x0b, 0x47, 0x07, 0x0b, 0x4b, 0x78, 0x30, 0xe5, 0x49,
	0xc8, 0x62, 0xaa, 0xd8, 0x5e, 0x2c, 0xf0, 0x43, 0x33, 0xcd, 0x50, 0xc1, 0x3c, 0xdd, 0x66, 0x88,
	0x3c, 0xaa, 0xfa, 0x89, 0xfc, 0x2d, 0x8c, 0xc1, 0xb7, 0x3f, 0x7f, 0x7f, 0x34, 0x1e, 0xe3, 0x9e,
	0x25, 0x59, 0xfc, 0xd9, 0x77, 0x99, 0xb4, 0xd2, 0x04, 0x5b, 0xde, 0xde, 0xf2, 0x0d, 0xe0, 0xdc,
	0x34, 0xcf, 0xd3, 0x49, 0xcb, 0x6e, 0x21, 0x94, 0xc2, 0x67, 0x98, 0xda, 0x70, 0x88, 0x9f, 0xd7,
	0x19, 0x5a, 0x51, 0x6e, 0xf2, 0x13, 0x41, 0xbb, 0x22, 0x8c, 0x78, 0x50, 0x98, 0x9c, 0xce, 0x2a,
	0x31, 0x6a, 0x9e, 0x30, 0x23, 0x1a, 0x6b, 0x22, 0x13, 0xbf, 0xac, 0x25, 0xfa, 0x72, 0x10, 0xa5,
	0xaf, 0xf8, 0x17, 0x82, 0xb6, 0x53, 0xcf, 0xe5, 0x9c, 0xc7, 0xf5, 0x46, 0x73, 0xbd, 0x22, 0xff,
	0xc5, 0x35, 0x41, 0x2f, 0x96, 0x17, 0x7a, 0xf7, 0xeb, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdd,
	0xd3, 0x28, 0x4e, 0x8c, 0x04, 0x00, 0x00,
}
