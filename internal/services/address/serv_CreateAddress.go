package address

import (
	"context"
	"errors"
	"github.com/SyaibanAhmadRamadhan/go-collection"
	"github.com/guregu/null/v5"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/models"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/repositories/biteship_api"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/repositories/shipping_addresses"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/util/primitive"
)

func (s *service) CreateAddress(ctx context.Context, input CreateAddressInput) (err error) {
	respAddresses, err := s.biteshipApiRepository.GetAddress(ctx, biteship_api.GetAddressInput{
		ID: input.ID,
	})
	if err != nil {
		if errors.Is(err, biteship_api.ErrFromInterServerBiteshipApi) {
			err = errors.Join(err, ErrFromInterServerBiteshipApi)
		} else if errors.Is(err, biteship_api.ErrBadRequestBiteshipApi) {
			err = errors.Join(err, ErrInvalidAddress)
		}
		return collection.Err(err)
	}

	if respAddresses.Data == nil || len(respAddresses.Data) <= 0 {
		return collection.Err(ErrInvalidAddress)
	}

	data := respAddresses.Data[0]
	err = s.shippingAddressRepository.Create(ctx, shipping_addresses.CreateInput{
		Data: models.ShippingAddress{
			UserID:          input.UserID,
			AddressSourceID: &data.ID,
			AddressSource:   string(primitive.AddressSourceBiteship),
			Name:            data.Name,
			Country:         data.CountryName,
			CountryCode:     data.CountryCode,
			Province:        data.AdministrativeDivisionLevel1Name,
			City:            data.AdministrativeDivisionLevel2Name,
			District:        data.AdministrativeDivisionLevel3Name,
			AddressNote:     input.AddressNote.Ptr(),
		},
	})
	if err != nil {
		return collection.Err(err)
	}

	return
}

type CreateAddressInput struct {
	ID          string
	UserID      int64
	AddressNote null.String
}
