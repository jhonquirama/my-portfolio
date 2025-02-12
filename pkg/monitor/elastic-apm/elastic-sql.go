package apm

import (
	"database/sql"

	apm "go.elastic.co/apm/module/apmsql/v2"
	_ "go.elastic.co/apm/module/apmsql/v2/pq" // nolint: revive
)

func Open(driverName, dataSourceName string) (*sql.DB, error) {
	return apm.Open(driverName, dataSourceName)
}
