package logging

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"

	kryptonsqlx "github.com/olireadcopper/sqlxprototype/pkg/sqlx"
	"github.com/olireadcopper/sqlxprototype/pkg/sqlx/nop"
	"github.com/olireadcopper/sqlxprototype/pkg/telemetry"
	"github.com/olireadcopper/sqlxprototype/pkg/telemetry/logging"
)

type DB struct {
	logger *slog.Logger
	inner  kryptonsqlx.DB
}

type DBOption func(db *DB)

func NewDB(opts ...DBOption) kryptonsqlx.DB {
	db := &DB{
		logger: logging.NopLogger,
		inner:  nop.NewDB(),
	}

	for _, opt := range opts {
		opt(db)
	}

	return db
}

// Begin logging implmentation of sqlx.Begin
func (db *DB) Begin() (*sql.Tx, error) {
	tx, err := db.inner.Begin()

	db.log(context.TODO(), "", err,
		logging.FieldMethod, "Begin",
	)

	return tx, err
}

// BeginTx logging implmentation of sqlx.BeginTx
func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	tx, err := db.inner.BeginTx(ctx, opts)

	db.log(ctx, "", err,
		logging.FieldMethod, "BeginTx",
		fieldTransaction, tx,
	)

	return tx, err
}

// BeginTxx logging implmentation of sqlx.BeginTxx
func (db *DB) BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	tx, err := db.inner.BeginTxx(ctx, opts)

	db.log(ctx, "", err,
		logging.FieldMethod, "BeginTxx",
		fieldTransaction, tx,
	)

	return tx, err
}

// Beginx logging implmentation of sqlx.Beginx
func (db *DB) Beginx() (*sqlx.Tx, error) {
	tx, err := db.inner.Beginx()

	db.log(context.TODO(), "", err,
		logging.FieldMethod, "Beginx",
		fieldTransaction, tx,
	)

	return tx, err
}

// BindNamed logging implmentation of sqlx.BindNamed
func (db *DB) BindNamed(query string, arg interface{}) (string, []interface{}, error) {
	res, args, err := db.inner.BindNamed(query, arg)

	db.log(context.TODO(), "", err,
		logging.FieldMethod, "BindNamed",
		fieldQuery, query,
		fieldArgs, arg,
		fieldResult, res,
	)

	return res, args, err
}

// Close logging implmentation of sqlx.Close
func (db *DB) Close() error {
	err := db.inner.Close()

	db.log(context.TODO(), "", err,
		logging.FieldMethod, "Close",
	)

	return nil
}

// Conn logging implmentation of sqlx.Conn
func (db *DB) Conn(ctx context.Context) (*sql.Conn, error) {
	conn, err := db.inner.Conn(ctx)

	db.log(ctx, "", err,
		logging.FieldMethod, "Conn",
		fieldConn, conn,
	)

	return conn, err
}

// Connx logging implmentation of sqlx.Connx
func (db *DB) Connx(ctx context.Context) (*sqlx.Conn, error) {
	conn, err := db.inner.Connx(ctx)

	db.log(ctx, "", err,
		logging.FieldMethod, "Connx",
		fieldConn, conn,
	)
	return conn, err
}

// Driver logging implmentation of sqlx.Driver
func (db *DB) Driver() driver.Driver {
	driver := db.inner.Driver()

	db.log(context.TODO(), "", nil,
		logging.FieldMethod, "Driver",
		"driver", driver,
	)

	return driver
}

// DriverName logging implmentation of sqlx.DriverName
func (db *DB) DriverName() string {
	driver := db.inner.DriverName()

	db.log(context.TODO(), "", nil,
		logging.FieldMethod, "DriverName",
		"driverName", driver,
	)

	return driver
}

// Exec logging implmentation of sqlx.Exec
func (db *DB) Exec(query string, args ...any) (sql.Result, error) {
	res, err := db.inner.Exec(query, args...)

	db.log(context.TODO(), "", err,
		logging.FieldMethod, "Exec",
		fieldQuery, query,
		fieldArgs, args,
		fieldResult, res,
	)

	return res, err
}

// ExecContext logging implmentation of sqlx.ExecContext
func (db *DB) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	res, err := db.inner.ExecContext(ctx, query, args...)

	db.log(ctx, "", err,
		logging.FieldMethod, "ExecContext",
		fieldQuery, query,
		fieldArgs, args,
		fieldResult, res,
	)

	return res, err
}

// Get logging implmentation of sqlx.Get
func (db *DB) Get(dest interface{}, query string, args ...interface{}) error {
	err := db.inner.Get(dest, query, args...)

	db.log(context.TODO(), "", err,
		logging.FieldMethod, "Get",
		fieldQuery, query,
		fieldArgs, args,
		fieldDest, dest,
	)

	return err
}

// GetContext logging implmentation of sqlx.GetContext
func (db *DB) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	err := db.inner.GetContext(ctx, dest, query, args...)

	db.log(ctx, "", err,
		logging.FieldMethod, "GetContext",
		fieldQuery, query,
		fieldArgs, args,
		fieldDest, dest,
	)

	return err
}

// MapperFunc logging implmentation of sqlx.MapperFunc
func (db *DB) MapperFunc(mf func(string) string) {
	db.inner.MapperFunc(mf)

	db.log(context.TODO(), "", nil,
		logging.FieldMethod, "MapperFunc",
	)
}

// MustBegin logging implmentation of sqlx.MustBegin
func (db *DB) MustBegin() *sqlx.Tx {
	tx := db.inner.MustBegin()

	db.log(context.TODO(), "", nil,
		logging.FieldMethod, "MustBegin",
		fieldTransaction, tx,
	)

	return tx
}

// MustBeginTx logging implmentation of sqlx.MustBeginTx
func (db *DB) MustBeginTx(ctx context.Context, opts *sql.TxOptions) *sqlx.Tx {
	tx := db.MustBegin()

	db.log(ctx, "", nil,
		logging.FieldMethod, "MustBeginTx",
		fieldTransaction, tx,
	)

	return tx
}

// MustExec logging implmentation of sqlx.MustExec
func (db *DB) MustExec(query string, args ...interface{}) sql.Result {
	res := db.inner.MustExec(query, args...)

	db.log(context.TODO(), "", nil,
		logging.FieldMethod, "MustExec",
		fieldQuery, query,
		fieldArgs, args,
		fieldResult, res,
	)

	return res
}

// MustExecContext logging implmentation of sqlx.MustExecContext
func (db *DB) MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result {
	res := db.inner.MustExecContext(ctx, query, args...)

	db.log(ctx, "", nil,
		logging.FieldMethod, "MustExecContext",
		fieldQuery, query,
		fieldArgs, args,
		fieldResult, res,
	)

	return res
}

// NamedExec logging implmentation of sqlx.NamedExec
func (db *DB) NamedExec(query string, arg interface{}) (sql.Result, error) {
	res, err := db.inner.NamedExec(query, arg)

	db.log(context.TODO(), "", err,
		logging.FieldMethod, "NamedExec",
		fieldQuery, query,
		fieldArgs, arg,
		fieldResult, res,
	)

	return res, err
}

// NamedExecContext logging implmentation of sqlx.NamedExecContext
func (db *DB) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	res, err := db.inner.NamedExecContext(ctx, query, arg)

	db.log(ctx, "", err,
		logging.FieldMethod, "NamedExecContext",
		fieldQuery, query,
		fieldArgs, arg,
		fieldResult, res,
	)

	return res, err
}

// NamedQuery logging implmentation of sqlx.NamedQuery
func (db *DB) NamedQuery(query string, arg interface{}) (*sqlx.Rows, error) {
	rows, err := db.inner.NamedQuery(query, arg)

	db.log(context.TODO(), "", err,
		logging.FieldMethod, "NamedQuery",
		fieldQuery, query,
		fieldArgs, arg,
		fieldRows, rows,
	)

	return rows, err
}

// NamedQueryContext logging implmentation of sqlx.NamedQueryContext
func (db *DB) NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error) {
	rows, err := db.inner.NamedQueryContext(ctx, query, arg)

	db.log(ctx, "", err,
		logging.FieldMethod, "NamedQueryContext",
		fieldQuery, query,
		fieldArgs, arg,
		fieldRows, rows,
	)

	return rows, err
}

// Ping logging implmentation of sqlx.Ping
func (db *DB) Ping() error {
	err := db.inner.Ping()

	db.log(context.TODO(), "", err,
		logging.FieldMethod, "Ping",
	)

	return err
}

// PingContext logging implmentation of sqlx.PingContext
func (db *DB) PingContext(ctx context.Context) error {
	err := db.inner.PingContext(ctx)

	db.log(ctx, "", err,
		logging.FieldMethod, "Ping",
	)

	return err
}

// Prepare logging implmentation of sqlx.Prepare
func (db *DB) Prepare(query string) (*sql.Stmt, error) {
	stmt, err := db.inner.Prepare(query)

	db.log(context.TODO(), "", err,
		logging.FieldMethod, "Prepare",
		fieldQuery, query,
		fieldStatement, stmt,
	)

	return stmt, err
}

// PrepareContext logging implmentation of sqlx.PrepareContext
func (db *DB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	stmt, err := db.inner.PrepareContext(ctx, query)

	db.log(ctx, "", err,
		logging.FieldMethod, "PrepareContext",
		fieldQuery, query,
		fieldStatement, stmt,
	)

	return stmt, err
}

// PrepareNamed logging implmentation of sqlx.PrepareNamed
func (db *DB) PrepareNamed(query string) (*sqlx.NamedStmt, error) {
	stmt, err := db.inner.PrepareNamed(query)

	db.log(context.TODO(), "", err,
		logging.FieldMethod, "PrepareNamed",
		fieldQuery, query,
		fieldStatement, stmt,
	)

	return stmt, err
}

// PrepareNamedContext logging implmentation of sqlx.PrepareNamedContext
func (db *DB) PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error) {
	stmt, err := db.inner.PrepareNamedContext(ctx, query)

	db.log(ctx, "", err,
		logging.FieldMethod, "PrepareNamedContext",
		fieldQuery, query,
		fieldStatement, stmt,
	)

	return stmt, err
}

// Preparex logging implmentation of sqlx.Preparex
func (db *DB) Preparex(query string) (*sqlx.Stmt, error) {
	stmt, err := db.inner.Preparex(query)

	db.log(context.TODO(), "", err,
		logging.FieldMethod, "Preparex",
		fieldQuery, query,
		fieldStatement, stmt,
	)

	return stmt, err
}

// PreparexContext logging implmentation of sqlx.PreparexContext
func (db *DB) PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error) {
	stmt, err := db.inner.PreparexContext(ctx, query)

	db.log(ctx, "", err,
		logging.FieldMethod, "PreparexContext",
		fieldQuery, query,
		fieldStatement, stmt,
	)

	return stmt, err
}

// Query logging implmentation of sqlx.Query
func (db *DB) Query(query string, args ...any) (*sql.Rows, error) {
	rows, err := db.inner.Query(query, args)

	db.log(context.TODO(), "", err,
		logging.FieldMethod, "Query",
		fieldQuery, query,
		fieldArgs, args,
		fieldRows, rows,
	)

	return rows, err
}

// QueryContext logging implmentation of sqlx.QueryContext
func (db *DB) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	rows, err := db.inner.QueryContext(ctx, query, args)

	db.log(ctx, "", err,
		logging.FieldMethod, "QueryContext",
		fieldQuery, query,
		fieldArgs, args,
		fieldRows, rows,
	)

	return rows, err
}

// QueryRow logging implmentation of sqlx.QueryRow
func (db *DB) QueryRow(query string, args ...any) *sql.Row {
	row := db.inner.QueryRow(query, args)

	db.log(context.TODO(), "", nil,
		logging.FieldMethod, "QueryRow",
		fieldQuery, query,
		fieldArgs, args,
		fieldRows, row,
	)

	return row
}

// QueryRowContext logging implmentation of sqlx.QueryRowContext
func (db *DB) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	row := db.inner.QueryRowContext(ctx, query, args)

	db.log(ctx, "", nil,
		logging.FieldMethod, "QueryRowContext",
		fieldQuery, query,
		fieldArgs, args,
		fieldRows, row,
	)

	return row
}

// QueryRowx logging implmentation of sqlx.QueryRowx
func (db *DB) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	row := db.inner.QueryRowx(query, args)

	db.log(context.TODO(), "", nil,
		logging.FieldMethod, "QueryRowx",
		fieldQuery, query,
		fieldArgs, args,
		fieldRows, row,
	)

	return row
}

// QueryRowxContext logging implmentation of sqlx.QueryRowxContext
func (db *DB) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	row := db.inner.QueryRowxContext(ctx, query, args)

	db.log(ctx, "", nil,
		logging.FieldMethod, "QueryRowxContext",
		fieldQuery, query,
		fieldArgs, args,
		fieldRows, row,
	)

	return row
}

// Queryx logging implmentation of sqlx.Queryx
func (db *DB) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	rows, err := db.inner.Queryx(query, args)

	db.log(context.TODO(), "", err,
		logging.FieldMethod, "Queryx",
		fieldQuery, query,
		fieldArgs, args,
		fieldRows, rows,
	)

	return rows, err
}

// QueryxContext logging implmentation of sqlx.QueryxContext
func (db *DB) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	rows, err := db.inner.QueryxContext(ctx, query, args)

	db.log(ctx, "", err,
		logging.FieldMethod, "QueryxContext",
		fieldQuery, query,
		fieldArgs, args,
		fieldRows, rows,
	)
	return rows, err
}

// Rebind logging implmentation of sqlx.Rebind
func (db *DB) Rebind(query string) string {
	rebind := db.inner.Rebind(query)

	db.log(context.TODO(), "", nil,
		logging.FieldMethod, "Rebind",
		fieldQuery, query,
		"rebind", rebind,
	)

	return rebind
}

// Select logging implmentation of sqlx.Select
func (db *DB) Select(dest interface{}, query string, args ...interface{}) error {
	err := db.inner.Select(dest, query, args...)

	db.log(context.TODO(), "", err,
		logging.FieldMethod, "Select",
		fieldQuery, query,
		fieldArgs, args,
		fieldDest, dest,
	)

	return err
}

// SelectContext logging implmentation of sqlx.SelectContext
func (db *DB) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	err := db.inner.SelectContext(ctx, dest, query, args...)

	db.log(ctx, "", err,
		logging.FieldMethod, "Select",
		fieldQuery, query,
		fieldArgs, args,
		fieldDest, dest,
	)

	return err
}

// SetConnMaxIdleTime logging implmentation of sqlx.SetConnMaxIdleTime
func (db *DB) SetConnMaxIdleTime(d time.Duration) {
	db.inner.SetConnMaxIdleTime(d)

	db.log(context.TODO(), "", nil,
		logging.FieldMethod, "SetConnMaxIdleTime",
		fieldDuration, d,
	)
}

// SetConnMaxLifetime logging implmentation of sqlx.SetConnMaxLifetime
func (db *DB) SetConnMaxLifetime(d time.Duration) {
	db.inner.SetConnMaxLifetime(d)

	db.log(context.TODO(), "", nil,
		logging.FieldMethod, "SetConnMaxLifetime",
		fieldDuration, d,
	)
}

// SetMaxIdleConns logging implmentation of sqlx.SetMaxIdleConns
func (db *DB) SetMaxIdleConns(n int) {
	db.inner.SetMaxIdleConns(n)

	db.log(context.TODO(), "", nil,
		logging.FieldMethod, "SetMaxIdleConns",
		fieldConnections, n,
	)
}

// SetMaxOpenConns logging implmentation of sqlx.SetMaxOpenConns
func (db *DB) SetMaxOpenConns(n int) {
	db.inner.SetMaxOpenConns(n)

	db.log(context.TODO(), "", nil,
		logging.FieldMethod, "SetMaxOpenConns",
		fieldConnections, n,
	)
}

// Stats logging implmentation of sqlx.Stats
func (db *DB) Stats() sql.DBStats {
	stats := db.inner.Stats()

	db.log(context.TODO(), "", nil,
		logging.FieldMethod, "Stats",
		"stats", stats,
	)

	return stats
}

// Unsafe logging implmentation of sqlx.Unsafe
func (db *DB) Unsafe() *sqlx.DB {
	unsafe := db.inner.Unsafe()

	db.log(context.TODO(), "", nil,
		logging.FieldMethod, "Unsafe",
	)

	return unsafe
}

func (db *DB) log(ctx context.Context, msg string, err error, args ...any) {
	if trace := ctx.Value(telemetry.ContextKeyTrace); trace != nil {
		args = append(args, logging.FieldTrace, trace)
	}

	if err != nil {
		db.logError(err, args...)
		return
	}

	db.logInfo(msg, args...)
}

func (db *DB) logInfo(msg string, args ...any) {
	db.logger.Info(msg, args...)
}

func (db *DB) logError(err error, args ...any) {
	args = append(args, logging.FieldErrorCode, errors.Unwrap(err).Error())

	db.logger.Error(err.Error(), args...)
}
