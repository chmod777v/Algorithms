package main

import "fmt"

func f1(s []int) {
	s[1] = 20
	s = append(s, 80)
}
func f2(s *[]int) {
	(*s)[1] = 10
	*s = append(*s, 40)
}

func main() {
	s := []int{1, 2, 3}
	f2(&s)
	f1(s)
	fmt.Println(s)
}
