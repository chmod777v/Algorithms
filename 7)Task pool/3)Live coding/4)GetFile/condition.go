package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"time"
)

// 1) Сколько будет выполнятся код?
// 2) Оптимизировать код (Функцию GetFile трогать нельзя)
func main() {
	start := time.Now()
	m, err := GetFiles(context.TODO(), "1", "2", "3", "4", "5")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(m)
	fmt.Println(time.Since(start))
}

// GetFiles пример функции, которую нужно оптимизировать.
func GetFiles(ctx context.Context, names ...string) (result map[string][]byte, err error) {
	if len(names) == 0 {
		return nil, nil
	}

	result = make(map[string][]byte, len(names))
	for _, name := range names {
		result[name], err = GetFile(ctx, name)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// GetFile является примером функции, которая относительно
// недолго выполняется при единичном вызове. Но достаточно
// долго если вызывать последовательно.
// Предположим, что это сторонняя функция, которую мы не можем оптимизировать.
func GetFile(ctx context.Context, name string) ([]byte, error) {
	if name == "" {
		return nil, fmt.Errorf("invalid name %q", name)
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-ticker.C:
	}

	if name == "invalid" {
		return nil, fmt.Errorf("invalid name %q", name)
	}

	b := make([]byte, 10)
	n, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("getting file %q: %w", name, err)
	}

	return b[:n], nil
}
