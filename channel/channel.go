package main

import (
	"fmt"
	"time"
)

func createWorker (id int) chan<- int{
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("Worker %d received %c\n",
				id, <-c)
		}
	}()
	return c
}


func chanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int)

	c <- 1
}

func main() {

}