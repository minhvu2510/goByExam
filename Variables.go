package main

import (
	"fmt"
	"reflect"
)

func main() {

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)

	d1 := 5.99
	d2 := 6
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(d1))
	fmt.Println(reflect.TypeOf(d2))
}
