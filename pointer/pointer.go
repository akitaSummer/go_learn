package main

import "fmt"

// go是值传递，对象传递的是地址值
func swap(a, b *int) {
	*a, *b = *b, *a
}

// go的指针不能运算
func main() {
	var a int = 2
	var pa *int = &a
	*pa = 3

	fmt.Println(a) // 3

	b, c := 3, 4
	fmt.Println(b, c) // 3 4
	swap(&b, &c)
	fmt.Println(b, c) // 4 3
}
