package driver

import (
	"context"
	"errors"

	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/interfaces"
	"github.com/jackc/pgx/v5/pgxpool"
)

var _ interfaces.DBAdapter = (*PostgresDriver)(nil)

func (p *PostgresDriver) Connect(ctx context.Context, config types.Config) error {

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
