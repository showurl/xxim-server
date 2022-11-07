// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.5
// source: conn.proto

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

type PushEvent int32

const (
	PushEvent_PushMsgDataList PushEvent = 0
)

// Enum value maps for PushEvent.
var (
	PushEvent_name = map[int32]string{
		0: "PushMsgDataList",
	}
	PushEvent_value = map[string]int32{
		"PushMsgDataList": 0,
	}
)

func (x PushEvent) Enum() *PushEvent {
	p := new(PushEvent)
	*p = x
	return p
}

func (x PushEvent) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PushEvent) Descriptor() protoreflect.EnumDescriptor {
	return file_conn_proto_enumTypes[0].Descriptor()
}

func (PushEvent) Type() protoreflect.EnumType {
	return &file_conn_proto_enumTypes[0]
}

func (x PushEvent) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PushEvent.Descriptor instead.
func (PushEvent) EnumDescriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{0}
}

// 推送事件
type ConnMQBody_Event int32

const (
	ConnMQBody_Unknown ConnMQBody_Event = 0
)

// Enum value maps for ConnMQBody_Event.
var (
	ConnMQBody_Event_name = map[int32]string{
		0: "Unknown",
	}
	ConnMQBody_Event_value = map[string]int32{
		"Unknown": 0,
	}
)

func (x ConnMQBody_Event) Enum() *ConnMQBody_Event {
	p := new(ConnMQBody_Event)
	*p = x
	return p
}

func (x ConnMQBody_Event) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConnMQBody_Event) Descriptor() protoreflect.EnumDescriptor {
	return file_conn_proto_enumTypes[1].Descriptor()
}

func (ConnMQBody_Event) Type() protoreflect.EnumType {
	return &file_conn_proto_enumTypes[1]
}

func (x ConnMQBody_Event) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConnMQBody_Event.Descriptor instead.
func (ConnMQBody_Event) EnumDescriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{0, 0}
}

// 服务器通过websocket发送给客户端的消息体
type ConnMQBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event ConnMQBody_Event `protobuf:"varint,1,opt,name=event,proto3,enum=pb.ConnMQBody_Event" json:"event"`
	Data  []byte           `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *ConnMQBody) Reset() {
	*x = ConnMQBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnMQBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnMQBody) ProtoMessage() {}

func (x *ConnMQBody) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnMQBody.ProtoReflect.Descriptor instead.
func (*ConnMQBody) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{0}
}

func (x *ConnMQBody) GetEvent() ConnMQBody_Event {
	if x != nil {
		return x.Event
	}
	return ConnMQBody_Unknown
}

func (x *ConnMQBody) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ConnParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string            `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId"`
	Token       string            `protobuf:"bytes,2,opt,name=token,proto3" json:"token"`
	DeviceId    string            `protobuf:"bytes,3,opt,name=deviceId,proto3" json:"deviceId"`
	Platform    string            `protobuf:"bytes,4,opt,name=platform,proto3" json:"platform"`
	Ips         string            `protobuf:"bytes,5,opt,name=ips,proto3" json:"ips"`
	NetworkUsed string            `protobuf:"bytes,6,opt,name=networkUsed,proto3" json:"networkUsed"`
	Headers     map[string]string `protobuf:"bytes,7,rep,name=headers,proto3" json:"headers" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	PodIp       string            `protobuf:"bytes,8,opt,name=podIp,proto3" json:"podIp"`
}

func (x *ConnParam) Reset() {
	*x = ConnParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnParam) ProtoMessage() {}

func (x *ConnParam) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnParam.ProtoReflect.Descriptor instead.
func (*ConnParam) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{1}
}

func (x *ConnParam) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ConnParam) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ConnParam) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *ConnParam) GetPlatform() string {
	if x != nil {
		return x.Platform
	}
	return ""
}

func (x *ConnParam) GetIps() string {
	if x != nil {
		return x.Ips
	}
	return ""
}

func (x *ConnParam) GetNetworkUsed() string {
	if x != nil {
		return x.NetworkUsed
	}
	return ""
}

func (x *ConnParam) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *ConnParam) GetPodIp() string {
	if x != nil {
		return x.PodIp
	}
	return ""
}

// 获取用户的连接 可以用userId、platforms、devices过滤
type GetUserConnReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds   []string `protobuf:"bytes,1,rep,name=userIds,proto3" json:"userIds"`
	Platforms []string `protobuf:"bytes,2,rep,name=platforms,proto3" json:"platforms"`
	Devices   []string `protobuf:"bytes,3,rep,name=devices,proto3" json:"devices"`
}

func (x *GetUserConnReq) Reset() {
	*x = GetUserConnReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserConnReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserConnReq) ProtoMessage() {}

func (x *GetUserConnReq) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserConnReq.ProtoReflect.Descriptor instead.
func (*GetUserConnReq) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserConnReq) GetUserIds() []string {
	if x != nil {
		return x.UserIds
	}
	return nil
}

func (x *GetUserConnReq) GetPlatforms() []string {
	if x != nil {
		return x.Platforms
	}
	return nil
}

func (x *GetUserConnReq) GetDevices() []string {
	if x != nil {
		return x.Devices
	}
	return nil
}

type GetUserConnResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonResp *CommonResp  `protobuf:"bytes,1,opt,name=commonResp,proto3" json:"commonResp"`
	ConnParams []*ConnParam `protobuf:"bytes,2,rep,name=connParams,proto3" json:"connParams"`
}

func (x *GetUserConnResp) Reset() {
	*x = GetUserConnResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserConnResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserConnResp) ProtoMessage() {}

func (x *GetUserConnResp) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserConnResp.ProtoReflect.Descriptor instead.
func (*GetUserConnResp) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserConnResp) GetCommonResp() *CommonResp {
	if x != nil {
		return x.CommonResp
	}
	return nil
}

func (x *GetUserConnResp) GetConnParams() []*ConnParam {
	if x != nil {
		return x.ConnParams
	}
	return nil
}

type KickUserConnReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GetUserConnReq *GetUserConnReq `protobuf:"bytes,1,opt,name=getUserConnReq,proto3" json:"getUserConnReq"` // 搜索用户的连接
}

func (x *KickUserConnReq) Reset() {
	*x = KickUserConnReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KickUserConnReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickUserConnReq) ProtoMessage() {}

func (x *KickUserConnReq) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickUserConnReq.ProtoReflect.Descriptor instead.
func (*KickUserConnReq) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{4}
}

func (x *KickUserConnReq) GetGetUserConnReq() *GetUserConnReq {
	if x != nil {
		return x.GetUserConnReq
	}
	return nil
}

type KickUserConnResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonResp *CommonResp `protobuf:"bytes,1,opt,name=commonResp,proto3" json:"commonResp"`
}

func (x *KickUserConnResp) Reset() {
	*x = KickUserConnResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KickUserConnResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KickUserConnResp) ProtoMessage() {}

func (x *KickUserConnResp) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KickUserConnResp.ProtoReflect.Descriptor instead.
func (*KickUserConnResp) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{5}
}

func (x *KickUserConnResp) GetCommonResp() *CommonResp {
	if x != nil {
		return x.CommonResp
	}
	return nil
}

type PushBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event PushEvent `protobuf:"varint,1,opt,name=event,proto3,enum=pb.PushEvent" json:"event"`
	Data  []byte    `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
}

func (x *PushBody) Reset() {
	*x = PushBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushBody) ProtoMessage() {}

func (x *PushBody) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushBody.ProtoReflect.Descriptor instead.
func (*PushBody) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{6}
}

func (x *PushBody) GetEvent() PushEvent {
	if x != nil {
		return x.Event
	}
	return PushEvent_PushMsgDataList
}

func (x *PushBody) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SendMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GetUserConnReq *GetUserConnReq `protobuf:"bytes,1,opt,name=getUserConnReq,proto3" json:"getUserConnReq"` // 搜索用户的连接
	Event          PushEvent       `protobuf:"varint,2,opt,name=event,proto3,enum=pb.PushEvent" json:"event"`
	Data           []byte          `protobuf:"bytes,3,opt,name=data,proto3" json:"data"` // 发送的数据
}

func (x *SendMsgReq) Reset() {
	*x = SendMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsgReq) ProtoMessage() {}

func (x *SendMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsgReq.ProtoReflect.Descriptor instead.
func (*SendMsgReq) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{7}
}

func (x *SendMsgReq) GetGetUserConnReq() *GetUserConnReq {
	if x != nil {
		return x.GetUserConnReq
	}
	return nil
}

func (x *SendMsgReq) GetEvent() PushEvent {
	if x != nil {
		return x.Event
	}
	return PushEvent_PushMsgDataList
}

func (x *SendMsgReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SendMsgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommonResp *CommonResp `protobuf:"bytes,1,opt,name=commonResp,proto3" json:"commonResp"`
	// 发送成功的连接
	SuccessConnParams []*ConnParam `protobuf:"bytes,2,rep,name=successConnParams,proto3" json:"successConnParams"`
	// 发送失败的连接
	FailedConnParams []*ConnParam `protobuf:"bytes,3,rep,name=failedConnParams,proto3" json:"failedConnParams"`
}

func (x *SendMsgResp) Reset() {
	*x = SendMsgResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conn_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsgResp) ProtoMessage() {}

func (x *SendMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_conn_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsgResp.ProtoReflect.Descriptor instead.
func (*SendMsgResp) Descriptor() ([]byte, []int) {
	return file_conn_proto_rawDescGZIP(), []int{8}
}

func (x *SendMsgResp) GetCommonResp() *CommonResp {
	if x != nil {
		return x.CommonResp
	}
	return nil
}

func (x *SendMsgResp) GetSuccessConnParams() []*ConnParam {
	if x != nil {
		return x.SuccessConnParams
	}
	return nil
}

func (x *SendMsgResp) GetFailedConnParams() []*ConnParam {
	if x != nil {
		return x.FailedConnParams
	}
	return nil
}

var File_conn_proto protoreflect.FileDescriptor

var file_conn_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x1a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62,
	0x0a, 0x0a, 0x43, 0x6f, 0x6e, 0x6e, 0x4d, 0x51, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x2a, 0x0a, 0x05,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x4d, 0x51, 0x42, 0x6f, 0x64, 0x79, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x14, 0x0a, 0x05,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x10, 0x00, 0x22, 0xad, 0x02, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x70, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x69, 0x70, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x55, 0x73, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x55, 0x73, 0x65, 0x64, 0x12, 0x34, 0x0a, 0x07, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x6f, 0x64, 0x49, 0x70, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x70, 0x6f, 0x64, 0x49, 0x70, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x22, 0x62, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x6e, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x70, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x0a, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2d, 0x0a, 0x0a, 0x63, 0x6f, 0x6e,
	0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x22, 0x4d, 0x0a, 0x0f, 0x4b, 0x69, 0x63, 0x6b,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x3a, 0x0a, 0x0e, 0x67,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x71, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x71, 0x52, 0x0e, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x71, 0x22, 0x42, 0x0a, 0x10, 0x4b, 0x69, 0x63, 0x6b, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0a, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x52,
	0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x43, 0x0a, 0x08, 0x50,
	0x75, 0x73, 0x68, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x75, 0x73, 0x68,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x81, 0x01, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12,
	0x3a, 0x0a, 0x0e, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65,
	0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x71, 0x52, 0x0e, 0x67, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x75, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x22, 0xb5, 0x01, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x3b, 0x0a, 0x11, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43,
	0x6f, 0x6e, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x11,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x6e, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x73, 0x12, 0x39, 0x0a, 0x10, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x52, 0x10, 0x66, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x2a, 0x20, 0x0a, 0x09,
	0x50, 0x75, 0x73, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x13, 0x0a, 0x0f, 0x50, 0x75, 0x73,
	0x68, 0x4d, 0x73, 0x67, 0x44, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x10, 0x00, 0x32, 0xac,
	0x01, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39,
	0x0a, 0x0c, 0x4b, 0x69, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x12, 0x13,
	0x2e, 0x70, 0x62, 0x2e, 0x4b, 0x69, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e,
	0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x4b, 0x69, 0x63, 0x6b, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x12, 0x36, 0x0a, 0x0b, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x2a, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x12, 0x0e, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x70,
	0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conn_proto_rawDescOnce sync.Once
	file_conn_proto_rawDescData = file_conn_proto_rawDesc
)

func file_conn_proto_rawDescGZIP() []byte {
	file_conn_proto_rawDescOnce.Do(func() {
		file_conn_proto_rawDescData = protoimpl.X.CompressGZIP(file_conn_proto_rawDescData)
	})
	return file_conn_proto_rawDescData
}

var file_conn_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_conn_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_conn_proto_goTypes = []interface{}{
	(PushEvent)(0),           // 0: pb.PushEvent
	(ConnMQBody_Event)(0),    // 1: pb.ConnMQBody.Event
	(*ConnMQBody)(nil),       // 2: pb.ConnMQBody
	(*ConnParam)(nil),        // 3: pb.ConnParam
	(*GetUserConnReq)(nil),   // 4: pb.GetUserConnReq
	(*GetUserConnResp)(nil),  // 5: pb.GetUserConnResp
	(*KickUserConnReq)(nil),  // 6: pb.KickUserConnReq
	(*KickUserConnResp)(nil), // 7: pb.KickUserConnResp
	(*PushBody)(nil),         // 8: pb.PushBody
	(*SendMsgReq)(nil),       // 9: pb.SendMsgReq
	(*SendMsgResp)(nil),      // 10: pb.SendMsgResp
	nil,                      // 11: pb.ConnParam.HeadersEntry
	(*CommonResp)(nil),       // 12: pb.CommonResp
}
var file_conn_proto_depIdxs = []int32{
	1,  // 0: pb.ConnMQBody.event:type_name -> pb.ConnMQBody.Event
	11, // 1: pb.ConnParam.headers:type_name -> pb.ConnParam.HeadersEntry
	12, // 2: pb.GetUserConnResp.commonResp:type_name -> pb.CommonResp
	3,  // 3: pb.GetUserConnResp.connParams:type_name -> pb.ConnParam
	4,  // 4: pb.KickUserConnReq.getUserConnReq:type_name -> pb.GetUserConnReq
	12, // 5: pb.KickUserConnResp.commonResp:type_name -> pb.CommonResp
	0,  // 6: pb.PushBody.event:type_name -> pb.PushEvent
	4,  // 7: pb.SendMsgReq.getUserConnReq:type_name -> pb.GetUserConnReq
	0,  // 8: pb.SendMsgReq.event:type_name -> pb.PushEvent
	12, // 9: pb.SendMsgResp.commonResp:type_name -> pb.CommonResp
	3,  // 10: pb.SendMsgResp.successConnParams:type_name -> pb.ConnParam
	3,  // 11: pb.SendMsgResp.failedConnParams:type_name -> pb.ConnParam
	6,  // 12: pb.connService.KickUserConn:input_type -> pb.KickUserConnReq
	4,  // 13: pb.connService.GetUserConn:input_type -> pb.GetUserConnReq
	9,  // 14: pb.connService.SendMsg:input_type -> pb.SendMsgReq
	7,  // 15: pb.connService.KickUserConn:output_type -> pb.KickUserConnResp
	5,  // 16: pb.connService.GetUserConn:output_type -> pb.GetUserConnResp
	10, // 17: pb.connService.SendMsg:output_type -> pb.SendMsgResp
	15, // [15:18] is the sub-list for method output_type
	12, // [12:15] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_conn_proto_init() }
func file_conn_proto_init() {
	if File_conn_proto != nil {
		return
	}
	file_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_conn_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnMQBody); i {
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
		file_conn_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnParam); i {
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
		file_conn_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserConnReq); i {
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
		file_conn_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserConnResp); i {
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
		file_conn_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KickUserConnReq); i {
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
		file_conn_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KickUserConnResp); i {
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
		file_conn_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushBody); i {
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
		file_conn_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsgReq); i {
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
		file_conn_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsgResp); i {
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
			RawDescriptor: file_conn_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_conn_proto_goTypes,
		DependencyIndexes: file_conn_proto_depIdxs,
		EnumInfos:         file_conn_proto_enumTypes,
		MessageInfos:      file_conn_proto_msgTypes,
	}.Build()
	File_conn_proto = out.File
	file_conn_proto_rawDesc = nil
	file_conn_proto_goTypes = nil
	file_conn_proto_depIdxs = nil
}