package main

import "fmt"

func main() {

	mapTest := []int{}
	for j := 0; j < 5; j++ {
		mapTest = append(mapTest, j+1)
	}
	fmt.Println(mapTest)
	for i, j := range mapTest {
		fmt.Println("it = ", i, j)
	}
	// map
	mapT := map[string]int{"dm": 1}
	for i, j := range mapT {
		fmt.Println("mapT = ", i, j)
	}

}
