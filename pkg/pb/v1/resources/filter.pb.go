// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/v1/resources/filter.proto

package resources

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

type Operator int32

const (
	Operator_EQ   Operator = 0
	Operator_NEQ  Operator = 1
	Operator_GT   Operator = 2
	Operator_GTE  Operator = 3
	Operator_LT   Operator = 4
	Operator_LTE  Operator = 5
	Operator_IN   Operator = 6
	Operator_LIKE Operator = 7
)

var Operator_name = map[int32]string{
	0: "EQ",
	1: "NEQ",
	2: "GT",
	3: "GTE",
	4: "LT",
	5: "LTE",
	6: "IN",
	7: "LIKE",
}

var Operator_value = map[string]int32{
	"EQ":   0,
	"NEQ":  1,
	"GT":   2,
	"GTE":  3,
	"LT":   4,
	"LTE":  5,
	"IN":   6,
	"LIKE": 7,
}

func (x Operator) String() string {
	return proto.EnumName(Operator_name, int32(x))
}

func (Operator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2c35b423a7ed9d39, []int{0}
}

type Match struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value"`
	Values               []string `protobuf:"bytes,3,rep,name=values,proto3" json:"values"`
	Operator             Operator `protobuf:"varint,4,opt,name=operator,proto3,enum=peekaboo.v1.resources.Operator" json:"operator"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Match) Reset()         { *m = Match{} }
func (m *Match) String() string { return proto.CompactTextString(m) }
func (*Match) ProtoMessage()    {}
func (*Match) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c35b423a7ed9d39, []int{0}
}

func (m *Match) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Match.Unmarshal(m, b)
}
func (m *Match) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Match.Marshal(b, m, deterministic)
}
func (m *Match) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Match.Merge(m, src)
}
func (m *Match) XXX_Size() int {
	return xxx_messageInfo_Match.Size(m)
}
func (m *Match) XXX_DiscardUnknown() {
	xxx_messageInfo_Match.DiscardUnknown(m)
}

var xxx_messageInfo_Match proto.InternalMessageInfo

func (m *Match) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *Match) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Match) GetValues() []string {
	if m != nil {
		return m.Values
	}
	return nil
}

func (m *Match) GetOperator() Operator {
	if m != nil {
		return m.Operator
	}
	return Operator_EQ
}

type Sort struct {
	Field                string   `protobuf:"bytes,1,opt,name=field,proto3" json:"field"`
	Desc                 bool     `protobuf:"varint,2,opt,name=desc,proto3" json:"desc"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sort) Reset()         { *m = Sort{} }
func (m *Sort) String() string { return proto.CompactTextString(m) }
func (*Sort) ProtoMessage()    {}
func (*Sort) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c35b423a7ed9d39, []int{1}
}

func (m *Sort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sort.Unmarshal(m, b)
}
func (m *Sort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sort.Marshal(b, m, deterministic)
}
func (m *Sort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sort.Merge(m, src)
}
func (m *Sort) XXX_Size() int {
	return xxx_messageInfo_Sort.Size(m)
}
func (m *Sort) XXX_DiscardUnknown() {
	xxx_messageInfo_Sort.DiscardUnknown(m)
}

var xxx_messageInfo_Sort proto.InternalMessageInfo

func (m *Sort) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *Sort) GetDesc() bool {
	if m != nil {
		return m.Desc
	}
	return false
}

type Filter struct {
	Fields               []string `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields"`
	Matches              []*Match `protobuf:"bytes,2,rep,name=matches,proto3" json:"matches"`
	Sorting              []*Sort  `protobuf:"bytes,3,rep,name=sorting,proto3" json:"sorting"`
	Limit                int32    `protobuf:"varint,4,opt,name=limit,proto3" json:"limit"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c35b423a7ed9d39, []int{2}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetFields() []string {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *Filter) GetMatches() []*Match {
	if m != nil {
		return m.Matches
	}
	return nil
}

func (m *Filter) GetSorting() []*Sort {
	if m != nil {
		return m.Sorting
	}
	return nil
}

func (m *Filter) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func init() {
	proto.RegisterEnum("peekaboo.v1.resources.Operator", Operator_name, Operator_value)
	proto.RegisterType((*Match)(nil), "peekaboo.v1.resources.Match")
	proto.RegisterType((*Sort)(nil), "peekaboo.v1.resources.Sort")
	proto.RegisterType((*Filter)(nil), "peekaboo.v1.resources.Filter")
}

func init() { proto.RegisterFile("pb/v1/resources/filter.proto", fileDescriptor_2c35b423a7ed9d39) }

var fileDescriptor_2c35b423a7ed9d39 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x5f, 0x4b, 0xf3, 0x30,
	0x18, 0xc5, 0xdf, 0xfe, 0xef, 0x32, 0x78, 0x09, 0x41, 0xa5, 0xe0, 0xc0, 0xb2, 0xab, 0x22, 0xd8,
	0xba, 0x89, 0xde, 0xec, 0x42, 0x14, 0xea, 0x18, 0xce, 0x8d, 0xc5, 0x5e, 0x79, 0xd7, 0x76, 0xd9,
	0x56, 0xd6, 0x99, 0x92, 0x64, 0xfb, 0x0c, 0x7e, 0x11, 0xbf, 0xa7, 0x24, 0x5d, 0x7b, 0x21, 0xee,
	0xee, 0x39, 0x87, 0xf3, 0x90, 0xdf, 0x09, 0x0f, 0xe8, 0x55, 0x59, 0x74, 0x18, 0x44, 0x8c, 0x70,
	0xba, 0x67, 0x39, 0xe1, 0xd1, 0xaa, 0x28, 0x05, 0x61, 0x61, 0xc5, 0xa8, 0xa0, 0xe8, 0xbc, 0x22,
	0x64, 0x9b, 0x66, 0x94, 0x86, 0x87, 0x41, 0xd8, 0x66, 0xfa, 0x5f, 0x1a, 0xb0, 0xde, 0x52, 0x91,
	0x6f, 0xd0, 0x19, 0xb0, 0x56, 0x05, 0x29, 0x97, 0x9e, 0xe6, 0x6b, 0x41, 0x07, 0xd7, 0x42, 0xba,
	0x87, 0xb4, 0xdc, 0x13, 0x4f, 0xaf, 0x5d, 0x25, 0xd0, 0x05, 0xb0, 0xd5, 0xc0, 0x3d, 0xc3, 0x37,
	0x82, 0x0e, 0x3e, 0x2a, 0x34, 0x02, 0x2e, 0xad, 0x08, 0x4b, 0x05, 0x65, 0x9e, 0xe9, 0x6b, 0xc1,
	0xff, 0xe1, 0x55, 0xf8, 0xe7, 0xbb, 0xe1, 0xfc, 0x18, 0xc3, 0xed, 0x42, 0xff, 0x16, 0x98, 0xef,
	0x94, 0x89, 0x13, 0x20, 0x08, 0x98, 0x4b, 0xc2, 0x73, 0xc5, 0xe1, 0x62, 0x35, 0xf7, 0xbf, 0x35,
	0x60, 0xbf, 0xa8, 0x92, 0x92, 0x48, 0xe5, 0xb8, 0xa7, 0xd5, 0x44, 0xb5, 0x42, 0x0f, 0xc0, 0xd9,
	0xc9, 0x7a, 0x84, 0x7b, 0xba, 0x6f, 0x04, 0xdd, 0x61, 0xef, 0x04, 0x90, 0xfa, 0x04, 0xdc, 0x84,
	0xd1, 0x3d, 0x70, 0x38, 0x65, 0xa2, 0xf8, 0x5c, 0xab, 0x8a, 0xdd, 0xe1, 0xe5, 0x89, 0x3d, 0x89,
	0x8c, 0x9b, 0xac, 0x64, 0x2f, 0x8b, 0x5d, 0x21, 0x54, 0x7b, 0x0b, 0xd7, 0xe2, 0x7a, 0x0e, 0xdc,
	0xa6, 0x2f, 0xb2, 0x81, 0x1e, 0x2f, 0xe0, 0x3f, 0xe4, 0x00, 0x63, 0x16, 0x2f, 0xa0, 0x26, 0x8d,
	0x71, 0x02, 0x75, 0x69, 0x8c, 0x93, 0x18, 0x1a, 0xd2, 0x98, 0x26, 0xd0, 0x94, 0xc6, 0x34, 0x89,
	0xa1, 0x25, 0x8d, 0xc9, 0x0c, 0xda, 0xc8, 0x05, 0xe6, 0x74, 0xf2, 0x1a, 0x43, 0xe7, 0xf9, 0xe9,
	0xe3, 0x71, 0x5d, 0x88, 0xcd, 0x3e, 0x0b, 0x73, 0xba, 0x8b, 0x1a, 0xb0, 0x9b, 0x32, 0xcd, 0x78,
	0xab, 0xa2, 0x6a, 0xbb, 0x8e, 0x7e, 0x9d, 0xc4, 0xa8, 0x9d, 0x32, 0x5b, 0x9d, 0xc5, 0xdd, 0x4f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xa5, 0xa4, 0xc9, 0x36, 0x02, 0x00, 0x00,
}