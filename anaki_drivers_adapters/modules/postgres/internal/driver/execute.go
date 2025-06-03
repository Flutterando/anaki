package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/flutterando/anaki/anaki_drivers_adapters/modules/postgres/internal/sqlparams"
	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/pkg/driver"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (p *Driver) Execute(ctx context.Context, query string, args map[string]interface{}) (*driver.ExecuteResult, error) {

	if query == "" {
		return nil, errors.New("query is required")
	}

	var queryResult string
	var argsResult []interface{}
	var err error

	if len(args) > 0 {
		queryResult, argsResult, err = sqlparams.ConvertNamedParamsToDollar(query, args)
		if err != nil {
			return nil, err
		}
	} else {
		queryResult = query
		argsResult = nil
	}

	isSelect := strings.HasPrefix(strings.ToUpper(strings.TrimSpace(queryResult)), "SELECT")

	if isSelect {
		var rows pgx.Rows
		if argsResult == nil {
			rows, err = p.Conn.Query(ctx, queryResult)
		} else {
			rows, err = p.Conn.Query(ctx, queryResult, argsResult...)
		}
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		processedRows, err := p.processRow(rows)
		if err != nil {
			return nil, err
		}

		return &driver.ExecuteResult{
			Rows:         processedRows,
			RowsAffected: 0,
		}, nil
	}

	var result pgconn.CommandTag
	if argsResult == nil {
		result, err = p.Conn.Exec(ctx, queryResult)
	} else {
		result, err = p.Conn.Exec(ctx, queryResult, argsResult...)
	}

	if err != nil {
		return nil, err
	}

	return &driver.ExecuteResult{
		Rows:         []map[string]interface{}{{}},
		RowsAffected: result.RowsAffected(),
	}, nil
}

func (p *Driver) processRow(rows pgx.Rows) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	fieldDescriptions := rows.FieldDescriptions()

	for rows.Next() {
		values := make([]interface{}, len(fieldDescriptions))
		valuePtrs := make([]interface{}, len(fieldDescriptions))

		for i := range fieldDescriptions {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		rowData := make(map[string]interface{})
		for i, field := range fieldDescriptions {
			rowData[string(field.Name)] = values[i]
		}

		result = append(result, rowData)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return result, nil
}
