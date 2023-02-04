// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.12
// source: appmgmt.proto

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

//AppMgmtConfig app管理系统 配置
type AppMgmtConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group          string `protobuf:"bytes,1,opt,name=group,proto3" json:"group"`
	K              string `protobuf:"bytes,2,opt,name=k,proto3" json:"k"`
	V              string `protobuf:"bytes,3,opt,name=v,proto3" json:"v"`
	Type           string `protobuf:"bytes,4,opt,name=type,proto3" json:"type"`
	Name           string `protobuf:"bytes,5,opt,name=name,proto3" json:"name"`
	ScopePlatforms string `protobuf:"bytes,6,opt,name=scopePlatforms,proto3" json:"scopePlatforms"`
}

func (x *AppMgmtConfig) Reset() {
	*x = AppMgmtConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appmgmt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppMgmtConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppMgmtConfig) ProtoMessage() {}

func (x *AppMgmtConfig) ProtoReflect() protoreflect.Message {
	mi := &file_appmgmt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppMgmtConfig.ProtoReflect.Descriptor instead.
func (*AppMgmtConfig) Descriptor() ([]byte, []int) {
	return file_appmgmt_proto_rawDescGZIP(), []int{0}
}

func (x *AppMgmtConfig) GetGroup() string {
	if x != nil {
		return x.Group
	}
	return ""
}

func (x *AppMgmtConfig) GetK() string {
	if x != nil {
		return x.K
	}
	return ""
}

func (x *AppMgmtConfig) GetV() string {
	if x != nil {
		return x.V
	}
	return ""
}

func (x *AppMgmtConfig) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *AppMgmtConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AppMgmtConfig) GetScopePlatforms() string {
	if x != nil {
		return x.ScopePlatforms
	}
	return ""
}

//GetAllAppMgmtConfigReq 获取所有配置
type GetAllAppMgmtConfigReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonReq     *CommonReq `protobuf:"bytes,1,opt,name=commonReq,proto3" json:"commonReq"`
	ScopePlatform string     `protobuf:"bytes,2,opt,name=scopePlatform,proto3" json:"scopePlatform"`
}

func (x *GetAllAppMgmtConfigReq) Reset() {
	*x = GetAllAppMgmtConfigReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appmgmt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAppMgmtConfigReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAppMgmtConfigReq) ProtoMessage() {}

func (x *GetAllAppMgmtConfigReq) ProtoReflect() protoreflect.Message {
	mi := &file_appmgmt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAppMgmtConfigReq.ProtoReflect.Descriptor instead.
func (*GetAllAppMgmtConfigReq) Descriptor() ([]byte, []int) {
	return file_appmgmt_proto_rawDescGZIP(), []int{1}
}

func (x *GetAllAppMgmtConfigReq) GetCommonReq() *CommonReq {
	if x != nil {
		return x.CommonReq
	}
	return nil
}

func (x *GetAllAppMgmtConfigReq) GetScopePlatform() string {
	if x != nil {
		return x.ScopePlatform
	}
	return ""
}

//GetAllAppMgmtConfigResp 获取所有配置
type GetAllAppMgmtConfigResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonResp     *CommonResp      `protobuf:"bytes,1,opt,name=commonResp,proto3" json:"commonResp"`
	AppMgmtConfigs []*AppMgmtConfig `protobuf:"bytes,2,rep,name=appMgmtConfigs,proto3" json:"appMgmtConfigs"`
}

func (x *GetAllAppMgmtConfigResp) Reset() {
	*x = GetAllAppMgmtConfigResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appmgmt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllAppMgmtConfigResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllAppMgmtConfigResp) ProtoMessage() {}

func (x *GetAllAppMgmtConfigResp) ProtoReflect() protoreflect.Message {
	mi := &file_appmgmt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllAppMgmtConfigResp.ProtoReflect.Descriptor instead.
func (*GetAllAppMgmtConfigResp) Descriptor() ([]byte, []int) {
	return file_appmgmt_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllAppMgmtConfigResp) GetCommonResp() *CommonResp {
	if x != nil {
		return x.CommonResp
	}
	return nil
}

func (x *GetAllAppMgmtConfigResp) GetAppMgmtConfigs() []*AppMgmtConfig {
	if x != nil {
		return x.AppMgmtConfigs
	}
	return nil
}

//UpdateAppMgmtConfigReq 更新配置
type UpdateAppMgmtConfigReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonReq      *CommonReq       `protobuf:"bytes,1,opt,name=commonReq,proto3" json:"commonReq"`
	AppMgmtConfigs []*AppMgmtConfig `protobuf:"bytes,2,rep,name=appMgmtConfigs,proto3" json:"appMgmtConfigs"`
}

func (x *UpdateAppMgmtConfigReq) Reset() {
	*x = UpdateAppMgmtConfigReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appmgmt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAppMgmtConfigReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAppMgmtConfigReq) ProtoMessage() {}

func (x *UpdateAppMgmtConfigReq) ProtoReflect() protoreflect.Message {
	mi := &file_appmgmt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAppMgmtConfigReq.ProtoReflect.Descriptor instead.
func (*UpdateAppMgmtConfigReq) Descriptor() ([]byte, []int) {
	return file_appmgmt_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateAppMgmtConfigReq) GetCommonReq() *CommonReq {
	if x != nil {
		return x.CommonReq
	}
	return nil
}

func (x *UpdateAppMgmtConfigReq) GetAppMgmtConfigs() []*AppMgmtConfig {
	if x != nil {
		return x.AppMgmtConfigs
	}
	return nil
}

//UpdateAppMgmtConfigResp 更新配置
type UpdateAppMgmtConfigResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonResp *CommonResp `protobuf:"bytes,1,opt,name=commonResp,proto3" json:"commonResp"`
}

func (x *UpdateAppMgmtConfigResp) Reset() {
	*x = UpdateAppMgmtConfigResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_appmgmt_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAppMgmtConfigResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAppMgmtConfigResp) ProtoMessage() {}

func (x *UpdateAppMgmtConfigResp) ProtoReflect() protoreflect.Message {
	mi := &file_appmgmt_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAppMgmtConfigResp.ProtoReflect.Descriptor instead.
func (*UpdateAppMgmtConfigResp) Descriptor() ([]byte, []int) {
	return file_appmgmt_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAppMgmtConfigResp) GetCommonResp() *CommonResp {
	if x != nil {
		return x.CommonResp
	}
	return nil
}

var File_appmgmt_proto protoreflect.FileDescriptor

var file_appmgmt_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x70, 0x6d, 0x67, 0x6d, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91, 0x01,
	0x0a, 0x0d, 0x41, 0x70, 0x70, 0x4d, 0x67, 0x6d, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x14, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0c, 0x0a, 0x01, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x01, 0x6b, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01,
	0x76, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x73, 0x22, 0x6b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x70, 0x70, 0x4d, 0x67,
	0x6d, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x09, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x52, 0x09, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x63, 0x6f, 0x70,
	0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x50, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x22, 0x84,
	0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x70, 0x70, 0x4d, 0x67, 0x6d, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x0a,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x39, 0x0a, 0x0e, 0x61, 0x70,
	0x70, 0x4d, 0x67, 0x6d, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x70, 0x70, 0x4d, 0x67, 0x6d, 0x74, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x61, 0x70, 0x70, 0x4d, 0x67, 0x6d, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x70, 0x70, 0x4d, 0x67, 0x6d, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71,
	0x12, 0x2b, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x39, 0x0a,
	0x0e, 0x61, 0x70, 0x70, 0x4d, 0x67, 0x6d, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x70, 0x70, 0x4d, 0x67,
	0x6d, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x0e, 0x61, 0x70, 0x70, 0x4d, 0x67, 0x6d,
	0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x22, 0x49, 0x0a, 0x17, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x41, 0x70, 0x70, 0x4d, 0x67, 0x6d, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x32, 0xb0, 0x01, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x4d, 0x67, 0x6d, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x41, 0x70, 0x70, 0x4d, 0x67, 0x6d, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x70, 0x70, 0x4d, 0x67, 0x6d, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x41, 0x70, 0x70, 0x4d, 0x67, 0x6d, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x4e, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x70, 0x70, 0x4d, 0x67, 0x6d, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x2e,
	0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x4d, 0x67, 0x6d, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x70, 0x70, 0x4d, 0x67, 0x6d, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_appmgmt_proto_rawDescOnce sync.Once
	file_appmgmt_proto_rawDescData = file_appmgmt_proto_rawDesc
)

func file_appmgmt_proto_rawDescGZIP() []byte {
	file_appmgmt_proto_rawDescOnce.Do(func() {
		file_appmgmt_proto_rawDescData = protoimpl.X.CompressGZIP(file_appmgmt_proto_rawDescData)
	})
	return file_appmgmt_proto_rawDescData
}

var file_appmgmt_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_appmgmt_proto_goTypes = []interface{}{
	(*AppMgmtConfig)(nil),           // 0: pb.AppMgmtConfig
	(*GetAllAppMgmtConfigReq)(nil),  // 1: pb.GetAllAppMgmtConfigReq
	(*GetAllAppMgmtConfigResp)(nil), // 2: pb.GetAllAppMgmtConfigResp
	(*UpdateAppMgmtConfigReq)(nil),  // 3: pb.UpdateAppMgmtConfigReq
	(*UpdateAppMgmtConfigResp)(nil), // 4: pb.UpdateAppMgmtConfigResp
	(*CommonReq)(nil),               // 5: pb.CommonReq
	(*CommonResp)(nil),              // 6: pb.CommonResp
}
var file_appmgmt_proto_depIdxs = []int32{
	5, // 0: pb.GetAllAppMgmtConfigReq.commonReq:type_name -> pb.CommonReq
	6, // 1: pb.GetAllAppMgmtConfigResp.commonResp:type_name -> pb.CommonResp
	0, // 2: pb.GetAllAppMgmtConfigResp.appMgmtConfigs:type_name -> pb.AppMgmtConfig
	5, // 3: pb.UpdateAppMgmtConfigReq.commonReq:type_name -> pb.CommonReq
	0, // 4: pb.UpdateAppMgmtConfigReq.appMgmtConfigs:type_name -> pb.AppMgmtConfig
	6, // 5: pb.UpdateAppMgmtConfigResp.commonResp:type_name -> pb.CommonResp
	1, // 6: pb.appMgmtService.GetAllAppMgmtConfig:input_type -> pb.GetAllAppMgmtConfigReq
	3, // 7: pb.appMgmtService.UpdateAppMgmtConfig:input_type -> pb.UpdateAppMgmtConfigReq
	2, // 8: pb.appMgmtService.GetAllAppMgmtConfig:output_type -> pb.GetAllAppMgmtConfigResp
	4, // 9: pb.appMgmtService.UpdateAppMgmtConfig:output_type -> pb.UpdateAppMgmtConfigResp
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_appmgmt_proto_init() }
func file_appmgmt_proto_init() {
	if File_appmgmt_proto != nil {
		return
	}
	file_common_proto_init()
	file_conn_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_appmgmt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppMgmtConfig); i {
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
		file_appmgmt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAppMgmtConfigReq); i {
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
		file_appmgmt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllAppMgmtConfigResp); i {
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
		file_appmgmt_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAppMgmtConfigReq); i {
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
		file_appmgmt_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAppMgmtConfigResp); i {
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
			RawDescriptor: file_appmgmt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_appmgmt_proto_goTypes,
		DependencyIndexes: file_appmgmt_proto_depIdxs,
		MessageInfos:      file_appmgmt_proto_msgTypes,
	}.Build()
	File_appmgmt_proto = out.File
	file_appmgmt_proto_rawDesc = nil
	file_appmgmt_proto_goTypes = nil
	file_appmgmt_proto_depIdxs = nil
}
