package main

import "fmt"

type Car struct {
	Name     string
	Age      int
	ModelnNo int
}

func (c Car) Print() {
	fmt.Println(c)
}

func (c Car) Drive() {
	fmt.Println("driving")
}

func (c Car) GetName() string {
	return c.Name
}

func main() {
	// c := Car{} // {" " 0 0} -> nulltype
	// c := Car{"ah", 1, 2} // {"ah" 1 2}
	c := Car{
		Name:     "ah",
		Age:      1,
		ModelnNo: 2,
	} // same
	// var c1 Car
	// fmt.Println(c)
	c.Print()
	c.Drive()
	fmt.Println(c.GetName())
}
