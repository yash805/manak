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
func pGetRequest (){
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


func pPostRequest(){
	 todo := Todo {
		UserID:22,
		Title: "welcome",
		Completed: true
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error handle",err)
		return
	}

	jsonString := string(jsonData)

	jsonReader := strings.NewReader(jsonString)

    myURL := "https://jsonplaceholder.typicode.com/todos"

	res, err := http.Post(myURL, "application/json",jsonReader)
	if err != nil {
		fmt.Println("Error handling",err)
		return
	}
    
	defer res.Body.Close()

	// data, _ := ioutil.ReadAll(res.Body)
	// fmt.Println("Response :", string(data))

	fmt.Println(res.status)

}


func pUpdateRequest(){
	todo := Todo {
		UserID:22378,
		Title: "welcome here",
		Completed: false
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marsh", err)
		return
	}

	jsonString := string(jsonData)

	jsonReader := strings.NewReader(jsonString)

	const myurl := "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest(http.MethodPut, myurl, jsonReader)
	if err != nil {
		fmt.Println("Error marsh", err)
		return
	}
   req.Header.Set("Content-Type", "application/json")

   client := http.Client()

   res, err := client.Do(req)

   if err != nil {
	fmt.Println("Error marsh", err)
	return
    }

	defer res.Body.Close()

data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response :", string(data))

}
func main(){
	fmt.Println("learning")
}