package community

import (
	"context"
	"database/sql"
	"errors"
)

type Repository interface {
	Writer
	Reader
}

type Writer interface {
	Create(ctx context.Context, p *Post) (int, error)
}

type Reader interface {
	Get(ctx context.Context, id int) (*Post, error)
	GetByTitle(ctx context.Context, t string) (*Post, error)
	All(ctx context.Context) ([]Post, error)
}

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExist     = errors.New("row does not exist")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

func (s *Service) Create(p *Post) (sql.Result, error) {
	q := `INSERT INTO posts(title, content, reactions) VALUES($1, $2, $3);`
	r, err := s.db.Exec(q, p.Title, p.Content, p.Reactions)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (s *Service) Get(id int) (*Post, error) {
	var p Post
	q := `SELECT * FROM community_posts WHERE id = $1;`
	r := s.db.QueryRow(q, id)

	if err := r.Scan(&p.Id, &p.Title, &p.Content, &p.Reactions); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExist
		}
		return nil, err
	}

	return &p, nil
}
