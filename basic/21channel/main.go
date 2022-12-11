package main

import (
	"fmt"
	"time"
)

/*
不要通过共享内存来通信，而是用通信来共享内存
channel运用场景：

	1.消息传递，消息过滤
	2.消息广播
	3.事件订阅和广播
	4.任务分发
	5.结果汇总
	6.并发控制
	7.同步异步
	.......
*/
func main() {
	// 1. 简单使用
	var msg chan string
	if msg == nil {
		fmt.Println(msg)
	}
	msg = make(chan string, 1) //放值的数量超过容量会阻塞，无缓冲的channel要是没有goroutine消费容易造成阻塞。

	go func() { // happen-before机制，可以保障
		data := <-msg
		fmt.Println(data)
	}()
	msg <- "s" // 存值
	msg <- "s"
	//msg1 := <-msg // 取值
	//fmt.Println(msg1)

	//2. 有缓冲，无缓冲
	msg3 := make(chan string, 0) //无缓冲,适用于 通知。B要第一之间要知道A有没有完成
	msg4 := make(chan string, 1) //有缓冲，适用于消费者和生产者之前的通信

	fmt.Println(msg3)
	fmt.Println(msg4)

	//3. for range遍历channel
	msg5 := make(chan int, 2)
	//go func() { // happen-before机制，可以保障
	//	data := <-msg5
	//	data = <-msg5
	//	data = <-msg5 //下面只传了两个值，这边会阻塞
	//	fmt.Println("data", data)
	//}()
	go func() { // happen-before机制，可以保障
		for data := range msg5 {
			fmt.Println("data", data)
		}
		fmt.Println("all done") // 执行不到
	}()
	msg5 <- 1
	msg5 <- 2
	close(msg5) //可以关闭掉msg5的channel，让55行的range可以退出，然后执行到58行
	//已经关闭的channel可以继续取值，不能存值
	time.Sleep(time.Second)

	// 4. 单向channel
	//var ch1 chan int
	//var ch2 chan<- int // 单向，只能写入int数据
	//var ch3 <-chan int // 单向，只能读int数据
	c := make(chan int, 3)
	var send chan<- int = c
	var read <-chan int = c
	send <- 5
	r := <-read
	fmt.Println(r)

	c1 := make(chan int)
	go producer(c1)
	go consumer(c1)
	time.Sleep(time.Second)

	// 5. select语句：主要作用于多个channel
	go g1(done1)
	go g2(done2)
	// 监听channel
	// 哪一个channel就绪了就执行那个分支。如果两个都就绪了，随机执行。目的：防止饥饿

	timer := time.NewTimer(time.Second)
	select {
	case dd1 := <-done1:
		fmt.Println("done1", dd1)

	case dd2 := <-done2:
		fmt.Println("done2", dd2)
	case <-timer.C: // 一个time.Second之后就执行
		fmt.Println("阻塞了")
		return // 然后直接退出
	}

	time.Sleep(time.Second)
}

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
	}
	close(out)
}

func consumer(in <-chan int) {
	for num := range in {
		fmt.Println("num", num)
	}
}

// select
var done1 = make(chan struct{}) // channel 是多线程安全的
var done2 = make(chan struct{}) // channel 是多线程安全的

func g1(c chan struct{}) {
	fmt.Println("第一个")
	c <- struct{}{}
	c <- struct{}{}
}

func g2(c chan struct{}) {
	fmt.Println("第二个")
	c <- struct{}{}
}
