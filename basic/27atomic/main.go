package main

import (
	"fmt"
	"sync/atomic"
)

/*
sync/atomic 包提供了底层的原子操作，用于在多线程（goroutine）并发场景下安全地操作共享变量，而无需显式使用锁（如 sync.Mutex）。原子操作是由硬件支持的单指令操作，具有高效性和原子性（不可分割性），非常适合高并发环境下的简单变量操作。
*/
func main() {

	/*
		支持以下基本类型的原子操作
				int32
				int64
				uint32
				uint64
				uintptr
				unsafe.Pointer
				它还支持一种特殊类型 Value，用于存储任意类型的值（但需注意类型一致性）。

		原子操作的基本函数有：
			读取（Load）：以原子方式读取变量的值。
			存储（Store）：以原子方式写入变量的值。
			加法（Add）：以原子方式对变量执行加减操作。
			比较并交换（CompareAndSwap, CAS）：比较变量当前值与预期值，若相等则替换为新值。
			交换（Swap）：以原子方式将变量替换为新值，并返回旧值。
	*/

	// 如果需要保证线程安全，必须通过指针引用变量进行原子操作。直接使用普通变量会导致数据竞争
	var i1 int32 = 10
	newValue := atomic.AddInt32(&i1, 1)                // 原子加1
	fmt.Println(atomic.LoadInt32(&i1))                 // 原子读取
	fmt.Println(i1, newValue)                          // 11 11
	atomic.StoreInt32(&i1, 10)                         // 原子设置
	fmt.Println(atomic.LoadInt32(&i1))                 // 原子读取
	swapped := atomic.CompareAndSwapInt32(&i1, 10, 20) // 原子比较交换 旧的值是10，才会替换成功 swapped 才会是true 否则是false
	fmt.Println(swapped)                               // true
	old := atomic.SwapInt32(&i1, 30)                   // 原子交换 交换成功后 i1 的值是30，old 是10
	fmt.Println(old, i1)                               // 20

	var v atomic.Value
	v.Store("hello")      // 原子存储
	fmt.Println(v.Load()) // 原子读取
	// v.Store(100)        // 会报错  类型不一致
	swapped = v.CompareAndSwap("hello", "world") // 原子比较交换
	fmt.Println(swapped)                         // true

	// uintptr
	var p uintptr = 0
	atomic.StoreUintptr(&p, 100)               // 原子存储
	fmt.Println(atomic.LoadUintptr(&p))        // 原子读取
	atomic.CompareAndSwapUintptr(&p, 100, 200) // 原子比较交换

}
