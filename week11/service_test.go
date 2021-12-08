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

	//making a subscriber one
	ns.Subscribe("one", "ai")
	//making a subscriber two
	ns.Subscribe("two", "go")

	//Create some Mock Source
	mRCtx := context.Background() //background context for the new mock source
	m := NewMockSource("mock1")
	mCtx := m.Start(mRCtx) //starting m with the created context

	ns.Add(mCtx, m) //add the mock sources to the news service

	// Publish 10 Stories with mock news 1
	for i := 1; i <= 10; i++ {
		st := story{}
		st.body = "Mock 1 ai News " + fmt.Sprint(i)
		st.catagory = "ai"

		go m.Publish(mCtx, st)
	}

	//Create some Mock Source
	nRCtx := context.Background() //background context for the new mock source
	n := NewMockSource("mock1")
	nCtx := n.Start(nRCtx) //starting m with the created context

	ns.Add(nCtx, n) //add the mock sources to the news service

	// Publish 10 Stories with mock news 1
	for i := 1; i <= 10; i++ {
		st := story{}
		st.body = "Mock 2 Go News " + fmt.Sprint(i)
		st.catagory = "go"

		go n.Publish(mCtx, st)
	}

	//allowing some sleeping time to ensure all go routines get time to complete
	time.Sleep(5 * time.Millisecond)

	// fmt.Println(ns.history)

	h := ns.Search(1,2)

	fmt.Println(h)



	//stopping the news service.
	// m.Stop()

}
