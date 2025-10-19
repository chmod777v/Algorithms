package main

import (
	"fmt"
	"time"
)

// 1) Что выведет?   // Паника: panic: send on closed channel
//		Создаем небуферизованный канал
//		В горутине блокируемся, ждем получателя
//		Закрываем канал и т.к. горутина заблокирована и пишет в канал происходит запись в закрытый канал => паника

//  2. Исправить!
//     Просто создаем буфер в канале чтобы горутина не блочилась
//     Но лучше всего закрывать канал в горутине	(хорошая практика: закрывает тот кто пишет)

func main() {
	ch := make(chan int) //	make(chan int, 1)
	go func() {
		ch <- 1
		close(ch)
	}()
	time.Sleep(time.Millisecond * 500)

	for i := range ch {
		fmt.Println(i)
	}

	time.Sleep(time.Millisecond * 100)
}
