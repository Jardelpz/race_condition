package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int
	mu      sync.Mutex
	wg      sync.WaitGroup
)

func main() {
	maxIterations := 100

	for i := 0; i < maxIterations; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(1 * time.Second)
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Printf("Counter %v\n", counter)
}
