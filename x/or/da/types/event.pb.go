// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: finschia/or/da/v1beta1/event.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

func init() {
	proto.RegisterFile("finschia/or/da/v1beta1/event.proto", fileDescriptor_722be1466a7f29a1)
}

var fileDescriptor_722be1466a7f29a1 = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4a, 0xcb, 0xcc, 0x2b,
	0x4e, 0xce, 0xc8, 0x4c, 0xd4, 0xcf, 0x2f, 0xd2, 0x4f, 0x49, 0xd4, 0x2f, 0x33, 0x4c, 0x4a, 0x2d,
	0x49, 0x34, 0xd4, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12,
	0x83, 0xa9, 0xd1, 0xcb, 0x2f, 0xd2, 0x4b, 0x49, 0xd4, 0x83, 0xaa, 0x71, 0xf2, 0x38, 0xf1, 0x48,
	0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0,
	0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xbd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd,
	0xe4, 0xfc, 0x5c, 0x7d, 0x37, 0x98, 0x05, 0x30, 0x53, 0x74, 0x8b, 0x53, 0xb2, 0xf5, 0x2b, 0xa0,
	0xf6, 0x95, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x2d, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xef, 0xa6, 0x81, 0xe7, 0x8e, 0x00, 0x00, 0x00,
}
