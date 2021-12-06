package week11

import "context"

type Source interface {
	Name() string
	Publish(context.Context, story)
	News() chan story
}

//Think About auto generating news for mock source
