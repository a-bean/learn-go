package main

import (
	"fmt"
	"sync"
)

// 定义回调类型
type Callback func(msg string)

// MessageBroker 中介，负责存储主题和对应的订阅者
type MessageBroker struct {
	subscribers map[string][]Callback
	mu          sync.RWMutex
}

// NewMessageBroker 构造函数，创建一个新的消息中介
func NewMessageBroker() *MessageBroker {
	return &MessageBroker{
		subscribers: make(map[string][]Callback),
	}
}

// Subscribe 订阅主题并注册回调函数
func (b *MessageBroker) On(topic string, callback Callback) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.subscribers[topic] = append(b.subscribers[topic], callback)
}

// Publish 发布消息到主题
func (b *MessageBroker) Publish(topic, message string) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if subs, ok := b.subscribers[topic]; ok {
		for _, callback := range subs {
			// 调用回调函数，传递消息
			callback(message)
		}
	}
}

// Unsubscribe 取消订阅某个主题的回调
func (b *MessageBroker) Unsubscribe(topic string, callback Callback) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if subs, ok := b.subscribers[topic]; ok {
		for i, cb := range subs {
			if fmt.Sprintf("%p", cb) == fmt.Sprintf("%p", callback) {
				// 移除回调
				b.subscribers[topic] = append(subs[:i], subs[i+1:]...)
				break
			}
		}
	}
}

func main() {
	// 创建消息中介
	broker := NewMessageBroker()

	// 订阅者1订阅 "news" 主题，注册回调
	callback1 := func(msg string) {
		fmt.Println("Subscriber 1 received:", msg)
	}
	broker.On("news", callback1)

	// 订阅者2订阅 "news" 主题，注册回调
	callback2 := func(msg string) {
		fmt.Println("Subscriber 2 received:", msg)
	}
	broker.On("news", callback2)

	// 发布者发布消息到 "news" 主题
	broker.Publish("news", "Breaking news!")
	broker.Publish("news", "More news!")

	// 取消订阅者1的订阅
	broker.Unsubscribe("news", callback1)

	// 发布新消息，只有订阅者2会收到
	broker.Publish("news", "News after unsubscription")
}
