package main

import (
	"fmt"
	"math/rand"
	"time"
)

// task 4.1
func send(ch chan int) {
	for i := 1; i <= 10; i++ {
		fmt.Println("发送数字：", i)
		ch <- i
	}
}

func receive(ch chan int) {
	for v := range ch {
		fmt.Println("收到数字：", v)
	}
}

// task 4.2
func producer(ch chan int) {
	for i := 1; i <= 100; i++ {
		randomInt := rand.Intn(100)
		ch <- randomInt
		fmt.Printf("发送第%d个整数：%d\n", i, randomInt)
	}
}

func consumer(ch chan int) {
	i := 1
	for v := range ch {
		fmt.Printf("收到第%d个整数：%d\n", i, v)
		i++
	}
}

func main() {
	// task4.1
	// ch := make(chan int)
	// go send(ch)
	// go receive(ch)
	// time.Sleep(time.Second)

	// 4.2
	ch := make(chan int, 2)
	go producer(ch)
	go consumer(ch)
	time.Sleep(500 * time.Millisecond)

}
