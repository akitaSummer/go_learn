package main

import "fmt"

// 数组是值类型，arr会被拷贝
func printArray(arr [4]int) {
	arr[2] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func changePtrArray(arr *[4]int) {
	//(*arr)[2] = 100
	arr[2] = 200 // 可不使用(*arr)[2]
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8}
	var grid [4][5]int

	fmt.Println(arr1, arr2, arr3, grid)

	for i := 0; i < len(arr1); i++ {
		fmt.Println(arr1[i])
	}

	for i := range arr2 {
		fmt.Println(arr2[i])
	}

	printArray(arr3)

	// 数组是值类型, arr3[2]没有变为100，仍为6
	for _, v := range arr3 {
		fmt.Println(v)
	}

	changePtrArray(&arr3)

	for _, v := range arr3 {
		fmt.Println(v)
	}
}
