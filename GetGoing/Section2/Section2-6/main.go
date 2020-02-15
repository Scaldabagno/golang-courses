package main

import "fmt"

// Car Interface with methods Drive and Stop
type Car interface {
	Drive()
	Stop()
}

// Lambo struct
type Lambo struct {
	LamboModel string
}

// Chevy struct
type Chevy struct {
	ChevyModel string
}

// Forces to declare all the methods of the interface to a specific structure
func newModel(arg string) Car {
	return &Lambo{arg}
}

// Drive method for Lambo
func (l *Lambo) Drive() {
	fmt.Println("Lambo going...")
	fmt.Println(l.LamboModel)
}

// Drive method for Chevy
func (c *Chevy) Drive() {
	fmt.Println("Chevy going...")
	fmt.Println(c.ChevyModel)
}

// Stop method for Lambo
func (l *Lambo) Stop() {
	fmt.Println("Lambo stops...")
	fmt.Println(l.LamboModel)
}

// Stop method for Chevy
func (c *Chevy) Stop() {
	fmt.Println("Chevy stops...")
	fmt.Println(c.ChevyModel)
}

func main() {
	l := Lambo{"Gallardo"}
	c := Chevy{"C369"}
	l.Drive()
	c.Drive()
	l.Stop()
	c.Stop()
}
