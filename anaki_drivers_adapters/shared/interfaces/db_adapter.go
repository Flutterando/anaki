package interfaces

import (
	"context"

	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/types"
)

type ExecuteResult struct {
	Rows         []map[string]interface{}
	RowsAffected int64
}

type DBAdapter interface {
	Connect(ctx context.Context, config types.Config) error
	Disconnect() error
	Execute(ctx context.Context, query string, args map[string]interface{}) (*ExecuteResult, error)
}
