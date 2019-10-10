package main

import "fmt"

func main(){
	cards := wholeDeck()

	fmt.Println(cards.toString())
	cards.savetoFile("my_cards")
}
