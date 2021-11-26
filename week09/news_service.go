package week09

import (
	"context"
	"sync"
)

type NewsService struct {
	cancel  context.CancelFunc
	Subs    subscribers //this will keep track of all subscribers
	Src     sources     //this keep track of all sources
	errs 	chan error
	waiting chan stories
	feed    chan news //map[catagory][]chan news
	history []news
	stopped bool
	sync.Mutex
	StopOnce sync.Once

}

//The importable package should be able to create a news service, manage sources and subscribers, and save and load the state of the news service

//function start that starts newss service
func (ns *NewsService) Start(ctx context.Context, src sources)(context.Context, error){
	if len(src) <= 0 {
		return nil, ErrSourcesEmpty(len(src))
	}
	ctx, cancel := context.WithCancel(ctx)
	ns.cancel = cancel 

	go func(ctx context.Context) {
		<-ctx.Done()
		cancel()
		ns.Stop() // need to create a stop function
	}(ctx)
} 

//function to create news 
	//listen to the waiting channels for stories 
	//covert them to news
	// save them in history
	// and the send them in the feed

//function to add and remove source

//function to add and remove subscribers

//function to save the state of news service

//function to load the state of news service

//Method to stop news Service
func (ns *NewsService) Stop(){
	ns.cancel()

	ns.Lock()
	if ns.stopped {
		ns.Unlock()
		return
	}
	ns.Unlock()
	ns.StopOnce.Do(func(){
		if ns.waiting != nil {
			close(ns.waiting)
		}

		if ns.errs != nil {
			close(ns.errs)
		}

		if ns.feed != nil{
			close(ns.errs)
		}
	})
}
