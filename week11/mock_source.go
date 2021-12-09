package week11

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
	sync.RWMutex
}

func NewMockSource(s string) *MockSource {
	ms := &MockSource{
		name: s,
		news: make(chan story),
	}
	return ms
}

func (ms *MockSource) Start(ctx context.Context) context.Context {
	ms.Lock()
	ctx, ms.cancel = context.WithCancel(ctx)
	ms.Unlock()
	return ctx
}

func (ms *MockSource) Name() string {
	return ms.name
}

func (ms *MockSource) Publish(ctx context.Context, s story) {
	ms.RLock()
	if !ms.stopped {
		ms.news <- s
	}
	ms.RUnlock()

	<-ctx.Done()

	fmt.Println("Closing mock service channel")
}

func (ms *MockSource) News() chan story {
	return ms.news
}

func (ms *MockSource) Stop() {

	if ms.stopped {
		return
	}

	ms.Once.Do(func() {
		ms.Lock()
		defer ms.Unlock()
		ms.cancel()
		ms.stopped = true
		if ms.news != nil {

			close(ms.news)

		}
	})
}
