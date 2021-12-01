package week09

import (
	"sort"
	"sync"
)

type NewsService struct {
	// cancel     context.CancelFunc //for cancellation when required
	// Subs       subscribers        //this will keep track of all subscribers
	// Src        sources            //this keep track of all sources
	// errs       chan error         //any error that might come about
	// waiting    chan stories       //stories waiting to be turned to news
	// feed       chan news          //news that ready to be distributed //map[catagory][]chan news
	History    map[int]news //collection of all past news
	Stopped    bool         //state of News Service
	sync.Mutex              // to lock race condtions that might occur
	StopOnce   sync.Once    // to close news service
}

//The importable package should be able to create a news service, manage sources and subscribers, and save and load the state of the news service

//function start that starts newss service
// func (ns *NewsService) Start(ctx context.Context, src sources) (context.Context, error) {

// 	if len(src) <= 0 {
// 		return nil, ErrSourcesEmpty(len(src))
// 	}
// 	ctx, cancel := context.WithCancel(ctx)
// 	ns.cancel = cancel

// 	go func(ctx context.Context) {
// 		<-ctx.Done()
// 		cancel()
// 		ns.Stop() // need to create a stop function
// 	}(ctx)

// }

//Search saves all the given ids into a slice of keys, sorts it and then checks in the news history if any news with the given key is found if yes than it will push it to a history of news and return the history
func (ns *NewsService) Search(ids ...int) []news {
	history := make([]news, 0)
	ks := make([]int, 0)

	ks = append(ks, ids...)

	sort.Ints(ks)

	for i := 0; i <= len(ks)-1; i++ {
		v, ok := ns.History[ks[i]]
		if ok {
			history = append(history, v)
		}

	}
	return history
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
func (ns *NewsService) Stop() {
	ns.Lock()
	ns.Stopped = true

	if ns.Stopped {
		ns.Unlock()
		return
	}

	ns.Unlock()

	// ns.StopOnce.Do(func() {
	// 	if ns.waiting != nil {
	// 		close(ns.waiting)
	// 	}

	// 	if ns.errs != nil {
	// 		close(ns.errs)
	// 	}

	// 	if ns.feed != nil {
	// 		close(ns.errs)
	// 	}
	// })
}
