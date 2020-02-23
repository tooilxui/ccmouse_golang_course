package main

import (
	"fmt"
	"sync"
)

//列印完就往外傳一個 done = true
func doWork(id int, c chan int, wg *sync.WaitGroup) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		wg.Done()
	}
}

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}

	go doWork(id, w.in, wg)
	return w
}

func chanDemo() {
	var wg sync.WaitGroup
	wg.Add(20)

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	for _, worker := range workers {
		worker.in <- 'a' + 1
	}
	for _, worker := range workers {
		worker.in <- 'A' + 1
	}
	wg.Wait()
}

func main() {
	chanDemo()
}
