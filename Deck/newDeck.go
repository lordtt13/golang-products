package main

func main(){
	cardsNew := readfromFile("my_cards")
	cardsNew.shuffle()
	cardsNew.print()
}
