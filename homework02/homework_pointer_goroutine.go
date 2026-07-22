package main

import (
	"fmt"
	"time"
)

// 编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，
// 在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。

func add10(p *int) {
	*p += 10
}

func main() {
	// 测试 add10 函数
	// p := 10
	// add10(&p)
	// fmt.Println(p)

	// 测试 multiply2 函数
	// p2 := []int{1, 2, 3, 4, 5}
	// multiply2(&p2)
	// fmt.Println(p2)

	// 测试 printWithGoroutine 函数
	// printWithGoroutine()

	// 测试 taskScheduler 函数
	tasks := []func(){printOdd, printEven}
	taskScheduler(tasks)
	time.Sleep(10 * time.Second)
}

// 实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
func multiply2(p *[]int) {
	for i := range *p {
		(*p)[i] *= 2
	}
}

// 编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
func printOdd() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("奇数:", i)
	}
}

func printEven() {
	for i := 2; i <= 10; i += 2 {
		fmt.Println("偶数:", i)
	}
}

func printWithGoroutine() {
	go printOdd()
	go printEven()
	time.Sleep(10 * time.Second)
}

// 设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
func taskScheduler(tasks []func()) {
	for _, task := range tasks {
		go task()
	}
}
