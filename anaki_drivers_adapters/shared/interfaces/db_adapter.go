package interfaces

import (
	"context"

	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/types"
)

type DBAdapter interface {
	Connect(ctx context.Context, config types.Config) error
	Disconnect() error
}
