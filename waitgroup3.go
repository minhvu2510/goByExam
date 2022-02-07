package main

import (
	"fmt"
	"sync"
)

func read(wg *sync.WaitGroup, data <-chan string) {
	fmt.Println(<-data)
	wg.Done()
}

func write(data chan<- string) {
	data <- "Có làm thì mới có ăn"
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	unbuffer := make(chan string)

	go write(unbuffer)
	go read(wg, unbuffer)

	wg.Wait()

	//time.Sleep(1 * time.Second)
}
