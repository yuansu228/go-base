package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// task5.1
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *SafeCounter) GetCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

// task5.2
type NoLockCounter struct {
	count int32
}

func (c *NoLockCounter) Increment() {
	atomic.AddInt32(&c.count, 1)
}

func (c *NoLockCounter) GetCount() int32 {
	return atomic.LoadInt32(&c.count)
}

func main() {
	// task5.1
	// counter := SafeCounter{}
	// for i := 1; i <= 10; i++ {
	// 	go func() {
	// 		for j := 1; j <= 1000; j++ {
	// 			counter.Increment()
	// 		}
	// 	}()
	// }
	// time.Sleep(time.Second)
	// fmt.Printf("Final Count : %d\n", counter.GetCount())

	// task5.2
	var counter NoLockCounter
	for i := 1; i <= 10; i++ {
		go func() {
			for j := 1; j <= 1000; j++ {
				counter.Increment()
			}
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("Final Count : %d\n", counter.GetCount())
}
