package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,banana,orange"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str := "one two three four two three five two"
	count := strings.Count(str, "two")
	fmt.Println(count)

	str1 := "   hello              world    "
	trimmed := strings.TrimSpace(str1)
	fmt.Println(trimmed)

	st1 := "yash"
	st2 := "saini"

	result := strings.Join([]string{st1, "kumar", st2}, " ")
	println(result)

}

// contains
