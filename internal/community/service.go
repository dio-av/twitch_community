package community

import "errors"

func NewPost(t, c string) *Post {
	return &Post{
		Title:     t,
		Content:   c,
		Reactions: make(map[string]int),
	}
}

func (p *Post) Edit(c string) error {
	p.Content = c
	return nil
}

func (p *Post) AddReaction(rc ...string) error {
	if len(rc) < 1 {
		return errors.New("trying do add empty reaction")
	}
	for _, v := range rc {
		p.Reactions[v]++
	}
	return nil
}

func (p *Post) RemoveReaction(rc ...string) error {
	if len(rc) < 1 {
		return errors.New("trying do remove empty reaction")
	}
	for _, v := range rc {
		if p.Reactions[v]--; p.Reactions[v] < 0 {
			p.Reactions[v] = 0
		}
		p.Reactions[v]--
	}

	return nil
}
