// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.29.4
// source: proto/perm-service/perm-service.proto

package perm_service

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

type CheckUserEnableReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	GenTokenAt string `protobuf:"bytes,2,opt,name=genTokenAt,proto3" json:"genTokenAt,omitempty"`
}

func (x *CheckUserEnableReq) Reset() {
	*x = CheckUserEnableReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_perm_service_perm_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserEnableReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserEnableReq) ProtoMessage() {}

func (x *CheckUserEnableReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_perm_service_perm_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserEnableReq.ProtoReflect.Descriptor instead.
func (*CheckUserEnableReq) Descriptor() ([]byte, []int) {
	return file_proto_perm_service_perm_service_proto_rawDescGZIP(), []int{0}
}

func (x *CheckUserEnableReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CheckUserEnableReq) GetGenTokenAt() string {
	if x != nil {
		return x.GenTokenAt
	}
	return ""
}

type CheckUserEnableResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Language string `protobuf:"bytes,1,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *CheckUserEnableResp) Reset() {
	*x = CheckUserEnableResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_perm_service_perm_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserEnableResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserEnableResp) ProtoMessage() {}

func (x *CheckUserEnableResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_perm_service_perm_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserEnableResp.ProtoReflect.Descriptor instead.
func (*CheckUserEnableResp) Descriptor() ([]byte, []int) {
	return file_proto_perm_service_perm_service_proto_rawDescGZIP(), []int{1}
}

func (x *CheckUserEnableResp) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

type CheckUserPermReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	GenTokenAt string   `protobuf:"bytes,2,opt,name=genTokenAt,proto3" json:"genTokenAt,omitempty"`
	OrgId      string   `protobuf:"bytes,3,opt,name=orgId,proto3" json:"orgId,omitempty"`
	OneOfPerms []string `protobuf:"bytes,4,rep,name=oneOfPerms,proto3" json:"oneOfPerms,omitempty"`
}

func (x *CheckUserPermReq) Reset() {
	*x = CheckUserPermReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_perm_service_perm_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserPermReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserPermReq) ProtoMessage() {}

func (x *CheckUserPermReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_perm_service_perm_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserPermReq.ProtoReflect.Descriptor instead.
func (*CheckUserPermReq) Descriptor() ([]byte, []int) {
	return file_proto_perm_service_perm_service_proto_rawDescGZIP(), []int{2}
}

func (x *CheckUserPermReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CheckUserPermReq) GetGenTokenAt() string {
	if x != nil {
		return x.GenTokenAt
	}
	return ""
}

func (x *CheckUserPermReq) GetOrgId() string {
	if x != nil {
		return x.OrgId
	}
	return ""
}

func (x *CheckUserPermReq) GetOneOfPerms() []string {
	if x != nil {
		return x.OneOfPerms
	}
	return nil
}

type CheckUserPermResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsAdmin  bool   `protobuf:"varint,1,opt,name=isAdmin,proto3" json:"isAdmin,omitempty"`
	IsSystem bool   `protobuf:"varint,2,opt,name=isSystem,proto3" json:"isSystem,omitempty"`
	Language string `protobuf:"bytes,3,opt,name=language,proto3" json:"language,omitempty"`
}

func (x *CheckUserPermResp) Reset() {
	*x = CheckUserPermResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_perm_service_perm_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserPermResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserPermResp) ProtoMessage() {}

func (x *CheckUserPermResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_perm_service_perm_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserPermResp.ProtoReflect.Descriptor instead.
func (*CheckUserPermResp) Descriptor() ([]byte, []int) {
	return file_proto_perm_service_perm_service_proto_rawDescGZIP(), []int{3}
}

func (x *CheckUserPermResp) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *CheckUserPermResp) GetIsSystem() bool {
	if x != nil {
		return x.IsSystem
	}
	return false
}

func (x *CheckUserPermResp) GetLanguage() string {
	if x != nil {
		return x.Language
	}
	return ""
}

var File_proto_perm_service_perm_service_proto protoreflect.FileDescriptor

var file_proto_perm_service_perm_service_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x65, 0x72, 0x6d, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x4c, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73,
	0x65, 0x72, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x41, 0x74, 0x22, 0x31, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72,
	0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x41,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67, 0x65, 0x6e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x41, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x6f, 0x72, 0x67, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x6e, 0x65,
	0x4f, 0x66, 0x50, 0x65, 0x72, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x6f,
	0x6e, 0x65, 0x4f, 0x66, 0x50, 0x65, 0x72, 0x6d, 0x73, 0x22, 0x65, 0x0a, 0x11, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x12, 0x18,
	0x0a, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x73, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x32, 0xbb, 0x01, 0x0a, 0x0b, 0x50, 0x65, 0x72, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x58, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x20, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6e, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x21, 0x2e, 0x70, 0x65, 0x72, 0x6d, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x45, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0d, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x12, 0x1e, 0x2e, 0x70, 0x65,
	0x72, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x70, 0x65,
	0x72, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x65, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x32,
	0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x55, 0x6e, 0x69,
	0x63, 0x6f, 0x6d, 0x41, 0x49, 0x2f, 0x77, 0x61, 0x6e, 0x77, 0x75, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x65, 0x72, 0x6d, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_perm_service_perm_service_proto_rawDescOnce sync.Once
	file_proto_perm_service_perm_service_proto_rawDescData = file_proto_perm_service_perm_service_proto_rawDesc
)

func file_proto_perm_service_perm_service_proto_rawDescGZIP() []byte {
	file_proto_perm_service_perm_service_proto_rawDescOnce.Do(func() {
		file_proto_perm_service_perm_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_perm_service_perm_service_proto_rawDescData)
	})
	return file_proto_perm_service_perm_service_proto_rawDescData
}

var file_proto_perm_service_perm_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_perm_service_perm_service_proto_goTypes = []interface{}{
	(*CheckUserEnableReq)(nil),  // 0: perm_service.CheckUserEnableReq
	(*CheckUserEnableResp)(nil), // 1: perm_service.CheckUserEnableResp
	(*CheckUserPermReq)(nil),    // 2: perm_service.CheckUserPermReq
	(*CheckUserPermResp)(nil),   // 3: perm_service.CheckUserPermResp
}
var file_proto_perm_service_perm_service_proto_depIdxs = []int32{
	0, // 0: perm_service.PermService.CheckUserEnable:input_type -> perm_service.CheckUserEnableReq
	2, // 1: perm_service.PermService.CheckUserPerm:input_type -> perm_service.CheckUserPermReq
	1, // 2: perm_service.PermService.CheckUserEnable:output_type -> perm_service.CheckUserEnableResp
	3, // 3: perm_service.PermService.CheckUserPerm:output_type -> perm_service.CheckUserPermResp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_perm_service_perm_service_proto_init() }
func file_proto_perm_service_perm_service_proto_init() {
	if File_proto_perm_service_perm_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_perm_service_perm_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserEnableReq); i {
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
		file_proto_perm_service_perm_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserEnableResp); i {
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
		file_proto_perm_service_perm_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserPermReq); i {
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
		file_proto_perm_service_perm_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserPermResp); i {
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
			RawDescriptor: file_proto_perm_service_perm_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_perm_service_perm_service_proto_goTypes,
		DependencyIndexes: file_proto_perm_service_perm_service_proto_depIdxs,
		MessageInfos:      file_proto_perm_service_perm_service_proto_msgTypes,
	}.Build()
	File_proto_perm_service_perm_service_proto = out.File
	file_proto_perm_service_perm_service_proto_rawDesc = nil
	file_proto_perm_service_perm_service_proto_goTypes = nil
	file_proto_perm_service_perm_service_proto_depIdxs = nil
}
