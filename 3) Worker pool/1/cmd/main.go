package main

import (
	workerpool "cmd/main.go/worker-pool"
	"context"
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
	Stats()
}

var maxMessage = 10
var messagesCounter = 0

func getMessage() []Message {
	messagesCount := rand.Intn(maxMessage)
	messages := make([]Message, 0, messagesCount)

	for range messagesCount {
		messagesCounter++
		messages = append(messages, Message{Id: messagesCounter})
	}
	return messages
}
func proceessMessage(workerId int, msg Message) {
	time.Sleep(300 * time.Millisecond)
	fmt.Printf("Worker %d ptocessed message %d\n", workerId, msg.Id)
}
func main() {
	var pool IPool = workerpool.New(proceessMessage)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

l:
	for {
		pool.Create()
		select {
		case <-ctx.Done():
			break l
		default:
		}
		messages := getMessage()
		fmt.Println("\nMessages count:", messagesCounter)

		for _, message := range messages {
			pool.Handle(message)
		}
		pool.Wait()
	}

	pool.Stats()
	fmt.Println("Messages count:", messagesCounter)
}
