package main

import (
	"context"
	"time"
)

// Написать функцию, которая запрашивает URL из списка и в случае положительного кода 200 выводит
// в stdout в отдельной строке url: , code:
// В случае ошибки выводит в отдельной строке url: , code:
// Функция должна завершаться при отмене контекста.
// Доп. задание: реализовать ограничение количества одновременно запущенных горутин.

func fetchParallel(ctx context.Context, urls []string) {
	//напиши код здесь
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

	fetchParallel(ctx, urls)

}
