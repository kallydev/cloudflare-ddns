// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.12.3
// source: grpc.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type RefreshDNSRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Key  string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *RefreshDNSRecordRequest) Reset() {
	*x = RefreshDNSRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshDNSRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshDNSRecordRequest) ProtoMessage() {}

func (x *RefreshDNSRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshDNSRecordRequest.ProtoReflect.Descriptor instead.
func (*RefreshDNSRecordRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *RefreshDNSRecordRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RefreshDNSRecordRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type RefreshDNSRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *RefreshDNSRecordResponse) Reset() {
	*x = RefreshDNSRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshDNSRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshDNSRecordResponse) ProtoMessage() {}

func (x *RefreshDNSRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshDNSRecordResponse.ProtoReflect.Descriptor instead.
func (*RefreshDNSRecordResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *RefreshDNSRecordResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_grpc_proto protoreflect.FileDescriptor

var file_grpc_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x3f, 0x0a, 0x17, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x44, 0x4e, 0x53, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x22, 0x34, 0x0a, 0x18, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x44, 0x4e, 0x53, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x60, 0x0a, 0x0e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x46, 0x6c, 0x61, 0x72, 0x65, 0x44, 0x44, 0x4e, 0x53, 0x12, 0x4e, 0x0a, 0x11, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x44, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1b,
	0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x44, 0x4e, 0x53, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x62,
	0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x44, 0x4e, 0x53, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_rawDescOnce sync.Once
	file_grpc_proto_rawDescData = file_grpc_proto_rawDesc
)

func file_grpc_proto_rawDescGZIP() []byte {
	file_grpc_proto_rawDescOnce.Do(func() {
		file_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_rawDescData)
	})
	return file_grpc_proto_rawDescData
}

var file_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_proto_goTypes = []interface{}{
	(*RefreshDNSRecordRequest)(nil),  // 0: pb.RefreshDNSRecordRequest
	(*RefreshDNSRecordResponse)(nil), // 1: pb.RefreshDNSRecordResponse
}
var file_grpc_proto_depIdxs = []int32{
	0, // 0: pb.CloudFlareDDNS.RefreshDDNSRecord:input_type -> pb.RefreshDNSRecordRequest
	1, // 1: pb.CloudFlareDDNS.RefreshDDNSRecord:output_type -> pb.RefreshDNSRecordResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_proto_init() }
func file_grpc_proto_init() {
	if File_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshDNSRecordRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshDNSRecordResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_goTypes,
		DependencyIndexes: file_grpc_proto_depIdxs,
		MessageInfos:      file_grpc_proto_msgTypes,
	}.Build()
	File_grpc_proto = out.File
	file_grpc_proto_rawDesc = nil
	file_grpc_proto_goTypes = nil
	file_grpc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CloudFlareDDNSClient is the client API for CloudFlareDDNS service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CloudFlareDDNSClient interface {
	RefreshDDNSRecord(ctx context.Context, in *RefreshDNSRecordRequest, opts ...grpc.CallOption) (*RefreshDNSRecordResponse, error)
}

type cloudFlareDDNSClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudFlareDDNSClient(cc grpc.ClientConnInterface) CloudFlareDDNSClient {
	return &cloudFlareDDNSClient{cc}
}

func (c *cloudFlareDDNSClient) RefreshDDNSRecord(ctx context.Context, in *RefreshDNSRecordRequest, opts ...grpc.CallOption) (*RefreshDNSRecordResponse, error) {
	out := new(RefreshDNSRecordResponse)
	err := c.cc.Invoke(ctx, "/pb.CloudFlareDDNS/RefreshDDNSRecord", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudFlareDDNSServer is the server API for CloudFlareDDNS service.
type CloudFlareDDNSServer interface {
	RefreshDDNSRecord(context.Context, *RefreshDNSRecordRequest) (*RefreshDNSRecordResponse, error)
}

// UnimplementedCloudFlareDDNSServer can be embedded to have forward compatible implementations.
type UnimplementedCloudFlareDDNSServer struct {
}

func (*UnimplementedCloudFlareDDNSServer) RefreshDDNSRecord(context.Context, *RefreshDNSRecordRequest) (*RefreshDNSRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshDDNSRecord not implemented")
}

func RegisterCloudFlareDDNSServer(s *grpc.Server, srv CloudFlareDDNSServer) {
	s.RegisterService(&_CloudFlareDDNS_serviceDesc, srv)
}

func _CloudFlareDDNS_RefreshDDNSRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshDNSRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudFlareDDNSServer).RefreshDDNSRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.CloudFlareDDNS/RefreshDDNSRecord",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudFlareDDNSServer).RefreshDDNSRecord(ctx, req.(*RefreshDNSRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CloudFlareDDNS_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CloudFlareDDNS",
	HandlerType: (*CloudFlareDDNSServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RefreshDDNSRecord",
			Handler:    _CloudFlareDDNS_RefreshDDNSRecord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}
