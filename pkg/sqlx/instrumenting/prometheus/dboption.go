package prometheus

import "github.com/olireadcopper/sqlxprototype/pkg/sqlx"

func DBWithInnerDB(db sqlx.DB) DBOption {
	return func(d *DB) {
		d.inner = db
	}
}
