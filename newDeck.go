package main

import "fmt"

func main(){
	cards := wholeDeck()

	fmt.Println(cards.toString())
	cards.savetoFile("my_cards")

	cardsNew := readfromFile("my_cards")
	cardsNew.print()
}
