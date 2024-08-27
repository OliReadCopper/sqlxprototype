package sql

import (
	"log/slog"

	"github.com/olireadcopper/sqlxprototype/pkg/sqlx"
	"github.com/olireadcopper/sqlxprototype/pkg/sqlx/logging"
)

func SQLiteWithDB(db sqlx.DB) SQLiteRepositoryOption {
	return func(r *SQLiteRepository) {
		r.db = db
	}
}

func SQLiteWithLogging(l *slog.Logger) SQLiteRepositoryOption {
	return func(r *SQLiteRepository) {
		r.db = logging.NewDB(
			logging.DBWithInnerDB(r.db),
			logging.DBWithLogger(l),
		)
	}
}
