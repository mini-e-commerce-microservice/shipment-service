package address

import (
	"github.com/mini-e-commerce-microservice/shipment-service/internal/repositories/biteship_api"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/repositories/shipping_addresses"
)

type service struct {
	biteshipApiRepository     biteship_api.Repository
	shippingAddressRepository shipping_addresses.Repository
}

type Opt struct {
	BiteshipApiRepository     biteship_api.Repository
	ShippingAddressRepository shipping_addresses.Repository
}

func New(opt Opt) *service {
	return &service{
		biteshipApiRepository:     opt.BiteshipApiRepository,
		shippingAddressRepository: opt.ShippingAddressRepository,
	}
}
