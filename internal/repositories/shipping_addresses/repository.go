package shipping_addresses

import (
	"github.com/Masterminds/squirrel"
	wsqlx "github.com/SyaibanAhmadRamadhan/sqlx-wrapper"
)

type repository struct {
	db wsqlx.Rdbms
	sq squirrel.StatementBuilderType
}

func New(db wsqlx.Rdbms) *repository {
	return &repository{
		db: db,
		sq: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
