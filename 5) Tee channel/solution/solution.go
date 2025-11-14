package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// написать функцию tee, которая будет данные из 1 канала перекладывать в n кол-во других и возвращать масив с этими каналами
// доработать код
// прокинуть контекст
func tee(ctx context.Context, in <-chan int, numChans int) []chan int {
	chans := make([]chan int, numChans)
	for i := range numChans {
		chans[i] = make(chan int)
	}
	go func() {
		for i := range numChans {
			defer close(chans[i])
		}
		for {
			select {
			case <-ctx.Done():
				return
			case val, ok := <-in:
				if !ok {
					return
				}
				wg := &sync.WaitGroup{}
				for i := range numChans {
					wg.Add(1)
					go func() {
						defer wg.Done()
						select {
						case <-ctx.Done():
							return
						case chans[i] <- val:
						}
					}()
				}
				wg.Wait()
			}
		}
	}()
	return chans
}

func generate() chan int {
	ch := make(chan int)

	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func service(chans []chan int) {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func() {
		defer wg.Done()
		for v := range chans[0] {
			fmt.Println("Logging... ", v)
			time.Sleep(time.Second)
		}
	}()
	go func() {
		defer wg.Done()
		for v := range chans[1] {
			fmt.Println("Metrics... ", v)
		}
	}()
	go func() {
		defer wg.Done()
		for v := range chans[2] {
			fmt.Println("Sending... ", v)
		}
	}()
	wg.Wait()
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	chans := tee(ctx, generate(), 3)
	service(chans) //имитация передачи масива каналов сервисам
}
