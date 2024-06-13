package main
func main()  {

	student := make(map[string]int)

	student["yash"] =  54
	student["aman"] = 34
	student["amit"] = 64

	fmt.Println(student["amans"])	
	student["yash"] =  100
	delete(student, "yash")

	grades, exists := student["amit"]
	fmt.Println(grades)
	fmt.Println(exists)
    
	gradess, existss := student["aman"]
	fmt.Println(gradess)
	fmt.Println(existss)
    

	for index, value := range student{
        fmt.Printf("key is %s and marks is %d\n", index , value)
	}


	person := map[string] int{
		"sam": 30,
		"john": 50,
		"anexa": 32,
	}
	for index, value := range person{
        fmt.Printf("key is %s and marks is %d\n", index , value)
	}
}