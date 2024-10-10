package community

import "database/sql"

type Post struct {
	Id        int
	Title     string
	Content   string
	Reactions map[string]int
}

type Posts []*Post

type Service struct {
	repo Repository
	db   *sql.DB
}

// UseCase interface
type UseCase interface {
	Create(content string) error
	Delete(id int) error
	Edit(id int) error
}
