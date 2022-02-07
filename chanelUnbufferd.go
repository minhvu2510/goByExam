package main

import (
	"fmt"
)

func main() {
	unbuffer := make(chan string)
	unbuffer <- "Có làm thì mới có ăn"
	fmt.Println(<-unbuffer)
	// -> ko chay vi chenel can 2 go routine tro len, o day chi co 1 goroutin (2 goroutin can bat dau caung luc -> dua nhan len tren)
}

// // Cách 1: Move đoạn code đọc dữ liệu ra lên trước
// func main() {
// 	unbuffer := make(chan string)
// 	go func() {
// 		fmt.Println(<-unbuffer)
// 	}()
// 	unbuffer <- "Có làm thì mới có ăn"

// }

// // Cách 2: Wrap đoạn code đưa dữ liệu vào trong 1 goroutines
// func main() {
// 	unbuffer := make(chan string)
// 	go func() {
// 		unbuffer <- "Có làm thì mới có ăn"
// 	}()
// 	fmt.Println(<-unbuffer)

// }
