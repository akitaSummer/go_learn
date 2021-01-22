package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	//for {
	//	if n, ok := <-c; !ok {
	//		break
	//	} else {
	//		fmt.Printf("Worker %d received %c\n", id, n)
	//	}
	//}
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}

func CreateWorker(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("CreateWorker %d received %c\n", id, <-c)
		}
	}()
	return c
}

func chanDemo() {
	//var c chan int // c == nil
	//c := make(chan int)
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	var createChannels [10]chan<- int

	for i := 0; i < 10; i++ {
		createChannels[i] = CreateWorker(i)
	}

	for i := 0; i < 10; i++ {
		createChannels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

func bufferedChannel() {
	c := make(chan int, 3)

	go worker(0, c)

	c <- 'a'
	c <- 'b'
	c <- 'c'

	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int, 3)

	go worker(0, c)

	c <- 'a'
	c <- 'b'
	c <- 'c'

	close(c)

	time.Sleep(time.Millisecond)
}

func main() {
	chanDemo()

	bufferedChannel()

	channelClose()
}
