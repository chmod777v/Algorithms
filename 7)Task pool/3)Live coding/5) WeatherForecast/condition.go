package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// Имеется сторонний сервис погоды (его имитация - это функция WeatherForecast).
// Сторонний сервис работает за секунду, что для нас долго.
// На наш сервис идет большая нагрузка. Как доработать текущую реализацию?
// 1. Предложить и реализовать решение.
// 2. Дополнительное задание: сторонний сервис может давать данные не только по одному городу.
// Доработать реализацию из первого пункта с учетом этого факта.

func WeatherForecast() int {
	time.Sleep(1 * time.Second)
	return rand.Intn(70) - 30
}

func main() {
	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "{\"temperature\":%d} \n", WeatherForecast())
	})
	if err := http.ListenAndServe(":3333", nil); err != nil {
		panic(err)
	}
}
