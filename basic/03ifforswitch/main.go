package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// if
	a := 10
	if a < 10 {
		fmt.Println(a)
	} else if a == 10 {
		fmt.Println(a)
	} else {
		fmt.Println(a)
	}

	// for
	for i := 0; i < 10; i++ {

		if i == 2 {
			continue
		}

		if i == 5 {
			break
		}
		fmt.Println("i:", i)
	}

	// 相当于while
	j := 1
	for j < 10 {
		fmt.Println(j)
		j++
	}

	// for range
	for index, value := range "hello" {
		//fmt.Println(index, value)
		fmt.Printf("%d,%c\r\n", index, value)
	}

	// switch语句
	// 1. case表达式可以任意类型
	// 2. case表达式可以多个值
	// 3. switch可以是表达式或者函数调用
	// 4. default分支没有前后顺序的要求
	// 5. 常量case表达式不能重复,但是布尔常量除外(编辑器会提示,但是编译可以通过)
	// 6. switch的默认缺省值为true
	// 7. switch跟case的数据类型必须保持一致
	// 8. switch的case分支只会命中一个,如果需要继续执行其他分支，可以使用fallthrough关键字,
	//	  fallthrough关键字不能出现在最后一个case分支,
	//	  fallthrough关键字一定要在case分支代码块的最后一行
	// 9. case分支的变量作用域为当前case分支的代码块内

	// 第一种
	expr := 5
	switch expr {
	case 1:
		fmt.Println("1")
	case 2, 3:
		fmt.Println("2,3")
	default:
		fmt.Println("default")
	}

	// 第二种
	switch expr + 4 {
	case 4:
		fmt.Println("1")
	default:
		fmt.Println("default")
	}

	// 第三,四种
	switch time.Now().Weekday() {
	default:
		fmt.Println("default")
	case time.Sunday, time.Saturday:
		fmt.Println("1")
	}

	// 第五种
	switch false {
	case false: // 这种情况只会执行第一个,因为break了
		fmt.Println("false")
	//case false:
	//	fmt.Println("false")
	default:
		fmt.Println("default")
	}

	// 第六种
	switch {
	case true:
		fmt.Println("1")
	default:
		fmt.Println("default")
	}

	// 第七种
	switch n := rand.Intn(10); n {
	//case n > 1: 编译无法通过
	//	fmt.Println("default")
	default:
		fmt.Println("default")
	}

	// 第八种
	switch {
	case true:
		fmt.Println("fallthrough1")
		fallthrough // 会继续之前下面命中的case
	default:
		fmt.Println("fallthrough")
	}

}
