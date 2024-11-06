package courier

import "context"

type Service interface {
	CourierRates(ctx context.Context, input CourierRatesInput) (output CourierRatesOutput, err error)
}
