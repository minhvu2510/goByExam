package main

import "fmt"

func plus(a int, b int) int {

	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}
func dev(a, b, c int, e string) (int, string) {
	return a + b + c, e
}

func main() {

	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	res, dm := dev(1, 2, 3, "dm")
	fmt.Println("1+2+3 =", res, dm)
}
