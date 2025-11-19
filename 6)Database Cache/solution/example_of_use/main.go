package main

import (
	cache "example_of_use/main.go"
	"fmt"
	"time"
)

type RealDatabase struct {
	// Здесь могут быть подключения к БД, конфигурация и т.д.
}

func (r *RealDatabase) Get(key string) (string, error) {
	time.Sleep(300 * time.Millisecond)
	fmt.Printf("Запрос к реальной БД для ключа: %s\n", key)

	// В реальности здесь был бы запрос к БД
	data := map[string]string{
		"user:1": "Алексей",
		"user:2": "Мария",
		"user:3": "Иван",
		"user:4": "Саша",
		"user:5": "Таня",
	}
	if value, exists := data[key]; exists {
		return value, nil
	}
	return "", fmt.Errorf("ключ не найден: %s", key)

}
func (r *RealDatabase) Keys() ([]string, error) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Запрос всех ключей к реальной БД")
	return []string{"user:1", "user:2", "user:3", "user:4", "user:5"}, nil
}
func (r *RealDatabase) MGet(keys []string) ([]string, error) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Запрос значения по ключам к реальной БД")
	return []string{"1", "2", "3"}, nil
}

func main() {
	realDB := &RealDatabase{}

	var db cache.KVDatabase = cache.NewCache(realDB)

	// Первый запрос - пойдет в БД
	start := time.Now()
	user1, _ := db.Get("user:1")
	fmt.Printf("Результат: %s, Время: %v\n", user1, time.Since(start))

	// Второй запрос - возьмет из кэша (быстро)
	start = time.Now()
	user1, _ = db.Get("user:1")
	fmt.Printf("Результат: %s, Время: %v\n", user1, time.Since(start))

}
