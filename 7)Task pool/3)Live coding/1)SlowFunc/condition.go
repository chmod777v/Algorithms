package main

import (
	"fmt"
	"time"
)

// есть некая медленная функция (SlowFunc), и нужно написать для нее обертку,
// которая добавляет поддержку отмены через контекст, не изменяя саму функцию.

func SlowFunc() string {
	time.Sleep(time.Second * 5)
	return "Hello world"
}

func main() {
	fmt.Println(SlowFunc())

}
