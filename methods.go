package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	r.width = 1
	return r.width
}

func (r rect) perim() int {
	r.width = 1
	return r.width
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area: ", r.area())
	fmt.Println(r)
	fmt.Println("perim:", r.perim())

	r1 := rect{width: 10, height: 5}

	// rp := &r
	fmt.Println("perim:", r1.perim())
	fmt.Println("area: ", r1)
}

// area:  1
// {1 5}
// perim: 1
// perim: 1
// area:  {10 5}
