package driver

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/flutterando/anaki/anaki_drivers_adapters/modules/postgres/utils"
	"github.com/flutterando/anaki/anaki_drivers_adapters/shared/contracts"
	_ "github.com/go-sql-driver/mysql"
)

func (p *MySQLDriver) Execute(ctx context.Context, query string, args map[string]interface{}) (*contracts.ExecuteResult, error) {
	var emptyExecuteResult = &contracts.ExecuteResult{
		Rows:         []map[string]interface{}{{}},
		RowsAffected: 0,
	}

	if query == "" {
		return emptyExecuteResult, errors.New("query is required")
	}

	var queryResult string
	var argsResult []interface{}
	var err error

	if len(args) > 0 {
		queryResult, argsResult, err = utils.ConvertNamedParams(query, args)
		if err != nil {
			return emptyExecuteResult, err
		}
	} else {
		queryResult = query
		argsResult = nil
	}

	isSelect := strings.HasPrefix(strings.ToUpper(strings.TrimSpace(queryResult)), "SELECT")

	if isSelect {
		var rows *sql.Rows
		if argsResult == nil {
			rows, err = p.Conn.QueryContext(ctx, queryResult)
		} else {
			rows, err = p.Conn.QueryContext(ctx, queryResult, argsResult...)
		}
		if err != nil {
			return emptyExecuteResult, err
		}
		defer rows.Close()

		processedRows, err := p.processRow(rows)
		if err != nil {
			return emptyExecuteResult, err
		}

		return &contracts.ExecuteResult{
			Rows:         processedRows,
			RowsAffected: 0,
		}, nil
	}

	var result sql.Result
	if argsResult == nil {
		result, err = p.Conn.ExecContext(ctx, queryResult)
	} else {
		result, err = p.Conn.ExecContext(ctx, queryResult, argsResult...)
	}

	if err != nil {
		return emptyExecuteResult, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return emptyExecuteResult, err
	}

	return &contracts.ExecuteResult{
		Rows:         []map[string]interface{}{{}},
		RowsAffected: rowsAffected,
	}, nil
}

func (p *MySQLDriver) processRow(rows *sql.Rows) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %v", err)
	}

	for rows.Next() {
		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))

		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		rowData := make(map[string]interface{})
		for i, col := range columns {
			rowData[col] = values[i]
		}

		result = append(result, rowData)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %v", err)
	}

	return result, nil
}
