package main

import (
	"fmt"
	"unicode/utf8"
)

// strings库中包含大量字符串操作
// Fields(认识空格分隔), Split, Join
// Contains, Index
// ToLower, ToUpper
// Trim, TrimRight, TrimLeft

func noRepeating(s string) int { // 支持中文
	lastOccurred := make(map[rune]int)
	start := 0
	maxLangth := 0
	for i, ch := range []rune(s) {
		// lastOccurred[ch]不存在时，会取zeroValue，int为0，是合法值
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i-start+1 > maxLangth {
			maxLangth = i - start + 1
		}
		lastOccurred[ch] = i

	}
	return maxLangth
}

func main() {
	s := "Yes是!"
	fmt.Println(len(s)) // 中文使用utf-8，占3个字节，1 1 1 3 1，所以长度为7

	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()

	for i, ch := range s { // ch is rune urf-8转成了Unicode
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()

	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c)", i, ch)
	}
	fmt.Println()

	fmt.Println(noRepeating("黑化肥挥发发灰会花飞灰化肥挥发发黑会飞花"))
}
