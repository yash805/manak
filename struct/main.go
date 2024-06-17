package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	House int
	Area  string
	state string
}

type Employee struct {
	Person_details Person
	Person_Contact Contact
	Person_Address Address
}

func main() {
	var yash Person
	// fmt.Println("person", yash)
	yash.firstName = "yash"
	yash.lastName = "kumar"
	yash.Age = 25
	// fmt.Println(yash)

	// 2nd method
	amit := Person{
		firstName: "John",
		lastName:  "alan",
		Age:       34,
	}
	fmt.Println(amit)

	// 3rd method
	var person2 = new(Person)
	person2.firstName = "sam"
	person2.lastName = "clark"
	person2.Age = 22

	// fmt.Println(person2)

	// fmt.Println(yash.Age)

	var employee1 Employee
	employee1.Person_details = Person{
		firstName: "jack",
		lastName:  "conavy",
		Age:       24,
	}
	employee1.Person_Contact.Email = "conavy@gmail.com"
	employee1.Person_Contact.Phone = "9559594769"

	employee1.Person_Address = Address{
		House: 14,
		Area:  "North",
		state: "South",
	}

	fmt.Println(employee1)
}

// e:= Employee{
// 	person: Person{

// 	},
// 	position: "engineer"
