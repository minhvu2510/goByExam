package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 8; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	// close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

// có 3 worker được start và có 9 work item (9 job) được đẩy vào jobs channel.
// Mỗi worker sẽ bốc một jobs ra xử lý lệnh Println và đẩy kết quả vào results channel.
