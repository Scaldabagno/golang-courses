package main

import "fmt"

func main() {
	// <name> chan <type>
	// var c chan int // nil
	c := make(chan int)

	// send in a goroutine
	go func() {
		c <- 1 // sending something to a channel
	}()

	// sniff
	val := <-c // taking value from a channel
	fmt.Println(val)
	// send in a goroutine again
	go func() {
		c <- 2 // sending something to a channel
	}()

	// sniff again
	val = <-c // taking value from a channel
	fmt.Println(val)

	fmt.Println(c)
}

// NOTE: can't send something until the previous one was sniffed
