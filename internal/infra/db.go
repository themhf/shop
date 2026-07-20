package infra

import (
	"database/sql"
)

func ConnetDB(dsn string) (*sql.DB, error) {
	return sql.Open("postgres", dsn)
}
