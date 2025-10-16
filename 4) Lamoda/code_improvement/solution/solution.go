package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

//1) Что можно улучшить? 		//WaitGroup
//2) Улучшить код
//3) Исправить программу так, чтобы как только какая-нибуть горутина ответила с ошибкой, то программа завершилась
//4) Какие могут быть проблемы если в urls будет 100k урлов
//		1) Исчерпание лимита файловых дескрипторов
//		2) Нехватка доступных портов для исходящих соединений
//5) Как это можно обойти			//Worker pool; SetLimit в errgroup
//6) Написать самое простое решение			//SetLimit в errgroup

func main() {
	urls := []string{
		"https://www.lamoda.ru",
		"https://www.yandex.ru",
		"https://www.mail.ru",
		"https://www.google.com",
	}
	//wg := sync.WaitGroup{}
	g, ctx := errgroup.WithContext(context.Background())
	g.SetLimit(1)

	for _, url := range urls {
		//wg.Add(1)
		g.Go(func() error {
			//defer wg.Done()
			if ctx.Err() != nil {
				return ctx.Err()
			}
			fmt.Printf("Fetching %s...\n", url)

			err := fetchUrl(url)
			if err != nil {
				return fmt.Errorf("error fetching %s: %v", url, err)
			}
			fmt.Printf("Fetched %s\n", url)
			return nil
		})
	}

	fmt.Println("All requests launched!")
	//wg.Wait()
	if err := g.Wait(); err != nil {
		fmt.Printf("Program finished with error: %v\n", err)
		return
	}
	fmt.Println("Program finished successfully.")

}

func fetchUrl(url string) error {
	// Подробная реализация опущена и не относится к теме задачи
	time.Sleep(2 * time.Second)
	_, err := http.Get(url)
	return err
}
