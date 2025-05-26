package postgres

import (
	"context"
	"errors"

	"github.com/flutterando/anaki/anaki_drivers_adapters/pkg/driver"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (p *Driver) Connect(ctx context.Context, config driver.Config) error {

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

	p.Conn = conn
	return nil
}
