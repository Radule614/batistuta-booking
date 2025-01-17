// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.0
// source: booking/booking-service.proto

package booking

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

type AM_GetAllBookingRequests_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AM_GetAllBookingRequests_Request) Reset() {
	*x = AM_GetAllBookingRequests_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AM_GetAllBookingRequests_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AM_GetAllBookingRequests_Request) ProtoMessage() {}

func (x *AM_GetAllBookingRequests_Request) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AM_GetAllBookingRequests_Request.ProtoReflect.Descriptor instead.
func (*AM_GetAllBookingRequests_Request) Descriptor() ([]byte, []int) {
	return file_booking_booking_service_proto_rawDescGZIP(), []int{0}
}

type AM_GetAllBookingRequestsByUserId_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*BookingRequestsDTO `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *AM_GetAllBookingRequestsByUserId_Response) Reset() {
	*x = AM_GetAllBookingRequestsByUserId_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AM_GetAllBookingRequestsByUserId_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AM_GetAllBookingRequestsByUserId_Response) ProtoMessage() {}

func (x *AM_GetAllBookingRequestsByUserId_Response) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AM_GetAllBookingRequestsByUserId_Response.ProtoReflect.Descriptor instead.
func (*AM_GetAllBookingRequestsByUserId_Response) Descriptor() ([]byte, []int) {
	return file_booking_booking_service_proto_rawDescGZIP(), []int{1}
}

func (x *AM_GetAllBookingRequestsByUserId_Response) GetData() []*BookingRequestsDTO {
	if x != nil {
		return x.Data
	}
	return nil
}

type AM_GetAllBookingRequests_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*BookingRequestsDTO `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *AM_GetAllBookingRequests_Response) Reset() {
	*x = AM_GetAllBookingRequests_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AM_GetAllBookingRequests_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AM_GetAllBookingRequests_Response) ProtoMessage() {}

func (x *AM_GetAllBookingRequests_Response) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AM_GetAllBookingRequests_Response.ProtoReflect.Descriptor instead.
func (*AM_GetAllBookingRequests_Response) Descriptor() ([]byte, []int) {
	return file_booking_booking_service_proto_rawDescGZIP(), []int{2}
}

func (x *AM_GetAllBookingRequests_Response) GetData() []*BookingRequestsDTO {
	if x != nil {
		return x.Data
	}
	return nil
}

type AM_GetAllBookingRequestsByUserId_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AM_GetAllBookingRequestsByUserId_Request) Reset() {
	*x = AM_GetAllBookingRequestsByUserId_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AM_GetAllBookingRequestsByUserId_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AM_GetAllBookingRequestsByUserId_Request) ProtoMessage() {}

func (x *AM_GetAllBookingRequestsByUserId_Request) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AM_GetAllBookingRequestsByUserId_Request.ProtoReflect.Descriptor instead.
func (*AM_GetAllBookingRequestsByUserId_Request) Descriptor() ([]byte, []int) {
	return file_booking_booking_service_proto_rawDescGZIP(), []int{3}
}

func (x *AM_GetAllBookingRequestsByUserId_Request) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AM_BookingRequest_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccommodationId string `protobuf:"bytes,1,opt,name=accommodationId,proto3" json:"accommodationId,omitempty"`
	StartDate       string `protobuf:"bytes,2,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate         string `protobuf:"bytes,3,opt,name=endDate,proto3" json:"endDate,omitempty"`
	UserId          string `protobuf:"bytes,4,opt,name=userId,proto3" json:"userId,omitempty"`
	NumberOfGuests  int32  `protobuf:"varint,5,opt,name=numberOfGuests,proto3" json:"numberOfGuests,omitempty"`
}

func (x *AM_BookingRequest_Request) Reset() {
	*x = AM_BookingRequest_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AM_BookingRequest_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AM_BookingRequest_Request) ProtoMessage() {}

func (x *AM_BookingRequest_Request) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AM_BookingRequest_Request.ProtoReflect.Descriptor instead.
func (*AM_BookingRequest_Request) Descriptor() ([]byte, []int) {
	return file_booking_booking_service_proto_rawDescGZIP(), []int{4}
}

func (x *AM_BookingRequest_Request) GetAccommodationId() string {
	if x != nil {
		return x.AccommodationId
	}
	return ""
}

func (x *AM_BookingRequest_Request) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *AM_BookingRequest_Request) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *AM_BookingRequest_Request) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AM_BookingRequest_Request) GetNumberOfGuests() int32 {
	if x != nil {
		return x.NumberOfGuests
	}
	return 0
}

type AM_CreateBookingRequest_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AM_CreateBookingRequest_Response) Reset() {
	*x = AM_CreateBookingRequest_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AM_CreateBookingRequest_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AM_CreateBookingRequest_Response) ProtoMessage() {}

func (x *AM_CreateBookingRequest_Response) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AM_CreateBookingRequest_Response.ProtoReflect.Descriptor instead.
func (*AM_CreateBookingRequest_Response) Descriptor() ([]byte, []int) {
	return file_booking_booking_service_proto_rawDescGZIP(), []int{5}
}

func (x *AM_CreateBookingRequest_Response) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type BookingRequestsDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AccommodationId string `protobuf:"bytes,2,opt,name=accommodationId,proto3" json:"accommodationId,omitempty"`
	StartDate       string `protobuf:"bytes,3,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate         string `protobuf:"bytes,4,opt,name=endDate,proto3" json:"endDate,omitempty"`
	UserId          string `protobuf:"bytes,5,opt,name=userId,proto3" json:"userId,omitempty"`
	NumberOfGuests  int32  `protobuf:"varint,6,opt,name=numberOfGuests,proto3" json:"numberOfGuests,omitempty"`
}

func (x *BookingRequestsDTO) Reset() {
	*x = BookingRequestsDTO{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingRequestsDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingRequestsDTO) ProtoMessage() {}

func (x *BookingRequestsDTO) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingRequestsDTO.ProtoReflect.Descriptor instead.
func (*BookingRequestsDTO) Descriptor() ([]byte, []int) {
	return file_booking_booking_service_proto_rawDescGZIP(), []int{6}
}

func (x *BookingRequestsDTO) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *BookingRequestsDTO) GetAccommodationId() string {
	if x != nil {
		return x.AccommodationId
	}
	return ""
}

func (x *BookingRequestsDTO) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *BookingRequestsDTO) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

func (x *BookingRequestsDTO) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BookingRequestsDTO) GetNumberOfGuests() int32 {
	if x != nil {
		return x.NumberOfGuests
	}
	return 0
}

type AM_DeleteBookingRequest_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AM_DeleteBookingRequest_Request) Reset() {
	*x = AM_DeleteBookingRequest_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AM_DeleteBookingRequest_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AM_DeleteBookingRequest_Request) ProtoMessage() {}

func (x *AM_DeleteBookingRequest_Request) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AM_DeleteBookingRequest_Request.ProtoReflect.Descriptor instead.
func (*AM_DeleteBookingRequest_Request) Descriptor() ([]byte, []int) {
	return file_booking_booking_service_proto_rawDescGZIP(), []int{7}
}

func (x *AM_DeleteBookingRequest_Request) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AM_DeleteBookingRequest_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AM_DeleteBookingRequest_Response) Reset() {
	*x = AM_DeleteBookingRequest_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AM_DeleteBookingRequest_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AM_DeleteBookingRequest_Response) ProtoMessage() {}

func (x *AM_DeleteBookingRequest_Response) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AM_DeleteBookingRequest_Response.ProtoReflect.Descriptor instead.
func (*AM_DeleteBookingRequest_Response) Descriptor() ([]byte, []int) {
	return file_booking_booking_service_proto_rawDescGZIP(), []int{8}
}

type ReservationConfirm_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ReservationConfirm_Request) Reset() {
	*x = ReservationConfirm_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReservationConfirm_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReservationConfirm_Request) ProtoMessage() {}

func (x *ReservationConfirm_Request) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReservationConfirm_Request.ProtoReflect.Descriptor instead.
func (*ReservationConfirm_Request) Descriptor() ([]byte, []int) {
	return file_booking_booking_service_proto_rawDescGZIP(), []int{9}
}

func (x *ReservationConfirm_Request) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type EmptyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyMessage) Reset() {
	*x = EmptyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyMessage) ProtoMessage() {}

func (x *EmptyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyMessage.ProtoReflect.Descriptor instead.
func (*EmptyMessage) Descriptor() ([]byte, []int) {
	return file_booking_booking_service_proto_rawDescGZIP(), []int{10}
}

type AllReservationsForGuest_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AllReservationsForGuest_Request) Reset() {
	*x = AllReservationsForGuest_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllReservationsForGuest_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllReservationsForGuest_Request) ProtoMessage() {}

func (x *AllReservationsForGuest_Request) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllReservationsForGuest_Request.ProtoReflect.Descriptor instead.
func (*AllReservationsForGuest_Request) Descriptor() ([]byte, []int) {
	return file_booking_booking_service_proto_rawDescGZIP(), []int{11}
}

func (x *AllReservationsForGuest_Request) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AllReservationsForGuest_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*BookingRequestsDTO `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *AllReservationsForGuest_Response) Reset() {
	*x = AllReservationsForGuest_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllReservationsForGuest_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllReservationsForGuest_Response) ProtoMessage() {}

func (x *AllReservationsForGuest_Response) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_service_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllReservationsForGuest_Response.ProtoReflect.Descriptor instead.
func (*AllReservationsForGuest_Response) Descriptor() ([]byte, []int) {
	return file_booking_booking_service_proto_rawDescGZIP(), []int{12}
}

func (x *AllReservationsForGuest_Response) GetData() []*BookingRequestsDTO {
	if x != nil {
		return x.Data
	}
	return nil
}

type DeleteReservation_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteReservation_Request) Reset() {
	*x = DeleteReservation_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_booking_booking_service_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReservation_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReservation_Request) ProtoMessage() {}

func (x *DeleteReservation_Request) ProtoReflect() protoreflect.Message {
	mi := &file_booking_booking_service_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReservation_Request.ProtoReflect.Descriptor instead.
func (*DeleteReservation_Request) Descriptor() ([]byte, []int) {
	return file_booking_booking_service_proto_rawDescGZIP(), []int{13}
}

func (x *DeleteReservation_Request) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_booking_booking_service_proto protoreflect.FileDescriptor

var file_booking_booking_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x22, 0x0a, 0x20, 0x41, 0x4d, 0x5f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x54, 0x0a, 0x29, 0x41, 0x4d, 0x5f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x27, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x44, 0x54, 0x4f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x4c, 0x0a, 0x21, 0x41, 0x4d, 0x5f,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27,
	0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x44, 0x54,
	0x4f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x3a, 0x0a, 0x28, 0x41, 0x4d, 0x5f, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x5f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0xbd, 0x01, 0x0a, 0x19, 0x41, 0x4d, 0x5f, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x28, 0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x47, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x22, 0x32, 0x0a, 0x20, 0x41, 0x4d, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xc6, 0x01, 0x0a, 0x12, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x44, 0x54, 0x4f, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x28,
	0x0a, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x4f, 0x66, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x4f, 0x66, 0x47, 0x75, 0x65, 0x73, 0x74, 0x73,
	0x22, 0x31, 0x0a, 0x1f, 0x41, 0x4d, 0x5f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x22, 0x0a, 0x20, 0x41, 0x4d, 0x5f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2c, 0x0a, 0x1a, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x5f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x31, 0x0a, 0x1f, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x6f, 0x72, 0x47, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4b, 0x0a, 0x20, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x6f, 0x72, 0x47, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x44, 0x54, 0x4f, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2b, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x32, 0xee, 0x04, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12,
	0x21, 0x2e, 0x41, 0x4d, 0x5f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x41, 0x4d, 0x5f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x12, 0x4d, 0x61, 0x6b, 0x65,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x2e, 0x41, 0x4d, 0x5f, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x41, 0x4d, 0x5f,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x5d, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x41, 0x4d, 0x5f, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x41, 0x4d, 0x5f, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x61,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x29, 0x2e, 0x41, 0x4d, 0x5f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x41, 0x4d,
	0x5f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x49, 0x0a, 0x19, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x72, 0x6d, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x1a,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x46, 0x6f, 0x72, 0x47, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x6f, 0x72, 0x47,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x41,
	0x6c, 0x6c, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x46, 0x6f,
	0x72, 0x47, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x40, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_booking_booking_service_proto_rawDescOnce sync.Once
	file_booking_booking_service_proto_rawDescData = file_booking_booking_service_proto_rawDesc
)

func file_booking_booking_service_proto_rawDescGZIP() []byte {
	file_booking_booking_service_proto_rawDescOnce.Do(func() {
		file_booking_booking_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_booking_booking_service_proto_rawDescData)
	})
	return file_booking_booking_service_proto_rawDescData
}

var file_booking_booking_service_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_booking_booking_service_proto_goTypes = []interface{}{
	(*AM_GetAllBookingRequests_Request)(nil),          // 0: AM_GetAllBookingRequests_Request
	(*AM_GetAllBookingRequestsByUserId_Response)(nil), // 1: AM_GetAllBookingRequestsByUserId_Response
	(*AM_GetAllBookingRequests_Response)(nil),         // 2: AM_GetAllBookingRequests_Response
	(*AM_GetAllBookingRequestsByUserId_Request)(nil),  // 3: AM_GetAllBookingRequestsByUserId_Request
	(*AM_BookingRequest_Request)(nil),                 // 4: AM_BookingRequest_Request
	(*AM_CreateBookingRequest_Response)(nil),          // 5: AM_CreateBookingRequest_Response
	(*BookingRequestsDTO)(nil),                        // 6: BookingRequestsDTO
	(*AM_DeleteBookingRequest_Request)(nil),           // 7: AM_DeleteBookingRequest_Request
	(*AM_DeleteBookingRequest_Response)(nil),          // 8: AM_DeleteBookingRequest_Response
	(*ReservationConfirm_Request)(nil),                // 9: ReservationConfirm_Request
	(*EmptyMessage)(nil),                              // 10: EmptyMessage
	(*AllReservationsForGuest_Request)(nil),           // 11: AllReservationsForGuest_Request
	(*AllReservationsForGuest_Response)(nil),          // 12: AllReservationsForGuest_Response
	(*DeleteReservation_Request)(nil),                 // 13: DeleteReservation_Request
}
var file_booking_booking_service_proto_depIdxs = []int32{
	6,  // 0: AM_GetAllBookingRequestsByUserId_Response.data:type_name -> BookingRequestsDTO
	6,  // 1: AM_GetAllBookingRequests_Response.data:type_name -> BookingRequestsDTO
	6,  // 2: AllReservationsForGuest_Response.data:type_name -> BookingRequestsDTO
	0,  // 3: BookingService.GetAll:input_type -> AM_GetAllBookingRequests_Request
	4,  // 4: BookingService.MakeBookingRequest:input_type -> AM_BookingRequest_Request
	7,  // 5: BookingService.DeleteBookingRequest:input_type -> AM_DeleteBookingRequest_Request
	3,  // 6: BookingService.GetAllByUserId:input_type -> AM_GetAllBookingRequestsByUserId_Request
	9,  // 7: BookingService.ConfirmReservationRequest:input_type -> ReservationConfirm_Request
	11, // 8: BookingService.GetAllReservationsForGuest:input_type -> AllReservationsForGuest_Request
	13, // 9: BookingService.DeleteReservation:input_type -> DeleteReservation_Request
	2,  // 10: BookingService.GetAll:output_type -> AM_GetAllBookingRequests_Response
	5,  // 11: BookingService.MakeBookingRequest:output_type -> AM_CreateBookingRequest_Response
	8,  // 12: BookingService.DeleteBookingRequest:output_type -> AM_DeleteBookingRequest_Response
	2,  // 13: BookingService.GetAllByUserId:output_type -> AM_GetAllBookingRequests_Response
	10, // 14: BookingService.ConfirmReservationRequest:output_type -> EmptyMessage
	12, // 15: BookingService.GetAllReservationsForGuest:output_type -> AllReservationsForGuest_Response
	10, // 16: BookingService.DeleteReservation:output_type -> EmptyMessage
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_booking_booking_service_proto_init() }
func file_booking_booking_service_proto_init() {
	if File_booking_booking_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_booking_booking_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AM_GetAllBookingRequests_Request); i {
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
		file_booking_booking_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AM_GetAllBookingRequestsByUserId_Response); i {
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
		file_booking_booking_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AM_GetAllBookingRequests_Response); i {
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
		file_booking_booking_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AM_GetAllBookingRequestsByUserId_Request); i {
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
		file_booking_booking_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AM_BookingRequest_Request); i {
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
		file_booking_booking_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AM_CreateBookingRequest_Response); i {
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
		file_booking_booking_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingRequestsDTO); i {
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
		file_booking_booking_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AM_DeleteBookingRequest_Request); i {
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
		file_booking_booking_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AM_DeleteBookingRequest_Response); i {
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
		file_booking_booking_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReservationConfirm_Request); i {
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
		file_booking_booking_service_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyMessage); i {
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
		file_booking_booking_service_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllReservationsForGuest_Request); i {
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
		file_booking_booking_service_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllReservationsForGuest_Response); i {
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
		file_booking_booking_service_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReservation_Request); i {
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
			RawDescriptor: file_booking_booking_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_booking_booking_service_proto_goTypes,
		DependencyIndexes: file_booking_booking_service_proto_depIdxs,
		MessageInfos:      file_booking_booking_service_proto_msgTypes,
	}.Build()
	File_booking_booking_service_proto = out.File
	file_booking_booking_service_proto_rawDesc = nil
	file_booking_booking_service_proto_goTypes = nil
	file_booking_booking_service_proto_depIdxs = nil
}
