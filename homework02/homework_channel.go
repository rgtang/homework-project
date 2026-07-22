package main

import (
	"fmt"
	"sync"
)

// 编写一个程序，使用通道实现两个协程之间的通信。
// 一个协程生成从1到10的整数，并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
func generateNumbers(channel chan int) {
	for i := 1; i <= 10; i++ {
		channel <- i
	}
	close(channel) // 关闭通道
}

func printNumbers(channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for number := range channel {
		fmt.Println(number)
	}
}

func main() {
	// channel := make(chan int)
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go generateNumbers(channel)
	// go printNumbers(channel, &wg)
	// wg.Wait()

	channel := make(chan int, 100)
	var wg sync.WaitGroup
	wg.Add(1)
	go producer(channel)
	go consumer(channel, &wg)
	wg.Wait()
}

// 实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
func producer(channel chan int) {
	for i := 1; i <= 100; i++ {
		channel <- i
	}
	close(channel)
}

func consumer(channel chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for number := range channel {
		fmt.Println(number)
	}
}
