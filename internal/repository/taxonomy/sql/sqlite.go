package sql

import (
	"context"

	"github.com/olireadcopper/sqlxprototype/internal/model"
	"github.com/olireadcopper/sqlxprototype/internal/repository"
	"github.com/olireadcopper/sqlxprototype/internal/repository/taxonomy"
	"github.com/olireadcopper/sqlxprototype/pkg/sqlx"
	"github.com/olireadcopper/sqlxprototype/pkg/sqlx/nop"
)

const (
	queryCreateTaxonomy = `
		INSERT INTO taxonomy (name) VALUES ($1)`
	queryReadTaxonomyByID = `
		SELECT (id, name) FROM taxonomy WHERE id=($1) LIMIT $2 OFFSET $3`
)

type SQLiteRepositoryOption func(*SQLiteRepository)

type SQLiteRepository struct {
	db sqlx.DB
}

// NewSQLiteRepository constructor for a new SQL repository utilising SQLite as
// a driver
func NewSQLiteRepository(opts ...SQLiteRepositoryOption) *SQLiteRepository {
	r := SQLiteRepository{
		db: nop.NewDB(),
	}

	for _, opt := range opts {
		opt(&r)
	}

	return &r
}

// Create SQLite implementation for a taxonomy repository
func (r *SQLiteRepository) Create(ctx context.Context, query taxonomy.Query) error {
	if _, err := r.db.ExecContext(ctx, queryCreateTaxonomy, query.Taxonomy.Name); err != nil {
		return err
	}

	return nil
}

// Read SQLite implementation for a taxonomy repository
func (r *SQLiteRepository) Read(ctx context.Context, query taxonomy.Query) (taxonomy.Response, error) {
	rows, err := r.db.QueryxContext(ctx, queryReadTaxonomyByID)
	if err != nil {
		return taxonomy.Response{}, err
	}

	taxonomies := []model.Taxonomy{}

	for rows.Next() {
		t := model.Taxonomy{}

		if err := rows.StructScan(&t); err != nil {
			return taxonomy.Response{}, err
		}

		taxonomies = append(taxonomies, t)
	}

	return taxonomy.NewResponse(
		taxonomy.ResponseWithPagination(repository.Pagination{
			Limit:  query.Pagination.Limit,
			Offset: query.Pagination.Offset,
			Count:  uint(len(taxonomies)),
		}),
		taxonomy.ResponseWithResult(taxonomies...),
	), nil
}

// Update SQLite implementation for a taxonomy repository
func (r *SQLiteRepository) Update(ctx context.Context, query taxonomy.Query) error {
	return nil
}

// Delete SQLite implementation for a taxonomy repository
func (r *SQLiteRepository) Delete(ctx context.Context, query taxonomy.Query) error {
	return nil
}
