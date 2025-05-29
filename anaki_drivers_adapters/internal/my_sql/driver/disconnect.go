package driver

import "github.com/flutterando/anaki/anaki_drivers_adapters/pkg/driver"

var _ driver.DriverAdapter = (*MySQLDriver)(nil)

func (p *MySQLDriver) Disconnect() error {
	if p.Conn == nil {
		return nil
	}

	p.Conn.Close()
	p.Conn = nil
	return nil
}
