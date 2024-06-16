package main

import "fmt"

func modifyValue(a *int) {
	*a = *a * 2

}

func main() {
	// var num int
	num := 2
	// var ptr *int
	// ptr = &num
	ptr := &num
	// fmt.Println(num)
	fmt.Println(ptr)
	fmt.Println(*ptr)
	var pointer *int
	if pointer == nil {
		fmt.Println("not assigned")
	}

	value := 10
	modifyValue(&value)
	fmt.Println((value))

}
