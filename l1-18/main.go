package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	value int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func (c *Counter) Add(n int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value += n
}

func worker(c *Counter, iterations int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < iterations; i++ {
		c.Increment()
	}
}

func main() {
	var (
		counter Counter
		wg      sync.WaitGroup
	)

	workers := []struct {
		gorutines  int
		increments int
	}{
		{3, 100},
		{5, 30},
		{2, 60},
	}

	totalExpected := 0
	for _, w := range workers {
		totalExpected += w.gorutines * w.increments
		for i := 0; i < w.gorutines; i++ {
			wg.Add(1)
			go worker(&counter, w.increments, &wg)
		}
	}
	wg.Wait()

	fmt.Printf("Expected: %d\n", totalExpected)
	fmt.Printf("Actual: %d\n", counter.Value())
	fmt.Printf("Match: %t\n", totalExpected == counter.Value())
}
