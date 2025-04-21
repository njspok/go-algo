package singleflight

import "sync"

func New() *SingleFlight {
	return &SingleFlight{
		queue: make(map[string]*single),
	}
}

type SingleFlight struct {
	queue map[string]*single
	mu    sync.RWMutex
}

func (sf *SingleFlight) Do(key string, f func() int) int {
	sf.mu.RLock()
	s, ok := sf.queue[key]
	if ok {
		sf.mu.RUnlock()
		return s.do(f)
	}
	sf.mu.RUnlock()

	sf.mu.Lock()
	if s, ok := sf.queue[key]; ok {
		sf.mu.Unlock()
		return s.do(f)
	}

	s = &single{}
	sf.queue[key] = s
	sf.mu.Unlock()

	res := s.do(f)

	sf.mu.Lock()
	delete(sf.queue, key)
	sf.mu.Unlock()

	return res
}

type single struct {
	once   sync.Once
	result int
}

func (s *single) do(f func() int) int {
	s.once.Do(func() {
		s.result = f()
	})
	return s.result
}
