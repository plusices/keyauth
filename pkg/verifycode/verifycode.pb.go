//go:generate  mcube enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.1
// source: pkg/verifycode/pb/verifycode.proto

package verifycode

import (
	proto "github.com/golang/protobuf/proto"
	_ "github.com/infraboard/protoc-gen-go-ext/extension/tag"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Code todo
type Code struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=id,proto3" json:"id" bson:"_id"`
	Username      string `protobuf:"bytes,2,opt,name=username,proto3" json:"username" validate:"required"`
	Number        string `protobuf:"bytes,3,opt,name=number,proto3" json:"number" validate:"required"`
	IssueAt       int64  `protobuf:"varint,4,opt,name=issue_at,json=issueAt,proto3" json:"issue_at" bson:"issue_at"`
	ExpiredMinite uint32 `protobuf:"varint,5,opt,name=expired_minite,json=expiredMinite,proto3" json:"expired_minite" bson:"expired_minite"`
}

func (x *Code) Reset() {
	*x = Code{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_verifycode_pb_verifycode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Code) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Code) ProtoMessage() {}

func (x *Code) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_verifycode_pb_verifycode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Code.ProtoReflect.Descriptor instead.
func (*Code) Descriptor() ([]byte, []int) {
	return file_pkg_verifycode_pb_verifycode_proto_rawDescGZIP(), []int{0}
}

func (x *Code) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Code) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Code) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Code) GetIssueAt() int64 {
	if x != nil {
		return x.IssueAt
	}
	return 0
}

func (x *Code) GetExpiredMinite() uint32 {
	if x != nil {
		return x.ExpiredMinite
	}
	return 0
}

var File_pkg_verifycode_pb_verifycode_proto protoreflect.FileDescriptor

var file_pkg_verifycode_pb_verifycode_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x6b, 0x67, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x63, 0x6f, 0x64, 0x65,
	0x2f, 0x70, 0x62, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x63, 0x6f, 0x64, 0x65, 0x1a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x65, 0x78,
	0x74, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x67, 0x2f,
	0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd6, 0x02, 0x0a, 0x04, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x2a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1a,
	0xc2, 0xde, 0x1f, 0x16, 0x0a, 0x14, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x5f, 0x69, 0x64, 0x22,
	0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x45,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x29, 0xc2, 0xde, 0x1f, 0x25, 0x0a, 0x23, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3f, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x27, 0xc2, 0xde, 0x1f, 0x23, 0x0a, 0x21, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x22, 0x52, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x40, 0x0a, 0x08, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x25, 0xc2, 0xde, 0x1f, 0x21, 0x0a, 0x1f,
	0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x61, 0x74, 0x22, 0x20,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x61, 0x74, 0x22, 0x52,
	0x07, 0x69, 0x73, 0x73, 0x75, 0x65, 0x41, 0x74, 0x12, 0x58, 0x0a, 0x0e, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x65, 0x64, 0x5f, 0x6d, 0x69, 0x6e, 0x69, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x31, 0xc2, 0xde, 0x1f, 0x2d, 0x0a, 0x2b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x6d, 0x69, 0x6e, 0x69, 0x74, 0x65, 0x22, 0x20, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x6d, 0x69, 0x6e, 0x69,
	0x74, 0x65, 0x22, 0x52, 0x0d, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x4d, 0x69, 0x6e, 0x69,
	0x74, 0x65, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x63, 0x6f,
	0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_verifycode_pb_verifycode_proto_rawDescOnce sync.Once
	file_pkg_verifycode_pb_verifycode_proto_rawDescData = file_pkg_verifycode_pb_verifycode_proto_rawDesc
)

func file_pkg_verifycode_pb_verifycode_proto_rawDescGZIP() []byte {
	file_pkg_verifycode_pb_verifycode_proto_rawDescOnce.Do(func() {
		file_pkg_verifycode_pb_verifycode_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_verifycode_pb_verifycode_proto_rawDescData)
	})
	return file_pkg_verifycode_pb_verifycode_proto_rawDescData
}

var file_pkg_verifycode_pb_verifycode_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_verifycode_pb_verifycode_proto_goTypes = []interface{}{
	(*Code)(nil), // 0: keyauth.verifycode.Code
}
var file_pkg_verifycode_pb_verifycode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_verifycode_pb_verifycode_proto_init() }
func file_pkg_verifycode_pb_verifycode_proto_init() {
	if File_pkg_verifycode_pb_verifycode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_verifycode_pb_verifycode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Code); i {
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
			RawDescriptor: file_pkg_verifycode_pb_verifycode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_verifycode_pb_verifycode_proto_goTypes,
		DependencyIndexes: file_pkg_verifycode_pb_verifycode_proto_depIdxs,
		MessageInfos:      file_pkg_verifycode_pb_verifycode_proto_msgTypes,
	}.Build()
	File_pkg_verifycode_pb_verifycode_proto = out.File
	file_pkg_verifycode_pb_verifycode_proto_rawDesc = nil
	file_pkg_verifycode_pb_verifycode_proto_goTypes = nil
	file_pkg_verifycode_pb_verifycode_proto_depIdxs = nil
}
