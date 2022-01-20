package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}
func main() {
	done := make(chan bool) // tạo chanel dạng bool, truyền duwois dạng tham số vòa go rountines
	fmt.Println("Main going to call hello go goroutine")
	go hello(done) // gọi gp rt
	<-done         //  nhận dữ liệu từ chanel, phải chờ go routine ghi dữ liệu vòa chanel mới chuyển sang main
	fmt.Println("main function")
}
