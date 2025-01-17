package handlers

import (
	"context"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/accommodation_service/controller"
	"github.com/XWS-SmFoYcSNaQ/batistuta-booking/common/proto/accommodation"
)

type AccommodationHandler struct {
	accommodation.UnimplementedAccommodationServiceServer
	AccommodationController *controller.AccommodationController
	PeriodController        *controller.PeriodController
	DiscountController      *controller.DiscountController
}

func (h AccommodationHandler) GetAllAccommodations(ctx context.Context, request *accommodation.AM_GetAllAccommodations_Request) (*accommodation.AM_GetAllAccommodations_Response, error) {
	return h.AccommodationController.GetAll(ctx, request)
}

func (h AccommodationHandler) GetMyAccommodations(ctx context.Context, request *accommodation.AM_GetAllAccommodations_Request) (*accommodation.AM_GetAllAccommodations_Response, error) {
	return h.AccommodationController.GetAllByHost(ctx, request)
}

func (h AccommodationHandler) CreateAccommodation(ctx context.Context, request *accommodation.AM_CreateAccommodation_Request) (*accommodation.AM_CreateAccommodation_Response, error) {
	return h.AccommodationController.Create(ctx, request)
}

func (h AccommodationHandler) GetAccommodation(ctx context.Context, request *accommodation.AM_GetAccommodation_Request) (*accommodation.AM_GetAccommodation_Response, error) {
	return h.AccommodationController.GetById(ctx, request)
}

func (h AccommodationHandler) GetAllPeriodsByAccommodation(ctx context.Context, request *accommodation.AM_GetAllPeriodsByAccommodation_Request) (*accommodation.AM_GetAllPeriodsByAccommodation_Response, error) {
	return h.PeriodController.GetAllByAccommodation(ctx, request)
}

func (h AccommodationHandler) CreatePeriod(ctx context.Context, request *accommodation.AM_CreatePeriod_Request) (*accommodation.AM_CreatePeriod_Response, error) {
	return h.PeriodController.Create(ctx, request)
}

func (h AccommodationHandler) GetAllDiscountsByAccommodation(ctx context.Context, request *accommodation.AM_GetAllDiscountsByAccommodation_Request) (*accommodation.AM_GetAllDiscountsByAccommodation_Response, error) {
	return h.DiscountController.GetAllByAccommodation(ctx, request)
}

func (h AccommodationHandler) GetAllDiscountsByAccommodationAndInterval(ctx context.Context, request *accommodation.AM_GetAllDiscountsByAccommodationAndInterval_Request) (*accommodation.AM_GetAllDiscountsByAccommodationAndInterval_Response, error) {
	return h.DiscountController.GetAllByAccommodationAndInterval(ctx, request)
}

func (h AccommodationHandler) CreateDiscount(ctx context.Context, request *accommodation.AM_CreateDiscount_Request) (*accommodation.AM_CreateDiscount_Response, error) {
	return h.DiscountController.Create(ctx, request)
}
