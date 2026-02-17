package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

/*
Напишите worker pool. Есть n-количество горутин (воркер), которые обрабатывают таски,
если горутина выполнила таску, то она не закрывается, а ждет новую.
Учесть код, который будет закрывать wоrkег pool, при этом все таски должны быть выполнены до полного закрытия
*/
type WorkerPool struct {
	isClosed bool
	mu       sync.Mutex
	taskCh   chan func()
	doneCh   chan struct{}
}

func New(workerNum int) *WorkerPool {
	wp := &WorkerPool{
		isClosed: false,
		mu:       sync.Mutex{},
		taskCh:   make(chan func(), 3), //Можно не буферизированный, тогда не будет запаса в 3 задач если все воркеры заняты
		doneCh:   make(chan struct{}),
	}

	wgStartup := sync.WaitGroup{}
	wgStartup.Add(workerNum)

	wg := sync.WaitGroup{}
	wg.Add(workerNum)

	for i := 0; i < workerNum; i++ {
		go func() {
			wgStartup.Done()
			defer wg.Done()
			for task := range wp.taskCh {
				task()
			}
		}()
	}

	go func() {
		wg.Wait()
		close(wp.doneCh)
	}()
	wgStartup.Wait()
	return wp
}

func (wp *WorkerPool) AddTask(task func()) error {
	wp.mu.Lock()
	if wp.isClosed {
		return errors.New("Worker pool is closed")
	}
	wp.mu.Unlock()

	select {
	case wp.taskCh <- task:
		return nil
	default:
		return errors.New("Worker pool is full")
	}
}
func (wp *WorkerPool) Close() error {
	wp.mu.Lock()
	if wp.isClosed {
		return errors.New("Worker pool is closed")
	}
	wp.isClosed = true
	wp.mu.Unlock()

	close(wp.taskCh)
	<-wp.doneCh
	return nil
}

func main() { //пример реализации (не по заданию)
	pool := New(2)
	fmt.Println("START")
	for i := 0; i < 6; i++ {
		fmt.Println("Added")
		if err := pool.AddTask(Message); err != nil {
			fmt.Println(err)
		}
	}
	pool.Close()

	if err := pool.AddTask(Message); err != nil {
		fmt.Println(err)
	}
}

func Message() { //тест функция
	time.Sleep(time.Second * 3)
	fmt.Println("Message")
}
