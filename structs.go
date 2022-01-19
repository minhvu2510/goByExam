package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {

	p := person{name: name}
	p.age = 42
	return &p
}
func dmPerson(p *person) {
	p.name = "vudeptrai"
	p.age = 18
	// return p
}
func main() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name: "Alice", age: 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	// s := person{name: "Sean", age: 50}
	s := newPerson("dd")
	fmt.Println(s)
	fmt.Println(*s)
	fmt.Println(s.name)
	fmt.Println(s.name)
	fmt.Println(s.name)

	sp := &s

	dmPerson(s)
	fmt.Println("---", *sp)
	fmt.Println("---", s)

	sp1 := s
	fmt.Println(sp1.age)

	// sp.age = 51
	// fmt.Println(sp.age)
}
