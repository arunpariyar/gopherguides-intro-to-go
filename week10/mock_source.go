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
	sync.RWMutex
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
			ms.Lock()
			close(ms.news)
			ms.Unlock()
		}
	})
}
