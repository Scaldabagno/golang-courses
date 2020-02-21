package main

import (
	"fmt"
	"sync"
)

func main() {
	// wait groups
	wg := &sync.WaitGroup{}
	// mutexes
	// mut := &sync.Mutex{}
	// mut.Lock()
	// mut.Unlock()
	// add, done, wait
	wg.Add(2) // adds 2 funcs, after 2 done calls, it will go on
	// lambda/anonymous functions as goroutines
	go func() {
		fmt.Println("Hello")
		wg.Done() // close
	}()
	go func() {
		fmt.Println("World")
		wg.Done() // close
	}()

	wg.Wait() // unlock

	fmt.Println("End")
	// select {} // deadlock

}
