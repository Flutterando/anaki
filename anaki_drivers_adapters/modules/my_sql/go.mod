module github.com/flutterando/anaki/anaki_drivers_adapters/modules/my_sql

go 1.24.2

replace github.com/flutterando/anaki/anaki_drivers_adapters/shared => ../../shared

require github.com/go-sql-driver/mysql v1.9.2

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	golang.org/x/crypto v0.35.0 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/text v0.22.0 // indirect
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/flutterando/anaki/anaki_drivers_adapters/modules/postgres v0.0.0-20250505051425-03237c5b77d9
	github.com/flutterando/anaki/anaki_drivers_adapters/shared v0.0.0-20250505051425-03237c5b77d9
	github.com/jackc/pgx/v5 v5.7.4
)
