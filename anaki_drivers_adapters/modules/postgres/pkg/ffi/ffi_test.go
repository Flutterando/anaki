package ffi_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/flutterando/anaki/anaki_drivers_adapters/modules/postgres/pkg/ffi"
	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/pkg/driver"
	ffistatus "github.com/flutterando/anaki/anaki_drivers_adapters/shared/pkg/ffi"
)

func TestPostgresFFI(t *testing.T) {
	t.Run("DatabaseConnection", func(t *testing.T) {
		connStr := os.Getenv("POSTGRES_TEST_DATABASE_URL")
		result := ffi.SetupDatabaseConnection(driver.Config{
			URL: connStr,
		})

		if result != ffistatus.SQL_SUCCESS {
			t.Errorf("Database connection failed. Expected status %d (SQL_SUCCESS), got: %d", ffistatus.SQL_SUCCESS, result)
		}
	})

	t.Run("CreateTable", func(t *testing.T) {
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
	})

	t.Run("Insert", func(t *testing.T) {
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
	})

	t.Run("Select", func(t *testing.T) {
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

		var result struct {
			Rows []map[string]interface{} `json:"Rows"`
		}

		err = json.Unmarshal([]byte(resultJson), &result)
		if err != nil {
			t.Fatalf("Failed to unmarshal result: %v", err)
		}

		if len(result.Rows) != 2 {
			t.Errorf("Expected 2 results, got: %d", len(result.Rows))
		}
	})

	t.Run("DropTable", func(t *testing.T) {
		query := `DROP TABLE IF EXISTS users_test CASCADE`
		resultJson, status := ffi.SetupDatabaseExecute(query, "")

		if status != ffistatus.SQL_SUCCESS {
			t.Fatalf("Query execution failed. Expected status %d (SQL_SUCCESS), got: %d. Error: %s", ffistatus.SQL_SUCCESS, status, resultJson)
		}

		ffi.SetupDatabaseClose()
	})
}
