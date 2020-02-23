package main

import (
	"fmt"
	"time"
)

// DeadLockChanDemo
// channel 是拿來給 goroutine 間溝通的，必須透過 goroutine 收/放，否則會deadlock
func DeadLockChanDemo() {
	c := make(chan int)
	go func() {
		for {
			// 收1印1、收2印2，但來不及印出2就結束程式了，先加一個sleep讓他順利印完
			n:= <-c
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}
func main(){
	DeadLockChanDemo()
}
