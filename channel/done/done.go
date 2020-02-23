package main

import (
	"fmt"
	"sync"
)

//列印完就往外傳一個 done = true
func doWork(id int, c chan int, w worker) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		w.done()
	}
}

type worker struct {
	in   chan int
	done func()
}

// 原本讓doWork決定done要做什麼，改讓createWork做決定。
// 而doWork就只是把事情做完了，就讓worker done()，讓done這件事情抽象化。
func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in:   make(chan int),
		done: func() { wg.Done() }, // 這裡要用func(){} 包起來，決定這個done裡面到底做什麼。
	}

	go doWork(id, w.in, w)
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
