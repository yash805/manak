package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 40
	fmt.Println(num)
	fmt.Printf("Type of num is %T\n", num)

	var data float64 = float64(num)
	fmt.Printf("Type of num is %T", data)

	num = 123
	str := strconv.Itoa(num)
	fmt.Println(str)
	fmt.Printf("Type of str is %T", str)

	number := "1234"
	number_int, _ := strconv.Atoi(number)
	fmt.Println(number_int)
	fmt.Printf("Type of number_int is %T", number_int)

	num_string := "3.14"
	number_float, _ := strconv.ParseFloat(num_string, 64)
	fmt.Println(number_float)
	fmt.Printf("Type of number_float is %T", number_float)
}
