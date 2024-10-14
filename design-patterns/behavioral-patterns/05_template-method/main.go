package main

import "fmt"

// 抽象类
type Game interface {
	initialize()
	startPlay()
	endPlay()
}

// 定义一个模板结构体
type GameTemplate struct {
	Game
}

// 模板方法
func (g GameTemplate) Play() {
	g.initialize()
	g.startPlay()
	g.endPlay()
}

// 具体类
type Cricket struct{}

func (c Cricket) initialize() {
	fmt.Println("Cricket Game Initialized!")
}

func (c Cricket) startPlay() {
	fmt.Println("Cricket Game Started. Enjoy the game!")
}

func (c Cricket) endPlay() {
	fmt.Println("Cricket Game Finished!")
}

type Football struct{}

func (f Football) initialize() {
	fmt.Println("Football Game Initialized!")
}

func (f Football) startPlay() {
	fmt.Println("Football Game Started. Enjoy the game!")
}

func (f Football) endPlay() {
	fmt.Println("Football Game Finished!")
}

func main() {
	cricketGame := GameTemplate{Game: Cricket{}}
	footballGame := GameTemplate{Game: Football{}}

	cricketGame.Play()  // 输出 Cricket Game
	footballGame.Play() // 输出 Football Game
}
