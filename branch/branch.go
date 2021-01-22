package main

import (
	"fmt"
	"io/ioutil"
)

func readFile() {
	const filename = "branch/abc.txt"
	// contents, err 在if中，生存期仅在if中
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func eval(a, b int, op string) int {
	var result int
	// switch 默认加break, 除非有fallthrough
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op) // 报错
	}
	return result
}

func grade(score int) string {
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		return "F"
	case score < 80:
		return "C"
	case score < 90:
		return "B"
	case score <= 100:
		return "A"
	default:
		panic(fmt.Sprintf("Wrong score: %d", score))
	}
}

func main() {
	const filename = "branch/abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}

	readFile()

	fmt.Println(
		eval(1, 2, "+"),
		eval(1, 2, "-"),
		eval(1, 2, "*"),
		eval(1, 2, "/"),
	)

	fmt.Println(
		grade(0),
		grade(60),
		grade(70),
		grade(80),
		grade(90),
		grade(100),
	)
}
