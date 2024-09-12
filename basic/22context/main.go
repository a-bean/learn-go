package main

import (
	"fmt"
	"sync"
	"time"
)

// content 协程的上下文
// 场景：
//      信息传递
//      取消任务
//      超时控制
// 最佳实践
//      将context作为第一个参数传递
//      不要在内层函数创建context
//		及时取消context
//		不要在context中存储大量数据，敏感数据
//		不要滥用context

import (
	"context"
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
