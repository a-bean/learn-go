package main

import "fmt"

// 观察者模式: 定义了对象间的一对多依赖关系，当一个对象改变状态时，它的所有依赖者都会收到通知并自动更新。

// Observer 观察者接口
type Observer interface {
	Update(string)
}

// User 具体观察者
type User struct {
	id string
}

func (u *User) Update(message string) {
	fmt.Printf("User %s received message: %s\n", u.id, message)
}

// 主题
type Subject struct {
	observers []Observer
	message   string
}

func (s *Subject) Attach(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *Subject) Notify() {
	for _, observer := range s.observers {
		observer.Update(s.message)
	}
}

func (s *Subject) UpdateMessage(message string) {
	s.message = message
	s.Notify()
}

func main() {
	subject := &Subject{}

	user1 := &User{id: "1"}
	user2 := &User{id: "2"}

	subject.Attach(user1)
	subject.Attach(user2)

	subject.UpdateMessage("Hello, Observers!")
}
