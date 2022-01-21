package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for j1 := range jobs {
			j, more := <-jobs
			fmt.Println(j1)
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
		return
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	//close(jobs) // nếu đóng routine trên vẫn nhận từ chanel mà ko có du lieu
	fmt.Println("sent all jobs")

	<-done
}
