package fn

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	re := add(1, 3)
	if re != 4 {
		t.Errorf("expect %d,actual %d", 4, re)
	}
}

func TestAdd2(t *testing.T) {
	fmt.Println("运行到")
	if testing.Short() {
		t.Skip("short 模式下跳过。。。")
	}
	fmt.Println("没运行到")
	re := add(1, 4)
	if re != 5 {
		t.Errorf("expect %d,actual %d", 5, re)
	}
}

// 表格驱动测试
func TestAdd3(t *testing.T) {
	type data struct {
		a   int
		b   int
		out int
	}
	var dataset = []data{
		{1, 2, 3},
		{3, 4, 7},
		{6, 4, 10},
		{-1, 2, 1},
	}

	for _, value := range dataset {
		res := add(value.a, value.b)
		if res != value.out {
			t.Errorf("expect %d,actual %d", value.out, res)
		}
	}
}

// 性能测试
func BenchmarkAdd(b *testing.B) {
	a := 123
	c := 456
	d := 579

	for i := 0; i < b.N; i++ {
		if actual := add(a, c); actual != d {
			fmt.Printf("expect %d,actual %d", d, actual)
		}
	}
}

func BenchmarkAdd2(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < 10000; j++ {
			str = fmt.Sprintf("%s%d", str, j)
		}
	}
	b.StopTimer()
}

func BenchmarkAdd3(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < 10000; j++ {
			str += strconv.Itoa(j)
		}
	}
	b.StopTimer()
}

func BenchmarkAdd4(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		for j := 0; j < 10000; j++ {
			builder.WriteString(strconv.Itoa(j))
		}
		_ = builder.String()
	}
	b.StopTimer()
}
