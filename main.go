package main

import (
	"container/list"

	deque "github.com/edwingeng/deque/v2"
	deque2 "github.com/gammazero/deque"
)

type Message struct {
	id      int
	message string
}

func getMessage(i int) Message {
	return Message{id: i, message: "This is an example message."}
}

func linkedListBuffer(num int) {
	buffer := list.New()
	for i := 0; i < num; i++ {
		message := getMessage(i)
		buffer.PushFront(message)
	}
	for {
		message := buffer.Back()
		_ = message // for parity with other tests
		if message == nil {
			break
		}
		buffer.Remove(message)
	}
}

func channelBufferSelect(num int) {
	buffer := make(chan Message, num)
	for i := 0; i < num; i++ {
		message := getMessage(i)
		buffer <- message
	}
out:
	for {
		select {
		case message := <-buffer:
			_ = message
		default:
			break out
		}
	}
}

func channelBufferIfLen(num int) {
	buffer := make(chan Message, num)
	for i := 0; i < num; i++ {
		message := getMessage(i)
		buffer <- message
	}
	for {
		if len(buffer) == 0 {
			break
		}
		message := <-buffer
		_ = message // avoid unused identifier
	}
}

func sliceBuffer(num int) {
	var buffer []Message
	for i := 0; i < num; i++ {
		message := getMessage(i)
		buffer = append(buffer, message)
	}
	for {
		if len(buffer) == 0 {
			break
		}
		message := buffer[0]
		_ = message // for parity with other tests
		buffer = buffer[1:]
	}
}

func sliceBufferMake(num int) {
	buffer := make([]Message, 1000)
	for i := 0; i < num; i++ {
		message := getMessage(i)
		buffer = append(buffer, message)
	}
	for {
		if len(buffer) == 0 {
			break
		}
		message := buffer[0]
		_ = message // for parity with other tests
		buffer = buffer[1:]
	}
}

func dequeBuffer(num int) {
	buffer := deque.NewDeque[Message]()
	for i := 0; i < num; i++ {
		message := getMessage(i)
		buffer.PushFront(message)
	}
	for {
		if buffer.Len() == 0 {
			break
		}
		message := buffer.PopBack()
		_ = message
	}
}

func deque2Buffer(num int) {
	buffer := deque2.New[Message]()
	for i := 0; i < num; i++ {
		message := getMessage(i)
		buffer.PushFront(message)
	}
	for {
		if buffer.Len() == 0 {
			break
		}
		message := buffer.PopBack()
		_ = message
	}
}

func main() {
	channelBufferIfLen(10)
}
