package main

import "fmt"

// slice是引用类型
func updateSlice(s []int) {
	s[0] = 100
}
func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("arr[2:6] =", arr[2:6])
	fmt.Println("arr[:6] =", arr[:6])
	fmt.Println("arr[2:] =", arr[2:])
	fmt.Println("arr[:] =", arr[:])

	s := arr[2:]

	updateSlice(s)
	// arr也会更新
	fmt.Println("s:", s, arr[:])

	s2 := arr[2:]
	fmt.Println("s2 =", s2)
	s2 = s2[2:]
	fmt.Println("s2= ", s2)

	s3 := arr[2:6]
	// slice 会向后扩展，不会向前扩展
	s4 := s3[3:5]
	fmt.Println("s3 =", s3)
	fmt.Println("s4 =", s4)
	fmt.Printf("s3 = %v, len(s3) = %d, cap(s3)=%d\n", s3, len(s3), cap(s3))
	fmt.Printf("s4 = %v, len(s4) = %d, cap(s4)=%d\n", s4, len(s4), cap(s4))
}
