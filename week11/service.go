package week11

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

//Service is the core of the news application
type Service struct {
	subs    map[string]Categories
	srcs    []string
	sub_chl map[string]chan News
	src_chl map[string]chan Article
	History map[int]News
	Once    sync.Once
	stopped bool
	cancel  context.CancelFunc
	sync.RWMutex
}

//NewService is a constructor function that allows creation of a service
func NewService() *Service {
	s := &Service{
		subs:    make(map[string]Categories),   //subscriber name and news catagories
		srcs:    make([]string, 0),             //source name and catagories
		sub_chl: make(map[string]chan News),    // a channel to give to every subscriber THE WILL NOT BE REQUIRED AT ALL
		src_chl: make(map[string]chan Article), //a channels to listen from every source for stories
		History: make(map[int]News),
	}
	return s
}

//Start Methods allows a service to be started
//A Background context must be provides as an argument
func (ns *Service) Start(ctx context.Context) {
	ctx, ns.cancel = context.WithCancel(ctx)
	ns.LoadArchive()

	for _, ch := range ns.src_chl {
		go ns.SrcListener(ctx, ch)
	}

	go ns.Archive()

}

//Subscribe allows creation of a new subscriber for the news service
func (ns *Service) Subscribe(n string, cs ...Category) {
	cats := make([]Category, 0)
	cats = append(cats, cs...)
	//error checks must be added later
	ns.Lock()
	ns.subs[n] = cats
	ns.sub_chl[n] = make(chan News)
	ns.Unlock()

	go SubListener(ns.sub_chl[n])
}

//Unsubscribe allows an existing subscriber to opt out from the service
func (ns *Service) Unsubscribe(s string) error {
	for name, _ := range ns.subs {
		if s == name {
			ns.Lock()
			delete(ns.subs, s)
			delete(ns.sub_chl, s)
			ns.Unlock()
			fmt.Printf("%s has unsubscribed \n", s)
			return nil
		}
	}
	return fmt.Errorf("%v has not subscribed", s)
}

//SubListener keeps watch on subscriber channel
//SubListener is to be used as go routine
func SubListener(ch chan News) {
	for news := range ch {
		fmt.Println(news)
	}
}

//Add allows new sources to be added to the service
func (ns *Service) Add(ctx context.Context, s Source) {
	//error checks must be added later
	ns.Lock()
	ns.srcs = append(ns.srcs, s.Name())
	ns.src_chl[s.Name()] = s.News()
	ns.Unlock()
	go ns.SrcListener(ctx, s.News())
}

//Remove allows existing sources to be removed from the service
func (ns *Service) Remove(ctx context.Context, s Source) error {
	ok, index := Contains(ns.srcs, s.Name())
	if !ok {
		return fmt.Errorf("%v not found", s.Name())
	}
	_, cancel := context.WithCancel(ctx)
	cancel()

	ns.Lock()
	ns.srcs = RemoveIndex(ns.srcs, index) //remove from ns.srcs
	delete(ns.src_chl, s.Name())          //remove from src.chl
	ns.Unlock()

	return nil
}

//SrcListener listens for articles coming from sources
func (ns *Service) SrcListener(ctx context.Context, ch chan Article) {
	// convert story to news
	for st := range ch {
		func(st Article) { //not running this as a go routine otherwise it won't get the ids right
			ns.Lock()
			news := News{}
			news.ID = len(ns.History) + 1
			news.Body = st.Body
			news.Category = st.Category
			ns.Unlock()

			ns.Publish(news)
		}(st)
	}

	<-ctx.Done()
	// fmt.Println("Source Closing Down")
}

//Publish distributes news to the relevent subscribers
func (ns *Service) Publish(n News) {
	//save to history
	ns.Lock()
	ns.History[n.ID] = n
	ns.Unlock()

	ns.RLock()
	defer ns.RUnlock()
	//send to the subscrber
	for sub, cs := range ns.subs {
		for _, c := range cs {
			if string(n.Category) == string(c) {
				ch := ns.sub_chl[sub]
				ch <- n
			}
		}
	}

}

//Stop closes the operation of a Service
func (ns *Service) Stop() {
	ns.RLock()
	if ns.stopped {
		ns.RLock()
		return
	}
	ns.RUnlock()

	ns.Once.Do(func() {
		ns.Backup()
		ns.Lock()
		defer ns.Unlock()

		ns.cancel()
		ns.stopped = true

		//closing all subscribers channels
		for _, ch := range ns.sub_chl {
			if ch != nil {
				close(ch)
			}
		}
	})

}

//Search allows to read from the services' history
//Search takes in ids of int as parameters
func (ns *Service) Search(ids ...int) ([]News, error) {
	if len(ids) == 0 {
		return nil, fmt.Errorf("no ID's entered")
	}
	results := make([]News, 0)

	for _, id := range ids {
		ns.Lock()
		news, ok := ns.History[id]
		ns.Unlock()

		if !ok {
			return nil, fmt.Errorf("couldnt find news with ID: %d", id)
		}
		ns.Lock()
		results = append(results, news)
		ns.Unlock()
	}

	return results, nil
}

func (ns *Service) Backup() error {
	ns.Lock()
	bb, err := json.Marshal(ns.History)
	ns.Unlock()

	if err != nil {
		return err
	}
	ns.Lock()
	ioutil.WriteFile("./tmp/news_service.json", bb, 0644)
	ns.Unlock()
	return nil
}

//BackupTo allows the user to provide a desired name for creating a backup
func (ns *Service) BackupTo(s string) error {
	ns.Lock()
	bb, err := json.Marshal(ns.History)
	ns.Unlock()

	if err != nil {
		return err
	}
	ns.Lock()
	ioutil.WriteFile(s, bb, 0644)
	ns.Unlock()
	return nil
}

//Archive performs a reoccuring backup of the services' history
func (ns *Service) Archive() error {
	for {
		time.Sleep(4 * time.Millisecond)
		ns.Backup()
	}
}

//LoadArchive allows past backup to be loaded to the services' history
func (ns *Service) LoadArchive() error {
	bb, err := ioutil.ReadFile("./tmp/news_service.json")
	if err != nil {
		return err
	}

	backup := make(map[int]News)

	err = json.Unmarshal(bb, &backup)

	if err != nil {
		return err
	}

	for key, news := range backup {
		ns.History[key] = news
	}
	return nil
}

func (ns *Service) Clear() {
	clear := make(map[int]News)
	ns.Lock()
	ns.History = clear
	ns.Unlock()

	ns.Backup()
}

//Contains is a helper function that allows identifying if a element exists in a Slice
func Contains(srcs []string, src string) (bool, int) {
	for i, v := range srcs {
		if v == src {
			return true, i
		}
	}
	return false, 0
}

//Remove Index is a helper function that allows removal of an element from a slice
func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
