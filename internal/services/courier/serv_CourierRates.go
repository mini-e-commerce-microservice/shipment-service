package courier

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/SyaibanAhmadRamadhan/go-collection"
	"github.com/mini-e-commerce-microservice/shipment-service/generated/proto/sha_256_payload"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/repositories/biteship_api"
	"google.golang.org/protobuf/proto"
)

func (s *service) CourierRates(ctx context.Context, input CourierRatesInput) (output CourierRatesOutput, err error) {
	biteshipApiCourierRateItems := make([]biteship_api.CourierRateInputItem, 0, len(input.Items))
	payloadShaProductItem := make([]*sha_256_payload.CourierRateProductItem, 0, len(input.Items))
	for _, item := range input.Items {
		biteshipApiCourierRateItems = append(biteshipApiCourierRateItems, biteship_api.CourierRateInputItem{
			Name:     item.Name,
			Value:    item.Price,
			Length:   item.Length,
			Width:    item.Width,
			Weight:   item.Weight,
			Height:   item.Height,
			Quantity: item.Qty,
		})
		payloadShaProductItem = append(payloadShaProductItem, &sha_256_payload.CourierRateProductItem{
			Length:    int64(item.Length),
			Width:     int64(item.Width),
			Height:    int64(item.Height),
			Weight:    int64(item.Weight),
			Quantity:  int64(item.Qty),
			Price:     float32(item.Price),
			ProductId: item.ProductID,
			Name:      item.Name,
		})
	}

	outputCourierRate, err := s.biteshipApiRepository.CourierRate(ctx, biteship_api.CourierRateInput{
		OriginAreaID:      input.OriginAreaSourceID,
		DestinationAreaID: input.DestinationAreaID,
		CourierCode:       "paxel,jne,sicepat",
		Items:             biteshipApiCourierRateItems,
	})
	if err != nil {
		if errors.Is(err, biteship_api.ErrNoCourierAvailable) {
			err = collection.Err(ErrNoCourierAvailable)
		}
		if errors.Is(err, biteship_api.ErrInvalidPostalCode) {
			err = collection.Err(ErrInvalidAddress)
		}

		return output, collection.Err(err)
	}

	output = CourierRatesOutput{
		Items: make([]CourierRatesOutputItem, 0, len(outputCourierRate.Items)),
	}

	for _, item := range outputCourierRate.Items {
		payloadSha := &sha_256_payload.CourierRate{
			ProductItem:                  payloadShaProductItem,
			AvailableForCashOnDelivery:   item.AvailableForCOD,
			AvailableForProofOfDelivery:  item.AvailableForPOD,
			AvailableForInstantWaybillId: item.AvailableForInstantWaybillID,
			AvailableForInsurance:        item.AvailableForInsurance,
			Company:                      item.Company,
			CourierCode:                  item.CourierCode,
			CourierServiceCode:           item.CourierServiceCode,
			Duration:                     item.Duration,
			ShipmentDurationRange:        item.ShipmentDurationRange,
			ShipmentDurationUnit:         item.ShipmentDurationUnit,
			ServiceType:                  item.ServiceType,
			CourierPrice:                 float32(item.Price),
			Type:                         item.Type,
		}
		payloadShaMarshal, err := proto.Marshal(payloadSha)
		if err != nil {
			return output, collection.Err(err)
		}

		hash := sha256.New()
		combined := append([]byte(s.sha256Key.ShippmentServiceCourierRate+"|"), payloadShaMarshal...)
		hash.Write(combined)
		hashedData := hash.Sum(nil)

		output.Items = append(output.Items, CourierRatesOutputItem{
			ID:                           hex.EncodeToString(hashedData),
			AvailableForCOD:              item.AvailableForCOD,
			AvailableForPOD:              item.AvailableForPOD,
			AvailableForInstantWaybillID: item.AvailableForInstantWaybillID,
			AvailableForInsurance:        item.AvailableForInsurance,
			Company:                      item.Company,
			CourierName:                  item.CourierName,
			CourierCode:                  item.CourierCode,
			CourierServiceName:           item.CourierServiceName,
			CourierServiceCode:           item.CourierServiceCode,
			Description:                  item.Description,
			Duration:                     item.Duration,
			ShipmentDurationRange:        item.ShipmentDurationRange,
			ShipmentDurationUnit:         item.ShipmentDurationUnit,
			ServiceType:                  item.ServiceType,
			ShippingType:                 item.ShippingType,
			Price:                        item.Price,
			Type:                         item.Type,
		})
	}

	return
}

type CourierRatesInput struct {
	OriginAreaSourceID string
	DestinationAreaID  string
	Items              []CourierRatesInputItem
}

type CourierRatesInputItem struct {
	Name      string
	ProductID int64
	Price     float64
	Length    int
	Width     int
	Weight    int
	Height    int
	Qty       int
}

type CourierRatesOutput struct {
	Items []CourierRatesOutputItem
}

type CourierRatesOutputItem struct {
	ID                           string
	AvailableForCOD              bool
	AvailableForPOD              bool
	AvailableForInstantWaybillID bool
	AvailableForInsurance        bool
	Company                      string
	CourierName                  string
	CourierCode                  string
	CourierServiceName           string
	CourierServiceCode           string
	Description                  string
	Duration                     string
	ShipmentDurationRange        string
	ShipmentDurationUnit         string
	ServiceType                  string
	ShippingType                 string
	Price                        float64
	Type                         string
}
