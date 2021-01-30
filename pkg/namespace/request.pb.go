//go:generate  mcube enum -m -p
//
// Code generated by protoc-gen-go-ext. DO NOT EDIT.
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/namespace/pb/request.proto

package namespace

import (
	proto "github.com/golang/protobuf/proto"
	department "github.com/infraboard/keyauth/pkg/department"
	page "github.com/infraboard/mcube/pb/page"
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

// CreateNamespaceRequest 创建项目请求
type CreateNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 所属部门
	DepartmentId string `protobuf:"bytes,1,opt,name=department_id,json=departmentId,proto3" json:"department_id" bson:"department_id" validate:"required,lte=80"`
	// 项目名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name" bson:"name" validate:"required,lte=80"`
	// 项目描述图片
	Picture string `protobuf:"bytes,3,opt,name=picture,proto3" json:"picture,omitempty" bson:"picture"`
	// 禁用项目, 该项目所有人暂时都无法访问
	Enabled bool `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty" bson:"enabled"`
	// 项目所有者, PMO
	Owner string `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty" bson:"owner"`
	// 项目描述
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty" bson:"description"`
	// 补充的部门
	Department *department.Department `protobuf:"bytes,7,opt,name=department,proto3" json:"department,omitempty" bson:"-"`
}

func (x *CreateNamespaceRequest) Reset() {
	*x = CreateNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_namespace_pb_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNamespaceRequest) ProtoMessage() {}

func (x *CreateNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_namespace_pb_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNamespaceRequest.ProtoReflect.Descriptor instead.
func (*CreateNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_pkg_namespace_pb_request_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNamespaceRequest) GetDepartmentId() string {
	if x != nil {
		return x.DepartmentId
	}
	return ""
}

func (x *CreateNamespaceRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateNamespaceRequest) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *CreateNamespaceRequest) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *CreateNamespaceRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *CreateNamespaceRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateNamespaceRequest) GetDepartment() *department.Department {
	if x != nil {
		return x.Department
	}
	return nil
}

// QueryNamespaceRequest 查询应用列表
type QueryNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page              *page.PageRequest `protobuf:"bytes,1,opt,name=page,proto3" json:"page"`
	DepartmentId      string            `protobuf:"bytes,2,opt,name=department_id,json=departmentId,proto3" json:"department_id"`
	WithSubDepartment bool              `protobuf:"varint,3,opt,name=with_sub_department,json=withSubDepartment,proto3" json:"with_sub_department"`
	WithDepartment    bool              `protobuf:"varint,4,opt,name=with_department,json=withDepartment,proto3" json:"with_department"`
}

func (x *QueryNamespaceRequest) Reset() {
	*x = QueryNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_namespace_pb_request_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryNamespaceRequest) ProtoMessage() {}

func (x *QueryNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_namespace_pb_request_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryNamespaceRequest.ProtoReflect.Descriptor instead.
func (*QueryNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_pkg_namespace_pb_request_proto_rawDescGZIP(), []int{1}
}

func (x *QueryNamespaceRequest) GetPage() *page.PageRequest {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *QueryNamespaceRequest) GetDepartmentId() string {
	if x != nil {
		return x.DepartmentId
	}
	return ""
}

func (x *QueryNamespaceRequest) GetWithSubDepartment() bool {
	if x != nil {
		return x.WithSubDepartment
	}
	return false
}

func (x *QueryNamespaceRequest) GetWithDepartment() bool {
	if x != nil {
		return x.WithDepartment
	}
	return false
}

// DescriptNamespaceRequest 查询应用详情
type DescriptNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	WithDepartment bool   `protobuf:"varint,2,opt,name=with_department,json=withDepartment,proto3" json:"with_department"`
}

func (x *DescriptNamespaceRequest) Reset() {
	*x = DescriptNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_namespace_pb_request_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DescriptNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescriptNamespaceRequest) ProtoMessage() {}

func (x *DescriptNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_namespace_pb_request_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescriptNamespaceRequest.ProtoReflect.Descriptor instead.
func (*DescriptNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_pkg_namespace_pb_request_proto_rawDescGZIP(), []int{2}
}

func (x *DescriptNamespaceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DescriptNamespaceRequest) GetWithDepartment() bool {
	if x != nil {
		return x.WithDepartment
	}
	return false
}

// DeleteNamespaceRequest todo
type DeleteNamespaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DeleteNamespaceRequest) Reset() {
	*x = DeleteNamespaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_namespace_pb_request_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNamespaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNamespaceRequest) ProtoMessage() {}

func (x *DeleteNamespaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_namespace_pb_request_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNamespaceRequest.ProtoReflect.Descriptor instead.
func (*DeleteNamespaceRequest) Descriptor() ([]byte, []int) {
	return file_pkg_namespace_pb_request_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteNamespaceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_pkg_namespace_pb_request_proto protoreflect.FileDescriptor

var file_pkg_namespace_pb_request_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f,
	0x70, 0x62, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x1a, 0x22, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x2f, 0x70, 0x62, 0x2f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2d, 0x65, 0x78, 0x74,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x74,
	0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2f,
	0x6d, 0x63, 0x75, 0x62, 0x65, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x61, 0x67, 0x65, 0x2f, 0x70, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x04, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x6f, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x4a, 0xc2, 0xde, 0x1f, 0x46,
	0x0a, 0x44, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2c, 0x6c,
	0x74, 0x65, 0x3d, 0x38, 0x30, 0x22, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x4c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x38, 0xc2, 0xde, 0x1f, 0x34, 0x0a, 0x32, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x3a, 0x22, 0x72, 0x65, 0x71, 0x75,
	0x69, 0x72, 0x65, 0x64, 0x2c, 0x6c, 0x74, 0x65, 0x3d, 0x38, 0x30, 0x22, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x47, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x2d, 0xc2, 0xde, 0x1f, 0x29, 0x0a, 0x27, 0x62, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x47, 0x0a, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x2d, 0xc2, 0xde,
	0x1f, 0x29, 0x0a, 0x27, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65,
	0x64, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64,
	0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x07, 0x65, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x64, 0x12, 0x3f, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x29, 0xc2, 0xde, 0x1f, 0x25, 0x0a, 0x23, 0x62, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6f, 0x77,
	0x6e, 0x65, 0x72, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x57, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x35, 0xc2, 0xde, 0x1f, 0x31,
	0x0a, 0x2f, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x6a,
	0x0a, 0x0a, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x64, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65,
	0x6e, 0x74, 0x42, 0x2a, 0xc2, 0xde, 0x1f, 0x26, 0x0a, 0x24, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x2d, 0x22, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d,
	0x65, 0x6e, 0x74, 0x2c, 0x6f, 0x6d, 0x69, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x52, 0x0a,
	0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xab, 0x02, 0x0a, 0x15, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x61, 0x67, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x11, 0xc2, 0xde, 0x1f, 0x0d, 0x0a, 0x0b, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x70, 0x61, 0x67, 0x65, 0x22, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x3f,
	0x0a, 0x0d, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1a, 0xc2, 0xde, 0x1f, 0x16, 0x0a, 0x14, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x22, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x50, 0x0a, 0x13, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x73, 0x75, 0x62, 0x5f, 0x64, 0x65, 0x70, 0x61,
	0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42, 0x20, 0xc2, 0xde,
	0x1f, 0x1c, 0x0a, 0x1a, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x73,
	0x75, 0x62, 0x5f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x52, 0x11,
	0x77, 0x69, 0x74, 0x68, 0x53, 0x75, 0x62, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x45, 0x0a, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x42, 0x1c, 0xc2, 0xde, 0x1f, 0x18,
	0x0a, 0x16, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x64, 0x65, 0x70,
	0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x52, 0x0e, 0x77, 0x69, 0x74, 0x68, 0x44, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x18, 0x44, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0f, 0xc2, 0xde, 0x1f, 0x0b, 0x0a, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69,
	0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x12, 0x45, 0x0a, 0x0f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x64,
	0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x42,
	0x1c, 0xc2, 0xde, 0x1f, 0x18, 0x0a, 0x16, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x77, 0x69, 0x74,
	0x68, 0x5f, 0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x52, 0x0e, 0x77,
	0x69, 0x74, 0x68, 0x44, 0x65, 0x70, 0x61, 0x72, 0x74, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x39, 0x0a,
	0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0f, 0xc2, 0xde, 0x1f, 0x0b, 0x0a, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x69, 0x64, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x2f, 0x6b, 0x65, 0x79, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_namespace_pb_request_proto_rawDescOnce sync.Once
	file_pkg_namespace_pb_request_proto_rawDescData = file_pkg_namespace_pb_request_proto_rawDesc
)

func file_pkg_namespace_pb_request_proto_rawDescGZIP() []byte {
	file_pkg_namespace_pb_request_proto_rawDescOnce.Do(func() {
		file_pkg_namespace_pb_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_namespace_pb_request_proto_rawDescData)
	})
	return file_pkg_namespace_pb_request_proto_rawDescData
}

var file_pkg_namespace_pb_request_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_namespace_pb_request_proto_goTypes = []interface{}{
	(*CreateNamespaceRequest)(nil),   // 0: keyauth.namespace.CreateNamespaceRequest
	(*QueryNamespaceRequest)(nil),    // 1: keyauth.namespace.QueryNamespaceRequest
	(*DescriptNamespaceRequest)(nil), // 2: keyauth.namespace.DescriptNamespaceRequest
	(*DeleteNamespaceRequest)(nil),   // 3: keyauth.namespace.DeleteNamespaceRequest
	(*department.Department)(nil),    // 4: keyauth.department.Department
	(*page.PageRequest)(nil),         // 5: page.PageRequest
}
var file_pkg_namespace_pb_request_proto_depIdxs = []int32{
	4, // 0: keyauth.namespace.CreateNamespaceRequest.department:type_name -> keyauth.department.Department
	5, // 1: keyauth.namespace.QueryNamespaceRequest.page:type_name -> page.PageRequest
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_namespace_pb_request_proto_init() }
func file_pkg_namespace_pb_request_proto_init() {
	if File_pkg_namespace_pb_request_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_namespace_pb_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNamespaceRequest); i {
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
		file_pkg_namespace_pb_request_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryNamespaceRequest); i {
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
		file_pkg_namespace_pb_request_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DescriptNamespaceRequest); i {
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
		file_pkg_namespace_pb_request_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNamespaceRequest); i {
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
			RawDescriptor: file_pkg_namespace_pb_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_namespace_pb_request_proto_goTypes,
		DependencyIndexes: file_pkg_namespace_pb_request_proto_depIdxs,
		MessageInfos:      file_pkg_namespace_pb_request_proto_msgTypes,
	}.Build()
	File_pkg_namespace_pb_request_proto = out.File
	file_pkg_namespace_pb_request_proto_rawDesc = nil
	file_pkg_namespace_pb_request_proto_goTypes = nil
	file_pkg_namespace_pb_request_proto_depIdxs = nil
}
