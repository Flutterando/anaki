module github.com/flutterando/anaki/anaki_drivers_adapters/modules/my_sql

go 1.24.2

replace github.com/flutterando/anaki/anaki_drivers_adapters/shared => ../../shared

require github.com/go-sql-driver/mysql v1.9.2

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/flutterando/anaki/anaki_drivers_adapters/modules/postgres v0.0.0-20250505051425-03237c5b77d9
	github.com/flutterando/anaki/anaki_drivers_adapters/shared v0.0.0-20250505051425-03237c5b77d9
)
