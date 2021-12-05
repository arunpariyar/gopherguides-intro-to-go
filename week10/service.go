package week10

import "sync"

type service struct {
	subs    map[string]catagories
	srcs    []string
	sub_chl map[string]chan news
	src_chl map[string]chan story
	history []news
	sync.Mutex
}

func NewService() *service {
	s := &service{
		subs:    make(map[string]catagories), //subscriber name and news catagories
		srcs:    make([]string, 0),           //source name and catagories
		sub_chl: make(map[string]chan news),  // a channel to give to every subscriber
		src_chl: make(map[string]chan story), //a channels to listen from every source for stories
		history: make([]news, 0),
	}
	return s
}

func (s *service) Start() {
	for _, ch := range s.src_chl {
		go s.listen(ch)
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

func (ns *service) Add(s Source) {
	//error checks must be added later
	ns.Lock()
	ns.srcs = append(ns.srcs, s.Name())
	ns.src_chl[s.Name()] = s.News()
	ns.Unlock()

}

func (ns *service) listen(ch chan story) {
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
				ch := ns.sub_chl[sub]
				ch <- n
			}
		}
	}
}
