package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// условие задачи:
// реализовать функцию processParallel
// прокинуть контекст

func processData(ctx context.Context, v int) (int, error) {
	ch := make(chan struct{})
	go func() {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		close(ch)
	}()

	select {
	case <-ch:
		return v * 2, nil
	case <-ctx.Done():
		return 0, errors.New("Timeout")
	}

}

func main() {
	in := make(chan int)
	out := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	go func() {
		//time.Sleep(time.Second * 10)
		defer close(in)
		for i := range 10 {
			select {
			case in <- i:
			case <-ctx.Done():
				return
			}
		}

	}()

	start := time.Now()
	processParallel(ctx, in, out, 5)

	for v := range out {
		fmt.Println("v =", v)
	}
	fmt.Println("main duration:", time.Since(start))
}

func processParallel(ctx context.Context, in <-chan int, out chan<- int, numWorkers int) {
	wg := &sync.WaitGroup{}

	wg.Add(numWorkers)
	for range numWorkers {
		go worker(wg, ctx, in, out)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
}

func worker(wg *sync.WaitGroup, ctx context.Context, in <-chan int, out chan<- int) {
	defer wg.Done()
	for {
		select {
		case v, ok := <-in:
			if !ok { //канал закрыт и больше нет данных для чтения
				return
			}
			val, err := processData(ctx, v)
			if err != nil {
				return
			}
			select {
			case <-ctx.Done():
				return
			case out <- val:
			}
		case <-ctx.Done():
			return
		}
	}
}
