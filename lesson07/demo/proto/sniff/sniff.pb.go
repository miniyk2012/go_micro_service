// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.29.3
// source: sniff/sniff.proto

package sniff

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

type InstanceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ip
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	// 端口
	Port int32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	// 集群
	Cluster string `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// 机房
	Idc string `protobuf:"bytes,4,opt,name=idc,proto3" json:"idc,omitempty"`
	// psm
	Psm string `protobuf:"bytes,5,opt,name=psm,proto3" json:"psm,omitempty"`
}

func (x *InstanceInfo) Reset() {
	*x = InstanceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sniff_sniff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstanceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstanceInfo) ProtoMessage() {}

func (x *InstanceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_sniff_sniff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstanceInfo.ProtoReflect.Descriptor instead.
func (*InstanceInfo) Descriptor() ([]byte, []int) {
	return file_sniff_sniff_proto_rawDescGZIP(), []int{0}
}

func (x *InstanceInfo) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *InstanceInfo) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *InstanceInfo) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *InstanceInfo) GetIdc() string {
	if x != nil {
		return x.Idc
	}
	return ""
}

func (x *InstanceInfo) GetPsm() string {
	if x != nil {
		return x.Psm
	}
	return ""
}

type ReqResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Thrift Http1
	Proto string `protobuf:"bytes,2,opt,name=proto,proto3" json:"proto,omitempty"`
	// 请求源
	From *InstanceInfo `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	// 请求目标
	To *InstanceInfo `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	// Thrift request logid
	Logid string `protobuf:"bytes,5,opt,name=logid,proto3" json:"logid,omitempty"`
	// request原始字节流
	Request []byte `protobuf:"bytes,6,opt,name=request,proto3" json:"request,omitempty"`
	// response原始字节流
	Response []byte `protobuf:"bytes,7,opt,name=response,proto3" json:"response,omitempty"`
	// 压测标记
	StressTag   string `protobuf:"bytes,8,opt,name=stress_tag,json=stressTag,proto3" json:"stress_tag,omitempty"`
	DdpTag      string `protobuf:"bytes,9,opt,name=ddp_tag,json=ddpTag,proto3" json:"ddp_tag,omitempty"`
	RingHashKey string `protobuf:"bytes,10,opt,name=ring_hash_key,json=ringHashKey,proto3" json:"ring_hash_key,omitempty"`
	// 方法名
	Method string `protobuf:"bytes,11,opt,name=method,proto3" json:"method,omitempty"`
	// 时间戳
	Timestamp int64 `protobuf:"varint,12,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// 状态码
	StatusCode int32 `protobuf:"varint,13,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	// 额外meta信息
	Meta map[string]string `protobuf:"bytes,14,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// 源env
	FromEnv string `protobuf:"bytes,15,opt,name=from_env,json=fromEnv,proto3" json:"from_env,omitempty"`
	// 目标env
	ToEnv        string            `protobuf:"bytes,16,opt,name=to_env,json=toEnv,proto3" json:"to_env,omitempty"`
	RequestMeta  map[string]string `protobuf:"bytes,17,rep,name=request_meta,json=requestMeta,proto3" json:"request_meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ResponseMeta map[string]string `protobuf:"bytes,18,rep,name=response_meta,json=responseMeta,proto3" json:"response_meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	RequestHash  string            `protobuf:"bytes,19,opt,name=request_hash,json=requestHash,proto3" json:"request_hash,omitempty"`
	PacketbusId  int64             `protobuf:"varint,20,opt,name=packetbus_id,json=packetbusId,proto3" json:"packetbus_id,omitempty"`
	IsEgress     bool              `protobuf:"varint,21,opt,name=is_egress,json=isEgress,proto3" json:"is_egress,omitempty"`
	// sniffer config id
	ConfigId int64 `protobuf:"varint,255,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	// packetbus task id
	TaskId int64 `protobuf:"varint,256,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *ReqResp) Reset() {
	*x = ReqResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sniff_sniff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqResp) ProtoMessage() {}

func (x *ReqResp) ProtoReflect() protoreflect.Message {
	mi := &file_sniff_sniff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqResp.ProtoReflect.Descriptor instead.
func (*ReqResp) Descriptor() ([]byte, []int) {
	return file_sniff_sniff_proto_rawDescGZIP(), []int{1}
}

func (x *ReqResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ReqResp) GetProto() string {
	if x != nil {
		return x.Proto
	}
	return ""
}

func (x *ReqResp) GetFrom() *InstanceInfo {
	if x != nil {
		return x.From
	}
	return nil
}

func (x *ReqResp) GetTo() *InstanceInfo {
	if x != nil {
		return x.To
	}
	return nil
}

func (x *ReqResp) GetLogid() string {
	if x != nil {
		return x.Logid
	}
	return ""
}

func (x *ReqResp) GetRequest() []byte {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *ReqResp) GetResponse() []byte {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *ReqResp) GetStressTag() string {
	if x != nil {
		return x.StressTag
	}
	return ""
}

func (x *ReqResp) GetDdpTag() string {
	if x != nil {
		return x.DdpTag
	}
	return ""
}

func (x *ReqResp) GetRingHashKey() string {
	if x != nil {
		return x.RingHashKey
	}
	return ""
}

func (x *ReqResp) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *ReqResp) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ReqResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ReqResp) GetMeta() map[string]string {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *ReqResp) GetFromEnv() string {
	if x != nil {
		return x.FromEnv
	}
	return ""
}

func (x *ReqResp) GetToEnv() string {
	if x != nil {
		return x.ToEnv
	}
	return ""
}

func (x *ReqResp) GetRequestMeta() map[string]string {
	if x != nil {
		return x.RequestMeta
	}
	return nil
}

func (x *ReqResp) GetResponseMeta() map[string]string {
	if x != nil {
		return x.ResponseMeta
	}
	return nil
}

func (x *ReqResp) GetRequestHash() string {
	if x != nil {
		return x.RequestHash
	}
	return ""
}

func (x *ReqResp) GetPacketbusId() int64 {
	if x != nil {
		return x.PacketbusId
	}
	return 0
}

func (x *ReqResp) GetIsEgress() bool {
	if x != nil {
		return x.IsEgress
	}
	return false
}

func (x *ReqResp) GetConfigId() int64 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *ReqResp) GetTaskId() int64 {
	if x != nil {
		return x.TaskId
	}
	return 0
}

var File_sniff_sniff_proto protoreflect.FileDescriptor

var file_sniff_sniff_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x6e, 0x69, 0x66, 0x66, 0x2f, 0x73, 0x6e, 0x69, 0x66, 0x66, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x74, 0x0a, 0x0c, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x63, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x63, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x73, 0x6d, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x73, 0x6d, 0x22, 0x9e, 0x07, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x0a, 0x04, 0x66,
	0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x1d,
	0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x49, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x6f, 0x67, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f,
	0x67, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x72,
	0x65, 0x73, 0x73, 0x5f, 0x74, 0x61, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x74, 0x72, 0x65, 0x73, 0x73, 0x54, 0x61, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x64, 0x64, 0x70, 0x5f,
	0x74, 0x61, 0x67, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x64, 0x70, 0x54, 0x61,
	0x67, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x6b,
	0x65, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x69, 0x6e, 0x67, 0x48, 0x61,
	0x73, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x26, 0x0a, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x52, 0x65, 0x71,
	0x52, 0x65, 0x73, 0x70, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x65, 0x6e, 0x76,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x72, 0x6f, 0x6d, 0x45, 0x6e, 0x76, 0x12,
	0x15, 0x0a, 0x06, 0x74, 0x6f, 0x5f, 0x65, 0x6e, 0x76, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x74, 0x6f, 0x45, 0x6e, 0x76, 0x12, 0x3c, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x11, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x52,
	0x65, 0x71, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x12, 0x3f, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x12, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x52, 0x65,
	0x71, 0x52, 0x65, 0x73, 0x70, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x62, 0x75, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x62, 0x75, 0x73, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x73, 0x5f, 0x65, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x15, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x45, 0x67, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0xff, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69,
	0x64, 0x18, 0x80, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64,
	0x1a, 0x37, 0x0a, 0x09, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3e, 0x0a, 0x10, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3f, 0x0a, 0x11, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x12, 0x5a, 0x10, 0x64, 0x65,
	0x6d, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6e, 0x69, 0x66, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sniff_sniff_proto_rawDescOnce sync.Once
	file_sniff_sniff_proto_rawDescData = file_sniff_sniff_proto_rawDesc
)

func file_sniff_sniff_proto_rawDescGZIP() []byte {
	file_sniff_sniff_proto_rawDescOnce.Do(func() {
		file_sniff_sniff_proto_rawDescData = protoimpl.X.CompressGZIP(file_sniff_sniff_proto_rawDescData)
	})
	return file_sniff_sniff_proto_rawDescData
}

var file_sniff_sniff_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_sniff_sniff_proto_goTypes = []interface{}{
	(*InstanceInfo)(nil), // 0: InstanceInfo
	(*ReqResp)(nil),      // 1: ReqResp
	nil,                  // 2: ReqResp.MetaEntry
	nil,                  // 3: ReqResp.RequestMetaEntry
	nil,                  // 4: ReqResp.ResponseMetaEntry
}
var file_sniff_sniff_proto_depIdxs = []int32{
	0, // 0: ReqResp.from:type_name -> InstanceInfo
	0, // 1: ReqResp.to:type_name -> InstanceInfo
	2, // 2: ReqResp.meta:type_name -> ReqResp.MetaEntry
	3, // 3: ReqResp.request_meta:type_name -> ReqResp.RequestMetaEntry
	4, // 4: ReqResp.response_meta:type_name -> ReqResp.ResponseMetaEntry
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_sniff_sniff_proto_init() }
func file_sniff_sniff_proto_init() {
	if File_sniff_sniff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sniff_sniff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstanceInfo); i {
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
		file_sniff_sniff_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqResp); i {
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
			RawDescriptor: file_sniff_sniff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sniff_sniff_proto_goTypes,
		DependencyIndexes: file_sniff_sniff_proto_depIdxs,
		MessageInfos:      file_sniff_sniff_proto_msgTypes,
	}.Build()
	File_sniff_sniff_proto = out.File
	file_sniff_sniff_proto_rawDesc = nil
	file_sniff_sniff_proto_goTypes = nil
	file_sniff_sniff_proto_depIdxs = nil
}
