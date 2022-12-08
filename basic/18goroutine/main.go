package main

import (
	"fmt"
	"time"
)

func asyncPrint() {
	fmt.Println("goroutine")
}

func main() {
	// 主死随从
	//go asyncPrint()

	//go func() {
	//	for {
	//		time.Sleep(time.Second)
	//		fmt.Println("goroutine")
	//	}
	//}()
	for i := 0; i < 100; i++ {
		go func(j int) {
			time.Sleep(time.Second)
			fmt.Println("goroutine:", j)
		}(i)
	}

	time.Sleep(time.Second * 10)
}
