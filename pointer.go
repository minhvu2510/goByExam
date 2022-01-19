package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0 // trỏ đến địa chỉ iptr
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i) // truyền vào địa chỉ biến i
	fmt.Println("zeroptri:", i)

	fmt.Println("pointer:", &i)
	a := 5
	p := &a
	var p1 *int
	p1 = &a
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(p1)
	fmt.Println(*p1)

}
