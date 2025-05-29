package driver

import (
	"database/sql"

	"github.com/flutterando/anaki/anaki_drivers_adapters/pkg/driver"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLDriver struct {
	Conn *sql.DB
}

func NewMySQLDriver() *MySQLDriver {
	return &MySQLDriver{}
}

var _ driver.DriverAdapter = (*MySQLDriver)(nil)
