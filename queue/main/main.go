package main

import (
	"fmt"
	"learn_go/queue"

	"golang.org/x/tools/container/intsets"
)

func testSparse() {
	s := intsets.Sparse{}

	s.Insert(1)
	s.Insert(100)
	s.Insert(1000)
	s.Insert(10000000)

	fmt.Println(s.Has(100))
	fmt.Println(s.Has(10000))
}

func main() {
	q := queue.Queue{1}

	q.Push(2)
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	q.Push("abc")
	fmt.Println(q.Pop())
	testSparse()
}
