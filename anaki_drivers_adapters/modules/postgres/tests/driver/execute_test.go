package driver_test

import (
	"context"
	"os"
	"testing"

	"github.com/flutterando/anaki/anaki_drivers_adapters/modules/postgres/driver"
	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/contracts"
)

func setupDriver(t *testing.T) *driver.PostgresDriver {
	connStr := os.Getenv("POSTGRES_TEST_DATABASE_URL")
	if connStr == "" {
		t.Fatal("POSTGRES_TEST_DATABASE_URL is not set")
	}

	config := contracts.Config{
		URL: connStr,
	}

	ctx := context.Background()
	driver := &driver.PostgresDriver{}

	err := driver.Connect(ctx, config)
	if err != nil {
		t.Fatalf("expected no error connecting to database, got: %v", err)
	}

	t.Cleanup(func() {
		driver.Disconnect()
	})

	return driver
}

func TestPostgresDriver_CreateTable(t *testing.T) {
	driver := setupDriver(t)

	_, err := driver.Execute(context.Background(), `
	CREATE TABLE IF NOT EXISTS users_test (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		age INT NOT NULL
	)
`, nil)
	if err != nil {
		t.Fatalf("failed to create table: %v", err)
	}
}

func TestPostgresDriver_Insert(t *testing.T) {
	driver := setupDriver(t)

	result, err := driver.Execute(context.Background(), `
	INSERT INTO users_test (name, age) VALUES
		('Marcos', 25),
		('Jacob', 30),
		('Max', 25)
`, nil)
	if err != nil {
		t.Fatalf("failed to insert users_test: %v", err)
	}

	if result.RowsAffected != 3 {
		t.Fatalf("expected 3 rows affected, got: %d", result.RowsAffected)
	}
}

func TestPostgresDriver_Select(t *testing.T) {
	driver := setupDriver(t)

	query := "SELECT id, name FROM users_test WHERE age = :age"
	params := map[string]interface{}{
		"age": 25,
	}

	result, err := driver.Execute(context.Background(), query, params)
	if err != nil {
		t.Fatalf("expected no error on execute, got: %v", err)
	}

	if len(result.Rows) != 2 {
		t.Errorf("expected 2 users with age 25, got: %d", len(result.Rows))
	}

	for _, row := range result.Rows {
		name, ok := row["name"].(string)
		if !ok {
			t.Errorf("expected name to be string, got: %v", row["name"])
		}
		if name != "Marcos" && name != "Max" {
			t.Errorf("unexpected user name: %s", name)
		}
	}
}

func TestPostgresDriver_DropTable(t *testing.T) {
	driver := setupDriver(t)

	_, err := driver.Execute(context.Background(), `DROP TABLE IF EXISTS users_test CASCADE`, nil)
	if err != nil {
		t.Fatalf("failed to drop table: %v", err)
	}
}
