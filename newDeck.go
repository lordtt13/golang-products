package main

func main(){
	cards := wholeDeck()

	hand , remainingCards := deal(cards, 13)
	hand.print()
	remainingCards.print()
}
