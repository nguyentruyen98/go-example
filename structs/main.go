package main

import "fmt"

type person struct {
	lastName  string
	firstName string
}

func main() {
	var alex person
	alex.firstName = "Truyen"
	alex.lastName = "Nguyen"

	fmt.Printf("%+v", alex)
}
