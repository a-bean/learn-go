package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// lock 解决资源竞争问题。本质是将并行的代码串行化了
// 使用lock会影响性能，即使设计锁，也要尽量保持并行

/*

锁是并发编程中的同步原语，他可以保证多线程在访问同一片内存时不会出现竞争来保证并发安全。
对于获取锁，一般来讲有两种方案，一种是不断地自旋+CAS，另一种就是阻塞+唤醒。两种方式各
有优劣。Go语言结合了这两种方案，自动的判断当前锁的竞争情况，先尝试自旋几次，如果锁一直
没被释放，再加入阻塞队列。

锁竞争方案		优势							劣势									适用场景
阻塞/唤醒		精准打击，不浪费CPU时间片		需要挂起协程，进行上下文切换，操作较重	并发竞争激烈的场景
自旋+CAS			无需阻塞协程，短期来看操作较轻	长时间争而不得，会浪费CPU时间片		并发竞争强度低的场景



读写锁和互斥锁的一些规则
	1. 不可重入性: 不允许读锁之后在获取写锁，不允许获取写锁之后在获取写锁
	2. 写锁只有在读锁和写锁都处于未加锁的状态下才能成功加锁
	3. 加锁和解锁可以由不同的协程来执行
	4. 同一时间只有一个协程能获取写锁，读锁会阻塞写锁
	5. 读锁不会阻塞读锁
	6. 未上锁的情况下调用unlock解锁会报panic

原子操作与锁的区别:
	原子操作: 不可被其他线程中断的操作,要么执行完成,要么不执行,可以理解为变量级别的互斥锁,真正能够保证原子性执行的只有原子操作
	锁: 是一种同步机制,用于确保多个线程在访问共享资源时不会发生冲突
	区别:
		实现方式不同: 原子操作是通过底层的cpu指令完成的,由cpu提供芯片级别的支持.互斥锁是在软件层面实现的，由操作系统提供支持
		保护范围不同: 原子操作保护的对象是单个变量,锁可以保护一段代码片段
		性能表现不同: 原子操作是底层硬件支持,而且保护范围很小,所以性能更好

条件变量: 标准库cond用于解决等待/通知场景下的并发问题
	注意点:
		cond.Wait()的调用必须先加锁
		cond.Wait(),cond.Signal()不能同时在主goroutine调用
		cond不能被复制(得传地址)

*/

var total int64 = 0
var wg sync.WaitGroup
var lock sync.Mutex // 互斥锁

var num = 0
var rwLock sync.RWMutex // 读写锁

func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		//lock.Lock()
		//total += 1
		//lock.Unlock()

		// 原子化
		atomic.AddInt64(&total, 1)
	}
}
func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		//lock.Lock()
		//total--
		//lock.Unlock()

		// 原子化
		atomic.AddInt64(&total, -1)

	}
}

func addRw() {
	defer wg.Done()
	rwLock.Lock() // 加写锁，写锁会防止别的写锁写值，读锁取值
	defer rwLock.Unlock()
	num = 24
}
func subRw() {
	defer wg.Done()
	rwLock.RLock() // 加读锁，读锁不会阻止其他读锁，会阻止写锁
	defer rwLock.RUnlock()
	fmt.Println(num)
}

// Mutex 自旋锁的简单实现
type Mutex struct {
	state int32
}

const (
	mutexLocked = 1
)

// 自旋条件如下：GOMAXPROCS>1,否则会死锁
func canSpin() bool {
	if runtime.GOMAXPROCS(0) <= 1 {
		return false
	}
	return true
}

// 让协程忙等待
func doSpin() {
	for i, j := 0, 0; i < 30; i++ {
		j = j + 2
	}
	return
}

func (m *Mutex) Lock() {
	//通过原子操作加锁
	if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
		return
	}
	for m.state&mutexLocked == mutexLocked && canSpin() { // 满足自旋条件
		doSpin()
		if atomic.CompareAndSwapInt32(&m.state, 0, mutexLocked) {
			return
		}
	}
	panic("cant lock with one core cpu!")
}

func (m *Mutex) Unlock() {
	//通过原子操作解锁
	newVal := atomic.AddInt32(&m.state, -mutexLocked)
	if newVal == 0 {
		return
	}

	//解锁时倘若发现 Mutex 此前未加锁，直接panic
	if (newVal+mutexLocked)&mutexLocked == 0 {
		panic("unlock of unlocked mutex")
	}
}

var count int
var mu Mutex

func increment() {
	mu.Lock()
	count++
	mu.Unlock()
	wg.Done()
}

func main() {
	//wg.Add(2)
	//go add()
	//go sub()
	//wg.Wait()
	//fmt.Println("total", total)

	//go addRw()
	//time.Sleep(time.Second)
	//go subRw()
	//wg.Wait()

	// 自旋锁
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment()
	}
	wg.Wait()
	fmt.Println("Final count:", count)

}
