package main

import "fmt"

func main(){
	cards := []string{"Two of Diamonds", newCard()}
	cards = append(cards, "Three of Diamonds")

	for i, card := range cards{
		fmt.Println(i, card)
	}
}

func newCard() string{
	return "Ace of Diamonds"
}