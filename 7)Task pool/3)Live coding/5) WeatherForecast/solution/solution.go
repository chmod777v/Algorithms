package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// Имеется сторонний сервис погоды (его имитация - это функция WeatherForecast).
// Сторонний сервис работает за секунду, что для нас долго.
// На наш сервис идет большая нагрузка. Как доработать текущую реализацию?
// 1. Предложить и реализовать решение.
// 2. Дополнительное задание: сторонний сервис может давать данные не только по одному городу.
// Доработать реализацию из первого пункта с учетом этого факта.
type Data struct {
	Temperatures map[string]int
	mu           sync.RWMutex
}

func (d *Data) UpdateTemperature() {
	wg := &sync.WaitGroup{}
	for city := range d.Temperatures {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tmp := WeatherForecast(city)

			d.mu.Lock()
			d.Temperatures[city] = tmp
			d.mu.Unlock()
		}()
	}
	wg.Wait()
}

func (d *Data) GetTemperature(city string) (int, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	t, ok := d.Temperatures[city]
	if !ok {
		return 0, fmt.Errorf("City %s not found", city)
	}
	return t, nil

}

func NewData(interval time.Duration) *Data {
	ticker := time.NewTicker(interval)

	newData := &Data{}
	go func() {
		defer ticker.Stop()
		for range ticker.C {
			newData.UpdateTemperature()
		}
	}()
	return newData
}

func WeatherForecast(city string) int {
	time.Sleep(1 * time.Second)
	return rand.Intn(70) - 30
}

func main() {
	data := NewData(time.Second * 5)
	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		temp, err := data.GetTemperature("Moscow")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNoContent)
		}
		fmt.Fprintf(w, "{\"temperature\":%d} \n", temp)
	})
	if err := http.ListenAndServe(":3333", nil); err != nil {
		panic(err)
	}
}
