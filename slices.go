package main

import "fmt"

func main() {

	dev := make([]int, 5, 6)
	// liece := make([]int, 5)
	dev[4] = 0
	fmt.Println(dev)
	// var s []int
	s1 := []int{}
	s1 = append(s1, 5)
	fmt.Println("s1:", s1)

	s := make([]string, 3)
	fmt.Println("emp:", s)

	var s2 []int
	s2 = append(s2, 6)
	fmt.Println("s2:", s2)

	s3 := make([]int, 0)
	s3 = append(s3, 7)
	fmt.Println("s3:", s3)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
