package courier

import (
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/mini-e-commerce-microservice/shipment-service/generated/proto/secret_proto"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/repositories/biteship_api"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/repositories/shipping_addresses"
)

type service struct {
	biteshipApiRepository     biteship_api.Repository
	shippingAddressRepository shipping_addresses.Repository
	dbTransaction             wsqlx.Tx
	sha256Key                 *secret_proto.Sha256Key
}

type Opt struct {
	BiteshipApiRepository     biteship_api.Repository
	ShippingAddressRepository shipping_addresses.Repository
	DBTransaction             wsqlx.Tx
	Sha256Key                 *secret_proto.Sha256Key
}

func New(opt Opt) *service {
	return &service{
		biteshipApiRepository:     opt.BiteshipApiRepository,
		shippingAddressRepository: opt.ShippingAddressRepository,
		dbTransaction:             opt.DBTransaction,
		sha256Key:                 opt.Sha256Key,
	}
}
