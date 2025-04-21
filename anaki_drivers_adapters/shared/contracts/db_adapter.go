package contracts

import (
	"context"
)

type Config struct {
	URL string
}

type ExecuteResult struct {
	Rows         []map[string]interface{}
	RowsAffected int64
}

type DriverAdapter interface {
	Connect(ctx context.Context, config Config) error
	Disconnect() error
	Execute(ctx context.Context, query string, args map[string]interface{}) (*ExecuteResult, error)
}
