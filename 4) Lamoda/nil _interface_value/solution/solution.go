package main

import "fmt"

//1) Что выведет?	//Вывод: oops т.к.:
//		return err возвращает интерфейс error, который содержит:
//				*Тип: MyError (конкретный тип)
//				Значение: nil (нулевое значение для этого типа)
//		Интерфейс nil тогда когда и тип nil и значение nil, в данном примере тип не nil а MyError => интерфейс тоже не nil

//2) Что представляет из себя интерфейс	(кратко)?
//		Интерфейс внутри - это структура, которая содержит два указателя: 1) на таблицу методов и тип 2) на конкретное значение в памяти

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
	if i > 5 {
		return &MyError{data: "i > 5"}
	}
	return nil
}
