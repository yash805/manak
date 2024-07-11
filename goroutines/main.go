package main

import (
	"fmt"
	"sync"
)

func printNumbers(wg *sync.WaitGroup, p string) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println(p, i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumbers(&wg, "method 1")
	go printNumbers(&wg, "method 2")

	wg.Wait()
}

// func work(i int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("work %d started\n", i
// 	fmt.Printf("work %d end\n", i)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	for i := 1; i <= 3; i++ {
// 		wg.Add(1)
// 		go work(i, &wg)
// 	}
// 	wg.Wait()

// 	fmt.Println("complete")
// }
