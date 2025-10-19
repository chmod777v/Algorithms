package main

import "fmt"

//1) Что выведет?
//2) Что представляет из себя интерфейс	(кратко)?
//3) Исправить, чтобы работало как доложно работать на первый взгляд!

type MyError struct {
	data string
}

func (e MyError) Error() string {
	return e.data
}

func main() {
	err := foo(4)
	if err != nil {
		fmt.Println("oops")
	} else {
		fmt.Println("ok")
	}
}

func foo(i int) error {
	var err *MyError
	if i > 5 {
		err = &MyError{data: "i > 5"}
	}
	return err
}
