package main

import (
	"fmt"
	"sync"
)

type Worker struct {
	in   chan int
	done func()
}

func worker(id int, c chan int, done func()) {
	//for {
	//	if n, ok := <-c; !ok {
	//		break
	//	} else {
	//		fmt.Printf("Worker %d received %c\n", id, n)
	//	}
	//}
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		done()
	}
}

func CreateWorker(id int, wg *sync.WaitGroup) Worker {
	w := Worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go worker(id, w.in, w.done)
	return w
}

func chanDemo() {
	var workers [10]Worker
	var wg sync.WaitGroup
	for i := range workers {
		workers[i] = CreateWorker(i, &wg)
	}

	wg.Add(20)

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	wg.Wait()
}

func main() {
	chanDemo()
}
