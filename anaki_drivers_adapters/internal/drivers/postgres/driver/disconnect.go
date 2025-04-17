package driver

import (
	"anaki_drivers_adapters/adapter"
)

var _ adapter.DBAdapter = (*PostgresDriver)(nil)

func (p *PostgresDriver) Disconnect() error {
	if p.Conn == nil {
		return nil
	}

	p.Conn.Close()
	p.Conn = nil
	return nil
}
