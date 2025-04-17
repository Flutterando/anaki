package driver

import (
	"anaki_drivers_adapters/shared/interfaces"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDriver struct {
	Conn *pgxpool.Pool
}

func NewPostgresDriver() *PostgresDriver {
	return &PostgresDriver{}
}

var _ interfaces.DBAdapter = (*PostgresDriver)(nil)
