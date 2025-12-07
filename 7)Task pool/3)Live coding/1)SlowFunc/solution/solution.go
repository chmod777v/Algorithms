package main

import (
	"context"
	"fmt"
	"time"
)

// есть некая медленная функция (slow func), и нужно написать для нее обертку,
// которая добавляет поддержку отмены через контекст, не изменяя саму медленную функцию.

func SlowFunc() string {
	time.Sleep(time.Second * 3)
	return "Hello world"
}
func ctxFunc(ctx context.Context) (string, bool) {
	ch := make(chan string, 1)
	go func() {
		ch <- SlowFunc()
	}()
	select {
	case v := <-ch:
		return v, true
	case <-ctx.Done():
		return "", false
	}
}
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*4)
	defer cancel()
	if val, ok := ctxFunc(ctx); ok {
		fmt.Println(val)
	}
}
