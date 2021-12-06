package week11

type Subscriber interface {
	Name() string
	Catagories() catagories
	Receive(chan news)
}
