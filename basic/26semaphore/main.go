/*
信号量:
	信号量的概念是计算机科学家 Dijkstra （Dijkstra算法的发明者）提出来的，广泛应用在不同的操作系统中。
	系统中，会给每一个进程一个信号量，代表每个进程当前的状态，未得到控制权的进程，会在特定的地方被迫停下来，
	等待可以继续进行的信号到来。

	如果信号量是一个任意的整数，通常被称为计数信号量（Counting semaphore），或一般信号量（general semaphore）；
	如果信号量只有二进制的0或1，称为二进制信号量（binary semaphore）。在linux系统中，二进制信号量（binary semaphore）
	又称互斥锁（Mutex）


	运行方式：
	1. 初始化信号量，给与它一个非负数的整数值。
	2. 运行P（wait()），信号量S的值将被减少。企图进入临界区的进程，需要先运行P（wait()）。
		当信号量S减为负值时，进程会被阻塞住，不能继续；当信号量S不为负值时，进程可以获准进入临界区。
	3. 运行V（signal()），信号量S的值会被增加。结束离开临界区段的进程，将会运行V（signal()）。
		当信号量S不为负值时，先前被阻塞住的其他进程，将可获准进入临界区。
	4. 运行Acquire方法: 相当P操作,支持一次获取多个资源
	5. Release方法: 相当V操作，支持释放多个资源
	6. TryAcquire方法: 尝试获取多个资源,但是不会阻塞

	我们一般用信号量保护一组资源，比如数据库连接池、一组客户端的连接等等。每次获取资源时都会将信号量中的
	计数器减去对应的数值，在释放资源时重新加回来。当信号量没资源时尝试获取信号量的线程就会进入休眠，等待
	其他线程释放信号量。如果信号量是只有0和1的二进位信号量，那么，它的 P/V 就和互斥锁的 Lock/Unlock 就一样了。

golang.org/x/sync/semaphore 对外提供了四个方法：
	1. semaphore.NewWeighted 用于创建新的信号量，通过参数(n int64) 指定信号量的初始值。
	2. semaphore.Weighted.Acquire 阻塞地获取指定权重的资源，如果当前没有空闲资源，就会陷入休眠等待；
		相当于 P 操作，你可以一次获取多个资源，如果没有足够多的资源，调用者就会被阻塞。它的第一个参数是
		Context，这就意味着，你可以通过 Context 增加超时或者 cancel 的机制。如果是正常获取了资源，
		就返回 nil；否则，就返回 ctx.Err()，信号量不改变。
	3. semaphore.Weighted.Release 用于释放指定权重的资源；相当于 V 操作，可以将 n 个资源释放，返还给信号量。
	4. semaphore.Weighted.TryAcquire 非阻塞地获取指定权重的资源，如果当前没有空闲资源，就会直接返回 false；


*/

package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

func doSomething(u string) { // 模拟抓取任务的执行
	fmt.Println(u)
	time.Sleep(2 * time.Second)
}

const (
	Limit  = 3 // 同時并行运行的goroutine上限
	Weight = 1 // 每个goroutine获取信号量资源的权重
)

func main() {
	urls := []string{
		"http://www.example.com",
		"http://www.example.net",
		"http://www.example.net/foo",
		"http://www.example.net/bar",
		"http://www.example.net/baz",
	}
	s := semaphore.NewWeighted(Limit)
	var w sync.WaitGroup
	for _, u := range urls {
		w.Add(1)
		go func(u string) {
			err := s.Acquire(context.Background(), Weight)
			if err != nil {
				return
			}
			doSomething(u)
			s.Release(Weight)
			w.Done()
		}(u)
	}
	w.Wait()

	fmt.Println("All Done")
}
