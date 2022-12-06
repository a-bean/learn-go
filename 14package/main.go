package main

// 可以定义别名
import (
	"fmt"
	User "learn-go/14package/user"
)

// 引入不用
import (
	_ "learn-go/14package/user"
)

func main() {
	// package 用来组织源码，是多个go源码的集合，代码复用的基础
	// 每个源文件都必须申请所属的package
	// package的名称可以不跟所属文件夹名称一致
	// 同一个文件夹下的多个文件package名称都要一致
	// 同一个文件夹下的多个文件可以互相访问（前提是：变量大写开头）
	c := User.Course{
		Name: "math",
	}
	fmt.Println(User.GetCourse(c))
}
