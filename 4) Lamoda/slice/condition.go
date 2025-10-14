package main

import "fmt"

//1) Что выведет?
//2) Исправить, чтобы работало как доложно работать на первый взгляд!

func foo(src []int) {
	src = append(src, 5)
}

func main() {
	arr := []int{1, 2, 3}
	src := arr[:1]

	foo(src)
	fmt.Println(src)
	fmt.Println(arr)

}
