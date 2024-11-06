package saga

import "context"

type Service interface {
	OrderCourierRate(ctx context.Context) (err error)
}
