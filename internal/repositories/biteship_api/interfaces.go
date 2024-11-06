package biteship_api

import "context"

type Repository interface {
	GetAddresses(ctx context.Context, input GetAddressesInput) (output GetAddressesOutput, err error)
	GetAddress(ctx context.Context, input GetAddressInput) (output GetAddressOutput, err error)
	CourierRate(ctx context.Context, input CourierRateInput) (output CourierRateOutput, err error)
}
