package main

import (
	"fmt"
	"strconv"
	"sync"
)

// сколько тут проблем?
func FindMaxProblem() {
	var wg sync.WaitGroup
	ch := make(chan string, 5)
	mu := sync.Mutex{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(ch chan<- string, i int, grp *sync.WaitGroup) {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			msg := fmt.Sprintf("Goroutine %s", strconv.Itoa(i))
			ch <- msg
		}(ch, i, &wg)
	}
	for {
		select {
		case q := <-ch:
			fmt.Println(q)
		}
	}
	wg.Wait()
}

func main() {
	FindMaxProblem()
}
