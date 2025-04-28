package ffi_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/flutterando/anaki/anaki_drivers_adapters/modules/postgres/ffi"
	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/contracts"
	ffistatus "github.com/flutterando/anaki/anaki_drivers_adapters/shared/ffi"
)

func setupFFIDatabase(t *testing.T) {
	connStr := os.Getenv("POSTGRES_TEST_DATABASE_URL")
	if connStr == "" {
		t.Fatal("POSTGRES_TEST_DATABASE_URL is not set")
	}

	config := contracts.Config{
		URL: connStr,
	}

	result := ffi.SetupDatabaseConnection(config)
	if result != ffistatus.SQL_SUCCESS {
		t.Fatalf("Failed to connect to database via FFI, got status: %d", result)
	}

	t.Cleanup(func() {
		ffi.SetupDatabaseClose()
	})
}

func TestFFIDatabase_CreateTable(t *testing.T) {
	setupFFIDatabase(t)

	query := `
	CREATE TABLE IF NOT EXISTS users_test (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		age INT NOT NULL
	)
	`
	resultJson, status := ffi.SetupDatabaseExecute(query, "")

	if status != ffistatus.SQL_SUCCESS {
		t.Fatalf("Query execution failed. Expected status %d (SQL_SUCCESS), got: %d. Error: %s", ffistatus.SQL_SUCCESS, status, resultJson)
	}
}

func TestFFIDatabase_Insert(t *testing.T) {
	setupFFIDatabase(t)

	query := `
	INSERT INTO users_test (name, age) VALUES
		('Marcos', 25),
		('Jacob', 30),
		('Max', 25)
	`
	resultJson, status := ffi.SetupDatabaseExecute(query, "")

	if status != ffistatus.SQL_SUCCESS {
		t.Fatalf("Query execution failed. Expected status %d (SQL_SUCCESS), got: %d. Error: %s", ffistatus.SQL_SUCCESS, status, resultJson)
	}
}

func TestFFIDatabase_Select(t *testing.T) {
	setupFFIDatabase(t)

	query := "SELECT id, name FROM users_test WHERE age = :age"
	params := map[string]interface{}{
		"age": 25,
	}

	paramsJson, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("Failed to marshal params: %v", err)
	}

	resultJson, status := ffi.SetupDatabaseExecute(query, string(paramsJson))

	if status != ffistatus.SQL_SUCCESS {
		t.Fatalf("Query execution failed. Expected status %d (SQL_SUCCESS), got: %d. Error: %s", ffistatus.SQL_SUCCESS, status, resultJson)
	}

	t.Logf("Query result: %s", resultJson)
}

func TestFFIDatabase_DropTable(t *testing.T) {
	setupFFIDatabase(t)

	query := `DROP TABLE IF EXISTS users_test CASCADE`
	resultJson, status := ffi.SetupDatabaseExecute(query, "")

	if status != ffistatus.SQL_SUCCESS {
		t.Fatalf("Query execution failed. Expected status %d (SQL_SUCCESS), got: %d. Error: %s", ffistatus.SQL_SUCCESS, status, resultJson)
	}
}
