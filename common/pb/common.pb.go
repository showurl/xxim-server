// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.5
// source: common.proto

package pb

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

type ConvType int32

const (
	ConvType_SINGLE ConvType = 0 // 单聊
	ConvType_GROUP  ConvType = 1 // 群聊
)

// Enum value maps for ConvType.
var (
	ConvType_name = map[int32]string{
		0: "SINGLE",
		1: "GROUP",
	}
	ConvType_value = map[string]int32{
		"SINGLE": 0,
		"GROUP":  1,
	}
)

func (x ConvType) Enum() *ConvType {
	p := new(ConvType)
	*p = x
	return p
}

func (x ConvType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConvType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[0].Descriptor()
}

func (ConvType) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[0]
}

func (x ConvType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConvType.Descriptor instead.
func (ConvType) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

type ContentType int32

const (
	ContentType_UNKNOWN  ContentType = 0
	ContentType_TYPING   ContentType = 1   // 正在输入
	ContentType_READ     ContentType = 2   // 已读
	ContentType_REVOKE   ContentType = 3   // 撤回
	ContentType_TEXT     ContentType = 11  // 文本
	ContentType_IMAGE    ContentType = 12  // 图片
	ContentType_AUDIO    ContentType = 13  // 语音
	ContentType_VIDEO    ContentType = 14  // 视频
	ContentType_FILE     ContentType = 15  // 文件
	ContentType_LOCATION ContentType = 16  // 位置
	ContentType_CARD     ContentType = 17  // 名片
	ContentType_MERGE    ContentType = 18  // 合并
	ContentType_EMOJI    ContentType = 19  // 表情
	ContentType_COMMAND  ContentType = 20  // 命令
	ContentType_CUSTOM   ContentType = 100 // 自定义消息
)

// Enum value maps for ContentType.
var (
	ContentType_name = map[int32]string{
		0:   "UNKNOWN",
		1:   "TYPING",
		2:   "READ",
		3:   "REVOKE",
		11:  "TEXT",
		12:  "IMAGE",
		13:  "AUDIO",
		14:  "VIDEO",
		15:  "FILE",
		16:  "LOCATION",
		17:  "CARD",
		18:  "MERGE",
		19:  "EMOJI",
		20:  "COMMAND",
		100: "CUSTOM",
	}
	ContentType_value = map[string]int32{
		"UNKNOWN":  0,
		"TYPING":   1,
		"READ":     2,
		"REVOKE":   3,
		"TEXT":     11,
		"IMAGE":    12,
		"AUDIO":    13,
		"VIDEO":    14,
		"FILE":     15,
		"LOCATION": 16,
		"CARD":     17,
		"MERGE":    18,
		"EMOJI":    19,
		"COMMAND":  20,
		"CUSTOM":   100,
	}
)

func (x ContentType) Enum() *ContentType {
	p := new(ContentType)
	*p = x
	return p
}

func (x ContentType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContentType) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[1].Descriptor()
}

func (ContentType) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[1]
}

func (x ContentType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContentType.Descriptor instead.
func (ContentType) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

type CommonResp_Code int32

const (
	CommonResp_Success       CommonResp_Code = 0
	CommonResp_UnknownError  CommonResp_Code = 1 // 未知 error
	CommonResp_InternalError CommonResp_Code = 2 // 内部错误
	CommonResp_RequestError  CommonResp_Code = 3 // 请求错误
	CommonResp_AuthError     CommonResp_Code = 4 // 鉴权错误 // 应该退出登录
	CommonResp_ToastError    CommonResp_Code = 5 // toast 错误 只有 message
	CommonResp_AlertError    CommonResp_Code = 7 // alert 错误 只有一个确定按钮
	CommonResp_RetryError    CommonResp_Code = 8 // alert 错误 有一个取消按钮 和一个重试按钮
)

// Enum value maps for CommonResp_Code.
var (
	CommonResp_Code_name = map[int32]string{
		0: "Success",
		1: "UnknownError",
		2: "InternalError",
		3: "RequestError",
		4: "AuthError",
		5: "ToastError",
		7: "AlertError",
		8: "RetryError",
	}
	CommonResp_Code_value = map[string]int32{
		"Success":       0,
		"UnknownError":  1,
		"InternalError": 2,
		"RequestError":  3,
		"AuthError":     4,
		"ToastError":    5,
		"AlertError":    7,
		"RetryError":    8,
	}
)

func (x CommonResp_Code) Enum() *CommonResp_Code {
	p := new(CommonResp_Code)
	*p = x
	return p
}

func (x CommonResp_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommonResp_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_common_proto_enumTypes[2].Descriptor()
}

func (CommonResp_Code) Type() protoreflect.EnumType {
	return &file_common_proto_enumTypes[2]
}

func (x CommonResp_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommonResp_Code.Descriptor instead.
func (CommonResp_Code) EnumDescriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0, 0}
}

type CommonResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code CommonResp_Code `protobuf:"varint,1,opt,name=code,proto3,enum=pb.CommonResp_Code" json:"code"`
	Msg  *string         `protobuf:"bytes,2,opt,name=msg,proto3,oneof" json:"msg"`
	Data *string         `protobuf:"bytes,3,opt,name=data,proto3,oneof" json:"data"`
}

func (x *CommonResp) Reset() {
	*x = CommonResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResp) ProtoMessage() {}

func (x *CommonResp) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResp.ProtoReflect.Descriptor instead.
func (*CommonResp) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{0}
}

func (x *CommonResp) GetCode() CommonResp_Code {
	if x != nil {
		return x.Code
	}
	return CommonResp_Success
}

func (x *CommonResp) GetMsg() string {
	if x != nil && x.Msg != nil {
		return *x.Msg
	}
	return ""
}

func (x *CommonResp) GetData() string {
	if x != nil && x.Data != nil {
		return *x.Data
	}
	return ""
}

type Requester struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	Token      string `protobuf:"bytes,2,opt,name=token,proto3" json:"token"`
	AppVersion string `protobuf:"bytes,3,opt,name=appVersion,proto3" json:"appVersion"`
	Ip         string `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip"`
	// header
	Ua          string `protobuf:"bytes,5,opt,name=ua,proto3" json:"ua"`
	OsVersion   string `protobuf:"bytes,6,opt,name=osVersion,proto3" json:"osVersion"`
	Platform    string `protobuf:"bytes,7,opt,name=platform,proto3" json:"platform"`
	DeviceModel string `protobuf:"bytes,8,opt,name=deviceModel,proto3" json:"deviceModel"`
	DeviceId    string `protobuf:"bytes,9,opt,name=deviceId,proto3" json:"deviceId"`
}

func (x *Requester) Reset() {
	*x = Requester{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Requester) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Requester) ProtoMessage() {}

func (x *Requester) ProtoReflect() protoreflect.Message {
	mi := &file_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Requester.ProtoReflect.Descriptor instead.
func (*Requester) Descriptor() ([]byte, []int) {
	return file_common_proto_rawDescGZIP(), []int{1}
}

func (x *Requester) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Requester) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *Requester) GetAppVersion() string {
	if x != nil {
		return x.AppVersion
	}
	return ""
}

func (x *Requester) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Requester) GetUa() string {
	if x != nil {
		return x.Ua
	}
	return ""
}

func (x *Requester) GetOsVersion() string {
	if x != nil {
		return x.OsVersion
	}
	return ""
}

func (x *Requester) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *Requester) GetDeviceModel() string {
	if x != nil {
		return x.DeviceModel
	}
	return ""
}

func (x *Requester) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

var File_common_proto protoreflect.FileDescriptor

var file_common_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x82, 0x02, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x27, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x13, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x2e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x15, 0x0a, 0x03, 0x6d, 0x73,
	0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x88, 0x01,
	0x01, 0x12, 0x17, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48,
	0x01, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x88, 0x01, 0x01, 0x22, 0x89, 0x01, 0x0a, 0x04, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x10, 0x00,
	0x12, 0x10, 0x0a, 0x0c, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x03, 0x12, 0x0d, 0x0a, 0x09, 0x41, 0x75, 0x74, 0x68, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x10, 0x04, 0x12, 0x0e, 0x0a, 0x0a, 0x54, 0x6f, 0x61, 0x73, 0x74, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x6c, 0x65, 0x72, 0x74, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x10, 0x07, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x65, 0x74, 0x72, 0x79, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x10, 0x08, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x6d, 0x73, 0x67, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x22, 0xe9, 0x01, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x61,
	0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x70, 0x70, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x75,
	0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x75, 0x61, 0x12, 0x1c, 0x0a, 0x09, 0x6f,
	0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6f, 0x73, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61,
	0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4d,
	0x6f, 0x64, 0x65, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x2a, 0x21, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x76, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x0a, 0x0a, 0x06, 0x53, 0x49, 0x4e, 0x47, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x47,
	0x52, 0x4f, 0x55, 0x50, 0x10, 0x01, 0x2a, 0xb8, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x54, 0x59, 0x50, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12,
	0x08, 0x0a, 0x04, 0x52, 0x45, 0x41, 0x44, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x56,
	0x4f, 0x4b, 0x45, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x45, 0x58, 0x54, 0x10, 0x0b, 0x12,
	0x09, 0x0a, 0x05, 0x49, 0x4d, 0x41, 0x47, 0x45, 0x10, 0x0c, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x55,
	0x44, 0x49, 0x4f, 0x10, 0x0d, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x49, 0x44, 0x45, 0x4f, 0x10, 0x0e,
	0x12, 0x08, 0x0a, 0x04, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x0f, 0x12, 0x0c, 0x0a, 0x08, 0x4c, 0x4f,
	0x43, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x10, 0x12, 0x08, 0x0a, 0x04, 0x43, 0x41, 0x52, 0x44,
	0x10, 0x11, 0x12, 0x09, 0x0a, 0x05, 0x4d, 0x45, 0x52, 0x47, 0x45, 0x10, 0x12, 0x12, 0x09, 0x0a,
	0x05, 0x45, 0x4d, 0x4f, 0x4a, 0x49, 0x10, 0x13, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x4f, 0x4d, 0x4d,
	0x41, 0x4e, 0x44, 0x10, 0x14, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x10,
	0x64, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_common_proto_rawDescOnce sync.Once
	file_common_proto_rawDescData = file_common_proto_rawDesc
)

func file_common_proto_rawDescGZIP() []byte {
	file_common_proto_rawDescOnce.Do(func() {
		file_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_proto_rawDescData)
	})
	return file_common_proto_rawDescData
}

var file_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_common_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_common_proto_goTypes = []interface{}{
	(ConvType)(0),        // 0: pb.ConvType
	(ContentType)(0),     // 1: pb.ContentType
	(CommonResp_Code)(0), // 2: pb.CommonResp.Code
	(*CommonResp)(nil),   // 3: pb.CommonResp
	(*Requester)(nil),    // 4: pb.Requester
}
var file_common_proto_depIdxs = []int32{
	2, // 0: pb.CommonResp.code:type_name -> pb.CommonResp.Code
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_common_proto_init() }
func file_common_proto_init() {
	if File_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResp); i {
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
		file_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Requester); i {
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
	file_common_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_common_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_common_proto_goTypes,
		DependencyIndexes: file_common_proto_depIdxs,
		EnumInfos:         file_common_proto_enumTypes,
		MessageInfos:      file_common_proto_msgTypes,
	}.Build()
	File_common_proto = out.File
	file_common_proto_rawDesc = nil
	file_common_proto_goTypes = nil
	file_common_proto_depIdxs = nil
}
