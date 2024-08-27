package prometheus

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"

	"github.com/jmoiron/sqlx"
	kryptonsqlx "github.com/olireadcopper/sqlxprototype/pkg/sqlx"
	"github.com/olireadcopper/sqlxprototype/pkg/sqlx/nop"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type DB struct {
	inner kryptonsqlx.DB
}

type DBOption func(db *DB)

var (
	beginCount    prometheus.Counter
	beginErrors   prometheus.Counter
	beginDuration prometheus.Histogram
	execCount     *prometheus.CounterVec
	execErrors    *prometheus.CounterVec
	execDuration  *prometheus.HistogramVec
)

func init() {
	reg := prometheus.NewRegistry()
	factory := promauto.With(reg)

	beginCount = factory.NewCounter(
		prometheus.CounterOpts{
			Namespace: "copper.co",
			Name:      "sql_begin",
			Help:      "Number of calls to begin an SQL transaction",
		},
	)

	beginErrors = factory.NewCounter(
		prometheus.CounterOpts{
			Namespace: "copper.co",
			Name:      "sql_begin_errors",
			Help:      "Number of erors from trying to begin an SQL transaction",
		},
	)

	beginDuration = factory.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "copper.co",
			Name:      "sql_begin_errors",
			Help:      "Duration fo calls to begin an SQL transaction, measured in seconds",
			Buckets:   []float64{0.1, 0.2, 0.3, 0.5, 1},
		},
	)

	execCount = factory.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "copper.co",
			Name:      "sql_exec_count",
			Help:      "Number of calls to execute an SQL query",
		},
		[]string{"query"},
	)

	execErrors = factory.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "copper.co",
			Name:      "sql_exec_errors",
			Help:      "Number of erors from trying to execute an SQL query",
		},
		[]string{"query"},
	)

	execDuration = factory.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "copper.co",
			Name:      "sql_exec_duration",
			Help:      "Duration of execution of an SQL query, measured in seconds",
			Buckets:   []float64{0.1, 0.2, 0.3, 0.5, 1},
		},
		[]string{"query"},
	)
}

func NewDB(opts ...DBOption) kryptonsqlx.DB {
	db := &DB{
		inner: nop.NewDB(),
	}

	for _, opt := range opts {
		opt(db)
	}

	return db
}

// Begin prometheus instrumentation implementation of sqlx.Begin
func (db *DB) Begin() (*sql.Tx, error) {
	begin := time.Now()
	defer beginDuration.Observe(time.Since(begin).Seconds())

	beginCount.Inc()

	tx, err := db.inner.Begin()
	if err != nil {
		beginErrors.Inc()
	}

	return tx, err
}

// BeginTx prometheus instrumentation implementation of sqlx.BeginTx
func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	begin := time.Now()
	defer beginDuration.Observe(time.Since(begin).Seconds())

	beginCount.Inc()

	tx, err := db.inner.BeginTx(ctx, opts)
	if err != nil {
		beginErrors.Inc()
	}

	return tx, err
}

// BeginTxx prometheus instrumentation implementation of sqlx.BeginTxx
func (db *DB) BeginTxx(ctx context.Context, opts *sql.TxOptions) (*sqlx.Tx, error) {
	begin := time.Now()
	defer beginDuration.Observe(time.Since(begin).Seconds())

	beginCount.Inc()

	tx, err := db.inner.BeginTxx(ctx, opts)
	if err != nil {
		beginErrors.Inc()
	}

	return tx, err
}

// Beginx prometheus instrumentation implementation of sqlx.Beginx
func (db *DB) Beginx() (*sqlx.Tx, error) {
	beginCount.Inc()

	sql, err := db.inner.Beginx()
	if err != nil {
		beginErrors.Inc()
	}

	return sql, err
}

// BindNamed prometheus instrumentation implementation of sqlx.BindNamed, does not
// gather metrics as it's uneccesary
func (db *DB) BindNamed(query string, arg interface{}) (string, []interface{}, error) {
	return db.inner.BindNamed(query, arg)
}

// Close prometheus instrumentation implementation of sqlx.Close, does not
// gather metrics as it's uneccesary
func (db *DB) Close() error {
	return db.inner.Close()
}

// Conn prometheus instrumentation implementation of sqlx.Conn, does not
// gather metrics as it's uneccesary
func (db *DB) Conn(ctx context.Context) (*sql.Conn, error) {
	return db.inner.Conn(ctx)
}

// Connx prometheus instrumentation implementation of sqlx.Connx, does not
// gather metrics as it's uneccesary
func (db *DB) Connx(ctx context.Context) (*sqlx.Conn, error) {
	return db.inner.Connx(ctx)
}

// Driver prometheus instrumentation implementation of sqlx.Driver, does not
// gather metrics as it's uneccesary
func (db *DB) Driver() driver.Driver {
	return db.inner.Driver()

}

// DriverName prometheus instrumentation implementation of sqlx.DriverName, does not
// gather metrics as it's uneccesary
func (db *DB) DriverName() string {
	return db.inner.DriverName()

}

// Exec prometheus instrumentation implementation of sqlx.Exec
func (db *DB) Exec(query string, args ...any) (sql.Result, error) {
	begin := time.Now()
	defer execDuration.With(prometheus.Labels{
		"query": query,
	}).Observe(time.Since(begin).Seconds())

	execCount.With(prometheus.Labels{
		"query": query,
	}).Inc()

	res, err := db.inner.Exec(query, args...)

	if err != nil {
		execErrors.With(prometheus.Labels{
			"query": query,
		}).Inc()
	}

	return res, err
}

// ExecContext prometheus instrumentation implementation of sqlx.ExecContext
func (db *DB) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	begin := time.Now()
	defer execDuration.With(prometheus.Labels{
		"query": query,
	}).Observe(time.Since(begin).Seconds())

	execCount.With(prometheus.Labels{
		"query": query,
	}).Inc()

	res, err := db.inner.ExecContext(ctx, query, args...)

	if err != nil {
		execErrors.With(prometheus.Labels{
			"query": query,
		}).Inc()
	}

	return res, err
}

// Get prometheus instrumentation implementation of sqlx.Get
func (db *DB) Get(dest interface{}, query string, args ...interface{}) error {
	begin := time.Now()
	defer execDuration.With(prometheus.Labels{
		"query": query,
	}).Observe(time.Since(begin).Seconds())

	execCount.With(prometheus.Labels{
		"query": query,
	}).Inc()

	err := db.inner.Get(dest, query, args...)

	if err != nil {
		execErrors.With(prometheus.Labels{
			"query": query,
		}).Inc()
	}

	return err
}

// GetContext prometheus instrumentation implementation of sqlx.GetContext
func (db *DB) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	begin := time.Now()
	defer execDuration.With(prometheus.Labels{
		"query": query,
	}).Observe(time.Since(begin).Seconds())

	execCount.With(prometheus.Labels{
		"query": query,
	}).Inc()

	err := db.inner.GetContext(ctx, dest, query, args...)

	if err != nil {
		execErrors.With(prometheus.Labels{
			"query": query,
		}).Inc()
	}

	return err
}

// MapperFunc prometheus instrumentation implementation of sqlx.MapperFunc, does not
// gather metrics as it's uneccesary
func (db *DB) MapperFunc(mf func(string) string) {
	db.inner.MapperFunc(mf)

}

// MustBegin prometheus instrumentation implementation of sqlx.MustBegin
func (db *DB) MustBegin() *sqlx.Tx {
	begin := time.Now()
	defer beginDuration.Observe(time.Since(begin).Seconds())

	beginCount.Inc()

	return db.inner.MustBegin()
}

// MustBeginTx prometheus instrumentation implementation of sqlx.MustBeginTx
func (db *DB) MustBeginTx(ctx context.Context, opts *sql.TxOptions) *sqlx.Tx {
	begin := time.Now()
	defer beginDuration.Observe(time.Since(begin).Seconds())

	beginCount.Inc()

	return db.inner.MustBeginTx(ctx, opts)
}

// MustExec prometheus instrumentation implementation of sqlx.MustExec
func (db *DB) MustExec(query string, args ...interface{}) sql.Result {
	begin := time.Now()
	defer execDuration.With(prometheus.Labels{
		"query": query,
	}).Observe(time.Since(begin).Seconds())

	execCount.With(prometheus.Labels{
		"query": query,
	}).Inc()

	return db.inner.MustExec(query, args...)
}

// MustExecContext prometheus instrumentation implementation of sqlx.MustExecContext
func (db *DB) MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result {
	begin := time.Now()
	defer execDuration.With(prometheus.Labels{
		"query": query,
	}).Observe(time.Since(begin).Seconds())

	execCount.With(prometheus.Labels{
		"query": query,
	}).Inc()

	return db.inner.MustExecContext(ctx, query, args...)
}

// NamedExec prometheus instrumentation implementation of sqlx.NamedExec
func (db *DB) NamedExec(query string, arg interface{}) (sql.Result, error) {
	begin := time.Now()
	defer execDuration.With(prometheus.Labels{
		"query": query,
	}).Observe(time.Since(begin).Seconds())

	execCount.With(prometheus.Labels{
		"query": query,
	}).Inc()

	res, err := db.inner.NamedExec(query, arg)
	if err != nil {
		execErrors.With(prometheus.Labels{
			"query": query,
		}).Inc()
	}

	return res, err
}

// NamedExecContext prometheus instrumentation implementation of sqlx.NamedExecContext
func (db *DB) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	begin := time.Now()
	defer execDuration.With(prometheus.Labels{
		"query": query,
	}).Observe(time.Since(begin).Seconds())

	execCount.With(prometheus.Labels{
		"query": query,
	}).Inc()

	res, err := db.inner.NamedExecContext(ctx, query, arg)
	if err != nil {
		execErrors.With(prometheus.Labels{
			"query": query,
		}).Inc()
	}

	return res, err
}

// NamedQuery prometheus instrumentation implementation of sqlx.NamedQuery
func (db *DB) NamedQuery(query string, arg interface{}) (*sqlx.Rows, error) {
	begin := time.Now()
	defer execDuration.With(prometheus.Labels{
		"query": query,
	}).Observe(time.Since(begin).Seconds())

	execCount.With(prometheus.Labels{
		"query": query,
	}).Inc()

	res, err := db.inner.NamedQuery(query, arg)
	if err != nil {
		execErrors.With(prometheus.Labels{
			"query": query,
		}).Inc()
	}

	return res, err
}

// NamedQueryContext prometheus instrumentation implementation of sqlx.NamedQueryContext
func (db *DB) NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error) {
	begin := time.Now()
	defer execDuration.With(prometheus.Labels{
		"query": query,
	}).Observe(time.Since(begin).Seconds())

	execCount.With(prometheus.Labels{
		"query": query,
	}).Inc()

	res, err := db.inner.NamedQueryContext(ctx, query, arg)
	if err != nil {
		execErrors.With(prometheus.Labels{
			"query": query,
		}).Inc()
	}

	return res, err
}

// Ping prometheus instrumentation implementation of sqlx.Ping, does not
// gather metrics as it's uneccesary
func (db *DB) Ping() error {
	return db.inner.Ping()
}

// PingContext prometheus instrumentation implementation of sqlx.PingContext,
// does not gather metrics as it's uneccesary
func (db *DB) PingContext(ctx context.Context) error {
	return db.inner.PingContext(ctx)
}

// Prepare prometheus instrumentation implementation of sqlx.Prepare, does not
// gather metrics as it's uneccesary
func (db *DB) Prepare(query string) (*sql.Stmt, error) {
	return db.inner.Prepare(query)
}

// PrepareContext prometheus instrumentation implementation of sqlx.PrepareContext, does not
// gather metrics as it's uneccesary
func (db *DB) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	return db.inner.PrepareContext(ctx, query)
}

// PrepareNamed prometheus instrumentation implementation of sqlx.PrepareNamed, does not
// gather metrics as it's uneccesary
func (db *DB) PrepareNamed(query string) (*sqlx.NamedStmt, error) {
	return db.inner.PrepareNamed(query)
}

// PrepareNamedContext prometheus instrumentation implementation of sqlx.PrepareNamedContext, does not
// gather metrics as it's uneccesary
func (db *DB) PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error) {
	return db.inner.PrepareNamedContext(ctx, query)
}

// Preparex prometheus instrumentation implementation of sqlx.Preparex, does not
// gather metrics as it's uneccesary
func (db *DB) Preparex(query string) (*sqlx.Stmt, error) {
	return db.inner.Preparex(query)
}

// PreparexContext prometheus instrumentation implementation of sqlx.PreparexContext, does not
// gather metrics as it's uneccesary
func (db *DB) PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error) {
	return db.inner.PreparexContext(ctx, query)
}

// Query prometheus instrumentation implementation of sqlx.Query
func (db *DB) Query(query string, args ...any) (*sql.Rows, error) {
	begin := time.Now()
	defer execDuration.With(prometheus.Labels{
		"query": query,
	}).Observe(time.Since(begin).Seconds())

	execCount.With(prometheus.Labels{
		"query": query,
	}).Inc()

	res, err := db.inner.Query(query, args...)
	if err != nil {
		execErrors.With(prometheus.Labels{
			"query": query,
		}).Inc()
	}

	return res, err
}

// QueryContext prometheus instrumentation implementation of sqlx.QueryContext
func (db *DB) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	begin := time.Now()
	defer execDuration.With(prometheus.Labels{
		"query": query,
	}).Observe(time.Since(begin).Seconds())

	execCount.With(prometheus.Labels{
		"query": query,
	}).Inc()

	res, err := db.inner.QueryContext(ctx, query, args...)
	if err != nil {
		execErrors.With(prometheus.Labels{
			"query": query,
		}).Inc()
	}

	return res, err
}

// QueryRow prometheus instrumentation implementation of sqlx.QueryRow
func (db *DB) QueryRow(query string, args ...any) *sql.Row {
	begin := time.Now()
	defer execDuration.With(prometheus.Labels{
		"query": query,
	}).Observe(time.Since(begin).Seconds())

	execCount.With(prometheus.Labels{
		"query": query,
	}).Inc()

	return db.inner.QueryRow(query, args...)
}

// QueryRowContext prometheus instrumentation implementation of sqlx.QueryRowContext
func (db *DB) QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row {
	begin := time.Now()
	defer execDuration.With(prometheus.Labels{
		"query": query,
	}).Observe(time.Since(begin).Seconds())

	execCount.With(prometheus.Labels{
		"query": query,
	}).Inc()

	return db.inner.QueryRowContext(ctx, query, args...)
}

// QueryRowx prometheus instrumentation implementation of sqlx.QueryRowx
func (db *DB) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	begin := time.Now()
	defer execDuration.With(prometheus.Labels{
		"query": query,
	}).Observe(time.Since(begin).Seconds())

	execCount.With(prometheus.Labels{
		"query": query,
	}).Inc()

	return db.inner.QueryRowx(query, args...)
}

// QueryRowxContext prometheus instrumentation implementation of sqlx.QueryRowxContext
func (db *DB) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	begin := time.Now()
	defer execDuration.With(prometheus.Labels{
		"query": query,
	}).Observe(time.Since(begin).Seconds())

	execCount.With(prometheus.Labels{
		"query": query,
	}).Inc()

	return db.inner.QueryRowxContext(ctx, query, args...)
}

// Queryx prometheus instrumentation implementation of sqlx.Queryx
func (db *DB) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	begin := time.Now()
	defer execDuration.With(prometheus.Labels{
		"query": query,
	}).Observe(time.Since(begin).Seconds())

	execCount.With(prometheus.Labels{
		"query": query,
	}).Inc()

	res, err := db.inner.Queryx(query, args...)
	if err != nil {
		execErrors.With(prometheus.Labels{
			"query": query,
		}).Inc()
	}

	return res, err
}

// QueryxContext prometheus instrumentation implementation of sqlx.QueryxContext
func (db *DB) QueryxContext(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	begin := time.Now()
	defer execDuration.With(prometheus.Labels{
		"query": query,
	}).Observe(time.Since(begin).Seconds())

	execCount.With(prometheus.Labels{
		"query": query,
	}).Inc()

	res, err := db.inner.QueryxContext(ctx, query, args...)
	if err != nil {
		execErrors.With(prometheus.Labels{
			"query": query,
		}).Inc()
	}

	return res, err
}

// Rebind prometheus instrumentation implementation of sqlx.Rebind, does not
// gather metrics as it's uneccesary
func (db *DB) Rebind(query string) string {
	return db.inner.Rebind(query)
}

// Select prometheus instrumentation implementation of sqlx.Select
func (db *DB) Select(dest interface{}, query string, args ...interface{}) error {
	return db.inner.Select(dest, query, args)
}

// SelectContext prometheus instrumentation implementation of sqlx.SelectContext
func (db *DB) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return db.inner.SelectContext(ctx, dest, query, args)
}

// SetConnMaxIdleTime prometheus instrumentation implementation of sqlx.SetConnMaxIdleTime, does not
// gather metrics as it's uneccesary
func (db *DB) SetConnMaxIdleTime(d time.Duration) {
	db.inner.SetConnMaxIdleTime(d)
}

// SetConnMaxLifetime prometheus instrumentation implementation of sqlx.SetConnMaxLifetime, does not
// gather metrics as it's uneccesary
func (db *DB) SetConnMaxLifetime(d time.Duration) {
	db.inner.SetConnMaxLifetime(d)
}

// SetMaxIdleConns prometheus instrumentation implementation of sqlx.SetMaxIdleConns, does not
// gather metrics as it's uneccesary
func (db *DB) SetMaxIdleConns(n int) {
	db.inner.SetMaxIdleConns(n)
}

// SetMaxOpenConns prometheus instrumentation implementation of sqlx.SetMaxOpenConns, does not
// gather metrics as it's uneccesary
func (db *DB) SetMaxOpenConns(n int) {
	db.inner.SetMaxOpenConns(n)
}

// Stats prometheus instrumentation implementation of sqlx.Stats, does not
// gather metrics as it's uneccesary
func (db *DB) Stats() sql.DBStats {
	return db.inner.Stats()
}

// Unsafe prometheus instrumentation implementation of sqlx.Unsafe, does not
// gather metrics as it's uneccesary
func (db *DB) Unsafe() *sqlx.DB {
	return db.inner.Unsafe()
}
