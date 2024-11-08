package biteship_api

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/SyaibanAhmadRamadhan/go-collection"
	"net/http"
)

func (r *repository) CourierRate(ctx context.Context, input CourierRateInput) (output CourierRateOutput, err error) {
	endpoint := "https://api.biteship.com/v1/rates/couriers"

	req := r.client.R().
		SetContext(ctx).
		SetHeader("Authorization", r.conf.Token).
		SetHeader("Content-Type", "application/json").
		SetBody(input)

	resp, err := req.Post(endpoint)
	if err != nil {
		return output, collection.Err(err)
	}

	if resp.StatusCode() >= http.StatusInternalServerError {
		return output, errors.Join(err, ErrFromBiteshipApi)
	}

	err = json.Unmarshal(resp.Body(), &output)
	if err != nil {
		return output, collection.Err(err)
	}

	if output.Code == 40001001 {
		return output, ErrInvalidPostalCode
	}

	if output.Code == 40001002 {
		return output, ErrMissingParameter
	}

	if output.Code == 40001010 {
		return output, ErrNoCourierAvailable
	}
	return
}

type CourierRateInput struct {
	OriginAreaID      string                 `json:"origin_area_id"`
	DestinationAreaID string                 `json:"destination_area_id"`
	CourierCode       string                 `json:"couriers"`
	Items             []CourierRateInputItem `json:"items"`
}

type CourierRateInputItem struct {
	Name     string  `json:"name"`
	Value    float64 `json:"value"`
	Length   int     `json:"length"`
	Width    int     `json:"width"`
	Weight   int     `json:"weight"`
	Height   int     `json:"height"`
	Quantity int     `json:"quantity"`
}

type CourierRateOutput struct {
	Success     bool                    `json:"success"`
	Object      string                  `json:"object"`
	Message     string                  `json:"message"`
	Code        int64                   `json:"code"`
	Origin      CourierRateLocation     `json:"origin"`
	Destination CourierRateLocation     `json:"destination"`
	Items       []CourierRateOutputItem `json:"pricing"`
}

type CourierRateLocation struct {
	LocationID string  `json:"location_id"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	PostalCode int32   `json:"postal_code"`
	Address    string  `json:"address"`
}

type CourierRateOutputItem struct {
	AvailableForCOD              bool    `json:"available_for_cash_on_delivery"`
	AvailableForPOD              bool    `json:"available_for_proof_of_delivery"`
	AvailableForInstantWaybillID bool    `json:"available_for_instant_waybill_id"`
	AvailableForInsurance        bool    `json:"available_for_insurance"`
	Company                      string  `json:"company"`
	CourierName                  string  `json:"courier_name"`
	CourierCode                  string  `json:"courier_code"`
	CourierServiceName           string  `json:"courier_service_name"`
	CourierServiceCode           string  `json:"courier_service_code"`
	Description                  string  `json:"description"`
	Duration                     string  `json:"duration"`
	ShipmentDurationRange        string  `json:"shipment_duration_range"`
	ShipmentDurationUnit         string  `json:"shipment_duration_unit"`
	ServiceType                  string  `json:"service_type"`
	ShippingType                 string  `json:"shipping_type"`
	Price                        float64 `json:"price"`
	Type                         string  `json:"type"`
}
