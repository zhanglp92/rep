// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/zhanglp92/rep/api/pb/user/user.proto

package pb_user

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phone                string               `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
	Sex                  string               `protobuf:"bytes,4,opt,name=sex,proto3" json:"sex,omitempty"`
	Addr                 string               `protobuf:"bytes,5,opt,name=addr,proto3" json:"addr,omitempty"`
	Remark               string               `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark,omitempty"`
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	Screate              string               `protobuf:"bytes,8,opt,name=screate,proto3" json:"screate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
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

func (m *Item) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
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

func (m *Item) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func (m *Item) GetRemark() string {
	if m != nil {
		return m.Remark
	}
	return ""
}

func (m *Item) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Item) GetScreate() string {
	if m != nil {
		return m.Screate
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
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x90, 0xcf, 0x4e, 0xc4, 0x20,
	0x10, 0x87, 0xd3, 0x3f, 0xdb, 0xea, 0x34, 0x1a, 0x43, 0x8c, 0x21, 0x7b, 0xb1, 0xee, 0xc5, 0x9e,
	0x20, 0xa9, 0x27, 0xe3, 0xd5, 0x8b, 0x47, 0x1b, 0xef, 0x86, 0x6e, 0xc7, 0x96, 0xb8, 0x14, 0x42,
	0x69, 0x62, 0x7c, 0x4f, 0xdf, 0xc7, 0x00, 0xdb, 0x0b, 0x99, 0xef, 0x63, 0x06, 0x7e, 0x00, 0x7c,
	0x94, 0x6e, 0x5a, 0x7b, 0x76, 0xd4, 0x8a, 0xff, 0x4e, 0x62, 0x1e, 0x4f, 0xe6, 0xb9, 0xe5, 0x16,
	0x0d, 0x17, 0x46, 0x72, 0xd3, 0xf3, 0x75, 0x41, 0x1b, 0x16, 0x66, 0xac, 0x76, 0x9a, 0x94, 0xa6,
	0x67, 0x1e, 0xf7, 0xf7, 0xa3, 0xd6, 0xe3, 0x09, 0x79, 0xd0, 0xfd, 0xfa, 0xc5, 0x9d, 0x54, 0xb8,
	0x38, 0xa1, 0x4c, 0xec, 0x3c, 0xfc, 0x25, 0x90, 0xbf, 0x39, 0x54, 0xe4, 0x1a, 0x52, 0x39, 0xd0,
	0xa4, 0x4e, 0x9a, 0xcb, 0x2e, 0x95, 0x03, 0x21, 0x90, 0xcf, 0x42, 0x21, 0x4d, 0x83, 0x09, 0x35,
	0xb9, 0x85, 0x9d, 0x99, 0xf4, 0x8c, 0x34, 0x0b, 0x32, 0x02, 0xb9, 0x81, 0x6c, 0xc1, 0x1f, 0x9a,
	0x07, 0xe7, 0x4b, 0x3f, 0x2b, 0x86, 0xc1, 0xd2, 0x5d, 0x9c, 0xf5, 0x35, 0xb9, 0x83, 0xc2, 0xa2,
	0x12, 0xf6, 0x9b, 0x16, 0xc1, 0x9e, 0x89, 0xbc, 0x40, 0x75, 0xb4, 0x28, 0x1c, 0x7e, 0xfa, 0x68,
	0xb4, 0xac, 0x93, 0xa6, 0x6a, 0xf7, 0x2c, 0xe6, 0x66, 0x5b, 0x6e, 0xf6, 0xb1, 0xe5, 0xee, 0x20,
	0xb6, 0x7b, 0x41, 0x28, 0x94, 0x4b, 0x44, 0x7a, 0x11, 0x4e, 0xdd, 0xf0, 0xf0, 0x08, 0xe9, 0xeb,
	0x3b, 0x79, 0x80, 0x7c, 0x10, 0x4e, 0xd0, 0xa4, 0xce, 0x9a, 0xaa, 0xbd, 0x62, 0xe7, 0x6f, 0x61,
	0xfe, 0xc5, 0x5d, 0xd8, 0xea, 0x8b, 0x70, 0xc5, 0xd3, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbf,
	0xf7, 0xec, 0xd7, 0x64, 0x01, 0x00, 0x00,
}
