package adapter

import (
	"anaki_drivers_adapters/types"
	"context"
)

type DBAdapter interface {
	Connect(ctx context.Context, config types.Config) error
	Disconnect() error
}
