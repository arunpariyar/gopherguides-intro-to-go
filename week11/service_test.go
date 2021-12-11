package week11

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"
	"time"
)

func Test_Service_Subscribe_Start(t *testing.T) {
	t.Parallel()
	s := NewService()
	s.Start(context.Background())

	exp := false

	if s.stopped != false {
		t.Errorf("expected %v got %v", exp, s.stopped)
	}
}

func Test_Service_Unit(t *testing.T) {

	//create a new news serive
	ns := NewService()
	defer ns.Stop()
	//create a background context for ns
	nsBCtx := context.Background()

	//go start must only be called when all the sources have been added otherwise it wont work
	ns.Start(nsBCtx)

	//making a subscriber one
	ns.Subscribe("one", "ai")
	//making a subscriber two
	ns.Subscribe("two", "go")

	// Create some Mock Source
	mRCtx := context.Background() //background context for the new mock source
	m := NewMockSource("mock1")
	defer m.Stop()
	//starting m with the created context
	mCtx := m.Start(mRCtx)

	//add the mock sources to the news service
	ns.Add(mCtx, m)

	// Publish 10 Stories with mock news 1
	for i := 1; i <= 10; i++ {
		st := Article{}
		st.Body = "Mock 1 ai News " + fmt.Sprint(i)
		st.Category = "ai"

		go m.Publish(mCtx, st)
	}

	// Unsubscribing two
	err := ns.Unsubscribe("two")
	if err != nil {
		t.Error(err)
	}

	//Create some Mock Source
	nRCtx := context.Background() //background context for the new mock source
	nfs := NewFileSource("stories")
	defer nfs.Stop()
	nCtx := nfs.Start(nRCtx) //starting m with the created context

	ns.Add(nCtx, nfs) //add the mock sources to the news service
	// Publish 10 Stories with mock news 1

	nfs.PublishStories()

	err = ns.Remove(nCtx, nfs)
	if err != nil {
		t.Error(err)
	}

	//allowing some sleeping time to ensure all go routines get time to complete
	time.Sleep(5 * time.Millisecond)

	_, err = ns.Search(1, 2, 3, 4, 5)
	if err != nil {
		t.Error(err)
	}

	//clearing the history
	ns.Clear()
	//stopping the news service.
	ns.Stop()
}

func Test_Service_Remove_Fail(t *testing.T) {
	ns := NewService()
	s := NewMockSource("test")

	exp := "test not found"

	act := ns.Remove(context.Background(), s)

	if act.Error() != exp {
		t.Fatalf("expected %v got %v", exp, act)
	}
}

func Test_Service_Unsubscribe_Fail(t *testing.T) {
	ns := NewService()
	act := ns.Unsubscribe("test")
	exp := "test has not subscribed"
	if act.Error() != exp {
		t.Fatalf("expected %v got %v", exp, act)
	}
}

func Test_Service_Search_Fail(t *testing.T) {
	ns := NewService()
	_, act := ns.Search()
	exp := "no ID's entered"

	if act.Error() != exp {
		t.Fatalf("expected %v got %v", exp, act.Error())
	}
}

func Test_Service_Search_WrongID(t *testing.T) {
	ns := NewService()
	_, act := ns.Search(100)
	exp := "couldnt find news with ID: 100"
	if act.Error() != exp {
		t.Fatalf("expected %v got %v", exp, act.Error())
	}
}

func Test_Service_BackUpTo(t *testing.T) {
	ns := NewService()
	f := "./testignore.json"
	ns.BackupTo(f)
	_, err := ioutil.ReadFile("./testignore.json")
	if err != nil {
		t.Fatalf("%v not found", f)
	}

}
