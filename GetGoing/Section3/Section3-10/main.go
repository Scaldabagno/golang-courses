package main

import (
	"fmt"
	"os"
	"time"
)

// Select simulates select block behaviour
func Select(c chan int, quits chan struct{}) {
	// switch case
	//  select for channel sync
	for {
		time.Sleep(time.Second)
		select {
		case <-c:
			fmt.Println("Received")
		case <-quits:
			fmt.Println("Quit")
			os.Exit(0)
		}
	}
}

// All is executed one time -> Received - Quit <- close
func main() {
	c := make(chan int)
	quits := make(chan struct{})
	go func() {
		c <- 1
		quits <- struct{}{}
	}()

	go Select(c, quits)
	select {}
}
