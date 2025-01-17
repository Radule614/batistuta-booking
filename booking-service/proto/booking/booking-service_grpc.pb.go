// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.0
// source: booking/booking-service.proto

package booking

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BookingService_GetAll_FullMethodName                     = "/BookingService/GetAll"
	BookingService_MakeBookingRequest_FullMethodName         = "/BookingService/MakeBookingRequest"
	BookingService_DeleteBookingRequest_FullMethodName       = "/BookingService/DeleteBookingRequest"
	BookingService_GetAllByUserId_FullMethodName             = "/BookingService/GetAllByUserId"
	BookingService_ConfirmReservationRequest_FullMethodName  = "/BookingService/ConfirmReservationRequest"
	BookingService_GetAllReservationsForGuest_FullMethodName = "/BookingService/GetAllReservationsForGuest"
	BookingService_DeleteReservation_FullMethodName          = "/BookingService/DeleteReservation"
)

// BookingServiceClient is the client API for BookingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookingServiceClient interface {
	GetAll(ctx context.Context, in *AM_GetAllBookingRequests_Request, opts ...grpc.CallOption) (*AM_GetAllBookingRequests_Response, error)
	MakeBookingRequest(ctx context.Context, in *AM_BookingRequest_Request, opts ...grpc.CallOption) (*AM_CreateBookingRequest_Response, error)
	DeleteBookingRequest(ctx context.Context, in *AM_DeleteBookingRequest_Request, opts ...grpc.CallOption) (*AM_DeleteBookingRequest_Response, error)
	GetAllByUserId(ctx context.Context, in *AM_GetAllBookingRequestsByUserId_Request, opts ...grpc.CallOption) (*AM_GetAllBookingRequests_Response, error)
	ConfirmReservationRequest(ctx context.Context, in *ReservationConfirm_Request, opts ...grpc.CallOption) (*EmptyMessage, error)
	GetAllReservationsForGuest(ctx context.Context, in *AllReservationsForGuest_Request, opts ...grpc.CallOption) (*AllReservationsForGuest_Response, error)
	DeleteReservation(ctx context.Context, in *DeleteReservation_Request, opts ...grpc.CallOption) (*EmptyMessage, error)
}

type bookingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBookingServiceClient(cc grpc.ClientConnInterface) BookingServiceClient {
	return &bookingServiceClient{cc}
}

func (c *bookingServiceClient) GetAll(ctx context.Context, in *AM_GetAllBookingRequests_Request, opts ...grpc.CallOption) (*AM_GetAllBookingRequests_Response, error) {
	out := new(AM_GetAllBookingRequests_Response)
	err := c.cc.Invoke(ctx, BookingService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) MakeBookingRequest(ctx context.Context, in *AM_BookingRequest_Request, opts ...grpc.CallOption) (*AM_CreateBookingRequest_Response, error) {
	out := new(AM_CreateBookingRequest_Response)
	err := c.cc.Invoke(ctx, BookingService_MakeBookingRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) DeleteBookingRequest(ctx context.Context, in *AM_DeleteBookingRequest_Request, opts ...grpc.CallOption) (*AM_DeleteBookingRequest_Response, error) {
	out := new(AM_DeleteBookingRequest_Response)
	err := c.cc.Invoke(ctx, BookingService_DeleteBookingRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetAllByUserId(ctx context.Context, in *AM_GetAllBookingRequestsByUserId_Request, opts ...grpc.CallOption) (*AM_GetAllBookingRequests_Response, error) {
	out := new(AM_GetAllBookingRequests_Response)
	err := c.cc.Invoke(ctx, BookingService_GetAllByUserId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) ConfirmReservationRequest(ctx context.Context, in *ReservationConfirm_Request, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, BookingService_ConfirmReservationRequest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) GetAllReservationsForGuest(ctx context.Context, in *AllReservationsForGuest_Request, opts ...grpc.CallOption) (*AllReservationsForGuest_Response, error) {
	out := new(AllReservationsForGuest_Response)
	err := c.cc.Invoke(ctx, BookingService_GetAllReservationsForGuest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookingServiceClient) DeleteReservation(ctx context.Context, in *DeleteReservation_Request, opts ...grpc.CallOption) (*EmptyMessage, error) {
	out := new(EmptyMessage)
	err := c.cc.Invoke(ctx, BookingService_DeleteReservation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookingServiceServer is the server API for BookingService service.
// All implementations must embed UnimplementedBookingServiceServer
// for forward compatibility
type BookingServiceServer interface {
	GetAll(context.Context, *AM_GetAllBookingRequests_Request) (*AM_GetAllBookingRequests_Response, error)
	MakeBookingRequest(context.Context, *AM_BookingRequest_Request) (*AM_CreateBookingRequest_Response, error)
	DeleteBookingRequest(context.Context, *AM_DeleteBookingRequest_Request) (*AM_DeleteBookingRequest_Response, error)
	GetAllByUserId(context.Context, *AM_GetAllBookingRequestsByUserId_Request) (*AM_GetAllBookingRequests_Response, error)
	ConfirmReservationRequest(context.Context, *ReservationConfirm_Request) (*EmptyMessage, error)
	GetAllReservationsForGuest(context.Context, *AllReservationsForGuest_Request) (*AllReservationsForGuest_Response, error)
	DeleteReservation(context.Context, *DeleteReservation_Request) (*EmptyMessage, error)
	mustEmbedUnimplementedBookingServiceServer()
}

// UnimplementedBookingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBookingServiceServer struct {
}

func (UnimplementedBookingServiceServer) GetAll(context.Context, *AM_GetAllBookingRequests_Request) (*AM_GetAllBookingRequests_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedBookingServiceServer) MakeBookingRequest(context.Context, *AM_BookingRequest_Request) (*AM_CreateBookingRequest_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeBookingRequest not implemented")
}
func (UnimplementedBookingServiceServer) DeleteBookingRequest(context.Context, *AM_DeleteBookingRequest_Request) (*AM_DeleteBookingRequest_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBookingRequest not implemented")
}
func (UnimplementedBookingServiceServer) GetAllByUserId(context.Context, *AM_GetAllBookingRequestsByUserId_Request) (*AM_GetAllBookingRequests_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllByUserId not implemented")
}
func (UnimplementedBookingServiceServer) ConfirmReservationRequest(context.Context, *ReservationConfirm_Request) (*EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmReservationRequest not implemented")
}
func (UnimplementedBookingServiceServer) GetAllReservationsForGuest(context.Context, *AllReservationsForGuest_Request) (*AllReservationsForGuest_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllReservationsForGuest not implemented")
}
func (UnimplementedBookingServiceServer) DeleteReservation(context.Context, *DeleteReservation_Request) (*EmptyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReservation not implemented")
}
func (UnimplementedBookingServiceServer) mustEmbedUnimplementedBookingServiceServer() {}

// UnsafeBookingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookingServiceServer will
// result in compilation errors.
type UnsafeBookingServiceServer interface {
	mustEmbedUnimplementedBookingServiceServer()
}

func RegisterBookingServiceServer(s grpc.ServiceRegistrar, srv BookingServiceServer) {
	s.RegisterService(&BookingService_ServiceDesc, srv)
}

func _BookingService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AM_GetAllBookingRequests_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetAll(ctx, req.(*AM_GetAllBookingRequests_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_MakeBookingRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AM_BookingRequest_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).MakeBookingRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_MakeBookingRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).MakeBookingRequest(ctx, req.(*AM_BookingRequest_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_DeleteBookingRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AM_DeleteBookingRequest_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).DeleteBookingRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_DeleteBookingRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).DeleteBookingRequest(ctx, req.(*AM_DeleteBookingRequest_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetAllByUserId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AM_GetAllBookingRequestsByUserId_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetAllByUserId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_GetAllByUserId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetAllByUserId(ctx, req.(*AM_GetAllBookingRequestsByUserId_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_ConfirmReservationRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationConfirm_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).ConfirmReservationRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_ConfirmReservationRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).ConfirmReservationRequest(ctx, req.(*ReservationConfirm_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_GetAllReservationsForGuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllReservationsForGuest_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).GetAllReservationsForGuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_GetAllReservationsForGuest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).GetAllReservationsForGuest(ctx, req.(*AllReservationsForGuest_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _BookingService_DeleteReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReservation_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookingServiceServer).DeleteReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BookingService_DeleteReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookingServiceServer).DeleteReservation(ctx, req.(*DeleteReservation_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// BookingService_ServiceDesc is the grpc.ServiceDesc for BookingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BookingService",
	HandlerType: (*BookingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAll",
			Handler:    _BookingService_GetAll_Handler,
		},
		{
			MethodName: "MakeBookingRequest",
			Handler:    _BookingService_MakeBookingRequest_Handler,
		},
		{
			MethodName: "DeleteBookingRequest",
			Handler:    _BookingService_DeleteBookingRequest_Handler,
		},
		{
			MethodName: "GetAllByUserId",
			Handler:    _BookingService_GetAllByUserId_Handler,
		},
		{
			MethodName: "ConfirmReservationRequest",
			Handler:    _BookingService_ConfirmReservationRequest_Handler,
		},
		{
			MethodName: "GetAllReservationsForGuest",
			Handler:    _BookingService_GetAllReservationsForGuest_Handler,
		},
		{
			MethodName: "DeleteReservation",
			Handler:    _BookingService_DeleteReservation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "booking/booking-service.proto",
}
