package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func cpuInfo(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("退出cpu监控")
			return
		default:
			time.Sleep(time.Second * 2)
			fmt.Println("cpu信息")
		}
	}
}

var wg sync.WaitGroup

// context提供了四种函数：WithCancel,WithTimeout,WithValue,WithDeadline
func main() {
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	ctx1, _ := context.WithCancel(ctx) // context具有传递性，父级cancel了，所有的子集也都会被cancel

	ctx2, cancel2 := context.WithTimeout(context.Background(), 6*time.Second)
	fmt.Println(ctx2, cancel2)

	ctx3 := context.WithValue(ctx2, "kk", 45)
	fmt.Println(ctx3)

	go cpuInfo(ctx1)

	time.Sleep(time.Second * 6)
	cancel()
	wg.Wait()
}
