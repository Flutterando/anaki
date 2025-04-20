package driver_test

import (
	"context"
	"os"

	"testing"

	"github.com/flutterando/anaki/anaki_drivers_adapters/modules/postgres/driver"
	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/types"
)

func TestPostgresDriver_Disconnect_ShouldSucceed(t *testing.T) {
	connStr := os.Getenv("POSTGRES_TEST_DATABASE_URL")

	config := types.Config{
		URL: connStr,
	}

	ctx := context.Background()
	db := &driver.PostgresDriver{}

	if err := db.Connect(ctx, config); err != nil {
		t.Fatalf("expected no error when connecting, got: %v", err)
	}

	if db.Conn == nil {
		t.Fatal("expected connection pool to be initialized, got nil")
	}

	if err := db.Disconnect(); err != nil {
		t.Fatalf("expected no error when disconnecting, got: %v", err)
	}

	if db.Conn != nil {
		t.Fatal("expected connection pool to be nil after disconnect, but it was still initialized")
	}
}
