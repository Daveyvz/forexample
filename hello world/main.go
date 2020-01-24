package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPU", runtime.NumCPU())
	fmt.Println("GoR", runtime.NumGoroutine())
	var counter int64

	var wg sync.WaitGroup
	//var mu sync.Mutex

	const gs = 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			//mu.Lock()
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			runtime.Gosched() // allows the cpu to run something else
			//mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("GoR", runtime.NumGoroutine())

	fmt.Println(counter)
}
