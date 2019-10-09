package main

func main(){
	cards := deck{"Two of Diamonds", newCard()}
	cards = append(cards, "Three of Diamonds")

	cards.print()
}

func newCard() string{
	return "Ace of Diamonds"
}