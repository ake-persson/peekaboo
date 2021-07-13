// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: pb/v1/resources/filesystem.proto

package resources

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Filesystem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filesystem string `protobuf:"bytes,1,opt,name=filesystem,proto3" json:"filesystem" csv:"filesystem" yaml:"filesystem"`
	Type       string `protobuf:"bytes,2,opt,name=type,proto3" json:"type" csv:"type" yaml:"type"`
	// KB as in 1024 bytes
	SizeKb        uint64   `protobuf:"varint,3,opt,name=size_kb,json=sizeKb,proto3" json:"size_kb" csv:"size_kb" yaml:"size_kb"`
	UsedKb        uint64   `protobuf:"varint,4,opt,name=used_kb,json=usedKb,proto3" json:"used_kb" csv:"used_kb" yaml:"used_kb"`
	FreeKb        uint64   `protobuf:"varint,5,opt,name=free_kb,json=freeKb,proto3" json:"free_kb" csv:"free_kb" yaml:"free_kb"`
	UsedPct       float32  `protobuf:"fixed32,6,opt,name=used_pct,json=usedPct,proto3" json:"used_pct" csv:"used_pct" yaml:"used_pct"`
	Inodes        uint64   `protobuf:"varint,7,opt,name=inodes,proto3" json:"inodes" csv:"inodes" yaml:"inodes"`
	InodesUsed    uint64   `protobuf:"varint,8,opt,name=inodes_used,json=inodesUsed,proto3" json:"inodes_used" csv:"inodes_used" yaml:"inodes_used"`
	InodesFree    uint64   `protobuf:"varint,9,opt,name=inodes_free,json=inodesFree,proto3" json:"inodes_free" csv:"inodes_free" yaml:"inodes_free"`
	InodesUsedPct float32  `protobuf:"fixed32,10,opt,name=inodes_used_pct,json=inodesUsedPct,proto3" json:"inodes_used_pct" csv:"inodes_used_pct" yaml:"inodes_used_pct"`
	MountedOn     string   `protobuf:"bytes,11,opt,name=mounted_on,json=mountedOn,proto3" json:"mounted_on" csv:"mounted_on" yaml:"mounted_on"`
	MountOptions  []string `protobuf:"bytes,12,rep,name=mount_options,json=mountOptions,proto3" json:"mount_options" csv:"mount_options" yaml:"mount_options"`
	IsLocal       bool     `protobuf:"varint,13,opt,name=is_local,json=isLocal,proto3" json:"is_local" csv:"is_local" yaml:"is_local"`
	IsAutomounted bool     `protobuf:"varint,14,opt,name=is_automounted,json=isAutomounted,proto3" json:"is_automounted" csv:"is_automounted" yaml:"is_automounted"`
}

func (x *Filesystem) Reset() {
	*x = Filesystem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_v1_resources_filesystem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filesystem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filesystem) ProtoMessage() {}

func (x *Filesystem) ProtoReflect() protoreflect.Message {
	mi := &file_pb_v1_resources_filesystem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Filesystem.ProtoReflect.Descriptor instead.
func (*Filesystem) Descriptor() ([]byte, []int) {
	return file_pb_v1_resources_filesystem_proto_rawDescGZIP(), []int{0}
}

func (x *Filesystem) GetFilesystem() string {
	if x != nil {
		return x.Filesystem
	}
	return ""
}

func (x *Filesystem) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Filesystem) GetSizeKb() uint64 {
	if x != nil {
		return x.SizeKb
	}
	return 0
}

func (x *Filesystem) GetUsedKb() uint64 {
	if x != nil {
		return x.UsedKb
	}
	return 0
}

func (x *Filesystem) GetFreeKb() uint64 {
	if x != nil {
		return x.FreeKb
	}
	return 0
}

func (x *Filesystem) GetUsedPct() float32 {
	if x != nil {
		return x.UsedPct
	}
	return 0
}

func (x *Filesystem) GetInodes() uint64 {
	if x != nil {
		return x.Inodes
	}
	return 0
}

func (x *Filesystem) GetInodesUsed() uint64 {
	if x != nil {
		return x.InodesUsed
	}
	return 0
}

func (x *Filesystem) GetInodesFree() uint64 {
	if x != nil {
		return x.InodesFree
	}
	return 0
}

func (x *Filesystem) GetInodesUsedPct() float32 {
	if x != nil {
		return x.InodesUsedPct
	}
	return 0
}

func (x *Filesystem) GetMountedOn() string {
	if x != nil {
		return x.MountedOn
	}
	return ""
}

func (x *Filesystem) GetMountOptions() []string {
	if x != nil {
		return x.MountOptions
	}
	return nil
}

func (x *Filesystem) GetIsLocal() bool {
	if x != nil {
		return x.IsLocal
	}
	return false
}

func (x *Filesystem) GetIsAutomounted() bool {
	if x != nil {
		return x.IsAutomounted
	}
	return false
}

var File_pb_v1_resources_filesystem_proto protoreflect.FileDescriptor

var file_pb_v1_resources_filesystem_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x15, 0x70, 0x65, 0x65, 0x6b, 0x61, 0x62, 0x6f, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0xae, 0x03, 0x0a, 0x0a, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x73, 0x69, 0x7a, 0x65, 0x5f, 0x6b, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73,
	0x69, 0x7a, 0x65, 0x4b, 0x62, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x6b, 0x62,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x64, 0x4b, 0x62, 0x12, 0x17,
	0x0a, 0x07, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x6b, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x66, 0x72, 0x65, 0x65, 0x4b, 0x62, 0x12, 0x19, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x64, 0x5f,
	0x70, 0x63, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x75, 0x73, 0x65, 0x64, 0x50,
	0x63, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x06, 0x69, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0a, 0x69, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x55, 0x73, 0x65, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x5f, 0x66, 0x72, 0x65, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x69, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x46, 0x72, 0x65, 0x65, 0x12, 0x26, 0x0a, 0x0f,
	0x69, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x70, 0x63, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0d, 0x69, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x55, 0x73, 0x65,
	0x64, 0x50, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x5f,
	0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x64, 0x4f, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6c,
	0x6f, 0x63, 0x61, 0x6c, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x4c, 0x6f,
	0x63, 0x61, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x69, 0x73, 0x41,
	0x75, 0x74, 0x6f, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x64, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6b, 0x65, 0x2d, 0x70, 0x65, 0x72,
	0x73, 0x73, 0x6f, 0x6e, 0x2f, 0x70, 0x65, 0x65, 0x6b, 0x61, 0x62, 0x6f, 0x6f, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_pb_v1_resources_filesystem_proto_rawDescOnce sync.Once
	file_pb_v1_resources_filesystem_proto_rawDescData = file_pb_v1_resources_filesystem_proto_rawDesc
)

func file_pb_v1_resources_filesystem_proto_rawDescGZIP() []byte {
	file_pb_v1_resources_filesystem_proto_rawDescOnce.Do(func() {
		file_pb_v1_resources_filesystem_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_v1_resources_filesystem_proto_rawDescData)
	})
	return file_pb_v1_resources_filesystem_proto_rawDescData
}

var file_pb_v1_resources_filesystem_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pb_v1_resources_filesystem_proto_goTypes = []interface{}{
	(*Filesystem)(nil), // 0: peekaboo.v1.resources.Filesystem
}
var file_pb_v1_resources_filesystem_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_v1_resources_filesystem_proto_init() }
func file_pb_v1_resources_filesystem_proto_init() {
	if File_pb_v1_resources_filesystem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_v1_resources_filesystem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filesystem); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_v1_resources_filesystem_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_v1_resources_filesystem_proto_goTypes,
		DependencyIndexes: file_pb_v1_resources_filesystem_proto_depIdxs,
		MessageInfos:      file_pb_v1_resources_filesystem_proto_msgTypes,
	}.Build()
	File_pb_v1_resources_filesystem_proto = out.File
	file_pb_v1_resources_filesystem_proto_rawDesc = nil
	file_pb_v1_resources_filesystem_proto_goTypes = nil
	file_pb_v1_resources_filesystem_proto_depIdxs = nil
}
