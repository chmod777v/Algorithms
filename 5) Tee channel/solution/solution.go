package main

import (
	"fmt"
	"sync"
)

func tee(
	in <-chan int,
) (_, _ <-chan int) {
	out1 := make(chan int)
	out2 := make(chan int)
	go func() {
		defer close(out1)
		defer close(out2)
		for val := range in {
			var out1, out2 = out1, out2
			for i := 0; i < 2; i++ {
				select {
				case out1 <- val:
					out1 = nil
				case out2 <- val:
					out2 = nil
				}
			}
		}
	}()
	return out1, out2
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
func main() {
	ch1, ch2 := tee(generate())
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for v := range ch1 {
			fmt.Println("ch1: ", v)
		}
	}()
	go func() {
		defer wg.Done()
		for v := range ch2 {
			fmt.Println("ch2: ", v)
		}
	}()
	wg.Wait()
}
