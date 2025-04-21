package ffi_test

import (
	"os"
	"testing"

	"github.com/flutterando/anaki/anaki_drivers_adapters/modules/postgres/ffi"
	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/types"
)

func TestPostgresFFIDatabaseConnection(t *testing.T) {
	connStr := os.Getenv("POSTGRES_TEST_DATABASE_URL")

	result := ffi.SetupDatabaseConnection(types.Config{
		URL: connStr,
	})

	if result != "Connection successful" {
		t.Errorf("Database connection failed. Expected 'Connection successful', got: %s", result)
	}
}

func TestPostgresFFIDatabaseConnectionFailed(t *testing.T) {
	connStr := "postgres://invaliduser:invalidpass@localhost:5432/testdb"

	result := ffi.SetupDatabaseConnection(types.Config{
		URL: connStr,
	})

	if result == "Connection successful" {
		t.Errorf("Expected database connection to fail, but it succeeded. Result: %s", result)
	}
}
