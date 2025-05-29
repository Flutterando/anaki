package driver

import (
	"context"
	"database/sql"
	"errors"

	"github.com/flutterando/anaki/anaki_drivers_adapters/pkg/driver"
	_ "github.com/go-sql-driver/mysql"
)

var _ driver.DriverAdapter = (*MySQLDriver)(nil)

func (d *MySQLDriver) Connect(ctx context.Context, config driver.Config) error {
	if config.URL == "" {
		return errors.New("database url is required")
	}

	conn, err := sql.Open("mysql", config.URL)
	if err != nil {
		return err
	}

	if err := conn.PingContext(ctx); err != nil {
		conn.Close()
		return err
	}

	d.Conn = conn
	return nil
}
