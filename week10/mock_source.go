package week10

import (
	"context"
	"fmt"
	"sync"
)

type MockSource struct {
	name    string
	news    chan story
	cancel  context.CancelFunc
	stopped bool
	Once    sync.Once
	sync.Mutex
}

func NewMockSource(s string) *MockSource {
	ms := &MockSource{
		name: s,
		news: make(chan story),
		// stopped: false,
	}
	return ms
}

func (ms *MockSource) Start(ctx context.Context) context.Context {
	ctx, ms.cancel = context.WithCancel(ctx)
	return ctx
}

func (ms *MockSource) Name() string {
	return ms.name
}

func (ms *MockSource) Publish(ctx context.Context, s story) {
	if !ms.stopped {
		ms.news <- s
	}

	<-ctx.Done()

	fmt.Println("Closing mock service news channel")
}

func (ms *MockSource) News() chan story {
	return ms.news
}

func (ms *MockSource) Stop() {
	ms.cancel()
	if ms.stopped {
		return
	}
	ms.Lock()
	ms.stopped = true
	ms.Unlock()

	ms.Once.Do(func() {
		if ms.news != nil {
			close(ms.news)
		}
	})
}
