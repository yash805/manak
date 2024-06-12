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
	for index, value := range numbers {
		fmt.Printf("index of Number is %d, and value is %d\n", value)
	}
}
