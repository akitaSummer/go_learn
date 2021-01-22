package main

import (
	"bufio"
	"fmt"
	"learn_go/functional/fib"
	"os"
)

func tryDefer() {
	// defer是个栈
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred")
	fmt.Println(4)
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	//err = errors.New("this is a custom error")
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Printf("%s, %s, %s\n", pathError.Op, pathError.Path, pathError.Err)
		}
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer() // 3 2 1 error
	writeFile("./errhandling/defer/fib.txt")
}
