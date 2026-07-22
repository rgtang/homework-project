package main

import (
	"fmt"
	"sync"
)

// 编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。
// 启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
func incrementCounter(counter *int, lock *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		lock.Lock()
		*counter++
		lock.Unlock()
	}
}

func main() {
	counter := 0
	lock := &sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go incrementCounter(&counter, lock, wg)
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
