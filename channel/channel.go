package main

import "fmt"

// DeadLockChanDemo
// channel 是拿來給 goroutine 間溝通的，必須透過 goroutine 收/放，否則會deadlock
func DeadLockChanDemo() {
	c := make(chan int)
	c <- 1
	c <- 2
	n := <-c
	fmt.Println(n)
}
func main(){
	DeadLockChanDemo()
}
