package week10

import "context"

type Source interface {
	Name() string
	Publish(context.Context, story)
	News() chan story
}
