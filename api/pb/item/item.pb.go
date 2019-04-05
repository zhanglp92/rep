// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/zhanglp92/rep/api/pb/item/item.proto

package pb_item

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	user "github.com/zhanglp92/rep/api/pb/user"
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

// 数据最基础单元
type Item struct {
	Id                   int32      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Userid               int32      `protobuf:"varint,2,opt,name=userid,proto3" json:"userid,omitempty"`
	Type                 string     `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Colour               string     `protobuf:"bytes,4,opt,name=colour,proto3" json:"colour,omitempty"`
	Opendirection        string     `protobuf:"bytes,5,opt,name=opendirection,proto3" json:"opendirection,omitempty"`
	Size                 string     `protobuf:"bytes,6,opt,name=size,proto3" json:"size,omitempty"`
	Shape                string     `protobuf:"bytes,7,opt,name=shape,proto3" json:"shape,omitempty"`
	Glass                string     `protobuf:"bytes,8,opt,name=glass,proto3" json:"glass,omitempty"`
	Squaremeter          float32    `protobuf:"fixed32,9,opt,name=squaremeter,proto3" json:"squaremeter,omitempty"`
	Priceunit            float32    `protobuf:"fixed32,10,opt,name=priceunit,proto3" json:"priceunit,omitempty"`
	Priceset             float32    `protobuf:"fixed32,11,opt,name=priceset,proto3" json:"priceset,omitempty"`
	Price                float32    `protobuf:"fixed32,12,opt,name=price,proto3" json:"price,omitempty"`
	Remarks              string     `protobuf:"bytes,13,opt,name=remarks,proto3" json:"remarks,omitempty"`
	User                 *user.Item `protobuf:"bytes,14,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Item) Reset()         { *m = Item{} }
func (m *Item) String() string { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()    {}
func (*Item) Descriptor() ([]byte, []int) {
	return fileDescriptor_fda99000515d91e5, []int{0}
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

func (m *Item) GetUserid() int32 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *Item) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Item) GetColour() string {
	if m != nil {
		return m.Colour
	}
	return ""
}

func (m *Item) GetOpendirection() string {
	if m != nil {
		return m.Opendirection
	}
	return ""
}

func (m *Item) GetSize() string {
	if m != nil {
		return m.Size
	}
	return ""
}

func (m *Item) GetShape() string {
	if m != nil {
		return m.Shape
	}
	return ""
}

func (m *Item) GetGlass() string {
	if m != nil {
		return m.Glass
	}
	return ""
}

func (m *Item) GetSquaremeter() float32 {
	if m != nil {
		return m.Squaremeter
	}
	return 0
}

func (m *Item) GetPriceunit() float32 {
	if m != nil {
		return m.Priceunit
	}
	return 0
}

func (m *Item) GetPriceset() float32 {
	if m != nil {
		return m.Priceset
	}
	return 0
}

func (m *Item) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Item) GetRemarks() string {
	if m != nil {
		return m.Remarks
	}
	return ""
}

func (m *Item) GetUser() *user.Item {
	if m != nil {
		return m.User
	}
	return nil
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
	return fileDescriptor_fda99000515d91e5, []int{1}
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
	proto.RegisterType((*Item)(nil), "pb.item.Item")
	proto.RegisterType((*DQ)(nil), "pb.item.DQ")
}

func init() {
	proto.RegisterFile("github.com/zhanglp92/rep/api/pb/item/item.proto", fileDescriptor_fda99000515d91e5)
}

var fileDescriptor_fda99000515d91e5 = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0xc6, 0x95, 0x34, 0xfd, 0x77, 0xa5, 0x1d, 0x2c, 0x84, 0x4e, 0x15, 0x43, 0xa8, 0x90, 0xc8,
	0x94, 0x48, 0x65, 0x62, 0x67, 0x61, 0x24, 0x6f, 0xe0, 0x24, 0xa7, 0xd6, 0x22, 0x89, 0x8d, 0xed,
	0x0c, 0xf4, 0xa5, 0x79, 0x05, 0xe4, 0x4b, 0xcb, 0x9f, 0x89, 0xc5, 0xba, 0xdf, 0xf7, 0x7d, 0x77,
	0x89, 0xee, 0xa0, 0x38, 0x28, 0x7f, 0x1c, 0xaa, 0xbc, 0xd6, 0x5d, 0x71, 0x3a, 0xca, 0xfe, 0xd0,
	0x9a, 0xa7, 0x7d, 0x61, 0xc9, 0x14, 0xd2, 0xa8, 0xc2, 0x54, 0x85, 0xf2, 0xd4, 0xf1, 0x93, 0x1b,
	0xab, 0xbd, 0x16, 0x73, 0x53, 0xe5, 0x01, 0xb7, 0xff, 0x76, 0x0e, 0x8e, 0x2c, 0x3f, 0x63, 0xe7,
	0xee, 0x33, 0x86, 0xe4, 0xc5, 0x53, 0x27, 0x36, 0x10, 0xab, 0x06, 0xa3, 0x34, 0xca, 0xa6, 0x65,
	0xac, 0x1a, 0x71, 0x03, 0xb3, 0x10, 0x53, 0x0d, 0xc6, 0xac, 0x9d, 0x49, 0x08, 0x48, 0xfc, 0x87,
	0x21, 0x9c, 0xa4, 0x51, 0xb6, 0x2c, 0xb9, 0x0e, 0xd9, 0x5a, 0xb7, 0x7a, 0xb0, 0x98, 0xb0, 0x7a,
	0x26, 0x71, 0x0f, 0x6b, 0x6d, 0xa8, 0x6f, 0x94, 0xa5, 0xda, 0x2b, 0xdd, 0xe3, 0x94, 0xed, 0xbf,
	0x62, 0x98, 0xe8, 0xd4, 0x89, 0x70, 0x36, 0x4e, 0x0c, 0xb5, 0xb8, 0x86, 0xa9, 0x3b, 0x4a, 0x43,
	0x38, 0x67, 0x71, 0x84, 0xa0, 0x1e, 0x5a, 0xe9, 0x1c, 0x2e, 0x46, 0x95, 0x41, 0xa4, 0xb0, 0x72,
	0xef, 0x83, 0xb4, 0xd4, 0x91, 0x27, 0x8b, 0xcb, 0x34, 0xca, 0xe2, 0xf2, 0xb7, 0x24, 0x6e, 0x61,
	0x69, 0xac, 0xaa, 0x69, 0xe8, 0x95, 0x47, 0x60, 0xff, 0x47, 0x10, 0x5b, 0x58, 0x30, 0x38, 0xf2,
	0xb8, 0x62, 0xf3, 0x9b, 0xc3, 0x17, 0xb9, 0xc6, 0x2b, 0x36, 0x46, 0x10, 0x08, 0x73, 0x4b, 0x9d,
	0xb4, 0x6f, 0x0e, 0xd7, 0xfc, 0x27, 0x17, 0x14, 0x77, 0x90, 0x84, 0x3d, 0xe1, 0x26, 0x8d, 0xb2,
	0xd5, 0x7e, 0x9d, 0x9b, 0x2a, 0xe7, 0x65, 0x87, 0x15, 0x97, 0x6c, 0xed, 0x1e, 0x20, 0x7e, 0x7e,
	0x0d, 0xc1, 0x46, 0x7a, 0x89, 0x51, 0x3a, 0xb9, 0x04, 0xf9, 0x9e, 0x63, 0x30, 0x58, 0xd5, 0x8c,
	0x2f, 0xf4, 0xf8, 0x15, 0x00, 0x00, 0xff, 0xff, 0x4c, 0x99, 0x23, 0xa5, 0x0e, 0x02, 0x00, 0x00,
}