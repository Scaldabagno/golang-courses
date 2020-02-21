package main

import "fmt"

// Car type for channel
type Car struct {
	Model string
}

func main() {
	// c := make(chan int, 3)  // size of the buffer
	c := make(chan *Car, 3) // referenced to Car type

	// sending
	// go func() {
	// 	// after sniffing 1, go sends 4 in channel
	// 	c <- 1
	// 	c <- 2
	// 	c <- 3
	// 	c <- 4
	// 	close(c) // closes channel, it prevents a deadlock
	// }()

	// for i := range c {
	// 	fmt.Println(i)
	// }

	// sending
	go func() {
		// after sniffing 1, go sends 4 in channel
		c <- &Car{"1"}
		c <- &Car{"2"}
		c <- &Car{"3"}
		c <- &Car{"4"}
		close(c) // closes channel, it prevents a deadlock
	}()

	for i := range c {
		fmt.Println(i.Model)
	}
}
