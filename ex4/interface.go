package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

func (p person) String() string {
	return p.firstname + " -- " + p.lastname
}

func main() {
	p := person{firstname: "Thomas", lastname: "Dupont"}
	fmt.Println(p)
}
