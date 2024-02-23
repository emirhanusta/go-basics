package main

import "fmt"

func main() {
	var p1 Person = Person{"John", 25,
		Address{"123 Main St", "Chicago"}}
	p1.Name = "John Doe"
	p1.setName("Michael")
	p1.ToString()
}

type Person struct {
	Name    string
	Age     int
	Address Address
}
type Address struct {
	Street string
	City   string
}

func (person *Person) ToString() {
	fmt.Printf("%s %s %s %d \n", person.Name, person.Address.City, person.Address.Street, person.Age)
}

func (person *Person) setName(name string) {
	person.Name = "Michael"
}
