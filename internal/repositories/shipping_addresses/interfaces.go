package shipping_addresses

import "context"

type Repository interface {
	Create(ctx context.Context, input CreateInput) (err error)
	GetAddress(ctx context.Context, input GetAddressInput) (output GetAddressOutput, err error)
}
