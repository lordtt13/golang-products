package main

import (
	"fmt"
	"strings"
)

func swap(m1, m2 *int) {
	var temp int
	temp = *m2
	*m2 = *m1
	*m1 = temp
}


 func main() {
 	m1 := "my name"
 	m2 := "name"
 	fmt.Println(strings.Split(m1, " "), m1+m2)
	 m3 := 2
	 m4 := 3
	 fmt.Println(m3 + m4)
	 m5, m6 := 2, 3
	 fmt.Println(m5, m6)
	 swap(&m5, &m6)
	 fmt.Println(m5, m6)
 }
