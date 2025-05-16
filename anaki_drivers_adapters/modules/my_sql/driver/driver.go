package driver

import (
	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/contracts"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MySQLDriver struct {
	Conn *pgxpool.Pool
}

func NewMySQLDriver() *MySQLDriver {
	return &MySQLDriver{}
}

var _ contracts.DriverAdapter = (*MySQLDriver)(nil)
