package week11

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

type Service struct {
	subs    map[string]catagories
	srcs    []string
	sub_chl map[string]chan News
	src_chl map[string]chan story
	history map[int]News
	Once    sync.Once
	stopped bool
	cancel  context.CancelFunc
	sync.RWMutex
}

func NewService() *Service {
	s := &Service{
		subs:    make(map[string]catagories), //subscriber name and news catagories
		srcs:    make([]string, 0),           //source name and catagories
		sub_chl: make(map[string]chan News),  // a channel to give to every subscriber THE WILL NOT BE REQUIRED AT ALL
		src_chl: make(map[string]chan story), //a channels to listen from every source for stories
		history: make(map[int]News),
	}
	return s
}

func (ns *Service) Start(ctx context.Context) {
	ctx, ns.cancel = context.WithCancel(ctx)

	ns.LoadArchive()

	for _, ch := range ns.src_chl {
		go ns.listen(ctx, ch)
	}
	// this will archive the history every 2 seconds
	go ns.Archive()

}

//trying to remove subscriber all together
func (ns *Service) Subscribe(n string, cs ...catagory) {
	cats := make([]catagory, 0)
	cats = append(cats, cs...)
	//error checks must be added later
	ns.Lock()
	ns.subs[n] = cats
	ns.sub_chl[n] = make(chan News)
	ns.Unlock()
	//as soon as the subsriber subscribe start publishing news as well // no need to return a channel start displaying news

	// This must be lauched as a goroutine or it will block
	go Listen(ns.sub_chl[n])
}

//To unsubscribe from the Service
func (ns *Service) UnSubscribe(s string) error {

	//return not found if not found
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
	return fmt.Errorf("%#v subscriber not found in subscription", s)
}

// This function will automaticaly start listening to the channel once subscribed.
func Listen(ch chan News) {
	for news := range ch {
		fmt.Println(news)
	}
}

func (ns *Service) Add(ctx context.Context, s Source) {
	//error checks must be added later
	ns.Lock()
	ns.srcs = append(ns.srcs, s.Name())
	ns.src_chl[s.Name()] = s.News()
	ns.Unlock()
	go ns.listen(ctx, s.News())
}

//remove a source
func (ns *Service) Remove(ctx context.Context, s Source) error {
	ok, index := Contains(ns.srcs, s.Name())
	if !ok {
		return fmt.Errorf("%v is not a Source", s.Name())
	}
	_, cancel := context.WithCancel(ctx)
	cancel()

	ns.Lock()
	ns.srcs = RemoveIndex(ns.srcs, index) //remove from ns.srcs
	delete(ns.src_chl, s.Name())          //remove from src.chl
	ns.Unlock()

	return nil
}

func (ns *Service) listen(ctx context.Context, ch chan story) {
	// convert story to news
	for st := range ch {
		func(st story) { //not running this as a go routine otherwise it won't get the ids right
			ns.Lock()
			news := News{}
			news.Id = len(ns.history) + 1
			news.Body = st.Body
			news.Catagory = st.Catagory
			ns.Unlock()

			ns.Publish(news)
		}(st)
	}

	<-ctx.Done()
	fmt.Println("Source Closing Down")
}

func (ns *Service) Publish(n News) {
	//save to history
	ns.Lock()
	ns.history[n.Id] = n
	ns.Unlock()

	ns.RLock()
	defer ns.RUnlock()
	//send to the subscrber
	for sub, cs := range ns.subs {
		for _, c := range cs {
			if string(n.Catagory) == string(c) {
				ch := ns.sub_chl[sub]
				ch <- n
			}
		}
	}

}

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

//there must be an error case as well.
func (ns *Service) Search(ids ...int) ([]News, error) {
	if len(ids) == 0 {
		return nil, fmt.Errorf("no ID's entered")
	}
	results := make([]News, 0)

	for _, id := range ids {
		ns.Lock()
		news, ok := ns.history[id]
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
	bb, err := json.Marshal(ns.history)
	ns.Unlock()

	if err != nil {
		return err
	}
	ns.Lock()
	ioutil.WriteFile("./serviceBackup.json", bb, 0644)
	ns.Unlock()
	return nil
}

func (ns *Service) BackupTo(s string) error {
	ns.Lock()
	bb, err := json.Marshal(ns.history)
	ns.Unlock()

	if err != nil {
		return err
	}
	ns.Lock()
	ioutil.WriteFile(s, bb, 0644)
	ns.Unlock()
	return nil
}

func (ns *Service) Archive() error {
	for {
		time.Sleep(4 * time.Millisecond)
		ns.Backup()
	}
}

func (ns *Service) LoadArchive() error {
	bb, err := ioutil.ReadFile("./serviceBackup.json")
	if err != nil {
		return err
	}

	backup := make(map[int]News)

	err = json.Unmarshal(bb, &backup)

	if err != nil {
		return err
	}

	for key, news := range backup {
		ns.history[key] = news
	}
	return nil
}

func (ns *Service) Clear() {
	clear := make(map[int]News)
	ns.Lock()
	ns.history = clear
	ns.Unlock()

	ns.Backup()
}

func Contains(srcs []string, src string) (bool, int) {
	for i, v := range srcs {
		if v == src {
			return true, i
		}
	}
	return false, 0
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
