package main

import (
	"math/rand"
	"sync"
)

type SafeMap struct {
	mu sync.RWMutex
	m  map[int]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{m: make(map[int]int)}
}

func (s *SafeMap) Get(key int) (int, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value, ok := s.m[key]
	return value, ok
}

func (s *SafeMap) Set(key int, value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[key] = value
}

func (s *SafeMap) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.m)
}

func main() {
	sm := NewSafeMap()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sm.Set(rand.Intn(100), rand.Intn(100))
		}()
	}
	wg.Wait()
}
