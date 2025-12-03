package datastructure

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

/*
消息队列是FIFO，生产者和消费者模式
*/

type Message struct {
	ID      string
	Payload []byte
}

type Queue interface {
	Publish([]byte) error
	Subscrible() (<-chan *Message, error)
	Close() error
}

type memoryMessageQueue struct {
	buffer   chan *Message
	isClosed bool
	mu       sync.RWMutex
	cap      int
}

func NewMemoryMessageQueue(cap int) Queue {
	return &memoryMessageQueue{
		buffer: make(chan *Message, cap),
		cap:    cap,
	}
}

// 发送消息（生产者使用）
func (s *memoryMessageQueue) Publish(msgByte []byte) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.isClosed {
		fmt.Println("pub failed because of queue is closed")
		return errors.New("the message queue is closed")
	}

	msg := &Message{
		ID:      fmt.Sprintf("MSG-%v-%d", time.Now().Nanosecond(), rand.Int32()),
		Payload: msgByte,
	}

	s.buffer <- msg
	fmt.Printf("pub msg: %v succeed\n", string(msgByte))
	return nil
}

// 订阅消息（消费者使用）
func (s *memoryMessageQueue) Subscrible() (<-chan *Message, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.isClosed {
		fmt.Println("sub failed because of queue is closed")
		return nil, errors.New("the message queue is closed")
	}

	return s.buffer, nil
}

// 关闭消息队列
func (s *memoryMessageQueue) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.isClosed {
		fmt.Println("sub failed because of queue is closed")
		return errors.New("the message queue is already closed")
	}

	s.isClosed = true
	fmt.Println("the queue is closed now")
	close(s.buffer)
	return nil
}

func HandleMessageQueue() {
	mq := NewMemoryMessageQueue(20)
	defer mq.Close()

	// 模仿2个生产者
	for i := 1; i < 3; i++ {
		go func(pubId int) {
			fmt.Printf("this is the num: %d publisher\n", pubId)
			for j := 0; j < 5; j++ {
				msgText := fmt.Sprintf("the publiser: %d, the msgId is: %d\n", pubId, j)
				if err := mq.Publish([]byte(msgText)); err != nil {
					fmt.Printf("pub msg: %v failed\n", msgText)
					return
				}
				// 延迟300毫秒
				time.Sleep(300 & time.Millisecond)
			}
			fmt.Printf("pub publiser: %d succeed\n", pubId)
		}(i)
	}

	// 模仿3个消费者
	for i := 1; i < 4; i++ {
		go func(subId int) {
			fmt.Printf("the subcribler is: %v\n", subId)
			msgChan, err := mq.Subscrible()
			if err != nil {
				fmt.Printf("sub message failed for subcribler: %v\n", subId)
				return
			}

			for msg := range msgChan {
				fmt.Printf("the subcribler: %v receive: %v\n", subId, string(msg.Payload))
				// 延迟100毫秒
				time.Sleep(100 * time.Microsecond)
			}
			fmt.Printf("the subscribler: %v receive all messag succeed\n", subId)
		}(i)
	}

	time.Sleep(10 * time.Second)
}
