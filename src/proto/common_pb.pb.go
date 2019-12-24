// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common_pb.proto

package com_distkv_drpc_pb_generated

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

type Status int32

const (
	Status_OK            Status = 0
	Status_KEY_NOT_FOUND Status = 1
	Status_UNKNOWN_ERROR Status = 2
	// The key of the dict is not found.
	Status_DICT_KEY_NOT_FOUND Status = 3
)

var Status_name = map[int32]string{
	0: "OK",
	1: "KEY_NOT_FOUND",
	2: "UNKNOWN_ERROR",
	3: "DICT_KEY_NOT_FOUND",
}

var Status_value = map[string]int32{
	"OK":                 0,
	"KEY_NOT_FOUND":      1,
	"UNKNOWN_ERROR":      2,
	"DICT_KEY_NOT_FOUND": 3,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7d244d554200e9e0, []int{0}
}

func init() {
	proto.RegisterEnum("com.distkv.drpc.pb.generated.Status", Status_name, Status_value)
}

func init() { proto.RegisterFile("common_pb.proto", fileDescriptor_7d244d554200e9e0) }

var fileDescriptor_7d244d554200e9e0 = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0x8b, 0x2f, 0x48, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x49, 0xce, 0xcf,
	0xd5, 0x4b, 0xc9, 0x2c, 0x2e, 0xc9, 0x2e, 0xd3, 0x4b, 0x29, 0x2a, 0x48, 0xd6, 0x2b, 0x48, 0xd2,
	0x4b, 0x4f, 0xcd, 0x4b, 0x2d, 0x4a, 0x2c, 0x49, 0x4d, 0xd1, 0xf2, 0xe3, 0x62, 0x0b, 0x2e, 0x49,
	0x2c, 0x29, 0x2d, 0x16, 0x62, 0xe3, 0x62, 0xf2, 0xf7, 0x16, 0x60, 0x10, 0x12, 0xe4, 0xe2, 0xf5,
	0x76, 0x8d, 0x8c, 0xf7, 0xf3, 0x0f, 0x89, 0x77, 0xf3, 0x0f, 0xf5, 0x73, 0x11, 0x60, 0x04, 0x09,
	0x85, 0xfa, 0x79, 0xfb, 0xf9, 0x87, 0xfb, 0xc5, 0xbb, 0x06, 0x05, 0xf9, 0x07, 0x09, 0x30, 0x09,
	0x89, 0x71, 0x09, 0xb9, 0x78, 0x3a, 0x87, 0xc4, 0xa3, 0x2a, 0x65, 0x76, 0xd2, 0xe3, 0xc2, 0x6b,
	0x9f, 0x13, 0x9f, 0x33, 0xd8, 0x79, 0x01, 0x20, 0xa7, 0x25, 0xe7, 0xe7, 0x24, 0xb1, 0x81, 0x1d,
	0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x96, 0xc1, 0xf8, 0x39, 0xb7, 0x00, 0x00, 0x00,
}