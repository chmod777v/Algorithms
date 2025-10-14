package main

import "fmt"

func foo(src []int) {
	src = append(src, 5) //arr=[1,5,3]
}
func foo2(src *[]int) {
	*src = append(*src, 5)
}

func main() {
	arr := []int{1, 2, 3} //[1,2,3] len=3 cap=3
	//src := arr[:1]        //[1] len=1 cap=3

	src := make([]int, 1)
	copy(src, arr)

	foo2(&src)
	//foo(src)
	fmt.Println(src) //[1] т.к. хоть 2 элемент изменился, но длина осталась прежней
	fmt.Println(arr) //[1,5,3]

}
