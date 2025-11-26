package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Написать функцию, которая запрашивает URL из списка и в случае положительного кода 200 выводит
// в stdout в отдельной строке url: , code:
// В случае ошибки выводит в отдельной строке url: , code:
// Функция должна завершаться при отмене контекста.
// Доп. задание: реализовать ограничение количества одновременно запущенных горутин.

func fetchParallel(ctx context.Context, urls []string) {
	httpClient := &http.Client{}
	for _, url := range urls {
		req, _ := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
		resp, _ := httpClient.Do(req)
		fmt.Println(resp.Body)
	}
}

func main() {
	urls := []string{
		"https://www.yandex.ru",
		"https://www.maill.ru",
		"https://www.google.com",
		"https://mail.google.com",
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	httpClient := &http.Client{}
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, "https://www.google.com", nil)
	resp, _ := httpClient.Do(req)
	defer resp.Body.Close()
	val, _ := io.ReadAll(resp.Body)
	fmt.Printf("%s", val)

	//fetchParallel(ctx, urls)

}
