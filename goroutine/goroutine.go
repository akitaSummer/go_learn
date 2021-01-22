package main

import (
	"fmt"
	"runtime"
	"time"
)

// 协程 Coroutine
// 轻量级"线程"
// 非抢占式多任务处理，由协程主动交出控制权
// 编译器/解释器/虚拟机层面多任务
// 多个协程可能在一个或者多个线程上执行

// 子程序是协程的一个特例（如函数）
func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) { // race condition
			for {
				a[i]++
				runtime.Gosched()
			}
		}(i)
	}

	time.Sleep(time.Millisecond)

	fmt.Println(a)
}
