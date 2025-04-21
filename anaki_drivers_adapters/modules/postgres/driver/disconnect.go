package driver

import "github.com/flutterando/anaki/anaki_drivers_adapters/shared/contracts"

var _ contracts.DriverAdapter = (*PostgresDriver)(nil)

func (p *PostgresDriver) Disconnect() error {
	if p.Conn == nil {
		return nil
	}

	p.Conn.Close()
	p.Conn = nil
	return nil
}
