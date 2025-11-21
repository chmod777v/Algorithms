package main

import (
	"fmt"
)

// сколько тут проблем?
func FindMaxProblem() {
	var maxNum int
	for i := 1000; i > 0; i-- {
		go func() {
			if i%2 == 0 && i > maxNum {
				maxNum = i
			}
		}()
	}
	fmt.Printf("Maximum is %d", maxNum)
}
func main() {
	FindMaxProblem()
}
