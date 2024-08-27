package logging

import (
	"log/slog"

	"github.com/olireadcopper/sqlxprototype/pkg/sqlx"
	"github.com/olireadcopper/sqlxprototype/pkg/telemetry/logging"
)

func DBWithLogger(l *slog.Logger) DBOption {
	return func(db *DB) {
		l.With(
			logging.FieldComponent, "sqlx",
		)

		db.logger = l
	}
}

func DBWithInnerDB(db sqlx.DB) DBOption {
	return func(d *DB) {
		d.inner = db
	}
}
