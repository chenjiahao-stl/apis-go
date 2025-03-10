// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: enum/error_reason.proto

package enum

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ErrorReason int32

const (
	ErrorReason_GREETER_UNSPECIFIED ErrorReason = 0
	ErrorReason_USER_NOT_FOUND      ErrorReason = 1
)

var ErrorReason_name = map[int32]string{
	0: "GREETER_UNSPECIFIED",
	1: "USER_NOT_FOUND",
}

var ErrorReason_value = map[string]int32{
	"GREETER_UNSPECIFIED": 0,
	"USER_NOT_FOUND":      1,
}

func (x ErrorReason) String() string {
	return proto.EnumName(ErrorReason_name, int32(x))
}

func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3af07f5fee24bd9d, []int{0}
}

func init() {
	proto.RegisterEnum("greeter.enum.ErrorReason", ErrorReason_name, ErrorReason_value)
}

func init() { proto.RegisterFile("enum/error_reason.proto", fileDescriptor_3af07f5fee24bd9d) }

var fileDescriptor_3af07f5fee24bd9d = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xcd, 0x2b, 0xcd,
	0xd5, 0x4f, 0x2d, 0x2a, 0xca, 0x2f, 0x8a, 0x2f, 0x4a, 0x4d, 0x2c, 0xce, 0xcf, 0xd3, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x49, 0x2f, 0x4a, 0x4d, 0x2d, 0x49, 0x2d, 0xd2, 0x03, 0x29, 0xd0,
	0xb2, 0xe2, 0xe2, 0x76, 0x05, 0xa9, 0x09, 0x02, 0x2b, 0x11, 0x12, 0xe7, 0x12, 0x76, 0x0f, 0x72,
	0x75, 0x0d, 0x71, 0x0d, 0x8a, 0x0f, 0xf5, 0x0b, 0x0e, 0x70, 0x75, 0xf6, 0x74, 0xf3, 0x74, 0x75,
	0x11, 0x60, 0x10, 0x12, 0xe2, 0xe2, 0x0b, 0x0d, 0x76, 0x0d, 0x8a, 0xf7, 0xf3, 0x0f, 0x89, 0x77,
	0xf3, 0x0f, 0xf5, 0x73, 0x11, 0x60, 0x74, 0xf2, 0x3a, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39,
	0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x67, 0x3c, 0x96, 0x63, 0x88, 0xb2, 0x48, 0xcf, 0x2c, 0xd1, 0x4b,
	0x2b, 0x48, 0x49, 0x2d, 0xd3, 0x2b, 0x49, 0x4d, 0xce, 0xd0, 0x2f, 0xd0, 0x2d, 0x29, 0x4a, 0x4c,
	0xc9, 0xcc, 0x4b, 0xd7, 0x2d, 0xae, 0x2c, 0xd6, 0x4f, 0x2c, 0xc8, 0x2c, 0xd6, 0x4d, 0xcf, 0x07,
	0xd3, 0xfa, 0x69, 0xa5, 0x79, 0x29, 0xfa, 0x20, 0x27, 0x58, 0x83, 0x88, 0x24, 0x36, 0xb0, 0xe3,
	0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x63, 0xce, 0x72, 0xfd, 0xb7, 0x00, 0x00, 0x00,
}
