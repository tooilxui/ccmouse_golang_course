package main

import (
	"fmt"
)

//列印完就往外傳一個 done = true
func doWork(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		done <- true
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
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + 1
	}
	// 分兩個迴圈放跟取也可以，這樣比較單純 (doWork就不用開goroutine了)
	for _, worker := range workers {
		<-worker.done
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + 1
	}
	// 分兩個迴圈放跟取也可以，這樣比較單純 (doWork就不用開goroutine了)
	for _, worker := range workers {
		<-worker.done
	}
}

func main() {
	chanDemo()
}
