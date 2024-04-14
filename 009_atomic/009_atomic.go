package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)



func main()  {
	var counter uint64
	var wg sync.WaitGroup

	const numRoutines = 100
	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			atomic.AddUint64(&counter, 1)
		}()
	}
	wg.Wait()

	//在并发情况下，直接打印 counter 变量的值是不安全的，因为在打印的时候，其他的 goroutine 可能仍在修改 counter 的值，这样就会导致打印出来的值不准确。
	// fmt.Println("final counter value:", counter) 
	fmt.Println("final counter value:", atomic.LoadUint64(&counter))
}