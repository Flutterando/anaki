package postgres

func (p *Driver) Disconnect() error {
	if p.Conn == nil {
		return nil
	}

	p.Conn.Close()
	p.Conn = nil
	return nil
}
