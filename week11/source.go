package week11

import "context"

//Source interface allows to implement news sources that the news service will use
type Source interface {
	//Returns name of the source
	Name() string
	//Publish method requires the sources context an articel
	Publish(context.Context, Article)
	//Returns channel that news service can store and listen to
	News() chan Article
}
