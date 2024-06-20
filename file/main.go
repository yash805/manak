package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	/*
		file, err := os.Create("data.txt")
		if err != nil {
			fmt.Println("Error creating", err)
			return
		}
		defer file.Close()

		content := "hello world i am here"
		byte, error := io.WriteString(file, content)
		fmt.Println(byte)
		if error != nil {
			fmt.Println("Error writing", error)
		}

		fmt.Println("successfully created")
	*/
	/*
		file, err := os.Open("data.txt")
		if err != nil {
			fmt.Println("Error opening data", err)
			return
		}
		defer file.Close()

		// create a buffer
		buffer := make([]byte, 1024)

		for {
			n, error := file.Read(buffer)
			if error == io.EOF {
				break
			}
			if error != nil {
				fmt.Println("Error reading", error)
			}

			fmt.Println(string(buffer[:n]))
		}
	*/

	content, err := ioutil.ReadFile("data.txt")
	if err != nil {
		fmt.Println("Error reading", err)
		return
	}
	fmt.Println(string(content))

}
