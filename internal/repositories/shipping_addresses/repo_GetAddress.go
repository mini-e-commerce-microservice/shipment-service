package shipping_addresses

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/SyaibanAhmadRamadhan/go-collection"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/models"
	"github.com/mini-e-commerce-microservice/shipment-service/internal/repositories"
)

func (r *repository) GetAddress(ctx context.Context, input GetAddressInput) (output GetAddressOutput, err error) {
	query := r.sq.Select("*").From("shipping_addresses").Where(squirrel.Eq{"id": input.ID})

	err = r.db.QueryRowSq(ctx, query, wsqlx.QueryRowScanTypeStruct, &output.Data)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return output, collection.Err(repositories.ErrNoRecordRow)
		}
		return output, collection.Err(err)
	}
	return
}

type GetAddressInput struct {
	ID int64
}

type GetAddressOutput struct {
	Data models.ShippingAddress
}
