package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}

func (englishBot) getGreeting() string{
	return "Hi There!"
}

func (spanishBot) getGreeting() string{
	return "Hola!"
}

func printGreeting(b bot){
	fmt.Println(b.getGreeting())
}

type Car interface {
	Drive()
	Stop()
}
type Lambo struct {
	LamboModel string
}
func NewModel(arg string) Car {
	return &Lambo{arg}
}
type Chevy struct {
	ChevyModel string
}
func (l *Lambo) Stop() {
	fmt.Println("Stopping lambo")
}
func (l *Lambo) Drive() {
	fmt.Println("Lambo on the move")
	fmt.Println(l.LamboModel)
}
func (c *Chevy) Drive() {
	fmt.Println("Chevy on the move")
	fmt.Println(c.ChevyModel)
}

func main(){
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	l := NewModel("Gallardo")
	c := Chevy{"C369"}
	l.Drive()
	c.Drive()
}