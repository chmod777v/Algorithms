package main

import (
	"fmt"
	"math/rand"
	"time"
)

// условие задачи:
// реализовать функцию processParallel
// прокинуть контекст

func processData(v int) int {
	time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	return v * 2
}

func main() {
	in := make(chan int)
	out := make(chan int)

	go func() {
		for i := range 10 {
			in <- i
		}
		close(in)
	}()

	start := time.Now()
	processParallel(in, out, 5)

	for v := range out {
		fmt.Println("v =", v)
	}
	fmt.Println("main duration:", time.Since(start))
}

func processParallel(in, out chan int, numWorkers int) {
}
