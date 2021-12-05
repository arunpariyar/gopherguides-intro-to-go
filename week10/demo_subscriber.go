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

// func (s *subscriber) Start(ctx context.Context) context.Context{
// 	ctx, s.cancel = context.WithCancel(ctx)
// 	return ctx
// }

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
	// for {
	// 	select {
	// 	// case <-ctx.Done():
	// 	// 	fmt.Printf("%s is stopped", s.name)
	// 	// 	return
	// 	case news, ok := <-ch:
	// 		if !ok {
	// 			fmt.Printf("Channel closed")
	// 		}
	// 		fmt.Printf("News: %v \n", news)
	// 	}
	// }
}
