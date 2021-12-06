package week11

import (
	"context"
	"fmt"
	"sync"
)

type service struct {
	subs    map[string]catagories
	srcs    []string
	sub_chl map[string]chan news
	src_chl map[string]chan story
	history []news
	Once    sync.Once
	stopped bool
	cancel  context.CancelFunc
	sync.RWMutex
}

func NewService() *service {
	s := &service{
		subs:    make(map[string]catagories), //subscriber name and news catagories
		srcs:    make([]string, 0),           //source name and catagories
		sub_chl: make(map[string]chan news),  // a channel to give to every subscriber THE WILL NOT BE REQUIRED AT ALL
		src_chl: make(map[string]chan story), //a channels to listen from every source for stories
		history: make([]news, 0),
	}
	return s
}

func (ns *service) Start(ctx context.Context) {
	ctx, ns.cancel = context.WithCancel(ctx)

	for _, ch := range ns.src_chl {
		go ns.listen(ctx, ch)
	}
}

func (ns *service) Subscribe(s Subscriber) chan news {
	//error checks must be added later
	ns.Lock()
	ns.subs[s.Name()] = s.Catagories()
	ns.sub_chl[s.Name()] = make(chan news)
	ns.Unlock()
	return ns.sub_chl[s.Name()]
}

func (ns *service) Add(ctx context.Context, s Source) {
	//error checks must be added later
	ns.Lock()
	ns.srcs = append(ns.srcs, s.Name())
	ns.src_chl[s.Name()] = s.News() // dont save just launch it in a go routine.
	ns.Unlock()
	
	go ns.listen(ctx, s.News())
}

func (ns *service) listen(ctx context.Context, ch chan story) {
	// convert story to news
	for st := range ch {
		func(st story) { //not running this as a go routine otherwise it won't get the ids right
			ns.Lock()
			news := news{}
			news.id = len(ns.history) + 1
			news.body = st.body
			news.catagory = st.catagory
			ns.Unlock()

			ns.Publish(news)
		}(st)
	}
	<-ctx.Done()
	fmt.Println("News Service Closing Down")
}

func (ns *service) Publish(n news) {
	//save to history
	ns.Lock()
	ns.history = append(ns.history, n)
	ns.Unlock()
	//send to the subscrber
	for sub, cs := range ns.subs {
		for _, c := range cs {
			if string(n.catagory) == string(c) {
				ns.Lock()
				ch := ns.sub_chl[sub]
				ch <- n
				ns.Unlock()
			}
		}
	}
}

func (ns *service) Stop() {
	ns.cancel()
	if ns.stopped {
		return
	}

	ns.Lock()
	ns.stopped = true
	ns.Unlock()

	ns.Once.Do(func() {
		//closing all source channels
		for _, ch := range ns.src_chl {
			if ch != nil {
				ns.Lock()
				close(ch)
				ns.Unlock()
			}
		}

		//closing all subscribers channels
		for _, ch := range ns.sub_chl {
			if ch != nil {
				ns.Lock()
				close(ch)
				ns.Unlock()
			}
		}
	})

}
