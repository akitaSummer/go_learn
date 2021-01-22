package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 包内部变量
var aa = 3
var ss = "ss"

//bb := true // 函数外不能用:=声明变量

// 集中定义
var (
	bb = true
	cc = "ccc"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	// 可不声明类型，go会自己推断
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	// 函数内可用:=代替var
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

/**
go内建变量类型
bool, string
(u)int (u)int8 (u)int16 (u)int32 (u)int64 uintptr
byte rune(类似于char但是为32位)
float32 float64 complex64 complex128
*/

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
	// 欧拉公式
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	var a, b = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	// 强制类型转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	// go中常量通常不为全大写，因为在go中大写有特殊意义
	const filename string = "abc.txt"
	const a, b = 3, 4
	const (
		d = true
		e = 3.34
		f = 5 + 4i
	)
	var c int
	// 未定义a,b可做int也可做float，即文本，可做各种类型使用，在编译时会自动转换
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, a, b, c, d, e, f)
}

func enums() {
	// go中枚举的实现
	const (
		cpp = iota // 自增值
		_          // 被跳掉
		javascript
		python
		golang
	)

	// b, kb, mb, gb, tb, pb
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, javascript, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello World", aa, ss, bb, cc)

	variableZeroValue()

	variableInitialValue()

	variableTypeDeduction()

	variableShorter()

	euler()

	triangle()

	consts()

	enums()
}
