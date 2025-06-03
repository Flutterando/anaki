package driver

import "context"

type DriverAdapter interface {
	Connect(ctx context.Context, config Config) error
	Disconnect() error
	Execute(ctx context.Context, query string, args map[string]interface{}) (*ExecuteResult, error)
}
