// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/zhanglp92/rep/api/pb/user/user.proto

package pb_user

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 客户数据
type Item struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone                string   `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Sex                  string   `protobuf:"bytes,4,opt,name=sex,proto3" json:"sex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b99c6771d55b1c4, []int{0}
}

func (m *Item) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Item.Unmarshal(m, b)
}
func (m *Item) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Item.Marshal(b, m, deterministic)
}
func (m *Item) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Item.Merge(m, src)
}
func (m *Item) XXX_Size() int {
	return xxx_messageInfo_Item.Size(m)
}
func (m *Item) XXX_DiscardUnknown() {
	xxx_messageInfo_Item.DiscardUnknown(m)
}

var xxx_messageInfo_Item proto.InternalMessageInfo

func (m *Item) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Item) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Item) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *Item) GetSex() string {
	if m != nil {
		return m.Sex
	}
	return ""
}

type DQ struct {
	Data                 []*Item  `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DQ) Reset()         { *m = DQ{} }
func (m *DQ) String() string { return proto.CompactTextString(m) }
func (*DQ) ProtoMessage()    {}
func (*DQ) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b99c6771d55b1c4, []int{1}
}

func (m *DQ) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DQ.Unmarshal(m, b)
}
func (m *DQ) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DQ.Marshal(b, m, deterministic)
}
func (m *DQ) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DQ.Merge(m, src)
}
func (m *DQ) XXX_Size() int {
	return xxx_messageInfo_DQ.Size(m)
}
func (m *DQ) XXX_DiscardUnknown() {
	xxx_messageInfo_DQ.DiscardUnknown(m)
}

var xxx_messageInfo_DQ proto.InternalMessageInfo

func (m *DQ) GetData() []*Item {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Item)(nil), "pb.user.Item")
	proto.RegisterType((*DQ)(nil), "pb.user.DQ")
}

func init() {
	proto.RegisterFile("github.com/zhanglp92/rep/api/pb/user/user.proto", fileDescriptor_4b99c6771d55b1c4)
}

var fileDescriptor_4b99c6771d55b1c4 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x2c, 0x8e, 0x31, 0x0b, 0xc2, 0x30,
	0x10, 0x85, 0x49, 0x9a, 0x2a, 0x9e, 0x28, 0x12, 0x1c, 0x32, 0xd6, 0x2e, 0x76, 0x4a, 0xa0, 0x4e,
	0xee, 0x2e, 0x8e, 0xe6, 0x1f, 0x24, 0x36, 0xb4, 0x01, 0xdb, 0x1c, 0x6d, 0x0a, 0xe2, 0xaf, 0x97,
	0x46, 0x97, 0xe3, 0xbd, 0xef, 0x86, 0xef, 0x81, 0x6a, 0x7d, 0xec, 0x66, 0x2b, 0x9f, 0xa1, 0x57,
	0x9f, 0xce, 0x0c, 0xed, 0x0b, 0xaf, 0xb5, 0x1a, 0x1d, 0x2a, 0x83, 0x5e, 0xa1, 0x55, 0xf3, 0xe4,
	0xc6, 0x74, 0x24, 0x8e, 0x21, 0x06, 0xbe, 0x46, 0x2b, 0x97, 0x5a, 0x6a, 0x60, 0xf7, 0xe8, 0x7a,
	0xbe, 0x07, 0xea, 0x1b, 0x41, 0x0a, 0x52, 0xe5, 0x9a, 0xfa, 0x86, 0x73, 0x60, 0x83, 0xe9, 0x9d,
	0xa0, 0x05, 0xa9, 0x36, 0x3a, 0x65, 0x7e, 0x84, 0x1c, 0xbb, 0x30, 0x38, 0x91, 0x25, 0xf8, 0x2b,
	0xfc, 0x00, 0xd9, 0xe4, 0xde, 0x82, 0x25, 0xb6, 0xc4, 0xf2, 0x0c, 0xf4, 0xf6, 0xe0, 0x27, 0x60,
	0x8d, 0x89, 0x46, 0x90, 0x22, 0xab, 0xb6, 0xf5, 0x4e, 0xfe, 0x8d, 0x72, 0xd1, 0xe9, 0xf4, 0xb2,
	0xab, 0x34, 0xe6, 0xf2, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x29, 0xd5, 0x88, 0x26, 0xbf, 0x00, 0x00,
	0x00,
}