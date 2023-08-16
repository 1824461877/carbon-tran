package selenium

import (
	"sync"
	"time"
)

type TaskWatch struct {
	mu    sync.Mutex
	cache *Cache
	task  map[string]func(*Cache)
	Time  time.Duration
}

func (tw *TaskWatch) Add(key string, fs func(*Cache)) {
	tw.mu.Lock()
	defer tw.mu.Unlock()
	tw.task[key] = fs
}

func (tw *TaskWatch) Del(key string) {
	tw.mu.Lock()
	defer tw.mu.Lock()
	delete(tw.task, key)
}

func (tw *TaskWatch) Watch() {
	go func() {
		times := time.NewTimer(2 * time.Second)
		for {
			select {
			case <-times.C:
				for _, v := range tw.task {
					v(tw.cache)
				}
				times.Reset(tw.Time)
			}
		}
	}()
}

func NewTaskWatch(times int) *TaskWatch {
	return &TaskWatch{
		task:  make(map[string]func(*Cache), 10),
		Time:  time.Duration(times) * time.Second,
		cache: NewCache(),
	}
}
