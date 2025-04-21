package driver

import (
	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/contracts"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDriver struct {
	Conn *pgxpool.Pool
}

func NewPostgresDriver() *PostgresDriver {
	return &PostgresDriver{}
}

var _ contracts.DriverAdapter = (*PostgresDriver)(nil)
