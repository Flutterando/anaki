package postgres

import (
	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/pkg/driver"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Driver struct {
	Conn *pgxpool.Pool
}

func NewPostgresDriver() *Driver {
	return &Driver{}
}

var _ driver.DriverAdapter = (*Driver)(nil)
