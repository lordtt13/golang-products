package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func wholeDeck() deck{
	cards := deck{}

	suits := []string{"Hearts","Diamonds","Spades","Clubs"}
	values := []string{"Ace","Two","Three","Four","Five","Six","Seven","Eight","Nine","Ten","Jack","Queen","King"}

	for _,suit := range suits{
		for _,value := range values{
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

func (d deck) print(){
	for i, card := range d{
		fmt.Println(i, card)
	}
}

func deal(d deck, size int) (deck , deck){
	return d[:size],d[size:]
}

func (d deck) toString() string{
	return strings.Join([]string(d),",")
}

func (d deck) savetoFile(filename string) error{
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}