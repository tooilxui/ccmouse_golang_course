package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for {
		fmt.Printf("Worker %d received %d\n", id, <-c)
	}
}

// DeadLockChanDemo
// channel 是拿來給 goroutine 間溝通的，必須透過 goroutine 收/放，否則會deadlock
func DeadLockChanDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}
	// 雖然依序傳入但Print的順序會不一樣，因為每個Print都是一個I/O操作，goroutine 會進行調度。
	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + 1
	}
	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + 1
	}

	time.Sleep(time.Millisecond)
}
func main() {
	DeadLockChanDemo()
}
