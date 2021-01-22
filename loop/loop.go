package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func convertToBin(n int) string {
	if n == 0 {
		return "0"
	}
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	// 死循环
	for {
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),  // 101
		convertToBin(13), // 1011 --> 1101
		convertToBin(0),
		convertToBin(321423432),
	)

	printFile("branch/abc.txt")
}
