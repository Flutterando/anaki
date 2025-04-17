package interfaces

import (
	"anaki_drivers_adapters/shared/types"
	"context"
)

type DBAdapter interface {
	Connect(ctx context.Context, config types.Config) error
	Disconnect() error
}
