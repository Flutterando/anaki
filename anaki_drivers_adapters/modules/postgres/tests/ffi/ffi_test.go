package ffi_test

import (
	"os"
	"testing"

	"github.com/flutterando/anaki/anaki_drivers_adapters/modules/postgres/ffi"
	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/contracts"
	ffistatus "github.com/flutterando/anaki/anaki_drivers_adapters/shared/ffi"
)

func TestPostgresFFIDatabaseConnection(t *testing.T) {
	connStr := os.Getenv("POSTGRES_TEST_DATABASE_URL")

	result := ffi.SetupDatabaseConnection(contracts.Config{
		URL: connStr,
	})

	if result != ffistatus.SQL_SUCCESS {
		t.Errorf("Database connection failed. Expected status %d (SQL_SUCCESS), got: %d", ffistatus.SQL_SUCCESS, result)
	}
}

func TestPostgresFFIDatabaseConnectionFailed(t *testing.T) {
	connStr := "postgres://invaliduser:invalidpass@localhost:5432/testdb"

	result := ffi.SetupDatabaseConnection(contracts.Config{
		URL: connStr,
	})

	if result == ffistatus.SQL_SUCCESS {
		t.Errorf("Expected database connection to fail. Connection succeeded unexpectedly with status: %d", result)
	}
}
