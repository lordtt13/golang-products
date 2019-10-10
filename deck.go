package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func readfromFile(filename string) deck{
	bs, err := ioutil.ReadFile(filename)
	if err!= nil{
		fmt.Println("Error: ",err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs),","))
}

func (d deck) shuffle(){
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d{
		newPos := r.Intn(len(d))

		d[i], d[newPos] = d[newPos], d[i]
	}
}