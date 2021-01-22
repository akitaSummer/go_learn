package main

import "fmt"

// 寻找最长不含有重复字符的子串abcabcbb -> abc
func noRepeating(s string) int { // 仅支持英文
	lastOccurred := make(map[byte]int)
	start := 0
	maxLangth := 0
	for i, ch := range []byte(s) {
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

// 除了slice, map, function的内建类型均可作为key
// Struct不包含上述字段，也可作为key
func main() {
	m := map[string]string{
		"name":   "akita",
		"course": "golang",
		"site":   "beijing",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil

	fmt.Println(m, m2, m3)

	for k, v := range m {
		// k, v 是无序的，有时name在前，有时其他在前
		fmt.Println(k, v)
	}

	courseName, ok := m["course"]

	fmt.Println(courseName, ok)

	// 取不存在的属性则获得zeroValue
	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName)
	} else {
		fmt.Println("key does not exist")
	}

	fmt.Println(m)
	delete(m, "name")
	fmt.Println(m)

	fmt.Println(noRepeating("abcabcbb"))
}
