package main

import (
	"fmt"
	"sync"
)

func main() {
	var count = 0
	var wg sync.WaitGroup
	wg.Add(1000)
	for i:=0 ; i<1000 ; i++{
		go func() {
			defer wg.Done()
			count++    //非原子操作
		}()
	}
	wg.Wait()
	fmt.Println(count) //结果分别为960 961 972 965 958
}
