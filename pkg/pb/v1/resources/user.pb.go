// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/v1/resources/user.proto

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

type User struct {
	// This is the user's login name. It should not contain capital letters.
	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name"`
	// The privileged root login account (superuser) has the user ID 0.
	Uid uint64 `protobuf:"varint,3,opt,name=uid,proto3" json:"uid"`
	// This is the numeric primary group ID for this user. (Additional groups for the user are defined in the system group file; see group(5)).
	Gid uint64 `protobuf:"varint,4,opt,name=gid,proto3" json:"gid"`
	// Signed UID for Mac OS X.
	UidSigned int64 `protobuf:"varint,5,opt,name=uid_signed,json=uidSigned,proto3" json:"uid_signed"`
	// Signed GID for Mac OS X.
	GidSigned int64 `protobuf:"varint,6,opt,name=gid_signed,json=gidSigned,proto3" json:"gid_signed"`
	// Description.
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description"`
	// Home directory.
	Directory string `protobuf:"bytes,8,opt,name=directory,proto3" json:"directory"`
	// This is the program to run at login (if empty, use /bin/sh). If set to a nonexistent executable, the user will be unable to login through login(1). The value in this field is used to set the SHELL environment variable.
	Shell                string   `protobuf:"bytes,9,opt,name=shell,proto3" json:"shell"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_cefae016e8f0da92, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *User) GetGid() uint64 {
	if m != nil {
		return m.Gid
	}
	return 0
}

func (m *User) GetUidSigned() int64 {
	if m != nil {
		return m.UidSigned
	}
	return 0
}

func (m *User) GetGidSigned() int64 {
	if m != nil {
		return m.GidSigned
	}
	return 0
}

func (m *User) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *User) GetDirectory() string {
	if m != nil {
		return m.Directory
	}
	return ""
}

func (m *User) GetShell() string {
	if m != nil {
		return m.Shell
	}
	return ""
}

type UserList struct {
	Users                []*User  `protobuf:"bytes,1,rep,name=users,proto3" json:"users"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserList) Reset()         { *m = UserList{} }
func (m *UserList) String() string { return proto.CompactTextString(m) }
func (*UserList) ProtoMessage()    {}
func (*UserList) Descriptor() ([]byte, []int) {
	return fileDescriptor_cefae016e8f0da92, []int{1}
}

func (m *UserList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserList.Unmarshal(m, b)
}
func (m *UserList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserList.Marshal(b, m, deterministic)
}
func (m *UserList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserList.Merge(m, src)
}
func (m *UserList) XXX_Size() int {
	return xxx_messageInfo_UserList.Size(m)
}
func (m *UserList) XXX_DiscardUnknown() {
	xxx_messageInfo_UserList.DiscardUnknown(m)
}

var xxx_messageInfo_UserList proto.InternalMessageInfo

func (m *UserList) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "peekaboo.v1.resources.User")
	proto.RegisterType((*UserList)(nil), "peekaboo.v1.resources.UserList")
}

func init() { proto.RegisterFile("pb/v1/resources/user.proto", fileDescriptor_cefae016e8f0da92) }

var fileDescriptor_cefae016e8f0da92 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x89, 0xed, 0xe6, 0xfa, 0x76, 0x91, 0xa0, 0x10, 0xa6, 0x42, 0xd9, 0xa9, 0x17, 0x13,
	0xa6, 0x47, 0x11, 0xd1, 0xb3, 0xa7, 0x8a, 0x17, 0x2f, 0xd2, 0x36, 0x21, 0x7b, 0xac, 0x6b, 0x4a,
	0xd2, 0x0c, 0xfc, 0xd2, 0x7e, 0x06, 0x49, 0x8a, 0x9d, 0x88, 0xb7, 0xf7, 0xff, 0xfd, 0xfe, 0x10,
	0x5e, 0x1e, 0xac, 0xfa, 0x5a, 0x1c, 0x36, 0xc2, 0x2a, 0x67, 0xbc, 0x6d, 0x94, 0x13, 0xde, 0x29,
	0xcb, 0x7b, 0x6b, 0x06, 0x43, 0x2f, 0x7a, 0xa5, 0x76, 0x55, 0x6d, 0x0c, 0x3f, 0x6c, 0xf8, 0xd4,
	0x58, 0x7f, 0x11, 0x48, 0xdf, 0x9c, 0xb2, 0x74, 0x05, 0x8b, 0xd0, 0xee, 0xaa, 0xbd, 0x62, 0x24,
	0x27, 0x45, 0x56, 0x4e, 0x99, 0x52, 0x48, 0x23, 0x3f, 0x89, 0x3c, 0xce, 0xf4, 0x0c, 0x12, 0x8f,
	0x92, 0x25, 0x39, 0x29, 0xd2, 0x32, 0x8c, 0x81, 0x68, 0x94, 0x2c, 0x1d, 0x89, 0x46, 0x49, 0xaf,
	0x01, 0x3c, 0xca, 0x0f, 0x87, 0xba, 0x53, 0x92, 0xcd, 0x72, 0x52, 0x24, 0x65, 0xe6, 0x51, 0xbe,
	0x46, 0x10, 0xb4, 0x3e, 0xea, 0xf9, 0xa8, 0xf5, 0xa4, 0x73, 0x58, 0x4a, 0xe5, 0x1a, 0x8b, 0xfd,
	0x80, 0xa6, 0x63, 0xa7, 0xf1, 0xf1, 0xdf, 0x88, 0x5e, 0x41, 0x26, 0xd1, 0xaa, 0x66, 0x30, 0xf6,
	0x93, 0x2d, 0xa2, 0x3f, 0x02, 0x7a, 0x0e, 0x33, 0xb7, 0x55, 0x6d, 0xcb, 0xb2, 0x68, 0xc6, 0xb0,
	0x7e, 0x80, 0x45, 0xd8, 0xf7, 0x05, 0xdd, 0x40, 0x37, 0x30, 0x0b, 0x3b, 0x3a, 0x46, 0xf2, 0xa4,
	0x58, 0xde, 0x5e, 0xf2, 0x7f, 0xff, 0x88, 0x87, 0x7e, 0x39, 0x36, 0x9f, 0x9f, 0xde, 0x1f, 0x35,
	0x0e, 0x5b, 0x5f, 0xf3, 0xc6, 0xec, 0xc5, 0x4f, 0xff, 0xa6, 0xad, 0x6a, 0x37, 0x25, 0xd1, 0xef,
	0xb4, 0xf8, 0x73, 0x8a, 0xfb, 0x69, 0xaa, 0xe7, 0xf1, 0x20, 0x77, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0xd2, 0x8e, 0xd2, 0x8a, 0xae, 0x01, 0x00, 0x00,
}
