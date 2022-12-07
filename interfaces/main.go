package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) {
// 	fmt.Printf(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Printf(sb.getGreeting())
// }

func (eb englishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hi there!"
}
func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
