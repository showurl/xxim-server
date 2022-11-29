// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.5
// source: im.proto

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

// 推送事件
type ImMQBody_Event int32

const (
	ImMQBody_Unknown ImMQBody_Event = 0
)

// Enum value maps for ImMQBody_Event.
var (
	ImMQBody_Event_name = map[int32]string{
		0: "Unknown",
	}
	ImMQBody_Event_value = map[string]int32{
		"Unknown": 0,
	}
)

func (x ImMQBody_Event) Enum() *ImMQBody_Event {
	p := new(ImMQBody_Event)
	*p = x
	return p
}

func (x ImMQBody_Event) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ImMQBody_Event) Descriptor() protoreflect.EnumDescriptor {
	return file_im_proto_enumTypes[0].Descriptor()
}

func (ImMQBody_Event) Type() protoreflect.EnumType {
	return &file_im_proto_enumTypes[0]
}

func (x ImMQBody_Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ImMQBody_Event.Descriptor instead.
func (ImMQBody_Event) EnumDescriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{0, 0}
}

type ImMQBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event ImMQBody_Event `protobuf:"varint,1,opt,name=event,proto3,enum=pb.ImMQBody_Event" json:"event"`
	Data  []byte         `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *ImMQBody) Reset() {
	*x = ImMQBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ImMQBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ImMQBody) ProtoMessage() {}

func (x *ImMQBody) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ImMQBody.ProtoReflect.Descriptor instead.
func (*ImMQBody) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{0}
}

func (x *ImMQBody) GetEvent() ImMQBody_Event {
	if x != nil {
		return x.Event
	}
	return ImMQBody_Unknown
}

func (x *ImMQBody) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type BeforeConnectReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConnParam *ConnParam `protobuf:"bytes,1,opt,name=connParam,proto3" json:"connParam"`
}

func (x *BeforeConnectReq) Reset() {
	*x = BeforeConnectReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeforeConnectReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeforeConnectReq) ProtoMessage() {}

func (x *BeforeConnectReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeforeConnectReq.ProtoReflect.Descriptor instead.
func (*BeforeConnectReq) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{1}
}

func (x *BeforeConnectReq) GetConnParam() *ConnParam {
	if x != nil {
		return x.ConnParam
	}
	return nil
}

type BeforeConnectResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg"`
}

func (x *BeforeConnectResp) Reset() {
	*x = BeforeConnectResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeforeConnectResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeforeConnectResp) ProtoMessage() {}

func (x *BeforeConnectResp) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeforeConnectResp.ProtoReflect.Descriptor instead.
func (*BeforeConnectResp) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{2}
}

func (x *BeforeConnectResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *BeforeConnectResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetUserLatestConnReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId"`
}

func (x *GetUserLatestConnReq) Reset() {
	*x = GetUserLatestConnReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserLatestConnReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserLatestConnReq) ProtoMessage() {}

func (x *GetUserLatestConnReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserLatestConnReq.ProtoReflect.Descriptor instead.
func (*GetUserLatestConnReq) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserLatestConnReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserLatestConnResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId         string    `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId"`
	Ip             string    `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip"`
	IpRegion       *IpRegion `protobuf:"bytes,3,opt,name=ipRegion,proto3" json:"ipRegion"`
	ConnectedAt    string    `protobuf:"bytes,4,opt,name=connectedAt,proto3" json:"connectedAt"`
	DisconnectedAt string    `protobuf:"bytes,5,opt,name=disconnectedAt,proto3" json:"disconnectedAt"`
	Platform       string    `protobuf:"bytes,6,opt,name=platform,proto3" json:"platform"`
	DeviceId       string    `protobuf:"bytes,7,opt,name=deviceId,proto3" json:"deviceId"`
}

func (x *GetUserLatestConnResp) Reset() {
	*x = GetUserLatestConnResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserLatestConnResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserLatestConnResp) ProtoMessage() {}

func (x *GetUserLatestConnResp) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserLatestConnResp.ProtoReflect.Descriptor instead.
func (*GetUserLatestConnResp) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserLatestConnResp) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetUserLatestConnResp) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *GetUserLatestConnResp) GetIpRegion() *IpRegion {
	if x != nil {
		return x.IpRegion
	}
	return nil
}

func (x *GetUserLatestConnResp) GetConnectedAt() string {
	if x != nil {
		return x.ConnectedAt
	}
	return ""
}

func (x *GetUserLatestConnResp) GetDisconnectedAt() string {
	if x != nil {
		return x.DisconnectedAt
	}
	return ""
}

func (x *GetUserLatestConnResp) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *GetUserLatestConnResp) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

//batchGetUserLatestConnReq 批量获取用户最新连接信息
type BatchGetUserLatestConnReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds []string `protobuf:"bytes,1,rep,name=userIds,proto3" json:"userIds"`
}

func (x *BatchGetUserLatestConnReq) Reset() {
	*x = BatchGetUserLatestConnReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetUserLatestConnReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetUserLatestConnReq) ProtoMessage() {}

func (x *BatchGetUserLatestConnReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetUserLatestConnReq.ProtoReflect.Descriptor instead.
func (*BatchGetUserLatestConnReq) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{5}
}

func (x *BatchGetUserLatestConnReq) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

//batchGetUserLatestConnResp 批量获取用户最新连接信息
type BatchGetUserLatestConnResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonResp      *CommonResp              `protobuf:"bytes,1,opt,name=commonResp,proto3" json:"commonResp"`
	UserLatestConns []*GetUserLatestConnResp `protobuf:"bytes,2,rep,name=userLatestConns,proto3" json:"userLatestConns"`
}

func (x *BatchGetUserLatestConnResp) Reset() {
	*x = BatchGetUserLatestConnResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchGetUserLatestConnResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchGetUserLatestConnResp) ProtoMessage() {}

func (x *BatchGetUserLatestConnResp) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchGetUserLatestConnResp.ProtoReflect.Descriptor instead.
func (*BatchGetUserLatestConnResp) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{6}
}

func (x *BatchGetUserLatestConnResp) GetCommonResp() *CommonResp {
	if x != nil {
		return x.CommonResp
	}
	return nil
}

func (x *BatchGetUserLatestConnResp) GetUserLatestConns() []*GetUserLatestConnResp {
	if x != nil {
		return x.UserLatestConns
	}
	return nil
}

//MsgNotifyOpt 消息通知选项
type MsgNotifyOpt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoDisturb bool   `protobuf:"varint,1,opt,name=noDisturb,proto3" json:"noDisturb"`
	Preview   bool   `protobuf:"varint,2,opt,name=preview,proto3" json:"preview"`    // 是否预览
	Sound     bool   `protobuf:"varint,3,opt,name=sound,proto3" json:"sound"`        // 是否声音
	SoundName string `protobuf:"bytes,4,opt,name=soundName,proto3" json:"soundName"` // 声音名称
	Vibrate   bool   `protobuf:"varint,5,opt,name=vibrate,proto3" json:"vibrate"`    // 是否震动
}

func (x *MsgNotifyOpt) Reset() {
	*x = MsgNotifyOpt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MsgNotifyOpt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MsgNotifyOpt) ProtoMessage() {}

func (x *MsgNotifyOpt) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MsgNotifyOpt.ProtoReflect.Descriptor instead.
func (*MsgNotifyOpt) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{7}
}

func (x *MsgNotifyOpt) GetNoDisturb() bool {
	if x != nil {
		return x.NoDisturb
	}
	return false
}

func (x *MsgNotifyOpt) GetPreview() bool {
	if x != nil {
		return x.Preview
	}
	return false
}

func (x *MsgNotifyOpt) GetSound() bool {
	if x != nil {
		return x.Sound
	}
	return false
}

func (x *MsgNotifyOpt) GetSoundName() string {
	if x != nil {
		return x.SoundName
	}
	return ""
}

func (x *MsgNotifyOpt) GetVibrate() bool {
	if x != nil {
		return x.Vibrate
	}
	return false
}

//GetAppSystemConfigReq 获取系统配置
type GetAppSystemConfigReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonReq *CommonReq `protobuf:"bytes,1,opt,name=commonReq,proto3" json:"commonReq"`
}

func (x *GetAppSystemConfigReq) Reset() {
	*x = GetAppSystemConfigReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppSystemConfigReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppSystemConfigReq) ProtoMessage() {}

func (x *GetAppSystemConfigReq) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppSystemConfigReq.ProtoReflect.Descriptor instead.
func (*GetAppSystemConfigReq) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{8}
}

func (x *GetAppSystemConfigReq) GetCommonReq() *CommonReq {
	if x != nil {
		return x.CommonReq
	}
	return nil
}

//GetAppSystemConfigResp 获取系统配置
type GetAppSystemConfigResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonResp *CommonResp       `protobuf:"bytes,1,opt,name=commonResp,proto3" json:"commonResp"`
	Configs    map[string]string `protobuf:"bytes,2,rep,name=configs,proto3" json:"configs" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetAppSystemConfigResp) Reset() {
	*x = GetAppSystemConfigResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_im_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAppSystemConfigResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAppSystemConfigResp) ProtoMessage() {}

func (x *GetAppSystemConfigResp) ProtoReflect() protoreflect.Message {
	mi := &file_im_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAppSystemConfigResp.ProtoReflect.Descriptor instead.
func (*GetAppSystemConfigResp) Descriptor() ([]byte, []int) {
	return file_im_proto_rawDescGZIP(), []int{9}
}

func (x *GetAppSystemConfigResp) GetCommonResp() *CommonResp {
	if x != nil {
		return x.CommonResp
	}
	return nil
}

func (x *GetAppSystemConfigResp) GetConfigs() map[string]string {
	if x != nil {
		return x.Configs
	}
	return nil
}

var File_im_proto protoreflect.FileDescriptor

var file_im_proto_rawDesc = []byte{
	0x0a, 0x08, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0c,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0a, 0x63, 0x6f,
	0x6e, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x08, 0x49, 0x6d, 0x4d, 0x51,
	0x42, 0x6f, 0x64, 0x79, 0x12, 0x28, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6d, 0x4d, 0x51, 0x42, 0x6f, 0x64,
	0x79, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x14, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x55,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x10, 0x00, 0x22, 0x3f, 0x0a, 0x10, 0x42, 0x65, 0x66, 0x6f,
	0x72, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a, 0x09,
	0x63, 0x6f, 0x6e, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x09,
	0x63, 0x6f, 0x6e, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x22, 0x39, 0x0a, 0x11, 0x42, 0x65, 0x66,
	0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12,
	0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x73, 0x67, 0x22, 0x2e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0xeb, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x28, 0x0a, 0x08, 0x69, 0x70, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x70,
	0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x69, 0x70, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x64, 0x22, 0x35, 0x0a, 0x19, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x71, 0x12,
	0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x91, 0x01, 0x0a, 0x1a, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x0a, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x43, 0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72,
	0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x0f, 0x75, 0x73,
	0x65, 0x72, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x73, 0x22, 0x94, 0x01,
	0x0a, 0x0c, 0x4d, 0x73, 0x67, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x4f, 0x70, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x6e, 0x6f, 0x44, 0x69, 0x73, 0x74, 0x75, 0x72, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x6e, 0x6f, 0x44, 0x69, 0x73, 0x74, 0x75, 0x72, 0x62, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x70,
	0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x6f, 0x75, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x69,
	0x62, 0x72, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x76, 0x69, 0x62,
	0x72, 0x61, 0x74, 0x65, 0x22, 0x44, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x53, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x12, 0x2b, 0x0a,
	0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x52,
	0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x22, 0xc7, 0x01, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x41, 0x70, 0x70, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x41, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x70, 0x70, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x32, 0xc8, 0x04, 0x0a, 0x09, 0x69, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0d, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x43,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x42,
	0x65, 0x66, 0x6f, 0x72, 0x65, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x33, 0x0a, 0x0c, 0x41, 0x66, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x66, 0x74, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x65,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x39, 0x0a, 0x0f, 0x41, 0x66, 0x74, 0x65, 0x72, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x66,
	0x74, 0x65, 0x72, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x39, 0x0a, 0x0c, 0x4b, 0x69, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e,
	0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x4b, 0x69, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f,
	0x6e, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4b, 0x69, 0x63, 0x6b, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x36, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x13,
	0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x48, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x52,
	0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x57, 0x0a,
	0x16, 0x42, 0x61, 0x74, 0x63, 0x68, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e, 0x6e, 0x12, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x43,
	0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x1e, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73,
	0x67, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65,
	0x71, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x4b, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x53, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x70, 0x70, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x53,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x42,
	0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_im_proto_rawDescOnce sync.Once
	file_im_proto_rawDescData = file_im_proto_rawDesc
)

func file_im_proto_rawDescGZIP() []byte {
	file_im_proto_rawDescOnce.Do(func() {
		file_im_proto_rawDescData = protoimpl.X.CompressGZIP(file_im_proto_rawDescData)
	})
	return file_im_proto_rawDescData
}

var file_im_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_im_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_im_proto_goTypes = []interface{}{
	(ImMQBody_Event)(0),                // 0: pb.ImMQBody.Event
	(*ImMQBody)(nil),                   // 1: pb.ImMQBody
	(*BeforeConnectReq)(nil),           // 2: pb.BeforeConnectReq
	(*BeforeConnectResp)(nil),          // 3: pb.BeforeConnectResp
	(*GetUserLatestConnReq)(nil),       // 4: pb.GetUserLatestConnReq
	(*GetUserLatestConnResp)(nil),      // 5: pb.GetUserLatestConnResp
	(*BatchGetUserLatestConnReq)(nil),  // 6: pb.BatchGetUserLatestConnReq
	(*BatchGetUserLatestConnResp)(nil), // 7: pb.BatchGetUserLatestConnResp
	(*MsgNotifyOpt)(nil),               // 8: pb.MsgNotifyOpt
	(*GetAppSystemConfigReq)(nil),      // 9: pb.GetAppSystemConfigReq
	(*GetAppSystemConfigResp)(nil),     // 10: pb.GetAppSystemConfigResp
	nil,                                // 11: pb.GetAppSystemConfigResp.ConfigsEntry
	(*ConnParam)(nil),                  // 12: pb.ConnParam
	(*IpRegion)(nil),                   // 13: pb.IpRegion
	(*CommonResp)(nil),                 // 14: pb.CommonResp
	(*CommonReq)(nil),                  // 15: pb.CommonReq
	(*AfterConnectReq)(nil),            // 16: pb.AfterConnectReq
	(*AfterDisconnectReq)(nil),         // 17: pb.AfterDisconnectReq
	(*KickUserConnReq)(nil),            // 18: pb.KickUserConnReq
	(*GetUserConnReq)(nil),             // 19: pb.GetUserConnReq
	(*SendMsgReq)(nil),                 // 20: pb.SendMsgReq
	(*KickUserConnResp)(nil),           // 21: pb.KickUserConnResp
	(*GetUserConnResp)(nil),            // 22: pb.GetUserConnResp
	(*SendMsgResp)(nil),                // 23: pb.SendMsgResp
}
var file_im_proto_depIdxs = []int32{
	0,  // 0: pb.ImMQBody.event:type_name -> pb.ImMQBody.Event
	12, // 1: pb.BeforeConnectReq.connParam:type_name -> pb.ConnParam
	13, // 2: pb.GetUserLatestConnResp.ipRegion:type_name -> pb.IpRegion
	14, // 3: pb.BatchGetUserLatestConnResp.commonResp:type_name -> pb.CommonResp
	5,  // 4: pb.BatchGetUserLatestConnResp.userLatestConns:type_name -> pb.GetUserLatestConnResp
	15, // 5: pb.GetAppSystemConfigReq.commonReq:type_name -> pb.CommonReq
	14, // 6: pb.GetAppSystemConfigResp.commonResp:type_name -> pb.CommonResp
	11, // 7: pb.GetAppSystemConfigResp.configs:type_name -> pb.GetAppSystemConfigResp.ConfigsEntry
	2,  // 8: pb.imService.BeforeConnect:input_type -> pb.BeforeConnectReq
	16, // 9: pb.imService.AfterConnect:input_type -> pb.AfterConnectReq
	17, // 10: pb.imService.AfterDisconnect:input_type -> pb.AfterDisconnectReq
	18, // 11: pb.imService.KickUserConn:input_type -> pb.KickUserConnReq
	19, // 12: pb.imService.GetUserConn:input_type -> pb.GetUserConnReq
	4,  // 13: pb.imService.GetUserLatestConn:input_type -> pb.GetUserLatestConnReq
	6,  // 14: pb.imService.BatchGetUserLatestConn:input_type -> pb.BatchGetUserLatestConnReq
	20, // 15: pb.imService.SendMsg:input_type -> pb.SendMsgReq
	9,  // 16: pb.imService.GetAppSystemConfig:input_type -> pb.GetAppSystemConfigReq
	3,  // 17: pb.imService.BeforeConnect:output_type -> pb.BeforeConnectResp
	14, // 18: pb.imService.AfterConnect:output_type -> pb.CommonResp
	14, // 19: pb.imService.AfterDisconnect:output_type -> pb.CommonResp
	21, // 20: pb.imService.KickUserConn:output_type -> pb.KickUserConnResp
	22, // 21: pb.imService.GetUserConn:output_type -> pb.GetUserConnResp
	5,  // 22: pb.imService.GetUserLatestConn:output_type -> pb.GetUserLatestConnResp
	7,  // 23: pb.imService.BatchGetUserLatestConn:output_type -> pb.BatchGetUserLatestConnResp
	23, // 24: pb.imService.SendMsg:output_type -> pb.SendMsgResp
	10, // 25: pb.imService.GetAppSystemConfig:output_type -> pb.GetAppSystemConfigResp
	17, // [17:26] is the sub-list for method output_type
	8,  // [8:17] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_im_proto_init() }
func file_im_proto_init() {
	if File_im_proto != nil {
		return
	}
	file_common_proto_init()
	file_conn_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_im_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ImMQBody); i {
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
		file_im_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeforeConnectReq); i {
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
		file_im_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeforeConnectResp); i {
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
		file_im_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserLatestConnReq); i {
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
		file_im_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserLatestConnResp); i {
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
		file_im_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetUserLatestConnReq); i {
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
		file_im_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchGetUserLatestConnResp); i {
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
		file_im_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MsgNotifyOpt); i {
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
		file_im_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppSystemConfigReq); i {
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
		file_im_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAppSystemConfigResp); i {
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
			RawDescriptor: file_im_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_im_proto_goTypes,
		DependencyIndexes: file_im_proto_depIdxs,
		EnumInfos:         file_im_proto_enumTypes,
		MessageInfos:      file_im_proto_msgTypes,
	}.Build()
	File_im_proto = out.File
	file_im_proto_rawDesc = nil
	file_im_proto_goTypes = nil
	file_im_proto_depIdxs = nil
}
