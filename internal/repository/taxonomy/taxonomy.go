package taxonomy

import (
	"context"

	"github.com/olireadcopper/sqlxprototype/internal/model"
	"github.com/olireadcopper/sqlxprototype/internal/repository"
)

type Repository interface {
	Create(ctx context.Context, query Query) error
	Read(ctx context.Context, query Query) (Response, error)
	Update(ctx context.Context, query Query) error
	Delete(ctx context.Context, query Query) error
}

type Query struct {
	Pagination repository.Pagination
	Taxonomy   model.Taxonomy
}

type QueryOption func(*Query)

type ResponseOption func(*Response)

type Response struct {
	Pagination repository.Pagination
	Results    []model.Taxonomy
}

func NewQuery(opts ...QueryOption) Query {
	q := Query{}

	for _, opt := range opts {
		opt(&q)
	}

	return q
}

func NewResponse(opts ...ResponseOption) Response {
	r := Response{}

	for _, opt := range opts {
		opt(&r)
	}

	return r
}

func ResponseWithPagination(p repository.Pagination) ResponseOption {
	return func(r *Response) {
		r.Pagination = p
	}
}

func ResponseWithResult(t ...model.Taxonomy) ResponseOption {
	return func(r *Response) {
		r.Results = t
	}
}
