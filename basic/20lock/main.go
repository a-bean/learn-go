package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// lock 解决资源竞争问题。本质是将并行的代码串行化了
// 使用lock会影响性能，即使设计锁，也要尽量保持并行

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

func main() {
	wg.Add(2)
	//go add()
	//go sub()
	//wg.Wait()
	//fmt.Println("total", total)

	go addRw()
	time.Sleep(time.Second)
	go subRw()
	wg.Wait()
}
