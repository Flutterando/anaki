package driver

import (
	"context"
	"errors"

	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/contracts"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ contracts.DriverAdapter = (*MySQLDriver)(nil)

func (d *MySQLDriver) Connect(ctx context.Context, config contracts.Config) error {
	if config.URL == "" {
		return errors.New("database url is required")
	}

	conn, err := pgxpool.New(ctx, config.URL)
	if err != nil {
		return err
	}

	if err := conn.Ping(context.Background()); err != nil {
		conn.Close()
		return err
	}

	d.Conn = conn
	return nil
}
