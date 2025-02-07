package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var incrementer int64

	gs := 100
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			// uso de atomic para pequenos incrementos vai realizar um lock em incrementer
			// dispensando o uso de um Mutex (exemplo anterior)
			atomic.AddInt64(&incrementer, 1)
			r := atomic.LoadInt64(&incrementer)
			fmt.Println(r)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", incrementer)
}
