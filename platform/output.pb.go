// Code generated by protoc-gen-go. DO NOT EDIT.
// source: platform/output.proto

package platform

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

// You can customise this message to change the fields for
// the output value from your Deployment
type Deployment struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deployment) Reset()         { *m = Deployment{} }
func (m *Deployment) String() string { return proto.CompactTextString(m) }
func (*Deployment) ProtoMessage()    {}
func (*Deployment) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e5234dca2919ebb, []int{0}
}

func (m *Deployment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deployment.Unmarshal(m, b)
}
func (m *Deployment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deployment.Marshal(b, m, deterministic)
}
func (m *Deployment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deployment.Merge(m, src)
}
func (m *Deployment) XXX_Size() int {
	return xxx_messageInfo_Deployment.Size(m)
}
func (m *Deployment) XXX_DiscardUnknown() {
	xxx_messageInfo_Deployment.DiscardUnknown(m)
}

var xxx_messageInfo_Deployment proto.InternalMessageInfo

func (m *Deployment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Deployment)(nil), "platform.Deployment")
}

func init() {
	proto.RegisterFile("platform/output.proto", fileDescriptor_1e5234dca2919ebb)
}

var fileDescriptor_1e5234dca2919ebb = []byte{
	// 139 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0xc8, 0x49, 0x2c,
	0x49, 0xcb, 0x2f, 0xca, 0xd5, 0xcf, 0x2f, 0x2d, 0x29, 0x28, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x80, 0x09, 0x2b, 0xc9, 0x70, 0x71, 0xb9, 0xa4, 0x16, 0xe4, 0xe4, 0x57, 0xe6,
	0xa6, 0xe6, 0x95, 0x08, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0x31, 0x65, 0xa6, 0x38, 0x59, 0x46, 0x99, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7,
	0xe7, 0xea, 0xa7, 0x64, 0xe6, 0xa5, 0xe7, 0xa7, 0x26, 0x96, 0x64, 0xe6, 0xa5, 0xa7, 0x95, 0x56,
	0x55, 0xe9, 0x97, 0x27, 0x56, 0x16, 0xe4, 0x67, 0xe6, 0x95, 0xe8, 0x16, 0xe4, 0x94, 0xa6, 0x67,
	0xe6, 0xe9, 0x16, 0x64, 0xa7, 0xeb, 0xc3, 0x0c, 0x4e, 0x62, 0x03, 0xdb, 0x64, 0x0c, 0x08, 0x00,
	0x00, 0xff, 0xff, 0x53, 0xe4, 0x9d, 0xea, 0x82, 0x00, 0x00, 0x00,
}