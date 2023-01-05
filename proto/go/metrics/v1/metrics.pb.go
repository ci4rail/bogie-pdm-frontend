//
//Copyright © 2022 Ci4Rail GmbH <engineering@ci4rail.com>
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//http://www.apache.org/licenses/LICENSE-2.0
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.14.0
// source: metrics/v1/metrics.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Metrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ts          *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=ts,proto3" json:"ts,omitempty"`
	SteadyDrive *Metrics_SteadyDrive   `protobuf:"bytes,2,opt,name=steady_drive,json=steadyDrive,proto3" json:"steady_drive,omitempty"`
	Position    *Metrics_Position      `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
	Temperature *Metrics_Temperature   `protobuf:"bytes,4,opt,name=temperature,proto3" json:"temperature,omitempty"`
	Internet    *Metrics_Internet      `protobuf:"bytes,5,opt,name=internet,proto3" json:"internet,omitempty"`
	Cellular    *Metrics_Cellular      `protobuf:"bytes,6,opt,name=cellular,proto3" json:"cellular,omitempty"`
	GnssRaw     *Metrics_GnssRaw       `protobuf:"bytes,7,opt,name=gnss_raw,json=gnssRaw,proto3" json:"gnss_raw,omitempty"`
}

func (x *Metrics) Reset() {
	*x = Metrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metrics_v1_metrics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metrics) ProtoMessage() {}

func (x *Metrics) ProtoReflect() protoreflect.Message {
	mi := &file_metrics_v1_metrics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metrics.ProtoReflect.Descriptor instead.
func (*Metrics) Descriptor() ([]byte, []int) {
	return file_metrics_v1_metrics_proto_rawDescGZIP(), []int{0}
}

func (x *Metrics) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *Metrics) GetSteadyDrive() *Metrics_SteadyDrive {
	if x != nil {
		return x.SteadyDrive
	}
	return nil
}

func (x *Metrics) GetPosition() *Metrics_Position {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *Metrics) GetTemperature() *Metrics_Temperature {
	if x != nil {
		return x.Temperature
	}
	return nil
}

func (x *Metrics) GetInternet() *Metrics_Internet {
	if x != nil {
		return x.Internet
	}
	return nil
}

func (x *Metrics) GetCellular() *Metrics_Cellular {
	if x != nil {
		return x.Cellular
	}
	return nil
}

func (x *Metrics) GetGnssRaw() *Metrics_GnssRaw {
	if x != nil {
		return x.GnssRaw
	}
	return nil
}

type Metrics_SteadyDrive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Max []float64 `protobuf:"fixed64,1,rep,packed,name=max,proto3" json:"max,omitempty"`
	Rms []float64 `protobuf:"fixed64,2,rep,packed,name=rms,proto3" json:"rms,omitempty"`
}

func (x *Metrics_SteadyDrive) Reset() {
	*x = Metrics_SteadyDrive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metrics_v1_metrics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metrics_SteadyDrive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metrics_SteadyDrive) ProtoMessage() {}

func (x *Metrics_SteadyDrive) ProtoReflect() protoreflect.Message {
	mi := &file_metrics_v1_metrics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metrics_SteadyDrive.ProtoReflect.Descriptor instead.
func (*Metrics_SteadyDrive) Descriptor() ([]byte, []int) {
	return file_metrics_v1_metrics_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Metrics_SteadyDrive) GetMax() []float64 {
	if x != nil {
		return x.Max
	}
	return nil
}

func (x *Metrics_SteadyDrive) GetRms() []float64 {
	if x != nil {
		return x.Rms
	}
	return nil
}

type Metrics_Position struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid bool    `protobuf:"varint,1,opt,name=valid,proto3" json:"valid,omitempty"`
	Lat   float32 `protobuf:"fixed32,2,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon   float32 `protobuf:"fixed32,3,opt,name=lon,proto3" json:"lon,omitempty"`
	Alt   float32 `protobuf:"fixed32,4,opt,name=alt,proto3" json:"alt,omitempty"`
	Speed float32 `protobuf:"fixed32,5,opt,name=speed,proto3" json:"speed,omitempty"`
}

func (x *Metrics_Position) Reset() {
	*x = Metrics_Position{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metrics_v1_metrics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metrics_Position) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metrics_Position) ProtoMessage() {}

func (x *Metrics_Position) ProtoReflect() protoreflect.Message {
	mi := &file_metrics_v1_metrics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metrics_Position.ProtoReflect.Descriptor instead.
func (*Metrics_Position) Descriptor() ([]byte, []int) {
	return file_metrics_v1_metrics_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Metrics_Position) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

func (x *Metrics_Position) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Metrics_Position) GetLon() float32 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *Metrics_Position) GetAlt() float32 {
	if x != nil {
		return x.Alt
	}
	return 0
}

func (x *Metrics_Position) GetSpeed() float32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

type Metrics_Temperature struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InBox float32 `protobuf:"fixed32,1,opt,name=inBox,proto3" json:"inBox,omitempty"`
}

func (x *Metrics_Temperature) Reset() {
	*x = Metrics_Temperature{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metrics_v1_metrics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metrics_Temperature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metrics_Temperature) ProtoMessage() {}

func (x *Metrics_Temperature) ProtoReflect() protoreflect.Message {
	mi := &file_metrics_v1_metrics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metrics_Temperature.ProtoReflect.Descriptor instead.
func (*Metrics_Temperature) Descriptor() ([]byte, []int) {
	return file_metrics_v1_metrics_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Metrics_Temperature) GetInBox() float32 {
	if x != nil {
		return x.InBox
	}
	return 0
}

type Metrics_Internet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Connected bool `protobuf:"varint,1,opt,name=connected,proto3" json:"connected,omitempty"`
}

func (x *Metrics_Internet) Reset() {
	*x = Metrics_Internet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metrics_v1_metrics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metrics_Internet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metrics_Internet) ProtoMessage() {}

func (x *Metrics_Internet) ProtoReflect() protoreflect.Message {
	mi := &file_metrics_v1_metrics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metrics_Internet.ProtoReflect.Descriptor instead.
func (*Metrics_Internet) Descriptor() ([]byte, []int) {
	return file_metrics_v1_metrics_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Metrics_Internet) GetConnected() bool {
	if x != nil {
		return x.Connected
	}
	return false
}

type Metrics_Cellular struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Operator string  `protobuf:"bytes,1,opt,name=operator,proto3" json:"operator,omitempty"`
	Strength float32 `protobuf:"fixed32,2,opt,name=strength,proto3" json:"strength,omitempty"`
}

func (x *Metrics_Cellular) Reset() {
	*x = Metrics_Cellular{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metrics_v1_metrics_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metrics_Cellular) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metrics_Cellular) ProtoMessage() {}

func (x *Metrics_Cellular) ProtoReflect() protoreflect.Message {
	mi := &file_metrics_v1_metrics_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metrics_Cellular.ProtoReflect.Descriptor instead.
func (*Metrics_Cellular) Descriptor() ([]byte, []int) {
	return file_metrics_v1_metrics_proto_rawDescGZIP(), []int{0, 4}
}

func (x *Metrics_Cellular) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *Metrics_Cellular) GetStrength() float32 {
	if x != nil {
		return x.Strength
	}
	return 0
}

type Metrics_GnssRaw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lat     float32 `protobuf:"fixed32,1,opt,name=lat,proto3" json:"lat,omitempty"`
	Lon     float32 `protobuf:"fixed32,2,opt,name=lon,proto3" json:"lon,omitempty"`
	Alt     float32 `protobuf:"fixed32,3,opt,name=alt,proto3" json:"alt,omitempty"`
	Speed   float32 `protobuf:"fixed32,4,opt,name=speed,proto3" json:"speed,omitempty"`
	Eph     float32 `protobuf:"fixed32,5,opt,name=eph,proto3" json:"eph,omitempty"`
	Mode    int32   `protobuf:"varint,6,opt,name=mode,proto3" json:"mode,omitempty"`
	Numsats int32   `protobuf:"varint,7,opt,name=numsats,proto3" json:"numsats,omitempty"`
}

func (x *Metrics_GnssRaw) Reset() {
	*x = Metrics_GnssRaw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metrics_v1_metrics_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metrics_GnssRaw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metrics_GnssRaw) ProtoMessage() {}

func (x *Metrics_GnssRaw) ProtoReflect() protoreflect.Message {
	mi := &file_metrics_v1_metrics_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metrics_GnssRaw.ProtoReflect.Descriptor instead.
func (*Metrics_GnssRaw) Descriptor() ([]byte, []int) {
	return file_metrics_v1_metrics_proto_rawDescGZIP(), []int{0, 5}
}

func (x *Metrics_GnssRaw) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Metrics_GnssRaw) GetLon() float32 {
	if x != nil {
		return x.Lon
	}
	return 0
}

func (x *Metrics_GnssRaw) GetAlt() float32 {
	if x != nil {
		return x.Alt
	}
	return 0
}

func (x *Metrics_GnssRaw) GetSpeed() float32 {
	if x != nil {
		return x.Speed
	}
	return 0
}

func (x *Metrics_GnssRaw) GetEph() float32 {
	if x != nil {
		return x.Eph
	}
	return 0
}

func (x *Metrics_GnssRaw) GetMode() int32 {
	if x != nil {
		return x.Mode
	}
	return 0
}

func (x *Metrics_GnssRaw) GetNumsats() int32 {
	if x != nil {
		return x.Numsats
	}
	return 0
}

var File_metrics_v1_metrics_proto protoreflect.FileDescriptor

var file_metrics_v1_metrics_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdc, 0x06, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x3f, 0x0a, 0x0c,
	0x73, 0x74, 0x65, 0x61, 0x64, 0x79, 0x5f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x74, 0x65, 0x61, 0x64, 0x79, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x52, 0x0b, 0x73, 0x74, 0x65, 0x61, 0x64, 0x79, 0x44, 0x72, 0x69, 0x76, 0x65, 0x12, 0x35, 0x0a,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x2e, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6d, 0x65, 0x74, 0x72,
	0x69, 0x63, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x54, 0x65, 0x6d, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x12, 0x35, 0x0a, 0x08, 0x63,
	0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e,
	0x43, 0x65, 0x6c, 0x6c, 0x75, 0x6c, 0x61, 0x72, 0x52, 0x08, 0x63, 0x65, 0x6c, 0x6c, 0x75, 0x6c,
	0x61, 0x72, 0x12, 0x33, 0x0a, 0x08, 0x67, 0x6e, 0x73, 0x73, 0x5f, 0x72, 0x61, 0x77, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x47, 0x6e, 0x73, 0x73, 0x52, 0x61, 0x77, 0x52, 0x07,
	0x67, 0x6e, 0x73, 0x73, 0x52, 0x61, 0x77, 0x1a, 0x31, 0x0a, 0x0b, 0x53, 0x74, 0x65, 0x61, 0x64,
	0x79, 0x44, 0x72, 0x69, 0x76, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x78, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x01, 0x52, 0x03, 0x6d, 0x61, 0x78, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x01, 0x52, 0x03, 0x72, 0x6d, 0x73, 0x1a, 0x6c, 0x0a, 0x08, 0x50, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x6c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x6c, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x6f, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x61, 0x6c, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x61,
	0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x1a, 0x23, 0x0a, 0x0b, 0x54, 0x65, 0x6d, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x42, 0x6f, 0x78,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x69, 0x6e, 0x42, 0x6f, 0x78, 0x1a, 0x28, 0x0a,
	0x08, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x1a, 0x42, 0x0a, 0x08, 0x43, 0x65, 0x6c, 0x6c, 0x75,
	0x6c, 0x61, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x73, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x08, 0x73, 0x74, 0x72, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x1a, 0x95, 0x01, 0x0a, 0x07,
	0x47, 0x6e, 0x73, 0x73, 0x52, 0x61, 0x77, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x61, 0x6c, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x73, 0x70,
	0x65, 0x65, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x70, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x03, 0x65, 0x70, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x75, 0x6d,
	0x73, 0x61, 0x74, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6e, 0x75, 0x6d, 0x73,
	0x61, 0x74, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2f, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metrics_v1_metrics_proto_rawDescOnce sync.Once
	file_metrics_v1_metrics_proto_rawDescData = file_metrics_v1_metrics_proto_rawDesc
)

func file_metrics_v1_metrics_proto_rawDescGZIP() []byte {
	file_metrics_v1_metrics_proto_rawDescOnce.Do(func() {
		file_metrics_v1_metrics_proto_rawDescData = protoimpl.X.CompressGZIP(file_metrics_v1_metrics_proto_rawDescData)
	})
	return file_metrics_v1_metrics_proto_rawDescData
}

var file_metrics_v1_metrics_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_metrics_v1_metrics_proto_goTypes = []interface{}{
	(*Metrics)(nil),               // 0: metrics.Metrics
	(*Metrics_SteadyDrive)(nil),   // 1: metrics.Metrics.SteadyDrive
	(*Metrics_Position)(nil),      // 2: metrics.Metrics.Position
	(*Metrics_Temperature)(nil),   // 3: metrics.Metrics.Temperature
	(*Metrics_Internet)(nil),      // 4: metrics.Metrics.Internet
	(*Metrics_Cellular)(nil),      // 5: metrics.Metrics.Cellular
	(*Metrics_GnssRaw)(nil),       // 6: metrics.Metrics.GnssRaw
	(*timestamppb.Timestamp)(nil), // 7: google.protobuf.Timestamp
}
var file_metrics_v1_metrics_proto_depIdxs = []int32{
	7, // 0: metrics.Metrics.ts:type_name -> google.protobuf.Timestamp
	1, // 1: metrics.Metrics.steady_drive:type_name -> metrics.Metrics.SteadyDrive
	2, // 2: metrics.Metrics.position:type_name -> metrics.Metrics.Position
	3, // 3: metrics.Metrics.temperature:type_name -> metrics.Metrics.Temperature
	4, // 4: metrics.Metrics.internet:type_name -> metrics.Metrics.Internet
	5, // 5: metrics.Metrics.cellular:type_name -> metrics.Metrics.Cellular
	6, // 6: metrics.Metrics.gnss_raw:type_name -> metrics.Metrics.GnssRaw
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_metrics_v1_metrics_proto_init() }
func file_metrics_v1_metrics_proto_init() {
	if File_metrics_v1_metrics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metrics_v1_metrics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metrics); i {
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
		file_metrics_v1_metrics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metrics_SteadyDrive); i {
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
		file_metrics_v1_metrics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metrics_Position); i {
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
		file_metrics_v1_metrics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metrics_Temperature); i {
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
		file_metrics_v1_metrics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metrics_Internet); i {
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
		file_metrics_v1_metrics_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metrics_Cellular); i {
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
		file_metrics_v1_metrics_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metrics_GnssRaw); i {
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
			RawDescriptor: file_metrics_v1_metrics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_metrics_v1_metrics_proto_goTypes,
		DependencyIndexes: file_metrics_v1_metrics_proto_depIdxs,
		MessageInfos:      file_metrics_v1_metrics_proto_msgTypes,
	}.Build()
	File_metrics_v1_metrics_proto = out.File
	file_metrics_v1_metrics_proto_rawDesc = nil
	file_metrics_v1_metrics_proto_goTypes = nil
	file_metrics_v1_metrics_proto_depIdxs = nil
}