package main

import (
	"fmt"
	"net/http"
	"time"
)

//1) Что можно улучшить?
//2) Улучшить код
//3) Исправить программу так, чтобы как только какая-нибуть горутина ответила с ошибкой, то программа завершилась
//4) Какие могут быть проблемы если в urls будет 100k урлов
//5) Как это можно обойти
//6) Написать самое простое решение

func main() {
	urls := []string{
		"https://www.lamoda.ru",
		"https://www.yandex.ru",
		"https://www.maill.ru",
		"https://www.google.com",
	}

	for _, url := range urls {
		go func(url string) {
			fmt.Printf("Fetching %s...\n", url)

			err := fetchUrl(url)
			if err != nil {
				fmt.Printf("Error fetching %s: %v\n", url, err)
				return
			}
			fmt.Printf("Fetched %s\n", url)
		}(url)
	}

	fmt.Println("All requests launched!")
	time.Sleep(400 * time.Millisecond)
	fmt.Println("Program finished.")
}

func fetchUrl(url string) error {
	// Подробная реализация опущена и не относится к теме задачи
	_, err := http.Get(url)
	return err
}
