package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int){
	for {
		fmt.Printf("Worker %d received %d\n",id, <-c)
	}
}

// DeadLockChanDemo
// channel 是拿來給 goroutine 間溝通的，必須透過 goroutine 收/放，否則會deadlock
func DeadLockChanDemo() {
	c := make(chan int)
	go worker(0,c)
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}
func main(){
	DeadLockChanDemo()
}
