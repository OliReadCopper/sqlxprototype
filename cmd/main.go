package main

import (
	"context"
	"flag"
	"log/slog"
	"os"

	"github.com/jmoiron/sqlx"

	"github.com/olireadcopper/sqlxprototype/internal/model"
	"github.com/olireadcopper/sqlxprototype/internal/repository/taxonomy"
	"github.com/olireadcopper/sqlxprototype/internal/repository/taxonomy/sql"
	kryptonsqlx "github.com/olireadcopper/sqlxprototype/pkg/sqlx"
	kryptonsqlxprometheus "github.com/olireadcopper/sqlxprototype/pkg/sqlx/instrumenting/prometheus"
	krtpronsqlxlogging "github.com/olireadcopper/sqlxprototype/pkg/sqlx/logging"
)

func main() {
	logging := flag.Bool("logging", false, "enable logging in the application")
	instrumenting := flag.Bool("instrumenting", false, "enable instrumenting in the application")
	flag.Parse()

	pool, err := sqlx.Connect("sqllite3", "kryptonsdk")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	var repositoryDB kryptonsqlx.DB

	repositoryDB = pool

	if *logging {
		repositoryDB = krtpronsqlxlogging.NewDB(
			krtpronsqlxlogging.DBWithLogger(slog.Default()),
			krtpronsqlxlogging.DBWithInnerDB(repositoryDB),
		)
	}

	if *instrumenting {
		repositoryDB = kryptonsqlxprometheus.NewDB(
			kryptonsqlxprometheus.DBWithInnerDB(repositoryDB),
		)
	}

	taxonomyRepository := sql.NewSQLiteRepository(
		sql.SQLiteWithDB(repositoryDB),
	)

	taxonomyRepository.Read(context.Background(), taxonomy.Query{
		Taxonomy: model.Taxonomy{
			ID:   "id",
			Name: "taxonomy",
		},
	})
}
