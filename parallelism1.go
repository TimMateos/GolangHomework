package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, mu *sync.Mutex) {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}(wg, mu)
	}
	wg.Wait()
	fmt.Println(counter)
}
