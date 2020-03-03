package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	// if else, for, switch case, break continue
	f := true
	flag := &f
	if flag == nil {
		fmt.Println("Value is nil")
	} else if *flag {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	arr := []string{"Tanmay", "Thakur", " "}
	for i, value := range arr {
		fmt.Println(i)
		fmt.Println(value)
	}
	mymap := make(map[string]interface{})
	mymap["name"] = "Tanmay"
	mymap["age"] = 21
	for k, v := range mymap {
		fmt.Printf("Key: %s and Value: %v", k, v)
	}
}
