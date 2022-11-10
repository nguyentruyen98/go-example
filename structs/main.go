package main

import "fmt"

type contractInfo struct {
	email   string
	zipCode int
}

type person struct {
	lastName  string
	firstName string
	contractInfo
}

func main() {
	jim := person{firstName: "Jim", lastName: "Party", contractInfo: contractInfo{email: "jim@gmail.com", zipCode: 70000}}
	test := 123
	jimPointer := &jim
	fmt.Println(&test)

	jimPointer.updateName("Truyen")
	// jim.updateName("Truyen")
	jim.print()
}

func (pointerToPersion *person) updateName(newFirstName string) {
	fmt.Println(pointerToPersion)
	(*pointerToPersion).firstName = newFirstName

}

func (p person) print() {
	fmt.Printf("%+v", p)
}
