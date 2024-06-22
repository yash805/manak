package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("hello")
	myURL := "https://jsonplaceholder.typicode.com/todos"
	// myURL := "yash"

	parsedURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Error parsing")
		return
	}
	fmt.Printf("Type of URL : %T\n", parsedURL)

	fmt.Println(parsedURL.Scheme)
	fmt.Println(parsedURL.Host)
	fmt.Println(parsedURL.Path)
	fmt.Println(parsedURL.RawQuery)

	// modifying

	parsedURL.Path = "/newPath"
	parsedURL.RawQuery = "username=yash"

	newUrl := parsedURL.String()
	fmt.Println(newUrl)

}

// scheme : indicates the protocol used
// Host:  domain name
// Path : path
// RawQuery
