package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	//Try with a mutex instead of atomic for speed comparison
	//Atomic is faster
	var runningTotal int
	//Operations will have no clashes on threads EVER
	var safeTotal atomic.Int64
	//Similar to using a channel for waiting for something until it is finished
	//Set up a WaitGroup that you set the number of things you are waiting for (This case, 5 times)
	var wg sync.WaitGroup
	var emm sync.Mutex

	wg.Add(5)
	for i := 1; i <= 5; i++ {
		go func() {
			for i := 0; i < 50000; i++ {
				emm.Lock()
				runningTotal++
				emm.Unlock()
				//Cannot do ++ on atomic, must use .Add/.Subtract
				safeTotal.Add(1)
			}
			//++
			wg.Done()
		}()
	}

	//Wait until all 5 threads have finished
	wg.Wait()
	fmt.Println("runningTotal =", runningTotal)
	correctTotal := safeTotal.Load()
	fmt.Println("CorrectTotal =", correctTotal)
}
