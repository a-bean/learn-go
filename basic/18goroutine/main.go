package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
	goroutine与进程,线程的区别:

	进程,线程
		进程是可执行程序在运行中形成一个独立的内存体,操作系统会以进程为单位,分配系统资源
		线程是轻量级的进程,是cpu调度执行的最小单位
		操作系统将进程和线程都看作一个单独的执行单元
	进程,线程切换开销:
		涉及硬件的上下文切换
		内核栈的切换
		切换前保存执行流程的状态到寄存器
		会导致cpu高速缓存失效
	协程:
		协程将线程的切换从内核态转移到用户态(协程只在用户态工作,避免的内核态跟用户态转化的时间)
		协程可以理解为轻量级的线程
		优势:
			占用空间小,初始占用内存空间2k,可自适应增减或者缩小
			极大减少进程从内核态到用户态的切换,协程切换成本很低

	减少上下文切换的方法:
		尽量避免使用锁
		CAS算法(是一种原子操作): 不需要锁来保护共享资源,避免了锁的开销和线程阻塞
		减少线程数量
		使用协程(在用户空间实行上下文切换)


*/

// Goroutine泄漏是指在 Go 语言中使用 goroutine 时，某些 goroutine 没有按预期终止
// 或退出，导致它们继续在后台运行，浪费系统资源。随着时间推移，如果大量的 goroutine
// 泄漏，可能会导致内存占用增加、性能下降，甚至程序崩溃。
//
// Goroutine 泄漏的常见原因：
// 1. 阻塞在 channel 上：当 goroutine 被阻塞在一个永远不会关闭的 channel 上，或等待一个永远不会发送的值时。
// 2. 无限循环：goroutine 可能进入了一个不受控制的无限循环，消耗资源且无法终止。
// 3. 未处理的退出条件：没有正确处理 goroutine 需要退出的条件，导致其继续运行。
// 4. goroutine 被遗忘：开启了 goroutine 后，未能在合适的时机对它进行管理或清理，导致其一直运行。
// 5. 等待外部事件：goroutine 在等待某些外部资源（例如网络连接、文件 I/O）时没有超时机制，导致它们一直处于挂起状态。
// Goroutine 泄漏的预防措施：
// 1. 合理使用 channel 和 select：
// 如果 goroutine 依赖于 channel 进行通信，确保 channel 能够按预期关闭或发出消息，
// 避免 goroutine 无法收到信号而无限等待。
// 使用 select 语句来处理多个通信渠道，并且在必要时设置超时或退出条件。
func worker1(done chan bool) {
	select {
	case <-done:
		// 正常退出
		return
	case <-time.After(time.Second * 5):
		// 超时处理
		fmt.Println("Timeout!")
		return
	}
}

// 2. 使用 context 控制生命周期：
// Go 提供了 context 包，可以用于控制 goroutine 的生命周期。当 context
// 被取消时，关联的 goroutine 也会终止。
func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// 接收到信号后退出
			fmt.Println("Worker done")
			return
		default:
			// 执行任务
		}
	}
}

// 3. 添加超时机制：
// 在可能长时间运行的操作中，添加超时机制，防止 goroutine 长时间挂起。
// 例如，使用 time.After 在超时时间到达后强制终止操作。
func worker3() {
	select {
	case <-time.After(10 * time.Second):
		// 超时退出
		fmt.Println("Task timed out")
		return
	}
}

// 4. 检查退出条件：
// 在 goroutine 中的循环或长时间运行的任务中，定期检查是否需要退出，防止 goroutine 持续运行。
func worker4(done chan struct{}) {
	for {
		select {
		case <-done:
			// 收到退出信号
			return
		default:
			// 处理正常任务
		}
	}
}

//5. 避免创建未管理的 goroutine：
//确保每个 goroutine 都有明确的管理方式和生命周期，避免启动后没有清晰的退出机制或监控。
//6. 使用工具检测 goroutine 泄漏：
//可以使用一些调试工具或库（如 pprof）来检查是否有大量的 goroutine 持续运行，以及是否存在 goroutine 泄漏。

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

	// 2. 使用 context 控制生命周期：
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)
	// 取消上下文，通知 goroutine 退出
	cancel()

	//6. 使用工具检测 goroutine 泄漏：
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

}
