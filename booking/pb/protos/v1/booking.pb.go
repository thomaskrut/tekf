// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: protos/v1/booking.proto

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

type EventType int32

const (
	EventType_EVENT_TYPE_NONE_UNSPECIFIED    EventType = 0
	EventType_EVENT_TYPE_BOOKING_CREATED     EventType = 1
	EventType_EVENT_TYPE_BOOKING_UPDATED     EventType = 2
	EventType_EVENT_TYPE_BOOKING_DELETED     EventType = 3
	EventType_EVENT_TYPE_BOOKING_CHECKED_IN  EventType = 4
	EventType_EVENT_TYPE_BOOKING_CHECKED_OUT EventType = 5
)

// Enum value maps for EventType.
var (
	EventType_name = map[int32]string{
		0: "EVENT_TYPE_NONE_UNSPECIFIED",
		1: "EVENT_TYPE_BOOKING_CREATED",
		2: "EVENT_TYPE_BOOKING_UPDATED",
		3: "EVENT_TYPE_BOOKING_DELETED",
		4: "EVENT_TYPE_BOOKING_CHECKED_IN",
		5: "EVENT_TYPE_BOOKING_CHECKED_OUT",
	}
	EventType_value = map[string]int32{
		"EVENT_TYPE_NONE_UNSPECIFIED":    0,
		"EVENT_TYPE_BOOKING_CREATED":     1,
		"EVENT_TYPE_BOOKING_UPDATED":     2,
		"EVENT_TYPE_BOOKING_DELETED":     3,
		"EVENT_TYPE_BOOKING_CHECKED_IN":  4,
		"EVENT_TYPE_BOOKING_CHECKED_OUT": 5,
	}
)

func (x EventType) Enum() *EventType {
	p := new(EventType)
	*p = x
	return p
}

func (x EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_v1_booking_proto_enumTypes[0].Descriptor()
}

func (EventType) Type() protoreflect.EnumType {
	return &file_protos_v1_booking_proto_enumTypes[0]
}

func (x EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EventType.Descriptor instead.
func (EventType) EnumDescriptor() ([]byte, []int) {
	return file_protos_v1_booking_proto_rawDescGZIP(), []int{0}
}

type BookingEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata  *Metadata `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	EventType EventType `protobuf:"varint,4,opt,name=event_type,json=eventType,proto3,enum=protos.v1.EventType" json:"event_type,omitempty"`
	Booking   *Booking  `protobuf:"bytes,5,opt,name=booking,proto3" json:"booking,omitempty"`
}

func (x *BookingEvent) Reset() {
	*x = BookingEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_booking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingEvent) ProtoMessage() {}

func (x *BookingEvent) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_booking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingEvent.ProtoReflect.Descriptor instead.
func (*BookingEvent) Descriptor() ([]byte, []int) {
	return file_protos_v1_booking_proto_rawDescGZIP(), []int{0}
}

func (x *BookingEvent) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *BookingEvent) GetEventType() EventType {
	if x != nil {
		return x.EventType
	}
	return EventType_EVENT_TYPE_NONE_UNSPECIFIED
}

func (x *BookingEvent) GetBooking() *Booking {
	if x != nil {
		return x.Booking
	}
	return nil
}

type Booking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	From   string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To     string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	Guests int32  `protobuf:"varint,4,opt,name=guests,proto3" json:"guests,omitempty"`
	Name   string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	UnitId int32  `protobuf:"varint,6,opt,name=unit_id,json=unitId,proto3" json:"unit_id,omitempty"`
}

func (x *Booking) Reset() {
	*x = Booking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_booking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Booking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Booking) ProtoMessage() {}

func (x *Booking) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_booking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Booking.ProtoReflect.Descriptor instead.
func (*Booking) Descriptor() ([]byte, []int) {
	return file_protos_v1_booking_proto_rawDescGZIP(), []int{1}
}

func (x *Booking) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Booking) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Booking) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Booking) GetGuests() int32 {
	if x != nil {
		return x.Guests
	}
	return 0
}

func (x *Booking) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Booking) GetUnitId() int32 {
	if x != nil {
		return x.UnitId
	}
	return 0
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *Metadata) Reset() {
	*x = Metadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_booking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metadata) ProtoMessage() {}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_booking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metadata.ProtoReflect.Descriptor instead.
func (*Metadata) Descriptor() ([]byte, []int) {
	return file_protos_v1_booking_proto_rawDescGZIP(), []int{2}
}

func (x *Metadata) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type WriteBookingEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingEvent *BookingEvent `protobuf:"bytes,1,opt,name=booking_event,json=bookingEvent,proto3" json:"booking_event,omitempty"`
}

func (x *WriteBookingEventRequest) Reset() {
	*x = WriteBookingEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_booking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteBookingEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteBookingEventRequest) ProtoMessage() {}

func (x *WriteBookingEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_booking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteBookingEventRequest.ProtoReflect.Descriptor instead.
func (*WriteBookingEventRequest) Descriptor() ([]byte, []int) {
	return file_protos_v1_booking_proto_rawDescGZIP(), []int{3}
}

func (x *WriteBookingEventRequest) GetBookingEvent() *BookingEvent {
	if x != nil {
		return x.BookingEvent
	}
	return nil
}

type WriteBookingEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *WriteBookingEventResponse) Reset() {
	*x = WriteBookingEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_booking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteBookingEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteBookingEventResponse) ProtoMessage() {}

func (x *WriteBookingEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_booking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteBookingEventResponse.ProtoReflect.Descriptor instead.
func (*WriteBookingEventResponse) Descriptor() ([]byte, []int) {
	return file_protos_v1_booking_proto_rawDescGZIP(), []int{4}
}

func (x *WriteBookingEventResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ReadBookingEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastKnownEventId int32 `protobuf:"varint,2,opt,name=last_known_event_id,json=lastKnownEventId,proto3" json:"last_known_event_id,omitempty"`
}

func (x *ReadBookingEventsRequest) Reset() {
	*x = ReadBookingEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_booking_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadBookingEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadBookingEventsRequest) ProtoMessage() {}

func (x *ReadBookingEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_booking_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadBookingEventsRequest.ProtoReflect.Descriptor instead.
func (*ReadBookingEventsRequest) Descriptor() ([]byte, []int) {
	return file_protos_v1_booking_proto_rawDescGZIP(), []int{5}
}

func (x *ReadBookingEventsRequest) GetLastKnownEventId() int32 {
	if x != nil {
		return x.LastKnownEventId
	}
	return 0
}

type ReadBookingEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingEvent *BookingEvent `protobuf:"bytes,1,opt,name=booking_event,json=bookingEvent,proto3" json:"booking_event,omitempty"`
	Revision     int32         `protobuf:"varint,2,opt,name=revision,proto3" json:"revision,omitempty"`
}

func (x *ReadBookingEventsResponse) Reset() {
	*x = ReadBookingEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_booking_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadBookingEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadBookingEventsResponse) ProtoMessage() {}

func (x *ReadBookingEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_booking_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadBookingEventsResponse.ProtoReflect.Descriptor instead.
func (*ReadBookingEventsResponse) Descriptor() ([]byte, []int) {
	return file_protos_v1_booking_proto_rawDescGZIP(), []int{6}
}

func (x *ReadBookingEventsResponse) GetBookingEvent() *BookingEvent {
	if x != nil {
		return x.BookingEvent
	}
	return nil
}

func (x *ReadBookingEventsResponse) GetRevision() int32 {
	if x != nil {
		return x.Revision
	}
	return 0
}

var File_protos_v1_booking_proto protoreflect.FileDescriptor

var file_protos_v1_booking_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x0c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2f, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x08, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x33, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x2c, 0x0a, 0x07,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x82, 0x01, 0x0a, 0x07, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x75, 0x65, 0x73,
	0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x6e, 0x69, 0x74, 0x49, 0x64, 0x22,
	0x44, 0x0a, 0x08, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x38, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x58, 0x0a, 0x18, 0x57, 0x72, 0x69, 0x74, 0x65, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3c, 0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x22,
	0x35, 0x0a, 0x19, 0x57, 0x72, 0x69, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x49, 0x0a, 0x18, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x2d, 0x0a, 0x13, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x10, 0x6c, 0x61, 0x73, 0x74, 0x4b, 0x6e, 0x6f, 0x77, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x75, 0x0a, 0x19, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c,
	0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x0c,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x2a, 0xd3, 0x01, 0x0a, 0x09, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x52,
	0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x55, 0x50,
	0x44, 0x41, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1e, 0x0a, 0x1a, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x12, 0x21, 0x0a, 0x1d, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4b, 0x49, 0x4e, 0x47, 0x5f, 0x43, 0x48,
	0x45, 0x43, 0x4b, 0x45, 0x44, 0x5f, 0x49, 0x4e, 0x10, 0x04, 0x12, 0x22, 0x0a, 0x1e, 0x45, 0x56,
	0x45, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4b, 0x49, 0x4e, 0x47,
	0x5f, 0x43, 0x48, 0x45, 0x43, 0x4b, 0x45, 0x44, 0x5f, 0x4f, 0x55, 0x54, 0x10, 0x05, 0x32, 0xd7,
	0x01, 0x0a, 0x13, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x11, 0x57, 0x72, 0x69, 0x74, 0x65, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x11, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x23, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x42, 0x11, 0x5a, 0x0f, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protos_v1_booking_proto_rawDescOnce sync.Once
	file_protos_v1_booking_proto_rawDescData = file_protos_v1_booking_proto_rawDesc
)

func file_protos_v1_booking_proto_rawDescGZIP() []byte {
	file_protos_v1_booking_proto_rawDescOnce.Do(func() {
		file_protos_v1_booking_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_v1_booking_proto_rawDescData)
	})
	return file_protos_v1_booking_proto_rawDescData
}

var file_protos_v1_booking_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_v1_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protos_v1_booking_proto_goTypes = []interface{}{
	(EventType)(0),                    // 0: protos.v1.EventType
	(*BookingEvent)(nil),              // 1: protos.v1.BookingEvent
	(*Booking)(nil),                   // 2: protos.v1.Booking
	(*Metadata)(nil),                  // 3: protos.v1.Metadata
	(*WriteBookingEventRequest)(nil),  // 4: protos.v1.WriteBookingEventRequest
	(*WriteBookingEventResponse)(nil), // 5: protos.v1.WriteBookingEventResponse
	(*ReadBookingEventsRequest)(nil),  // 6: protos.v1.ReadBookingEventsRequest
	(*ReadBookingEventsResponse)(nil), // 7: protos.v1.ReadBookingEventsResponse
	(*timestamppb.Timestamp)(nil),     // 8: google.protobuf.Timestamp
}
var file_protos_v1_booking_proto_depIdxs = []int32{
	3, // 0: protos.v1.BookingEvent.metadata:type_name -> protos.v1.Metadata
	0, // 1: protos.v1.BookingEvent.event_type:type_name -> protos.v1.EventType
	2, // 2: protos.v1.BookingEvent.booking:type_name -> protos.v1.Booking
	8, // 3: protos.v1.Metadata.timestamp:type_name -> google.protobuf.Timestamp
	1, // 4: protos.v1.WriteBookingEventRequest.booking_event:type_name -> protos.v1.BookingEvent
	1, // 5: protos.v1.ReadBookingEventsResponse.booking_event:type_name -> protos.v1.BookingEvent
	4, // 6: protos.v1.BookingEventService.WriteBookingEvent:input_type -> protos.v1.WriteBookingEventRequest
	6, // 7: protos.v1.BookingEventService.ReadBookingEvents:input_type -> protos.v1.ReadBookingEventsRequest
	5, // 8: protos.v1.BookingEventService.WriteBookingEvent:output_type -> protos.v1.WriteBookingEventResponse
	7, // 9: protos.v1.BookingEventService.ReadBookingEvents:output_type -> protos.v1.ReadBookingEventsResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_protos_v1_booking_proto_init() }
func file_protos_v1_booking_proto_init() {
	if File_protos_v1_booking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_v1_booking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingEvent); i {
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
		file_protos_v1_booking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Booking); i {
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
		file_protos_v1_booking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metadata); i {
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
		file_protos_v1_booking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteBookingEventRequest); i {
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
		file_protos_v1_booking_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteBookingEventResponse); i {
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
		file_protos_v1_booking_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadBookingEventsRequest); i {
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
		file_protos_v1_booking_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadBookingEventsResponse); i {
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
			RawDescriptor: file_protos_v1_booking_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_v1_booking_proto_goTypes,
		DependencyIndexes: file_protos_v1_booking_proto_depIdxs,
		EnumInfos:         file_protos_v1_booking_proto_enumTypes,
		MessageInfos:      file_protos_v1_booking_proto_msgTypes,
	}.Build()
	File_protos_v1_booking_proto = out.File
	file_protos_v1_booking_proto_rawDesc = nil
	file_protos_v1_booking_proto_goTypes = nil
	file_protos_v1_booking_proto_depIdxs = nil
}
