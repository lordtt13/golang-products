package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

func main(){
	anu := person{
		firstName: "Anusha",
		lastName:  "Gupta",
	}

	fmt.Println(anu)
}

