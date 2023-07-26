// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: common/events/v1/service.proto

package pb_common_events

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

// Get single event by id.
type EventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
}

func (x *EventRequest) Reset() {
	*x = EventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_events_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventRequest) ProtoMessage() {}

func (x *EventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_events_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventRequest.ProtoReflect.Descriptor instead.
func (*EventRequest) Descriptor() ([]byte, []int) {
	return file_common_events_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *EventRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

type EventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Event `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EventResponse) Reset() {
	*x = EventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_events_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventResponse) ProtoMessage() {}

func (x *EventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_events_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventResponse.ProtoReflect.Descriptor instead.
func (*EventResponse) Descriptor() ([]byte, []int) {
	return file_common_events_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *EventResponse) GetData() *Event {
	if x != nil {
		return x.Data
	}
	return nil
}

// Get New Events since a given date.
type NewEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Since uint64 `protobuf:"varint,1,opt,name=since,proto3" json:"since,omitempty"`
}

func (x *NewEventsRequest) Reset() {
	*x = NewEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_events_v1_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewEventsRequest) ProtoMessage() {}

func (x *NewEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_events_v1_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewEventsRequest.ProtoReflect.Descriptor instead.
func (*NewEventsRequest) Descriptor() ([]byte, []int) {
	return file_common_events_v1_service_proto_rawDescGZIP(), []int{2}
}

func (x *NewEventsRequest) GetSince() uint64 {
	if x != nil {
		return x.Since
	}
	return 0
}

type NewEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Event `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *NewEventsResponse) Reset() {
	*x = NewEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_events_v1_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewEventsResponse) ProtoMessage() {}

func (x *NewEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_events_v1_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewEventsResponse.ProtoReflect.Descriptor instead.
func (*NewEventsResponse) Descriptor() ([]byte, []int) {
	return file_common_events_v1_service_proto_rawDescGZIP(), []int{3}
}

func (x *NewEventsResponse) GetData() *Event {
	if x != nil {
		return x.Data
	}
	return nil
}

// Record User Events stream.
type RecordUserEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Event `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *RecordUserEventsRequest) Reset() {
	*x = RecordUserEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_events_v1_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordUserEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordUserEventsRequest) ProtoMessage() {}

func (x *RecordUserEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_events_v1_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordUserEventsRequest.ProtoReflect.Descriptor instead.
func (*RecordUserEventsRequest) Descriptor() ([]byte, []int) {
	return file_common_events_v1_service_proto_rawDescGZIP(), []int{4}
}

func (x *RecordUserEventsRequest) GetData() []*Event {
	if x != nil {
		return x.Data
	}
	return nil
}

type RecordUserEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *RecordUserEventsResponse) Reset() {
	*x = RecordUserEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_events_v1_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordUserEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordUserEventsResponse) ProtoMessage() {}

func (x *RecordUserEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_events_v1_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordUserEventsResponse.ProtoReflect.Descriptor instead.
func (*RecordUserEventsResponse) Descriptor() ([]byte, []int) {
	return file_common_events_v1_service_proto_rawDescGZIP(), []int{5}
}

func (x *RecordUserEventsResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

// Events stream.
type EventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Event `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EventsRequest) Reset() {
	*x = EventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_events_v1_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventsRequest) ProtoMessage() {}

func (x *EventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_common_events_v1_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventsRequest.ProtoReflect.Descriptor instead.
func (*EventsRequest) Descriptor() ([]byte, []int) {
	return file_common_events_v1_service_proto_rawDescGZIP(), []int{6}
}

func (x *EventsRequest) GetData() *Event {
	if x != nil {
		return x.Data
	}
	return nil
}

type EventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Event `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EventsResponse) Reset() {
	*x = EventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_common_events_v1_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventsResponse) ProtoMessage() {}

func (x *EventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_common_events_v1_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventsResponse.ProtoReflect.Descriptor instead.
func (*EventsResponse) Descriptor() ([]byte, []int) {
	return file_common_events_v1_service_proto_rawDescGZIP(), []int{7}
}

func (x *EventsResponse) GetData() *Event {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_common_events_v1_service_proto protoreflect.FileDescriptor

var file_common_events_v1_service_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x1a, 0x1c, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x29, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x0d, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x28, 0x0a, 0x10, 0x4e, 0x65, 0x77,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x73, 0x69,
	0x6e, 0x63, 0x65, 0x22, 0x40, 0x0a, 0x11, 0x4e, 0x65, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x46, 0x0a, 0x17, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2a, 0x0a,
	0x18, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x3c, 0x0a, 0x0d, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3d, 0x0a, 0x0e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xee, 0x02, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x12, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x56, 0x0a, 0x09, 0x4e, 0x65, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x22,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4e, 0x65, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x65, 0x77, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x6b, 0x0a, 0x10, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x29, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x55, 0x73, 0x65, 0x72, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x4f, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x55, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x68, 0x65, 0x61, 0x72, 0x74, 0x6f, 0x66, 0x65, 0x64,
	0x75, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x73,
	0x2d, 0x72, 0x65, 0x70, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62,
	0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_common_events_v1_service_proto_rawDescOnce sync.Once
	file_common_events_v1_service_proto_rawDescData = file_common_events_v1_service_proto_rawDesc
)

func file_common_events_v1_service_proto_rawDescGZIP() []byte {
	file_common_events_v1_service_proto_rawDescOnce.Do(func() {
		file_common_events_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_common_events_v1_service_proto_rawDescData)
	})
	return file_common_events_v1_service_proto_rawDescData
}

var file_common_events_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_common_events_v1_service_proto_goTypes = []interface{}{
	(*EventRequest)(nil),             // 0: common.events.v1.EventRequest
	(*EventResponse)(nil),            // 1: common.events.v1.EventResponse
	(*NewEventsRequest)(nil),         // 2: common.events.v1.NewEventsRequest
	(*NewEventsResponse)(nil),        // 3: common.events.v1.NewEventsResponse
	(*RecordUserEventsRequest)(nil),  // 4: common.events.v1.RecordUserEventsRequest
	(*RecordUserEventsResponse)(nil), // 5: common.events.v1.RecordUserEventsResponse
	(*EventsRequest)(nil),            // 6: common.events.v1.EventsRequest
	(*EventsResponse)(nil),           // 7: common.events.v1.EventsResponse
	(*Event)(nil),                    // 8: common.events.v1.Event
}
var file_common_events_v1_service_proto_depIdxs = []int32{
	8, // 0: common.events.v1.EventResponse.data:type_name -> common.events.v1.Event
	8, // 1: common.events.v1.NewEventsResponse.data:type_name -> common.events.v1.Event
	8, // 2: common.events.v1.RecordUserEventsRequest.data:type_name -> common.events.v1.Event
	8, // 3: common.events.v1.EventsRequest.data:type_name -> common.events.v1.Event
	8, // 4: common.events.v1.EventsResponse.data:type_name -> common.events.v1.Event
	0, // 5: common.events.v1.EventService.Event:input_type -> common.events.v1.EventRequest
	2, // 6: common.events.v1.EventService.NewEvents:input_type -> common.events.v1.NewEventsRequest
	4, // 7: common.events.v1.EventService.RecordUserEvents:input_type -> common.events.v1.RecordUserEventsRequest
	6, // 8: common.events.v1.EventService.Events:input_type -> common.events.v1.EventsRequest
	1, // 9: common.events.v1.EventService.Event:output_type -> common.events.v1.EventResponse
	3, // 10: common.events.v1.EventService.NewEvents:output_type -> common.events.v1.NewEventsResponse
	5, // 11: common.events.v1.EventService.RecordUserEvents:output_type -> common.events.v1.RecordUserEventsResponse
	7, // 12: common.events.v1.EventService.Events:output_type -> common.events.v1.EventsResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_common_events_v1_service_proto_init() }
func file_common_events_v1_service_proto_init() {
	if File_common_events_v1_service_proto != nil {
		return
	}
	file_common_events_v1_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_common_events_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventRequest); i {
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
		file_common_events_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventResponse); i {
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
		file_common_events_v1_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewEventsRequest); i {
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
		file_common_events_v1_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewEventsResponse); i {
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
		file_common_events_v1_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordUserEventsRequest); i {
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
		file_common_events_v1_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordUserEventsResponse); i {
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
		file_common_events_v1_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventsRequest); i {
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
		file_common_events_v1_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventsResponse); i {
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
			RawDescriptor: file_common_events_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_common_events_v1_service_proto_goTypes,
		DependencyIndexes: file_common_events_v1_service_proto_depIdxs,
		MessageInfos:      file_common_events_v1_service_proto_msgTypes,
	}.Build()
	File_common_events_v1_service_proto = out.File
	file_common_events_v1_service_proto_rawDesc = nil
	file_common_events_v1_service_proto_goTypes = nil
	file_common_events_v1_service_proto_depIdxs = nil
}
