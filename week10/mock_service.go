package week10

type MockSource struct {
	name string
	news chan story
}

func NewMockSource(s string) MockSource {
	ms := MockSource{
		name: s,
		news: make(chan story),
	}
	return ms
}

func (ms MockSource) Name() string {
	return ms.name
}

func (ms MockSource) Publish(s story) {
	ms.news <- s
}

func (ms MockSource) News() chan story {
	return ms.news
}
