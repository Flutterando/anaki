package driver

import (
	"context"
	"database/sql"
	"errors"

	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/contracts"
	_ "github.com/go-sql-driver/mysql"
)

var _ contracts.DriverAdapter = (*MySQLDriver)(nil)

func (d *MySQLDriver) Connect(ctx context.Context, config contracts.Config) error {
	if config.URL == "" {
		return errors.New("database url is required")
	}

	db, err := sql.Open("mysql", config.URL)
	if err != nil {
		return err
	}

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return err
	}

	d.Conn = db
	return nil
}
