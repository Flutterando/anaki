package driver_test

import (
	"anaki_postgres/driver"
	"context"
	"os"

	"anaki_drivers_adapters/shared/types"

	"testing"
)

func TestPostgresDriver_Disconnect_Success(t *testing.T) {
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

	err = db.Disconnect()

	if err != nil {
		t.Fatalf("Disconnect() error = %v", err)
	}

	if db.Conn != nil {
		t.Fatal("Disconnect() did not properly close connection")
	}
}
