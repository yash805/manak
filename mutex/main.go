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
	runGo := 5
	wg.Add(runGo)
	for i := 0; i < runGo; i++ {
		go incCounter(i)
	}

	wg.Wait()
}

func incCounter(id int) {
	defer wg.Done()

	for j := 0; j < 5; j++ {
		mu.Lock()
		counter++
		fmt.Printf("run %d inc counter%d\n", id, counter)
		mu.Unlock()
		time.Sleep(time.Millisecond * 10)
	}
}
