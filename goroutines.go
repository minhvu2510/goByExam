package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func main() {
	go hello()
	time.Sleep(1 * time.Second) // khi go hello chạy, hệ thống ko chờ goroutines chạy xong mà chạy tiếp luôn, và kết thúc nên cần sleep
	fmt.Println("main function")
	go func(msg string) {
		fmt.Println(msg)
	}("going....")
	time.Sleep(time.Second)
	fmt.Println("done")

}

// https://techmaster.vn/posts/35065/series-golang-co-ban-phan-21-goroutines
