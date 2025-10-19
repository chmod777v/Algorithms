package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func merge(ch ...<-chan int) <-chan int {
	out := make(chan int)
	wg := sync.WaitGroup{}

	for _, i := range ch {
		wg.Add(1)
		go func(i <-chan int) {
			defer wg.Done()

			for k := range i {
				out <- k
			}
		}(i)
	}

	go func() {
		wg.Wait()
		close((out))
	}()
	return out
}

func source(sourceFunc func(int) int) <-chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- sourceFunc(i)
			time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
		}
	}()

	return ch
}

func main() {
	in1 := source(func(_ int) int {
		return rand.Int()
	})

	in2 := source(func(i int) int {
		return i
	})
	out := merge(in1, in2)

	for val := range out {
		fmt.Println("Value: ", val)
	}
}
