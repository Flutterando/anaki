package driver_test

import (
	"anaki_postgres/driver"
	"context"

	"anaki_drivers_adapters/types"

	"testing"
)

func TestPostgresDriver_Connect_Success(t *testing.T) {

	ctx := context.Background()
	db := &driver.PostgresDriver{}

	config := types.Config{
		URL: "",
	}

	err := db.Connect(ctx, config)
	if err != nil {
		t.Fatalf("Connect() error = %v", err)
	}

	if db.Conn == nil {
		t.Fatal("Connect() did not initialize connection pool")
	}
}
