// package main

// import "fmt"

// func max(slice []int) (int,error) {
// 	if len(slice) == 0 {
// 		return 0, errors.New("slice is empty")
// 	}
// 	max := slice[0]
// 	for _, value := range slice {
// 		if value > max {
// 			max = value
// 		}
// 	}
// 	return max, nil
// }

// func main() {
// 	numbers := []int{}
// 	max, err := max(numbers);
// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(max)
// 	}
// }