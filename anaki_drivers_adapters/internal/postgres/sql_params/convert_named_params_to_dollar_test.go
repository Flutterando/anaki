package sqlparams_test

import (
	"reflect"
	"testing"

	sqlparams "github.com/flutterando/anaki/anaki_drivers_adapters/internal/postgres/sql_params"
)

func TestConvertNamedParams(t *testing.T) {
	query := "INSERT INTO users (name, age) VALUES (:name, :age)"
	params := map[string]interface{}{
		"name": "Marcos",
		"age":  25,
	}

	expectedQuery := "INSERT INTO users (name, age) VALUES ($1, $2)"
	expectedArgs := []interface{}{"Marcos", 25}

	finalQuery, args, err := sqlparams.ConvertNamedParamsToDollar(query, params)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if finalQuery != expectedQuery {
		t.Errorf("expected query %q, got %q", expectedQuery, finalQuery)
	}

	if !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("expected args %v, got %v", expectedArgs, args)
	}
}
