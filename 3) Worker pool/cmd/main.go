package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	Id    int
	Title string
	Text  string
}

type IPool interface {
	Create()
	Handle(Message)
	Wait()
}

var maxMessage = 10
var messagesCounter = 0

func getMessage() []Message {
	messagesCount := rand.Intn(maxMessage)
	messages := make([]Message, messagesCount)

	for range messagesCount {
		messagesCounter++
		messages = append(messages, Message{Id: messagesCounter})
	}
	return messages
}
func proceessMessage(workerId int, msg Message) {
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Worker %d ptocessed message %d", workerId, msg.Id)
}
func main() {
	var pool IPool
	pool = New(proceessMessage)
	for {
		messages := getMessage()

		pool.Create()

		for _, message := range messages {
			pool.Handle(message)
		}

		pool.Wait()
	}
}
