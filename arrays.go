package main

// arr kich thuoc co dinh, cung kieu du lieu
import "fmt"

func main() {

	a := []int{1, 2, 3}
	fmt.Println(a)

	var b [1]int
	fmt.Println(len(b))
	b[0] = 1

	// b[5] = 0
	fmt.Println(b)

	// ko caanf kich thuoc
	numbers := []int{}
	fmt.Println(numbers)
	fmt.Printf("Don't know type %T\n", numbers)
	// numbers[2] = 3
	// numbers[0] = 1

}
