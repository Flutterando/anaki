package driver_test

import (
	"context"
	"os"

	"testing"

	"github.com/flutterando/anaki/anaki_drivers_adapters/modules/postgres/driver"
	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/types"
)

func TestPostgresDriver_Connect_Success(t *testing.T) {
	connStr := os.Getenv("POSTGRES_TEST_DATABASE_URL")

	config := types.Config{
		URL: connStr,
	}

	ctx := context.Background()
	db := &driver.PostgresDriver{}

	err := db.Connect(ctx, config)
	if err != nil {
		t.Fatalf("Connect() error = %v", err)
	}

	if db.Conn == nil {
		t.Fatal("Connect() did not initialize connection pool")
	}
}
