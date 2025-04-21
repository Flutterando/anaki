package driver_test

import (
	"context"
	"os"

	"testing"

	"github.com/flutterando/anaki/anaki_drivers_adapters/modules/postgres/driver"
	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/contracts"
)

func TestPostgresDriver_Connect_ShouldSucceed(t *testing.T) {
	connStr := os.Getenv("POSTGRES_TEST_DATABASE_URL")

	config := contracts.Config{
		URL: connStr,
	}

	ctx := context.Background()
	db := &driver.PostgresDriver{}

	err := db.Connect(ctx, config)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if db.Conn == nil {
		t.Fatal("expected connection pool to be initialized, got nil")
	}
}

func TestPostgresDriver_Connect_ShouldFailWithInvalidCredentials(t *testing.T) {
	connStr := "postgresql://wrong:wrong123@localhost:5432/wrongdb"

	config := contracts.Config{
		URL: connStr,
	}

	ctx := context.Background()
	db := &driver.PostgresDriver{}

	err := db.Connect(ctx, config)
	if err == nil {
		t.Fatal("expected an error when connecting with invalid credentials, got nil")
	}

	if db.Conn != nil {
		t.Fatal("expected connection pool to be nil after failed connection, but it was initialized")
	}
}
