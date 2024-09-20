package main

import (
	"fmt"
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


*/

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
