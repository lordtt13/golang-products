package main

import "fmt"

type contactInfo struct {
	email string
	zip int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main(){
	anu := person{
		firstName: "Anusha",
		lastName:  "Gupta",
		contact:contactInfo{
			email: "whatever@gmail.com",
			zip:   632014,
		},
	}

	fmt.Println(anu)
	fmt.Printf("%+v",anu)
}

