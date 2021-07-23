package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstname: "jim",
		lastname:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}

	jim.print()
	jim.updateName("jimmy")
	fmt.Printf("%+v", jim)

}

func (p *person) updateName(newFirstName string) {
	p.firstname = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
