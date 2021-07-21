// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdexcore/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
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

func init() { proto.RegisterFile("comdexcore/query.proto", fileDescriptor_8a40dfb76e519534) }

var fileDescriptor_8a40dfb76e519534 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0xce, 0xcf, 0x4d,
	0x49, 0xad, 0x48, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x92, 0x87, 0x88, 0xfb, 0xe7, 0xa5, 0xea, 0x21, 0x54, 0x20, 0x31, 0xa5, 0x64,
	0xd2, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b,
	0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0x21, 0xda, 0xa5, 0xb4, 0x92, 0xf3, 0x8b, 0x73, 0xf3, 0x8b,
	0xf5, 0x93, 0x12, 0x8b, 0xa1, 0xe6, 0xea, 0x97, 0x19, 0x26, 0xa5, 0x96, 0x24, 0x1a, 0xea, 0x17,
	0x24, 0xa6, 0x67, 0xe6, 0x81, 0x15, 0x43, 0xd4, 0x1a, 0xb1, 0x73, 0xb1, 0x06, 0x82, 0x54, 0x38,
	0xf9, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e,
	0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x51, 0x7a, 0x66, 0x49,
	0x46, 0x69, 0x12, 0xc8, 0x7e, 0x7d, 0xb8, 0xc3, 0xa0, 0x2c, 0x67, 0x90, 0xd3, 0x2b, 0xf4, 0x91,
	0xfc, 0x51, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x36, 0xdd, 0x18, 0x10, 0x00, 0x00, 0xff,
	0xff, 0x17, 0x19, 0x4d, 0x25, 0xe2, 0x00, 0x00, 0x00,
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
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "comdexOne.comdexcore.comdexcore.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "comdexcore/query.proto",
}
