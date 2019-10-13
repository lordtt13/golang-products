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

	anu.update("Anushaa")
	anu.print()
}

func (pointertoPerson *person) update(firstName string){
	(*pointertoPerson).firstName = firstName
}

func (p person) print() {
	fmt.Println(p)
	fmt.Printf("%+v",p)
}