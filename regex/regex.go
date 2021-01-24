package main

import (
	"fmt"
	"regexp"
)

const text = `
My email is 644171127@qq.com
My email is aaa@qq.kk.com
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9.]+)(\.[a-zA-Z0-9]+)`)

	//match := re.FindString(text)
	//match := re.FindAllString(text, -1)

	//fmt.Println(match)
	match := re.FindAllStringSubmatch(text, -1)

	for _, m := range match {
		fmt.Println(m)
	}
}
