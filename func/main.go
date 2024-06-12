package main

import "fmt"

func divide(a, b float64) (float64, string) {
	if b == 0 {

		return 0, "denominator must not be zero"
	}
	return a / b, "nil"
}

func main() {
	fmt.Println("hello")
	ans, err := divide(10, 0)
	fmt.Println("Division of number is ", ans)
	fmt.Println("Division of number is ", err)
}
