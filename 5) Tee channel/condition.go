package main

import (
	"fmt"
	"time"
)

//написать функцию tee, которая будет данные из 1 канала перекладывать в n-ое кол-во других и возвращать масив с этими каналами
//доработать код
//прокинуть контекст

func tee(in <-chan int, numChans int) []chan int {

	return nil
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
	go func() {
		for v := range chans[0] {
			fmt.Println("Logging... ", v)
			time.Sleep(time.Second)
		}
	}()
	go func() {
		for v := range chans[1] {
			fmt.Println("Metrics... ", v)
		}
	}()
	go func() {
		for v := range chans[2] {
			fmt.Println("Sending... ", v)
		}
	}()
}
func main() {
	chans := tee(generate(), 3)

	service(chans) //имитация передачи масива каналов сервисам
	time.Sleep(time.Second * 10)
}
