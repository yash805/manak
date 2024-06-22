package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("learning web")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos")
	if err != nil {
		fmt.Println("Error getting", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("Type of response: %T\n", res)
	
	// Read  the response body
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading", err)
		return
	}
	fmt.Println(string(data))
}
