package week11

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func Test_Service_Unit(t *testing.T) {

	//create a new news serive
	ns := NewService()
	//create a background context for ns
	nsBCtx := context.Background()

	//go start must only be called when all the sources have been added otherwise it wont work
	ns.Start(nsBCtx)

	//creating a new Demo Subscriber
	ds := NewDemoSubscriber("demo_subscriber", []catagory{"go", "ai"})
	//start receiving
	

	//subscribe to the news service channel (returns channel to listen)
	ch := ns.Subscribe(ds)
	// save it the the demo subsriber news channnel
	ds.nChl = ch
	// demo subscriber receives the news and prints it to the terminal
	go ds.Receive(ds.nChl)
	//Create some Mock Source
	m := NewMockSource("mock1")

	//background context for the new mock source
	mRCtx := context.Background()

	//starting m with the created context
	mCtx := m.Start(mRCtx)
	
	

	

	//add the mock sources to the news service
	// ns.Add(m)
	ns.Add(mCtx, m)

	// Publish 10 Stories with mock news 1
	for i := 0; i <= 10; i++ {
		st := story{}
		st.body = "Mock 1 Go News " + fmt.Sprint(i)
		st.catagory = "ai"

		go m.Publish(mCtx, st)
	}



	

	//allowing some sleeping time to ensure all go routines get time to complete
	time.Sleep(5 * time.Millisecond)

	//stopping the news service.
	// ns.Stop()
}
