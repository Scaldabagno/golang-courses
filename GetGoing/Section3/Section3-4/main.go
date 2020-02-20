package main

import (
	"fmt"
	"time"
)

func heavy() {
	time.Sleep(time.Second * 5)
}

func main() {
	// heavy() // normal call
	go heavy() // goroutine -> way to call a function
	fmt.Println("End")
}
