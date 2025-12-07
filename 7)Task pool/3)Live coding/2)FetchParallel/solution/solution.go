package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Написать функцию, которая запрашивает URL из списка и в случае положительного кода 200 выводит
// в stdout в отдельной строке url: , code:
// В случае ошибки выводит в отдельной строке url: , code:
// Вывод должен быть синхронный(в каком порядке горутины выводят в таком он и должен выводится)
// Функция должна завершаться при отмене контекста.
// Реализовать ограничение количества одновременно запущенных горутин.

func fetchParallel(ctx context.Context, urls []string) {
	const concurrentLimit = 3
	sem := make(chan struct{}, concurrentLimit)
	httpClient := &http.Client{}
	wg := sync.WaitGroup{}
	mu := &sync.Mutex{}

	for _, url := range urls {
		select {
		case <-ctx.Done():
			return
		case sem <- struct{}{}:
		}
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			defer func() { <-sem }()
			req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
			if err != nil {
				fmt.Printf("err:%v\n", err.Error())
				return
			}
			resp, err := httpClient.Do(req)
			if err != nil {
				fmt.Printf("err:%v\n", err.Error())
				return
			}
			defer resp.Body.Close()

			mu.Lock()
			fmt.Printf("url: %s		code: %v\n", url, resp.Status)
			mu.Unlock()
		}(url)
	}
	wg.Wait()
}

func main() {
	urls := []string{
		"https://www.yandex.ru",
		"https://www.mail.ru",
		"https://www.google.com",
		"https://mail.google.com",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	fetchParallel(ctx, urls)
}
