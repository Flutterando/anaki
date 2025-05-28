package driver

import (
	"database/sql"

	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/contracts"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLDriver struct {
	Conn *sql.DB
}

func NewMySQLDriver() *MySQLDriver {
	return &MySQLDriver{}
}

var _ contracts.DriverAdapter = (*MySQLDriver)(nil)
