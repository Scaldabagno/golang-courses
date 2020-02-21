package main

import (
	"fmt"
	"time"
)

func heavy() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Heavy")
	}
	// time.Sleep(time.Second * 5)
}

func superHeavy() {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("Super Heavy")
	}
}

func main() {
	// heavy() // normal call
	go heavy()      // goroutine -> way to call a function, background
	go superHeavy() // as a goroutine, it does not block main execution
	fmt.Println("End")
	// time.Sleep(time.Second * 5) // time for running goroutine (bad way)
	select {} // right way
}

//  goroutine are executed concurrently
