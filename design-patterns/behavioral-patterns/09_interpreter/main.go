// 解释器模式: 提供了评估语言的语法或表达式的方式。
package main

import "fmt"

// 表达式接口
type Expression interface {
	Interpret(context string) bool
}

// 终结符表达式
type TerminalExpression struct {
	data string
}

func (t *TerminalExpression) Interpret(context string) bool {
	return context == t.data
}

// 或表达式
type OrExpression struct {
	expr1, expr2 Expression
}

func (o *OrExpression) Interpret(context string) bool {
	return o.expr1.Interpret(context) || o.expr2.Interpret(context)
}

// 且表达式
type AndExpression struct {
	expr1, expr2 Expression
}

func (a *AndExpression) Interpret(context string) bool {
	return a.expr1.Interpret(context) && a.expr2.Interpret(context)
}

func main() {
	robert := &TerminalExpression{data: "Robert"}
	john := &TerminalExpression{data: "John"}

	isMale := &OrExpression{expr1: robert, expr2: john}

	julie := &TerminalExpression{data: "Julie"}
	married := &TerminalExpression{data: "Married"}

	isMarriedWoman := &AndExpression{expr1: julie, expr2: married}

	fmt.Println("John is male?", isMale.Interpret("John"))
	fmt.Println("Julie is a married woman?", isMarriedWoman.Interpret("Married Julie"))
}
