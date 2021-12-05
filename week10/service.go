package week10

type service struct {
	subs    map[string]catagories
	srcs    []string
	sub_chl map[string]chan news
	src_chl map[string]chan story
	// err     chan error
	feed    []news
	history []news
}

func NewService() *service {
	s := &service{
		subs:    make(map[string]catagories), //subscriber name and news catagories
		srcs:    make([]string, 0),           //source name and catagories
		sub_chl: make(map[string]chan news),  // a channel to give to every subscriber
		src_chl: make(map[string]chan story), //a channels to listen from every source for stories
		feed:    make([]news, 0),
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
	ns.subs[s.Name()] = s.Catagories()
	ns.sub_chl[s.Name()] = make(chan news)
	return ns.sub_chl[s.Name()]
}

func (ns *service) Add(s Source) {
	//error checks must be added later
	ns.srcs = append(ns.srcs, s.Name())
	ns.src_chl[s.Name()] = s.News()

}

func (s *service) listen(ch chan story) {
	// convert story to news
	for st := range ch {
		func(st story) { //not running this as a go routine otherwise it won't get the ids right
			news := news{}
			news.id = len(s.history) + 1
			news.body = st.body
			news.catagory = st.catagory
			s.Publish(news)
		}(st)
	}
}

func (s *service) Publish(n news) {
	//save to history
	s.history = append(s.history, n)
	//send to the subscrber
	for sub, cs := range s.subs {
		for _, c := range cs {
			if string(n.catagory) == string(c) {
				ch := s.sub_chl[sub]
				ch <- n
			}
		}
	}
}
