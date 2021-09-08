package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex
	Count int64
}
func main() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)
	for i:=0;i<10;i++{
		go func() {
			defer wg.Done()
			counter.mu.Lock()
			counter.Count++
			counter.mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count)
}
