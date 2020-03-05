package main

import "fmt"

func main() {
	test("hello",",world")
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
		fmt.Print(" ")
	}

	flag2 := true
	// continue, break, switch case,
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			flag2 = false
			break
		} else if i == 1 {
			continue
		}
	}
	fmt.Println(flag2)

	day := "Fri"
	switch day {
	case "Fri":
		fmt.Println("TGIF")
	case "Mon", "Tue", "Wed":
		fmt.Println("Boring")
	default:
		fmt.Println("default")
	}

	switch {
	case day == "Fri":
		fmt.Println("TGIF")
		break
	}
}

func test(i ...interface{}){
	fmt.Print(i...)
}
