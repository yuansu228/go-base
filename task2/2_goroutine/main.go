package main

import (
	"fmt"
	"time"
)

// task2.1
func printOdd() {
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			time.Sleep(time.Second)
			fmt.Println(i)
		}
	}
}

func printEven() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			time.Sleep(time.Second)
			fmt.Println(i)
		}
	}
}

// task2.2
func taskSchedule(tasks []func(int)) {
	for k, v := range tasks {
		go func(id int, task func(int)) {
			start := time.Now()
			task(id)
			cost := time.Since(start)
			fmt.Println("task:", k, "cost:", cost)
		}(k, v)
	}
}

func main() {
	// task2.1
	// go printOdd()
	// go printEven()
	// time.Sleep(10 * time.Second)

	// task2.2
	tasks := []func(int){
		func(i int) {
			time.Sleep(time.Second * 3)
			fmt.Println("task", i)
		},
		func(i int) {
			time.Sleep(time.Second * 3)
			fmt.Println("task", i)
		},
		func(i int) {
			time.Sleep(time.Second * 3)
			fmt.Println("task", i)
		},
	}
	taskSchedule(tasks)
	time.Sleep(time.Second * 5)
}
