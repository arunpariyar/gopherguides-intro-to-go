package week09

import "fmt"

// news service will build this
type news struct {
	id        int
	story     story
	publisher string
	cat       catagory
}

func (n news) String() string {
	return fmt.Sprintf("id:%d\nstory:%spublisher:%s\ncatagory:%s\n", n.id, n.story, n.publisher, n.cat)
}
