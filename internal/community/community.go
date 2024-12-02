package community

type Post struct {
	Id        int            `db:"id"`
	Title     string         `db:"title"`
	Content   string         `db:"content"`
	Reactions map[string]int `db:"reactions"`
}

// type Posts []*Post

// UseCase interface
type UseCase interface {
	Create(content string) error
	Delete(id int) error
	Edit(id int) error
}
