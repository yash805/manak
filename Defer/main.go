package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("john")
	data := add(2, 3)
	defer fmt.Println(data)
	defer fmt.Println("sam")
	fmt.Println("clark")

}
