package main

import (
	"context"
	"fmt"
	"time"
)

// 一、Context 的背景与设计目标
// 在 Go 中，goroutine 是并发的基本单位，但如何优雅地管理 goroutine 的生命周期（例如取消任务、传递请求范围内的值、设置超时）曾经是一个挑战。context 包的引入（Go 1.7）解决了这些问题，提供了统一的机制来处理：

// 取消信号：通知 goroutine 停止执行。
// 超时和截止时间：控制任务的执行时限。
// 请求范围的数据传递：在调用链中传递上下文信息。
// context 的设计灵感来源于 Google 内部的分布式系统实践，旨在标准化 goroutine 的协作和生命周期管理。

// content 协程的上下文
// 场景：
//
//	信息传递
//	取消任务
//	超时控制
//
// 最佳实践
//
//	将context作为第一个参数传递
//	不要在内层函数创建context
//	及时取消context
//	不要在context中存储大量数据，敏感数据
//	不要滥用context
//	每个 goroutine 应该使用自己的 Context，避免复用父 Context。
//	避免过度使用 Value
func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func slowOperation(ctx context.Context) (string, error) {
	select {
	case <-time.After(3 * time.Second): // 模拟耗时操作
		return "Operation completed", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

func handler(ctx context.Context) {
	if userID, ok := ctx.Value("userID").(string); ok {
		fmt.Println("Handling request for user:", userID)
	}
}

func main() {
	// 1. 取消 goroutine: WithCancel
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)
	time.Sleep(2 * time.Second)
	cancel()                    // 手动取消
	time.Sleep(1 * time.Second) // 等待 worker 退出

	// 2. 控制任务执行时间：WithTimeout
	ctx1, cancel1 := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel1() // 确保释放资源
	result, err := slowOperation(ctx1)
	fmt.Println("Result:", result, "Error:", err)
	// 输出: Result:  Error: context deadline exceeded

	// 3. 设置截止时间
	deadline := time.Now().Add(2 * time.Second)
	ctx2, cancel2 := context.WithDeadline(context.Background(), deadline)
	defer cancel2()
	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Operation completed")
	case <-ctx.Done():
		fmt.Println("Deadline exceeded:", ctx2.Err())
	}

	// 4. 传递数据
	ctx3 := context.WithValue(context.Background(), "userID", "12345")
	handler(ctx3)
	// 输出: Handling request for user: 12345

	// 与 Channel 结合
	ctx4, cancel4 := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel4()
	ch := make(chan int)
	go worker1(ctx4, ch)
	ch <- 1
	time.Sleep(3 * time.Second)
}

func worker1(ctx context.Context, ch chan int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopped:", ctx.Err())
			return
		case v := <-ch:
			fmt.Println("Received:", v)
		}
	}
}
