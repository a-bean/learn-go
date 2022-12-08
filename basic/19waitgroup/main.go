package main

import (
	"fmt"
	"sync"
)

// sync.WaitGroup 主要用于goroutine的执行等待。

func main() {
	var wg sync.WaitGroup
	wg.Add(100) // 必须与wg.Done一块使用
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
	fmt.Println("Wait")
}
