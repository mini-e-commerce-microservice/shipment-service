package address

import "context"

type Service interface {
	CreateAddress(ctx context.Context, input CreateAddressInput) (err error)
}
