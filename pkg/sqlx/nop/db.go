package nop

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"

	"github.com/jmoiron/sqlx"
)

type Nop struct{}

func NewDB() *Nop {
	return &Nop{}
}

// Begin no-operation implementation of sqlx.Begin
func (*Nop) Begin() (*sql.Tx, error) {
	return nil, nil
}

// BeginTx no-operation implementation of sqlx.BeginTx
func (*Nop) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return nil, nil
}

// BeginTxx no-operation implementation of sqlx.BeginTxx
func (*Nop) BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	return nil, nil
}

// Beginx no-operation implementation of sqlx.Beginx
func (*Nop) Beginx() (*sqlx.Tx, error) {
	return nil, nil
}

// BindNamed no-operation implementation of sqlx.BindNamed
func (*Nop) BindNamed(query string, arg interface{}) (string, []interface{}, error) {
	return "", nil, nil
}

// Close no-operation implementation of sqlx.Close
func (*Nop) Close() error {
	return nil
}

// Conn no-operation implementation of sqlx.Conn
func (*Nop) Conn(ctx context.Context) (*sql.Conn, error) {
	return nil, nil
}

// Connx no-operation implementation of sqlx.Connx
func (*Nop) Connx(ctx context.Context) (*sqlx.Conn, error) {
	return nil, nil
}

// Driver no-operation implementation of sqlx.Driver
func (*Nop) Driver() driver.Driver {
	// TODO implement a nop driver.Driver to prevent SIG_SEGV errors
	return nil
}

// DriverName no-operation implementation of sqlx.DriverName
func (*Nop) DriverName() string {
	return ""
}

// Exec no-operation implementation of sqlx.Exec
func (*Nop) Exec(query string, args ...any) (sql.Result, error) {
	return nil, nil
}

// ExecContext no-operation implementation of sqlx.ExecContext
func (*Nop) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return nil, nil
}

// Get no-operation implementation of sqlx.Get
func (*Nop) Get(dest interface{}, query string, args ...interface{}) error {
	return nil
}

// GetContext no-operation implementation of sqlx.GetContext
func (*Nop) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return nil
}

// MapperFunc no-operation implementation of sqlx.MapperFunc
func (*Nop) MapperFunc(mf func(string) string) {
	return
}

// MustBegin no-operation implementation of sqlx.MustBegin
func (*Nop) MustBegin() *sqlx.Tx {
	return nil
}

// MustBeginTx no-operation implementation of sqlx.MustBeginTx
func (*Nop) MustBeginTx(ctx context.Context, opts *sql.TxOptions) *sqlx.Tx {
	return nil
}

// MustExec no-operation implementation of sqlx.MustExec
func (*Nop) MustExec(query string, args ...interface{}) sql.Result {
	return nil
}

// MustExecContext no-operation implementation of sqlx.MustExecContext
func (*Nop) MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result {
	return nil
}

// NamedExec no-operation implementation of sqlx.NamedExec
func (*Nop) NamedExec(query string, arg interface{}) (sql.Result, error) {
	return nil, nil
}

// NamedExecContext no-operation implementation of sqlx.NamedExecContext
func (*Nop) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return nil, nil
}

// NamedQuery no-operation implementation of sqlx.NamedQuery
func (*Nop) NamedQuery(query string, arg interface{}) (*sqlx.Rows, error) {
	return nil, nil
}

// NamedQueryContext no-operation implementation of sqlx.NamedQueryContext
func (*Nop) NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error) {
	return nil, nil
}

// Ping no-operation implementation of sqlx.Ping
func (*Nop) Ping() error {
	return nil
}

// PingContext no-operation implementation of sqlx.PingContext
func (*Nop) PingContext(ctx context.Context) error {
	return nil
}

// Prepare no-operation implementation of sqlx.Prepare
func (*Nop) Prepare(query string) (*sql.Stmt, error) {
	return nil, nil
}

// PrepareContext no-operation implementation of sqlx.PrepareContext
func (*Nop) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return nil, nil
}

// PrepareNamed no-operation implementation of sqlx.PrepareNamed
func (*Nop) PrepareNamed(query string) (*sqlx.NamedStmt, error) {
	return nil, nil
}

// PrepareNamedContext no-operation implementation of sqlx.PrepareNamedContext
func (*Nop) PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error) {
	return nil, nil
}

// Preparex no-operation implementation of sqlx.Preparex
func (*Nop) Preparex(query string) (*sqlx.Stmt, error) {
	return nil, nil
}

// PreparexContext no-operation implementation of sqlx.PreparexContext
func (*Nop) PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error) {
	return nil, nil
}

// Query no-operation implementation of sqlx.Query
func (*Nop) Query(query string, args ...any) (*sql.Rows, error) {
	return nil, nil
}

// QueryContext no-operation implementation of sqlx.QueryContext
func (*Nop) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return nil, nil
}

// QueryRow no-operation implementation of sqlx.QueryRow
func (*Nop) QueryRow(query string, args ...any) *sql.Row {
	return nil
}

// QueryRowContext no-operation implementation of sqlx.QueryRowContext
func (*Nop) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	return nil
}

// QueryRowx no-operation implementation of sqlx.QueryRowx
func (*Nop) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	return nil
}

// QueryRowxContext no-operation implementation of sqlx.QueryRowxContext
func (*Nop) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	return nil
}

// Queryx no-operation implementation of sqlx.Queryx
func (*Nop) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	return nil, nil
}

// QueryxContext no-operation implementation of sqlx.QueryxContext
func (*Nop) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	return nil, nil
}

// Rebind no-operation implementation of sqlx.Rebind
func (*Nop) Rebind(query string) string {
	return ""
}

// Select no-operation implementation of sqlx.Select
func (*Nop) Select(dest interface{}, query string, args ...interface{}) error {
	return nil
}

// SelectContext no-operation implementation of sqlx.SelectContext
func (*Nop) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return nil
}

// SetConnMaxIdleTime no-operation implementation of sqlx.SetConnMaxIdleTime
func (*Nop) SetConnMaxIdleTime(d time.Duration) {
}

// SetConnMaxLifetime no-operation implementation of sqlx.SetConnMaxLifetime
func (*Nop) SetConnMaxLifetime(d time.Duration) {
}

// SetMaxIdleConns no-operation implementation of sqlx.SetMaxIdleConns
func (*Nop) SetMaxIdleConns(n int) {
}

// SetMaxOpenConns no-operation implementation of sqlx.SetMaxOpenConns
func (*Nop) SetMaxOpenConns(n int) {
}

// Stats no-operation implementation of sqlx.Stats
func (*Nop) Stats() sql.DBStats {
	return sql.DBStats{}
}

// Unsafe no-operation implementation of sqlx.Unsafe
func (*Nop) Unsafe() *sqlx.DB {
	return nil
}
