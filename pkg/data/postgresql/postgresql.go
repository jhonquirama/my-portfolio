package postgresql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/lib/pq"

	apm "github.com/jhonquirama/my-portfolio/pkg/monitor/elastic-apm"
)

//go:generate mockery --name DB
//go:generate mockery --name SqlRows
type (
	ExecStmt func(*sql.Tx) error

	db struct {
		*sql.DB
	}

	DB interface {
		QueryContext(ctx context.Context, query string, args ...interface{}) (SQLRows, error)
		ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
		DoTransactionWithContext(ctx context.Context, fnStmt ExecStmt) error
		PingContext(context.Context) error
		Close(rows SQLRows)
		Array(array interface{}) AnyArray
	}

	SQLRows interface {
		Next() bool
		Err() error
		Scan(dest ...interface{}) error
		Close() error
	}

	AnyArray interface {
		driver.Valuer
		sql.Scanner
	}

	Config interface {
		User() string
		Password() string
		Port() int64
		Host() string
		DbName() string
		SslMode() string
		SetMaxOpenConns() int
		SetMaxIdleConns() int
		SetConnMaxLifetime() int
	}
)

func NewPostgresqlDB(config Config) (DB, error) {
	var (
		dataSourceName string
		postgresql     db
		err            error
	)

	dataSourceName = fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=%s`,
		config.Host(),
		config.Port(),
		config.User(),
		config.Password(),
		config.DbName(),
		config.SslMode(),
	)

	if postgresql.DB, err = apm.Open("postgres", dataSourceName); err != nil {
		return nil, err
	}

	postgresql.DB.SetMaxOpenConns(config.SetMaxOpenConns())
	postgresql.DB.SetMaxIdleConns(config.SetMaxIdleConns())
	postgresql.DB.SetConnMaxLifetime(time.Duration(config.SetConnMaxLifetime()) * time.Second)

	return &postgresql, nil
}

func (pg *db) QueryContext(ctx context.Context, query string, args ...interface{}) (SQLRows, error) {
	return pg.DB.QueryContext(ctx, query, args...)
}

func (pg *db) DoTransactionWithContext(ctx context.Context, fnStmt ExecStmt) error {
	var (
		tx  *sql.Tx
		err error
	)

	if tx, err = pg.DB.BeginTx(ctx, nil); err != nil {
		return err
	}

	if err = fnStmt(tx); err != nil {
		if rollbackError := tx.Rollback(); rollbackError != nil {
			return rollbackError
		}

		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (pg *db) Close(rows SQLRows) {
	if rows != nil {
		rows.Close()
	}
}

func (pg *db) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return pg.DB.ExecContext(ctx, query, args...)
}

func (pg *db) PingContext(ctx context.Context) error {
	return pg.DB.PingContext(ctx)
}

func (pg *db) Array(array interface{}) AnyArray {
	return pq.Array(array)
}
