package week10

type Source interface {
	Name() string
	Publish(story)
	News() chan story
}
