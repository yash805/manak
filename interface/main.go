package main

import "fmt"

type values struct {
	first  int
	second int
}

type mathTest interface {
	add(a, b int) int
	multiply(a, b int) int
}

func (v values) add(a, b int) int {
	return a + b + v.first + v.second
}

func (v values) multiply(a, b int) int {
	return a + b + v.first + v.second
}

func main() {
	var te mathTest = values{1, 2}
	fmt.Println(te.add(1, 2))
	fmt.Println(te.multiply(1, 2))

	var test interface{}
	test = "some str"
	val, ok := test.(string)
	if ok {
		fmt.Println(val)
	}

}
