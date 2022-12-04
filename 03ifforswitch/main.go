package main

import "fmt"

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
	switch {
	case expr > 2:
		fmt.Println("1")
	default:
		fmt.Println("default")
	}

}
