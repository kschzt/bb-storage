// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: pkg/proto/replicator/replicator.proto

package replicator

import (
	context "context"
	v2 "github.com/bazelbuild/remote-apis/build/bazel/remote/execution/v2"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReplicateBlobsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InstanceName string       `protobuf:"bytes,1,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	BlobDigests  []*v2.Digest `protobuf:"bytes,2,rep,name=blob_digests,json=blobDigests,proto3" json:"blob_digests,omitempty"`
}

func (x *ReplicateBlobsRequest) Reset() {
	*x = ReplicateBlobsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_proto_replicator_replicator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplicateBlobsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplicateBlobsRequest) ProtoMessage() {}

func (x *ReplicateBlobsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_proto_replicator_replicator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplicateBlobsRequest.ProtoReflect.Descriptor instead.
func (*ReplicateBlobsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_proto_replicator_replicator_proto_rawDescGZIP(), []int{0}
}

func (x *ReplicateBlobsRequest) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *ReplicateBlobsRequest) GetBlobDigests() []*v2.Digest {
	if x != nil {
		return x.BlobDigests
	}
	return nil
}

var File_pkg_proto_replicator_replicator_proto protoreflect.FileDescriptor

var file_pkg_proto_replicator_replicator_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x70, 0x6c,
	0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61,
	0x72, 0x6e, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x1a, 0x36, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2f, 0x72, 0x65, 0x6d, 0x6f, 0x74,
	0x65, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x32, 0x2f, 0x72,
	0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x88, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65,
	0x42, 0x6c, 0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d,
	0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x4a, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x62, 0x5f, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2e,
	0x62, 0x61, 0x7a, 0x65, 0x6c, 0x2e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x2e, 0x65, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x32, 0x2e, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x62, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x73, 0x32, 0x63, 0x0a,
	0x0a, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x55, 0x0a, 0x0e, 0x52,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x62, 0x73, 0x12, 0x2b, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x65, 0x42, 0x6c,
	0x6f, 0x62, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x62, 0x61, 0x72, 0x6e, 0x2f, 0x62, 0x62, 0x2d, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pkg_proto_replicator_replicator_proto_rawDescOnce sync.Once
	file_pkg_proto_replicator_replicator_proto_rawDescData = file_pkg_proto_replicator_replicator_proto_rawDesc
)

func file_pkg_proto_replicator_replicator_proto_rawDescGZIP() []byte {
	file_pkg_proto_replicator_replicator_proto_rawDescOnce.Do(func() {
		file_pkg_proto_replicator_replicator_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_proto_replicator_replicator_proto_rawDescData)
	})
	return file_pkg_proto_replicator_replicator_proto_rawDescData
}

var file_pkg_proto_replicator_replicator_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_proto_replicator_replicator_proto_goTypes = []interface{}{
	(*ReplicateBlobsRequest)(nil), // 0: buildbarn.replicator.ReplicateBlobsRequest
	(*v2.Digest)(nil),             // 1: build.bazel.remote.execution.v2.Digest
	(*emptypb.Empty)(nil),         // 2: google.protobuf.Empty
}
var file_pkg_proto_replicator_replicator_proto_depIdxs = []int32{
	1, // 0: buildbarn.replicator.ReplicateBlobsRequest.blob_digests:type_name -> build.bazel.remote.execution.v2.Digest
	0, // 1: buildbarn.replicator.Replicator.ReplicateBlobs:input_type -> buildbarn.replicator.ReplicateBlobsRequest
	2, // 2: buildbarn.replicator.Replicator.ReplicateBlobs:output_type -> google.protobuf.Empty
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_proto_replicator_replicator_proto_init() }
func file_pkg_proto_replicator_replicator_proto_init() {
	if File_pkg_proto_replicator_replicator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_proto_replicator_replicator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplicateBlobsRequest); i {
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
			RawDescriptor: file_pkg_proto_replicator_replicator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_proto_replicator_replicator_proto_goTypes,
		DependencyIndexes: file_pkg_proto_replicator_replicator_proto_depIdxs,
		MessageInfos:      file_pkg_proto_replicator_replicator_proto_msgTypes,
	}.Build()
	File_pkg_proto_replicator_replicator_proto = out.File
	file_pkg_proto_replicator_replicator_proto_rawDesc = nil
	file_pkg_proto_replicator_replicator_proto_goTypes = nil
	file_pkg_proto_replicator_replicator_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReplicatorClient is the client API for Replicator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReplicatorClient interface {
	ReplicateBlobs(ctx context.Context, in *ReplicateBlobsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type replicatorClient struct {
	cc grpc.ClientConnInterface
}

func NewReplicatorClient(cc grpc.ClientConnInterface) ReplicatorClient {
	return &replicatorClient{cc}
}

func (c *replicatorClient) ReplicateBlobs(ctx context.Context, in *ReplicateBlobsRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/buildbarn.replicator.Replicator/ReplicateBlobs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReplicatorServer is the server API for Replicator service.
type ReplicatorServer interface {
	ReplicateBlobs(context.Context, *ReplicateBlobsRequest) (*emptypb.Empty, error)
}

// UnimplementedReplicatorServer can be embedded to have forward compatible implementations.
type UnimplementedReplicatorServer struct {
}

func (*UnimplementedReplicatorServer) ReplicateBlobs(context.Context, *ReplicateBlobsRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReplicateBlobs not implemented")
}

func RegisterReplicatorServer(s grpc.ServiceRegistrar, srv ReplicatorServer) {
	s.RegisterService(&_Replicator_serviceDesc, srv)
}

func _Replicator_ReplicateBlobs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplicateBlobsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReplicatorServer).ReplicateBlobs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildbarn.replicator.Replicator/ReplicateBlobs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReplicatorServer).ReplicateBlobs(ctx, req.(*ReplicateBlobsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Replicator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "buildbarn.replicator.Replicator",
	HandlerType: (*ReplicatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReplicateBlobs",
			Handler:    _Replicator_ReplicateBlobs_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/replicator/replicator.proto",
}
