package main

import "fmt"

type stu struct {
	Name   string
	Grade1 float64
	Grade2 float64
	Grade3 float64
}

func (s stu) GPA() float64 {
	return (s.Grade1 + s.Grade2 + s.Grade3) / 3
}

func main() {
	student := stu{Name: "sam", Grade1: 90, Grade2: 80, Grade3: 70}
	fmt.Printf("%s %.2f\n", student.Name, student.GPA())
}

// func long(s string) int {
// 	m := make(map[byte]int)
// 	right, left, maxLength := 0, 0, 0

// 	for right < len(s) {
// 		if _, ok := m[s[right]]; ok {
// 			left = max(left, m[s[right]]+1)
// 		}
// 		m[s[right]] = right
// 		maxLength = max(maxLength, right-left+1)
// 		right++
// 	}
// 	return maxLength
// }

// func max(x, y int) int {
// 	if x > y {
// 		return x
// 	}
// 	return y
// }

// func main() {
// 	s := "abcabcbb"
// 	fmt.Println(long(s))
// }
