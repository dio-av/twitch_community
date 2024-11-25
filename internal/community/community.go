package community

type Post struct {
	Id        int
	Title     string
	Content   string
	Reactions map[string]int
}

type Posts []*Post

// UseCase interface
type UseCase interface {
	Create(content string) error
	Delete(id int) error
	Edit(id int) error
}
