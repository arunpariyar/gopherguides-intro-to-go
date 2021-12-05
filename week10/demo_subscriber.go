package week10

import "fmt"

type subscriber struct {
	name       string
	catagories []catagory
	nChl       chan news
}

func NewDemoSubscriber(n string, cs catagories) *subscriber {
	ds := &subscriber{}
	ds.name = n
	ds.catagories = cs
	ds.nChl = make(chan news)
	return ds
}

func (s subscriber) Name() string {
	return s.name
}

func (s subscriber) Catagories() catagories {
	return s.catagories
}

func (s subscriber) Receive(ch chan news) {
	for news := range ch {
		fmt.Println(news)
	}
}
