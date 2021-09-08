package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	Mu sync.Mutex
	Count int
}
func main() {
	var (
		counter Counter
		wg sync.WaitGroup
	)
	wg.Add(10)
	for i := 0;i < 10 ;i++ {
		go func() {
			defer wg.Done()
			for j:=0;j<100000000;j++{
				counter.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count)

}

func (c *Counter) Incr(){
	c.Mu.Lock()
	c.Count++
	c.Mu.Unlock()
}

func (c *Counter) Counts() int{
	c.Mu.Lock()
	defer c.Mu.Unlock()
	return c.Count
}
