package driver

import "github.com/flutterando/anaki/anaki_drivers_adapters/shared/interfaces"

var _ interfaces.DBAdapter = (*PostgresDriver)(nil)

func (p *PostgresDriver) Disconnect() error {
	if p.Conn == nil {
		return nil
	}

	p.Conn.Close()
	p.Conn = nil
	return nil
}
