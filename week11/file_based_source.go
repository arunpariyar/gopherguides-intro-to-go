package week11

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sync"
)

type FileSource struct {
	name    string
	news    chan Article
	cancel  context.CancelFunc
	stopped bool
	Once    sync.Once
	sync.RWMutex
}

func NewFileSource(s string) *FileSource {
	nfs := &FileSource{
		name: s,
		news: make(chan Article),
	}
	return nfs
}

func (nfs *FileSource) Name() string {
	return nfs.name
}

func (nfs *FileSource) Start(ctx context.Context) context.Context {
	nfs.Lock()
	ctx, nfs.cancel = context.WithCancel(ctx)
	nfs.Unlock()
	return ctx
}

func (nfs *FileSource) Publish(ctx context.Context, s Article) {
	nfs.RLock()
	if !nfs.stopped {
		nfs.news <- s
	}
	nfs.RUnlock()
	<-ctx.Done()
	fmt.Println("Closing file service channel")
}

func (nfs *FileSource) PublishStories() error {
	stories, err := nfs.LoadFile()
	if err != nil {
		return err
	}
	if !nfs.stopped {
		for _, story := range stories {
			nfs.RLock()
			nfs.news <- story
			nfs.RUnlock()
		}
	}
	return nil

}

func (nfs *FileSource) News() chan Article {
	return nfs.news
}

func (nfs *FileSource) Stop() {

	if nfs.stopped {
		return
	}

	nfs.Once.Do(func() {
		nfs.Lock()
		defer nfs.Unlock()
		nfs.cancel()
		nfs.stopped = true
		if nfs.news != nil {

			close(nfs.news)

		}
	})
}

func (ns *FileSource) LoadFile() (Articles, error) {
	bb, err := ioutil.ReadFile("./stories/stories.json")
	if err != nil {
		return nil, err
	}

	articles := make([]Article, 0)

	err = json.Unmarshal(bb, &articles)

	if err != nil {
		return nil, err
	}
	return articles, nil
}
