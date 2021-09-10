package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu sync.Mutex
	count uint64
}

// Incr 对计数值加一
func (c *Counter) Incr(){
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

// Count 获取当前计数值
func (c *Counter) Count() uint64{
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
func Worker(c *Counter,wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(time.Second)
	c.Incr()
}
func main() {
	var counter Counter
	var wg sync.WaitGroup //声明waiigroup变量初始值为0
	wg.Add(10)	  //需要编排10个goroutine
	for i:=0;i<10;i++{
		go Worker(&counter,&wg)
	}
	wg.Wait()             //调用wait阻塞等待
	fmt.Println(counter.Count())
}
