package main

import (
     "fmt"
	 "net/http"
	 "io/ioutil"
	//  "encoding/json"
)


type Todo struct {
	UserId int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Comple bool `json:"comple"`
}
func main(){
	fmt.Println("learning")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("error occured:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("error occured:", res.Status)
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error occured:", err)
		return
	}

	fmt.Println(string(data))

	// var todo Todo 
	// err = json.NewDecoder(res.Body).Decode(&todo)
	// if err != nil {
	// 	fmt.Println("error occured:", err)
	// 	return
	// }
	
	// fmt.Println(todo)

}