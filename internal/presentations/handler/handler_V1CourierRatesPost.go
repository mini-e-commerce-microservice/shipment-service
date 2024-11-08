package handler

import (
	"errors"
	"github.com/mini-e-commerce-microservice/shipment-service/generated/api"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/services/courier"
	"net/http"
)

func (h *handler) V1CourierRatesPost(w http.ResponseWriter, r *http.Request) {
	req := api.V1CourierRatesPostRequestBody{}
	if !h.httpOtel.BindBodyRequest(w, r, &req) {
		return
	}

	_, ok := h.getUserFromBearerAuth(w, r, false)
	if !ok {
		return
	}

	courierRatesInputItems := make([]courier.CourierRatesInputItem, 0, len(req.ProductItems))
	for _, item := range req.ProductItems {
		courierRatesInputItems = append(courierRatesInputItems, courier.CourierRatesInputItem{
			Name:      item.Name,
			ProductID: item.Id,
			Price:     item.Price,
			Length:    item.Length,
			Width:     item.Width,
			Weight:    item.Weight,
			Height:    item.Height,
			Qty:       item.Qty,
		})
	}

	outputCourierRates, err := h.serv.courierService.CourierRates(r.Context(), courier.CourierRatesInput{
		OriginAreaSourceID: req.OriginAddressSourceId,
		DestinationAreaID:  req.DestinationAddressSourceId,
		Items:              courierRatesInputItems,
	})
	if err != nil {
		if errors.Is(err, courier.ErrNoCourierAvailable) {
			h.httpOtel.WriteJson(w, r, http.StatusOK, map[string]string{
				"message": "no courier available",
			})
		} else if errors.Is(err, courier.ErrInvalidAddress) {
			h.httpOtel.Err(w, r, http.StatusBadRequest, err, "invalid address")
		} else {
			h.httpOtel.Err(w, r, http.StatusInternalServerError, err)
		}
		return
	}

	resp := api.V1CourierRatesPostResponseBody{
		Destination: api.V1CourierRatesPostResponseBodyLocation{
			Address:    outputCourierRates.Destination.Address,
			Latitude:   outputCourierRates.Destination.Latitude,
			LocationId: outputCourierRates.Destination.LocationId,
			Longitude:  outputCourierRates.Destination.Longitude,
			PostalCode: outputCourierRates.Destination.PostalCode,
		},
		Origin: api.V1CourierRatesPostResponseBodyLocation{
			Address:    outputCourierRates.Origin.Address,
			Latitude:   outputCourierRates.Origin.Latitude,
			LocationId: outputCourierRates.Origin.LocationId,
			Longitude:  outputCourierRates.Origin.Longitude,
			PostalCode: outputCourierRates.Origin.PostalCode,
		},
		Items: make([]api.V1CourierRatesPostResponseBodyItem, 0, len(outputCourierRates.Items)),
	}

	for _, item := range outputCourierRates.Items {
		resp.Items = append(resp.Items, api.V1CourierRatesPostResponseBodyItem{
			AvailableForCashOnDelivery:   item.AvailableForCOD,
			AvailableForInstantWaybillId: item.AvailableForInstantWaybillID,
			AvailableForInsurance:        item.AvailableForInsurance,
			AvailableForProofOfDelivery:  item.AvailableForPOD,
			Company:                      item.Company,
			CourierCode:                  item.CourierCode,
			CourierName:                  item.CourierName,
			CourierServiceCode:           item.CourierServiceCode,
			CourierServiceName:           item.CourierServiceName,
			Description:                  item.Description,
			Duration:                     item.Duration,
			Id:                           item.ID,
			Price:                        item.Price,
			ServiceType:                  item.ServiceType,
			ShipmentDurationRange:        item.ShipmentDurationRange,
			ShipmentDurationUnit:         item.ShipmentDurationUnit,
			ShippingType:                 item.ShippingType,
			Type:                         item.Type,
		})
	}

	h.httpOtel.WriteJson(w, r, http.StatusOK, resp)
}
