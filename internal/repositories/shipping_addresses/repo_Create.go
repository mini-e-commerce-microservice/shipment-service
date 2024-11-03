package shipping_addresses

import (
	"context"
	"github.com/SyaibanAhmadRamadhan/go-collection"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/models"
)

func (r *repository) Create(ctx context.Context, input CreateInput) (err error) {
	columns, values := collection.GetTagsWithValues(input.Data, "db", "id")
	query := r.sq.Insert("shipping_addresses").Columns(columns...).Values(values...)

	rdbms := input.Tx
	if input.Tx == nil {
		rdbms = r.db
	}

	_, err = rdbms.ExecSq(ctx, query)
	if err != nil {
		return collection.Err(err)
	}

	return
}

type CreateInput struct {
	Tx   wsqlx.WriterCommand
	Data models.ShippingAddress
}
