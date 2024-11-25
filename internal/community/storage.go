package community

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExist     = errors.New("row does not exist")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type Repository interface {
	Writer
	Reader
	Service
}

type Writer interface {
	Create(ctx context.Context, p *Post) (sql.Result, error)
}

type Reader interface {
	Get(ctx context.Context, id int) (*Post, error)
	GetByTitle(ctx context.Context, t string) (*Post, error)
	All(ctx context.Context) ([]Post, error)
}

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
}
