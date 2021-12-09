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
	// mRCtx := context.Background() //background context for the new mock source
	// m := NewMockSource("mock1")
	// mCtx := m.Start(mRCtx) //starting m with the created context

	// ns.Add(mCtx, m) //add the mock sources to the news service

	// // Publish 10 Stories with mock news 1
	// for i := 1; i <= 10; i++ {
	// 	st := story{}
	// 	st.body = "Mock 1 ai News " + fmt.Sprint(i)
	// 	st.catagory = "ai"

	// 	go m.Publish(mCtx, st)
	// }

	//Unsubscribing two
	// err := ns.UnSubscribe("two")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	//Create some Mock Source
	nRCtx := context.Background() //background context for the new mock source
	n := NewFileSource("stories")
	nCtx := n.Start(nRCtx) //starting m with the created context

	ns.Add(nCtx, n) //add the mock sources to the news service
	// Publish 10 Stories with mock news 1
	stories, err := n.LoadFile()
	if err != nil {
		fmt.Println(err)
	}
	for _, st := range stories {
		go n.Publish(nRCtx, st)
	}

	n.Stop()

	// // err := ns.Remove(nCtx, n)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	fmt.Println(ns.srcs, ns.src_chl)

	//allowing some sleeping time to ensure all go routines get time to complete
	time.Sleep(5 * time.Millisecond)

	// fmt.Println(ns.history)
	// res, err := ns.Search(1,2,3,4,5, 500)

	// fmt.Println(res, err)

	//stopping the news service.
	// ns.Stop()
	//clearing the history
	// ns.Clear()

	// fmt.Println(ns.history)

}
