package driver

type ExecuteResult struct {
	Rows         []map[string]interface{}
	RowsAffected int64
}
