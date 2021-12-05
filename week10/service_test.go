package week10

import (
	"fmt"
	"testing"
)

func Test_Service_Unit(t *testing.T) {

	//create a new news serive
	ns := NewService()

	//creating a new Demo Subscriber
	ds := NewDemoSubscriber("demo_subscriber", []catagory{"go"})

	//subscribe to the news service channel (returns channel to listen)
	ch := ns.Subscribe(ds)
	ds.nChl = ch

	// save it the the demo subsriber news channnel
	//Create some Mock Source
	m := NewMockSource("mock1")
	n := NewMockSource("mock2")

	//add the mock sources to the news service
	ns.Add(m)
	ns.Add(n)

	//go start must only be called when all the sources have been added otherwise it wont work
	ns.Start()

	//Publish 10 Stories with mock news 1
	for i := 0; i <= 10; i++ {
		st := story{}
		st.body = "Mock 1 Go News " + fmt.Sprint(i)
		st.catagory = "go"

		go m.Publish(st)
	}
	//Publish 5 news with mock news 2
	for i := 0; i <= 5; i++ {
		st := story{}
		st.body = "Mock 2 Go News " + fmt.Sprint(i)
		st.catagory = "go"

		go n.Publish(st)
	}

	//demo subscriber receives the news and prints it to the terminal
	ds.Receive(ds.nChl)
}
