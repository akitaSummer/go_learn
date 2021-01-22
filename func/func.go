package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
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
		q, _ := div(a, b)
		return q, nil
	default:
		return 0, fmt.Errorf("unsupported operator: %s", op)
	}
	return result, nil
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

//func div(a, b int) (int, int) {
//	return a / b, a % b
//}

//func div(a, b int) (q, r int) {
//	q = a / b
//	r = a % b
//	return
//}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// go没有默认参数, 可选参数, 只有默认参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	if result, err := eval(3, 4, "x"); err != nil {
		fmt.Println(result)
	} else {
		fmt.Println("Error: ", err)
	}
	a, b := div(34, 3)
	fmt.Println(
		a,
		b,
	)

	fmt.Println(apply(pow, 3, 4))

	fmt.Println(apply(
		func(a, b int) int { // 匿名函数
			return int(math.Pow(float64(a), float64(b)))
		},
		a,
		b))

	fmt.Println(sum(1, 2, 3, 4, 5))
}
