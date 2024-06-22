package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	IsOk bool   `json:"is_Ok"`
}

func main() {
	person := Person{Name: "sam", Age: 23, IsOk: true}
	fmt.Println(person)
	// convert person into jSON(Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Marshalling", err)
		return
	}
	fmt.Println(string(jsonData))

	// Decoding(unMarshalling)

	var personData Person
	errr := json.Unmarshal(jsonData, &personData)
		if errr != nil {
		fmt.Println("Marshalling", errr)
		return
	}
	fmt.Println(personData)

}
