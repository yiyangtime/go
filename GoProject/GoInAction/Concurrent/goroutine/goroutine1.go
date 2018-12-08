package goroutine

import (
	"fmt"
	"runtime"
	"sync"
)

// wg 用来等待程序完成
var wg sync.WaitGroup

// main 是所有 Go 程序的入口
func main() {
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1) // 计数加 2，表示要等待两个 goroutine
	wg.Add(2)

	// 创建两个 goroutine
	fmt.Println("Create Goroutines")
	go printPrime("A")
	go printPrime("B")

	// 等待 goroutine 结束
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

// printPrime 显示 50000 以内的素数值
func printPrime(prefix string) {
	// 在函数退出时调用 Done 来通知 main 函数工作已经完成
	defer wg.Done()

next:
	for outer := 2; outer < 50000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer%inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}
