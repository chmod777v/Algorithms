package main

import (
	"fmt"
	"time"
)

type Vehicle interface {
	Start()
	Stop()
	Time() time.Duration
}

type Car struct {
	isMoving  bool
	startTime time.Time
}

func (c *Car) Start() {
	if c.isMoving {
		fmt.Println("Машина уже в движении")
		return
	}

	c.isMoving = true
	c.startTime = time.Now()
	fmt.Println("Машина начала движение")

}
func (c *Car) Stop() {
	if !c.isMoving {
		fmt.Println("Уже остановлена")
		return
	}
	elapsed := time.Since(c.startTime)
	c.isMoving = false

	fmt.Printf("Машина остановлена. Общее время в пути: %v\n", elapsed)
}
func (c *Car) Time() time.Duration {
	if c.isMoving {
		return time.Since(c.startTime)
	}
	return 0
}

func main() {
	car := &Car{}
	var vehicle Vehicle = car

	vehicle.Start()
	time.Sleep(time.Second * 5)
	fmt.Println(vehicle.Time())
	time.Sleep(time.Second * 3)
	vehicle.Stop()

}
