package main

import (
	"fmt"
)

//列印完就往外傳一個 done = true
func doWork(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		// 追加一個goroutine來丟done，這樣就不用等人拿走done才能往done放東西了
		// 因為每個goroutine都是獨立的
		go func() { done <- true }()
	}
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}

	go doWork(id, w.in, w.done)
	return w
}

func chanDemo() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}
	//// 這個寫法會導致兩個迴圈會依序print
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + 1
	}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + 1
	}

	// wait for all of them
	// 上面每個worker都會跑兩次for來丟資料, 所以每個都要收兩次
	// 但這個寫法會deadlock
	// why ?
	// createWorker會創造goroutine，不斷的等資料傳入，以觸發doWork
	// 當第一個for迴圈丟完資料時，每個doWork會往done丟true，但直到下面這個for才有人收done，
	//
	// 而golang channel是阻塞式的，沒有人收done，而第二個for迴圈又開始往同一個work丟資料，
	// 就會觸發doWork，而往done再次丟資料，導致阻塞。
	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}
}

func main() {
	chanDemo()
}
