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
	m := NewMockSource("meta")

	//add the mock sources to the news service
	ns.Add(m)

	//go start must only be called when all the sources have been added otherwise it wont work
	ns.Start()
	
	//Publish 10 Stories
	for i := 0; i <= 10; i++ {
		st := &story{}
		st.body = "Go News " + fmt.Sprint(i)  
		st.catagory = "go"

		go m.Publish(*st)
	}

	ds.Receive(ds.nChl)
}
