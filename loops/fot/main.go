package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	counter := 0
	for {
		fmt.Println("infinite loop")
		counter++
		if counter == 3 {
			break
		}
	}

	numbers := []int{1, 2, 3, 4, 5}
	for _, value := range numbers {
		fmt.Printf("value is %d\n", value)
	}
}
