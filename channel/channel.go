package main

import (
	"fmt"
	"time"
)

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}
func createWorker(id int) chan<- int { // 加一個 <- 代表只能傳資料給chan
	c := make(chan int)
	go worker(id, c)
	return c
}

// DeadLockChanDemo
// channel 是拿來給 goroutine 間溝通的，必須透過 goroutine 收/放，否則會deadlock
func DeadLockChanDemo() {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
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

// buffered channel : channel with size
func bufferedChannel() {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd' // size=3 傳入4個? 只要取的速度夠快，保持不要超過size都可以運行
	time.Sleep(time.Millisecond)
}

// channel close後, goroutine的無限迴圈仍然持續向channel要數據,
// 但因為已經close了所以只能收到空數據, int default 0
func channelClose() {
	c := make(chan int)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}
func main() {
	//DeadLockChanDemo()
	//bufferedChannel()
	channelClose()
}
