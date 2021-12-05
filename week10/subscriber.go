package week10

type Subscriber interface {
	Name() string
	Catagories() catagories
	Receive(chan news)
}
